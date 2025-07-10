<script setup lang="ts">
import { Input } from '@/core/components/ui/input'
import { EditIcon, Search } from 'lucide-vue-next'
import { buttonVariants } from '@/core/components/ui/button'
import { cn } from '@/core/lib/utils'
import { onMounted } from 'vue'
import { Draggable } from '@fullcalendar/interaction'
import ItemsPaginator from '../paginator/ItemsPaginator.vue'
import DestinationFormDialog from '@/destination/components/DestinationFormDialog.vue'

onMounted(() => {
  new Draggable(document.getElementById('destination-draggable')!, {
    itemSelector: '.destination',
    eventData: function (element) {
      const data = JSON.parse(element.getAttribute('props') ?? '{}')
      return {
        title: 'event 1',
        date: '2025-04-04',
        backgroundColor: '#e5f1fe',
        //borderColor: '#d8e6f7',
        borderColor: '#005bbf',
        textColor: '#005bbf',
        ...data,
      }
    },
  })
})

const data = {
  title: 'Copacabana',
}
</script>

<template>
  <div class="relative w-full items-center">
    <Input id="search" type="text" placeholder="Search..." class="pl-10" />
    <span class="absolute start-0 inset-y-0 flex items-center justify-center px-2">
      <Search class="size-6 text-muted-foreground" />
    </span>
  </div>

  <div id="destination-draggable" class="grid gap-2">
    <div
      v-for="i in 2"
      :key="i"
      class="destination grid border border-gray-200 rounded-md"
      :props="JSON.stringify(data)"
    >
      <div class="grid gap-2 p-4 bg-muted/50">
        <ul class="grid gap-1">
          <li class="text-lg font-bold">Copacabana</li>
          <li class="text-sm text-muted-foreground">Description:</li>
        </ul>
        <div>
          <DestinationFormDialog> <EditIcon /> Edit </DestinationFormDialog>
        </div>
      </div>
    </div>
  </div>

  <ItemsPaginator :items-per-page="2" :total="10" />
</template>
