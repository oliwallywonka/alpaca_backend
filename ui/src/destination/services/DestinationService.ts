import type { Ref } from 'vue'
import type PocketBase from 'pocketbase'
import { useMutation, useQuery } from '@tanstack/vue-query'

import type { ParamsRequest } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { Destination } from '../interfaces/destination'

class destinationService {
  public api: PocketBase
  destinationTable = 'destinations'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<Partial<ParamsRequest>>) {
    return useQuery({
      queryKey: ['destinations', params?.value],
      queryFn: async () => {
        return await this.api
          .collection<Destination>(this.destinationTable)
          .getList(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter || '',
            expand: 'parent',
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['destination', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api.collection<Destination>(this.destinationTable).getOne(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['destination_create'],
      mutationFn: async (data: Partial<Destination>) => {
        return await this.api.collection(this.destinationTable).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['destination_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<Destination> }) => {
        return await this.api.collection(this.destinationTable).update(id, data)
      },
    })
  }
}

const DestinationService = new destinationService()
export { DestinationService }
