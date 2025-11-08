import type { Ref } from 'vue'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'

import type { QueryParams } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { ItineraryCommentary } from '../interfaces/itineraryCommentary'

class itineraryCommentaryService {
  public api: PocketBase
  table = 'itineraryCommentaries'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<QueryParams>) {
    return useQuery({
      queryKey: ['itineraryCommentaries', params],
      queryFn: async () => {
        return await this.api
          .collection(this.table)
          .getList<ItineraryCommentary>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['itineraryCommentaries', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api.collection(this.table).getOne<ItineraryCommentary>(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['itineraryCommentary_create'],
      mutationFn: async (data: Partial<ItineraryCommentary> | FormData) => {
        return await this.api.collection(this.table).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['itineraryCommentary_update'],
      mutationFn: async ({
        id,
        data,
      }: {
        id: string
        data: Partial<ItineraryCommentary> | FormData
      }) => {
        return await this.api.collection(this.table).update(id, data)
      },
    })
  }

  useDelete() {
    return useMutation({
      mutationKey: ['itineraryCommentary_delete'],
      mutationFn: async (id: string) => {
        return await this.api.collection(this.table).delete(id)
      },
    })
  }
}

const ItineraryCommentaryService = new itineraryCommentaryService()
export { ItineraryCommentaryService }
