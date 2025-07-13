import type { ContactField } from '@/core/interfaces/fields'
import type { Role } from '@/role/interfaces/role'

export interface User {
  id: string
  role: string
  email: string
  password: string
  passwordConfirm: string
  name: string
  avatar: null | string | string[]
  isActive: boolean
  created: number
  updated: number
  contacts: ContactField[]
  expand: {
    role: Role
  }
}

export interface Auth {
  toke: string
  record: User
}
