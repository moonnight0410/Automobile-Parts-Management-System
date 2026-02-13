/**
 * 供应链相关类型定义
 * 定义供应链管理模块的数据类型和接口
 */

// 供应链订单
export interface SupplyOrder {
  orderID: string      // 订单ID
  buyer: string        // 采购方ID（车企）
  seller: string       // 销售方ID（零部件厂商）
  partID: string       // 零部件ID
  quantity: number     // 数量
  batchNo: string      // 涉及批次
  agreedTime: string   // 约定交付时间
  status: string       // 状态（待发货/已发货/已签收）
  createTime: string   // 订单创建时间
}

// 物流数据
export interface LogisticsData {
  logisticsID: string  // 物流记录ID
  orderID: string      // 关联订单ID
  carrier: string      // 物流商ID
  startTime: string    // 出发时间
  endTime: string      // 送达时间
  gpsHash: string      // GPS轨迹数据哈希
  receiver: string     // 签收人ID
}

// 对账信息
export interface Reconciliation {
  reconID: string      // 对账ID
  orderID: string      // 关联订单
  amount: string       // 金额
  status: string       // 状态（待对账/已确认/已结算）
  settleTime: string   // 结算时间
}

// 供应链阶段
export interface SupplyChainStage {
  chainID: string          // 供应链环节唯一ID
  partID: string           // 关联零部件
  batchNo: string          // 关联批次
  orderID: string          // 关联采购订单
  logisticsID: string      // 关联物流记录
  stageType: string        // 环节类型
  stageStatus: string      // 环节状态
  participator: string     // 环节参与方ID
  participatorRole: string // 参与方角色
  quantity: number         // 环节涉及数量
  operateTime: string      // 环节操作时间
  operator: string         // 操作人ID
  remark: string           // 环节备注
}

// 创建供应链订单请求
export interface CreateSupplyOrderRequest {
  orderID: string
  buyer: string
  seller: string
  partID: string
  quantity: number
  batchNo: string
  agreedTime: string
  status?: string
}

// 创建物流数据请求
export interface CreateLogisticsDataRequest {
  logisticsID: string
  orderID: string
  carrier: string
  startTime?: string
  endTime?: string
  gpsHash?: string
  receiver?: string
}
