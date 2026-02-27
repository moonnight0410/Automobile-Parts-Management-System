<!--
  物流记录创建页面
  设计：与生产数据创建页面风格统一
-->

<template>
  <div class="logistics-create">
    <!-- 顶部导航栏 -->
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">添加物流记录</h1>
          <span class="page-subtitle">填写物流信息与配送详情</span>
        </div>
      </div>
      <div class="top-bar-right">
        <a-tag :color="formStatus.color" class="status-badge">
          <component v-if="formStatus.icon" :is="formStatus.icon" />
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
              <EnvironmentOutlined />
            </div>
            <div class="header-text">
              <h3>物流信息</h3>
              <p>请填写以下信息添加物流记录</p>
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
            <!-- 物流ID和订单ID -->
            <div class="form-row">
              <a-form-item label="物流ID" name="logisticsID" class="form-item half">
                <a-input
                  v-model:value="form.logisticsID"
                  placeholder="请输入物流ID"
                  :maxlength="50"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><NumberOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
              <a-form-item label="订单ID" name="orderID" class="form-item half">
                <a-input
                  v-model:value="form.orderID"
                  placeholder="请输入订单ID"
                  :maxlength="50"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><ShoppingCartOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
            </div>

            <!-- 承运商和出发时间 -->
            <div class="form-row">
              <a-form-item label="承运商" name="carrier" class="form-item half">
                <a-select
                  v-model:value="form.carrier"
                  placeholder="请选择承运商"
                  :options="carrierOptions"
                  class="custom-select"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                  :dropdown-style="{ zIndex: 1050 }"
                />
              </a-form-item>
              <a-form-item label="出发时间" name="startTime" class="form-item half">
                <a-date-picker
                  v-model:value="form.startTime"
                  placeholder="请选择出发时间"
                  show-time
                  format="YYYY-MM-DD HH:mm:ss"
                  class="custom-date-picker"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                />
              </a-form-item>
            </div>

            <!-- 送达时间和签收人 -->
            <div class="form-row">
              <a-form-item label="送达时间" name="endTime" class="form-item half">
                <a-date-picker
                  v-model:value="form.endTime"
                  placeholder="请选择送达时间"
                  show-time
                  format="YYYY-MM-DD HH:mm:ss"
                  class="custom-date-picker"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                />
              </a-form-item>
              <a-form-item label="签收人" name="receiver" class="form-item half">
                <a-input
                  v-model:value="form.receiver"
                  placeholder="请输入签收人"
                  :maxlength="50"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><UserOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
            </div>

            <!-- GPS轨迹数据哈希 -->
            <a-form-item label="GPS轨迹数据哈希" name="gpsHash" class="form-item">
              <a-textarea
                v-model:value="form.gpsHash"
                placeholder="请输入GPS轨迹数据哈希（可选）"
                :rows="3"
                :maxlength="200"
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
                {{ submitting ? '提交中...' : '提交物流记录' }}
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
            <span>物流说明</span>
          </div>
          <ul class="guide-list">
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>物流ID可自动生成或手动输入</span>
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
              <span>请确保订单ID和承运商信息准确</span>
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
              <span class="preview-label">物流ID</span>
              <span class="preview-value">{{ form.logisticsID || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">订单ID</span>
              <span class="preview-value">{{ form.orderID || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">承运商</span>
              <span class="preview-value">{{ form.carrier || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">出发时间</span>
              <span class="preview-value">{{ formatDateTime(form.startTime) || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">送达时间</span>
              <span class="preview-value">{{ formatDateTime(form.endTime) || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">签收人</span>
              <span class="preview-value">{{ form.receiver || '-' }}</span>
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
            <div class="quick-action-item" @click="generateLogisticsID">
              <NumberOutlined class="action-icon" />
              <span>生成物流ID</span>
            </div>
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
        <h2 class="success-title">物流记录添加成功</h2>
        <p class="success-desc">物流信息已成功提交到区块链</p>
        <div class="success-logistics-id">
          <span class="label">物流ID：</span>
          <span class="value">{{ createdLogisticsID }}</span>
        </div>
        <div class="success-actions">
          <a-button type="primary" @click="handleCreateAnother">
            <template #icon><PlusOutlined /></template>
            继续添加
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
import { createLogisticsData } from '@/services/supply.service'
import dayjs, { Dayjs } from 'dayjs'
import {
  ArrowLeftOutlined,
  CheckOutlined,
  ReloadOutlined,
  InfoCircleOutlined,
  CheckCircleOutlined,
  EyeOutlined,
  EnvironmentOutlined,
  NumberOutlined,
  ShoppingCartOutlined,
  UserOutlined,
  ThunderboltOutlined,
  PlusOutlined,
  UnorderedListOutlined,
  CheckCircleFilled
} from '@ant-design/icons-vue'

const router = useRouter()
const formRef = ref<FormInstance>()
const submitting = ref(false)
const showSuccessModal = ref(false)
const createdLogisticsID = ref('')

const form = reactive({
  logisticsID: '',
  orderID: '',
  carrier: '',
  startTime: null as Dayjs | null,
  endTime: null as Dayjs | null,
  receiver: '',
  gpsHash: ''
})

const carrierOptions = ref<SelectProps['options']>([
  { label: '顺丰物流', value: '顺丰物流' },
  { label: '京东物流', value: '京东物流' },
  { label: '中通快递', value: '中通快递' },
  { label: '圆通速递', value: '圆通速递' },
  { label: '申通快递', value: '申通快递' },
  { label: '韵达速递', value: '韵达速递' },
  { label: '德邦物流', value: '德邦物流' }
])

const formStatus = computed(() => {
  if (submitting.value) {
    return { text: '提交中', color: 'processing', icon: ReloadOutlined }
  }
  if (isFormValid.value) {
    return { text: '可提交', color: 'success', icon: CheckCircleOutlined }
  }
  return { text: '待填写', color: 'warning', icon: EnvironmentOutlined }
})

const isFormValid = computed(() => {
  return (
    form.logisticsID.trim() !== '' &&
    form.orderID.trim() !== '' &&
    form.carrier.trim() !== ''
  )
})

const formRules = {
  logisticsID: [{ required: true, message: '请输入物流ID', trigger: 'blur' }],
  orderID: [{ required: true, message: '请输入订单ID', trigger: 'blur' }],
  carrier: [{ required: true, message: '请选择承运商', trigger: 'change' }]
}

function formatDateTime(date: Dayjs | null) {
  if (!date) return ''
  return date.format('YYYY-MM-DD HH:mm:ss')
}

function goBack() {
  router.push('/supply/logistics')
}

function generateLogisticsID() {
  const timestamp = Date.now().toString().slice(-8)
  form.logisticsID = `LOG-${timestamp}`
  message.success('已生成物流ID')
}

function fillTestData() {
  form.logisticsID = `LOG-${Date.now().toString().slice(-8)}`
  form.orderID = `ORDER-${Date.now().toString().slice(-8)}`
  form.carrier = '顺丰物流'
  form.startTime = dayjs()
  form.endTime = dayjs().add(3, 'day')
  form.receiver = '张三'
  form.gpsHash = 'GPS-HASH-' + Date.now().toString().slice(-6)
  message.success('已填充测试数据')
}

async function handleSubmit() {
  if (submitting.value) return
  submitting.value = true

  try {
    await createLogisticsData({
      logisticsID: form.logisticsID,
      orderID: form.orderID,
      carrier: form.carrier,
      startTime: form.startTime ? formatDateTime(form.startTime) : '',
      endTime: form.endTime ? formatDateTime(form.endTime) : '',
      receiver: form.receiver,
      gpsHash: form.gpsHash
    })
    createdLogisticsID.value = form.logisticsID
    showSuccessModal.value = true
    message.success('物流记录添加成功')
  } catch (error: any) {
    message.error(error.message || '添加失败')
  } finally {
    submitting.value = false
  }
}

function handleReset() {
  formRef.value?.resetFields()
  form.logisticsID = ''
  form.orderID = ''
  form.carrier = ''
  form.startTime = null
  form.endTime = null
  form.receiver = ''
  form.gpsHash = ''
}

function handleCreateAnother() {
  showSuccessModal.value = false
  createdLogisticsID.value = ''
  handleReset()
  generateLogisticsID()
}
</script>

<style scoped>
.logistics-create {
  min-height: 100vh;
  background: var(--bg-color);
  padding: 0;
  display: flex;
  flex-direction: column;
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
  margin-bottom: 24px;
}

.header-icon {
  width: 56px;
  height: 56px;
  border-radius: 12px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  color: white;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.header-text h3 {
  margin: 0 0 4px 0;
  font-size: 20px;
  font-weight: 600;
  color: var(--text-color);
}

.header-text p {
  margin: 0;
  font-size: 14px;
  color: var(--text-color-secondary);
}

.form-divider {
  margin: 0 0 32px 0;
  border-color: var(--border-color);
}

.create-form {
  max-width: 100%;
}

.form-row {
  display: flex;
  gap: 20px;
  margin-bottom: 0;
}

.form-row .form-item {
  flex: 1;
}

.form-item {
  margin-bottom: 24px;
}

.form-item :deep(.ant-form-item-label > label) {
  font-weight: 500;
  color: var(--text-color);
  font-size: 14px;
}

.custom-input {
  height: 48px;
  border-radius: 12px;
  border: 2px solid var(--border-color);
  background: var(--bg-color-tertiary);
  transition: all 0.3s;
  font-size: 14px;
}

.custom-input:hover,
.custom-input:focus {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(24, 144, 255, 0.1);
}

.custom-input :deep(.ant-input) {
  background: transparent;
  height: 100%;
}

.input-icon {
  font-size: 16px;
  color: var(--text-color-secondary);
}

.custom-select {
  height: 48px;
}

.custom-select :deep(.ant-select-selector) {
  height: 48px !important;
  border-radius: 12px !important;
  border: 2px solid var(--border-color) !important;
  background: var(--bg-color-tertiary) !important;
  transition: all 0.3s;
  font-size: 14px;
}

.custom-select:hover :deep(.ant-select-selector),
.custom-select :deep(.ant-select-focused .ant-select-selector) {
  border-color: var(--primary-color) !important;
  box-shadow: 0 0 0 3px rgba(24, 144, 255, 0.1);
}

.custom-date-picker {
  width: 100%;
  height: 48px;
}

.custom-date-picker :deep(.ant-picker) {
  width: 100%;
  height: 48px;
  border-radius: 12px;
  border: 2px solid var(--border-color);
  background: var(--bg-color-tertiary);
  transition: all 0.3s;
}

.custom-date-picker :deep(.ant-picker:hover),
.custom-date-picker :deep(.ant-picker-focused) {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(24, 144, 255, 0.1);
}

.custom-textarea {
  border-radius: 12px;
  border: 2px solid var(--border-color);
  background: var(--bg-color-tertiary);
  transition: all 0.3s;
  font-size: 14px;
  resize: vertical;
}

.custom-textarea:hover,
.custom-textarea:focus {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(24, 144, 255, 0.1);
}

.form-actions {
  display: flex;
  gap: 12px;
  margin-top: 32px;
  padding-top: 24px;
  border-top: 1px solid var(--border-color);
}

.submit-btn {
  height: 48px;
  padding: 0 32px;
  border-radius: 12px;
  font-weight: 600;
  font-size: 16px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
  transition: all 0.3s;
}

.submit-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(102, 126, 234, 0.4);
}

.submit-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.secondary-btn {
  height: 48px;
  padding: 0 32px;
  border-radius: 12px;
  font-weight: 500;
  font-size: 16px;
  border: 2px solid var(--border-color);
  background: transparent;
  transition: all 0.3s;
}

.secondary-btn:hover {
  border-color: var(--primary-color);
  color: var(--primary-color);
}

.info-container {
  width: 380px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.info-card {
  border-radius: 16px;
  box-shadow: var(--shadow-sm);
  background: var(--bg-color-secondary);
  overflow: hidden;
}

.info-card :deep(.ant-card-body) {
  padding: 20px;
}

.card-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 16px;
  font-weight: 600;
  color: var(--text-color);
  margin-bottom: 16px;
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
  font-size: 14px;
  color: var(--text-color-secondary);
}

.guide-list li:last-child {
  border-bottom: none;
}

.guide-icon {
  font-size: 16px;
  flex-shrink: 0;
  margin-top: 2px;
}

.guide-icon.success {
  color: #52c41a;
}

.guide-icon.warning {
  color: #faad14;
}

.guide-list em {
  color: #ff4d4f;
  font-style: normal;
  font-weight: 600;
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
  padding: 10px 0;
  border-bottom: 1px solid var(--border-color);
}

.preview-item:last-child {
  border-bottom: none;
}

.preview-label {
  font-size: 13px;
  color: var(--text-color-secondary);
  font-weight: 500;
}

.preview-value {
  font-size: 14px;
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
  border-radius: 10px;
  background: var(--bg-color-tertiary);
  cursor: pointer;
  transition: all 0.3s;
  font-size: 14px;
  color: var(--text-color);
  font-weight: 500;
}

.quick-action-item:hover {
  background: var(--primary-color);
  color: white;
  transform: translateX(4px);
}

.action-icon {
  font-size: 18px;
  color: var(--primary-color);
  transition: all 0.3s;
}

.quick-action-item:hover .action-icon {
  color: white;
}

.success-modal :deep(.ant-modal-content) {
  border-radius: 20px;
  overflow: hidden;
}

.success-modal :deep(.ant-modal-body) {
  padding: 40px 32px;
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
  border-radius: 50%;
  background: linear-gradient(135deg, #52c41a 0%, #73d13d 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 24px;
  box-shadow: 0 8px 24px rgba(82, 196, 26, 0.3);
}

.success-icon {
  font-size: 40px;
  color: white;
}

.success-title {
  font-size: 24px;
  font-weight: 600;
  color: var(--text-color);
  margin: 0 0 12px 0;
}

.success-desc {
  font-size: 14px;
  color: var(--text-color-secondary);
  margin: 0 0 24px 0;
}

.success-logistics-id {
  padding: 16px 24px;
  background: var(--bg-color-tertiary);
  border-radius: 12px;
  margin-bottom: 24px;
  font-size: 14px;
}

.success-logistics-id .label {
  color: var(--text-color-secondary);
  margin-right: 8px;
}

.success-logistics-id .value {
  color: var(--primary-color);
  font-weight: 600;
  font-size: 16px;
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
  font-weight: 500;
}

@media (max-width: 1200px) {
  .main-content {
    flex-direction: column;
  }

  .info-container {
    width: 100%;
  }
}

@media (max-width: 768px) {
  .top-bar {
    padding: 12px 16px;
  }

  .page-title {
    font-size: 16px;
  }

  .main-content {
    padding: 16px;
  }

  .form-row {
    flex-direction: column;
    gap: 0;
  }

  .form-actions {
    flex-direction: column;
  }

  .submit-btn,
  .secondary-btn {
    width: 100%;
  }
}
</style>
