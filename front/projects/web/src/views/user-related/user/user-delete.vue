<script lang="ts" setup>
import { cloneDeep } from "lodash-es"
import { UserInfo } from "@/types"

const emits = defineEmits<{
  (e: "refresh"): void
}>()

const { t } = useI18n()
const dialogInfo = ref({
  show: false,
  info: {} as UserInfo
})

function open(info: UserInfo) {
  dialogInfo.value.info = cloneDeep(info)
  dialogInfo.value.show = true
}

function confirmDelete() {}

defineExpose({ open })
</script>

<template>
  <v-dialog v-model="dialogInfo.show" width="600px">
    <template #header>{{ t("header.deleteUser") }}</template>
    <div class="my-3">
      <strong v-html="t('notify.deleteUser', { name: dialogInfo.info.username })" />
    </div>
    <template #footer>
      <el-button @click="dialogInfo.show = false">{{ t("btn.cancel") }}</el-button>
      <el-button type="danger" @click="confirmDelete">{{ t("btn.delete") }}</el-button>
    </template>
  </v-dialog>
</template>

<style lang="scss" scoped></style>
