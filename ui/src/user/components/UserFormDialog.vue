<script setup lang="ts">
import { computed, toRef, watch } from 'vue'
import { toast } from 'vue-sonner'
import { useQueryClient } from '@tanstack/vue-query'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'
import { useForm } from 'vee-validate'

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/core/components/ui/dialog'
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/core/components/ui/form'

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/core/components/ui/select'
import { Input } from '@/core/components/ui/input'
import { Button } from '@/core/components/ui/button'

import { UserService } from '../services/userService'
import { RoleService } from '@/role/services/roleService'
import ContactField from '../../core/components/fields/ContactField.vue'
import { PlusIcon } from 'lucide-vue-next'

const isOpen = defineModel<boolean>({ default: false })

const props = defineProps<{
  userID?: string
}>()

const startFetching = computed(() => !!props.userID && isOpen.value)

const queryClient = useQueryClient()
const { data, isLoading } = RoleService.useGetAll()
const { data: userData } = UserService.useGetOne(toRef(props.userID!), startFetching)
const { mutateAsync: createMutate } = UserService.useCreate()
const { mutateAsync: updateMutate } = UserService.useUpdate()

const formSchema = toTypedSchema(
  z
    .object({
      role: z.string().min(1, 'Role is required'),
      name: z.string().min(1, 'Name is required'),
      email: z.string().email('Invalid email').min(1, 'Email is required'),
      password: z.string().min(1, 'Password is required'),
      passwordConfirm: z.string().min(1, 'Confirm Password is required'),
      contacts: z.array(z.object({ type: z.string(), value: z.string() })),
    })
    .refine((data) => data.password === data.passwordConfirm, {
      message: 'Passwords do not match',
      path: ['Password', 'ConfirmPassword'],
    }),
)

const { handleSubmit,values, setFieldValue, setValues } = useForm({
  validationSchema: formSchema,
  initialValues: {
    name: userData?.value?.name || '',
    role: userData.value?.role || '',
    email: userData.value?.email || '',
    password: userData?.value?.password || '',
    passwordConfirm: userData?.value?.password || '',
    contacts: userData?.value?.contacts || [],
  },
})

const onSubmit = handleSubmit(async (values) => {
  try {
    if (props.userID) {
      await updateMutate({ id: props.userID!, data: values })
      toast.success('User updated successfully')
    }

    if (!props.userID) {
      await createMutate(values)
      toast.success('User created successfully')
    }
  } catch (err) {
    console.log(err)
    toast.error('Error creating user')
  } finally {
    queryClient.invalidateQueries({ queryKey: ['users'] })
    isOpen.value = false
  }
})

watch([isOpen, userData], ([newIsOpen, newUserData]) => {
  if (!newIsOpen || !newUserData) return
  setValues({
    name: newUserData?.name || '',
    role: newUserData?.role || '',
    email: newUserData?.email || '',
    password: '',
    passwordConfirm: '',
    contacts: newUserData?.contacts || [],
  })
})
</script>
<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="outline"><PlusIcon /> {{ props.userID ? 'Edit User' : 'New User' }} </Button>
    </DialogTrigger>
    <DialogContent class="w-full max-w-2xl">
      <DialogHeader>
        <DialogTitle>{{ props.userID ? 'Edit User' : 'New User' }} </DialogTitle>
        <DialogDescription>
          Fill the form to create a new user here. Click save when you're done.
          {{ values }}
        </DialogDescription>
      </DialogHeader>
      <form id="user-form" class="grid grid-cols-2 gap-4" @submit.prevent="onSubmit">
        <FormField v-slot="{ componentField }" name="name">
          <FormItem>
            <FormLabel>*Full Name</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Full Name" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="email">
          <FormItem>
            <FormLabel>*Email</FormLabel>
            <FormControl>
              <Input type="text" placeholder="example@example.com" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="password">
          <FormItem>
            <FormLabel>Password</FormLabel>
            <FormControl>
              <Input type="password" placeholder="" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="passwordConfirm">
          <FormItem>
            <FormLabel>Confirm Password</FormLabel>
            <FormControl>
              <Input type="password" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="role">
          <FormItem class="col-span-2">
            <FormLabel>*Role</FormLabel>
            <FormControl>
              <p v-if="isLoading">Loading Roles...</p>
              <Select v-if="!isLoading" v-bind="componentField">
                <SelectTrigger class="w-full">
                  <SelectValue placeholder="Select a role" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    <SelectItem v-for="role in data?.items" :key="role.id" :value="role.id">
                      {{ role.name }}
                    </SelectItem>
                  </SelectGroup>
                </SelectContent>
              </Select>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="contacts">
          <FormItem class="col-span-2">
            <FormLabel>Contacts</FormLabel>
            <FormControl>
              <ContactField
                v-model="componentField.modelValue"
                @update:model-value="setFieldValue('contacts', $event)"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </form>
      <DialogFooter>
        <Button form="user-form" type="submit"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
