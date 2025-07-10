import type { RouteRecordRaw } from 'vue-router'

export const cashRoutes: RouteRecordRaw[] = [
  {
    path: '/cash-registers',
    name: 'cash-registers',
    component: () => import('./pages/CashRegistersPage.vue'),
  },
  {
    path: '/cash-registers/:registerID/movements',
    name: 'movements',
    component: () => import('./pages/MovementsPage.vue'),
  },
]
