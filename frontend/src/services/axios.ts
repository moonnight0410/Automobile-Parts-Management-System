/**
 * Axios HTTP客户端配置
 * 配置请求/响应拦截器，处理认证和错误
 */

import axios from 'axios'
import { message } from 'ant-design-vue'
import { API_BASE_URL, STORAGE_KEYS } from '../constants'
import type { ApiResponse } from '../types'

// 创建Axios实例
const axiosInstance = axios.create({
  baseURL: API_BASE_URL,
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
axiosInstance.interceptors.request.use(
  config => {
    // 从localStorage获取token
    const token = localStorage.getItem(STORAGE_KEYS.TOKEN)
    
    // 如果token存在，添加到请求头
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`
    }
    
    // 打印请求信息（开发环境）
    if (import.meta.env.DEV) {
      console.log(`[API Request] ${config.method?.toUpperCase()} ${config.url}`, config.data)
    }
    
    return config
  },
  error => {
    console.error('[API Request Error]', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
axiosInstance.interceptors.response.use(
  response => {
    // 打印响应信息（开发环境）
    if (import.meta.env.DEV) {
      console.log(`[API Response] ${response.config.method?.toUpperCase()} ${response.config.url}`, response.data)
    }
    
    // 检查业务状态
    if (response.data && response.data.code !== 0 && response.data.code !== undefined) {
      message.error(response.data.message || '请求失败')
      return Promise.reject(new Error(response.data.message || '请求失败'))
    }
    
    return response
  },
  error => {
    // 处理HTTP错误
    if (error.response) {
      const { status, data } = error.response
      
      switch (status) {
        case 401:
          // 未授权，清除token并跳转到登录页
          localStorage.removeItem(STORAGE_KEYS.TOKEN)
          localStorage.removeItem(STORAGE_KEYS.USER)
          message.error('登录已过期，请重新登录')
          window.location.href = '/login'
          break
        case 403:
          message.error('没有权限访问该资源')
          break
        case 404:
          message.error('请求的资源不存在')
          break
        case 500:
          message.error(data?.message || '服务器内部错误')
          break
        default:
          message.error(data?.message || `请求失败 (${status})`)
      }
    } else if (error.request) {
      // 请求已发出但没有收到响应
      message.error('网络错误，请检查网络连接')
    } else {
      // 其他错误
      message.error(error.message || '请求失败')
    }
    
    console.error('[API Response Error]', error)
    return Promise.reject(error)
  }
)

// 封装GET请求
export const get = <T = any>(url: string, config?: any): Promise<any> => {
  return axiosInstance.get(url, config).then(res => res.data)
}

export const post = <T = any>(url: string, data?: any, config?: any): Promise<any> => {
  return axiosInstance.post(url, data, config).then(res => res.data)
}

export const put = <T = any>(url: string, data?: any, config?: any): Promise<any> => {
  return axiosInstance.put(url, data, config).then(res => res.data)
}

export const del = <T = any>(url: string, config?: any): Promise<any> => {
  return axiosInstance.delete(url, config).then(res => res.data)
}

export default axiosInstance
