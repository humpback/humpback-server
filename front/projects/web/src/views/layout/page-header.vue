<script lang="ts" setup>
import { ChangeEventType, SendChannelMessage } from "@/utils"

enum Menu {
  Logout = "logout",
  UserProfile = "userProfile",
  Help = "help"
}

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const pageStore = usePageStore()
const userStore = useUserStore()

const userMenu = computed<
  {
    label: string
    value: string
    icon: any
    color: string
    component?: any
    hide?: boolean
  }[]
>(() => {
  return [
    { label: "menu.userProfile", value: Menu.UserProfile, icon: IconMdiAccount, color: "green" },
    { label: "menu.logout", value: Menu.Logout, icon: IconMdiLogoutVariant, color: "var(--el-color-info)" },
    {
      label: "menu.help",
      value: Menu.Help,
      icon: IconMdiHelpCircle,
      color: "var(--el-color-default)",
      hide: !pageStore.isSmallScreen
    }
  ]
})

function handleUserMenuClick(v: string) {
  switch (v) {
    case Menu.Logout:
      userService.logout().finally(() => {
        userStore.clearUserInfo()
        router.push({ name: "login", query: { redirectURL: route.fullPath } })
        SendChannelMessage(ChangeEventType.Logout)
      })
      return
    case Menu.UserProfile:
      router.push({ name: "userProfile" })
  }
}
</script>

<template>
  <div class="header-box">
    <div class="d-flex">
      <strong> {{ t("menu.header." + (route.name as string)) }} </strong>
    </div>

    <div class="d-flex">
      <div class="d-flex gap-5 mr-5">
        <el-button v-if="!pageStore.isSmallScreen" link> {{ t("btn.help") }}</el-button>
        <Language show-type="icon" trigger="hover" />
        <v-role-admin v-if="!pageStore.isSmallScreen" :role="userStore.userInfo.role" />
      </div>

      <el-dropdown :show-timeout="0" placement="bottom-end" trigger="hover" @command="handleUserMenuClick">
        <el-button class="user-btn" link>
          <el-icon :size="20">
            <IconMdiUserCircleOutline />
          </el-icon>
          <span v-if="userStore.userInfo.username" class="username overflow_div"> {{ userStore.userInfo.username }}</span>
          <el-icon :size="20">
            <IconMdiChevronDown />
          </el-icon>
        </el-button>

        <template #dropdown>
          <el-dropdown-menu>
            <template v-for="item in userMenu" :key="item.value">
              <el-dropdown-item v-if="!item.hide" :command="item.value" :style="{ color: item.color }">
                <el-icon :size="20">
                  <component :is="item.icon" />
                </el-icon>
                <component :is="item.component" v-if="item.component"></component>
                <span v-else>
                  {{ t(item.label) }}
                </span>
              </el-dropdown-item>
            </template>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.header-box {
  height: 100%;
  width: 100%;
  box-sizing: border-box;
  padding: 0 10px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  color: inherit;

  .user-btn {
    padding: 0;
    margin-right: 4px;

    &:hover {
      border: 1px solid var(--el-border-color-darker);
      border-radius: 9999px;
    }
  }

  .username {
    max-width: 200px;
    line-height: 24px;
  }
}
</style>
