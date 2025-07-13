import type { Ref } from 'vue'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'

import type { ParamsRequest } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { TourDestination } from '../interfaces/tourDestination'

class tourDestinationService {
  public api: PocketBase
  tourDestinationTable = 'tourDestinations'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<ParamsRequest>) {
    return useQuery({
      queryKey: ['tourDestinations', params?.value],
      queryFn: async () => {
        return await this.api
          .collection(this.tourDestinationTable)
          .getList<TourDestination>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
            expand: params?.value.expand,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['tourDestination', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api
          .collection(this.tourDestinationTable)
          .getOne<TourDestination>(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['tourDestination_create'],
      mutationFn: async (data: Partial<TourDestination>) => {
        return await this.api.collection(this.tourDestinationTable).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['tourDestination_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<TourDestination> }) => {
        return await this.api.collection(this.tourDestinationTable).update(id, data)
      },
    })
  }

  useDelete() {
    return useMutation({
      mutationKey: ['tourDestination_delete'],
      mutationFn: async (id: string) => {
        return await this.api.collection(this.tourDestinationTable).delete(id)
      },
    })
  }
}

const TourDestinationService = new tourDestinationService()
export { TourDestinationService }
