<!--
  主布局组件
  包含侧边栏、顶部导航、内容区域
-->

<template>
  <a-layout class="main-layout">
    <!-- 侧边栏 -->
    <a-layout-sider
      v-model:collapsed="sidebarCollapsed"
      :trigger="null"
      collapsible
      :width="240"
      :collapsed-width="80"
      class="sidebar"
    >
      <!-- Logo -->
      <div class="logo">
        <img src="/vite.svg" alt="Logo" class="logo-img" />
        <span v-show="!sidebarCollapsed" class="logo-text">零部件管理系统</span>
      </div>
      
      <!-- 导航菜单 -->
      <a-menu
        v-model:selectedKeys="selectedKeys"
        v-model:openKeys="openKeys"
        mode="inline"
        :theme="theme"
        @click="handleMenuClick"
      >
        <!-- 仪表盘 -->
        <a-menu-item key="/dashboard">
          <template #icon>
            <DashboardOutlined />
          </template>
          <span>仪表盘</span>
        </a-menu-item>
        
        <!-- 零部件管理 -->
        <a-sub-menu key="parts">
          <template #icon>
            <AppstoreOutlined />
          </template>
          <template #title>零部件管理</template>
          <a-menu-item key="/parts/list">零部件列表</a-menu-item>
          <a-menu-item key="/parts/create" v-if="canCreatePart">创建零部件</a-menu-item>
        </a-sub-menu>
        
        <!-- BOM管理 -->
        <a-sub-menu key="bom" v-if="canManageBOM">
          <template #icon>
            <FileTextOutlined />
          </template>
          <template #title>BOM管理</template>
          <a-menu-item key="/bom/list">BOM列表</a-menu-item>
          <a-menu-item key="/bom/create" v-if="isManufacturer">创建BOM</a-menu-item>
          <a-menu-item key="/bom/compare">BOM比较</a-menu-item>
        </a-sub-menu>
        
        <!-- 生产管理 -->
        <a-sub-menu key="production" v-if="isManufacturer">
          <template #icon>
            <ToolOutlined />
          </template>
          <template #title>生产管理</template>
          <a-menu-item key="/production/data">生产数据</a-menu-item>
          <a-menu-item key="/production/quality">质检管理</a-menu-item>
        </a-sub-menu>
        
        <!-- 供应链管理 -->
        <a-sub-menu key="supply" v-if="canManageSupply">
          <template #icon>
            <SwapOutlined />
          </template>
          <template #title>供应链管理</template>
          <a-menu-item key="/supply/orders">采购订单</a-menu-item>
          <a-menu-item key="/supply/logistics">物流跟踪</a-menu-item>
        </a-sub-menu>
        
        <!-- 售后管理 -->
        <a-sub-menu key="aftersale" v-if="canManageAftersale">
          <template #icon>
            <CustomerServiceOutlined />
          </template>
          <template #title>售后管理</template>
          <a-menu-item key="/aftersale/fault">故障报告</a-menu-item>
          <a-menu-item key="/aftersale/recall">召回记录</a-menu-item>
          <a-menu-item key="/aftersale/assistant">智能售后助手</a-menu-item>
        </a-sub-menu>
        
        <!-- 区块链浏览器 -->
        <a-menu-item key="/blockchain">
          <template #icon>
            <LinkOutlined />
          </template>
          <span>区块链浏览器</span>
        </a-menu-item>
        
        <!-- 系统管理 -->
        <a-sub-menu key="system">
          <template #icon>
            <SettingOutlined />
          </template>
          <template #title>系统管理</template>
          <a-menu-item key="/system/users">用户管理</a-menu-item>
          <a-menu-item key="/system/settings">系统设置</a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
    
    <!-- 右侧内容区 -->
    <a-layout class="main-content" :style="{ marginLeft: sidebarCollapsed ? '80px' : '240px' }">
      <!-- 顶部导航栏 -->
      <a-layout-header class="header">
        <div class="header-left">
          <!-- 折叠按钮 -->
          <span class="trigger" @click="toggleSidebar">
            <MenuUnfoldOutlined v-if="sidebarCollapsed" />
            <MenuFoldOutlined v-else />
          </span>
          
          <!-- 面包屑导航 -->
          <a-breadcrumb class="breadcrumb">
            <a-breadcrumb-item v-for="item in breadcrumbs" :key="item.path">
              {{ item.title }}
            </a-breadcrumb-item>
          </a-breadcrumb>
        </div>
        
        <div class="header-right">
          <!-- 主题切换 -->
          <button class="theme-toggle" @click="toggleTheme" :title="theme === 'light' ? '切换到深色模式' : '切换到浅色模式'">
            <svg v-if="theme === 'light'" class="theme-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <circle cx="12" cy="12" r="5" stroke="currentColor" stroke-width="2"/>
              <line x1="12" y1="1" x2="12" y2="3" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="12" y1="21" x2="12" y2="23" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="4.22" y1="4.22" x2="5.64" y2="5.64" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="18.36" y1="18.36" x2="19.78" y2="19.78" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="1" y1="12" x2="3" y2="12" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="21" y1="12" x2="23" y2="12" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="4.22" y1="19.78" x2="5.64" y2="18.36" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="18.36" y1="5.64" x2="19.78" y2="4.22" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
            </svg>
            <svg v-else class="theme-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </button>
          
          <!-- 用户信息 -->
          <a-dropdown>
            <div class="user-info">
              <a-avatar :style="{ backgroundColor: '#1890ff' }">
                {{ user?.username?.charAt(0).toUpperCase() }}
              </a-avatar>
              <span class="username">{{ user?.username }}</span>
              <DownOutlined />
            </div>
            <template #overlay>
              <a-menu>
                <a-menu-item key="profile">
                  <UserOutlined />
                  <span>个人信息</span>
                </a-menu-item>
                <a-menu-divider />
                <a-menu-item key="logout" @click="handleLogout">
                  <LogoutOutlined />
                  <span>退出登录</span>
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </div>
      </a-layout-header>
      
      <!-- 内容区域 -->
      <a-layout-content class="content">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </a-layout-content>
      
      <!-- 底部 -->
      <a-layout-footer class="footer">
        汽车零部件区块链管理系统 ©2024 Created by Blockchain Team
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { message } from 'ant-design-vue'
import { useAuthStore, useAppStore } from '../stores'
import { UserRole } from '../types'

