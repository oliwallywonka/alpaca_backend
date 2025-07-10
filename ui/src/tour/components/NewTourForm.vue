<script setup lang="ts">
import { watch } from 'vue'
import { useForm } from 'vee-validate'
import { useQueryClient } from '@tanstack/vue-query'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'

import { Button } from '@/core/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/core/components/ui/dialog'
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/core/components/ui/form'
import { Input } from '@/core/components/ui/input'
import { toast } from 'vue-sonner'
import { PlusIcon } from 'lucide-vue-next'
import { TourService } from '../services/tourService'
import { generateCode, generateSlug } from '@/core/lib/text'

const isOpen = defineModel({ default: false })

const tourDefaultValues = {
  transport: {
    en: '',
  },
  accommodation: {
    en: '',
  },
  team: {
    en: '',
  },
  groupSize: '',
  days: 0,
  shortDescription: {
    en: '',
  },
  longDescription: {
    en: '',
  },
}

const createQuery = TourService.useCreate()
const queryClient = useQueryClient()

const formSchema = toTypedSchema(
  z.object({
    name: z.string().min(3, 'Name must have at least 3 characters'),
    slug: z.string(),
    code: z
      .string()
      .min(5, 'Code must have at least 5 characters')
      .max(15, 'Code must have at most 10 characters'),
  }),
)

const { values, handleSubmit, setFieldValue, handleReset } = useForm({
  validationSchema: formSchema,
  initialValues: {
    code: generateCode(),
    name: '',
    slug: '',
  },
})

const onSubmit = handleSubmit(async (values) => {
  // TODO: Verify if Name already exists
  try {
    await createQuery.mutateAsync({
      code: values.code,
      name: {
        en: values.name,
      },
      slug: {
        en: values.slug,
      },
      ...tourDefaultValues,
    })
    toast.success('Tour created successfully')
    queryClient.invalidateQueries({ queryKey: ['tours'] })
  } catch (err) {
    console.log(err)
    toast.error('Something went wrong')
  } finally {
    isOpen.value = false
  }
})

watch(isOpen, (newOpen) => {
  if (!newOpen) return
  handleReset()
})

watch(
  () => values.name,
  () => {
    if (!values.name) return
    setFieldValue('slug', generateSlug(values.name))
  },
)
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="outline">
        <PlusIcon class="w-4 h-4" />
        New Tour
      </Button>
    </DialogTrigger>
    <DialogContent class="sm:max-w-4xl">
      <DialogHeader>
        <DialogTitle>New Tour</DialogTitle>
        <DialogDescription>
          Create a new tour by filling out the form below. Make sure to provide all necessary
          details.
        </DialogDescription>
      </DialogHeader>

      <form id="dialogForm" @submit.prevent="onSubmit" class="grid gap-4">
        <FormField v-slot="{ componentField }" name="code">
          <FormItem>
            <FormLabel>Code</FormLabel>
            <FormControl>
              <div class="grid grid-cols-[1fr_auto] gap-1">
                <Input type="text" placeholder="Tour code" v-bind="componentField" />
                <Button
                  variant="ghost"
                  type="button"
                  @click="() => setFieldValue('code', generateCode())"
                >
                  Generate Code
                </Button>
              </div>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
        <FormField v-slot="{ componentField }" name="name">
          <FormItem>
            <FormLabel>En Name</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Tour name" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="slug">
          <FormItem>
            <FormLabel>En Slug</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Tour slug" v-bind="componentField" disabled />
            </FormControl>
            <FormDescription> This is the tour url </FormDescription>
            <FormMessage />
          </FormItem>
        </FormField>
      </form>

      <DialogFooter>
        <Button type="submit" form="dialogForm"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
