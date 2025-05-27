<template>
  <div class="app-container" :style="backgroundStyle">
    <div class="content-container">
      <TabNav />
      
      <div class="main-content">
        <div class="content-top">
          <div class="featured-live">
            <div class="featured-live-inner">
              <div class="thumbnail">
                <img src="../assets/play-icon.webp" alt="播放图标" class="preview-img" loading="lazy">
              </div>
              <div class="live-info">
                <h2 class="title">录播演示</h2>
                <p class="time">直播时间: 2025.05.17 22:05:00</p>
                <div class="action-buttons">
                  <button class="action-btn enter" @click="handleReplay">进入回放</button>
                  <button class="action-btn end">直播已结束</button>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <div class="content-bottom">
          <div class="content-left">
            <div class="detail-tabs">
              <div class="tab-header">
                <div class="tab active">详情</div>
                <div class="tab">直播互动</div>
              </div>
              <div class="tab-content">
                <!-- 详情内容 -->
                <div class="detail-content">
                  <img src="../assets/img_empty.webp" alt="详情图标" loading="lazy">
                </div>
              </div>
            </div>
          </div>

          <div class="category-replay-section">
              
              <div class="person-item">
                <img src="../assets/chicken.webp" alt="人物头像" loading="lazy">
                <div class="person-detail">小黑子2号</div>
                <button class="follow-btn">进店逛逛</button>
              </div>


              <div class="replay-list">
                <div class="list-title">相关推荐</div>
                <div class="replay-item">
                   <div class="replay-thumbnail">
                    <img src="../assets/cat.webp" alt="播放图标" loading="lazy">
                  </div>

                  <div class="replay-info">
                    <h4>录播演示2</h4>
                    <p>直播时间: 2025.05.17 22:05:00</p>
                  </div>
                  </div>

                <div class="replay-item" v-for="n in 8" :key="n">
                  <div class="replay-thumbnail">
                    <img src="../assets/play-icon.webp" alt="播放图标" loading="lazy">
                  </div>
                  <div class="replay-info">
                    <h4>录播演示{{ n + 2 }}</h4>
                    <p>直播时间: 2025.05.17 22:05:00</p>
                  </div>
                </div>
              </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import TabNav from '../components/TabNav.vue';

const router = useRouter();

const handleReplay = () => {
  router.push('/video');
};

</script>

<style scoped>
.app-container {
  min-height: 100vh;
  width: 100%;
}

.content-container {
  width: 100%;
  border-radius: 8px;
  overflow: hidden;
}

.main-content {
  display: flex;
  flex-direction: column;
  width: 75%;
  margin-left: auto;
  margin-right: auto;
}

.content-top {
  flex: 1;
  min-width: 100%;
}

.content-left {
  width: 75%;
  padding: 20px;
  background-color: white;
  margin-right: 3px;
  height: 1300px;
  margin-top: 10px;
}

/* 主要直播展示区 */
.featured-live {
  height: 300px;
  padding: 20px;
  background-color: white;
}

.featured-live-inner {
  display: flex;
  align-items: flex-start;
  gap: 20px;
}

.thumbnail {
  flex-shrink: 0;
  width: 400px;
  height: 230px;
  background-color: #f0f0f0;
  border-radius: 4px;
  overflow: hidden;
  position: relative;
}

.preview-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.live-info {
  flex: 1;
}

.title {
  font-size: 20px;
  margin-bottom: 8px;
}

.time {
  color: #666;
  font-size: 14px;
  margin-bottom: 15px;
}

.action-buttons {
  display: flex;
  gap: 10px;
  margin-top: 90px;
}

.action-btn {
  padding: 8px 16px;
  border-radius: 4px;
  font-size: 14px;
  cursor: pointer;
  margin-top: 45px;
}

.enter {
  background-color: #15bf84ff;
  color: white;
  border: none;
}

.end {
  background-color: white;
  color: #666;
  border: none;
}

.content-bottom {
  margin-top: 5px;
  display: flex;
  flex-direction: row;
  width: 100%;
}

/* 分类和回放列表 */
.category-replay-section {
  width: 25%;
  background-color: #f0f0f0;
  margin-left: 10px;
}

.category-title h3 {
  margin-bottom: 10px;
  font-size: 16px;
}

.category-filter {
  margin-bottom: 15px;
}

.category-filter p {
  margin-bottom: 8px;
  font-size: 14px;
}

.category-items {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.category-item {
  border-radius: 15px;
  background-color: #f0f0f0;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s;
}

.category-item:hover {
  background-color: #e0e0e0;
}

.category-item.active {
  background-color: #2196f3;
  color: white;
}

.person-item{
  margin-top: 10px;
  /* display: flex; */
  width: 100%;
  height: 200px;
  background-color: white;
}

.person-item img{
  margin-top: 20px;
  width: 85px;
  height: 85px;
  margin-left: 130px;
}

.person-detail{
  margin-top: 5px;
  margin-left: 120px;
  font-size: 20px;
}

.follow-btn{
  margin-top: 10px;
  margin-left: 100px;
  background-color: #f0f0f0;
  height: 30px;
  width: 150px;
}


/* 回放列表 */
.replay-list {
  margin-top: 15px;
  background-color: white;
  height: 100%;
  position: relative;
  padding-top: 20px;
}

.list-title{
  font-size: 20px;
  margin-left: 20px;
}

.replay-item {
  display: flex;
  padding: 10px 0;
  border-bottom: 1px solid #f0f0f0;
}

.replay-thumbnail {
  width: 100px;
  height: 80px;
  margin-right: 20px;
}

.replay-thumbnail img {
  width: 100px;
  height: 80px;
  margin-left: 10px;
}

.replay-info h4 {
  font-size: 14px;
  margin-bottom: 40px;
}

.replay-info p {
  font-size: 12px;
  color: #666;
}

/* 详情标签 */
.detail-tabs {
  background-color: white;
  border-radius: 4px;
  overflow: hidden;
}

.tab-header {
  display: flex;
  border-bottom: 1px solid #eee;
}

.tab {
  padding: 10px 20px;
  cursor: pointer;
  position: relative;
  font-size: 14px;
}

.tab.active {
  color: #15bf84ff;
}

.tab.active::after {
  content: '';
  position: absolute;
  bottom: -1px;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: #15bf84ff;
}

.tab-content {
  padding: 15px;
  min-height: 200px;
}

.detail-content{
  display: flex;
  justify-content: center;
  align-items: center;
}
.detail-content img{
  margin-top: 100px;
  height: 200px;
  width: 200px;
}

@media (max-width: 768px) {
  .main-content {
    flex-direction: column;
  }
  
  .content-left {
    border-right: none;
  }
  
  .content-right {
    width: 100%;
  }
  
  .featured-live-inner {
    flex-direction: column;
  }
  
  .thumbnail {
    width: 100%;
  }
}
</style> 