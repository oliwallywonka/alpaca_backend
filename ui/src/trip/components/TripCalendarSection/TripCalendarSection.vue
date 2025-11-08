<script setup lang="ts">
import type { EventApi } from '@fullcalendar/core'

import ToursDraggable from './ToursDraggable/ToursDraggable.vue'
import TripCalendar from './TripCalendar.vue'
import { ref } from 'vue'

const currentEvent = ref()
function handleOnReceive(event: EventApi) {
  console.log('EVENT RECEIVED', event)
}
</script>

<template>
  <div class="grid grid-cols-2 gap-4 items-start">
    <TripCalendar :initialEvents="[]" @receive:event="handleOnReceive" />
    <ToursDraggable />
  </div>
</template>
