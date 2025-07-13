<script setup lang="ts">
import { PlusIcon, RefreshCcwIcon } from 'lucide-vue-next'
import { toast } from 'vue-sonner'

import { SearchInput } from '@/core/components/fields'
import DestinationFormDialog from './components/DestinationFormDialog.vue'
import DestinationsTable from './components/DestinationsTable.vue'
import { DestinationService } from '@/destination/services/DestinationService'
import { useParams } from '@/core/hooks/useParams'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
import { Button } from '@/core/components/ui/button'
import type { Destination } from './interfaces/destination'

const props = defineProps<{
  showSelect?: boolean
}>()

const emit = defineEmits<{
  (e: 'select:destination', destination: Destination | undefined): void
}>()

const { params, setParams } = useParams()
const { data, isLoading, refetch } = DestinationService.useGetAll(params)
const { mutateAsync } = DestinationService.useUpdate()

const handleSearch = (value: string) => {
  setParams({
    ...params,
    filter: `name~"${value}" || description~"${value}" || parent.name~"${value}"`,
  })
  refetch()
}

const handleSelect = (destination: Destination | undefined) => {
  emit('select:destination', destination)
}

const updateStatus = async (id: string, status: boolean) => {
  try {
    await mutateAsync({ id, data: { isActive: status } })
    await refetch()
    toast.message('Destination updated successifully')
  } catch (e) {
    console.log(e)
    toast.error('Error updating destination')
  }
}
</script>

<template>
  <SearchInput @change:debounce="handleSearch" :isLoading="isLoading" />
  <div class="flex gap-2">
    <DestinationFormDialog><PlusIcon class="w-4 h-4" /> New Destination</DestinationFormDialog>

    <Button @click="refetch" :disabled="isLoading">
      <RefreshCcwIcon :class="{ 'animate-spin': isLoading }" />
    </Button>
  </div>
  <p v-if="isLoading">Loading...</p>
  <template v-else>
    <DestinationsTable
      :destinations="data?.items || []"
      :showSelect="props.showSelect"
      @select:destination="handleSelect"
      @click:status="updateStatus"
    />
  </template>
  <ItemsPaginator
    :items-per-page="params.perPage"
    :total="data?.totalItems || 0"
    @update:page="
      async (page) => {
        setParams({ page })
        await refetch()
      }
    "
  />
</template>
