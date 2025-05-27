# 视频直播回放服务

这是一个用Go实现的简单视频直播回放服务，提供视频文件的流式传输和播放功能。

## 功能特点

- 支持视频文件的流式传输
- 自动识别常见视频格式
- 提供视频列表API
- 支持视频范围请求（用于拖动进度条）
- 简单的Web界面展示可用视频

## 使用方法

### 准备工作

1. 确保已安装Go环境（建议Go 1.18或更高版本）
2. 克隆本仓库到本地

### 放置视频文件

1. 在项目根目录下创建一个名为`videos`的文件夹
2. 将您的视频文件（mp4、webm、ogg等格式）放入该文件夹

### 运行服务

```bash
go run main.go
```

服务器将在 http://localhost:8080 启动，您可以通过浏览器访问此地址来查看和播放视频。

### API接口

- `GET /` - 网站首页，显示可用视频列表
- `GET /video/{filename}` - 获取指定视频文件的流
- `GET /api/videos` - 获取所有可用视频的JSON列表

## 与前端集成

您可以使用以下方式在前端获取视频URL：

```javascript
// 获取视频列表
fetch('http://localhost:8080/api/videos')
  .then(response => response.json())
  .then(videos => {
    // 使用第一个视频或特定视频
    const videoUrl = `http://localhost:8080/video/${videos[0]}`;
    
    // 初始化视频播放器
    player.value = videojs(videoPlayer.value, {
      controls: true,
      autoplay: false,
      preload: 'auto',
      fluid: true,
      responsive: true,
      sources: [{
        src: videoUrl,
        type: 'video/mp4'
      }]
    });
  });
```

## 业务流程

1. 进入直播间
2. 请求接口数据
   - 成功：进入视频播放器
   - 失败：显示异常提示 