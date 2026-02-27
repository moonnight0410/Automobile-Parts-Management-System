import { get, post, del } from './axios'
import type { ApiResponse } from '../types'

export interface QualityInspection {
  inspectionID: string
  partID: string
  batchNo: string
  inspectionType: string
  inspector: string
  handler: string
  inspectionTime: string
  inspectionResult: string
  defectDescription?: string
  handlingMethod?: string
  status: string
}

export interface QualityInspectionDTO {
  inspectionID: string
  partID: string
  batchNo: string
  inspectionType: string
  inspector: string
  handler: string
  inspectionTime: string
  inspectionResult: string
  defectDescription?: string
  handlingMethod?: string
  status: string
}

export const createQualityInspection = (data: QualityInspectionDTO): Promise<ApiResponse<void>> => {
  return post('/api/quality', data)
}

export const listQualityInspections = (): Promise<ApiResponse<QualityInspection[]>> => {
  return get('/api/quality')
}

type Dict = Record<string, string>

export interface BackendQualityDTO {
  inspectionID: string
  partID: string
  batchNo: string
  indicators?: Dict
  result: string
  handler: string
  handleTime: string
  disposal?: string
}

export const createBackendQualityInspection = (data: BackendQualityDTO): Promise<ApiResponse<void>> => {
  return post('/api/quality', data)
}

export const createFabricQualityInspection = (data: BackendQualityDTO): Promise<ApiResponse<void>> => {
  return post('/api/fabric/quality', data)
}

export const createQualityInspectionWithChain = async (data: BackendQualityDTO): Promise<ApiResponse<void>> => {
  await createBackendQualityInspection(data)
  const resp = await createFabricQualityInspection(data)
  return resp
}

export const deleteQualityInspection = (inspectionID: string): Promise<ApiResponse<void>> => {
  return del(`/api/fabric/quality/${inspectionID}`)
}
