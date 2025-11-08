import type { Tour } from '../../tour/interfaces/tour'

export interface Itinerary {
  id: string
  tour: string
  minPersons: number
  maxPersons: number
  finalPrice: number
  isActive: boolean
  isTemplate: boolean
  created: string
  updated: string
  expand: {
    tour: Partial<Tour>
  }
}

export interface ItinerarySummary extends Itinerary {
  startDate?: string | Date
  endDate?: string | Date
}
