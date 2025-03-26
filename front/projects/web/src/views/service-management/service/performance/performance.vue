<script lang="ts" setup>
import * as echarts from "echarts/core"
import { GridComponent, GridComponentOption, LegendComponent, LegendComponentOption, TooltipComponent, TooltipComponentOption } from "echarts/components"
import { LineChart, LineSeriesOption } from "echarts/charts"
import { UniversalTransition } from "echarts/features"
import { CanvasRenderer } from "echarts/renderers"
import { SetWebTitle, TimestampToTime } from "@/utils"
import { refreshData } from "@/views/service-management/service/common.ts"
import { ContainerPerformance, ServiceInfo } from "@/types"
import { filter, find, findIndex, map } from "lodash-es"

echarts.use([GridComponent, LegendComponent, TooltipComponent, LineChart, CanvasRenderer, UniversalTransition])

type EChartsOption = echarts.ComposeOption<GridComponentOption | LineSeriesOption | TooltipComponentOption | LegendComponentOption>

const { t } = useI18n()
const route = useRoute()
const stateStore = useStateStore()

const groupId = ref(route.params.groupId as string)
const serviceId = ref(route.params.serviceId as string)
const serviceInfo = computed<ServiceInfo | undefined>(() => stateStore.getService(serviceId.value))

const isLoading = ref(false)
const timer = ref<any>(null)

const cpuRef = useTemplateRef<any>("cpuRef")
const memoryRef = useTemplateRef<any>("memoryRef")
const networkRef = useTemplateRef<any>("networkRef")
const ioRef = useTemplateRef<any>("ioRef")

let cpuChart: echarts.ECharts
let memoryChart: echarts.ECharts
let networkChart: echarts.ECharts
let ioChart: echarts.ECharts

let cpuOptions = ref<EChartsOption | any>({
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
      data: []
    }
  ],
  yAxis: [{ type: "value", axisLabel: { formatter: "{value} %" } }],
  series: []
})

const memoryOptions = ref<EChartsOption | any>({
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
      data: []
    }
  ],
  yAxis: [{ type: "value", axisLabel: { formatter: "{value} MB" } }],
  series: []
})

const networkOptions = ref<EChartsOption | any>({
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
  xAxis: [{ type: "category", data: [] }],
  yAxis: [
    {
      type: "value",
      axisLabel: { formatter: "{value} MB" },
      position: "left",
      name: "RX",
      nameTextStyle: { align: "right", fontWeight: "bold" }
    },
    {
      type: "value",
      axisLabel: { formatter: "{value} MB" },
      position: "right",
      name: "TX",
      nameTextStyle: { align: "left", fontWeight: "bold" }
    }
  ],
  series: []
})

const ioOptions = ref<EChartsOption | any>({
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
  xAxis: [{ type: "category", data: [] }],
  yAxis: [
    {
      type: "value",
      axisLabel: { formatter: "{value} B" },
      position: "left",
      name: "Read",
      nameTextStyle: { align: "right", fontWeight: "bold" }
    },
    {
      type: "value",
      axisLabel: { formatter: "{value} B" },
      position: "right",
      name: "Write",
      nameTextStyle: { align: "left", fontWeight: "bold" }
    }
  ],
  series: []
})

function resize() {
  cpuChart?.resize()
  memoryChart?.resize()
  networkChart?.resize()
  ioChart?.resize()
}

function parseCpuOption(container: ServiceContainerStatusInfo, info: ContainerPerformance) {
  cpuOptions.value.xAxis?.[0].data.push(TimestampToTime(info.stats?.statsAt, 2))
  const titleIndex = cpuOptions.value.xAxis?.[0].data.length || 0
  const containerIndex = findIndex(cpuOptions.value?.series || [], (x: any) => x.name === container.containerName)
  const data = containerIndex !== -1 ? cpuOptions.value?.series[containerIndex].data : []
  for (let i = data.length; i < titleIndex; i++) {
    data[i] = i + 1 === titleIndex ? info.stats?.cpuPercent : 0
  }
  if (containerIndex !== -1) {
    cpuOptions.value!.series[containerIndex].data = data
  } else {
    cpuOptions.value?.series.push({
      name: container.containerName,
      type: "line",
      data: data
    })
  }
}

function parseMemoryOptions(container: ServiceContainerStatusInfo, info: ContainerPerformance) {
  memoryOptions.value.xAxis?.[0].data.push(TimestampToTime(info.stats?.statsAt, 2))
  const titleIndex = memoryOptions.value.xAxis?.[0].data.length || 0
  const containerIndex = findIndex(memoryOptions.value?.series || [], (x: any) => x.name === container.containerName)
  const data = containerIndex !== -1 ? memoryOptions.value?.series[containerIndex].data : []
  for (let i = data.length; i < titleIndex; i++) {
    data[i] = i + 1 === titleIndex ? (info.stats!.memoryUsed / 1024 / 1024).toFixed(2) : 0
  }
  if (containerIndex !== -1) {
    memoryOptions.value!.series[containerIndex].data = data
  } else {
    memoryOptions.value?.series.push({
      name: container.containerName,
      type: "line",
      data: data
    })
  }
}

