<script setup lang="ts">
import { useParams } from '@/core/hooks/useParams'
import ExpensesTable from './ExpensesTable.vue'
import { TourResourcesService } from '@/tour/services/tourResources'
import { Button } from '@/core/components/ui/button'
import { RefreshCcwIcon } from 'lucide-vue-next'
import type { TourResource } from '@/tour/interfaces/tourResources'
import { toast } from 'vue-sonner'

const props = defineProps<{
  tourVariantID: string
}>()

const { params } = useParams({
  perPage: 200,
  filter: `tourVariant="${props.tourVariantID}"`,
  expand:
    'resourceProvider.resource, resourceProvider.provider, resourceProvider.user, tourVariant',
})
const { data, isLoading, refetch } = TourResourcesService.useGetAll(params)
const { mutateAsync } = TourResourcesService.useUpdate()

const handleChangeTourResource = async (tourResource: Partial<TourResource>) => {
  try {
    await mutateAsync({
      id: tourResource.id!,
      data: tourResource,
    })
    await refetch()
    toast.success('Tour resource updated successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error updating tour resource')
  }
}
</script>

<template>
  <div>
    <Button @click="refetch"
      ><RefreshCcwIcon :class="{ 'animate-spin': isLoading }" /> Refresh</Button
    >
  </div>
  <p v-if="isLoading" class="text-muted-foreground text-sm">Loading...</p>
  <p v-if="!data || data.items.length === 0" class="text-muted-foreground text-sm">
    No expenses found
  </p>
  <ExpensesTable :resources="data?.items || []" @change:tourResource="handleChangeTourResource" />
</template>
