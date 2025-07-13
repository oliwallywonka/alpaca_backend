<script setup lang="ts">
import { PlusIcon, RefreshCcwIcon } from 'lucide-vue-next'
import CustomersTable from '@/customer/components/CustomersTable.vue'
import { CustomerService } from '../services/CustomerService'
import { useParams } from '@/core/hooks/useParams'
import SearchInput from '@/core/components/fields/SearchInput.vue'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
import { Button } from '@/core/components/ui/button'
import CustomerFormDIalog from '../components/CustomerFormDIalog.vue'

const { params, setParams } = useParams()
const { data, isLoading, refetch } = CustomerService.useGetAll(params)

const handleSearch = async (value: string) => {
  setParams({
    ...params,
    filter: `title~"${value}" || firstName~"${value}" || middleName~"${value}" || lastName~"${value}" || dateOfBirth~"${value}"`,
  })
  await refetch()
}
</script>

<template>
  <h1 class="text-xl font-bold">Customers</h1>
  <SearchInput @change:debounce="handleSearch" :isLoading="isLoading" />
  <div class="flex gap-2">
    <CustomerFormDIalog><PlusIcon /> New</CustomerFormDIalog>
    <Button @click="refetch" :disabled="isLoading">
      <RefreshCcwIcon :class="{ 'animate-spin': isLoading }" />
    </Button>
  </div>
  <p v-if="isLoading" class="text-muted-foreground text-sm">Loading...</p>
  <p v-if="!data || data.items.length === 0" class="text-muted-foreground text-sm">
    No customers found
  </p>
  <CustomersTable :customers="data?.items || []" />
  <ItemsPaginator :itemsPerPage="params.perPage" :total="data?.totalItems || 0" />
</template>
