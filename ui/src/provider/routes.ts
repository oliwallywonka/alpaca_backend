import type { RouteRecordRaw } from 'vue-router'

export const providerRoutes: RouteRecordRaw[] = [
  {
    path: '/providers',
    name: 'providers',
    component: () => import('@/provider/pages/ProviderPage.vue'),
  },
  {
    path: '/providers/:id/resources',
    name: 'provider-resources',
    component: () => import('@/provider/pages/ResourcePricesPage.vue')
  },
]
