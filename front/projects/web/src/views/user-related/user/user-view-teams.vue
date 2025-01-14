<script lang="ts" setup>
import { TeamInfo, UserInfo } from "@/types"

const { t } = useI18n()

const dialogInfo = ref({
  show: false,
  info: {} as UserInfo
})
const list = ref({
  total: 0,
  data: [] as Array<TeamInfo>
})

function open(info: UserInfo) {
  list.value = {
    total: 0,
    data: [] as Array<TeamInfo>
  }
  dialogInfo.value.info = info
  dialogInfo.value.show = true
}

defineExpose({ open })
</script>

<template>
  <v-dialog v-model="dialogInfo.show" width="600px">
    <template #header>
      <span v-html="t('header.userTeams', { username: dialogInfo.info.username })"></span>
    </template>
    <v-table :data="list.data" :max-height="600" :show-header="false" :total="list.total">
      <el-table-column :label="t('label.name')" min-width="140" prop="name" />
      <el-table-column :label="t('label.description')" min-width="140" prop="description">
        <template #default="scope">
          <v-table-column-none :text="scope.row.description" />
        </template>
      </el-table-column>
    </v-table>
  </v-dialog>
</template>

<style lang="scss" scoped></style>
