import type { ResourceProvider } from '@/user/interfaces/resourcePrices'
import type { Tour } from '@/tour/interfaces/tour'
import type { Itinerary } from '../../itinerary/interfaces/itinerary'

export type Currencies = 'USD' | 'BOL'
export interface ItineraryResource {
  id: string
  tour: string
  itinerary: string
  resourceProvider: string
  startDate: string
  endDate: string
  isAllDay: boolean
  originalCost: number
  minPersons: number
  maxPersons: number
  profitPercent: number
  currency: string
  dollarChangeRate: number
  quantity: number
  isVisible: boolean
  created: string
  updated: string
  expand: {
    tour: Tour
    itinerary: Itinerary
    resourceProvider: ResourceProvider
  }
}
