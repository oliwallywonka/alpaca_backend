<script setup lang="ts">
import { ref } from 'vue'

import FullCalendarComponent from '@fullcalendar/vue3'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import type { CalendarOptions } from '@fullcalendar/core'
import type { EventImpl } from '@fullcalendar/core/internal'

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/core/components/ui/dialog'
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/core/components/ui/accordion'
import { Button } from '@/core/components/ui/button'
import { Input } from '@/core/components/ui/input'
import { Label } from '@/core/components/ui/label'

import ServiceDraggable from '@/core/components/calendarDraggables/ServiceDraggable.vue'
import DestinationDraggable from '@/core/components/calendarDraggables/DestinationDraggable.vue'
import ActivityDraggable from '@/core/components/calendarDraggables/ActivityDraggable.vue'
import MealDraggable from '@/core/components/calendarDraggables/MealDraggable.vue'
import HotelDraggable from '@/core/components/calendarDraggables/HotelDraggable.vue'
import ServiceItem from '@/core/components/tourItems/ServiceItem.vue'
import ActivityItem from '@/core/components/tourItems/ActivityItem.vue'

const isOpen = ref(false)
const currentEvent = ref<EventImpl | null>(null)
const itineraryOptions = ref<CalendarOptions>({
  plugins: [interactionPlugin, timeGridPlugin, dayGridPlugin],
  initialView: 'customView',
  headerToolbar: {
    left: 'prev,next',
    center: 'title',
    right: 'dayGridMonth,customView',
  },
  scrollTime: '10:00:00',
  views: {
    customView: {
      type: 'timeGrid',
      duration: { days: 1 },
      buttonText: 'days',
    },
  },
  eventReceive: function (info) {
    console.log('RECIEVING EVENT', info.event.toJSON())
  },
  eventResize: function (info) {
    console.log('RESIZING EVENT', info.event.toJSON())
  },
  eventDrop: function (info) {
    console.log('DROPPING EVENT', info.event.toJSON())
    info.event.remove()
  },
  eventClick: function (info) {
    console.log('CLICKING EVENT', info.event.toJSON())
    currentEvent.value = info.event
    isOpen.value = true
  },
  dateClick: function (info) {
    alert('Clicked on: ' + info.dateStr)
  },
  height: '100%',
  editable: true,
  droppable: true,
  events: [],
})

function deleteEvent() {
  currentEvent.value?.remove()
  currentEvent.value = null
  isOpen.value = false
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>{{ currentEvent?.title }}</DialogTitle>
        <DialogDescription>
          Make changes to your profile here. Click save when you're done.pc i7core i7
        </DialogDescription>
      </DialogHeader>
      <div class="grid gap-4 py-4">
        <div class="grid grid-cols-4 items-center gap-4">
          <Label for="name" class="text-right"> Name </Label>
          <Input id="name" value="Pedro Duarte" class="col-span-3" />
        </div>
        <div class="grid grid-cols-4 items-center gap-4">
          <Label for="username" class="text-right"> Username </Label>
          <Input id="username" value="@peduarte" class="col-span-3" />
        </div>
      </div>
      <DialogFooter>
        <Button @click="deleteEvent"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>

  <Accordion type="single" collapsible class="w-full" default-value="item-1">
    <h2 class="font-semibold">All Tour Features</h2>
    <AccordionItem value="item-1">
      <AccordionTrigger>Services</AccordionTrigger>
      <AccordionContent class="grid gap-2">
        <ServiceDraggable />
      </AccordionContent>
    </AccordionItem>

    <AccordionItem value="item-2">
      <AccordionTrigger>Destinations</AccordionTrigger>
      <AccordionContent class="grid gap-2">
        <DestinationDraggable />
      </AccordionContent>
    </AccordionItem>
    <AccordionItem value="item-3">
      <AccordionTrigger>Activities</AccordionTrigger>
      <AccordionContent class="grid gap-2">
        <ActivityDraggable />
      </AccordionContent>
    </AccordionItem>
    <AccordionItem value="item-4">
      <AccordionTrigger>Meals</AccordionTrigger>
      <AccordionContent class="grid gap-2"> <MealDraggable /> </AccordionContent>
    </AccordionItem>
    <AccordionItem value="item-5">
      <AccordionTrigger>Hotels</AccordionTrigger>
      <AccordionContent> <HotelDraggable /></AccordionContent>
    </AccordionItem>
  </Accordion>
  <FullCalendarComponent :options="itineraryOptions" />

  <Accordion type="single" collapsible class="w-full" default-value="item-1">
    <h2 class="font-semibold">Tour Template</h2>
    <AccordionItem value="item-1">
      <AccordionTrigger>Day 1</AccordionTrigger>
      <AccordionContent class="grid gap-2">
        <ServiceItem />
        <ActivityItem />
      </AccordionContent>
    </AccordionItem>

    <AccordionItem value="item-2">
      <AccordionTrigger>Day 2</AccordionTrigger>
      <AccordionContent class="grid gap-2">
        <DestinationDraggable />
      </AccordionContent>
    </AccordionItem>
    <AccordionItem value="item-2">
      <AccordionTrigger>Day 2</AccordionTrigger>
      <AccordionContent class="grid gap-2">
        <DestinationDraggable />
      </AccordionContent>
    </AccordionItem>
  </Accordion>
</template>
