import { UserInfo } from "#/index.ts"

class UserService {
  async login(data: any) {
    return await httpClient.post<UserInfo>("/webapi/user/login", data).then(res => res.data)
  }

  async logout() {
    return await httpClient.post("/webapi/user/logout", null, { disableErrMsg: true }).then(res => res.data)
  }

  async changePassword(data: any) {
    return await httpClient.put("/webapi/user/me/change-psd", data).then(res => res.data)
  }

  async getUserInfo(startup?: boolean) {
    return await httpClient.get<UserInfo>("/webapi/user/me", { params: startup ? { startup: "true" } : undefined }).then(res => res.data)
  }

  async updateUserInfo(data: any) {
    return await httpClient.put("/webapi/user/me", data).then(res => res.data)
  }
}

export const userService = new UserService()
