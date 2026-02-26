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
        <template v-for="menu in accessibleMenus" :key="menu.key">
          <!-- 单级菜单 -->
          <a-menu-item v-if="!menu.children || menu.children.length === 0" :key="menu.key">
            <template #icon>
              <component :is="getIconComponent(menu.icon)" />
            </template>
            <span>{{ menu.title }}</span>
          </a-menu-item>
          
          <!-- 多级菜单 -->
          <a-sub-menu v-else :key="menu.key">
            <template #icon>
              <component :is="getIconComponent(menu.icon)" />
            </template>
            <template #title>{{ menu.title }}</template>
            <a-menu-item
              v-for="child in menu.children"
              :key="child.key"
            >
              {{ child.title }}
            </a-menu-item>
          </a-sub-menu>
        </template>
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
          <a-dropdown
            :trigger="['hover']"
            placement="bottomRight"
            :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
            :mouse-enter-delay="0.1"
            :mouse-leave-delay="0.1"
          >
            <div class="user-info">
              <a-avatar :style="{ backgroundColor: '#6366f1' }">
                {{ user?.username?.charAt(0).toUpperCase() }}
              </a-avatar>
              <span class="username">{{ user?.username }}</span>
              <span class="role-tag">{{ roleName }}</span>
              <DownOutlined class="dropdown-arrow" />
            </div>
            <template #overlay>
              <a-menu class="user-dropdown-menu">
                <a-menu-item key="profile" class="dropdown-item">
                  <UserOutlined />
                  <span>个人信息</span>
                </a-menu-item>
                <a-menu-divider />
                <a-menu-item key="logout" @click="handleLogout" class="dropdown-item logout-item">
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
        汽车零部件区块链管理系统 ©2026 Created by JianYu Zou
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, shallowRef } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { message, Modal } from 'ant-design-vue'
import {
  DashboardOutlined,
  AppstoreOutlined,
  FileTextOutlined,
  ToolOutlined,
  SwapOutlined,
  CustomerServiceOutlined,
  LinkOutlined,
  SettingOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  DownOutlined,
  UserOutlined,
  LogoutOutlined
} from '@ant-design/icons-vue'
import { useAuthStore, useAppStore } from '../stores'
import { UserRole } from '../types'
import { getAccessibleMenus, getRoleName } from '../utils/permission'
import { getAndClearNoPermissionMessage } from '../middleware/auth'
import type { MenuConfig } from '../config/permissions'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()
const appStore = useAppStore()

const selectedKeys = ref<string[]>([route.path])
const openKeys = ref<string[]>([])

const sidebarCollapsed = computed({
  get: () => appStore.sidebarCollapsed,
  set: (value) => appStore.setSidebarCollapsed(value)
})

const theme = computed(() => appStore.theme)
const user = computed(() => authStore.user)
const userRole = computed(() => authStore.userRole)
const roleName = computed(() => userRole.value ? getRoleName(userRole.value) : '')

const accessibleMenus = computed(() => getAccessibleMenus())

const breadcrumbs = computed(() => {
  const matched = route.matched.filter(item => item.meta && item.meta.title)
  return matched.map(item => ({
    path: item.path,
    title: item.meta.title as string
  }))
})

const iconMap: Record<string, any> = {
  DashboardOutlined,
  AppstoreOutlined,
  FileTextOutlined,
  ToolOutlined,
  SwapOutlined,
  CustomerServiceOutlined,
  LinkOutlined,
  SettingOutlined
}

function getIconComponent(iconName?: string) {
  if (!iconName) return null
  return iconMap[iconName] || null
}

function toggleSidebar() {
  appStore.toggleSidebar()
}

function toggleTheme() {
  appStore.toggleTheme()
}

function handleMenuClick({ key }: { key: string }) {
  router.push(key)
}

function handleLogout() {
  authStore.logout()
  message.success('已退出登录')
  router.push('/login')
}

