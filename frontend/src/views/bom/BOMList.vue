<!--
  BOM列表页面
  展示BOM列表，支持搜索和查看详情
  设计：与Dashboard风格统一
-->

<template>
  <div class="bom-list">
    <!-- 页面标题 -->
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <FileTextOutlined />
          </div>
          <div class="title-text">
            <h2>BOM列表</h2>
            <p class="subtitle">管理物料清单信息</p>
          </div>
        </div>
        <a-button type="primary" class="create-btn" @click="goCreate">
          <template #icon><PlusOutlined /></template>
          创建BOM
        </a-button>
      </div>
      <!-- 统计概览 -->
      <div class="stats-overview">
        <div class="stat-item">
          <div class="stat-value">{{ tableData.length }}</div>
          <div class="stat-label">总数量</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-published">{{ getStatusCount('已发布') }}</div>
          <div class="stat-label">已发布</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-draft">{{ getStatusCount('草稿') }}</div>
          <div class="stat-label">草稿</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-archived">{{ getStatusCount('已归档') }}</div>
          <div class="stat-label">已归档</div>
        </div>
      </div>
    </div>
    
    <!-- 搜索筛选区 -->
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
          <a-tag color="blue" class="filter-tag">
            <template #icon><SearchOutlined /></template>
            高级搜索
          </a-tag>
        </div>
      </div>
      <a-form layout="inline" :model="searchForm" class="search-form">
        <a-row :gutter="[20, 20]" style="width: 100%">
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <NumberOutlined class="label-icon" />
                <span>BOM ID</span>
              </div>
              <a-input
                v-model:value="searchForm.bomID"
                placeholder="请输入BOM ID"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <BarcodeOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <AppstoreOutlined class="label-icon" />
                <span>BOM类型</span>
              </div>
              <a-select
                v-model:value="searchForm.bomType"
                placeholder="请选择类型"
                allow-clear
                class="filter-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                :dropdown-style="{ zIndex: 1050 }"
              >
                <a-select-option value="生产BOM">生产BOM</a-select-option>
                <a-select-option value="研发BOM">研发BOM</a-select-option>
                <a-select-option value="工程BOM">工程BOM</a-select-option>
              </a-select>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <CarOutlined class="label-icon" />
                <span>车型</span>
              </div>
              <a-input
                v-model:value="searchForm.productModel"
                placeholder="请输入车型"
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
                <a-select-option value="已发布">
                  <span class="status-option"><span class="status-dot dot-published"></span>已发布</span>
                </a-select-option>
                <a-select-option value="草稿">
                  <span class="status-option"><span class="status-dot dot-draft"></span>草稿</span>
                </a-select-option>
                <a-select-option value="已归档">
                  <span class="status-option"><span class="status-dot dot-archived"></span>已归档</span>
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
    
    <!-- 数据表格 -->
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
            <a-button 
              class="toolbar-btn refresh-btn" 
              @click="handleRefresh"
              :loading="refreshing"
            >
              <template #icon><ReloadOutlined :spin="refreshing" /></template>
              {{ refreshing ? '刷新中...' : '刷新数据' }}
            </a-button>
            <a-button 
              class="toolbar-btn compare-btn" 
              @click="goCompare"
            >
              <template #icon><DiffOutlined /></template>
              BOM比较
            </a-button>
          </a-space>
        </div>
      </div>
      <a-table
        :columns="columns"
        :data-source="tableData"
        row-key="bomID"
        class="custom-table"
        :scroll="{ x: 1000 }"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'bomID'">
            <div class="bom-id-cell">
              <span class="bom-id">{{ record.bomID }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'bomType'">
            <a-tag :color="getTypeColor(record.bomType)" class="type-tag">
              {{ record.bomType }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'status'">
            <div class="status-cell">
              <span class="status-dot" :class="getStatusDotClass(record.status)"></span>
              <span class="status-text">{{ record.status }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'version'">
            <a-tag color="purple" class="version-tag">{{ record.version }}</a-tag>
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
              <a-button type="link" size="small" @click="editBOM(record)" class="action-btn edit-btn">
                <EditOutlined />
                编辑
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { listBOMs, type BOM } from '../../services/bom.service'
import {
  FileTextOutlined,
  PlusOutlined,
  FilterOutlined,
  SearchOutlined,
  ReloadOutlined,
  NumberOutlined,
  AppstoreOutlined,
  CarOutlined,
  CheckCircleOutlined,
  InfoCircleOutlined,
  UnorderedListOutlined,
  DatabaseOutlined,
  DiffOutlined,
  ClockCircleOutlined,
  EyeOutlined,
  EditOutlined,
  BarcodeOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const refreshing = ref(false)
const loading = ref(false)

const searchForm = reactive({
  bomID: '',
  bomType: '',
  productModel: '',
  status: ''
})

const tableData = ref<any[]>([])

const columns = [
  { title: 'BOM ID', dataIndex: 'bomID', key: 'bomID', width: 120 },
  { title: '类型', dataIndex: 'bomType', key: 'bomType', width: 120 },
  { title: '车型', dataIndex: 'productModel', key: 'productModel', width: 120 },
  { title: '版本', dataIndex: 'version', key: 'version', width: 100 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 100 },
  { title: '创建时间', dataIndex: 'createTime', key: 'createTime', width: 180 },
  { title: '操作', key: 'action', width: 150, fixed: 'right' as const }
]

function getStatusCount(status: string) {
  return tableData.value.filter(item => item.status === status).length
}

function getTypeColor(type: string) {
  const colors: Record<string, string> = {
    '生产BOM': 'blue',
    '研发BOM': 'green',
    '工程BOM': 'orange'
  }
  return colors[type] || 'default'
}

function getStatusDotClass(status: string) {
  const classes: Record<string, string> = {
    '已发布': 'dot-published',
    '草稿': 'dot-draft',
    '已归档': 'dot-archived'
  }
  return classes[status] || ''
}

async function handleSearch() {
  loading.value = true
  try {
    const response = await listBOMs()
    if (response.code === 0 && response.data) {
      let boms = response.data
      
      if (searchForm.bomID) {
        boms = boms.filter(b => b.bomID.includes(searchForm.bomID))
      }
      if (searchForm.bomType) {
        boms = boms.filter(b => b.bomType === searchForm.bomType)
      }
      if (searchForm.productModel) {
        boms = boms.filter(b => b.productModel.includes(searchForm.productModel))
      }
      if (searchForm.status) {
        boms = boms.filter(b => b.status === searchForm.status)
      }
      
      tableData.value = boms.map(b => ({
        ...b,
        createTime: b.createTime
      }))
    }
  } catch (error: any) {
    message.error(error.message || '查询失败')
    tableData.value = []
  } finally {
    loading.value = false
  }
}

function handleReset() {
  searchForm.bomID = ''
  searchForm.bomType = ''
  searchForm.productModel = ''
  searchForm.status = ''
  handleSearch()
}

async function handleRefresh() {
  refreshing.value = true
  try {
    await handleSearch()
    message.success('数据刷新成功')
  } catch (error: any) {
    message.error('刷新失败')
  } finally {
    setTimeout(() => {
      refreshing.value = false
    }, 500)
  }
}

function goCreate() {
  router.push('/bom/create')
}

function goCompare() {
  router.push('/bom/compare')
}

function viewDetail(record: any) {
  router.push(`/bom/detail/${record.bomID}`)
}

function editBOM(record: any) {
  router.push(`/bom/edit/${record.bomID}`)
}

onMounted(() => {
  handleSearch()
})
</script>

<style scoped>
.bom-list {
  padding: 0;
  position: relative;
}

.bom-list :deep(.ant-select-dropdown) {
  z-index: 1050 !important;
}

/* ==================== 页面头部样式 ==================== */
.page-header {
  background: linear-gradient(135deg, #1890ff 0%, #722ed1 100%);
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 24px;
  box-shadow: 0 4px 20px rgba(24, 144, 255, 0.3);
  position: relative;
  z-index: 5;
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
  color: #fff;
  font-size: 24px;
}

.title-text h2 {
  font-size: 24px;
  font-weight: 700;
  color: #fff;
  margin: 0;
}

.subtitle {
  color: rgba(255, 255, 255, 0.8);
  margin: 4px 0 0 0;
  font-size: 14px;
}

.create-btn {
  height: 40px;
  padding: 0 24px;
  border-radius: 8px;
  font-weight: 600;
  background: rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: #fff;
  backdrop-filter: blur(10px);
}

.create-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  border-color: rgba(255, 255, 255, 0.5);
  color: #fff;
}

.stats-overview {
  display: flex;
  gap: 24px;
  padding-top: 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.2);
}

.stat-item {
  text-align: center;
}

.stat-value {
  font-size: 28px;
  font-weight: 700;
  color: #fff;
  line-height: 1;
}

.stat-label {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.75);
  margin-top: 6px;
}

.status-published { color: #52c41a; }
.status-draft { color: #faad14; }
.status-archived { color: #8c8c8c; }

/* ==================== 搜索卡片样式 ==================== */
.search-card {
  margin-bottom: 24px;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.06);
  overflow: visible;
  background: linear-gradient(180deg, #ffffff 0%, #fafbfc 100%);
  border: 1px solid rgba(24, 144, 255, 0.1);
  position: relative;
  z-index: 10;
}

.search-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 24px;
  margin-bottom: 8px;
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.05) 0%, rgba(114, 46, 209, 0.05) 100%);
  border-bottom: 1px solid rgba(24, 144, 255, 0.1);
}

.header-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.filter-icon-wrapper {
  width: 44px;
  height: 44px;
  background: linear-gradient(135deg, #1890ff 0%, #722ed1 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 20px;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3);
}

.header-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.header-title {
  font-size: 18px;
  font-weight: 700;
  color: #262626;
  letter-spacing: 0.5px;
}

.header-desc {
  font-size: 13px;
  color: #8c8c8c;
}

.header-right {
  display: flex;
  align-items: center;
}

.filter-tag {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 14px;
  border-radius: 20px;
  font-weight: 500;
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.1) 0%, rgba(114, 46, 209, 0.1) 100%);
  border: 1px solid rgba(24, 144, 255, 0.2);
}

.search-form {
  width: 100%;
  padding: 20px 24px;
}

.filter-item {
  width: 100%;
}

.filter-label {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 10px;
  font-weight: 600;
  color: #262626;
  font-size: 14px;
}

.label-icon {
  color: #1890ff;
  font-size: 15px;
}

.filter-input,
.filter-select {
  width: 100%;
  height: 42px;
}

.input-prefix-icon {
  font-size: 14px;
}

.status-option {
  display: flex;
  align-items: center;
  gap: 8px;
}

.status-option .status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.dot-published { background: #52c41a; }
.dot-draft { background: #faad14; }
.dot-archived { background: #8c8c8c; }

.filter-select :deep(.ant-select-dropdown) {
  z-index: 1050 !important;
  border-radius: 8px;
  box-shadow: 0 6px 16px 0 rgba(0, 0, 0, 0.12);
}

.filter-select :deep(.ant-select-item) {
  padding: 8px 12px;
  transition: all 0.3s;
  border-radius: 4px;
  margin: 4px;
}

.filter-select :deep(.ant-select-item:hover) {
  background: #e6f7ff;
}

.filter-select :deep(.ant-select-item-option-selected) {
  background: #bae7ff;
  font-weight: 600;
}

.search-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 24px;
  padding: 20px 24px;
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.03) 0%, rgba(114, 46, 209, 0.03) 100%);
  border-top: 1px solid rgba(24, 144, 255, 0.1);
  border-radius: 0 0 16px 16px;
  margin-left: -24px;
  margin-right: -24px;
  margin-bottom: -24px;
}

.action-left {
  display: flex;
  align-items: center;
}

.action-hint {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #8c8c8c;
  font-size: 13px;
  background: rgba(0, 0, 0, 0.02);
  padding: 8px 16px;
  border-radius: 20px;
}

.action-right {
  display: flex;
  gap: 12px;
}

.reset-btn {
  height: 42px;
  padding: 0 24px;
  border-radius: 10px;
  font-weight: 600;
  border: 2px solid #e8e8e8;
  background: #fff;
  color: #595959;
  transition: all 0.3s;
}

.reset-btn:hover {
  border-color: #1890ff;
  color: #1890ff;
  background: rgba(24, 144, 255, 0.05);
}

.search-btn {
  height: 42px;
  padding: 0 32px;
  border-radius: 10px;
  background: linear-gradient(135deg, #1890ff 0%, #722ed1 100%);
  border: none;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3);
  transition: all 0.3s;
}

.search-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(24, 144, 255, 0.4);
  background: linear-gradient(135deg, #40a9ff 0%, #9254de 100%);
}

/* ==================== 表格卡片样式 ==================== */
.table-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  overflow: visible;
  position: relative;
}

