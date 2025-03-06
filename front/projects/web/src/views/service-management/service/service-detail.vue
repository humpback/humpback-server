<script lang="ts" setup>
import BasicInfo from "./basic-info/basic-info.vue"
import Application from "./application/application.vue"
import Deployment from "./deployment/deployment.vue"
import Instances from "./instance/instances.vue"
import Log from "./log/log.vue"
import Performance from "./performance/performance.vue"
import ServiceDelete from "./action/service-delete.vue"
import { shallowRef } from "vue"
import { find } from "lodash-es"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const stateStore = useStateStore()

const groupId = ref(route.params.groupId as string)
const serviceId = ref(route.params.serviceId as string)
const serviceInfo = computed(() => stateStore.getService())

const applicationCompleted = computed(() => !!serviceInfo.value?.meta)
const deploymentCompleted = computed(() => !!serviceInfo.value?.deployment)

const deleteRef = useTemplateRef<InstanceType<typeof ServiceDelete>>("deleteRef")

const activeMenu = ref(route.params.mode as string)

const menuOptions = ref<any[]>([
  {
    i18nLabel: "label.setting",
    iconClass: "icon_mdi--settings-outline",
    isGroup: true
  },
  { i18nLabel: "label.basicInfo", value: PageServiceDetail.BasicInfo, isRequired: true, component: shallowRef(BasicInfo) },
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
  { i18nLabel: "label.instances", value: PageServiceDetail.Instances, isRequired: false, component: shallowRef(Instances) },
  { i18nLabel: "label.log", value: PageServiceDetail.Log, isRequired: false, component: shallowRef(Log) },
  { i18nLabel: "label.performance", value: PageServiceDetail.Performance, isRequired: false, component: shallowRef(Performance) }
])

const actionOptions = ref<
  Array<{
    action: "Start" | "Stop" | "Restart" | "Enable" | "Disable"
    type: "default" | "info" | "success" | "primary" | "text" | "warning" | "danger"
    i18nLabel: string
    icon: any
    isAction: boolean
  }>
>([
  { action: "Enable", type: "primary", i18nLabel: "btn.enable", icon: shallowRef(IconMdiPlay), isAction: false },
  { action: "Disable", type: "info", i18nLabel: "btn.disable", icon: shallowRef(IconMdiSquare), isAction: false },
  { action: "Start", type: "success", i18nLabel: "btn.start", icon: shallowRef(IconMdiPlay), isAction: false },
  { action: "Restart", type: "success", i18nLabel: "btn.restart", icon: shallowRef(IconMdiRestart), isAction: false },
  { action: "Stop", type: "primary", i18nLabel: "btn.stop", icon: shallowRef(IconMdiSquare), isAction: false }
])

function showAction(action: "Start" | "Stop" | "Restart" | "Enable" | "Disable") {
  if (!serviceInfo.value) {
    return false
  }
  switch (action) {
    case "Start":
    case "Stop":
    case "Restart":
    case "Disable":
      {
        if (serviceInfo.value.isEnabled) {
          return true
        }
      }
      break
    case "Enable": {
      if (!serviceInfo.value.isEnabled && applicationCompleted.value && deploymentCompleted.value) {
        return true
      }
    }
  }
  return false
}

function showIncomplete(v: string) {
  if (v === PageServiceDetail.Application && !applicationCompleted.value) {
    return true
  }
  return v === PageServiceDetail.Deployment && !deploymentCompleted.value
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

async function operateService(action: "Start" | "Stop" | "Restart" | "Enable" | "Disable") {
  const actinInfo = find(actionOptions.value, x => x.action === action)
  if (!actinInfo) {
    return false
  }
  actinInfo.isAction = true
  await serviceService
    .operate(stateStore.getGroup()!.groupId, { serviceId: serviceInfo.value!.serviceId, action: action })
    .then(async () => {
      return await search()
    })
    .finally(() => (actinInfo.isAction = false))
  ShowSuccessMsg(t("message.succeed"))
}

async function deleteService() {
  deleteRef.value?.open(serviceInfo.value!, true)
}
</script>

<template>
  <div class="header">
    <v-page-title :title="serviceInfo?.serviceName" show-breadcrumbs />

    <div class="header-actions">
      <template v-for="item in actionOptions" :key="item.action">
        <el-button v-if="showAction(item.action)" v-loading="item.isAction" :type="item.type" @click="operateService(item.action)">
          <el-icon :size="16">
            <IconMdiSquare />
          </el-icon>
          {{ t(item.i18nLabel) }}
        </el-button>
      </template>

      <el-button plain type="primary">
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
      <div class="mb-2">
        <v-service-status-tag :is-enabled="serviceInfo?.isEnabled" :status="serviceInfo?.status" />
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

    <v-card class="body-content">
      <template v-for="(item, index) in menuOptions" :key="index">
        <div v-if="!item.isGroup && item.value === activeMenu">
          <component :is="item.component" />
        </div>
      </template>
    </v-card>
  </div>

  <service-delete ref="deleteRef" />
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
  gap: 20px;

  .body-menu {
    flex: 25%;
    max-width: 25%;
    min-width: 240px;
    background-color: #ffffff;
    border: 1px solid var(--el-border-color);
    border-radius: 4px;
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
  }
}
</style>
