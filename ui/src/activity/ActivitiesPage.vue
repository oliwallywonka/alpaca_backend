<script setup lang="ts">
import SearchInput from '@/core/components/fields/SearchInput.vue'
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
  TableCaption,
} from '@/core/components/ui/table'
import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationFirst,
  PaginationItem,
  PaginationLast,
  PaginationNext,
  PaginationPrevious,
} from '@/core/components/ui/pagination'
import { Badge } from '@/core/components/ui/badge'
import { Button } from '@/core/components/ui/button'
import { Checkbox } from '@/core/components/ui/checkbox'
</script>

<template>
  <SearchInput />
  <Table class="min-w-[800px] w-full">
    <TableCaption>Activities registered</TableCaption>
    <TableHeader>
      <TableRow>
        <TableHead class="sticky left-0.5 bg-inherit min-w-5"><Checkbox class="size-5" /></TableHead>
        <TableHead>Title</TableHead>
        <TableHead>Description</TableHead>
        <TableHead>Ref Price</TableHead>
        <TableHead>images</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow class="hover:bg-gray-100 hover:cursor-pointer">
        <TableCell class="sticky left-0.5 bg-inherit min-w-5"><Checkbox class="size-5" /></TableCell>
        <TableCell>Uyuni Salar 4 Days</TableCell>
        <TableCell>Roberto Alonso</TableCell>
        <TableCell><Badge>Uyuni</Badge></TableCell>
        <TableCell>
          <Badge>123545.54 | 1325466.5</Badge>
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>

  <Pagination
    v-slot="{ page }"
    :items-per-page="10"
    :total="80"
    :sibling-count="10"
    show-edges
    :default-page="1"
  >
    <PaginationContent v-slot="{ items }" class="grid md:grid-cols-[auto_1fr_auto] gap-2">
      <div>
        <PaginationFirst />
        <PaginationPrevious />
      </div>
      <div>
        <template v-for="(item, index) in items">
          <PaginationItem v-if="item.type === 'page'" :key="index" :value="item.value" as-child>
            <Button class="w-10 h-10 p-0" :variant="item.value === page ? 'default' : 'outline'">
              {{ item.value }}
            </Button>
          </PaginationItem>
          <PaginationEllipsis v-else :key="item.type" :index="index" />
        </template>
      </div>

      <div>
        <PaginationNext />
        <PaginationLast />
      </div>
    </PaginationContent>
  </Pagination>
</template>
