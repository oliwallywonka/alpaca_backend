import type { RouteRecordRaw } from 'vue-router'

export const hotelRoutes: RouteRecordRaw[] = [
  {
    path: '/hotels',
    name: 'hotels',
    component: () => import('./HotelsPage.vue'),
  },
]
