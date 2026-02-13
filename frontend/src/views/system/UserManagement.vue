<!--
  用户管理页面
-->

<template>
  <div class="user-management">
    <div class="page-header">
      <h2>用户管理</h2>
    </div>
    
    <a-card :bordered="false">
      <a-table :columns="columns" :data-source="mockData" row-key="id">
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'role'">
            <a-tag :color="record.role === '制造商' ? 'blue' : record.role === '车企' ? 'green' : 'orange'">
              {{ record.role }}
            </a-tag>
          </template>
          <template v-if="column.key === 'action'">
            <a-button type="link" size="small">编辑</a-button>
            <a-button type="link" size="small" danger>删除</a-button>
          </template>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const columns = [
  { title: '用户ID', dataIndex: 'id' },
  { title: '用户名', dataIndex: 'username' },
  { title: '组织', dataIndex: 'org' },
  { title: '角色', dataIndex: 'role', key: 'role' },
  { title: '注册时间', dataIndex: 'registerTime' },
  { title: '操作', key: 'action' }
]

const mockData = ref([
  { id: 'user-001', username: 'manufacturer_user', org: 'Org1MSP', role: '制造商', registerTime: '2024-01-01' },
  { id: 'user-002', username: 'automaker_user', org: 'Org2MSP', role: '车企', registerTime: '2024-01-02' },
  { id: 'user-003', username: 'aftersale_user', org: 'Org3MSP', role: '售后中心', registerTime: '2024-01-03' }
])
</script>

<style scoped>
.user-management { padding: 0; }
.page-header { margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
