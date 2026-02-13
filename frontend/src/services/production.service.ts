/**
 * 生产管理API服务
 * 提供生产数据和质检相关的API调用方法
 */

import { post } from './axios'
import type { 
  ProductionData,
  QualityInspection,
  ApiResponse 
} from '../types'

// API路径前缀
const API_PREFIX = '/api/fabric'

/**
 * 创建生产数据
 * @param data 生产数据
 * @returns 创建结果
 */
export const createProductionData = (data: ProductionData): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/production`, data)
}

/**
 * 创建质检数据
 * @param data 质检数据
 * @returns 创建结果
 */
export const createQualityInspection = (data: QualityInspection): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/quality`, data)
}
