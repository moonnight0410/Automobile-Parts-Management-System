import { get, post, del } from './axios'
import type { ApiResponse } from '../types'

export interface ProductionData {
  productionID: string
  partID: string
  batchNo: string
  productionLine: string
  operator: string
  finishTime: string
  params: Record<string, string>
}

export interface ProductionDataDTO {
  productionID: string
  partID: string
  batchNo: string
  productionLine: string
  operator: string
  finishTime: string
  params: Record<string, string>
}

export const createProductionData = (data: ProductionDataDTO): Promise<ApiResponse<void>> => {
  return post('/api/production', data)
}

export const listProductionData = (): Promise<ApiResponse<ProductionData[]>> => {
  return get('/api/production')
}

export const deleteProductionData = (productionID: string): Promise<ApiResponse<void>> => {
  return del(`/api/fabric/production/${productionID}`)
}
