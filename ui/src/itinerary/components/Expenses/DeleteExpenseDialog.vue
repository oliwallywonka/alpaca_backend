<script setup lang="ts">
import { useQueryClient } from '@tanstack/vue-query'
import { Trash } from 'lucide-vue-next'
import { toast } from 'vue-sonner'

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/core/components/ui/alert-dialog'
import { Button } from '@/core/components/ui/button'
import { ItineraryResourceService } from '@/itinerary/services/itineraryResources'

const isOpen = defineModel({ default: false })
const props = defineProps<{
  itineraryResourceID: string
}>()

const { mutateAsync } = ItineraryResourceService.useDelete()
const query = useQueryClient()

const handleDelete = async () => {
  try {
    await mutateAsync(props.itineraryResourceID)
    await query.invalidateQueries({ queryKey: ['itineraryResources'] })
    toast.success('Resource deleted successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error deleting resource')
  }
}
</script>

<template>
  <AlertDialog v-model:open="isOpen">
    <AlertDialogTrigger as-child>
      <Button variant="outline"> <Trash class="text-red-600" /> </Button>
    </AlertDialogTrigger>
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
        <AlertDialogDescription>
          This action cannot be undone. This will permanently remove the resource from the table.
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Cancel</AlertDialogCancel>
        <AlertDialogAction @click="handleDelete">Continue</AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>
