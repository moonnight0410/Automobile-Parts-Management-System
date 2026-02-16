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
        <a-button type="primary" class="create-btn" @click="showCreateModal = true">
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
              <a-range-picker
                v-model:value="searchForm.dateRange"
                class="filter-date"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
              />
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

    <a-modal
      v-model:open="showCreateModal"
      title="创建采购订单"
      :width="600"
      @ok="handleCreate"
      @cancel="showCreateModal = false"
      class="create-modal"
    >
      <a-form :model="createForm" layout="vertical" class="create-form">
        <a-form-item label="采购方" required>
          <a-input v-model:value="createForm.buyer" placeholder="请输入采购方名称" />
        </a-form-item>
        <a-form-item label="销售方" required>
          <a-input v-model:value="createForm.seller" placeholder="请输入销售方名称" />
        </a-form-item>
        <a-form-item label="零部件ID" required>
          <a-input v-model:value="createForm.partID" placeholder="请输入零部件ID" />
        </a-form-item>
        <a-form-item label="数量" required>
          <a-input-number v-model:value="createForm.quantity" :min="1" style="width: 100%" placeholder="请输入数量" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { message } from 'ant-design-vue'

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
  dateRange: null as any
})

const createForm = ref({
  buyer: '',
  seller: '',
  partID: '',
  quantity: 1
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
  if (!createForm.value.buyer || !createForm.value.seller || !createForm.value.partID) {
    message.error('请填写完整信息')
    return
  }
  mockData.value.unshift({
    orderID: `ORDER-${String(mockData.value.length + 1).padStart(3, '0')}`,
    buyer: createForm.value.buyer,
    seller: createForm.value.seller,
    partID: createForm.value.partID,
    quantity: createForm.value.quantity,
    status: '待处理',
    createTime: new Date().toISOString().split('T')[0]
  })
  showCreateModal.value = false
  createForm.value = { buyer: '', seller: '', partID: '', quantity: 1 }
  message.success('订单创建成功')
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
}

.header-title {
  display: flex;
  align-items: center;
  gap: 16px;
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
