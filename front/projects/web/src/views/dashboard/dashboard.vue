<script lang="ts" setup>
import * as echarts from "echarts/core"
import { GridComponent, GridComponentOption, LegendComponent, LegendComponentOption, TooltipComponent, TooltipComponentOption } from "echarts/components"
import { LineChart, LineSeriesOption } from "echarts/charts"
import { CanvasRenderer } from "echarts/renderers"
import { UniversalTransition } from "echarts/features"

echarts.use([GridComponent, LegendComponent, TooltipComponent, LineChart, CanvasRenderer, UniversalTransition])

type EChartsOption = echarts.ComposeOption<GridComponentOption | LineSeriesOption | TooltipComponentOption | LegendComponentOption>

const { t } = useI18n()
const userStore = useUserStore()

let incrementChart: echarts.ECharts
const incrementRef = useTemplateRef<HTMLDivElement>("incrementRef")

const greetings = computed(() => {
  const currentHour = new Date().getHours()
  if (currentHour >= 6 && currentHour < 9) {
    return {
      i18nLabel: "tips.breakfastTips",
      icon: IconNotoSalutingFace
    }
  } else if (currentHour >= 9 && currentHour < 12) {
    return {
      i18nLabel: "tips.morningTips",
      icon: IconNotoGrinningFace
    }
  } else if (currentHour >= 12 && currentHour < 14) {
    return {
      i18nLabel: "tips.middayTips",
      icon: IconNotoWinkingFace
    }
  } else if (currentHour >= 14 && currentHour < 18) {
    return {
      i18nLabel: "tips.afternoonTips",
      icon: IconNotoMeltingFace
    }
  } else if (currentHour >= 18 && currentHour < 20) {
    return {
      i18nLabel: "tips.eveningTips",
      icon: IconNotoSmilingFaceWithSunglasses
    }
  } else if (currentHour >= 20 && currentHour < 23) {
    return {
      i18nLabel: "tips.eveningAfterTips",
      icon: IconNotoYawningFace
    }
  } else {
    return {
      i18nLabel: "tips.morningBeforeTips",
      icon: IconNotoSleepingFace
    }
  }
})

const statistics = ref({
  group: {
    total: 200,
    owner: 10
  },
  service: {
    total: 1000,
    owner: 200
  },
  node: {
    total: 1000
  },
  user: {
    total: 500
  }
})

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
    <div>
      <el-row :gutter="20">
        <el-col :span="6">
          <v-card>
            <div class="total-title">
              <div>{{ t("label.groups") }}</div>
            </div>
            <div class="total-content"> {{ statistics.group.total }}</div>
          </v-card>
        </el-col>
        <el-col :span="6">
          <v-card>
            <div class="total-title"> {{ t("label.services") }}</div>
            <div class="total-content"> {{ statistics.service.total }}</div>
          </v-card>
        </el-col>
        <el-col :span="6">
          <v-card>
            <div class="total-title"> {{ t("label.nodes") }}</div>
            <div class="total-content"> {{ statistics.node.total }}</div>
          </v-card>
        </el-col>
        <el-col :span="6">
          <v-card>
            <div class="total-title"> {{ t("label.users") }}</div>
            <div class="total-content"> {{ statistics.user.total }}</div>
          </v-card>
        </el-col>
      </el-row>
    </div>
    <v-card class="mt-5">
      <div class="greeting-box">
        <div>
          <el-icon :size="60">
            <component :is="greetings.icon" />
          </el-icon>
        </div>
        <div class="greeting-title">
          <div class="greeting-title-left">
            <div> {{ t(greetings.i18nLabel, { name: userStore.userInfo.username }) }}</div>
            <el-text style="font-weight: normal">{{ t("tips.thankUseTips") }}</el-text>
          </div>
          <div class="greeting-title-right">
            <div class="greeting-title-right-item">
              <div>{{ t("label.groups") }}</div>
              <div class="mt-3 f-semiBold">{{ statistics.group.owner }}</div>
            </div>
            <div>
              <div>{{ t("label.services") }}</div>
              <div class="mt-3 f-semiBold">{{ statistics.service.owner }}</div>
            </div>
          </div>
        </div>
      </div>
    </v-card>

    <div class="mt-5">
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

    <div class="mt-5">
      <v-card></v-card>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.greeting-box {
  display: flex;
  align-items: start;
  gap: 12px;

  .greeting-title {
    flex: 1;
    display: flex;
    align-items: start;
    gap: 20px;
    flex-wrap: wrap;
    padding-top: 8px;

    .greeting-title-left {
      flex: 1;
      font-size: 20px;
      font-weight: bold;
    }

    .greeting-title-right {
      display: flex;
      align-items: start;
      gap: 20px;
      padding-right: 50px;

      .greeting-title-right-item {
        min-width: 80px;
      }
    }
  }
}

.total-title {
  font-weight: bold;
  font-size: 20px;
}

.total-content {
  margin-top: 20px;
  font-size: 20px;
}

.chart {
  height: 360px;
  width: 100%;
}
</style>
