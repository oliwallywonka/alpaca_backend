<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch, defineExpose, defineProps, useId } from 'vue'
import {
  createChart,
  type IChartApi,
  type ISeriesApi,
  LineSeries,
  AreaSeries,
  BarSeries,
  CandlestickSeries,
  HistogramSeries,
  BaselineSeries,
  type LineData,
  type AreaData,
  type BarData,
  type CandlestickData,
  type HistogramData,
  type BaselineData,
  type ChartOptions,
  type DeepPartial,
  type SeriesType,
} from 'lightweight-charts'

export type ChartType = 'line' | 'area' | 'histogram'

type Props = {
  type?: ChartType
  data: Array<LineData | AreaData | HistogramData>
  autosize?: boolean
  chartOptions?: DeepPartial<ChartOptions>
  seriesOptions?: object
  timeScaleOptions?: object
  priceScaleOptions?: object
}

const props = defineProps<Props>()

function getChartSeriesDefinition(type: ChartType) {
  switch (type.toLowerCase()) {
    case 'line':
      return LineSeries
    case 'area':
      return AreaSeries
    case 'bar':
      return BarSeries
    case 'candlestick':
      return CandlestickSeries
    case 'histogram':
      return HistogramSeries
    case 'baseline':
      return BaselineSeries
    default:
      return LineSeries
  }
}

let chart: IChartApi | null = null
let series: ISeriesApi<SeriesType> | null = null

const chartId = useId()

const chartContainer = ref<HTMLDivElement | null>(null)

const fitContent = () => {
  if (!chart) return
  chart.timeScale().fitContent()
}

const updateData = (
  newData: LineData | AreaData | BarData | CandlestickData | HistogramData | BaselineData,
) => {
  if (series) {
    series.update(newData)
  }
}

const getChart = () => chart

defineExpose({ fitContent, getChart, updateData })

const resizeHandler = () => {
  if (!chart || !chartContainer.value) return
  const dimensions = chartContainer.value.getBoundingClientRect()
  chart.resize(dimensions.width, dimensions.height)
}

const addSeries = (props: Props) => {
  const seriesDefinition = getChartSeriesDefinition(props.type || 'line')
  series = chart!.addSeries(seriesDefinition, props.seriesOptions!) as ISeriesApi<SeriesType>
  series.setData(props.data)
}

onMounted(() => {
  if (!chartContainer.value || !props.data) return
  chart = createChart(chartContainer.value, props.chartOptions)
  chart.applyOptions({
    layout: { attributionLogo: false, textColor: '#808080' },
    grid: { vertLines: { visible: false }, horzLines: { visible: false } },
    rightPriceScale: {
      borderColor: '#e5e7eb',
    },
    timeScale: {
      borderColor: '#e5e7eb',
    },
  })
  addSeries(props)

  if (props.priceScaleOptions) {
    chart.priceScale(chartId).applyOptions(props.priceScaleOptions)
  }

  if (props.timeScaleOptions) {
    chart.timeScale().applyOptions(props.timeScaleOptions)
  }

  chart.timeScale().fitContent()

  if (props.autosize) {
    window.addEventListener('resize', resizeHandler)
  }
})

onUnmounted(() => {
  if (chart) {
    chart.remove()
    chart = null
  }
  series = null
  window.removeEventListener('resize', resizeHandler)
})

/* watch(
  () => props.baseData,
  (newData) => {
    console.log(props.baseData)
    if (series) {
      series.setData(newData)
    }
  },
  {
    deep: true,
  },
) */

watch(
  () => props.autosize,
  (enabled) => {
    if (!enabled) {
      window.removeEventListener('resize', resizeHandler)
      return
    }
    window.addEventListener('resize', resizeHandler)
  },
)

watch(
  () => props.type,
  () => {
    if (series && chart) {
      chart.removeSeries(series)
    }
    addSeries(props)
  },
)

watch(
  () => props.chartOptions,
  (newOptions) => {
    if (chart && newOptions) {
      chart.applyOptions(newOptions)
    }
  },
)

watch(
  () => props.seriesOptions,
  (newOptions) => {
    if (series && newOptions) {
      series.applyOptions(newOptions)
    }
  },
)

watch(
  () => props.priceScaleOptions,
  (newOptions) => {
    if (chart && newOptions) {
      chart.priceScale(chartId).applyOptions(newOptions)
    }
  },
)

watch(
  () => props.timeScaleOptions,
  (newOptions) => {
    if (chart && newOptions) {
      chart.timeScale().applyOptions(newOptions)
    }
  },
)
</script>

<template>
  <div class="h-full w-full" ref="chartContainer"></div>
</template>
