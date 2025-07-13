import type { RouteRecordRaw } from 'vue-router'

export const tripRoutes: RouteRecordRaw[] = [
  {
    path: '/trips',
    name: 'trips',
    component: () => import('./pages/TripsPage/TripsPage.vue'),
  },
  {
    path: '/trips/:tripID',
    name: 'trip',
    component: () => import('./pages/TripDetailPage/TripDetailPage.vue'),
  },
]
