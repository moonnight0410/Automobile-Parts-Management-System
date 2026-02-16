<!--
  用户管理页面
-->

<template>
  <div class="user-management">
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <TeamOutlined />
          </div>
          <div class="title-text">
            <h2>用户管理</h2>
            <p class="subtitle">管理系统用户与权限分配</p>
          </div>
        </div>
        <a-button type="primary" class="create-btn" @click="showCreateModal = true">
          <template #icon><UserAddOutlined /></template>
          添加用户
        </a-button>
      </div>
      <div class="stats-overview">
        <div class="stat-item">
          <div class="stat-value">{{ totalCount }}</div>
          <div class="stat-label">总用户数</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-manufacturer">{{ manufacturerCount }}</div>
          <div class="stat-label">制造商</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-automaker">{{ automakerCount }}</div>
          <div class="stat-label">车企</div>
        </div>
        <div class="stat-item">
          <div class="stat-value status-aftersale">{{ aftersaleCount }}</div>
          <div class="stat-label">售后中心</div>
        </div>
      </div>
    </div>

    <a-card :bordered="false" class="search-card">
      <div class="search-header">
        <div class="header-left">
          <div class="filter-icon-wrapper">
            <FilterOutlined />
          </div>
          <div class="header-info">
            <span class="header-title">筛选条件</span>
            <span class="header-desc">支持多条件组合查询</span>
          </div>
        </div>
        <div class="header-right">
          <a-tag color="cyan" class="filter-tag">
            <template #icon><SearchOutlined /></template>
            用户查询
          </a-tag>
        </div>
      </div>
      <a-form layout="inline" :model="searchForm" class="search-form">
        <a-row :gutter="[20, 20]" style="width: 100%">
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <IdcardOutlined class="label-icon" />
                <span>用户ID</span>
              </div>
              <a-input
                v-model:value="searchForm.userID"
                placeholder="请输入用户ID"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <UserOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <UserOutlined class="label-icon" />
                <span>用户名</span>
              </div>
              <a-input
                v-model:value="searchForm.username"
                placeholder="请输入用户名"
                allow-clear
                class="filter-input"
              >
                <template #prefix>
                  <UserOutlined class="input-prefix-icon" />
                </template>
              </a-input>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <ApartmentOutlined class="label-icon" />
                <span>组织</span>
              </div>
              <a-select
                v-model:value="searchForm.org"
                placeholder="请选择组织"
                allow-clear
                class="filter-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                :dropdown-style="{ zIndex: 1050 }"
              >
                <a-select-option value="Org1MSP">Org1MSP</a-select-option>
                <a-select-option value="Org2MSP">Org2MSP</a-select-option>
                <a-select-option value="Org3MSP">Org3MSP</a-select-option>
              </a-select>
            </div>
          </a-col>
          <a-col :xs="24" :sm="12" :md="6">
            <div class="filter-item">
              <div class="filter-label">
                <SafetyOutlined class="label-icon" />
                <span>角色</span>
              </div>
              <a-select
                v-model:value="searchForm.role"
                placeholder="请选择角色"
                allow-clear
                class="filter-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                :dropdown-style="{ zIndex: 1050 }"
              >
                <a-select-option value="制造商">
                  <span class="status-option"><span class="status-dot dot-manufacturer"></span>制造商</span>
                </a-select-option>
                <a-select-option value="车企">
                  <span class="status-option"><span class="status-dot dot-automaker"></span>车企</span>
                </a-select-option>
                <a-select-option value="售后中心">
                  <span class="status-option"><span class="status-dot dot-aftersale"></span>售后中心</span>
                </a-select-option>
              </a-select>
            </div>
          </a-col>
        </a-row>
        <div class="search-actions">
          <div class="action-left">
            <span class="action-hint">
              <InfoCircleOutlined />
              可同时使用多个条件进行精确筛选
            </span>
          </div>
          <div class="action-right">
            <a-button @click="handleReset" class="reset-btn">
              <template #icon><ReloadOutlined /></template>
              重置
            </a-button>
            <a-button type="primary" @click="handleSearch" class="search-btn">
              <template #icon><SearchOutlined /></template>
              开始搜索
            </a-button>
          </div>
        </div>
      </a-form>
    </a-card>

    <a-card :bordered="false" class="table-card">
      <div class="table-header">
        <div class="table-header-left">
          <span class="table-title">
            <UnorderedListOutlined class="title-icon" />
            数据列表
          </span>
          <a-tag color="blue" class="record-count">
            <DatabaseOutlined />
            共 {{ tableData.length }} 条记录
          </a-tag>
        </div>
        <div class="table-header-right">
          <a-space :size="12">
            <a-button class="toolbar-btn refresh-btn" @click="handleRefresh" :loading="refreshing">
              <template #icon><ReloadOutlined :spin="refreshing" /></template>
              {{ refreshing ? '刷新中...' : '刷新数据' }}
            </a-button>
            <a-button class="toolbar-btn export-btn" @click="handleExport">
              <template #icon><ExportOutlined /></template>
              导出数据
            </a-button>
          </a-space>
        </div>
      </div>
      <a-table
        :columns="columns"
        :data-source="tableData"
        row-key="id"
        class="custom-table"
        :scroll="{ x: 1000 }"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'id'">
            <div class="user-id-cell">
              <IdcardOutlined class="user-icon" />
              <span>{{ record.id }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'username'">
            <div class="username-cell">
              <a-avatar :style="{ backgroundColor: getAvatarColor(record.role) }" class="user-avatar">
                {{ record.username.charAt(0).toUpperCase() }}
              </a-avatar>
              <span class="username">{{ record.username }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'org'">
            <a-tag color="geekblue" class="org-tag">
              <ApartmentOutlined />
              {{ record.org }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'role'">
            <a-tag :color="getRoleColor(record.role)" class="role-tag">
              {{ record.role }}
            </a-tag>
          </template>
          <template v-else-if="column.key === 'status'">
            <div class="status-cell">
              <span class="status-dot" :class="record.status === '在线' ? 'dot-online' : 'dot-offline'"></span>
              <span class="status-text">{{ record.status }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'registerTime'">
            <div class="time-cell">
              <ClockCircleOutlined class="time-icon" />
              <span>{{ record.registerTime }}</span>
            </div>
          </template>
          <template v-else-if="column.key === 'action'">
            <a-space class="action-cell">
              <a-button type="link" size="small" @click="viewDetail(record)" class="action-btn detail-btn">
                <EyeOutlined />
                详情
              </a-button>
              <a-button type="link" size="small" @click="editUser(record)" class="action-btn edit-btn">
                <EditOutlined />
                编辑
              </a-button>
              <a-popconfirm
                title="确定要删除该用户吗？"
                ok-text="确定"
                cancel-text="取消"
                @confirm="deleteUser(record)"
              >
                <a-button type="link" size="small" danger class="action-btn delete-btn">
                  <DeleteOutlined />
                  删除
                </a-button>
              </a-popconfirm>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <a-modal
      v-model:open="showCreateModal"
      title="添加用户"
      :width="550"
      @ok="handleCreateUser"
      @cancel="showCreateModal = false"
      class="create-modal"
    >
      <a-form :model="createForm" layout="vertical" class="create-form">
        <a-form-item label="用户名" required>
          <a-input v-model:value="createForm.username" placeholder="请输入用户名" />
        </a-form-item>
        <a-form-item label="所属组织" required>
          <a-select v-model:value="createForm.org" placeholder="请选择组织">
            <a-select-option value="Org1MSP">Org1MSP</a-select-option>
            <a-select-option value="Org2MSP">Org2MSP</a-select-option>
            <a-select-option value="Org3MSP">Org3MSP</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="角色" required>
          <a-select v-model:value="createForm.role" placeholder="请选择角色">
            <a-select-option value="制造商">制造商</a-select-option>
            <a-select-option value="车企">车企</a-select-option>
            <a-select-option value="售后中心">售后中心</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="邮箱">
          <a-input v-model:value="createForm.email" placeholder="请输入邮箱地址" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { message } from 'ant-design-vue'
import {
  TeamOutlined,
  UserAddOutlined,
  FilterOutlined,
  SearchOutlined,
  IdcardOutlined,
  UserOutlined,
  ApartmentOutlined,
  SafetyOutlined,
  InfoCircleOutlined,
  ReloadOutlined,
  UnorderedListOutlined,
  DatabaseOutlined,
  ExportOutlined,
  ClockCircleOutlined,
  EyeOutlined,
  EditOutlined,
  DeleteOutlined
} from '@ant-design/icons-vue'

const refreshing = ref(false)
const showCreateModal = ref(false)

const searchForm = ref({
  userID: '',
  username: '',
  org: undefined as string | undefined,
  role: undefined as string | undefined
})

const createForm = ref({
  username: '',
  org: '',
  role: '',
  email: ''
})

const columns = [
  { title: '用户ID', key: 'id', width: 120 },
  { title: '用户名', key: 'username', width: 180 },
  { title: '组织', key: 'org', width: 130 },
  { title: '角色', key: 'role', width: 100 },
  { title: '状态', key: 'status', width: 100 },
  { title: '注册时间', key: 'registerTime', width: 150 },
  { title: '操作', key: 'action', width: 180, fixed: 'right' as const }
]

const tableData = ref([
  { id: 'user-001', username: 'manufacturer_user', org: 'Org1MSP', role: '制造商', status: '在线', registerTime: '2024-01-01 10:30' },
  { id: 'user-002', username: 'automaker_user', org: 'Org2MSP', role: '车企', status: '离线', registerTime: '2024-01-02 14:20' },
  { id: 'user-003', username: 'aftersale_user', org: 'Org3MSP', role: '售后中心', status: '在线', registerTime: '2024-01-03 09:15' },
  { id: 'user-004', username: 'parts_manager', org: 'Org1MSP', role: '制造商', status: '在线', registerTime: '2024-01-04 16:45' },
  { id: 'user-005', username: 'quality_checker', org: 'Org2MSP', role: '车企', status: '离线', registerTime: '2024-01-05 11:00' }
])

const totalCount = computed(() => tableData.value.length)
const manufacturerCount = computed(() => tableData.value.filter(item => item.role === '制造商').length)
const automakerCount = computed(() => tableData.value.filter(item => item.role === '车企').length)
const aftersaleCount = computed(() => tableData.value.filter(item => item.role === '售后中心').length)

const getRoleColor = (role: string) => {
  const colorMap: Record<string, string> = {
    '制造商': 'blue',
    '车企': 'green',
    '售后中心': 'orange'
  }
  return colorMap[role] || 'default'
}

const getAvatarColor = (role: string) => {
  const colorMap: Record<string, string> = {
    '制造商': '#3b82f6',
    '车企': '#10b981',
    '售后中心': '#f59e0b'
  }
  return colorMap[role] || '#6b7280'
}

const handleSearch = () => {
  message.success('搜索完成')
}

const handleReset = () => {
  searchForm.value = {
    userID: '',
    username: '',
    org: undefined,
    role: undefined
  }
}

const handleRefresh = () => {
  refreshing.value = true
  setTimeout(() => {
    refreshing.value = false
    message.success('数据已刷新')
  }, 1000)
}

const handleExport = () => {
  message.success('用户数据导出成功')
}

const viewDetail = (record: any) => {
  message.info(`查看用户详情: ${record.username}`)
}

const editUser = (record: any) => {
  message.info(`编辑用户: ${record.username}`)
}

const deleteUser = (record: any) => {
  const index = tableData.value.findIndex(item => item.id === record.id)
  if (index > -1) {
    tableData.value.splice(index, 1)
    message.success(`用户 ${record.username} 已删除`)
  }
}

const handleCreateUser = () => {
  if (!createForm.value.username || !createForm.value.org || !createForm.value.role) {
    message.warning('请填写必填项')
    return
  }
  const newUser = {
    id: `user-${String(tableData.value.length + 1).padStart(3, '0')}`,
    username: createForm.value.username,
    org: createForm.value.org,
    role: createForm.value.role,
    status: '离线',
    registerTime: new Date().toLocaleString()
  }
  tableData.value.unshift(newUser)
  showCreateModal.value = false
  message.success('用户添加成功')
}
</script>

<style scoped>
.user-management {
  padding: 0;
  min-height: 100%;
}

.page-header {
  margin-bottom: 24px;
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  border-radius: 12px;
  padding: 24px 32px;
  box-shadow: 0 4px 20px rgba(6, 182, 212, 0.25);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.header-title {
  display: flex;
  align-items: center;
  gap: 16px;
}

.title-icon {
  width: 48px;
  height: 48px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  color: #fff;
}

.title-text h2 {
  font-size: 24px;
  font-weight: 600;
  margin: 0;
  color: #fff;
}

.title-text .subtitle {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.85);
  margin: 4px 0 0;
}

.create-btn {
  height: 40px;
  border-radius: 10px;
  font-size: 14px;
  background: rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: #fff;
}

.create-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  border-color: rgba(255, 255, 255, 0.4);
  color: #fff;
}

.stats-overview {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
}

.stat-item {
  background: rgba(255, 255, 255, 0.15);
  border-radius: 10px;
  padding: 16px;
  text-align: center;
}

.stat-value {
  font-size: 28px;
  font-weight: 700;
  color: #fff;
  line-height: 1;
}

.stat-value.status-manufacturer {
  color: #60a5fa;
}

.stat-value.status-automaker {
  color: #34d399;
}

.stat-value.status-aftersale {
  color: #fbbf24;
}

.stat-label {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.85);
  margin-top: 6px;
}

.search-card {
  border-radius: 12px;
  margin-bottom: 24px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.search-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.filter-icon-wrapper {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: #fff;
}

.header-info {
  display: flex;
  flex-direction: column;
}

.header-title {
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
}

.header-desc {
  font-size: 13px;
  color: #6b7280;
  margin-top: 2px;
}

.filter-tag {
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 12px;
}

.search-form {
  width: 100%;
}

.filter-item {
  width: 100%;
}

.filter-label {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 8px;
  font-size: 13px;
  font-weight: 500;
  color: #374151;
}

.label-icon {
  color: #06b6d4;
}

.filter-input,
.filter-select {
  width: 100%;
}

.input-prefix-icon {
  color: #9ca3af;
}

.status-option {
  display: flex;
  align-items: center;
  gap: 8px;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.dot-manufacturer {
  background: #3b82f6;
}

.dot-automaker {
  background: #10b981;
}

.dot-aftersale {
  background: #f59e0b;
}

.search-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 20px;
  padding-top: 16px;
  border-top: 1px solid #f3f4f6;
}

.action-hint {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #9ca3af;
}

.action-right {
  display: flex;
  gap: 12px;
}

.reset-btn,
.search-btn {
  height: 38px;
  border-radius: 8px;
  font-size: 14px;
}

.reset-btn {
  border: 1px solid #e5e7eb;
  color: #6b7280;
}

.reset-btn:hover {
  border-color: #06b6d4;
  color: #06b6d4;
}

.search-btn {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  border: none;
}

.search-btn:hover {
  background: linear-gradient(135deg, #0891b2 0%, #0e7490 100%);
}

.table-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.table-header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.table-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 15px;
  font-weight: 600;
  color: #1f2937;
}

.title-icon {
  color: #06b6d4;
}

.record-count {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 12px;
}

.toolbar-btn {
  height: 36px;
  border-radius: 8px;
  font-size: 13px;
  border: 1px solid #e5e7eb;
  color: #6b7280;
}

.toolbar-btn:hover {
  border-color: #06b6d4;
  color: #06b6d4;
}

.custom-table :deep(.ant-table-thead > tr > th) {
  background: #f9fafb;
  font-weight: 600;
  color: #374151;
  border-bottom: 2px solid #e5e7eb;
}

.custom-table :deep(.ant-table-tbody > tr > td) {
  border-bottom: 1px solid #f3f4f6;
}

.custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #ecfeff;
}

.user-id-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.user-icon {
  color: #06b6d4;
}

.username-cell {
  display: flex;
  align-items: center;
  gap: 10px;
}

.user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  font-size: 14px;
}

.username {
  font-weight: 500;
  color: #1f2937;
}

.org-tag {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  border-radius: 6px;
}

.role-tag {
  border-radius: 6px;
}

.status-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.dot-online {
  background: #10b981;
}

.dot-offline {
  background: #9ca3af;
}

.status-text {
  font-size: 13px;
  color: #374151;
}

.time-cell {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #6b7280;
}

.time-icon {
  font-size: 12px;
}

.action-cell {
  gap: 4px;
}

.action-btn {
  padding: 0 8px;
  height: 28px;
  font-size: 12px;
}

.detail-btn {
  color: #6b7280;
}

.detail-btn:hover {
  color: #06b6d4;
}

.edit-btn {
  color: #3b82f6;
}

.edit-btn:hover {
  color: #2563eb;
}

.delete-btn {
  color: #ef4444;
}

.create-form :deep(.ant-input),
.create-form :deep(.ant-select-selector) {
  border-radius: 8px;
}

[data-theme='dark'] .page-header {
  background: linear-gradient(135deg, #0e7490 0%, #155e75 100%);
  box-shadow: 0 4px 20px rgba(14, 116, 144, 0.3);
}

[data-theme='dark'] .title-text h2 {
  color: #fff;
}

[data-theme='dark'] .title-text .subtitle {
  color: rgba(255, 255, 255, 0.8);
}

[data-theme='dark'] .create-btn {
  background: rgba(255, 255, 255, 0.15);
  border-color: rgba(255, 255, 255, 0.25);
  color: #fff;
}

[data-theme='dark'] .stat-item {
  background: rgba(255, 255, 255, 0.1);
}

[data-theme='dark'] .stat-value {
  color: #fff;
}

[data-theme='dark'] .stat-label {
  color: rgba(255, 255, 255, 0.8);
}

[data-theme='dark'] .search-card,
[data-theme='dark'] .table-card {
  background: #1f2937;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

[data-theme='dark'] .header-title,
[data-theme='dark'] .table-title {
  color: #f3f4f6;
}

[data-theme='dark'] .header-desc,
[data-theme='dark'] .action-hint {
  color: #9ca3af;
}

[data-theme='dark'] .filter-label {
  color: #e5e7eb;
}

[data-theme='dark'] .reset-btn,
[data-theme='dark'] .toolbar-btn {
  border-color: #374151;
  color: #9ca3af;
}

[data-theme='dark'] .reset-btn:hover,
[data-theme='dark'] .toolbar-btn:hover {
  border-color: #22d3ee;
  color: #22d3ee;
}

[data-theme='dark'] .search-btn {
  background: linear-gradient(135deg, #0891b2 0%, #0e7490 100%);
}

[data-theme='dark'] .custom-table :deep(.ant-table-thead > tr > th) {
  background: #111827;
  color: #e5e7eb;
  border-bottom-color: #374151;
}

[data-theme='dark'] .custom-table :deep(.ant-table-tbody > tr > td) {
  background: #1f2937;
  border-bottom-color: #374151;
  color: #e5e7eb;
}

[data-theme='dark'] .custom-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #111827;
}

[data-theme='dark'] .username {
  color: #f3f4f6;
}

[data-theme='dark'] .status-text {
  color: #e5e7eb;
}

[data-theme='dark'] .time-cell {
  color: #9ca3af;
}

[data-theme='dark'] .action-btn.detail-btn {
  color: #9ca3af;
}

[data-theme='dark'] .action-btn.detail-btn:hover {
  color: #22d3ee;
}
</style>
