import type { ContactField } from '@/core/interfaces/fields'

export interface Provider {
  id: string
  fullName: string
  description: string
  contacts: ContactField[]
  isActive: boolean
  types: string[]
  created: string
  updated: string
}
