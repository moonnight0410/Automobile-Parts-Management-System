/**
 * AI智能售后助手API服务
 * 提供AI问答相关的API调用方法
 */

import { get, post } from './axios'
import type { ApiResponse } from '../types'

// AI问答请求
export interface AIQuestionRequest {
  question: string
  userID: string
  partID?: string
}

// AI问答响应
export interface AIQuestionResponse {
  qaid: string
  question: string
  intent: string
  confidence: number
  answer: string
  relatedActions?: Array<{
    label: string
    action: string
    params?: Record<string, any>
  }>
  relatedFaqs?: string[]
}

// AI统计信息
export interface AIStats {
  conversationCount: number
  userCount: number
}

// 对话历史
export interface ConversationHistory {
  messages: ConversationMessage[]
}

// 对话消息
export interface ConversationMessage {
  role: string
  content: string
  timestamp: number
  metadata?: Record<string, any>
}

/**
 * 发送AI问题
 * @param data 问题请求数据
 * @returns AI回答
 */
export const sendAIQuestion = async (data: AIQuestionRequest): Promise<ApiResponse<AIQuestionResponse>> => {
  return post<AIQuestionResponse>('/api/ai/question', data)
}

/**
 * 查询对话历史
 * @param userID 用户ID
 * @param n 返回的对话轮数
 * @returns 对话历史
 */
export const getConversationHistory = async (userID: string, n: number = 10): Promise<ApiResponse<ConversationHistory>> => {
  return get<ConversationHistory>('/api/ai/conversation', { params: { userID, n } })
}

/**
 * 清空对话历史
 * @param userID 用户ID
 * @returns 操作结果
 */
export const clearConversationHistory = async (userID: string): Promise<ApiResponse<null>> => {
  return del<null>('/api/ai/conversation', { params: { userID } })
}

/**
 * 获取统计信息
 * @returns 统计数据
 */
export const getAIStats = async (): Promise<ApiResponse<AIStats>> => {
  return get<AIStats>('/api/ai/stats')
}

