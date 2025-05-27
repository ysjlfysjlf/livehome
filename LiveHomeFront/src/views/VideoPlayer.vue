<template>
  <div class="video-container">

    <div class="main-content">
      <!-- 左侧视频区域 -->
      <div class="video-wrapper" :class="{ 'loading': loading }">
        <div v-if="loading" class="loading-indicator">
          加载中...
        </div>
        <div v-else-if="error" class="error-message">
          {{ error }}
          <div class="action-buttons">
            <button @click="retry" class="retry-button">重试</button>
            <button @click="goBack" class="back-button">返回</button>
          </div>
        </div>
        
        <div v-else class="player-area">
          <!-- 使用HTML5 video标签 -->
          <video
            ref="videoPlayer"
            class="video-element"
            @timeupdate="onTimeUpdate"
            @loadedmetadata="onVideoLoaded"
            @play="onPlay"
            @pause="onPause"
            @ended="onEnded"
            @error="onError"
            @waiting="onWaiting"
            @canplay="onCanPlay"
            preload="auto"
            controlslist="nodownload nofullscreen noremoteplayback"
            nocontrols
          >
            <source :src="videoUrl" type="video/mp4">
            <source :src="videoUrl" type="video/webm">
            <source :src="videoUrl" type="application/x-mpegURL">
            您的浏览器不支持HTML5视频播放，请更换浏览器尝试。
          </video>
          
          <!-- 自定义视频控制条 -->
          <div class="custom-controls">
            <div class="progress-container">
              <div class="progress-bar" ref="progressBar" @click="seek">
                <div class="progress-filled" :style="{ width: progress + '%' }"></div>
              </div>
            </div>
            
            <div class="controls-bottom">
              <div class="left-controls">
                <button class="control-btn play-btn" @click="togglePlay">
                  <span class="control-icon">{{ isPlaying ? '⏸' : '▶' }}</span>
                </button>
                <button class="control-btn restart-btn" @click="restart">
                  <span class="control-icon">↻</span>
                </button>
                <div class="time-display">{{ currentTimeFormatted }} / {{ durationFormatted }}</div>
              </div>
              
              <div class="right-controls">
                <div class="volume-control">
                  <button class="control-btn volume-btn" @click="toggleMute">
                    <span class="control-icon">{{ isMuted ? '🔇' : '🔊' }}</span>
                  </button>
                  <input 
                    type="range" 
                    class="volume-slider" 
                    min="0" 
                    max="1" 
                    step="0.1" 
                    v-model="volume"
                    @input="updateVolume"
                  >
                </div>
                <select class="playback-rate" v-model="playbackRate" @change="updatePlaybackRate">
                  <option value="0.5">0.5x</option>
                  <option value="1.0">1.0x</option>
                  <option value="1.5">1.5x</option>
                  <option value="2.0">2.0x</option>
                </select>
                <button class="control-btn fullscreen-btn" @click="toggleFullscreen">
                  <span class="control-icon">⛶</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 右侧讨论区域 -->

  
      <div class="top-tabs">
        <div 
          v-for="tab in tabs" 
          :key="tab"
          :class="['tab', currentTab === tab ? 'active' : '']"
          @click="switchTab(tab)"
        >
          {{ tab }}
        </div>
      </div>

      <div class="sidebar">
        <!-- 讨论区域 -->
        <div v-if="currentTab === '讨论'" class="chat-area">

          <div class="assistant-name">
            <img src="../assets/assistant-name.webp" alt="人物头像" loading="lazy">直播助手</div>
          <div class="assistant-tip">
            <div class="upload-text">欢迎进入直播间</div>
            <div class="upload-text">1、请自行调节手机音量至合适的状态</div>
          <div class="upload-text">2、直播界面显示讲师发布的内容，听众发言可以在讨论区进行或以弹幕形式查看</div>
          <div class="upload-text">3、直播结束后，你可以随时回看全部内容</div>
          </div>


          
          <div class="chat-messages" ref="chatMessages">
            <div v-for="(msg, index) in messages" :key="index" class="message">
              <span class="username">{{ msg.username }}:</span>
              <span class="message-content">{{ msg.content }}</span>
            </div>
          </div>

          <div class="comment-input">
            <div class="input-container">
              <input 
                type="text" 
                v-model="currentMessage" 
                @keyup.enter="sendMessage"
                placeholder="请输入您的评论内容" 
                class="comment-box" 
              />
              <button class="send-btn" @click="sendMessage">发送</button>
            </div>
          </div>
        </div>

        <!-- 讲解区域 -->
        <div v-else-if="currentTab === '讲解'" class="lecture-area">
          <!-- 暂时空白 -->
        </div>

        <!-- 文件区域 -->
        <div v-else-if="currentTab === '文件'" class="file-area">
          <div class="file-img">
            <img src="../assets/courseFileEmpty.webp" alt="文件图标" class="file-empty-icon" loading="lazy">
            <div class="file-empty-text">暂无共享文件夹</div>
          </div>
          <!-- 暂时空白 -->
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const videoPlayer = ref(null);
const progressBar = ref(null);
const loading = ref(true);
const error = ref(null);

