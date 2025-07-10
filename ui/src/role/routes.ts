import type { RouteRecordRaw } from 'vue-router'

export const rolesRoutes: RouteRecordRaw[] = [
  {
    path: '/roles',
    name: 'roles',
    component: () => import('./RolesPage.vue'),
  },
]
