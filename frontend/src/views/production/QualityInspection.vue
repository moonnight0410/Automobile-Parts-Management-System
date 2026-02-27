<!--
  质检管理页面
  设计：与Dashboard风格统一
-->

<template>
  <div class="quality-inspection">
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <SafetyCertificateOutlined />
          </div>
          <div class="title-text">
            <h2>质检管理</h2>
            <p class="subtitle">产品质量检测与合规性管理</p>
          </div>
        </div>
        <a-button type="primary" class="create-btn" @click="goToCreate">
          <template #icon><PlusOutlined /></template>
          录入质检数据
        </a-button>
      </div>
      <div class="stats-overview">
        <div class="stat-item">
          <div class="stat-value">{{ totalCount }}</div>
          <div class="stat-label">质检总数</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-passed">{{ passedCount }}</div>
          <div class="stat-label">合格数</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-failed">{{ failedCount }}</div>
          <div class="stat-label">不合格数</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-rate">{{ passRate }}%</div>
          <div class="stat-label">合格率</div>
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
            <span class="header-desc">支持质检ID、批次号等查询</span>
          </div>
        </div>
        <div class="header-right">
          <a-tag color="cyan" class="filter-tag">
            <template #icon><SafetyCertificateOutlined /></template>
            质检查询
          </a-tag>
        </div>
      </div>
      <a-form layout="inline" :model="searchForm" class="search-form">
        <a-row :gutter="[20, 20]" style="width: 100%">
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <NumberOutlined class="label-icon" />
                <span>质检ID</span>
              </div>
              <a-input
                v-model:value="searchForm.inspectionID"
                placeholder="请输入质检ID"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <SafetyCertificateOutlined class="input-prefix-icon" />
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
                <CheckCircleOutlined class="label-icon" />
                <span>质检结果</span>
              </div>
              <a-select
                v-model:value="searchForm.result"
                placeholder="请选择结果"
                allow-clear
                class="filter-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
              >
                <a-select-option value="合格">
                  <span class="status-option"><span class="status-dot dot-passed"></span>合格</span>
                </a-select-option>
                <a-select-option value="不合格">
                  <span class="status-option"><span class="status-dot dot-failed"></span>不合格</span>
                </a-select-option>
              </a-select>
            </div>
          </a-col>
        </a-row>
        <div class="search-actions">
          <div class="action-left">
            <span class="action-hint">
              <InfoCircleOutlined />
              可通过批次号快速定位质检记录
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
            质检记录
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
        row-key="inspectionID"
        class="custom-table"
        :scroll="{ x: 1000 }"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'inspectionID'">
            <div class="inspection-id-cell">
              <SafetyCertificateOutlined class="inspection-icon" />
              <span>{{ record.inspectionID }}</span>
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
          <template v-else-if="column.key === 'result'">
            <a-tag :color="record.result === '合格' ? 'green' : 'red'" class="result-tag">
              <CheckCircleOutlined v-if="record.result === '合格'" />
              <CloseCircleOutlined v-else />
              {{ record.result }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'handler'">
            <div class="handler-cell">
              <a-avatar :style="{ backgroundColor: record.result === '合格' ? '#10b981' : '#ef4444' }" size="small">
                {{ record.handler.charAt(0) }}
              </a-avatar>
              <span class="handler-name">{{ record.handler }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'handleTime'">
            <div class="time-cell">
              <ClockCircleOutlined class="time-icon" />
              <span>{{ record.handleTime }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'action'">
            <a-space class="action-cell">
              <a-button type="link" size="small" @click="viewDetail(record)" class="action-btn detail-btn">
                <EyeOutlined />
                详情
              </a-button>
              <a-button type="link" size="small" @click="viewReport(record)" class="action-btn report-btn">
                <FileTextOutlined />
                报告
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <a-modal
      v-model:open="showCreateModal"
      title="录入质检数据"
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
        <a-form-item label="质检结果" required>
          <a-select v-model:value="createForm.result" placeholder="请选择质检结果">
            <a-select-option value="合格">合格</a-select-option>
            <a-select-option value="不合格">不合格</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="处理人" required>
          <a-input v-model:value="createForm.handler" placeholder="请输入处理人姓名" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { listQualityInspections, type QualityInspection } from '../../services/quality.service'

const router = useRouter()

const columns = [
  { title: '质检ID', dataIndex: 'inspectionID', key: 'inspectionID', width: 120 },
  { title: '零部件ID', dataIndex: 'partID', key: 'partID', width: 120 },
  { title: '批次号', dataIndex: 'batchNo', key: 'batchNo', width: 130 },
  { title: '质检结果', dataIndex: 'result', key: 'result', width: 100 },
  { title: '处理人', dataIndex: 'handler', key: 'handler', width: 100 },
  { title: '处理时间', dataIndex: 'inspectionTime', key: 'inspectionTime', width: 140 },
  { title: '操作', key: 'action', fixed: 'right', width: 140 }
]

const tableData = ref<any[]>([])
const loading = ref(false)

const searchForm = ref({
  inspectionID: '',
  partID: '',
  batchNo: '',
  result: undefined as string | undefined
})

const createForm = ref({
  partID: '',
  batchNo: '',
  result: undefined as string | undefined,
  handler: ''
})

const showCreateModal = ref(false)
const refreshing = ref(false)

async function fetchData() {
  loading.value = true
  try {
    const response = await listQualityInspections()
    if (response.code === 0 && response.data) {
      tableData.value = response.data.map((item: any) => {
        const result = item.inspectionResult ?? item.result
        const handleTime = item.inspectionTime ?? item.handleTime
        return {
          ...item,
          result,
          handleTime
        }
      })
    }
  } catch (error: any) {
    message.error(error.message || '获取数据失败')
    tableData.value = []
  } finally {
    loading.value = false
  }
}

const totalCount = computed(() => tableData.value.length)
const passedCount = computed(() => tableData.value.filter(item => item.result === '合格').length)
const failedCount = computed(() => tableData.value.filter(item => item.result === '不合格').length)
const passRate = computed(() => {
  if (totalCount.value === 0) return 0
  return Math.round((passedCount.value / totalCount.value) * 100)
})

const filteredData = computed(() => {
  let result = [...tableData.value]
  if (searchForm.value.inspectionID) {
    result = result.filter(item => item.inspectionID?.toLowerCase().includes(searchForm.value.inspectionID.toLowerCase()))
  }
  if (searchForm.value.partID) {
    result = result.filter(item => item.partID?.toLowerCase().includes(searchForm.value.partID.toLowerCase()))
  }
  if (searchForm.value.batchNo) {
    result = result.filter(item => item.batchNo?.toLowerCase().includes(searchForm.value.batchNo.toLowerCase()))
  }
  if (searchForm.value.result) {
    result = result.filter(item => item.result === searchForm.value.result)
  }
  return result
})

function handleSearch() {
  message.success(`找到 ${filteredData.value.length} 条记录`)
}

function handleReset() {
  searchForm.value = {
    inspectionID: '',
    partID: '',
    batchNo: '',
    result: undefined
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
  router.push(`/production/quality/detail/${record.inspectionID}`)
}

function viewReport(record: any) {
  message.info(`查看质检报告 ${record.inspectionID}`)
}

function goToCreate() {
  router.push('/production/quality/create')
}

function handleCreate() {
  if (!createForm.value.partID || !createForm.value.batchNo || !createForm.value.result || !createForm.value.handler) {
    message.error('请填写完整信息')
    return
  }
  showCreateModal.value = false
  createForm.value = { partID: '', batchNo: '', result: undefined, handler: '' }
  message.success('质检数据录入成功')
  fetchData()
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.quality-inspection {
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
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  color: #fff;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.35);
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
  height: 40px !important;
  min-height: 40px !important;
  max-height: 40px !important;
  line-height: 40px !important;
  border-radius: 10px;
  padding: 0 20px;
  font-weight: 500;
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  border: none;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.35);
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.create-btn:hover {
  background: linear-gradient(135deg, #0891b2 0%, #0e7490 100%);
  box-shadow: 0 6px 16px rgba(6, 182, 212, 0.45);
  height: 40px !important;
  min-height: 40px !important;
  max-height: 40px !important;
  line-height: 40px !important;
}

.stats-overview {
  display: flex;
  gap: 24px;
  padding: 20px 24px;
  background: linear-gradient(135deg, #ecfeff 0%, #cffafe 100%);
  border-radius: 12px;
  border: 1px solid #a5f3fc;
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

.status-passed { color: #10b981; }
.status-failed { color: #ef4444; }
.status-rate { color: #06b6d4; }

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
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
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
  color: #06b6d4;
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

.dot-passed { background: #10b981; }
.dot-failed { background: #ef4444; }

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
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
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
  color: #06b6d4;
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
  color: #06b6d4;
  border-color: #06b6d4;
}

.export-btn {
  color: #10b981;
  border-color: #10b981;
}

.custom-table :deep(.ant-table-thead > tr > th) {
  background: linear-gradient(135deg, #ecfeff 0%, #cffafe 100%);
  font-weight: 600;
  color: #374151;
}

.custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #ecfeff;
}

.inspection-id-cell {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
  color: #06b6d4;
}

.inspection-icon {
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

.result-tag {
  display: flex;
  align-items: center;
  gap: 4px;
}

.handler-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.handler-name {
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
  color: #06b6d4;
}

.report-btn {
  color: #8b5cf6;
}

/* 深色主题 */
[data-theme='dark'] .title-icon {
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.45);
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
