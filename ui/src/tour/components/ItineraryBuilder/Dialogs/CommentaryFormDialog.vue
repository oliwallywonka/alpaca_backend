<script setup lang="ts">
import { computed, toRef, watch } from 'vue'
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import * as z from 'zod'
import { toast } from 'vue-sonner'
import { useQueryClient } from '@tanstack/vue-query'

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
import { ImagesInput, LanguageTextArea } from '@/core/components/fields'
import { TourCommentaryService } from '@/tour/services/tourCommentaryService'
import type { TourCommentary } from '@/tour/interfaces/tourCommentary'

const isOpen = defineModel<boolean>({ default: false })

const props = defineProps<{
  commentary: Partial<TourCommentary>
}>()

const startFetching = computed(() => !!props.commentary.id || isOpen.value)

const { data } = TourCommentaryService.useGetOne(
  toRef(props.commentary.id!),
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

const { handleSubmit, setFieldValue, setValues, resetForm } = useForm({
  validationSchema: formSchema,
  initialValues: {
    commentary: {
      en: '',
    },
    images: {
      newFiles: [],
      deletedFiles: [],
      uploadedFiles: data.value?.images || [],
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
  if (props.commentary.tour) {
    form.append('tour', props.commentary.tour)
  }
  if (props.commentary.tourVariant) {
    form.append('tourVariant', props.commentary.tourVariant)
  }
  form.append('isAllDay', props.commentary.isAllDay ? 'true' : 'false')
  form.append('startDate', props.commentary.startDate || '')
  form.append('endDate', props.commentary.endDate || '')
  try {
    if (props.commentary.id) {
      await update.mutateAsync({
        id: props.commentary.id,
        data: form,
      })
      toast.success('Commentary updated successfully')
      query.invalidateQueries({ queryKey: ['tourCommentaries'] })
    }
    if (!props.commentary.id) {
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

// TODO: fix reactivity
watch([isOpen, data], ([newIsOpen, newData]) => {
  if (!newIsOpen || !newData) return
  setValues({
    commentary: newData.commentary,
    images: {
      newFiles: [],
      deletedFiles: [],
      uploadedFiles: newData.images,
    },
  })
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="max-w-6xl">
      <DialogHeader>
        <DialogTitle>{{ props.commentary.id ? 'Edit Commentary' : 'New Commentary' }}</DialogTitle>
        <DialogDescription>
          Make changes to your profile here. Click save when you're done.
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
