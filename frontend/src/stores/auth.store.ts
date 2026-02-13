/**
 * 认证状态管理
 * 管理用户登录状态、用户信息、Token等
 */

import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { mockLogin, registerUser } from '../services/auth.service'
import { STORAGE_KEYS } from '../constants'
import type { User, LoginRequest, RegisterRequest, UserRole } from '../types'

export const useAuthStore = defineStore('auth', () => {
  // ==================== 状态定义 ====================
  
  // 用户Token
  const token = ref<string | null>(localStorage.getItem(STORAGE_KEYS.TOKEN))
  
  // 用户信息
  const user = ref<User | null>(null)
  
  // 加载状态
  const loading = ref(false)
  
  // 错误信息
  const error = ref<string | null>(null)
  
  // ==================== 计算属性 ====================
  
  // 是否已登录
  const isAuthenticated = computed(() => !!token.value)
  
  // 用户角色
  const userRole = computed(() => user.value?.role || null)
  
  // 用户组织
  const userOrg = computed(() => user.value?.org || null)
  
  // ==================== Actions ====================
  
  /**
   * 用户登录
   * @param username 用户名
   * @param password 密码
   * @returns 登录结果
   */
  async function login(username: string, password: string) {
    loading.value = true
    error.value = null
    
    try {
      const request: LoginRequest = { username, password }
      const response = await mockLogin(request)
      
      if (response.success && response.data) {
        // 保存Token和用户信息
        token.value = response.data.token
        user.value = response.data.user
        
        // 持久化到localStorage
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
   * @param data 注册数据
   * @returns 注册结果
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
    
    // 清除localStorage
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
   * @param role 角色列表
   * @returns 是否有权限
   */
  function hasRole(role: UserRole | UserRole[]): boolean {
    if (!user.value) return false
    
    const roles = Array.isArray(role) ? role : [role]
    return roles.includes(user.value.role)
  }
  
  // 初始化时恢复用户信息
  restoreUser()
  
  // ==================== 导出 ====================
  
  return {
    // 状态
    token,
    user,
    loading,
    error,
    
    // 计算属性
    isAuthenticated,
    userRole,
    userOrg,
    
    // Actions
    login,
    register,
    logout,
    restoreUser,
    hasRole
  }
})
