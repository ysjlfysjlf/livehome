<template>
  <div class="tab-nav">
    <div class="hand-icon">
      <img src="../assets/hand-icon.webp" alt="Icon" loading="lazy">
    </div>

    <div class="tabs">
      <div 
        v-for="(tab, index) in tabs" 
        :key="index" 
        :class="['tab-item', activeTab === index ? 'active' : '']"
        @click="selectTab(index)"
      >
        {{ tab }}
      </div>
    </div>

    <!-- 新增容器包裹右侧内容 -->
    <div class="right-content">
      <div class="search-box">
        <input type="text" placeholder="搜索">
      </div>
      <div class="vertical-line"></div>
      <div class="learn-buttons">
        <div class="learn-button">学习记录</div>
        <div class="vertical-line"></div>
        <div class="learn-button">我的订阅</div>
      </div>
      <div class="login-button">登录</div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const tabs = ['首页', '我的学习', '我的课程', '报考系统'];
const activeTab = ref(0);

const selectTab = (index) => {
  activeTab.value = index;
};
</script>

<style scoped>
.tab-nav {
  display: flex;
  align-items: center;
  padding: 15px 25px;
  background-color: white;
  margin-bottom: 3px;
  width: 100%;
}

.hand-icon {
  display: flex;
  margin-left: 250px;
}

.hand-icon img {
  width: 40px; 
  height: 40px;
}

.tabs {
  margin-left: 100px;
  display: flex;
  gap: 50px;
}

.tab-item {
  font-size: 20px;
  color: #333;
  padding: 8px 0;
  cursor: pointer;
  position: relative;
}

.tab-item.active {
  color: #15bf84ff;
  font-weight: bold;
}

.tab-item.active::after {
  content: "";
  position: absolute;
  bottom: -1px;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: #15bf84ff;
}

/* 新增 right-content 样式 */
.right-content {
  display: flex;
  align-items: center;
  margin-left: auto; /* 靠右对齐 */
}

.search-box {
  position: relative;
  margin-right: 20px;
}

.search-box input {
  padding: 8px 15px;
  border-radius: 20px;
  border: none;
  outline: none;
  width: 100px;
  padding-left: 35px;
  background-color: #f5f5f5;
}

.search-box::before {
  content: "";
  position: absolute;
  left: 10px;
  top: 50%;
  transform: translateY(-50%);
  width: 16px; /* 设置宽度 */
  height: 16px; /* 设置高度 */
  background-image: url('@/assets/search.png'); /* 使用 @ 别名加载静态资源 */
  background-size: contain;
  background-repeat: no-repeat;
  color: #999;
}

.learn-buttons {
  display: flex;
  gap: 10px;
  font-size: 20px;
}

.learn-button {
  padding: 8px 15px;
  border-radius: 20px;
  outline: none;
  cursor: pointer;
  transition: all 0.3s;
}

.vertical-line{
  width: 1px;
  height: 24px;
  background-color: #999;
  margin: 0 5px;
  align-self: center;
  opacity: 0.6;
}

.login-button {
  margin-right: 250px;
  padding: 8px 15px;
  border-radius: 5px;
  border: 1px solid #999;
  color: #15bf84ff;
  cursor: pointer;
  transition: all 0.3s;
  width: 100px;
  display: flex;
  justify-content: center;
}
</style>