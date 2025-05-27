# 视频回放服务 API 文档

本文档描述了视频回放服务提供的API接口，供前端开发人员集成。

## 基础URL

所有API的基础URL为：`http://localhost:8080`

## API 端点

### 1. 获取所有可用回放列表

**请求**:
```
GET /api/replays
```

**响应**:
```json
[
  {
    "id": "abc123",
    "title": "示例直播回放",
    "filename": "example.mp4",
    "duration": 1800,
    "size": 104857600,
    "created_time": "2023-09-25T14:30:00Z",
    "url": "http://localhost:8080/replay/abc123"
  },
  ...
]
```

### 2. 获取特定回放视频信息

**请求**:
```
GET /replay/{videoId}
```

**参数**:
- `videoId`: 视频的唯一标识符

**响应**:
```json
{
  "id": "abc123",
  "title": "示例直播回放",
  "filename": "example.mp4",
  "duration": 1800,
  "size": 104857600,
  "created_time": "2023-09-25T14:30:00Z",
  "url": "http://localhost:8080/replay/abc123/stream"
}
```

### 3. 获取视频流

**请求**:
```
GET /replay/{videoId}/stream
```

**参数**:
- `videoId`: 视频的唯一标识符

**响应**:
视频文件的二进制流，可以直接用于`<video>`标签的`src`属性或播放器的源URL。

### 4. 获取所有视频列表（简单版本）

**请求**:
```
GET /api/videos
```

**响应**:
```json
[
  "video1.mp4",
  "video2.mp4",
  ...
]
```

### 5. 获取特定视频文件

**请求**:
```
GET /video/{filename}
```

**参数**:
- `filename`: 视频文件名（包括扩展名）

**响应**:
视频文件的二进制流。

## 使用示例

### 获取视频列表并播放

```javascript
// 获取所有可用回放
fetch('http://localhost:8080/api/replays')
  .then(response => response.json())
  .then(videos => {
    if (videos.length > 0) {
      // 获取第一个视频的信息
      const video = videos[0];
      
      // 视频流URL
      const videoUrl = `${video.url}/stream`;
      
      // 初始化播放器
      const player = videojs('videoPlayer', {
        controls: true,
        autoplay: false,
        preload: 'auto',
        sources: [{
          src: videoUrl,
          type: 'video/mp4'
        }]
      });
      
      // 更新视频信息显示
      document.getElementById('videoTitle').textContent = video.title;
      document.getElementById('videoDuration').textContent = 
        `${Math.floor(video.duration / 60)}分${video.duration % 60}秒`;
    }
  });
```

### 使用专用播放器页面

您也可以直接使用我们提供的播放器页面：

```
http://localhost:8080/player/replay/{videoId}
```

这将打开一个专用的视频播放页面，无需自行开发播放器界面。

## 错误处理

所有API在发生错误时返回适当的HTTP状态码和错误消息：

- `400 Bad Request`: 请求参数无效
- `404 Not Found`: 请求的资源不存在
- `500 Internal Server Error`: 服务器内部错误

错误响应格式：

```json
{
  "error": "错误描述信息"
}
``` 