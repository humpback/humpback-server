export interface GroupInfo {
  groupId: string
  name: string
  description: string
  createdAt: number
  updatedAt: number
  users: string[]
}

export function NewGroupEmptyInfo(): GroupInfo {
  return {
    groupId: "",
    name: "",
    description: "",
    createdAt: 0,
    updatedAt: 0,
    users: [] as Array<string>
  }
}
