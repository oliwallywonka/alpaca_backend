<script setup lang="ts">
import { toast } from 'vue-sonner'
import { useQueryClient } from '@tanstack/vue-query'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/core/components/ui/alert-dialog'
import { ItineraryDestinationService } from '@/itinerary/services/ItineraryDestinationService'

const isOpen = defineModel<boolean>({ default: false })
const props = defineProps<{
  tourDestinationID?: string
}>()
const { mutateAsync } = ItineraryDestinationService.useDelete()
const query = useQueryClient()

async function handleDelete() {
  if (!props.tourDestinationID) return
  try {
    await mutateAsync(props.tourDestinationID)
    query.invalidateQueries({ queryKey: ['itineraryDestinations'] })
    toast.success('Destination deleted successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error deleting Destination')
  } finally {
    isOpen.value = false
  }
}
</script>

<template>
  <AlertDialog v-model:open="isOpen">
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
        <AlertDialogDescription>
          This action cannot be undone. This will permanently delete this Destination.
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Cancel</AlertDialogCancel>
        <AlertDialogAction @click="handleDelete"> Continue </AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>
