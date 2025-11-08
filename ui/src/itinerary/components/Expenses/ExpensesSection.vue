<script setup lang="ts">
import { toast } from 'vue-sonner'
import { useQueryClient } from '@tanstack/vue-query'
import { RefreshCcwIcon } from 'lucide-vue-next'

import { useParams } from '@/core/hooks/useParams'
import { Button } from '@/core/components/ui/button'
import { ItineraryResourceService } from '@/itinerary/services/itineraryResources'
import type { ItineraryResource } from '@/itinerary/interfaces/itineraryResource'
import ExpensesTable from './ExpensesTable.vue'

const props = defineProps<{
  itineraryID: string
}>()

const query = useQueryClient()

const { params } = useParams({
  perPage: 200,
  filter: `itinerary="${props.itineraryID}"`,
  expand: 'resourceProvider.resource, resourceProvider.provider, resourceProvider.user, itinerary',
})
const { data, isLoading, isRefetching, refetch } = ItineraryResourceService.useGetAll(params)
const { mutateAsync } = ItineraryResourceService.useUpdate()

const handleChangeTourResource = async (itineraryResource: Partial<ItineraryResource>) => {
  try {
    await mutateAsync({
      id: itineraryResource.id!,
      data: itineraryResource,
    })
    await query.invalidateQueries({ queryKey: ['itineraryResources'] })
    await query.invalidateQueries({ queryKey: ['itineraries'] })
    toast.success('Tour resource updated successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error updating tour resource')
  }
}
</script>

<template>
  <div>
    <Button @click="refetch">
      <RefreshCcwIcon :class="{ 'animate-spin': isLoading || isRefetching }" /> Refresh
    </Button>
  </div>
  <p v-if="isLoading" class="text-muted-foreground text-sm">Loading...</p>
  <p v-if="!data || data.items.length === 0" class="text-muted-foreground text-sm">
    No expenses found
  </p>
  <ExpensesTable
    :resources="data?.items || []"
    @change:itineraryResource="handleChangeTourResource"
  />
</template>
