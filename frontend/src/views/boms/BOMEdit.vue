<!--
  BOM编辑页面
  设计：与BOM创建页面和详情页面风格统一
-->

<template>
  <div class="bom-edit">
    <!-- 顶部导航栏 -->
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回详情
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">编辑BOM</h1>
          <span class="page-subtitle">修改物料清单信息</span>
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
              <EditOutlined />
            </div>
            <div class="header-text">
              <h3>BOM信息</h3>
              <p>请修改以下信息更新物料清单</p>
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
            class="edit-form"
          >
            <!-- BOM ID（只读） -->
            <a-form-item label="BOM ID" name="bomID" class="form-item">
              <a-input
                v-model:value="form.bomID"
                placeholder="BOM ID"
                :maxlength="50"
                class="custom-input"
                disabled
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

            <!-- 版本号和零部件批次号 -->
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
              <a-form-item label="零部件批次号" name="partBatchNo" class="form-item half">
                <a-input
                  v-model:value="form.partBatchNo"
                  placeholder="请输入零部件批次号"
                  :maxlength="50"
                  class="custom-input"
                >
                  <template #prefix>
                    <span class="input-icon"><BarcodeOutlined /></span>
                  </template>
                </a-input>
              </a-form-item>
            </div>

            <!-- 状态 -->
            <a-form-item label="状态" name="status" class="form-item">
              <a-select
                v-model:value="form.status"
                placeholder="请选择状态"
                :options="statusOptions"
                class="custom-select"
                :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                :dropdown-style="{ zIndex: 1050 }"
              />
            </a-form-item>

            <!-- 物料列表 -->
            <div class="form-item">
              <label class="ant-form-item-required">物料列表</label>
              <div class="material-list">
                <div v-for="(material, index) in form.materialList" :key="index" class="material-item">
                  <div class="material-header">
                    <span class="material-index">物料 {{ index + 1 }}</span>
                    <a-button
                      type="text"
                      danger
                      size="small"
                      @click="removeMaterial(index)"
                      v-if="form.materialList.length > 1"
                    >
                      <template #icon><MinusCircleOutlined /></template>
                      删除
                    </a-button>
                  </div>
                  <div class="material-fields">
                    <a-input
                      v-model:value="material.materialID"
                      placeholder="物料ID"
                      class="material-input"
                    />
                    <a-input
                      v-model:value="material.materialName"
                      placeholder="物料名称"
                      class="material-input"
                    />
                    <a-input
                      v-model:value="material.quantity"
                      placeholder="物料数量"
                      class="material-input"
                    />
                  </div>
                </div>
                <a-button type="dashed" @click="addMaterial" class="add-material-btn">
                  <template #icon><PlusOutlined /></template>
                  添加物料
                </a-button>
              </div>
            </div>

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
                {{ submitting ? '保存中...' : '保存修改' }}
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
        <!-- 编辑说明卡片 -->
        <a-card :bordered="false" class="info-card guide-card">
          <div class="card-title">
            <InfoCircleOutlined />
            <span>编辑说明</span>
          </div>
          <ul class="guide-list">
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>BOM ID不可修改</span>
            </li>
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>带 <em>*</em> 的为必填项</span>
            </li>
            <li>
              <CheckCircleOutlined class="guide-icon success" />
              <span>修改后数据将更新到区块链</span>
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
            <div class="preview-item">
              <span class="preview-label">物料数量</span>
              <span class="preview-value">{{ form.materialList.length }} 项</span>
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
            <div class="quick-action-item" @click="resetToOriginal">
              <ReloadOutlined class="action-icon" />
              <span>恢复原始数据</span>
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
        <h2 class="success-title">修改成功</h2>
        <p class="success-desc">BOM数据已成功更新到区块链</p>
        <div class="success-part-id">
          <span class="label">BOM ID：</span>
          <span class="value">{{ form.bomID }}</span>
        </div>
        <div class="success-actions">
          <a-button type="primary" @click="goBack">
            <template #icon><EyeOutlined /></template>
            查看详情
          </a-button>
        </div>
      </div>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { message } from 'ant-design-vue'
