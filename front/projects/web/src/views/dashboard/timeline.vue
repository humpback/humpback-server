<script lang="ts" setup>
import * as echarts from "echarts/core"
import {
  GridComponent,
  GridComponentOption,
  ToolboxComponent,
  ToolboxComponentOption,
  LegendComponent,
  LegendComponentOption,
  TooltipComponent,
  TooltipComponentOption
} from "echarts/components"
import { LineChart, BarChart, LineSeriesOption, BarSeriesOption } from "echarts/charts"
import { CanvasRenderer } from "echarts/renderers"
import { UniversalTransition } from "echarts/features"
import { map } from "lodash-es"

echarts.use([GridComponent, ToolboxComponent, LegendComponent, BarChart, TooltipComponent, LineChart, CanvasRenderer, UniversalTransition])

type EChartsOption = echarts.ComposeOption<
  GridComponentOption | ToolboxComponentOption | BarSeriesOption | LineSeriesOption | TooltipComponentOption | LegendComponentOption
>

const { t, locale } = useI18n()

let incrementChart: echarts.ECharts
const incrementRef = useTemplateRef<HTMLDivElement>("incrementRef")

let incrementOptions = computed<EChartsOption>(() => ({
  tooltip: {
    trigger: "axis",
    formatter: params => {
      let content = ""
      params = Array.isArray(params) ? params : [params]
      content += `${params[0].axisValue}<br/>`
      map(params, item => {
        switch (item.seriesIndex) {
          case 0:
            content += `${item.marker}${item.seriesName}：${item.data}<br/>`
            break
          case 1:
            content += `${item.marker}${item.seriesName}：${item.data}<br/>`
            break
          case 2:
            content += `${item.marker}${item.seriesName}：${item.data}<br/>`
            break
        }
      })
      return content
    }
  },
  legend: {
    right: "50%",
    selected: {
      "Group": true,
      "Service": true,
      "Deploy": true
    }
  },
  grid: {
    left: "1%",
    right: "1%",
    bottom: "3%",
    containLabel: true
  },
  toolbox: {
    show: true,
    feature: {
      dataView: { show: true, readOnly: false },
      magicType: { show: true, type: ["line", "bar", "stack"] },
      restore: { show: true }
    },
    right: "1%"
  },
  calculable: true,
  xAxis: [
    {
      type: "category",
      data: ["3-1", "3-2", "3-3", "3-4", "4-5", "5-5", "6-5", "7-5", "8-5"]
    }
  ],
  yAxis: [{ type: "value" }],
  series: [
    {
      name: t("label.group"),
      type: "bar",
      seriesLayoutBy: "column",
      data: [1, 2, 3, 4, 2, 3, 1, 9, 0]
    },
    {
      name: t("label.service"),
      type: "bar",
      seriesLayoutBy: "column",
      data: [1, 5, 8, 4, 10, 3, 1, 0, 0]
    },
    {
      name: t("label.deploy"),
      type: "line",
      seriesLayoutBy: "column",
      data: [5, 6, 7, 4, 20, 25, 11, 19, 16]
    }
  ]
}))

function resize() {
  incrementChart?.resize()
}

watch(locale, () => {
  incrementChart?.setOption(incrementOptions.value)
})

onMounted(() => {
  incrementChart = echarts.init(incrementRef.value)
  incrementChart?.setOption(incrementOptions.value)
  window.addEventListener("resize", resize)
})

onBeforeUnmount(() => {
  incrementChart?.dispose()
  window.removeEventListener("resize", resize)
})
</script>

<template>
  <v-card>
    <div class="title">{{ t("header.newDataOf30Days") }}</div>
    <div ref="incrementRef" class="chart" />
  </v-card>
</template>

<style lang="scss" scoped>
.title {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 20px;
  padding-left: 4px;
}

.chart {
  height: 360px;
  width: 100%;
}
</style>
