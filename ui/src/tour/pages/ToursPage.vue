<script setup lang="ts">
import { SearchInput } from '@/core/components/fields'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
import ToursTable from '../components/ToursTable.vue'
import NewTourForm from '../components/NewTourForm.vue'
import { TourService } from '../services/tourService'
import { useParams } from '@/core/hooks/useParams'

const { params, setParams } = useParams()
const { data, isLoading, refetch } = TourService.useGetAll(params)

const handleSearch = async (value: string) => {
  setParams({ filter: `code~"${value}" || name.en~"${value}" || slug.en~"${value}"` })
  await refetch()
}
</script>

<template>
  <div class="grid gap-4">
    <h1 class="text-xl font-semibold">Lista de tours</h1>

    <SearchInput @change:debounce="handleSearch" :isLoading="isLoading" />
    <NewTourForm />
    <ToursTable :tours="data?.items || []" />

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
