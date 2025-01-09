<script lang="ts" setup>
import { RuleFormatErrEmailOption, RuleFormatErrPhone, RulePleaseEnter } from "@/utils"
import { LimitEmail, LimitNotes } from "@/models"
import { FormInstance } from "element-plus"

const { t } = useI18n()
const userStore = useUserStore()

const userInfo = ref<UserInfo>(NewUserEmptyInfo())
const tableRef = useTemplateRef<FormInstance>("tableRef")

const rules = {
  username: [
    { required: true, validator: RulePleaseEnter("label.username"), trigger: "blur" },
    { required: true, validator: RuleLimitRange(LimitUserName.Min, LimitUserName.Max), trigger: "blur" }
  ],
  email: [
    { validator: RuleLimitMax(LimitEmail.Max), trigger: "blur" },
    { validator: RuleFormatErrEmailOption(), trigger: "blur" }
  ],
  phone: [
    { validator: RuleLimitMax(LimitEmail.Max), trigger: "blur" },
    { validator: RuleFormatErrPhone(), trigger: "blur" }
  ],
  description: [{ validator: RuleLimitMax(LimitNotes.Max), trigger: "blur" }]
}

async function getUserInfo() {
  return await userService.getUserInfo().then(data => {
    userInfo.value = data
    userStore.setUserInfo(data)
  })
}

async function save() {
  if (!(await tableRef.value?.validate())) {
    return
  }
  await userService.updateUserInfo({
    username: userInfo.value.username,
    email: userInfo.value.email,
    phone: userInfo.value.phone,
    description: userInfo.value.description
  })
  await getUserInfo()
  ShowSuccessMsg(t("message.saveSuccess"))
}

onMounted(async () => {
  await getUserInfo()
})
</script>

<template>
  <v-card>
    <div>
      <v-role-admin :role="userInfo.role" />
      <div class="mt-1 pl-1">
        <el-text size="small" type="info">
          {{ t("label.createDate") }}:
          <v-date-view :timestamp="userInfo.createdAt" />
        </el-text>
      </div>
    </div>
    <div class="mb-3 mt-2">
      <el-alert :closable="false" class="alert" show-icon type="info">{{ t("tips.usernameChangeTips") }}</el-alert>
    </div>
    <el-form ref="tableRef" :model="userInfo" :rules="rules" label-position="top" label-width="auto">
      <el-form-item :label="t('label.username')" prop="username">
        <v-username-input v-model="userInfo.username" />
      </el-form-item>
      <el-form-item :label="t('label.description')" prop="description">
        <v-notes-input v-model="userInfo.description" />
      </el-form-item>
      <el-form-item :label="t('label.email')" prop="email">
        <v-email-input v-model="userInfo.email" />
      </el-form-item>
      <el-form-item :label="t('label.phone')" prop="phone">
        <v-phone-input v-model="userInfo.phone" />
      </el-form-item>
      <el-form-item>
        <div class="text-align-right w-100">
          <el-button type="primary" @click="save()">{{ t("btn.save") }}</el-button>
        </div>
      </el-form-item>
    </el-form>
  </v-card>
</template>

<style lang="scss" scoped>
:deep(.alert) {
  &.el-alert {
    padding: 4px 8px;
  }

  .el-alert__icon {
    font-size: 20px;
    width: 20px;
    margin-right: 8px;
  }

  font-size: 12px;
}
</style>
