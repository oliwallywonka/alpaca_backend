<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { toast } from 'vue-sonner'
import type { DateClickArg } from '@fullcalendar/interaction/index.js'
import type { EventImpl } from '@fullcalendar/core/internal'

import { useParams } from '@/core/hooks/useParams'
import { colors } from '@/core/constants/colors'
import FullCalendar from '@/core/components/calendar/FullCalendar.vue'
import { ItineraryCommentaryService } from '@/itinerary/services/itineraryCommentaryService'
import { ItineraryResourceService } from '@/itinerary/services/itineraryResources'
import { ItineraryDestinationService } from '@/itinerary/services/ItineraryDestinationService'
import type { ItineraryCommentary } from '@/itinerary/interfaces/itineraryCommentary'
import type { ItineraryDestination } from '@/itinerary/interfaces/itineraryDestination'
import type { ItineraryResource } from '@/itinerary/interfaces/itineraryResource'
import EventTypeDialog from '@/itinerary/components/Dialogs/EventTypeDialog.vue'
import CommentaryFormDialog from '@/itinerary/components/Dialogs/CommentaryFormDialog.vue'
import ResourceFormDialog from '@/itinerary/components/Dialogs/ResourceFormDialog.vue'
import DestinationFormDialog from '@/itinerary/components/Dialogs/DestinationFormDialog.vue'
import DestinationDeleteDialog from '@/itinerary/components/Dialogs/DestinationDeleteDialog.vue'
import ResourceDetailsDialog from '@/itinerary/components/Dialogs/ResourceDetailsDialog.vue'

const props = defineProps<{
  itineraryID: string
  tourID: string
}>()

const tourID = computed(() => props.tourID)
const itineraryID = computed(() => props.itineraryID)
const currentEvent = ref<EventImpl | null>(null)
const currentDate = ref<DateClickArg>()
const formattedDate = computed(() => {
  if (currentEvent.value && currentEvent.value.start) return currentEvent.value.start.toISOString()
  if (!currentDate.value) return new Date().toISOString()
  return currentDate.value.date.toISOString()
})

const eventTypeDialog = ref(false)
const commentaryDialog = ref({ form: false, delete: false })
const destinationDialog = ref({ form: false, delete: false })
const resourceDialog = ref({ form: false, details: false })

const { params: commentaryParams, setParams: setCommentaryParams } = useParams({
  filter: `itinerary="${itineraryID.value}"`,
  orderBy: '+startDate',
})
const commentaries = ItineraryCommentaryService.useGetAll(commentaryParams)
const updateCommentary = ItineraryCommentaryService.useUpdate()

const { params: destinationParms, setParams: setDestinationParams } = useParams({
  filter: `itinerary="${itineraryID.value}"`,
  orderBy: '+startDate',
  expand: 'destination',
})
const destinations = ItineraryDestinationService.useGetAll(destinationParms)
const updateDestination = ItineraryDestinationService.useUpdate()

const { params: resourceParms, setParams: setResourceParams } = useParams({
  filter: `itinerary="${itineraryID.value}"`,
  orderBy: '+startDate',
  expand: 'resourceProvider.resource, resourceProvider.provider, resourceProvider.user',
})
const resources = ItineraryResourceService.useGetAll(resourceParms)
const updateResource = ItineraryResourceService.useUpdate()

const resourceEvents = computed(() => resourcesToEvent(resources.data.value?.items || [])) 
const destinationEvents = computed(() => destinationsToEvent(destinations.data.value?.items || []))
const commentaryEvents = computed(() => commentariesToEvent(commentaries.data.value?.items || []))
const events = computed(() => [...commentaryEvents.value, ...destinationEvents.value, ...resourceEvents.value])

function handleEventType(type: string) {
  if (type === 'commentary') {
    commentaryDialog.value.form = true
  }
  if (type === 'destination') {
    destinationDialog.value.form = true
  }
  if (type === 'resource') {
    resourceDialog.value.form = true
  }
}

function handleClickDate(dateArg: DateClickArg) {
  currentEvent.value = null
  eventTypeDialog.value = true
  currentDate.value = dateArg
}

