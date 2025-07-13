import { type Ref } from 'vue'
import type PocketBase from 'pocketbase'
import { useMutation, useQuery } from '@tanstack/vue-query'

import { API } from '@/core/services/pocketbase'
import type { ParamsRequest } from '@/core/interfaces/api'
import type { ResourceProvider } from './ResourceProvider'

class resourceProviderService {
  public api: PocketBase
  resourceProviderTable = 'resourceProviders'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<ParamsRequest>) {
    return useQuery({
      queryKey: ['provierResources', params?.value],
      queryFn: async () => {
        return await this.api
          .collection(this.resourceProviderTable)
          .getList<ResourceProvider>(params?.value?.page || 1, params?.value?.perPage || 10, {
            filter: params?.value?.filter || '',
            sort: params?.value?.orderBy || '-created',
            expand: params?.value?.expand || '',
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['providerResource', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api
          .collection(this.resourceProviderTable)
          .getOne<ResourceProvider>(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['providerResource_create'],
      mutationFn: async (data: Partial<ResourceProvider>) => {
        return await this.api.collection(this.resourceProviderTable).create<ResourceProvider>(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['providerResource_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<ResourceProvider> }) => {
        return await this.api
          .collection(this.resourceProviderTable)
          .update<ResourceProvider>(id, data)
      },
    })
  }

  useDelete() {
    return useMutation({
      mutationKey: ['providerResource_delete'],
      mutationFn: async (id: string) => {
        return await this.api.collection(this.resourceProviderTable).delete(id)
      },
    })
  }
}

const ResourceProviderService = new resourceProviderService()
export { ResourceProviderService }
