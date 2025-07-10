import type { Resource } from '@/resource/interfaces/resource'
import type { User } from './user'
import type { Provider } from '@/provider/interfaces/provider'

export interface Price {
  minPersons: number
  maxPersons: number
  price: number
  currency: string
  isPerPerson: boolean
}
export interface ResourceProvider {
  id: string
  resource: string
  user: string
  provider: string
  type: string
  refPrices: Price[]
  expand: {
    resource: Resource
    provider: Provider
    user: User
  }
}
