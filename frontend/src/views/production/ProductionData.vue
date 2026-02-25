<!--
  生产数据页面
  设计：与Dashboard风格统一
-->

<template>
  <div class="production-data">
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <SettingOutlined />
          </div>
          <div class="title-text">
            <h2>生产数据</h2>
            <p class="subtitle">管理生产流程数据与批次记录</p>
          </div>
        </div>
        <a-button type="primary" class="create-btn" @click="goToCreate">
          <template #icon><PlusOutlined /></template>
          录入生产数据
        </a-button>
      </div>
      <div class="stats-overview">
        <div class="stat-item">
          <div class="stat-value">{{ totalCount }}</div>
          <div class="stat-label">生产批次</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-production">{{ todayCount }}</div>
          <div class="stat-label">今日生产</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-line">{{ activeLines }}</div>
          <div class="stat-label">活跃产线</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-output">{{ totalOutput }}</div>
          <div class="stat-label">总产量(件)</div>
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
            <span class="header-desc">支持批次号、产线等多维度查询</span>
          </div>
        </div>
        <div class="header-right">
          <a-tag color="orange" class="filter-tag">
            <template #icon><SettingOutlined /></template>
            生产查询
          </a-tag>
        </div>
      </div>
      <a-form layout="inline" :model="searchForm" class="search-form">
        <a-row :gutter="[20, 20]" style="width: 100%">
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <NumberOutlined class="label-icon" />
                <span>生产ID</span>
              </div>
              <a-input
                v-model:value="searchForm.productionID"
                placeholder="请输入生产ID"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <SettingOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <AppstoreOutlined class="label-icon" />
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
                <TagOutlined class="label-icon" />
                <span>批次号</span>
              </div>
              <a-input
                v-model:value="searchForm.batchNo"
                placeholder="请输入批次号"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <NumberOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <TeamOutlined class="label-icon" />
                <span>生产线</span>
              </div>
              <a-select
                v-model:value="searchForm.productionLine"
                placeholder="请选择生产线"
                allow-clear
                class="filter-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
              >
                <a-select-option value="A线">A线</a-select-option>
                <a-select-option value="B线">B线</a-select-option>
                <a-select-option value="C线">C线</a-select-option>
              </a-select>
            </div>
          </a-col>
        </a-row>
        <div class="search-actions">
          <div class="action-left">
            <span class="action-hint">
              <InfoCircleOutlined />
              可通过批次号快速定位生产记录
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
            生产记录
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
        row-key="productionID"
        class="custom-table"
        :scroll="{ x: 1000 }"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'productionID'">
            <div class="production-id-cell">
              <SettingOutlined class="production-icon" />
              <span>{{ record.productionID }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'partID'">
            <a-tag color="purple" class="part-tag">
              <AppstoreOutlined />
              {{ record.partID }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'batchNo'">
            <div class="batch-cell">
              <TagOutlined class="batch-icon" />
              <span>{{ record.batchNo }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'productionLine'">
            <a-tag :color="getLineColor(record.productionLine)" class="line-tag">
              <TeamOutlined />
              {{ record.productionLine }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'operator'">
            <div class="operator-cell">
              <a-avatar :style="{ backgroundColor: '#f59e0b' }" size="small">
                {{ record.operator.charAt(0) }}
              </a-avatar>
              <span class="operator-name">{{ record.operator }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'finishTime'">
            <div class="time-cell">
              <ClockCircleOutlined class="time-icon" />
              <span>{{ record.finishTime }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'action'">
            <a-space class="action-cell">
              <a-button type="link" size="small" @click="viewDetail(record)" class="action-btn detail-btn">
                <EyeOutlined />
                详情
              </a-button>
              <a-button type="link" size="small" @click="viewQuality(record)" class="action-btn quality-btn">
                <SafetyCertificateOutlined />
                质检
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <a-modal
      v-model:open="showCreateModal"
      title="录入生产数据"
      :width="600"
      @ok="handleCreate"
      @cancel="showCreateModal = false"
      class="create-modal"
    >
      <a-form :model="createForm" layout="vertical" class="create-form">
        <a-form-item label="零部件ID" required>
          <a-input v-model:value="createForm.partID" placeholder="请输入零部件ID" />
        </a-form-item>
        <a-form-item label="批次号" required>
          <a-input v-model:value="createForm.batchNo" placeholder="请输入批次号" />
        </a-form-item>
        <a-form-item label="生产线" required>
          <a-select v-model:value="createForm.productionLine" placeholder="请选择生产线">
            <a-select-option value="A线">A线</a-select-option>
            <a-select-option value="B线">B线</a-select-option>
            <a-select-option value="C线">C线</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="操作员" required>
          <a-input v-model:value="createForm.operator" placeholder="请输入操作员姓名" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { listProductionData, type ProductionData } from '../../services/production.service'

const router = useRouter()

const columns = [
  { title: '生产ID', dataIndex: 'recordID', key: 'recordID', width: 130 },
  { title: '零部件ID', dataIndex: 'partID', key: 'partID', width: 130 },
  { title: '批次号', dataIndex: 'batchNo', key: 'batchNo', width: 130 },
  { title: '生产线', dataIndex: 'productionLine', key: 'productionLine', width: 100 },
  { title: '操作员', dataIndex: 'operator', key: 'operator', width: 100 },
  { title: '完成时间', dataIndex: 'operationTime', key: 'operationTime', width: 140 },
  { title: '操作', key: 'action', fixed: 'right', width: 140 }
]

const tableData = ref<any[]>([])
const loading = ref(false)

const searchForm = ref({
  productionID: '',
  partID: '',
  batchNo: '',
  productionLine: undefined as string | undefined
})

const createForm = ref({
  partID: '',
  batchNo: '',
  productionLine: undefined as string | undefined,
  operator: ''
})

const showCreateModal = ref(false)
const refreshing = ref(false)

const totalCount = computed(() => tableData.value.length)
const todayCount = computed(() => {
  const today = new Date().toISOString().split('T')[0]
  return tableData.value.filter(item => item.operationTime?.startsWith(today)).length
})
const activeLines = computed(() => {
  const lines = new Set(tableData.value.map(item => item.productionLine))
  return lines.size
})
const totalOutput = computed(() => tableData.value.length * 100)

async function fetchData() {
  loading.value = true
  try {
    const response = await listProductionData()
    if (response.code === 0 && response.data) {
      tableData.value = response.data.map((item: ProductionData) => ({
        ...item,
        productionID: item.recordID,
        finishTime: item.operationTime
      }))
    }
  } catch (error: any) {
    message.error(error.message || '获取数据失败')
    tableData.value = []
  } finally {
    loading.value = false
  }
}

const filteredData = computed(() => {
  let result = [...tableData.value]
  if (searchForm.value.productionID) {
    result = result.filter(item => item.recordID?.toLowerCase().includes(searchForm.value.productionID.toLowerCase()))
  }
  if (searchForm.value.partID) {
    result = result.filter(item => item.partID?.toLowerCase().includes(searchForm.value.partID.toLowerCase()))
  }
  if (searchForm.value.batchNo) {
    result = result.filter(item => item.batchNo?.toLowerCase().includes(searchForm.value.batchNo.toLowerCase()))
  }
  if (searchForm.value.productionLine) {
    result = result.filter(item => item.productionLine === searchForm.value.productionLine)
  }
  return result
})

function getLineColor(line: string) {
  const colors: Record<string, string> = {
    'A线': 'blue',
    'B线': 'green',
    'C线': 'purple'
  }
  return colors[line] || 'default'
}

function handleSearch() {
  message.success(`找到 ${filteredData.value.length} 条记录`)
}

function handleReset() {
  searchForm.value = {
    productionID: '',
    partID: '',
    batchNo: '',
    productionLine: undefined
  }
  fetchData()
}

async function handleRefresh() {
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

function handleExport() {
  message.success('数据导出成功')
}

function viewDetail(record: any) {
  message.info(`查看生产记录 ${record.recordID} 详情`)
}

function viewQuality(record: any) {
  message.info(`查看生产记录 ${record.recordID} 的质检信息`)
}

function goToCreate() {
  router.push('/production/data/create')
}

function handleCreate() {
  if (!createForm.value.partID || !createForm.value.batchNo || !createForm.value.productionLine || !createForm.value.operator) {
    message.error('请填写完整信息')
    return
  }
  showCreateModal.value = false
  createForm.value = { partID: '', batchNo: '', productionLine: undefined, operator: '' }
  message.success('生产数据录入成功')
  fetchData()
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.production-data {
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
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  color: #fff;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.35);
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
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  border: none;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.35);
}

.create-btn:hover {
  background: linear-gradient(135deg, #d97706 0%, #b45309 100%);
  box-shadow: 0 6px 16px rgba(245, 158, 11, 0.45);
}

.stats-overview {
  display: flex;
  gap: 24px;
  padding: 20px 24px;
  background: linear-gradient(135deg, #fffbeb 0%, #fef3c7 100%);
  border-radius: 12px;
  border: 1px solid #fde68a;
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

.status-production { color: #f59e0b; }
.status-line { color: #3b82f6; }
.status-output { color: #10b981; }

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
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
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
  color: #f59e0b;
}

.filter-input,
.filter-select {
  width: 100%;
  height: 38px;
}

.input-prefix-icon {
  color: #9ca3af;
}

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
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
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
  color: #f59e0b;
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
  color: #f59e0b;
  border-color: #f59e0b;
}

.export-btn {
  color: #10b981;
  border-color: #10b981;
}

.custom-table :deep(.ant-table-thead > tr > th) {
  background: linear-gradient(135deg, #fffbeb 0%, #fef3c7 100%);
  font-weight: 600;
  color: #374151;
}

.custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #fffbeb;
}

.production-id-cell {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
  color: #f59e0b;
}

.production-icon {
  font-size: 14px;
}

.part-tag {
  display: flex;
  align-items: center;
  gap: 4px;
}

.batch-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.batch-icon {
  color: #6366f1;
}

.line-tag {
  display: flex;
  align-items: center;
  gap: 4px;
}

.operator-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.operator-name {
  font-weight: 500;
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
  color: #f59e0b;
}

.quality-btn {
  color: #10b981;
}

/* 深色主题 */
[data-theme='dark'] .title-icon {
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.45);
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
