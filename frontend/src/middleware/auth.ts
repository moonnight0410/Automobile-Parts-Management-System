/**
 * 路由权限中间件
 * 处理路由导航守卫，验证用户权限
 */

import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { useAuthStore } from '../stores/auth.store'
import { STORAGE_KEYS } from '../constants'

/**
 * 认证中间件
 * 检查用户是否已登录，以及是否有权限访问目标路由
 */
export default function AuthMiddleware(
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  next: NavigationGuardNext
) {
  // 获取路由元信息
  const requiresAuth = to.meta.requiresAuth !== false // 默认需要认证
  const requiredRoles = to.meta.roles as string[] | undefined
  
  // 如果不需要认证，直接放行
  if (!requiresAuth) {
    return next()
  }
  
  // 检查是否已登录
  const token = localStorage.getItem(STORAGE_KEYS.TOKEN)
  if (!token) {
    // 未登录，跳转到登录页
    return next({
      path: '/login',
      query: { redirect: to.fullPath }
    })
  }
  
  // 如果路由指定了需要的角色，检查用户角色
  if (requiredRoles && requiredRoles.length > 0) {
    // 从localStorage获取用户信息
    const userStr = localStorage.getItem(STORAGE_KEYS.USER)
    if (!userStr) {
      return next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    }
    
    try {
      const user = JSON.parse(userStr)
      if (!requiredRoles.includes(user.role)) {
        // 用户角色不符合要求，跳转到首页
        console.warn(`用户角色 ${user.role} 无权访问路由 ${to.path}`)
        return next({ path: '/' })
      }
    } catch (e) {
      console.error('解析用户信息失败:', e)
      return next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    }
  }
  
  // 通过所有检查，放行
  next()
}

/**
 * 设置路由元信息类型
 */
declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth?: boolean
    roles?: string[]
    title?: string
    icon?: string
  }
}
