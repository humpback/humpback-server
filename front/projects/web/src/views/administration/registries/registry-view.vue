<script lang="ts" setup>
import { NewRegistryEmptyInfo, RegistryInfo } from "@/types"
import { cloneDeep } from "lodash-es"

const { t } = useI18n()

const isLoading = ref(false)
const dialogInfo = ref({
  show: false,
  info: {} as RegistryInfo
})

function open(info: RegistryInfo) {
  dialogInfo.value.info = info ? cloneDeep(info) : NewRegistryEmptyInfo()
  dialogInfo.value.show = true
  isLoading.value = true
  registryService
    .info(info.registryId, true)
    .then(data => {
      dialogInfo.value.info = data
    })
    .catch(() => (dialogInfo.value.show = false))
    .finally(() => (isLoading.value = false))
}

defineExpose({ open })
</script>

<template>
  <v-dialog v-model="dialogInfo.show" width="600px">
    <template #header>
      {{ t("header.viewRegistry") }}
    </template>
    <div v-loading="isLoading" class="my-3">
      <el-descriptions :column="1" border label-width="160px">
        <el-descriptions-item :label="t('label.url')">
          <span>{{ dialogInfo.info.url }}</span>
          <el-tag v-if="dialogInfo.info.isDefault" class="ml-3" effect="dark" round size="small" type="warning">
            {{ t("label.default") }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item :label="t('label.username')">{{ dialogInfo.info.username }}</el-descriptions-item>
        <el-descriptions-item :label="t('label.password')">{{ dialogInfo.info.password }}</el-descriptions-item>
      </el-descriptions>
    </div>
  </v-dialog>
</template>

<style lang="scss" scoped></style>
