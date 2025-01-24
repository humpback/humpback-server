import { RegistryInfo, ResponseQuery, ResponseSuccess } from "@/types"

class RegistryService {
  async info(id: string, hasAuth?: boolean) {
    const url = hasAuth ? `/webapi/registry/info/${id}?hasAuth=true` : `/webapi/registry/info/${id}`
    return await httpClient.get<RegistryInfo>(url).then(res => res.data)
  }

  async query(data: any) {
    return await httpClient.post<ResponseQuery<RegistryInfo>>("/webapi/registry/query", data).then(res => res.data)
  }

  async create(data: any) {
    return await httpClient.post<string>("/webapi/registry", data).then(res => res.data)
  }

  async update(data: any) {
    return await httpClient.put<string>("/webapi/registry", data).then(res => res.data)
  }

  async delete(id: string) {
    return await httpClient.delete<ResponseSuccess>(`/webapi/registry/${id}`).then(res => res.data)
  }
}

export const registryService = new RegistryService()
