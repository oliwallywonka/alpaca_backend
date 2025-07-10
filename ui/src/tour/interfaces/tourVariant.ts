import type { Tour } from "./tour"

export interface TourVariant {
  id: string
  tour: string
  minPersons: number
  maxPersons: number
  totalPrice: number
  isActive: boolean
  created: string
  updated: string
  expand: {
    tour: Partial<Tour>
  }
}
