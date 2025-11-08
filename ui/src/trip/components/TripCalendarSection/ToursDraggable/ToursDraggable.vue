<script setup lang="ts">
import { SearchInput } from '@/core/components/fields'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
import ToursTable from './ToursTable.vue'
import { useParams } from '@/core/hooks/useParams'
import { ItineraryService } from '@/itinerary/services/itineraryService'

const { params, setParams } = useParams({
  expand: `tour`,
})
const { data, refetch } = ItineraryService.useGetSummaries(params)
const handleDebounce = async (value: string) => {
  setParams({ filter: `tour.name.en~"${value}" || minPersons~"${value}" || maxPersons~"${value}"` })
  await refetch()
}
</script>

<template>
  <div class="grid gap-4">
    <h1 class="text-xl font-semibold">Lista de tours</h1>
    <SearchInput @change:debounce="handleDebounce" />
    <ToursTable :data="data?.items || []" />

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
