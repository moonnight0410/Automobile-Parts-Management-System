/**
 * 通用类型定义
 * 定义系统中通用的数据类型和接口
 */

// API响应通用结构
export interface ApiResponse<T = any> {
  code: number
  message: string
  data?: T
}

// 分页请求参数
export interface PaginationParams {
  page: number
  pageSize: number
  keyword?: string
}

// 分页响应数据
export interface PaginationData<T> {
  list: T[]
  total: number
  page: number
  pageSize: number
}

// 用户角色枚举
export enum UserRole {
  MANUFACTURER = 'manufacturer', // 制造商
  AUTOMAKER = 'automaker',       // 车企
  AFTERSALE = 'aftersale'        // 售后中心
}

// 用户信息
export interface UserInfo {
  id: string
  username: string
  role: UserRole
  org: string
  token: string
}

// 主题类型
export type ThemeMode = 'light' | 'dark'

// 加载状态
export interface LoadingState {
  loading: boolean
  error: string | null
}
