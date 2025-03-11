class ContainerService {
  async operate(groupId: string, data: { nodeId: string; containerId: string; action: "Start" | "Stop" | "Restart" }) {
    return await httpClient.put<any>(`/webapi/group/${groupId}/container/operate`, data).then(res => res.data)
  }
}

export const containerService = new ContainerService()
