<script setup lang="ts">
import { computed, toRef, watch } from 'vue'
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import * as z from 'zod'
import { hourEnd, hourStart } from '@formkit/tempo'

import { Button } from '@/core/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/core/components/ui/dialog'
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/core/components/ui/form'
import { toast } from 'vue-sonner'
import { ImagesInput, LanguageTextArea } from '@/core/components/fields'
import { TourCommentaryService } from '@/tour/services/tourCommentaryService'
import { useQueryClient } from '@tanstack/vue-query'

const isOpen = defineModel<boolean>({ default: false })

const props = defineProps<{
  commentaryID?: string
  isAllDay?: boolean
  tourID: string
  tourVariantID: string
  endDate: Date
}>()

const startFetching = computed(() => !!props.commentaryID || isOpen.value)

const commentary = TourCommentaryService.useGetOne(
  toRef(props.commentaryID as string),
  startFetching,
)
const create = TourCommentaryService.useCreate()
const update = TourCommentaryService.useUpdate()
const query = useQueryClient()

const formSchema = toTypedSchema(
  z.object({
    commentary: z.record(z.string(), z.string().min(1, 'Every lang must has at least 1 character')),
    images: z.object({
      newFiles: z.array(
        z.custom<File>((file) => file instanceof File, {
          message: 'File is required',
        }),
      ),
      deletedFiles: z.array(z.string()),
      uploadedFiles: z.array(z.string()),
    }),
  }),
)

const { handleSubmit, values, setFieldValue, resetForm } = useForm({
  validationSchema: formSchema,
  initialValues: {
    commentary: {
      en: '',
    },
    images: {
      newFiles: [],
      deletedFiles: [],
      uploadedFiles: [],
    },
  },
})

const onSubmit = handleSubmit(async (values) => {
  const form = new FormData()
  form.append('commentary', JSON.stringify(values.commentary))
  values.images.newFiles.forEach((file) => {
    form.append('images+', file)
  })
  values.images.deletedFiles.forEach((file) => {
    form.append('images-', file)
  })
  form.append('tour', props.tourID)
  form.append('tourVariant', props.tourVariantID)
  if (props.isAllDay) {
    
  }
  form.append('startDate', props.endDate.toISOString())
  form.append('endDate', props.endDate.toISOString())
  try {
    if (props.commentaryID) {
      await update.mutateAsync({
        id: props.commentaryID,
        data: form,
      })
      toast.success('Commentary updated successfully')
      query.invalidateQueries({ queryKey: ['tourCommentaries'] })
    }
    if (!props.commentaryID) {
      await create.mutateAsync(form)
      toast.success('Commentary created successfully')
      query.invalidateQueries({ queryKey: ['tourCommentaries'] })
    }
  } catch (err) {
    console.log(err)
    toast.error('Error creating commentary')
  } finally {
    isOpen.value = false
  }
})

watch(isOpen, (newIsOpen) => {
  if (!newIsOpen) return
  resetForm()
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="max-w-6xl">
      <DialogHeader>
        <DialogTitle>{{ props.commentaryID ? 'Edit Commentary' : 'New Commentary' }}</DialogTitle>
        <DialogDescription>
          Make changes to your profile here. Click save when you're done.
          {{ props.endDate.toISOString() }} {{ values }}
        </DialogDescription>
      </DialogHeader>

      <form id="commentaryForm" class="grid md:grid-cols-2 gap-2" @submit.prevent="onSubmit">
        <FormField v-slot="{ componentField }" name="commentary">
          <FormItem class="self-start">
            <FormLabel>Commentary</FormLabel>
            <FormControl>
              <LanguageTextArea
                v-model="componentField.modelValue"
                @update:model-value="setFieldValue('commentary', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="images">
          <FormItem>
            <FormLabel>Images</FormLabel>
            <FormControl>
              <ImagesInput
                v-model="componentField.modelValue"
                @update:model-value="setFieldValue('images', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </form>

      <DialogFooter>
        <Button type="submit" form="commentaryForm"> Save </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
