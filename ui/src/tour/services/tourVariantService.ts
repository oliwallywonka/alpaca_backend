import type { Ref } from 'vue'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'

import type { ParamsRequest } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { TourVariant } from '../interfaces/tourVariant'

class tourVariantService {
  public api: PocketBase
  tourVariantTable = 'tourVariants'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<ParamsRequest>) {
    return useQuery({
      queryKey: ['tourVariants', params?.value],
      queryFn: async () => {
        return await this.api
          .collection(this.tourVariantTable)
          .getList<TourVariant>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
            expand: params?.value.expand,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['tourVariant', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api.collection(this.tourVariantTable).getOne<TourVariant>(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['tourVariant_create'],
      mutationFn: async (data: Partial<TourVariant>) => {
        return await this.api.collection(this.tourVariantTable).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['tourVariant_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<TourVariant> | FormData }) => {
        return await this.api.collection(this.tourVariantTable).update(id, data)
      },
    })
  }
}

const TourVariantService = new tourVariantService()
export { TourVariantService }
