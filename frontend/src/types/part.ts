/**
 * 零部件相关类型定义
 * 定义零部件管理模块的数据类型和接口
 */

// 零部件信息
export interface Part {
  partID: string        // 零部件唯一标识
  vin: string           // 关联车辆VIN码
  batchNo: string       // 生产批次号
  name: string          // 零部件名称
  type: string          // 类型
  manufacturer: string  // 生产厂商ID
  createTime: string    // 创建时间戳
  status: string        // 状态（在产/在售/召回/报废等）
}

// 零部件生命周期
export interface PartLifecycle {
  partID: string
  bomInfo?: BOMReference
  productionInfo?: ProductionData
  qualityInfo?: QualityInspection
  supplyChainInfo?: SupplyChainData[]
  aftersaleInfo?: AftersaleRecord[]
}

// BOM引用信息
export interface BOMReference {
  bomID: string
  version: string
  type: string
}

// 生产数据
export interface ProductionData {
  productionID: string
  partID: string
  batchNo: string
  params: Record<string, string>
  productionLine: string
  operator: string
  finishTime: string
}

// 质检数据
export interface QualityInspection {
  inspectionID: string
  partID: string
  batchNo: string
  indicators: Record<string, string>
  result: string
  handler: string
  handleTime: string
  disposal: string
}

// 供应链数据
export interface SupplyChainData {
  chainID: string
  partID: string
  batchNo: string
  orderID: string
  logisticsID: string
  stageType: string
  stageStatus: string
  participator: string
  participatorRole: string
  quantity: number
  operateTime: string
  operator: string
  remark: string
}

// 售后记录
export interface AftersaleRecord {
  aftersaleID: string
  partID: string
  batchNo: string
  vin: string
  faultID: string
  recallID: string
  aftersaleType: string
  aftersaleStatus: string
  handlerOrgID: string
  handlerID: string
  description: string
  processResult: string
  processTime: string
  cost: string
  remark: string
}

// 创建零部件请求
export interface CreatePartRequest {
  partID: string
  vin?: string
  batchNo: string
  name: string
  type: string
  manufacturer: string
  status?: string
}

// 更新零部件状态请求
export interface UpdatePartStatusRequest {
  partID: string
  status: string
}
