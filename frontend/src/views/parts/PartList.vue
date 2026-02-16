<!--
  零部件列表页面
  展示零部件列表，支持搜索、筛选、查看详情
-->

<template>
  <div class="part-list">
    <!-- 页面标题 -->
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <AppstoreOutlined />
          </div>
          <div class="title-text">
            <h2>零部件列表</h2>
            <p class="subtitle">管理和追踪所有零部件信息</p>
          </div>
        </div>
        <a-button type="primary" class="create-btn" @click="goCreate" v-if="canCreate">
          <template #icon><PlusOutlined /></template>
          创建零部件
        </a-button>
      </div>
      <!-- 统计概览 -->
      <div class="stats-overview">
        <div class="stat-item">
          <div class="stat-value">{{ tableData.length }}</div>
          <div class="stat-label">总数量</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-normal">{{ getStatusCount('NORMAL') }}</div>
          <div class="stat-label">正常</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-production">{{ getStatusCount('IN_PRODUCTION') }}</div>
          <div class="stat-label">在产</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-recalled">{{ getStatusCount('RECALLED') }}</div>
          <div class="stat-label">已召回</div>
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
                <IdcardOutlined class="label-icon" />
                <span>零部件ID</span>
              </div>
              <a-input
                v-model:value="searchForm.partID"
                placeholder="请输入零部件ID"
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
                  <BarcodeOutlined class="input-prefix-icon" />
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
                  <AimOutlined class="input-prefix-icon" />
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
                <a-select-option value="NORMAL">
                  <span class="status-option"><span class="status-dot dot-normal"></span>正常</span>
                </a-select-option>
                <a-select-option value="IN_PRODUCTION">
                  <span class="status-option"><span class="status-dot dot-production"></span>在产</span>
                </a-select-option>
                <a-select-option value="RECALLED">
                  <span class="status-option"><span class="status-dot dot-recalled"></span>已召回</span>
                </a-select-option>
                <a-select-option value="SCRAPPED">
                  <span class="status-option"><span class="status-dot dot-scrapped"></span>已报废</span>
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
            共 {{ pagination.total }} 条记录
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
              class="toolbar-btn export-btn" 
              @click="handleExport" 
              v-if="canCreate"
            >
              <template #icon><ExportOutlined /></template>
              导出数据
            </a-button>
          </a-space>
        </div>
      </div>
      <a-table
        :columns="columns"
        :data-source="tableData"
        :loading="loading"
        :pagination="pagination"
        @change="handleTableChange"
        row-key="partID"
        class="custom-table"
        :scroll="{ x: 1200 }"
      >
        <!-- 状态列 -->
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'partID'">
            <div class="part-id-cell">
              <span class="part-id">{{ record.partID }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'name'">
            <div class="name-cell">
              <div class="part-type-icon" :style="{ background: getTypeColor(record.type) }">
                <component :is="getTypeIcon(record.type)" />
              </div>
              <span class="part-name">{{ record.name }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'type'">
            <a-tag :color="getTypeTagColor(record.type)" class="type-tag">
              {{ record.type }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'status'">
            <div class="status-cell">
              <span class="status-dot" :class="getStatusDotClass(record.status)"></span>
              <span class="status-text">{{ getStatusText(record.status) }}</span>
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
              <a-button type="link" size="small" @click="viewDetail(record.partID)" class="action-btn detail-btn">
                <EyeOutlined />
                详情
              </a-button>
              <a-button type="link" size="small" @click="viewLifecycle(record.partID)" class="action-btn lifecycle-btn">
                <HistoryOutlined />
                生命周期
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>
    
    <!-- 生命周期抽屉 -->
    <a-drawer
      v-model:open="lifecycleVisible"
      title="零部件生命周期"
      :width="720"
      placement="right"
    >
      <a-descriptions :column="1" bordered v-if="lifecycleData">
        <a-descriptions-item label="零部件ID">
          {{ lifecycleData.partID }}
        </a-descriptions-item>
        <a-descriptions-item label="BOM信息">
          {{ lifecycleData.bomInfo ? `BOM-${lifecycleData.bomInfo.bomID} v${lifecycleData.bomInfo.version}` : '暂无' }}
        </a-descriptions-item>
        <a-descriptions-item label="生产信息">
          {{ lifecycleData.productionInfo ? `批次: ${lifecycleData.productionInfo.batchNo}` : '暂无' }}
        </a-descriptions-item>
        <a-descriptions-item label="质检信息">
          {{ lifecycleData.qualityInfo ? `结果: ${lifecycleData.qualityInfo.result}` : '暂无' }}
        </a-descriptions-item>
        <a-descriptions-item label="供应链信息">
          {{ lifecycleData.supplyChainInfo?.length || 0 }} 条记录
        </a-descriptions-item>
        <a-descriptions-item label="售后信息">
          {{ lifecycleData.aftersaleInfo?.length || 0 }} 条记录
        </a-descriptions-item>
      </a-descriptions>
    </a-drawer>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { useAuthStore, usePartStore } from '../../stores'
import { UserRole } from '../../types'
import type { Part, PartLifecycle } from '../../types'
import {
  AppstoreOutlined,
  PlusOutlined,
  SearchOutlined,
  ReloadOutlined,
  FilterOutlined,
  EyeOutlined,
  HistoryOutlined,
  ClockCircleOutlined,
  ThunderboltOutlined,
  AlertOutlined,
  SettingOutlined,
  CompassOutlined,
  BulbOutlined,
  CarOutlined,
  IdcardOutlined,
  TagOutlined,
  CheckCircleOutlined,
  NumberOutlined,
  BarcodeOutlined,
  AimOutlined,
  InfoCircleOutlined,
  UnorderedListOutlined,
  ExportOutlined,
  DatabaseOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const authStore = useAuthStore()
const partStore = usePartStore()

const loading = ref(false)
const refreshing = ref(false)

const searchForm = reactive({
  partID: '',
  batchNo: '',
  vin: '',
  status: ''
})

const tableData = ref<Part[]>([])

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  showSizeChanger: true,
  showQuickJumper: true
})

const lifecycleVisible = ref(false)
const lifecycleData = ref<PartLifecycle | null>(null)

const canCreate = computed(() => authStore.userRole === UserRole.MANUFACTURER)

const columns = [
  {
    title: '零部件ID',
    dataIndex: 'partID',
    key: 'partID',
    width: 150
  },
  {
    title: '名称',
    dataIndex: 'name',
    key: 'name',
    width: 150
  },
  {
    title: '类型',
    dataIndex: 'type',
    key: 'type',
    width: 120
  },
  {
    title: '批次号',
    dataIndex: 'batchNo',
    key: 'batchNo',
    width: 150
  },
  {
    title: 'VIN码',
    dataIndex: 'vin',
    key: 'vin',
    width: 180
  },
  {
    title: '状态',
    dataIndex: 'status',
    key: 'status',
    width: 100
  },
  {
    title: '创建时间',
    dataIndex: 'createTime',
    key: 'createTime',
    width: 180
  },
  {
    title: '操作',
    key: 'action',
    width: 180,
    fixed: 'right' as const
  }
]

function getStatusCount(status: string) {
  return tableData.value.filter(item => item.status === status).length
}

function getTypeColor(type: string) {
  const colors: Record<string, string> = {
    '发动机部件': 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
    '制动系统': 'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)',
    '传动系统': 'linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)',
    '悬挂系统': 'linear-gradient(135deg, #43e97b 0%, #38f9d7 100%)',
    '电气系统': 'linear-gradient(135deg, #fa709a 0%, #fee140 100%)',
    '车身部件': 'linear-gradient(135deg, #a8edea 0%, #fed6e3 100%)'
  }
  return colors[type] || 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)'
}

function getTypeTagColor(type: string) {
  const colors: Record<string, string> = {
    '发动机部件': 'purple',
    '制动系统': 'magenta',
    '传动系统': 'blue',
    '悬挂系统': 'green',
    '电气系统': 'orange',
    '车身部件': 'cyan'
  }
  return colors[type] || 'default'
}

function getTypeIcon(type: string) {
  const icons: Record<string, any> = {
    '发动机部件': ThunderboltOutlined,
    '制动系统': AlertOutlined,
    '传动系统': SettingOutlined,
    '悬挂系统': CompassOutlined,
    '电气系统': BulbOutlined,
    '车身部件': CarOutlined
  }
  return icons[type] || AppstoreOutlined
}

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    'NORMAL': 'green',
    'IN_PRODUCTION': 'blue',
    'QUALITY_INSPECTED': 'cyan',
    'IN_SUPPLY_CHAIN': 'purple',
    'RECALLED': 'orange',
    'SCRAPPED': 'default'
  }
  return colors[status] || 'default'
}

