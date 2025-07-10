import type { RouteRecordRaw } from 'vue-router'

export const userRoutes: RouteRecordRaw[] = [
  {
    path: '/users',
    name: 'users',
    component: () => import('@/user/pages/UsersPage.vue'),
  },
  {
    path: '/users/:id/resources',
    name: 'users-resources',
    component: () => import('@/user/pages/UserResources.vue'),
  }
]
