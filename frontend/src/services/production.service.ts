import { get, post } from './axios'
import type { ApiResponse } from '../types'

export interface ProductionData {
  recordID: string
  partID: string
  batchNo: string
  productionLine: string
  operator: string
  operationType: string
  operationTime: string
  operationResult: string
  qualityStatus: string
  remarks?: string
}

export interface ProductionDataDTO {
  recordID: string
  partID: string
  batchNo: string
  productionLine: string
  operator: string
  operationType: string
  operationTime: string
  operationResult: string
  qualityStatus: string
  remarks?: string
}

export const createProductionData = (data: ProductionDataDTO): Promise<ApiResponse<void>> => {
  return post('/api/production', data)
}

export const listProductionData = (): Promise<ApiResponse<ProductionData[]>> => {
  return get('/api/production')
}
