/**
 * 常量定义文件
 * 定义系统中使用的常量
 */

// API基础URL
export const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

// 本地存储键名
export const STORAGE_KEYS = {
  TOKEN: 'token',
  USER: 'user',
  THEME: 'theme',
  LANGUAGE: 'language'
} as const

// 零部件状态
export const PART_STATUS = {
  NORMAL: 'NORMAL',                    // 正常
  IN_PRODUCTION: 'IN_PRODUCTION',      // 在产
  QUALITY_INSPECTED: 'QUALITY_INSPECTED', // 质检完成
  IN_SUPPLY_CHAIN: 'IN_SUPPLY_CHAIN',  // 供应链中
  RECALLED: 'RECALLED',                // 已召回
  SCRAPPED: 'SCRAPPED'                 // 已报废
} as const

// BOM状态
export const BOM_STATUS = {
  DRAFT: '草稿',
  PUBLISHED: '已发布',
  CHANGED: '已变更',
  OBSOLETE: '已作废'
} as const

// BOM类型
export const BOM_TYPE = {
  RD: '研发BOM',
  PRODUCTION: '生产BOM'
} as const

// 供应链订单状态
export const ORDER_STATUS = {
  PENDING: '待发货',
  SHIPPED: '已发货',
  RECEIVED: '已签收'
} as const

// 故障报告状态
export const FAULT_STATUS = {
  PENDING: '待审核',
  REVIEWED: '已审核',
  RECALLED: '已召回'
} as const

// 召回记录状态
export const RECALL_STATUS = {
  PENDING: '待通知',
  NOTIFYING: '通知中',
  COMPLETED: '已完成'
} as const

// 售后类型
export const AFTERSALE_TYPE = {
  FAULT_REPORT: '故障报修',
  REPAIR: '维修处理',
  RECALL: '召回登记',
  SCRAP: '报废处理'
} as const

// 售后状态
export const AFTERSALE_STATUS = {
  PENDING: '待审核',
  PROCESSING: '处理中',
  COMPLETED: '已完成'
} as const

// 供应链环节类型
export const STAGE_TYPE = {
  ORDER: '采购下单',
  SHIP: '物流发货',
  WAREHOUSE_IN: '仓库入库',
  WAREHOUSE_OUT: '仓库出库',
  RECEIVE: '车企签收'
} as const

// 用户角色名称
export const ROLE_NAMES = {
  manufacturer: '制造商',
  automaker: '车企',
  aftersale: '售后中心'
} as const

// 分页配置
export const PAGINATION_CONFIG = {
  DEFAULT_PAGE: 1,
  DEFAULT_PAGE_SIZE: 10,
  PAGE_SIZE_OPTIONS: ['10', '20', '50', '100']
} as const
