import { NewUserEmptyInfo, UserInfo } from "#/index.ts"

export default defineStore("user", () => {
  const userInfo = ref<UserInfo>(NewUserEmptyInfo())

  const isLogged = computed(() => !!userInfo.value.userId)

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
    isLogged,
    init,
    setUserInfo,
    clearUserInfo
  }
})
