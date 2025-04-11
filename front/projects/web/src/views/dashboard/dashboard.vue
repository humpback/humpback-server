<script lang="ts" setup>
import TotalPage from "./total.vue"
import TimeLinePage from "./timeline.vue"
import GreetingPage from "./greetings.vue"
import ActivitiesPage from "./activities.vue"
import ExceptionServicesPage from "./exception-services.vue"
import AbnormalNodesPage from "./abnormal-nodes.vue"
import { DashboardResourceStatisticsInfo, NewDashboardResourceStatisticsInfo } from "@/types"

const timelineRef = useTemplateRef<InstanceType<typeof TimeLinePage>>("timelineRef")
const resourceStatisticsLoading = ref(false)
const resourceStatisticsInfo = ref<DashboardResourceStatisticsInfo>(NewDashboardResourceStatisticsInfo())
function get30DaysAgoMidnightTimestamp() {
  const now = new Date()
  now.setHours(0, 0, 0, 0)
  const pastDate = new Date(now.getTime() - 29 * 24 * 60 * 60 * 1000)
  return pastDate.getTime() // 返回毫秒时间戳
}

async function getResourceStatisticsInfo() {
  resourceStatisticsLoading.value = true
  return await dashboardService
    .getResourceStatistics()
    .then(statistics => {
      resourceStatisticsInfo.value = statistics
    })
    .finally(() => (resourceStatisticsLoading.value = false))
}

async function getStatisticsCountData() {
  const data = {
    startAt: get30DaysAgoMidnightTimestamp()
  }
  return await statisticsCountService.query(data).then(statistics => {
    timelineRef.value?.setData(data.startAt, statistics)
  })
}

onMounted(async () => {
  await Promise.all([getResourceStatisticsInfo(), getStatisticsCountData()])
})
</script>

<template>
  <div>
    <div>
      <total-page :info="resourceStatisticsInfo" :is-loading="resourceStatisticsLoading" />
    </div>

    <div class="mt-5">
      <greeting-page
        :is-loading="resourceStatisticsLoading"
        :owner-groups="resourceStatisticsInfo.ownGroups"
        :owner-services="resourceStatisticsInfo.ownServices" />
    </div>

    <div class="mt-5">
      <time-line-page ref="timelineRef" />
    </div>

    <div class="mt-5">
      <el-row :gutter="20">
        <el-col :span="8">
          <activities-page />
        </el-col>
        <el-col :span="8">
          <exception-services-page
            :data="resourceStatisticsInfo.exceptionServices"
            :enabled-services="resourceStatisticsInfo.enableOwnServices"
            :is-loading="resourceStatisticsLoading" />
        </el-col>
        <el-col :span="8">
          <abnormal-nodes-page
            :data="resourceStatisticsInfo.abnormalNodes"
            :enabled-nodes="resourceStatisticsInfo.enableOwnNodes"
            :is-loading="resourceStatisticsLoading" />
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.total-title {
  font-weight: bold;
  font-size: 20px;
}

.total-content {
  margin-top: 20px;
  font-size: 20px;
}
</style>
