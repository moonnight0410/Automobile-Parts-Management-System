/**
 * 认证状态管理
 * 管理用户登录状态、用户信息、Token等
 */

import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login as loginApi, register as registerApi } from '../services/auth.service'
import { STORAGE_KEYS } from '../constants'
import type { User, LoginRequest, RegisterRequest, UserRole } from '../types'
import { hasPermission as checkPermission } from '../utils/permission'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem(STORAGE_KEYS.TOKEN))
  const user = ref<User | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)
  
  const isAuthenticated = computed(() => !!token.value)
  const userRole = computed(() => user.value?.role || null)
  const userOrg = computed(() => user.value?.org || null)
  
  const isManufacturer = computed(() => userRole.value === 'manufacturer')
  const isAutomaker = computed(() => userRole.value === 'automaker')
  const isAftersale = computed(() => userRole.value === 'aftersale')
  
  async function login(username: string, password: string) {
    loading.value = true
    error.value = null
    
    try {
      const request: LoginRequest = { username, password }
      const response = await loginApi(request)
      
      if (response.code === 0 && response.data) {
        token.value = response.data.token
        user.value = {
          id: response.data.userID,
          username: response.data.userID,
          role: response.data.role as UserRole,
          org: ''
        }
        
        localStorage.setItem(STORAGE_KEYS.TOKEN, response.data.token)
        localStorage.setItem(STORAGE_KEYS.USER, JSON.stringify(user.value))
        
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
  
  async function register(data: RegisterRequest) {
    loading.value = true
    error.value = null
    
    try {
      const response = await registerApi(data)
      
      if (response.code !== 0) {
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
  
  function logout() {
    token.value = null
    user.value = null
    
    localStorage.removeItem(STORAGE_KEYS.TOKEN)
    localStorage.removeItem(STORAGE_KEYS.USER)
  }
  
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
  
  function hasRole(role: UserRole | UserRole[]): boolean {
    if (!user.value) return false
    
    const roles = Array.isArray(role) ? role : [role]
    return roles.includes(user.value.role)
  }
  
  function hasPermission(permission: string): boolean {
    return checkPermission(permission)
  }
  
  restoreUser()
  
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
