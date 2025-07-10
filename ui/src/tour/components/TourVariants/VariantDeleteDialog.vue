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
import { TourVariantService } from '@/tour/services/tourVariantService'
import { useQueryClient } from '@tanstack/vue-query';
import { toast } from 'vue-sonner';

const props = defineProps<{
  variantID: string
}>()

const update = TourVariantService.useUpdate()
const query = useQueryClient()

const onSubmit = async () => {
    try {
      await update.mutateAsync({ id: props.variantID, data: { isActive: false } })
      toast.success('Variant deleted successfully')
      query.invalidateQueries({ queryKey: ['tourVariants'] })
    } catch (err) {
      console.log(err)
      toast.error('Error deleting variant')
    }   
}
</script>

<template>
  <AlertDialog>
    <AlertDialogTrigger as-child>
      <slot>
        <Button variant="outline"> Open </Button>
      </slot>
    </AlertDialogTrigger>
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
        <AlertDialogDescription>
          This action cannot be undone. This will delete this tour variant and remove all itineraries and operative expenses items ASSOCIATED .
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Cancel</AlertDialogCancel>
        <AlertDialogAction @click="onSubmit">Continue</AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>