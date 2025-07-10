<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import { Input } from '@/core/components/ui/input'
import { Label } from '@/core/components/ui/label'
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
        <div class="grid gap-4">
          <div class="grid gap-2">
            <Label for="email">Email</Label>
            <Input id="email" type="email" placeholder="m@example.com" required />
          </div>
          <div class="grid gap-2">
            <div class="flex items-center">
              <Label for="password">Password</Label>
              <a href="/forgot-password" class="ml-auto inline-block text-sm underline">
                Forgot your password?
              </a>
            </div>
            <Input id="password" type="password" required />
          </div>
          <Button type="submit" class="w-full"> Login </Button>
        </div>
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
