import { NewUserEmptyInfo, UserInfo } from "#/index.ts"
import { GetUserRole, IsAdmin, IsNormal, IsSupperAdmin } from "@/utils"

export default defineStore("user", () => {
  const userInfo = ref<UserInfo>(NewUserEmptyInfo())

  const isLogged = computed(() => !!userInfo.value.userId)
  const userRole = computed(() => GetUserRole(userInfo.value.role))
  const isAdmin = computed(() => IsAdmin(userInfo.value.role))
  const isSupperAdmin = computed(() => IsSupperAdmin(userInfo.value.role))
  const isNormal = computed(() => IsNormal(userInfo.value.role))

  const init = async () => {
    return await userService.getUserInfo(true).then(data => {
      userInfo.value = data
    })
  }

  function setUserInfo(info: UserInfo) {
    userInfo.value = info
  }

  function clearUserInfo() {
    userInfo.value = NewUserEmptyInfo()
  }

  return {
    userInfo,
    userRole,
    isAdmin,
    isSupperAdmin,
    isNormal,
    isLogged,
    init,
    setUserInfo,
    clearUserInfo
  }
})
