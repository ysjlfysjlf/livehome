import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import VideoPlayer from '../views/VideoPlayer.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/video',
    name: 'VideoPlayer',
    component: VideoPlayer
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router; 