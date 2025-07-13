<script setup lang="ts">
import { toRef } from 'vue'
import { useRoute } from 'vue-router'
import ServiceProviderFormDialog from '@/provider/components/ResourceProviderFormDialog.vue'
import ResourcesPricesTable from '@/user/components/ResourcesPricesTable/ResourcesPricesTable.vue'
import { PlusIcon, RefreshCwIcon } from 'lucide-vue-next'
import { ProviderService } from '../services/providerService'
import { Button } from '@/core/components/ui/button'

const route = useRoute()
const { data, isLoading, refetch } = ProviderService.useGetResources(
  toRef(route.params.id as string),
)

const { data: providerData } = ProviderService.useGetOne(toRef(route.params.id as string), toRef(true))
</script>

<template>
  <h1 class="text-xl font-semibold">Provider resource Prices {{ providerData?.fullName }}</h1>
  <div class="flex gap-2">
    <ServiceProviderFormDialog :providerID="route.params.id as string">
      <PlusIcon class="w-4 h-4" />
      New Service
    </ServiceProviderFormDialog>
    <Button @click="refetch"><RefreshCwIcon :class="{ 'animate-spin': isLoading }" /> </Button>
  </div>
  <ResourcesPricesTable :resources="data?.items || []" />
</template>
