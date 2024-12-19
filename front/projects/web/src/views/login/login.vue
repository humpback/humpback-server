<script lang="ts" setup>
import type { FormInstance, FormRules } from "element-plus"

const { t } = useI18n()

const formRef = ref<FormInstance>()
const formData = reactive({
  name: "",
  password: ""
})

const formRules = reactive<FormRules>({
  name: [{ required: true, trigger: "blur" }],
  password: [{ required: true, trigger: "blur" }]
})

function login() {}
</script>

<template>
  <div class="pub-page">
    <el-card class="v-content">
      <template #header>
        <v-logo />
      </template>
      <el-form ref="formRef" :model="formData" :rules="formRules" @submit.prevent="login">
        <el-form-item prop="name">
          <v-name-input v-model="formData.name" :placeholder="t('placeholder.name')" size="large" />
        </el-form-item>
        <el-form-item prop="password">
          <v-password-input v-model="formData.password" :placeholder="t('placeholder.password')" size="large" />
        </el-form-item>
        <el-form-item>
          <el-button class="w-100" native-type="submit" size="large" type="primary">
            <el-icon>
              <icon-mdi-login-variant />
            </el-icon>
            {{ t("btn.signIn") }}
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
      <el-text size="small" type="info">©2024 Humpback. {{ t("tips.allRightsReserved") }}</el-text>
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
    padding: 0 0 40px 0;
  }

  .el-card__body {
    padding: 0;
  }

  .el-card__footer {
    padding: 40px 0 0 0;
  }
}

.copy-right {
  margin: 0 auto;
  text-align: center;
  font-size: 12px;
  font-family: Poppins, "Noto Sans SC", "Microsoft YaHei", "Ping Fang SC", Helvetica, Arial, sans-serif;
}
</style>
