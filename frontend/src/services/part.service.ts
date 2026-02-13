/**
 * 零部件管理API服务
 * 提供零部件相关的API调用方法
 */

import { get, post, put } from './axios'
import type { 
  Part, 
  PartLifecycle, 
  CreatePartRequest, 
  UpdatePartStatusRequest,
  ApiResponse 
} from '../types'

// API路径前缀
const API_PREFIX = '/api/fabric'

/**
 * 创建零部件
 * @param data 零部件创建请求数据
 * @returns 创建结果
 */
export const createPart = (data: CreatePartRequest): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/parts`, data)
}

/**
 * 查询零部件详情
 * @param partID 零部件ID
 * @returns 零部件详情
 */
export const queryPart = (partID: string): Promise<ApiResponse<Part>> => {
  return get(`${API_PREFIX}/parts/${partID}`)
}

/**
 * 查询零部件生命周期
 * @param partID 零部件ID
 * @returns 零部件生命周期信息
 */
export const queryPartLifecycle = (partID: string): Promise<ApiResponse<PartLifecycle>> => {
  return get(`${API_PREFIX}/parts/${partID}/lifecycle`)
}

/**
 * 按批次号查询零部件
 * @param batchNo 批次号
 * @returns 零部件列表
 */
export const queryPartByBatchNo = (batchNo: string): Promise<ApiResponse<Part[]>> => {
  return get(`${API_PREFIX}/parts/batch/${batchNo}`)
}

/**
 * 按VIN码查询零部件
 * @param vin VIN码
 * @returns 零部件列表
 */
export const queryPartByVIN = (vin: string): Promise<ApiResponse<Part[]>> => {
  return get(`${API_PREFIX}/parts/vin/${vin}`)
}

/**
 * 更新零部件状态
 * @param data 状态更新请求数据
 * @returns 更新结果
 */
export const updatePartStatus = (data: UpdatePartStatusRequest): Promise<ApiResponse<void>> => {
  return put(`${API_PREFIX}/parts/${data.partID}/status`, { status: data.status })
}

/**
 * 初始化账本
 * @returns 初始化结果
 */
export const initLedger = (): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/init`)
}

/**
 * 健康检查
 * @returns 健康状态
 */
export const healthCheck = (): Promise<ApiResponse<{ status: string }>> => {
  return get(`${API_PREFIX}/health`)
}
