<!--
  BOM创建页面
  设计：与Dashboard风格统一
-->

<template>
  <div class="bom-create">
    <!-- 顶部导航栏 -->
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">创建BOM</h1>
          <span class="page-subtitle">填写物料清单信息</span>
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
              <FileAddOutlined />
            </div>
            <div class="header-text">
              <h3>BOM信息</h3>
              <p>请填写以下信息创建新的物料清单</p>
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
            <!-- BOM ID -->
            <a-form-item label="BOM ID" name="bomID" class="form-item">
              <a-input
                v-model:value="form.bomID"
                placeholder="请输入BOM ID"
                :maxlength="50"
                class="custom-input"
              >
                <template #prefix>
                  <span class="input-icon"><NumberOutlined /></span>
                </template>
              </a-input>
            </a-form-item>

            <!-- BOM类型和车型 -->
            <div class="form-row">
              <a-form-item label="BOM类型" name="bomType" class="form-item half">
                <a-select
                  v-model:value="form.bomType"
                  placeholder="请选择类型"
                  :options="bomTypeOptions"
                  class="custom-select"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                  :dropdown-style="{ zIndex: 1050 }"
                />
              </a-form-item>
              <a-form-item label="车型" name="productModel" class="form-item half">
                <a-input
                  v-model:value="form.productModel"
                  placeholder="请输入车型"
                  :maxlength="50"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><CarOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
            </div>

            <!-- 版本号和状态 -->
            <div class="form-row">
              <a-form-item label="版本号" name="version" class="form-item half">
                <a-input
                  v-model:value="form.version"
                  placeholder="例如: v1.0"
                  :maxlength="20"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><TagOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
              <a-form-item label="状态" name="status" class="form-item half">
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
                {{ submitting ? '提交中...' : '提交创建' }}
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
              <span>BOM ID需要唯一标识</span>
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
              <span>草稿状态可继续编辑</span>
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
              <span class="preview-label">BOM ID</span>
              <span class="preview-value">{{ form.bomID || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">类型</span>
              <span class="preview-value">{{ form.bomType || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">车型</span>
              <span class="preview-value">{{ form.productModel || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">版本</span>
              <span class="preview-value">{{ form.version || '-' }}</span>
            </div>
            <div class="preview-item">
              <span class="preview-label">状态</span>
              <a-tag :color="getStatusColor(form.status)" size="small">
                {{ getStatusText(form.status) }}
              </a-tag>
            </div>
          </div>
        </a-card>
      </div>
    </div>
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
  FileAddOutlined,
  NumberOutlined,
  TagOutlined,
  CarOutlined,
  ClockCircleOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const formRef = ref<FormInstance>()
const submitting = ref(false)

const form = reactive({
  bomID: '',
  bomType: '',
  productModel: '',
  version: '',
  status: '草稿',
  remark: ''
})

const formStatus = computed(() => {
  if (submitting.value) {
    return { text: '提交中', color: 'processing', icon: ReloadOutlined }
  }
  if (isFormValid.value) {
    return { text: '可提交', color: 'success', icon: CheckCircleOutlined }
  }
  return { text: '待填写', color: 'warning', icon: ClockCircleOutlined }
})

const bomTypeOptions = computed<SelectProps['options']>(() => [
  { value: '生产BOM', label: '生产BOM' },
  { value: '研发BOM', label: '研发BOM' },
  { value: '工程BOM', label: '工程BOM' }
])

const statusOptions = computed<SelectProps['options']>(() => [
  { value: '草稿', label: '草稿' },
  { value: '已发布', label: '已发布' }
])

const isFormValid = computed(() => {
  return (
    form.bomID.trim() !== '' &&
    form.bomType !== '' &&
    form.productModel.trim() !== '' &&
    form.version.trim() !== '' &&
    form.status !== ''
  )
})

const formRules = {
  bomID: [{ required: true, message: '请输入BOM ID', trigger: 'blur' }],
  bomType: [{ required: true, message: '请选择BOM类型', trigger: 'change' }],
  productModel: [{ required: true, message: '请输入车型', trigger: 'blur' }],
  version: [{ required: true, message: '请输入版本号', trigger: 'blur' }],
  status: [{ required: true, message: '请选择状态', trigger: 'change' }]
}

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    '草稿': 'orange',
    '已发布': 'green'
  }
  return colors[status] || 'default'
}

function getStatusText(status: string) {
  return status || '草稿'
}

function goBack() {
  router.push('/bom/list')
}

async function handleSubmit() {
  if (submitting.value) return
  submitting.value = true
  
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('BOM创建成功')
    router.push('/bom/list')
  } catch (error: any) {
    message.error(error.message || '创建失败')
  } finally {
    submitting.value = false
  }
}

function handleReset() {
  formRef.value?.resetFields()
  form.status = '草稿'
  message.success('表单已重置')
}
</script>

<style scoped>
.bom-create {
  min-height: 100vh;
  background: var(--bg-color);
  padding: 0;
  display: flex;
  flex-direction: column;
}

.bom-create :deep(.ant-select-dropdown) {
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

[data-theme='dark'] .custom-input :deep(.ant-input-affix-wrapper),
[data-theme='dark'] .custom-select :deep(.ant-select-selector),
[data-theme='dark'] .custom-textarea :deep(textarea) {
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

[data-theme='dark'] .preview-item {
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
</style>
