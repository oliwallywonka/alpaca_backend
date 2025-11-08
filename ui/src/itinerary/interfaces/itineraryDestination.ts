import type { LanguageField } from '@/core/interfaces/fields'

import type { Destination } from '@/destination/interfaces/destination'
export interface ItineraryDestination {
  id: string
  tour: string
  itinerary: string
  destination: string
  descritption: LanguageField
  startDate: string
  endDate: string
  isAllDay: boolean
  created: string
  updated: string
  expand: {
    destination: Destination
  }
}
