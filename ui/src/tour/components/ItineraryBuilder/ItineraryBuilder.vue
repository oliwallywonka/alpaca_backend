<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRoute } from 'vue-router'
import { toast } from 'vue-sonner'
import type { DateClickArg } from '@fullcalendar/interaction/index.js'

import type { TourCommentary } from '@/tour/interfaces/tourCommentary'
import { CardHeader, CardTitle, Card, CardContent } from '@/core/components/ui/card'
import { Button } from '@/core/components/ui/button'
import { useParams } from '@/core/hooks/useParams'
import { TourCommentaryService } from '@/tour/services/tourCommentaryService'
import EventTypeDialog from './Dialogs/EventTypeDialog.vue'
import CommentaryFormDialog from './Dialogs/CommentaryFormDialog.vue'
import ItineraryCalendar from './ItineraryCalendar.vue'
import { colors } from '@/core/constants/colors'
import type { EventImpl } from '@fullcalendar/core/internal'
import DestinationFormDialog from './Dialogs/DestinationFormDialog.vue'
import type { TourDestination } from '@/tour/interfaces/tourDestination'
import { TourDestinationService } from '@/tour/services/tourDestinationService'
import ResourceFormDialog from './Dialogs/ResourceFormDialog.vue'
import { TourResourcesService } from '@/tour/services/tourResources'
import type { TourResource } from '@/tour/interfaces/tourResources'
import ResourceDeleteDialog from './Dialogs/ResourceDeleteDialog.vue'
import CommentaryDeleteDialog from './Dialogs/CommentaryDeleteDialog.vue'
import DestinationDeleteDialog from './Dialogs/DestinationDeleteDialog.vue'

const props = defineProps<{
  tourVariantID: string
}>()

const route = useRoute()
const tourID = route.params.tourID as string
const variantID = computed(() => props.tourVariantID)

const eventTypeDialog = ref(false)

const commentaryDialog = ref(false)
const commentaryDeleteDialog = ref(false)

const destinationDialog = ref(false)
const destinationDeleteDialog = ref(false)

const resourceDialog = ref(false)
const resourceDeleteDialog = ref(false)

const currentEvent = ref<(EventImpl & DateClickArg) | EventImpl | DateClickArg | null>(null)

const { params: commentaryParams } = useParams({
  filter: `tourVariant="${variantID.value}"`,
  orderBy: '+startDate',
})
const commentaries = TourCommentaryService.useGetAll(commentaryParams)
const updateCommentary = TourCommentaryService.useUpdate()

const { params: destinationParms } = useParams({
  filter: `tourVariant="${variantID.value}"`,
  orderBy: '+startDate',
  expand: 'destination',
})
const destinations = TourDestinationService.useGetAll(destinationParms)
const updateDestination = TourDestinationService.useUpdate()

const { params: resourceParms } = useParams({
  filter: `tourVariant="${variantID.value}"`,
  orderBy: '+startDate',
  expand: 'resourceProvider.resource, resourceProvider.provider, resourceProvider.user',
})
const resources = TourResourcesService.useGetAll(resourceParms)
const updateResource = TourResourcesService.useUpdate()

function handleEventType(type: string) {
  if (type === 'commentary') {
    commentaryDialog.value = true
  }
  if (type === 'destination') {
    destinationDialog.value = true
  }
  if (type === 'resource') {
    resourceDialog.value = true
  }
}

function handleClickDate(event: DateClickArg) {
  eventTypeDialog.value = true
  currentEvent.value = event
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
  } catch (err) {
    console.log(err)
    toast.error('Error updating event')
  }
}

function handleClickEvent(event: EventImpl) {
  currentEvent.value = event
  if (event.extendedProps.type === 'commentary') {
    commentaryDeleteDialog.value = true
  }
  if (event.extendedProps.type === 'destination') {
    destinationDeleteDialog.value = true
  }
  if (event.extendedProps.type === 'resource') {
    resourceDeleteDialog.value = true
  }
}

function handleRefresh() {
  commentaries.refetch()
  destinations.refetch()
  resources.refetch()
}

function commentariesToEvent(commentaries: TourCommentary[]) {
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

function destinationsToEvent(destinations: TourDestination[]) {
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

function resourcesToEvent(resources: TourResource[]) {
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
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle>Itinerary Builder </CardTitle>
      <Button @click="handleRefresh"> reload </Button>
    </CardHeader>
    <CardContent class="grid min-h-[85vh] gap-2 p-4">
      <EventTypeDialog v-model="eventTypeDialog" @click:event="handleEventType" />

      <ResourceDeleteDialog
        v-model="resourceDeleteDialog"
        :tourResource="currentEvent?.id"
      />
      <CommentaryDeleteDialog
        v-model="commentaryDeleteDialog"
        :commentaryID="currentEvent?.id"
      />
      <DestinationDeleteDialog
        v-model="destinationDeleteDialog"
        :tourDestinationID="currentEvent?.id"
      />

      <DestinationFormDialog
        v-model="destinationDialog"
        :tourDestination="{
          tour: tourID,
          tourVariant: props.tourVariantID,
          startDate:
            currentEvent && (currentEvent?.start?.toISOString() || currentEvent.date.toISOString()),
          endDate: (currentEvent && currentEvent?.end?.toISOString()) || '',
          isAllDay: currentEvent?.allDay || false,
        }"
      />

      <ResourceFormDialog
        v-model="resourceDialog"
        :tourResource="{
          tour: tourID,
          tourVariant: props.tourVariantID,
          startDate:
            currentEvent && (currentEvent?.start?.toISOString() || currentEvent.date.toISOString()),
          endDate: currentEvent && currentEvent?.end?.toISOString(),
          isAllDay: currentEvent?.allDay,
        }"
      />

      <CommentaryFormDialog
        v-model="commentaryDialog"
        :commentary="{
          id: currentEvent && currentEvent?.id,
          tourVariant: props.tourVariantID,
          tour: tourID,
          startDate:
            currentEvent && (currentEvent?.start?.toISOString() || currentEvent.date.toISOString()),
          endDate: currentEvent && currentEvent?.end?.toISOString(),
          isAllDay: currentEvent?.allDay,
        }"
      />

      <ItineraryCalendar
        :commentaries="commentariesToEvent(commentaries.data.value?.items || [])"
        :destinations="destinationsToEvent(destinations.data.value?.items || [])"
        :resources="resourcesToEvent(resources.data.value?.items || [])"
        :initialDate="new Date(commentaries.data.value?.items?.[0]?.startDate || Date.now())"
        @click:date="handleClickDate"
        @click:event="handleClickEvent"
        @event:drop="handleUpdateEvent"
        @event:resize="handleUpdateEvent"
      />
    </CardContent>
  </Card>
</template>
