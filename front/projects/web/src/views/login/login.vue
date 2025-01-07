<script lang="ts" setup>
import type { FormInstance, FormRules } from "element-plus"
import { ChangeEventType, IsValidEmail, RulePleaseEnter, SendChannelMessage } from "@/utils"
import { RSAEncrypt } from "utils/rsa.ts"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const isLogin = ref(false)
const formRef = useTemplateRef<FormInstance>("formRef")
const keyMax = computed(() => Math.max(LimitEmail.Max, LimitUserName.Max))
const formData = reactive({
  name: "",
  password: ""
})

const formRules = reactive<FormRules>({
  name: [
    { required: true, validator: RulePleaseEnter("placeholder.nameOrEmail"), trigger: "blur" },
    { validator: checkName, trigger: "blur" }
  ],
  password: [{ required: true, validator: RulePleaseEnter("placeholder.password"), trigger: "blur" }]
})

const disabled = computed(() => formData.name.length < LimitUserName.Min || formData.password.length < LimitPassword.Min)

function checkName(rule: any, value: any, callback: any) {
  const str = value as string
  if (str.includes("@")) {
    if (!IsValidEmail(str)) {
      return callback(new Error("rules.formatErrEmail"))
    }
  }
  callback()
}

async function login() {
  if (!(await formRef.value?.validate())) {
    return
  }
  isLogin.value = true
  const body = {
    name: RSAEncrypt(formData.name),
    password: RSAEncrypt(formData.password)
  }

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
    .finally(() => (isLogin.value = false))
}
</script>

<template>
  <div class="pub-page">
    <el-card class="v-content">
      <template #header>
        <v-logo />
      </template>
      <el-form ref="formRef" :model="formData" :rules="formRules" @submit.prevent="login">
        <el-form-item prop="name">
          <v-input v-model="formData.name" :maxlength="keyMax" :placeholder="t('placeholder.nameOrEmail')" size="large" />
        </el-form-item>
        <el-form-item prop="password">
          <v-password-input v-model="formData.password" :placeholder="t('placeholder.password')" size="large" />
        </el-form-item>
        <el-form-item>
          <el-button :disabled="disabled || isLogin" class="w-100 mt-3" native-type="submit" size="large" type="primary">
            <el-icon>
              <icon-mdi-login-variant />
            </el-icon>
            {{ isLogin ? t("logging") : t("btn.login") }}
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
