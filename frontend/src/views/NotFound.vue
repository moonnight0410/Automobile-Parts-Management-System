<template>
  <div class="not-found">
    <div class="not-found-content">
      <div class="error-icon">
        <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2"/>
          <path d="M12 8V12" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          <circle cx="12" cy="16" r="1" fill="currentColor"/>
        </svg>
      </div>
      <h1 class="error-code">404</h1>
      <h2 class="error-title">页面未找到</h2>
      <p class="error-desc">抱歉，您访问的页面不存在或已被移除</p>
      <div class="error-actions">
        <a-button type="primary" class="home-btn" @click="goHome">
          <template #icon><HomeOutlined /></template>
          返回首页
        </a-button>
        <a-button class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回上一页
        </a-button>
      </div>
      <div class="error-suggestions">
        <div class="suggestion-title">您可以尝试：</div>
        <div class="suggestion-list">
          <div class="suggestion-item" @click="goTo('/parts/list')">
            <AppstoreOutlined class="suggestion-icon" />
            <span>浏览零部件</span>
          </div>
          <div class="suggestion-item" @click="goTo('/bom/list')">
            <FileTextOutlined class="suggestion-icon" />
            <span>查看BOM</span>
          </div>
          <div class="suggestion-item" @click="goTo('/dashboard')">
            <DashboardOutlined class="suggestion-icon" />
            <span>返回仪表盘</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import {
  HomeOutlined,
  ArrowLeftOutlined,
  AppstoreOutlined,
  FileTextOutlined,
  DashboardOutlined
} from '@ant-design/icons-vue'

const router = useRouter()

function goHome() {
  router.push('/')
}

function goBack() {
  router.back()
}

function goTo(path: string) {
  router.push(path)
}
</script>

<style scoped>
.not-found {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  padding: 40px 20px;
}

.not-found-content {
  text-align: center;
  max-width: 500px;
}

.error-icon {
  width: 80px;
  height: 80px;
  margin: 0 auto 24px;
  color: #6366f1;
  animation: pulse 2s ease-in-out infinite;
}

.error-icon svg {
  width: 100%;
  height: 100%;
}

@keyframes pulse {
  0%, 100% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.05);
    opacity: 0.8;
  }
}

.error-code {
  font-size: 120px;
  font-weight: 800;
  margin: 0;
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1;
  letter-spacing: -4px;
}

.error-title {
  font-size: 28px;
  font-weight: 600;
  color: #1f2937;
  margin: 16px 0 8px;
}

.error-desc {
  font-size: 16px;
  color: #6b7280;
  margin: 0 0 32px;
}

.error-actions {
  display: flex;
  justify-content: center;
  gap: 16px;
  margin-bottom: 48px;
}

.home-btn {
  height: 44px;
  padding: 0 28px;
  border-radius: 12px;
  font-size: 15px;
  font-weight: 500;
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  border: none;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.35);
}

.home-btn:hover {
  background: linear-gradient(135deg, #4f46e5 0%, #4338ca 100%);
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(99, 102, 241, 0.4);
}

.back-btn {
  height: 44px;
  padding: 0 28px;
  border-radius: 12px;
  font-size: 15px;
  font-weight: 500;
  border: 2px solid #e5e7eb;
  color: #6b7280;
}

.back-btn:hover {
  border-color: #6366f1;
  color: #6366f1;
}

.error-suggestions {
  background: #fff;
  border-radius: 16px;
  padding: 24px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
}

.suggestion-title {
  font-size: 14px;
  font-weight: 500;
  color: #9ca3af;
  margin-bottom: 16px;
}

.suggestion-list {
  display: flex;
  justify-content: center;
  gap: 24px;
}

.suggestion-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 16px 24px;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.2s;
}

.suggestion-item:hover {
  background: #f3f4f6;
  transform: translateY(-2px);
}

.suggestion-icon {
  font-size: 24px;
  color: #6366f1;
}

.suggestion-item span {
  font-size: 13px;
  color: #374151;
  font-weight: 500;
}

[data-theme='dark'] .not-found {
  background: linear-gradient(135deg, #111827 0%, #0f172a 100%);
}

[data-theme='dark'] .error-code {
  background: linear-gradient(135deg, #818cf8 0%, #6366f1 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

[data-theme='dark'] .error-title {
  color: #e5e7eb;
}

[data-theme='dark'] .error-desc {
  color: #9ca3af;
}

[data-theme='dark'] .back-btn {
  border-color: #374151;
  color: #9ca3af;
}

[data-theme='dark'] .back-btn:hover {
  border-color: #818cf8;
  color: #818cf8;
}

[data-theme='dark'] .error-suggestions {
  background: #1f2937;
}

[data-theme='dark'] .suggestion-item:hover {
  background: #374151;
}

[data-theme='dark'] .suggestion-icon {
  color: #818cf8;
}

[data-theme='dark'] .suggestion-item span {
  color: #e5e7eb;
}
</style>
