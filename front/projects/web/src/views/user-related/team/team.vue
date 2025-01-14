<script lang="ts" setup>
import { TeamInfo } from "@/types"
import { TableHeight } from "@/utils"
import { Action } from "@/models"
import TeamEdit from "./team-edit.vue"
import TeamDelete from "./team-delete.vue"
import TeamViewUsers from "./team-view-users.vue"

const { t } = useI18n()

const tableHeight = computed(() => TableHeight(258))

const queryInfo = ref<QueryInfo>({
  keywords: "",
  mode: "name",
  pageInfo: {
    index: 1,
    size: 20
  },
  filter: {
    role: 0
  },
  sortInfo: {
    field: "name",
    order: "asc"
  }
})

const tableList = ref({
  total: 0,
  data: [] as Array<TeamInfo>
})

const teamEditRef = useTemplateRef<InstanceType<typeof TeamEdit>>("teamEditRef")
const teamDeleteRef = useTemplateRef<InstanceType<typeof TeamDelete>>("teamDeleteRef")
const teamViewUsersRef = useTemplateRef<InstanceType<typeof TeamViewUsers>>("teamViewUsersRef")

function search() {}

function openAction(action: string, info?: TeamInfo) {
  switch (action) {
    case Action.Add:
    case Action.Edit:
      teamEditRef.value?.open(info)
      break
    case Action.Delete:
      teamDeleteRef.value?.open(info!)
      break
    case Action.View:
      teamViewUsersRef.value?.open(info!)
      break
  }
}

onMounted(() => {
  const data: Array<TeamInfo> = [
    {
      teamId: "supper admin",
      name: "John Doe",
      description: "Developer",
      createdAt: 1736744969984,
      updatedAt: 1736744969984,
      users: []
    },
    {
      teamId: "admin",
      name: "John Doe",
      description: "Developer",
      createdAt: 1736744969984,
      updatedAt: 1736744969984,
      users: []
    },
    {
      teamId: "user",
      name: "John Doe",
      description: "Developer",
      createdAt: 1736744969984,
      updatedAt: 1736744969984,
      users: ["sfljsf", "flsjfls"]
    }
  ]
  for (let i = 0; i < 10; i++) {
    tableList.value.data.push(...data)
  }
  tableList.value.total = tableList.value.data.length
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
          <el-button plain type="primary" @click="openAction(Action.Add)">
            <template #icon>
              <el-icon :size="20">
                <IconMdiAdd />
              </el-icon>
            </template>
            {{ t("btn.addTeam") }}
          </el-button>
        </div>
      </div>
    </el-form-item>
  </el-form>
  <v-table
    v-model:page-info="queryInfo.pageInfo"
    v-model:sort-info="queryInfo.sortInfo"
    :data="tableList.data"
    :max-height="tableHeight"
    :total="tableList.total"
    @page-change="search"
    @sort-change="search">
    <el-table-column :label="t('label.name')" fixed="left" min-width="140" prop="name" sortable="custom" />
    <el-table-column :label="t('label.description')" min-width="140" prop="description">
      <template #default="scope">
        <v-table-column-none :text="scope.row.description" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.users')" min-width="120">
      <template #default="scope">
        <el-button v-if="Array.isArray(scope.row.users) && scope.row.users.length > 0" link type="primary" @click="openAction(Action.View, scope.row)">
          {{ t("label.totalUsers", { total: scope.row.users.length }) }}
        </el-button>
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

  <team-delete ref="teamDeleteRef" @refresh="search()" />

  <team-edit ref="teamEditRef" @refresh="search()" />

  <team-view-users ref="teamViewUsersRef" />
</template>

<style lang="scss" scoped></style>