import type { FormInstance, SelectProps } from 'ant-design-vue'
import { getBOM, updateBOM } from '../../services/bom.service'
import type { BOMDTO, BOMMaterial } from '../../services/bom.service'
import {
  ArrowLeftOutlined,
  CheckOutlined,
  ReloadOutlined,
  InfoCircleOutlined,
  CheckCircleOutlined,
  EyeOutlined,
  EditOutlined,
  NumberOutlined,
  TagOutlined,
  CarOutlined,
  ClockCircleOutlined,
  ThunderboltOutlined,
  PlusOutlined,
  MinusCircleOutlined,
  CheckCircleFilled,
  BarcodeOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()
const bomID = route.params.id as string

const formRef = ref<FormInstance>()
const loading = ref(false)
const submitting = ref(false)
const showSuccessModal = ref(false)
const originalData = ref<BOMDTO | null>(null)

const form = reactive<BOMDTO>({
  bomID: '',
  bomType: '',
  productModel: '',
  partBatchNo: '',
  version: '',
  creator: '',
  createTime: '',
  status: 'draft',
  materialList: [
    {
      materialID: '',
      materialName: '',
      quantity: 0,
      unit: '个'
    }
  ]
})

const formStatus = computed(() => {
  if (submitting.value) {
    return { color: 'processing', text: '保存中...', icon: 'LoadingOutlined' }
  }
  return { color: 'default', text: '编辑中', icon: 'EditOutlined' }
})

const isFormValid = computed(() => {
  return form.bomID && 
         form.bomType && 
         form.productModel && 
         form.partBatchNo && 
         form.version && 
         form.status &&
         form.materialList.length > 0 &&
         form.materialList.every(m => m.materialID && m.materialName && m.quantity > 0)
})

const bomTypeOptions: SelectProps['options'] = [
  { label: '研发BOM', value: '研发BOM' },
  { label: '生产BOM', value: '生产BOM' }
]

const statusOptions: SelectProps['options'] = [
  { label: '草稿', value: 'draft' },
  { label: '已发布', value: 'published' },
  { label: '已归档', value: 'archived' }
]

const formRules = {
  bomType: [{ required: true, message: '请选择BOM类型', trigger: 'change' }],
  productModel: [{ required: true, message: '请输入车型', trigger: 'blur' }],
  partBatchNo: [{ required: true, message: '请输入零部件批次号', trigger: 'blur' }],
  version: [{ required: true, message: '请输入版本号', trigger: 'blur' }],
  status: [{ required: true, message: '请选择状态', trigger: 'change' }]
}

const getStatusColor = (status: string) => {
  const colorMap: Record<string, string> = {
    draft: 'default',
    published: 'success',
    archived: 'default'
  }
  return colorMap[status] || 'default'
}

const getStatusText = (status: string) => {
  const textMap: Record<string, string> = {
    draft: '草稿',
    published: '已发布',
    archived: '已归档'
  }
  return textMap[status] || status
}

const loadBOMData = async () => {
  loading.value = true
  try {
    const response = await getBOM(bomID)
    if (response.code === 0 && response.data) {
      const data = response.data
      
      const statusMap: Record<string, string> = {
        '草稿': 'draft',
        '已发布': 'published',
        '已归档': 'archived'
      }
      
      form.bomID = data.bomID
      form.bomType = data.bomType
      form.productModel = data.productModel
      form.partBatchNo = data.partBatchNo
      form.version = data.version
      form.creator = data.creator
      form.createTime = data.createTime
      form.status = statusMap[data.status] || data.status
      form.materialList = data.materialList && data.materialList.length > 0 
        ? data.materialList.map((m: any) => ({
            materialID: m.materialID,
            materialName: m.materialName || m.materialID,
            quantity: Number(m.quantity) || 0,
            unit: m.unit || '个',
            specifications: m.spec || m.specifications || ''
          }))
        : [
            {
              materialID: '',
              materialName: '',
              quantity: 0,
              unit: '个'
            }
          ]
      originalData.value = JSON.parse(JSON.stringify(form))
    }
  } catch (error: any) {
    message.error(error.message || '加载BOM数据失败')
  } finally {
    loading.value = false
  }
}

const addMaterial = () => {
  form.materialList.push({
    materialID: '',
    materialName: '',
    quantity: 0,
    unit: '个'
  })
}

const removeMaterial = (index: number) => {
  if (form.materialList.length > 1) {
    form.materialList.splice(index, 1)
  }
}

const fillTestData = () => {
  form.bomType = '研发BOM'
  form.productModel = 'MODEL-A-2026'
  form.partBatchNo = 'BATCH-MA-001'
  form.version = 'V1.0'
  form.status = 'draft'
  form.materialList = [
    {
      materialID: 'PART-MA-001',
      materialName: '发动机总成',
      quantity: 1,
      unit: '个'
    },
    {
      materialID: 'PART-MA-003',
      materialName: '变速箱',
      quantity: 4,
      unit: '个'
    }
  ]
  message.success('测试数据已填充')
}

const resetToOriginal = () => {
  if (originalData.value) {
    Object.assign(form, JSON.parse(JSON.stringify(originalData.value)))
    message.success('已恢复原始数据')
  }
}

const handleReset = () => {
  resetToOriginal()
}

const handleSubmit = async () => {
  submitting.value = true
  try {
    const submitData: BOMDTO = {
      bomID: form.bomID,
      bomType: form.bomType,
      productModel: form.productModel,
      partBatchNo: form.partBatchNo,
      version: form.version,
      creator: form.creator,
      createTime: form.createTime,
      status: form.status,
      materialList: form.materialList.map(m => ({
        ...m,
        quantity: Number(m.quantity)
      }))
    }
    await updateBOM(bomID, submitData)
    showSuccessModal.value = true
  } catch (error: any) {
    message.error(error.message || '更新失败')
  } finally {
    submitting.value = false
  }
}

const goBack = () => {
  router.push(`/bom/detail/${bomID}`)
}

onMounted(() => {
  loadBOMData()
})
</script>

<style scoped>
.bom-edit {
  min-height: 100%;
  background: var(--bg-color);
  padding: 0;
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
}

.top-bar-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.back-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--text-color-secondary);
  font-size: 14px;
  padding: 8px 16px;
  border-radius: 8px;
  transition: all 0.3s;
}

