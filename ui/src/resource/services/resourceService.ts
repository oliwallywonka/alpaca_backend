import { type Ref } from 'vue'
import type PocketBase from 'pocketbase'
import { useMutation, useQuery } from '@tanstack/vue-query'

import { API } from '@/core/services/pocketbase'
import type { ParamsRequest } from '@/core/interfaces/api'
import type { Resource } from '../interfaces/resource'

class resourceService {
  public api: PocketBase
  resourceTable = 'resources'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<ParamsRequest>) {
    return useQuery({
      queryKey: ['resources', params?.value],
      queryFn: async () => {
        return await this.api
          .collection(this.resourceTable)
          .getList<Resource>(params?.value?.page || 1, params?.value?.perPage || 10, {
            filter: params?.value?.filter || '',
            sort: params?.value?.orderBy || '-created',
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['resource', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api.collection(this.resourceTable).getOne<Resource>(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['resource_create'],
      mutationFn: async (data: Partial<Resource>) => {
        return await this.api.collection(this.resourceTable).create<Resource>(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['resource_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<Resource> }) => {
        return await this.api.collection(this.resourceTable).update<Resource>(id, data)
      },
    })
  }
}

const ResourceService = new resourceService()
export { ResourceService }