function parseNetworkOptions(container: ServiceContainerStatusInfo, info: ContainerPerformance) {
  networkOptions.value.xAxis?.[0].data.push(TimestampToTime(info.stats?.statsAt, 2))
  const titleIndex = networkOptions.value.xAxis?.[0].data.length || 0
  map(info.stats?.networks, network => {
    const containerReadName = `${container.containerName} - ${network.name} - RX`
    const containerWriteName = `${container.containerName} - ${network.name} - TX`
    const readIndex = findIndex(networkOptions.value?.series || [], (x: any) => x.name === containerReadName && x.yAxisIndex === 0)
    const writeIndex = findIndex(networkOptions.value?.series || [], (x: any) => x.name === containerWriteName && x.yAxisIndex === 1)
    const readData = readIndex !== -1 ? networkOptions.value?.series[readIndex].data : []
    const writeData = writeIndex !== -1 ? networkOptions.value?.series[writeIndex].data : []

    for (let i = readData.length; i < titleIndex; i++) {
      readData[i] = i + 1 === titleIndex ? (network.rxBytes / 1024 / 1024).toFixed(2) : 0
    }

    for (let i = writeData.length; i < titleIndex; i++) {
      writeData[i] = i + 1 === titleIndex ? (network.txBytes / 1024 / 1024).toFixed(2) : 0
    }

    if (readIndex !== -1) {
      networkOptions.value!.series![readIndex].data = readData
    } else {
      ;(networkOptions.value?.series as any[]).push({
        name: containerReadName,
        type: "line",
        yAxisIndex: 0,
        data: readData
      })
    }

    if (writeIndex !== -1) {
      networkOptions.value!.series![writeIndex].data = writeData
    } else {
      ;(networkOptions.value?.series as any[]).push({
        name: containerWriteName,
        type: "line",
        yAxisIndex: 1,
        data: writeData
      })
    }
  })
}

function parseIoOptions(container: ServiceContainerStatusInfo, info: ContainerPerformance) {
  ioOptions.value.xAxis?.[0].data.push(TimestampToTime(info.stats?.statsAt, 2))
  const titleIndex = ioOptions.value.xAxis?.[0].data.length || 0
  const readName = `${container.containerName} - Read`
  const writeName = `${container.containerName} - Write`
  const readIndex = findIndex(ioOptions.value?.series || [], (x: any) => x.name === readName && x.yAxisIndex === 0)
  const writeIndex = findIndex(ioOptions.value?.series || [], (x: any) => x.name === writeName && x.yAxisIndex === 1)
  const readData = readIndex !== -1 ? ioOptions.value?.series[readIndex].data : []
  const writeData = writeIndex !== -1 ? ioOptions.value?.series[writeIndex].data : []

  for (let i = readData.length; i < titleIndex; i++) {
    readData[i] = i + 1 === titleIndex ? info.stats!.ioRead : 0
  }
  for (let i = writeData.length; i < titleIndex; i++) {
    writeData[i] = i + 1 === titleIndex ? info.stats!.ioRead : 0
  }

  if (readIndex !== -1) {
    ioOptions.value!.series[readIndex].data = readData
  } else {
    ioOptions.value?.series.push({
      name: readName,
      type: "line",
      yAxisIndex: 0,
      data: readData
    })
  }

  if (writeIndex !== -1) {
    ioOptions.value!.series[writeIndex].data = writeData
  } else {
    ioOptions.value?.series.push({
      name: writeName,
      type: "line",
      yAxisIndex: 1,
      data: writeData
    })
  }
}

function parseStatsToChart(statsList: ContainerPerformance[]) {
  map(statsList, item => {
    const container = find(serviceInfo.value?.containers || [], x => x.containerId === item.containerId)
    if (!container || !item.isSuccess) {
      return
    }
    parseCpuOption(container, item)
    parseMemoryOptions(container, item)
    parseNetworkOptions(container, item)
    parseIoOptions(container, item)
  })
  resetChartData()
}

async function getPerformance() {
  const validContainers = filter(serviceInfo.value?.containers || [], x => !!x.containerId)
  const containers = map(validContainers, x => ({ nodeId: x.nodeId, containerId: x.containerId }))
  if (!containers?.length) {
    return
  }
  const statsList = await groupContainerService.performance(groupId.value, containers)
  parseStatsToChart(statsList)
}

function loopSearchPerformance() {
  timer.value = setTimeout(async () => {
    if (serviceInfo.value?.isEnabled) {
      await getPerformance().catch(() => {})
    }
    loopSearchPerformance()
  }, 5000)
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
  window.addEventListener("resize", resize)
  await getPerformance()
  loopSearchPerformance()
})

onUnmounted(() => {
  window.removeEventListener("resize", resize)
  cpuChart?.dispose()
  memoryChart?.dispose()
  networkChart?.dispose()
  ioChart?.dispose()

  if (timer.value) {
    clearTimeout(timer.value)
    timer.value = null
  }
})
</script>

<template>
  <div class="d-flex gap-3">
    <div class="header-icon">
      <el-icon :size="18">
        <IconMdiPerformance />
      </el-icon>
    </div>
    <el-text class="f-bold" size="large">{{ t("label.performance") }}</el-text>
  </div>

  <div v-loading="isLoading">
    <el-row :gutter="20">
      <el-col :md="12" :span="24" class="mt-5">
        <div class="chart-box">
          <div class="chart-header">
            <el-text>{{ t("label.cpuUsage") }}</el-text>
          </div>
          <div ref="cpuRef" class="chart" />
        </div>
      </el-col>

      <el-col :md="12" :span="24" class="mt-5">
        <div class="chart-box">
          <div class="chart-header">
            <el-text>{{ t("label.memoryUsage") }}</el-text>
          </div>
          <div ref="memoryRef" class="chart" />
        </div>
      </el-col>

      <el-col :md="12" :span="24" class="mt-5">
        <div class="chart-box">
          <div class="chart-header">
            <el-text>{{ t("label.networkUsage") }}</el-text>
          </div>
          <div ref="networkRef" class="chart" />
        </div>
      </el-col>

      <el-col :md="12" :span="24" class="mt-5">
        <div class="chart-box">
          <div class="chart-header">
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
  }

  .chart {
    height: 400px;
    width: 100%;
  }
}
</style>
