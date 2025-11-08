<script setup lang="ts">
import { computed, watch } from 'vue'
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
import { ItineraryCommentaryService } from '@/itinerary/services/itineraryCommentaryService'
import type { ItineraryCommentary } from '@/itinerary/interfaces/itineraryCommentary'
import { getFileNameFromURL, getFileURL } from '@/core/lib/fileURL'
import ConfirmDeleteDialog from '@/core/components/Dialogs/ConfirmDeleteDialog.vue'
import { Trash } from 'lucide-vue-next'

const isOpen = defineModel<boolean>({ default: false })

const props = defineProps<{
  commentary: Partial<ItineraryCommentary>
}>()

const startFetching = computed(() => Boolean(props.commentary.id) && isOpen.value)
const commentaryID = computed(() => props.commentary.id || '')

const { data } = ItineraryCommentaryService.useGetOne(commentaryID, startFetching)
const create = ItineraryCommentaryService.useCreate()
const update = ItineraryCommentaryService.useUpdate()
const remove = ItineraryCommentaryService.useDelete()
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

const derivedData = computed(() =>
  data.value
    ? {
        ...data.value,
        images: data.value.images.map((image: string) =>
          getFileURL({
            collectionName: 'itineraryCommentaries',
            recordID: commentaryID.value,
            fileName: image,
          }),
        ),
      }
    : null,
)

const initialValues = {
  commentary: data.value?.commentary || { en: '' },
  images: {
    newFiles: [],
    deletedFiles: [],
    uploadedFiles: derivedData.value?.images || [],
  },
}

const { handleSubmit, setFieldValue, setValues, resetForm } = useForm({
  validationSchema: formSchema,
  initialValues,
})

const onSubmit = handleSubmit(async (values) => {
  const form = new FormData()
  form.append('commentary', JSON.stringify(values.commentary))
  values.images.newFiles.forEach((file) => {
    form.append('images+', file)
  })
  values.images.deletedFiles.forEach((file) => {
    form.append('images-', getFileNameFromURL(file))
  })
  if (props.commentary.tour) {
    form.append('tour', props.commentary.tour)
  }
  if (props.commentary.itinerary) {
    form.append('itinerary', props.commentary.itinerary)
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
      query.invalidateQueries({ queryKey: ['itineraryCommentaries'] })
    }
    if (!props.commentary.id) {
      await create.mutateAsync(form)
      toast.success('Commentary created successfully')
      query.invalidateQueries({ queryKey: ['itineraryCommentaries'] })
    }
  } catch (err) {
    console.log(err)
    toast.error('Error creating commentary')
  } finally {
    isOpen.value = false
  }
})

async function handleDelete() {
  if (!commentaryID.value) return
  try {
    await remove.mutateAsync(commentaryID.value)
    query.invalidateQueries({ queryKey: ['itineraryCommentaries'] })
    toast.success('Commentary deleted successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error deleting commentary')
  } finally {
    isOpen.value = false
  }
}

watch(isOpen, (newIsOpen) => {
  if (!newIsOpen) return
  resetForm()
})

watch([isOpen, data], ([newIsOpen, newData]) => {
  if (!newIsOpen || !newData) return
  setValues({
    commentary: newData.commentary,
    images: {
      newFiles: [],
      deletedFiles: [],
      uploadedFiles: derivedData.value?.images || [],
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
          Make changes to your profile here. Click save when you're done. {{ props.commentary
          }}{{ startFetching }}
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
        <ConfirmDeleteDialog @confirm:delete="handleDelete">
          <Button variant="destructive"><Trash /> Delete</Button>
        </ConfirmDeleteDialog>
        <Button type="submit" form="commentaryForm"> Save </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
