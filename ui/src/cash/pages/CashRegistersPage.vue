<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import * as z from 'zod'

import { ArrowUpRight } from 'lucide-vue-next'
import Badge from '@/core/components/ui/badge/Badge.vue'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableHeader,
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
import { Button, buttonVariants } from '@/core/components/ui/button'
import { cn } from '@/core/lib/utils'
import { FormField, FormItem, FormLabel, FormMessage } from '@/core/components/ui/form'
import { toast } from 'vue-sonner'
import { Input } from '@/core/components/ui/input'

const formSchema = toTypedSchema(
  z.object({
    email: z.string({
      required_error: 'Please select a date to display.',
    }),
  }),
)

const { handleSubmit } = useForm({
  validationSchema: formSchema,
})

const onSubmit = handleSubmit(() => {
  toast('Event has been created', {
    description: 'Sunday, December 03, 2023 at 9:00 AM',
    action: {
      label: 'Undo',
      onClick: () => console.log('Undo'),
    },
    duration: Infinity,
  })
})
</script>
<template>
  <section class="flex flex-wrap gap-4">
    <Button class="h-auto grid gap-2" variant="outline">
      <div class="flex items-center gap-2">
        <img
          src="https://images.unsplash.com/photo-1554629947-334ff61d85dc?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=387&q=80"
          alt="Tour Image"
          class="w-6 h-6 rounded-full"
        />
        <span class="font-semibold">Roberto Alonso</span>
      </div>
      <div class="flex gap-4">
        <Badge variant="green">1 Open</Badge>
        <Badge variant="red">23 closed</Badge>
      </div>
    </Button>
  </section>

  <form class="flex gap-4" @submit="onSubmit">
    <FormField v-slot="{ componentField }" name="email">
      <FormItem>
        <FormLabel>From</FormLabel>
        <Input type="date" v-bind="componentField" />
        <FormMessage />
      </FormItem>
    </FormField>

    <FormField v-slot="{ componentField }" name="email">
      <FormItem>
        <FormLabel>To</FormLabel>
        <Input type="date" v-bind="componentField" />
        <FormMessage />
      </FormItem>
    </FormField>
    <Button type="submit" class="self-end"> Filter </Button>
  </form>

  <section class="w-full h-full">
    <Table class="min-w-[800px] w-full">
      <TableCaption>Open and Closed Cash Movements</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Nº</TableHead>
          <TableHead>Open amount</TableHead>
          <TableHead>Closed amount</TableHead>
          <TableHead>Opened by</TableHead>
          <TableHead>Closed by</TableHead>
          <TableHead>Status</TableHead>
          <TableHead>Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="index in 4" :key="index">
          <TableCell>1</TableCell>
          <TableCell>USD 4.250</TableCell>
          <TableCell>USD 0</TableCell>
          <TableCell>
            <div class="flex gap-2 items-center">
              <img
                src="https://images.unsplash.com/photo-1554629947-334ff61d85dc?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=387&q=80"
                alt="Tour Image"
                class="w-6 h-6 rounded-full"
              />
              <div>
                <span class="text-sm text-black/60">Roberto Alonso</span>
                <p class="text-sm text-black/60">20 Jun. 2022 11:00</p>
              </div>
            </div>
          </TableCell>
          <TableCell>
            <div class="flex gap-2 items-center">
              <img
                src="https://images.unsplash.com/photo-1554629947-334ff61d85dc?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=387&q=80"
                alt="Tour Image"
                class="w-6 h-6 rounded-full"
              />
              <div>
                <span class="text-sm text-black/60">Roberto Alonso</span>
                <p class="text-sm text-black/60">20 Jun. 2022 11:12</p>
              </div>
            </div>
          </TableCell>
          <TableCell>
            <Badge variant="green">Open</Badge>
          </TableCell>
          <TableCell>
            <div class="grid gap-2">
              <Button variant="outline" size="sm"> <ArrowUpRight class="w-4 h-4" /> Close </Button>
              <RouterLink :to="`/cash-registers/${1}/movements`" size="sm" :class="cn(buttonVariants())">
                <ArrowUpRight class="w-4 h-4" /> Details
              </RouterLink>
            </div>
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
  </section>
</template>
