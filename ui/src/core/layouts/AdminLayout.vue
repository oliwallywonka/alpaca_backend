<script setup lang="ts">
import { useRouter } from 'vue-router'

import AppSidebar from '@/core/components/AppSidebar.vue'
import { Separator } from '@/core/components/ui/separator'
import { SidebarInset, SidebarProvider, SidebarTrigger } from '@/core/components/ui/sidebar'
import { UserService } from '@/user/services/userService'
import { watch, ref, onMounted } from 'vue'

const router = useRouter()

const authStore = ref(UserService.authStore)

onMounted(() => {
  if (!UserService.authStore.isValid) {
    router.push('/login')
  }
})

watch(authStore, (newAuthStore) => {
  if (!newAuthStore.isValid ) {
    router.push('/login')
  }
}, {
  immediate: true
})
</script>
<template>
  <SidebarProvider>
    <AppSidebar />
    <SidebarInset>
      <header
        class="flex pt-2 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12"
      >
        <div class="flex items-center gap-2 px-4">
          <SidebarTrigger class="-ml-1" />
          <Separator orientation="vertical" class="mr-2 h-4" />
        </div>
      </header>
      <main class="mx-4 grid gap-4">
        <RouterView />
      </main>
    </SidebarInset>
  </SidebarProvider>
</template>
