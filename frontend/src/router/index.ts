/**
 * Vue Router 路由配置
 * 定义应用的路由规则和导航守卫
 */

import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import AuthMiddleware from '../middleware/auth'

// ==================== 路由配置 ====================

const routes: Array<RouteRecordRaw> = [
  // 登录页面
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: {
      requiresAuth: false,
      title: '登录'
    }
  },
  
  // 主布局
  {
    path: '/',
    name: 'Root',
    component: () => import('../layouts/MainLayout.vue'),
    beforeEnter: AuthMiddleware,
    children: [
      // 仪表盘
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
          roles: ['manufacturer', 'automaker', 'aftersale']
        }
      },
      
      // 零部件管理
      {
        path: 'parts',
        name: 'Parts',
        redirect: '/parts/list',
        meta: {
          title: '零部件管理',
          icon: 'AppstoreOutlined',
          roles: ['manufacturer', 'automaker', 'aftersale']
        },
        children: [
          {
            path: 'list',
            name: 'PartList',
            component: () => import('../views/parts/PartList.vue'),
            meta: {
              requiresAuth: true,
              title: '零部件列表',
              roles: ['manufacturer', 'automaker', 'aftersale']
            }
          },
          {
            path: 'create',
            name: 'PartCreate',
            component: () => import('../views/parts/PartCreate.vue'),
            meta: {
              requiresAuth: true,
              title: '创建零部件',
              roles: ['manufacturer']
            }
          },
          {
            path: 'detail/:id',
            name: 'PartDetail',
            component: () => import('../views/parts/PartDetail.vue'),
            meta: {
              requiresAuth: true,
              title: '零部件详情',
              roles: ['manufacturer', 'automaker', 'aftersale']
            }
          }
        ]
      },
      
      // BOM管理
      {
        path: 'bom',
        name: 'BOM',
        redirect: '/bom/list',
        meta: {
          title: 'BOM管理',
          icon: 'FileTextOutlined',
          roles: ['manufacturer', 'automaker']
        },
        children: [
          {
            path: 'list',
            name: 'BOMList',
            component: () => import('../views/bom/BOMList.vue'),
            meta: {
              requiresAuth: true,
              title: 'BOM列表',
              roles: ['manufacturer', 'automaker']
            }
          },
          {
            path: 'create',
            name: 'BOMCreate',
            component: () => import('../views/bom/BOMCreate.vue'),
            meta: {
              requiresAuth: true,
              title: '创建BOM',
              roles: ['manufacturer']
            }
          },
          {
            path: 'compare',
            name: 'BOMCompare',
            component: () => import('../views/bom/BOMCompare.vue'),
            meta: {
              requiresAuth: true,
              title: 'BOM比较',
              roles: ['manufacturer', 'automaker']
            }
          }
        ]
      },
      
      // 生产管理
      {
        path: 'production',
        name: 'Production',
        redirect: '/production/data',
        meta: {
          title: '生产管理',
          icon: 'ToolOutlined',
          roles: ['manufacturer']
        },
        children: [
          {
            path: 'data',
            name: 'ProductionData',
            component: () => import('../views/production/ProductionData.vue'),
            meta: {
              requiresAuth: true,
              title: '生产数据',
              roles: ['manufacturer']
            }
          },
          {
            path: 'quality',
            name: 'QualityInspection',
            component: () => import('../views/production/QualityInspection.vue'),
            meta: {
              requiresAuth: true,
              title: '质检管理',
              roles: ['manufacturer']
            }
          }
        ]
      },
      
      // 供应链管理
      {
        path: 'supply',
        name: 'Supply',
        redirect: '/supply/orders',
        meta: {
          title: '供应链管理',
          icon: 'SwapOutlined',
          roles: ['manufacturer', 'automaker']
        },
        children: [
          {
            path: 'orders',
            name: 'SupplyOrders',
            component: () => import('../views/supply/SupplyOrders.vue'),
            meta: {
              requiresAuth: true,
              title: '采购订单',
              roles: ['manufacturer', 'automaker']
            }
          },
          {
            path: 'logistics',
            name: 'Logistics',
            component: () => import('../views/supply/Logistics.vue'),
            meta: {
              requiresAuth: true,
              title: '物流跟踪',
              roles: ['manufacturer', 'automaker']
            }
          }
        ]
      },
      
      // 售后管理
      {
        path: 'aftersale',
        name: 'Aftersale',
        redirect: '/aftersale/fault',
        meta: {
          title: '售后管理',
          icon: 'CustomerServiceOutlined',
          roles: ['aftersale', 'automaker']
        },
        children: [
          {
            path: 'fault',
            name: 'FaultReport',
            component: () => import('../views/aftersale/FaultReport.vue'),
            meta: {
              requiresAuth: true,
              title: '故障报告',
              roles: ['aftersale', 'automaker']
            }
          },
          {
            path: 'recall',
            name: 'RecallRecord',
            component: () => import('../views/aftersale/RecallRecord.vue'),
            meta: {
              requiresAuth: true,
              title: '召回记录',
              roles: ['aftersale', 'automaker']
            }
          },
          {
            path: 'assistant',
            name: 'AIAssistant',
            component: () => import('../views/aftersale/AIAssistant.vue'),
            meta: {
              requiresAuth: true,
              title: '智能售后助手',
              roles: ['aftersale', 'automaker']
            }
          }
        ]
      },
      
      // 区块链浏览器
      {
        path: 'blockchain',
        name: 'Blockchain',
        component: () => import('../views/Blockchain.vue'),
        meta: {
          requiresAuth: true,
          title: '区块链浏览器',
          icon: 'LinkOutlined',
          roles: ['manufacturer', 'automaker', 'aftersale']
        }
      },
      
      // 系统管理
      {
        path: 'system',
        name: 'System',
        redirect: '/system/users',
        meta: {
          title: '系统管理',
          icon: 'SettingOutlined',
          roles: ['manufacturer', 'automaker', 'aftersale']
        },
        children: [
          {
            path: 'users',
            name: 'UserManagement',
            component: () => import('../views/system/UserManagement.vue'),
            meta: {
              requiresAuth: true,
              title: '用户管理',
              roles: ['manufacturer', 'automaker', 'aftersale']
            }
          },
          {
            path: 'settings',
            name: 'SystemSettings',
            component: () => import('../views/system/SystemSettings.vue'),
            meta: {
              requiresAuth: true,
              title: '系统设置',
              roles: ['manufacturer', 'automaker', 'aftersale']
            }
          }
        ]
      }
    ]
  },
  
  // 404页面
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

// ==================== 创建路由实例 ====================

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

// ==================== 全局路由守卫 ====================

router.beforeEach((to, from, next) => {
  // 设置页面标题
  if (to.meta.title) {
    document.title = `${to.meta.title} - 汽车零部件区块链管理系统`
  }
  
  next()
})

export default router
