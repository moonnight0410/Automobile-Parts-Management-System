/**
 * 应用全局状态管理
 * 管理主题、侧边栏、加载状态等全局状态
 */

import { defineStore } from 'pinia'
import { ref, watch } from 'vue'
import { STORAGE_KEYS } from '../constants'
import type { ThemeMode } from '../types'

export const useAppStore = defineStore('app', () => {
  // ==================== 状态定义 ====================
  
  // 当前主题
  const theme = ref<ThemeMode>(
    (localStorage.getItem(STORAGE_KEYS.THEME) as ThemeMode) || 'light'
  )
  
  // 侧边栏折叠状态
  const sidebarCollapsed = ref(false)
  
  // 全局加载状态
  const globalLoading = ref(false)
  
  // 页面标题
  const pageTitle = ref('汽车零部件区块链管理系统')
  
  // ==================== 监听主题变化 ====================
  
  watch(theme, (newTheme) => {
    // 保存到localStorage
    localStorage.setItem(STORAGE_KEYS.THEME, newTheme)
    
    // 更新HTML class
    document.documentElement.setAttribute('data-theme', newTheme)
    
    // 更新body class（用于Ant Design Vue主题切换）
    if (newTheme === 'dark') {
      document.body.classList.add('dark')
      document.body.classList.remove('light')
    } else {
      document.body.classList.add('light')
      document.body.classList.remove('dark')
    }
  }, { immediate: true })
  
  // ==================== Actions ====================
  
  /**
   * 切换主题
   */
  function toggleTheme() {
    theme.value = theme.value === 'light' ? 'dark' : 'light'
  }
  
  /**
   * 设置主题
   * @param newTheme 新主题
   */
  function setTheme(newTheme: ThemeMode) {
    theme.value = newTheme
  }
  
  /**
   * 切换侧边栏折叠状态
   */
  function toggleSidebar() {
    sidebarCollapsed.value = !sidebarCollapsed.value
  }
  
  /**
   * 设置侧边栏折叠状态
   * @param collapsed 是否折叠
   */
  function setSidebarCollapsed(collapsed: boolean) {
    sidebarCollapsed.value = collapsed
  }
  
  /**
   * 设置全局加载状态
   * @param loading 是否加载中
   */
  function setGlobalLoading(loading: boolean) {
    globalLoading.value = loading
  }
  
  /**
   * 设置页面标题
   * @param title 标题
   */
  function setPageTitle(title: string) {
    pageTitle.value = title
    document.title = title
  }
  
  // ==================== 导出 ====================
  
  return {
    // 状态
    theme,
    sidebarCollapsed,
    globalLoading,
    pageTitle,
    
    // Actions
    toggleTheme,
    setTheme,
    toggleSidebar,
    setSidebarCollapsed,
    setGlobalLoading,
    setPageTitle
  }
})
