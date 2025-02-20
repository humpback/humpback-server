<script lang="ts" setup>
import PageServices from "@/views/service-management/service/services.vue"
import PageNodes from "@/views/service-management/node/nodes.vue"
import { TabPaneName } from "element-plus"
import { PageGroupDetail } from "@/models"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const stateStore = useStateStore()

const activeTab = ref<TabPaneName | undefined>(route.params?.mode as string)

const groupInfo = computed(() => stateStore.getGroup())

const options = reactive<{ name: string; label: string; component: any }[]>([
  { name: PageGroupDetail.Services, label: "header.services", component: shallowRef(PageServices) },
  { name: PageGroupDetail.Nodes, label: "header.nodes", component: shallowRef(PageNodes) }
])

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
      <el-tab-pane v-for="item in options" :key="item.name" :name="item.name">
        <template #label>
          <el-badge :offset="[10, 2]" class="mr-3" color="#28c3d7">
            <template #content>{{ groupInfo?.total[item.name] || 0 }}</template>
            <strong>{{ t(item.label) }}</strong>
          </el-badge>
        </template>
        <component :is="item.component" v-if="activeTab === item.name" />
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
