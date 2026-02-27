<template>
  <div class="production-detail">
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">生产记录详情</h1>
          <span class="page-subtitle">查看生产记录完整信息</span>
        </div>
      </div>
      <div class="top-bar-right">
        <a-tag :color="getQualityStatusColor(productionData?.params?.status || '')" class="status-badge">
          {{ getQualityStatusText(productionData?.params?.status || '') }}
        </a-tag>
      </div>
    </div>

    <a-spin :spinning="loading">
      <div class="main-content">
        <div class="info-container">
          <a-card :bordered="false" class="info-card">
            <div class="card-header">
              <div class="header-icon">
                <ToolOutlined />
              </div>
              <div class="header-text">
                <h3>基本信息</h3>
                <p>生产记录核心属性</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="info-grid">
              <div class="info-item">
                <div class="info-label">
                  <NumberOutlined class="label-icon" />
                  <span>记录ID</span>
                </div>
                <div class="info-value record-id">{{ productionData?.productionID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <AppstoreOutlined class="label-icon" />
                  <span>零部件ID</span>
                </div>
                <div class="info-value part-id">{{ productionData?.partID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <BarcodeOutlined class="label-icon" />
                  <span>批次号</span>
                </div>
                <div class="info-value batch-no">{{ productionData?.batchNo || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <NumberOutlined class="label-icon" />
                  <span>生产数量</span>
                </div>
                <div class="info-value">{{ productionData?.params?.quantity || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <SettingOutlined class="label-icon" />
                  <span>生产线</span>
                </div>
                <div class="info-value">{{ productionData?.productionLine || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <UserOutlined class="label-icon" />
                  <span>操作员</span>
                </div>
                <div class="info-value">{{ productionData?.operator || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ThunderboltOutlined class="label-icon" />
                  <span>操作类型</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getOperationTypeColor('生产')" class="type-tag">
                    生产
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CheckCircleOutlined class="label-icon" />
                  <span>质量状态</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getQualityStatusColor(productionData?.params?.status || '')" class="status-tag">
                    {{ getQualityStatusText(productionData?.params?.status || '') }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ClockCircleOutlined class="label-icon" />
                  <span>操作时间</span>
                </div>
                <div class="info-value">{{ productionData?.finishTime || '-' }}</div>
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
              <div class="action-item" @click="editRecord">
                <EditOutlined class="action-icon-item" />
                <span>编辑记录</span>
              </div>
              <div class="action-item" @click="printRecord">
                <PrinterOutlined class="action-icon-item" />
                <span>打印</span>
              </div>
              <div class="action-item" @click="exportData">
                <ExportOutlined class="action-icon-item" />
                <span>导出数据</span>
              </div>
              <div class="action-item" @click="viewPart">
                <AppstoreOutlined class="action-icon-item" />
                <span>查看零部件</span>
              </div>
              <div class="action-item delete-action" @click="handleDelete" :class="{ 'deleting': deleting }">
                <DeleteOutlined class="action-icon-item" :class="{ 'spinning': deleting }" />
                <span>{{ deleting ? '删除中...' : '删除该数据' }}</span>
              </div>
            </div>
          </a-card>
        </div>

        <div class="details-container">
          <a-card :bordered="false" class="details-card">
            <div class="card-header">
              <div class="header-icon details-icon">
                <FileTextOutlined />
              </div>
              <div class="header-text">
                <h3>操作详情</h3>
                <p>生产操作详细信息</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="operation-details">
              <div class="info-item">
                <div class="info-label">操作结果</div>
                <div class="info-value">
                  <a-tag :color="'success'" class="result-tag">
                    成功
                  </a-tag>
                </div>
              </div>
              <div class="info-item" v-if="productionData?.params?.remark">
                <div class="info-label">备注信息</div>
                <div class="info-value remarks">{{ productionData.params.remark }}</div>
              </div>
            </div>
          </a-card>

          <a-card :bordered="false" class="details-card">
            <div class="card-header">
              <div class="header-icon details-icon">
                <HistoryOutlined />
              </div>
              <div class="header-text">
                <h3>操作时间线</h3>
                <p>生产操作时间记录</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <a-timeline mode="left" class="custom-timeline">
              <a-timeline-item color="blue">
                <template #dot>
                  <div class="timeline-dot blue">
                    <PlayCircleOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">操作开始</span>
                    <span class="timeline-time">{{ productionData?.finishTime }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <UserOutlined />
                      操作员: {{ productionData?.operator }}
                    </div>
                    <div class="timeline-desc">
                      <SettingOutlined />
                      生产线: {{ productionData?.productionLine }}
                    </div>
                  </div>
                </div>
              </a-timeline-item>

              <a-timeline-item color="green">
                <template #dot>
                  <div class="timeline-dot green">
                    <CheckCircleOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">操作完成</span>
                    <span class="timeline-time">{{ productionData?.finishTime }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <CheckCircleOutlined />
                      操作结果: 
                      <a-tag :color="'success'" size="small">
                        成功
                      </a-tag>
                    </div>
                    <div class="timeline-desc">
                      <SafetyCertificateOutlined />
                      质量状态: 
                      <a-tag :color="getQualityStatusColor(productionData?.params?.status || '')" size="small">
                        {{ getQualityStatusText(productionData?.params?.status || '') }}
                      </a-tag>
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
import { listProductionData, deleteProductionData } from '../../services/production.service'
import type { ProductionData } from '../../services/production.service'
import {
  ArrowLeftOutlined,
  NumberOutlined,
  AppstoreOutlined,
  BarcodeOutlined,
  SettingOutlined,
  UserOutlined,
  CheckCircleOutlined,
  ClockCircleOutlined,
  ThunderboltOutlined,
  ToolOutlined,
  EditOutlined,
  PrinterOutlined,
  ExportOutlined,
  DeleteOutlined,
  FileTextOutlined,
  HistoryOutlined,
  PlayCircleOutlined,
  SafetyCertificateOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()

const loading = ref(false)
const deleting = ref(false)
const productionData = ref<ProductionData | null>(null)

function goBack() {
  router.push('/production/data')
}

function getQualityStatusColor(status: string) {
  const colors: Record<string, string> = {
    '合格': 'success',
    '不合格': 'error',
    '待检': 'warning',
    '免检': 'default'
  }
  return colors[status] || 'default'
}

function getQualityStatusText(status: string) {
  return status || status
}

function getOperationTypeColor(type: string) {
  const colors: Record<string, string> = {
    '生产': 'blue',
    '装配': 'cyan',
    '测试': 'purple',
    '包装': 'green'
  }
  return colors[type] || 'default'
}

function editRecord() {
  message.info('编辑功能开发中')
}

function printRecord() {
  message.success('打印任务已提交')
}

function exportData() {
  message.success('数据导出成功')
}

function viewPart() {
  if (productionData.value?.partID) {
    router.push(`/parts/detail/${productionData.value.partID}`)
  } else {
    message.warning('零部件ID不存在')
  }
}

async function handleDelete() {
  const recordID = route.params.id as string
  if (!recordID) {
    message.error('记录ID不存在')
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
        `您确定要删除生产记录 ${recordID} 吗？`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `零部件ID：${productionData.value?.partID || '-'}`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `批次号：${productionData.value?.batchNo || '-'}`
      ]),
      h('p', { style: { fontSize: '13px', color: '#999' } }, [
        '删除后将无法恢复，请谨慎操作。'
      ])
    ]),
    okText: '确认删除',
    okType: 'danger',
    cancelText: '取消',
    okButtonProps: {
      loading: deleting.value,
      disabled: deleting.value
    },
    onOk: async () => {
      deleting.value = true
      try {
        const response = await deleteProductionData(recordID)
        if (response.success) {
          message.success('删除成功')
          goBack()
        } else {
          message.error(response.message || '删除失败')
        }
      } catch (error: any) {
        message.error(error.message || '删除失败，请稍后重试')
      } finally {
        deleting.value = false
      }
    }
  })
}

async function fetchProductionData() {
  loading.value = true
  try {
    const response = await listProductionData()
    if (response.code === 0 && response.data) {
      const recordID = route.params.id as string
      const record = response.data.find((item: ProductionData) => item.productionID === recordID)
      if (record) {
        productionData.value = record
      } else {
        message.error('未找到该生产记录')
      }
    } else {
      message.error(response.message || '获取生产记录失败')
    }
  } catch (error: any) {
    message.error(error.message || '获取生产记录失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchProductionData()
})
</script>

<style scoped>
.production-detail {
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

.details-container {
  flex: 1;
  min-width: 0;
}

.info-card,
.action-card,
.details-card {
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

.header-icon.details-icon {
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
  gap: 6px;
  font-size: 12px;
  color: var(--text-color-secondary);
  font-weight: 500;
  margin-bottom: 6px;
}

.label-icon {
  font-size: 14px;
  color: var(--text-color-tertiary);
}

.info-value {
  font-size: 14px;
  color: var(--text-color);
  font-weight: 500;
  word-break: break-all;
}

.record-id,
.part-id,
.batch-no {
  font-family: 'Courier New', monospace;
  color: var(--primary-color);
  font-weight: 600;
}

.action-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
  padding: 4px;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 16px;
  background: var(--bg-color);
  border-radius: 12px;
  border: 1px solid var(--border-color);
  cursor: pointer;
  transition: all 0.3s;
}

.action-item:hover {
  border-color: var(--primary-color);
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.15);
  transform: translateY(-2px);
}

.action-item.delete-action:hover {
  border-color: #ff4d4f;
  box-shadow: 0 2px 8px rgba(255, 77, 79, 0.15);
}

.action-item.delete-action.deleting {
  opacity: 0.6;
  cursor: not-allowed;
  pointer-events: none;
}

.action-item.delete-action.deleting:hover {
  transform: none;
  box-shadow: none;
}

.action-icon-item {
  font-size: 20px;
  color: var(--primary-color);
  transition: all 0.3s;
}

.action-icon-item.spinning {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

.action-item.delete-action .action-icon-item {
  color: #ff4d4f;
}

.action-item span {
  font-size: 13px;
  color: var(--text-color);
  font-weight: 500;
}

.operation-details {
  padding: 4px;
}

.remarks {
  white-space: pre-wrap;
  line-height: 1.6;
}

.timeline-wrapper {
  padding: 4px;
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
  color: var(--text-color);
}

.timeline-time {
  font-size: 12px;
  color: var(--text-color-tertiary);
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
  color: var(--text-color-secondary);
}

.timeline-desc .anticon {
  font-size: 14px;
  color: var(--text-color-tertiary);
}

@media (max-width: 1200px) {
  .main-content {
    flex-direction: column;
  }
  
  .info-container {
    width: 100%;
  }
}

@media (max-width: 768px) {
  .production-detail {
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