// ==================== 组合式API ====================

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()
const appStore = useAppStore()

// ==================== 响应式状态 ====================

// 选中的菜单项
const selectedKeys = ref<string[]>([route.path])

// 展开的子菜单
const openKeys = ref<string[]>([])

// ==================== 计算属性 ====================

// 侧边栏折叠状态
const sidebarCollapsed = computed({
  get: () => appStore.sidebarCollapsed,
  set: (value) => appStore.setSidebarCollapsed(value)
})

// 当前主题
const theme = computed(() => appStore.theme)

// 当前用户
const user = computed(() => authStore.user)

// 用户角色
const userRole = computed(() => authStore.userRole)

// 是否是制造商
const isManufacturer = computed(() => userRole.value === UserRole.MANUFACTURER)

// 是否是车企
const isAutomaker = computed(() => userRole.value === UserRole.AUTOMAKER)

// 是否是售后中心
const isAftersale = computed(() => userRole.value === UserRole.AFTERSALE)

// 是否可以创建零部件
const canCreatePart = computed(() => isManufacturer.value)

// 是否可以管理BOM
const canManageBOM = computed(() => isManufacturer.value || isAutomaker.value)

// 是否可以管理供应链
const canManageSupply = computed(() => isManufacturer.value || isAutomaker.value)

// 是否可以管理售后
const canManageAftersale = computed(() => isAftersale.value || isAutomaker.value)

// 面包屑导航
const breadcrumbs = computed(() => {
  const matched = route.matched.filter(item => item.meta && item.meta.title)
  return matched.map(item => ({
    path: item.path,
    title: item.meta.title as string
  }))
})

// ==================== 方法 ====================

/**
 * 切换侧边栏折叠状态
 */
function toggleSidebar() {
  appStore.toggleSidebar()
}

/**
 * 切换主题
 */
function toggleTheme() {
  appStore.toggleTheme()
}

/**
 * 处理菜单点击
 */
function handleMenuClick({ key }: { key: string }) {
  router.push(key)
}

/**
 * 处理退出登录
 */
function handleLogout() {
  authStore.logout()
  message.success('已退出登录')
  router.push('/login')
}

// ==================== 监听路由变化 ====================

watch(
  () => route.path,
  (path) => {
    selectedKeys.value = [path]
    
    // 自动展开父菜单
    const parentPath = '/' + path.split('/')[1]
    if (parentPath && !openKeys.value.includes(parentPath)) {
      openKeys.value.push(parentPath)
    }
  },
  { immediate: true }
)
</script>

<style scoped>
.main-layout {
  min-height: 100vh;
}

.main-content {
  transition: margin-left 0.2s;
}

