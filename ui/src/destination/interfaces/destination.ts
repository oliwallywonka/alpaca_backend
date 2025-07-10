import type { LanguageField, LocationField } from '@/core/interfaces/fields'

export interface Destination {
  id: string
  name: LanguageField
  description: LanguageField
  parent: string
  location: LocationField
  isActive: boolean
  expand: {
    parent?: Destination
  }
}