.back-btn:hover {
  background: var(--bg-color-tertiary);
  color: var(--primary-color);
}

.page-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.page-title {
  margin: 0;
  font-size: 20px;
  font-weight: 600;
  color: var(--text-color);
}

.page-subtitle {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.divider {
  height: 24px;
  background: var(--border-color);
}

.top-bar-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.status-badge {
  font-size: 13px;
  padding: 4px 12px;
  border-radius: 16px;
  font-weight: 500;
}

.main-content {
  display: flex;
  gap: 24px;
  padding: 24px 32px;
  max-width: 1600px;
  margin: 0 auto;
}

.form-container {
  flex: 1;
  min-width: 0;
}

.form-card {
  border-radius: 16px;
  box-shadow: var(--shadow-sm);
  background: var(--bg-color-secondary);
}

.form-header {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 16px 0;
}

.header-icon {
  width: 44px;
  height: 44px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 20px;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.header-text h3 {
  margin: 0 0 2px 0;
  font-size: 16px;
  font-weight: 600;
  color: var(--text-color);
}

.header-text p {
  margin: 0;
  font-size: 13px;
  color: var(--text-color-secondary);
}

.form-divider {
  margin: 20px 0;
}

.edit-form {
  padding: 4px;
}

.form-item {
  margin-bottom: 20px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

.form-item.half {
  margin-bottom: 20px;
}

.custom-input {
  border-radius: 6px;
}

.custom-input :deep(.ant-input) {
  border-radius: 6px;
}

.custom-select {
  width: 100%;
  background: #fff;
}

.custom-select :deep(.ant-select-selector) {
  border-radius: 6px;
  background: #fff !important;
}

.input-icon {
  color: var(--text-color-tertiary);
  font-size: 14px;
}

.material-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  padding: 16px;
  background: var(--bg-color);
  border-radius: 12px;
  border: 1px dashed var(--border-color);
}

.material-item {
  background: var(--bg-color-secondary);
  border-radius: 12px;
  padding: 16px;
  border: 1px solid var(--border-color);
  transition: all 0.3s;
}

.material-item:hover {
  border-color: var(--primary-color);
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.1);
}

.material-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.material-index {
  font-size: 13px;
  font-weight: 500;
  color: var(--text-color);
}

.material-fields {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  gap: 8px;
}

.material-input {
  border-radius: 6px;
}

.add-material-btn {
  width: 100%;
  border-radius: 12px;
  border-style: dashed;
}

.form-actions {
  display: flex;
  gap: 12px;
  padding-top: 24px;
  border-top: 1px solid var(--border-color);
}

.submit-btn {
  flex: 1;
  height: 40px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
}

.secondary-btn {
  height: 40px;
  border-radius: 8px;
  font-size: 14px;
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
}

.card-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
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
  align-items: start;
  gap: 8px;
  margin-bottom: 12px;
  font-size: 13px;
  color: var(--text-color-secondary);
  line-height: 1.6;
}

