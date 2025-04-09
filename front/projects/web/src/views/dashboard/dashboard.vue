<script lang="ts" setup>
import TotalPage from "./total.vue"
import TimeLinePage from "./timeline.vue"
import GreetingPage from "./greetings.vue"
import ActivitiesPage from "./activities.vue"
import ExceptionServicesPage from "./exception-services.vue"
import AbnormalNodesPage from "./abnormal-nodes.vue"

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
</script>

<template>
  <div>
    <div>
      <total-page />
    </div>

    <div class="mt-5">
      <greeting-page :owner-groups="statistics.group.owner" :owner-services="statistics.service.owner" />
    </div>

    <div class="mt-5">
      <time-line-page />
    </div>

    <div class="mt-5">
      <el-row :gutter="20">
        <el-col :span="8">
          <activities-page />
        </el-col>
        <el-col :span="8">
          <exception-services-page />
        </el-col>
        <el-col :span="8">
          <abnormal-nodes-page />
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
