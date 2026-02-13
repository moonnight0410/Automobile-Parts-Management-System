<!--
  BOM列表页面
  展示BOM列表，支持搜索和查看详情
-->

<template>
  <div class="bom-list">
    <div class="page-header">
      <h2>BOM列表</h2>
      <a-button type="primary" @click="goCreate">
        <template #icon><PlusOutlined /></template>
        创建BOM
      </a-button>
    </div>
    
    <a-card :bordered="false">
      <a-table :columns="columns" :data-source="mockData" row-key="bomID">
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'status'">
            <a-tag :color="record.status === '已发布' ? 'green' : 'blue'">{{ record.status }}</a-tag>
          </template>
          <template v-if="column.key === 'action'">
            <a-button type="link" size="small">查看详情</a-button>
          </template>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const columns = [
  { title: 'BOM ID', dataIndex: 'bomID', key: 'bomID' },
  { title: '类型', dataIndex: 'bomType', key: 'bomType' },
  { title: '车型', dataIndex: 'productModel', key: 'productModel' },
  { title: '版本', dataIndex: 'version', key: 'version' },
  { title: '状态', dataIndex: 'status', key: 'status' },
  { title: '创建时间', dataIndex: 'createTime', key: 'createTime' },
  { title: '操作', key: 'action' }
]

const mockData = ref([
  { bomID: 'BOM-001', bomType: '生产BOM', productModel: 'Model-A', version: 'v1.0', status: '已发布', createTime: '2024-01-15' },
  { bomID: 'BOM-002', bomType: '研发BOM', productModel: 'Model-B', version: 'v2.0', status: '草稿', createTime: '2024-01-16' }
])

function goCreate() {
  router.push('/bom/create')
}
</script>

<style scoped>
.bom-list { padding: 0; }
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
