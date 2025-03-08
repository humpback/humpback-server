<script lang="ts" setup>
import BasicInfo from "./basic-info/basic-info.vue"
import Application from "./application/application.vue"
import Deployment from "./deployment/deployment.vue"
import Instances from "./instance/instances.vue"
import Log from "./log/log.vue"
import Performance from "./performance/performance.vue"
import ServiceDelete from "./action/service-delete.vue"
import ServiceClone from "./action/servcie-clone.vue"
import { shallowRef } from "vue"
import { find } from "lodash-es"
import { ActionOptions, refreshData, showAction } from "@/views/service-management/service/common.ts"
import VLoading from "@/components/common/v-loading/VLoading.vue"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const stateStore = useStateStore()

const groupId = ref(route.params.groupId as string)
const serviceId = ref(route.params.serviceId as string)
const serviceInfo = computed(() => stateStore.getService())

const isReset = ref(false)
const isLoading = ref(false)
const isLoadingComponent = ref("")

const timer = ref<any>(null)

const compRef = ref<any>()
const cloneRef = useTemplateRef<InstanceType<typeof ServiceClone>>("cloneRef")
const deleteRef = useTemplateRef<InstanceType<typeof ServiceDelete>>("deleteRef")

const activeMenu = ref(route.params.mode as string)

const menuOptions = ref<any[]>([
  {
    i18nLabel: "label.setting",
    iconClass: "icon_mdi--settings-outline",
    isGroup: true
  },
  {
    i18nLabel: "label.basicInfo",
    value: PageServiceDetail.BasicInfo,
    isRequired: true,
    component: shallowRef(BasicInfo)
  },
  {
    i18nLabel: "label.application",
    value: PageServiceDetail.Application,
    isRequired: true,
    component: shallowRef(Application)
  },
  {
    i18nLabel: "label.deployment",
    value: PageServiceDetail.Deployment,
    isRequired: true,
    component: shallowRef(Deployment)
  },
  {
    i18nLabel: "label.monitor",
    iconClass: "icon_mdi--gauge",
    isGroup: true
  },
  {
    i18nLabel: "label.instances",
    value: PageServiceDetail.Instances,
    isRequired: false,
    component: shallowRef(Instances)
  },
  { i18nLabel: "label.log", value: PageServiceDetail.Log, isRequired: false, component: shallowRef(Log) },
  {
    i18nLabel: "label.performance",
    value: PageServiceDetail.Performance,
    isRequired: false,
    component: shallowRef(Performance)
  }
])

function showIncomplete(v: string) {
  if (v === PageServiceDetail.Application && !serviceInfo.value?.meta) {
    return true
  }
  return v === PageServiceDetail.Deployment && !serviceInfo.value?.deployment
}

function menuChange(v: string) {
  activeMenu.value = v
  router.replace({ params: Object.assign({}, route.params, { mode: v }) })
}

async function getGroupInfo() {
  return await groupService.info(groupId.value).then(info => {
    stateStore.setGroup(groupId.value, info)
  })
}

async function getServiceInfo() {
  return await serviceService.info(groupId.value, serviceId.value).then(info => {
    stateStore.setService(serviceId.value, info)
  })
}

async function search() {
  await Promise.all([getGroupInfo(), getServiceInfo()])
}

function loopSearch() {
  const tempPage = [PageServiceDetail.Instances, PageServiceDetail.Log, PageServiceDetail.Performance]
  timer.value = setTimeout(async () => {
    if (!find(tempPage, x => x === activeMenu.value && serviceInfo.value?.isEnabled)) {
      isLoading.value = true
      await refreshData(groupId.value, serviceId.value, "global")
        .catch(() => {})
        .finally(() => (isLoading.value = false))
    }
    loopSearch()
  }, 10000)
}

async function operateService(action: "Start" | "Stop" | "Restart" | "Enable" | "Disable") {
  const actinInfo = find(ActionOptions, x => x.action === action)
  if (!actinInfo) {
    return false
  }

  isLoadingComponent.value = action
  await serviceService
    .operate(stateStore.getGroup()!.groupId, {
      serviceId: serviceInfo.value!.serviceId,
      action: action
    })
    .then(async () => {
      await search()
      if (activeMenu.value == PageServiceDetail.Instances && compRef.value) {
        compRef.value?.resetLoopSearch()
      }
    })
    .finally(() => (isLoadingComponent.value = ""))
  ShowSuccessMsg(t("message.succeed"))
}

async function cloneService() {
  cloneRef.value?.open(serviceInfo.value!)
}

async function deleteService() {
  deleteRef.value?.open(serviceInfo.value!, true)
}