.table-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 18px 24px;
  border-bottom: 1px solid #f0f0f0;
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.02) 0%, rgba(114, 46, 209, 0.02) 100%);
}

.table-header-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.table-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 16px;
  font-weight: 600;
  color: #262626;
}

.table-title .title-icon {
  color: #1890ff;
  font-size: 18px;
}

.record-count {
  font-size: 13px;
  padding: 6px 14px;
  border-radius: 16px;
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.1) 0%, rgba(114, 46, 209, 0.1) 100%);
  border: 1px solid rgba(24, 144, 255, 0.2);
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-weight: 500;
}

.table-header-right {
  display: flex;
  align-items: center;
}

.toolbar-btn {
  height: 36px;
  padding: 0 16px;
  border-radius: 8px;
  font-size: 13px;
  font-weight: 500;
  border: 1px solid #e8e8e8;
  color: #595959;
  background: #fff;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  display: inline-flex;
  align-items: center;
  gap: 6px;
}

.toolbar-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.refresh-btn {
  border-color: rgba(24, 144, 255, 0.3);
  color: #1890ff;
}

.refresh-btn:hover {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.08) 0%, rgba(114, 46, 209, 0.08) 100%);
  border-color: #1890ff;
  color: #40a9ff;
}

.compare-btn {
  border-color: rgba(114, 46, 209, 0.3);
  color: #722ed1;
}

