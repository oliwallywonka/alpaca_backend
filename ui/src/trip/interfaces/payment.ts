import type { Customer } from '@/customer/interfaces/customer'
import type { User } from '@/user/interfaces/user'

export interface Payment {
  id: string
  trip: string
  customer: string
  registeredBy: string
  amount: number
  currency: string
  changeRate: number
  method: string
  created: string
  updated: string
  expand: {
    customer: Customer
    registeredBy: User
  }
}
