/**
 * BOM（物料清单）相关类型定义
 * 定义BOM管理模块的数据类型和接口
 */

// 物料明细项
export interface MaterialItem {
  materialID: string     // 物料唯一ID
  materialName: string   // 物料名称
  spec: string           // 规格型号
  quantity: number       // 数量
  supplierID: string     // 供应商ID
}

// BOM信息
export interface BOM {
  bomID: string              // BOM唯一ID
  bomType: string            // BOM类型：研发BOM/生产BOM
  productModel: string       // 对应车型
  partBatchNo: string        // 零部件批次号
  version: string            // BOM版本号
  creator: string            // 创建人
  createTime: string         // 创建时间
  status: string             // 状态：草稿/已发布/已变更/已作废
  materialList: MaterialItem[] // 物料明细列表
}

// BOM比较请求
export interface BOMCompareRequest {
  productionBOMID: string
  rdBOMID: string
}

// BOM变更请求
export interface BOMChangeRequest {
  bomID: string
  changeID: string
  changeReason: string
  oldVersion: string
  newVersion: string
  changeTime: string
}

// 创建BOM请求
export interface CreateBOMRequest {
  bomID: string
  bomType: string
  productModel: string
  partBatchNo: string
  version: string
  creator: string
  status?: string
  materialList: MaterialItem[]
}

// 更新BOM请求
export interface UpdateBOMRequest extends CreateBOMRequest {}
