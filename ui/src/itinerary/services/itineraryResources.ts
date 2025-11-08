import type { Ref } from 'vue'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'

import type { QueryParams } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { ItineraryResource } from '../interfaces/itineraryResource'

class itineraryResources {
  public api: PocketBase
  itineraryResourcesTable = 'itineraryResources'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<QueryParams>) {
    return useQuery({
      queryKey: ['itineraryResources', params],
      enabled: true,
      queryFn: async () => {
        return await this.api
          .collection(this.itineraryResourcesTable)
          .getList<ItineraryResource>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
            sort: params?.value.orderBy,
            expand: params?.value.expand,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>, params?: Ref<QueryParams>) {
    return useQuery({
      queryKey: ['itineraryResource', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api
          .collection<ItineraryResource>(this.itineraryResourcesTable)
          .getOne(id.value, { ...params?.value })
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['itineraryResource_create'],
      mutationFn: async (data: Partial<ItineraryResource>) => {
        return await this.api
          .collection<ItineraryResource>(this.itineraryResourcesTable)
          .create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['itineraryResource_update'],
      mutationFn: async ({
        id,
        data,
      }: {
        id: string
        data: Partial<ItineraryResource> | FormData
      }) => {
        return await this.api.collection(this.itineraryResourcesTable).update(id, data)
      },
    })
  }

  useDelete() {
    return useMutation({
      mutationKey: ['itineraryResource_delete'],
      mutationFn: async (id: string) => {
        return await this.api.collection(this.itineraryResourcesTable).delete(id)
      },
    })
  }
}

const ItineraryResourceService = new itineraryResources()
export { ItineraryResourceService }
