import { get, post } from './axios'
import type { ApiResponse } from '../types'

export interface Part {
  partID: string
  name: string
  type: string
  manufacturer: string
  batchNo: string
  productionDate: string
  vin?: string
  status: string
  qualityGrade?: string
  specifications?: string
  materialInfo?: string
  supplierInfo?: string
  createTime?: string
}

export interface PartDTO {
  partID: string
  name: string
  type: string
  manufacturer: string
  batchNo: string
  productionDate: string
  vin?: string
  status: string
  qualityGrade?: string
  specifications?: string
  materialInfo?: string
  supplierInfo?: string
}

export const createPart = (data: PartDTO): Promise<ApiResponse<void>> => {
  return post('/api/parts', data)
}

export const getPart = (partID: string): Promise<ApiResponse<Part>> => {
  return get(`/api/parts/${partID}`)
}

export const listParts = (params: { batchNo?: string; vin?: string }): Promise<ApiResponse<Part[]>> => {
  return get('/api/parts', { params })
}

export const listMyParts = (): Promise<ApiResponse<Part[]>> => {
  return get('/api/parts/my')
}

export const queryPart = getPart

export const queryPartByBatchNo = (batchNo: string): Promise<ApiResponse<Part[]>> => {
  return listParts({ batchNo })
}

export const queryPartByVIN = (vin: string): Promise<ApiResponse<Part[]>> => {
  return listParts({ vin })
}

export const queryPartLifecycle = (partID: string): Promise<ApiResponse<any>> => {
  return get(`/api/parts/${partID}/lifecycle`)
}

export const updatePartStatus = (data: { partID: string; status: string }): Promise<ApiResponse<void>> => {
  return post(`/api/parts/${data.partID}/status`, data)
}
