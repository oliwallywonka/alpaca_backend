<script setup lang="ts">
import { ref } from 'vue'
import { Search } from 'lucide-vue-next'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableHeader,
  TableCaption,
} from '@/core/components/ui/table'
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
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/core/components/ui/tabs'
import { Badge } from '@/core/components/ui/badge'
import { Button, buttonVariants } from '@/core/components/ui/button'
import { ToggleGroup, ToggleGroupItem } from '@/core/components/ui/toggle-group'
import { Input } from '@/core/components/ui/input'

import { cn } from '@/core/lib/utils'
import FullCalendarComponent from '@fullcalendar/vue3'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import type { CalendarOptions } from '@fullcalendar/core'
import type { EventImpl } from '@fullcalendar/core/internal'
import { colors } from '@/core/constants/colors'

const currentEvent = ref<EventImpl | null>(null)

const options = ref<CalendarOptions>({
  plugins: [interactionPlugin, timeGridPlugin, dayGridPlugin],
  initialView: 'dayGridMonth',
  headerToolbar: {
    left: 'prev,next today',
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
  datesSet: function (info) {
    console.log('DATES SET', info)
  },
  height: '100%',
  editable: true,
  eventDurationEditable: false,
  droppable: true,
  events: [
    {
      title: 'Tour Uyuni 3 days - Roberto Figueroa',
      start: '2025-05-10T10:00:00',
      end: '2025-05-13T10:00:00',
      ...colors.blue,
    },
    {
      title: 'Tour Copacabana Simple - Miguel Angel',
      start: '2025-05-09T10:00:00',
      end: '2025-05-10T10:00:00',
      ...colors.green,
    },
    {
      title: 'Tour Valle de la Luna - Alberto Garcia',
      start: '2025-05-01T10:00:00',
      end: '2025-05-05T10:00:00',
      ...colors.purple,
    },
  ],
})
</script>

<template>
  <ToggleGroup type="single" class="w-full grid grid-cols-5 gap-2">
    <ToggleGroupItem value="all" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Trips</span>
      <Badge variant="outline">All</Badge>
      <span>Total: 10</span>
    </ToggleGroupItem>
    <ToggleGroupItem value="booked" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Trips</span>
      <Badge variant="blue">Booked</Badge>
      <span>Total: 4</span>
    </ToggleGroupItem>

    <ToggleGroupItem value="inProgress" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Trips</span>
      <Badge variant="green">In Progress</Badge>
      <span>Total: 2</span>
    </ToggleGroupItem>

    <ToggleGroupItem value="cancelled" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Trips</span>
      <Badge variant="red">Cancelled</Badge>
      <span>Total: 5</span>
    </ToggleGroupItem>

    <ToggleGroupItem value="finished" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Trips</span>
      <Badge variant="purple">Finished</Badge>
      <span>Total: 3</span>
    </ToggleGroupItem>
  </ToggleGroup>

  <RouterLink to="/travels/new" :class="cn(buttonVariants(), 'w-48')">+ New Trip</RouterLink>

  <Tabs default-value="account">
    <TabsList>
      <TabsTrigger value="account"> Trips List </TabsTrigger>
      <TabsTrigger value="password"> Trips Calendar </TabsTrigger>
    </TabsList>
    <TabsContent value="account">
      <section class="grid gap-4">
        <div class="relative w-full items-center">
          <Input id="search" type="text" placeholder="Search..." class="pl-10" />
          <span class="absolute start-0 inset-y-0 flex items-center justify-center px-2">
            <Search class="size-6 text-muted-foreground" />
          </span>
        </div>
        <Table class="min-w-[800px] w-full">
          <TableCaption>Trips registered</TableCaption>
          <TableHeader>
            <TableRow>
              <TableHead>Tour</TableHead>
              <TableHead>Customer Lead</TableHead>
              <TableHead>Nº of persons</TableHead>
              <TableHead>Booking date</TableHead>
              <TableHead>Tour Start</TableHead>
              <TableHead>Tour End</TableHead>
              <TableHead>State</TableHead>
              <TableHead>Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="index in 4" :key="index">
              <TableCell>Uyuni Salar 4 Days</TableCell>
              <TableCell>Roberto Alonso</TableCell>
              <TableCell>5</TableCell>
              <TableCell>06/06/2022: 10:00</TableCell>
              <TableCell>06/06/2022: 10:00</TableCell>
              <TableCell>06/06/2022: 10:00</TableCell>
              <TableCell>
                <Badge variant="green">Confirmed</Badge>
              </TableCell>
              <TableCell class="grid gap-2">
                <RouterLink :to="`/trips/${1}`" :class="cn(buttonVariants())"> Details </RouterLink>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>

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
                <PaginationItem
                  v-if="item.type === 'page'"
                  :key="index"
                  :value="item.value"
                  as-child
                >
                  <Button
                    class="w-10 h-10 p-0"
                    :variant="item.value === page ? 'default' : 'outline'"
                  >
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
      </section>
    </TabsContent>
    <TabsContent value="password" class="min-h-[80vh]">
      <FullCalendarComponent :options="options" />
    </TabsContent>
  </Tabs>
</template>
