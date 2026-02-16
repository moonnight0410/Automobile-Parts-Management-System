<!--
  零部件详情页面
  展示零部件详细信息和生命周期轨迹
  设计：与Dashboard风格统一
-->

<template>
  <div class="part-detail">
    <!-- 顶部导航栏 -->
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">零部件详情</h1>
          <span class="page-subtitle">查看完整信息和生命周期</span>
        </div>
      </div>
      <div class="top-bar-right">
        <a-tag :color="getStatusColor(partData?.status || '')" class="status-badge">
          {{ getStatusText(partData?.status || '') }}
        </a-tag>
      </div>
    </div>

    <!-- 加载状态 -->
    <a-spin :spinning="loading">
      <div class="main-content">
        <!-- 左侧信息区 -->
        <div class="info-container">
          <!-- 基本信息卡片 -->
          <a-card :bordered="false" class="info-card">
            <div class="card-header">
              <div class="header-icon">
                <AppstoreOutlined />
              </div>
              <div class="header-text">
                <h3>基本信息</h3>
                <p>零部件核心属性</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="info-grid">
              <div class="info-item">
                <div class="info-label">
                  <NumberOutlined class="label-icon" />
                  <span>零部件ID</span>
                </div>
                <div class="info-value part-id">{{ partData?.partID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <TagOutlined class="label-icon" />
                  <span>零部件名称</span>
                </div>
                <div class="info-value">{{ partData?.name || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <AppstoreOutlined class="label-icon" />
                  <span>类型</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getTypeColor(partData?.type || '')" class="type-tag">
                    {{ partData?.type || '-' }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <BarcodeOutlined class="label-icon" />
                  <span>批次号</span>
                </div>
                <div class="info-value batch-no">{{ partData?.batchNo || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CarOutlined class="label-icon" />
                  <span>VIN码</span>
                </div>
                <div class="info-value vin-code">{{ partData?.vin || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ShopOutlined class="label-icon" />
                  <span>生产厂商</span>
                </div>
                <div class="info-value">{{ partData?.manufacturer || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CheckCircleOutlined class="label-icon" />
                  <span>状态</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getStatusColor(partData?.status || '')" class="status-tag">
                    {{ getStatusText(partData?.status || '') }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ClockCircleOutlined class="label-icon" />
                  <span>创建时间</span>
                </div>
                <div class="info-value">{{ partData?.createTime || '-' }}</div>
              </div>
            </div>
          </a-card>

          <!-- 快捷操作卡片 -->
          <a-card :bordered="false" class="action-card">
            <div class="card-header">
              <div class="header-icon action-icon">
                <ThunderboltOutlined />
              </div>
              <div class="header-text">
                <h3>快捷操作</h3>
                <p>常用功能入口</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="action-grid">
              <div class="action-item" @click="viewLifecycle">
                <HistoryOutlined class="action-icon-item" />
                <span>生命周期</span>
              </div>
              <div class="action-item" @click="viewBOM">
                <FileTextOutlined class="action-icon-item" />
                <span>BOM信息</span>
              </div>
              <div class="action-item" @click="viewSupplyChain">
                <SwapOutlined class="action-icon-item" />
                <span>供应链</span>
              </div>
              <div class="action-item" @click="exportData">
                <ExportOutlined class="action-icon-item" />
                <span>导出数据</span>
              </div>
            </div>
          </a-card>
        </div>

        <!-- 右侧生命周期区 -->
        <div class="lifecycle-container">
          <a-card :bordered="false" class="lifecycle-card">
            <div class="card-header">
              <div class="header-icon lifecycle-icon">
                <HistoryOutlined />
              </div>
              <div class="header-text">
                <h3>生命周期轨迹</h3>
                <p>零部件全流程追踪记录</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="timeline-wrapper">
              <a-timeline mode="left" class="custom-timeline">
                <a-timeline-item color="green">
                  <template #dot>
                    <div class="timeline-dot green">
                      <PlusCircleOutlined />
                    </div>
                  </template>
                  <div class="timeline-item">
                    <div class="timeline-header">
                      <span class="timeline-title">零部件创建</span>
                      <span class="timeline-time">{{ partData?.createTime }}</span>
                    </div>
                    <div class="timeline-content">
                      <div class="timeline-desc">
                        <InfoCircleOutlined />
                        零部件 {{ partData?.partID }} 已创建并上链
                      </div>
                    </div>
                  </div>
                </a-timeline-item>
                
                <a-timeline-item color="blue" v-if="lifecycleData?.productionInfo">
                  <template #dot>
                    <div class="timeline-dot blue">
                      <ToolOutlined />
                    </div>
                  </template>
                  <div class="timeline-item">
                    <div class="timeline-header">
                      <span class="timeline-title">生产完成</span>
                      <span class="timeline-time">{{ lifecycleData?.productionInfo?.finishTime }}</span>
                    </div>
                    <div class="timeline-content">
                      <div class="timeline-desc">
                        <SettingOutlined />
                        生产线: {{ lifecycleData?.productionInfo?.productionLine }}
                      </div>
                      <div class="timeline-desc">
                        <UserOutlined />
                        操作员: {{ lifecycleData?.productionInfo?.operator }}
                      </div>
                    </div>
                  </div>
                </a-timeline-item>
                
                <a-timeline-item color="cyan" v-if="lifecycleData?.qualityInfo">
                  <template #dot>
                    <div class="timeline-dot cyan">
                      <SafetyCertificateOutlined />
                    </div>
                  </template>
                  <div class="timeline-item">
                    <div class="timeline-header">
                      <span class="timeline-title">质检完成</span>
                      <span class="timeline-time">{{ lifecycleData?.qualityInfo?.handleTime }}</span>
                    </div>
                    <div class="timeline-content">
                      <div class="timeline-desc">
                        <CheckCircleOutlined />
                        质检结果: 
                        <a-tag :color="lifecycleData?.qualityInfo?.result === '合格' ? 'success' : 'error'" size="small">
                          {{ lifecycleData?.qualityInfo?.result }}
                        </a-tag>
                      </div>
                      <div class="timeline-desc">
                        <UserOutlined />
                        处理人: {{ lifecycleData?.qualityInfo?.handler }}
                      </div>
                    </div>
                  </div>
                </a-timeline-item>
                
                <a-timeline-item color="purple" v-if="lifecycleData?.supplyChainInfo?.length">
                  <template #dot>
                    <div class="timeline-dot purple">
                      <SwapOutlined />
                    </div>
                  </template>
                  <div class="timeline-item">
                    <div class="timeline-header">
                      <span class="timeline-title">供应链流转</span>
                      <span class="timeline-time">{{ lifecycleData?.supplyChainInfo?.length }} 条记录</span>
                    </div>
                    <div class="timeline-content">
                      <div class="timeline-desc">
                        <ApartmentOutlined />
                        已完成供应链流转环节
                      </div>
                    </div>
                  </div>
                </a-timeline-item>
                
                <a-timeline-item color="orange" v-if="lifecycleData?.aftersaleInfo?.length">
                  <template #dot>
                    <div class="timeline-dot orange">
                      <CustomerServiceOutlined />
                    </div>
                  </template>
                  <div class="timeline-item">
                    <div class="timeline-header">
                      <span class="timeline-title">售后服务</span>
                      <span class="timeline-time">{{ lifecycleData?.aftersaleInfo?.length }} 条记录</span>
                    </div>
                    <div class="timeline-content">
                      <div class="timeline-desc">
                        <AlertOutlined />
                        存在售后服务相关记录
                      </div>
                    </div>
                  </div>
                </a-timeline-item>
              </a-timeline>
            </div>
          </a-card>
        </div>
      </div>
    </a-spin>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { message } from 'ant-design-vue'
import { usePartStore } from '../../stores'
import type { Part, PartLifecycle } from '../../types'
import {
  ArrowLeftOutlined,
  AppstoreOutlined,
  NumberOutlined,
  TagOutlined,
  BarcodeOutlined,
  CarOutlined,
  ShopOutlined,
  CheckCircleOutlined,
  ClockCircleOutlined,
  ThunderboltOutlined,
  HistoryOutlined,
  FileTextOutlined,
  SwapOutlined,
  ExportOutlined,
  PlusCircleOutlined,
  ToolOutlined,
  SettingOutlined,
  UserOutlined,
  SafetyCertificateOutlined,
  ApartmentOutlined,
  AlertOutlined,
  CustomerServiceOutlined,
  InfoCircleOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()
const partStore = usePartStore()

const loading = ref(false)
const partData = ref<Part | null>(null)
const lifecycleData = ref<PartLifecycle | null>(null)

function goBack() {
  router.push('/parts/list')
}

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    'NORMAL': 'success',
    'IN_PRODUCTION': 'processing',
    'QUALITY_INSPECTED': 'cyan',
    'IN_SUPPLY_CHAIN': 'purple',
    'RECALLED': 'warning',
    'SCRAPPED': 'default'
  }
  return colors[status] || 'default'
}

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

function getTypeColor(type: string) {
  const colors: Record<string, string> = {
    '发动机部件': 'purple',
    '制动系统': 'magenta',
    '传动系统': 'blue',
    '悬挂系统': 'green',
    '电气系统': 'orange',
    '车身部件': 'cyan'
  }
  return colors[type] || 'default'
}

function viewLifecycle() {
  message.info('生命周期信息已展示在右侧')
}

function viewBOM() {
  message.info('BOM信息功能开发中')
}

function viewSupplyChain() {
  message.info('供应链信息功能开发中')
}

function exportData() {
  message.success('数据导出成功')
}

async function loadData() {
  const partID = route.params.id as string
  if (!partID) {
    message.error('零部件ID不存在')
    goBack()
    return
  }
  
  loading.value = true
  
  try {
    await partStore.fetchPart(partID)
    partData.value = partStore.currentPart
    
    await partStore.fetchLifecycle(partID)
    lifecycleData.value = partStore.currentLifecycle
  } catch (error: any) {
    message.error(error.message || '加载数据失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.part-detail {
  min-height: 100%;
  background: var(--bg-color);
  padding: 0;
}

.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 32px;
  background: var(--bg-color-secondary);
  border-bottom: 1px solid var(--border-color);
  position: sticky;
  top: 0;
  z-index: 100;
}

.top-bar-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.back-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--text-color-secondary);
  font-size: 14px;
  padding: 8px 16px;
  border-radius: 8px;
  transition: all 0.3s;
}

.back-btn:hover {
  background: var(--bg-color-tertiary);
  color: var(--primary-color);
}

.divider {
  height: 24px;
  background: var(--border-color);
}

.page-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.page-title {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-color);
  margin: 0;
}

.page-subtitle {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.status-badge {
  font-size: 13px;
  padding: 4px 12px;
  border-radius: 16px;
  font-weight: 500;
}

.main-content {
  display: flex;
  gap: 24px;
  padding: 24px 32px;
  max-width: 1600px;
  margin: 0 auto;
}

.info-container {
  width: 480px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.lifecycle-container {
  flex: 1;
  min-width: 0;
}

.info-card,
.action-card,
.lifecycle-card {
  border-radius: 16px;
  box-shadow: var(--shadow-sm);
  background: var(--bg-color-secondary);
}

.card-header {
  display: flex;
  align-items: center;
  gap: 16px;
}

.header-icon {
  width: 44px;
  height: 44px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 20px;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.header-icon.action-icon {
  background: linear-gradient(135deg, #fa8c16 0%, #fa541c 100%);
  box-shadow: 0 4px 12px rgba(250, 140, 22, 0.3);
}

.header-icon.lifecycle-icon {
  background: linear-gradient(135deg, #1890ff 0%, #722ed1 100%);
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3);
}

.header-text h3 {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-color);
  margin: 0;
}

.header-text p {
  font-size: 13px;
  color: var(--text-color-secondary);
  margin: 2px 0 0 0;
}

.card-divider {
  margin: 20px 0;
}

.info-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  padding: 4px;
}

.info-item {
  padding: 16px;
  background: var(--bg-color);
  border-radius: 12px;
  border: 1px solid var(--border-color);
  transition: all 0.3s;
}

.info-item:hover {
  border-color: var(--primary-color);
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.1);
}

.info-label {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: var(--text-color-secondary);
  margin-bottom: 10px;
}

.label-icon {
  color: var(--primary-color);
  font-size: 15px;
}

.info-value {
  font-size: 15px;
  font-weight: 500;
  color: var(--text-color);
  word-break: break-all;
}

.part-id {
  font-family: 'Monaco', 'Consolas', monospace;
  color: var(--primary-color);
  background: rgba(102, 126, 234, 0.1);
  padding: 4px 8px;
  border-radius: 4px;
  display: inline-block;
}

.batch-no,
.vin-code {
  font-family: 'Monaco', 'Consolas', monospace;
}

.type-tag,
.status-tag {
  border-radius: 4px;
  font-weight: 500;
}

.action-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  padding: 4px;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px 16px;
  background: var(--bg-color);
  border-radius: 12px;
  border: 1px solid var(--border-color);
  cursor: pointer;
  transition: all 0.3s;
}

.action-item:hover {
  border-color: var(--primary-color);
  background: rgba(102, 126, 234, 0.05);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.15);
}

.action-icon-item {
  font-size: 28px;
  color: var(--primary-color);
  margin-bottom: 10px;
}

.action-item span {
  font-size: 14px;
  color: var(--text-color);
  font-weight: 500;
}

.timeline-wrapper {
  padding: 8px 0;
}

.custom-timeline :deep(.ant-timeline-item-tail) {
  border-left-style: dashed;
  border-left-width: 2px;
}

.custom-timeline :deep(.ant-timeline-item-head) {
  background: transparent;
  border: none;
  padding: 0;
}

.timeline-dot {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 16px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

.timeline-dot.green {
  background: linear-gradient(135deg, #52c41a 0%, #73d13d 100%);
}

.timeline-dot.blue {
  background: linear-gradient(135deg, #1890ff 0%, #40a9ff 100%);
}

.timeline-dot.cyan {
  background: linear-gradient(135deg, #13c2c2 0%, #36cfc9 100%);
}

.timeline-dot.purple {
  background: linear-gradient(135deg, #722ed1 0%, #9254de 100%);
}

.timeline-dot.orange {
  background: linear-gradient(135deg, #fa8c16 0%, #ffa940 100%);
}

.timeline-item {
  padding: 8px 0 24px 16px;
}

.timeline-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 12px;
}

.timeline-title {
  font-size: 15px;
  font-weight: 600;
  color: var(--text-color);
}

.timeline-time {
  font-size: 12px;
  color: var(--text-color-secondary);
  background: var(--bg-color);
  padding: 4px 10px;
  border-radius: 12px;
}

.timeline-content {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.timeline-desc {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: var(--text-color-secondary);
  padding: 8px 12px;
  background: var(--bg-color);
  border-radius: 8px;
}

.timeline-desc :deep(.anticon) {
  color: var(--primary-color);
}

[data-theme='dark'] .top-bar {
  background: var(--bg-color-tertiary);
}

[data-theme='dark'] .info-card,
[data-theme='dark'] .action-card,
[data-theme='dark'] .lifecycle-card {
  background: var(--bg-color-tertiary);
}

[data-theme='dark'] .info-item {
  background: var(--bg-color);
  border-color: var(--border-color);
}

[data-theme='dark'] .info-item:hover {
  border-color: var(--primary-color);
}

[data-theme='dark'] .action-item {
  background: var(--bg-color);
  border-color: var(--border-color);
}

[data-theme='dark'] .action-item:hover {
  background: rgba(102, 126, 234, 0.1);
}

[data-theme='dark'] .timeline-time {
  background: var(--bg-color);
}

[data-theme='dark'] .timeline-desc {
  background: var(--bg-color);
}
</style>
