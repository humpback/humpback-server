<script lang="ts" setup>
import * as echarts from "echarts/core"
import { GridComponent, GridComponentOption, TooltipComponent, TooltipComponentOption, LegendComponent, LegendComponentOption } from "echarts/components"
import { LineChart, LineSeriesOption } from "echarts/charts"
import { UniversalTransition } from "echarts/features"
import { CanvasRenderer } from "echarts/renderers"
import { SetWebTitle } from "@/utils"
import { refreshData } from "@/views/service-management/service/common.ts"
import { ServiceInfo } from "@/types"

echarts.use([GridComponent, LegendComponent, TooltipComponent, LineChart, CanvasRenderer, UniversalTransition])

type EChartsOption = echarts.ComposeOption<GridComponentOption | LineSeriesOption | TooltipComponentOption | LegendComponentOption>

const { t } = useI18n()
const route = useRoute()
const stateStore = useStateStore()

const groupId = ref(route.params.groupId as string)
const serviceId = ref(route.params.serviceId as string)
const serviceInfo = computed<ServiceInfo | undefined>(() => stateStore.getService(serviceId.value))

const isLoading = ref(false)

const cpuRef = useTemplateRef<any>("cpuRef")
const memoryRef = useTemplateRef<any>("memoryRef")
const networkRef = useTemplateRef<any>("networkRef")
const ioRef = useTemplateRef<any>("ioRef")

let cpuChart: echarts.ECharts
let memoryChart: echarts.ECharts
let networkChart: echarts.ECharts
let ioChart: echarts.ECharts

let cpuOptions = ref<EChartsOption>({
  tooltip: {
    trigger: "axis",
    valueFormatter: value => `${value}%`
  },
  grid: {
    left: "6%",
    right: "6%",
    bottom: "10%",
    containLabel: true
  },
  xAxis: [
    {
      type: "category",
      data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
    }
  ],
  yAxis: [{ type: "value", axisLabel: { formatter: "{value} %" } }],
  series: [
    {
      name: "Service",
      type: "line",
      data: [150, 230, 224, 218, 135, 147, 260]
    },
    {
      name: "Version",
      type: "line",
      data: [150, 230, 224, 218, 135, 147, 260]
    },
    {
      name: "DeployService",
      type: "line",
      data: [150, 230, 224, 218, 135, 147, 260]
    }
  ]
})

const memoryOptions = ref<EChartsOption>({
  tooltip: {
    trigger: "axis",
    valueFormatter: value => `${value} MB`
  },
  grid: {
    left: "6%",
    right: "6%",
    bottom: "10%",
    containLabel: true
  },
  xAxis: [
    {
      type: "category",
      data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
    }
  ],
  yAxis: [{ type: "value", axisLabel: { formatter: "{value} MB" } }],
  series: [
    {
      name: "Service",
      type: "line",
      data: [150, 230, 224, 218, 135, 147, 260]
    },
    {
      name: "Version",
      type: "line",
      data: [150, 230, 224, 218, 135, 147, 260]
    },
    {
      name: "DeployService",
      type: "line",
      data: [150, 230, 224, 218, 135, 147, 260]
    }
  ]
})

const networkOptions = ref<EChartsOption>({
  tooltip: {
    trigger: "axis",
    valueFormatter: value => `${value} MB`
  },
  grid: {
    left: "6%",
    right: "6%",
    bottom: "10%",
    containLabel: true
  },
  xAxis: [
    {
      type: "category",
      data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
    }
  ],
  yAxis: [
    { type: "value", axisLabel: { formatter: "{value} MB" }, position: "left" },
    { type: "value", axisLabel: { formatter: "{value} MB" }, position: "right" }
  ],
  series: [
    {
      name: "Service",
      type: "line",
      yAxisIndex: 0,
      data: [150, 230, 224, 218, 135, 147, 260]
    },
    {
      name: "Version",
      type: "line",
      yAxisIndex: 1,
      data: [150, 230, 224, 218, 135, 147, 260]
    },
    {
      name: "DeployService",
      type: "line",
      yAxisIndex: 0,
      data: [150, 230, 224, 218, 135, 147, 260]
    },
    {
      name: "DeployServiceaa",
      type: "line",
      yAxisIndex: 1,
      data: [150, 230, 224, 218, 135, 147, 300]
    }
  ]
})

