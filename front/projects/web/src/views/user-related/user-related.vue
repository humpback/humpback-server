<script lang="ts" setup>
import { UserInfo } from "@/types"
import { IsUser, TableHeight } from "@/utils"

const { t } = useI18n()

const tableHeight = computed(() => TableHeight(209))

const queryInfo = ref<QueryInfo>({
  keywords: "",
  mode: "username",
  pageInfo: {
    index: 1,
    size: 20
  },
  filter: {
    role: 0
  },
  sortInfo: {
    field: "username",
    order: "asc"
  }
})

const tableList = ref({
  total: 0,
  data: [] as Array<UserInfo>
})

function search() {}

function editUser(info?: UserInfo) {}

function deleteUser(info: UserInfo) {}

onMounted(() => {
  const data: Array<UserInfo> = [
    {
      userId: "supper admin",
      username: "John Doe",
      email: "john@example.com",
      phone: "17721865797",
      description: "Developer",
      groups: [],
      password: "",
      role: 1,
      createdAt: 1736744969984,
      updatedAt: 1736744969984
    },
    {
      userId: "admin",
      username: "John Doe",
      email: "john@example.com",
      phone: "17721865797",
      description: "Developer",
      groups: [],
      password: "",
      role: 2,
      createdAt: 1736744969984,
      updatedAt: 1736744969984
    },
    {
      userId: "user",
      username: "John Doe",
      email: "john@example.com",
      phone: "17721865797",
      description: "Developer",
      groups: [],
      password: "",
      role: 3,
      createdAt: 1736744969984,
      updatedAt: 1736744969984
    }
  ]
  for (let i = 0; i < 10; i++) {
    tableList.value.data.push(...data)
  }
  tableList.value.total = tableList.value.data.length
})
</script>

<template>
  <v-card>
    <el-form @submit.prevent="search">
      <el-form-item>
        <div class="d-flex gap-3 w-100 flex-wrap">
          <div>
            <v-role-select v-model="queryInfo.filter.role" :placeholder="t('placeholder.all')" @change="search()" />
          </div>
          <div class="flex-1" style="min-width: 300px">
            <v-input v-model="queryInfo.keywords">
              <template #prepend>
                <el-select v-model="queryInfo.mode" placeholder="" style="width: 140px">
                  <el-option :label="t('label.username')" value="username" />
                  <el-option :label="t('label.email')" value="email" />
                  <el-option :label="t('label.phone')" value="phone" />
                </el-select>
              </template>
            </v-input>
          </div>
          <div>
            <el-button native-type="submit" type="primary">{{ t("btn.search") }}</el-button>
            <el-button plain type="primary">
              <template #icon>
                <el-icon :size="20">
                  <IconMdiAdd />
                </el-icon>
              </template>
              {{ t("btn.createUser") }}
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
      <el-table-column :label="t('label.username')" fixed="left" min-width="140" prop="username" sortable="custom" />
      <el-table-column :label="t('label.role')" min-width="120" prop="groupList">
        <template #default="scope">
          <v-role-view :role="scope.row.role" />
        </template>
      </el-table-column>
      <el-table-column :label="t('label.description')" min-width="140" prop="description">
        <template #default="scope">
          <v-table-column-none :text="scope.row.description" />
        </template>
      </el-table-column>
      <el-table-column :label="t('label.email')" min-width="140" prop="email">
        <template #default="scope">
          <v-table-column-none :text="scope.row.email" />
        </template>
      </el-table-column>
      <el-table-column :label="t('label.phone')" min-width="120" prop="phone">
        <template #default="scope">
          <v-table-column-none :text="scope.row.phone" />
        </template>
      </el-table-column>
      <el-table-column :label="t('label.groups')" min-width="120">
        <template #default="scope">
          <span v-if="!IsUser(scope.row.role)">{{ t("label.allGroups") }}</span>
          <span v-else-if="Array.isArray(scope.row.groups) && scope.row.groups.length > 0">
            {{ t("label.totalGroups", { total: scope.row.groups.length }) }}
          </span>
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
          <el-button link type="primary" @click="editUser(scope.row)">{{ t("btn.edit") }}</el-button>
          <el-button link type="danger" @click="deleteUser(scope.row)">{{ t("btn.delete") }}</el-button>
        </template>
      </el-table-column>
    </v-table>
  </v-card>
</template>

<style lang="scss" scoped></style>
