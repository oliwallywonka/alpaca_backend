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
import { TourResourcesService } from '@/tour/services/tourResources'
import { useQueryClient } from '@tanstack/vue-query'
import { toast } from 'vue-sonner'

const isOpen = defineModel<boolean>({ default: false })
const props = defineProps<{
  tourResource?: string
}>()
const { mutateAsync } = TourResourcesService.useDelete()
const query = useQueryClient()

async function handleDelete() {
  if (!props.tourResource) return
  try {
    await mutateAsync(props.tourResource)
    query.invalidateQueries({ queryKey: ['tourResources'] })
    toast.success('Resource deleted successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error deleting Resource')
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
          This action cannot be undone. This will permanently delete this Resource.
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Cancel</AlertDialogCancel>
        <AlertDialogAction @click="handleDelete"> Continue </AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>
