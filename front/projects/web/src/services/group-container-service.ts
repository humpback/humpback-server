import { ContainerPerformance } from "@/types"

class GroupContainerService {
  async operate(groupId: string, data: { nodeId: string; containerId: string; action: "Start" | "Stop" | "Restart" }) {
    return await httpClient.put<any>(`/webapi/group/${groupId}/container/operate`, data).then(res => res.data)
  }

  async performance(groupId: string, data: Array<{ nodeId: string; containerId: string }>) {
    return await httpClient.post<ContainerPerformance[]>(`/webapi/group/${groupId}/container/performance`, data).then(res => res.data)
  }

  async logs(groupId: string, data: { nodeId: string; containerId: string; line?: number; startAt?: number; endAt?: number; showTimestamp?: boolean }) {
    return await httpClient.post<string[]>(`/webapi/group/${groupId}/container/logs`, data).then(res => res.data)
  }
}

export const groupContainerService = new GroupContainerService()
