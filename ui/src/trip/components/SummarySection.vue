<script setup lang="ts">
import { Card, CardContent } from '@/core/components/ui/card'
import { Badge } from '@/core/components/ui/badge'
import { ArrowDownLeft, ArrowUpRight } from 'lucide-vue-next'
</script>

<template>
  <section class="grid grid-cols-4 gap-4">
    <Card>
      <CardContent class="grid gap-1">
        <span class="text-sm text-black/60">Total Payments:</span>
        <div class="flex gap-2">
          <span class="text-xl font-semibold">SUS 4.250</span>
          <Badge variant="green">40% <ArrowUpRight class="text-green-500" /></Badge>
        </div>
        <span class="text-sm text-black/60">SUS 1243 Left</span>
      </CardContent>
    </Card>

    <Card>
      <CardContent class="grid gap-1">
        <span class="text-sm text-black/60">Total Expenses:</span>
        <div class="flex gap-2">
          <span class="text-xl font-semibold">SUS 4.250</span>
          <Badge variant="red">+10% <ArrowDownLeft /></Badge>
        </div>
        <span class="text-sm text-black/60">SUS 150 Left</span>
      </CardContent>
    </Card>
    <Card>
      <CardContent class="grid gap-1">
        <span class="text-sm text-black/60">Total Revenue:</span>
        <div class="flex gap-2">
          <span class="text-xl font-semibold">SUS 465</span>
          <Badge variant="green">+10% <ArrowUpRight /></Badge>
        </div>
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
</template>
