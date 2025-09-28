import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { message } from 'ant-design-vue'
import { useUserStore } from '@/stores/user'
import { PermissionService } from '@/services/permission'
import { walletInfo } from '@/services/web3'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('@/layouts/BasicLayout.vue'),
    redirect: '/dashboard',
    children: [
      // 测试页面
      {
        path: 'test',
        name: 'TestPage',
        component: () => import('@/views/TestPage.vue'),
        meta: { title: '测试页面', hidden: true }
      },
      // 概览页面
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/EnhancedDashboard.vue'),
        meta: { title: '概览', icon: 'DashboardOutlined' }
      },
      // 简化Dashboard页面
      {
        path: 'dashboard-simple',
        name: 'SimpleDashboard',
        component: () => import('@/views/dashboard/NewTestDashboard.vue'),
        meta: { title: '简化Dashboard', hidden: true }
      },

      // 用户相关路由
      {
        path: 'user/profile',
        name: 'UserProfile',
        component: () => import('@/views/user/Profile.vue'),
        meta: { title: '个人资料', requiresAuth: true }
      },
      {
        path: 'user/management',
        name: 'UserManagement',
        component: () => import('@/views/user/Management.vue'),
        meta: { title: '用户管理', requiresRegulator: true }
      },

      // 项目相关路由
      {
        path: 'project/overview',
        name: 'ProjectOverview',
        component: () => import('@/views/project/Overview.vue'),
        meta: { title: '项目总览' }
      },
      {
        path: 'project/my',
        name: 'MyProjects',
        component: () => import('@/views/project/MyProjects.vue'),
        meta: { title: '我的项目', requiresProjectOwner: true }
      },
      {
        path: 'project/:id',
        name: 'ProjectDetail',
        component: () => import('@/views/project/Detail.vue'),
        meta: { title: '项目详情', hidden: true }
      },

      // 审核相关路由
      {
        path: 'audit/market',
        name: 'AuditMarket',
        component: () => import('@/views/audit/Market.vue'),
        meta: { title: '审核市场', requiresVerifier: true }
      },
      {
        path: 'audit/my',
        name: 'MyAudits',
        component: () => import('@/views/audit/MyAudits.vue'),
        meta: { title: '我的审核', requiresVerifier: true }
      },

      // 碳信用相关路由
      {
        path: 'credit/market',
        name: 'CreditMarket',
        component: () => import('@/views/credit/Market.vue'),
        meta: { title: '碳信用交易市场' }
      },
      {
        path: 'credit/my',
        name: 'MyCredits',
        component: () => import('@/views/credit/MyCredits.vue'),
        meta: { title: '我的碳信用', requiresProjectOwnerOrBuyer: true }
      },
      {
        path: 'credit/:id',
        name: 'CreditDetail',
        component: () => import('@/views/credit/Detail.vue'),
        meta: { title: '碳信用详情', hidden: true }
      },

      // 证书相关路由
      {
        path: 'certificate/overview',
        name: 'CertificateOverview',
        component: () => import('@/views/certificate/Overview.vue'),
        meta: { title: '证书总览' }
      },
      {
        path: 'certificate/:id',
        name: 'CertificateDetail',
        component: () => import('@/views/certificate/Detail.vue'),
        meta: { title: '证书详情', hidden: true }
      },

      // 兼容旧路由的重定向
      {
        path: 'users',
        redirect: '/user/management'
      },
      {
        path: 'projects',
        redirect: '/project/overview'
      },
      {
        path: 'credits',
        redirect: '/credit/market'
      },
      {
        path: 'certificates',
        redirect: '/certificate/overview'
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫（基于角色与连接状态的权限检查）
router.beforeEach((to, from, next) => {
  console.log('路由守卫：', to.path, to.meta)
  
  // 临时禁用所有权限检查，直接放行
  next()
  
  // const userStore = useUserStore()

  // const requiresAuth = Boolean(to.meta && (to.meta.requiresAuth || to.meta.requiresRegulator || to.meta.requiresProjectOwner || to.meta.requiresVerifier || to.meta.requiresProjectOwnerOrBuyer))

  // // 需要认证但未连接或无用户信息
  // if (requiresAuth) {
  //   if (!walletInfo.isConnected) {
  //     message.warning('请先连接钱包')
  //     return next('/dashboard')
  //   }
  //   if (!userStore.currentUser) {
  //     message.warning('请先完善个人资料')
  //     return next('/dashboard')
  //   }
  // }

  // // 角色细分校验
  // const user = userStore.currentUser
  // if (to.meta?.requiresRegulator && !PermissionService.isRegulator(user)) {
  //   message.error('无权限访问：需要监管者权限')
  //   return next('/dashboard')
  // }
  // if (to.meta?.requiresProjectOwner && !PermissionService.isProjectOwner(user)) {
  //   message.error('无权限访问：需要项目方权限')
  //   return next('/dashboard')
  // }
  // if (to.meta?.requiresVerifier && !PermissionService.isVerifier(user)) {
  //   message.error('无权限访问：需要审核员权限')
  //   return next('/dashboard')
  // }
  // if (to.meta?.requiresProjectOwnerOrBuyer && !(PermissionService.isProjectOwner(user) || PermissionService.isBuyer(user))) {
  //   message.error('无权限访问：需要项目方或购买方权限')
  //   return next('/dashboard')
  // }

  // next()
})

export default router
