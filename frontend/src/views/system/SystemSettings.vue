<!--
  系统设置页面
-->

<template>
  <div class="system-settings">
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <SettingOutlined />
          </div>
          <div class="title-text">
            <h2>系统设置</h2>
            <p class="subtitle">配置系统参数与个性化选项</p>
          </div>
        </div>
      </div>
    </div>

    <div class="settings-container">
      <div class="settings-left">
        <a-card :bordered="false" class="settings-card">
          <div class="card-header">
            <div class="header-left">
              <div class="card-icon-wrapper theme-icon">
                <BgColorsOutlined />
              </div>
              <div class="card-info">
                <span class="card-title">主题设置</span>
                <span class="card-desc">自定义系统外观主题</span>
              </div>
            </div>
          </div>
          <a-divider class="card-divider" />
          <div class="theme-options">
            <div
              class="theme-option"
              :class="{ active: theme === 'light' }"
              @click="selectTheme('light')"
            >
              <div class="theme-preview light-preview">
                <div class="preview-header"></div>
                <div class="preview-sidebar"></div>
                <div class="preview-content"></div>
              </div>
              <div class="theme-info">
                <span class="theme-name">浅色主题</span>
                <span class="theme-desc">明亮清新的界面风格</span>
              </div>
              <div class="theme-check" v-if="theme === 'light'">
                <CheckCircleOutlined />
              </div>
            </div>
            <div
              class="theme-option"
              :class="{ active: theme === 'dark' }"
              @click="selectTheme('dark')"
            >
              <div class="theme-preview dark-preview">
                <div class="preview-header"></div>
                <div class="preview-sidebar"></div>
                <div class="preview-content"></div>
              </div>
              <div class="theme-info">
                <span class="theme-name">深色主题</span>
                <span class="theme-desc">护眼舒适的暗色模式</span>
              </div>
              <div class="theme-check" v-if="theme === 'dark'">
                <CheckCircleOutlined />
              </div>
            </div>
          </div>
        </a-card>

        <a-card :bordered="false" class="settings-card">
          <div class="card-header">
            <div class="header-left">
              <div class="card-icon-wrapper api-icon">
                <ApiOutlined />
              </div>
              <div class="card-info">
                <span class="card-title">API配置</span>
                <span class="card-desc">系统接口与区块链连接配置</span>
              </div>
            </div>
            <a-tag color="blue" class="status-tag">已连接</a-tag>
          </div>
          <a-divider class="card-divider" />
          <div class="config-list">
            <div class="config-item">
              <div class="config-label">
                <GlobalOutlined class="config-icon" />
                <span>API地址</span>
              </div>
              <div class="config-value">
                <span class="value-text">http://localhost:8080</span>
                <a-tag color="green" size="small">在线</a-tag>
              </div>
            </div>
            <div class="config-item">
              <div class="config-label">
                <LinkOutlined class="config-icon" />
                <span>Fabric通道</span>
              </div>
              <div class="config-value">
                <span class="value-text">mychannel</span>
              </div>
            </div>
            <div class="config-item">
              <div class="config-label">
                <CodeOutlined class="config-icon" />
                <span>链码名称</span>
              </div>
              <div class="config-value">
                <span class="value-text">parts</span>
              </div>
            </div>
            <div class="config-item">
              <div class="config-label">
                <ClusterOutlined class="config-icon" />
                <span>节点数量</span>
              </div>
              <div class="config-value">
                <span class="value-text">4 个Peer节点</span>
                <a-tag color="purple" size="small">正常</a-tag>
              </div>
            </div>
          </div>
        </a-card>
      </div>

      <div class="settings-right">
        <a-card :bordered="false" class="settings-card about-card">
          <div class="card-header">
            <div class="header-left">
              <div class="card-icon-wrapper about-icon">
                <InfoCircleOutlined />
              </div>
              <div class="card-info">
                <span class="card-title">关于系统</span>
                <span class="card-desc">系统信息与版本说明</span>
              </div>
            </div>
          </div>
          <a-divider class="card-divider" />
          <div class="about-content">
            <div class="system-logo">
              <div class="logo-icon">
                <CarOutlined />
              </div>
              <div class="logo-text">
                <h3>汽车零部件区块链管理系统</h3>
                <span class="version">Version 1.0.0</span>
              </div>
            </div>
            <div class="info-list">
              <div class="info-item">
                <span class="info-label">技术栈</span>
                <span class="info-value">Vue 3 + TypeScript + Ant Design Vue</span>
              </div>
              <div class="info-item">
                <span class="info-label">区块链</span>
                <span class="info-value">Hyperledger Fabric</span>
              </div>
              <div class="info-item">
                <span class="info-label">开发团队</span>
                <span class="info-value">Automobile Parts Dev Team</span>
              </div>
              <div class="info-item">
                <span class="info-label">更新日期</span>
                <span class="info-value">2024-01-15</span>
              </div>
            </div>
            <div class="system-links">
              <a-button type="link" class="link-btn">
                <template #icon><BookOutlined /></template>
                使用文档
              </a-button>
              <a-button type="link" class="link-btn">
                <template #icon><GithubOutlined /></template>
                源码仓库
              </a-button>
              <a-button type="link" class="link-btn">
                <template #icon><QuestionCircleOutlined /></template>
                帮助中心
              </a-button>
            </div>
          </div>
        </a-card>

        <a-card :bordered="false" class="settings-card stats-card">
          <div class="card-header">
            <div class="header-left">
              <div class="card-icon-wrapper stats-icon">
                <DashboardOutlined />
              </div>
              <div class="card-info">
                <span class="card-title">系统状态</span>
                <span class="card-desc">实时监控与运行指标</span>
              </div>
            </div>
          </div>
          <a-divider class="card-divider" />
          <div class="stats-grid">
            <div class="stat-item">
              <div class="stat-icon cpu">
                <DashboardOutlined />
              </div>
              <div class="stat-content">
                <span class="stat-value">32%</span>
                <span class="stat-label">CPU使用率</span>
              </div>
            </div>
            <div class="stat-item">
              <div class="stat-icon memory">
                <DatabaseOutlined />
              </div>
              <div class="stat-content">
                <span class="stat-value">2.4GB</span>
                <span class="stat-label">内存占用</span>
              </div>
            </div>
            <div class="stat-item">
              <div class="stat-icon disk">
                <HddOutlined />
              </div>
              <div class="stat-content">
                <span class="stat-value">128GB</span>
                <span class="stat-label">存储空间</span>
              </div>
            </div>
            <div class="stat-item">
              <div class="stat-icon network">
                <WifiOutlined />
              </div>
              <div class="stat-content">
                <span class="stat-value">稳定</span>
                <span class="stat-label">网络状态</span>
              </div>
            </div>
          </div>
        </a-card>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { message } from 'ant-design-vue'
