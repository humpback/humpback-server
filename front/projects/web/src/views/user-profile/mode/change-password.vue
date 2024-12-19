<script lang="ts" setup>
import type { FormInstance } from "element-plus"
import { RSAEncrypt } from "utils/rsa.ts"

const { t } = useI18n()
const userStore = useUserStore()

const passwordInfo = reactive({
  oldPassword: "",
  newPassword: "",
  confirmNewPassword: ""
})

const formRef = ref<FormInstance>()
const formRules = reactive({
  oldPassword: [{ required: true, trigger: "blur" }],
  newPassword: [{ required: true, trigger: "blur" }],
  confirmNewPassword: [{ validator: checkConfirmNewPassword, trigger: "blur" }]
})

function checkConfirmNewPassword(rule: any, value: any, callback: any) {
  let psd = (value as string).trim()
  if (!psd) {
    return callback(new Error(`${t("label.confirmTheNewPassword")} ${t("rules.isRequired")}`))
  }
  if (psd !== passwordInfo.newPassword) {
    return callback(new Error(t("rules.formatErrConfirmNewPassword")))
  }
  callback()
}

async function updatePassword() {
  if (!(await formRef.value?.validate())) {
    return
  }
  await userService.resetPasswordByPsd({
    oldPassword: RSAEncrypt(passwordInfo.oldPassword),
    newPassword: RSAEncrypt(passwordInfo.newPassword)
  })
  userStore.clearUserInfo()
  ShowSuccessMsg(t("message.updatePasswordSucceed"))
}
</script>

<template>
  <el-form ref="formRef" :model="passwordInfo" :rules="formRules" label-position="top">
    <el-row :gutter="12">
      <el-col :span="24">
        <el-form-item :label="t('label.oldPassword')" prop="oldPassword">
          <v-password-input v-model="passwordInfo.oldPassword" />
        </el-form-item>
      </el-col>
      <el-col :span="24">
        <el-form-item :label="t('label.newPassword')" prop="newPassword">
          <v-password-input v-model="passwordInfo.newPassword" />
        </el-form-item>
      </el-col>
      <el-col :span="24">
        <el-form-item :label="t('label.confirmTheNewPassword')" prop="confirmNewPassword">
          <v-password-input v-model="passwordInfo.confirmNewPassword" />
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
  <div class="text-align-right mt-2">
    <el-button type="primary" @click.stop="updatePassword"> {{ t("btn.save") }}</el-button>
  </div>
</template>

<style lang="scss" scoped></style>
