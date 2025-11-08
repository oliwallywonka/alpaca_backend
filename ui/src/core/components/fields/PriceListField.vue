<script setup lang="ts">
import { useVModel } from '@vueuse/core'
import { Switch } from '@/core/components/ui/switch'
import { Badge } from '@/core/components/ui/badge'
import type { PriceField } from '@/core/interfaces/fields'

const props = withDefaults(
  defineProps<{
    modelValue: PriceField | null
    priceList: PriceField[]
  }>(),
  {
    modelValue: null,
    priceList: () => [],
  },
)
const emit = defineEmits<{ (e: 'update:modelValue', value: PriceField): void }>()
const data = useVModel(props, 'modelValue', emit, { passive: true })

function includesPrice(item: PriceField) {
  if (!data.value) return false
  return data.value?.minPersons === item.minPersons && data.value?.maxPersons === item.maxPersons
}

function togglePrice(item: PriceField) {
  data.value = item
}
</script>
<template>
  <ul class="flex flex-col gap-2">
    <li
      v-for="(item, index) in props.priceList"
      :key="`${item.minPersons}-${item.maxPersons}-${index}`"
      class="flex items-center gap-2"
    >
      <Switch :modelValue="includesPrice(item)" @update:modelValue="togglePrice(item)" />
      <Badge>{{ `Persons min: ${item.minPersons} max: ${item.maxPersons}` }}</Badge>
      <Badge>{{ `Price: ${item.currency} ${item.price}` }}</Badge>
    </li>
  </ul>
</template>
