<script setup lang="ts">
import { RefreshCcwIcon } from 'lucide-vue-next'

import { useParams } from '@/core/hooks/useParams'
import SearchInput from '@/core/components/fields/SearchInput.vue'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
import { TripService } from '@/trip/services/TripService'
import TripsTable from './TripsTable.vue'
import TripFormDialog from '@/trip/pages/TripsPage/components/TripFormDialog.vue'
import { Button } from '@/core/components/ui/button'
import type { TripSummary } from '@/trip/interfaces/trip'

const props = defineProps<{
  showSelect?: boolean
}>()

const emit = defineEmits<{
  (e: 'select:trip', trip: TripSummary | undefined): void
}>()

const { params, setParams } = useParams({
  expand: 'tour',
})
const { data, isLoading, isRefetching, refetch } = TripService.useGetSummaries(params)
const handleSearch = (value: string) => {
  setParams({
    ...params,
    filter: `tour.name.en~"${value}"`,
  })
  refetch()
}
</script>

<template>
  <div class="grid gap-2">
    <SearchInput @change:debounce="handleSearch" :isLoading="isLoading || isRefetching" />
    <div class="flex gap-2">
      <TripFormDialog />
      <Button @click="async () => await refetch({})">
        <RefreshCcwIcon :class="{ 'animate-spin': isLoading || isRefetching }" /> Refresh
      </Button>
    </div>
    <p v-if="isLoading">Loading...</p>
    <TripsTable
      v-if="!isLoading"
      :trips="data?.items || []"
      :showSelect="props.showSelect"
      @select:trip="(e) => emit('select:trip', e)"
    />
    <ItemsPaginator
      :itemsPerPage="params.perPage"
      :total="data?.totalItems || 0"
      @update:page="
        async ($event) => {
          setParams({ page: $event })
          await refetch()
        }
      "
    />
  </div>
</template>
