package main

// PlayerHandler 处理视频播放器页面
type PlayerHandler struct {
	baseURL string
}

// NewPlayerHandler 创建一个新的播放器处理器
func NewPlayerHandler(baseURL string) *PlayerHandler {
	return &PlayerHandler{
		baseURL: baseURL,
	}
}
