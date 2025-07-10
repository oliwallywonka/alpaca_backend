import type { RouteRecordRaw } from 'vue-router'

export const tourRoutes: RouteRecordRaw[] = [
  {
    path: '/tours',
    name: 'tour',
    component: () => import('./pages/ToursPage.vue'),
  },
  {
    path: '/tours/:tourID',
    name: 'tour-edit',
    component: () => import('./pages/TourDetailPage.vue'),
  },
]
