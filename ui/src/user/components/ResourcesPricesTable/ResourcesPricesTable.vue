<script lang="ts" setup>
import { computed, h } from 'vue'
import Badge from '@/core/components/ui/badge/Badge.vue'

import type { Price, ResourceProvider } from '@/user/interfaces/resourcePrices'

import {
  ClientSideRowModelModule,
  ColumnApiModule,
  ModuleRegistry,
  NumberEditorModule,
  TextEditorModule,
  ValidationModule,
  type ColDef,
  type ICellRendererParams,
  type IHeaderParams,
  type ValueFormatterParams,
  type ValueSetterParams,
} from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue3'
import ServiceProviderDeleteDialog from './ResourceDeleteDialog.vue'
import ServiceProviderFormDialog from './ResourceFormDialog.vue'
import { PlusIcon } from 'lucide-vue-next'

ModuleRegistry.registerModules([
  ClientSideRowModelModule,
  TextEditorModule,
  NumberEditorModule,
  ValidationModule,
  ColumnApiModule,
])

const props = withDefaults(
  defineProps<{
    resources: ResourceProvider[]
  }>(),
  {
    resources: () => [],
  },
)

const DeleteButton = {
  setup({ params }: { params: ICellRendererParams<ResourceProvider> }) {
    return () => h(ServiceProviderDeleteDialog, { serviceProviderId: params.data?.id || '' })
  },
}

const FormModal = {
  setup({ params }: { params: IHeaderParams<string, ResourceProvider[]> }) {
    console.log(params)
    return () =>
      h(
        ServiceProviderFormDialog,
        {
          prices: params.context?.[0]?.refPrices,
          providerID: params.context?.[0]?.provider,
          userID: params.context?.[0]?.user,
        },
        () => [h(PlusIcon), 'Add Row'],
      )
  },
}

const ResourceCell = {
  setup({ params }: { params: ICellRendererParams<ResourceProvider> }) {
    return () =>
      h('p', { class: 'flex gap-2 items-center' }, [
        params.value,
        h(
          Badge,
          {},
          {
            default: () => params.data?.type,
          },
        ),
      ])
  },
}

function generatePriceKey(prices: Price[]) {
  // Sort by minPersons to ensure order consistency
  return prices
    .slice()
    .sort((a, b) => a.minPersons - b.minPersons)
    .map((p) => `${p.minPersons}-${p.maxPersons}`)
    .join('|')
}

function groupResourcesByPrices(services: ResourceProvider[]) {
  const groups: Record<string, ResourceProvider[]> = {}

  for (const service of services) {
    const key = generatePriceKey(service.refPrices)
    if (!groups[key]) {
      groups[key] = []
    }
    groups[key].push(service)
  }

  return Object.values(groups)
}

const resourcesGrouped = computed(() => groupResourcesByPrices(props.resources))

function generateColumns(service: ResourceProvider) {
  const pricesColumns: ColDef[] = service.refPrices.map((price, index) => {
    return {
      field: `refPrices.${index}.price`,
      flex: 1,
      editable: true,
      minWidth: 100,
      headerName: `${price.minPersons} - ${price.maxPersons} Pax`,
      valueFormatter: (params: ValueFormatterParams<ResourceProvider>) => {
        const indexPrice = params.data?.refPrices[index]
        return `${params.value} ${indexPrice?.currency}`
      },
      valueSetter: (params: ValueSetterParams<ResourceProvider>) => {
        params.data.refPrices[index].price = params.newValue
        return true
      },
    }
  })

  const columns: ColDef[] = [
    {
      width: 300,
      field: 'expand.resource.name.en',
      lockPinned: true,
      editable: true,
      headerName: 'Resource',
      cellRenderer: ResourceCell,
      headerComponent: FormModal,
    },
    ...pricesColumns,
    {
      field: 'actions',
      pinned: 'right',
      cellRenderer: DeleteButton,
    },
  ]
  return columns
}
</script>
<template>
  <AgGridVue
    v-for="(group, index) in resourcesGrouped"
    :key="index"
    :context="group"
    :rowData="group"
    :columnDefs="generateColumns(group[0])"
    domLayout="autoHeight"
  ></AgGridVue>
</template>
