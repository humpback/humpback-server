<script lang="ts" setup>
import { TeamInfo, UserInfo } from "@/types"

const { t } = useI18n()

const dialogInfo = ref({
  show: false,
  info: {} as TeamInfo
})
const list = ref({
  total: 0,
  data: [] as Array<UserInfo>
})

function open(info: TeamInfo) {
  list.value = {
    total: 0,
    data: [] as Array<UserInfo>
  }
  dialogInfo.value.info = info
  dialogInfo.value.show = true
}

defineExpose({ open })
</script>

<template>
  <v-dialog v-model="dialogInfo.show" width="800px">
    <template #header>
      <span v-html="t('header.teamUsers', { teamName: dialogInfo.info.name })"></span>
    </template>
    <v-table :data="list.data" :max-height="600" :show-header="false" :total="list.total">
      <el-table-column :label="t('label.username')" min-width="140" prop="username" />
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
      <el-table-column :label="t('label.role')" min-width="160" prop="groupList">
        <template #default="scope">
          <v-role-view :role="scope.row.role" />
        </template>
      </el-table-column>
    </v-table>
  </v-dialog>
</template>

<style lang="scss" scoped></style>
