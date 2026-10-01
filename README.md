## 📌 Description

小鹅通公司和学院联合举办的为期两周的项目实战，是一款基于 Go + WebSocket + Vue3 + Vite + Node + Nginx + Docker 的网站。

## 🖥 效果展示

首页效果图，点击进入回放按钮，可进入播放视频页面。

<img width="1525" height="907" alt="image" src="https://github.com/user-attachments/assets/6e34dbeb-366e-47b0-8ad0-ae9ed5445008" />

播放视频页面，支持视频播放的倍速、前进、后退、暂停的能力。

<img width="1532" height="913" alt="image" src="https://github.com/user-attachments/assets/6408ac42-4753-44f3-8176-ee439b118b4d" />

聊天模块，使用 WebSocket 实现多人实时聊天功能。

<img width="263" height="858" alt="image" src="https://github.com/user-attachments/assets/47673bb3-5e58-49ed-a266-7d8f2fb1adc1" />

### 🏆 项目成果

- 独立完成前后端全部开发
- 在 40+ 人参加的比赛中获得「最佳实现奖」

<img src="https://github.com/ysjlfysjlf/livehome/blob/main/xet.jpg?raw=true" width="250" />

## 📦 项目提交说明

### 技术栈

- 前端：Vite + Vue3
- 后端：Go
- 部署：Docker & Docker Compose（阿里云服务器 Ubuntu）

### 1. 构建与运行（使用 Docker Compose）

服务器安装 Docker 和 Docker Compose。

```bash
docker-compose up --build -d
