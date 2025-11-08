import type { LanguageField } from '@/core/interfaces/fields'

export interface ItineraryCommentary {
  id: string
  tour: string
  itinerary: string
  commentary: LanguageField
  images: string[]
  startDate: string
  endDate: string
  isAllDay: boolean
  created: string
  updated: string
}
