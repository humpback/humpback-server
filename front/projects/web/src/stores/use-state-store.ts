export default defineStore("group", () => {
  const groups = ref<Map<string, GroupInfo>>(new Map<string, GroupInfo>())
  const nodes = ref<Map<string, NodeInfo>>(new Map<string, NodeInfo>())

  function setGroup(id: string, groupInfo: GroupInfo) {
    groups.value.set(id, groupInfo)
  }

  function getGroup(id: string): GroupInfo | undefined {
    return groups.value.get(id)
  }

  function setNode(id: string, nodeInfo: NodeInfo) {
    nodes.value.set(id, nodeInfo)
  }

  function getNode(id: string): NodeInfo | undefined {
    return nodes.value.get(id)
  }

  return {
    setGroup,
    setNode,
    getNode,
    getGroup
  }
})
