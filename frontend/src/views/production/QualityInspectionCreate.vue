<!--
  质检数据录入页面
  设计：与BOMCreate风格统一
-->

<template>
  <div class="quality-create">
    <!-- 顶部导航栏 -->
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">录入质检数据</h1>
          <span class="page-subtitle">填写质量检验信息</span>
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
            <div class="header-icon quality">
              <SafetyCertificateOutlined />
            </div>
            <div class="header-text">
              <h3>质检信息</h3>
              <p>请填写以下信息录入质检数据</p>
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
            <!-- 质检ID -->
            <a-form-item label="质检ID" name="inspectionID" class="form-item">
              <a-input
                v-model:value="form.inspectionID"
                placeholder="自动生成或手动输入"
                :maxlength="50"
                class="custom-input"
              >
                <template #prefix>
                  <span class="input-icon"><NumberOutlined /></span>
                </template>
              </a-input>
            </a-form-item>

            <!-- 零部件ID和批次号 -->
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
              <a-form-item label="批次号" name="batchNo" class="form-item half">
                <a-input
                  v-model:value="form.batchNo"
                  placeholder="请输入批次号"
                  :maxlength="30"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><TagOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
            </div>

            <!-- 质检结果和处理人 -->
            <div class="form-row">
              <a-form-item label="质检结果" name="result" class="form-item half">
                <a-select
                  v-model:value="form.result"
                  placeholder="请选择质检结果"
                  :options="resultOptions"
                  class="custom-select"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                  :dropdown-style="{ zIndex: 1050 }"
                />
              </a-form-item>
              <a-form-item label="处理人" name="handler" class="form-item half">
                <a-input
                  v-model:value="form.handler"
                  placeholder="请输入处理人姓名"
                  :maxlength="20"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><UserOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
            </div>

            <!-- 检验类型和检验项目 -->
            <div class="form-row">
              <a-form-item label="检验类型" name="inspectionType" class="form-item half">
                <a-select
                  v-model:value="form.inspectionType"
                  placeholder="请选择检验类型"
                  :options="inspectionTypeOptions"
                  class="custom-select"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                  :dropdown-style="{ zIndex: 1050 }"
                />
              </a-form-item>
              <a-form-item label="检验项目" name="inspectionItem" class="form-item half">
                <a-input
                  v-model:value="form.inspectionItem"
                  placeholder="请输入检验项目"
                  :maxlength="50"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><FileSearchOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
            </div>

            <!-- 不良数量 -->
            <a-form-item label="不良数量" name="defectCount" class="form-item">
              <a-input-number
                v-model:value="form.defectCount"
                placeholder="请输入不良数量"
                :min="0"
                :max="99999"
                class="custom-input-number"
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
                {{ submitting ? '提交中...' : '提交录入' }}
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
            <span>录入说明</span>
          </div>
          <ul class="guide-list">
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>质检ID可自动生成或手动输入</span>
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
              <span>不合格品需进行后续处理</span>
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
              <span class="preview-label">质检ID</span>
              <span class="preview-value">{{ form.inspectionID || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">零部件ID</span>
              <span class="preview-value">{{ form.partID || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">批次号</span>
              <span class="preview-value">{{ form.batchNo || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">质检结果</span>
              <a-tag :color="getResultColor(form.result)" size="small">
                {{ form.result || '-' }}
              </a-tag>
            </div>
            <div class="preview-item">
              <span class="preview-label">处理人</span>
              <span class="preview-value">{{ form.handler || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">检验类型</span>
              <span class="preview-value">{{ form.inspectionType || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">不良数量</span>
              <span class="preview-value">{{ form.defectCount || 0 }} 件</span>
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
            <div class="quick-action-item" @click="generateInspectionID">
              <NumberOutlined class="action-icon" />
              <span>生成质检ID</span>
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
        <div class="success-icon-wrapper" :class="form.result === '合格' ? 'success' : 'warning'">
          <CheckCircleFilled v-if="form.result === '合格'" class="success-icon" />
          <WarningFilled v-else class="warning-icon" />
        </div>
        <h2 class="success-title">录入成功</h2>
        <p class="success-desc">质检数据已成功提交到区块链</p>
        <div class="success-part-id">
          <span class="label">质检ID：</span>
          <span class="value">{{ createdInspectionID }}</span>
        </div>
        <div class="success-result">
          <span class="label">质检结果：</span>
          <a-tag :color="getResultColor(form.result)" size="small">
            {{ form.result }}
          </a-tag>
        </div>
        <div class="success-actions">
          <a-button type="primary" @click="handleCreateAnother">
            <template #icon><PlusOutlined /></template>
            继续录入
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
import { createQualityInspectionWithChain, type BackendQualityDTO } from '@/services/quality.service'
import {
  ArrowLeftOutlined,
  CheckOutlined,
  ReloadOutlined,
  InfoCircleOutlined,
  CheckCircleOutlined,
  EyeOutlined,
  SafetyCertificateOutlined,
  NumberOutlined,
  AppstoreOutlined,
  TagOutlined,
  UserOutlined,
  FileSearchOutlined,
  ThunderboltOutlined,
  PlusOutlined,
  UnorderedListOutlined,
  CheckCircleFilled,
  WarningFilled
} from '@ant-design/icons-vue'

const router = useRouter()
const formRef = ref<FormInstance>()
const submitting = ref(false)
const showSuccessModal = ref(false)
const createdInspectionID = ref('')

const form = reactive({
  inspectionID: '',
  partID: '',
  batchNo: '',
  result: '',
  handler: '',
  inspectionType: '',
  inspectionItem: '',
  defectCount: 0,
  remark: ''
})

const formStatus = computed(() => {
  if (submitting.value) {
    return { text: '提交中', color: 'processing', icon: ReloadOutlined }
  }
  if (isFormValid.value) {
    return { text: '可提交', color: 'success', icon: CheckCircleOutlined }
  }
  return { text: '待填写', color: 'warning', icon: SafetyCertificateOutlined }
})

const resultOptions = computed<SelectProps['options']>(() => [
  { value: '合格', label: '合格' },
  { value: '不合格', label: '不合格' }
])

const inspectionTypeOptions = computed<SelectProps['options']>(() => [
  { value: '来料检验', label: '来料检验' },
  { value: '过程检验', label: '过程检验' },
  { value: '成品检验', label: '成品检验' },
  { value: '出货检验', label: '出货检验' }
])

const isFormValid = computed(() => {
  return (
    form.partID.trim() !== '' &&
    form.batchNo.trim() !== '' &&
    form.result !== '' &&
    form.handler.trim() !== '' &&
    form.inspectionType !== ''
  )
})

const formRules = {
  partID: [{ required: true, message: '请输入零部件ID', trigger: 'blur' }],
  batchNo: [{ required: true, message: '请输入批次号', trigger: 'blur' }],
  result: [{ required: true, message: '请选择质检结果', trigger: 'change' }],
  handler: [{ required: true, message: '请输入处理人姓名', trigger: 'blur' }],
  inspectionType: [{ required: true, message: '请选择检验类型', trigger: 'change' }]
}

function getResultColor(result: string) {
  if (result === '合格') return 'green'
  if (result === '不合格') return 'red'
  return 'default'
}

function goBack() {
  router.push('/production/quality')
}

function generateInspectionID() {
  const timestamp = Date.now().toString().slice(-8)
  form.inspectionID = `QC-${timestamp}`
  message.success('已生成质检ID')
}

function fillTestData() {
  form.partID = 'PART-' + Date.now().toString().slice(-6)
  form.batchNo = 'BATCH-' + Date.now().toString().slice(-6)
  form.result = '合格'
  form.handler = '测试质检员'
  form.inspectionType = '成品检验'
  form.inspectionItem = '外观检验'
  form.defectCount = 0
  message.success('已填充测试数据')
}

async function handleSubmit() {
  if (submitting.value) return
  submitting.value = true

  try {
    const id = form.inspectionID || `QC-${Date.now().toString().slice(-8)}`
    const payload: BackendQualityDTO = {
      inspectionID: id,
      partID: form.partID.trim(),
      batchNo: form.batchNo.trim(),
      indicators: {},
      result: form.result,
      handler: form.handler.trim(),
      handleTime: Math.floor(Date.now() / 1000).toString(),
      disposal: form.result === '不合格' ? '' : ''
    }
    await createQualityInspectionWithChain(payload)
    createdInspectionID.value = id
    showSuccessModal.value = true
    message.success('质检数据录入成功')
  } catch (error: any) {
    message.error(error.message || '录入失败')
  } finally {
    submitting.value = false
  }
}

function handleReset() {
  formRef.value?.resetFields()
  form.inspectionID = ''
  form.partID = ''
  form.batchNo = ''
  form.result = ''
  form.handler = ''
  form.inspectionType = ''
  form.inspectionItem = ''
  form.defectCount = 0
  form.remark = ''
}

function handleCreateAnother() {
  showSuccessModal.value = false
  createdInspectionID.value = ''
  handleReset()
  generateInspectionID()
}
</script>

<style scoped>
.quality-create {
  min-height: 100vh;
  background: var(--bg-color);
  padding: 0;
  display: flex;
  flex-direction: column;
}

.quality-create :deep(.ant-select-dropdown) {
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

.header-icon.quality {
  background: linear-gradient(135deg, #52c41a 0%, #389e0d 100%);
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

.success-icon-wrapper.warning {
  background: linear-gradient(135deg, #faad14 0%, #d48806 100%);
}

.success-icon {
  font-size: 36px;
  color: #fff;
}

.warning-icon {
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
  margin-bottom: 24px;
}

.success-result .label {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.success-actions {
  display: flex;
  gap: 12px;
  justify-content: center;
}

.success-actions .ant-btn {
  min-width: 100px;
  height: 40px;
  border-radius: 8px;
}

/* 深色主题 */
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

[data-theme='dark'] .header-icon.quality {
  background: linear-gradient(135deg, #49aa19 0%, #237804 100%);
}

[data-theme='dark'] .custom-input :deep(.ant-input-affix-wrapper),
[data-theme='dark'] .custom-select :deep(.ant-select-selector),
[data-theme='dark'] .custom-textarea :deep(textarea),
[data-theme='dark'] .custom-input-number :deep(.ant-input-number) {
  background: var(--bg-color);
  border-color: var(--border-color);
}

[data-theme='dark'] .custom-input :deep(.ant-input-affix-wrapper:hover),
[data-theme='dark'] .custom-select :deep(.ant-select-selector:hover),
[data-theme='dark'] .custom-textarea :deep(textarea:hover) {
  border-color: #177ddc;
}

[data-theme='dark'] .custom-input :deep(.ant-input-affix-wrapper:focus),
[data-theme='dark'] .custom-input :deep(.ant-input-affix-wrapper-focused),
[data-theme='dark'] .custom-select :deep(.ant-select-focused .ant-select-selector),
[data-theme='dark'] .custom-textarea :deep(textarea:focus) {
  border-color: #177ddc;
  box-shadow: 0 0 0 3px rgba(23, 125, 220, 0.15);
}

[data-theme='dark'] .custom-input :deep(.ant-input-affix-wrapper:focus-within) .input-icon {
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

[data-theme='dark'] .custom-select :deep(.ant-select-item:hover) {
  background: rgba(129, 140, 248, 0.15);
  color: #a5b4fc;
}

[data-theme='dark'] .custom-select :deep(.ant-select-item-option-selected) {
  background: rgba(129, 140, 248, 0.2);
  color: #a5b4fc;
}

/* 响应式 */
@media (max-width: 1024px) {
  .main-content {
    flex-direction: column;
    padding: 16px;
  }

  .info-container {
    width: 100%;
  }
}

@media (max-width: 640px) {
  .top-bar {
    padding: 12px 16px;
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }

  .top-bar-right {
    width: 100%;
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
