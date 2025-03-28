<script lang="ts" setup>
import { PageActivity } from "@/models"
import { TabPaneName } from "element-plus"
import PageGroups from "./groups/groups.vue"
import PageServices from "./services/services.vue"
import PageConfigs from "./configs/configs.vue"
import PageNodes from "./nodes/nodes.vue"
import PageUsers from "./user-related/users.vue"
import PageTeams from "./user-related/teams.vue"
import { find } from "lodash-es"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const activeTab = ref<TabPaneName | undefined>(route.params?.mode as string)

async function changeTab(name: TabPaneName) {
  await router.replace({ params: { mode: name } })
  activeTab.value = name
}

const options = reactive<Array<{ name: string; label: string; component: any; limitAdmin?: boolean }>>([
  { name: PageActivity.Groups, label: "header.groups", component: shallowRef(PageGroups) },
  { name: PageActivity.Services, label: "header.services", component: shallowRef(PageServices) },
  { name: PageActivity.Configs, label: "header.configs", component: shallowRef(PageConfigs) },
  { name: PageActivity.Nodes, label: "header.nodes", component: shallowRef(PageNodes), limitAdmin: true },
  { name: PageActivity.Users, label: "header.users", component: shallowRef(PageUsers), limitAdmin: true },
  { name: PageActivity.Teams, label: "header.teams", component: shallowRef(PageTeams), limitAdmin: true }
])

onMounted(() => {
  const item = find(options, x => x.name === (route.params.mode as string))
  if (!item || (item.limitAdmin && userStore.isUser)) {
    changeTab(PageActivity.Groups)
  }
})
</script>

<template>
  <v-card class="tab-box">
    <v-page-title :title="t('label.activities')" />
    <el-tabs :model-value="activeTab" class="tab-box" @update:modelValue="changeTab">
      <template v-for="item in options" :key="item.name">
        <el-tab-pane v-if="!item.limitAdmin || !userStore.isUser" :label="t(item.label)" :name="item.name">
          <template #label>
            <strong>{{ t(item.label) }}</strong>
          </template>
          <component :is="item.component" v-if="item.name === activeTab" />
        </el-tab-pane>
      </template>
    </el-tabs>
  </v-card>
</template>

<style lang="scss" scoped>
.tab-box {
  :deep(.el-tabs__header.is-top) {
    margin-bottom: 0;
  }

  :deep(.el-tabs__content) {
    padding-top: 20px;
  }
}
</style>
