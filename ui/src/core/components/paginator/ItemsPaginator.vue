<script setup lang="ts">
import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationFirst,
  PaginationItem,
  PaginationLast,
  PaginationNext,
  PaginationPrevious,
} from '@/core/components/ui/pagination'

import { Button } from '@/core/components/ui/button'

const emit = defineEmits<{
  (e: 'update:page', page: number): void
}>()

const props = withDefaults(defineProps<{
  itemsPerPage: number
  total: number
}>(), {
  itemsPerPage: 10,
  total: 0,
})

const handleClick = (page: number) => {
  emit('update:page', page)
}
</script>
<template>
  <Pagination
    v-slot="{ page }"
    :items-per-page="props.itemsPerPage"
    :total="props.total"
    show-edges
    :default-page="1"
    @update:page="handleClick"
  >
    <PaginationContent v-slot="{ items }" class="grid md:grid-cols-[auto_1fr_auto] gap-2">
      <div>
        <PaginationFirst />
        <PaginationPrevious />
      </div>
      <div>
        <template v-for="(item, index) in items">
          <PaginationItem v-if="item.type === 'page'" :key="index" :value="item.value" as-child>
            <Button class="w-10 h-10 p-0" :variant="item.value === page ? 'default' : 'outline'">
              {{ item.value }}
            </Button>
          </PaginationItem>
          <PaginationEllipsis v-else :key="item.type" :index="index" />
        </template>
      </div>

      <div>
        <PaginationNext />
        <PaginationLast />
      </div>
    </PaginationContent>
  </Pagination>
</template>
