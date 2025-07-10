import type { RouteRecordRaw } from 'vue-router'

export const mealRoutes: RouteRecordRaw[] = [
  {
    path: '/activities',
    name: 'activities',
    component: () => import('./MealsPage.vue'),
  },
]
