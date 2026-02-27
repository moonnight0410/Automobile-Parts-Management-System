<template>
  <div class="quality-detail">
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">质检记录详情</h1>
          <span class="page-subtitle">查看质检记录完整信息</span>
        </div>
      </div>
      <div class="top-bar-right">
        <a-tag :color="getResultColor(qualityData?.inspectionResult || '')" class="status-badge">
          {{ qualityData?.inspectionResult || '-' }}
        </a-tag>
      </div>
    </div>

    <a-spin :spinning="loading">
      <div class="main-content">
        <div class="info-container">
          <a-card :bordered="false" class="info-card">
            <div class="card-header">
              <div class="header-icon">
                <SafetyCertificateOutlined />
              </div>
              <div class="header-text">
                <h3>基本信息</h3>
                <p>质检记录核心属性</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="info-grid">
              <div class="info-item">
                <div class="info-label">
                  <NumberOutlined class="label-icon" />
                  <span>质检ID</span>
                </div>
                <div class="info-value inspection-id">{{ qualityData?.inspectionID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <AppstoreOutlined class="label-icon" />
                  <span>零部件ID</span>
                </div>
                <div class="info-value part-id">{{ qualityData?.partID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <BarcodeOutlined class="label-icon" />
                  <span>批次号</span>
                </div>
                <div class="info-value batch-no">{{ qualityData?.batchNo || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ExperimentOutlined class="label-icon" />
                  <span>质检类型</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getInspectionTypeColor(qualityData?.inspectionType || '')" class="type-tag">
                    {{ qualityData?.inspectionType || '-' }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <UserOutlined class="label-icon" />
                  <span>质检员</span>
                </div>
                <div class="info-value">{{ qualityData?.inspector || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CheckCircleOutlined class="label-icon" />
                  <span>处理人</span>
                </div>
                <div class="info-value">{{ qualityData?.handler || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CheckCircleOutlined class="label-icon" />
                  <span>质检结果</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getResultColor(qualityData?.inspectionResult || '')" class="status-tag">
                    {{ qualityData?.inspectionResult || '-' }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ClockCircleOutlined class="label-icon" />
                  <span>质检时间</span>
                </div>
                <div class="info-value">{{ qualityData?.inspectionTime || '-' }}</div>
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
                <h3>质检详情</h3>
                <p>质检详细信息</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="quality-details">
              <div class="info-item" v-if="qualityData?.defectDescription">
                <div class="info-label">缺陷描述</div>
                <div class="info-value defect">{{ qualityData.defectDescription }}</div>
              </div>
              <div class="info-item" v-if="qualityData?.handlingMethod">
                <div class="info-label">处理方法</div>
                <div class="info-value handling">{{ qualityData.handlingMethod }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">状态</div>
                <div class="info-value">
                  <a-tag :color="getStatusColor(qualityData?.status || '')" class="status-tag">
                    {{ qualityData?.status || '-' }}
                  </a-tag>
                </div>
              </div>
            </div>
          </a-card>

          <a-card :bordered="false" class="details-card">
            <div class="card-header">
              <div class="header-icon details-icon">
                <HistoryOutlined />
              </div>
              <div class="header-text">
                <h3>质检时间线</h3>
                <p>质检流程记录</p>
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
                    <span class="timeline-title">质检开始</span>
                    <span class="timeline-time">{{ qualityData?.inspectionTime }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <UserOutlined />
                      质检员: {{ qualityData?.inspector }}
                    </div>
                    <div class="timeline-desc">
                      <ExperimentOutlined />
                      质检类型: {{ qualityData?.inspectionType }}
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
                    <span class="timeline-title">质检完成</span>
                    <span class="timeline-time">{{ qualityData?.inspectionTime }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <CheckCircleOutlined />
                      质检结果: 
                      <a-tag :color="getResultColor(qualityData?.inspectionResult || '')" size="small">
                        {{ qualityData?.inspectionResult }}
                      </a-tag>
                    </div>
                    <div class="timeline-desc">
                      <UserOutlined />
                      处理人: {{ qualityData?.handler }}
                    </div>
                  </div>
                </div>
              </a-timeline-item>

              <a-timeline-item color="orange" v-if="qualityData?.defectDescription">
                <template #dot>
                  <div class="timeline-dot orange">
                    <AlertOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">缺陷处理</span>
                    <span class="timeline-time">{{ qualityData?.inspectionTime }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <AlertOutlined />
                      缺陷描述: {{ qualityData?.defectDescription }}
                    </div>
                    <div class="timeline-desc">
                      <ToolOutlined />
                      处理方法: {{ qualityData?.handlingMethod }}
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
import { listQualityInspections, deleteQualityInspection } from '../../services/quality.service'
import type { QualityInspection } from '../../services/quality.service'
import {
  ArrowLeftOutlined,
  NumberOutlined,
  AppstoreOutlined,
  BarcodeOutlined,
  UserOutlined,
  CheckCircleOutlined,
  ClockCircleOutlined,
  ThunderboltOutlined,
  SafetyCertificateOutlined,
  EditOutlined,
  PrinterOutlined,
  ExportOutlined,
  DeleteOutlined,
  FileTextOutlined,
  HistoryOutlined,
  PlayCircleOutlined,
  ExperimentOutlined,
  AlertOutlined,
  ToolOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()

const loading = ref(false)
const deleting = ref(false)
const qualityData = ref<QualityInspection | null>(null)

function goBack() {
  router.push('/production/quality')
}

function getResultColor(result: string) {
  const colors: Record<string, string> = {
    '合格': 'success',
    '不合格': 'error',
    '待检': 'warning'
  }
  return colors[result] || 'default'
}

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    '已完成': 'success',
    '处理中': 'processing',
    '待处理': 'warning'
  }
  return colors[status] || 'default'
}

function getInspectionTypeColor(type: string) {
  const colors: Record<string, string> = {
    '入库质检': 'blue',
    '出库质检': 'cyan',
    '过程质检': 'purple',
    '最终质检': 'green'
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
  if (qualityData.value?.partID) {
    router.push(`/parts/detail/${qualityData.value.partID}`)
  } else {
    message.warning('零部件ID不存在')
  }
}

async function handleDelete() {
  const inspectionID = route.params.id as string
  if (!inspectionID) {
    message.error('质检ID不存在')
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
        `您确定要删除质检记录 ${inspectionID} 吗？`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `零部件ID：${qualityData.value?.partID || '-'}`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `批次号：${qualityData.value?.batchNo || '-'}`
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
        const response = await deleteQualityInspection(inspectionID)
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

async function fetchQualityData() {
  loading.value = true
  try {
    const response = await listQualityInspections()
    if (response.code === 0 && response.data) {
      const inspectionID = route.params.id as string
      const record = response.data.find((item: QualityInspection) => item.inspectionID === inspectionID)
      if (record) {
        qualityData.value = record
      } else {
        message.error('未找到该质检记录')
      }
    } else {
      message.error(response.message || '获取质检记录失败')
    }
  } catch (error: any) {
    message.error(error.message || '获取质检记录失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchQualityData()
})
</script>

<style scoped>
.quality-detail {
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
  max-width: 100%;
  overflow: hidden;
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

.inspection-id,
.part-id,
.batch-no {
  font-family: 'Courier New', monospace;
  color: #667eea;
  font-weight: 600;
}

.type-tag,
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

.action-item.delete-action.deleting {
  opacity: 0.6;
  cursor: not-allowed;
  pointer-events: none;
}

.action-item.delete-action.deleting:hover {
  background: linear-gradient(135deg, #fee2e2 0%, #fecaca 100%);
  border-color: #ef4444;
}

.action-item.delete-action .action-icon-item {
  color: #ef4444;
  transition: all 0.3s;
}

.action-item.delete-action .action-icon-item.spinning {
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

.details-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
  max-width: 100%;
  overflow: hidden;
}

.details-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.quality-details {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 0 24px 24px;
}

.defect,
.handling {
  white-space: pre-wrap;
  line-height: 1.6;
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
  .quality-detail {
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
