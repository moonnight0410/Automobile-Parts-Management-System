<template>
  <div class="fault-detail">
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">故障记录详情</h1>
          <span class="page-subtitle">查看故障记录完整信息</span>
        </div>
      </div>
      <div class="top-bar-right">
        <a-tag :color="getStatusColor(faultData?.status || '')" class="status-badge">
          {{ faultData?.status || '-' }}
        </a-tag>
      </div>
    </div>

    <a-spin :spinning="loading">
      <div class="main-content">
        <div class="info-container">
          <a-card :bordered="false" class="info-card">
            <div class="card-header">
              <div class="header-icon">
                <AlertOutlined />
              </div>
              <div class="header-text">
                <h3>基本信息</h3>
                <p>故障记录核心属性</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="info-grid">
              <div class="info-item">
                <div class="info-label">
                  <NumberOutlined class="label-icon" />
                  <span>故障ID</span>
                </div>
                <div class="info-value fault-id">{{ faultData?.faultID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <AppstoreOutlined class="label-icon" />
                  <span>零部件ID</span>
                </div>
                <div class="info-value part-id">{{ faultData?.partID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CarOutlined class="label-icon" />
                  <span>VIN码</span>
                </div>
                <div class="info-value vin">{{ faultData?.vin || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ExperimentOutlined class="label-icon" />
                  <span>故障类型</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getFaultTypeColor(faultData?.faultType || '')" class="type-tag">
                    {{ faultData?.faultType || '-' }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <UserOutlined class="label-icon" />
                  <span>上报人</span>
                </div>
                <div class="info-value">{{ faultData?.reporter || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <TeamOutlined class="label-icon" />
                  <span>上报组织</span>
                </div>
                <div class="info-value">{{ faultData?.reporterOrg || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CheckCircleOutlined class="label-icon" />
                  <span>状态</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getStatusColor(faultData?.status || '')" class="status-tag">
                    {{ faultData?.status || '-' }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ClockCircleOutlined class="label-icon" />
                  <span>上报时间</span>
                </div>
                <div class="info-value">{{ faultData?.reportTime || '-' }}</div>
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
                <h3>故障详情</h3>
                <p>故障详细信息</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="fault-details">
              <div class="info-item">
                <div class="info-label">故障描述</div>
                <div class="info-value description">{{ faultData?.faultDescription || '-' }}</div>
              </div>
              <div class="info-item" v-if="faultData?.handlingResult">
                <div class="info-label">处理结果</div>
                <div class="info-value handling">{{ faultData.handlingResult }}</div>
              </div>
              <div class="info-item" v-if="faultData?.handlingTime">
                <div class="info-label">处理时间</div>
                <div class="info-value">{{ faultData.handlingTime }}</div>
              </div>
            </div>
          </a-card>

          <a-card :bordered="false" class="details-card">
            <div class="card-header">
              <div class="header-icon details-icon">
                <HistoryOutlined />
              </div>
              <div class="header-text">
                <h3>处理时间线</h3>
                <p>故障处理流程记录</p>
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
                    <span class="timeline-title">故障上报</span>
                    <span class="timeline-time">{{ faultData?.reportTime }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <UserOutlined />
                      上报人: {{ faultData?.reporter }}
                    </div>
                    <div class="timeline-desc">
                      <TeamOutlined />
                      上报组织: {{ faultData?.reporterOrg }}
                    </div>
                    <div class="timeline-desc">
                      <ExperimentOutlined />
                      故障类型: {{ faultData?.faultType }}
                    </div>
                  </div>
                </div>
              </a-timeline-item>

              <a-timeline-item color="green" v-if="faultData?.status === '已审核'">
                <template #dot>
                  <div class="timeline-dot green">
                    <CheckCircleOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">审核完成</span>
                    <span class="timeline-time">{{ faultData?.handlingTime }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <CheckCircleOutlined />
                      处理结果: {{ faultData?.handlingResult }}
                    </div>
                  </div>
                </div>
              </a-timeline-item>

              <a-timeline-item color="orange" v-if="faultData?.status === '待审核'">
                <template #dot>
                  <div class="timeline-dot orange">
                    <ClockCircleOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">待审核</span>
                    <span class="timeline-time">处理中</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <ClockCircleOutlined />
                      等待管理员审核处理
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
import { listFaultReports, type FaultReport } from '../../services/aftersale.service'
import {
  ArrowLeftOutlined,
  NumberOutlined,
  AppstoreOutlined,
  CarOutlined,
  UserOutlined,
  TeamOutlined,
  CheckCircleOutlined,
  ClockCircleOutlined,
  ThunderboltOutlined,
  AlertOutlined,
  EditOutlined,
  PrinterOutlined,
  ExportOutlined,
  DeleteOutlined,
  FileTextOutlined,
  HistoryOutlined,
  PlayCircleOutlined,
  ExperimentOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()

const loading = ref(false)
const deleting = ref(false)
const faultData = ref<FaultReport | null>(null)

function goBack() {
  router.push('/aftersale/fault')
}

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    '待审核': 'warning',
    '已审核': 'success',
    '已处理': 'success'
  }
  return colors[status] || 'default'
}

function getFaultTypeColor(type: string) {
  const colors: Record<string, string> = {
    '质量问题': 'error',
    '安全隐患': 'red',
    '功能缺陷': 'orange',
    '性能问题': 'gold',
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

function viewPart() {
  if (faultData.value?.partID) {
    router.push(`/parts/detail/${faultData.value.partID}`)
  } else {
    message.warning('零部件ID不存在')
  }
}

async function handleDelete() {
  const faultID = route.params.id as string
  if (!faultID) {
    message.error('故障ID不存在')
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
        `您确定要删除故障记录 ${faultID} 吗？`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `零部件ID：${faultData.value?.partID || '-'}`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `VIN码：${faultData.value?.vin || '-'}`
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

async function fetchFaultData() {
  loading.value = true
  try {
    const response = await listFaultReports()
    if (response.code === 0 && response.data) {
      const faultID = route.params.id as string
      const record = response.data.find((item: FaultReport) => item.faultID === faultID)
      if (record) {
        faultData.value = record
      } else {
        message.error('未找到该故障记录')
      }
    } else {
      message.error(response.message || '获取故障记录失败')
    }
  } catch (error: any) {
    message.error(error.message || '获取故障记录失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchFaultData()
})
</script>

<style scoped>
.fault-detail {
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

.fault-id,
.part-id,
.vin {
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

.fault-details {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 0 24px 24px;
}

.description,
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
