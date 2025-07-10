<script setup lang="ts">
import { toast } from 'vue-sonner'
import { PlusIcon, RefreshCcwIcon } from 'lucide-vue-next'
import { Button } from '@/core/components/ui/button'
import ResourceTable from '@/resource/components/ResourcesTable.vue'
import { ResourceService } from './services/resourceService'
import { useParams } from '@/core/hooks/useParams'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
import ResourceFormDialog from './components/ResourceFormDialog.vue'
import SearchInput from '@/core/components/fields/SearchInput.vue'

const { params, setParams } = useParams()
const { data, isLoading, refetch } = ResourceService.useGetAll(params)
const { mutateAsync } = ResourceService.useUpdate()

const handleStatus = async (id: string, status: boolean) => {
  try {
    await mutateAsync({ id, data: { isActive: status } })
    await refetch()
    toast.success(`Service status updated to ${status ? 'active' : 'inactive'}`)
  } catch (error) {
    console.error('Error updating service status:', error)
    toast.error('Error updating service status')
  }
}

const handleSearch = async (filter: string) => {
  setParams({ filter: `name~"${filter}" || description.en~"${filter}" || types~"${filter}"` })
  await refetch()
}
</script>

<template>
  <h1 class="text-xl font-semibold">Resources list</h1>
  <SearchInput @change:debounce="handleSearch" />
  <p v-if="isLoading" class="text-gray-500">Loading...</p>
  <div class="flex gap-2">
    <ResourceFormDialog>
      <PlusIcon class="w-4 h-4" />
      New resource
    </ResourceFormDialog>
    <Button @click="refetch" :disabled="isLoading">
      <RefreshCcwIcon :class="{ 'animate-spin': isLoading }" />
    </Button>
  </div>
  <ResourceTable :services="data?.items" @click:status="handleStatus" />
  <ItemsPaginator
    :itemsPerPage="params.perPage"
    :total="data?.totalItems || 0"
    @update:page="
      async (page) => {
        setParams({ page })
        await refetch()
      }
    "
  />
</template>
