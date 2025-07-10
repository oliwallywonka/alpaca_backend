<script setup lang="ts">
import type { SidebarProps } from '@/core/components/ui/sidebar'

import NavMain from '@/core/components/NavMain.vue'
import NavUser from '@/core/components/NavUser.vue'
import TeamSwitcher from '@/core/components/TeamSwitcher.vue'
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from '@/core/components/ui/sidebar'

import { sidebarData } from '@/core/router/sidebarLinks'

const props = withDefaults(defineProps<SidebarProps>(), {
  collapsible: 'icon',
})
</script>

<template>
  <Sidebar v-bind="props">
    <SidebarHeader>
      <TeamSwitcher :teams="sidebarData.teams" />
    </SidebarHeader>
    <SidebarContent>
      <NavMain :items="sidebarData.navMain" />
    </SidebarContent>
    <SidebarFooter>
      <NavUser :user="sidebarData.user" />
    </SidebarFooter>
    <SidebarRail />
  </Sidebar>
</template>
