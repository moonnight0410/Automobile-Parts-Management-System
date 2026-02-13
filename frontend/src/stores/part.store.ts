/**
 * 零部件状态管理
 * 管理零部件列表、当前零部件详情等状态
 */

import { defineStore } from 'pinia'
import { ref } from 'vue'
import { 
  createPart, 
  queryPart, 
  queryPartLifecycle, 
  queryPartByBatchNo, 
  queryPartByVIN,
  updatePartStatus 
} from '../services/part.service'
import type { Part, PartLifecycle, CreatePartRequest, UpdatePartStatusRequest } from '../types'

export const usePartStore = defineStore('part', () => {
  // ==================== 状态定义 ====================
  
  // 零部件列表
  const partList = ref<Part[]>([])
  
  // 当前零部件详情
  const currentPart = ref<Part | null>(null)
  
  // 当前零部件生命周期
  const currentLifecycle = ref<PartLifecycle | null>(null)
  
  // 加载状态
  const loading = ref(false)
  
  // 错误信息
  const error = ref<string | null>(null)
  
  // ==================== Actions ====================
  
  /**
   * 创建零部件
   * @param data 零部件数据
   */
  async function create(data: CreatePartRequest) {
    loading.value = true
    error.value = null
    
    try {
      const response = await createPart(data)
      
      if (!response.success) {
        throw new Error(response.message || '创建零部件失败')
      }
      
      return response
    } catch (err: any) {
      error.value = err.message || '创建零部件失败'
      throw err
    } finally {
      loading.value = false
    }
  }
  
  /**
   * 查询零部件详情
   * @param partID 零部件ID
   */
  async function fetchPart(partID: string) {
    loading.value = true
    error.value = null
    
    try {
      const response = await queryPart(partID)
      
      if (response.success && response.data) {
        currentPart.value = response.data
        return response.data
      } else {
        throw new Error(response.message || '查询零部件失败')
      }
    } catch (err: any) {
      error.value = err.message || '查询零部件失败'
      throw err
    } finally {
      loading.value = false
    }
  }
  
  /**
   * 查询零部件生命周期
   * @param partID 零部件ID
   */
  async function fetchLifecycle(partID: string) {
    loading.value = true
    error.value = null
    
    try {
      const response = await queryPartLifecycle(partID)
      
      if (response.success && response.data) {
        currentLifecycle.value = response.data
        return response.data
      } else {
        throw new Error(response.message || '查询生命周期失败')
      }
    } catch (err: any) {
      error.value = err.message || '查询生命周期失败'
      throw err
    } finally {
      loading.value = false
    }
  }
  
  /**
   * 按批次号查询零部件
   * @param batchNo 批次号
   */
  async function fetchByBatchNo(batchNo: string) {
    loading.value = true
    error.value = null
    
    try {
      const response = await queryPartByBatchNo(batchNo)
      
      if (response.success && response.data) {
        partList.value = response.data
        return response.data
      } else {
        throw new Error(response.message || '查询零部件失败')
      }
    } catch (err: any) {
      error.value = err.message || '查询零部件失败'
      throw err
    } finally {
      loading.value = false
    }
  }
  
  /**
   * 按VIN码查询零部件
   * @param vin VIN码
   */
  async function fetchByVIN(vin: string) {
    loading.value = true
    error.value = null
    
    try {
      const response = await queryPartByVIN(vin)
      
      if (response.success && response.data) {
        partList.value = response.data
        return response.data
      } else {
        throw new Error(response.message || '查询零部件失败')
      }
    } catch (err: any) {
      error.value = err.message || '查询零部件失败'
      throw err
    } finally {
      loading.value = false
    }
  }
  
  /**
   * 更新零部件状态
   * @param data 状态更新数据
   */
  async function updateStatus(data: UpdatePartStatusRequest) {
    loading.value = true
    error.value = null
    
    try {
      const response = await updatePartStatus(data)
      
      if (!response.success) {
        throw new Error(response.message || '更新状态失败')
      }
      
      // 更新当前零部件状态
      if (currentPart.value && currentPart.value.partID === data.partID) {
        currentPart.value.status = data.status
      }
      
      return response
    } catch (err: any) {
      error.value = err.message || '更新状态失败'
      throw err
    } finally {
      loading.value = false
    }
  }
  
  /**
   * 清除当前零部件
   */
  function clearCurrentPart() {
    currentPart.value = null
    currentLifecycle.value = null
  }
  
  /**
   * 清除错误
   */
  function clearError() {
    error.value = null
  }
  
  // ==================== 导出 ====================
  
  return {
    // 状态
    partList,
    currentPart,
    currentLifecycle,
    loading,
    error,
    
    // Actions
    create,
    fetchPart,
    fetchLifecycle,
    fetchByBatchNo,
    fetchByVIN,
    updateStatus,
    clearCurrentPart,
    clearError
  }
})
