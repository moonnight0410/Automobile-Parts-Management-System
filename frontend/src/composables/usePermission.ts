/**
 * 权限组合式函数
 * 提供页面组件内的权限检查功能
 */

import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Modal } from 'ant-design-vue'
import { useAuthStore } from '../stores/auth.store'
import { hasPermission, hasRoutePermission, getCurrentRole } from '../utils/permission'
import { UserRole } from '../types'

/**
 * 权限检查组合式函数
 */
export function usePermission() {
  const router = useRouter()
  const authStore = useAuthStore()
  
  const currentRole = computed(() => authStore.userRole)
  const isAuthenticated = computed(() => authStore.isAuthenticated)
  
  /**
   * 检查是否有指定权限
   * @param permission 权限标识
   * @returns 是否有权限
   */
  function checkPermission(permission: string): boolean {
    return hasPermission(permission)
  }
  
  /**
   * 检查当前路由是否有权限
   * @param path 路由路径
   * @returns 是否有权限
   */
  function checkRoutePermission(path: string): boolean {
    return hasRoutePermission(path)
  }
  
  /**
   * 显示无权限提示并跳转
   * @param permission 权限标识（可选）
   */
  function showNoPermissionAndRedirect(permission?: string) {
    Modal.warning({
      title: '权限提示',
      content: '您没有访问权限',
      okText: '确定',
      onOk: () => {
        if (window.history.length > 1) {
          router.back()
        } else {
          router.push('/dashboard')
        }
      },
    })
  }
  
  /**
   * 页面权限检查
   * 在页面组件的onMounted中调用
   * @param permission 权限标识
   * @returns 是否有权限
   */
  function checkPagePermission(permission: string): boolean {
    if (!isAuthenticated.value) {
      router.push('/login')
      return false
    }
    
    if (!checkPermission(permission)) {
      showNoPermissionAndRedirect(permission)
      return false
    }
    
    return true
  }
  
  /**
   * 检查是否是制造商
   */
  function isManufacturer(): boolean {
    return currentRole.value === UserRole.MANUFACTURER
  }
  
  /**
   * 检查是否是车企
   */
  function isAutomaker(): boolean {
    return currentRole.value === UserRole.AUTOMAKER
  }
  
  /**
   * 检查是否是售后中心
   */
  function isAftersale(): boolean {
    return currentRole.value === UserRole.AFTERSALE
  }
  
  return {
    currentRole,
    isAuthenticated,
    checkPermission,
    checkRoutePermission,
    checkPagePermission,
    showNoPermissionAndRedirect,
    isManufacturer,
    isAutomaker,
    isAftersale,
  }
}

/**
 * 页面权限守卫组合式函数
 * 用于在页面组件内进行权限二次检查
 * @param permission 所需权限标识
 */
export function usePageGuard(permission: string) {
  const { checkPagePermission, currentRole } = usePermission()
  const hasAccess = ref(false)
  
  onMounted(() => {
    hasAccess.value = checkPagePermission(permission)
  })
  
  return {
    hasAccess,
    currentRole,
  }
}
