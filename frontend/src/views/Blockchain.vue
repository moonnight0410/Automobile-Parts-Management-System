<!--
  区块链浏览器页面
  设计：与Dashboard风格统一
-->

<template>
  <div class="blockchain">
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <LinkOutlined />
          </div>
          <div class="title-text">
            <h2>区块链浏览器</h2>
            <p class="subtitle">查看区块链网络状态与交易记录</p>
          </div>
        </div>
        <div class="header-actions">
          <a-button class="action-btn" @click="handleRefresh" :loading="refreshing">
            <template #icon><ReloadOutlined :spin="refreshing" /></template>
            刷新状态
          </a-button>
        </div>
      </div>
    </div>

    <a-row :gutter="20" class="stats-row">
      <a-col :span="6">
        <a-card :bordered="false" class="stat-card block-card">
          <div class="stat-content">
            <div class="stat-icon-wrapper">
              <BlockOutlined class="stat-icon" />
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ blockHeight }}</div>
              <div class="stat-label">区块高度</div>
            </div>
          </div>
          <div class="stat-footer">
            <span class="trend up"><ArrowUpOutlined />+12 今日</span>
          </div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card :bordered="false" class="stat-card tx-card">
          <div class="stat-content">
            <div class="stat-icon-wrapper">
              <TransactionOutlined class="stat-icon" />
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ txCount }}</div>
              <div class="stat-label">交易数量</div>
            </div>
          </div>
          <div class="stat-footer">
            <span class="trend up"><ArrowUpOutlined />+56 今日</span>
          </div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card :bordered="false" class="stat-card node-card">
          <div class="stat-content">
            <div class="stat-icon-wrapper">
              <ClusterOutlined class="stat-icon" />
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ nodeCount }}</div>
              <div class="stat-label">网络节点</div>
            </div>
          </div>
          <div class="stat-footer">
            <span class="status online"><CheckCircleOutlined /> 全部在线</span>
          </div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card :bordered="false" class="stat-card channel-card">
          <div class="stat-content">
            <div class="stat-icon-wrapper">
              <ApiOutlined class="stat-icon" />
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ channelCount }}</div>
              <div class="stat-label">通道数量</div>
            </div>
          </div>
          <div class="stat-footer">
            <span class="status active"><ThunderboltOutlined /> 运行中</span>
          </div>
        </a-card>
      </a-col>
    </a-row>

    <a-card :bordered="false" class="search-card">
      <div class="search-header">
        <div class="header-left">
          <div class="filter-icon-wrapper">
            <SearchOutlined />
          </div>
          <div class="header-info">
            <span class="header-title">区块/交易查询</span>
            <span class="header-desc">输入区块高度或交易哈希进行查询</span>
          </div>
        </div>
      </div>
      <div class="search-content">
        <a-input-search
          v-model:value="searchQuery"
          placeholder="输入区块高度、区块哈希或交易ID..."
          enter-button="搜索"
          size="large"
          class="search-input"
          @search="handleSearch"
        />
        <div class="quick-search">
          <span class="quick-label">快捷查询：</span>
          <a-tag v-for="tag in quickTags" :key="tag" class="quick-tag" @click="searchQuery = tag; handleSearch(tag)">
            {{ tag }}
          </a-tag>
        </div>
      </div>
    </a-card>

    <a-card :bordered="false" class="table-card">
      <div class="table-header">
        <div class="table-header-left">
          <span class="table-title">
            <UnorderedListOutlined class="title-icon" />
            最近区块
          </span>
          <a-tag color="blue" class="record-count">
            <DatabaseOutlined />
            实时更新
          </a-tag>
        </div>
        <div class="table-header-right">
          <a-button class="toolbar-btn export-btn" @click="handleExport">
            <template #icon><ExportOutlined /></template>
            导出数据
          </a-button>
        </div>
      </div>
      <a-table
        :columns="columns"
        :data-source="mockData"
        row-key="blockNumber"
        class="custom-table"
        :scroll="{ x: 1000 }"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'blockNumber'">
            <div class="block-number-cell">
              <BlockOutlined class="block-icon" />
              <span>{{ record.blockNumber }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'txCount'">
            <a-tag color="blue" class="tx-tag">
              <TransactionOutlined />
              {{ record.txCount }} 笔
            </a-tag>
          </template>
          <template v-else-if="column.key === 'blockHash'">
            <div class="hash-cell">
              <span class="hash-text">{{ record.blockHash }}</span>
              <a-button type="link" size="small" class="copy-btn" @click="copyHash(record.blockHash)">
                <CopyOutlined />
              </a-button>
            </div>
          </template>
          <template v-else-if="column.key === 'timestamp'">
            <div class="time-cell">
              <ClockCircleOutlined class="time-icon" />
              <span>{{ record.timestamp }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'action'">
            <a-space class="action-cell">
              <a-button type="link" size="small" @click="viewBlock(record)" class="action-btn view-btn">
                <EyeOutlined />
                查看
              </a-button>
              <a-button type="link" size="small" @click="viewTransactions(record)" class="action-btn tx-btn">
                <TransactionOutlined />
                交易
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <a-card :bordered="false" class="network-card">
      <div class="network-header">
        <span class="network-title">
          <ClusterOutlined class="title-icon" />
          网络状态
        </span>
      </div>
      <a-row :gutter="20">
        <a-col :span="8" v-for="org in organizations" :key="org.name">
          <div class="org-card">
            <div class="org-header">
              <div class="org-icon" :style="{ background: org.color }">
                <TeamOutlined />
              </div>
              <div class="org-info">
                <div class="org-name">{{ org.name }}</div>
                <div class="org-status">
                  <span class="status-dot" :class="org.status"></span>
                  {{ org.statusText }}
                </div>
              </div>
            </div>
            <div class="org-stats">
              <div class="org-stat">
                <span class="stat-label">节点数</span>
                <span class="stat-value">{{ org.nodes }}</span>
              </div>
              <div class="org-stat">
                <span class="stat-label">交易数</span>
                <span class="stat-value">{{ org.txs }}</span>
              </div>
            </div>
          </div>
        </a-col>
      </a-row>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { message } from 'ant-design-vue'
import { getBlockchainInfo, getRecentBlocks, type BlockchainInfo, type BlockInfo } from '../services/blockchain.service'

const columns = [
  { title: '区块高度', dataIndex: 'blockNumber', key: 'blockNumber', width: 120 },
  { title: '交易数', dataIndex: 'txCount', key: 'txCount', width: 100 },
  { title: '区块哈希', dataIndex: 'blockHash', key: 'blockHash', width: 300 },
  { title: '时间戳', dataIndex: 'timestamp', key: 'timestamp', width: 180 },
  { title: '操作', key: 'action', fixed: 'right', width: 140 }
]

const mockData = ref<BlockInfo[]>([])
const organizations = ref([
  { name: '制造商组织 (Org1MSP)', color: 'linear-gradient(135deg, #6366f1 0%, #4f46e5 100%)', status: 'online', statusText: '在线', nodes: 2, txs: 18234 },
  { name: '车企组织 (Org2MSP)', color: 'linear-gradient(135deg, #10b981 0%, #059669 100%)', status: 'online', statusText: '在线', nodes: 1, txs: 14567 },
  { name: '售后组织 (Org3MSP)', color: 'linear-gradient(135deg, #f59e0b 0%, #d97706 100%)', status: 'online', statusText: '在线', nodes: 1, txs: 12877 }
])

const quickTags = ref(['15234', '15233', '0x0000...'])

const blockHeight = ref(0)
const txCount = ref(0)
const nodeCount = ref(0)
const channelCount = ref(0)
const searchQuery = ref('')
const refreshing = ref(false)
let refreshInterval: number | null = null

async function fetchData() {
  try {
    const [infoRes, blocksRes] = await Promise.all([
      getBlockchainInfo(),
      getRecentBlocks(10)
    ])

    if (infoRes.success && infoRes.data) {
      blockHeight.value = infoRes.data.blockHeight
      txCount.value = infoRes.data.txCount
      nodeCount.value = infoRes.data.nodeCount
      channelCount.value = infoRes.data.channelCount
    }

    if (blocksRes.success && blocksRes.data) {
      mockData.value = blocksRes.data
    }
  } catch (error: any) {
    console.error('获取区块链数据失败:', error)
    message.error('获取区块链数据失败')
  }
}

function handleRefresh() {
  refreshing.value = true
  fetchData().finally(() => {
    refreshing.value = false
  })
}

function handleSearch(value: string) {
  if (!value) {
    message.warning('请输入查询内容')
    return
  }
  message.info(`正在查询: ${value}`)
}

function handleExport() {
  message.success('数据导出成功')
}

function copyHash(hash: string) {
  navigator.clipboard.writeText(hash)
  message.success('哈希已复制到剪贴板')
}

function viewBlock(record: any) {
  message.info(`查看区块 ${record.blockNumber} 详情`)
}

function viewTransactions(record: any) {
  message.info(`查看区块 ${record.blockNumber} 的交易列表`)
}

onMounted(() => {
  fetchData()
  refreshInterval = window.setInterval(fetchData, 30000)
})

onUnmounted(() => {
  if (refreshInterval) {
    clearInterval(refreshInterval)
  }
})
</script>

<style scoped>
.blockchain {
  padding: 0;
}

.page-header {
  margin-bottom: 24px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-title {
  display: flex;
  align-items: center;
  gap: 16px;
}

.title-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  color: #fff;
  box-shadow: 0 4px 12px rgba(139, 92, 246, 0.35);
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
  color: #8b5cf6;
  border-color: #8b5cf6;
}

.action-btn:hover {
  color: #7c3aed;
  border-color: #7c3aed;
}

.stats-row {
  margin-bottom: 20px;
}

.stat-card {
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  overflow: hidden;
}

.stat-content {
  display: flex;
  align-items: center;
  gap: 16px;
}

.stat-icon-wrapper {
  width: 56px;
  height: 56px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  color: #fff;
}

.block-card .stat-icon-wrapper {
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
}

.tx-card .stat-icon-wrapper {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.node-card .stat-icon-wrapper {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
}

.channel-card .stat-icon-wrapper {
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%);
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 28px;
  font-weight: 700;
  color: #1e293b;
  margin-bottom: 4px;
}

.stat-label {
  font-size: 14px;
  color: #64748b;
}

.stat-footer {
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px solid #f0f0f0;
}

.trend {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 13px;
}

.trend.up {
  color: #10b981;
}

.status {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 13px;
}

.status.online {
  color: #10b981;
}

.status.active {
  color: #f59e0b;
}

.search-card {
  border-radius: 12px;
  margin-bottom: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.search-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.filter-icon-wrapper {
  width: 36px;
  height: 36px;
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%);
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

.search-content {
  padding: 8px 0;
}

.search-input {
  margin-bottom: 12px;
}

.search-input :deep(.ant-input) {
  border-radius: 10px;
  height: 48px;
}

.search-input :deep(.ant-input-search-button) {
  border-radius: 0 10px 10px 0;
  height: 48px;
}

.quick-search {
  display: flex;
  align-items: center;
  gap: 8px;
}

.quick-label {
  font-size: 13px;
  color: #6b7280;
}

.quick-tag {
  cursor: pointer;
  transition: all 0.3s ease;
}

.quick-tag:hover {
  background: #8b5cf6;
  color: #fff;
  border-color: #8b5cf6;
}

.table-card {
  border-radius: 12px;
  margin-bottom: 20px;
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
  color: #8b5cf6;
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
  color: #8b5cf6;
  border-color: #8b5cf6;
}

.custom-table :deep(.ant-table-thead > tr > th) {
  background: linear-gradient(135deg, #f5f3ff 0%, #ede9fe 100%);
  font-weight: 600;
  color: #374151;
}

.custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #f5f3ff;
}

.block-number-cell {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
  color: #8b5cf6;
}

.block-icon {
  font-size: 14px;
}

.tx-tag {
  display: flex;
  align-items: center;
  gap: 4px;
}

.hash-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.hash-text {
  font-family: 'Monaco', 'Menlo', monospace;
  font-size: 12px;
  color: #6b7280;
}

.copy-btn {
  padding: 0 4px;
  height: auto;
  color: #8b5cf6;
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

.view-btn {
  color: #8b5cf6;
}

.tx-btn {
  color: #10b981;
}

.network-card {
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.network-header {
  margin-bottom: 20px;
}

.network-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 15px;
  font-weight: 600;
  color: var(--text-color);
}

.org-card {
  padding: 16px;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border-radius: 12px;
  border: 1px solid #e2e8f0;
}

.org-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}

.org-icon {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 18px;
}

.org-info {
  flex: 1;
}

.org-name {
  font-weight: 600;
  color: #1e293b;
  margin-bottom: 4px;
}

.org-status {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: #6b7280;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.status-dot.online {
  background: #10b981;
}

.org-stats {
  display: flex;
  gap: 20px;
}

.org-stat {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.org-stat .stat-label {
  font-size: 12px;
  color: #6b7280;
}

.org-stat .stat-value {
  font-size: 18px;
  font-weight: 600;
  color: #1e293b;
}

/* 深色主题 */
[data-theme='dark'] .title-icon {
  box-shadow: 0 4px 12px rgba(139, 92, 246, 0.45);
}

[data-theme='dark'] .stat-value {
  color: #f1f5f9;
}

[data-theme='dark'] .stat-footer {
  border-top-color: #374151;
}

[data-theme='dark'] .custom-table :deep(.ant-table-thead > tr > th) {
  background: linear-gradient(135deg, #1e293b 0%, #0f172a 100%);
  color: #e5e7eb;
}

[data-theme='dark'] .custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #1e293b;
}

[data-theme='dark'] .org-card {
  background: linear-gradient(135deg, #1e293b 0%, #0f172a 100%);
  border-color: #334155;
}

[data-theme='dark'] .org-name {
  color: #f1f5f9;
}

[data-theme='dark'] .org-stat .stat-value {
  color: #f1f5f9;
}
</style>
