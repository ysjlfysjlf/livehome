package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

// LiveVideo 表示一个直播视频
type LiveVideo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Filename    string    `json:"filename"`
	Duration    int       `json:"duration"` // 视频时长（秒）
	Size        int64     `json:"size"`     // 文件大小（字节）
	CreatedTime time.Time `json:"created_time"`
	URL         string    `json:"url,omitempty"` // 只在API响应中填充
}

// LiveReplayHandler 处理直播回放的相关请求
type LiveReplayHandler struct {
	videoDir string
	baseURL  string
}

// NewLiveReplayHandler 创建一个新的回放处理器
func NewLiveReplayHandler(videoDir, baseURL string) *LiveReplayHandler {
	return &LiveReplayHandler{
		videoDir: videoDir,
		baseURL:  baseURL,
	}
}

// ListReplays 返回所有可用的回放视频
func (h *LiveReplayHandler) ListReplays(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许CORS请求

	// 获取所有可用的回放视频
	videos, err := h.getAvailableVideos()
	if err != nil {
		log.Printf("获取视频列表失败: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "获取视频列表失败"})
		return
	}

	// 添加访问URL
	for i := range videos {
		videos[i].URL = fmt.Sprintf("%s/replay/%s", h.baseURL, videos[i].ID)
	}

	// 返回JSON响应
	json.NewEncoder(w).Encode(videos)
}

// GetReplay 获取特定回放视频信息
func (h *LiveReplayHandler) GetReplay(w http.ResponseWriter, r *http.Request) {
	// 从URL中提取视频ID
	videoID := strings.TrimPrefix(r.URL.Path, "/replay/")
	if strings.Contains(videoID, "/") {
		// 如果是请求视频流，处理视频流请求
		parts := strings.SplitN(videoID, "/", 2)
		if len(parts) == 2 && parts[1] == "stream" {
			h.StreamReplay(w, r, parts[0])
			return
		}
	}

	if videoID == "" {
		http.Error(w, "视频ID不能为空", http.StatusBadRequest)
		return
	}

	// 获取视频信息
	videos, err := h.getAvailableVideos()
	if err != nil {
		http.Error(w, "获取视频信息失败", http.StatusInternalServerError)
		return
	}

	// 查找指定ID的视频
	var targetVideo *LiveVideo
	for _, video := range videos {
		if video.ID == videoID {
			targetVideo = &video
			break
		}
	}

	if targetVideo == nil {
		http.Error(w, "视频不存在", http.StatusNotFound)
		return
	}

	// 添加视频URL
	targetVideo.URL = fmt.Sprintf("%s/replay/%s/stream", h.baseURL, targetVideo.ID)

	// 返回视频信息
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(targetVideo)
}

// StreamReplay 直接提供视频文件流
func (h *LiveReplayHandler) StreamReplay(w http.ResponseWriter, r *http.Request, videoID string) {
	// 设置CORS头
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Range, Accept")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Range, Content-Type")

	// 处理OPTIONS请求
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	log.Printf("接收到视频流请求: videoID=%s", videoID)

	// 获取视频信息
	videos, err := h.getAvailableVideos()
	if err != nil {
		log.Printf("获取视频列表失败: %v", err)
		http.Error(w, "获取视频信息失败", http.StatusInternalServerError)
		return
	}

	// 查找指定ID的视频
	var filename string
	for _, video := range videos {
		if video.ID == videoID {
			filename = video.Filename
			break
		}
	}

	if filename == "" {
		log.Printf("未找到视频: videoID=%s", videoID)
		http.Error(w, "视频不存在", http.StatusNotFound)
		return
	}

	// 安全检查：防止目录遍历攻击
	videoPath := filepath.Join(h.videoDir, filename)
	log.Printf("视频文件路径: %s", videoPath)

	// 检查文件是否存在
	fileInfo, err := os.Stat(videoPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("视频文件不存在: %s", videoPath)
			http.Error(w, "视频文件不存在", http.StatusNotFound)
		} else {
			log.Printf("无法访问视频文件: %v", err)
			http.Error(w, "无法访问视频文件", http.StatusInternalServerError)
		}
		return
	}

	// 检查是否是文件
	if fileInfo.IsDir() {
		log.Printf("路径指向的是目录而不是文件: %s", videoPath)
		http.Error(w, "不是有效的视频文件", http.StatusBadRequest)
		return
	}

	// 设置适当的Content-Type
	contentType := "video/mp4"
	if strings.HasSuffix(filename, ".webm") {
		contentType = "video/webm"
	} else if strings.HasSuffix(filename, ".ogg") {
		contentType = "video/ogg"
	}
	w.Header().Set("Content-Type", contentType)

	// 支持范围请求（用于视频流）
	http.ServeFile(w, r, videoPath)
	log.Printf("视频流请求完成: %s", videoPath)
}