watch(
  () => route.path,
  (path) => {
    selectedKeys.value = [path]
    
    const parentPath = '/' + path.split('/')[1]
    if (parentPath && !openKeys.value.includes(parentPath)) {
      openKeys.value.push(parentPath)
    }
    
    const noPermMsg = getAndClearNoPermissionMessage()
    if (noPermMsg) {
      Modal.warning({
        title: '权限提示',
        content: noPermMsg,
        okText: '确定'
      })
    }
    
    if (route.meta.noPermissionMessage) {
      Modal.warning({
        title: '权限提示',
        content: route.meta.noPermissionMessage as string,
        okText: '确定'
      })
    }
  },
  { immediate: true }
)

onMounted(() => {
  const parentPath = '/' + route.path.split('/')[1]
  if (parentPath) {
    openKeys.value = [parentPath]
  }
  
  const noPermMsg = getAndClearNoPermissionMessage()
  if (noPermMsg) {
    noPermissionMessage.value = noPermMsg
    showNoPermissionModal.value = true
  }
})
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
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border: 1px solid transparent;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.04);
  position: relative;
  overflow: hidden;
}

.sidebar :deep(.ant-menu-item::before) {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.1) 0%, rgba(64, 169, 255, 0.1) 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
  z-index: -1;
}

.sidebar :deep(.ant-menu-item:hover) {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.1) 0%, rgba(64, 169, 255, 0.1) 100%);
  border-color: rgba(24, 144, 255, 0.3);
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.15);
  color: #1890ff;
  transform: translateX(4px);
}

.sidebar :deep(.ant-menu-item:hover::before) {
  opacity: 1;
}

.sidebar :deep(.ant-menu-item:hover .anticon) {
  color: #1890ff;
  transform: scale(1.1);
}

.sidebar :deep(.ant-menu-item .anticon) {
  transition: all 0.3s ease;
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
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border: 1px solid transparent;
  position: relative;
  overflow: hidden;
}

.sidebar :deep(.ant-menu-submenu-title:hover) {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.1) 0%, rgba(64, 169, 255, 0.1) 100%);
  border-color: rgba(24, 144, 255, 0.3);
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.15);
  color: #1890ff;
  transform: translateX(4px);
}

.sidebar :deep(.ant-menu-submenu-title:hover .anticon) {
  color: #1890ff;
  transform: scale(1.1);
}

.sidebar :deep(.ant-menu-submenu-title .anticon) {
  transition: all 0.3s ease;
}

.sidebar :deep(.ant-menu-submenu-selected > .ant-menu-submenu-title) {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.15) 0%, rgba(64, 169, 255, 0.15) 100%);
  color: #1890ff;
  font-weight: 600;
  border-color: rgba(24, 144, 255, 0.2);
}

.sidebar :deep(.ant-menu-sub) {
  background: transparent !important;
}

.sidebar :deep(.ant-menu-sub .ant-menu-item) {
  padding-left: 48px !important;
  margin: 2px 8px;
  border-radius: 6px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border: 1px solid transparent;
}

.sidebar :deep(.ant-menu-sub .ant-menu-item:hover) {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.08) 0%, rgba(64, 169, 255, 0.08) 100%);
  border-color: rgba(24, 144, 255, 0.25);
  box-shadow: 0 2px 8px rgba(24, 144, 255, 0.1);
  color: #1890ff;
  transform: translateX(2px);
}

.sidebar :deep(.ant-menu-sub .ant-menu-item-selected) {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.15) 0%, rgba(64, 169, 255, 0.15) 100%) !important;
  color: #1890ff !important;
  font-weight: 600;
  border-color: rgba(24, 144, 255, 0.3);
  box-shadow: 0 2px 8px rgba(24, 144, 255, 0.15);
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
  border-radius: 10px;
  transition: all 0.3s ease;
  height: 40px;
  background: transparent;
  border: 1px solid transparent;
}

.user-info:hover {
  background: rgba(99, 102, 241, 0.08);
  border-color: rgba(99, 102, 241, 0.2);
}

.dropdown-arrow {
  font-size: 10px;
  color: #6b7280;
  transition: transform 0.3s ease;
}

