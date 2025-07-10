<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRoute } from 'vue-router'

import type { TourCommentary } from '@/tour/interfaces/tourCommentary'

import { CardHeader, CardTitle, Card, CardContent } from '@/core/components/ui/card'
import { useParams } from '@/core/hooks/useParams'
import { TourCommentaryService } from '@/tour/services/tourCommentaryService'
import EventDetailDialog from './EventDetailDialog.vue'
import EventNewDialog from './EventNewDialog.vue'
import CommentaryFormDialog from './CommentaryFormDialog.vue'
import ItineraryCalendar from './ItineraryCalendar.vue'
import { colors } from '@/core/constants/colors'
import type { EventImpl } from '@fullcalendar/core/internal'
import { toast } from 'vue-sonner'

const props = defineProps<{
  tourVariantID: string
}>()

const route = useRoute()
const tourID = route.params.tourID as string
const variantID = computed(() => props.tourVariantID)

const eventDialogIsOpen = ref(false)
const eventNewDialogIsOpen = ref(false)
const commentaryDialogIsOpen = ref(false)

const currentDate = ref<Date>(new Date())
const currentEvent = ref<EventImpl | null>(null)

const commetariesParams = useParams({
  filter: `tourVariant="${variantID.value}"`,
  orderBy: '+startDate',
})
const commentaries = TourCommentaryService.useGetAll(commetariesParams.params)
const update = TourCommentaryService.useUpdate()

function handleEventType(type: string) {
  if (type === 'commentary') {
    commentaryDialogIsOpen.value = true
  }
}

function handleClickDate(date: Date) {
  eventNewDialogIsOpen.value = true
  currentDate.value = date
}

async function handleUpdateEvent(event: EventImpl) {
  currentEvent.value = event
  try {
    await update.mutateAsync({
      id: event.id,
      data: {
        startDate: event.start?.toISOString(),
        endDate: event.end?.toISOString(),
      },
    })
    await commentaries.refetch()
  } catch (err) {
    console.log(err)
    toast.error('Error updating event')
  }
}

function commentariesToEvents(commentaries: TourCommentary[]) {
  return commentaries.map((commentary) => {
    return {
      id: commentary.id,
      title: commentary.commentary.en,
      start: commentary.startDate,
      end: commentary.endDate,
      extendedProps: {
        type: 'commentary',
      },
      ...colors.blue,
    }
  })
}
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle>Itinerary Builder </CardTitle>
      <Button @click="commentaries.refetch"> reload </Button>
    </CardHeader>
    <CardContent class="grid min-h-[85vh] gap-2 p-4">
      <EventNewDialog v-model="eventNewDialogIsOpen" @click:event="handleEventType" />
      <EventDetailDialog v-model="eventDialogIsOpen" :event="null" />
      <CommentaryFormDialog
        v-model="commentaryDialogIsOpen"
        :endDate="currentDate"
        :tourID="tourID"
        :tourVariantID="props.tourVariantID"
      />

      <ItineraryCalendar
        :commentaries="commentariesToEvents(commentaries.data.value?.items || [])"
        :initialDate="new Date(commentaries.data.value?.items?.[0]?.startDate || Date.now())"
        @click:date="handleClickDate"
        @event:drop="handleUpdateEvent"
        @event:resize="handleUpdateEvent"
      />
    </CardContent>
  </Card>
</template>