import { useAppStore } from '../../stores'
import {
  SettingOutlined,
  BgColorsOutlined,
  CheckCircleOutlined,
  ApiOutlined,
  GlobalOutlined,
  LinkOutlined,
  CodeOutlined,
  ClusterOutlined,
  InfoCircleOutlined,
  CarOutlined,
  BookOutlined,
  GithubOutlined,
  QuestionCircleOutlined,
  DashboardOutlined,
  DatabaseOutlined,
  HdOutlined,
  WifiOutlined,
  HddOutlined
} from '@ant-design/icons-vue'

const appStore = useAppStore()
const theme = ref(appStore.theme)

const selectTheme = (newTheme: string) => {
  theme.value = newTheme
  appStore.setTheme(newTheme)
  message.success(`已切换到${newTheme === 'light' ? '浅色' : '深色'}主题`)
}
</script>

<style scoped>
.system-settings {
  padding: 0;
  min-height: 100%;
}

.page-header {
  margin-bottom: 24px;
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  border-radius: 12px;
  padding: 24px 32px;
  box-shadow: 0 4px 20px rgba(99, 102, 241, 0.25);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-title {
  display: flex;
  align-items: center;
  gap: 16px;
}

.title-icon {
  width: 48px;
  height: 48px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  color: #fff;
}

.title-text h2 {
  font-size: 24px;
  font-weight: 600;
  margin: 0;
  color: #fff;
}

.title-text .subtitle {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.85);
  margin: 4px 0 0;
}

.settings-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
}

