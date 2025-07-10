import type { RouteRecordRaw } from 'vue-router';

export const dashboardRoutes: RouteRecordRaw[] = [
  {
    path: '/dashboard',
    name: 'dashboard',
    component: () => import('./DashboardPage.vue'),
  },
  {
    path: '/',
    name: 'dashboard',
    component: () => import('./DashboardPage.vue'),
  },
]