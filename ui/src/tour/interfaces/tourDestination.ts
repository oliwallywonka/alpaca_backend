import type { LanguageField } from '@/core/interfaces/fields'
import type { Destination } from '@/destination/interfaces/destination'

export interface TourDestination {
  id: string
  tour: string
  tourVariant: string
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
