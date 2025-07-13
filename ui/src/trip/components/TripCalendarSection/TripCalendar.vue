<script setup lang="ts">
import { ref } from 'vue'

import FullCalendarComponent from '@fullcalendar/vue3'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import type { CalendarOptions, EventApi, EventSourceInput } from '@fullcalendar/core'
import type { EventImpl } from '@fullcalendar/core/internal'

const props = defineProps<{
  initialEvents: EventSourceInput
}>()
const emit = defineEmits<{
  (e: 'drop:event', event: EventImpl): void
  (e: 'receive:event', event: EventApi): void
}>()

const travelCalendarOptions = ref<CalendarOptions>({
  plugins: [interactionPlugin, timeGridPlugin, dayGridPlugin],
  initialView: 'dayGridMonth',
  headerToolbar: {
    left: 'prev,next',
    center: 'title',
  },
  eventReceive: function (info) {
    console.log('RECIEVING EVENT', info.event.toJSON())
    emit('receive:event', info.event)
  },
  eventDrop: function (info) {
    console.log('DROPPING EVENT', info.event.toJSON())
    emit('drop:event', info.event)
  },
  height: '100%',
  droppable: true,
  editable: true,
  eventDurationEditable: false,
  eventStartEditable: true,
  events: props.initialEvents,
})
</script>

<template>
  <div class="h-[80vh]">
    <FullCalendarComponent :options="travelCalendarOptions" />
  </div>
</template>