// 视频状态
const isPlaying = ref(false);
const isMuted = ref(false);
const progress = ref(0);
const currentTime = ref(0);
const duration = ref(0);
const volume = ref(1);
const playbackRate = ref('1.0');

// 视频源URL - 添加调试日志和备用视频源
const apiBaseUrl = '/api';
const videoId = ref('1'); // 默认使用ID为1的视频

// 初始化视频URL
const videoUrl = ref(`${apiBaseUrl}/replay/${videoId.value}/stream`);
console.log('尝试加载视频:', videoUrl.value);

// 切换到备用视频源
const switchToBackupSource = () => {
  console.log('切换到备用视频源:', backupVideoUrl.value);
  videoUrl.value = backupVideoUrl.value;
  if (videoPlayer.value) {
    videoPlayer.value.load();
  }
};

// 检查视频源是否可访问
const checkVideoSource = async () => {
  try {
    console.log('正在检查视频源是否可访问...');
    const response = await fetch(`${apiBaseUrl}/replay/${videoId.value}`);
    if (!response.ok) {
      throw new Error(`视频源检查失败: ${response.status}`);
    }
    const videoData = await response.json();
    console.log('获取到视频数据:', videoData);
    
    // 确保URL是完整的
    if (!videoData.url.startsWith('http')) {
      videoUrl.value = `${apiBaseUrl}${videoData.url}`;
    } else {
      videoUrl.value = videoData.url;
    }
    
    console.log('设置视频URL:', videoUrl.value);
    loading.value = false;
  } catch (e) {
    console.error('视频源检查失败:', e);
    error.value = '无法连接到视频服务器，请检查网络连接或服务器状态';
    loading.value = false;
  }
};

