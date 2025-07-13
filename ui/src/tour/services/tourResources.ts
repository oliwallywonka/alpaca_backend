import type { Ref } from 'vue'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'

import type { ParamsRequest } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { TourResource } from '../interfaces/tourResources'

class tourResources {
  public api: PocketBase
  tourResourcesTable = 'tourResources'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<ParamsRequest>) {
    return useQuery({
      queryKey: ['tourResources', params?.value],
      queryFn: async () => {
        return await this.api
          .collection(this.tourResourcesTable)
          .getList<TourResource>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
            sort: params?.value.orderBy,
            expand: params?.value.expand,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['tourResource', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api
          .collection<TourResource>(this.tourResourcesTable)
          .getOne(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['tourResource_create'],
      mutationFn: async (data: Partial<TourResource>) => {
        return await this.api.collection<TourResource>(this.tourResourcesTable).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['tourResource_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<TourResource> | FormData }) => {
        return await this.api.collection(this.tourResourcesTable).update(id, data)
      },
    })
  }

  useDelete() {
    return useMutation({
      mutationKey: ['tourResource_delete'],
      mutationFn: async (id: string) => {
        return await this.api.collection(this.tourResourcesTable).delete(id)
      },
    })
  }
}

const TourResourceService = new tourResources()
export { TourResourceService as TourResourcesService }
