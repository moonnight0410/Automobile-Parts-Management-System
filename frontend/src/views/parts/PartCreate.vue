<!--
  零部件创建页面
  创建新的零部件并提交到区块链
-->

<template>
  <div class="part-create">
    <!-- 页面标题 -->
    <div class="page-header">
      <h2>创建零部件</h2>
      <a-button @click="goBack">
        <template #icon><ArrowLeftOutlined /></template>
        返回
      </a-button>
    </div>
    
    <!-- 创建表单 -->
    <a-card :bordered="false" class="form-card">
      <a-form
        ref="formRef"
        :model="form"
        :rules="rules"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 14 }"
        @finish="handleSubmit"
      >
        <a-form-item label="零部件ID" name="partID">
          <a-input
            v-model:value="form.partID"
            placeholder="请输入零部件唯一标识，如 PART-001"
          />
        </a-form-item>
        
        <a-form-item label="零部件名称" name="name">
          <a-input
            v-model:value="form.name"
            placeholder="请输入零部件名称"
          />
        </a-form-item>
        
        <a-form-item label="零部件类型" name="type">
          <a-select
            v-model:value="form.type"
            placeholder="请选择零部件类型"
          >
            <a-select-option value="发动机部件">发动机部件</a-select-option>
            <a-select-option value="传动系统">传动系统</a-select-option>
            <a-select-option value="制动系统">制动系统</a-select-option>
            <a-select-option value="悬挂系统">悬挂系统</a-select-option>
            <a-select-option value="电气系统">电气系统</a-select-option>
            <a-select-option value="车身部件">车身部件</a-select-option>
          </a-select>
        </a-form-item>
        
        <a-form-item label="批次号" name="batchNo">
          <a-input
            v-model:value="form.batchNo"
            placeholder="请输入生产批次号"
          />
        </a-form-item>
        
        <a-form-item label="VIN码" name="vin">
          <a-input
            v-model:value="form.vin"
            placeholder="请输入关联车辆VIN码（可选）"
          />
        </a-form-item>
        
        <a-form-item label="生产厂商" name="manufacturer">
          <a-input
            v-model:value="form.manufacturer"
            placeholder="生产厂商"
            disabled
          />
        </a-form-item>
        
        <a-form-item label="初始状态" name="status">
          <a-select
            v-model:value="form.status"
            placeholder="请选择初始状态"
          >
            <a-select-option value="IN_PRODUCTION">在产</a-select-option>
            <a-select-option value="NORMAL">正常</a-select-option>
          </a-select>
        </a-form-item>
        
        <a-form-item :wrapper-col="{ span: 14, offset: 6 }">
          <a-space>
            <a-button type="primary" html-type="submit" :loading="loading">
              <template #icon><CheckOutlined /></template>
              提交
            </a-button>
            <a-button @click="handleReset">
              <template #icon><ReloadOutlined /></template>
              重置
            </a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { useAuthStore, usePartStore } from '../../stores'
import type { FormInstance } from 'ant-design-vue'

// ==================== 组合式API ====================

const router = useRouter()
const authStore = useAuthStore()
const partStore = usePartStore()

// ==================== 响应式状态 ====================

const formRef = ref<FormInstance>()
const loading = ref(false)

// 表单数据
const form = reactive({
  partID: '',
  name: '',
  type: '',
  batchNo: '',
  vin: '',
  manufacturer: computed(() => authStore.userOrg || 'Org1MSP'),
  status: 'IN_PRODUCTION'
})

// 表单验证规则
const rules = {
  partID: [
    { required: true, message: '请输入零部件ID', trigger: 'blur' },
    { pattern: /^PART-[A-Z0-9]+$/, message: '零部件ID格式为 PART-XXX', trigger: 'blur' }
  ],
  name: [
    { required: true, message: '请输入零部件名称', trigger: 'blur' }
  ],
  type: [
    { required: true, message: '请选择零部件类型', trigger: 'change' }
  ],
  batchNo: [
    { required: true, message: '请输入批次号', trigger: 'blur' }
  ],
  status: [
    { required: true, message: '请选择初始状态', trigger: 'change' }
  ]
}

// ==================== 方法 ====================

/**
 * 返回上一页
 */
function goBack() {
  router.back()
}

/**
 * 提交表单
 */
async function handleSubmit() {
  loading.value = true
  
  try {
    await partStore.create({
      partID: form.partID,
      name: form.name,
      type: form.type,
      batchNo: form.batchNo,
      vin: form.vin || undefined,
      manufacturer: form.manufacturer,
      status: form.status
    })
    
    message.success('零部件创建成功')
    router.push('/parts/list')
  } catch (error: any) {
    message.error(error.message || '创建失败')
  } finally {
    loading.value = false
  }
}

/**
 * 重置表单
 */
function handleReset() {
  formRef.value?.resetFields()
}

// ==================== 生命周期 ====================

onMounted(() => {
  // 自动生成零部件ID
  form.partID = `PART-${Date.now().toString(36).toUpperCase()}`
})
</script>

<style scoped>
.part-create {
  padding: 0;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.page-header h2 {
  font-size: 20px;
  font-weight: 600;
  margin: 0;
}

.form-card {
  border-radius: 8px;
  max-width: 800px;
}
</style>
