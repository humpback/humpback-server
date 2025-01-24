class CommonService {
  async config() {
    return await httpClient.get<{ [key: string]: any }>("/webapi/common/config").then(res => res.data)
  }
}

export const commonService = new CommonService()
