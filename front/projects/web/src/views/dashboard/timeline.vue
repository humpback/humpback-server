<script lang="ts" setup>
import * as echarts from "echarts/core"
import { GridComponent, GridComponentOption, LegendComponent, LegendComponentOption, TooltipComponent, TooltipComponentOption } from "echarts/components"
import { LineChart, LineSeriesOption } from "echarts/charts"
import { CanvasRenderer } from "echarts/renderers"
import { UniversalTransition } from "echarts/features"

echarts.use([GridComponent, LegendComponent, TooltipComponent, LineChart, CanvasRenderer, UniversalTransition])

type EChartsOption = echarts.ComposeOption<GridComponentOption | LineSeriesOption | TooltipComponentOption | LegendComponentOption>

const { t } = useI18n()

let incrementChart: echarts.ECharts
const incrementRef = useTemplateRef<HTMLDivElement>("incrementRef")

let incrementOptions = ref<EChartsOption | any>({
  tooltip: {
    trigger: "axis"
  },
  grid: {
    left: "2%",
    right: "2%",
    bottom: "4%",
    top: "8%",
    containLabel: true
  },
  xAxis: {
    type: "category",
    boundaryGap: false,
    data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
  },
  yAxis: {
    type: "value"
  },
  series: [
    {
      data: [820, 932, 901, 934, 1290, 1330, 1320],
      type: "line",
      lineStyle: {
        color: "var(--el-color-primary)",
        width: 4
      },
      areaStyle: {
        color: "#e2ebf0"
      }
    }
  ]
})

function resize() {
  incrementChart?.resize()
}

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
  <div>
    <el-row :gutter="20">
      <el-col :span="18">
        <v-card>
          <div ref="incrementRef" class="chart" />
        </v-card>
      </el-col>
      <el-col :span="6">
        <v-card></v-card>
      </el-col>
    </el-row>
  </div>
</template>

<style lang="scss" scoped>
.chart {
  height: 360px;
  width: 100%;
}
</style>
