<!--
  故障报告页面
-->

<template>
  <div class="fault-report">
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <AlertOutlined />
          </div>
          <div class="title-text">
            <h2>故障报告</h2>
            <p class="subtitle">管理零部件故障信息与风险预警</p>
          </div>
        </div>
        <a-button type="primary" class="create-btn" @click="router.push('/aftersale/fault/create')">
          <template #icon><PlusOutlined /></template>
          上报故障
        </a-button>
      </div>
      <div class="stats-overview">
        <div class="stat-item">
          <div class="stat-value">{{ totalCount }}</div>
          <div class="stat-label">总报告数</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-pending">{{ pendingCount }}</div>
          <div class="stat-label">待审核</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-reviewed">{{ reviewedCount }}</div>
          <div class="stat-label">已审核</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-high">{{ highRiskCount }}</div>
          <div class="stat-label">高风险</div>
        </div>
      </div>
    </div>

    <a-card :bordered="false" class="search-card">
      <div class="search-header">
        <div class="header-left">
          <div class="filter-icon-wrapper">
            <FilterOutlined />
          </div>
          <div class="header-info">
            <span class="header-title">筛选条件</span>
            <span class="header-desc">支持多条件组合查询</span>
          </div>
        </div>
        <div class="header-right">
          <a-tag color="orange" class="filter-tag">
            <template #icon><SearchOutlined /></template>
            故障查询
          </a-tag>
        </div>
      </div>
      <a-form layout="inline" :model="searchForm" class="search-form">
        <a-row :gutter="[20, 20]" style="width: 100%">
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <NumberOutlined class="label-icon" />
                <span>故障ID</span>
              </div>
              <a-input
                v-model:value="searchForm.faultID"
                placeholder="请输入故障ID"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <AlertOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <BarcodeOutlined class="label-icon" />
                <span>零部件ID</span>
              </div>
              <a-input
                v-model:value="searchForm.partID"
                placeholder="请输入零部件ID"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <AppstoreOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <CarOutlined class="label-icon" />
                <span>VIN码</span>
              </div>
              <a-input
                v-model:value="searchForm.vin"
                placeholder="请输入VIN码"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <CarOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <CheckCircleOutlined class="label-icon" />
                <span>状态</span>
              </div>
              <a-select
                v-model:value="searchForm.status"
                placeholder="请选择状态"
                allow-clear
                class="filter-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                :dropdown-style="{ zIndex: 1050 }"
              >
                <a-select-option value="待审核">
                  <span class="status-option"><span class="status-dot dot-pending"></span>待审核</span>
                </a-select-option>
                <a-select-option value="已审核">
                  <span class="status-option"><span class="status-dot dot-reviewed"></span>已审核</span>
                </a-select-option>
              </a-select>
            </div>
          </a-col>
        </a-row>
        <div class="search-actions">
          <div class="action-left">
            <span class="action-hint">
              <InfoCircleOutlined />
              可同时使用多个条件进行精确筛选
            </span>
          </div>
          <div class="action-right">
            <a-button @click="handleReset" class="reset-btn">
              <template #icon><ReloadOutlined /></template>
              重置
            </a-button>
            <a-button type="primary" @click="handleSearch" class="search-btn">
              <template #icon><SearchOutlined /></template>
              开始搜索
            </a-button>
          </div>
        </div>
      </a-form>
    </a-card>

    <a-card :bordered="false" class="table-card">
      <div class="table-header">
        <div class="table-header-left">
          <span class="table-title">
            <UnorderedListOutlined class="title-icon" />
            数据列表
          </span>
          <a-tag color="blue" class="record-count">
            <DatabaseOutlined />
            共 {{ tableData.length }} 条记录
          </a-tag>
        </div>
        <div class="table-header-right">
          <a-space :size="12">
            <a-button class="toolbar-btn refresh-btn" @click="handleRefresh" :loading="refreshing">
              <template #icon><ReloadOutlined :spin="refreshing" /></template>
              {{ refreshing ? '刷新中...' : '刷新数据' }}
            </a-button>
            <a-button class="toolbar-btn export-btn" @click="handleExport">
              <template #icon><ExportOutlined /></template>
              导出报告
            </a-button>
          </a-space>
        </div>
      </div>
      <a-table
        :columns="columns"
        :data-source="tableData"
        row-key="faultID"
        class="custom-table"
        :scroll="{ x: 1200 }"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'faultID'">
            <div class="fault-id-cell">
              <AlertOutlined class="fault-icon" />
              <span>{{ record.faultID }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'partID'">
            <div class="part-id-cell">
              <BarcodeOutlined class="part-icon" />
              <span>{{ record.partID }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'faultType'">
            <a-tag :color="getFaultTypeColor(record.faultType)" class="type-tag">
              {{ record.faultType }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'riskProb'">
            <div class="risk-cell">
              <a-progress
                :percent="parseInt(record.riskProb)"
                :stroke-color="getRiskColor(record.riskProb)"
                size="small"
                :show-info="false"
                class="risk-progress"
              />
              <span class="risk-value" :style="{ color: getRiskColor(record.riskProb) }">
                {{ record.riskProb }}
              </span>
            </div>
          </template>
          <template v-else-if="column.key === 'status'">
            <div class="status-cell">
              <span class="status-dot" :class="getStatusDotClass(record.status)"></span>
              <span class="status-text">{{ record.status }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'reportTime'">
            <div class="time-cell">
              <ClockCircleOutlined class="time-icon" />
              <span>{{ record.reportTime }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'action'">
            <a-space class="action-cell">
              <a-button type="link" size="small" @click="viewDetail(record)" class="action-btn detail-btn">
                <EyeOutlined />
                详情
              </a-button>
              <a-button
                type="link"
                size="small"
                @click="handleReview(record)"
                :disabled="record.status === '已审核'"
                class="action-btn review-btn"
              >
                <CheckCircleOutlined />
                审核
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { listFaultReports, type FaultReport } from '../../services/aftersale.service'
import {
  AlertOutlined,
  PlusOutlined,
  FilterOutlined,
  SearchOutlined,
  NumberOutlined,
  BarcodeOutlined,
  AppstoreOutlined,
  CarOutlined,
  CheckCircleOutlined,
  InfoCircleOutlined,
  ReloadOutlined,
  UnorderedListOutlined,
  DatabaseOutlined,
  ExportOutlined,
  ClockCircleOutlined,
  EyeOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const refreshing = ref(false)
const loading = ref(false)

const searchForm = ref({
  faultID: '',
  partID: '',
  vin: '',
  status: undefined as string | undefined
})

const columns = [
  { title: '故障ID', key: 'faultID', width: 140 },
  { title: '零部件ID', key: 'partID', width: 140 },
  { title: 'VIN码', dataIndex: 'vin', width: 160 },
  { title: '故障类型', key: 'faultType', width: 120 },
  { title: '风险概率', key: 'riskProb', width: 150 },
  { title: '状态', key: 'status', width: 100 },
  { title: '上报时间', key: 'reportTime', width: 150 },
  { title: '操作', key: 'action', width: 150, fixed: 'right' as const }
]

const tableData = ref<any[]>([])

async function fetchData() {
  loading.value = true
  try {
    const response = await listFaultReports()
    if (response.code === 0 && response.data) {
      tableData.value = response.data.map((item: FaultReport) => ({
        ...item,
        riskProb: item.riskProbability ? `${item.riskProbability}%` : '0%'
      }))
    }
  } catch (error: any) {
    message.error(error.message || '获取数据失败')
    tableData.value = []
  } finally {
    loading.value = false
  }
}

const totalCount = computed(() => tableData.value.length)
const pendingCount = computed(() => tableData.value.filter(item => item.status === '待审核' || item.status === 'PENDING').length)
const reviewedCount = computed(() => tableData.value.filter(item => item.status === '已审核' || item.status === 'REVIEWED').length)
const highRiskCount = computed(() => tableData.value.filter(item => parseInt(item.riskProb) >= 70).length)

const getFaultTypeColor = (type: string) => {
  const colorMap: Record<string, string> = {
    '制动故障': 'red',
    '转向故障': 'orange',
    '发动机故障': 'purple',
    '电气故障': 'blue',
    '其他': 'default'
  }
  return colorMap[type] || 'default'
}

const getRiskColor = (risk: string) => {
  const value = parseInt(risk)
  if (value >= 70) return '#ef4444'
  if (value >= 40) return '#f59e0b'
  return '#10b981'
}

const getStatusDotClass = (status: string) => {
  return status === '已审核' || status === 'REVIEWED' ? 'dot-reviewed' : 'dot-pending'
}

const handleSearch = () => {
  message.success('搜索完成')
}

const handleReset = () => {
  searchForm.value = {
    faultID: '',
    partID: '',
    vin: '',
    status: undefined
  }
  fetchData()
}

const handleRefresh = async () => {
  refreshing.value = true
  try {
    await fetchData()
    message.success('数据已刷新')
  } catch (error: any) {
    message.error('刷新失败')
  } finally {
    setTimeout(() => {
      refreshing.value = false
    }, 500)
  }
}

const handleExport = () => {
  message.success('故障报告导出成功')
}

const viewDetail = (record: any) => {
  message.info(`查看故障详情: ${record.faultID}`)
}

const handleReview = (record: any) => {
  if (record.status === '已审核' || record.status === 'REVIEWED') return
  record.status = '已审核'
  message.success(`故障 ${record.faultID} 已审核通过`)
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.fault-report {
  padding: 0;
  min-height: 100%;
}

.page-header {
  margin-bottom: 24px;
  background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);
  border-radius: 12px;
  padding: 24px 32px;
  box-shadow: 0 4px 20px rgba(249, 115, 22, 0.25);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.header-title {
  display: flex;
  align-items: center;
  gap: 16px;
}

.title-icon {
  width: 48px;
  height: 48px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  color: #fff;
}

.title-text h2 {
  font-size: 24px;
  font-weight: 600;
  margin: 0;
  color: #fff;
}

.title-text .subtitle {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.85);
  margin: 4px 0 0;
}

.create-btn {
  height: 40px;
  border-radius: 10px;
  font-size: 14px;
  background: rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: #fff;
}

.create-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  border-color: rgba(255, 255, 255, 0.4);
  color: #fff;
}

.stats-overview {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
}

.stat-item {
  background: rgba(255, 255, 255, 0.15);
  border-radius: 10px;
  padding: 16px;
  text-align: center;
}

.stat-value {
  font-size: 28px;
  font-weight: 700;
  color: #fff;
  line-height: 1;
}

.stat-value.status-pending {
  color: #fbbf24;
}

.stat-value.status-reviewed {
  color: #34d399;
}

.stat-value.status-high {
  color: #f87171;
}

.stat-label {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.85);
  margin-top: 6px;
}

.search-card {
  border-radius: 12px;
  margin-bottom: 24px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.search-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.filter-icon-wrapper {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: #fff;
}

.header-info {
  display: flex;
  flex-direction: column;
}

.header-title {
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
}

.header-desc {
  font-size: 13px;
  color: #6b7280;
  margin-top: 2px;
}

.filter-tag {
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 12px;
}

.search-form {
  width: 100%;
}

.filter-item {
  width: 100%;
}

.filter-label {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 8px;
  font-size: 13px;
  font-weight: 500;
  color: #374151;
}

.label-icon {
  color: #f97316;
}

.filter-input,
.filter-select {
  width: 100%;
}

.input-prefix-icon {
  color: #9ca3af;
}

.status-option {
  display: flex;
  align-items: center;
  gap: 8px;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.dot-pending {
  background: #f59e0b;
}

.dot-reviewed {
  background: #10b981;
}

.search-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 20px;
  padding-top: 16px;
  border-top: 1px solid #f3f4f6;
}

.action-hint {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #9ca3af;
}

.action-right {
  display: flex;
  gap: 12px;
}

.reset-btn,
.search-btn {
  height: 38px;
  border-radius: 8px;
  font-size: 14px;
}

.reset-btn {
  border: 1px solid #e5e7eb;
  color: #6b7280;
}

.reset-btn:hover {
  border-color: #f97316;
  color: #f97316;
}

.search-btn {
  background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);
  border: none;
}

.search-btn:hover {
  background: linear-gradient(135deg, #ea580c 0%, #c2410c 100%);
}

.table-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.table-header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.table-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 15px;
  font-weight: 600;
  color: #1f2937;
}

.title-icon {
  color: #f97316;
}

.record-count {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 12px;
}

.toolbar-btn {
  height: 36px;
  border-radius: 8px;
  font-size: 13px;
  border: 1px solid #e5e7eb;
  color: #6b7280;
}

.toolbar-btn:hover {
  border-color: #f97316;
  color: #f97316;
}

.custom-table :deep(.ant-table-thead > tr > th) {
  background: #f9fafb;
  font-weight: 600;
  color: #374151;
  border-bottom: 2px solid #e5e7eb;
}

.custom-table :deep(.ant-table-tbody > tr > td) {
  border-bottom: 1px solid #f3f4f6;
}

.custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #fff7ed;
}

.fault-id-cell,
.part-id-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.fault-icon {
  color: #f97316;
}

.part-icon {
  color: #6b7280;
}

.type-tag {
  border-radius: 6px;
}

.risk-cell {
  display: flex;
  align-items: center;
  gap: 12px;
}

.risk-progress {
  width: 80px;
}

.risk-value {
  font-weight: 600;
  font-size: 13px;
}

.status-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.status-dot.dot-pending {
  background: #f59e0b;
}

.status-dot.dot-reviewed {
  background: #10b981;
}

.status-text {
  font-size: 13px;
  color: #374151;
}

.time-cell {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #6b7280;
}

.time-icon {
  font-size: 12px;
}

.action-cell {
  gap: 4px;
}

.action-btn {
  padding: 0 8px;
  height: 28px;
  font-size: 12px;
}

.detail-btn {
  color: #6b7280;
}

.detail-btn:hover {
  color: #f97316;
}

.review-btn {
  color: #10b981;
}

.review-btn:hover:not(:disabled) {
  color: #059669;
}

.review-btn:disabled {
  color: #d1d5db;
  cursor: not-allowed;
}

.report-form :deep(.ant-input),
.report-form :deep(.ant-select-selector),
.report-form :deep(.ant-input-number) {
  border-radius: 8px;
}

[data-theme='dark'] .page-header {
  background: linear-gradient(135deg, #c2410c 0%, #9a3412 100%);
  box-shadow: 0 4px 20px rgba(194, 65, 12, 0.3);
}

[data-theme='dark'] .title-text h2 {
  color: #fff;
}

[data-theme='dark'] .title-text .subtitle {
  color: rgba(255, 255, 255, 0.8);
}

[data-theme='dark'] .create-btn {
  background: rgba(255, 255, 255, 0.15);
  border-color: rgba(255, 255, 255, 0.25);
  color: #fff;
}

[data-theme='dark'] .stat-item {
  background: rgba(255, 255, 255, 0.1);
}

[data-theme='dark'] .stat-value {
  color: #fff;
}

[data-theme='dark'] .stat-label {
  color: rgba(255, 255, 255, 0.8);
}

[data-theme='dark'] .search-card,
[data-theme='dark'] .table-card {
  background: #1f2937;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

[data-theme='dark'] .header-title,
[data-theme='dark'] .table-title {
  color: #f3f4f6;
}

[data-theme='dark'] .header-desc,
[data-theme='dark'] .action-hint {
  color: #9ca3af;
}

[data-theme='dark'] .filter-label {
  color: #e5e7eb;
}

[data-theme='dark'] .reset-btn,
[data-theme='dark'] .toolbar-btn {
  border-color: #374151;
  color: #9ca3af;
}

[data-theme='dark'] .reset-btn:hover,
[data-theme='dark'] .toolbar-btn:hover {
  border-color: #f97316;
  color: #fb923c;
}

[data-theme='dark'] .search-btn {
  background: linear-gradient(135deg, #ea580c 0%, #c2410c 100%);
}

[data-theme='dark'] .custom-table :deep(.ant-table-thead > tr > th) {
  background: #111827;
  color: #e5e7eb;
  border-bottom-color: #374151;
}

[data-theme='dark'] .custom-table :deep(.ant-table-tbody > tr > td) {
  background: #1f2937;
  border-bottom-color: #374151;
  color: #e5e7eb;
}

[data-theme='dark'] .custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #111827;
}

[data-theme='dark'] .status-text {
  color: #e5e7eb;
}

[data-theme='dark'] .time-cell {
  color: #9ca3af;
}

[data-theme='dark'] .action-btn.detail-btn {
  color: #9ca3af;
}

[data-theme='dark'] .action-btn.detail-btn:hover {
  color: #fb923c;
}
</style>