function resetInit() {
  groupId.value = route.params.groupId as string
  serviceId.value = route.params.serviceId as string
  activeMenu.value = PageServiceDetail.BasicInfo
  isReset.value = true
  setTimeout(() => {
    isReset.value = false
  }, 500)
}

onMounted(() => {
  loopSearch()
})

onUnmounted(() => {
  if (timer.value) {
    clearTimeout(timer.value)
    timer.value = null
  }
})
</script>

<template>
  <div>
    <div class="header">
      <v-page-title :title="serviceInfo?.serviceName" show-breadcrumbs />

      <div class="header-actions">
        <template v-for="item in ActionOptions" :key="item.action">
          <el-button
            v-if="showAction(serviceInfo, item.action)"
            :loading="isLoadingComponent === item.action"
            :type="item.type"
            @click="operateService(item.action)">
            <el-icon :size="16">
              <IconMdiSquare />
            </el-icon>
            {{ t(item.i18nLabel) }}
          </el-button>
        </template>

        <el-button plain type="primary" @click="cloneService()">
          <el-icon :size="16">
            <IconMdiCheckboxMultipleBlank />
          </el-icon>
          {{ t("btn.clone") }}
        </el-button>
        <el-button type="danger" @click="deleteService()">
          <el-icon :size="16">
            <IconMdiTrash />
          </el-icon>
          {{ t("btn.delete") }}
        </el-button>
      </div>
    </div>

    <div class="body">
      <div class="body-menu">
        <div class="mb-2 d-flex gap-1">
          <v-service-status-tag :is-enabled="serviceInfo?.isEnabled" :is-loading="isLoading" :status="serviceInfo?.status" />
          <v-loading v-if="serviceInfo?.isEnabled && serviceInfo.status !== ServiceStatus.ServiceStatusRunning" />
        </div>
        <div v-for="(item, index) in menuOptions" :key="index" class="menu-group">
          <div v-if="item.isGroup" class="menu-group-title">
            <span :class="item.iconClass" style="width: 18px; height: 18px" />
            <el-text>{{ t(item.i18nLabel) }}</el-text>
          </div>
          <div v-else :class="activeMenu === item.value && 'is-active'" class="menu-group-item" @click.stop="menuChange(item.value)">
            <div class="flex-1">
              <el-text :type="activeMenu === item.value ? 'info' : ''">{{ t(item.i18nLabel) }}</el-text>
              <el-text v-if="item.isRequired" type="danger"> *</el-text>
            </div>
            <div v-if="showIncomplete(item.value)" class="pr-3">
              <el-text type="danger">{{ t("label.incomplete") }}</el-text>
            </div>
          </div>
        </div>
      </div>

      <v-card v-loading="isReset" class="body-content">
        <template v-for="(item, index) in menuOptions" :key="index">
          <div v-if="!item.isGroup && item.value === activeMenu">
            <div v-if="isReset" class="reset-loading">
              <v-loading :size="80" />
            </div>
            <component :is="item.component" v-else :ref="(el: any) => (compRef = el)" />
          </div>
        </template>
      </v-card>
    </div>

    <service-clone ref="cloneRef" @refresh="resetInit()" />

    <service-delete ref="deleteRef" />
  </div>
</template>

<style lang="scss" scoped>
.header {
  .header-actions {
    display: flex;
    align-items: center;
    gap: 12px;
    flex-wrap: wrap;

    .el-button {
      margin: 0;
    }
  }
}

.body {
  margin-top: 12px;
  display: flex;
  align-items: start;
  gap: 12px;

  .body-menu {
    flex: 25%;
    max-width: 25%;
    min-width: 240px;
    background-color: #ffffff;
    //border: 1px solid var(--el-border-color);
    border-radius: 8px;
    padding: 20px 20px 28px 20px;
    box-sizing: border-box;

    .menu-group-title {
      display: flex;
      align-items: center;
      gap: 4px;
      background-color: #f5f5f5;
      padding: 8px;
      font-weight: 700;
      margin: 12px 0;
      border-radius: 4px;
    }

    .menu-group-item {
      display: flex;
      align-items: center;
      gap: 8px;
      height: 40px;
      padding-left: 16px;
      font-size: 14px;
      cursor: pointer;
      box-sizing: border-box;
      margin: 1px 0;

      &.is-active {
        border: 1px solid #e8e8e9;
        border-left: 4px solid #26b4ff;
        border-radius: 4px;
        background-color: #ecf0f5;
        box-sizing: border-box;
      }

      &:not(.is-active):hover {
        background-color: #f1f1f1;
        opacity: 0.7;
      }
    }
  }

  .body-content {
    flex: 75%;

    .reset-loading {
      display: flex;
      align-items: center;
      justify-content: center;
      min-height: 60vh;
    }
  }
}
</style>
