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
  AlertDialogTrigger,
} from '@/core/components/ui/alert-dialog'
import { Button } from '@/core/components/ui/button'

const isOpen = defineModel<boolean>({ default: false })
const props = defineProps<{
  showTrigger?: boolean
}>()
const emit = defineEmits<{
  (e: 'confirm:delete'): void
}>()
function handleConfirm() {
  emit('confirm:delete')
}
</script>

<template>
  <AlertDialog v-model="isOpen">
    <AlertDialogTrigger v-if="props.showTrigger" as-child>
      <slot>
        <Button variant="outline"> Show Dialog </Button>
      </slot>
    </AlertDialogTrigger>
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>
          <slot name="title">Are you absolutely sure?</slot>
        </AlertDialogTitle>
        <AlertDialogDescription>
          <slot name="description"> This action cannot be undone. </slot>
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Cancel</AlertDialogCancel>
        <AlertDialogAction @click="handleConfirm"> Continue </AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>
