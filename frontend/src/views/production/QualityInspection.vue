<!--
  质检管理页面
-->

<template>
  <div class="quality-inspection">
    <div class="page-header">
      <h2>质检管理</h2>
      <a-button type="primary">
        <template #icon><PlusOutlined /></template>
        录入质检数据
      </a-button>
    </div>
    
    <a-card :bordered="false">
      <a-table :columns="columns" :data-source="mockData" row-key="inspectionID">
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'result'">
            <a-tag :color="record.result === '合格' ? 'green' : 'red'">{{ record.result }}</a-tag>
          </template>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const columns = [
  { title: '质检ID', dataIndex: 'inspectionID' },
  { title: '零部件ID', dataIndex: 'partID' },
  { title: '批次号', dataIndex: 'batchNo' },
  { title: '质检结果', dataIndex: 'result', key: 'result' },
  { title: '处理人', dataIndex: 'handler' },
  { title: '处理时间', dataIndex: 'handleTime' }
]

const mockData = ref([
  { inspectionID: 'QC-001', partID: 'PART-001', batchNo: 'BATCH-001', result: '合格', handler: '李四', handleTime: '2024-01-15' },
  { inspectionID: 'QC-002', partID: 'PART-002', batchNo: 'BATCH-001', result: '不合格', handler: '李四', handleTime: '2024-01-15' }
])
</script>

<style scoped>
.quality-inspection { padding: 0; }
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
