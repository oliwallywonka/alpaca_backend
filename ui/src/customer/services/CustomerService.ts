import type { Ref } from 'vue'

import type PocketBase from 'pocketbase'
import { useMutation, useQuery } from '@tanstack/vue-query'

import { API } from '@/core/services/pocketbase'
import type { ParamsRequest } from '@/core/interfaces/api'
import type { Customer } from '../interfaces/customer'

class customerService {
  public api: PocketBase
  customerTable = 'customers'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<ParamsRequest>) {
    return useQuery({
      queryKey: ['customers', params?.value],
      queryFn: async () => {
        return await this.api
          .collection<Customer>(this.customerTable)
          .getList(params?.value?.page || 1, params?.value?.perPage || 20, {
            filter: params?.value?.filter || '',
            sort: params?.value?.orderBy || 'created',
            order: params?.value?.orderDirection || 'DESC',
            expand: params?.value.expand || '',
          })
      },
    })
  }

  useGetOne(id: Ref<string>, enabled: Ref<boolean>) {
    return useQuery({
      queryKey: ['role', id],
      enabled,
      queryFn: async () => {
        return await this.api.collection<Customer>(this.customerTable).getOne(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['role_create'],
      mutationFn: async (data: Partial<Customer>) => {
        return await this.api.collection(this.customerTable).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['role'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<Customer> }) => {
        return await this.api.collection(this.customerTable).update(id, data)
      },
    })
  }
}

const CustomerService = new customerService()
export { CustomerService }
