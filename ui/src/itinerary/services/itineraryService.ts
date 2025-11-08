import type { Ref } from 'vue'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'

import type { QueryParams } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { Itinerary, ItinerarySummary } from '../interfaces/itinerary'

class itineraryVariantService {
  public api: PocketBase
  itineraryTable = 'itineraries'
  itinerarySummaryTable = 'itinerariesSummary'
  constructor() {
    this.api = API
  }

  useGetSummaries(params?: Ref<QueryParams>) {
    return useQuery({
      queryKey: ['itinerariesSummary', params?.value],
      queryFn: async () => {
        return await this.api
          .collection(this.itinerarySummaryTable)
          .getList<ItinerarySummary>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
            expand: params?.value.expand,
          })
      },
    })
  }

  useGetAll(params?: Ref<QueryParams>) {
    return useQuery({
      queryKey: ['itineraries'],
      queryFn: async () => {
        return await this.api
          .collection(this.itineraryTable)
          .getList<Itinerary>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
            expand: params?.value.expand,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['itinerary', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api.collection(this.itineraryTable).getOne<Itinerary>(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['itinerary_create'],
      mutationFn: async (data: Partial<Itinerary>) => {
        return await this.api.collection(this.itineraryTable).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['itinerary_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<Itinerary> | FormData }) => {
        return await this.api.collection(this.itineraryTable).update(id, data)
      },
    })
  }
}

const ItineraryService = new itineraryVariantService()
export { ItineraryService }