// 格式化时间
const formatTime = (seconds) => {
  const mins = Math.floor(seconds / 60);
  const secs = Math.floor(seconds % 60);
  return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`;
};

const currentTimeFormatted = ref('00:00');
const durationFormatted = ref('00:00');

// 视频控制方法
const togglePlay = () => {
  if (!videoPlayer.value) return;
  if (isPlaying.value) {
    videoPlayer.value.pause();
  } else {
    videoPlayer.value.play();
  }
};

const restart = () => {
  if (!videoPlayer.value) return;
  videoPlayer.value.currentTime = 0;
  videoPlayer.value.play();
};

const seek = (event) => {
  if (!videoPlayer.value || !progressBar.value) return;
  const rect = progressBar.value.getBoundingClientRect();
  const pos = (event.clientX - rect.left) / rect.width;
  videoPlayer.value.currentTime = pos * videoPlayer.value.duration;
};

const toggleMute = () => {
  if (!videoPlayer.value) return;
  videoPlayer.value.muted = !videoPlayer.value.muted;
  isMuted.value = videoPlayer.value.muted;
};

const updateVolume = () => {
  if (!videoPlayer.value) return;
  videoPlayer.value.volume = volume.value;
  isMuted.value = volume.value === 0;
};

const updatePlaybackRate = () => {
  if (!videoPlayer.value) return;
  videoPlayer.value.playbackRate = parseFloat(playbackRate.value);
};

const toggleFullscreen = () => {
  if (!videoPlayer.value) return;
  if (document.fullscreenElement) {
    document.exitFullscreen();
  } else {
    videoPlayer.value.requestFullscreen();
  }
};

// 视频事件处理
const onTimeUpdate = () => {
  if (!videoPlayer.value) return;
  currentTime.value = videoPlayer.value.currentTime;
  duration.value = videoPlayer.value.duration;
  progress.value = (currentTime.value / duration.value) * 100;
  currentTimeFormatted.value = formatTime(currentTime.value);
  durationFormatted.value = formatTime(duration.value);
};

const onWaiting = () => {
  console.log('视频正在缓冲...');
};

const onCanPlay = () => {
  console.log('视频可以开始播放');
  loading.value = false;
};

const onVideoLoaded = () => {
  console.log('视频元数据已加载，视频时长:', videoPlayer.value?.duration);
  if (!videoPlayer.value) return;
  duration.value = videoPlayer.value.duration;
  durationFormatted.value = formatTime(duration.value);
  loading.value = false;
};

const onPlay = () => {
  isPlaying.value = true;
};

const onPause = () => {
  isPlaying.value = false;
};

const onEnded = () => {
  isPlaying.value = false;
  progress.value = 100;
};



const goBack = () => {
  router.push('/');
};

// 聊天相关的状态
const messages = ref([]);
const currentMessage = ref('');
const username = ref('');
const chatMessages = ref(null);

// 标签相关
const tabs = ['讲解', '讨论', '文件']
const currentTab = ref('讨论')

const switchTab = (tab) => {
  currentTab.value = tab
}

// 生成随机用户名
const generateUsername = () => {
  const randomNum = Math.floor(Math.random() * 100);
  return `用户${randomNum}`;
};

// 发送消息
const sendMessage = async () => {
  if (!currentMessage.value.trim()) {
    return;
  }

  // 每次发送消息时生成新的随机用户名
  const currentUsername = generateUsername();

  const messageData = {
    username: currentUsername,
    content: currentMessage.value.trim(),
    timestamp: new Date().getTime()
  };

  // 先添加到本地消息列表
  messages.value.push(messageData);
  
  // 清空输入框
  currentMessage.value = '';

  // 滚动到最新消息
  nextTick(() => {
    if (chatMessages.value) {
      chatMessages.value.scrollTop = chatMessages.value.scrollHeight;
    }
  });


  // 通过 WebSocket 发送消息
  if (ws.value && ws.value.readyState === 1) {
    ws.value.send(JSON.stringify(messageData));
  }
  currentMessage.value = '';
};

const ws = ref(null)
// 在组件挂载时生成用户名
onMounted(() => {
  username.value = generateUsername();
  console.log('组件已挂载，检查视频元素:', videoPlayer.value);
  
  // 先检查视频源
  checkVideoSource();
  
  if (videoPlayer.value) {
    videoPlayer.value.volume = volume.value;
    
    // 添加视频网络状态监听
    videoPlayer.value.addEventListener('waiting', () => {
      console.log('视频等待数据中...');
    });
    
    videoPlayer.value.addEventListener('canplay', () => {
      console.log('视频可以播放');
      loading.value = false;
    });
    
    videoPlayer.value.addEventListener('stalled', () => {
      console.log('视频加载停滞');
    });
  } else {
    console.error('视频元素引用获取失败');
  }
  
  // 30秒后如果仍在加载，尝试自动重试
  const loadingTimeout = setTimeout(() => {
    if (loading.value) {
      console.log('视频加载超时，尝试重新加载');
      retry();
    }
    clearTimeout(loadingTimeout);
  }, 30000);


  ws.value = new WebSocket('ws://47.105.118.106/api/chat');
  ws.value.onmessage = (event) =>{
    const msg =JSON.parse(event.data);
    messages.value.push(msg);
    nextTick(() => {
      if(chatMessages.value){
        chatMessages.value.scrollTop = chatMessages.value.scrollHeight
      }
    });
  };
});

onBeforeUnmount(() => {
  if (videoPlayer.value) {
    videoPlayer.value.pause();
  }
});


</script>

<style scoped>
.video-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background-color: #000;
  color: #fff;
}



.top-tabs {
  display: flex;
  position: absolute;
  height: 60px;
  background-color: #2b2f38ff;
  right: 0;
  width: 17%;
}

.tab {
  padding: 0 15px;
  cursor: pointer;
  font-size: 16px;
  margin-right: 40px;
  margin-top: 25px;
}

.tab.active {
  position: relative;
}

.tab.active::after {
  content: '';
  position: absolute;
  bottom: -5px;
  left: 10px;
  width: 80%;
  height: 2px;
  background-color: #fff;
}

.main-content {
  display: flex;
  flex: 1;
  height: calc(100vh - 40px);
}

.video-wrapper {
  flex: 3;
  position: relative;
  background-color: #000;
}

.player-area {
  width: 100%;
  height: 100%;
  position: relative;
}

.video-element {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.custom-controls {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  background: linear-gradient(transparent, rgba(0, 0, 0, 0.7));
  padding: 10px;
  opacity: 1;
  transition: opacity 0.3s;
}

.player-area:hover .custom-controls {
  opacity: 1;
}

.progress-container {
  padding: 10px 0;
}

.progress-bar {
  height: 4px;
  background: rgba(255, 255, 255, 0.2);
  cursor: pointer;
  position: relative;
}

.progress-bar:hover {
  height: 6px;
}

.progress-filled {
  background: #3498db;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  transition: width 0.1s;
}

.controls-bottom {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.left-controls, .right-controls {
  display: flex;
  align-items: center;
  gap: 10px;
}

.control-btn {
  background: none;
  border: none;
  color: white;
  cursor: pointer;
  padding: 5px;
  font-size: 16px;
}

.time-display {
  font-size: 14px;
  color: black;
}

.volume-control {
  display: flex;
  align-items: center;
  gap: 5px;
}

.volume-slider {
  width: 60px;
  height: 4px;
}

.playback-rate {
  background: rgba(255, 255, 255, 0.1);
  color: white;
  border: none;
  padding: 3px 5px;
  border-radius: 3px;
  cursor: pointer;
  background-color: #000;
}

.sidebar {
  position: relative;
  display: flex;
  flex-direction: column;
  background-color: black;
  width: 17%;
  height: 90%;
  margin-top: 68px;
}

.file-upload-area {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 200px;
  flex-direction: column;
  padding: 20px;
  margin-top: 150px;
}

.upload-icon {
  font-size: 32px;
  margin-bottom: 10px;
  width: 60px;
  height: 60px;
  background-color: #333;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.upload-icon img{
  width: 120px;
  height: 120px;
}



.comment-input {
  position: fixed;
  bottom: 0;
  right: 0;
  width: 17%;
  padding: 10px;
  background-color: #181a1fff;
  border-top: 1px solid #333;
}

.input-container {
  display: flex;
  align-items: center;
  background: #333;
  border-radius: 4px;
  margin-bottom: 10px;
}

.comment-box {
  flex: 1;
  background: transparent;
  border: none;
  color: #fff;
  padding: 8px 10px;
  outline: none;
}

.send-btn {
  height: 30px;
  padding: 0 10px;
  cursor: pointer;
  background: #4CAF50;
  color: white;
  border: none;
  border-radius: 5px;
}


.send-btn:hover {
  background: #45a049;
}

.send-btn:active {
  background: #3d8b40;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  margin-top: 60px;
  margin-bottom: 120px; /* 增加底部间距，防止消息被输入框遮挡 */
}

.message {
  margin-bottom: 12px;
  line-height: 1.4;
}

.username {
  color: #fff;
  font-weight: bold;
  margin-right: 8px;
}

.message-content {
  color: #fff;
  word-break: break-all;
}

.lecture-area,
.file-area {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #666;
  font-size: 14px;
}

.file-empty-icon{
  width: 120px;
  height: 120px;
}

.file-empty-text{
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bolder;
}

/* 确保聊天区域占满整个侧边栏 */
.chat-area {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.assistant-tip{
  width: 80%;
  height: 170px;
  margin-left:30px;
  background-color: #3e79f0ff;
  margin-top: 10px;
  border-radius: 5px;
}

.assistant-name{
  font-size: 20px;
  font-weight: bolder;
  margin-left: 30px;
  margin-top: 10px;
}

.assistant-name img{
  width: 30px;
  height: 30px;
  margin-right: 10px;
}

.upload-text{
  margin-left: 7px;
}

@media (max-width: 768px) {
  .main-content {
    flex-direction: column;
  }
  
  .sidebar {
    flex: none;
    height: 300px;
    border-left: none;
    border-top: 1px solid #333;
  }
}
</style>