function getStatusText(status: string) {
  const texts: Record<string, string> = {
    'NORMAL': '正常',
    'IN_PRODUCTION': '在产',
    'QUALITY_INSPECTED': '质检完成',
    'IN_SUPPLY_CHAIN': '供应链中',
    'RECALLED': '已召回',
    'SCRAPPED': '已报废'
  }
  return texts[status] || status
}

function getStatusDotClass(status: string) {
  const classes: Record<string, string> = {
    'NORMAL': 'dot-normal',
    'IN_PRODUCTION': 'dot-production',
    'QUALITY_INSPECTED': 'dot-quality',
    'IN_SUPPLY_CHAIN': 'dot-supply',
    'RECALLED': 'dot-recalled',
    'SCRAPPED': 'dot-scrapped'
  }
  return classes[status] || ''
}

async function handleSearch() {
  loading.value = true
  
  try {
    if (searchForm.batchNo) {
      await partStore.fetchByBatchNo(searchForm.batchNo)
      tableData.value = partStore.partList
    } else if (searchForm.vin) {
      await partStore.fetchByVIN(searchForm.vin)
      tableData.value = partStore.partList
    } else {
      tableData.value = generateMockData()
    }
    
    pagination.total = tableData.value.length
  } catch (error: any) {
    message.error(error.message || '查询失败')
  } finally {
    loading.value = false
  }
}

