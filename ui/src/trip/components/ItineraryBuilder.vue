<script setup lang="ts">
import { toRef } from 'vue'
import { useRoute } from 'vue-router'

import ItineraryCalendar from '@/itinerary/components/Calendar/ItineraryCalendar.vue'
import { TripService } from '../services/TripService'
import { CardDescription } from '@/core/components/ui/card'
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from '@/core/components/ui/sheet'
import ItineraryDayList from '@/trip/components/ItineraryDayList/ItineraryDayList.vue'

const route = useRoute()
const tripID = route.params.tripID as string

const { data, isLoading } = TripService.useGetOne(toRef(tripID), toRef(true))
</script>

<template>
  <p v-if="isLoading" class="text-gray-500">Loading...</p>
  <div v-if="data && !isLoading">
    <div class="grid gap-2">
      <CardDescription>Click on a date to add a service</CardDescription>

      <Sheet :open="true">
        <SheetTrigger>Open</SheetTrigger>
        <SheetContent class="sm:max-w-full w-[90%] overflow-y-auto">
          <SheetHeader>
            <SheetTitle class="text-4xl font-bold">Itineray</SheetTitle>
            <SheetDescription>
              This is a preview of the itinerary. You can edit the itinerary by clicking on each
              item.
            </SheetDescription>
          </SheetHeader>

          <ItineraryDayList />
        </SheetContent>
      </Sheet>
      <ItineraryCalendar :itineraryID="data.itinerary" :tourID="data.tour" />
    </div>
  </div>
</template>