const ioOptions = ref<EChartsOption>({
  tooltip: {
    trigger: "axis",
    valueFormatter: value => `${value} B`
  },
  grid: {
    left: "6%",
    right: "6%",
    bottom: "10%",
    containLabel: true
  },
  xAxis: [
    {
      type: "category",
      data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
    }
  ],
  yAxis: [
    { type: "value", axisLabel: { formatter: "{value} B" }, position: "left" },
    { type: "value", axisLabel: { formatter: "{value} B" }, position: "right" }
  ],
  series: [
    {
      name: "Service",
      type: "line",
      yAxisIndex: 0,
      data: [150, 230, 224, 218, 135, 147, 260]
    },
    {
      name: "Version",
      type: "line",
      yAxisIndex: 1,
      data: [150, 230, 224, 218, 135, 147, 260]
    },
    {
      name: "DeployService",
      type: "line",
      yAxisIndex: 0,
      data: [150, 230, 224, 218, 135, 147, 260]
    },
    {
      name: "DeployServiceaa",
      type: "line",
      yAxisIndex: 1,
      data: [150, 230, 224, 218, 135, 147, 300]
    }
  ]
})

function resize() {
  cpuChart?.resize()
  memoryChart?.resize()
  networkChart?.resize()
  ioChart?.resize()
}

async function search() {
  isLoading.value = true
  await refreshData(groupId.value, serviceId.value, "instances").finally(() => (isLoading.value = false))
}

function resetChartData() {
  cpuChart.setOption(cpuOptions.value)
  memoryChart.setOption(memoryOptions.value)
  networkChart.setOption(networkOptions.value)
  ioChart.setOption(ioOptions.value)
}

onMounted(async () => {
  await search()
  SetWebTitle(`${t("webTitle.serviceInfo")} - ${stateStore.getService()?.serviceName}`)

  cpuChart = echarts.init(cpuRef.value)
  memoryChart = echarts.init(memoryRef.value)
  networkChart = echarts.init(networkRef.value)
  ioChart = echarts.init(ioRef.value)

  resetChartData()

  window.addEventListener("resize", resize)
})

onUnmounted(() => {
  window.removeEventListener("resize", resize)
  cpuChart?.dispose()
  memoryChart?.dispose()
  networkChart?.dispose()
  ioChart?.dispose()
})
</script>

<template>
  <div class="d-flex gap-3">
    <div class="header-icon">
      <el-icon :size="18">
        <IconMdiPerformance />
      </el-icon>
    </div>
    <strong>
      <el-text size="large">{{ t("label.performance") }}</el-text>
    </strong>
  </div>

  <div v-loading="isLoading">
    <el-row :gutter="20">
      <el-col :md="12" :span="24" class="mt-5">
        <div class="chart-box">
          <div class="chart-header">
            <el-icon :size="18">
              <IconMdiPerformance />
            </el-icon>
            <el-text>{{ t("label.cpuUsage") }}</el-text>
          </div>
          <div ref="cpuRef" class="chart" />
        </div>
      </el-col>

      <el-col :md="12" :span="24" class="mt-5">
        <div class="chart-box">
          <div class="chart-header">
            <el-icon :size="18">
              <IconMdiPerformance />
            </el-icon>
            <el-text>{{ t("label.memoryUsage") }}</el-text>
          </div>
          <div ref="memoryRef" class="chart" />
        </div>
      </el-col>

      <el-col :md="12" :span="24" class="mt-5">
        <div class="chart-box">
          <div class="chart-header">
            <el-icon :size="18">
              <IconMdiPerformance />
            </el-icon>
            <el-text>{{ t("label.networkUsage") }}</el-text>
          </div>
          <div ref="networkRef" class="chart" />
        </div>
      </el-col>

      <el-col :md="12" :span="24" class="mt-5">
        <div class="chart-box">
          <div class="chart-header">
            <el-icon :size="18">
              <IconMdiPerformance />
            </el-icon>
            <el-text>{{ t("label.ioUsage") }}</el-text>
          </div>
          <div ref="ioRef" class="f chart" />
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<style lang="scss" scoped>
.header-icon {
  color: var(--el-color-primary);
  background-color: var(--el-color-primary-light-9);
  border-radius: 50%;
  padding: 8px;
  width: 18px;
  height: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.chart-box {
  border: 1px solid var(--el-border-color);
  border-radius: 8px;
  box-sizing: border-box;

  .chart-header {
    padding: 16px 16px 0 20px;
    display: flex;
    align-items: center;
    gap: 12px;

    .el-icon {
      color: var(--el-color-primary);
      background-color: var(--el-color-primary-light-9);
      border-radius: 50%;
      padding: 4px;
      width: 14px;
      height: 14px;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }

  .chart {
    height: 400px;
    width: 100%;
  }
}
</style>
