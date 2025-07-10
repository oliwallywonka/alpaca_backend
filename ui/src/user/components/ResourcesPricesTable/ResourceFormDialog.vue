<script setup lang="ts">
import { watch } from 'vue'
import { toast } from 'vue-sonner'
import { useQueryClient } from '@tanstack/vue-query'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'
import { useForm } from 'vee-validate'

import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/core/components/ui/dialog'

import {
  Combobox,
  ComboboxAnchor,
  ComboboxEmpty,
  ComboboxGroup,
  ComboboxInput,
  ComboboxItem,
  ComboboxItemIndicator,
  ComboboxList,
  ComboboxTrigger,
} from '@/core/components/ui/combobox'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/core/components/ui/select'
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
  FormFieldArray,
} from '@/core/components/ui/form'

import { Input } from '@/core/components/ui/input'
import { Button } from '@/core/components/ui/button'
import { Switch } from '@/core/components/ui/switch'
import { PlusIcon, Check, ChevronsUpDown, X } from 'lucide-vue-next'
import { cn } from '@/core/lib/utils'

import type { Price } from '@/user/interfaces/resourcePrices'
import { ResourceService } from '@/resource/services/resourceService'
import { useParams } from '@/core/hooks/useParams'
import type { Resource } from '@/resource/interfaces/resource'
import { UserService } from '@/user/services/userService'
import { ProviderService } from '@/provider/services/providerService'
import { ProviderTypes } from '@/core/constants/providerTypes'

const isOpen = defineModel<boolean>({ default: false })

const props = defineProps<{
  resourceProviderID?: string
  providerID?: string
  userID?: string
  prices?: Price[]
}>()

const { params } = useParams()
const servicesQuery = ResourceService.useGetAll(params)
const createUser = UserService.useCreateResource()
const createProvider = ProviderService.useCreateResource()
const queryClient = useQueryClient()

const formSchema = toTypedSchema(
  z.object({
    resource: z.string().min(1, 'Service is required'),
    type: z.string().min(1, 'Type is required'),
    refPrices: z.array(
      z.object({
        minPersons: z.number().min(1, 'Min persons is required'),
        maxPersons: z.number().min(1, 'Max persons is required'),
        price: z.number().min(1, 'Price is required and cannot be zero'),
        currency: z.string().min(1, 'Currency is required'),
        isPerPerson: z.boolean(),
      }),
    ),
  }),
)

const initialValues = {
  resource: '',
  type: '',
  refPrices: props.prices || [],
}

const { handleSubmit, setFieldValue, resetForm } = useForm({
  validationSchema: formSchema,
  initialValues: initialValues,
})

const onSubmit = handleSubmit(async (values) => {
  console.log(values)
  try {
    if (props.userID) {
      await createUser.mutateAsync({
        userID: props.userID || '',
        data: {
          ...values,
          user: props.userID || '',
        },
      })
      toast.success('Resource Prices added successfully')
      queryClient.invalidateQueries({ queryKey: ['user_resources'] })
    }
    if (props.providerID) {
      await createProvider.mutateAsync({
        providerID: props.providerID || '',
        data: {
          ...values,
          provider: props.providerID || '',
        },
      })
      toast.success('Resource Prices added successfully')
      queryClient.invalidateQueries({ queryKey: ['provider_resources'] })
    }
  } catch (error) {
    console.error('Error submitting form:', error)
    toast.error('Failed to save service prices')
  } finally {
    isOpen.value = false
  }
})

