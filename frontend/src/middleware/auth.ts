/**
 * 路由权限中间件
 * 处理路由导航守卫，验证用户权限
 */

import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { STORAGE_KEYS } from '../constants'
import { hasPermission, isAuthenticated, getCurrentRole } from '../utils/permission'
import { ROUTE_PERMISSION_MAP } from '../config/permissions'
import { UserRole } from '../types'

const NO_PERMISSION_KEY = 'no_permission_message'

/**
 * 设置无权限消息到sessionStorage
 */
export function setNoPermissionMessage(msg: string) {
  sessionStorage.setItem(NO_PERMISSION_KEY, msg)
}

/**
 * 获取并清除无权限消息
 */
export function getAndClearNoPermissionMessage(): string | null {
  const msg = sessionStorage.getItem(NO_PERMISSION_KEY)
  if (msg) {
    sessionStorage.removeItem(NO_PERMISSION_KEY)
  }
  return msg
}

/**
 * 检查路由权限
 */
function checkRoutePermission(path: string): boolean {
  const permission = ROUTE_PERMISSION_MAP[path]
  if (!permission) {
    return true
  }
  return hasPermission(permission)
}

/**
 * 认证中间件
 * 检查用户是否已登录，以及是否有权限访问目标路由
 */
export default function AuthMiddleware(
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  next: NavigationGuardNext
) {
  const requiresAuth = to.meta.requiresAuth !== false
  
  if (!requiresAuth) {
    return next()
  }
  
  if (!isAuthenticated()) {
    return next({
      path: '/login',
      query: { redirect: to.fullPath }
    })
  }
  
  const currentRole = getCurrentRole()
  const requiredRoles = to.meta.roles as UserRole[] | undefined
  
  if (requiredRoles && requiredRoles.length > 0 && currentRole) {
    if (!requiredRoles.includes(currentRole)) {
      setNoPermissionMessage('您没有访问权限')
      return next({ path: '/dashboard' })
    }
  }
  
  const path = to.path
  if (!checkRoutePermission(path)) {
    setNoPermissionMessage('您没有访问权限')
    return next({ path: '/dashboard' })
  }
  
  next()
}

/**
 * 设置路由元信息类型
 */
declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth?: boolean
    roles?: UserRole[]
    title?: string
    icon?: string
    permission?: string
    noPermissionMessage?: string
  }
}
