<script lang="ts" setup>
import type { FormInstance, FormRules } from "element-plus"
import { ChangeEventType, globalLoading, RulePleaseEnter, SendChannelMessage } from "@/utils"
import { RSAEncrypt } from "utils/rsa.ts"
import VUsernameInput from "@/components/business/v-name/VUsernameInput.vue"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const formRef = useTemplateRef<FormInstance>("formRef")
const formData = reactive({
  username: "",
  password: ""
})

const formRules = reactive<FormRules>({
  username: [{ required: true, validator: RulePleaseEnter("placeholder.username"), trigger: "blur" }],
  password: [{ required: true, validator: RulePleaseEnter("placeholder.password"), trigger: "blur" }]
})

async function login() {
  if (!(await formRef.value?.validate())) {
    return
  }
  const body = {
    username: RSAEncrypt(formData.username),
    password: RSAEncrypt(formData.password)
  }
  globalLoading.show(t("message.loggingIn"))
  userService
    .login(body)
    .then(data => {
      userStore.setUserInfo(data)
      SendChannelMessage(ChangeEventType.Login, data)
      ShowSuccessMsg(t("message.loginSuccess"))
      if (route.query?.redirectUrl) {
        router.push(route.query.redirectUrl as string)
        return
      }
      router.push({ name: "workspace" })
    })
    .finally(() => {
      globalLoading.close()
    })
}
</script>

<template>
  <div class="pub-page">
    <el-card class="v-content">
      <template #header>
        <v-logo />
      </template>
      <el-form ref="formRef" :model="formData" :rules="formRules" @submit.prevent="login()">
        <el-form-item prop="username">
          <v-username-input v-model="formData.username" :clearable="false" :placeholder="t('placeholder.username')" :show-word-limit="false" size="large" />
        </el-form-item>
        <el-form-item prop="password">
          <v-password-input v-model="formData.password" :placeholder="t('placeholder.password')" size="large" />
        </el-form-item>
        <el-form-item>
          <el-button class="w-100 mt-3" native-type="submit" size="large" type="primary">
            <el-icon>
              <icon-mdi-login-variant />
            </el-icon>
            {{ t("btn.login") }}
          </el-button>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="text-align-right">
          <Language />
        </div>
      </template>
    </el-card>
    <div class="copy-right">
      <el-text size="small" type="info">{{ t("tips.allRightsReserved") }}</el-text>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.pub-page {
  box-sizing: border-box;
  min-height: calc(100vh);
  width: 100%;
  padding-top: 16vh;
  background-image: linear-gradient(-20deg, #fffeff 50%, #d7fffe 100%);

  :deep(.el-card__footer) {
    border: none;
  }
}

:deep(.v-content) {
  box-sizing: border-box;
  box-shadow: var(--el-box-shadow-lighter);
  border: none;
  margin: 0 auto 48px auto;
  padding: 24px;
  width: 400px;
  border-radius: 8px;
  background-color: rgba(0, 0, 0, 0);
  background-color: #ffffff;

  .el-card__header {
    border: none;
    padding: 0 0 30px 0;
  }

  .el-card__body {
    padding: 0;
  }

  .el-card__footer {
    padding: 30px 0 0 0;
  }
}

.copy-right {
  margin: 0 auto;
  text-align: center;
  font-size: 12px;
  font-family: Poppins, "Noto Sans SC", "Microsoft YaHei", "Ping Fang SC", Helvetica, Arial, sans-serif;
}
</style>
