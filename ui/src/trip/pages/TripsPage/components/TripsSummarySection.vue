<script setup lang="ts">
import { ToggleGroup, ToggleGroupItem } from '@/core/components/ui/toggle-group'
import { Badge } from '@/core/components/ui/badge'
</script>

<template>
  <ToggleGroup type="single" class="w-full grid grid-cols-5 gap-2">
    <ToggleGroupItem value="all" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Trips</span>
      <Badge variant="outline">All</Badge>
      <span>Total: 10</span>
    </ToggleGroupItem>
    <ToggleGroupItem value="booked" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Trips</span>
      <Badge variant="blue">Booked</Badge>
      <span>Total: 4</span>
    </ToggleGroupItem>

    <ToggleGroupItem value="inProgress" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Trips</span>
      <Badge variant="green">In Progress</Badge>
      <span>Total: 2</span>
    </ToggleGroupItem>

    <ToggleGroupItem value="cancelled" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Trips</span>
      <Badge variant="red">Cancelled</Badge>
      <span>Total: 5</span>
    </ToggleGroupItem>

    <ToggleGroupItem value="finished" class="grid justify-items-start gap-2 border h-auto p-2">
      <span>Trips</span>
      <Badge variant="purple">Finished</Badge>
      <span>Total: 3</span>
    </ToggleGroupItem>
  </ToggleGroup>
</template>
