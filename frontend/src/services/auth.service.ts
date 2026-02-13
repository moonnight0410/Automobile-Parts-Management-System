/**
 * 认证API服务
 * 提供用户认证相关的API调用方法
 */

import { get, post } from './axios'
import type { 
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  UserIdentity,
  QAInteraction,
  CreateQAInteractionRequest,
  ApiResponse 
} from '../types'

// API路径前缀
const API_PREFIX = '/api/fabric'

/**
 * 用户注册
 * @param data 注册请求数据
 * @returns 注册结果
 */
export const registerUser = (data: RegisterRequest): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/users`, data)
}

/**
 * 查询用户权限
 * @param userID 用户ID
 * @returns 用户身份信息
 */
export const queryUserPermissions = (userID: string): Promise<ApiResponse<UserIdentity>> => {
  return get(`${API_PREFIX}/users/${userID}/permissions`)
}

/**
 * 记录QA交互
 * @param data QA交互请求数据
 * @returns 记录结果
 */
export const recordQAInteraction = (data: CreateQAInteractionRequest): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/qa`, data)
}

/**
 * 查询QA交互记录
 * @param qaid QA交互ID
 * @returns QA交互详情
 */
export const queryQAInteraction = (qaid: string): Promise<ApiResponse<QAInteraction>> => {
  return get(`${API_PREFIX}/qa/${qaid}`)
}

/**
 * 模拟登录（临时使用）
 * 注意：实际项目中应该调用真实的登录API
 * @param data 登录请求数据
 * @returns 登录响应
 */
export const mockLogin = (data: LoginRequest): Promise<ApiResponse<LoginResponse>> => {
  // 模拟登录响应
  return new Promise((resolve) => {
    setTimeout(() => {
      const mockResponse: ApiResponse<LoginResponse> = {
        success: true,
        message: '登录成功',
        data: {
          token: 'mock-jwt-token-' + Date.now(),
          user: {
            id: 'user-' + data.username,
            username: data.username,
            role: 'manufacturer' as any,
            org: 'Org1MSP'
          }
        }
      }
      resolve(mockResponse as any)
    }, 500)
  })
}
