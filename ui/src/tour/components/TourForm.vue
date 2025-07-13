<script setup lang="ts">
import { useRoute } from 'vue-router'
import { toRef, watch } from 'vue'
import { toast } from 'vue-sonner'
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'

import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/core/components/ui/form'
import {
  CardDescription,
  CardHeader,
  CardTitle,
  Card,
  CardContent,
} from '@/core/components/ui/card'
import { Switch } from '@/core/components/ui/switch'
import { Button } from '@/core/components/ui/button'
import { LanguageInput, LanguageTextArea } from '@/core/components/fields'
import Input from '@/core/components/ui/input/Input.vue'
import { TourService } from '../services/tourService'
import { parseDate } from '@/core/lib/date'

const route = useRoute()

const { data, isLoading, refetch } = TourService.useGetOne(
  toRef(route.params.tourID as string),
  toRef(true),
)
const { mutateAsync, isPending } = TourService.useUpdate()

const formSchema = toTypedSchema(
  z.object({
    name: z.record(z.string(), z.string().min(1, 'Every lang must has at least 1 character')),
    slug: z.record(z.string(), z.string().min(1, 'Every lang must has at least 1 character')),
    shortDescription: z.record(
      z.string(),
      z.string().min(1, 'Every lang must has at least 1 character'),
    ),
    longDescription: z.record(
      z.string(),
      z.string().min(1, 'Every lang must has at least 1 character'),
    ),
    team: z.record(z.string(), z.string().min(1, 'Every lang must has at least 1 character')),
    transport: z.record(z.string(), z.string().min(1, 'Every lang must has at least 1 character')),
    accommodation: z.record(
      z.string(),
      z.string().min(1, 'Every lang must has at least 1 character'),
    ),
    groupSize: z.string(),
    daysDuration: z.number(),
    isPublic: z.boolean(),
    isActive: z.boolean(),
  }),
)

const form = useForm({
  validationSchema: formSchema,
  initialValues: {
    name: data.value?.name,
    slug: data.value?.slug,
    shortDescription: data.value?.shortDescription,
    longDescription: data.value?.longDescription,
    transport: data.value?.transport,
    team: data.value?.team,
    accommodation: data.value?.accommodation,
    groupSize: data.value?.groupSize,
    daysDuration: data.value?.days,
    isPublic: data.value?.isPublic,
    isActive: data.value?.isActive,
  },
})

const onSubmit = form.handleSubmit(async (values) => {
  if (!data.value) return
  try {
    await mutateAsync({
      id: data.value?.id,
      data: values,
    })
    await refetch()
    toast.success('Tour updated successfully')
  } catch (err) {
    console.log(err)
    toast.error('Something went wrong')
  }
})

watch(data, (newData) => {
  if (!newData) return
  form.setValues({
    name: newData.name,
    slug: newData.slug,
    shortDescription: newData.shortDescription,
    longDescription: newData.longDescription,
    transport: newData.transport,
    team: newData.team,
    accommodation: newData.accommodation,
    groupSize: newData.groupSize,
    daysDuration: newData.days,
    isPublic: newData.isPublic,
    isActive: newData.isActive,
  })
})
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle>Tour Basic Information </CardTitle>
      <CardTitle>{{ data ? data.name.en.toUpperCase() : '' }}</CardTitle>
      <CardDescription>
        Last Modification:
        {{ data ? parseDate(data.updated) : '' }} {{ form.values }}
      </CardDescription>
    </CardHeader>
    <CardContent class="grid gap-2">
      <p v-if="isLoading" class="text-gray-500">Loading...</p>
      <p v-if="!data" class="text-gray-500">The tour has not been found</p>
      <form v-if="data" @submit="onSubmit" class="md:cols-span-2 grid md:grid-cols-2 gap-4">
        <fieldset class="md:col-span-2">
          <Button type="submit" :disabled="isPending"> Submit </Button>
        </fieldset>
        <FormField v-slot="{ componentField }" name="isPublic">
          <FormItem>
            <FormLabel>Publish</FormLabel>
            <FormControl>
              <Switch
                :model-value="componentField.modelValue"
                @update:model-value="form.setFieldValue('isPublic', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="isActive">
          <FormItem>
            <FormLabel>Active</FormLabel>
            <FormControl>
              <Switch
                :model-value="componentField.modelValue"
                @update:model-value="form.setFieldValue('isActive', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="name">
          <FormItem>
            <FormLabel>Tour Name</FormLabel>
            <FormControl>
              <LanguageInput
                v-model="componentField.modelValue"
                @update:model-value="form.setFieldValue('name', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="slug">
          <FormItem>
            <FormLabel>Slug URL</FormLabel>
            <FormControl>
              <LanguageInput
                v-model="componentField.modelValue"
                @update:model-value="form.setFieldValue('slug', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
        <FormField v-slot="{ componentField }" name="transport">
          <FormItem>
            <FormLabel>Transport</FormLabel>
            <FormControl>
              <LanguageInput
                v-model="componentField.modelValue"
                @update:model-value="form.setFieldValue('transport', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="accommodation">
          <FormItem>
            <FormLabel>Accomodation</FormLabel>
            <FormControl>
              <LanguageInput
                v-model="componentField.modelValue"
                @update:model-value="form.setFieldValue('accommodation', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="team">
          <FormItem>
            <FormLabel>Tour Team</FormLabel>
            <FormControl>
              <LanguageInput
                v-model="componentField.modelValue"
                @update:model-value="form.setFieldValue('team', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <div></div>

        <FormField v-slot="{ componentField }" name="groupSize">
          <FormItem>
            <FormLabel>Group Size</FormLabel>
            <FormControl>
              <Input v-bind="componentField"/>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="daysDuration">
          <FormItem>
            <FormLabel>Days Duration</FormLabel>
            <FormControl>
              <Input v-bind="componentField" type="number" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="shortDescription">
          <FormItem>
            <FormLabel>Short Description</FormLabel>
            <FormControl>
              <LanguageTextArea
                v-model="componentField.modelValue"
                @update:model-value="form.setFieldValue('shortDescription', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="longDescription">
          <FormItem>
            <FormLabel>Long Description</FormLabel>
            <FormControl>
              <LanguageTextArea
                v-model="componentField.modelValue"
                @update:model-value="form.setFieldValue('longDescription', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </form>
    </CardContent>
  </Card>
</template>
