<!--
  生产数据页面
-->

<template>
  <div class="production-data">
    <div class="page-header">
      <h2>生产数据</h2>
      <a-button type="primary">
        <template #icon><PlusOutlined /></template>
        录入生产数据
      </a-button>
    </div>
    
    <a-card :bordered="false">
      <a-table :columns="columns" :data-source="mockData" row-key="productionID" />
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const columns = [
  { title: '生产ID', dataIndex: 'productionID' },
  { title: '零部件ID', dataIndex: 'partID' },
  { title: '批次号', dataIndex: 'batchNo' },
  { title: '生产线', dataIndex: 'productionLine' },
  { title: '操作员', dataIndex: 'operator' },
  { title: '完成时间', dataIndex: 'finishTime' }
]

const mockData = ref([
  { productionID: 'PROD-001', partID: 'PART-001', batchNo: 'BATCH-001', productionLine: 'A线', operator: '张三', finishTime: '2024-01-15' }
])
</script>

<style scoped>
.production-data { padding: 0; }
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
