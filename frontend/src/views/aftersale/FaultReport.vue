<!--
  故障报告页面
-->

<template>
  <div class="fault-report">
    <div class="page-header">
      <h2>故障报告</h2>
      <a-button type="primary">
        <template #icon><PlusOutlined /></template>
        上报故障
      </a-button>
    </div>
    
    <a-card :bordered="false">
      <a-table :columns="columns" :data-source="mockData" row-key="faultID">
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'status'">
            <a-tag :color="record.status === '已审核' ? 'green' : 'orange'">{{ record.status }}</a-tag>
          </template>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const columns = [
  { title: '故障ID', dataIndex: 'faultID' },
  { title: '零部件ID', dataIndex: 'partID' },
  { title: 'VIN码', dataIndex: 'vin' },
  { title: '故障类型', dataIndex: 'faultType' },
  { title: '风险概率', dataIndex: 'riskProb' },
  { title: '状态', dataIndex: 'status', key: 'status' },
  { title: '上报时间', dataIndex: 'reportTime' }
]

const mockData = ref([
  { faultID: 'FAULT-001', partID: 'PART-001', vin: 'VIN-001', faultType: '制动故障', riskProb: '85%', status: '待审核', reportTime: '2024-01-15' }
])
</script>

<style scoped>
.fault-report { padding: 0; }
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
