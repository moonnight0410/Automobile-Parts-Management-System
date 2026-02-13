<!--
  物流跟踪页面
-->

<template>
  <div class="logistics">
    <div class="page-header">
      <h2>物流跟踪</h2>
    </div>
    
    <a-card :bordered="false">
      <a-table :columns="columns" :data-source="mockData" row-key="logisticsID">
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'action'">
            <a-button type="link" size="small">查看轨迹</a-button>
          </template>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const columns = [
  { title: '物流ID', dataIndex: 'logisticsID' },
  { title: '订单ID', dataIndex: 'orderID' },
  { title: '物流商', dataIndex: 'carrier' },
  { title: '出发时间', dataIndex: 'startTime' },
  { title: '送达时间', dataIndex: 'endTime' },
  { title: '签收人', dataIndex: 'receiver' },
  { title: '操作', key: 'action' }
]

const mockData = ref([
  { logisticsID: 'LOG-001', orderID: 'ORDER-001', carrier: '顺丰物流', startTime: '2024-01-15', endTime: '2024-01-17', receiver: '王五' }
])
</script>

<style scoped>
.logistics { padding: 0; }
.page-header { margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