.compare-btn:hover {
  background: linear-gradient(135deg, rgba(114, 46, 209, 0.08) 0%, rgba(24, 144, 255, 0.08) 100%);
  border-color: #722ed1;
  color: #9254de;
}

/* ==================== 表格样式 ==================== */
.custom-table :deep(.ant-table-thead > tr > th) {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.04) 0%, rgba(114, 46, 209, 0.04) 100%);
  font-weight: 600;
  color: #262626;
  border-bottom: 2px solid rgba(24, 144, 255, 0.1);
  padding: 14px 16px;
}

.custom-table :deep(.ant-table-tbody > tr > td) {
  padding: 14px 16px;
  border-bottom: 1px solid #f0f0f0;
  transition: all 0.3s;
}

.custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: rgba(24, 144, 255, 0.04);
}

.bom-id-cell {
  display: flex;
  align-items: center;
}

.bom-id {
  font-weight: 600;
  color: #1890ff;
  font-family: 'Monaco', 'Menlo', monospace;
}

.type-tag {
  border-radius: 6px;
  font-weight: 500;
}

.version-tag {
  border-radius: 6px;
  font-weight: 500;
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

.status-text {
  font-size: 14px;
  color: #595959;
}

.time-cell {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #8c8c8c;
  font-size: 13px;
}

.time-icon {
  color: #bfbfbf;
}

.action-cell {
  gap: 8px;
}

.action-btn {
  padding: 4px 8px;
  height: auto;
  font-size: 13px;
  border-radius: 6px;
}

.detail-btn {
  color: #1890ff;
}

.detail-btn:hover {
  background: rgba(24, 144, 255, 0.1);
}

.edit-btn {
  color: #52c41a;
}

.edit-btn:hover {
  background: rgba(82, 196, 26, 0.1);
}

/* ==================== 深色主题 ==================== */
[data-theme='dark'] .page-header {
  background: linear-gradient(135deg, #177ddc 0%, #642ab5 100%);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.4);
}

[data-theme='dark'] .search-card {
  background: var(--bg-color-secondary);
  border-color: rgba(255, 255, 255, 0.1);
}

[data-theme='dark'] .search-header {
  background: linear-gradient(135deg, rgba(23, 125, 220, 0.1) 0%, rgba(100, 42, 181, 0.1) 100%);
  border-color: rgba(255, 255, 255, 0.1);
}

[data-theme='dark'] .header-title {
  color: var(--text-color);
}

[data-theme='dark'] .header-desc {
  color: var(--text-color-secondary);
}

[data-theme='dark'] .filter-label {
  color: var(--text-color);
}

[data-theme='dark'] .search-actions {
  background: linear-gradient(135deg, rgba(23, 125, 220, 0.05) 0%, rgba(100, 42, 181, 0.05) 100%);
  border-color: rgba(255, 255, 255, 0.1);
}

[data-theme='dark'] .action-hint {
  background: rgba(255, 255, 255, 0.05);
  color: var(--text-color-secondary);
}

[data-theme='dark'] .reset-btn {
  background: var(--bg-color);
  border-color: var(--border-color);
  color: var(--text-color);
}

[data-theme='dark'] .reset-btn:hover {
  border-color: #177ddc;
  color: #177ddc;
  background: rgba(23, 125, 220, 0.1);
}

[data-theme='dark'] .search-btn {
  background: linear-gradient(135deg, #177ddc 0%, #642ab5 100%);
}

[data-theme='dark'] .search-btn:hover {
  background: linear-gradient(135deg, #3c9ae8 0%, #7a3db8 100%);
}

[data-theme='dark'] .table-card {
  background: var(--bg-color-secondary);
}

[data-theme='dark'] .table-header {
  background: linear-gradient(135deg, rgba(23, 125, 220, 0.05) 0%, rgba(100, 42, 181, 0.05) 100%);
  border-color: rgba(255, 255, 255, 0.1);
}

[data-theme='dark'] .table-title {
  color: var(--text-color);
}

[data-theme='dark'] .custom-table :deep(.ant-table-thead > tr > th) {
  background: linear-gradient(135deg, rgba(23, 125, 220, 0.08) 0%, rgba(100, 42, 181, 0.08) 100%);
  color: var(--text-color);
  border-color: rgba(255, 255, 255, 0.1);
}

[data-theme='dark'] .custom-table :deep(.ant-table-tbody > tr > td) {
  border-color: rgba(255, 255, 255, 0.08);
  color: var(--text-color);
}

[data-theme='dark'] .custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: rgba(23, 125, 220, 0.08);
}

[data-theme='dark'] .bom-id {
  color: #177ddc;
}

[data-theme='dark'] .status-text {
  color: var(--text-color-secondary);
}

[data-theme='dark'] .time-cell {
  color: var(--text-color-secondary);
}

[data-theme='dark'] .toolbar-btn {
  background: var(--bg-color);
  border-color: var(--border-color);
  color: var(--text-color);
}

[data-theme='dark'] .refresh-btn:hover,
[data-theme='dark'] .compare-btn:hover {
  background: rgba(23, 125, 220, 0.15);
}
</style>