.sidebar {
  box-shadow: 2px 0 8px rgba(0, 0, 0, 0.1);
  z-index: 10;
  height: 100vh;
  position: fixed;
  left: 0;
  top: 0;
  overflow: auto;
}

.sidebar :deep(.ant-layout-sider-children) {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.sidebar :deep(.ant-menu) {
  border-inline-end: none !important;
}

.sidebar :deep(.ant-menu-item) {
  margin: 4px 8px;
  border-radius: 8px;
  transition: all 0.3s ease;
  border: 1px solid rgba(0, 0, 0, 0.06);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.04);
}

.sidebar :deep(.ant-menu-item:hover) {
  border-color: rgba(0, 0, 0, 0.12);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.08);
}

.sidebar :deep(.ant-menu-item-selected) {
  background: linear-gradient(135deg, #1890ff 0%, #40a9ff 100%) !important;
  color: #fff !important;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.35);
  border-color: transparent;
}

.sidebar :deep(.ant-menu-item-selected)::after {
  display: none;
}

.sidebar :deep(.ant-menu-item-selected .anticon) {
  color: #fff !important;
}

.sidebar :deep(.ant-menu-submenu-title) {
  margin: 4px 8px;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.sidebar :deep(.ant-menu-submenu-selected > .ant-menu-submenu-title) {
  background: rgba(24, 144, 255, 0.15);
  color: #1890ff;
  font-weight: 600;
}

.sidebar :deep(.ant-menu-sub) {
  background: transparent !important;
}

.sidebar :deep(.ant-menu-sub .ant-menu-item) {
  padding-left: 48px !important;
}

.logo {
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 16px;
  background: rgba(255, 255, 255, 0.1);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.logo-img {
  width: 32px;
  height: 32px;
}

.logo-text {
  margin-left: 12px;
  font-size: 16px;
  font-weight: bold;
  color: #141414;
  white-space: nowrap;
}

[data-theme='dark'] .logo-text {
  color: #fff;
}

.header {
  background: var(--bg-color-secondary);
  padding: 0 24px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  z-index: 9;
}

.header-left {
  display: flex;
  align-items: center;
}

.trigger {
  font-size: 18px;
  cursor: pointer;
  transition: color 0.3s;
  padding: 0 16px;
}

.trigger:hover {
  color: #1890ff;
}

.breadcrumb {
  margin-left: 16px;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 16px;
  height: 100%;
}

.theme-toggle {
  width: 40px;
  height: 40px;
  border: 2px solid #e0e0e0;
  border-radius: 50%;
  background: #ffffff;
  color: #666;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  flex-shrink: 0;
}

.theme-toggle:hover {
  border-color: #1890ff;
  color: #1890ff;
  transform: rotate(15deg);
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.25);
}

.theme-icon {
  width: 20px;
  height: 20px;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  padding: 8px 12px;
  border-radius: 4px;
  transition: background 0.3s;
  height: 40px;
}

.user-info:hover {
  background: rgba(0, 0, 0, 0.025);
}

.username {
  max-width: 100px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.content {
  margin: 24px;
  padding: 24px;
  background: var(--bg-color-secondary);
  border-radius: 8px;
  min-height: calc(100vh - 64px - 70px - 48px);
  overflow: auto;
}

.footer {
  text-align: center;
  background: transparent;
  padding: 16px 50px;
  color: var(--text-color-secondary);
}

/* 过渡动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* 深色主题样式 */
[data-theme='dark'] .sidebar {
  background: var(--bg-color-secondary);
}

[data-theme='dark'] .logo {
  background: rgba(255, 255, 255, 0.05);
  border-bottom-color: rgba(255, 255, 255, 0.1);
}

[data-theme='dark'] .header {
  background: var(--bg-color-secondary);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
}

[data-theme='dark'] .user-info:hover {
  background: rgba(255, 255, 255, 0.05);
}

[data-theme='dark'] .theme-toggle {
  background: var(--bg-color-tertiary);
  border-color: rgba(255, 255, 255, 0.2);
  color: #aaa;
}

[data-theme='dark'] .theme-toggle:hover {
  border-color: #1890ff;
  color: #1890ff;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.35);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-item-selected) {
  background: linear-gradient(135deg, #1890ff 0%, #40a9ff 100%) !important;
  color: #fff !important;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.45);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-item) {
  border-color: rgba(255, 255, 255, 0.08);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.15);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-item:hover) {
  border-color: rgba(255, 255, 255, 0.15);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.25);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-submenu-selected > .ant-menu-submenu-title) {
  background: rgba(24, 144, 255, 0.25);
  color: #40a9ff;
}
</style>
