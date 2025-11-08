import type { Itinerary } from "@/itinerary/interfaces/itinerary"
import type { Tour } from "@/tour/interfaces/tour"

export type TripState = 'draft' | 'booked' | 'inProgress' | 'completed' | 'cancelled'
export enum TripStateEnum {
  draft = 'draft',
  booked = 'booked',
  inProgress = 'inProgress',
  completed = 'completed',
  cancelled = 'cancelled',
}

export interface Trip {
  id: string
  tour: string
  originalItinerary: string
  itinerary: string
  registeredBy: string
  coordinator: string
  customerLead: string
  state: TripState
  expand: {
    tour: Tour
    itinerary: Itinerary
    originalItinerary: Itinerary
  }
}

export interface TripSummary extends Trip {
  startDate: Date
  endDate: Date
}

export interface NewTrip {
  itineraryID: string
  transferResources: boolean
  startDate: Date | string | null
}
