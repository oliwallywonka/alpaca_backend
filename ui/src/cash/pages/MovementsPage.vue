<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import * as z from 'zod'
import { toast } from 'vue-sonner'

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
import Button from '@/core/components/ui/button/Button.vue'
import { Card } from '@/core/components/ui/card'
import CardContent from '@/core/components/ui/card/CardContent.vue'
import { FormField, FormItem, FormLabel, FormMessage } from '@/core/components/ui/form'
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
  <section class="grid grid-cols-2 gap-2">
    <Card>
      <CardContent class="font-semibold">
        <span>Dolar Flow (USD)</span>
        <p>Incomes: <Badge variant="green">14.250</Badge></p>
        <p>Outcomes: <Badge variant="red">14.250</Badge></p>
        <p>Balance: <Badge variant="green">0.00</Badge></p>
      </CardContent>
    </Card>
    <Card>
      <CardContent class="font-semibold">
        <span>Bolivian Flow (Bs)</span>
        <p>Incomes: <Badge variant="green">14.250</Badge></p>
        <p>Outcomes: <Badge variant="red">14.250</Badge></p>
        <p>Balance: <Badge variant="green">0.00</Badge></p>
      </CardContent>
    </Card>
  </section>

  <section class="flex gap-2">
    <Button>+ Income</Button>
    <Button>- Outcome</Button>
    <Button>History</Button>
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
    <Button type="submit" class="self-end"> Search </Button>
  </form>

  <section>
    <Table class="min-w-[800px] w-full">
      <TableCaption>Cash Movements</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Date</TableHead>
          <TableHead>Operation type</TableHead>
          <TableHead>Gloss</TableHead>
          <TableHead>Quantity</TableHead>
          <TableHead>Currency</TableHead>
          <TableHead>Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="index in 4" :key="index">
          <TableCell>20 Jun. 2022 11:00</TableCell>
          <TableCell><Badge variant="green">Outcome</Badge></TableCell>
          <TableCell>First Cash Outcome</TableCell>
          <TableCell>250</TableCell>
          <TableCell> <Badge variant="green">Bs</Badge></TableCell>
          <TableCell class="grid gap-2">
            <Button variant="outline" size="sm" class="w-24"> Delete </Button>
            <Button variant="outline" size="sm" class="w-24"> Edit </Button>
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
