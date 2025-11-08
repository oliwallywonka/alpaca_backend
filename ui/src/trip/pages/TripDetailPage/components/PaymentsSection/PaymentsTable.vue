<script setup lang="ts">
import { Badge } from '@/core/components/ui/badge'
import PDFVoucher from '@/trip/components/PDFVoucher.vue'
import {
  ClientSideRowModelModule,
  ColumnAutoSizeModule,
  ModuleRegistry,
  ValidationModule,
  type GridOptions,
  type ICellRendererParams,
} from 'ag-grid-community'
import type { Payment } from '@/trip/interfaces/payment'
import { AgGridVue } from 'ag-grid-vue3'
import { h } from 'vue'
import { dateToSuggar } from '@/core/lib/date'

ModuleRegistry.registerModules([ClientSideRowModelModule, ColumnAutoSizeModule, ValidationModule])

const props = defineProps<{
  payments: Payment[]
}>()

const amountCell = {
  setup({ params }: { params: ICellRendererParams<Payment> }) {
    return () =>
      h(
        Badge,
        {
          variant: 'green',
        },
        () => `${params.data?.amount} ${params.data?.currency.toLocaleLowerCase() || ''}`,
      )
  },
}

const methodCell = {
  setup({ params }: { params: ICellRendererParams<Payment> }) {
    return () => h(Badge, {}, () => params.value)
  },
}

const actionCell = {
  setup({ params }: { params: ICellRendererParams<Payment> }) {
    return () => h(PDFVoucher)
  },
}

const gridOptions: GridOptions<Payment> = {
  domLayout: 'autoHeight',
  autoSizeStrategy: {
    type: 'fitCellContents',
  },
  defaultColDef: {
    minWidth: 120,
  },
  columnDefs: [
    {
      field: 'created',
      headerName: 'Date',
      valueFormatter: ({ value }) => dateToSuggar(value, true),
    },
    {
      field: 'expand.registeredBy.name',
      headerName: 'Registered By',
    },
    {
      headerName: 'Customer',
      valueFormatter: ({ data }) =>
        `${data?.expand.customer.firstName} ${data?.expand.customer.lastName}`,
    },
    {
      headerName: 'Amount',
      cellRenderer: amountCell,
    },
    {
      field: 'method',
      cellRenderer: methodCell,
    },
    {
      headerName: 'Actions',
      pinned: 'right',
      cellRenderer: actionCell,
    },
  ],
}
</script>

<template>
  <AgGridVue :gridOptions="gridOptions" :rowData="props.payments" />
</template>