async function handleUpdateEvent(event: EventImpl) {
  currentEvent.value = event
  const data = {
    startDate: event.start?.toISOString(),
    endDate: event.end?.toISOString(),
    isAllDay: event.allDay,
  }
  try {
    if (event.extendedProps.type === 'commentary') {
      await updateCommentary.mutateAsync({
        id: event.id,
        data,
      })
      await commentaries.refetch()
    }
    if (event.extendedProps.type === 'destination') {
      await updateDestination.mutateAsync({
        id: event.id,
        data,
      })
      await destinations.refetch()
    }
    if (event.extendedProps.type === 'resource') {
      await updateResource.mutateAsync({
        id: event.id,
        data,
      })
      await resources.refetch()
    }
    toast.success('Event updated successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error updating event')
  }
}

function handleClickEvent(event: EventImpl) {
  currentEvent.value = event
  if (currentEvent.value.extendedProps.type === 'commentary') {
    commentaryDialog.value.form = true
  }
  if (currentEvent.value.extendedProps.type === 'destination') {
    destinationDialog.value.form = true
  }
  if (currentEvent.value.extendedProps.type === 'resource') {
    resourceDialog.value.details = true
  }
}

/* function handleRefresh() {
  commentaries.refetch()
  destinations.refetch()
  resources.refetch()
} */

function commentariesToEvent(commentaries: ItineraryCommentary[]) {
  return commentaries.map((commentary) => {
    return {
      id: commentary.id,
      title: commentary.commentary.en,
      start: commentary.startDate,
      end: commentary.endDate,
      allDay: commentary.isAllDay,
      extendedProps: {
        type: 'commentary',
      },
      ...colors.black,
    }
  })
}

function destinationsToEvent(destinations: ItineraryDestination[]) {
  return destinations.map((destination) => {
    return {
      id: destination.id,
      title: destination.expand.destination.name.en,
      start: destination.startDate,
      end: destination.endDate,
      allDay: destination.isAllDay,
      extendedProps: {
        type: 'destination',
      },
      ...colors.blue,
    }
  })
}

function resourcesToEvent(resources: ItineraryResource[]) {
  return resources.map((resource) => {
    return {
      id: resource.id,
      title: resource.expand.resourceProvider.expand.resource.name.en,
      start: resource.startDate,
      end: resource.endDate,
      allDay: resource.isAllDay,
      extendedProps: {
        type: 'resource',
      },
      ...colors.red,
    }
  })
}

watch(itineraryID, async () => {
  setResourceParams({ filter: `itinerary="${itineraryID.value}"` })
  setDestinationParams({ filter: `itinerary="${itineraryID.value}"` })
  setCommentaryParams({ filter: `itinerary="${itineraryID.value}"` })
})
</script>

<template>
  <EventTypeDialog v-model="eventTypeDialog" @click:event="handleEventType" />

  <DestinationDeleteDialog
    v-model="destinationDialog.delete"
    :itineraryDestinationID="currentEvent?.id"
  />

  <ResourceDetailsDialog
    v-if="
      currentEvent &&
      currentEvent.extendedProps.type === 'resource' &&
      currentEvent.id &&
      resourceDialog.details
    "
    v-model="resourceDialog.details"
    :itineraryResourceID="currentEvent?.id"
  />

  <DestinationFormDialog
    v-model="destinationDialog.form"
    :itineraryDestination="{
      tour: tourID,
      itinerary: itineraryID,
      startDate: formattedDate,
      endDate: formattedDate,
      isAllDay: currentDate?.allDay,
    }"
  />

  <ResourceFormDialog
    v-model="resourceDialog.form"
    :itineraryResource="{
      tour: tourID,
      itinerary: itineraryID,
      startDate: formattedDate,
      endDate: formattedDate,
      isAllDay: currentDate?.allDay,
    }"
  />

  <CommentaryFormDialog
    v-model="commentaryDialog.form"
    :commentary="{
      id: currentEvent?.extendedProps.type === 'commentary' ? currentEvent?.id : undefined,
      tour: tourID,
      itinerary: itineraryID,
      startDate: formattedDate,
      endDate: formattedDate,
      isAllDay: currentEvent?.allDay || currentDate?.allDay,
    }"
  />
  <FullCalendar
    :events="events"
    @click:date="handleClickDate"
    @click:event="handleClickEvent"
    @event:drop="handleUpdateEvent"
    @event:resize="handleUpdateEvent"
  />
</template>
