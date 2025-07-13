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
  TextEditorModule,
  ValidationModule,
  ClientSideRowModelApiModule,
  HighlightChangesModule,
  type GridOptions,
  type IAggFuncParams,
  type ICellRendererParams,
} from 'ag-grid-community'
import { RowGroupingModule } from 'ag-grid-enterprise'
import { h } from 'vue'
import Badge from '@/core/components/ui/badge/Badge.vue'
import type { TourResource } from '@/tour/interfaces/tourResources'
import type { ResourceProvider } from '@/providerResource/ResourceProvider'
import DeleteExpenseDialog from './DeleteExpenseDialog.vue'

ModuleRegistry.registerModules([
  ClientSideRowModelModule,
  NumberEditorModule,
  TextEditorModule,
  ColumnApiModule,
  HighlightChangesModule,
  CellStyleModule,
  RowGroupingModule,
  ClientSideRowModelApiModule,
  RowAutoHeightModule,
  PinnedRowModule,
  ColumnAutoSizeModule,
  ValidationModule,
])

const props = defineProps<{
  resources: TourResource[]
}>()

const emit = defineEmits<{
  (e: 'change:tourResource', tourResource: Partial<TourResource>): void
}>()

const DescriptionCell = {
  setup({ params }: { params: ICellRendererParams<TourResource, ResourceProvider> }) {
    if (params.node.level === -1 && params.node.footer) {
      return () => h('span', { class: 'font-bold' }, 'Total')
    }
    return () =>
      h('p', { class: 'flex gap-2 items-center' }, [
        params.value?.expand.resource.name.en,
        h(Badge, {}, () => [
          params.value?.expand.provider?.fullName || params.value?.expand.user?.name,
        ]),
      ])
  },
}

const CurrencyCell = {
  setup({ params }: { params: ICellRendererParams }) {
    return () => h(Badge, () => [params.value])
  },
}

const ActionsCell = {
  setup({ params }: { params: ICellRendererParams }) {
    if (params.node.level === -1 && params.node.footer) {
      return () => {}
    }
    return () => h(DeleteExpenseDialog, { tourResourceID: params.data?.id })
  },
}

const sumHelper = (params: IAggFuncParams) => {
  const sum = params.values.reduce((acc, val) => acc + val, 0)
  return Math.round(sum * 100) / 100
}

const gridOptions: GridOptions<TourResource> = {
  defaultColDef: {
    minWidth: 90,
    enableCellChangeFlash: true,
    cellStyle: (params) => {
      return {
        backgroundColor: params.colDef.editable ? '#e2e8f0' : '#f9fafb',
      }
    },
  },
  onCellValueChanged: (params) => {
    const changedData = [params.data]
    params.api.applyTransaction({ update: changedData })
  },
  autoSizeStrategy: {
    type: 'fitCellContents',
  },
  columnDefs: [
    {
      headerName: 'Description',
      field: 'expand.resourceProvider',
      pinned: 'left',
      cellRenderer: DescriptionCell,
    },
    {
      headerName: 'Original Cost',
      field: 'originalCost',
      editable: true,
      onCellValueChanged: (params) => {
        emit('change:tourResource', {
          id: params.data?.id,
          originalCost: params.newValue,
        })
      },
      aggFunc: sumHelper,
    },
    {
      field: 'currency',
      cellRenderer: CurrencyCell,
    },
    {
      field: 'dollarChangeRate',
      valueFormatter: (params) => `${params.value}`,
      editable: true,
      onCellValueChanged: (params) => {
        emit('change:tourResource', {
          id: params.data?.id,
          dollarChangeRate: params.newValue,
        })
      },
      aggFunc: () => '',
    },
    {
      colId: 'COSTSUS',
      headerName: 'Cost SUS',
      valueGetter: (params) => {
        const costSUS = (params.data?.originalCost || 0) / (params.data?.dollarChangeRate || 0)
        return Math.round(costSUS * 100) / 100
      },
      aggFunc: sumHelper,
    },
    {
      field: 'quantity',
      editable: true,
      onCellValueChanged: (params) => {
        emit('change:tourResource', {
          id: params.data?.id,
          quantity: params.newValue,
        })
      },
      aggFunc: sumHelper,
    },
    {
      colId: 'SUBTOTALCOST',
      headerName: 'SubTotal Cost',
      valueGetter: (params) => {
        const costSUS = params.getValue('COSTSUS') ?? 1
        const quantity = params.data?.quantity || 1
        return Math.round(costSUS * quantity * 100) / 100
      },
    },
    /* {
      headerName: 'PP Cost',
      valueGetter: (params) => {
        // PPCOST = COSTSUS / MAX_PERSONS
        const costSUS = params.getValue('COSTSUS') ?? 1
        const maxPersons = params.data?.expand?.tourVariant?.maxPersons || 1
        return Math.round((costSUS / maxPersons) * 100) / 100
      },
      aggFunc: 'sum',
    }, */
    {
      field: 'profitPercent',
      valueFormatter: (params) => `${params.value}%`,
      aggFunc: () => '',
      editable: true,
      onCellValueChanged: (params) => {
        emit('change:tourResource', {
          id: params.data?.id,
          profitPercent: params.newValue,
        })
      },
    },
    {
      colId: 'PROFIT',
      headerName: 'Profit',
      valueGetter: (params) => {
        const subTotalCost = params.getValue('SUBTOTALCOST') ?? 1
        const profitPercent = params.data?.profitPercent || 0
        const profit = Math.round((profitPercent / 100) * subTotalCost * 100) / 100
        return Math.round(profit * 100) / 100
      },
      aggFunc: sumHelper,
    },
    /* {
      headerName: 'PP Profit',
      valueGetter: (params) => {
        const costSUS = params.getValue('COSTSUS') ?? 1
        const maxPersons = params.data?.expand?.tourVariant?.maxPersons || 1
        const profitPercent = params.data?.profitPercent || 0
        const profit = Math.round((costSUS / maxPersons) * profitPercent * 100) / 100
        return profit
      },
    }, */
    /* {
      headerName: 'SubTotal Profit',
      valueGetter: (params) => {
        const costSUS = params.getValue('COSTSUS') ?? 1
        const profit = ((params.data?.profitPercent || 0) * costSUS) / 100
        return profit
      }
    }, */
    /* {
      headerName: 'Price',
      valueGetter: (params) => {
        const subTotalCost = params.getValue('SUBTOTALCOST') ?? 1
        const profit = params.getValue('PROFIT') ?? 0
        return Math.round((subTotalCost + profit) * 100) / 100
      },
      aggFunc: 'sum',
    }, */
    {
      headerName: 'SubTotal Price',
      valueGetter: (params) => {
        const subTotalCost = params.getValue('SUBTOTALCOST') ?? 1
        const profit = params.getValue('PROFIT') ?? 0
        return Math.round((subTotalCost + profit) * 100) / 100
      },
      aggFunc: sumHelper,
      pinned: 'right',
    },
    {
      headerName: 'Actions',
      pinned: 'right',
      cellRenderer: ActionsCell,
    },
  ],
  // TODO: Total row are not updating when editing fix it
  grandTotalRow: 'pinnedBottom',
  enableRowPinning: true,
  domLayout: 'autoHeight',
}
</script>

<template>
  <AgGridVue :gridOptions="gridOptions" :rowData="props.resources" />
</template>
