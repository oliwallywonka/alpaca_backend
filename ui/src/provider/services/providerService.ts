import type { Ref } from 'vue'
import type PocketBase from 'pocketbase'
import { useMutation, useQuery } from '@tanstack/vue-query'

import { API } from '@/core/services/pocketbase'
import type { ParamsRequest } from '@/core/interfaces/api'
import type { Provider } from '../interfaces/provider'
import type { ResourceProvider } from '@/user/interfaces/resourcePrices'

class providerService {
  public api: PocketBase
  providerTable = 'providers'
  resourceProviderTable = 'resourceProviders'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<ParamsRequest>) {
    return useQuery({
      queryKey: ['providers', params?.value],
      queryFn: async () => {
        return await this.api
          .collection<Provider>(this.providerTable)
          .getList(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter || '',
            sort: params?.value.orderBy,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['provider', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api.collection<Provider>(this.providerTable).getOne(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['provider_create'],
      mutationFn: async (data: Partial<Provider>) => {
        return await this.api.collection(this.providerTable).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['provider_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<Provider> }) => {
        return await this.api.collection(this.providerTable).update(id, data)
      },
    })
  }

  useGetResources(providerID: Ref<string>) {
    return useQuery({
      queryKey: ['provider_resources', providerID],
      queryFn: async () => {
        return await this.api
          .collection<ResourceProvider>(this.resourceProviderTable)
          .getList(1, 100, {
            filter: `provider.id="${providerID.value}"`,
            expand: 'provider,resource',
          })
      },
    })
  }

  useCreateResource() {
    return useMutation({
      mutationKey: ['provider_resource_create'],
      mutationFn: async (data: { providerID: string; data: Partial<ResourceProvider> }) => {
        return await this.api
          .collection(this.resourceProviderTable)
          .create({ ...data.data, provider: data.providerID })
      },
    })
  }
  useUpdateService() {
    return useMutation({
      mutationKey: ['provider_resource_update'],
      mutationFn: async (data: {
        providerID: string
        resourceProviderID: string
        data: Partial<ResourceProvider>
      }) => {
        return await this.api
          .collection(this.resourceProviderTable)
          .update(data.resourceProviderID, { ...data.data, provider: data.providerID })
      },
    })
  }
  useDeleteService() {
    return useMutation({
      mutationKey: ['provider_service_delete'],
      mutationFn: async (data: { providerID: string; resourceProviderID: string }) => {
        return await this.api.collection(this.resourceProviderTable).delete(data.resourceProviderID)
      },
    })
  }
}

const ProviderService = new providerService()
export { ProviderService }
