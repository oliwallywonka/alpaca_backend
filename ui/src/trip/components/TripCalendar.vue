<script setup lang="ts">
import { ref } from 'vue'
import { Search } from 'lucide-vue-next'

import FullCalendarComponent from '@fullcalendar/vue3'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import type { CalendarOptions } from '@fullcalendar/core'
import type { EventImpl } from '@fullcalendar/core/internal'

import { Badge } from '@/core/components/ui/badge'
import { Input } from '@/core/components/ui/input'
import { Button, buttonVariants } from '@/core/components/ui/button'
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
import { cn } from '@/core/lib/utils'

const currentEvent = ref<EventImpl | null>(null)

const travelCalendarOptions = ref<CalendarOptions>({
  plugins: [interactionPlugin, timeGridPlugin, dayGridPlugin],
  initialView: 'dayGridMonth',
  headerToolbar: {
    left: 'prev,next',
    center: 'title',
  },
  eventReceive: function (info) {
    console.log('RECIEVING EVENT', info.event.toJSON())
  },
  eventResize: function (info) {
    console.log('RESIZING EVENT', info.event.toJSON())
  },
  eventDrop: function (info) {
    console.log('DROPPING EVENT', info.event.toJSON())
  },
  eventClick: function (info) {
    console.log('CLICKING EVENT', info.event.toJSON())
    currentEvent.value = info.event
  },
  dateClick: function (info) {
    alert('Clicked on: ' + info.dateStr)
  },
  height: '100%',
  editable: true,
  eventDurationEditable: false,
  droppable: true,
  events: [],
})
</script>

<template>
  <FullCalendarComponent :options="travelCalendarOptions" />

  <div class="grid gap-4">
    <div class="relative w-full items-center">
      <Input id="search" type="text" placeholder="Search..." class="pl-10" />
      <span class="absolute start-0 inset-y-0 flex items-center justify-center px-2">
        <Search class="size-6 text-muted-foreground" />
      </span>
    </div>

    <div id="travel-draggable" class="grid gap-4">
      <div class="travel grid grid-cols-[150px_1fr] border border-gray-200 rounded-md">
        <figure class="h-full">
          <img
            class="object-cover object-center w-full h-full rounded-l-md"
            src="https://images.unsplash.com/photo-1554629947-334ff61d85dc?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=387&q=80"
            alt="Tour Image"
          />
        </figure>
        <div class="grid gap-2 p-4 bg-muted/50">
          <ul class="grid gap-1">
            <li class="text-lg font-bold">Tour Name</li>
            <li class="text-sm text-muted-foreground">Slug:</li>
            <li class="text-sm text-muted-foreground flex flex-wrap gap-0.5">
              Destinations: <Badge v-for="(_, index) in 8" :key="index">La Paz</Badge>
            </li>
            <li class="text-sm text-muted-foreground">Short Description:</li>
            <li class="text-sm text-muted-foreground">Long Description:</li>
            <li class="text-sm text-muted-foreground">Ref price pp:</li>
            <li class="text-sm text-muted-foreground">Discount:</li>
            <li class="text-sm text-muted-foreground">Days:</li>
            <li class="text-sm text-muted-foreground">Group Size:</li>
            <li class="text-sm text-muted-foreground">Transportation:</li>
            <li class="text-sm text-muted-foreground">Accommodation:</li>
            <li class="text-sm text-muted-foreground">Team:</li>
            <li class="text-sm text-muted-foreground">Public:</li>
          </ul>
          <div class="flex flex-wrap gap-2">
            <Button>Itinerary</Button>
            <RouterLink to="/tours/1" :class="cn(buttonVariants())">Edit</RouterLink>
            <Button>Delete</Button>
          </div>
        </div>
      </div>
    </div>

    <Pagination
      v-slot="{ page }"
      :items-per-page="10"
      :total="80"
      :sibling-count="10"
      show-edges
      :default-page="1"
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
  </div>
</template>
