<script setup lang="ts">
import AppSidebar from '@/core/components/AppSidebar.vue'
import { Separator } from '@/core/components/ui/separator'
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from '@/core/components/ui/sidebar'
</script>
<template>
  <SidebarProvider>
    <AppSidebar />
    <SidebarInset>
      <header class="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12">
        <div class="flex items-center gap-2 px-4">
          <SidebarTrigger class="-ml-1" />
          <Separator orientation="vertical" class="mr-2 h-4" />
        </div>
      </header>
      <main class="m-4 grid gap-4">
        <RouterView />
      </main>
    </SidebarInset>
  </SidebarProvider>
</template>