import type { Ref } from 'vue'

import type PocketBase from 'pocketbase'
import { useMutation, useQuery } from '@tanstack/vue-query'

import { API } from '@/core/services/pocketbase'
import type { QueryParams } from '@/core/interfaces/api'
import type { Role } from '../interfaces/role'

class roleService {
  public api: PocketBase
  roleTable = 'roles'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<QueryParams>) {
    return useQuery({
      queryKey: ['roles', params?.value],
      queryFn: async () => {
        return await this.api
          .collection<Role>(this.roleTable)
          .getList(params?.value?.page || 1, params?.value?.perPage || 20, {
            filter: params?.value?.filter || '',
            sort: params?.value?.orderBy || 'created',
            order: params?.value?.orderDirection || 'DESC',
          })
      },
    })
  }

  useGetOne(id: Ref<string>, enabled: Ref<boolean>) {
    return useQuery({
      queryKey: ['role', id],
      enabled,
      queryFn: async () => {
        return await this.api.collection<Role>(this.roleTable).getOne(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['role_create'],
      mutationFn: async (data: Partial<Role>) => {
        return await this.api.collection(this.roleTable).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['role'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<Role> }) => {
        return await this.api.collection(this.roleTable).update(id, data)
      },
    })
  }
}

const RoleService = new roleService()
export { RoleService }
