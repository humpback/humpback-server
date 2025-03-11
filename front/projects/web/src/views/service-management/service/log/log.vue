<script lang="ts" setup>
import VMonacoView from "@/components/common/v-monaco/VMonacoView.vue"
import { SetWebTitle } from "@/utils"
import { refreshData } from "@/views/service-management/service/common.ts"
import { RuleLength } from "@/models"

const { t } = useI18n()
const route = useRoute()
const stateStore = useStateStore()

const isLoading = ref(false)
const isAction = ref(false)

const groupId = ref(route.params.groupId as string)
const serviceId = ref(route.params.serviceId as string)
const serviceInfo = computed<ServiceInfo | undefined>(() => stateStore.getService(serviceId.value))
const containers = ref<ServiceContainerStatusInfo[]>([])

const searchInfo = ref({
  instance: "",
  line: 1000,
  startAt: 0,
  endAt: 0,
  showTimestamp: false
})

const logViewRef = useTemplateRef<InstanceType<typeof VMonacoView>>("logViewRef")

async function search() {
  isLoading.value = true
  await refreshData(groupId.value, serviceId.value, "log").finally(() => (isLoading.value = false))
}

onMounted(async () => {
  await search()
  containers.value = serviceInfo.value?.containers || []
  SetWebTitle(`${t("webTitle.serviceInfo")} - ${stateStore.getService()?.serviceName}`)
})
</script>

<template>
  <div>
    <div class="d-flex gap-3">
      <div class="header-icon">
        <el-icon :size="18">
          <IconMdiFileDocumentOutline />
        </el-icon>
      </div>
      <strong>
        <el-text size="large">{{ t("label.instanceLogs") }}</el-text>
      </strong>
    </div>
    <div class="d-flex gap-3 flex-wrap mt-5">
      <div class="flex-1" style="min-width: 300px; max-width: 400px">
        <v-select v-model="searchInfo.instance" :out-label="t('label.instance')" out-label-width="100px" placeholder="" show-out-label>
          <el-option v-for="item in containers" :key="item.containerId" :label="item.containerName" :value="item.containerId" />
        </v-select>
      </div>
      <div class="flex-1" style="min-width: 500px; max-width: 600px">
        <v-log-search-time-range
          v-model:end-time="searchInfo.endAt"
          v-model:start-time="searchInfo.startAt"
          :out-label="t('label.timeRange')"
          out-label-width="140px"
          show-out-label />
      </div>
      <div style="width: 200px">
        <el-input v-model.number="searchInfo.line" :max="RuleLength?.LogsLine?.Max" :min="RuleLength?.LogsLine?.Min" type="number">
          <template #prepend>
            <el-text>{{ t("label.lines") }}</el-text>
          </template>
        </el-input>
      </div>

      <el-button v-loading="isAction" plain type="primary">{{ t("btn.refresh") }}</el-button>
    </div>
  </div>

  <div class="mt-5" style="height: 600px">
    <v-monaco-view ref="logViewRef">
      <template #title>
        <div class="d-flex gap-5 pl-3">
          <el-checkbox v-model="searchInfo.showTimestamp">{{ t("label.showTimestamp") }}</el-checkbox>
        </div>
      </template>
    </v-monaco-view>
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
</style>