function handleReset() {
  searchForm.partID = ''
  searchForm.batchNo = ''
  searchForm.vin = ''
  searchForm.status = ''
  handleSearch()
}

function handleTableChange(pag: any) {
  pagination.current = pag.current
  pagination.pageSize = pag.pageSize
  handleSearch()
}

async function handleRefresh() {
  refreshing.value = true
  try {
    await handleSearch()
    message.success('数据刷新成功')
  } catch (error: any) {
    message.error('刷新失败: ' + (error.message || '未知错误'))
  } finally {
    setTimeout(() => {
      refreshing.value = false
    }, 500)
  }
}

function handleExport() {
  if (tableData.value.length === 0) {
    message.warning('暂无数据可导出')
    return
  }
  
  const hide = message.loading('正在导出数据...', 0)
  
  setTimeout(() => {
    hide()
    
    const csvContent = generateCSV()
    const blob = new Blob(['\ufeff' + csvContent], { type: 'text/csv;charset=utf-8;' })
    const link = document.createElement('a')
    const url = URL.createObjectURL(blob)
    
    link.setAttribute('href', url)
    link.setAttribute('download', `零部件数据_${new Date().toISOString().slice(0, 10)}.csv`)
    link.style.visibility = 'hidden'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    
    message.success('数据导出成功')
  }, 1000)
}

function generateCSV(): string {
  const headers = ['零部件ID', '名称', '类型', '批次号', 'VIN码', '状态', '创建时间']
  const rows = tableData.value.map(item => [
    item.partID,
    item.name,
    item.type,
    item.batchNo,
    item.vin,
    getStatusText(item.status),
    item.createTime
  ])
  
  return [headers.join(','), ...rows.map(row => row.join(','))].join('\n')
}

function goCreate() {
  router.push('/parts/create')
}

