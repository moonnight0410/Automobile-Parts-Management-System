<!--
  采购订单页面
  设计：与Dashboard风格统一
-->

<template>
  <div class="supply-orders">
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <ShoppingCartOutlined />
          </div>
          <div class="title-text">
            <h2>采购订单</h2>
            <p class="subtitle">管理供应链采购订单与交付跟踪</p>
          </div>
        </div>
        <a-button type="primary" class="create-btn" @click="handleCreate">
          <template #icon><PlusOutlined /></template>
          创建订单
        </a-button>
      </div>
      <div class="stats-overview">
        <div class="stat-item">
          <div class="stat-value">{{ totalCount }}</div>
          <div class="stat-label">总订单数</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-pending">{{ pendingCount }}</div>
          <div class="stat-label">待处理</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-shipping">{{ shippingCount }}</div>
          <div class="stat-label">运输中</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-completed">{{ completedCount }}</div>
          <div class="stat-label">已完成</div>
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
          <a-tag color="purple" class="filter-tag">
            <template #icon><SearchOutlined /></template>
            订单查询
          </a-tag>
        </div>
      </div>
      <a-form layout="inline" :model="searchForm" class="search-form">
        <a-row :gutter="[20, 20]" style="width: 100%">
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <IdcardOutlined class="label-icon" />
                <span>订单ID</span>
              </div>
              <a-input
                v-model:value="searchForm.orderID"
                placeholder="请输入订单ID"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <ShoppingOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <UserOutlined class="label-icon" />
                <span>采购方</span>
              </div>
              <a-input
                v-model:value="searchForm.buyer"
                placeholder="请输入采购方名称"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <UserOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <TagOutlined class="label-icon" />
                <span>订单状态</span>
              </div>
              <a-select
                v-model:value="searchForm.status"
                placeholder="请选择状态"
                allow-clear
                class="filter-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
              >
                <a-select-option value="待处理">
                  <span class="status-option"><span class="status-dot dot-pending"></span>待处理</span>
                </a-select-option>
                <a-select-option value="运输中">
                  <span class="status-option"><span class="status-dot dot-shipping"></span>运输中</span>
                </a-select-option>
                <a-select-option value="已签收">
                  <span class="status-option"><span class="status-dot dot-completed"></span>已签收</span>
                </a-select-option>
              </a-select>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <CalendarOutlined class="label-icon" />
                <span>创建时间</span>
              </div>
              <div class="date-range-container">
                <a-date-picker
                  v-model:value="searchForm.startDate"
                  placeholder="开始日期"
                  class="date-input"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                />
                <span class="date-separator">至</span>
                <a-date-picker
                  v-model:value="searchForm.endDate"
                  placeholder="结束日期"
                  class="date-input"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                />
              </div>
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
            共 {{ filteredData.length }} 条记录
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
              导出数据
            </a-button>
          </a-space>
        </div>
      </div>
      <a-table
        :columns="columns"
        :data-source="filteredData"
        row-key="orderID"
        class="custom-table"
        :scroll="{ x: 1200 }"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'orderID'">
            <div class="order-id-cell">
              <ShoppingOutlined class="order-icon" />
              <span>{{ record.orderID }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'buyer'">
            <div class="user-cell">
              <a-avatar :style="{ backgroundColor: '#6366f1' }" size="small">
                {{ record.buyer.charAt(0) }}
              </a-avatar>
              <span class="user-name">{{ record.buyer }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'seller'">
            <div class="user-cell">
              <a-avatar :style="{ backgroundColor: '#8b5cf6' }" size="small">
                {{ record.seller.charAt(0) }}
              </a-avatar>
              <span class="user-name">{{ record.seller }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'quantity'">
            <a-tag color="cyan" class="quantity-tag">
              <NumberOutlined />
              {{ record.quantity }} 件
            </a-tag>
          </template>
          <template v-else-if="column.key === 'status'">
            <a-tag :color="getStatusColor(record.status)" class="status-tag">
              <component :is="getStatusIcon(record.status)" />
              {{ record.status }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'createTime'">
            <div class="time-cell">
              <ClockCircleOutlined class="time-icon" />
              <span>{{ record.createTime }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'action'">
            <a-space class="action-cell">
              <a-button type="link" size="small" @click="viewDetail(record)" class="action-btn detail-btn">
                <EyeOutlined />
                详情
              </a-button>
              <a-button type="link" size="small" @click="trackLogistics(record)" class="action-btn track-btn">
                <EnvironmentOutlined />
                物流
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'

const router = useRouter()

const columns = [
  { title: '订单ID', dataIndex: 'orderID', key: 'orderID', width: 140 },
  { title: '采购方', dataIndex: 'buyer', key: 'buyer', width: 120 },
  { title: '销售方', dataIndex: 'seller', key: 'seller', width: 120 },
  { title: '零部件ID', dataIndex: 'partID', width: 120 },
  { title: '数量', dataIndex: 'quantity', key: 'quantity', width: 100 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 100 },
  { title: '创建时间', dataIndex: 'createTime', key: 'createTime', width: 140 },
  { title: '操作', key: 'action', fixed: 'right', width: 140 }
]

const mockData = ref([
  { orderID: 'ORDER-001', buyer: '车企A', seller: '制造商B', partID: 'PART-001', quantity: 100, status: '已签收', createTime: '2024-01-15' },
  { orderID: 'ORDER-002', buyer: '车企A', seller: '制造商C', partID: 'PART-002', quantity: 50, status: '运输中', createTime: '2024-01-16' },
  { orderID: 'ORDER-003', buyer: '车企B', seller: '制造商B', partID: 'PART-003', quantity: 200, status: '待处理', createTime: '2024-01-17' },
  { orderID: 'ORDER-004', buyer: '车企A', seller: '制造商D', partID: 'PART-004', quantity: 75, status: '已签收', createTime: '2024-01-18' },
  { orderID: 'ORDER-005', buyer: '车企C', seller: '制造商B', partID: 'PART-005', quantity: 150, status: '运输中', createTime: '2024-01-19' }
])

const searchForm = ref({
  orderID: '',
  buyer: '',
  status: undefined as string | undefined,
  startDate: null as any,
  endDate: null as any
})

const showCreateModal = ref(false)
const refreshing = ref(false)

const totalCount = computed(() => mockData.value.length)
const pendingCount = computed(() => mockData.value.filter(item => item.status === '待处理').length)
const shippingCount = computed(() => mockData.value.filter(item => item.status === '运输中').length)
const completedCount = computed(() => mockData.value.filter(item => item.status === '已签收').length)

const filteredData = computed(() => {
  let result = [...mockData.value]
  if (searchForm.value.orderID) {
    result = result.filter(item => item.orderID.toLowerCase().includes(searchForm.value.orderID.toLowerCase()))
  }
  if (searchForm.value.buyer) {
    result = result.filter(item => item.buyer.toLowerCase().includes(searchForm.value.buyer.toLowerCase()))
  }
  if (searchForm.value.status) {
    result = result.filter(item => item.status === searchForm.value.status)
  }
  return result
})

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    '待处理': 'orange',
    '运输中': 'blue',
    '已签收': 'green'
  }
  return colors[status] || 'default'
}

function getStatusIcon(status: string) {
  const icons: Record<string, string> = {
    '待处理': 'ClockCircleOutlined',
    '运输中': 'CarOutlined',
    '已签收': 'CheckCircleOutlined'
  }
  return icons[status] || 'TagOutlined'
}

function handleSearch() {
  message.success(`找到 ${filteredData.value.length} 条记录`)
}

function handleReset() {
  searchForm.value = {
    orderID: '',
    buyer: '',
    status: undefined,
    dateRange: null
  }
}

function handleRefresh() {
  refreshing.value = true
  setTimeout(() => {
    refreshing.value = false
    message.success('数据已刷新')
  }, 1000)
}

function handleExport() {
  message.success('数据导出成功')
}

function viewDetail(record: any) {
  message.info(`查看订单 ${record.orderID} 详情`)
}

function trackLogistics(record: any) {
  message.info(`追踪订单 ${record.orderID} 物流`)
}

function handleCreate() {
  router.push('/supply/order/create')
}

function handleDateChange() {
  if (searchForm.value.dateRange && searchForm.value.dateRange.length === 2) {
    datePickerOpen.value = false
  }
}
</script>

<style scoped>
.supply-orders {
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
  width: 100%;
  gap: 16px;
}

.header-title {
  display: flex;
  align-items: center;
  gap: 16px;
  flex: 1;
  min-width: 0;
}

.title-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  color: #fff;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.35);
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

.create-btn {
  height: 40px;
  border-radius: 10px;
  padding: 0 20px;
  font-weight: 500;
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  border: none;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.35);
  flex-shrink: 0;
  white-space: nowrap;
}

.create-btn:hover {
  background: linear-gradient(135deg, #4f46e5 0%, #4338ca 100%);
  box-shadow: 0 6px 16px rgba(99, 102, 241, 0.45);
}

.stats-overview {
  display: flex;
  gap: 24px;
  padding: 20px 24px;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border-radius: 12px;
  border: 1px solid #e2e8f0;
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

.status-pending { color: #f59e0b; }
.status-shipping { color: #3b82f6; }
.status-completed { color: #10b981; }

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
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
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
  color: #6366f1;
}

.filter-input,
.filter-select,
.filter-date {
  width: 100%;
  height: 38px;
}

.date-range-container {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
}

.date-input {
  flex: 1;
  height: 38px;
}

.date-separator {
  color: #9ca3af;
  font-size: 14px;
  white-space: nowrap;
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
.dot-shipping { background: #3b82f6; }
.dot-completed { background: #10b981; }

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
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
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
  color: #6366f1;
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

.refresh-btn {
  color: #6366f1;
  border-color: #6366f1;
}

.export-btn {
  color: #10b981;
  border-color: #10b981;
}

.custom-table :deep(.ant-table-thead > tr > th) {
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  font-weight: 600;
  color: #374151;
}

.custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #f8fafc;
}

.order-id-cell {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
  color: #6366f1;
}

.order-icon {
  font-size: 14px;
}

.user-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.user-name {
  font-weight: 500;
}

.quantity-tag {
  display: flex;
  align-items: center;
  gap: 4px;
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

.detail-btn {
  color: #6366f1;
}

.track-btn {
  color: #10b981;
}

/* 深色主题 */
[data-theme='dark'] .title-icon {
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.45);
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
</style>

<style>
.ant-picker-dropdown {
  max-width: 600px !important;
  max-height: 400px !important;
  overflow: hidden !important;
  border-radius: 12px !important;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.15) !important;
  background: #ffffff !important;
  padding: 0 !important;
  animation: datePickerFadeIn 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes datePickerFadeIn {
  from {
    opacity: 0;
    transform: translateY(-8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.ant-picker-range-wrapper {
  max-width: 100% !important;
}

.ant-picker-panel-container {
  max-width: 100% !important;
}

.ant-picker-panel {
  max-width: 280px !important;
  border: none !important;
  background: #ffffff !important;
}

.ant-picker-header {
  border-bottom: 1px solid #f0f0f0 !important;
  padding: 12px 16px !important;
  background: linear-gradient(135deg, #fafbfc 0%, #f5f7fa 100%) !important;
  display: flex !important;
  align-items: center !important;
  justify-content: space-between !important;
  flex-wrap: nowrap !important;
  gap: 8px !important;
  min-height: 60px !important;
  box-sizing: border-box !important;
}

.ant-picker-header > button {
  flex-shrink: 0 !important;
  min-width: 36px !important;
  width: 36px !important;
  height: 36px !important;
  border-radius: 10px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  transition: all 0.2s ease !important;
  color: #000000 !important;
  cursor: pointer !important;
}

.ant-picker-header > button:hover {
  background: #e0f2fe !important;
  color: #1890ff !important;
}

.ant-picker-header-view {
  font-weight: 600 !important;
  font-size: 15px !important;
  color: #000000 !important;
  flex: 1 !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  gap: 8px !important;
  white-space: nowrap !important;
}

.ant-picker-header-view button {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 15px !important;
  transition: all 0.2s ease !important;
  padding: 6px 12px !important;
  border-radius: 6px !important;
  min-width: auto !important;
  width: auto !important;
  height: auto !important;
  white-space: nowrap !important;
  display: inline-flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.ant-picker-header-view button:hover {
  color: #1890ff !important;
  background: #e0f2fe !important;
}

.ant-picker-body {
  padding: 12px 16px 16px !important;
  background: #ffffff !important;
}

.ant-picker-content {
  border-collapse: separate !important;
  border-spacing: 4px !important;
}

.ant-picker-content th {
  font-size: 12px !important;
  font-weight: 600 !important;
  color: #9ca3af !important;
  padding: 8px 0 !important;
  text-transform: uppercase !important;
  letter-spacing: 0.5px !important;
}

.ant-picker-cell {
  padding: 4px 0 !important;
  transition: all 0.2s ease !important;
}

.ant-picker-cell-inner {
  width: 36px !important;
  height: 36px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  border-radius: 10px !important;
  font-size: 14px !important;
  font-weight: 500 !important;
  color: #374151 !important;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1) !important;
  cursor: pointer !important;
  position: relative !important;
  overflow: hidden !important;
}

.ant-picker-cell-inner::before {
  content: '' !important;
  position: absolute !important;
  top: 50% !important;
  left: 50% !important;
  width: 0 !important;
  height: 0 !important;
  border-radius: 50% !important;
  background: rgba(24, 144, 255, 0.1) !important;
  transform: translate(-50%, -50%) !important;
  transition: all 0.3s ease !important;
}

.ant-picker-cell:hover .ant-picker-cell-inner::before {
  width: 100% !important;
  height: 100% !important;
  border-radius: 10px !important;
}

.ant-picker-cell:hover .ant-picker-cell-inner {
  background: #e0f2fe !important;
  color: #1890ff !important;
  transform: scale(1.05) !important;
}

.ant-picker-cell-selected .ant-picker-cell-inner {
  background: linear-gradient(135deg, #1890ff 0%, #096dd9 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3) !important;
}

.ant-picker-cell-selected:hover .ant-picker-cell-inner {
  background: linear-gradient(135deg, #40a9ff 0%, #1890ff 100%) !important;
  box-shadow: 0 6px 16px rgba(24, 144, 255, 0.4) !important;
}

.ant-picker-cell-today .ant-picker-cell-inner {
  border: 2px solid #1890ff !important;
  font-weight: 600 !important;
}

.ant-picker-cell-today:hover .ant-picker-cell-inner {
  border-color: #40a9ff !important;
}

.ant-picker-cell-range-start .ant-picker-cell-inner,
.ant-picker-cell-range-end .ant-picker-cell-inner {
  background: linear-gradient(135deg, #1890ff 0%, #096dd9 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3) !important;
}

.ant-picker-cell-in-range .ant-picker-cell-inner {
  background: #e0f2fe !important;
  color: #1890ff !important;
}

.ant-picker-cell-in-range:hover .ant-picker-cell-inner {
  background: #bae7ff !important;
}

.ant-picker-cell-disabled .ant-picker-cell-inner {
  color: #d1d5db !important;
  background: #f9fafb !important;
  cursor: not-allowed !important;
}

.ant-picker-cell-disabled:hover .ant-picker-cell-inner {
  background: #f9fafb !important;
  transform: none !important;
}

.ant-picker-footer {
  border-top: 1px solid #f0f0f0 !important;
  padding: 12px 16px !important;
  background: #fafbfc !important;
  display: flex !important;
  justify-content: flex-end !important;
  gap: 8px !important;
}

.ant-picker-footer .ant-btn {
  border-radius: 8px !important;
  height: 32px !important;
  font-size: 13px !important;
  font-weight: 500 !important;
  transition: all 0.2s ease !important;
}

.ant-picker-footer .ant-btn-primary {
  background: linear-gradient(135deg, #1890ff 0%, #096dd9 100%) !important;
  border: none !important;
  box-shadow: 0 2px 8px rgba(24, 144, 255, 0.25) !important;
}

.ant-picker-footer .ant-btn-primary:hover {
  background: linear-gradient(135deg, #40a9ff 0%, #1890ff 100%) !important;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.35) !important;
  transform: translateY(-1px) !important;
}

.ant-picker-year-panel .ant-picker-cell-inner,
.ant-picker-month-panel .ant-picker-cell-inner {
  width: 60px !important;
  height: 36px !important;
  border-radius: 8px !important;
}

.ant-picker-decade-panel .ant-picker-cell-inner {
  width: 80px !important;
  height: 36px !important;
  border-radius: 8px !important;
}

@media (max-width: 768px) {
  .ant-picker-dropdown {
    max-width: 320px !important;
    max-height: 350px !important;
  }
  
  .ant-picker-panel {
    max-width: 280px !important;
  }
  
  .ant-picker-cell-inner {
    width: 32px !important;
    height: 32px !important;
    font-size: 13px !important;
  }
  
  .ant-picker-header {
    padding: 10px 12px !important;
    min-height: 50px !important;
  }
  
  .ant-picker-header > button {
    width: 32px !important;
    height: 32px !important;
    min-width: 32px !important;
  }
  
  .ant-picker-header-view {
    font-size: 14px !important;
  }
  
  .ant-picker-header-view button {
    font-size: 14px !important;
    padding: 4px 8px !important;
  }
}
</style>
