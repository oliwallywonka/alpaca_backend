<script setup lang="ts">
import { computed, watch } from 'vue'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/core/components/ui/dialog'
import { Input } from '@/core/components/ui/input'
import { Button } from '@/core/components/ui/button'
import { Select, SelectContent, SelectItem, SelectTrigger } from '@/core/components/ui/select'
import { Badge } from '@/core/components/ui/badge'

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/core/components/ui/table'
import { toTypedSchema } from '@vee-validate/zod'
import { z } from 'zod'
import CustomerPicker from './CustomerPicker.vue'
import { useForm } from 'vee-validate'
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/core/components/ui/form'
import { PaymentService } from '@/trip/services/TripPaymentService'
import { useParams } from '@/core/hooks/useParams'
import { ItineraryResourceService } from '@/itinerary/services/itineraryResources'
import type { Customer } from '@/customer/interfaces/customer'
import { toast } from 'vue-sonner'
import { UserService } from '@/user/services/userService'
import { useQueryClient } from '@tanstack/vue-query'

const CURRENCIES = [
  { value: 'USD', label: 'US Dollar' },
  { value: 'BOL', label: 'Bolivian Boliviano' },
]

const open = defineModel({ default: false })
const props = defineProps<{
  paymentID?: string
  tripID: string
  itinearyID: string
}>()

const userID = UserService.authStore.record?.id

const queryClient = useQueryClient()

const paymentParams = useParams({
  filter: `trip="${props.tripID}"`,
})
const { data: payments, isLoading: isLoadingPayments } = PaymentService.useGetAll(
  paymentParams.params,
)

const itineraryResourceParams = useParams({
  filter: `itinerary="${props.itinearyID}"`,
})
const { data: resources, isLoading: isLoadingResources } = ItineraryResourceService.useGetAll(
  itineraryResourceParams.params,
)
const createPayment = PaymentService.useCreate()

const totalPayments = computed(() => {
  return (
    payments.value?.items.reduce((acc, payment) => acc + payment.amount / payment.changeRate, 0) ??
    0
  )
})

const tripCost = computed(() => {
  return (
    resources.value?.items.reduce(
      (acc, resource) =>
        acc + (resource.originalCost * resource.quantity) / resource.dollarChangeRate,
      0,
    ) ?? 0
  )
})

const schema = toTypedSchema(
  z.object({
    customer: z.object({
      id: z.string(),
      title: z.enum(['Mr', 'Mrs', 'Miss', 'Ms', 'Dr']),
      firstName: z.string(),
      middleName: z.string().optional(),
      lastName: z.string(),
      dateOfBirth: z.string(),
    }),
    amount: z.coerce.number().min(0, 'Amount must be greater than 0'),
    currency: z.enum(['USD', 'BOL']),
    changeRate: z.coerce.number().min(1, 'changeRate must be equal or greater than 1'),
  }),
)

const initialValues: {
  customer?: Customer
  amount: number
  currency: 'USD' | 'BOL'
  changeRate: number
} = {
  customer: undefined,
  amount: 1,
  currency: 'USD',
  changeRate: 1,
}

const { values, handleSubmit, setFieldValue, resetForm } = useForm({
  validationSchema: schema,
  initialValues: initialValues,
})

const paymentLeft = computed(() => {
  return roundHelper(tripCost.value - totalPayments.value - values.amount! / values.changeRate!)
})

watch(
  () => values.currency,
  (newCurrency) => {
    if (newCurrency === 'USD') setFieldValue('changeRate', 1)
  },
)

watch(open, (newIsOpen) => {
  if (!newIsOpen) return
  resetForm()
})

const roundHelper = (number: number) => Math.round(number * 100) / 100

const completePaymentLeft = () => {
  const prePaymentLeft = (tripCost.value - totalPayments.value) * values.changeRate!
  setFieldValue('amount', roundHelper(prePaymentLeft))
}

const onSubmit = handleSubmit(async (values) => {
  try {
    await createPayment.mutateAsync({
      registeredBy: userID!,
      trip: props.tripID,
      customer: values.customer!.id,
      amount: values.amount!,
      currency: values.currency!,
      changeRate: values.changeRate!,
    })
    await queryClient.invalidateQueries({ queryKey: ['payments'] })
    toast.success('Payment submitted successfully')
    open.value = false
  } catch (error) {
    console.error('Failed to submit payment:', error)
    toast.error('Failed to submit payment. Please try again.')
  }
})
</script>

