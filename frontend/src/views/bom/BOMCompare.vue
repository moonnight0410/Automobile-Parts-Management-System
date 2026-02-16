<!--
  BOM比较页面
-->

<template>
  <div class="bom-compare">
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <DiffOutlined />
          </div>
          <div class="title-text">
            <h2>BOM比较</h2>
            <p class="subtitle">对比不同BOM的差异内容</p>
          </div>
        </div>
        <div class="header-actions">
          <a-tag color="purple" class="mode-tag">
            <template #icon><SwapOutlined /></template>
            差异对比模式
          </a-tag>
        </div>
      </div>
    </div>

    <div class="compare-container">
      <div class="select-section">
        <a-card :bordered="false" class="select-card">
          <div class="select-header">
            <div class="header-left">
              <div class="select-icon-wrapper">
                <SelectOutlined />
              </div>
              <div class="select-info">
                <span class="select-title">选择BOM</span>
                <span class="select-desc">选择两个BOM进行差异对比分析</span>
              </div>
            </div>
          </div>

          <div class="select-content">
            <div class="select-row">
              <div class="select-item">
                <div class="item-label">
                  <span class="label-icon production"><AppstoreOutlined /></span>
                  <span>生产BOM</span>
                  <a-tag color="blue" size="small" class="source-tag">源</a-tag>
                </div>
                <a-input
                  v-model:value="selectedProductionBOM"
                  placeholder="请输入生产BOM ID"
                  class="custom-input"
                  allow-clear
                >
                  <template #prefix>
                    <span class="input-icon"><NumberOutlined /></span>
                  </template>
                </a-input>
              </div>

              <div class="compare-arrow">
                <div class="arrow-icon">
                  <SwapOutlined />
                </div>
                <span class="arrow-text">VS</span>
              </div>

              <div class="select-item">
                <div class="item-label">
                  <span class="label-icon research"><ExperimentOutlined /></span>
                  <span>研发BOM</span>
                  <a-tag color="green" size="small" class="source-tag">目标</a-tag>
                </div>
                <a-input
                  v-model:value="selectedResearchBOM"
                  placeholder="请输入研发BOM ID"
                  class="custom-input"
                  allow-clear
                >
                  <template #prefix>
                    <span class="input-icon"><NumberOutlined /></span>
                  </template>
                </a-input>
              </div>
            </div>

            <div class="select-actions">
              <a-button @click="handleReset" class="reset-btn">
                <template #icon><ReloadOutlined /></template>
                重置选择
              </a-button>
              <a-button 
                type="primary" 
                :disabled="!canCompare" 
                @click="handleCompare"
                :loading="comparing"
                class="compare-btn"
              >
                <template #icon><DiffOutlined /></template>
                {{ comparing ? '比较中...' : '开始比较' }}
              </a-button>
            </div>
          </div>
        </a-card>
      </div>

      <div class="result-section" v-if="showResult">
        <a-card :bordered="false" class="result-card">
          <div class="result-header">
            <div class="header-left">
              <div class="result-icon-wrapper">
                <BarChartOutlined />
              </div>
              <div class="result-info">
                <span class="result-title">比较结果</span>
                <span class="result-desc">发现 {{ differenceCount }} 处差异</span>
              </div>
            </div>
            <div class="header-right">
              <a-space :size="12">
                <a-button class="action-btn" @click="exportResult">
                  <template #icon><ExportOutlined /></template>
                  导出报告
                </a-button>
                <a-button class="action-btn" @click="handleReset">
                  <template #icon><ReloadOutlined /></template>
                  重新比较
                </a-button>
              </a-space>
            </div>
          </div>

          <a-divider class="result-divider" />

          <div class="result-summary">
            <div class="summary-item added">
              <div class="summary-icon">
                <PlusCircleOutlined />
              </div>
              <div class="summary-content">
                <span class="summary-value">{{ addedCount }}</span>
                <span class="summary-label">新增项</span>
              </div>
            </div>
            <div class="summary-item removed">
              <div class="summary-icon">
                <MinusCircleOutlined />
              </div>
              <div class="summary-content">
                <span class="summary-value">{{ removedCount }}</span>
                <span class="summary-label">删除项</span>
              </div>
            </div>
            <div class="summary-item modified">
              <div class="summary-icon">
                <EditOutlined />
              </div>
              <div class="summary-content">
                <span class="summary-value">{{ modifiedCount }}</span>
                <span class="summary-label">修改项</span>
              </div>
            </div>
            <div class="summary-item unchanged">
              <div class="summary-icon">
                <CheckCircleOutlined />
              </div>
              <div class="summary-content">
                <span class="summary-value">{{ unchangedCount }}</span>
                <span class="summary-label">未变化</span>
              </div>
            </div>
          </div>

          <div class="result-table">
            <a-table
              :columns="resultColumns"
              :data-source="resultData"
              :pagination="{ pageSize: 10 }"
              class="compare-table"
              :scroll="{ x: 1200 }"
            >
              <template #bodyCell="{ column, record }">
                <template v-if="column.key === 'changeType'">
                  <a-tag :color="getChangeTypeColor(record.changeType)" class="change-tag">
                    <template #icon>
                      <component :is="getChangeTypeIcon(record.changeType)" />
                    </template>
                    {{ record.changeType }}
                  </a-tag>
                </template>
                <template v-else-if="column.key === 'partID'">
                  <div class="part-id-cell">
                    <BarcodeOutlined class="part-icon" />
                    <span>{{ record.partID }}</span>
                  </div>
                </template>
                <template v-else-if="column.key === 'productionValue'">
                  <div class="value-cell" :class="{ 'value-old': record.changeType === '修改' }">
                    {{ record.productionValue || '-' }}
                  </div>
                </template>
                <template v-else-if="column.key === 'researchValue'">
                  <div class="value-cell" :class="{ 'value-new': record.changeType === '修改' }">
                    {{ record.researchValue || '-' }}
                  </div>
                </template>
              </template>
            </a-table>
          </div>
        </a-card>
      </div>

      <div class="empty-section" v-else>
        <a-card :bordered="false" class="empty-card">
          <a-empty description="请选择两个BOM进行比较">
            <template #image>
              <div class="empty-icon">
                <DiffOutlined />
              </div>
            </template>
            <div class="empty-content">
              <p class="empty-title">尚未开始比较</p>
              <p class="empty-desc">请在上方选择两个BOM，点击"开始比较"按钮查看差异分析结果</p>
              <div class="empty-tips">
                <div class="tip-item">
                  <InfoCircleOutlined />
                  <span>支持生产BOM与研发BOM的对比</span>
                </div>
                <div class="tip-item">
                  <InfoCircleOutlined />
                  <span>比较结果将显示新增、删除、修改的项目</span>
                </div>
              </div>
            </div>
          </a-empty>
        </a-card>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { message } from 'ant-design-vue'
