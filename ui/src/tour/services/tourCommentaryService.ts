import type { Ref } from 'vue'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'

import type { ParamsRequest } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { TourCommentary } from '../interfaces/tourCommentary'

class tourCommentaryService {
  public api: PocketBase
  tourCommentaryTable = 'tourCommentaries'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<ParamsRequest>) {
    return useQuery({
      queryKey: ['tourCommentaries', params?.value],
      queryFn: async () => {
        return await this.api
          .collection(this.tourCommentaryTable)
          .getList<TourCommentary>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['tourCommentaries', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api.collection(this.tourCommentaryTable).getOne<TourCommentary>(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['tourCommentary_create'],
      mutationFn: async (data: Partial<TourCommentary> | FormData) => {
        return await this.api.collection(this.tourCommentaryTable).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['tourCommentary_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<TourCommentary> | FormData }) => {
        return await this.api.collection(this.tourCommentaryTable).update(id, data)
      },
    })
  }

  useDelete() {
    return useMutation({
      mutationKey: ['tourCommentary_delete'],
      mutationFn: async (id: string) => {
        return await this.api.collection(this.tourCommentaryTable).delete(id)
      },
    })
  }
}

const TourCommentaryService = new tourCommentaryService()
export { TourCommentaryService }
