<h3><b>Description</b></h3>
小鹅通公司和学院联合举办的为期两周的项目实战，是一款基于Go+WebSocket+Vue3+Vite+Node+Nginx+Dokcer的网站Go+WebSocket+Vue3+Vite+Node+Nginx+Dokcer的网站
<h3><b>效果展示：</b></h3>


https://github.com/user-attachments/assets/9dd365bb-556e-4a7e-8039-d9cce1285f7c

首页效果图，点击进入回放按钮，可进入播放视频页面。
<img width="1525" height="907" alt="image" src="https://github.com/user-attachments/assets/6e34dbeb-366e-47b0-8ad0-ae9ed5445008" />

播放视频页面，支持视频播放的倍速、前进、后退、暂停的能力
<img width="1532" height="913" alt="image" src="https://github.com/user-attachments/assets/6408ac42-4753-44f3-8176-ee439b118b4d" />

聊天模块，使用WebSocket实现多人实时聊天功能

<img width="263" height="858" alt="image" src="https://github.com/user-attachments/assets/47673bb3-5e58-49ed-a266-7d8f2fb1adc1" />


</br></br>
<h2><b>以下为项目提交说明（无关）</b></h2>
## 技术栈

- 前端：Vite + Vue3
- 后端：Go
- 部署：Docker & Docker Compose（阿里云服务器ubuntu）


### 1. 构建与运行（使用 Docker Compose）

服务器安装 Docker 和 Docker Compose。

```bash
docker-compose up --build -d
```

- 前端服务将运行在 `80` 端口
- 后端服务将运行在 `8080` 端口

  ### 2. 访问方式

- 前端页面：http://47.105.118.106
- 后端 API：http://47.105.118.106:8080/replay/1/stream


## 已实现功能

1. **直播间播放器能力**
   - 支持视频播放的倍速切换、前进、后退、暂停等操作。
   - 前端自定义了视频控制条，用户可自由调整播放进度和速度。
  
2. **前后端服务部署与优化**
   - 前后端均已通过 Docker 容器化，键部署到自有服务器（阿里云ubuntu22.04LTS）。
   - 通过将图片压缩并转换为webp格式，优化页面打开速度。例如页面中的小猫的图片，下载下来时是114kb，但是经过压缩并转换为webp格式，最后变为了48kb， 
     其他图片压缩的更小。
   - 前端同时使用图片懒加载优化页面打开速度。


 3. **直播间评论聊天功能**
   - 已集成 WebSocket，实现直播间的实时评论与消息推送。
   - 前端评论区支持消息即时显示，后端 Go 服务负责消息广播。





