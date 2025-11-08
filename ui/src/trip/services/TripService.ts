import type { Ref } from 'vue'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'

import type { QueryParams } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { NewTrip, Trip, TripSummary } from '@/trip/interfaces/trip'

class tripService {
  public api: PocketBase
  tripTable = 'trips'
  tripSummaryTable = 'tripsSummary'
  constructor() {
    this.api = API
  }

  useGetSummaries(params?: Ref<QueryParams>) {
    return useQuery({
      queryKey: ['tripsSummary', params?.value],
      queryFn: async () => {
        return await this.api
          .collection(this.tripSummaryTable)
          .getList<TripSummary>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
            expand: params?.value.expand,
          })
      },
    })
  }

  useGetAll(params?: Ref<QueryParams>) {
    return useQuery({
      queryKey: ['trips', params?.value],
      queryFn: async () => {
        return await this.api
          .collection(this.tripTable)
          .getList<Trip>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
            orderBy: params?.value.orderBy,
            expand: params?.value.expand,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['trip', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api.collection(this.tripTable).getOne<Trip>(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['trip_create'],
      mutationFn: async (data: NewTrip) => {
        return await this.api.send<Trip>('trips', {
          method: 'POST',
          body: data,
        })
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['trip_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<Trip> | FormData }) => {
        return await this.api.collection(this.tripTable).update(id, data)
      },
    })
  }
}

const TripService = new tripService()
export { TripService }