.user-info:hover .dropdown-arrow {
  color: #6366f1;
  transform: rotate(180deg);
}

.username {
  max-width: 100px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-weight: 500;
  color: #374151;
}

.user-info:hover .username {
  color: #6366f1;
}

.role-tag {
  font-size: 12px;
  padding: 2px 8px;
  border-radius: 4px;
  background: rgba(99, 102, 241, 0.1);
  color: #6366f1;
}

.user-dropdown-menu {
  min-width: 180px;
  padding: 8px 0;
  border-radius: 12px;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.12);
  border: 1px solid #f0f0f0;
  overflow: hidden;
}

.dropdown-item {
  padding: 12px 16px !important;
  margin: 0 !important;
  border-radius: 0 !important;
  display: flex;
  align-items: center;
  gap: 10px;
  transition: all 0.2s ease;
}

.dropdown-item:hover {
  background: #f5f3ff !important;
  color: #6366f1 !important;
}

.dropdown-item :deep(.anticon) {
  font-size: 14px;
}

.logout-item:hover {
  background: #fef2f2 !important;
  color: #ef4444 !important;
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

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

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
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.2) 0%, rgba(64, 169, 255, 0.2) 100%);
  border-color: rgba(24, 144, 255, 0.4);
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.25);
  color: #40a9ff;
  transform: translateX(4px);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-item:hover .anticon) {
  color: #40a9ff;
  transform: scale(1.1);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-submenu-title) {
  border-color: rgba(255, 255, 255, 0.08);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-submenu-title:hover) {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.2) 0%, rgba(64, 169, 255, 0.2) 100%);
  border-color: rgba(24, 144, 255, 0.4);
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.25);
  color: #40a9ff;
  transform: translateX(4px);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-submenu-title:hover .anticon) {
  color: #40a9ff;
  transform: scale(1.1);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-submenu-selected > .ant-menu-submenu-title) {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.25) 0%, rgba(64, 169, 255, 0.25) 100%);
  color: #40a9ff;
  border-color: rgba(24, 144, 255, 0.3);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-sub .ant-menu-item) {
  border-color: rgba(255, 255, 255, 0.05);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-sub .ant-menu-item:hover) {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.15) 0%, rgba(64, 169, 255, 0.15) 100%);
  border-color: rgba(24, 144, 255, 0.35);
  box-shadow: 0 2px 8px rgba(24, 144, 255, 0.2);
  color: #40a9ff;
  transform: translateX(2px);
}

[data-theme='dark'] .sidebar :deep(.ant-menu-sub .ant-menu-item-selected) {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.25) 0%, rgba(64, 169, 255, 0.25) 100%) !important;
  color: #40a9ff !important;
  border-color: rgba(24, 144, 255, 0.4);
  box-shadow: 0 2px 8px rgba(24, 144, 255, 0.25);
}

[data-theme='dark'] .user-info:hover {
  background: rgba(167, 139, 250, 0.15);
  border-color: rgba(167, 139, 250, 0.3);
}

[data-theme='dark'] .username {
  color: #e5e7eb;
}

[data-theme='dark'] .user-info:hover .username {
  color: #a78bfa;
}

[data-theme='dark'] .dropdown-arrow {
  color: #9ca3af;
}

[data-theme='dark'] .user-info:hover .dropdown-arrow {
  color: #a78bfa;
}

[data-theme='dark'] .user-dropdown-menu {
  background: #1f2937;
  border-color: #374151;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.4);
}

[data-theme='dark'] .dropdown-item {
  color: #e5e7eb;
}

[data-theme='dark'] .dropdown-item:hover {
  background: rgba(167, 139, 250, 0.15) !important;
  color: #a78bfa !important;
}

[data-theme='dark'] .logout-item:hover {
  background: rgba(239, 68, 68, 0.15) !important;
  color: #f87171 !important;
}

[data-theme='dark'] .role-tag {
  background: rgba(167, 139, 250, 0.2);
  color: #a78bfa;
}
</style>
