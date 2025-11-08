<script setup lang="ts">
// TODO: Use forms and create Field Relation Component for Relational fields
import { ref } from 'vue'
import { toast } from 'vue-sonner'
import { useQueryClient } from '@tanstack/vue-query'

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/core/components/ui/dialog'
import { Button } from '@/core/components/ui/button'
import { Badge } from '@/core/components/ui/badge'

import type { ItineraryDestination } from '@/itinerary/interfaces/itineraryDestination'
import type { Destination } from '@/destination/interfaces/destination'
import DestinationPage from '@/destination/DestinationPage.vue'
import { ItineraryDestinationService } from '@/itinerary/services/ItineraryDestinationService'

const isOpen = defineModel<boolean>({ default: false })
const props = defineProps<{
  itineraryDestination: Partial<ItineraryDestination>
}>()
const currrentDestination = ref<Destination | undefined>()

const create = ItineraryDestinationService.useCreate()
const query = useQueryClient()

const handleSelect = (destination: Destination | undefined) => {
  currrentDestination.value = destination
}

const handleSubmit = async () => {
  if (!currrentDestination.value) {
    toast.error('Please select a destination')
    return
  }
  try {
    const data = {
      ...props.itineraryDestination,
      destination: currrentDestination.value?.id,
      // TODO: for some reason the description is not being saved
      descritption: {
        en: '123',
      },
    }
    console.log(data)
    await create.mutateAsync(data)
    await query.invalidateQueries({ queryKey: ['itineraryDestinations'] })
  } catch (err) {
    console.log(err)
    toast.error('Error creating destination')
  } finally {
    isOpen.value = false
  }
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="max-w-6xl">
      <DialogHeader>
        <DialogTitle>Destination</DialogTitle>
        <DialogDescription>
          Select the destination of the itinerary "this will be showed on the summary section"
        </DialogDescription>
      </DialogHeader>

      <p>
        <span class="font-semibold text-sm">* Selected destination: </span>
        <span v-if="!currrentDestination" class="text-muted-foreground text-sm">
          No destination selected
        </span>
        <Badge v-else>{{ currrentDestination.name.en }}</Badge>
      </p>

      <DestinationPage :showSelect="true" @select:destination="handleSelect" />

      <DialogFooter>
        <Button type="submit" form="destinationForm" @click="handleSubmit"> Save </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
