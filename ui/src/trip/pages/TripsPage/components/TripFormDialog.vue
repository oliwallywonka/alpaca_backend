<script setup lang="ts">
import { ref, watch } from 'vue'
import { useForm } from 'vee-validate'
import { z } from 'zod'
import { toTypedSchema } from '@vee-validate/zod'
import { PlusIcon } from 'lucide-vue-next'
import { toast } from 'vue-sonner'
import { addDay, diffDays } from '@formkit/tempo'
import { useQueryClient } from '@tanstack/vue-query'
import type { DateClickArg } from '@fullcalendar/interaction/index.js'

import { Button } from '@/core/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/core/components/ui/dialog'
import { Badge } from '@/core/components/ui/badge'
import { Switch } from '@/core/components/ui/switch'
import {
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
  FormControl
} from '@/core/components/ui/form'
import FullCalendar from '@/core/components/calendar/FullCalendar.vue'
import TourVariantPickerDialog from '@/trip/components/TourVariantPickerDialog/TourItineraryPickerDialog.vue'

import type { ItinerarySummary } from '@/itinerary/interfaces/itinerary'
import { TripService } from '@/trip/services/TripService'
import { colors } from '@/core/constants/colors'
import { dateToSuggar } from '@/core/lib/date'

const isOpen = defineModel<boolean>({ default: false })

const isTourVariantOpen = ref<boolean>(false)
const currentItinerary = ref<ItinerarySummary>()
const startDate = ref<Date>()

const query = useQueryClient()
const create = TripService.useCreate()

const schema = toTypedSchema(
  z.object({
    transferResources: z.boolean(),
    itineraryID: z.string().min(1, 'Tour variant is required'),
  }),
)

const initialValues = {
  transferResources: true,
  itineraryID: '',
}
const { handleSubmit, setFieldValue, setValues } = useForm({
  validationSchema: schema,
  initialValues,
})

const handleSelectTourVariant = (itinerary: ItinerarySummary | undefined) => {
  if (!startDate.value || !itinerary) return
  const days = diffDays(startDate.value, itinerary?.startDate || new Date(Date.now()))
  const newStartDate = addDay(itinerary?.startDate || new Date(Date.now()), days)
  const newEndDate = addDay(itinerary?.endDate || new Date(Date.now()), days)

  currentItinerary.value = {
    ...itinerary,
    startDate: newStartDate,
    endDate: newEndDate,
  }
}

const handleClickDate = (date: DateClickArg) => {
  isTourVariantOpen.value = true
  startDate.value = date.date
}

function itinerarysToEvent(itineraries: ItinerarySummary[]) {
  return itineraries?.map((itinerary) => {
    return {
      id: itinerary.id,
      title: itinerary.expand.tour.name?.en,
      start: itinerary.startDate,
      end: itinerary.endDate || itinerary.startDate,
      allDay: true,
      ...colors.blue,
    }
  })
}



const onSubmit = handleSubmit(async (values) => {
  try {
    await create.mutateAsync({
      ...values,
      startDate: startDate.value as Date,
    })
    await query.invalidateQueries({ queryKey: ['trips'] })
    toast.message('Trip created successfully')
  } catch (error) {
    console.log(error)
    toast.message('Error creating trip')
  } finally {
    isOpen.value = false
  }
})

watch(isOpen, (newIsOpen) => {
  if (!newIsOpen) return
  setValues(initialValues)
  currentItinerary.value = undefined
})

watch(currentItinerary, (newTourVariant) => {
  if (!newTourVariant) return
  setFieldValue('itineraryID', newTourVariant.id)
})
</script>

<template>
  <Dialog id="new-trip-dialog" v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="outline"> <PlusIcon /> New Trip </Button>
    </DialogTrigger>
    <DialogContent class="sm:max-w-[920px]">
      <DialogHeader>
        <DialogTitle>New Trip </DialogTitle>
        <DialogDescription> Select a tour itinerary to be imported to the trip </DialogDescription>
      </DialogHeader>
      <div class="grid grid-cols-[1fr_2fr] gap-2">
        <form id="new-trip-form" @submit="onSubmit" class="flex flex-col gap-2">
          <FormField name="itineraryID">
            <FormItem>
              <FormLabel>* Selected Tour Variant</FormLabel>
              <span v-if="!currentItinerary" class="text-muted-foreground text-sm">
                Click into a date to choose a tour Itinerary.
              </span>

              <div v-if="currentItinerary" class="flex flex-wrap gap-1">
                <Badge>{{ currentItinerary.expand.tour.name?.en }}</Badge>
                <Badge>{{
                  `min: ${currentItinerary.minPersons}, max: ${currentItinerary.maxPersons}`
                }}</Badge>
                <Badge>{{ `Start: ${dateToSuggar(currentItinerary.startDate)}` }}</Badge>
                <Badge>{{ `End: ${dateToSuggar(currentItinerary.endDate)}` }}</Badge>
                <Badge>{{ `Price: ${currentItinerary.finalPrice} usd` }}</Badge>
              </div>
              <FormMessage />
            </FormItem>
          </FormField>
          <FormField v-slot="{ value, handleChange }" name="transferResources">
            <FormItem>
              <FormLabel>Transfer all resources</FormLabel>
              <FormDescription>Transfer all resources from tour itinerary to trip.</FormDescription>
              <FormControl>
                <Switch :modelValue="value" @update:modelValue="handleChange" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
        </form>
        <FullCalendar
          :events="itinerarysToEvent(currentItinerary ? [currentItinerary] : [])"
          :reRender="isOpen"
          :isEditable="false"
          @click:date="handleClickDate"
          @event:drop="
            (event) => {
              if (!currentItinerary) return
              startDate = event.start as Date
              currentItinerary = {
                ...currentItinerary,
                startDate: event.start as Date,
                endDate: event.end || (event.start as Date),
              }
            }
          "
        />
      </div>
      <DialogFooter>
        <Button type="submit" form="new-trip-form"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>

  <TourVariantPickerDialog
    v-model="isTourVariantOpen"
    @select:itinerary="handleSelectTourVariant"
  />
</template>