.guide-list li:last-child {
  margin-bottom: 0;
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

.guide-list em {
  color: #ff4d4f;
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
  padding: 8px 0;
  border-bottom: 1px solid var(--border-color);
}

.preview-item:last-child {
  border-bottom: none;
}

.preview-label {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.preview-value {
  font-size: 13px;
  font-weight: 500;
  color: var(--text-color);
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
  padding: 16px;
  background: var(--bg-color);
  border-radius: 12px;
  border: 1px solid var(--border-color);
  cursor: pointer;
  transition: all 0.3s;
}

.quick-action-item:hover {
  border-color: var(--primary-color);
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.15);
  transform: translateY(-2px);
}

.action-icon {
  font-size: 20px;
  color: var(--primary-color);
}

.quick-action-item span {
  font-size: 13px;
  color: var(--text-color);
  font-weight: 500;
}

.success-modal :deep(.ant-modal-content) {
  border-radius: 12px;
  padding: 32px;
}

.success-content {
  text-align: center;
}

.success-icon-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 80px;
  height: 80px;
  margin: 0 auto 24px;
  background: linear-gradient(135deg, #52c41a 0%, #389e0d 100%);
  border-radius: 50%;
  box-shadow: 0 8px 24px rgba(82, 196, 26, 0.3);
}

.success-icon {
  font-size: 40px;
  color: #fff;
}

.success-title {
  font-size: 24px;
  font-weight: 600;
  color: var(--text-color);
  margin: 0 0 8px 0;
}

.success-desc {
  font-size: 14px;
  color: var(--text-color-secondary);
  margin: 0 0 24px 0;
}

.success-part-id {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
  padding: 12px;
  background: var(--bg-color);
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
  font-family: 'Courier New', monospace;
}

.success-actions {
  display: flex;
  justify-content: center;
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
    flex-direction: column;
    gap: 12px;
    align-items: flex-start;
  }
  
  .top-bar-left {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
  
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .material-fields {
    grid-template-columns: 1fr;
  }
  
  .form-actions {
    flex-direction: column;
  }
}
</style>
