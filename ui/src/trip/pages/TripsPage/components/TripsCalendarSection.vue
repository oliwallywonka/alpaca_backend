<script setup lang="ts">
import { ref } from 'vue'

import FullCalendarComponent from '@fullcalendar/vue3'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import type { CalendarOptions } from '@fullcalendar/core'
import { colors } from '@/core/constants/colors.ts'
import type { EventImpl } from '@fullcalendar/core/internal'

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
  <FullCalendarComponent :options="options" />
</template>
