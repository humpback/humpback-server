<script lang="ts" setup>
import { TabPaneName } from "element-plus"
import TeamPage from "./team/teams.vue"
import UserPage from "./user/users.vue"
import { find, toLower } from "lodash-es"

enum UserRelatedType {
  Users = "users",
  Teams = "teams"
}

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const activeTab = ref<TabPaneName | undefined>()

async function changeTab(name: TabPaneName) {
  await router.replace({ params: { mode: name } })
  activeTab.value = name
}

const options = reactive<{ name: UserRelatedType; label: string; component: any }[]>([
  { name: UserRelatedType.Users, label: "header.users", component: shallowRef(UserPage) },
  { name: UserRelatedType.Teams, label: "header.teams", component: shallowRef(TeamPage) }
])

onBeforeMount(async () => {
  const t = find(options, x => x.name === toLower(route.params["mode"] as string))
  if (!t) {
    await router.push({ name: "404" })
    return
  }
  await changeTab(t.name as UserRelatedType)
})
</script>

<template>
  <v-card class="tab-card">
    <el-tabs :model-value="activeTab" class="tab-box" @update:modelValue="changeTab">
      <template v-for="item in options" :key="item.name">
        <el-tab-pane :label="t(item.label)" :name="item.name">
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
    margin-bottom: 20px;
  }
}
</style>
