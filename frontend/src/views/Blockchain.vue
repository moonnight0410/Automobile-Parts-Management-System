<!--
  区块链浏览器页面
-->

<template>
  <div class="blockchain">
    <div class="page-header">
      <h2>区块链浏览器</h2>
    </div>
    
    <a-row :gutter="16">
      <a-col :span="8">
        <a-card title="区块高度" :bordered="false">
          <a-statistic :value="1234" suffix="个区块" />
        </a-card>
      </a-col>
      <a-col :span="8">
        <a-card title="交易数量" :bordered="false">
          <a-statistic :value="5678" suffix="笔交易" />
        </a-card>
      </a-col>
      <a-col :span="8">
        <a-card title="网络状态" :bordered="false">
          <a-statistic title="节点数" :value="4" suffix="个节点" />
        </a-card>
      </a-col>
    </a-row>
    
    <a-card title="最近区块" :bordered="false" style="margin-top: 16px">
      <a-table :columns="columns" :data-source="mockData" row-key="blockNumber" />
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const columns = [
  { title: '区块高度', dataIndex: 'blockNumber' },
  { title: '交易数', dataIndex: 'txCount' },
  { title: '区块哈希', dataIndex: 'blockHash' },
  { title: '时间戳', dataIndex: 'timestamp' }
]

const mockData = ref([
  { blockNumber: 1234, txCount: 5, blockHash: '0x1234...abcd', timestamp: '2024-01-15 10:30:00' },
  { blockNumber: 1233, txCount: 3, blockHash: '0x1233...efgh', timestamp: '2024-01-15 10:25:00' }
])
</script>

<style scoped>
.blockchain { padding: 0; }
.page-header { margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
