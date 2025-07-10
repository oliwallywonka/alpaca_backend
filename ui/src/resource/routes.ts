import type { RouteRecordRaw } from 'vue-router'

export const resourceRoutes: RouteRecordRaw[] = [
  {
    path: '/resources',
    name: 'resources',
    component: () => import('./ResourcePage.vue'),
  },
]
