import type { LanguageField } from '@/core/interfaces/fields'

export interface Tour {
  id: string
  code: string
  name: LanguageField
  slug: LanguageField
  discount: number
  days: number
  groupSize: string
  transport: LanguageField
  accommodation: LanguageField
  team: LanguageField
  shortDescription: LanguageField
  longDescription: LanguageField
  banner: string
  images: string[]
  isPublic: boolean
  isActive: boolean
  created: string
  updated: string
}
