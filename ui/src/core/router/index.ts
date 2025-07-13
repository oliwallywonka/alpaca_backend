import { createRouter, createWebHistory } from 'vue-router'
import AdminLayout from '@/core/layouts/AdminLayout.vue'
import { userRoutes } from '@/user/routes'
import { dashboardRoutes } from '@/dashboard/routes'
import { tourRoutes } from '@/tour/routes'
import { resourceRoutes } from '@/resource/routes'
import { destinationRoutes } from '@/destination/routes'
import { cashRoutes } from '@/cash/routes'
import { rolesRoutes } from '@/role/routes'
import { tripRoutes } from '@/trip/routes'
import { authRoutes } from '@/auth/routes'
import { providerRoutes } from '@/provider/routes'
import { API } from '../services/pocketbase'
import { customerRoutes } from '@/customer/route'

const appRoutes = [
  ...dashboardRoutes,
  ...userRoutes,
  ...providerRoutes,
  ...tourRoutes,
  ...tripRoutes,
  ...resourceRoutes,
  ...destinationRoutes,
  ...rolesRoutes,
  ...cashRoutes,
  ...customerRoutes,
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    ...authRoutes,
    {
      path: '/',
      name: 'home',
      component: AdminLayout,
      children: appRoutes,
    },
  ],
})

router.beforeEach((to) => {
  const canAccess = API.authStore.isValid
  if (!canAccess && to.name !== 'login') return '/login'
})

export default router
