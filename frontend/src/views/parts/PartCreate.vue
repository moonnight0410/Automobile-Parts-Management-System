<!--
  零部件创建页面
  功能：创建新的零部件并提交到区块链
  设计：现代化卡片布局，与Dashboard风格统一
-->

<template>
  <div class="part-create-page">
    <!-- 顶部导航栏 -->
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">创建零部件</h1>
          <span class="page-subtitle">填写信息提交到区块链</span>
        </div>
      </div>
      <div class="top-bar-right">
        <a-tag :color="formStatus.color" class="status-badge">
          <component :is="formStatus.icon" />
          {{ formStatus.text }}
        </a-tag>
      </div>
    </div>

    <!-- 主内容区 -->
    <div class="main-content">
      <!-- 左侧表单区 -->
      <div class="form-container">
        <a-card :bordered="false" class="form-card">
          <!-- 表单头部 -->
          <div class="form-header">
            <div class="header-icon">
              <PlusCircleOutlined />
            </div>
            <div class="header-text">
              <h3>零部件信息</h3>
              <p>请填写以下信息创建新的零部件记录</p>
            </div>
          </div>

          <a-divider class="form-divider" />

          <!-- 表单内容 -->
          <a-form
            ref="formRef"
            :model="form"
            :rules="formRules"
            layout="vertical"
            @finish="handleSubmit"
            class="create-form"
          >
            <!-- 零部件ID -->
            <a-form-item label="零部件ID" name="partID" class="form-item">
              <div class="input-wrapper">
                <a-input
                  v-model:value="form.partID"
                  placeholder="自动生成或手动输入"
                  :maxlength="50"
                  @change="handlePartIDChange"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><NumberOutlined /></span>
                  </template>
                  <template #suffix>
                    <a-tooltip title="重新生成ID">
                      <span class="refresh-id-btn" @click="regeneratePartID">
                        <SyncOutlined :spin="generatingID" />
                      </span>
                    </a-tooltip>
                  </template>
                </a-input>
              </div>
            </a-form-item>

            <!-- 零部件名称 -->
            <a-form-item label="零部件名称" name="name" class="form-item">
              <a-input
                v-model:value="form.name"
                placeholder="请输入零部件名称"
                :maxlength="100"
                class="custom-input"
              >
                <template #prefix>
                  <span class="input-icon"><TagOutlined /></span>
                </template>
              </a-input>
            </a-form-item>

            <!-- 零部件类型和状态 -->
            <div class="form-row">
              <a-form-item label="零部件类型" name="type" class="form-item half">
                <a-select
                  v-model:value="form.type"
                  placeholder="请选择类型"
                  :options="partTypeOptions"
                  class="custom-select"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                  :dropdown-style="{ zIndex: 1050 }"
                />
              </a-form-item>
              <a-form-item label="初始状态" name="status" class="form-item half">
                <a-select
                  v-model:value="form.status"
                  placeholder="请选择状态"
                  :options="statusOptions"
                  class="custom-select"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                  :dropdown-style="{ zIndex: 1050 }"
                />
              </a-form-item>
            </div>

            <!-- 批次号和生产厂商 -->
            <div class="form-row">
              <a-form-item label="批次号" name="batchNo" class="form-item half">
                <a-input
                  v-model:value="form.batchNo"
                  placeholder="请输入批次号"
                  :maxlength="50"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><BarcodeOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
              <a-form-item label="生产厂商" name="manufacturer" class="form-item half">
                <a-input
                  v-model:value="form.manufacturer"
                  placeholder="自动获取"
                  disabled
                  class="custom-input disabled"
                >
                  <template #prefix>
                    <span class="input-icon"><ShopOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
            </div>

            <!-- VIN码 -->
            <a-form-item label="VIN码" name="vin" class="form-item">
              <a-input
                v-model:value="form.vin"
                placeholder="请输入关联车辆VIN码（可选）"
                :maxlength="17"
                class="custom-input"
              >
                <template #prefix>
                  <span class="input-icon"><CarOutlined /></span>
                </template>
              </a-input>
              <template #extra>
                <span class="field-hint">17位车辆识别码，不含字母I、O、Q</span>
              </template>
            </a-form-item>

            <!-- 操作按钮 -->
            <div class="form-actions">
              <a-button
                type="primary"
                html-type="submit"
                :loading="submitting"
                :disabled="!isFormValid"
                class="submit-btn"
              >
                <template #icon><CheckOutlined /></template>
                {{ submitting ? '提交中...' : '提交创建' }}
              </a-button>
              <a-button @click="handleReset" class="secondary-btn">
                <template #icon><ReloadOutlined /></template>
                重置
              </a-button>
              <a-button @click="handleSaveDraft" class="secondary-btn">
                <template #icon><SaveOutlined /></template>
                保存草稿
              </a-button>
            </div>
          </a-form>
        </a-card>
      </div>

      <!-- 右侧信息区 -->
      <div class="info-container">
        <!-- 创建说明卡片 -->
        <a-card :bordered="false" class="info-card guide-card">
          <div class="card-title">
            <InfoCircleOutlined />
            <span>创建说明</span>
          </div>
          <ul class="guide-list">
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>零部件ID自动生成，也可手动修改</span>
            </li>
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>带 <em>*</em> 的为必填项</span>
            </li>
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>提交后数据将记录到区块链</span>
            </li>
            <li>
              <CheckCircleOutlined class="guide-icon warning" />
              <span>数据提交后不可更改，请仔细核对</span>
            </li>
          </ul>
        </a-card>

        <!-- 表单预览卡片 -->
        <a-card :bordered="false" class="info-card preview-card">
          <div class="card-title">
            <EyeOutlined />
            <span>数据预览</span>
          </div>
          <div class="preview-content">
            <div class="preview-item">
              <span class="preview-label">零部件ID</span>
              <span class="preview-value">{{ form.partID || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">名称</span>
              <span class="preview-value">{{ form.name || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">类型</span>
              <span class="preview-value">{{ form.type || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">批次号</span>
              <span class="preview-value">{{ form.batchNo || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">状态</span>
              <a-tag :color="getStatusColor(form.status)" size="small">
                {{ getStatusText(form.status) }}
              </a-tag>
            </div>
          </div>
        </a-card>

        <!-- 快捷操作卡片 -->
        <a-card :bordered="false" class="info-card quick-card">
          <div class="card-title">
            <ThunderboltOutlined />
            <span>快捷操作</span>
          </div>
          <div class="quick-actions-grid">
            <div class="quick-action-item" @click="fillTestData">
              <ThunderboltOutlined class="action-icon" />
              <span>填充测试数据</span>
            </div>
            <div class="quick-action-item" @click="handleSaveDraft">
              <SaveOutlined class="action-icon" />
              <span>保存草稿</span>
            </div>
          </div>
        </a-card>
      </div>
    </div>

    <!-- 成功弹窗 -->
    <a-modal
      v-model:visible="showSuccessModal"
      :footer="null"
      :closable="false"
      centered
      class="success-modal"
      width="420px"
    >
      <div class="success-content">
        <div class="success-icon-wrapper">
          <CheckCircleFilled class="success-icon" />
        </div>
        <h2 class="success-title">创建成功</h2>
        <p class="success-desc">零部件已成功提交到区块链</p>
        <div class="success-part-id">
          <span class="label">零部件ID：</span>
          <span class="value">{{ createdPartID }}</span>
        </div>
        <div class="success-actions">
          <a-button type="primary" @click="handleCreateAnother">
            <template #icon><PlusOutlined /></template>
            继续创建
          </a-button>
          <a-button @click="goToDetail">
            <template #icon><EyeOutlined /></template>
            查看详情
          </a-button>
          <a-button @click="goBack">
            <template #icon><UnorderedListOutlined /></template>
            返回列表
          </a-button>
        </div>
      </div>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { message, Modal } from 'ant-design-vue'
import { useAuthStore, usePartStore } from '../../stores'
import type { FormInstance, SelectProps } from 'ant-design-vue'
import type { CreatePartRequest } from '../../types'
import {
  ArrowLeftOutlined,
  CheckOutlined,
  ReloadOutlined,
  SaveOutlined,
  PlusOutlined,
  EyeOutlined,
  UnorderedListOutlined,
  InfoCircleOutlined,
  ClockCircleOutlined,
  CheckCircleFilled,
  CheckCircleOutlined,
  ThunderboltOutlined,
  NumberOutlined,
  TagOutlined,
  BarcodeOutlined,
  ShopOutlined,
  CarOutlined,
  SyncOutlined,
  PlusCircleOutlined,
  ExclamationCircleOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const authStore = useAuthStore()
const partStore = usePartStore()

const formRef = ref<FormInstance>()
const submitting = ref(false)
const showSuccessModal = ref(false)
const createdPartID = ref('')
const generatingID = ref(false)

const form = reactive({
  partID: '',
  name: '',
  type: '',
  batchNo: '',
  vin: '',
  manufacturer: '',
  status: 'IN_PRODUCTION'
})

const originalForm = reactive({ ...form })

const formStatus = computed(() => {
  if (submitting.value) {
    return { text: '提交中', color: 'processing', icon: SyncOutlined }
  }
  if (isFormValid.value) {
    return { text: '可提交', color: 'success', icon: CheckCircleOutlined }
  }
  return { text: '待填写', color: 'warning', icon: ClockCircleOutlined }
})

const partTypeOptions = computed<SelectProps['options']>(() => [
  { value: '发动机部件', label: '发动机部件' },
  { value: '传动系统', label: '传动系统' },
  { value: '制动系统', label: '制动系统' },
  { value: '悬挂系统', label: '悬挂系统' },
  { value: '电气系统', label: '电气系统' },
  { value: '车身部件', label: '车身部件' },
  { value: '内饰件', label: '内饰件' },
  { value: '外饰件', label: '外饰件' }
])

const statusOptions = computed<SelectProps['options']>(() => [
  { value: 'IN_PRODUCTION', label: '在产' },
  { value: 'NORMAL', label: '正常' }
])

const isFormValid = computed(() => {
  return (
    form.partID.trim() !== '' &&
    form.name.trim() !== '' &&
    form.type !== '' &&
    form.batchNo.trim() !== '' &&
    form.status !== ''
  )
})

const hasFormChanged = computed(() => {
  return JSON.stringify(form) !== JSON.stringify(originalForm)
})

const validatePartID = async (_rule: any, value: string) => {
  if (!value) return Promise.reject('请输入零部件ID')
  if (!/^PART-[A-Z0-9-]+$/.test(value)) {
    return Promise.reject('格式：PART-XXX（大写字母、数字、横线）')
  }
  return Promise.resolve()
}

const validateVIN = async (_rule: any, value: string) => {
  if (value && !/^[A-HJ-NPR-Z0-9]{17}$/.test(value)) {
    return Promise.reject('VIN码为17位，不含I、O、Q')
  }
  return Promise.resolve()
}

const formRules = {
  partID: [{ required: true, validator: validatePartID, trigger: 'blur' }],
  name: [
    { required: true, message: '请输入零部件名称', trigger: 'blur' },
    { min: 2, max: 100, message: '名称长度2-100字符', trigger: 'blur' }
  ],
  type: [{ required: true, message: '请选择零部件类型', trigger: 'change' }],
  batchNo: [
    { required: true, message: '请输入批次号', trigger: 'blur' },
    { max: 50, message: '批次号不超过50字符', trigger: 'blur' }
  ],
  vin: [{ validator: validateVIN, trigger: 'blur' }],
  status: [{ required: true, message: '请选择初始状态', trigger: 'change' }]
}

function generatePartID(): string {
  const timestamp = Date.now().toString(36).toUpperCase()
  const random = Math.random().toString(36).substring(2, 6).toUpperCase()
  return `PART-${timestamp}-${random}`
}

function regeneratePartID() {
  generatingID.value = true
  setTimeout(() => {
    form.partID = generatePartID()
    generatingID.value = false
  }, 300)
}

function handlePartIDChange() {
  form.partID = form.partID.toUpperCase().replace(/[^A-Z0-9-]/g, '')
}

function filterOption(input: string, option: any) {
  return option.label.toLowerCase().includes(input.toLowerCase())
}

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    'IN_PRODUCTION': 'blue',
    'NORMAL': 'green',
    'IN_SALE': 'cyan',
    'RECALL': 'orange',
    'SCRAP': 'red'
  }
  return colors[status] || 'default'
}

function getStatusText(status: string) {
  const texts: Record<string, string> = {
    'IN_PRODUCTION': '在产',
    'NORMAL': '正常',
    'IN_SALE': '在售',
    'RECALL': '召回',
    'SCRAP': '报废'
  }
  return texts[status] || status
}

function goBack() {
  if (hasFormChanged.value && !showSuccessModal.value) {
    Modal.confirm({
      title: '确认离开',
      content: '表单数据已修改，离开将丢失未保存的内容',
      okText: '确定离开',
      cancelText: '取消',
      onOk: () => router.push('/parts/list')
    })
  } else {
    router.push('/parts/list')
  }
}

async function handleSubmit() {
  if (submitting.value) return
  submitting.value = true
  
  try {
    const requestData: CreatePartRequest = {
      partID: form.partID,
      name: form.name,
      type: form.type,
      batchNo: form.batchNo,
      vin: form.vin || undefined,
      manufacturer: form.manufacturer,
      status: form.status
    }
    
    await partStore.create(requestData)
    
    createdPartID.value = form.partID
    showSuccessModal.value = true
    clearDraft()
  } catch (error: any) {
    handleSubmissionError(error)
  } finally {
    submitting.value = false
  }
}

function handleSubmissionError(error: any) {
  const errorMessage = error.message || '创建失败'
  
  if (errorMessage.includes('已存在')) {
    message.error('该零部件ID已存在')
    formRef.value?.setFields({
      partID: { value: form.partID, errors: [new Error('零部件ID已存在')] }
    })
  } else if (errorMessage.includes('网络')) {
    Modal.error({
      title: '网络错误',
      content: '网络连接失败，请检查网络后重试',
      okText: '重试',
      onOk: () => handleSubmit()
    })
  } else {
    message.error(errorMessage)
  }
}

function handleReset() {
  Modal.confirm({
    title: '确认重置',
    content: '确定要重置表单吗？所有未保存的内容将丢失。',
    okText: '确定',
    cancelText: '取消',
    onOk: () => {
      formRef.value?.resetFields()
      form.partID = generatePartID()
      form.manufacturer = authStore.userOrg || 'Org1MSP'
      form.status = 'IN_PRODUCTION'
      Object.assign(originalForm, form)
      message.success('表单已重置')
    }
  })
}

function handleSaveDraft() {
  try {
    const draftData = { ...form, savedAt: new Date().toISOString() }
    localStorage.setItem('part_create_draft', JSON.stringify(draftData))
    message.success('草稿已保存')
  } catch {
    message.error('保存草稿失败')
  }
}

function loadDraft() {
  try {
    const draftStr = localStorage.getItem('part_create_draft')
    if (draftStr) {
      const draft = JSON.parse(draftStr)
      Modal.confirm({
        title: '发现未完成的草稿',
        content: `上次保存：${new Date(draft.savedAt).toLocaleString()}`,
        okText: '恢复',
        cancelText: '不恢复',
        onOk: () => {
          Object.assign(form, draft)
          Object.assign(originalForm, draft)
          message.success('草稿已恢复')
        },
        onCancel: clearDraft
      })
    }
  } catch {
    clearDraft()
  }
}

function clearDraft() {
  localStorage.removeItem('part_create_draft')
}

function fillTestData() {
  form.name = '测试零部件-' + Date.now().toString().slice(-4)
  form.type = '发动机部件'
  form.batchNo = 'BATCH-' + Date.now().toString().slice(-6)
  form.vin = 'LSVNV2182E21' + Date.now().toString().slice(-5)
  message.success('已填充测试数据')
}

function handleCreateAnother() {
  showSuccessModal.value = false
  createdPartID.value = ''
  formRef.value?.resetFields()
  form.partID = generatePartID()
  form.manufacturer = authStore.userOrg || 'Org1MSP'
  form.status = 'IN_PRODUCTION'
  Object.assign(originalForm, form)
}

function goToDetail() {
  router.push(`/parts/detail/${createdPartID.value}`)
}

onMounted(() => {
  form.partID = generatePartID()
  form.manufacturer = authStore.userOrg || 'Org1MSP'
  Object.assign(originalForm, form)
  loadDraft()
})

onUnmounted(() => {
  if (!showSuccessModal.value && hasFormChanged.value) {
    handleSaveDraft()
  }
})

window.addEventListener('beforeunload', (e) => {
  if (hasFormChanged.value && !showSuccessModal.value) {
    e.preventDefault()
    e.returnValue = ''
  }
})
</script>

<style scoped>
.part-create-page {
  min-height: 100vh;
  background: var(--bg-color);
  padding: 0;
  display: flex;
  flex-direction: column;
}

.part-create-page :deep(.ant-select-dropdown) {
  z-index: 1050 !important;
}

.part-create-page :deep(.ant-picker-dropdown) {
  z-index: 1050 !important;
}

.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 32px;
  background: var(--bg-color-secondary);
  border-bottom: 1px solid var(--border-color);
  position: sticky;
  top: 0;
  z-index: 100;
  flex-shrink: 0;
}

.top-bar-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.back-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  color: var(--text-color-secondary);
  font-weight: 500;
  padding: 8px 12px;
  border-radius: 8px;
  transition: all 0.3s;
}

.back-btn:hover {
  background: var(--bg-color-tertiary);
  color: var(--primary-color);
}

.divider {
  height: 24px;
  margin: 0;
  background: var(--border-color);
}

.page-info {
  display: flex;
  flex-direction: column;
}

.page-title {
  font-size: 18px;
  font-weight: 600;
  margin: 0;
  color: var(--text-color);
}

.page-subtitle {
  font-size: 13px;
  color: var(--text-color-secondary);
  margin: 0;
}

.status-badge {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 14px;
  border-radius: 20px;
  font-size: 13px;
  font-weight: 500;
}

.main-content {
  display: flex;
  gap: 24px;
  padding: 24px 32px;
  max-width: 1400px;
  margin: 0 auto;
  flex: 1;
  width: 100%;
  box-sizing: border-box;
  overflow: visible;
  position: relative;
}

.form-container {
  flex: 1;
  min-width: 0;
  overflow: visible;
  position: relative;
}

.form-card {
  border-radius: 16px;
  box-shadow: var(--shadow-sm);
  background: var(--bg-color-secondary);
  min-height: 100%;
  overflow: visible;
  position: relative;
  z-index: 10;
}

.form-card :deep(.ant-card-body) {
  padding: 28px;
}

.form-header {
  display: flex;
  align-items: center;
  gap: 16px;
  padding-bottom: 20px;
}

.header-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, var(--primary-color) 0%, #722ed1 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  color: #fff;
}

.header-text h3 {
  font-size: 18px;
  font-weight: 600;
  margin: 0 0 4px 0;
  color: var(--text-color);
}

.header-text p {
  font-size: 13px;
  color: var(--text-color-secondary);
  margin: 0;
}

.form-divider {
  margin: 0 0 24px 0;
  border-color: var(--border-color);
}

.create-form {
  padding: 0;
}

.form-item {
  margin-bottom: 20px;
}

.form-item :deep(.ant-form-item-label > label) {
  font-size: 14px;
  font-weight: 500;
  color: var(--text-color);
}

.form-row {
  display: flex;
  gap: 20px;
}

.form-item.half {
  flex: 1;
}

.input-wrapper {
  position: relative;
}

.custom-input,
.custom-select {
  width: 100%;
  height: 42px;
}

.custom-input :deep(.ant-input),
.custom-select :deep(.ant-select-selector) {
  height: 42px;
  border-radius: 10px;
  border: 2px solid var(--border-color);
  background: var(--bg-color-secondary);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  font-size: 14px;
}

.custom-input :deep(.ant-input) {
  padding-left: 40px;
}

.custom-input :deep(.ant-input-affix-wrapper) {
  display: flex;
  align-items: center;
  height: 42px;
  padding: 0 12px;
  width: 100%;
  gap: 8px;
  box-sizing: border-box;
  border-radius: 10px;
  border: 2px solid var(--border-color);
  background: var(--bg-color-secondary);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.custom-input :deep(.ant-input-prefix) {
  margin-right: 8px;
  color: var(--text-color-tertiary);
}

.input-prefix-icon {
  font-size: 14px;
}

.custom-input :deep(.ant-input-affix-wrapper:hover),
.custom-select :deep(.ant-select-selector:hover) {
  border-color: var(--primary-color);
  box-shadow: 0 2px 8px rgba(24, 144, 255, 0.15);
}

.custom-input :deep(.ant-input-affix-wrapper:focus),
.custom-input :deep(.ant-input-affix-wrapper-focused),
.custom-select :deep(.ant-select-focused .ant-select-selector) {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(24, 144, 255, 0.12);
}

.custom-select :deep(.ant-select-selector) {
  padding: 0 12px !important;
  border: 2px solid var(--border-color) !important;
  display: flex !important;
  align-items: center !important;
}

.custom-select :deep(.ant-select-selection-search) {
  display: flex !important;
  align-items: center !important;
}

.custom-select :deep(.ant-select-selection-search-input) {
  height: 38px !important;
}

.custom-select :deep(.ant-select-selection-item),
.custom-select :deep(.ant-select-selection-placeholder) {
  line-height: 38px !important;
  display: flex !important;
  align-items: center !important;
}

.custom-select :deep(.ant-select-dropdown) {
  z-index: 1050 !important;
  border-radius: 8px;
  box-shadow: 0 6px 16px 0 rgba(0, 0, 0, 0.12);
}

.custom-select :deep(.ant-select-item) {
  padding: 10px 12px;
  transition: all 0.2s ease;
  border-radius: 8px;
  margin: 2px 0;
  display: flex;
  align-items: center;
  justify-content: flex-start;
}

.custom-select :deep(.ant-select-item:hover) {
  background: #f5f3ff;
  color: #6366f1;
}

.custom-select :deep(.ant-select-item-option-selected) {
  background: #ede9fe;
  color: #6366f1;
  font-weight: 500;
}

.custom-select :deep(.ant-select-item-option-active) {
  background: #f5f3ff;
}

.custom-select :deep(.ant-select-item-option-disabled) {
  color: rgba(0, 0, 0, 0.25);
  cursor: not-allowed;
  background: transparent;
}

.custom-input.disabled :deep(.ant-input-affix-wrapper) {
  background: var(--bg-color-tertiary);
  cursor: not-allowed;
}

.input-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: var(--text-color-tertiary);
  font-size: 14px;
}

.custom-input :deep(.ant-input-affix-wrapper:focus-within) .input-icon {
  color: var(--primary-color);
}

.refresh-id-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: var(--text-color-tertiary);
  cursor: pointer;
  padding: 4px;
  border-radius: 4px;
  transition: all 0.3s;
}

.refresh-id-btn:hover {
  color: var(--primary-color);
  background: rgba(24, 144, 255, 0.1);
}

.field-hint {
  font-size: 12px;
  color: var(--text-color-tertiary);
  margin-top: 4px;
}

.form-actions {
  display: flex;
  gap: 12px;
  padding-top: 24px;
  margin-top: 8px;
  border-top: 1px solid var(--border-color);
}

.submit-btn {
  min-width: 140px;
  height: 44px;
  border-radius: 10px;
  font-weight: 500;
  font-size: 14px;
  background: linear-gradient(135deg, var(--primary-color) 0%, #722ed1 100%);
  border: none;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3);
}

.submit-btn:hover:not(:disabled) {
  background: linear-gradient(135deg, #40a9ff 0%, #9254de 100%);
  box-shadow: 0 6px 16px rgba(24, 144, 255, 0.4);
  transform: translateY(-1px);
}

.submit-btn:disabled {
  opacity: 0.5;
  box-shadow: none;
}

.secondary-btn {
  height: 44px;
  border-radius: 10px;
  font-weight: 500;
  font-size: 14px;
  border: 2px solid var(--border-color);
  background: transparent;
  color: var(--text-color);
}

.secondary-btn:hover {
  border-color: var(--primary-color);
  color: var(--primary-color);
}

.info-container {
  width: 340px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.info-card {
  border-radius: 16px;
  box-shadow: var(--shadow-sm);
  background: var(--bg-color-secondary);
  padding: 20px;
}

.card-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  font-weight: 600;
  color: var(--text-color);
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--border-color);
}

.card-title .anticon {
  font-size: 16px;
  color: var(--primary-color);
}

.guide-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.guide-list li {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  padding: 8px 0;
  font-size: 13px;
  color: var(--text-color-secondary);
}

.guide-list li em {
  color: var(--error-color);
  font-style: normal;
}

.guide-icon {
  font-size: 14px;
  margin-top: 2px;
}

.guide-icon.success {
  color: var(--success-color);
}

.guide-icon.warning {
  color: var(--warning-color);
}

.preview-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.preview-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 12px;
  background: var(--bg-color);
  border-radius: 8px;
}

.preview-label {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.preview-value {
  font-size: 13px;
  font-weight: 500;
  color: var(--text-color);
  max-width: 150px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.quick-actions-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.quick-action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  padding: 20px 16px;
  background: var(--bg-color);
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s;
  border: 1px solid transparent;
}

.quick-action-item:hover {
  background: var(--bg-color-tertiary);
  border-color: var(--primary-color);
}

.quick-action-item .action-icon {
  font-size: 20px;
  color: var(--primary-color);
}

.quick-action-item span {
  font-size: 12px;
  color: var(--text-color-secondary);
}

:global(.ant-modal-root:has(.success-modal)) {
  position: fixed !important;
  inset: 0 !important;
  z-index: 1000 !important;
}

:global(.ant-modal-root:has(.success-modal) .ant-modal-mask) {
  background-color: rgba(0, 0, 0, 0.85) !important;
}

:global(.ant-modal-wrap:has(.success-modal)) {
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  height: 100% !important;
}

:global(.ant-modal-wrap:has(.success-modal) .ant-modal) {
  top: 0 !important;
  padding-bottom: 0 !important;
  display: inline-block;
}

.success-modal :deep(.ant-modal-content) {
  border-radius: 16px;
  padding: 0;
  overflow: hidden;
}

.success-content {
  padding: 40px 32px;
  text-align: center;
}

.success-icon-wrapper {
  width: 72px;
  height: 72px;
  margin: 0 auto 24px;
  background: linear-gradient(135deg, #52c41a 0%, #73d13d 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.success-icon {
  font-size: 36px;
  color: #fff;
}

.success-title {
  font-size: 22px;
  font-weight: 600;
  margin: 0 0 8px 0;
  color: var(--text-color);
}

.success-desc {
  font-size: 14px;
  color: var(--text-color-secondary);
  margin: 0 0 20px 0;
}

.success-part-id {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  background: var(--bg-color-tertiary);
  border-radius: 8px;
  margin-bottom: 24px;
}

.success-part-id .label {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.success-part-id .value {
  font-size: 14px;
  font-weight: 600;
  color: var(--primary-color);
  font-family: monospace;
}

.success-actions {
  display: flex;
  justify-content: center;
  gap: 12px;
}

.success-actions .ant-btn {
  height: 40px;
  border-radius: 8px;
  font-weight: 500;
}

@media (max-width: 1024px) {
  .main-content {
    flex-direction: column;
    padding: 16px;
  }
  
  .info-container {
    width: 100%;
    flex-direction: row;
    flex-wrap: wrap;
  }
  
  .info-card {
    flex: 1;
    min-width: 280px;
  }
}

@media (max-width: 768px) {
  .top-bar {
    padding: 12px 16px;
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
  
  .form-row {
    flex-direction: column;
    gap: 0;
  }
  
  .form-actions {
    flex-direction: column;
  }
  
  .form-actions .ant-btn {
    width: 100%;
  }
}

[data-theme='dark'] .top-bar {
  background: var(--bg-color-tertiary);
}

[data-theme='dark'] .form-card,
[data-theme='dark'] .info-card {
  background: var(--bg-color-tertiary);
}

[data-theme='dark'] .header-icon {
  background: linear-gradient(135deg, #177ddc 0%, #642ab5 100%);
}

[data-theme='dark'] .custom-input,
[data-theme='dark'] .custom-select {
  background: var(--bg-color);
  border-color: var(--border-color);
}

[data-theme='dark'] .custom-input:hover,
[data-theme='dark'] .custom-select:hover {
  border-color: #177ddc;
}

[data-theme='dark'] .custom-input:focus,
[data-theme='dark'] .custom-select:focus-within {
  border-color: #177ddc;
  box-shadow: 0 0 0 3px rgba(23, 125, 220, 0.15);
}

[data-theme='dark'] .custom-input:focus-within .input-icon {
  color: #177ddc;
}

[data-theme='dark'] .submit-btn {
  background: linear-gradient(135deg, #177ddc 0%, #642ab5 100%);
}

[data-theme='dark'] .submit-btn:hover:not(:disabled) {
  background: linear-gradient(135deg, #3c9ae8 0%, #7a3db8 100%);
}

[data-theme='dark'] .preview-item,
[data-theme='dark'] .quick-action-item {
  background: var(--bg-color);
}

[data-theme='dark'] .quick-action-item:hover {
  background: var(--bg-color-secondary);
}

[data-theme='dark'] .success-part-id {
  background: var(--bg-color);
}

[data-theme='dark'] .custom-select :deep(.ant-select-dropdown) {
  background: var(--bg-color-tertiary);
  box-shadow: 0 6px 16px 0 rgba(0, 0, 0, 0.45);
}

[data-theme='dark'] .custom-select :deep(.ant-select-item) {
  color: #e5e7eb;
}

[data-theme='dark'] .custom-select :deep(.ant-select-item:hover) {
  background: rgba(129, 140, 248, 0.15);
  color: #a5b4fc;
}

[data-theme='dark'] .custom-select :deep(.ant-select-item-option-selected) {
  background: rgba(129, 140, 248, 0.2);
  color: #a5b4fc;
}

[data-theme='dark'] .custom-select :deep(.ant-select-item-option-active) {
  background: rgba(129, 140, 248, 0.1);
}

[data-theme='dark'] .custom-select :deep(.ant-select-item-option-disabled) {
  color: rgba(255, 255, 255, 0.3);
  background: transparent;
}
</style>
