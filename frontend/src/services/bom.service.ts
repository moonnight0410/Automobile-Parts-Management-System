import { get, post, put, del } from './axios'
import type { ApiResponse } from '../types'

export interface BOMMaterial {
  materialID: string
  materialName: string
  quantity: number
  unit: string
  specifications?: string
}

export interface BOM {
  bomID: string
  bomType: string
  productModel: string
  partBatchNo: string
  version: string
  creator: string
  createTime: string
  status: string
  materialList: BOMMaterial[]
  rdBOMFileInfo?: {
    fileName: string
    fileURL: string
    uploadTime: string
  }
  productionBOMInfo?: {
    productionLine: string
    processFlow: string
    qualityStandard: string
  }
  changeRecords?: {
    changeID: string
    changeTime: string
    changeContent: string
    changer: string
  }[]
}

export interface BOMDTO {
  bomID: string
  bomType: string
  productModel: string
  partBatchNo: string
  version: string
  creator: string
  createTime: string
  status: string
  materialList: BOMMaterial[]
  rdBOMFileInfo?: {
    fileName: string
    fileURL: string
    uploadTime: string
  }
  productionBOMInfo?: {
    productionLine: string
    processFlow: string
    qualityStandard: string
  }
}

export const createBOM = (data: BOMDTO): Promise<ApiResponse<void>> => {
  return post('/api/boms', data)
}

export const getBOM = (bomID: string): Promise<ApiResponse<BOM>> => {
  return get(`/api/boms/${bomID}`)
}

export const listBOMs = (): Promise<ApiResponse<BOM[]>> => {
  return get('/api/boms')
}

export const deleteBOM = (bomID: string): Promise<ApiResponse<void>> => {
  return del(`/api/boms/${bomID}`)
}

export const updateBOM = (bomID: string, data: BOMDTO): Promise<ApiResponse<void>> => {
  return put(`/api/boms/${bomID}`, data)
}
