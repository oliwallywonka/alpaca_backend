<script setup lang="ts">
import { useForm } from 'vee-validate'
import { z } from 'zod'
import { toast } from 'vue-sonner'
import { useRouter } from 'vue-router'

import { UserService } from '@/user/services/userService'

import { Button } from '@/core/components/ui/button'
import { Input } from '@/core/components/ui/input'
import { toTypedSchema } from '@vee-validate/zod'
import { FormControl, FormField, FormItem, FormMessage, FormLabel } from '@/core/components/ui/form'

const router = useRouter()
const { isPending, mutateAsync } = UserService.useAuth()

const loginSchema = toTypedSchema(
  z.object({
    identity: z.string().email(),
    password: z.string().min(6),
  }),
)

const { handleSubmit } = useForm({
  validationSchema: loginSchema,
  initialValues: {
    identity: '',
    password: '',
  },
})

const onSubmit = handleSubmit(async (values) => {
  try {
    await mutateAsync(values)
    toast.success('Login Successful')
    router.push('/')
  } catch (err) {
    console.log(err)
    toast.error('Error logging in')
  }
})
</script>
<template>
  <div class="w-full lg:grid lg:min-h-[600px] lg:grid-cols-2 xl:min-h-[800px]">
    <div class="flex items-center justify-center py-12">
      <div class="mx-auto grid w-[350px] gap-6">
        <div class="grid gap-2 text-center">
          <h1 class="text-3xl font-bold">Travel Agency Management System</h1>
          <p class="text-balance text-muted-foreground">
            Welcome!👌 Login to your account to start.
          </p>
        </div>
        <form id="loginForm" class="grid gap-4" @submit.prevent="onSubmit">
          <FormField v-slot="{ componentField }" name="identity">
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input type="email" v-bind="componentField" />
                <FormMessage />
              </FormControl>
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="password">
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input type="password" v-bind="componentField" />
                <FormMessage />
              </FormControl>
            </FormItem>
          </FormField>

          <Button type="submit" class="w-full" :disabled="isPending"> Login </Button>
        </form>
      </div>
    </div>
    <div class="hidden lg:block w-full h-screen lg:relative">
      <span
        class="absolute top-0 left-0 z-10 flex h-full w-full items-center justify-center text-6xl font-bold text-white"
      >
        ALPACA FACE
      </span>
      <figure class="p-4 h-full w-full">
        <img
          src="/login-image.jpg"
          alt="Image"
          width="1920"
          height="1080"
          class="h-full w-full object-cover brightness-90 dark:brightness-[0.2] dark:grayscale rounded-lg"
        />
      </figure>
    </div>
  </div>
</template>
