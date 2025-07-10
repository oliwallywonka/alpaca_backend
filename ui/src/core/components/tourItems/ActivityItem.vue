<script setup lang="ts">
import { cn } from '@/core/lib/utils'
import { buttonVariants } from '@/core/components/ui/button'
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from '@/core/components/ui/carousel'
</script>
<template>
  <div class="activity grid border border-gray-200 rounded-md">
    <Carousel class="relative w-full">
      <CarouselContent>
        <CarouselItem v-for="(_, index) in 5" :key="index">
          <figure class="h-[250px]">
            <img
              class="object-cover object-center w-full h-full rounded-t-md"
              src="https://images.unsplash.com/photo-1554629947-334ff61d85dc?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=387&q=80"
              alt="Tour Image"
            />
          </figure>
        </CarouselItem>
      </CarouselContent>
      <CarouselPrevious class="absolute left-0 bottom-0" />
      <CarouselNext class="absolute right-0 bottom-0" />
    </Carousel>
    <div class="grid gap-2 p-4 bg-muted/50">
      <ul class="grid gap-1">
        <li class="text-lg font-bold">Activity Title</li>
        <li class="text-sm text-muted-foreground">Description:</li>
        <li class="text-sm text-muted-foreground flex flex-wrap gap-0.5">
          Destination: <Badge>La Paz</Badge>
        </li>
        <li class="text-sm text-muted-foreground">Ref price:</li>
      </ul>
      <div class="flex flex-wrap gap-2">
        <RouterLink to="/tours/1" :class="cn(buttonVariants())">Edit</RouterLink>
      </div>
    </div>
  </div>
</template>
