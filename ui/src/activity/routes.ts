import type { RouteRecordRaw } from 'vue-router'

export const activityRoutes: RouteRecordRaw[] = [
  {
    path: '/activities',
    name: 'activities',
    component: () => import('./ActivitiesPage.vue'),
  },
]
