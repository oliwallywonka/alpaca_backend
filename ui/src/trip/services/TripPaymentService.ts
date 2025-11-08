import type { Ref } from 'vue'
import { useMutation, useQuery } from '@tanstack/vue-query'
import type PocketBase from 'pocketbase'

import type { QueryParams } from '@/core/interfaces/api'
import { API } from '@/core/services/pocketbase'
import type { Payment } from '@/trip/interfaces/payment'

class paymentService {
  public api: PocketBase
  payment = 'tripPayments'
  constructor() {
    this.api = API
  }

  useGetAll(params?: Ref<QueryParams>) {
    return useQuery({
      queryKey: ['payments', params?.value],
      queryFn: async () => {
        return await this.api
          .collection(this.payment)
          .getList<Payment>(params?.value.page, params?.value.perPage, {
            filter: params?.value.filter,
            orderBy: params?.value.orderBy,
            expand: params?.value.expand,
          })
      },
    })
  }

  useGetOne(id: Ref<string>, startFetching: Ref<boolean>) {
    return useQuery({
      queryKey: ['payment', id],
      enabled: startFetching,
      queryFn: async () => {
        return await this.api.collection(this.payment).getOne<Payment>(id.value)
      },
    })
  }

  useCreate() {
    return useMutation({
      mutationKey: ['payment_create'],
      mutationFn: async (data: Partial<Payment>) => {
        return await this.api.collection(this.payment).create(data)
      },
    })
  }

  useUpdate() {
    return useMutation({
      mutationKey: ['payment_update'],
      mutationFn: async ({ id, data }: { id: string; data: Partial<Payment> | FormData }) => {
        return await this.api.collection(this.payment).update(id, data)
      },
    })
  }
}

const PaymentService = new paymentService()
export { PaymentService }
