<script setup lang="ts">
import { onMounted, ref, useTemplateRef, watch } from 'vue'

import { Calendar, type EventSourceInput, type EventInput } from '@fullcalendar/core'
import type { EventImpl } from '@fullcalendar/core/internal'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin, { type DateClickArg } from '@fullcalendar/interaction'

import { Button } from '@/core/components/ui/button'
import { ChevronLeft, ChevronRight } from 'lucide-vue-next'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/core/components/ui/select'

const props = withDefaults(
  defineProps<{
    events: EventSourceInput
    isEditable?: boolean
    scrollContainer?: string
  }>(),
  {
    isEditable: true,
  },
)

const emit = defineEmits<{
  (e: 'click:event', event: EventImpl): void
  (e: 'click:date', date: DateClickArg): void
  (e: 'event:drop', event: EventImpl): void
  (e: 'event:resize', event: EventImpl): void
}>()

const calendarRef = useTemplateRef<HTMLElement | null>('calendarRef')
const calendar = ref<Calendar | null>(null)
const currentDate = ref<string>('')
const currentView = ref<'dayGridMonth' | 'timeGridWeek'>('dayGridMonth')

watch(
  calendarRef,
  (newCalendar) => {
    if (!newCalendar) return
    calendar.value = new Calendar(newCalendar, {
      plugins: [interactionPlugin, timeGridPlugin, dayGridPlugin],
      initialView: currentView.value,
      headerToolbar: false,
      initialDate: getFirstDate(props.events as EventInput[]),
      events: props.events,
      editable: true,
      droppable: true,
      eventDurationEditable: props.isEditable,
      eventStartEditable: true,
      fixedMirrorParent: document.body,
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
      datesSet() {
        currentDate.value = formatDateRange(calendar.value?.getDate() || new Date())
      },
    })
    calendar.value.render()
  },
  {
    immediate: true,
  },
)

watch(
  () => props.events,
  (newEvents) => {
    if (!calendar.value) return
    calendar.value.setOption('events', newEvents)
    calendar.value.gotoDate(getFirstDate(newEvents as EventInput[]))
  },
)

onMounted(() => {
  if (!calendar.value) return
  calendar.value.destroy()
  calendar.value = null
})

const handleNext = () => {
  if (!calendar.value) return
  calendar.value.next()
}

const handlePrev = () => {
  if (!calendar.value) return
  calendar.value.prev()
}

const changeView = (view: 'dayGridMonth' | 'timeGridWeek') => {
  if (!calendar.value) return
  calendar.value.changeView(view)
  currentView.value = view
}

function formatDateRange(date: Date): string {
  const opts = { month: 'long' } as const

  const dateAdjusted = new Date(date.getTime()) // end is exclusive
  const dateStr = date.toLocaleDateString('en-US', opts)

  return `${dateStr} - ${dateAdjusted.getFullYear()}`
}

function getFirstDate(events: EventInput[]) {
  if (events.length === 0) return new Date()
  const sortedEvents = events.slice().sort((a, b) => {
    const aStart =
      a.start instanceof Date ? a.start.getTime() : typeof a.start === 'number' ? a.start : 0
    const bStart =
      b.start instanceof Date ? b.start.getTime() : typeof b.start === 'number' ? b.start : 0
    return aStart - bStart
  })
  return sortedEvents[0]?.start ?? new Date()
}
</script>

<template>
  <section class="relative grid gap-2">
    <div class="flex justify-between items-center gap-2">
      <div class="flex items-center gap-2">
        <Button @click="handlePrev">
          <ChevronLeft />
        </Button>
        <Button @click="handleNext">
          <ChevronRight />
        </Button>
        <Button @click="() => calendar?.gotoDate(new Date())"> Today </Button>
        <Button
          @click="
            () =>
              Array.isArray(props.events) &&
              calendar?.gotoDate(getFirstDate(props.events as EventInput[]))
          "
        >
          Go first
        </Button>
        <span class="text-xl xl:text-3xl font-semibold">
          {{ currentDate }}
        </span>
      </div>
      <Select @update:model-value="(view) => changeView(view as 'dayGridMonth' | 'timeGridWeek')">
        <SelectTrigger>
          <SelectValue :placeholder="currentView" />
        </SelectTrigger>
        <SelectContent>
          <SelectGroup>
            <SelectItem value="dayGridMonth"> Month </SelectItem>
            <SelectItem value="timeGridWeek"> Week </SelectItem>
          </SelectGroup>
        </SelectContent>
      </Select>
    </div>
    <div class="h-[70vh]" ref="calendarRef" />
  </section>
</template>
