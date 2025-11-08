<script setup lang="ts">
import { SearchInput } from '@/core/components/fields'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
import { useParams } from '@/core/hooks/useParams'
import { ItineraryService } from '@/itinerary/services/itineraryService'
import TourItineraryTable from './TourItineraryTable.vue'
import type { Itinerary } from '@/itinerary/interfaces/itinerary'

const emit = defineEmits<{
  (e: 'select:itinerary', itinerary: Itinerary | undefined): void
}>()

const { params, setParams } = useParams({
  expand: 'tour',
  filter: 'isTemplate=true',
})
const { data, refetch } = ItineraryService.useGetSummaries(params)
</script>

<template>
  <div class="grid gap-4">
    <h1 class="text-xl font-semibold">Lista de tours</h1>
    <SearchInput />
    <TourItineraryTable
      :data="data?.items || []"
      @select:itinerary="emit('select:itinerary', $event)"
    />

    <ItemsPaginator
      :total="data?.totalItems || 0"
      :itemsPerPage="params.perPage"
      @update:page="
        async ($event) => {
          setParams({ page: $event })
          await refetch()
        }
      "
    />
  </div>
</template>
