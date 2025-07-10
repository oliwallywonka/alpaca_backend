<script setup lang="ts">
import { computed, toRef, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import { z } from 'zod'
import { toast } from 'vue-sonner'

import { Card, CardContent, CardHeader, CardTitle, CardFooter } from '@/core/components/ui/card'
import { ImageInput, ImagesInput } from '@/core/components/fields'
import { FormControl, FormDescription, FormField, FormItem, FormLabel } from '@/core/components/ui/form'
import { Button } from '@/core/components/ui/button'
import FormMessage from '@/core/components/ui/form/FormMessage.vue'
import { TourService } from '@/tour/services/tourService'
import { getFileURL, getFileNameFromURL } from '@/core/lib/fileURL'

const route = useRoute()
const tourID = route.params.tourID as string
const { data, isLoading, refetch } = TourService.useGetOne(
  toRef(route.params.tourID as string),
  toRef(true),
)
const { mutateAsync, isPending } = TourService.useUpdate()

const derivedData = computed(() =>
  data.value
    ? {
        ...data.value,
        banner: data.value.banner && getFileURL({
          collectionName: 'tours',
          recordID: tourID,
          fileName: data.value.banner,
        }),
        images: data.value.images.map((image) =>
          getFileURL({
            collectionName: 'tours',
            recordID: tourID,
            fileName: image,
          }),
        ),
      }
    : null,
)

const formSchema = toTypedSchema(
  z.object({
    banner: z
      .object({
        newFile: z.custom<File | null>((file) => file instanceof File || file === null),
        uploadedFile: z.string(),
      })
      /* .refine((banner) => !banner.newFile && !banner.uploadedFile, {
        message: 'Banner its mandatory',
      }) */,
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
const { handleSubmit, setFieldValue, setValues } = useForm({
  validationSchema: formSchema,
  initialValues: {
    banner: {
      newFile: null,
      uploadedFile: derivedData.value?.banner ?? '',
    },
    images: {
      newFiles: [],
      deletedFiles: [],
      uploadedFiles: derivedData.value?.images ?? [],
    },
  },
})

const onSubmit = handleSubmit(async (values) => {
  if (!derivedData.value) return
  const form = new FormData()
  if (values.banner.newFile) {
    form.append('banner+', values.banner.newFile)
  }
  if (values.banner.uploadedFile && values.banner.newFile) {
    form.append('banner-', getFileNameFromURL(values.banner.uploadedFile))
  }
  values.images.newFiles.forEach((file) => {
    form.append('images+', file)
  })
  values.images.deletedFiles.forEach((file) => {
    form.append('images-', getFileNameFromURL(file))
  })

  try {
    await mutateAsync({
      id: tourID,
      data: form,
    })
    toast.success('Tour updated successfully')
    await refetch()
  } catch (err) {
    console.log(err)
    toast.error('Something went wrong')
  }
})

watch(derivedData, (newData) => {
  if (!newData) return
  setValues({
    banner: {
      newFile: null,
      uploadedFile: newData.banner,
    },
    images: {
      newFiles: [],
      deletedFiles: [],
      uploadedFiles: newData.images,
    },
  })
})
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle>Tour Images</CardTitle>
    </CardHeader>
    <CardContent>
      <p v-if="isLoading" class="text-gray-500">Loading...</p>
      <p v-if="!derivedData" class="text-gray-500">The tour has not been found</p>
      <form v-if="derivedData" class="grid md:grid-cols-2 gap-2 min-h-[40vh]" @submit.prevent="onSubmit">
        <FormField v-slot="{ componentField }" name="banner">
          <FormItem class="flex flex-col">
            <FormLabel>Banner</FormLabel>
            <FormDescription>Recommended aspect ratio: 16:9</FormDescription>
            <FormControl>
              <ImageInput
                v-model="componentField.modelValue"
                @update:model-value="setFieldValue('banner', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="images">
          <FormItem class="flex flex-col">
            <FormLabel>Images</FormLabel>
            <FormDescription>Recommended aspect ratio: 4:4</FormDescription>
            <FormControl>
              <ImagesInput
                v-model="componentField.modelValue"
                @update:model-value="setFieldValue('images', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
        <Button type="submit" :disabled="isPending">Submit all images</Button>
      </form>
    </CardContent>
    <CardFooter> </CardFooter>
  </Card>
</template>
