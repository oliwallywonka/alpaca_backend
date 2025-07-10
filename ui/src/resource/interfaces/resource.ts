import type { LanguageField, LocationField } from '@/core/interfaces/fields'

export interface Resource {
  id: string
  destination: string
  location: LocationField,
  name: LanguageField
  description: LanguageField
  images: string[]
  types: string[]
  isActive: boolean
  created: number
  updated: number
}
