<script setup lang="ts">
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
import { TourDestinationService } from '@/tour/services/tourDestinationService'
import { useQueryClient } from '@tanstack/vue-query';
import { toast } from 'vue-sonner'

const isOpen = defineModel<boolean>({ default: false })
const props = defineProps<{
  tourDestinationID?: string
}>()
const { mutateAsync } = TourDestinationService.useDelete()
const query = useQueryClient()

async function handleDelete() {
  if (!props.tourDestinationID) return
  try {
    await mutateAsync(props.tourDestinationID)
    query.invalidateQueries({ queryKey: ['tourDestinations'] })
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
