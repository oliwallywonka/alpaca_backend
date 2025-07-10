<script setup lang="ts">
import { cn } from '@/core/lib/utils'
import { buttonVariants } from '@/core/components/ui/button'
</script>
<template>
  <div class="hotel grid border border-gray-200 rounded-md">
    <div class="grid gap-2 p-4 bg-muted/50">
      <ul class="grid gap-1">
        <li class="text-lg font-bold">Hotel name</li>
        <li class="text-sm text-muted-foreground">Direction:</li>
        <li class="text-sm text-muted-foreground">Phone:</li>
      </ul>
      <div class="flex flex-wrap gap-2">
        <RouterLink to="/tours/1" :class="cn(buttonVariants())">Edit</RouterLink>
      </div>
    </div>
  </div>
</template>
