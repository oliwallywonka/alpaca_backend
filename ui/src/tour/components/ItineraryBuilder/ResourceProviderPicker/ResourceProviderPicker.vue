<script setup lang="ts">
import { SearchInput } from '@/core/components/fields'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
import { useParams } from '@/core/hooks/useParams'
import { ResourceProviderService } from '@/providerResource/providerResourcesService'
import type { ResourceProvider } from '@/providerResource/ResourceProvider'
import ResourceProviderTable from './ResourceProviderTable.vue'

const emit = defineEmits<{
  (e: 'select:resourceProvider', resource: ResourceProvider | undefined): void
}>()

const { params, setParams } = useParams({
  expand: 'resource, provider, user',
})
const { data, isLoading, refetch } = ResourceProviderService.useGetAll(params)

const handleSearch = async (value: string) => {
  setParams({
    ...params,
    filter: `resource.name.en~"${value}" || user.name~"${value}" || provider.fullName~"${value}" || refPrices~"${value}"`,
  })
  await refetch()
}
</script>
<template>
  <SearchInput @change:debounce="handleSearch" :isLoading="isLoading" />
  <p v-if="isLoading">Loading...</p>
  <ResourceProviderTable
    :data="data?.items || []"
    @select:resourceProvider="emit('select:resourceProvider', $event)"
  />
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
