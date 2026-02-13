<!--
  召回记录页面
-->

<template>
  <div class="recall-record">
    <div class="page-header">
      <h2>召回记录</h2>
      <a-button type="primary">
        <template #icon><PlusOutlined /></template>
        发起召回
      </a-button>
    </div>
    
    <a-card :bordered="false">
      <a-table :columns="columns" :data-source="mockData" row-key="recallID">
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'status'">
            <a-tag :color="record.status === '已完成' ? 'green' : 'blue'">{{ record.status }}</a-tag>
          </template>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const columns = [
  { title: '召回ID', dataIndex: 'recallID' },
  { title: '涉及批次', dataIndex: 'batchNos' },
  { title: '召回原因', dataIndex: 'reason' },
  { title: '受影响零部件', dataIndex: 'affectedCount' },
  { title: '状态', dataIndex: 'status', key: 'status' },
  { title: '发起时间', dataIndex: 'createTime' }
]

const mockData = ref([
  { recallID: 'RECALL-001', batchNos: 'BATCH-001, BATCH-002', reason: '制动系统潜在隐患', affectedCount: 150, status: '进行中', createTime: '2024-01-15' }
])
</script>

<style scoped>
.recall-record { padding: 0; }
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
