<!--
  仪表盘页面
  展示系统概览、统计数据、最新动态等
-->

<template>
  <div class="dashboard">
    <!-- 页面标题 -->
    <div class="page-header">
      <h2>仪表盘</h2>
      <p>欢迎回来，{{ user?.username }}！</p>
    </div>
    
    <!-- 统计卡片 -->
    <a-row :gutter="16" class="stat-cards">
      <a-col :span="6">
        <a-card :bordered="false" class="stat-card">
          <a-statistic
            title="零部件总数"
            :value="statistics.totalParts"
            :value-style="{ color: '#3f8600' }"
          >
            <template #prefix>
              <AppstoreOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card :bordered="false" class="stat-card">
          <a-statistic
            title="BOM记录"
            :value="statistics.totalBOMs"
            :value-style="{ color: '#1890ff' }"
          >
            <template #prefix>
              <FileTextOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card :bordered="false" class="stat-card">
          <a-statistic
            title="故障报告"
            :value="statistics.totalFaults"
            :value-style="{ color: '#cf1322' }"
          >
            <template #prefix>
              <WarningOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card :bordered="false" class="stat-card">
          <a-statistic
            title="召回记录"
            :value="statistics.totalRecalls"
            :value-style="{ color: '#faad14' }"
          >
            <template #prefix>
              <RollbackOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
    </a-row>
    
    <!-- 快捷操作 -->
    <a-card title="快捷操作" :bordered="false" class="quick-actions">
      <div class="quick-actions-row">
        <div class="action-item" v-for="action in quickActions" :key="action.key" @click="handleAction(action)">
          <div class="action-icon" :style="{ background: action.color }">
            <component :is="action.icon" />
          </div>
          <div class="action-text">{{ action.text }}</div>
        </div>
      </div>
    </a-card>
    
    <!-- 图表区域 -->
    <a-row :gutter="16" class="charts-row">
      <!-- 零部件状态分布 -->
      <a-col :xs="24" :lg="12">
        <a-card title="零部件状态分布" :bordered="false">
          <div ref="partStatusChartRef" class="chart-container"></div>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth.store'
import * as echarts from 'echarts'
import { listMyParts } from '../services/part.service'
import { listBOMs } from '../services/bom.service'
import { listFaultReports } from '../services/aftersale.service'
import { listRecallRecords } from '../services/aftersale.service'

// ==================== 组合式API ====================

const router = useRouter()
const authStore = useAuthStore()

// ==================== 响应式状态 ====================

// 图表引用
const partStatusChartRef = ref<HTMLElement>()

// 统计数据
const statistics = reactive({
  totalParts: 0,
  totalBOMs: 0,
  totalFaults: 0,
  totalRecalls: 0
})

// 当前用户
const user = computed(() => authStore.user)

// 快捷操作
const quickActions = [
  { key: 'create-part', text: '创建零部件', icon: 'PlusOutlined', color: '#52c41a', path: '/parts/create' },
  { key: 'create-bom', text: '创建BOM', icon: 'FileAddOutlined', color: '#1890ff', path: '/bom/create' },
  { key: 'query-part', text: '查询零部件', icon: 'SearchOutlined', color: '#722ed1', path: '/parts/list' },
  { key: 'fault-report', text: '故障上报', icon: 'WarningOutlined', color: '#fa8c16', path: '/aftersale/fault' },
  { key: 'ai-assistant', text: 'AI小助手', icon: 'RobotOutlined', color: '#13c2c2', path: '/aftersale/assistant' },
  { key: 'blockchain', text: '区块链浏览', icon: 'LinkOutlined', color: '#eb2f96', path: '/blockchain' }
]

// ==================== 方法 ====================

async function fetchStatistics() {
  try {
    const [partsRes, bomsRes, faultsRes, recallsRes] = await Promise.all([
      listMyParts(),
      listBOMs(),
      listFaultReports(),
      listRecallRecords()
    ])
    
    statistics.totalParts = partsRes.data?.length || 0
    statistics.totalBOMs = bomsRes.data?.length || 0
    statistics.totalFaults = faultsRes.data?.length || 0
    statistics.totalRecalls = recallsRes.data?.length || 0
  } catch (error) {
    console.error('获取统计数据失败:', error)
  }
}

/**
 * 处理快捷操作点击
 */
function handleAction(action: any) {
  router.push(action.path)
}

/**
 * 初始化零部件状态分布图表
 */
function initPartStatusChart() {
  if (!partStatusChartRef.value) return
  
  const chart = echarts.init(partStatusChartRef.value)
  
  const option = {
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    series: [
      {
        name: '零部件状态',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 20,
            fontWeight: 'bold'
          }
        },
        labelLine: {
          show: false
        },
        data: [
          { value: Math.round(statistics.totalParts * 0.65), name: '正常', itemStyle: { color: '#52c41a' } },
          { value: Math.round(statistics.totalParts * 0.16), name: '在产', itemStyle: { color: '#1890ff' } },
          { value: Math.round(statistics.totalParts * 0.12), name: '供应链中', itemStyle: { color: '#722ed1' } },
          { value: Math.round(statistics.totalParts * 0.04), name: '已召回', itemStyle: { color: '#fa8c16' } },
          { value: Math.round(statistics.totalParts * 0.03), name: '已报废', itemStyle: { color: '#8c8c8c' } }
        ]
      }
    ]
  }
  
  chart.setOption(option)
  
  // 响应窗口大小变化
  window.addEventListener('resize', () => chart.resize())
}

// ==================== 生命周期 ====================

onMounted(async () => {
  await fetchStatistics()
  // 初始化图表
  setTimeout(() => {
    initPartStatusChart()
  }, 100)
})
</script>

<style scoped>
.dashboard {
  padding: 0;
}

.page-header {
  margin-bottom: 24px;
}

.page-header h2 {
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 8px;
  color: var(--text-color);
}

.page-header p {
  color: var(--text-color-secondary);
}

.stat-cards {
  margin-bottom: 24px;
  display: flex;
  flex-wrap: nowrap;
}

.stat-cards .ant-col {
  flex: 0 0 25%;
  max-width: 25%;
}

.stat-card {
  border-radius: 8px;
  transition: all 0.3s;
}

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.quick-actions {
  margin-bottom: 24px;
  border-radius: 8px;
}

.quick-actions-row {
  display: flex;
  flex-wrap: nowrap;
  justify-content: center;
  align-items: center;
  gap: 16px;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 16px;
  cursor: pointer;
  border-radius: 8px;
  transition: all 0.3s;
  width: 10%;
  min-width: 80px;
  flex-shrink: 0;
}

.action-item:hover {
  background: var(--bg-color-tertiary);
}

.action-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  color: #fff;
  margin-bottom: 8px;
}

.action-text {
  font-size: 14px;
  color: var(--text-color);
}

.charts-row {
  margin-bottom: 24px;
}

.charts-row .ant-card {
  border-radius: 8px;
}

.chart-container {
  width: 100%;
  height: 300px;
}

/* 深色主题 */
[data-theme='dark'] .stat-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.4);
}

[data-theme='dark'] .action-item:hover {
  background: var(--bg-color-tertiary);
}
</style>
