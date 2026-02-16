<!--
  召回记录页面
-->

<template>
  <div class="recall-record">
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <WarningOutlined />
          </div>
          <div class="title-text">
            <h2>召回记录</h2>
            <p class="subtitle">管理产品召回与安全事件处理</p>
          </div>
        </div>
        <a-button type="primary" class="create-btn" @click="showRecallModal = true">
          <template #icon><PlusOutlined /></template>
          发起召回
        </a-button>
      </div>
      <div class="stats-overview">
        <div class="stat-item">
          <div class="stat-value">{{ totalCount }}</div>
          <div class="stat-label">总召回数</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-progress">{{ inProgressCount }}</div>
          <div class="stat-label">进行中</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-completed">{{ completedCount }}</div>
          <div class="stat-label">已完成</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-affected">{{ totalAffectedCount }}</div>
          <div class="stat-label">受影响零部件</div>
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
          <a-tag color="red" class="filter-tag">
            <template #icon><SearchOutlined /></template>
            召回查询
          </a-tag>
        </div>
      </div>
      <a-form layout="inline" :model="searchForm" class="search-form">
        <a-row :gutter="[20, 20]" style="width: 100%">
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <NumberOutlined class="label-icon" />
                <span>召回ID</span>
              </div>
              <a-input
                v-model:value="searchForm.recallID"
                placeholder="请输入召回ID"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <WarningOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <AppstoreOutlined class="label-icon" />
                <span>涉及批次</span>
              </div>
              <a-input
                v-model:value="searchForm.batchNo"
                placeholder="请输入批次号"
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
                <a-select-option value="进行中">
                  <span class="status-option"><span class="status-dot dot-progress"></span>进行中</span>
                </a-select-option>
                <a-select-option value="已完成">
                  <span class="status-option"><span class="status-dot dot-completed"></span>已完成</span>
                </a-select-option>
              </a-select>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <CalendarOutlined class="label-icon" />
                <span>时间范围</span>
              </div>
              <a-range-picker
                v-model:value="searchForm.dateRange"
                class="filter-date"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                :dropdown-style="{ zIndex: 1050 }"
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
        row-key="recallID"
        class="custom-table"
        :scroll="{ x: 1200 }"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'recallID'">
            <div class="recall-id-cell">
              <WarningOutlined class="recall-icon" />
              <span>{{ record.recallID }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'batchNos'">
            <div class="batch-cell">
              <a-tag v-for="(batch, index) in record.batchNos.split(', ')" :key="index" color="blue" class="batch-tag">
                {{ batch }}
              </a-tag>
            </div>
          </template>
          <template v-else-if="column.key === 'reason'">
            <a-tooltip :title="record.reason">
              <span class="reason-text">{{ record.reason }}</span>
            </a-tooltip>
          </template>
          <template v-else-if="column.key === 'affectedCount'">
            <div class="affected-cell">
              <span class="affected-value">{{ record.affectedCount }}</span>
              <span class="affected-unit">件</span>
            </div>
          </template>
          <template v-else-if="column.key === 'progress'">
            <div class="progress-cell">
              <a-progress
                :percent="record.progress"
                :stroke-color="record.status === '已完成' ? '#10b981' : '#3b82f6'"
                size="small"
                class="recall-progress"
              />
            </div>
          </template>
          <template v-else-if="column.key === 'status'">
            <div class="status-cell">
              <span class="status-dot" :class="getStatusDotClass(record.status)"></span>
              <span class="status-text">{{ record.status }}</span>
            </div>
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
              <a-button
                type="link"
                size="small"
                @click="handleComplete(record)"
                :disabled="record.status === '已完成'"
                class="action-btn complete-btn"
              >
                <CheckCircleOutlined />
                完成
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <a-modal
      v-model:open="showRecallModal"
      title="发起召回"
      :width="650"
      @ok="handleSubmitRecall"
      @cancel="showRecallModal = false"
      class="recall-modal"
    >
      <a-form :model="recallForm" layout="vertical" class="create-form">
        <a-form-item label="涉及批次" required>
          <a-select
            v-model:value="recallForm.batchNos"
            mode="multiple"
            placeholder="请选择涉及批次"
            :options="batchOptions"
          />
        </a-form-item>
        <a-form-item label="召回原因" required>
          <a-textarea v-model:value="recallForm.reason" placeholder="请详细说明召回原因" :rows="3" />
        </a-form-item>
        <a-form-item label="受影响零部件数量" required>
          <a-input-number v-model:value="recallForm.affectedCount" :min="1" style="width: 100%" addon-after="件" />
        </a-form-item>
        <a-form-item label="召回级别">
          <a-select v-model:value="recallForm.level" placeholder="请选择召回级别">
            <a-select-option value="一级">一级（紧急）</a-select-option>
            <a-select-option value="二级">二级（重要）</a-select-option>
            <a-select-option value="三级">三级（一般）</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="处理措施">
          <a-textarea v-model:value="recallForm.measures" placeholder="请描述处理措施" :rows="3" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { message } from 'ant-design-vue'
import {
  WarningOutlined,
  PlusOutlined,
  FilterOutlined,
  SearchOutlined,
  NumberOutlined,
  AppstoreOutlined,
  CheckCircleOutlined,
  CalendarOutlined,
  InfoCircleOutlined,
  ReloadOutlined,
  UnorderedListOutlined,
  DatabaseOutlined,
  ExportOutlined,
  ClockCircleOutlined,
  EyeOutlined
} from '@ant-design/icons-vue'

const refreshing = ref(false)
const showRecallModal = ref(false)

const searchForm = ref({
  recallID: '',
  batchNo: '',
  status: undefined as string | undefined,
  dateRange: null as any
})

const recallForm = ref({
  batchNos: [] as string[],
  reason: '',
  affectedCount: 100,
  level: '二级',
  measures: ''
})

const batchOptions = ref([
  { value: 'BATCH-001', label: 'BATCH-001' },
  { value: 'BATCH-002', label: 'BATCH-002' },
  { value: 'BATCH-003', label: 'BATCH-003' },
  { value: 'BATCH-004', label: 'BATCH-004' }
])

const columns = [
  { title: '召回ID', key: 'recallID', width: 140 },
  { title: '涉及批次', key: 'batchNos', width: 200 },
  { title: '召回原因', key: 'reason', width: 200 },
  { title: '受影响零部件', key: 'affectedCount', width: 120 },
  { title: '进度', key: 'progress', width: 150 },
  { title: '状态', key: 'status', width: 100 },
  { title: '发起时间', key: 'createTime', width: 150 },
  { title: '操作', key: 'action', width: 150, fixed: 'right' as const }
]

const tableData = ref([
  { recallID: 'RECALL-001', batchNos: 'BATCH-001, BATCH-002', reason: '制动系统潜在隐患，可能导致制动效能下降', affectedCount: 150, progress: 75, status: '进行中', createTime: '2024-01-15 10:30' },
  { recallID: 'RECALL-002', batchNos: 'BATCH-003', reason: '转向系统存在异响问题', affectedCount: 80, progress: 100, status: '已完成', createTime: '2024-01-10 14:20' },
  { recallID: 'RECALL-003', batchNos: 'BATCH-004, BATCH-005', reason: '安全气囊传感器故障风险', affectedCount: 200, progress: 45, status: '进行中', createTime: '2024-01-08 09:15' },
  { recallID: 'RECALL-004', batchNos: 'BATCH-006', reason: '发动机舱线束磨损隐患', affectedCount: 60, progress: 100, status: '已完成', createTime: '2024-01-05 16:45' },
  { recallID: 'RECALL-005', batchNos: 'BATCH-007, BATCH-008', reason: '燃油泵工作异常问题', affectedCount: 120, progress: 30, status: '进行中', createTime: '2024-01-02 11:00' }
])

const totalCount = computed(() => tableData.value.length)
const inProgressCount = computed(() => tableData.value.filter(item => item.status === '进行中').length)
const completedCount = computed(() => tableData.value.filter(item => item.status === '已完成').length)
const totalAffectedCount = computed(() => tableData.value.reduce((sum, item) => sum + item.affectedCount, 0))

const getStatusDotClass = (status: string) => {
  return status === '已完成' ? 'dot-completed' : 'dot-progress'
}

const handleSearch = () => {
  message.success('搜索完成')
}

const handleReset = () => {
  searchForm.value = {
    recallID: '',
    batchNo: '',
    status: undefined,
    dateRange: null
  }
}

const handleRefresh = () => {
  refreshing.value = true
  setTimeout(() => {
    refreshing.value = false
    message.success('数据已刷新')
  }, 1000)
}

const handleExport = () => {
  message.success('召回报告导出成功')
}

const viewDetail = (record: any) => {
  message.info(`查看召回详情: ${record.recallID}`)
}

const handleComplete = (record: any) => {
  if (record.status === '已完成') return
  record.status = '已完成'
  record.progress = 100
  message.success(`召回 ${record.recallID} 已标记为完成`)
}

const handleSubmitRecall = () => {
  if (!recallForm.value.batchNos.length || !recallForm.value.reason || !recallForm.value.affectedCount) {
    message.warning('请填写必填项')
    return
  }
  const newRecall = {
    recallID: `RECALL-${String(tableData.value.length + 1).padStart(3, '0')}`,
    batchNos: recallForm.value.batchNos.join(', '),
    reason: recallForm.value.reason,
    affectedCount: recallForm.value.affectedCount,
    progress: 0,
    status: '进行中',
    createTime: new Date().toLocaleString()
  }
  tableData.value.unshift(newRecall)
  showRecallModal.value = false
  message.success('召回发起成功')
}
</script>

<style scoped>
.recall-record {
  padding: 0;
  min-height: 100%;
}

.page-header {
  margin-bottom: 24px;
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
  border-radius: 12px;
  padding: 24px 32px;
  box-shadow: 0 4px 20px rgba(239, 68, 68, 0.25);
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

.stat-value.status-progress {
  color: #60a5fa;
}

.stat-value.status-completed {
  color: #34d399;
}

.stat-value.status-affected {
  color: #fbbf24;
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
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
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
  color: #ef4444;
}

.filter-input,
.filter-select,
.filter-date {
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

.dot-progress {
  background: #3b82f6;
}

.dot-completed {
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
  border-color: #ef4444;
  color: #ef4444;
}

.search-btn {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
  border: none;
}

.search-btn:hover {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
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
  color: #ef4444;
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
  border-color: #ef4444;
  color: #ef4444;
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
  background: #fef2f2;
}

.recall-id-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.recall-icon {
  color: #ef4444;
}

.batch-cell {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.batch-tag {
  margin: 0;
  font-size: 11px;
}

.reason-text {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  font-size: 13px;
  color: #374151;
}

.affected-cell {
  display: flex;
  align-items: baseline;
  gap: 4px;
}

.affected-value {
  font-size: 16px;
  font-weight: 600;
  color: #ef4444;
}

.affected-unit {
  font-size: 12px;
  color: #6b7280;
}

.progress-cell {
  width: 100%;
}

.recall-progress {
  width: 100%;
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

.status-dot.dot-progress {
  background: #3b82f6;
}

.status-dot.dot-completed {
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
  color: #ef4444;
}

.complete-btn {
  color: #10b981;
}

.complete-btn:hover:not(:disabled) {
  color: #059669;
}

.complete-btn:disabled {
  color: #d1d5db;
  cursor: not-allowed;
}

.recall-form :deep(.ant-input),
.recall-form :deep(.ant-select-selector),
.recall-form :deep(.ant-input-number) {
  border-radius: 8px;
}

[data-theme='dark'] .page-header {
  background: linear-gradient(135deg, #b91c1c 0%, #991b1b 100%);
  box-shadow: 0 4px 20px rgba(185, 28, 28, 0.3);
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
  border-color: #f87171;
  color: #f87171;
}

[data-theme='dark'] .search-btn {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
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

[data-theme='dark'] .reason-text {
  color: #e5e7eb;
}

[data-theme='dark'] .affected-value {
  color: #f87171;
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
  color: #f87171;
}
</style>
