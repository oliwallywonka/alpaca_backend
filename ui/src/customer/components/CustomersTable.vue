<script setup lang="ts">
import { computed, h } from 'vue'

import {
  ClientSideRowModelModule,
  ModuleRegistry,
  RowSelectionModule,
  ValidationModule,
  type GridOptions,
  type ICellRendererParams,
  type SelectionChangedEvent,
} from 'ag-grid-community'
import type { Customer } from '../interfaces/customer'
import { AgGridVue } from 'ag-grid-vue3'
import { parseDate } from '@/core/lib/date'
import type { ContactField } from '@/core/interfaces/fields'
import Badge from '@/core/components/ui/badge/Badge.vue'
import CustomerFormDIalog from '../components/CustomerFormDIalog.vue'

ModuleRegistry.registerModules([ClientSideRowModelModule, RowSelectionModule, ValidationModule])

const props = defineProps<{
  customers: Customer[]
  showSelect?: boolean
}>()

const emit = defineEmits<{
  (e: 'select:customer', customer: Customer | undefined): void
}>()

const contactsCell = {
  setup({ params }: { params: ICellRendererParams<Customer, ContactField[]> }) {
    return () =>
      h('div', { class: 'flex flex-row gap-1' }, () => [
        params.value?.map((contact) => h(Badge, {}, [contact.type, ': ', contact.value])),
      ])
  },
}

const actionsCell = {
  setup({ params }: { params: ICellRendererParams }) {
    return () => h(CustomerFormDIalog, { customerID: params.data?.id })
  },
}

const selectOptions = computed<Partial<GridOptions>>(() => {
  if (!props.showSelect) return {}
  return {
    rowSelection: {
      mode: 'singleRow',
      checkboxes: true,
    },
    selectionColumnDef: {
      pinned: 'left',
    },
    onSelectionChanged: (event: SelectionChangedEvent<Customer>) => {
      if (!event.selectedNodes || event.selectedNodes?.length === 0) {
        emit('select:customer', undefined)
        return
      }
      emit('select:customer', event.selectedNodes[0].data)
    },
  }
})

const gridOptions: GridOptions<Customer> = {
  domLayout: 'autoHeight',
  defaultColDef: {
    flex: 1,
    minWidth: 100,
  },
  ...selectOptions.value,
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
