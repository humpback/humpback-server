<script lang="ts" setup>
import { GroupInfo } from "@/types"
import { TableHeight } from "@/utils"
import { Action } from "@/models"
import { QueryGroupsInfo } from "./common.ts"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const userStore = useUserStore()
const stateStore = useStateStore()

const tableHeight = computed(() => TableHeight(252))

const groupId = ref(route.params?.groupId as string)

const resourceTotal = defineModel<{ services: number; nodes: number }>()

const isLoading = ref(false)
const queryInfo = ref<QueryGroupsInfo>(new QueryGroupsInfo(route.query))

const tableList = ref({
  total: 0,
  data: [] as Array<GroupInfo>
})

const isAdmin = computed(() => userStore.isAdmin || userStore.isSupperAdmin)

function getGroupInfo() {
  return groupService.info(groupId.value).then(info => {
    stateStore.setGroup(groupId, info)
  })
}

async function search() {
  await router.replace(queryInfo.value.urlQuery())
  isLoading.value = true
  return await groupService
    .query(queryInfo.value.searchParams())
    .then(res => {
      tableList.value.data = res.list
      tableList.value.total = res.total
    })
    .finally(() => (isLoading.value = false))
}

function openAction(action: string, info?: GroupInfo) {}

onMounted(async () => {
  await getGroupInfo()
})
</script>

<template>
  <el-form @submit.prevent="search">
    <el-form-item>
      <div class="d-flex gap-3 w-100 flex-wrap">
        <div class="flex-1" style="min-width: 300px">
          <v-input v-model="queryInfo.keywords">
            <template #prepend>
              <span>{{ t("label.name") }}</span>
            </template>
          </v-input>
        </div>
        <div>
          <el-button native-type="submit" type="primary">{{ t("btn.search") }}</el-button>
          <el-button v-if="isAdmin" plain type="primary" @click="openAction(Action.Add)">
            <template #icon>
              <el-icon :size="20">
                <IconMdiAdd />
              </el-icon>
            </template>
            {{ t("btn.addGroup") }}
          </el-button>
        </div>
      </div>
    </el-form-item>
  </el-form>

  <v-table
    v-loading="isLoading"
    v-model:page-info="queryInfo.pageInfo"
    v-model:sort-info="queryInfo.sortInfo"
    :data="tableList.data"
    :max-height="tableHeight"
    :total="tableList.total"
    @page-change="search"
    @sort-change="search">
    <el-table-column :label="t('label.group')" fixed="left" min-width="200" prop="groupName" sortable="custom">
      <template #default="scope">
        <v-router-link :href="`/ws/group/${scope.row.groupId}/detail`" :text="scope.row.groupName" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.description')" min-width="200" prop="description">
      <template #default="scope">
        <v-table-column-none :text="scope.row.description" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.nodes')" min-width="100" prop="description">
      <template #default="scope">
        <v-router-link
          v-if="scope.row.nodes.length > 0"
          :href="`/ws/group/${scope.row.groupId}/detail?tab=nodes`"
          :text="t('label.totalNodes', { total: scope.row.nodes.length })" />
        <span v-else>--</span>
      </template>
    </el-table-column>
    <el-table-column :label="t('label.updateDate')" min-width="140" prop="updatedAt" sortable="custom">
      <template #default="scope">
        <v-date-view :timestamp="scope.row.updatedAt" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.createDate')" min-width="140" prop="createdAt" sortable="custom">
      <template #default="scope">
        <v-date-view :timestamp="scope.row.createdAt" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.action')" align="right" fixed="right" header-align="center" width="130">
      <template #default="scope">
        <el-button link type="primary" @click="openAction(Action.Edit, scope.row)">{{ t("btn.edit") }}</el-button>
        <el-button link type="danger" @click="openAction(Action.Delete, scope.row)">{{ t("btn.delete") }}</el-button>
      </template>
    </el-table-column>
  </v-table>
</template>

<style lang="scss" scoped>
.ellipsis {
  display: inline-block;
  width: 100%;
  //max-width: 150px; /* 设置最大宽度 */
  white-space: nowrap; /* 禁止换行 */
  overflow: hidden; /* 隐藏溢出内容 */
  text-overflow: ellipsis; /* 使用省略号表示溢出 */
}
</style>