import {
  DiffOutlined,
  SwapOutlined,
  SelectOutlined,
  AppstoreOutlined,
  ExperimentOutlined,
  ReloadOutlined,
  BarChartOutlined,
  ExportOutlined,
  PlusCircleOutlined,
  MinusCircleOutlined,
  EditOutlined,
  CheckCircleOutlined,
  BarcodeOutlined,
  InfoCircleOutlined,
  PlusOutlined,
  MinusOutlined,
  NumberOutlined
} from '@ant-design/icons-vue'

const selectedProductionBOM = ref<string | undefined>(undefined)
const selectedResearchBOM = ref<string | undefined>(undefined)
const comparing = ref(false)
const showResult = ref(false)

const canCompare = computed(() => {
  return selectedProductionBOM.value && selectedResearchBOM.value
})

const addedCount = ref(5)
const removedCount = ref(3)
const modifiedCount = ref(8)
const unchangedCount = ref(42)
const differenceCount = computed(() => addedCount.value + removedCount.value + modifiedCount.value)

const resultColumns = [
  {
    title: '变更类型',
    key: 'changeType',
    width: 120,
    fixed: 'left' as const
  },
  {
    title: '零部件ID',
    key: 'partID',
    width: 150
  },
  {
    title: '零部件名称',
    dataIndex: 'partName',
    key: 'partName',
    width: 180
  },
  {
    title: '生产BOM值',
    key: 'productionValue',
    width: 150
  },
  {
    title: '研发BOM值',
    key: 'researchValue',
    width: 150
  },
  {
    title: '差异说明',
    dataIndex: 'difference',
    key: 'difference',
    width: 200
  }
]

