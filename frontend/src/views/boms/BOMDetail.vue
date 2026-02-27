<template>
  <div class="bom-detail">
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">BOM详情</h1>
          <span class="page-subtitle">查看物料清单完整信息</span>
        </div>
      </div>
      <div class="top-bar-right">
        <a-tag :color="getStatusColor(bomData?.status || '')" class="status-badge">
          {{ getStatusText(bomData?.status || '') }}
        </a-tag>
      </div>
    </div>

    <a-spin :spinning="loading">
      <div class="main-content">
        <div class="info-container">
          <a-card :bordered="false" class="info-card">
            <div class="card-header">
              <div class="header-icon">
                <FileTextOutlined />
              </div>
              <div class="header-text">
                <h3>基本信息</h3>
                <p>BOM核心属性</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="info-grid">
              <div class="info-item">
                <div class="info-label">
                  <NumberOutlined class="label-icon" />
                  <span>BOM ID</span>
                </div>
                <div class="info-value bom-id">{{ bomData?.bomID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <TagOutlined class="label-icon" />
                  <span>BOM类型</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getBOMTypeColor(bomData?.bomType || '')" class="type-tag">
                    {{ bomData?.bomType || '-' }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CarOutlined class="label-icon" />
                  <span>对应车型</span>
                </div>
                <div class="info-value">{{ bomData?.productModel || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <BarcodeOutlined class="label-icon" />
                  <span>零部件批次号</span>
                </div>
                <div class="info-value batch-no">{{ bomData?.partBatchNo || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CodeOutlined class="label-icon" />
                  <span>版本号</span>
                </div>
                <div class="info-value version">{{ bomData?.version || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <UserOutlined class="label-icon" />
                  <span>创建人</span>
                </div>
                <div class="info-value">{{ bomData?.creator || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CheckCircleOutlined class="label-icon" />
                  <span>状态</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getStatusColor(bomData?.status || '')" class="status-tag">
                    {{ getStatusText(bomData?.status || '') }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ClockCircleOutlined class="label-icon" />
                  <span>创建时间</span>
                </div>
                <div class="info-value">{{ bomData?.createTime || '-' }}</div>
              </div>
            </div>
          </a-card>

          <a-card :bordered="false" class="action-card">
            <div class="card-header">
              <div class="header-icon action-icon">
                <ThunderboltOutlined />
              </div>
              <div class="header-text">
                <h3>快捷操作</h3>
                <p>常用功能入口</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="action-grid">
              <div class="action-item" @click="editBOM">
                <EditOutlined class="action-icon-item" />
                <span>编辑BOM</span>
              </div>
              <div class="action-item" @click="printBOM">
                <PrinterOutlined class="action-icon-item" />
                <span>打印</span>
              </div>
              <div class="action-item" @click="exportData">
                <ExportOutlined class="action-icon-item" />
                <span>导出数据</span>
              </div>
              <div class="action-item" @click="compareBOM">
                <SwapOutlined class="action-icon-item" />
                <span>BOM对比</span>
              </div>
              <div class="action-item delete-action" @click="handleDelete">
                <DeleteOutlined class="action-icon-item" />
                <span>删除该数据</span>
              </div>
            </div>
          </a-card>
        </div>

        <div class="details-container">
          <a-card :bordered="false" class="details-card">
            <div class="card-header">
              <div class="header-icon details-icon">
                <UnorderedListOutlined />
              </div>
              <div class="header-text">
                <h3>物料明细</h3>
                <p>BOM物料清单</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <a-table
              :columns="materialColumns"
              :data-source="bomData?.materialList || []"
              :pagination="false"
              :scroll="{ x: 800 }"
              class="material-table"
            >
              <template #bodyCell="{ column, record }">
                <template v-if="column.key === 'materialID'">
                  <span class="material-id">{{ record.materialID }}</span>
                </template>
                <template v-else-if="column.key === 'quantity'">
                  <span class="quantity">{{ record.quantity }}</span>
                </template>
              </template>
            </a-table>
          </a-card>

          <a-card :bordered="false" class="details-card" v-if="bomData?.bomType === '研发BOM'">
            <div class="card-header">
              <div class="header-icon details-icon">
                <FileOutlined />
              </div>
              <div class="header-text">
                <h3>研发BOM文件</h3>
                <p>相关文档信息</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="file-info">
              <div class="info-item">
                <div class="info-label">文件名</div>
                <div class="info-value">{{ bomData?.rdBOMFileInfo?.fileName || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">上传时间</div>
                <div class="info-value">{{ bomData?.rdBOMFileInfo?.uploadTime || '-' }}</div>
              </div>
              <a-button type="link" @click="downloadFile(bomData?.rdBOMFileInfo?.fileURL)" v-if="bomData?.rdBOMFileInfo?.fileURL">
                <DownloadOutlined /> 下载文件
              </a-button>
            </div>
          </a-card>

          <a-card :bordered="false" class="details-card" v-if="bomData?.bomType === '生产BOM'">
            <div class="card-header">
              <div class="header-icon details-icon">
                <SettingOutlined />
              </div>
              <div class="header-text">
                <h3>生产BOM信息</h3>
                <p>生产工艺相关</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="production-info">
              <div class="info-item">
                <div class="info-label">生产线</div>
                <div class="info-value">{{ bomData?.productionBOMInfo?.productionLine || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">工艺流程</div>
                <div class="info-value">{{ bomData?.productionBOMInfo?.processFlow || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">质量标准</div>
                <div class="info-value">{{ bomData?.productionBOMInfo?.qualityStandard || '-' }}</div>
              </div>
            </div>
          </a-card>

          <a-card :bordered="false" class="details-card" v-if="bomData?.changeRecords?.length">
            <div class="card-header">
              <div class="header-icon details-icon">
                <HistoryOutlined />
              </div>
              <div class="header-text">
                <h3>变更记录</h3>
                <p>BOM版本变更历史</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <a-timeline mode="left" class="custom-timeline">
              <a-timeline-item v-for="record in bomData.changeRecords" :key="record.changeID">
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">版本变更</span>
                    <span class="timeline-time">{{ record.changeTime }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <UserOutlined />
                      变更人: {{ record.changer }}
                    </div>
                    <div class="timeline-desc">
                      <EditOutlined />
                      变更内容: {{ record.changeContent }}
                    </div>
                  </div>
                </div>
              </a-timeline-item>
            </a-timeline>
          </a-card>
        </div>
      </div>
    </a-spin>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, h } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { message, Modal } from 'ant-design-vue'
import { getBOM, deleteBOM } from '../../services/bom.service'
import type { BOM } from '../../services/bom.service'
import {
  ArrowLeftOutlined,
  NumberOutlined,
  TagOutlined,
  CarOutlined,
  BarcodeOutlined,
  CodeOutlined,
  UserOutlined,
  CheckCircleOutlined,
  ClockCircleOutlined,
  ThunderboltOutlined,
  FileTextOutlined,
  EditOutlined,
  PrinterOutlined,
  ExportOutlined,
  SwapOutlined,
  DeleteOutlined,
  UnorderedListOutlined,
  FileOutlined,
  DownloadOutlined,
  SettingOutlined,
  HistoryOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()

const loading = ref(false)
const bomData = ref<BOM | null>(null)

const materialColumns = [
  {
    title: '物料ID',
    dataIndex: 'materialID',
    key: 'materialID',
    width: 200
  },
  {
    title: '物料名称',
    dataIndex: 'materialName',
    key: 'materialName',
    width: 200
  },
  {
    title: '规格型号',
    dataIndex: 'specifications',
    key: 'specifications',
    width: 150
  },
  {
    title: '数量',
    dataIndex: 'quantity',
    key: 'quantity',
    width: 100
  },
  {
    title: '单位',
    dataIndex: 'unit',
    key: 'unit',
    width: 100
  }
]

function goBack() {
  router.push('/bom/list')
}

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    '草稿': 'default',
    '已发布': 'success',
    '已变更': 'warning',
    '已作废': 'error'
  }
  return colors[status] || 'default'
}

function getStatusText(status: string) {
  return status || status
}

function getBOMTypeColor(type: string) {
  const colors: Record<string, string> = {
    '研发BOM': 'purple',
    '生产BOM': 'blue'
  }
  return colors[type] || 'default'
}

function editBOM() {
  const bomID = route.params.id as string
  router.push(`/bom/edit/${bomID}`)
}

function printBOM() {
  window.print()
}

function exportData() {
  if (!bomData.value) {
    message.error('没有可导出的数据')
    return
  }
  
  try {
    const bomID = route.params.id as string
    const data = bomData.value
    
    let csvContent = '\uFEFF'
    csvContent += 'BOM信息\n'
    csvContent += `BOM ID,${data.bomID}\n`
    csvContent += `BOM类型,${data.bomType}\n`
    csvContent += `对应车型,${data.productModel}\n`
    csvContent += `零部件批次号,${data.partBatchNo}\n`
    csvContent += `版本号,${data.version}\n`
    csvContent += `状态,${data.status}\n`
    csvContent += `创建人,${data.creator}\n`
    csvContent += `创建时间,${data.createTime}\n\n`
    
    csvContent += '物料明细\n'
    csvContent += '物料ID,物料名称,规格型号,数量,单位\n'
    
    if (data.materialList && data.materialList.length > 0) {
      data.materialList.forEach(material => {
        csvContent += `${material.materialID},${material.materialName || ''},${material.specifications || ''},${material.quantity},${material.unit || ''}\n`
      })
    }
    
    const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' })
    const link = document.createElement('a')
    const url = URL.createObjectURL(blob)
    
    link.setAttribute('href', url)
    link.setAttribute('download', `BOM_${bomID}_${data.version}.csv`)
    link.style.visibility = 'hidden'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    
    message.success('数据导出成功')
  } catch (error) {
    console.error('导出失败:', error)
    message.error('数据导出失败')
  }
}

function compareBOM() {
  const bomID = route.params.id as string
  router.push({ 
    path: '/bom/compare', 
    query: { bomId: bomID } 
  })
}

function downloadFile(url: string) {
  window.open(url, '_blank')
}

async function handleDelete() {
  const bomID = route.params.id as string
  if (!bomID) {
    message.error('BOM ID不存在')
    return
  }
  
  const user = JSON.parse(localStorage.getItem('user') || '{}')
  const userRole = user.role || user.userType || ''
  const hasDeletePermission = userRole === 'admin' || userRole === 'manufacturer' || userRole === 'manufacturer_user'
  
  if (!hasDeletePermission) {
    message.error('您没有权限删除该数据')
    return
  }
  
  Modal.confirm({
    title: '确认删除',
    icon: () => h(DeleteOutlined, { style: { color: '#ff4d4f' } }),
    maskClosable: false,
    maskStyle: { backgroundColor: 'rgba(0, 0, 0, 0.6)' },
    content: h('div', { style: { lineHeight: '1.8' } }, [
      h('p', { style: { marginBottom: '8px' } }, [
        h('strong', { style: { color: '#ff4d4f' } }, '警告：此操作不可恢复！')
      ]),
      h('p', { style: { marginBottom: '8px' } }, [
        `您确定要删除BOM ${bomID} 吗？`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `BOM类型：${bomData.value?.bomType || '-'}`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `版本号：${bomData.value?.version || '-'}`
      ]),
      h('p', { style: { fontSize: '13px', color: '#999' } }, [
        '删除后将无法恢复，请谨慎操作。'
      ])
    ]),
    okText: '确认删除',
    okType: 'danger',
    cancelText: '取消',
    onOk: async () => {
      try {
        await deleteBOM(bomID)
        message.success('删除成功')
        setTimeout(() => {
          goBack()
        }, 1000)
      } catch (error: any) {
        message.error(error.message || '删除失败')
      }
    }
  })
}

async function fetchBOMData() {
  const bomID = route.params.id as string
  if (!bomID) {
    message.error('BOM ID不存在')
    return
  }
  
  loading.value = true
  try {
    const response = await getBOM(bomID)
    if (response.code === 0 && response.data) {
      bomData.value = response.data
    } else {
      message.error(response.message || '获取BOM信息失败')
    }
  } catch (error: any) {
    message.error(error.message || '获取BOM信息失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchBOMData()
})
</script>

<style scoped>
.bom-detail {
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

.divider {
  height: 24px;
  background: var(--border-color);
}

.page-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.page-title {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-color);
  margin: 0;
}

.page-subtitle {
  font-size: 13px;
  color: var(--text-color-secondary);
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

.info-container {
  width: 480px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.details-container {
  flex: 1;
  min-width: 0;
}

.info-card,
.action-card,
.details-card {
  border-radius: 16px;
  box-shadow: var(--shadow-sm);
  background: var(--bg-color-secondary);
}

.card-header {
  display: flex;
  align-items: center;
  gap: 16px;
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

.header-icon.action-icon {
  background: linear-gradient(135deg, #fa8c16 0%, #fa541c 100%);
  box-shadow: 0 4px 12px rgba(250, 140, 22, 0.3);
}

.header-icon.details-icon {
  background: linear-gradient(135deg, #1890ff 0%, #722ed1 100%);
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3);
}

.header-icon.lifecycle-icon {
  background: linear-gradient(135deg, #1890ff 0%, #722ed1 100%);
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3);
}

.header-text h3 {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-color);
  margin: 0;
}

.header-text p {
  font-size: 13px;
  color: var(--text-color-secondary);
  margin: 2px 0 0 0;
}

.card-divider {
  margin: 20px 0;
}

.info-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  padding: 4px;
}

.info-item {
  padding: 16px;
  background: var(--bg-color);
  border-radius: 12px;
  border: 1px solid var(--border-color);
  transition: all 0.3s;
}

.info-item:hover {
  border-color: var(--primary-color);
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.1);
}

.info-label {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: var(--text-color-secondary);
  font-weight: 500;
  margin-bottom: 6px;
}

.label-icon {
  font-size: 14px;
  color: var(--text-color-tertiary);
}

.info-value {
  font-size: 14px;
  color: var(--text-color);
  font-weight: 500;
  word-break: break-all;
}

.bom-id,
.version {
  font-family: 'Courier New', monospace;
  color: var(--primary-color);
  font-weight: 600;
}

.action-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
  padding: 4px;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 16px;
  background: var(--bg-color);
  border-radius: 12px;
  border: 1px solid var(--border-color);
  cursor: pointer;
  transition: all 0.3s;
}

.action-item:hover {
  border-color: var(--primary-color);
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.15);
  transform: translateY(-2px);
}

.action-item.delete-action:hover {
  border-color: #ff4d4f;
  box-shadow: 0 2px 8px rgba(255, 77, 79, 0.15);
}

.action-icon-item {
  font-size: 20px;
  color: var(--primary-color);
}

.action-item.delete-action .action-icon-item {
  color: #ff4d4f;
}

.action-item span {
  font-size: 13px;
  color: var(--text-color);
  font-weight: 500;
}

.material-table {
  padding: 4px;
}

.material-id {
  font-family: 'Courier New', monospace;
  color: var(--primary-color);
  font-weight: 600;
}

.quantity {
  font-weight: 600;
  color: #fa8c16;
}

.file-info,
.production-info {
  padding: 4px;
}

.timeline-wrapper {
  padding: 4px;
}

.custom-timeline {
  margin-top: 16px;
}

.timeline-item {
  padding-left: 8px;
}

.timeline-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.timeline-title {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-color);
}

.timeline-time {
  font-size: 12px;
  color: var(--text-color-tertiary);
}

.timeline-content {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.timeline-content p {
  font-size: 13px;
  color: var(--text-color-secondary);
  margin: 0;
}

.timeline-desc {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: var(--text-color-secondary);
}

.timeline-desc .anticon {
  font-size: 14px;
  color: var(--text-color-tertiary);
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
  
  .action-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .material-table :deep(.ant-table) {
    font-size: 12px;
  }
}
</style>