// getAvailableVideos 获取所有可用的视频列表
func (h *LiveReplayHandler) getAvailableVideos() ([]LiveVideo, error) {
	files, err := os.ReadDir(h.videoDir)
	if err != nil {
		return nil, err
	}

	var videos []LiveVideo
	for _, file := range files {
		if !file.IsDir() {
			name := file.Name()
			if isVideoFile(name) {
				info, err := file.Info()
				if err != nil {
					continue
				}

				// 从文件名中提取ID（去掉.mp4后缀）
				id := strings.TrimSuffix(name, filepath.Ext(name))

				video := LiveVideo{
					ID:          id,                      // 直接使用文件名中的数字作为ID
					Title:       fmt.Sprintf("视频%s", id), // 标题使用"视频+数字"的格式
					Filename:    name,
					Size:        info.Size(),
					CreatedTime: info.ModTime(),
					// 简单假设每MB的视频有10秒钟
					Duration: int(info.Size() / (1024 * 1024) * 10),
				}
				videos = append(videos, video)
			}
		}
	}

	return videos, nil
}

// 替换原有的 broadcast 通道类型
type BroadcastMsg struct {
	Sender *websocket.Conn
	Data   []byte
}

// 替换原有的全局变量
var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan BroadcastMsg)
	clientsMu sync.Mutex
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // 允许跨域
	}
)

// 聊天 WebSocket 处理
func chatWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket 升级失败:", err)
		return
	}
	defer ws.Close()

	clientsMu.Lock()
	clients[ws] = true
	clientsMu.Unlock()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			clientsMu.Lock()
			delete(clients, ws)
			clientsMu.Unlock()
			break
		}
		// 广播时带上发送者
		broadcast <- BroadcastMsg{Sender: ws, Data: msg}
	}
}

// 广播消息给除发送者外的所有客户端
func handleChatMessages() {
	for {
		bmsg := <-broadcast
		clientsMu.Lock()
		for client := range clients {
			if client == bmsg.Sender {
				continue // 跳过自己
			}
			err := client.WriteMessage(websocket.TextMessage, bmsg.Data)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		clientsMu.Unlock()
	}
}

// 配置视频目录
const videoDir = "./videos"

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	// 确保视频目录存在
	if err := os.MkdirAll(videoDir, 0755); err != nil {
		log.Fatalf("创建视频目录失败: %v", err)
	}

	// 创建直播回放处理器
	baseURL := "http://47.105.118.106:8080"
	replayHandler := NewLiveReplayHandler(videoDir, baseURL)

	// 设置路由
	http.HandleFunc("/api/videos", listVideosHandler)
	http.HandleFunc("/video/", videoHandler)

	// 设置直播回放相关路由
	http.HandleFunc("/replay/", replayHandler.GetReplay)
	http.HandleFunc("/api/replays", replayHandler.ListReplays)

	// 聊天 WebSocket 路由
	http.HandleFunc("/api/chat", chatWebSocketHandler)
	go handleChatMessages()

	// 启动服务器
	port := "8080"
	fmt.Printf("服务器启动在 http://47.105.118.106:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

// videoHandler 提供视频文件流
func videoHandler(w http.ResponseWriter, r *http.Request) {
	// 从URL中提取视频文件名
	filename := strings.TrimPrefix(r.URL.Path, "/video/")
	if filename == "" {
		http.Error(w, "视频文件名不能为空", http.StatusBadRequest)
		return
	}

	// 安全检查：防止目录遍历攻击
	cleanFilename := filepath.Base(filename)
	videoPath := filepath.Join(videoDir, cleanFilename)

	// 检查文件是否存在
	fileInfo, err := os.Stat(videoPath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "视频文件不存在", http.StatusNotFound)
		} else {
			http.Error(w, "无法访问视频文件", http.StatusInternalServerError)
		}
		return
	}

	// 检查是否是文件
	if fileInfo.IsDir() {
		http.Error(w, "不是有效的视频文件", http.StatusBadRequest)
		return
	}

	// 设置适当的Content-Type
	contentType := "video/mp4"
	if strings.HasSuffix(filename, ".webm") {
		contentType = "video/webm"
	} else if strings.HasSuffix(filename, ".ogg") {
		contentType = "video/ogg"
	}
	w.Header().Set("Content-Type", contentType)

	// 支持范围请求（用于视频流）
	http.ServeFile(w, r, videoPath)
}

// listVideosHandler 返回可用视频列表
func listVideosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许CORS请求

	files, err := os.ReadDir(videoDir)
	if err != nil {
		log.Printf("读取视频目录失败: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "读取视频目录失败"}`))
		return
	}

	// 只列出视频文件
	videos := []string{}
	for _, file := range files {
		if !file.IsDir() {
			name := file.Name()
			if isVideoFile(name) {
				videos = append(videos, name)
			}
		}
	}

	// 构建JSON响应
	result := "["
	for i, video := range videos {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf(`"%s"`, video)
	}
	result += "]"

	w.Write([]byte(result))
}

// 判断文件是否是视频文件
func isVideoFile(filename string) bool {
	validExtensions := []string{".mp4", ".webm", ".ogg", ".mov", ".avi", ".mkv"}
	ext := strings.ToLower(filepath.Ext(filename))

	for _, validExt := range validExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}
