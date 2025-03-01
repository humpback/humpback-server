<script lang="ts" setup>
import { cloneDeep } from "lodash-es"
import { ServiceInfo } from "@/types"

const emits = defineEmits<{
  (e: "refresh"): void
}>()

const { t } = useI18n()

const isAction = ref(false)
const isChecked = ref(false)
const dialogInfo = ref({
  show: false,
  info: {} as ServiceInfo
})

function open(info: ServiceInfo) {
  isChecked.value = false
  dialogInfo.value.info = cloneDeep(info)
  dialogInfo.value.show = true
}

async function confirmDelete() {
  if (!isChecked.value) {
    return
  }
  // isAction.value = true
  // return await serviceService
  //   .delete(dialogInfo.value.info.serviceId)
  //   .then(() => {
  //     ShowSuccessMsg(t("message.deleteSuccess"))
  //     dialogInfo.value.show = false
  //     emits("refresh")
  //   })
  //   .finally(() => (isAction.value = false))
}

defineExpose({ open })
</script>

<template>
  <v-dialog v-model="dialogInfo.show" width="600px">
    <template #header>{{ t("header.deleteService") }}</template>
    <v-alert type="warning">{{ t("tips.deleteServiceTips") }}</v-alert>
    <div class="my-3">
      <strong>{{ t("notify.delete") }}</strong>
    </div>
    <v-delete-input-continue v-model="isChecked" :keywords="dialogInfo.info.serviceName" class="mt-5" />
    <template #footer>
      <el-button @click="dialogInfo.show = false">{{ t("btn.cancel") }}</el-button>
      <el-button :disabled="!isChecked" :loading="isAction" type="danger" @click="confirmDelete">{{ t("btn.delete") }}</el-button>
    </template>
  </v-dialog>
</template>

<style lang="scss" scoped></style>
