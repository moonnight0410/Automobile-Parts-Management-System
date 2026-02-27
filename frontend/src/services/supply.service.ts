import { get, post, del } from './axios'
import type { ApiResponse } from '../types'

export interface SupplyOrder {
  orderID: string
  buyer: string
  seller: string
  partID: string
  quantity: number
  batchNo: string
  agreedTime: string
  status: string
  createTime: string
}

export interface SupplyOrderDTO {
  orderID: string
  buyer: string
  seller: string
  partID: string
  quantity: number
  batchNo: string
  agreedTime: string
  status: string
  createTime?: string
}

export interface LogisticsData {
  logisticsID: string
  orderID: string
  carrier: string
  startTime: string
  endTime: string
  gpsHash: string
  receiver: string
}

export interface LogisticsDataDTO {
  logisticsID: string
  orderID: string
  carrier: string
  startTime?: string
  endTime?: string
  gpsHash?: string
  receiver?: string
}

export const createSupplyOrder = (data: SupplyOrderDTO): Promise<ApiResponse<void>> => {
  return post('/api/orders', data)
}

export const listSupplyOrders = (): Promise<ApiResponse<SupplyOrder[]>> => {
  return get('/api/orders')
}

export const createLogisticsData = (data: LogisticsDataDTO): Promise<ApiResponse<void>> => {
  return post('/api/logistics', data)
}

export const listLogisticsData = (): Promise<ApiResponse<LogisticsData[]>> => {
  return get('/api/logistics')
}

export const deleteLogisticsData = (id: string): Promise<ApiResponse<void>> => {
  return del(`/api/logistics/${id}`)
}
