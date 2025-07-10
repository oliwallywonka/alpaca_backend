<script setup lang="ts">
import { ref, watch, useAttrs } from 'vue'
import { useVModel } from '@vueuse/core'

import * as L from 'leaflet'
import 'leaflet/dist/leaflet.css'
import { Map } from 'lucide-vue-next'
import { useQuery, useQueryClient } from '@tanstack/vue-query'

import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/core/components/ui/tooltip'
import { Input } from '@/core/components/ui/input'
import { Button } from '@/core/components/ui/button'
import { SearchInput } from '@/core/components/fields'
import { Card, CardContent } from '@/core/components/ui/card'

const props = withDefaults(defineProps<{ modelValue?: { lat: number; lon: number } }>(), {
  modelValue: () => ({ lat: 0, lon: 0 }),
})
const emit = defineEmits<{
  (e: 'update:modelValue', value: { lat: number; lon: number }): void
}>()
const attr = useAttrs()

const point = useVModel(props, 'modelValue', emit, { passive: true })

const map = ref<L.Map>()
const marker = ref<L.Marker | null>(null)
const container = ref<HTMLElement | null>(null)

const searchValue = ref('')
const showMap = ref(false)

const queryClient = useQueryClient()

const { data, isLoading, refetch } = useQuery({
  queryKey: ['mapData'],
  refetchOnWindowFocus: false,
  refetchOnMount: false,
  refetchOnReconnect: false,
  staleTime: Infinity,
  queryFn: async () => {
    return await fetch(
      'https://nominatim.openstreetmap.org/search.php?format=jsonv2&q=' +
        encodeURIComponent(searchValue.value),
      {
        signal: AbortSignal.timeout(1000),
      },
    )
      .then(async (response) => await response.json())
      .then((data) => {
        return data as {
          place_id: string
          lat: string
          lon: string
          display_name: string
        }[]
      })
  },
})

watch(container, (newContainer) => {
  if (!newContainer || !showMap.value) return
  map.value = L.map(container.value!).setView(
    [point.value.lat === 0 ? -16 : point.value.lat, point.value.lon === 0 ? -67 : point.value.lon],
    6,
  )
  L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
  }).addTo(map.value)
  marker.value = L.marker([point.value.lat, point.value.lon]).addTo(map.value!)
  map.value.on('contextmenu', (e: L.LeafletMouseEvent) => {
    if (!marker.value) return
    marker.value.setLatLng(e.latlng)
    point.value.lat = normalizeCoordinate(e.latlng.lat)
    point.value.lon = normalizeCoordinate(e.latlng.lng)
  })
})

function toogleMap() {
  showMap.value = !showMap.value
}

function handleSearch(inputValue: string) {
  if (!inputValue) {
    queryClient.setQueryData(['mapData'], null)
    return
  }
  searchValue.value = inputValue
  refetch()
}

function handleSelect(item: { lat: string; lon: string }) {
  point.value.lat = normalizeCoordinate(Number(item.lat))
  point.value.lon = normalizeCoordinate(Number(item.lon))
  queryClient.setQueryData(['mapData'], null)
}

function normalizeCoordinate(coord: number) {
  return +(+coord).toFixed(6)
}

watch(
  point,
  (newPoint) => {
    if (!map.value) return
    marker.value?.setLatLng([newPoint.lat, newPoint.lon]).addTo(map.value!)
    map.value?.setView([newPoint.lat, newPoint.lon])
  },
  { deep: true },
)
</script>
<template>
  <div class="grid grid-cols-[1fr_1fr_auto] gap-4">
    <Input placeholder="Latitude 0" type="number" v-model="point.lat" v-bind="attr" step="any" />
    <Input placeholder="Longitude 0" type="number" v-model="point.lon" v-bind="attr" step="any" />
    <div class="flex gap-1">
      <TooltipProvider>
        <Tooltip>
          <TooltipTrigger as-child>
            <Button @click="toogleMap" type="button"> <Map /></Button>
          </TooltipTrigger>
          <TooltipContent>
            <p>Toogle Map</p>
          </TooltipContent>
        </Tooltip>
      </TooltipProvider>
    </div>
  </div>
  <div class="relative h-[400px] bg-green-200" v-if="showMap">
    <div ref="container" class="w-full h-full z-1" />
    <div class="z-1 absolute top-2 w-[70%] max-w-[600px] ml-[15%] flex flex-col gap-2">
      <SearchInput
        placeholder="Search address..."
        class="bg-white opacity-80"
        @change:debounce="(s) => handleSearch(s)"
        :debounceMilis="1100"
        :isLoading="isLoading"
      />
      <Card v-if="data && data.length > 0" class="rounded-sm py-0">
        <CardContent class="p-2 max-h-[200px] overflow-y-auto overflow-x-hidden">
          <Button
            v-for="item in data"
            :key="item.place_id"
            @click="() => handleSelect(item)"
            type="button"
            variant="ghost"
            class="w-full h-auto whitespace-normal justify-start text-left p-1"
          >
            {{ item.display_name }}
          </Button>
        </CardContent>
      </Card>
    </div>
  </div>
</template>
