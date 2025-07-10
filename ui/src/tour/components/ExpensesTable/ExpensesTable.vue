<script setup lang="ts">
import { AgGridVue } from 'ag-grid-vue3'
import {
  CellStyleModule,
  ClientSideRowModelModule,
  ColumnApiModule,
  ColumnAutoSizeModule,
  ModuleRegistry,
  NumberEditorModule,
  PinnedRowModule,
  RowAutoHeightModule,
  ValidationModule,
  type GridOptions,
  type ICellRendererParams,
} from 'ag-grid-community'
import { RowGroupingModule } from 'ag-grid-enterprise'
import { h } from 'vue'
import Badge from '@/core/components/ui/badge/Badge.vue'

ModuleRegistry.registerModules([
  ClientSideRowModelModule,
  NumberEditorModule,
  ColumnApiModule,
  CellStyleModule,
  RowGroupingModule,
  RowAutoHeightModule,
  PinnedRowModule,
  ColumnAutoSizeModule,
  ValidationModule,
])

const DescriptionCell = {
  setup({ params }: { params: ICellRendererParams }) {
    if (params.node.level === -1 && params.node.footer) {
      return () => h('span', { class: 'font-bold' }, 'Total')
    }
    return () =>
      h('p', [params.value, h(Badge, { variant: 'yellow' }, () => ['Service'])])
  },
}

const CurrencyCell = {
  setup({ params }: { params: ICellRendererParams }) {
    return () => h(Badge, () => [params.value])
  },
}

const gridOptions: GridOptions = {
  defaultColDef: {
    width: 90,
  },
  autoSizeStrategy: {
    type: 'fitCellContents',
  },
  columnDefs: [
    {
      field: 'description',
      pinned: 'left',
      cellRenderer: DescriptionCell,
    },
    {
      headerName: 'Original Cost',
    },
    {
      field: 'currency',
      cellRenderer: CurrencyCell,
    },
    {
      field: 'changeRate',
    },
    {
      field: 'cost1',
      valueFormatter: (params) => `$${params.value}`,
      headerName: 'Cost $',
      aggFunc: 'sum',
    },
    {
      field: 'quantity',
      aggFunc: 'sum',
      editable: true,
    },
    {
      headerName: 'PP Cost',
    },
    {
      field: 'profitPercent',
      valueFormatter: (params) => `${params.value}%`,
      aggFunc: () => '',
    },
    {
      headerName: 'PP Profit',
    },
    {
      headerName: 'SubTotal Profit',
    },
    {
      headerName: 'PP Price',
    },
    {
      headerName: 'SubTotal Price',
      pinned: 'right',
    },
    {
      headerName: 'Actions',
      pinned: 'right',
    },
  ],
  rowData: [
    {
      description: 'Transfer Aereopuerto IN / OUT',
      userProvider: 'Fernando camacho',
      currency: 'SUS',
      changeRate: 1.0,
      quantity: 1,
      cost1: 20.0,
      profitPercent: 30,
    },
  ],
  grandTotalRow: 'pinnedBottom',
  enableRowPinning: true,
  domLayout: 'autoHeight',
}
</script>

<template>
  <AgGridVue :gridOptions="gridOptions" />

</template>
