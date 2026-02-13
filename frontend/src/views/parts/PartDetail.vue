<!--
  零部件详情页面
  展示零部件详细信息和生命周期轨迹
-->

<template>
  <div class="part-detail">
    <!-- 页面标题 -->
    <div class="page-header">
      <h2>零部件详情</h2>
      <a-button @click="goBack">
        <template #icon><ArrowLeftOutlined /></template>
        返回
      </a-button>
    </div>
    
    <!-- 加载状态 -->
    <a-spin :spinning="loading">
      <!-- 基本信息 -->
      <a-card title="基本信息" :bordered="false" class="info-card">
        <a-descriptions :column="2" bordered>
          <a-descriptions-item label="零部件ID">
            {{ partData?.partID }}
          </a-descriptions-item>
          <a-descriptions-item label="零部件名称">
            {{ partData?.name }}
          </a-descriptions-item>
          <a-descriptions-item label="类型">
            {{ partData?.type }}
          </a-descriptions-item>
          <a-descriptions-item label="批次号">
            {{ partData?.batchNo }}
          </a-descriptions-item>
          <a-descriptions-item label="VIN码">
            {{ partData?.vin || '-' }}
          </a-descriptions-item>
          <a-descriptions-item label="生产厂商">
            {{ partData?.manufacturer }}
          </a-descriptions-item>
          <a-descriptions-item label="状态">
            <a-tag :color="getStatusColor(partData?.status || '')">
              {{ getStatusText(partData?.status || '') }}
            </a-tag>
          </a-descriptions-item>
          <a-descriptions-item label="创建时间">
            {{ partData?.createTime }}
          </a-descriptions-item>
        </a-descriptions>
      </a-card>
      
      <!-- 生命周期时间轴 -->
      <a-card title="生命周期轨迹" :bordered="false" class="lifecycle-card">
        <a-timeline mode="left">
          <a-timeline-item color="green">
            <template #dot><ClockCircleOutlined style="font-size: 16px" /></template>
            <div class="timeline-item">
              <div class="timeline-title">零部件创建</div>
              <div class="timeline-time">{{ partData?.createTime }}</div>
              <div class="timeline-desc">零部件 {{ partData?.partID }} 已创建并上链</div>
            </div>
          </a-timeline-item>
          
          <a-timeline-item color="blue" v-if="lifecycleData?.productionInfo">
            <template #dot><ToolOutlined style="font-size: 16px" /></template>
            <div class="timeline-item">
              <div class="timeline-title">生产完成</div>
              <div class="timeline-time">{{ lifecycleData?.productionInfo?.finishTime }}</div>
              <div class="timeline-desc">
                生产线: {{ lifecycleData?.productionInfo?.productionLine }}，
                操作员: {{ lifecycleData?.productionInfo?.operator }}
              </div>
            </div>
          </a-timeline-item>
          
          <a-timeline-item color="cyan" v-if="lifecycleData?.qualityInfo">
            <template #dot><SafetyCertificateOutlined style="font-size: 16px" /></template>
            <div class="timeline-item">
              <div class="timeline-title">质检完成</div>
              <div class="timeline-time">{{ lifecycleData?.qualityInfo?.handleTime }}</div>
              <div class="timeline-desc">
                质检结果: {{ lifecycleData?.qualityInfo?.result }}，
                处理人: {{ lifecycleData?.qualityInfo?.handler }}
              </div>
            </div>
          </a-timeline-item>
          
          <a-timeline-item color="purple" v-if="lifecycleData?.supplyChainInfo?.length">
            <template #dot><SwapOutlined style="font-size: 16px" /></template>
            <div class="timeline-item">
              <div class="timeline-title">供应链流转</div>
              <div class="timeline-desc">
                共 {{ lifecycleData?.supplyChainInfo?.length }} 条流转记录
              </div>
            </div>
          </a-timeline-item>
          
          <a-timeline-item color="orange" v-if="lifecycleData?.aftersaleInfo?.length">
            <template #dot><CustomerServiceOutlined style="font-size: 16px" /></template>
            <div class="timeline-item">
              <div class="timeline-title">售后服务</div>
              <div class="timeline-desc">
                共 {{ lifecycleData?.aftersaleInfo?.length }} 条售后记录
              </div>
            </div>
          </a-timeline-item>
        </a-timeline>
      </a-card>
    </a-spin>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { message } from 'ant-design-vue'
import { usePartStore } from '../../stores'
import type { Part, PartLifecycle } from '../../types'

// ==================== 组合式API ====================

const router = useRouter()
const route = useRoute()
const partStore = usePartStore()

// ==================== 响应式状态 ====================

const loading = ref(false)
const partData = ref<Part | null>(null)
const lifecycleData = ref<PartLifecycle | null>(null)

// ==================== 方法 ====================

/**
 * 返回上一页
 */
function goBack() {
  router.back()
}

/**
 * 获取状态颜色
 */
function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    'NORMAL': 'green',
    'IN_PRODUCTION': 'blue',
    'QUALITY_INSPECTED': 'cyan',
    'IN_SUPPLY_CHAIN': 'purple',
    'RECALLED': 'orange',
    'SCRAPPED': 'default'
  }
  return colors[status] || 'default'
}

/**
 * 获取状态文本
 */
function getStatusText(status: string) {
  const texts: Record<string, string> = {
    'NORMAL': '正常',
    'IN_PRODUCTION': '在产',
    'QUALITY_INSPECTED': '质检完成',
    'IN_SUPPLY_CHAIN': '供应链中',
    'RECALLED': '已召回',
    'SCRAPPED': '已报废'
  }
  return texts[status] || status
}

/**
 * 加载数据
 */
async function loadData() {
  const partID = route.params.id as string
  if (!partID) {
    message.error('零部件ID不存在')
    goBack()
    return
  }
  
  loading.value = true
  
  try {
    // 获取零部件详情
    await partStore.fetchPart(partID)
    partData.value = partStore.currentPart
    
    // 获取生命周期
    await partStore.fetchLifecycle(partID)
    lifecycleData.value = partStore.currentLifecycle
  } catch (error: any) {
    message.error(error.message || '加载数据失败')
  } finally {
    loading.value = false
  }
}

// ==================== 生命周期 ====================

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.part-detail {
  padding: 0;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.page-header h2 {
  font-size: 20px;
  font-weight: 600;
  margin: 0;
}

.info-card,
.lifecycle-card {
  border-radius: 8px;
  margin-bottom: 16px;
}

.timeline-item {
  padding: 4px 0;
}

.timeline-title {
  font-weight: 500;
  font-size: 14px;
  margin-bottom: 4px;
}

.timeline-time {
  font-size: 12px;
  color: var(--text-color-secondary);
  margin-bottom: 4px;
}

.timeline-desc {
  font-size: 13px;
  color: var(--text-color-secondary);
}
</style>
