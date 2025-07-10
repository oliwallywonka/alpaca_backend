<script setup lang="ts">
import { Input } from '@/core/components/ui/input'
import { EditIcon, Search } from 'lucide-vue-next'
import { onMounted, ref } from 'vue'
import { Draggable } from '@fullcalendar/interaction'
import ItemsPaginator from '../paginator/ItemsPaginator.vue'
import ResourceFormDialog from '@/resource/components/ResourceFormDialog.vue'

const element = ref<Draggable>()

onMounted(() => {
  element.value = new Draggable(document.getElementById('service-draggable')!, {
    itemSelector: '.service',
    eventData: function (element) {
      const data = JSON.parse(element.getAttribute('props') ?? '{}')
      return {
        title: 'event 1',
        backgroundColor: '#fefbe8',
        //borderColor: '#eae3d1',
        borderColor: '#c58220',
        textColor: '#c58220',
        ...data,
      }
    },
  })
})

const data1 = {
  title: 'Transfer Copacabana IN / OUT',
  description: 'A bus transfer from the bus station to the hotel in copacabana',
}
</script>

<template>
  <div class="relative w-full items-center">
    <Input id="search" type="text" placeholder="Search..." class="pl-10" />
    <span class="absolute start-0 inset-y-0 flex items-center justify-center px-2">
      <Search class="size-6 text-muted-foreground" />
    </span>
  </div>

  <div class="grid gap-2" id="service-draggable">
    <div v-for="i in 2" :key="i" class="grid border border-gray-200 rounded-md">
      <div class="grid gap-2 p-4 bg-muted/50">
        <ul class="grid gap-1">
          <li class="text-lg font-bold">Transfer Copacabana IN / OUT</li>
          <li class="text-sm text-muted-foreground">
            Description: A bus transfer from the bus station to the hotel in copacabana
          </li>
        </ul>
        <ul class="flex flex-wrap gap-2">
          <li v-for="i in 4" :key="i" class="service text-sm text-muted-foreground p-2 border rounded-sm">
            Provider/User: Fernando Camacho:
          </li>
        </ul>
        <div class="flex flex-wrap gap-2">
          <ResourceFormDialog> <EditIcon /> Edit </ResourceFormDialog>
        </div>
      </div>
    </div>
  </div>

  <ItemsPaginator :items-per-page="2" :total="10" />
</template>
