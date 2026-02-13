/**
 * 售后管理API服务
 * 提供售后相关的API调用方法
 */

import { get, post } from './axios'
import type { 
  FaultReport,
  RecallRecord,
  AftersaleRecord,
  CreateFaultReportRequest,
  CreateRecallRecordRequest,
  CreateAftersaleRecordRequest,
  Part,
  ApiResponse 
} from '../types'

// API路径前缀
const API_PREFIX = '/api/fabric'

/**
 * 创建故障报告
 * @param data 故障报告创建请求数据
 * @returns 创建结果
 */
export const createFaultReport = (data: CreateFaultReportRequest): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/faults`, data)
}

/**
 * 查询故障处理进度
 * @param faultID 故障ID
 * @returns 故障报告详情
 */
export const queryFaultProgress = (faultID: string): Promise<ApiResponse<FaultReport>> => {
  return get(`${API_PREFIX}/faults/${faultID}`)
}

/**
 * 创建召回记录
 * @param data 召回记录创建请求数据
 * @returns 创建结果
 */
export const createRecallRecord = (data: CreateRecallRecordRequest): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/recalls`, data)
}

/**
 * 查询受影响零部件
 * @param batchNo 批次号
 * @returns 受影响零部件列表
 */
export const queryAffectedParts = (batchNo: string): Promise<ApiResponse<Part[]>> => {
  return get(`${API_PREFIX}/recalls/affected/${batchNo}`)
}

/**
 * 创建售后记录
 * @param data 售后记录创建请求数据
 * @returns 创建结果
 */
export const createAftersaleRecord = (data: CreateAftersaleRecordRequest): Promise<ApiResponse<void>> => {
  return post(`${API_PREFIX}/aftersales`, data)
}

/**
 * 查询售后记录
 * @param aftersaleID 售后记录ID
 * @returns 售后记录详情
 */
export const queryAftersaleRecord = (aftersaleID: string): Promise<ApiResponse<AftersaleRecord>> => {
  return get(`${API_PREFIX}/aftersales/${aftersaleID}`)
}
