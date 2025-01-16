<script lang="ts" setup>
import { TeamInfo, UserInfo } from "@/types"

const { t } = useI18n()

const dialogInfo = ref({
  show: false,
  info: {} as UserInfo
})

const isLoading = ref(false)
const list = ref<Array<TeamInfo>>([])

function open(info: UserInfo) {
  list.value = []
  dialogInfo.value.info = info
  dialogInfo.value.show = true
  isLoading.value = true
  teamService
    .queryByUserId(info.userId)
    .then(data => {
      list.value = data
    })
    .finally(() => (isLoading.value = false))
}

defineExpose({ open })
</script>

<template>
  <v-dialog v-model="dialogInfo.show" width="600px">
    <template #header>
      <span v-html="t('header.userTeams', { username: dialogInfo.info.username })"></span>
    </template>
    <v-table v-loading="isLoading" :data="list" :max-height="600" :show-header="true" :total="list.length">
      <el-table-column :label="t('label.name')" min-width="140" prop="name" />
      <el-table-column :label="t('label.description')" min-width="140" prop="description">
        <template #default="scope">
          <v-table-column-none :text="scope.row.description" />
        </template>
      </el-table-column>
    </v-table>
    <div class="mt-5">
      <el-text v-if="list.length > 0">{{ t("label.totalTeams", { total: list.length }) }}</el-text>
    </div>
  </v-dialog>
</template>

<style lang="scss" scoped></style>
