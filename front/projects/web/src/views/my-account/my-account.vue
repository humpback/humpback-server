<script lang="ts" setup>
import ChangePassword from "./change-password.vue"
import { RuleFormatErrEmailOption, RuleFormatErrPhone, RulePleaseEnter } from "@/utils"
import { RuleLength } from "@/models"
import { FormInstance, FormRules } from "element-plus"
import { cloneDeep } from "lodash-es"
import { QueryInfo } from "@/types"

const { t } = useI18n()
const userStore = useUserStore()

const loading = ref(false)
const isLoadingActivities = ref(false)
const userInfo = ref<UserInfo>(NewUserEmptyInfo())

const queryUserActivityInfo = ref<QueryInfo>(new QueryInfo({}, ["keywords"], { size: 10, index: 1 }, { field: "", order: "desc" }))
const activities = ref({
  total: 0,
  data: [] as Array<any>
})

const tableRef = useTemplateRef<FormInstance>("tableRef")

const rules = ref<FormRules>({
  username: [
    { required: true, validator: RulePleaseEnter("label.username"), trigger: "blur" },
    { required: true, validator: RuleLimitRange(RuleLength.Username.Min, RuleLength.Username.Max), trigger: "blur" }
  ],
  email: [
    { validator: RuleLimitMax(RuleLength.Email.Max), trigger: "blur" },
    { validator: RuleFormatErrEmailOption(), trigger: "blur" }
  ],
  phone: [
    { validator: RuleLimitMax(RuleLength.Email.Max), trigger: "blur" },
    { validator: RuleFormatErrPhone(), trigger: "blur" }
  ],
  description: [{ validator: RuleLimitMax(RuleLength.Description.Max), trigger: "blur" }]
})

async function getUserInfo() {
  loading.value = true
  return await userService
    .getMe()
    .then(data => {
      userInfo.value = data
      userStore.setUserInfo(cloneDeep(data))
    })
    .finally(() => {
      loading.value = false
    })
}

async function getUserActivities() {}

async function save() {
  if (!(await tableRef.value?.validate())) {
    return
  }
  await userService.updateMeInfo({
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
  <div>
    <v-card v-loading="loading">
      <div>
        <v-role-admin :role="userInfo.role" size="default" />
        <div class="d-flex gap-1 mt-2 pl-1 mb-3">
          <el-text type="info">
            {{ t("label.createDate") }}:
            <v-date-view :timestamp="userInfo.createdAt" />
          </el-text>
          <el-divider direction="vertical" />
          <div>
            <change-password />
          </div>
        </div>
      </div>

      <div class="mb-5 mt-2">
        <v-alert> {{ t("tips.usernameChangeTips") }}</v-alert>
      </div>
      <el-form ref="tableRef" :model="userInfo" :rules="rules" label-position="top" label-width="auto">
        <el-row :gutter="12">
          <el-col>
            <el-form-item :label="t('label.username')" prop="username">
              <v-username-input v-model="userInfo.username" />
            </el-form-item>
          </el-col>
          <el-col>
            <el-form-item :label="t('label.description')" prop="description">
              <v-description-input v-model="userInfo.description" />
            </el-form-item>
          </el-col>
          <el-col :md="12" :span="24">
            <el-form-item :label="t('label.email')" prop="email">
              <v-email-input v-model="userInfo.email" />
            </el-form-item>
          </el-col>
          <el-col :md="12" :span="24">
            <el-form-item :label="t('label.phone')" prop="phone">
              <v-phone-input v-model="userInfo.phone" />
            </el-form-item>
          </el-col>
          <el-col>
            <el-form-item>
              <div class="text-align-right w-100">
                <el-button type="primary" @click="save()">{{ t("btn.save") }}</el-button>
              </div>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </v-card>

    <v-card v-loading="isLoadingActivities" class="mt-5">
      <div class="d-flex gap-1">
        <div class="f-bold">
          {{ t("label.activities") }}
        </div>
        <el-button :title="t('label.refresh')" link type="primary">
          <el-icon :size="20">
            <IconMdiRefresh />
          </el-icon>
        </el-button>
      </div>
      <v-table
        :data="activities.data"
        :page-info="queryUserActivityInfo.pageInfo"
        :show-header="false"
        :total="activities.total"
        @page-change="getUserActivities()">
        <el-table-column>
          <template #default></template>
        </el-table-column>
      </v-table>
    </v-card>
  </div>
</template>

<style lang="scss" scoped></style>