const resultData = ref([
  {
    key: '1',
    changeType: '新增',
    partID: 'PART-NEW-001',
    partName: '新型传感器组件',
    productionValue: '-',
    researchValue: 'A-2024',
    difference: '研发版本新增零部件'
  },
  {
    key: '2',
    changeType: '删除',
    partID: 'PART-OLD-015',
    partName: '旧版控制模块',
    productionValue: 'V2.0',
    researchValue: '-',
    difference: '研发版本已移除该零部件'
  },
  {
    key: '3',
    changeType: '修改',
    partID: 'PART-MOD-008',
    partName: '发动机支架',
    productionValue: '材质: 钢',
    researchValue: '材质: 铝合金',
    difference: '材质由钢改为铝合金'
  },
  {
    key: '4',
    changeType: '修改',
    partID: 'PART-MOD-012',
    partName: '刹车盘',
    productionValue: '直径: 300mm',
    researchValue: '直径: 320mm',
    difference: '尺寸规格调整'
  },
  {
    key: '5',
    changeType: '新增',
    partID: 'PART-NEW-003',
    partName: '智能控制单元',
    productionValue: '-',
    researchValue: 'ICU-2024',
    difference: '新增智能控制功能'
  }
])

const getChangeTypeColor = (type: string) => {
  const colorMap: Record<string, string> = {
    '新增': 'green',
    '删除': 'red',
    '修改': 'orange'
  }
  return colorMap[type] || 'default'
}

const getChangeTypeIcon = (type: string) => {
  const iconMap: Record<string, any> = {
    '新增': PlusOutlined,
    '删除': MinusOutlined,
    '修改': EditOutlined
  }
  return iconMap[type] || InfoCircleOutlined
}

const handleCompare = () => {
  if (!canCompare.value) {
    message.warning('请选择两个BOM进行比较')
    return
  }
  
  comparing.value = true
  setTimeout(() => {
    comparing.value = false
    showResult.value = true
    message.success('比较完成')
  }, 1500)
}

const handleReset = () => {
  selectedProductionBOM.value = undefined
  selectedResearchBOM.value = undefined
  showResult.value = false
}

const exportResult = () => {
  message.success('比较报告导出成功')
}
</script>

<style scoped>
.bom-compare {
  padding: 0;
  min-height: 100%;
}

.page-header {
  margin-bottom: 24px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  padding: 24px 32px;
  box-shadow: 0 4px 20px rgba(102, 126, 234, 0.25);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
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

.mode-tag {
  padding: 6px 16px;
  border-radius: 20px;
  font-size: 13px;
  background: rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: #fff;
}

.compare-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.select-section .select-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.select-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.select-icon-wrapper {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: #fff;
}

.select-info {
  display: flex;
  flex-direction: column;
}

.select-title {
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
}

.select-desc {
  font-size: 13px;
  color: #6b7280;
  margin-top: 2px;
}

.select-content {
  background: #f9fafb;
  border-radius: 12px;
  padding: 24px;
}

.select-row {
  display: flex;
  align-items: flex-end;
  gap: 24px;
  margin-bottom: 24px;
}

.select-item {
  flex: 1;
}

.item-label {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 12px;
  font-size: 14px;
  font-weight: 500;
  color: #374151;
}

.label-icon {
  width: 28px;
  height: 28px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
}

.label-icon.production {
  background: #dbeafe;
  color: #2563eb;
}

.label-icon.research {
  background: #d1fae5;
  color: #059669;
}

.source-tag {
  margin-left: auto;
}

.custom-input {
  width: 100%;
  height: 42px;
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
  border: 2px solid #e5e7eb;
  background: #fff;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.custom-input :deep(.ant-input) {
  font-size: 14px;
  color: #1f2937;
}

.custom-input :deep(.ant-input::placeholder) {
  color: #9ca3af;
}

.custom-input :deep(.ant-input-prefix) {
  margin-right: 8px;
  color: #9ca3af;
}

.input-icon {
  font-size: 14px;
  color: #9ca3af;
  transition: color 0.3s;
}

.custom-input :deep(.ant-input-affix-wrapper:hover) {
  border-color: #667eea;
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.15);
}

.custom-input :deep(.ant-input-affix-wrapper:focus),
.custom-input :deep(.ant-input-affix-wrapper-focused) {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.12);
}

