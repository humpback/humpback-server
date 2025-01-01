<script lang="ts" setup>
import BasicInfo from "./mode/basic-info.vue"
import ChangePassword from "./mode/change-password.vue"
import { TabPaneName } from "element-plus"
import { find, toLower } from "lodash"

enum UserMode {
  BasicInfo = "basic-info",
  ChangePsd = "change-psd"
}

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const pageStore = usePageStore()

const activeTab = ref<TabPaneName | undefined>()

async function changeTab(name: TabPaneName) {
  await router.replace({ params: { "mode": name.toString() } })
  activeTab.value = name
}

const modeList = reactive<{ name: UserMode; label: string; component: any }[]>([
  { name: UserMode.BasicInfo, label: "header.basicInfo", component: shallowRef(BasicInfo) },
  { name: UserMode.ChangePsd, label: "header.changePassword", component: shallowRef(ChangePassword) }
])

onMounted(() => {
  const configType = toLower(route.params["mode"] as string)
  const tabInfo = find(modeList, x => configType === x.name)
  changeTab(tabInfo ? tabInfo.name : UserMode.BasicInfo)
})
</script>

<template>
  <v-card>
    <el-tabs
      :class="pageStore.isSmallScreen ? '' : 'user-page-tab'"
      :modelValue="activeTab"
      :stretch="!pageStore.isSmallScreen"
      :tab-position="pageStore.isSmallScreen ? 'top' : 'left'"
      @update:modelValue="changeTab">
      <el-tab-pane v-for="item in modeList" :key="item.name" :label="t(item.label)" :name="item.name">
        <h3 v-if="!pageStore.isSmallScreen" style="margin-top: 0">{{ t(item.label) }}</h3>
        <div :class="pageStore.isSmallScreen ? '' : 'content-box'">
          <keep-alive>
            <component :is="item.component" />
          </keep-alive>
        </div>
      </el-tab-pane>
    </el-tabs>
  </v-card>
</template>

<style lang="scss" scoped>
.user-page-tab {
  :deep(.el-tabs__content) {
    padding: 0 10px;
  }

  :deep(.el-tabs__nav) {
    min-height: 600px;

    .el-tabs__active-bar.is-left {
      background-color: rgba(64, 158, 255, 0.1);
      width: 100%;
      border-right: 2px solid var(--el-color-primary);
    }

    .el-tabs__item {
      justify-content: left;
      width: 200px;
    }
  }
}

.content-box {
  padding: 10px 60px 0 60px;
}

.notify-box {
  font-size: 16px;
}
</style>
