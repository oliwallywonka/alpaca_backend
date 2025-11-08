import type { Ref } from 'vue'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'

import type { QueryParams } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { ItineraryDestination } from '../interfaces/itineraryDestination'

class itineraryDestinationService {
  public api: PocketBase
  itineraryDestinationTable = 'itineraryDestinations'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<QueryParams>) {
    return useQuery({
      queryKey: ['itineraryDestinations', params],
      queryFn: async () => {
        return await this.api
          .collection(this.itineraryDestinationTable)
          .getList<ItineraryDestination>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
            expand: params?.value.expand,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['itineraryDestination', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api
          .collection(this.itineraryDestinationTable)
          .getOne<ItineraryDestination>(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['itineraryDestination_create'],
      mutationFn: async (data: Partial<ItineraryDestination>) => {
        return await this.api.collection(this.itineraryDestinationTable).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['itineraryDestination_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<ItineraryDestination> }) => {
        return await this.api.collection(this.itineraryDestinationTable).update(id, data)
      },
    })
  }

  useDelete() {
    return useMutation({
      mutationKey: ['itineraryDestination_delete'],
      mutationFn: async (id: string) => {
        return await this.api.collection(this.itineraryDestinationTable).delete(id)
      },
    })
  }
}

const ItineraryDestinationService = new itineraryDestinationService()
export { ItineraryDestinationService }
