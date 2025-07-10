<script setup lang="ts">
import { ref } from 'vue'

import LWChart, { type ChartType } from '@/core/components/charts/LWChart.vue'
import { colors } from '@/core/constants/colors'
import Badge from '@/core/components/ui/badge/Badge.vue'
import { ArrowDownLeft, ArrowUpRight } from 'lucide-vue-next'
import { Card, CardContent } from '@/core/components/ui/card'

const options = {
  color: '#005bbf',
  topColor: colors.blue.textColor,
  bottomColor: colors.blue.backgroundColor,
  lineColor: colors.blue.textColor,
}
/* const lwChart = ref<InstanceType<typeof LWChart> | null>(null) */
const chartOptions = ref({})

function toDateString(unix: number): string {
  const date = new Date(unix * 1000)
  return date.toISOString().split('T')[0]
}
const data = ref([
  { value: 1, time: toDateString(1642425322) },
  { value: 8, time: toDateString(1642511722) },
  { value: 10, time: toDateString(1642598122) },
  { value: 20, time: toDateString(1642684522) },
  { value: 43, time: toDateString(1642857322) },
  { value: 43, time: toDateString(1643030122) },
  { value: 56, time: toDateString(1643116522) },
])
const chartType = ref<ChartType>('area')

/* function addData() {
  const newData = {
    time: '2024-04-05',
    value: Math.random() * 100 * 5,
  }
  lwChart.value?.updateData(newData)
} */
</script>

<template>
  <section class="grid grid-cols-4 gap-4">
    <Card>
      <CardContent class="grid gap-1">
        <span class="text-sm text-black/60">Tours Revenue:</span>
        <div class="flex gap-2">
          <span class="text-xl font-semibold">SUS 4.250</span>
          <Badge variant="green">+10% <ArrowUpRight class="text-green-500" /></Badge>
        </div>
        <span class="text-sm text-black/60">+SUS1243 Revenue</span>
      </CardContent>
    </Card>

    <Card>
      <CardContent class="grid gap-1">
        <span class="text-sm text-black/60">Current Tours:</span>
        <div class="flex gap-2">
          <span class="text-xl font-semibold">80</span>
          <Badge variant="red">+10% <ArrowDownLeft /></Badge>
        </div>
        <span class="text-sm text-black/60">+16 Tours</span>
      </CardContent>
    </Card>
    <Card>
      <CardContent class="grid gap-1">
        <span class="text-sm text-black/60">Cash Flow:</span>
        <div class="flex gap-2">
          <span class="text-xl font-semibold">SUS 4.250</span>
          <Badge variant="green">+10% <ArrowUpRight /></Badge>
        </div>
        <span class="text-sm text-black/60">+SUS 450</span>
      </CardContent>
    </Card>
    <Card>
      <CardContent class="grid gap-1">
        <span class="text-sm text-black/60">Dolar Blue:</span>
        <div class="flex gap-2">
          <span class="text-xl font-semibold">Bs 13.20</span>
          <Badge variant="green">+10% <ArrowUpRight /></Badge>
        </div>
        <span class="text-sm text-black/60">+Bs 0.60 than yesterday</span>
      </CardContent>
    </Card>
  </section>
  <section class="grid gap-4">
    <Card>
      <CardContent class="grid gap-16">
        <div class="h-[300px] w-full">
          <span class="font-bold">Tours Revenue:</span>
          <LWChart
            :type="chartType"
            :seriesOptions="options"
            :data="data"
            :autosize="true"
            :chartOptions="chartOptions"
          />
        </div>

        <div class="h-[300px] w-full">
          <span class="font-bold">AVG: Cash Flow:</span>
          <LWChart
            type="histogram"
            :seriesOptions="options"
            :data="data"
            :autosize="true"
            :chartOptions="chartOptions"
          />
        </div>
      </CardContent>
    </Card>
  </section>
</template>
