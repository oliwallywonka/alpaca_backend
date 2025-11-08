<script setup lang="ts">
import { Badge } from '@/core/components/ui/badge'
import { Button } from '@/core/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/core/components/ui/dialog'
import { FormField, FormItem, FormMessage } from '@/core/components/ui/form'
import FormLabel from '@/core/components/ui/form/FormLabel.vue'
import { Select, SelectValue, SelectContent, SelectItem, SelectTrigger } from '@/core/components/ui/select'
import { Input } from '@/core/components/ui/input'

const props = defineProps<{
  itineraryResourceID: string
  resourcePaymentID?: string
}>()
</script>

<template>
  <Dialog>
    <DialogTrigger as-child>
      <Button variant="ghost"> 120.0 usd. </Button>
    </DialogTrigger>
    <DialogContent class="sm:max-w-2xl">
      <DialogHeader>
        <DialogTitle>
          {{ props.resourcePaymentID ? 'Edit' : 'New' }}
          Payment
        </DialogTitle>
        <DialogDescription>
          Fill out the form to add a new payment to cover this expense.
        </DialogDescription>
      </DialogHeader>
      <div class="grid grid-cols-2 gap-2">
        <div class="grid grid-cols-[auto_1fr] gap-y-2 gap-x-4 items-start place-content-start">
          <Badge class="w-full">Resource: </Badge>
          <span class="-mt-1"
            >Transfer in/out Airport El Alto (day) <Badge>Cost: 2000.0 bol.</Badge></span
          >
          <Badge class="w-full">Provider: </Badge>
          <span class="-mt-1">Transporte Turistico Boliviano Emperador</span>
          <Badge class="w-full">Tour: </Badge>
          <span class="-mt-1"
            >Salar 2 Days <Badge>Persons Min: 1 Max: 4</Badge>
            <Badge>Price: 120.0 usd.</Badge></span
          >
        </div>
        <form class="px-2 bg-white grid gap-4 place-content-start">
          <div class="flex gap-2">
            <FormField name="amount">
              <FormItem>
                <FormLabel>Amount</FormLabel>
                <Input />
              </FormItem>
            </FormField>
            <FormField v-slot="{ componentField }" name="email">
              <FormItem>
                <FormLabel>Currency</FormLabel>

                <Select v-bind="componentField">
                  <SelectTrigger>
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="m@example.com"> BOL </SelectItem>
                    <SelectItem value="m@google.com"> USD </SelectItem>
                  </SelectContent>
                </Select>
                <FormMessage />
              </FormItem>
            </FormField>
            <FormField name="changeRate">
              <FormItem>
                <FormLabel>CR</FormLabel>
                <Input />
              </FormItem>
            </FormField>
          </div>
            <FormField name="changeRate">
              <FormItem>
                <FormLabel>Left Amount</FormLabel>
                <Input disabled value="120.0 usd." />
              </FormItem>
            </FormField>
        </form>
      </div>
      <DialogFooter>
        <Button type="submit"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