function viewDetail(partID: string) {
  router.push(`/parts/detail/${partID}`)
}

async function viewLifecycle(partID: string) {
  try {
    await partStore.fetchLifecycle(partID)
    lifecycleData.value = partStore.currentLifecycle
    lifecycleVisible.value = true
  } catch (error: any) {
    message.error(error.message || '查询生命周期失败')
  }
}

function generateMockData(): Part[] {
  return [
    {
      partID: 'PART-001',
      vin: 'LSVNV2182E2100001',
      batchNo: 'BATCH-2024-001',
      name: '发动机活塞',
      type: '发动机部件',
      manufacturer: 'Org1MSP',
      createTime: '2024-01-15 10:30:00',
      status: 'NORMAL'
    },
    {
      partID: 'PART-002',
      vin: 'LSVNV2182E2100002',
      batchNo: 'BATCH-2024-001',
      name: '制动盘',
      type: '制动系统',
      manufacturer: 'Org1MSP',
      createTime: '2024-01-16 14:20:00',
      status: 'IN_SUPPLY_CHAIN'
    },
    {
      partID: 'PART-003',
      vin: 'LSVNV2182E2100003',
      batchNo: 'BATCH-2024-002',
      name: '变速箱齿轮',
      type: '传动系统',
      manufacturer: 'Org1MSP',
      createTime: '2024-01-17 09:15:00',
      status: 'RECALLED'
    }
  ]
}

onMounted(() => {
  handleSearch()
})
</script>

<style scoped>
.part-list {
  padding: 0;
  position: relative;
}

.part-list :deep(.ant-select-dropdown) {
  z-index: 1050 !important;
}

.part-list :deep(.ant-picker-dropdown) {
  z-index: 1050 !important;
}

/* ==================== 页面头部样式 ==================== */
.page-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 24px;
  box-shadow: 0 4px 20px rgba(102, 126, 234, 0.3);
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

.status-normal { color: #52c41a; }
.status-production { color: #1890ff; }
.status-recalled { color: #faad14; }

/* ==================== 搜索卡片样式 ==================== */
.search-card {
  margin-bottom: 24px;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.06);
  overflow: visible;
  background: linear-gradient(180deg, #ffffff 0%, #fafbfc 100%);
  border: 1px solid rgba(102, 126, 234, 0.1);
  position: relative;
  z-index: 10;
}

.search-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 24px;
  margin-bottom: 8px;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
}

.header-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.filter-icon-wrapper {
  width: 44px;
  height: 44px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 20px;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
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
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.1) 0%, rgba(102, 126, 234, 0.1) 100%);
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
  color: #667eea;
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
  background: rgba(102, 126, 234, 0.08);
}

.filter-select :deep(.ant-select-item-option-selected) {
  background: rgba(102, 126, 234, 0.12);
  font-weight: 600;
}

.filter-select :deep(.ant-select-item-option-active) {
  background: rgba(102, 126, 234, 0.06);
}

.search-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 24px;
  padding: 20px 24px;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.03) 0%, rgba(118, 75, 162, 0.03) 100%);
  border-top: 1px solid rgba(102, 126, 234, 0.1);
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
  border-color: #667eea;
  color: #667eea;
  background: rgba(102, 126, 234, 0.05);
}

.search-btn {
  height: 42px;
  padding: 0 32px;
  border-radius: 10px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
  transition: all 0.3s;
}

.search-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
  background: linear-gradient(135deg, #5a6fd6 0%, #6a4196 100%);
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
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.02) 0%, rgba(118, 75, 162, 0.02) 100%);
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
  color: #667eea;
  font-size: 18px;
}

.record-count {
  font-size: 13px;
  padding: 6px 14px;
  border-radius: 16px;
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.1) 0%, rgba(102, 126, 234, 0.1) 100%);
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
  border-color: rgba(102, 126, 234, 0.3);
  color: #667eea;
}

.refresh-btn:hover {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.08) 0%, rgba(118, 75, 162, 0.08) 100%);
  border-color: #667eea;
  color: #5a6fd6;
}

