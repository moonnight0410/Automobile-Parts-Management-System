<!--
  故障上报页面
  设计：与生产数据创建页面风格统一
-->

<template>
  <div class="fault-create">
    <!-- 顶部导航栏 -->
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">上报故障</h1>
          <span class="page-subtitle">填写故障信息与风险预警</span>
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
              <AlertOutlined />
            </div>
            <div class="header-text">
              <h3>故障信息</h3>
              <p>请填写以下信息上报故障</p>
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
            <!-- 零部件ID和VIN码 -->
            <div class="form-row">
              <a-form-item label="零部件ID" name="partID" class="form-item half">
                <a-input
                  v-model:value="form.partID"
                  placeholder="请输入零部件ID"
                  :maxlength="50"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><AppstoreOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
              <a-form-item label="VIN码" name="vin" class="form-item half">
                <a-input
                  v-model:value="form.vin"
                  placeholder="请输入VIN码"
                  :maxlength="17"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><CarOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
            </div>

            <!-- 故障类型和风险概率 -->
            <div class="form-row">
              <a-form-item label="故障类型" name="faultType" class="form-item half">
                <a-select
                  v-model:value="form.faultType"
                  placeholder="请选择故障类型"
                  :options="faultTypeOptions"
                  class="custom-select"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                  :dropdown-style="{ zIndex: 1050 }"
                />
              </a-form-item>
              <a-form-item label="风险概率" name="riskProb" class="form-item half">
                <a-input-number
                  v-model:value="form.riskProb"
                  placeholder="请输入风险概率"
                  :min="0"
                  :max="100"
                  :addon-after="'%'"
                  class="custom-input-number"
                />
              </a-form-item>
            </div>

            <!-- 故障描述 -->
            <a-form-item label="故障描述" name="description" class="form-item">
              <a-textarea
                v-model:value="form.description"
                placeholder="请详细描述故障情况"
                :rows="4"
                :maxlength="500"
                class="custom-textarea"
              />
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
                {{ submitting ? '提交中...' : '提交上报' }}
              </a-button>
              <a-button @click="handleReset" class="secondary-btn">
                <template #icon><ReloadOutlined /></template>
                重置
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
            <span>上报说明</span>
          </div>
          <ul class="guide-list">
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>请确保零部件ID和VIN码准确</span>
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
              <span>风险概率范围为0-100%</span>
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
              <span class="preview-label">VIN码</span>
              <span class="preview-value">{{ form.vin || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">故障类型</span>
              <a-tag :color="getFaultTypeColor(form.faultType)" size="small">
                {{ form.faultType || '-' }}
              </a-tag>
            </div>
            <div class="preview-item">
              <span class="preview-label">风险概率</span>
              <a-tag :color="getRiskColor(form.riskProb)" size="small">
                {{ form.riskProb || 0 }}%
              </a-tag>
            </div>
            <div class="preview-item">
              <span class="preview-label">描述</span>
              <span class="preview-value">{{ form.description || '-' }}</span>
            </div>
          </div>
        </a-card>

        <!-- 快捷操作卡片 -->
        <a-card :bordered="false" class="info-card quick-card">
          <div class="card-title">
            <ThunderboltOutlined />
            <span>快捷操作</span>
          </div>
          <div class="quick-actions">
            <div class="quick-action-item" @click="fillTestData">
              <ThunderboltOutlined class="action-icon" />
              <span>填充测试数据</span>
            </div>
          </div>
        </a-card>
      </div>
    </div>

    <!-- 成功弹窗 -->
    <a-modal
      v-model:open="showSuccessModal"
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
        <h2 class="success-title">上报成功</h2>
        <p class="success-desc">故障信息已成功提交到区块链</p>
        <div class="success-fault-id">
          <span class="label">故障ID：</span>
          <span class="value">{{ createdFaultID }}</span>
        </div>
        <div class="success-actions">
          <a-button type="primary" @click="handleCreateAnother">
            <template #icon><PlusOutlined /></template>
            继续上报
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
import { ref, reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import type { FormInstance, SelectProps } from 'ant-design-vue'
import {
  ArrowLeftOutlined,
  CheckOutlined,
  ReloadOutlined,
  InfoCircleOutlined,
  CheckCircleOutlined,
  EyeOutlined,
  AlertOutlined,
  AppstoreOutlined,
  CarOutlined,
  ThunderboltOutlined,
  PlusOutlined,
  UnorderedListOutlined,
  CheckCircleFilled
} from '@ant-design/icons-vue'

const router = useRouter()
const formRef = ref<FormInstance>()
const submitting = ref(false)
const showSuccessModal = ref(false)
const createdFaultID = ref('')

const form = reactive({
  partID: '',
  vin: '',
  faultType: '',
  riskProb: 50,
  description: ''
})

const formStatus = computed(() => {
  if (submitting.value) {
    return { text: '提交中', color: 'processing', icon: ReloadOutlined }
  }
  if (isFormValid.value) {
    return { text: '可提交', color: 'success', icon: CheckCircleOutlined }
  }
  return { text: '待填写', color: 'warning', icon: AlertOutlined }
})

const faultTypeOptions = computed<SelectProps['options']>(() => [
  { value: '制动故障', label: '制动故障' },
  { value: '转向故障', label: '转向故障' },
  { value: '发动机故障', label: '发动机故障' },
  { value: '电气故障', label: '电气故障' },
  { value: '悬挂故障', label: '悬挂故障' },
  { value: '其他', label: '其他' }
])

const isFormValid = computed(() => {
  return (
    form.partID.trim() !== '' &&
    form.vin.trim() !== '' &&
    form.faultType !== '' &&
    form.description.trim() !== ''
  )
})

const formRules = {
  partID: [{ required: true, message: '请输入零部件ID', trigger: 'blur' }],
  vin: [{ required: true, message: '请输入VIN码', trigger: 'blur' }],
  faultType: [{ required: true, message: '请选择故障类型', trigger: 'change' }],
  description: [{ required: true, message: '请输入故障描述', trigger: 'blur' }]
}

function getFaultTypeColor(type: string) {
  const colors: Record<string, string> = {
    '制动故障': 'red',
    '转向故障': 'orange',
    '发动机故障': 'red',
    '电气故障': 'orange',
    '悬挂故障': 'yellow',
    '其他': 'default'
  }
  return colors[type] || 'default'
}

function getRiskColor(prob: number) {
  if (prob >= 80) return 'red'
  if (prob >= 60) return 'orange'
  if (prob >= 40) return 'yellow'
  return 'green'
}

function goBack() {
  router.push('/aftersale/fault')
}

function fillTestData() {
  form.partID = 'PART-' + Date.now().toString().slice(-6)
  form.vin = 'VIN' + Date.now().toString().slice(-13)
  form.faultType = '制动故障'
  form.riskProb = 75
  form.description = '测试故障描述：在正常行驶过程中出现制动不灵敏的情况，需要更长的制动距离。'
  message.success('已填充测试数据')
}

async function handleSubmit() {
  if (submitting.value) return
  submitting.value = true

  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    createdFaultID.value = `FAULT-${Date.now().toString().slice(-8)}`
    showSuccessModal.value = true
    message.success('故障上报成功')
  } catch (error: any) {
    message.error(error.message || '上报失败')
  } finally {
    submitting.value = false
  }
}

function handleReset() {
  formRef.value?.resetFields()
  form.partID = ''
  form.vin = ''
  form.faultType = ''
  form.riskProb = 50
  form.description = ''
}

function handleCreateAnother() {
  showSuccessModal.value = false
  createdFaultID.value = ''
  handleReset()
  fillTestData()
}
</script>

<style scoped>
.fault-create {
  min-height: 100vh;
  background: var(--bg-color);
  padding: 0;
  display: flex;
  flex-direction: column;
}

.fault-create :deep(.ant-select-dropdown) {
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
  background: linear-gradient(135deg, #ff4d4f 0%, #ff7875 100%);
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

.custom-input-number {
  width: 100%;
}

.custom-input-number :deep(.ant-input-number) {
  width: 100%;
  height: 42px;
  border-radius: 10px;
  border: 2px solid var(--border-color);
  background: var(--bg-color-secondary);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.custom-input-number :deep(.ant-input-number:hover) {
  border-color: var(--primary-color);
  box-shadow: 0 2px 8px rgba(24, 144, 255, 0.15);
}

.custom-input-number :deep(.ant-input-number-focused) {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(24, 144, 255, 0.12);
}

.custom-input-number :deep(.ant-input-number-input) {
  height: 38px;
  font-size: 14px;
  padding: 0 12px;
  border: none;
  outline: none;
  background: transparent;
}

.custom-input-number :deep(.ant-input-number-input-wrap) {
  height: 100%;
  display: flex;
  align-items: center;
}

.custom-input-number :deep(.ant-input-number-handler-wrap) {
  display: none;
}

.custom-textarea {
  width: 100%;
}

.custom-textarea :deep(textarea) {
  border-radius: 10px;
  border: 2px solid var(--border-color);
  background: var(--bg-color-secondary);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  font-size: 14px;
  padding: 12px;
}

.custom-textarea :deep(textarea:hover) {
  border-color: var(--primary-color);
}

.custom-textarea :deep(textarea:focus) {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(24, 144, 255, 0.12);
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
  font-size: 15px;
  font-weight: 500;
  background: linear-gradient(135deg, var(--primary-color) 0%, #722ed1 100%);
  border: none;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
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
  width: 360px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.info-card {
  border-radius: 16px;
  box-shadow: var(--shadow-sm);
  background: var(--bg-color-secondary);
}

.info-card :deep(.ant-card-body) {
  padding: 20px;
}

.card-title {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 15px;
  font-weight: 600;
  color: var(--text-color);
  margin-bottom: 16px;
}

.card-title svg {
  font-size: 18px;
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
  padding: 10px 0;
  border-bottom: 1px solid var(--border-color);
  font-size: 13px;
  color: var(--text-color-secondary);
  line-height: 1.6;
}

.guide-list li:last-child {
  border-bottom: none;
}

.guide-icon {
  flex-shrink: 0;
  margin-top: 2px;
}

.guide-icon.success {
  color: #52c41a;
}

.guide-icon.warning {
  color: #faad14;
}

.guide-list li em {
  color: var(--primary-color);
  font-style: normal;
  font-weight: 500;
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
  background: var(--bg-color-tertiary);
  border-radius: 8px;
}

.preview-label {
  font-size: 13px;
  color: var(--text-color-secondary);
  font-weight: 500;
}

.preview-value {
  font-size: 13px;
  color: var(--text-color);
  font-weight: 500;
  text-align: right;
  max-width: 200px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.quick-actions {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.quick-action-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 16px;
  background: var(--bg-color-tertiary);
  border-radius: 10px;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  font-size: 14px;
  color: var(--text-color);
  font-weight: 500;
}

.quick-action-item:hover {
  background: var(--primary-color);
  color: #fff;
  transform: translateX(4px);
}

.action-icon {
  font-size: 18px;
}

.success-modal :deep(.ant-modal-content) {
  border-radius: 20px;
  padding: 32px;
  text-align: center;
}

.success-content {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.success-icon-wrapper {
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, #52c41a 0%, #73d13d 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 20px;
  box-shadow: 0 8px 24px rgba(82, 196, 26, 0.3);
}

.success-icon {
  font-size: 48px;
  color: #fff;
}

.success-title {
  font-size: 24px;
  font-weight: 600;
  margin: 0 0 8px 0;
  color: var(--text-color);
}

.success-desc {
  font-size: 14px;
  color: var(--text-color-secondary);
  margin: 0 0 20px 0;
}

.success-fault-id {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  background: var(--bg-color-tertiary);
  border-radius: 10px;
  margin-bottom: 24px;
}

.success-fault-id .label {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.success-fault-id .value {
  font-size: 15px;
  font-weight: 600;
  color: var(--primary-color);
}

.success-actions {
  display: flex;
  gap: 12px;
  width: 100%;
}

.success-actions .ant-btn {
  flex: 1;
  height: 44px;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 500;
}

@media (max-width: 1200px) {
  .main-content {
    flex-direction: column;
  }

  .info-container {
    width: 100%;
  }

  .form-row {
    flex-direction: column;
    gap: 20px;
  }

  .form-item.half {
    flex: 1;
  }
}
</style>
