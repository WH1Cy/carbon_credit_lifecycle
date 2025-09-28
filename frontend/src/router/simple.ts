import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/dashboard-simple'
  },
  {
    path: '/test',
    name: 'TestPage',
    component: () => import('@/views/TestPage.vue'),
    meta: { title: '测试页面' }
  },
  {
    path: '/dashboard-simple',
    name: 'SimpleDashboard',
    component: () => import('@/views/dashboard/NewTestDashboard.vue'),
    meta: { title: '简化Dashboard' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