.refresh-btn:active {
  transform: translateY(0);
}

.export-btn {
  border-color: rgba(82, 196, 26, 0.3);
  color: #52c41a;
}

.export-btn:hover {
  background: linear-gradient(135deg, rgba(82, 196, 26, 0.08) 0%, rgba(135, 208, 104, 0.08) 100%);
  border-color: #52c41a;
  color: #389e0d;
}

.export-btn:active {
  transform: translateY(0);
}

.custom-table :deep(.ant-table-thead > tr > th) {
  background: linear-gradient(135deg, #f5f7fa 0%, #e8ecf1 100%);
  font-weight: 600;
  color: #262626;
  border-bottom: 2px solid #e8e8e8;
  padding: 16px;
}

.custom-table :deep(.ant-table-tbody > tr > td) {
  padding: 16px;
  border-bottom: 1px solid #f0f0f0;
  transition: all 0.3s;
}

.custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #fafafa;
}

.custom-table :deep(.ant-table-tbody > tr) {
  transition: all 0.3s;
}

.custom-table :deep(.ant-table-tbody > tr:hover) {
  transform: scale(1.002);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.part-id-cell {
  display: flex;
  align-items: center;
}

.part-id {
  font-family: 'Monaco', 'Consolas', monospace;
  font-weight: 600;
  color: #667eea;
  background: rgba(102, 126, 234, 0.1);
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 13px;
}

.name-cell {
  display: flex;
  align-items: center;
  gap: 12px;
}

.part-type-icon {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 16px;
}

.part-name {
  font-weight: 500;
  color: #262626;
}

.type-tag {
  border-radius: 4px;
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
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}

.dot-normal { background: #52c41a; }
.dot-production { background: #1890ff; }
.dot-quality { background: #13c2c2; }
.dot-supply { background: #722ed1; }
.dot-recalled { background: #faad14; }
.dot-scrapped { background: #8c8c8c; }

.status-text {
  font-size: 13px;
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
  font-size: 14px;
}

.action-cell {
  gap: 8px;
}

.action-btn {
  padding: 4px 8px;
  border-radius: 4px;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 4px;
}

.detail-btn {
  color: #667eea;
}

.detail-btn:hover {
  background: rgba(102, 126, 234, 0.1);
  color: #5a6fd6;
}

.lifecycle-btn {
  color: #722ed1;
}

.lifecycle-btn:hover {
  background: rgba(114, 46, 209, 0.1);
  color: #642ab5;
}

/* ==================== 分页样式 ==================== */
.custom-table :deep(.ant-pagination) {
  margin: 16px 0;
}

.custom-table :deep(.ant-pagination-item) {
  border-radius: 4px;
}

.custom-table :deep(.ant-pagination-item-active) {
  border-color: #667eea;
}

.custom-table :deep(.ant-pagination-item-active a) {
  color: #667eea;
}

/* ==================== 响应式布局 ==================== */
@media (max-width: 768px) {
  .page-header {
    padding: 16px;
  }

  .header-content {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .create-btn {
    width: 100%;
  }

  .stats-overview {
    flex-wrap: wrap;
    gap: 16px;
  }

  .stat-item {
    flex: 1 1 45%;
  }

  .search-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
    padding: 16px;
  }

  .header-right {
    width: 100%;
  }

  .filter-tag {
    width: 100%;
    justify-content: center;
  }

  .search-form {
    padding: 16px;
  }

  .search-actions {
    flex-direction: column;
    gap: 16px;
    padding: 16px;
  }

  .action-left {
    width: 100%;
  }

  .action-hint {
    width: 100%;
    justify-content: center;
  }

  .action-right {
    width: 100%;
    flex-direction: column;
  }

  .search-btn,
  .reset-btn {
    width: 100%;
  }

  .table-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
    padding: 16px;
  }

  .table-header-right {
    width: 100%;
  }

  .toolbar-btn {
    flex: 1;
  }
}

/* ==================== 暗色主题适配 ==================== */
[data-theme='dark'] .page-header {
  background: linear-gradient(135deg, #1a1f3c 0%, #2d1b4e 100%);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.4);
}

[data-theme='dark'] .search-card,
[data-theme='dark'] .table-card {
  background: linear-gradient(180deg, #1f1f1f 0%, #1a1a1a 100%);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.4);
  border-color: rgba(102, 126, 234, 0.2);
}

[data-theme='dark'] .search-header {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-bottom-color: rgba(102, 126, 234, 0.2);
}

[data-theme='dark'] .header-title {
  color: #fff;
}

[data-theme='dark'] .header-desc {
  color: #8c8c8c;
}

[data-theme='dark'] .filter-label {
  color: #e8e8e8;
}

[data-theme='dark'] .search-actions {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
  border-top-color: rgba(102, 126, 234, 0.2);
}

[data-theme='dark'] .action-hint {
  background: rgba(255, 255, 255, 0.05);
  color: #8c8c8c;
}

[data-theme='dark'] .reset-btn {
  background: #2a2a2a;
  border-color: #434343;
  color: #bfbfbf;
}

[data-theme='dark'] .reset-btn:hover {
  border-color: #667eea;
  color: #667eea;
  background: rgba(102, 126, 234, 0.1);
}

[data-theme='dark'] .table-header {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
  border-bottom-color: rgba(102, 126, 234, 0.2);
}

[data-theme='dark'] .table-title {
  color: #fff;
}

[data-theme='dark'] .toolbar-btn {
  background: #2a2a2a;
  border-color: #434343;
  color: #bfbfbf;
}

[data-theme='dark'] .toolbar-btn:hover {
  border-color: #667eea;
  color: #667eea;
  background: rgba(102, 126, 234, 0.1);
}

[data-theme='dark'] .refresh-btn {
  border-color: rgba(102, 126, 234, 0.4);
  color: #8b9ef3;
}

[data-theme='dark'] .refresh-btn:hover {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.15) 0%, rgba(118, 75, 162, 0.15) 100%);
  border-color: #8b9ef3;
  color: #a8b4f5;
}

[data-theme='dark'] .export-btn {
  border-color: rgba(82, 196, 26, 0.4);
  color: #73d13d;
}

[data-theme='dark'] .export-btn:hover {
  background: linear-gradient(135deg, rgba(82, 196, 26, 0.15) 0%, rgba(135, 208, 104, 0.15) 100%);
  border-color: #73d13d;
  color: #95de64;
}

[data-theme='dark'] .record-count {
  background: linear-gradient(135deg, rgba(24, 144, 255, 0.15) 0%, rgba(102, 126, 234, 0.15) 100%);
  border-color: rgba(24, 144, 255, 0.3);
  color: #91d5ff;
}

[data-theme='dark'] .custom-table :deep(.ant-table-thead > tr > th) {
  background: linear-gradient(135deg, #2a2a2a 0%, #1f1f1f 100%);
  color: #fff;
  border-bottom-color: #333;
}

[data-theme='dark'] .custom-table :deep(.ant-table-tbody > tr > td) {
  border-bottom-color: #333;
  background: #1f1f1f;
  color: #bfbfbf;
}

[data-theme='dark'] .custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #2a2a2a;
}

[data-theme='dark'] .part-id {
  background: rgba(102, 126, 234, 0.2);
}

[data-theme='dark'] .part-name {
  color: #fff;
}

[data-theme='dark'] .status-text {
  color: #bfbfbf;
}

[data-theme='dark'] .ant-select-dropdown {
  background: #2a2a2a;
  box-shadow: 0 6px 16px 0 rgba(0, 0, 0, 0.45);
}

[data-theme='dark'] .ant-select-item {
  color: #bfbfbf;
}

[data-theme='dark'] .ant-select-item:hover {
  background: rgba(102, 126, 234, 0.15);
}

[data-theme='dark'] .ant-select-item-option-selected {
  background: rgba(102, 126, 234, 0.2);
  color: #fff;
}
</style>
