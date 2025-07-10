import type { Ref } from 'vue'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'

import type { ParamsRequest } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { Tour } from '../interfaces/tour'

class tourService {
  public api: PocketBase
  tourTable = 'tours'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<ParamsRequest>) {
    return useQuery({
      queryKey: ['tours', params?.value],
      queryFn: async () => {
        return await this.api
          .collection(this.tourTable)
          .getList<Tour>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['tour', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api.collection(this.tourTable).getOne<Tour>(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['tour_create'],
      mutationFn: async (data: Partial<Tour>) => {
        return await this.api.collection(this.tourTable).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['tour_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<Tour> | FormData }) => {
        return await this.api.collection(this.tourTable).update(id, data)
      },
    })
  }
}

const TourService = new tourService()
export { TourService }
