<script setup lang="ts">
import {
  Pagination,
  PaginationEllipsis,
  PaginationFirst,
  PaginationLast,
  PaginationContent,
  PaginationItem,
  PaginationNext,
  PaginationPrevious,
} from '@/core/components/ui/pagination'
import { Input } from '@/core/components/ui/input'
import { Search } from 'lucide-vue-next'
import { Button, buttonVariants } from '@/core/components/ui/button'
import { cn } from '@/core/lib/utils'
import { onMounted } from 'vue'
import { Draggable } from '@fullcalendar/interaction'

onMounted(() => {
  new Draggable(document.getElementById('hotel-draggable')!, {
    itemSelector: '.hotel',
    eventData: function (element) {
      const data = JSON.parse(element.getAttribute('props') ?? '{}')
      return {
        title: 'event 1',
        date: '2025-04-04',
        backgroundColor: '#e8f9ef',
        //borderColor: '#ddf5e9',
        borderColor: '#14a051',
        textColor: '#14a051',
        ...data,
      }
    },
  })
})
</script>

<template>
  <div class="relative w-full items-center">
    <Input id="search" type="text" placeholder="Search..." class="pl-10" />
    <span class="absolute start-0 inset-y-0 flex items-center justify-center px-2">
      <Search class="size-6 text-muted-foreground" />
    </span>
  </div>

  <div id="hotel-draggable" class="grid gap-2">
    <div
      v-for="i in 2"
      :key="i"
      class="hotel grid border border-gray-200 rounded-md"
      :props="JSON.stringify({ title: 'Hotel Fuente de la Cruz' })"
    >
      <div class="grid gap-2 p-4 bg-muted/50">
        <ul class="grid gap-1">
          <li class="text-lg font-bold">Hotel Fuente de la Cruz</li>
          <li class="text-sm text-muted-foreground">
            Direction: Copacabana calle Alonzo de Ibáñez #44
          </li>
          <li class="text-sm text-muted-foreground">Phone: +56 9123456789</li>
        </ul>
        <div class="flex flex-wrap gap-2">
          <RouterLink to="/tours/1" :class="cn(buttonVariants())">Edit</RouterLink>
        </div>
      </div>
    </div>
  </div>

  <Pagination
    v-slot="{ page }"
    :items-per-page="2"
    :total="80"
    :sibling-count="1"
    show-edges
    :default-page="1"
  >
    <PaginationContent v-slot="{ items }">
      <PaginationFirst />
      <PaginationPrevious />
      <template v-for="(item, index) in items">
        <PaginationItem v-if="item.type === 'page'" :key="index" :value="item.value" as-child>
          <Button class="w-10 h-10 p-0" :variant="item.value === page ? 'default' : 'outline'">
            {{ item.value }}
          </Button>
        </PaginationItem>
        <PaginationEllipsis v-else :key="item.type" :index="index" />
      </template>

      <PaginationNext />
      <PaginationLast />
    </PaginationContent>
  </Pagination>
</template>
