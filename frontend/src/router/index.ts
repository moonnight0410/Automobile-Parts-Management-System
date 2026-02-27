/**
 * Vue Router 路由配置
 * 定义应用的路由规则和导航守卫
 */

import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import AuthMiddleware from '../middleware/auth'
import { UserRole } from '../types'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: {
      requiresAuth: false,
      title: '登录'
    }
  },
  
  {
    path: '/',
    name: 'Root',
    component: () => import('../layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        redirect: '/dashboard'
      },
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('../views/Dashboard.vue'),
        meta: {
          requiresAuth: true,
          title: '仪表盘',
          icon: 'DashboardOutlined',
          permission: 'dashboard',
          roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE]
        }
      },
      
      {
        path: 'parts',
        name: 'Parts',
        redirect: '/parts/list',
        meta: {
          title: '零部件管理',
          icon: 'AppstoreOutlined',
          roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE]
        },
        children: [
          {
            path: 'list',
            name: 'PartList',
            component: () => import('../views/parts/PartList.vue'),
            meta: {
              requiresAuth: true,
              title: '零部件列表',
              permission: 'partsList',
              roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER]
            }
          },
          {
            path: 'create',
            name: 'PartCreate',
            component: () => import('../views/parts/PartCreate.vue'),
            meta: {
              requiresAuth: true,
              title: '创建零部件',
              permission: 'partsCreate',
              roles: [UserRole.MANUFACTURER]
            }
          },
          {
            path: 'detail/:id',
            name: 'PartDetail',
            component: () => import('../views/parts/PartDetail.vue'),
            meta: {
              requiresAuth: true,
              title: '零部件详情',
              permission: 'partsDetail',
              roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE]
            }
          }
        ]
      },
      
      {
        path: 'bom',
        name: 'BOM',
        redirect: '/bom/list',
        meta: {
          title: 'BOM管理',
          icon: 'FileTextOutlined',
          roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER]
        },
        children: [
          {
            path: 'list',
            name: 'BOMList',
            component: () => import('../views/bom/BOMList.vue'),
            meta: {
              requiresAuth: true,
              title: 'BOM列表',
              permission: 'bomList',
              roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER]
            }
          },
          {
            path: 'create',
            name: 'BOMCreate',
            component: () => import('../views/bom/BOMCreate.vue'),
            meta: {
              requiresAuth: true,
              title: '创建BOM',
              permission: 'bomCreate',
              roles: [UserRole.MANUFACTURER]
            }
          },
          {
            path: 'compare',
            name: 'BOMCompare',
            component: () => import('../views/bom/BOMCompare.vue'),
            meta: {
              requiresAuth: true,
              title: 'BOM比较',
              permission: 'bomCompare',
              roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER]
            }
          },
          {
            path: 'detail/:id',
            name: 'BOMDetail',
            component: () => import('../views/boms/BOMDetail.vue'),
            meta: {
              requiresAuth: true,
              title: 'BOM详情',
              permission: 'bomDetail',
              roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER]
            }
          },
          {
            path: 'edit/:id',
            name: 'BOMEdit',
            component: () => import('../views/boms/BOMEdit.vue'),
            meta: {
              requiresAuth: true,
              title: '编辑BOM',
              permission: 'bomEdit',
              roles: [UserRole.MANUFACTURER]
            }
          }
        ]
      },
      
      {
        path: 'production',
        name: 'Production',
        redirect: '/production/data',
        meta: {
          title: '生产管理',
          icon: 'ToolOutlined',
          roles: [UserRole.MANUFACTURER]
        },
        children: [
          {
            path: 'data',
            name: 'ProductionData',
            component: () => import('../views/production/ProductionData.vue'),
            meta: {
              requiresAuth: true,
              title: '生产数据',
              permission: 'productionData',
              roles: [UserRole.MANUFACTURER]
            }
          },
          {
            path: 'data/create',
            name: 'ProductionDataCreate',
            component: () => import('../views/production/ProductionDataCreate.vue'),
            meta: {
              requiresAuth: true,
              title: '录入生产数据',
              permission: 'productionDataCreate',
              roles: [UserRole.MANUFACTURER]
            }
          },
          {
            path: 'quality',
            name: 'QualityInspection',
            component: () => import('../views/production/QualityInspection.vue'),
            meta: {
              requiresAuth: true,
              title: '质检管理',
              permission: 'qualityInspection',
              roles: [UserRole.MANUFACTURER]
            }
          },
          {
            path: 'quality/create',
            name: 'QualityInspectionCreate',
            component: () => import('../views/production/QualityInspectionCreate.vue'),
            meta: {
              requiresAuth: true,
              title: '录入质检数据',
              permission: 'qualityInspectionCreate',
              roles: [UserRole.MANUFACTURER]
            }
          },
          {
            path: 'detail/:id',
            name: 'ProductionDetail',
            component: () => import('../views/production/ProductionDetail.vue'),
            meta: {
              requiresAuth: true,
              title: '生产记录详情',
              permission: 'productionDetail',
              roles: [UserRole.MANUFACTURER]
            }
          },
          {
            path: 'quality/detail/:id',
            name: 'QualityDetail',
            component: () => import('../views/quality/QualityDetail.vue'),
            meta: {
              requiresAuth: true,
              title: '质检记录详情',
              permission: 'qualityDetail',
              roles: [UserRole.MANUFACTURER]
            }
          }
        ]
      },
      
      {
        path: 'supply',
        name: 'Supply',
        redirect: '/supply/orders',
        meta: {
          title: '供应链管理',
          icon: 'SwapOutlined',
          roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER]
        },
        children: [
          {
            path: 'orders',
            name: 'SupplyOrders',
            component: () => import('../views/supply/SupplyOrders.vue'),
            meta: {
              requiresAuth: true,
              title: '采购订单',
              permission: 'supplyOrders',
              roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER]
            }
          },
          {
            path: 'order/create',
            name: 'SupplyOrderCreate',
            component: () => import('../views/supply/SupplyOrderCreate.vue'),
            meta: {
              requiresAuth: true,
              title: '创建采购订单',
              permission: 'supplyOrderCreate',
              roles: [UserRole.AUTOMAKER]
            }
          },
          {
            path: 'order/detail/:id',
            name: 'PurchaseDetail',
            component: () => import('../views/orders/PurchaseDetail.vue'),
            meta: {
              requiresAuth: true,
              title: '采购订单详情',
              permission: 'purchaseDetail',
              roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER]
            }
          },
          {
            path: 'logistics',
            name: 'Logistics',
            component: () => import('../views/supply/Logistics.vue'),
            meta: {
              requiresAuth: true,
              title: '物流跟踪',
              permission: 'logistics',
              roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER]
            }
          },
          {
            path: 'logistics/detail/:id',
            name: 'LogisticsDetail',
            component: () => import('../views/logistics/LogisticsDetail.vue'),
            meta: {
              requiresAuth: true,
              title: '物流跟踪详情',
              permission: 'logisticsDetail',
              roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER]
            }
          }
        ]
      },
      
      {
        path: 'aftersale',
        name: 'Aftersale',
        redirect: '/aftersale/fault',
        meta: {
          title: '售后管理',
          icon: 'CustomerServiceOutlined',
          roles: [UserRole.AUTOMAKER, UserRole.AFTERSALE]
        },
        children: [
          {
            path: 'fault',
            name: 'FaultReport',
            component: () => import('../views/aftersale/FaultReport.vue'),
            meta: {
              requiresAuth: true,
              title: '故障报告',
              permission: 'faultReport',
              roles: [UserRole.AUTOMAKER, UserRole.AFTERSALE]
            }
          },
          {
            path: 'fault/create',
            name: 'FaultReportCreate',
            component: () => import('../views/aftersale/FaultReportCreate.vue'),
            meta: {
              requiresAuth: true,
              title: '上报故障',
              permission: 'faultReport',
              roles: [UserRole.AUTOMAKER, UserRole.AFTERSALE]
            }
          },
          {
            path: 'fault/detail/:id',
            name: 'FaultReportDetail',
            component: () => import('../views/aftersale/FaultReportDetail.vue'),
            meta: {
              requiresAuth: true,
              title: '故障记录详情',
              permission: 'faultReportDetail',
              roles: [UserRole.AUTOMAKER, UserRole.AFTERSALE]
            }
          },
          {
            path: 'recall',
            name: 'RecallRecord',
            component: () => import('../views/aftersale/RecallRecord.vue'),
            meta: {
              requiresAuth: true,
              title: '召回记录',
              permission: 'recallRecord',
              roles: [UserRole.AUTOMAKER, UserRole.AFTERSALE]
            }
          },
          {
            path: 'recall/create',
            name: 'RecallRecordCreate',
            component: () => import('../views/aftersale/RecallRecordCreate.vue'),
            meta: {
              requiresAuth: true,
              title: '发起召回',
              permission: 'recallRecord',
              roles: [UserRole.AUTOMAKER, UserRole.AFTERSALE]
            }
          },
          {
            path: 'recall/detail/:id',
            name: 'RecallRecordDetail',
            component: () => import('../views/aftersale/RecallRecordDetail.vue'),
            meta: {
              requiresAuth: true,
              title: '召回记录详情',
              permission: 'recallRecordDetail',
              roles: [UserRole.AUTOMAKER, UserRole.AFTERSALE]
            }
          },
          {
            path: 'assistant',
            name: 'AIAssistant',
            component: () => import('../views/aftersale/AIAssistant.vue'),
            meta: {
              requiresAuth: true,
              title: '智能售后助手',
              permission: 'aiAssistant',
              roles: [UserRole.AUTOMAKER, UserRole.AFTERSALE]
            }
          }
        ]
      },
      
      {
        path: 'blockchain',
        name: 'Blockchain',
        component: () => import('../views/Blockchain.vue'),
        meta: {
          requiresAuth: true,
          title: '区块链浏览器',
          icon: 'LinkOutlined',
          permission: 'blockchain',
          roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE]
        }
      },
      
      {
        path: 'system',
        name: 'System',
        redirect: '/system/users',
        meta: {
          title: '系统管理',
          icon: 'SettingOutlined',
          roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE]
        },
        children: [
          {
            path: 'users',
            name: 'UserManagement',
            component: () => import('../views/system/UserManagement.vue'),
            meta: {
              requiresAuth: true,
              title: '用户管理',
              permission: 'userManagement',
              roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE]
            }
          },
          {
            path: 'settings',
            name: 'SystemSettings',
            component: () => import('../views/system/SystemSettings.vue'),
            meta: {
              requiresAuth: true,
              title: '系统设置',
              permission: 'systemSettings',
              roles: [UserRole.MANUFACTURER, UserRole.AUTOMAKER, UserRole.AFTERSALE]
            }
          }
        ]
      }
    ]
  },
  
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('../views/NotFound.vue'),
    meta: {
      requiresAuth: false,
      title: '页面未找到'
    }
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = `${to.meta.title} - 汽车零部件区块链管理系统`
  }
  
  next()
})

router.beforeEach(AuthMiddleware)

export default router
