<script lang="ts" setup>
import { SetWebTitle } from "@/utils"
import VContainerStatus from "@/components/business/v-container/VContainerStatus.vue"

const { t } = useI18n()
const route = useRoute()
const stateStore = useStateStore()

const isLoading = ref(false)

const groupId = ref(route.params.groupId as string)
const serviceId = ref(route.params.serviceId as string)
const serviceInfo = ref<ServiceInfo>(NewServiceEmptyInfo())

async function getGroupInfo() {
  return await groupService.info(groupId.value).then(info => {
    stateStore.setGroup(groupId.value, info)
  })
}

async function getServiceInfo() {
  return await serviceService.info(groupId.value, serviceId.value).then(info => {
    serviceInfo.value = info
    stateStore.setService(serviceId.value, info)
    serviceInfo.value.containers = [
      {
        containerId: "FSJLFJLFJPQJT03QWGHQOGLANG;AHNGAHNG;KLANG",
        containerName: "humpback-losjfljsofwoj-lsfjlsj-jfslf",
        nodeId: "SLFJOPQW2JRFO",
        status: "running",
        statusInfo: "",
        errorMsg: "",
        image: "",
        command: "",
        network: "",
        created: 1740818013735,
        started: 1740818013735,
        lastHeartbeat: 1740818013735,
        labels: { test: "skyler" },
        env: [],
        mounts: [],
        ports: []
      }
    ]
  })
}

async function search() {
  isLoading.value = true
  await Promise.all([getGroupInfo(), getServiceInfo()]).finally(() => (isLoading.value = false))
}

onMounted(async () => {
  await search()
  SetWebTitle(`${t("webTitle.serviceInfo")} - ${stateStore.getService()?.serviceName}`)
})
</script>

<template>
  <div class="d-flex gap-3">
    <strong>
      <el-text>{{ t("label.instanceOverview") }}</el-text>
    </strong>
    <div>
      <el-button plain size="small" type="success">{{ t("btn.refresh") }}</el-button>
      <el-button plain size="small" type="primary">{{ t("btn.viewMonitor") }}</el-button>
    </div>
  </div>
  <v-table :data="serviceInfo.containers" border class="mt-5">
    <el-table-column type="expand">
      <template>
        <div class="expand-content">
          <el-row :gutter="12">
            <el-col :span="12">aaa</el-col>
            <el-col :span="12"></el-col>
          </el-row>
        </div>
      </template>
    </el-table-column>
    <el-table-column :label="t('label.name')" min-width="300" prop="containerName" sortable />
    <el-table-column :label="t('label.ip')" min-width="240"></el-table-column>
    <el-table-column :label="t('label.status')" min-width="200">
      <template #default="scope">
        <v-container-status :status="scope.row.status" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.action')" width="200">
      <template>
        <el-button :title="t('label.restart')" link type="success">
          <el-icon :size="16">
            <IconMdiRestart />
          </el-icon>
        </el-button>
        <el-button :title="t('label.start')" link type="success">
          <el-icon :size="16">
            <IconMdiPlay />
          </el-icon>
        </el-button>
        <el-button :title="t('label.stop')" link type="danger">
          <el-icon :size="16">
            <IconMdiSquare />
          </el-icon>
        </el-button>
        <el-button :title="t('label.log')" link type="primary">
          <el-icon :size="16">
            <IconMdiNoteText />
          </el-icon>
        </el-button>
      </template>
    </el-table-column>
  </v-table>
</template>

<style lang="scss" scoped>
.expand-content {
  box-sizing: border-box;
  padding: 20px;
}
</style>
