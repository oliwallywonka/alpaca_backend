<script setup lang="ts">
import { computed, toRef } from 'vue'
import { useRoute } from 'vue-router'
import PaymentsTable from './PaymentsTable.vue'
import PaymentForm from './PaymentForm.vue'
import { TripService } from '@/trip/services/TripService'
import { PaymentService } from '@/trip/services/TripPaymentService'
import { useParams } from '@/core/hooks/useParams'
import { Button } from '@/core/components/ui/button'

const route = useRoute()
const tripID = computed(() => route.params.tripID as string)
const trip = TripService.useGetOne(toRef(tripID), toRef(true))
const paymentParams = useParams({
  filter: `trip="${tripID.value}"`,
  expand: 'customer, registeredBy',
})
const payments = PaymentService.useGetAll(paymentParams.params)
</script>

<template>
  <div class="grid gap-2">
    <PaymentForm
      v-if="trip.data.value?.itinerary && tripID"
      :tripID="tripID"
      :itinearyID="trip.data.value?.itinerary"
    />
    <Button @click="payments.refetch">Refresh </Button>
    <PaymentsTable
      v-if="!trip.isLoading.value && !payments.isLoading.value"
      :payments="payments.data?.value?.items || []"
    />
    <p v-if="trip.isLoading.value || payments.isLoading.value">Loading...</p>
  </div>
</template>
