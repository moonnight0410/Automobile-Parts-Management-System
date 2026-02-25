import { get, post } from './axios'
import type { ApiResponse } from '../types'

export interface SupplyOrderItem {
  partID: string
  partName: string
  quantity: number
  unitPrice: number
}

export interface SupplyOrder {
  orderID: string
  buyer: string
  seller: string
  orderDate: string
  deliveryDate: string
  status: string
  totalAmount: number
  items: SupplyOrderItem[]
  remarks?: string
}

export interface SupplyOrderDTO {
  orderID: string
  buyer: string
  seller: string
  orderDate: string
  deliveryDate: string
  status: string
  totalAmount: number
  items: SupplyOrderItem[]
  remarks?: string
}

export interface LogisticsData {
  logisticsID: string
  orderID: string
  carrier: string
  trackingNo: string
  departureLocation: string
  arrivalLocation: string
  departureTime: string
  estimatedArrivalTime: string
  actualArrivalTime?: string
  status: string
  currentLocation?: string
}

export interface LogisticsDataDTO {
  logisticsID: string
  orderID: string
  carrier: string
  trackingNo: string
  departureLocation: string
  arrivalLocation: string
  departureTime: string
  estimatedArrivalTime: string
  actualArrivalTime?: string
  status: string
  currentLocation?: string
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
