import { get, post, put } from './axios'
import type { ApiResponse } from '../types'

export interface FaultReport {
  faultID: string
  vin: string
  partID: string
  faultType: string
  faultDescription: string
  reportTime: string
  reporter: string
  reporterOrg: string
  status: string
  handlingResult?: string
  handlingTime?: string
}

export interface FaultReportDTO {
  faultID: string
  vin: string
  partID: string
  faultType: string
  faultDescription: string
  reportTime: string
  reporter: string
  reporterOrg: string
  status: string
}

export interface RecallRecord {
  recallID: string
  recallType: string
  affectedParts: string[]
  affectedVins: string[]
  recallReason: string
  recallDate: string
  initiator: string
  status: string
  handlingMeasures?: string
}

export interface RecallRecordDTO {
  recallID: string
  recallType: string
  affectedParts: string[]
  affectedVins: string[]
  recallReason: string
  recallDate: string
  initiator: string
  status: string
  handlingMeasures?: string
}

export interface AftersaleRecord {
  recordID: string
  vin: string
  partID: string
  serviceType: string
  serviceDescription: string
  serviceTime: string
  handlerOrgID: string
  handlerName: string
  cost: number
  status: string
}

export interface AftersaleRecordDTO {
  recordID: string
  vin: string
  partID: string
  serviceType: string
  serviceDescription: string
  serviceTime: string
  handlerOrgID: string
  handlerName: string
  cost: number
  status: string
}

export const createFaultReport = (data: FaultReportDTO): Promise<ApiResponse<void>> => {
  return post('/api/faults', data)
}

export const listFaultReports = (): Promise<ApiResponse<FaultReport[]>> => {
  return get('/api/faults')
}

export const createRecallRecord = (data: RecallRecordDTO): Promise<ApiResponse<void>> => {
  return post('/api/recalls', data)
}

export const listRecallRecords = (): Promise<ApiResponse<RecallRecord[]>> => {
  return get('/api/recalls')
}

export const createAftersaleRecord = (data: AftersaleRecordDTO): Promise<ApiResponse<void>> => {
  return post('/api/aftersale-records', data)
}

export const listAftersaleRecords = (): Promise<ApiResponse<AftersaleRecord[]>> => {
  return get('/api/aftersale-records')
}

export const updateFaultReportStatus = (faultID: string, status: string): Promise<ApiResponse<void>> => {
  return put('/api/faults/status', { faultID, status })
}