watch([props.prices, isOpen], ([newPrices, newIsOpen]) => {
  if (!newPrices || !newIsOpen) return
  resetForm()
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="outline">
        <slot> Add </slot>
      </Button>
    </DialogTrigger>
    <DialogContent class="w-full max-w-5xl">
      <DialogHeader>
        <DialogTitle>
          {{
            props.resourceProviderID ? 'Update Resource Prices List' : 'Create Resource Prices List'
          }}
        </DialogTitle>
        <DialogDescription> * Are mandatory fields {{ props }} </DialogDescription>
      </DialogHeader>
      <form id="resource-provider-form" @submit.prevent="onSubmit" class="grid grid-cols-3 gap-4">
        <fieldset class="flex flex-col gap-2">
          <FormField name="resource">
            <FormItem class="flex flex-col">
              <FormLabel>Resource</FormLabel>
              <Combobox by="label">
                <FormControl>
                  <ComboboxAnchor class="w-full">
                    <div class="relative w-full items-center">
                      <ComboboxInput
                        :display-value="
                          (val: Resource) => {
                            return val?.name?.en ?? ''
                          }
                        "
                        placeholder="Select a service..."
                      />
                      <ComboboxTrigger
                        class="absolute end-0 inset-y-0 flex items-center justify-center px-3"
                      >
                        <ChevronsUpDown class="size-4 text-muted-foreground" />
                      </ComboboxTrigger>
                    </div>
                  </ComboboxAnchor>
                </FormControl>
                <ComboboxList class="w-[300px]">
                  <ComboboxEmpty> Nothing found. </ComboboxEmpty>
                  <ComboboxGroup>
                    <ComboboxItem
                      v-for="item in servicesQuery.data.value?.items || []"
                      :key="item.id"
                      :value="item"
                      @select="
                        () => {
                          setFieldValue('resource', item.id)
                        }
                      "
                    >
                      {{ item.name.en }}
                      <ComboboxItemIndicator>
                        <Check :class="cn('ml-auto h-4 w-4')" />
                      </ComboboxItemIndicator>
                    </ComboboxItem>
                  </ComboboxGroup>
                </ComboboxList>
              </Combobox>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="type">
            <FormItem>
              <FormLabel>Type</FormLabel>
              <Select v-bind="componentField" defaultValue="service">
                <FormControl>
                  <SelectTrigger class="w-full">
                    <SelectValue placeholder="Select a resource Type" />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  <SelectGroup>
                    <SelectItem v-for="type of ProviderTypes" :key="type.value" :value="type.value">
                      {{ type.label }}
                    </SelectItem>
                  </SelectGroup>
                </SelectContent>
              </Select>
              <FormMessage />
            </FormItem>
          </FormField>
        </fieldset>

        <div class="col-span-2 grid gap-2 py-4">
          <FormFieldArray name="refPrices" v-slot="{ fields, push, remove }">
            <Button
              @click="
                push({
                  minPersons: 1,
                  maxPersons: 4,
                  price: 0,
                  currency: 'BOL',
                  isPerPerson: false,
                })
              "
              class="mb-4"
            >
              <PlusIcon /> Add Price
            </Button>
            <fieldset
              v-for="(field, index) in fields"
              :key="field.key"
              class="flex gap-2 items-start"
            >
              <FormField :name="`refPrices.${index}.minPersons`" v-slot="{ componentField }">
                <FormItem>
                  <FormLabel>Min Persons</FormLabel>
                  <FormControl>
                    <Input type="number" v-bind="componentField" />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField :name="`refPrices[${index}].maxPersons`" v-slot="{ componentField }">
                <FormItem>
                  <FormLabel>Max Persons</FormLabel>
                  <FormControl>
                    <Input type="number" v-bind="componentField" />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField :name="`refPrices[${index}].price`" v-slot="{ componentField }">
                <FormItem>
                  <FormLabel>Price</FormLabel>
                  <FormControl>
                    <Input type="number" v-bind="componentField" />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField :name="`refPrices[${index}].currency`" v-slot="{ componentField }">
                <FormItem>
                  <FormLabel>Currency</FormLabel>
                  <Select v-bind="componentField" defaultValue="BOL">
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select currency" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectGroup>
                        <SelectItem value="USD"> USD </SelectItem>
                        <SelectItem value="BOL"> BOL </SelectItem>
                      </SelectGroup>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              </FormField>

              <FormField v-slot="{ value, handleChange }" :name="`refPrices[${index}].isPerPerson`">
                <FormItem>
                  <FormLabel> Is Per Person</FormLabel>
                  <FormControl>
                    <Switch :model-value="value" aria-readonly @update:model-value="handleChange" />
                  </FormControl>
                </FormItem>
              </FormField>

              <Button class="self-center" type="button" @click="remove(index)"><X /></Button>
            </fieldset>
          </FormFieldArray>
        </div>
      </form>
      <DialogFooter>
        <DialogClose as-child>
          <Button variant="outline"> Cancel </Button>
        </DialogClose>
        <Button type="submit" form="resource-provider-form"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
