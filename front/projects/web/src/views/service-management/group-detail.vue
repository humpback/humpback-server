<script lang="ts" setup>
import PageServices from "@/views/service-management/service/services.vue"
import PageNodes from "@/views/service-management/node/nodes.vue"
import { TabPaneName } from "element-plus"

enum DetailMode {
  Services = "services",
  Nodes = "nodes"
}

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const stateStore = useStateStore()

const activeTab = ref<TabPaneName | undefined>(route.params?.mode as string)
const resourceTotal = ref({
  services: 0,
  nodes: 0
})

const groupInfo = computed(() => {
  const info = stateStore.getGroup(route.params.groupId as string)
  console.log(info)
  return info
})

async function changeTab(name: TabPaneName) {
  await router.replace({ params: { mode: name } })
  activeTab.value = name
}
</script>

<template>
  <v-card>
    <template #bodyTitle>
      <div class="card-title">
        <span v-if="groupInfo">{{ groupInfo?.groupName }}</span>
        <el-button v-else link loading />
      </div>
    </template>

    <el-tabs v-model="activeTab" class="tab-box" type="card" @tab-change="changeTab">
      <el-tab-pane :name="DetailMode.Services">
        <template #label>
          <el-badge :offset="[10, 2]" class="mr-3" color="#28c3d7">
            <template #content>{{ resourceTotal.services }}</template>
            <strong>{{ t("header.services") }}</strong>
          </el-badge>
        </template>
        <PageServices v-if="activeTab === DetailMode.Services" v-model="resourceTotal" />
      </el-tab-pane>
      <el-tab-pane :name="DetailMode.Nodes">
        <template #label>
          <el-badge :offset="[10, 2]" class="mr-3" color="#28c3d7">
            <template #content>{{ resourceTotal.nodes }}</template>
            <strong>{{ t("header.nodes") }}</strong>
          </el-badge>
        </template>
        <PageNodes v-if="activeTab === DetailMode.Nodes" v-model="resourceTotal" />
      </el-tab-pane>
    </el-tabs>
  </v-card>
</template>

<style lang="scss" scoped>
.card-title {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 20px;
}

.tab-box {
  :deep(.el-tabs__header.is-top) {
    margin-bottom: 0;
  }

  :deep(.el-tabs__content) {
    padding: 20px;
    border: 1px solid var(--el-border-color);
    border-top: none;
    border-bottom-left-radius: 4px;
    border-bottom-right-radius: 4px;
    border-top-right-radius: 4px;
  }
}
</style>