<template>
  <Dialog v-model:open="open">
    <DialogTrigger as-child>
      <Button variant="outline"> {{ props.paymentID ? 'Edit' : 'Add' }} Payment </Button>
    </DialogTrigger>
    <DialogContent class="sm:max-w-4xl">
      <DialogHeader>
        <DialogTitle>{{ props.paymentID ? 'Edit' : 'Add' }} Payment</DialogTitle>
        <DialogDescription> Fill out the form to add a new payment. </DialogDescription>
      </DialogHeader>
      <p v-if="isLoadingPayments || isLoadingResources">Loading...</p>
      <form
        v-if="!isLoadingPayments && !isLoadingResources"
        id="trip-payment-form"
        @submit.prevent="onSubmit"
      >
        <FormField v-slot="{ componentField }" name="customer">
          <FormItem>
            <FormLabel>Contacts </FormLabel>
            <FormControl>
              <CustomerPicker
                v-model="componentField.modelValue"
                @update:model-value="setFieldValue('customer', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <Table>
          <TableHeader class="bg-white">
            <TableRow>
              <TableHead> </TableHead>
              <TableHead>Amount</TableHead>
              <TableHead>Currency</TableHead>
              <TableHead>Change Rate</TableHead>
              <TableHead class="text-right"> SubTotal </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell class="font-medium"> Trip cost: </TableCell>
              <TableCell>{{ roundHelper(tripCost).toFixed(2) }}</TableCell>
              <TableCell><Badge variant="green">USD</Badge></TableCell>
              <TableCell>1.00</TableCell>
              <TableCell class="text-right">
                {{ roundHelper(tripCost).toFixed(2) }}
                <Badge variant="green">USD</Badge>
              </TableCell>
            </TableRow>

            <TableRow>
              <TableCell class="font-medium"> Payments: </TableCell>
              <TableCell>{{ roundHelper(totalPayments).toFixed(2) }}</TableCell>
              <TableCell><Badge variant="green">USD</Badge></TableCell>
              <TableCell>1.00</TableCell>
              <TableCell class="text-right">
                {{ roundHelper(totalPayments).toFixed(2) }} <Badge variant="green">USD</Badge>
              </TableCell>
            </TableRow>

            <TableRow>
              <TableCell class="font-medium">
                Payment:
                <Button size="sm" @click="completePaymentLeft">fill</Button>
              </TableCell>
              <TableCell>
                <FormField v-slot="{ componentField }" name="amount">
                  <FormItem>
                    <FormControl>
                      <Input
                        type="number"
                        inputmode="decimal"
                        step="any"
                        placeholder="0.00"
                        v-bind="componentField"
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                </FormField>
              </TableCell>
              <TableCell>
                <FormField v-slot="{ componentField, value }" name="currency">
                  <FormItem>
                    <Select v-bind="componentField">
                      <FormControl>
                        <SelectTrigger>
                          {{ value }}
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent class="min-w-24">
                        <SelectItem
                          v-for="item in CURRENCIES"
                          :key="item.value"
                          :value="item.value"
                        >
                          {{ item.value }}
                          <span class="text-muted-foreground">{{ item.label }}</span>
                        </SelectItem>
                      </SelectContent>
                    </Select>
                    <FormMessage />
                  </FormItem>
                </FormField>
              </TableCell>
              <TableCell>
                <FormField v-slot="{ componentField }" name="changeRate">
                  <FormItem>
                    <FormControl>
                      <Input
                        :disabled="values.currency === 'USD'"
                        type="number"
                        inputmode="decimal"
                        step="any"
                        placeholder="0.00"
                        v-bind="componentField"
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                </FormField>
              </TableCell>
              <TableCell class="text-right">
                {{ roundHelper(values.amount! / values.changeRate!).toFixed(2) }}
                <Badge variant="green">USD</Badge>
              </TableCell>
            </TableRow>

            <TableRow>
              <TableCell class="font-medium"> Payment Left: </TableCell>
              <TableCell colSpan="4" class="text-right">
                {{ paymentLeft.toFixed(2) }} <Badge variant="green">USD</Badge>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </form>

      <DialogFooter>
        <Button type="submit" form="trip-payment-form"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
