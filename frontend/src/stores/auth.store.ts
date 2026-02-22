/**
 * 认证状态管理
 * 管理用户登录状态、用户信息、Token等
 */

import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { mockLogin, registerUser } from '../services/auth.service'
import { STORAGE_KEYS } from '../constants'
import type { User, LoginRequest, RegisterRequest, UserRole } from '../types'
import { hasPermission as checkPermission } from '../utils/permission'

export const useAuthStore = defineStore('auth', () => {
  // ==================== 状态定义 ====================
  
  const token = ref<string | null>(localStorage.getItem(STORAGE_KEYS.TOKEN))
  const user = ref<User | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)
  
  // ==================== 计算属性 ====================
  
  const isAuthenticated = computed(() => !!token.value)
  const userRole = computed(() => user.value?.role || null)
  const userOrg = computed(() => user.value?.org || null)
  
  const isManufacturer = computed(() => userRole.value === 'manufacturer')
  const isAutomaker = computed(() => userRole.value === 'automaker')
  const isAftersale = computed(() => userRole.value === 'aftersale')
  
  // ==================== Actions ====================
  
  /**
   * 用户登录
   */
  async function login(username: string, password: string) {
    loading.value = true
    error.value = null
    
    try {
      const request: LoginRequest = { username, password }
      const response = await mockLogin(request)
      
      if (response.success && response.data) {
        token.value = response.data.token
        user.value = response.data.user
        
        localStorage.setItem(STORAGE_KEYS.TOKEN, response.data.token)
        localStorage.setItem(STORAGE_KEYS.USER, JSON.stringify(response.data.user))
        
        return response
      } else {
        throw new Error(response.message || '登录失败')
      }
    } catch (err: any) {
      error.value = err.message || '登录失败'
      throw err
    } finally {
      loading.value = false
    }
  }
  
  /**
   * 用户注册
   */
  async function register(data: RegisterRequest) {
    loading.value = true
    error.value = null
    
    try {
      const response = await registerUser(data)
      
      if (!response.success) {
        throw new Error(response.message || '注册失败')
      }
      
      return response
    } catch (err: any) {
      error.value = err.message || '注册失败'
      throw err
    } finally {
      loading.value = false
    }
  }
  
  /**
   * 用户登出
   */
  function logout() {
    token.value = null
    user.value = null
    
    localStorage.removeItem(STORAGE_KEYS.TOKEN)
    localStorage.removeItem(STORAGE_KEYS.USER)
  }
  
  /**
   * 从localStorage恢复用户信息
   */
  function restoreUser() {
    const savedUser = localStorage.getItem(STORAGE_KEYS.USER)
    if (savedUser) {
      try {
        user.value = JSON.parse(savedUser)
      } catch (e) {
        console.error('Failed to parse saved user:', e)
        logout()
      }
    }
  }
  
  /**
   * 检查用户是否有指定角色
   */
  function hasRole(role: UserRole | UserRole[]): boolean {
    if (!user.value) return false
    
    const roles = Array.isArray(role) ? role : [role]
    return roles.includes(user.value.role)
  }
  
  /**
   * 检查用户是否有指定权限
   */
  function hasPermission(permission: string): boolean {
    return checkPermission(permission)
  }
  
  restoreUser()
  
  // ==================== 导出 ====================
  
  return {
    token,
    user,
    loading,
    error,
    
    isAuthenticated,
    userRole,
    userOrg,
    isManufacturer,
    isAutomaker,
    isAftersale,
    
    login,
    register,
    logout,
    restoreUser,
    hasRole,
    hasPermission
  }
})
