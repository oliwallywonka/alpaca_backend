import type { RouteRecordRaw } from 'vue-router'

export const destinationRoutes: RouteRecordRaw[] = [
  {
    path: '/destinations',
    name: 'destinations',
    component: () => import('./DestinationPage.vue'),
  },
]
