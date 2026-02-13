/**
 * 售后相关类型定义
 * 定义售后管理模块的数据类型和接口
 */

// 故障报告
export interface FaultReport {
  faultID: string      // 故障ID
  partID: string       // 关联零部件
  vin: string          // 车辆VIN码
  reporter: string     // 上报者ID
  description: string  // 故障描述
  faultType: string    // AI分类结果
  riskProb: number     // 风险概率
  reportTime: string   // 上报时间
  status: string       // 处理状态
}

// 召回记录
export interface RecallRecord {
  recallID: string       // 召回ID
  batchNos: string[]     // 涉及批次列表
  reason: string         // 召回原因
  affectedParts: string[] // 受影响零部件ID列表
  status: string         // 进度
  createTime: string     // 召回发起时间
  finishTime: string     // 完成时间
}

// 售后记录
export interface AftersaleRecord {
  aftersaleID: string      // 售后记录唯一ID
  partID: string           // 关联零部件
  batchNo: string          // 关联批次
  vin: string              // 关联车辆
  faultID: string          // 关联故障上报
  recallID: string         // 关联召回记录
  aftersaleType: string    // 售后类型
  aftersaleStatus: string  // 处理状态
  handlerOrgID: string     // 处理机构ID
  handlerID: string        // 处理人ID
  description: string      // 售后描述
  processResult: string    // 处理结果
  processTime: string      // 处理完成时间
  cost: string             // 售后成本
  remark: string           // 售后备注
}

// 创建故障报告请求
export interface CreateFaultReportRequest {
  faultID: string
  partID: string
  vin: string
  reporter: string
  description: string
  faultType?: string
  riskProb?: number
  status?: string
}

// 创建召回记录请求
export interface CreateRecallRecordRequest {
  recallID: string
  batchNos: string[]
  reason: string
  affectedParts: string[]
  status?: string
}

// 创建售后记录请求
export interface CreateAftersaleRecordRequest {
  aftersaleID: string
  partID: string
  batchNo: string
  vin: string
  faultID?: string
  recallID?: string
  aftersaleType: string
  aftersaleStatus?: string
  handlerOrgID: string
  handlerID: string
  description: string
  processResult?: string
  cost?: string
  remark?: string
}
