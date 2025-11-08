import type { QueryParams } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'
import type { Ref } from 'vue'
import type { User } from '../interfaces/user'
import type { ResourceProvider } from '../interfaces/resourcePrices'
import type { BaseAuthStore } from 'pocketbase'

class userService {
  public api: PocketBase
  public authStore: BaseAuthStore
  user = 'users'
  resourceProvider = 'resourceProviders'
  constructor() {
    this.api = API
    this.authStore = this.api.authStore
  }

  useAuth() {
    return useMutation({
      mutationKey: ['user_auth'],
      mutationFn: async ({ identity, password }: { identity: string; password: string }) => {
        const res = await this.api
          .collection<User>(this.user)
          .authWithPassword(identity, password, {
            expand: 'role',
          })
        console.log(res)
        return res
      },
    })
  }

  logOut() {
    this.authStore.clear()
  }

  useGetAll(params?: Ref<QueryParams>) {
    return useQuery({
      queryKey: ['users', params?.value],
      queryFn: async () => {
        return await this.api
          .collection<User>(this.user)
          .getList(params?.value.page || 1, params?.value?.perPage || 20, {
            filter: params?.value?.filter || '',
            expand: 'role',
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['user', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api.collection<User>(this.user).getOne(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['user_create'],
      mutationFn: async (data: Partial<User>) => {
        return await this.api.collection(this.user).create({ ...data, emailVisibility: true })
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['user_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<User> }) => {
        return await this.api.collection(this.user).update(id, data)
      },
    })
  }

  useGetResources(userID: Ref<string>) {
    return useQuery({
      queryKey: ['user_resources', userID],
      queryFn: async () => {
        return await this.api.collection<ResourceProvider>(this.resourceProvider).getList(1, 100, {
          filter: `user.id="${userID.value}"`,
          expand: 'user,resource',
        })
      },
    })
  }

  useCreateResource() {
    return useMutation({
      mutationKey: ['user_resource_create'],
      mutationFn: async (data: { userID: string; data: Partial<ResourceProvider> }) => {
        return await this.api
          .collection<ResourceProvider>(this.resourceProvider)
          .create({ ...data.data, user: data.userID })
      },
    })
  }

  useUpdateResource() {
    return useMutation({
      mutationKey: ['user_resource_update'],
      mutationFn: async (data: {
        userID: string
        resourceProviderID: string
        data: Partial<ResourceProvider>
      }) => {
        return await this.api
          .collection<ResourceProvider>(this.resourceProvider)
          .update(data.resourceProviderID, { ...data.data, user: data.userID })
      },
    })
  }
  useDeleteResource() {
    return useMutation({
      mutationKey: ['user_resource_delete'],
      mutationFn: async (data: { userID: string; resourceProviderID: string }) => {
        return await this.api
          .collection<ResourceProvider>(this.resourceProvider)
          .delete(data.resourceProviderID)
      },
    })
  }
}

const UserService = new userService()
export { UserService }
