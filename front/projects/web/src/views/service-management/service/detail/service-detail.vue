<script lang="ts" setup>
import BasicInfo from "./basic-info.vue"
import Application from "./application.vue"
import Deployment from "./deployment.vue"
import Instances from "./instances.vue"
import Log from "./log.vue"
import Performance from "./performance.vue"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const stateStore = useStateStore()

const serviceInfo = computed(() => stateStore.getService())
const completedInfo = computed(() => {
  return {
    application: !!serviceInfo.value?.meta,
    deployment: !!serviceInfo.value?.deployment
  }
})

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
    isCompleted: completedInfo.value.application,
    isRequired: true,
    component: shallowRef(Application)
  },
  {
    i18nLabel: "label.deployment",
    value: PageServiceDetail.Deployment,
    isCompleted: completedInfo.value.deployment,
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

function menuChange(v: string) {
  activeMenu.value = v
  router.replace({ params: Object.assign({}, route.params, { mode: v }) })
}
</script>

<template>
  <div class="header">
    <div class="card-title">
      <span v-if="serviceInfo">{{ serviceInfo?.serviceName }}</span>
      <el-button v-else link loading />
    </div>
    <div class="header-actions">
      <el-button type="info">
        <el-icon :size="16">
          <IconMdiSquare />
        </el-icon>
        {{ t("btn.disable") }}
      </el-button>
      <el-button type="primary">
        <el-icon :size="16">
          <IconMdiPlay />
        </el-icon>
        {{ t("btn.enable") }}
      </el-button>
      <el-button type="success">
        <el-icon :size="16">
          <IconMdiRestart />
        </el-icon>
        {{ t("btn.restart") }}
      </el-button>
      <el-button type="success">
        <el-icon :size="16">
          <IconMdiPlay />
        </el-icon>
        {{ t("btn.start") }}
      </el-button>
      <el-button type="primary">
        <el-icon :size="16">
          <IconMdiSquare />
        </el-icon>
        {{ t("btn.stop") }}
      </el-button>
      <el-button plain type="primary">
        <el-icon :size="16">
          <IconMdiCheckboxMultipleBlank />
        </el-icon>
        {{ t("btn.clone") }}
      </el-button>
      <el-button type="danger">
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
        <el-button size="small" style="padding: 0 8px" type="success">{{ t("label.enabled") }}</el-button>
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
</template>

<style lang="scss" scoped>
.header {
  .card-title {
    font-size: 20px;
    font-weight: 600;
    margin-bottom: 20px;
  }

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
