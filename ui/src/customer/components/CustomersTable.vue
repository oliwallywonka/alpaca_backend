<script setup lang="ts">
import { h } from 'vue'

import {
  ClientSideRowModelModule,
  ModuleRegistry,
  RowSelectionModule,
  ValidationModule,
  type GridOptions,
  type ICellRendererParams,
} from 'ag-grid-community'
import type { Customer } from '../interfaces/customer'
import { AgGridVue } from 'ag-grid-vue3'
import { parseDate } from '@/core/lib/date'
import type { ContactField } from '@/core/interfaces/fields'
import Badge from '@/core/components/ui/badge/Badge.vue'
import CustomerFormDIalog from '../components/CustomerFormDIalog.vue'

const props = defineProps<{
  customers: Customer[]
}>()

ModuleRegistry.registerModules([ClientSideRowModelModule, RowSelectionModule, ValidationModule])

const contactsCell = {
  setup({ params }: { params: ICellRendererParams<Customer, ContactField[]> }) {
    return () =>
      h('div', { class: 'flex flex-row gap-1' }, [
        params.value?.map((contact) => h(Badge, {}, [contact.type, ': ', contact.value])),
      ])
  },
}

const actionsCell = {
  setup({ params }: { params: ICellRendererParams }) {
    return () => h(CustomerFormDIalog, { customerID: params.data?.id })
  },
}

const gridOptions: GridOptions<Customer> = {
  domLayout: 'autoHeight',
  defaultColDef: {
    flex: 1,
    minWidth: 100,
  },
  columnDefs: [
    {
      headerName: 'Full Name',
      valueFormatter: (params) => {
        return `${params.data?.title} ${params.data?.firstName} ${params.data?.middleName} ${params.data?.lastName}`
      },
    },
    {
      field: 'dateOfBirth',
      headerName: 'Date of Birth',
      valueFormatter: (params) => parseDate(params.value),
    },
    {
      field: 'contacts',
      cellRenderer: contactsCell,
    },
    {
      headerName: 'Actions',
      cellRenderer: actionsCell,
    },
  ],
}
</script>

<template>
  <AgGridVue :gridOptions="gridOptions" :rowData="props.customers" />
</template>
