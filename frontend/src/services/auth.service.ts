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

const API_PREFIX = '/api/auth'

export const login = (data: LoginRequest): Promise<ApiResponse<LoginResponse>> => {
  return post(`${API_PREFIX}/login`, data)
}

export const register = (data: RegisterRequest): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/register`, data)
}

export const queryUserPermissions = (userID: string): Promise<ApiResponse<UserIdentity>> => {
  return get(`/api/fabric/users/${userID}/permissions`)
}

export const recordQAInteraction = (data: CreateQAInteractionRequest): Promise<ApiResponse<void>> => {
  return post('/api/fabric/qa', data)
}

export const queryQAInteraction = (qaid: string): Promise<ApiResponse<QAInteraction>> => {
  return get(`/api/fabric/qa/${qaid}`)
}
