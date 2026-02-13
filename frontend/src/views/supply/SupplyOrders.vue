<!--
  采购订单页面
-->

<template>
  <div class="supply-orders">
    <div class="page-header">
      <h2>采购订单</h2>
      <a-button type="primary">
        <template #icon><PlusOutlined /></template>
        创建订单
      </a-button>
    </div>
    
    <a-card :bordered="false">
      <a-table :columns="columns" :data-source="mockData" row-key="orderID">
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'status'">
            <a-tag :color="record.status === '已签收' ? 'green' : 'blue'">{{ record.status }}</a-tag>
          </template>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const columns = [
  { title: '订单ID', dataIndex: 'orderID' },
  { title: '采购方', dataIndex: 'buyer' },
  { title: '销售方', dataIndex: 'seller' },
  { title: '零部件ID', dataIndex: 'partID' },
  { title: '数量', dataIndex: 'quantity' },
  { title: '状态', dataIndex: 'status', key: 'status' },
  { title: '创建时间', dataIndex: 'createTime' }
]

const mockData = ref([
  { orderID: 'ORDER-001', buyer: '车企A', seller: '制造商B', partID: 'PART-001', quantity: 100, status: '已签收', createTime: '2024-01-15' }
])
</script>

<style scoped>
.supply-orders { padding: 0; }
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