.custom-input :deep(.ant-input-affix-wrapper:focus-within) .input-icon {
  color: #667eea;
}

.compare-arrow {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 0 16px;
  min-width: 80px;
}

.arrow-icon {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: #fff;
  margin-bottom: 8px;
}

.arrow-text {
  font-size: 14px;
  font-weight: 600;
  color: #667eea;
}

.select-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

.reset-btn {
  height: 40px;
  border-radius: 10px;
  font-size: 14px;
  border: 1px solid #e5e7eb;
  color: #6b7280;
}

.reset-btn:hover {
  border-color: #667eea;
  color: #667eea;
}

.compare-btn {
  height: 40px;
  border-radius: 10px;
  font-size: 14px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  padding: 0 32px;
}

.compare-btn:hover:not(:disabled) {
  background: linear-gradient(135deg, #5a67d8 0%, #6b46c1 100%);
}

.compare-btn:disabled {
  background: #e5e7eb;
  color: #9ca3af;
}

.result-section .result-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.result-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.result-icon-wrapper {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: #fff;
}

.result-info {
  display: flex;
  flex-direction: column;
}

.result-title {
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
}

.result-desc {
  font-size: 13px;
  color: #6b7280;
  margin-top: 2px;
}

.action-btn {
  height: 36px;
  border-radius: 8px;
  font-size: 13px;
  border: 1px solid #e5e7eb;
  color: #6b7280;
}

.action-btn:hover {
  border-color: #667eea;
  color: #667eea;
}

.result-divider {
  margin: 20px 0;
}

.result-summary {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
  margin-bottom: 24px;
}

.summary-item {
  background: #f9fafb;
  border-radius: 12px;
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 16px;
  transition: all 0.3s ease;
}

.summary-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.summary-item.added {
  border-left: 4px solid #10b981;
}

.summary-item.added .summary-icon {
  background: #d1fae5;
  color: #059669;
}

.summary-item.removed {
  border-left: 4px solid #ef4444;
}

.summary-item.removed .summary-icon {
  background: #fee2e2;
  color: #dc2626;
}

.summary-item.modified {
  border-left: 4px solid #f59e0b;
}

.summary-item.modified .summary-icon {
  background: #fef3c7;
  color: #d97706;
}

.summary-item.unchanged {
  border-left: 4px solid #6b7280;
}

.summary-item.unchanged .summary-icon {
  background: #e5e7eb;
  color: #4b5563;
}

.summary-icon {
  width: 44px;
  height: 44px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
}

.summary-content {
  display: flex;
  flex-direction: column;
}

.summary-value {
  font-size: 28px;
  font-weight: 700;
  color: #1f2937;
  line-height: 1;
}

.summary-label {
  font-size: 13px;
  color: #6b7280;
  margin-top: 4px;
}

.result-table {
  margin-top: 16px;
}

.compare-table :deep(.ant-table-thead > tr > th) {
  background: #f9fafb;
  font-weight: 600;
  color: #374151;
  border-bottom: 2px solid #e5e7eb;
}

.compare-table :deep(.ant-table-tbody > tr > td) {
  border-bottom: 1px solid #f3f4f6;
}

.compare-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #f9fafb;
}

.change-tag {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 4px 12px;
  border-radius: 6px;
  font-size: 13px;
}

.part-id-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.part-icon {
  color: #6b7280;
}

.value-cell {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 13px;
}

.value-old {
  background: #fef3c7;
  color: #92400e;
}

.value-new {
  background: #d1fae5;
  color: #065f46;
}

.empty-section .empty-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.empty-icon {
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 36px;
  color: #fff;
  margin: 0 auto;
}

.empty-content {
  margin-top: 24px;
}

.empty-title {
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 8px;
}

.empty-desc {
  font-size: 14px;
  color: #6b7280;
  margin-bottom: 20px;
}

.empty-tips {
  display: flex;
  flex-direction: column;
  gap: 8px;
  align-items: center;
}

.tip-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #9ca3af;
}

