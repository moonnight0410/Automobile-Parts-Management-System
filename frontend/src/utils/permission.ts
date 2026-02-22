/**
 * 权限工具函数
 * 提供权限检查相关的工具函数
 */

import { UserRole } from '../types'
import { PAGE_PERMISSIONS, MENU_CONFIG, ROUTE_PERMISSION_MAP, type MenuConfig } from '../config/permissions'
import { STORAGE_KEYS } from '../constants'

/**
 * 获取当前用户信息
 */
export function getCurrentUser(): { role: UserRole; username: string; org: string } | null {
  const userStr = localStorage.getItem(STORAGE_KEYS.USER)
  if (!userStr) return null
  
  try {
    return JSON.parse(userStr)
  } catch (e) {
    console.error('解析用户信息失败:', e)
    return null
  }
}

/**
 * 获取当前用户角色
 */
export function getCurrentRole(): UserRole | null {
  const user = getCurrentUser()
  return user?.role || null
}

/**
 * 检查用户角色是否在允许的角色列表中
 * @param roles 允许的角色列表
 * @returns 是否有权限
 */
export function hasRole(roles: UserRole[]): boolean {
  const currentRole = getCurrentRole()
  if (!currentRole) return false
  return roles.includes(currentRole)
}

/**
 * 检查是否有权限访问某页面
 * @param permission 权限标识
 * @returns 是否有权限
 */
export function hasPermission(permission: string): boolean {
  const allowedRoles = PAGE_PERMISSIONS[permission]
  if (!allowedRoles) {
    console.warn(`未找到权限配置: ${permission}`)
    return false
  }
  return hasRole(allowedRoles)
}

/**
 * 检查路由路径是否有权限访问
 * @param path 路由路径
 * @returns 是否有权限
 */
export function hasRoutePermission(path: string): boolean {
  const permission = ROUTE_PERMISSION_MAP[path]
  if (!permission) {
    return true
  }
  return hasPermission(permission)
}

/**
 * 过滤菜单项，只保留有权限的菜单
 * @param menus 菜单配置列表
 * @returns 过滤后的菜单列表
 */
export function filterMenusByPermission(menus: MenuConfig[]): MenuConfig[] {
  const result: MenuConfig[] = []
  
  for (const menu of menus) {
    if (menu.children && menu.children.length > 0) {
      const filteredChildren = filterMenusByPermission(menu.children)
      if (filteredChildren.length > 0) {
        result.push({
          ...menu,
          children: filteredChildren,
        })
      }
    } else if (menu.permission) {
      if (hasPermission(menu.permission)) {
        result.push(menu)
      }
    } else {
      result.push(menu)
    }
  }
  
  return result
}

/**
 * 获取用户可访问的菜单列表
 * @returns 可访问的菜单列表
 */
export function getAccessibleMenus(): MenuConfig[] {
  return filterMenusByPermission(MENU_CONFIG)
}

/**
 * 检查用户是否已登录
 */
export function isAuthenticated(): boolean {
  const token = localStorage.getItem(STORAGE_KEYS.TOKEN)
  return !!token
}

/**
 * 获取权限标识对应的路由路径
 * @param permission 权限标识
 * @returns 路由路径
 */
export function getRouteByPermission(permission: string): string | undefined {
  for (const [path, perm] of Object.entries(ROUTE_PERMISSION_MAP)) {
    if (perm === permission) {
      return path
    }
  }
  return undefined
}

/**
 * 获取用户角色的中文名称
 * @param role 用户角色
 * @returns 角色中文名称
 */
export function getRoleName(role: UserRole): string {
  const roleNames: Record<UserRole, string> = {
    [UserRole.MANUFACTURER]: '制造商',
    [UserRole.AUTOMAKER]: '车企',
    [UserRole.AFTERSALE]: '售后中心',
  }
  return roleNames[role] || '未知角色'
}
