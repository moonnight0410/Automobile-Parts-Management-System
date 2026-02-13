/**
 * AI智能售后助手API服务
 * 提供AI问答相关的API调用方法
 * 注意：当前使用Mock数据，后续需要对接真实的AI后端
 */

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
  answer: string
  relatedActions?: Array<{
    label: string
    action: string
    params?: Record<string, any>
  }>
}

// Mock知识库数据
const MOCK_KNOWLEDGE_BASE: Record<string, { intent: string; answer: string; actions?: any[] }> = {
  '故障上报': {
    intent: '故障上报咨询',
    answer: '您可以通过以下步骤上报零部件故障：\n1. 进入"售后管理 > 故障报告"页面\n2. 点击"新建故障报告"按钮\n3. 填写零部件ID、VIN码、故障描述等信息\n4. 提交后系统会自动分配故障ID并开始处理流程',
    actions: [
      { label: '立即上报故障', action: 'navigate', params: { path: '/aftersale/fault' } }
    ]
  },
  '故障进度': {
    intent: '故障进度查询',
    answer: '您可以通过故障ID查询处理进度。请提供您的故障ID，我将为您查询当前处理状态。',
    actions: [
      { label: '查询故障进度', action: 'query', params: { type: 'fault' } }
    ]
  },
  '召回': {
    intent: '召回相关查询',
    answer: '召回查询需要提供零部件批次号。系统会自动检测该批次是否在召回范围内，并显示召回详情和处理流程。',
    actions: [
      { label: '查询召回信息', action: 'query', params: { type: 'recall' } }
    ]
  },
  '质保': {
    intent: '售后政策咨询',
    answer: '零部件质保政策如下：\n- 标准质保期：12个月或2万公里（以先到为准）\n- 延保服务：可购买延保套餐，最长延至36个月\n- 质保范围：正常使用条件下的材料或工艺缺陷\n- 不在质保范围：人为损坏、事故损坏、自然磨损等',
    actions: [
      { label: '查看完整政策', action: 'navigate', params: { path: '/aftersale/policy' } }
    ]
  },
  '维修': {
    intent: '售后政策咨询',
    answer: '故障零部件维修后的质保政策：\n- 维修更换的零部件享有3个月质保期\n- 维修服务需在授权售后中心进行\n- 维修记录将自动上链存证，可追溯查询',
    actions: [
      { label: '预约维修', action: 'navigate', params: { path: '/aftersale/repair' } }
    ]
  }
}

/**
 * 发送AI问题
 * @param data 问题请求数据
 * @returns AI回答
 */
export const sendAIQuestion = async (data: AIQuestionRequest): Promise<ApiResponse<AIQuestionResponse>> => {
  // 模拟AI处理延迟
  await new Promise(resolve => setTimeout(resolve, 800))
  
  // 简单的关键词匹配
  let matchedKey = '故障上报' // 默认回复
  for (const key of Object.keys(MOCK_KNOWLEDGE_BASE)) {
    if (data.question.includes(key)) {
      matchedKey = key
      break
    }
  }
  
  const knowledge = MOCK_KNOWLEDGE_BASE[matchedKey]
  
  // 构造Mock响应
  const response: ApiResponse<AIQuestionResponse> = {
    success: true,
    message: '查询成功',
    data: {
      qaid: `QA-${Date.now()}`,
      question: data.question,
      intent: knowledge.intent,
      answer: knowledge.answer,
      relatedActions: knowledge.actions
    }
  }
  
  return response
}

/**
 * 查询故障处理进度（AI辅助）
 * @param faultID 故障ID
 * @returns 处理进度信息
 */
export const queryFaultProgressAI = async (faultID: string): Promise<ApiResponse<string>> => {
  await new Promise(resolve => setTimeout(resolve, 500))
  
  // Mock故障进度数据
  const mockProgress = {
    'FAULT-001': '您的故障报告已审核通过，已标记为召回批次，待更换零部件。',
    'FAULT-002': '您的故障报告正在审核中，预计1-2个工作日内完成审核。',
    'FAULT-003': '您的故障报告已完成处理，零部件已更换，请前往售后中心取车。'
  }
  
  const response: ApiResponse<string> = {
    success: true,
    message: '查询成功',
    data: mockProgress[faultID] || `未找到故障ID为 ${faultID} 的处理记录，请确认ID是否正确。`
  }
  
  return response
}

/**
 * 查询召回信息（AI辅助）
 * @param batchNo 批次号
 * @returns 召回信息
 */
export const queryRecallInfoAI = async (batchNo: string): Promise<ApiResponse<string>> => {
  await new Promise(resolve => setTimeout(resolve, 500))
  
  // Mock召回信息
  const response: ApiResponse<string> = {
    success: true,
    message: '查询成功',
    data: batchNo.startsWith('RECALL') 
      ? `批次 ${batchNo} 在召回范围内。召回原因：制动系统潜在隐患。请尽快联系售后中心安排免费更换。`
      : `批次 ${batchNo} 不在召回范围内。如有疑问，请联系客服。`
  }
  
  return response
}
