<script setup lang="ts">
import { computed } from 'vue'

import type { CalendarOptions, EventSourceInput } from '@fullcalendar/core'
import type { EventImpl } from '@fullcalendar/core/internal'

import FullCalendarComponent from '@fullcalendar/vue3'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin, { type DateClickArg } from '@fullcalendar/interaction'

const props = defineProps<{
  commentaries: EventSourceInput[]
  destinations: EventSourceInput[]
  resources: EventSourceInput[]
  initialDate?: Date
}>()

const emit = defineEmits<{
  (e: 'click:event', event: EventImpl): void
  (e: 'click:date', date: DateClickArg): void
  (e: 'event:drop', event: EventImpl): void
  (e: 'event:resize', event: EventImpl): void
}>()

const options = computed<CalendarOptions>(() => ({
  plugins: [interactionPlugin, timeGridPlugin, dayGridPlugin],
  initialView: 'customView',
  headerToolbar: {
    left: 'prev,next today',
    center: 'title',
    right: 'dayGridMonth,customView',
  },
  initialDate: props.initialDate,
  views: {
    customView: {
      type: 'timeGrid',
      duration: { days: 3 },
      buttonText: 'days',
    },
  },
  events: [...props.commentaries, ...props.destinations, ...props.resources],
  height: '100%',
  editable: true,
  droppable: true,
  /* eventReceive: function (info) {
    console.log('RECIEVING EVENT', info.event.toJSON())
  }, */
  eventResize: function (info) {
    console.log('RESIZING EVENT', info.event.toJSON())
    emit('event:resize', info.event as EventImpl)
  },
  eventDrop: function (info) {
    console.log('DROPPING EVENT', info.event.toJSON())
    emit('event:drop', info.event as EventImpl)
  },
  eventClick: function (info) {
    console.log('CLICKING EVENT', info.event.toJSON())
    emit('click:event', info.event as EventImpl)
  },
  dateClick: function (info) {
    console.log('CLICKING DATE', info)
    emit('click:date', info)
  },
}))
</script>

<template>
  <FullCalendarComponent :options="options" />
</template>
