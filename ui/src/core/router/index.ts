import { createRouter, createWebHistory } from 'vue-router'
import AdminLayout from '@/core/layouts/AdminLayout.vue'
import { userRoutes } from '@/user/routes'
import { dashboardRoutes } from '@/dashboard/routes'
import { tourRoutes } from '@/tour/routes'
import { resourceRoutes } from '@/resource/routes'
import { destinationRoutes } from '@/destination/routes'
import { activityRoutes } from '@/activity/routes'
import { cashRoutes } from '@/cash/routes'
import { hotelRoutes } from '@/hotel/routes'
import { mealRoutes } from '@/meal/routes'
import { rolesRoutes } from '@/role/routes'
import { tripRoutes } from '@/trip/routes'
import { authRoutes } from '@/auth/routes'
import { providerRoutes } from '@/provider/routes'

const appRoutes = [
  ...dashboardRoutes,
  ...userRoutes,
  ...providerRoutes,
  ...tourRoutes,
  ...tripRoutes,
  ...resourceRoutes,
  ...destinationRoutes,
  ...activityRoutes,
  ...mealRoutes,
  ...hotelRoutes,
  ...rolesRoutes,
  ...cashRoutes,
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

export default router
