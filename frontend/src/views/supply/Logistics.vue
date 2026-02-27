<!--
  物流跟踪页面
  设计：与Dashboard风格统一
-->

<template>
  <div class="logistics">
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <EnvironmentOutlined />
          </div>
          <div class="title-text">
            <h2>物流跟踪</h2>
            <p class="subtitle">实时追踪货物运输状态与配送进度</p>
          </div>
        </div>
        <div class="header-actions">
          <a-button class="action-btn" @click="handleRefresh" :loading="refreshing">
            <template #icon><ReloadOutlined :spin="refreshing" /></template>
            刷新状态
          </a-button>
        </div>
      </div>
      <div class="stats-overview">
        <div class="stat-item">
          <div class="stat-value">{{ totalCount }}</div>
          <div class="stat-label">总物流单</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-transit">{{ inTransitCount }}</div>
          <div class="stat-label">运输中</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-delivered">{{ deliveredCount }}</div>
          <div class="stat-label">已送达</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-pending">{{ pendingCount }}</div>
          <div class="stat-label">待发货</div>
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
            <span class="header-desc">快速定位物流信息</span>
          </div>
        </div>
        <div class="header-right">
          <a-tag color="green" class="filter-tag">
            <template #icon><CarOutlined /></template>
            物流查询
          </a-tag>
        </div>
      </div>
      <a-form layout="inline" :model="searchForm" class="search-form">
        <a-row :gutter="[20, 20]" style="width: 100%">
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <NumberOutlined class="label-icon" />
                <span>物流ID</span>
              </div>
              <a-input
                v-model:value="searchForm.logisticsID"
                placeholder="请输入物流ID"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <EnvironmentOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <ShoppingOutlined class="label-icon" />
                <span>订单ID</span>
              </div>
              <a-input
                v-model:value="searchForm.orderID"
                placeholder="请输入订单ID"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <ShoppingCartOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <CarOutlined class="label-icon" />
                <span>物流商</span>
              </div>
              <a-select
                v-model:value="searchForm.carrier"
                placeholder="请选择物流商"
                allow-clear
                class="filter-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
              >
                <a-select-option value="顺丰物流">顺丰物流</a-select-option>
                <a-select-option value="京东物流">京东物流</a-select-option>
                <a-select-option value="中通快递">中通快递</a-select-option>
                <a-select-option value="圆通速递">圆通速递</a-select-option>
              </a-select>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <TagOutlined class="label-icon" />
                <span>运输状态</span>
              </div>
              <a-select
                v-model:value="searchForm.status"
                placeholder="请选择状态"
                allow-clear
                class="filter-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
              >
                <a-select-option value="待发货">
                  <span class="status-option"><span class="status-dot dot-pending"></span>待发货</span>
                </a-select-option>
                <a-select-option value="运输中">
                  <span class="status-option"><span class="status-dot dot-transit"></span>运输中</span>
                </a-select-option>
                <a-select-option value="已送达">
                  <span class="status-option"><span class="status-dot dot-delivered"></span>已送达</span>
                </a-select-option>
              </a-select>
            </div>
          </a-col>
        </a-row>
        <div class="search-actions">
          <div class="action-left">
            <span class="action-hint">
              <InfoCircleOutlined />
              输入物流ID或订单ID进行精确查询
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
            物流列表
          </span>
          <a-tag color="blue" class="record-count">
            <DatabaseOutlined />
            共 {{ filteredData.length }} 条记录
          </a-tag>
        </div>
        <div class="table-header-right">
          <a-space :size="12">
            <a-button class="toolbar-btn export-btn" @click="handleExport">
              <template #icon><ExportOutlined /></template>
              导出数据
            </a-button>
          </a-space>
        </div>
      </div>
      <a-table
        :columns="columns"
        :data-source="filteredData"
        row-key="logisticsID"
        class="custom-table"
        :scroll="{ x: 1200 }"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'logisticsID'">
            <div class="logistics-id-cell">
              <EnvironmentOutlined class="logistics-icon" />
              <span>{{ record.logisticsID }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'orderID'">
            <a-tag color="purple" class="order-tag">
              <ShoppingOutlined />
              {{ record.orderID }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'carrier'">
            <div class="carrier-cell">
              <CarOutlined class="carrier-icon" />
              <span>{{ record.carrier }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'status'">
            <a-tag :color="getStatusColor(record.status)" class="status-tag">
              <component :is="getStatusIcon(record.status)" />
              {{ record.status }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'startTime'">
            <div class="time-cell">
              <ClockCircleOutlined class="time-icon" />
              <span>{{ record.startTime }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'endTime'">
            <div class="time-cell">
              <CheckCircleOutlined class="time-icon success" v-if="record.status === '已送达'" />
              <ClockCircleOutlined class="time-icon" v-else />
              <span>{{ record.endTime || '预计中' }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'action'">
            <a-space class="action-cell">
              <a-button type="link" size="small" @click="viewTrack(record)" class="action-btn track-btn">
                <EnvironmentOutlined />
                轨迹
              </a-button>
              <a-button type="link" size="small" @click="viewDetail(record)" class="action-btn detail-btn">
                <EyeOutlined />
                详情
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <a-modal
      v-model:open="showTrackModal"
      title="物流轨迹"
      :width="700"
      :footer="null"
      class="track-modal"
    >
      <div class="track-content" v-if="currentLogistics">
        <div class="track-header">
          <div class="track-info">
            <span class="track-label">物流单号：</span>
            <span class="track-value">{{ currentLogistics.logisticsID }}</span>
          </div>
          <a-tag :color="getStatusColor(currentLogistics.status)">
            {{ currentLogistics.status }}
          </a-tag>
        </div>
        <a-timeline class="track-timeline">
          <a-timeline-item color="green">
            <div class="timeline-item">
              <div class="timeline-title">已签收</div>
              <div class="timeline-time">{{ currentLogistics.endTime }} 14:30</div>
              <div class="timeline-desc">您的包裹已由 {{ currentLogistics.receiver }} 签收</div>
            </div>
          </a-timeline-item>
          <a-timeline-item color="blue">
            <div class="timeline-item">
              <div class="timeline-title">派送中</div>
              <div class="timeline-time">{{ currentLogistics.endTime }} 09:00</div>
              <div class="timeline-desc">快递员正在为您派送</div>
            </div>
          </a-timeline-item>
          <a-timeline-item color="blue">
            <div class="timeline-item">
              <div class="timeline-title">到达目的地</div>
              <div class="timeline-time">{{ currentLogistics.endTime }} 06:00</div>
              <div class="timeline-desc">已到达目的地网点</div>
            </div>
          </a-timeline-item>
          <a-timeline-item color="gray">
            <div class="timeline-item">
              <div class="timeline-title">运输中</div>
              <div class="timeline-time">{{ currentLogistics.startTime }} 18:00</div>
              <div class="timeline-desc">包裹正在运往目的地</div>
            </div>
          </a-timeline-item>
          <a-timeline-item color="gray">
            <div class="timeline-item">
              <div class="timeline-title">已揽收</div>
              <div class="timeline-time">{{ currentLogistics.startTime }} 10:00</div>
              <div class="timeline-desc">快递员已揽收包裹</div>
            </div>
          </a-timeline-item>
        </a-timeline>
      </div>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { listLogisticsData, type LogisticsData } from '../../services/supply.service'

const router = useRouter()

const columns = [
  { title: '物流ID', dataIndex: 'logisticsID', key: 'logisticsID', width: 130 },
  { title: '订单ID', dataIndex: 'orderID', key: 'orderID', width: 130 },
  { title: '物流商', dataIndex: 'carrier', key: 'carrier', width: 120 },
  { title: '出发时间', dataIndex: 'startTime', key: 'startTime', width: 130 },
  { title: '送达时间', dataIndex: 'endTime', key: 'endTime', width: 130 },
  { title: '签收人', dataIndex: 'receiver', width: 100 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 100 },
  { title: '操作', key: 'action', fixed: 'right', width: 140 }
]

const tableData = ref<any[]>([])
const loading = ref(false)

const searchForm = ref({
  logisticsID: '',
  orderID: '',
  carrier: undefined as string | undefined,
  status: undefined as string | undefined
})

const showTrackModal = ref(false)
const currentLogistics = ref<any>(null)
const refreshing = ref(false)

async function fetchData() {
  loading.value = true
  try {
    const response = await listLogisticsData()
    if (response.code === 0 && response.data) {
      tableData.value = response.data.map((item: LogisticsData) => ({
        ...item,
        startTime: item.departureTime,
        endTime: item.arrivalTime
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
const inTransitCount = computed(() => tableData.value.filter(item => item.status === '运输中' || item.status === 'IN_TRANSIT').length)
const deliveredCount = computed(() => tableData.value.filter(item => item.status === '已送达' || item.status === 'DELIVERED').length)
const pendingCount = computed(() => tableData.value.filter(item => item.status === '待发货' || item.status === 'PENDING').length)

const filteredData = computed(() => {
  let result = [...tableData.value]
  if (searchForm.value.logisticsID) {
    result = result.filter(item => item.logisticsID?.toLowerCase().includes(searchForm.value.logisticsID.toLowerCase()))
  }
  if (searchForm.value.orderID) {
    result = result.filter(item => item.orderID?.toLowerCase().includes(searchForm.value.orderID.toLowerCase()))
  }
  if (searchForm.value.carrier) {
    result = result.filter(item => item.carrier === searchForm.value.carrier)
  }
  if (searchForm.value.status) {
    result = result.filter(item => item.status === searchForm.value.status)
  }
  return result
})

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    '待发货': 'orange',
    '运输中': 'blue',
    '已送达': 'green',
    'PENDING': 'orange',
    'IN_TRANSIT': 'blue',
    'DELIVERED': 'green'
  }
  return colors[status] || 'default'
}

function getStatusIcon(status: string) {
  const icons: Record<string, string> = {
    '待发货': 'ClockCircleOutlined',
    '运输中': 'CarOutlined',
    '已送达': 'CheckCircleOutlined'
  }
  return icons[status] || 'TagOutlined'
}

function handleSearch() {
  message.success(`找到 ${filteredData.value.length} 条记录`)
}

function handleReset() {
  searchForm.value = {
    logisticsID: '',
    orderID: '',
    carrier: undefined,
    status: undefined
  }
  fetchData()
}

async function handleRefresh() {
  refreshing.value = true
  try {
    await fetchData()
    message.success('物流状态已刷新')
  } catch (error: any) {
    message.error('刷新失败')
  } finally {
    setTimeout(() => {
      refreshing.value = false
    }, 500)
  }
}

function handleExport() {
  message.success('数据导出成功')
}

function viewTrack(record: any) {
  currentLogistics.value = record
  showTrackModal.value = true
}

function viewDetail(record: any) {
  router.push(`/supply/logistics/detail/${record.logisticsID}`)
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.logistics {
  padding: 0;
}

.page-header {
  margin-bottom: 24px;
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
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  color: #fff;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.35);
}

.title-text h2 {
  font-size: 22px;
  font-weight: 600;
  margin: 0;
  color: var(--text-color);
}

.subtitle {
  color: var(--text-color-secondary);
  margin: 4px 0 0;
  font-size: 14px;
}

.action-btn {
  height: 40px;
  border-radius: 10px;
  padding: 0 20px;
  font-weight: 500;
  color: #10b981;
  border-color: #10b981;
}

.action-btn:hover {
  color: #059669;
  border-color: #059669;
}

.stats-overview {
  display: flex;
  gap: 24px;
  padding: 20px 24px;
  background: linear-gradient(135deg, #f0fdf4 0%, #dcfce7 100%);
  border-radius: 12px;
  border: 1px solid #bbf7d0;
}

.stat-item {
  flex: 1;
  text-align: center;
  padding: 8px 0;
}

.stat-value {
  font-size: 28px;
  font-weight: 700;
  color: #1e293b;
  margin-bottom: 4px;
}

.stat-label {
  font-size: 13px;
  color: #64748b;
}

.status-transit { color: #3b82f6; }
.status-delivered { color: #10b981; }
.status-pending { color: #f59e0b; }

.search-card {
  border-radius: 12px;
  margin-bottom: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.search-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 16px;
  border-bottom: 1px solid #f0f0f0;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.filter-icon-wrapper {
  width: 36px;
  height: 36px;
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 16px;
}

.header-info {
  display: flex;
  flex-direction: column;
}

.header-title {
  font-size: 15px;
  font-weight: 600;
  color: var(--text-color);
}

.header-desc {
  font-size: 12px;
  color: var(--text-color-secondary);
}

.filter-tag {
  display: flex;
  align-items: center;
  gap: 6px;
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
  color: #10b981;
}

.filter-input,
.filter-select {
  width: 100%;
  height: 38px;
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

.dot-pending { background: #f59e0b; }
.dot-transit { background: #3b82f6; }
.dot-delivered { background: #10b981; }

.search-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 20px;
  padding-top: 16px;
  border-top: 1px solid #f0f0f0;
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

.reset-btn {
  height: 36px;
  border-radius: 8px;
}

.search-btn {
  height: 36px;
  border-radius: 8px;
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  border: none;
}

.table-card {
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
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
  color: var(--text-color);
}

.title-icon {
  color: #10b981;
}

.record-count {
  display: flex;
  align-items: center;
  gap: 6px;
}

.toolbar-btn {
  height: 34px;
  border-radius: 8px;
  font-size: 13px;
}

.export-btn {
  color: #10b981;
  border-color: #10b981;
}

.custom-table :deep(.ant-table-thead > tr > th) {
  background: linear-gradient(135deg, #f0fdf4 0%, #dcfce7 100%);
  font-weight: 600;
  color: #374151;
}

.custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #f0fdf4;
}

.logistics-id-cell {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
  color: #10b981;
}

.logistics-icon {
  font-size: 14px;
}

.order-tag {
  display: flex;
  align-items: center;
  gap: 4px;
}

.carrier-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.carrier-icon {
  color: #6366f1;
}

.status-tag {
  display: flex;
  align-items: center;
  gap: 4px;
}

.time-cell {
  display: flex;
  align-items: center;
  gap: 6px;
  color: #6b7280;
}

.time-icon {
  font-size: 12px;
}

.time-icon.success {
  color: #10b981;
}

.action-cell {
  display: flex;
  gap: 4px;
}

.action-btn {
  padding: 0 8px;
  height: 28px;
  display: flex;
  align-items: center;
  gap: 4px;
}

.track-btn {
  color: #10b981;
}

.detail-btn {
  color: #6366f1;
}

.track-content {
  padding: 16px 0;
}

.track-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid #f0f0f0;
}

.track-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.track-label {
  color: #6b7280;
}

.track-value {
  font-weight: 600;
  font-size: 16px;
}

.track-timeline {
  padding-left: 8px;
}

.timeline-item {
  padding-left: 8px;
}

.timeline-title {
  font-weight: 600;
  color: #1e293b;
  margin-bottom: 4px;
}

.timeline-time {
  font-size: 12px;
  color: #6b7280;
  margin-bottom: 4px;
}

.timeline-desc {
  font-size: 13px;
  color: #475569;
}

/* 深色主题 */
[data-theme='dark'] .title-icon {
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.45);
}

[data-theme='dark'] .stats-overview {
  background: linear-gradient(135deg, #1e293b 0%, #0f172a 100%);
  border-color: #334155;
}

[data-theme='dark'] .stat-value {
  color: #f1f5f9;
}

[data-theme='dark'] .stat-label {
  color: #94a3b8;
}

[data-theme='dark'] .search-header {
  border-bottom-color: #374151;
}

[data-theme='dark'] .filter-label {
  color: #e5e7eb;
}

[data-theme='dark'] .search-actions {
  border-top-color: #374151;
}

[data-theme='dark'] .custom-table :deep(.ant-table-thead > tr > th) {
  background: linear-gradient(135deg, #1e293b 0%, #0f172a 100%);
  color: #e5e7eb;
}

[data-theme='dark'] .custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #1e293b;
}

[data-theme='dark'] .track-header {
  border-bottom-color: #374151;
}

[data-theme='dark'] .timeline-title {
  color: #f1f5f9;
}

[data-theme='dark'] .timeline-desc {
  color: #94a3b8;
}
</style>
