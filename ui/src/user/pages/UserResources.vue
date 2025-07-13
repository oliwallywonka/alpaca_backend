<script setup lang="ts">
import { toRef } from 'vue'
import { useRoute } from 'vue-router'
import { PlusIcon, RefreshCcwIcon } from 'lucide-vue-next'

import ResourcesPricesTable from '@/user/components/ResourcesPricesTable/ResourcesPricesTable.vue'
import ResourceFormDialog from '@/user/components/ResourcesPricesTable/ResourceFormDialog.vue'
import { UserService } from '../services/userService'
import { Button } from '@/core/components/ui/button'

const route = useRoute()
const { data, isLoading, refetch } = UserService.useGetResources(toRef(route.params.id as string))
const { data: userData } = UserService.useGetOne(toRef(route.params.id as string), toRef(true))
async function handleRefetch() {
  await refetch()
}
</script>

<template>
  <h1 class="text-xl font-semibold">User Resources Prices {{ userData?.name }}</h1>
  <div class="w-full flex gap-2">
    <ResourceFormDialog :userID="route.params.id as string">
      <PlusIcon class="w-4 h-4" />
      New Resource
    </ResourceFormDialog>
    <Button @click="handleRefetch">
      <RefreshCcwIcon :class="{ 'animate-spin': isLoading }" />
    </Button>
  </div>
  <ResourcesPricesTable :resources="data?.items || []" />
</template>
