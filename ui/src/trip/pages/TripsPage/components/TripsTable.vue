<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { cn } from '@/core/lib/utils'
import { Search } from 'lucide-vue-next'

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableHeader,
  TableCaption,
} from '@/core/components/ui/table'
import { buttonVariants } from '@/core/components/ui/button'
import { Input } from '@/core/components/ui/input'
import { Badge } from '@/core/components/ui/badge'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
</script>
<template>
  <section class="grid gap-4">
    <div class="relative w-full items-center">
      <Input id="search" type="text" placeholder="Search..." class="pl-10" />
      <span class="absolute start-0 inset-y-0 flex items-center justify-center px-2">
        <Search class="size-6 text-muted-foreground" />
      </span>
    </div>
    <Table class="min-w-[800px] w-full">
      <TableCaption>Trips registered</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Tour</TableHead>
          <TableHead>Customer Lead</TableHead>
          <TableHead>Nº of persons</TableHead>
          <TableHead>Booking date</TableHead>
          <TableHead>Tour Start</TableHead>
          <TableHead>Tour End</TableHead>
          <TableHead>State</TableHead>
          <TableHead>Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="index in 4" :key="index">
          <TableCell>Uyuni Salar 4 Days</TableCell>
          <TableCell>Roberto Alonso</TableCell>
          <TableCell>5</TableCell>
          <TableCell>06/06/2022: 10:00</TableCell>
          <TableCell>06/06/2022: 10:00</TableCell>
          <TableCell>06/06/2022: 10:00</TableCell>
          <TableCell>
            <Badge variant="green">Confirmed</Badge>
          </TableCell>
          <TableCell class="grid gap-2">
            <RouterLink :to="`/trips/${1}`" :class="cn(buttonVariants())"> Details </RouterLink>
          </TableCell>
        </TableRow>
      </TableBody>
    </Table>

    <ItemsPaginator :items-per-page="10" :total="80" />
  </section>
</template>
