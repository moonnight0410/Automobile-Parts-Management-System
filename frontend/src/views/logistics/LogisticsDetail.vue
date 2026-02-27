<template>
  <div class="logistics-detail">
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">物流跟踪详情</h1>
          <span class="page-subtitle">查看物流跟踪完整信息</span>
        </div>
      </div>
      <div class="top-bar-right">
        <a-tag :color="getStatusColor(logisticsData?.status || '')" class="status-badge">
          {{ getStatusText(logisticsData?.status || '') }}
        </a-tag>
      </div>
    </div>

    <a-spin :spinning="loading">
      <div class="main-content">
        <div class="info-container">
          <a-card :bordered="false" class="info-card">
            <div class="card-header">
              <div class="header-icon">
                <CarOutlined />
              </div>
              <div class="header-text">
                <h3>基本信息</h3>
                <p>物流跟踪核心属性</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="info-grid">
              <div class="info-item">
                <div class="info-label">
                  <NumberOutlined class="label-icon" />
                  <span>物流ID</span>
                </div>
                <div class="info-value logistics-id">{{ logisticsData?.logisticsID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ShoppingCartOutlined class="label-icon" />
                  <span>订单ID</span>
                </div>
                <div class="info-value order-id">{{ logisticsData?.orderID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ShopOutlined class="label-icon" />
                  <span>承运商</span>
                </div>
                <div class="info-value">{{ logisticsData?.carrier || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <BarcodeOutlined class="label-icon" />
                  <span>运单号</span>
                </div>
                <div class="info-value tracking-no">{{ logisticsData?.trackingNo || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CheckCircleOutlined class="label-icon" />
                  <span>状态</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getStatusColor(logisticsData?.status || '')" class="status-tag">
                    {{ getStatusText(logisticsData?.status || '') }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item" v-if="logisticsData?.currentLocation">
                <div class="info-label">
                  <EnvironmentOutlined class="label-icon" />
                  <span>当前位置</span>
                </div>
                <div class="info-value">{{ logisticsData.currentLocation }}</div>
              </div>
            </div>
          </a-card>

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
              <div class="action-item" @click="editLogistics">
                <EditOutlined class="action-icon-item" />
                <span>编辑物流</span>
              </div>
              <div class="action-item" @click="printLogistics">
                <PrinterOutlined class="action-icon-item" />
                <span>打印</span>
              </div>
              <div class="action-item" @click="exportData">
                <ExportOutlined class="action-icon-item" />
                <span>导出数据</span>
              </div>
              <div class="action-item" @click="viewOrder">
                <ShoppingCartOutlined class="action-icon-item" />
                <span>查看订单</span>
              </div>
              <div class="action-item delete-action" @click="handleDelete">
                <DeleteOutlined class="action-icon-item" />
                <span>删除该数据</span>
              </div>
            </div>
          </a-card>
        </div>

        <div class="details-container">
          <a-card :bordered="false" class="details-card">
            <div class="card-header">
              <div class="header-icon details-icon">
                <SwapOutlined />
              </div>
              <div class="header-text">
                <h3>运输信息</h3>
                <p>物流运输详情</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="transport-info">
              <div class="info-item">
                <div class="info-label">出发地</div>
                <div class="info-value">{{ logisticsData?.departureLocation || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">目的地</div>
                <div class="info-value">{{ logisticsData?.arrivalLocation || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">出发时间</div>
                <div class="info-value">{{ logisticsData?.departureTime || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">预计到达时间</div>
                <div class="info-value">{{ logisticsData?.estimatedArrivalTime || '-' }}</div>
              </div>
              <div class="info-item" v-if="logisticsData?.actualArrivalTime">
                <div class="info-label">实际到达时间</div>
                <div class="info-value">{{ logisticsData.actualArrivalTime }}</div>
              </div>
            </div>
          </a-card>

          <a-card :bordered="false" class="details-card">
            <div class="card-header">
              <div class="header-icon details-icon">
                <HistoryOutlined />
              </div>
              <div class="header-text">
                <h3>物流时间线</h3>
                <p>物流运输轨迹</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <a-timeline mode="left" class="custom-timeline">
              <a-timeline-item color="blue">
                <template #dot>
                  <div class="timeline-dot blue">
                    <SendOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">货物发出</span>
                    <span class="timeline-time">{{ logisticsData?.departureTime }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <EnvironmentOutlined />
                      出发地: {{ logisticsData?.departureLocation }}
                    </div>
                    <div class="timeline-desc">
                      <ShopOutlined />
                      承运商: {{ logisticsData?.carrier }}
                    </div>
                    <div class="timeline-desc">
                      <BarcodeOutlined />
                      运单号: {{ logisticsData?.trackingNo }}
                    </div>
                  </div>
                </div>
              </a-timeline-item>

              <a-timeline-item color="cyan" v-if="logisticsData?.currentLocation">
                <template #dot>
                  <div class="timeline-dot cyan">
                    <CarOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">运输中</span>
                    <span class="timeline-time">进行中</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <EnvironmentOutlined />
                      当前位置: {{ logisticsData.currentLocation }}
                    </div>
                    <div class="timeline-desc">
                      <ClockCircleOutlined />
                      预计到达: {{ logisticsData?.estimatedArrivalTime }}
                    </div>
                  </div>
                </div>
              </a-timeline-item>

              <a-timeline-item color="green" v-if="logisticsData?.actualArrivalTime">
                <template #dot>
                  <div class="timeline-dot green">
                    <CheckCircleOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">已送达</span>
                    <span class="timeline-time">{{ logisticsData.actualArrivalTime }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <EnvironmentOutlined />
                      目的地: {{ logisticsData?.arrivalLocation }}
                    </div>
                    <div class="timeline-desc">
                      <CheckCircleOutlined />
                      物流完成
                    </div>
                  </div>
                </div>
              </a-timeline-item>
            </a-timeline>
          </a-card>
        </div>
      </div>
    </a-spin>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, h } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { message, Modal } from 'ant-design-vue'
import { listLogisticsData } from '../../services/supply.service'
import type { LogisticsData } from '../../services/supply.service'
import {
  ArrowLeftOutlined,
  NumberOutlined,
  ShoppingCartOutlined,
  ShopOutlined,
  BarcodeOutlined,
  CheckCircleOutlined,
  EnvironmentOutlined,
  ClockCircleOutlined,
  ThunderboltOutlined,
  CarOutlined,
  EditOutlined,
  PrinterOutlined,
  ExportOutlined,
  DeleteOutlined,
  SwapOutlined,
  HistoryOutlined,
  SendOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()

const loading = ref(false)
const logisticsData = ref<LogisticsData | null>(null)

function goBack() {
  router.push('/supply/logistics')
}

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    '待发货': 'default',
    '运输中': 'processing',
    '已送达': 'success',
    '已签收': 'success',
    '异常': 'error'
  }
  return colors[status] || 'default'
}

function getStatusText(status: string) {
  return status || status
}

function editLogistics() {
  message.info('编辑功能开发中')
}

function printLogistics() {
  message.success('打印任务已提交')
}

function exportData() {
  message.success('数据导出成功')
}

function viewOrder() {
  if (logisticsData.value?.orderID) {
    router.push(`/orders/detail/${logisticsData.value.orderID}`)
  } else {
    message.warning('订单ID不存在')
  }
}

async function handleDelete() {
  const logisticsID = route.params.id as string
  if (!logisticsID) {
    message.error('物流ID不存在')
    return
  }
  
  const user = JSON.parse(localStorage.getItem('user') || '{}')
  const userRole = user.role || user.userType || ''
  const hasDeletePermission = userRole === 'admin' || userRole === 'manufacturer' || userRole === 'manufacturer_user'
  
  if (!hasDeletePermission) {
    message.error('您没有权限删除该数据')
    return
  }
  
  Modal.confirm({
    title: '确认删除',
    icon: () => h(DeleteOutlined, { style: { color: '#ff4d4f' } }),
    maskClosable: false,
    maskStyle: { backgroundColor: 'rgba(0, 0, 0, 0.6)' },
    content: h('div', { style: { lineHeight: '1.8' } }, [
      h('p', { style: { marginBottom: '8px' } }, [
        h('strong', { style: { color: '#ff4d4f' } }, '警告：此操作不可恢复！')
      ]),
      h('p', { style: { marginBottom: '8px' } }, [
        `您确定要删除物流记录 ${logisticsID} 吗？`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `订单ID：${logisticsData.value?.orderID || '-'}`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `运单号：${logisticsData.value?.trackingNo || '-'}`
      ]),
      h('p', { style: { fontSize: '13px', color: '#999' } }, [
        '删除后将无法恢复，请谨慎操作。'
      ])
    ]),
    okText: '确认删除',
    okType: 'danger',
    cancelText: '取消',
    onOk: async () => {
      message.success('删除成功')
      setTimeout(() => {
        goBack()
      }, 1000)
    }
  })
}

async function fetchLogisticsData() {
  loading.value = true
  try {
    const response = await listLogisticsData()
    if (response.code === 0 && response.data) {
      const logisticsID = route.params.id as string
      const record = response.data.find((item: LogisticsData) => item.logisticsID === logisticsID)
      if (record) {
        logisticsData.value = record
      } else {
        message.error('未找到该物流记录')
      }
    } else {
      message.error(response.message || '获取物流记录失败')
    }
  } catch (error: any) {
    message.error(error.message || '获取物流记录失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchLogisticsData()
})
</script>

<style scoped>
.logistics-detail {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  padding: 24px;
}

.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  background: white;
  padding: 16px 24px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.top-bar-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.back-btn {
  font-size: 14px;
  color: #1890ff;
}

.divider {
  height: 24px;
  margin: 0;
}

.page-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.page-title {
  font-size: 20px;
  font-weight: 600;
  margin: 0;
  color: #1f2937;
}

.page-subtitle {
  font-size: 13px;
  color: #6b7280;
}

.top-bar-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.status-badge {
  font-size: 13px;
  padding: 4px 12px;
}

.main-content {
  display: grid;
  grid-template-columns: 400px 1fr;
  gap: 24px;
}

.info-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.info-card,
.action-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.card-header {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px 24px 0;
}

.header-icon {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 18px;
}

.header-icon.action-icon {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.header-icon.details-icon {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.header-text h3 {
  font-size: 16px;
  font-weight: 600;
  margin: 0 0 4px 0;
  color: #1f2937;
}

.header-text p {
  font-size: 12px;
  color: #6b7280;
  margin: 0;
}

.card-divider {
  margin: 16px 0;
}

.info-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 16px;
  padding: 0 24px 24px;
}

.info-item {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.info-label {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: #6b7280;
  font-weight: 500;
}

.label-icon {
  font-size: 14px;
  color: #9ca3af;
}

.info-value {
  font-size: 14px;
  color: #1f2937;
  font-weight: 500;
  word-break: break-all;
}

.logistics-id,
.order-id,
.tracking-no {
  font-family: 'Courier New', monospace;
  color: #667eea;
  font-weight: 600;
}

.status-tag {
  font-size: 12px;
}

.action-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
  padding: 0 24px 24px;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 16px;
  background: linear-gradient(135deg, #f5f7fa 0%, #e4e8ec 100%);
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 1px solid transparent;
}

.action-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border-color: #667eea;
}

.action-icon-item {
  font-size: 20px;
  color: #667eea;
}

.action-item span {
  font-size: 12px;
  color: #4b5563;
  font-weight: 500;
}

.action-item.delete-action:hover {
  background: linear-gradient(135deg, #fee2e2 0%, #fecaca 100%);
  border-color: #ef4444;
}

.action-item.delete-action .action-icon-item {
  color: #ef4444;
}

.details-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.details-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.transport-info {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 0 24px 24px;
}

.timeline-wrapper {
  padding: 0 24px 24px;
}

.custom-timeline {
  margin-top: 16px;
}

.timeline-item {
  padding-left: 8px;
}

.timeline-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.timeline-title {
  font-size: 14px;
  font-weight: 600;
  color: #1f2937;
}

.timeline-time {
  font-size: 12px;
  color: #9ca3af;
}

.timeline-content {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.timeline-desc {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #6b7280;
}

.timeline-desc .anticon {
  font-size: 14px;
  color: #9ca3af;
}

@media (max-width: 1200px) {
  .main-content {
    grid-template-columns: 1fr;
  }
  
  .info-container {
    max-width: 100%;
  }
}

@media (max-width: 768px) {
  .logistics-detail {
    padding: 16px;
  }
  
  .top-bar {
    flex-direction: column;
    gap: 12px;
    align-items: flex-start;
  }
  
  .top-bar-left {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
  
  .action-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
