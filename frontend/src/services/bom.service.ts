/**
 * BOM管理API服务
 * 提供BOM相关的API调用方法
 */

import { get, post, put } from './axios'
import type { 
  BOM, 
  CreateBOMRequest, 
  UpdateBOMRequest,
  BOMCompareRequest,
  BOMChangeRequest,
  ApiResponse 
} from '../types'

// API路径前缀
const API_PREFIX = '/api/fabric'

/**
 * 创建BOM
 * @param data BOM创建请求数据
 * @returns 创建结果
 */
export const createBOM = (data: CreateBOMRequest): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/bom`, data)
}

/**
 * 查询BOM详情
 * @param bomID BOM ID
 * @returns BOM详情
 */
export const queryBOM = (bomID: string): Promise<ApiResponse<BOM>> => {
  return get(`${API_PREFIX}/bom/${bomID}`)
}

/**
 * 更新BOM
 * @param data BOM更新请求数据
 * @returns 更新结果
 */
export const updateBOM = (data: UpdateBOMRequest): Promise<ApiResponse<void>> => {
  return put(`${API_PREFIX}/bom/${data.bomID}`, data)
}

/**
 * BOM比较
 * @param data BOM比较请求数据
 * @returns 比较结果
 */
export const compareBOM = (data: BOMCompareRequest): Promise<ApiResponse<string>> => {
  return post(`${API_PREFIX}/bom/compare`, data)
}

/**
 * 提交BOM变更
 * @param data BOM变更请求数据
 * @returns 变更结果
 */
export const submitBOMChange = (data: BOMChangeRequest): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/bom/change`, data)
}
