/**
 * 认证相关类型定义
 * 定义用户认证和授权的数据类型和接口
 */

import { UserRole } from './common'

// 登录请求
export interface LoginRequest {
  username: string
  password: string
}

// 登录响应
export interface LoginResponse {
  token: string
  userID: string
  role: string
}

// 用户信息
export interface User {
  id: string
  username: string
  role: UserRole
  org: string
  permissions?: string[]
}

// 注册请求
export interface RegisterRequest {
  userID: string
  password: string
  org: string
  role: UserRole
  certHash?: string
  permissions?: string[]
}

// 用户身份
export interface UserIdentity {
  userID: string
  org: string
  role: UserRole
  certHash: string
  permissions: string[]
  registerTime: string
}

// QA交互记录
export interface QAInteraction {
  qaid: string           // 问答记录ID
  userID: string         // 提问用户ID
  question: string       // 问题内容
  intent: string         // 意图分类
  partID: string         // 关联零部件ID
  answer: string         // 回答内容
  queryTime: string      // 提问时间
  contractMethod: string // 调用的智能合约方法
}

// 创建QA交互请求
export interface CreateQAInteractionRequest {
  qaid: string
  userID: string
  question: string
  intent?: string
  partID?: string
  answer: string
  contractMethod?: string
}