.settings-left,
.settings-right {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.settings-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.card-icon-wrapper {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: #fff;
}

.theme-icon {
  background: linear-gradient(135deg, #ec4899 0%, #db2777 100%);
}

.api-icon {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
}

.about-icon {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.stats-icon {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
}

.card-info {
  display: flex;
  flex-direction: column;
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
}

.card-desc {
  font-size: 13px;
  color: #6b7280;
  margin-top: 2px;
}

.status-tag {
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 12px;
}

.card-divider {
  margin: 16px 0;
}

.theme-options {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.theme-option {
  padding: 16px;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
}

.theme-option:hover {
  border-color: #6366f1;
  background: #f9fafb;
}

.theme-option.active {
  border-color: #6366f1;
  background: #eef2ff;
}

.theme-preview {
  width: 100%;
  height: 80px;
  border-radius: 8px;
  overflow: hidden;
  display: grid;
  grid-template-columns: 60px 1fr;
  grid-template-rows: 20px 1fr;
  margin-bottom: 12px;
}

.light-preview {
  background: #f3f4f6;
}

.light-preview .preview-header {
  grid-column: 1 / -1;
  background: #fff;
  border-bottom: 1px solid #e5e7eb;
}

.light-preview .preview-sidebar {
  background: #fff;
  border-right: 1px solid #e5e7eb;
}

.light-preview .preview-content {
  background: #f9fafb;
}

.dark-preview {
  background: #1f2937;
}

.dark-preview .preview-header {
  grid-column: 1 / -1;
  background: #111827;
  border-bottom: 1px solid #374151;
}

.dark-preview .preview-sidebar {
  background: #111827;
  border-right: 1px solid #374151;
}

.dark-preview .preview-content {
  background: #1f2937;
}

.theme-info {
  display: flex;
  flex-direction: column;
}

.theme-name {
  font-size: 14px;
  font-weight: 600;
  color: #1f2937;
}

.theme-desc {
  font-size: 12px;
  color: #6b7280;
  margin-top: 2px;
}

.theme-check {
  position: absolute;
  top: 12px;
  right: 12px;
  color: #6366f1;
  font-size: 18px;
}

.config-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.config-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: #f9fafb;
  border-radius: 10px;
}

.config-label {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
  color: #374151;
}

.config-icon {
  color: #6366f1;
}

.config-value {
  display: flex;
  align-items: center;
  gap: 8px;
}

.value-text {
  font-size: 14px;
  font-weight: 500;
  color: #1f2937;
  font-family: 'Monaco', 'Consolas', monospace;
}

.about-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.system-logo {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 16px;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%);
  border-radius: 12px;
}

.logo-icon {
  width: 56px;
  height: 56px;
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 28px;
  color: #fff;
}

.logo-text h3 {
  font-size: 16px;
  font-weight: 600;
  margin: 0;
  color: #1f2937;
}

.logo-text .version {
  font-size: 13px;
  color: #6b7280;
}

.info-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 0;
  border-bottom: 1px solid #f3f4f6;
}

.info-item:last-child {
  border-bottom: none;
}

.info-label {
  font-size: 13px;
  color: #6b7280;
}

.info-value {
  font-size: 13px;
  font-weight: 500;
  color: #374151;
}

.system-links {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.link-btn {
  padding: 0 12px;
  height: 32px;
  border-radius: 8px;
  font-size: 13px;
  color: #6366f1;
}

.link-btn:hover {
  background: #eef2ff;
  color: #4f46e5;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px;
  background: #f9fafb;
  border-radius: 10px;
}

.stat-icon {
  width: 44px;
  height: 44px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: #fff;
}

.stat-icon.cpu {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
}

.stat-icon.memory {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.stat-icon.disk {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
}

.stat-icon.network {
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%);
}

.stat-content {
  display: flex;
  flex-direction: column;
}

.stat-value {
  font-size: 18px;
  font-weight: 700;
  color: #1f2937;
}

.stat-label {
  font-size: 12px;
  color: #6b7280;
  margin-top: 2px;
}

[data-theme='dark'] .page-header {
  background: linear-gradient(135deg, #4338ca 0%, #3730a3 100%);
  box-shadow: 0 4px 20px rgba(67, 56, 202, 0.3);
}

[data-theme='dark'] .title-text h2 {
  color: #fff;
}

[data-theme='dark'] .title-text .subtitle {
  color: rgba(255, 255, 255, 0.8);
}

[data-theme='dark'] .settings-card {
  background: #1f2937;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

[data-theme='dark'] .card-title {
  color: #f3f4f6;
}

[data-theme='dark'] .card-desc {
  color: #9ca3af;
}

[data-theme='dark'] .theme-option {
  border-color: #374151;
  background: #111827;
}

[data-theme='dark'] .theme-option:hover {
  border-color: #818cf8;
  background: #1f2937;
}

[data-theme='dark'] .theme-option.active {
  border-color: #818cf8;
  background: rgba(129, 140, 248, 0.1);
}

[data-theme='dark'] .theme-name {
  color: #f3f4f6;
}

[data-theme='dark'] .theme-desc {
  color: #9ca3af;
}

[data-theme='dark'] .theme-check {
  color: #a78bfa;
}

[data-theme='dark'] .config-item {
  background: #111827;
}

[data-theme='dark'] .config-label {
  color: #e5e7eb;
}

[data-theme='dark'] .value-text {
  color: #f3f4f6;
}

[data-theme='dark'] .system-logo {
  background: linear-gradient(135deg, #111827 0%, #1f2937 100%);
}

[data-theme='dark'] .logo-text h3 {
  color: #f3f4f6;
}

[data-theme='dark'] .info-item {
  border-bottom-color: #374151;
}

[data-theme='dark'] .info-label {
  color: #9ca3af;
}

[data-theme='dark'] .info-value {
  color: #e5e7eb;
}

[data-theme='dark'] .link-btn {
  color: #a78bfa;
}

[data-theme='dark'] .link-btn:hover {
  background: rgba(129, 140, 248, 0.1);
  color: #c4b5fd;
}

[data-theme='dark'] .stat-item {
  background: #111827;
}

[data-theme='dark'] .stat-value {
  color: #f3f4f6;
}

[data-theme='dark'] .stat-label {
  color: #9ca3af;
}
</style>
