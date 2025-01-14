<script lang="ts" setup>
import { cloneDeep } from "lodash-es"
import { FormInstance, FormRules } from "element-plus"
import { RulePleaseEnter } from "@/utils"
import { LimitDescription } from "@/models"
import { NewTeamEmptyInfo, TeamInfo, UserInfo } from "@/types"

const emits = defineEmits<{
  (e: "refresh"): void
}>()

const { t } = useI18n()

const dialogInfo = ref({
  show: false,
  info: {} as TeamInfo
})

const userOptions = ref<UserInfo[]>([])

const formRef = useTemplateRef<FormInstance>("formRef")
const rules = ref<FormRules>({
  name: [
    { required: true, validator: RulePleaseEnter("label.name"), trigger: "blur" },
    { required: true, validator: RuleLimitRange(LimitUserName.Min, LimitUserName.Max), trigger: "blur" }
  ],
  description: [{ validator: RuleLimitMax(LimitDescription.Max), trigger: "blur" }]
})

function open(info?: TeamInfo) {
  dialogInfo.value.info = info ? cloneDeep(info) : NewTeamEmptyInfo()
  dialogInfo.value.show = true
}

async function save() {
  if (!(await formRef.value?.validate())) {
    return
  }
}

defineExpose({ open })
</script>

<template>
  <v-dialog v-model="dialogInfo.show" width="800px">
    <template #header>{{ dialogInfo.info.teamId ? t("header.editTeam") : t("header.addTeam") }}</template>
    <div class="my-3">
      <el-form ref="formRef" :model="dialogInfo.info" :rules="rules" label-position="top" label-width="auto">
        <el-form-item :label="t('label.name')" prop="name">
          <v-username-input v-model="dialogInfo.info.name" />
        </el-form-item>
        <el-form-item :label="t('label.description')" prop="description">
          <v-description-input v-model="dialogInfo.info.description" />
        </el-form-item>
        <el-form-item :label="t('label.users')" prop="teams">
          <v-user-use-options-select v-model="dialogInfo.info.users" :options="userOptions" />
        </el-form-item>
      </el-form>
    </div>
    <template #footer>
      <el-button @click="dialogInfo.show = false">{{ t("btn.cancel") }}</el-button>
      <el-button type="primary" @click="save">{{ t("btn.save") }}</el-button>
    </template>
  </v-dialog>
</template>

<style lang="scss" scoped></style>
