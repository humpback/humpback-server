<script lang="ts" setup>
import { cloneDeep } from "lodash-es"
import { GroupInfo } from "@/types"

const emits = defineEmits<{
  (e: "refresh"): void
}>()

const { t } = useI18n()

const isAction = ref(false)
const dialogInfo = ref({
  show: false,
  info: {} as GroupInfo
})

function open(info: GroupInfo) {
  dialogInfo.value.info = cloneDeep(info)
  dialogInfo.value.show = true
}

async function confirmDelete() {
  isAction.value = true
  return await groupService
    .delete(dialogInfo.value.info.groupId)
    .then(() => {
      ShowSuccessMsg(t("message.deleteSuccess"))
      dialogInfo.value.show = false
      emits("refresh")
    })
    .finally(() => (isAction.value = false))
}

defineExpose({ open })
</script>

<template>
  <v-dialog v-model="dialogInfo.show" width="600px">
    <template #header>{{ t("header.deleteGroup") }}</template>
    <v-alert type="warning">{{ t("tips.deleteGroupTips") }}</v-alert>
    <div class="my-3">
      <strong v-html="t('notify.deleteGroup', { name: dialogInfo.info.groupName })" />
    </div>
    <template #footer>
      <el-button @click="dialogInfo.show = false">{{ t("btn.cancel") }}</el-button>
      <el-button :loading="isAction" type="danger" @click="confirmDelete">{{ t("btn.delete") }}</el-button>
    </template>
  </v-dialog>
</template>

<style lang="scss" scoped></style>
