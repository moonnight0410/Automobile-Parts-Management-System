/**
 * 供应链管理API服务
 * 提供供应链相关的API调用方法
 */

import { get, post, put } from './axios'
import type { 
  SupplyOrder,
  LogisticsData,
  Reconciliation,
  SupplyChainStage,
  CreateSupplyOrderRequest,
  CreateLogisticsDataRequest,
  ApiResponse 
} from '../types'

// API路径前缀
const API_PREFIX = '/api/fabric'

/**
 * 创建供应链订单
 * @param data 订单创建请求数据
 * @returns 创建结果
 */
export const createSupplyOrder = (data: CreateSupplyOrderRequest): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/supply/orders`, data)
}

/**
 * 创建物流数据
 * @param data 物流数据创建请求数据
 * @returns 创建结果
 */
export const createLogisticsData = (data: CreateLogisticsDataRequest): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/supply/logistics`, data)
}

/**
 * 更新供应链阶段
 * @param data 阶段更新请求数据
 * @returns 更新结果
 */
export const updateSupplyChainStage = (data: SupplyChainStage): Promise<ApiResponse<void>> => {
  return put(`${API_PREFIX}/supply/stage`, data)
}

/**
 * 创建对账记录
 * @param data 对账创建请求数据
 * @returns 创建结果
 */
export const createReconciliation = (data: Reconciliation): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/supply/reconciliation`, data)
}
