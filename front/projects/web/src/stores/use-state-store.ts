import { cloneDeep } from "lodash-es"
import { GroupInfo } from "@/types"
import { PageGroupDetail, SessionStorageCurrentGroupId } from "@/models"

type TotalKey = PageGroupDetail.Services | PageGroupDetail.Nodes

interface StateGroupInfo extends GroupInfo {
  total: {
    [key in TotalKey]: number
  }
}

export default defineStore("state", () => {
  const route = useRoute()
  const groups = ref<{ [key: string]: StateGroupInfo }>({})
  const services = ref<{ [key: string]: ServiceInfo }>({})
  const nodes = ref<{ [key: string]: NodeInfo }>({})

  function setGroup(id: string, groupInfo: GroupInfo) {
    sessionStorage.setItem(SessionStorageCurrentGroupId, id)
    const info = groups.value[id]
    const total = { services: info?.total.services || 0, nodes: groupInfo.nodes.length }
    groups.value[id] = { ...cloneDeep(groupInfo), total: total } as StateGroupInfo
  }

  function setGroupTotal(id?: string, serviceTotal?: number, nodeTotal?: number) {
    const key = id || (route.params["groupId"] as string)
    const info = cloneDeep(groups.value[key])
    if (info) {
      info.total = {
        services: typeof serviceTotal === "undefined" ? info.total.services : serviceTotal,
        nodes: typeof nodeTotal === "undefined" ? info.total.nodes : nodeTotal
      }
      groups.value[key] = info
    }
  }

  function getGroup(id?: string): StateGroupInfo | undefined {
    return groups.value[id || (route.params["groupId"] as string)]
  }

  function setService(id: string, serviceInfo: ServiceInfo) {
    services.value[id] = cloneDeep(serviceInfo)
  }

  function getService(id?: string): ServiceInfo | undefined {
    return services.value[id || (route.params["serviceId"] as string)]
  }

  function setNode(id: string, nodeInfo: NodeInfo) {
    nodes.value[id] = cloneDeep(nodeInfo)
  }

  function getNode(id?: string): NodeInfo | undefined {
    return nodes.value[id || (route.params["nodeId"] as string)]
  }

  return {
    setGroup,
    setGroupTotal,
    getGroup,
    setService,
    getService,
    setNode,
    getNode
  }
})
