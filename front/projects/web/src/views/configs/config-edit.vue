<script lang="ts" setup>
import { ConfigInfo, NewConfigEmptyInfo } from "@/types"
import { cloneDeep } from "lodash-es"
import { FormInstance, FormRules } from "element-plus"
import { RuleFormatErrEmailOption, RulePleaseEnter } from "@/utils"
import { LimitConfigValue, LimitDescription } from "@/models"

const { t } = useI18n()

const isAction = ref(false)
const dialogInfo = ref({
  show: false,
  info: {} as ConfigInfo
})

const formRef = useTemplateRef<FormInstance>("formRef")
const rules = ref<FormRules>({
  configName: [
    { required: true, validator: RulePleaseEnter("label.name"), trigger: "blur" },
    { required: true, validator: RuleLimitRange(LimitUserName.Min, LimitUserName.Max), trigger: "blur" }
  ],
  configType: [{ validator: RuleFormatErrEmailOption(), trigger: "blur" }],
  configValue: [{ validator: RuleLimitRange(LimitConfigValue.Min, LimitConfigValue.Max), trigger: "blur" }],
  description: [{ validator: RuleLimitMax(LimitDescription.Max), trigger: "blur" }]
})

function open(info?: ConfigInfo) {
  dialogInfo.value.info = info ? cloneDeep(info) : NewConfigEmptyInfo()
  dialogInfo.value.show = true
}

function save() {}

defineExpose({ open })
</script>

<template>
  <v-dialog v-model="dialogInfo.show" :close-on-press-escape="false">
    <template #header>{{ dialogInfo.info.configId ? t("header.editConfig") : t("header.addConfig") }}</template>
    <div class="my-3">
      <el-form ref="formRef" :model="dialogInfo.info" :rules="rules" label-position="top" label-width="auto">
        <el-form-item :label="t('label.name')" prop="configName">
          <v-username-input v-model="dialogInfo.info.configName" />
        </el-form-item>

        <el-form-item :label="t('label.description')" prop="description">
          <v-description-input v-model="dialogInfo.info.description" />
        </el-form-item>

        <el-form-item :label="t('label.value')" prop="configType">
          <el-radio-group v-model="dialogInfo.info.configType">
            <el-radio :value="ConfigType.Static">{{ t("label.static") }}</el-radio>
            <el-radio :value="ConfigType.Volume">{{ t("label.volume") }}</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item v-if="dialogInfo.info.configType === ConfigType.Static" prop="configType">
          <v-input
            v-model="dialogInfo.info.configValue"
            :autosize="{ minRows: 2, maxRows: 4 }"
            :maxlength="LimitConfigValue.Max / 2"
            resize="vertical"
            show-word-limit
            type="textarea" />
        </el-form-item>

        <el-form-item v-if="dialogInfo.info.configType === ConfigType.Volume" prop="configType">
          <div style="width: 100%; height: 400px">
            <v-monaco-edit />
          </div>
        </el-form-item>
      </el-form>
    </div>
    <template #footer>
      <el-button @click="dialogInfo.show = false">{{ t("btn.cancel") }}</el-button>
      <el-button :loading="isAction" type="primary" @click="save">{{ t("btn.save") }}</el-button>
    </template>
  </v-dialog>
</template>

<style lang="scss" scoped></style>
