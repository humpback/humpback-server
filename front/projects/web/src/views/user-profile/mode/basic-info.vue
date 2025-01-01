<script lang="ts" setup>
import type { FormInstance } from "element-plus"
import { cloneDeep } from "lodash"
import { TimestampToTime } from "@/utils"

const { t } = useI18n()
const userStore = useUserStore()

const userFormRef = ref<FormInstance>()
const userFormInfo = reactive({
  name: "",
  email: "",
  phone: "",
  notes: "",
  companyName: "",
  enterpriseCode: ""
})

const userFormRules = reactive({
  name: [{ required: true, trigger: "blur" }],
  companyName: [{ required: true, trigger: "blur" }],
  email: [{ required: true, trigger: "blur" }],
  phone: [{ required: true, trigger: "blur" }],
  enterpriseCode: [{ required: true, trigger: "blur" }]
})

async function getUserInfo() {
  const userInfo = await userService.getUserInfo()
  userStore.setUserInfo(cloneDeep(userInfo))
  userFormInfo.name = userInfo.name
}

async function updateBasicInfo() {
  if (!(await userFormRef.value?.validate())) {
    return
  }
  const body: any = {
    name: userFormInfo.name,
    phone: userFormInfo.phone
  }

  userService.updateUserInfo(body).then(() => {
    ShowSuccessMsg(t("message.userInfoUpdateSucceed"))
    getUserInfo()
  })
}

onMounted(() => {
  getUserInfo()
})
</script>

<template>
  <div class="d-flex flex-wrap gap-3 mb-5">
    <el-text size="small"> {{ t("label.createDate") }}: {{ TimestampToTime(userStore.userInfo.createdUserInfo?.timeAt, 4) }}</el-text>
  </div>

  <el-form ref="userFormRef" :model="userFormInfo" :rules="userFormRules" label-position="top">
    <el-row :gutter="12">
      <el-col :span="24">
        <el-form-item :label="t('label.userName')" prop="name">
          <v-name-input v-model="userFormInfo.name" />
        </el-form-item>
      </el-col>
      <el-col :span="24">
        <el-form-item :label="t('label.introduction')" prop="notes">
          <v-notes-input v-model="userFormInfo.notes" :placeholder="t('placeholder.enterIntroduction')"></v-notes-input>
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>

  <div class="text-align-right mt-2">
    <el-button type="primary" @click.stop="updateBasicInfo()"> {{ t("btn.save") }}</el-button>
  </div>
</template>

<style lang="scss" scoped>
.help :deep(.el-form-item__label) {
  display: flex;
  align-items: center;
  justify-content: left;
}
</style>
