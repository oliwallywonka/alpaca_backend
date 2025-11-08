<script setup lang="ts">
import { ref } from 'vue'
import { ToggleGroup, ToggleGroupItem } from '@/core/components/ui/toggle-group'
import { Badge } from '@/core/components/ui/badge'

const tripTypes = [
  {
    value: 'all',
    variant: 'outline',
  },
  {
    value: 'draft',
    variant: 'outline',
  },
  {
    value: 'booked',
    variant: 'blue',
  },
  {
    value: 'inProgress',
    variant: 'green',
  },
  {
    value: 'cancelled',
    variant: 'red',
  },
  {
    value: 'finished',
    variant: 'purple',
  },
] as const

const defaultValue = ref('all')
</script>

<template>
  <ToggleGroup v-model="defaultValue" type="single" class="w-full grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-2">
    <ToggleGroupItem
      v-for="type of tripTypes"
      :key="type.value"
      :value="type.value"
      class="grid justify-items-start gap-2 border h-auto p-2"
    >
      <span>Trips</span>
      <Badge :variant="type.variant">{{ type.value }}</Badge>
      <span>Total: 10</span>
    </ToggleGroupItem>
  </ToggleGroup>
</template>
