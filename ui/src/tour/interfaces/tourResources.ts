import type { ResourceProvider } from '@/user/interfaces/resourcePrices'
import type { Tour } from './tour'
import type { TourVariant } from './tourVariant'

export type Currencies = 'USD' | 'BOL'
export interface TourResource {
  id: string
  tour: string
  tourVariant: string
  resourceProvider: string
  startDate: string
  endDate: string
  isAllDay: boolean
  originalCost: number
  profitPercent: number
  currency: string
  dollarChangeRate: number
  quantity: number
  isVisible: boolean
  created: string
  updated: string
  expand: {
    tour: Tour
    tourVariant: TourVariant
    resourceProvider: ResourceProvider
  }
}
