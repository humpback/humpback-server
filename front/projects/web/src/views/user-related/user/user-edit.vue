<script lang="ts" setup>
import { cloneDeep, find } from "lodash-es"
import { FormInstance, FormRules } from "element-plus"
import { RuleFormatErrEmailOption, RuleFormatErrPhone, RulePleaseEnter } from "@/utils"
import { LimitDescription, LimitEmail, LimitPassword } from "@/models"
import { TeamInfo, UserInfo } from "@/types"

const emits = defineEmits<{
  (e: "refresh"): void
}>()

const { t } = useI18n()
const userStore = useUserStore()

const dialogInfo = ref({
  show: false,
  info: {} as UserInfo
})

const teamsOptions = ref<TeamInfo[]>([])

const formRef = useTemplateRef<FormInstance>("formRef")
const rules = ref<FormRules>({
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
  password: [
    { required: true, validator: RulePleaseEnter("label.password"), trigger: "blur" },
    { validator: RuleLimitRange(LimitPassword.Min, LimitPassword.Max), trigger: "blur" }
  ],
  role: [{ required: true, validator: checkRole, trigger: "change" }],
  description: [{ validator: RuleLimitMax(LimitDescription.Max), trigger: "blur" }]
})

function checkRole(rule: any, value: any, callback: any) {
  const options = userStore.isSupperAdmin ? [UserRole.User, UserRole.Admin] : [UserRole.User]
  if (!find(options, x => x === value)) {
    return callback(new Error(t("rules.invalidRole")))
  }
  callback()
}

function open(info?: UserInfo) {
  dialogInfo.value.info = info ? cloneDeep(info) : NewUserEmptyInfo()
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
    <template #header>
      {{ dialogInfo.info.userId ? t("header.editUser") : t("header.addUser") }}
    </template>
    <div class="my-3">
      <el-form ref="formRef" :model="dialogInfo.info" :rules="rules" label-position="top" label-width="auto">
        <el-row :gutter="12">
          <el-col>
            <el-form-item :label="t('label.username')" prop="username">
              <v-username-input v-model="dialogInfo.info.username" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item :label="t('label.role')" prop="role">
              <v-role-select
                v-model="dialogInfo.info.role"
                :clearable="false"
                :only-show-roles="userStore.isSupperAdmin ? [UserRole.User, UserRole.Admin] : [UserRole.User]" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item :label="t('label.password')" prop="password">
              <v-password-input v-model="dialogInfo.info.password" />
            </el-form-item>
          </el-col>
          <el-col>
            <el-form-item :label="t('label.description')" prop="description">
              <v-description-input v-model="dialogInfo.info.description" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item :label="t('label.email')" prop="email">
              <v-email-input v-model="dialogInfo.info.email" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item :label="t('label.phone')" prop="phone">
              <v-phone-input v-model="dialogInfo.info.phone" />
            </el-form-item>
          </el-col>

          <el-col>
            <el-form-item :label="t('label.teams')" prop="teams">
              <v-team-use-option-select v-model="dialogInfo.info.teams" :options="teamsOptions" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>
    <template #footer>
      <el-button @click="dialogInfo.show = false">{{ t("btn.cancel") }}</el-button>
      <el-button type="primary" @click="save">{{ t("btn.save") }}</el-button>
    </template>
  </v-dialog>
</template>

<style lang="scss" scoped></style>
