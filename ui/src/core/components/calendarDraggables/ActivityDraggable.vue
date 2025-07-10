<script setup lang="ts">
import { onMounted } from 'vue'
import { Draggable } from '@fullcalendar/interaction'
import { Search } from 'lucide-vue-next'

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
import { Button, buttonVariants } from '@/core/components/ui/button'
import { Badge } from '@/core/components/ui/badge'
import { cn } from '@/core/lib/utils'

onMounted(() => {
  new Draggable(document.getElementById('activity-draggable')!, {
    itemSelector: '.activity',
    eventData: function (element) {
      const data = JSON.parse(element.getAttribute('props') ?? '{}')
      return {
        title: 'event 1',
        date: '2025-04-04',
        backgroundColor: '#f2e9fa',
        //borderColor: '#e9e3f0',
        borderColor: '#63209e',
        textColor: '#63209e',
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

  <div id="activity-draggable" class="grid gap-2">

    <div
      v-for="i in 2"
      :key="i"
      :props="JSON.stringify({ title: 'Museo Virgen de Copacabana' })"
      class="activity grid border border-gray-200 rounded-md"
    >
      <div class="grid gap-2 p-4 bg-muted/50">
        <ul class="grid gap-1">
          <li class="text-lg font-bold">Museo Virgen de Copacabana</li>
          <li class="text-sm text-muted-foreground">Description: Ingreso al museo</li>
          <li class="text-sm text-muted-foreground flex flex-wrap gap-0.5">
            Destination: <Badge>La Paz</Badge>
          </li>
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
