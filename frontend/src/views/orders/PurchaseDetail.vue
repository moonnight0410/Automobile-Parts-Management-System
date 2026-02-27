<template>
  <div class="purchase-detail">
    <div class="top-bar">
      <div class="top-bar-left">
        <a-button type="text" class="back-btn" @click="goBack">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <a-divider type="vertical" class="divider" />
        <div class="page-info">
          <h1 class="page-title">采购订单详情</h1>
          <span class="page-subtitle">查看采购订单完整信息</span>
        </div>
      </div>
      <div class="top-bar-right">
        <a-tag :color="getStatusColor(orderData?.status || '')" class="status-badge">
          {{ getStatusText(orderData?.status || '') }}
        </a-tag>
      </div>
    </div>

    <a-spin :spinning="loading">
      <div class="main-content">
        <div class="info-container">
          <a-card :bordered="false" class="info-card">
            <div class="card-header">
              <div class="header-icon">
                <ShoppingCartOutlined />
              </div>
              <div class="header-text">
                <h3>基本信息</h3>
                <p>采购订单核心属性</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="info-grid">
              <div class="info-item">
                <div class="info-label">
                  <NumberOutlined class="label-icon" />
                  <span>订单ID</span>
                </div>
                <div class="info-value order-id">{{ orderData?.orderID || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <UserOutlined class="label-icon" />
                  <span>采购方</span>
                </div>
                <div class="info-value">{{ orderData?.buyer || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ShopOutlined class="label-icon" />
                  <span>供应商</span>
                </div>
                <div class="info-value">{{ orderData?.seller || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CheckCircleOutlined class="label-icon" />
                  <span>状态</span>
                </div>
                <div class="info-value">
                  <a-tag :color="getStatusColor(orderData?.status || '')" class="status-tag">
                    {{ getStatusText(orderData?.status || '') }}
                  </a-tag>
                </div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <CalendarOutlined class="label-icon" />
                  <span>订单日期</span>
                </div>
                <div class="info-value">{{ orderData?.orderDate || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <ClockCircleOutlined class="label-icon" />
                  <span>交货日期</span>
                </div>
                <div class="info-value">{{ orderData?.deliveryDate || '-' }}</div>
              </div>
              <div class="info-item">
                <div class="info-label">
                  <DollarOutlined class="label-icon" />
                  <span>订单金额</span>
                </div>
                <div class="info-value amount">¥{{ formatAmount(orderData?.totalAmount || 0) }}</div>
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
              <div class="action-item" @click="editOrder">
                <EditOutlined class="action-icon-item" />
                <span>编辑订单</span>
              </div>
              <div class="action-item" @click="printOrder">
                <PrinterOutlined class="action-icon-item" />
                <span>打印</span>
              </div>
              <div class="action-item" @click="exportData">
                <ExportOutlined class="action-icon-item" />
                <span>导出数据</span>
              </div>
              <div class="action-item" @click="viewLogistics">
                <CarOutlined class="action-icon-item" />
                <span>物流跟踪</span>
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
                <h3>订单明细</h3>
                <p>采购物品清单</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <a-table
              :columns="itemColumns"
              :data-source="orderData?.items || []"
              :pagination="false"
              :scroll="{ x: 800 }"
              class="item-table"
            >
              <template #bodyCell="{ column, record }">
                <template v-if="column.key === 'partID'">
                  <span class="part-id">{{ record.partID }}</span>
                </template>
                <template v-else-if="column.key === 'quantity'">
                  <span class="quantity">{{ record.quantity }}</span>
                </template>
                <template v-else-if="column.key === 'unitPrice'">
                  <span class="price">¥{{ record.unitPrice.toFixed(2) }}</span>
                </template>
                <template v-else-if="column.key === 'total'">
                  <span class="total">¥{{ (record.quantity * record.unitPrice).toFixed(2) }}</span>
                </template>
              </template>
              <template #summary>
                <a-table-summary>
                  <a-table-summary-row>
                    <a-table-summary-cell :index="0" :col-span="3">总计</a-table-summary-cell>
                    <a-table-summary-cell :index="1">
                      <span class="total-amount">¥{{ formatAmount(orderData?.totalAmount || 0) }}</span>
                    </a-table-summary-cell>
                  </a-table-summary-row>
                </a-table-summary>
              </template>
            </a-table>
          </a-card>

          <a-card :bordered="false" class="details-card" v-if="orderData?.remarks">
            <div class="card-header">
              <div class="header-icon details-icon">
                <FileTextOutlined />
              </div>
              <div class="header-text">
                <h3>备注信息</h3>
                <p>订单备注</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <div class="remarks-content">
              {{ orderData.remarks }}
            </div>
          </a-card>

          <a-card :bordered="false" class="details-card">
            <div class="card-header">
              <div class="header-icon details-icon">
                <HistoryOutlined />
              </div>
              <div class="header-text">
                <h3>订单时间线</h3>
                <p>订单流程记录</p>
              </div>
            </div>
            <a-divider class="card-divider" />
            <a-timeline mode="left" class="custom-timeline">
              <a-timeline-item color="blue">
                <template #dot>
                  <div class="timeline-dot blue">
                    <PlusCircleOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">订单创建</span>
                    <span class="timeline-time">{{ orderData?.orderDate }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <ShoppingCartOutlined />
                      订单已创建
                    </div>
                    <div class="timeline-desc">
                      <UserOutlined />
                      采购方: {{ orderData?.buyer }}
                    </div>
                    <div class="timeline-desc">
                      <ShopOutlined />
                      供应商: {{ orderData?.seller }}
                    </div>
                  </div>
                </div>
              </a-timeline-item>

              <a-timeline-item color="green">
                <template #dot>
                  <div class="timeline-dot green">
                    <CheckCircleOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">订单确认</span>
                    <span class="timeline-time">{{ orderData?.orderDate }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <CheckCircleOutlined />
                      订单金额: ¥{{ formatAmount(orderData?.totalAmount || 0) }}
                    </div>
                    <div class="timeline-desc">
                      <CalendarOutlined />
                      交货日期: {{ orderData?.deliveryDate }}
                    </div>
                  </div>
                </div>
              </a-timeline-item>

              <a-timeline-item color="orange">
                <template #dot>
                  <div class="timeline-dot orange">
                    <CarOutlined />
                  </div>
                </template>
                <div class="timeline-item">
                  <div class="timeline-header">
                    <span class="timeline-title">物流配送</span>
                    <span class="timeline-time">{{ orderData?.deliveryDate }}</span>
                  </div>
                  <div class="timeline-content">
                    <div class="timeline-desc">
                      <CarOutlined />
                      物流配送中
                    </div>
                    <div class="timeline-desc">
                      <ClockCircleOutlined />
                      预计交货日期: {{ orderData?.deliveryDate }}
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
import { listSupplyOrders } from '../../services/supply.service'
import type { SupplyOrder } from '../../services/supply.service'
import {
  ArrowLeftOutlined,
  NumberOutlined,
  UserOutlined,
  ShopOutlined,
  CheckCircleOutlined,
  CalendarOutlined,
  ClockCircleOutlined,
  DollarOutlined,
  ThunderboltOutlined,
  ShoppingCartOutlined,
  EditOutlined,
  PrinterOutlined,
  ExportOutlined,
  DeleteOutlined,
  UnorderedListOutlined,
  FileTextOutlined,
  HistoryOutlined,
  PlusCircleOutlined,
  CarOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()

const loading = ref(false)
const orderData = ref<SupplyOrder | null>(null)

const itemColumns = [
  {
    title: '零部件ID',
    dataIndex: 'partID',
    key: 'partID',
    width: 200
  },
  {
    title: '零部件名称',
    dataIndex: 'partName',
    key: 'partName',
    width: 200
  },
  {
    title: '数量',
    dataIndex: 'quantity',
    key: 'quantity',
    width: 100
  },
  {
    title: '单价',
    dataIndex: 'unitPrice',
    key: 'unitPrice',
    width: 120
  },
  {
    title: '小计',
    key: 'total',
    width: 120
  }
]

function goBack() {
  router.push('/supply/orders')
}

function getStatusColor(status: string) {
  const colors: Record<string, string> = {
    '待确认': 'warning',
    '已确认': 'processing',
    '配送中': 'blue',
    '已完成': 'success',
    '已取消': 'default'
  }
  return colors[status] || 'default'
}

function getStatusText(status: string) {
  return status || status
}

function formatAmount(amount: number) {
  return amount.toLocaleString('zh-CN', { minimumFractionDigits: 2, maximumFractionDigits: 2 })
}

function editOrder() {
  message.info('编辑功能开发中')
}

function printOrder() {
  message.success('打印任务已提交')
}

function exportData() {
  message.success('数据导出成功')
}

function viewLogistics() {
  message.info('物流跟踪功能开发中')
}

async function handleDelete() {
  const orderID = route.params.id as string
  if (!orderID) {
    message.error('订单ID不存在')
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
        `您确定要删除采购订单 ${orderID} 吗？`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `采购方：${orderData.value?.buyer || '-'}`
      ]),
      h('p', { style: { marginBottom: '8px', fontSize: '13px', color: '#999' } }, [
        `供应商：${orderData.value?.seller || '-'}`
      ]),
      h('p', { style: { fontSize: '13px', color: '#999' } }, [
        '删除后将无法恢复，请谨慎操作。'
      ])
    ]),
    okText: '确认删除',
    okType: 'danger',
    cancelText: '取消',
    onOk: async () => {
      message.success('删除成功')
      setTimeout(() => {
        goBack()
      }, 1000)
    }
  })
}

async function fetchOrderData() {
  loading.value = true
  try {
    const response = await listSupplyOrders()
    if (response.code === 0 && response.data) {
      const orderID = route.params.id as string
      const order = response.data.find((item: SupplyOrder) => item.orderID === orderID)
      if (order) {
        orderData.value = order
      } else {
        message.error('未找到该采购订单')
      }
    } else {
      message.error(response.message || '获取采购订单失败')
    }
  } catch (error: any) {
    message.error(error.message || '获取采购订单失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchOrderData()
})
</script>

<style scoped>
.purchase-detail {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  padding: 24px;
}

.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  background: white;
  padding: 16px 24px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.top-bar-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.back-btn {
  font-size: 14px;
  color: #1890ff;
}

.divider {
  height: 24px;
  margin: 0;
}

.page-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.page-title {
  font-size: 20px;
  font-weight: 600;
  margin: 0;
  color: #1f2937;
}

.page-subtitle {
  font-size: 13px;
  color: #6b7280;
}

.top-bar-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.status-badge {
  font-size: 13px;
  padding: 4px 12px;
}

.main-content {
  display: grid;
  grid-template-columns: 400px 1fr;
  gap: 24px;
}

.info-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.info-card,
.action-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.card-header {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px 24px 0;
}

.header-icon {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 18px;
}

.header-icon.action-icon {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.header-icon.details-icon {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.header-text h3 {
  font-size: 16px;
  font-weight: 600;
  margin: 0 0 4px 0;
  color: #1f2937;
}

.header-text p {
  font-size: 12px;
  color: #6b7280;
  margin: 0;
}

.card-divider {
  margin: 16px 0;
}

.info-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 16px;
  padding: 0 24px 24px;
}

.info-item {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.info-label {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: #6b7280;
  font-weight: 500;
}

.label-icon {
  font-size: 14px;
  color: #9ca3af;
}

.info-value {
  font-size: 14px;
  color: #1f2937;
  font-weight: 500;
  word-break: break-all;
}

.order-id {
  font-family: 'Courier New', monospace;
  color: #667eea;
  font-weight: 600;
}

.amount {
  font-size: 16px;
  color: #f59e0b;
  font-weight: 600;
}

.status-tag {
  font-size: 12px;
}

.action-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
  padding: 0 24px 24px;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 16px;
  background: linear-gradient(135deg, #f5f7fa 0%, #e4e8ec 100%);
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 1px solid transparent;
}

.action-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border-color: #667eea;
}

.action-icon-item {
  font-size: 20px;
  color: #667eea;
}

.action-item span {
  font-size: 12px;
  color: #4b5563;
  font-weight: 500;
}

.action-item.delete-action:hover {
  background: linear-gradient(135deg, #fee2e2 0%, #fecaca 100%);
  border-color: #ef4444;
}

.action-item.delete-action .action-icon-item {
  color: #ef4444;
}

.details-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.details-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.item-table {
  padding: 0 24px 24px;
}

.part-id {
  font-family: 'Courier New', monospace;
  color: #667eea;
  font-weight: 600;
}

.quantity {
  font-weight: 600;
  color: #f59e0b;
}

.price,
.total {
  font-weight: 600;
  color: #10b981;
}

.total-amount {
  font-size: 16px;
  font-weight: 600;
  color: #f59e0b;
}

.remarks-content {
  padding: 0 24px 24px;
  white-space: pre-wrap;
  line-height: 1.6;
  color: #4b5563;
}

.timeline-wrapper {
  padding: 0 24px 24px;
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
  color: #1f2937;
}

.timeline-time {
  font-size: 12px;
  color: #9ca3af;
}

.timeline-content {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.timeline-desc {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #6b7280;
}

.timeline-desc .anticon {
  font-size: 14px;
  color: #9ca3af;
}

@media (max-width: 1200px) {
  .main-content {
    grid-template-columns: 1fr;
  }
  
  .info-container {
    max-width: 100%;
  }
}

@media (max-width: 768px) {
  .purchase-detail {
    padding: 16px;
  }
  
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
  
  .item-table :deep(.ant-table) {
    font-size: 12px;
  }
}
</style>
