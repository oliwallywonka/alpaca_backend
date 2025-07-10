<script setup lang="ts">
import { toast } from 'vue-sonner'
import { PlusIcon, RefreshCwIcon } from 'lucide-vue-next'

import SearchInput from '@/core/components/fields/SearchInput.vue'
import ProvidersTable from '@/provider/components/ProvidersTable.vue'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'

import { ProviderService } from '@/provider/services/providerService'
import { useParams } from '@/core/hooks/useParams'
import ProviderFormDialog from '../components/ProviderFormDialog.vue'

import { Button } from '@/core/components/ui/button'

const { params, setParams } = useParams()
const { data, refetch } = ProviderService.useGetAll(params)
const { mutateAsync } = ProviderService.useUpdate()

const handleStatus = async (id: string, status: boolean) => {
  try {
    await mutateAsync({ id, data: { isActive: status } })
    await refetch()
    toast.success('Provider status updated successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error updating provider status')
  }
}

const handleSearchInput = async (value: string) => {
  setParams({ filter: `fullName~"${value}" || contacts~"${value}"` })
  await refetch()
}
</script>

<template>
  <h1 class="text-xl font-bold">Providers management</h1>
  <SearchInput @change:debounce="handleSearchInput" />
  <div class="flex gap-2">
    <ProviderFormDialog><PlusIcon class="w-4 h-4" />Add Provider</ProviderFormDialog>
    <Button @click="refetch"><RefreshCwIcon /></Button>
  </div>
  <ProvidersTable :providers="data?.items" @click:status="handleStatus" />
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
