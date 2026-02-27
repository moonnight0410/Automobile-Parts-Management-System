<!--
  召回发起页面
  设计：与生产数据创建页面风格统一
-->

<template>
  <div class="recall-create">
    <!-- 顶部导航栏 -->
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">发起召回</h1>
          <span class="page-subtitle">填写召回信息与处理方案</span>
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
              <WarningOutlined />
            </div>
            <div class="header-text">
              <h3>召回信息</h3>
              <p>请填写以下信息发起召回</p>
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
            <!-- 召回ID -->
            <a-form-item label="召回ID" name="recallID" class="form-item">
              <a-input
                v-model:value="form.recallID"
                placeholder="自动生成或手动输入"
                :maxlength="50"
                class="custom-input"
              >
                <template #prefix>
                  <span class="input-icon"><NumberOutlined /></span>
                </template>
              </a-input>
            </a-form-item>

            <!-- 涉及批次 -->
            <a-form-item label="涉及批次" name="batchNos" class="form-item enhanced-form-item">
              <a-select
                v-model:value="form.batchNos"
                mode="tags"
                placeholder="请输入批次号，按回车添加多个批次"
                :max-tag-count="5"
                :token-separators="[',']"
                :options="[]"
                class="custom-select enhanced-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                :dropdown-style="{ zIndex: 1050 }"
                :open="false"
              >
                <template #suffixIcon>
                  <AppstoreOutlined class="select-suffix-icon" />
                </template>
              </a-select>
            </a-form-item>

            <!-- 受影响零部件 -->
            <a-form-item label="受影响零部件" name="affectedParts" class="form-item enhanced-form-item">
              <a-select
                v-model:value="form.affectedParts"
                mode="tags"
                placeholder="请输入零部件ID，按回车添加多个零部件"
                :max-tag-count="5"
                :token-separators="[',']"
                :options="[]"
                class="custom-select enhanced-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                :dropdown-style="{ zIndex: 1050 }"
                :open="false"
              >
                <template #suffixIcon>
                  <AppstoreOutlined class="select-suffix-icon" />
                </template>
              </a-select>
            </a-form-item>

            <!-- 召回原因 -->
            <a-form-item label="召回原因" name="reason" class="form-item">
              <a-textarea
                v-model:value="form.reason"
                placeholder="请详细描述召回原因"
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
                {{ submitting ? '提交中...' : '提交召回' }}
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
            <span>召回说明</span>
          </div>
          <ul class="guide-list">
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>召回ID可自动生成或手动输入</span>
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
              <span>请确保零部件ID和批次号准确</span>
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
              <span class="preview-label">召回ID</span>
              <span class="preview-value">{{ form.recallID || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">涉及批次</span>
              <span class="preview-value">{{ form.batchNos.join(', ') || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">受影响零部件</span>
              <span class="preview-value">{{ form.affectedParts.join(', ') || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">状态</span>
              <a-tag color="orange" size="small">
                {{ form.status }}
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
          <div class="quick-actions">
            <div class="quick-action-item" @click="generateRecallID">
              <NumberOutlined class="action-icon" />
              <span>生成召回ID</span>
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
        <h2 class="success-title">召回发起成功</h2>
        <p class="success-desc">召回信息已成功提交到区块链</p>
        <div class="success-recall-id">
          <span class="label">召回ID：</span>
          <span class="value">{{ createdRecallID }}</span>
        </div>
        <div class="success-actions">
          <a-button type="primary" @click="handleCreateAnother">
            <template #icon><PlusOutlined /></template>
            继续发起
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
import { createRecallRecord } from '@/services/aftersale.service'
import {
  ArrowLeftOutlined,
  CheckOutlined,
  ReloadOutlined,
  InfoCircleOutlined,
  CheckCircleOutlined,
  EyeOutlined,
  WarningOutlined,
  NumberOutlined,
  ThunderboltOutlined,
  PlusOutlined,
  UnorderedListOutlined,
  CheckCircleFilled,
  AppstoreOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const formRef = ref<FormInstance>()
const submitting = ref(false)
const showSuccessModal = ref(false)
const createdRecallID = ref('')

const form = reactive({
  recallID: '',
  batchNos: [] as string[],
  reason: '',
  affectedParts: [] as string[],
  status: '待通知'
})

const formStatus = computed(() => {
  if (submitting.value) {
    return { text: '提交中', color: 'processing', icon: ReloadOutlined }
  }
  if (isFormValid.value) {
    return { text: '可提交', color: 'success', icon: CheckCircleOutlined }
  }
  return { text: '待填写', color: 'warning', icon: WarningOutlined }
})

const isFormValid = computed(() => {
  return (
    form.recallID.trim() !== '' &&
    form.batchNos.length > 0 &&
    form.batchNos.every(b => b.trim() !== '') &&
    form.reason.trim() !== ''
  )
})

const formRules = {
  recallID: [{ required: true, message: '请输入召回ID', trigger: 'blur' }],
  reason: [{ required: true, message: '请输入召回原因', trigger: 'blur' }]
}

function getRecallTypeColor(type: string) {
  const colors: Record<string, string> = {
    '安全召回': 'red',
    '质量召回': 'orange',
    '环保召回': 'green',
    '其他': 'default'
  }
  return colors[type] || 'default'
}

function getPriorityColor(priority: string) {
  const colors: Record<string, string> = {
    '紧急': 'red',
    '高': 'orange',
    '中': 'blue',
    '低': 'green'
  }
  return colors[priority] || 'default'
}

function goBack() {
  router.push('/aftersale/recall')
}

function generateRecallID() {
  const timestamp = Date.now().toString().slice(-8)
  form.recallID = `RECALL-${timestamp}`
  message.success('已生成召回ID')
}

function fillTestData() {
  form.recallID = `RECALL-${Date.now().toString().slice(-8)}`
  form.batchNos = ['BATCH-' + Date.now().toString().slice(-6), 'BATCH-' + (Date.now() + 1).toString().slice(-6)]
  form.affectedParts = ['PART-' + Date.now().toString().slice(-6)]
  form.reason = '测试召回原因：发现零部件存在安全隐患，可能影响车辆正常使用。'
  message.success('已填充测试数据')
}

async function handleSubmit() {
  if (submitting.value) return
  submitting.value = true

  try {
    await createRecallRecord({
      recallID: form.recallID,
      batchNos: form.batchNos,
      reason: form.reason,
      affectedParts: form.affectedParts,
      status: form.status
    })
    createdRecallID.value = form.recallID
    showSuccessModal.value = true
    message.success('召回发起成功')
  } catch (error: any) {
    message.error(error.message || '发起失败')
  } finally {
    submitting.value = false
  }
}

function handleReset() {
  formRef.value?.resetFields()
  form.recallID = ''
  form.batchNos = []
  form.affectedParts = []
  form.reason = ''
  form.status = '待通知'
}

function handleCreateAnother() {
  showSuccessModal.value = false
  createdRecallID.value = ''
  handleReset()
  generateRecallID()
}
</script>

<style scoped>
.recall-create {
  min-height: 100vh;
  background: var(--bg-color);
  padding: 0;
  display: flex;
  flex-direction: column;
}

.recall-create :deep(.ant-select-dropdown) {
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
  background: linear-gradient(135deg, #faad14 0%, #ffc53d 100%);
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

/* ==================== 增强输入框样式 ==================== */
.enhanced-form-item {
  position: relative;
  margin-bottom: 24px;
}

.enhanced-form-item :deep(.ant-form-item-label > label) {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-color);
  margin-bottom: 8px;
  display: flex;
  align-items: center;
  gap: 6px;
}

.enhanced-form-item :deep(.ant-form-item-label > label::after) {
  content: '';
  display: none;
}

.enhanced-form-item :deep(.ant-form-item-explain) {
  margin-top: 8px;
  font-size: 12px;
  line-height: 1.5;
  padding: 8px 12px;
  background: rgba(250, 84, 28, 0.08);
  border-radius: 6px;
  border-left: 3px solid #fa541c;
  color: #fa541c;
}

.enhanced-form-item :deep(.ant-form-item-has-error .ant-form-item-explain) {
  background: rgba(250, 84, 28, 0.08);
  border-left-color: #ff4d4f;
  color: #ff4d4f;
}

.enhanced-form-item :deep(.ant-form-item-has-warning .ant-form-item-explain) {
  background: rgba(250, 173, 20, 0.08);
  border-left-color: #faad14;
  color: #faad14;
}

.enhanced-form-item :deep(.ant-form-item-has-success .ant-form-item-explain) {
  background: rgba(82, 196, 26, 0.08);
  border-left-color: #52c41a;
  color: #52c41a;
}

.enhanced-select {
  position: relative;
}

.enhanced-select :deep(.ant-select-selector) {
  height: auto !important;
  min-height: 48px !important;
  border-radius: 12px !important;
  border: 2px solid var(--border-color) !important;
  background: linear-gradient(135deg, var(--bg-color-secondary) 0%, var(--bg-color-tertiary) 100%) !important;
  transition: all 0.35s cubic-bezier(0.4, 0, 0.2, 1) !important;
  font-size: 14px;
  padding: 6px 16px !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.04);
  overflow: hidden !important;
}

.enhanced-select :deep(.ant-select-selection-overflow) {
  display: flex !important;
  flex-wrap: wrap !important;
  align-items: center !important;
  gap: 4px !important;
}

.enhanced-select :deep(.ant-select-selection-overflow-item) {
  flex-shrink: 0 !important;
}

.enhanced-select :deep(.ant-select-selection-overflow-item-suffix) {
  flex: 1 !important;
  min-width: 100px !important;
}

.enhanced-select :deep(.ant-select-selector:hover) {
  border-color: #6366f1 !important;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.2) !important;
  transform: translateY(-1px);
}

.enhanced-select :deep(.ant-select-focused .ant-select-selector) {
  border-color: #6366f1 !important;
  box-shadow: 0 0 0 4px rgba(99, 102, 241, 0.15) !important;
}

.enhanced-select :deep(.ant-select-selection-search) {
  height: 100% !important;
  display: flex !important;
  align-items: center !important;
}

.enhanced-select :deep(.ant-select-selection-search-input) {
  height: 44px !important;
  font-size: 14px;
  color: var(--text-color);
  background: transparent;
}

.enhanced-select :deep(.ant-select-selection-search-input::placeholder) {
  color: var(--text-color-tertiary);
  font-size: 13px;
}

.enhanced-select :deep(.ant-select-selection-placeholder) {
  height: 100% !important;
  display: flex !important;
  align-items: center !important;
  color: var(--text-color-tertiary) !important;
  font-size: 13px;
}

.enhanced-select :deep(.ant-select-selection-item) {
  height: 32px;
  line-height: 32px !important;
  display: inline-flex !important;
  align-items: center !important;
  padding: 0 10px;
  margin: 4px 6px 4px 0;
  background: linear-gradient(135deg, #e0e7ff 0%, #c7d2fe 100%);
  border: 1px solid #a78bfa;
  border-radius: 8px;
  color: #5b21b6;
  font-size: 13px;
  font-weight: 500;
  transition: all 0.3s ease;
  max-width: 100%;
  flex-shrink: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.enhanced-select :deep(.ant-select-selection-item-content) {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  max-width: 200px;
}

.enhanced-select :deep(.ant-select-selection-item:hover) {
  background: linear-gradient(135deg, #c7d2fe 0%, #a78bfa 100%);
  border-color: #8b5cf6;
  color: #4c1d95;
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(139, 92, 246, 0.3);
}

.enhanced-select :deep(.ant-select-selection-item-remove) {
  color: #7c3aed;
  margin-left: 6px;
  font-size: 14px;
  transition: all 0.2s ease;
}

.enhanced-select :deep(.ant-select-selection-item-remove:hover) {
  color: #dc2626;
  transform: scale(1.2);
}

.enhanced-select :deep(.ant-select-arrow) {
  display: none;
}

.select-suffix-icon {
  font-size: 16px;
  color: var(--text-color-tertiary);
  transition: all 0.3s ease;
}

.enhanced-select:hover .select-suffix-icon {
  color: #6366f1;
  transform: scale(1.1);
}

.enhanced-select :deep(.ant-select-focused) .select-suffix-icon {
  color: #6366f1;
  transform: scale(1.1);
}

/* 下拉菜单增强样式 */
.enhanced-select :deep(.ant-select-dropdown) {
  z-index: 1050 !important;
  border-radius: 12px !important;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15) !important;
  border: 1px solid var(--border-color);
  padding: 8px;
  background: var(--bg-color-secondary);
  animation: slideDown 0.2s ease-out;
}

@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.enhanced-select :deep(.ant-select-item) {
  padding: 12px 16px !important;
  border-radius: 8px !important;
  margin: 2px 0 !important;
  transition: all 0.2s ease !important;
  font-size: 14px;
  color: var(--text-color);
  display: flex !important;
  align-items: center !important;
}

.enhanced-select :deep(.ant-select-item:hover) {
  background: linear-gradient(135deg, #f5f3ff 0%, #ede9fe 100%) !important;
  color: #6366f1 !important;
  transform: translateX(4px);
}

.enhanced-select :deep(.ant-select-item-option-selected) {
  background: linear-gradient(135deg, #ede9fe 0%, #ddd6fe 100%) !important;
  color: #6366f1 !important;
  font-weight: 600 !important;
}

.enhanced-select :deep(.ant-select-item-option-active) {
  background: linear-gradient(135deg, #f5f3ff 0%, #ede9fe 100%) !important;
  color: #6366f1 !important;
}

/* 错误状态样式 */
.enhanced-form-item :deep(.ant-form-item-has-error .enhanced-select .ant-select-selector) {
  border-color: #ff4d4f !important;
  box-shadow: 0 0 0 3px rgba(255, 77, 79, 0.1) !important;
  animation: shake 0.3s ease-in-out;
}

@keyframes shake {
  0%, 100% { transform: translateX(0); }
  25% { transform: translateX(-4px); }
  75% { transform: translateX(4px); }
}

/* 警告状态样式 */
.enhanced-form-item :deep(.ant-form-item-has-warning .enhanced-select .ant-select-selector) {
  border-color: #faad14 !important;
  box-shadow: 0 0 0 3px rgba(250, 173, 20, 0.1) !important;
}

/* 成功状态样式 */
.enhanced-form-item :deep(.ant-form-item-has-success .enhanced-select .ant-select-selector) {
  border-color: #52c41a !important;
  box-shadow: 0 0 0 3px rgba(82, 196, 26, 0.1) !important;
}

/* 输入框聚焦时的标签动画 */
.enhanced-form-item :deep(.ant-form-item-label) {
  transition: all 0.3s ease;
}

.enhanced-form-item :deep(.ant-form-item-has-feedback .ant-form-item-label > label) {
  color: var(--primary-color);
}

/* 深色主题适配 */
[data-theme='dark'] .enhanced-select :deep(.ant-select-selector) {
  background: linear-gradient(135deg, #1f2937 0%, #111827 100%) !important;
  border-color: #4b5563 !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  height: auto !important;
  min-height: 48px !important;
  padding: 6px 16px !important;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-selector:hover) {
  border-color: #818cf8 !important;
  box-shadow: 0 4px 12px rgba(129, 140, 248, 0.3) !important;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-focused .ant-select-selector) {
  border-color: #818cf8 !important;
  box-shadow: 0 0 0 4px rgba(129, 140, 248, 0.2) !important;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-selection-search-input) {
  color: #e5e7eb;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-selection-search-input::placeholder) {
  color: #6b7280;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-selection-placeholder) {
  color: #6b7280 !important;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-selection-item) {
  background: linear-gradient(135deg, #4c1d95 0%, #5b21b6 100%);
  border-color: #7c3aed;
  color: #e9d5ff;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-selection-item:hover) {
  background: linear-gradient(135deg, #5b21b6 0%, #6d28d9 100%);
  border-color: #8b5cf6;
  color: #fff;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-selection-item-remove) {
  color: #c4b5fd;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-selection-item-remove:hover) {
  color: #f87171;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-dropdown) {
  background: #1f2937;
  border-color: #374151;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.4);
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-item) {
  color: #e5e7eb;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-item:hover) {
  background: linear-gradient(135deg, rgba(129, 140, 248, 0.2) 0%, rgba(139, 92, 246, 0.2) 100%) !important;
  color: #a5b4fc !important;
}

[data-theme='dark'] .enhanced-select :deep(.ant-select-item-option-selected) {
  background: linear-gradient(135deg, rgba(129, 140, 248, 0.3) 0%, rgba(139, 92, 246, 0.3) 100%) !important;
  color: #a5b4fc !important;
}

[data-theme='dark'] .enhanced-form-item :deep(.ant-form-item-has-error .enhanced-select .ant-select-selector) {
  border-color: #f87171 !important;
  box-shadow: 0 0 0 3px rgba(248, 113, 113, 0.2) !important;
}

[data-theme='dark'] .enhanced-form-item :deep(.ant-form-item-has-warning .enhanced-select .ant-select-selector) {
  border-color: #fbbf24 !important;
  box-shadow: 0 0 0 3px rgba(251, 191, 36, 0.2) !important;
}

[data-theme='dark'] .enhanced-form-item :deep(.ant-form-item-has-success .enhanced-select .ant-select-selector) {
  border-color: #34d399 !important;
  box-shadow: 0 0 0 3px rgba(52, 211, 153, 0.2) !important;
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

.success-recall-id {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  background: var(--bg-color-tertiary);
  border-radius: 10px;
  margin-bottom: 24px;
}

.success-recall-id .label {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.success-recall-id .value {
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
