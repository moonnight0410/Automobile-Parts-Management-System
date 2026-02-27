/**
 * 权限配置文件
 * 集中管理所有页面权限
 */

import { UserRole } from '../types'

/**
 * 页面权限配置
 * key: 页面标识（与路由meta.permission对应）
 * value: 允许访问的角色列表
 */
export const PAGE_PERMISSIONS: Record<string, UserRole[]> = {
  // 仪表盘
  dashboard: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE],

  // 零部件管理
  partsList: [UserRole.MANUFACTURER, UserRole.AUTOMAKER],
  partsCreate: [UserRole.MANUFACTURER],
  partsDetail: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE],

  // BOM管理
  bomList: [UserRole.MANUFACTURER, UserRole.AUTOMAKER],
  bomCreate: [UserRole.MANUFACTURER],
  bomCompare: [UserRole.MANUFACTURER, UserRole.AUTOMAKER],

  // 生产管理
  productionData: [UserRole.MANUFACTURER],
  productionDataCreate: [UserRole.MANUFACTURER],
  qualityInspection: [UserRole.MANUFACTURER],
  qualityInspectionCreate: [UserRole.MANUFACTURER],

  // 供应链管理
  supplyOrders: [UserRole.MANUFACTURER, UserRole.AUTOMAKER],
  supplyOrderCreate: [UserRole.AUTOMAKER],
  logistics: [UserRole.MANUFACTURER, UserRole.AUTOMAKER],
  logisticsCreate: [UserRole.MANUFACTURER],

  // 售后管理
  faultReport: [UserRole.AUTOMAKER, UserRole.AFTERSALE],
  recallRecord: [UserRole.AUTOMAKER, UserRole.AFTERSALE],
  aiAssistant: [UserRole.AUTOMAKER, UserRole.AFTERSALE],

  // 区块链浏览器
  blockchain: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE],

  // 系统管理
  userManagement: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE],
  systemSettings: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE],
}

/**
 * 菜单配置
 * 用于动态生成导航菜单
 */
export interface MenuConfig {
  key: string
  title: string
  icon?: string
  permission?: string
  children?: MenuConfig[]
}

/**
 * 菜单权限配置
 * 定义左侧导航菜单结构及权限
 */
export const MENU_CONFIG: MenuConfig[] = [
  {
    key: '/dashboard',
    title: '仪表盘',
    icon: 'DashboardOutlined',
    permission: 'dashboard',
  },
  {
    key: '/parts',
    title: '零部件管理',
    icon: 'AppstoreOutlined',
    children: [
      { key: '/parts/list', title: '零部件列表', permission: 'partsList' },
      { key: '/parts/create', title: '创建零部件', permission: 'partsCreate' },
    ],
  },
  {
    key: '/bom',
    title: 'BOM管理',
    icon: 'FileTextOutlined',
    children: [
      { key: '/bom/list', title: 'BOM列表', permission: 'bomList' },
      { key: '/bom/create', title: '创建BOM', permission: 'bomCreate' },
      { key: '/bom/compare', title: 'BOM比较', permission: 'bomCompare' },
    ],
  },
  {
    key: '/production',
    title: '生产管理',
    icon: 'ToolOutlined',
    children: [
      { key: '/production/data', title: '生产数据', permission: 'productionData' },
      { key: '/production/quality', title: '质检管理', permission: 'qualityInspection' },
    ],
  },
  {
    key: '/supply',
    title: '供应链管理',
    icon: 'SwapOutlined',
    children: [
      { key: '/supply/orders', title: '采购订单', permission: 'supplyOrders' },
      { key: '/supply/order/create', title: '创建采购订单', permission: 'supplyOrderCreate' },
      { key: '/supply/logistics', title: '物流跟踪', permission: 'logistics' },
      { key: '/supply/logistics/create', title: '添加物流记录', permission: 'logisticsCreate' },
    ],
  },
  {
    key: '/aftersale',
    title: '售后管理',
    icon: 'CustomerServiceOutlined',
    children: [
      { key: '/aftersale/fault', title: '故障报告', permission: 'faultReport' },
      { key: '/aftersale/recall', title: '召回记录', permission: 'recallRecord' },
      { key: '/aftersale/assistant', title: '智能售后助手', permission: 'aiAssistant' },
    ],
  },
  {
    key: '/blockchain',
    title: '区块链浏览器',
    icon: 'LinkOutlined',
    permission: 'blockchain',
  },
  {
    key: '/system',
    title: '系统管理',
    icon: 'SettingOutlined',
    children: [
      { key: '/system/users', title: '用户管理', permission: 'userManagement' },
      { key: '/system/settings', title: '系统设置', permission: 'systemSettings' },
    ],
  },
]

/**
 * 路由路径到权限key的映射
 */
export const ROUTE_PERMISSION_MAP: Record<string, string> = {
  '/dashboard': 'dashboard',
  '/parts/list': 'partsList',
  '/parts/create': 'partsCreate',
  '/parts/detail': 'partsDetail',
  '/bom/list': 'bomList',
  '/bom/create': 'bomCreate',
  '/bom/compare': 'bomCompare',
  '/production/data': 'productionData',
  '/production/data/create': 'productionDataCreate',
  '/production/quality': 'qualityInspection',
  '/production/quality/create': 'qualityInspectionCreate',
  '/supply/orders': 'supplyOrders',
  '/supply/order/create': 'supplyOrderCreate',
  '/supply/logistics': 'logistics',
  '/supply/logistics/create': 'logisticsCreate',
  '/aftersale/fault': 'faultReport',
  '/aftersale/recall': 'recallRecord',
  '/aftersale/assistant': 'aiAssistant',
  '/blockchain': 'blockchain',
  '/system/users': 'userManagement',
  '/system/settings': 'systemSettings',
}
