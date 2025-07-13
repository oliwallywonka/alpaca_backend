<script setup lang="ts">
import { SearchInput } from '@/core/components/fields'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
import ToursTable from './ToursTable.vue'
import { useParams } from '@/core/hooks/useParams'
import { TourVariantService } from '@/tour/services/tourVariantService'

const { params, setParams } = useParams({
  expand: 'tour',
})
const { data, refetch } = TourVariantService.useGetAll(params)
</script>

<template>
  <div class="grid gap-4">
    <h1 class="text-xl font-semibold">Lista de tours</h1>
    <SearchInput />
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
