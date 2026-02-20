<!--
  采购订单创建页面
  设计：与QualityInspectionCreate风格统一
-->

<template>
  <div class="order-create">
    <!-- 顶部导航栏 -->
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">创建采购订单</h1>
          <span class="page-subtitle">填写订单信息</span>
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
            <div class="header-icon order">
              <ShoppingCartOutlined />
            </div>
            <div class="header-text">
              <h3>订单信息</h3>
              <p>请填写以下信息创建采购订单</p>
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
            <!-- 订单ID -->
            <a-form-item label="订单ID" name="orderID" class="form-item">
              <a-input
                v-model:value="form.orderID"
                placeholder="自动生成或手动输入"
                :maxlength="50"
                class="custom-input"
              >
                <template #prefix>
                  <span class="input-icon"><NumberOutlined /></span>
                </template>
              </a-input>
            </a-form-item>

            <!-- 采购方和供应商 -->
            <div class="form-row">
              <a-form-item label="采购方" name="buyer" class="form-item half">
                <a-input
                  v-model:value="form.buyer"
                  placeholder="请输入采购方名称"
                  :maxlength="50"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><UserOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
              <a-form-item label="供应商" name="supplier" class="form-item half">
                <a-input
                  v-model:value="form.supplier"
                  placeholder="请输入供应商名称"
                  :maxlength="50"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><ShopOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
            </div>

            <!-- 订单状态和订单金额 -->
            <div class="form-row">
              <a-form-item label="订单状态" name="status" class="form-item half">
                <a-select
                  v-model:value="form.status"
                  placeholder="请选择订单状态"
                  :options="statusOptions"
                  class="custom-select"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                  :dropdown-style="{ zIndex: 1050 }"
                />
              </a-form-item>
              <a-form-item label="订单金额" name="amount" class="form-item half">
                <a-input-number
                  v-model:value="form.amount"
                  placeholder="请输入订单金额"
                  :min="0"
                  :max="9999999"
                  :precision="2"
                  class="custom-input-number"
                />
              </a-form-item>
            </div>

            <!-- 交货日期 -->
            <a-form-item label="交货日期" name="deliveryDate" class="form-item">
              <a-date-picker
                v-model:value="form.deliveryDate"
                placeholder="请选择交货日期"
                class="custom-date-picker"
                style="width: 100%"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
              />
            </a-form-item>

            <!-- 备注 -->
            <a-form-item label="备注" name="remark" class="form-item">
              <a-textarea
                v-model:value="form.remark"
                placeholder="请输入备注信息（可选）"
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
                {{ submitting ? '提交中...' : '提交订单' }}
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
            <span>创建说明</span>
          </div>
          <ul class="guide-list">
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>订单ID可自动生成或手动输入</span>
            </li>
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>带 <em>*</em> 的为必填项</span>
            </li>
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>提交后订单将记录到区块链</span>
            </li>
            <li>
              <CheckCircleOutlined class="guide-icon warning" />
              <span>订单金额单位为元</span>
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
              <span class="preview-label">订单ID</span>
              <span class="preview-value">{{ form.orderID || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">采购方</span>
              <span class="preview-value">{{ form.buyer || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">供应商</span>
              <span class="preview-value">{{ form.supplier || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">订单状态</span>
              <a-tag :color="getStatusColor(form.status)" size="small">
                {{ form.status || '-' }}
              </a-tag>
            </div>
            <div class="preview-item">
              <span class="preview-label">订单金额</span>
              <span class="preview-value">{{ form.amount ? `¥${form.amount.toFixed(2)}` : '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">交货日期</span>
              <span class="preview-value">{{ formatDate(form.deliveryDate) }}</span>
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
            <div class="quick-action-item" @click="generateOrderID">
              <NumberOutlined class="action-icon" />
              <span>生成订单ID</span>
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
        <h2 class="success-title">创建成功</h2>
        <p class="success-desc">采购订单已成功提交到区块链</p>
        <div class="success-part-id">
          <span class="label">订单ID：</span>
          <span class="value">{{ createdOrderID }}</span>
        </div>
        <div class="success-result">
          <span class="label">订单金额：</span>
          <span class="value">¥{{ form.amount?.toFixed(2) || '0.00' }}</span>
        </div>
        <div class="success-actions">
          <a-button type="primary" @click="handleCreateAnother">
            <template #icon><PlusOutlined /></template>
            继续创建
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
import dayjs, { Dayjs } from 'dayjs'
import {
  ArrowLeftOutlined,
  CheckOutlined,
  ReloadOutlined,
  InfoCircleOutlined,
  CheckCircleOutlined,
  EyeOutlined,
  ShoppingCartOutlined,
  NumberOutlined,
  UserOutlined,
  ShopOutlined,
  ThunderboltOutlined,
  PlusOutlined,
  UnorderedListOutlined,
  CheckCircleFilled
} from '@ant-design/icons-vue'

const router = useRouter()
const formRef = ref<FormInstance>()
const submitting = ref(false)
const showSuccessModal = ref(false)
const createdOrderID = ref('')

const form = reactive({
  orderID: '',
  buyer: '',
  supplier: '',
  status: '',
  amount: 0,
  deliveryDate: null as Dayjs | null,
  remark: ''
})

const formStatus = computed(() => {
  if (submitting.value) {
    return { text: '提交中', color: 'processing', icon: ReloadOutlined }
  }
  if (isFormValid.value) {
    return { text: '可提交', color: 'success', icon: CheckCircleOutlined }
  }
  return { text: '待填写', color: 'warning', icon: ShoppingCartOutlined }
})

const statusOptions = computed<SelectProps['options']>(() => [
  { value: '待处理', label: '待处理' },
  { value: '运输中', label: '运输中' },
  { value: '已签收', label: '已签收' }
])

const isFormValid = computed(() => {
  return (
    form.buyer.trim() !== '' &&
    form.supplier.trim() !== '' &&
    form.status !== '' &&
    form.amount > 0 &&
    form.deliveryDate !== null
  )
})

const formRules = {
  buyer: [{ required: true, message: '请输入采购方名称', trigger: 'blur' }],
  supplier: [{ required: true, message: '请输入供应商名称', trigger: 'blur' }],
  status: [{ required: true, message: '请选择订单状态', trigger: 'change' }],
  amount: [
    { required: true, message: '请输入订单金额', trigger: 'blur' },
    { type: 'number', min: 0.01, message: '订单金额必须大于0', trigger: 'blur' }
  ],
  deliveryDate: [{ required: true, message: '请选择交货日期', trigger: 'change' }]
}

function getStatusColor(status: string) {
  if (status === '待处理') return 'orange'
  if (status === '运输中') return 'blue'
  if (status === '已签收') return 'green'
  return 'default'
}

function formatDate(date: Dayjs | null) {
  if (!date) return '-'
  return date.format('YYYY-MM-DD')
}

function goBack() {
  router.push('/supply/orders')
}

function generateOrderID() {
  const timestamp = Date.now().toString().slice(-8)
  form.orderID = `ORD-${timestamp}`
  message.success('已生成订单ID')
}

function fillTestData() {
  form.buyer = '测试采购方'
  form.supplier = '测试供应商'
  form.status = '待处理'
  form.amount = 10000
  form.deliveryDate = dayjs().add(7, 'day')
  message.success('已填充测试数据')
}

async function handleSubmit() {
  if (submitting.value) return
  submitting.value = true

  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    createdOrderID.value = form.orderID || `ORD-${Date.now().toString().slice(-8)}`
    showSuccessModal.value = true
    message.success('采购订单创建成功')
  } catch (error: any) {
    message.error(error.message || '创建失败')
  } finally {
    submitting.value = false
  }
}

function handleReset() {
  formRef.value?.resetFields()
  form.orderID = ''
  form.buyer = ''
  form.supplier = ''
  form.status = ''
  form.amount = 0
  form.deliveryDate = null
  form.remark = ''
}

function handleCreateAnother() {
  showSuccessModal.value = false
  createdOrderID.value = ''
  handleReset()
  generateOrderID()
}
</script>

<style scoped>
.order-create {
  min-height: 100vh;
  background: var(--bg-color);
  padding: 0;
  display: flex;
  flex-direction: column;
}

.order-create :deep(.ant-select-dropdown) {
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

.header-icon.order {
  background: linear-gradient(135deg, #fa8c16 0%, #d46b08 100%);
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
.custom-select,
.custom-date-picker {
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

.custom-input :deep(.ant-input-affix-wrapper:hover),
.custom-select :deep(.ant-select-selector:hover),
.custom-date-picker :deep(.ant-picker:hover .ant-picker-input) {
  border-color: var(--primary-color);
  box-shadow: 0 2px 8px rgba(24, 144, 255, 0.15);
}

.custom-input :deep(.ant-input-affix-wrapper:focus),
.custom-input :deep(.ant-input-affix-wrapper-focused),
.custom-select :deep(.ant-select-focused .ant-select-selector),
.custom-date-picker :deep(.ant-picker-focused) {
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

.custom-date-picker :deep(.ant-picker) {
  width: 100%;
  height: 42px;
  border-radius: 10px;
  border: 2px solid var(--border-color);
  background: var(--bg-color-secondary);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.custom-date-picker :deep(.ant-picker-input) {
  height: 38px;
  font-size: 14px;
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

.custom-date-picker :deep(.ant-picker-dropdown) {
  z-index: 1050 !important;
  border-radius: 12px;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.15);
  background: #ffffff;
  padding: 0;
  animation: datePickerFadeIn 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes datePickerFadeIn {
  from {
    opacity: 0;
    transform: translateY(-8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.custom-date-picker :deep(.ant-picker-panel-container) {
  border-radius: 12px;
  overflow: hidden;
  background: #ffffff;
  border: none;
}

.custom-date-picker :deep(.ant-picker-panel) {
  border: none;
  background: #ffffff;
}

.custom-date-picker :deep(.ant-picker-header) {
  border-bottom: 1px solid #f0f0f0;
  padding: 12px 16px;
  background: linear-gradient(135deg, #fafbfc 0%, #f5f7fa 100%);
}

.custom-date-picker :deep(.ant-picker-header-view) {
  font-weight: 600;
  font-size: 15px;
  color: #1f2937;
}

.custom-date-picker :deep(.ant-picker-header-view button) {
  color: #1f2937;
  font-weight: 600;
  transition: color 0.2s ease;
}

.custom-date-picker :deep(.ant-picker-header-view button:hover) {
  color: #1890ff;
}

.custom-date-picker :deep(.ant-picker-header button) {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
  color: #6b7280;
}

.custom-date-picker :deep(.ant-picker-header button:hover) {
  background: #e0f2fe;
  color: #1890ff;
}

.custom-date-picker :deep(.ant-picker-body) {
  padding: 12px 16px 16px;
  background: #ffffff;
}

.custom-date-picker :deep(.ant-picker-content) {
  border-collapse: separate;
  border-spacing: 4px;
}

.custom-date-picker :deep(.ant-picker-content th) {
  font-size: 12px;
  font-weight: 600;
  color: #9ca3af;
  padding: 8px 0;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.custom-date-picker :deep(.ant-picker-cell) {
  padding: 4px 0;
  transition: all 0.2s ease;
}

.custom-date-picker :deep(.ant-picker-cell-inner) {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 500;
  color: #374151;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.custom-date-picker :deep(.ant-picker-cell-inner::before) {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 0;
  height: 0;
  border-radius: 50%;
  background: rgba(24, 144, 255, 0.1);
  transform: translate(-50%, -50%);
  transition: all 0.3s ease;
}

.custom-date-picker :deep(.ant-picker-cell:hover .ant-picker-cell-inner::before) {
  width: 100%;
  height: 100%;
  border-radius: 10px;
}

.custom-date-picker :deep(.ant-picker-cell:hover .ant-picker-cell-inner) {
  background: #e0f2fe;
  color: #1890ff;
  transform: scale(1.05);
}

.custom-date-picker :deep(.ant-picker-cell-selected .ant-picker-cell-inner) {
  background: linear-gradient(135deg, #1890ff 0%, #096dd9 100%);
  color: #ffffff;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3);
}

.custom-date-picker :deep(.ant-picker-cell-selected:hover .ant-picker-cell-inner) {
  background: linear-gradient(135deg, #40a9ff 0%, #1890ff 100%);
  transform: scale(1.05);
}

.custom-date-picker :deep(.ant-picker-cell-today .ant-picker-cell-inner) {
  border: 2px solid #1890ff;
  color: #1890ff;
  font-weight: 600;
}

.custom-date-picker :deep(.ant-picker-cell-today:hover .ant-picker-cell-inner) {
  background: #e0f2fe;
  border-color: #1890ff;
}

.custom-date-picker :deep(.ant-picker-cell-in-view) {
  color: #374151;
}

.custom-date-picker :deep(.ant-picker-cell-disabled) {
  opacity: 0.4;
  cursor: not-allowed;
}

.custom-date-picker :deep(.ant-picker-cell-disabled .ant-picker-cell-inner) {
  background: transparent;
  color: #d1d5db;
}

.custom-date-picker :deep(.ant-picker-cell-disabled:hover .ant-picker-cell-inner) {
  background: transparent;
  transform: none;
}

.custom-date-picker :deep(.ant-picker-range-wrapper) {
  background: #ffffff;
  border-radius: 12px;
}

.custom-date-picker :deep(.ant-picker-footer) {
  border-top: 1px solid #f0f0f0;
  padding: 12px 16px;
  background: #fafbfc;
  border-radius: 0 0 12px 12px;
}

.custom-date-picker :deep(.ant-picker-footer-extra) {
  font-size: 13px;
  color: #6b7280;
}

.custom-date-picker :deep(.ant-picker-today-btn) {
  font-weight: 500;
  color: #1890ff;
  transition: all 0.2s ease;
}

.custom-date-picker :deep(.ant-picker-today-btn:hover) {
  color: #40a9ff;
}

.custom-date-picker :deep(.ant-picker-year-panel),
.custom-date-picker :deep(.ant-picker-month-panel) {
  background: #ffffff;
}

.custom-date-picker :deep(.ant-picker-year-panel .ant-picker-cell-inner),
.custom-date-picker :deep(.ant-picker-month-panel .ant-picker-cell-inner) {
  width: 60px;
  height: 36px;
  border-radius: 8px;
}

.custom-date-picker :deep(.ant-picker-decade-panel .ant-picker-cell-inner) {
  width: 80px;
  height: 36px;
  border-radius: 8px;
}

@media (max-width: 768px) {
  .custom-date-picker :deep(.ant-picker-dropdown) {
    max-width: calc(100vw - 32px);
  }
  
  .custom-date-picker :deep(.ant-picker-panel-container) {
    width: 100%;
  }
  
  .custom-date-picker :deep(.ant-picker-cell-inner) {
    width: 32px;
    height: 32px;
    font-size: 13px;
  }
}

.custom-input-number :deep(.ant-input-number-handler-wrap) {
  display: none;
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

.quick-actions {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.quick-action-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 16px;
  background: var(--bg-color);
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
  font-size: 13px;
  color: var(--text-color-secondary);
}

.quick-action-item:hover {
  background: var(--bg-color-tertiary);
  color: var(--primary-color);
}

.action-icon {
  font-size: 16px;
}

.success-content {
  text-align: center;
  padding: 24px;
}

.success-icon-wrapper {
  width: 72px;
  height: 72px;
  background: linear-gradient(135deg, #52c41a 0%, #389e0d 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
}

.success-icon {
  font-size: 36px;
  color: #fff;
}

.success-title {
  font-size: 22px;
  font-weight: 600;
  color: var(--text-color);
  margin: 0 0 8px;
}

.success-desc {
  font-size: 14px;
  color: var(--text-color-secondary);
  margin: 0 0 20px;
}

.success-part-id {
  background: var(--bg-color);
  padding: 12px 20px;
  border-radius: 8px;
  margin-bottom: 12px;
}

.success-part-id .label {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.success-part-id .value {
  font-size: 16px;
  font-weight: 600;
  color: var(--primary-color);
  margin-left: 8px;
}

.success-result {
  background: var(--bg-color);
  padding: 12px 20px;
  border-radius: 8px;
  margin-bottom: 20px;
}

.success-result .label {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.success-result .value {
  font-size: 16px;
  font-weight: 600;
  color: var(--success-color);
  margin-left: 8px;
}

.success-actions {
  display: flex;
  gap: 12px;
  justify-content: center;
}

.success-actions .ant-btn {
  height: 40px;
  border-radius: 10px;
  font-weight: 500;
  font-size: 14px;
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

  .form-card :deep(.ant-card-body) {
    padding: 20px;
  }
}
</style>

<style>
.ant-picker-dropdown {
  z-index: 1050 !important;
  border-radius: 12px !important;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.15) !important;
  background: #ffffff !important;
  padding: 0 !important;
  animation: datePickerFadeIn 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes datePickerFadeIn {
  from {
    opacity: 0;
    transform: translateY(-8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.ant-picker-panel-container {
  border-radius: 12px !important;
  overflow: hidden !important;
  background: #ffffff !important;
  border: none !important;
}

.ant-picker-panel {
  border: none !important;
  background: #ffffff !important;
}

.ant-picker-header {
  border-bottom: 1px solid #f0f0f0 !important;
  padding: 12px 16px !important;
  background: linear-gradient(135deg, #fafbfc 0%, #f5f7fa 100%) !important;
  display: flex !important;
  align-items: center !important;
  justify-content: space-between !important;
  flex-wrap: nowrap !important;
  gap: 8px !important;
  min-height: 60px !important;
  box-sizing: border-box !important;
}

.ant-picker-header > button {
  flex-shrink: 0 !important;
  min-width: 36px !important;
  width: 36px !important;
  height: 36px !important;
  border-radius: 10px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  transition: all 0.2s ease !important;
  color: #000000 !important;
  cursor: pointer !important;
}

.ant-picker-header > button:hover {
  background: #e0f2fe !important;
  color: #1890ff !important;
}

.ant-picker-header-view {
  font-weight: 600 !important;
  font-size: 15px !important;
  color: #000000 !important;
  flex: 1 !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  gap: 8px !important;
  white-space: nowrap !important;
}

.ant-picker-header-view button {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 15px !important;
  transition: all 0.2s ease !important;
  padding: 6px 12px !important;
  border-radius: 6px !important;
  min-width: auto !important;
  width: auto !important;
  height: auto !important;
  white-space: nowrap !important;
  display: inline-flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.ant-picker-header-view button:hover {
  color: #1890ff !important;
  background: #e0f2fe !important;
}

.ant-picker-body {
  padding: 12px 16px 16px !important;
  background: #ffffff !important;
}

.ant-picker-content {
  border-collapse: separate !important;
  border-spacing: 4px !important;
}

.ant-picker-content th {
  font-size: 12px !important;
  font-weight: 600 !important;
  color: #9ca3af !important;
  padding: 8px 0 !important;
  text-transform: uppercase !important;
  letter-spacing: 0.5px !important;
}

.ant-picker-cell {
  padding: 4px 0 !important;
  transition: all 0.2s ease !important;
}

.ant-picker-cell-inner {
  width: 36px !important;
  height: 36px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  border-radius: 10px !important;
  font-size: 14px !important;
  font-weight: 500 !important;
  color: #374151 !important;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1) !important;
  cursor: pointer !important;
  position: relative !important;
  overflow: hidden !important;
}

.ant-picker-cell-inner::before {
  content: '' !important;
  position: absolute !important;
  top: 50% !important;
  left: 50% !important;
  width: 0 !important;
  height: 0 !important;
  border-radius: 50% !important;
  background: rgba(24, 144, 255, 0.1) !important;
  transform: translate(-50%, -50%) !important;
  transition: all 0.3s ease !important;
}

.ant-picker-cell:hover .ant-picker-cell-inner::before {
  width: 100% !important;
  height: 100% !important;
  border-radius: 10px !important;
}

.ant-picker-cell:hover .ant-picker-cell-inner {
  background: #e0f2fe !important;
  color: #1890ff !important;
  transform: scale(1.05) !important;
}

.ant-picker-cell-selected .ant-picker-cell-inner {
  background: linear-gradient(135deg, #1890ff 0%, #096dd9 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3) !important;
}

.ant-picker-cell-selected:hover .ant-picker-cell-inner {
  background: linear-gradient(135deg, #40a9ff 0%, #1890ff 100%) !important;
  transform: scale(1.05) !important;
}

.ant-picker-cell-today .ant-picker-cell-inner {
  border: 2px solid #1890ff !important;
  color: #1890ff !important;
  font-weight: 600 !important;
}

.ant-picker-cell-today:hover .ant-picker-cell-inner {
  background: #e0f2fe !important;
  border-color: #1890ff !important;
}

.ant-picker-cell-in-view {
  color: #374151 !important;
}

.ant-picker-cell-disabled {
  opacity: 0.4 !important;
  cursor: not-allowed !important;
}

.ant-picker-cell-disabled .ant-picker-cell-inner {
  background: transparent !important;
  color: #d1d5db !important;
}

.ant-picker-cell-disabled:hover .ant-picker-cell-inner {
  background: transparent !important;
  transform: none !important;
}

.ant-picker-range-wrapper {
  background: #ffffff !important;
  border-radius: 12px !important;
}

.ant-picker-footer {
  border-top: 1px solid #f0f0f0 !important;
  padding: 12px 16px !important;
  background: #fafbfc !important;
  border-radius: 0 0 12px 12px !important;
}

.ant-picker-footer-extra {
  font-size: 13px !important;
  color: #6b7280 !important;
}

.ant-picker-today-btn {
  font-weight: 500 !important;
  color: #1890ff !important;
  transition: all 0.2s ease !important;
}

.ant-picker-today-btn:hover {
  color: #40a9ff !important;
}

.ant-picker-year-panel,
.ant-picker-month-panel {
  background: #ffffff !important;
}

.ant-picker-year-panel .ant-picker-cell-inner,
.ant-picker-month-panel .ant-picker-cell-inner {
  width: 60px !important;
  height: 36px !important;
  border-radius: 8px !important;
}

.ant-picker-decade-panel .ant-picker-cell-inner {
  width: 80px !important;
  height: 36px !important;
  border-radius: 8px !important;
}

@media (max-width: 768px) {
  .ant-picker-dropdown {
    max-width: calc(100vw - 32px) !important;
  }
  
  .ant-picker-panel-container {
    width: 100% !important;
  }
  
  .ant-picker-cell-inner {
    width: 32px !important;
    height: 32px !important;
    font-size: 13px !important;
  }
}
</style>
