<script setup lang="ts">
import { ref } from 'vue'

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogClose,
} from '@/core/components/ui/dialog'
import { Badge } from '@/core/components/ui/badge'
import type { Itinerary } from '@/itinerary/interfaces/itinerary'
import TourItineraryPicker from '@/trip/components/TourItineraryPicker/TourItineraryPicker.vue'
import { Button } from '@/core/components/ui/button'

const isOpen = defineModel<boolean>({ default: false })
const emit = defineEmits<{
  (e: 'select:itinerary', itinerary: Itinerary | undefined): void
}>()
const currentTourItinerary = ref<Itinerary>()
const handleSelectTourItinerary = (itinerary: Itinerary | undefined) => {
  currentTourItinerary.value = itinerary
}
const handleClose = () => {
  emit('select:itinerary', currentTourItinerary.value)
  currentTourItinerary.value = undefined
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="sm:max-w-[840px]">
      <DialogHeader>
        <DialogTitle>Pick a tour variant</DialogTitle>
        <DialogDescription> Select a tour variant to be imported to the trip </DialogDescription>
      </DialogHeader>
      <p>
        <span class="font-semibold text-sm">* Selected Variant: </span>
        <span v-if="!currentTourItinerary" class="text-muted-foreground text-sm">
          No variant selected
        </span>
      </p>
      <div v-if="currentTourItinerary" class="flex flex-wrap gap-1">
        <Badge>
          {{ currentTourItinerary.expand.tour.name?.en }}
        </Badge>
        <Badge>
          {{ `min: ${currentTourItinerary.minPersons}, max: ${currentTourItinerary.maxPersons}` }}
        </Badge>
        <Badge>
          Price: {{ currentTourItinerary.finalPrice }} usd
        </Badge>
      </div>
      <TourItineraryPicker @select:itinerary="handleSelectTourItinerary" />
      <DialogFooter>
        <DialogClose as-child>
          <Button @click="handleClose"> Save changes </Button>
        </DialogClose>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
