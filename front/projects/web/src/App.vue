<script lang="ts" setup>
import { ChangeEventType, CloseChannelMessage, GetChannelMessage } from "utils/index.ts"
import { GetUILocale } from "@/locales/index.ts"

const { t } = useI18n()
const pageStore = usePageStore()
const userStore = useUserStore()

const handleChannelMessage = (data: any) => {
  switch (data?.type) {
    case ChangeEventType.Login:
      if (!data.value || data.value?.userId != userStore.userInfo.userId) {
        pageStore.refreshPage.needRefresh = true
        pageStore.refreshPage.type = ChangeEventType.Login
      }
      return
    case ChangeEventType.Logout:
      pageStore.refreshPage.needRefresh = true
      pageStore.refreshPage.type = ChangeEventType.Logout
  }
}

const handleFocus = () => {
  if (pageStore.refreshPage.needRefresh) {
    if (pageStore.refreshPage.type === ChangeEventType.Login) {
      ShowWarningMsg(t("message.loginUserChangeEvent"))
    }
    location.reload()
  }
}

onMounted(() => {
  pageStore.setScreen()
  GetChannelMessage(handleChannelMessage)
  window.addEventListener("resize", pageStore.setScreen)
  window.addEventListener("focus", handleFocus)
})

onBeforeUnmount(() => {
  CloseChannelMessage()
  window.removeEventListener("resize", pageStore.setScreen)
  window.removeEventListener("focus", handleFocus)
})
</script>

<template>
  <el-config-provider :locale="GetUILocale('elementPlus')">
    <router-view />
  </el-config-provider>
</template>

<style scoped></style>
