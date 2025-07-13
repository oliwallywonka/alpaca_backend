import type { ParamsRequest } from '@/core/interfaces/api'
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
  userTable = 'users'
  resourceProviderTable = 'resourceProviders'
  constructor() {
    this.api = API
    this.authStore = this.api.authStore
  }

  useAuth() {
    return useMutation({
      mutationKey: ['user_auth'],
      mutationFn: async ({ identity, password }: { identity: string; password: string }) => {
        return await this.api.collection<User>(this.userTable).authWithPassword(identity, password, {
          expand: 'role',
        })
      },
    })
  }

  logOut() {
    this.authStore.clear()
  }

  useGetAll(params?: Ref<ParamsRequest>) {
    return useQuery({
      queryKey: ['users', params?.value],
      queryFn: async () => {
        return await this.api
          .collection<User>(this.userTable)
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
        return await this.api.collection<User>(this.userTable).getOne(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['user_create'],
      mutationFn: async (data: Partial<User>) => {
        return await this.api.collection(this.userTable).create({ ...data, emailVisibility: true })
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['user_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<User> }) => {
        return await this.api.collection(this.userTable).update(id, data)
      },
    })
  }

  useGetResources(userID: Ref<string>) {
    return useQuery({
      queryKey: ['user_resources', userID],
      queryFn: async () => {
        return await this.api
          .collection<ResourceProvider>(this.resourceProviderTable)
          .getList(1, 100, {
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
          .collection<ResourceProvider>(this.resourceProviderTable)
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
          .collection<ResourceProvider>(this.resourceProviderTable)
          .update(data.resourceProviderID, { ...data.data, user: data.userID })
      },
    })
  }
  useDeleteResource() {
    return useMutation({
      mutationKey: ['user_resource_delete'],
      mutationFn: async (data: { userID: string; resourceProviderID: string }) => {
        return await this.api
          .collection<ResourceProvider>(this.resourceProviderTable)
          .delete(data.resourceProviderID)
      },
    })
  }
}

const UserService = new userService()
export { UserService }
