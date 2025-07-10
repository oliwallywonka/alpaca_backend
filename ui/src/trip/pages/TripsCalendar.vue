<script setup lang="ts">
import { ref } from 'vue'

import FullCalendarComponent from '@fullcalendar/vue3'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import type { CalendarOptions } from '@fullcalendar/core'
import type { EventImpl } from '@fullcalendar/core/internal'

import { Badge } from '@/core/components/ui/badge'
import { ToggleGroup, ToggleGroupItem } from '@/core/components/ui/toggle-group'

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
  events: [],
})
</script>

<template>
  <ToggleGroup type="single" class="w-full grid grid-cols-5 gap-2">
    <ToggleGroupItem value="all" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Travels</span>
      <Badge variant="outline">All</Badge>
      <span>Total: 10</span>
    </ToggleGroupItem>
    <ToggleGroupItem value="booked" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Travels</span>
      <Badge variant="blue">Booked</Badge>
      <span>Total: 10</span>
    </ToggleGroupItem>

    <ToggleGroupItem value="inProgress" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Travels</span>
      <Badge variant="green">In Progress</Badge>
      <span>Total: 10</span>
    </ToggleGroupItem>

    <ToggleGroupItem value="cancelled" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Travels</span>
      <Badge variant="red">Cancelled</Badge>
      <span>Total: 10</span>
    </ToggleGroupItem>

    <ToggleGroupItem value="finished" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Travels</span>
      <Badge variant="purple">Finished</Badge>
      <span>Total: 10</span>
    </ToggleGroupItem>
  </ToggleGroup>

  <section class="min-h-[80vh]">
    <FullCalendarComponent :options="options" />
  </section>
</template>
