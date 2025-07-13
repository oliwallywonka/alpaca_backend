import type { PriceField } from '@/core/interfaces/fields'
import type { Provider } from '@/provider/interfaces/provider'
import type { Resource } from '@/resource/interfaces/resource'
import type { User } from '@/user/interfaces/user'

export interface ResourceProvider {
  id: string
  resource: string
  provider?: string
  user?: string
  refPrices: PriceField[]
  type: string
  created: string
  updated: string
  expand: {
    resource: Resource
    provider: Provider
    user: User
  }
}