[data-theme='dark'] .page-header {
  background: linear-gradient(135deg, #4c51bf 0%, #553c9a 100%);
  box-shadow: 0 4px 20px rgba(76, 81, 191, 0.3);
}

[data-theme='dark'] .title-text h2 {
  color: #fff;
}

[data-theme='dark'] .title-text .subtitle {
  color: rgba(255, 255, 255, 0.8);
}

[data-theme='dark'] .mode-tag {
  background: rgba(255, 255, 255, 0.15);
  border-color: rgba(255, 255, 255, 0.25);
  color: #fff;
}

[data-theme='dark'] .select-card,
[data-theme='dark'] .result-card,
[data-theme='dark'] .empty-card {
  background: #1f2937;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

[data-theme='dark'] .select-title,
[data-theme='dark'] .result-title,
[data-theme='dark'] .empty-title {
  color: #f3f4f6;
}

[data-theme='dark'] .select-desc,
[data-theme='dark'] .result-desc,
[data-theme='dark'] .empty-desc {
  color: #9ca3af;
}

[data-theme='dark'] .select-content {
  background: #111827;
}

[data-theme='dark'] .item-label {
  color: #e5e7eb;
}

[data-theme='dark'] .custom-input :deep(.ant-input-affix-wrapper) {
  background: #1f2937;
  border-color: #374151;
}

[data-theme='dark'] .custom-input :deep(.ant-input) {
  color: #f3f4f6;
}

[data-theme='dark'] .custom-input :deep(.ant-input::placeholder) {
  color: #6b7280;
}

[data-theme='dark'] .custom-input :deep(.ant-input-prefix) {
  color: #6b7280;
}

[data-theme='dark'] .input-icon {
  color: #6b7280;
}

[data-theme='dark'] .custom-input :deep(.ant-input-affix-wrapper:hover) {
  border-color: #818cf8;
}

[data-theme='dark'] .custom-input :deep(.ant-input-affix-wrapper:focus),
[data-theme='dark'] .custom-input :deep(.ant-input-affix-wrapper-focused) {
  border-color: #818cf8;
  box-shadow: 0 0 0 3px rgba(129, 140, 248, 0.15);
}

[data-theme='dark'] .custom-input :deep(.ant-input-affix-wrapper:focus-within) .input-icon {
  color: #818cf8;
}

[data-theme='dark'] .reset-btn,
[data-theme='dark'] .action-btn {
  border-color: #374151;
  color: #9ca3af;
}

[data-theme='dark'] .reset-btn:hover,
[data-theme='dark'] .action-btn:hover {
  border-color: #667eea;
  color: #a78bfa;
}

[data-theme='dark'] .compare-btn:disabled {
  background: #374151;
  color: #6b7280;
}

[data-theme='dark'] .result-summary {
  background: transparent;
}

[data-theme='dark'] .summary-item {
  background: #111827;
}

[data-theme='dark'] .summary-item.added {
  border-left-color: #34d399;
}

[data-theme='dark'] .summary-item.removed {
  border-left-color: #f87171;
}

[data-theme='dark'] .summary-item.modified {
  border-left-color: #fbbf24;
}

[data-theme='dark'] .summary-item.unchanged {
  border-left-color: #9ca3af;
}

[data-theme='dark'] .summary-value {
  color: #f3f4f6;
}

[data-theme='dark'] .summary-label {
  color: #9ca3af;
}

[data-theme='dark'] .compare-table :deep(.ant-table-thead > tr > th) {
  background: #111827;
  color: #e5e7eb;
  border-bottom-color: #374151;
}

[data-theme='dark'] .compare-table :deep(.ant-table-tbody > tr > td) {
  background: #1f2937;
  border-bottom-color: #374151;
  color: #e5e7eb;
}

[data-theme='dark'] .compare-table :deep(.ant-table-tbody > tr:hover > td) {
  background: #111827;
}

[data-theme='dark'] .part-icon {
  color: #9ca3af;
}

[data-theme='dark'] .value-old {
  background: rgba(251, 191, 36, 0.2);
  color: #fbbf24;
}

[data-theme='dark'] .value-new {
  background: rgba(52, 211, 153, 0.2);
  color: #34d399;
}

[data-theme='dark'] .empty-icon {
  background: linear-gradient(135deg, #4c51bf 0%, #553c9a 100%);
}

[data-theme='dark'] .tip-item {
  color: #6b7280;
}

[data-theme='dark'] .label-icon.production {
  background: rgba(37, 99, 235, 0.2);
  color: #60a5fa;
}

[data-theme='dark'] .label-icon.research {
  background: rgba(5, 150, 105, 0.2);
  color: #34d399;
}
</style>
