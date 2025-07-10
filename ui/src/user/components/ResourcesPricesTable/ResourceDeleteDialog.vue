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
import { UserService } from '@/user/services/userService'
import { useQueryClient } from '@tanstack/vue-query'

import { Trash } from 'lucide-vue-next'
import { toast } from 'vue-sonner'

const props = defineProps<{
  serviceProviderId: string
}>()

const { mutateAsync: deleteService } = UserService.useDeleteResource()
const queryClient = useQueryClient()

const onSubmit = async () => {
  try {
    await deleteService({ userID: '123', resourceProviderID: props.serviceProviderId })
    toast.success('Service deleted successfully')
    queryClient.invalidateQueries({ queryKey: ['user_resources'] })
    queryClient.invalidateQueries({ queryKey: ['provider_resources'] })
  } catch (error) {
    toast.error('Error deleting service')
    console.error('Error deleting service:', error)
  }
}
</script>

<template>
  <AlertDialog>
    <AlertDialogTrigger><Trash class="w-5 h-5" /></AlertDialogTrigger>
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>Are you absolutely sure? {{ props.serviceProviderId }}</AlertDialogTitle>
        <AlertDialogDescription>
          This action cannot be undone. This will permanently delete the row selected.
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Cancel</AlertDialogCancel>
        <AlertDialogAction @click="onSubmit">Continue</AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>
