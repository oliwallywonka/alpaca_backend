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
import { TourCommentaryService } from '@/tour/services/tourCommentaryService'
import { useQueryClient } from '@tanstack/vue-query';
import { toast } from 'vue-sonner'

const isOpen = defineModel<boolean>({ default: false })
const props = defineProps<{
  commentaryID?: string
}>()
const { mutateAsync } = TourCommentaryService.useDelete()
const query = useQueryClient()

async function handleDelete() {
  if (!props.commentaryID) return
  try {
    await mutateAsync(props.commentaryID)
    query.invalidateQueries({ queryKey: ['tourCommentaries'] })
    toast.success('Commentary deleted successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error deleting commentary')
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
          This action cannot be undone. This will permanently delete this commentary and the files
          related.
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Cancel</AlertDialogCancel>
        <AlertDialogAction @click="handleDelete"> Continue </AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>
