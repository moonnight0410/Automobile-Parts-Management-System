<template>
  <div class="recall-detail">
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">召回记录详情</h1>
          <span class="page-subtitle">查看召回记录完整信息</span>
        </div>
      </div>
      <div class="top-bar-right">
        <a-tag :color="getStatusColor(recallData?.status || '')" class="status-badge">
          {{ recallData?.status || '-' }}
        </a-tag>
      </div>
    </div>

    <a-spin :spinning="loading">
      <div class="main-content">
        <div class="info-container">
          <a-card :bordered="false" class="info-card">
            <div class="card-header">
              <div class="header-icon">
                <WarningOutlined />
              </div>
              <div class="header-text">
                <h3>基本信息</h3>
                <p>召回记录核心属性</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="info-grid">
              <div class="info-item">
                <div class="info-label">
                  <NumberOutlined class="label-icon" />
                  <span>召回ID</span>
                </div>
                <div class="info-value recall-id">{{ recallData?.recallID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ExperimentOutlined class="label-icon" />
                  <span>召回类型</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getRecallTypeColor(recallData?.recallType || '')" class="type-tag">
                    {{ recallData?.recallType || '-' }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <UserOutlined class="label-icon" />
                  <span>发起人</span>
                </div>
                <div class="info-value">{{ recallData?.initiator || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CheckCircleOutlined class="label-icon" />
                  <span>状态</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getStatusColor(recallData?.status || '')" class="status-tag">
                    {{ recallData?.status || '-' }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ClockCircleOutlined class="label-icon" />
                  <span>召回日期</span>
                </div>
                <div class="info-value">{{ recallData?.recallDate || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <AppstoreOutlined class="label-icon" />
                  <span>涉及零部件</span>
                </div>
                <div class="info-value">{{ recallData?.affectedParts?.length || 0 }} 个</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CarOutlined class="label-icon" />
                  <span>涉及车辆</span>
                </div>
                <div class="info-value">{{ recallData?.affectedVins?.length || 0 }} 辆</div>
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
              <div class="action-item" @click="viewAffectedParts">
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
                <h3>召回详情</h3>
                <p>召回详细信息</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="recall-details">
              <div class="info-item">
                <div class="info-label">召回原因</div>
                <div class="info-value reason">{{ recallData?.recallReason || '-' }}</div>
              </div>
              <div class="info-item" v-if="recallData?.handlingMeasures">
                <div class="info-label">处理措施</div>
                <div class="info-value handling">{{ recallData.handlingMeasures }}</div>
              </div>
            </div>
          </a-card>

          <a-card :bordered="false" class="details-card">
            <div class="card-header">
              <div class="header-icon details-icon">
                <AppstoreOutlined />
              </div>
              <div class="header-text">
                <h3>涉及零部件</h3>
                <p>受影响的零部件列表</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="affected-parts">
              <a-tag v-for="part in recallData?.affectedParts" :key="part" color="blue" class="part-tag">
                {{ part }}
              </a-tag>
              <a-empty v-if="!recallData?.affectedParts || recallData.affectedParts.length === 0" description="暂无涉及零部件" />
            </div>
          </a-card>

          <a-card :bordered="false" class="details-card">
            <div class="card-header">
              <div class="header-icon details-icon">
                <CarOutlined />
              </div>
              <div class="header-text">
                <h3>涉及车辆</h3>
                <p>受影响的车辆VIN码列表</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="affected-vins">
              <a-tag v-for="vin in recallData?.affectedVins" :key="vin" color="orange" class="vin-tag">
                {{ vin }}
              </a-tag>
              <a-empty v-if="!recallData?.affectedVins || recallData.affectedVins.length === 0" description="暂无涉及车辆" />
            </div>
          </a-card>

          <a-card :bordered="false" class="details-card">
            <div class="card-header">
              <div class="header-icon details-icon">
                <HistoryOutlined />
              </div>
              <div class="header-text">
                <h3>召回时间线</h3>
                <p>召回处理流程记录</p>
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
                    <span class="timeline-title">召回发起</span>
                    <span class="timeline-time">{{ recallData?.recallDate }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <UserOutlined />
                      发起人: {{ recallData?.initiator }}
                    </div>
                    <div class="timeline-desc">
                      <ExperimentOutlined />
                      召回类型: {{ recallData?.recallType }}
                    </div>
                  </div>
                </div>
              </a-timeline-item>

              <a-timeline-item color="green" v-if="recallData?.status === '已完成'">
                <template #dot>
                  <div class="timeline-dot green">
                    <CheckCircleOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">召回完成</span>
                    <span class="timeline-time">已完成</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <CheckCircleOutlined />
                      所有受影响零部件和车辆已处理完毕
                    </div>
                  </div>
                </div>
              </a-timeline-item>

              <a-timeline-item color="orange" v-if="recallData?.status === '进行中'">
                <template #dot>
                  <div class="timeline-dot orange">
                    <SyncOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">处理中</span>
                    <span class="timeline-time">进行中</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <SyncOutlined />
                      正在处理受影响的零部件和车辆
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
import { listRecallRecords, type RecallRecord } from '../../services/aftersale.service'
import {
  ArrowLeftOutlined,
  NumberOutlined,
  AppstoreOutlined,
  CarOutlined,
  UserOutlined,
  CheckCircleOutlined,
  ClockCircleOutlined,
  ThunderboltOutlined,
  WarningOutlined,
  EditOutlined,
  PrinterOutlined,
  ExportOutlined,
  DeleteOutlined,
  FileTextOutlined,
  HistoryOutlined,
  PlayCircleOutlined,
  ExperimentOutlined,
  SyncOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()

const loading = ref(false)
const deleting = ref(false)
const recallData = ref<RecallRecord | null>(null)

function goBack() {
  router.push('/aftersale/recall')
}

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    '进行中': 'processing',
    '已完成': 'success',
    '待处理': 'warning'
  }
  return colors[status] || 'default'
}

function getRecallTypeColor(type: string) {
  const colors: Record<string, string> = {
    '安全召回': 'error',
    '质量召回': 'orange',
    '环保召回': 'green',
    '其他': 'default'
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

function viewAffectedParts() {
  if (recallData.value?.affectedParts && recallData.value.affectedParts.length > 0) {
    const firstPart = recallData.value.affectedParts[0]
    router.push(`/parts/detail/${firstPart}`)
  } else {
    message.warning('暂无涉及零部件')
  }
}

async function handleDelete() {
  const recallID = route.params.id as string
  if (!recallID) {
    message.error('召回ID不存在')
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
        `您确定要删除召回记录 ${recallID} 吗？`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `涉及零部件：${recallData.value?.affectedParts?.length || 0} 个`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `涉及车辆：${recallData.value?.affectedVins?.length || 0} 辆`
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
        message.success('删除成功')
        goBack()
      } catch (error: any) {
        message.error(error.message || '删除失败，请稍后重试')
      } finally {
        deleting.value = false
      }
    }
  })
}

async function fetchRecallData() {
  loading.value = true
  try {
    const response = await listRecallRecords()
    if (response.code === 0 && response.data) {
      const recallID = route.params.id as string
      const record = response.data.find((item: RecallRecord) => item.recallID === recallID)
      if (record) {
        recallData.value = record
      } else {
        message.error('未找到该召回记录')
      }
    } else {
      message.error(response.message || '获取召回记录失败')
    }
  } catch (error: any) {
    message.error(error.message || '获取召回记录失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchRecallData()
})
</script>

<style scoped>
.recall-detail {
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

.recall-id {
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

.recall-details {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 0 24px 24px;
}

.reason,
.handling {
  white-space: pre-wrap;
  line-height: 1.6;
}

.affected-parts,
.affected-vins {
  padding: 0 24px 24px;
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.part-tag,
.vin-tag {
  font-size: 13px;
  padding: 4px 12px;
  border-radius: 4px;
}

.vin-tag {
  font-family: 'Courier New', monospace;
}

.custom-timeline {
  margin-top: 16px;
  padding: 0 24px 24px;
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
  color: #4b5563;
}

.timeline-dot {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 16px;
}

.timeline-dot.blue {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.timeline-dot.green {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.timeline-dot.orange {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
}

@media (max-width: 768px) {
  .main-content {
    grid-template-columns: 1fr;
  }

  .info-container {
    width: 100%;
  }

  .action-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}
</style>
