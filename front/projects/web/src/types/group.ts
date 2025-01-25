import { BaseInfo, NewBaseEmptyInfo } from "#/base.ts"

export interface GroupInfo extends BaseInfo {
  groupId: string
  name: string
  description: string
  users: string[]
}

export function NewGroupEmptyInfo(): GroupInfo {
  return {
    ...NewBaseEmptyInfo(),
    groupId: "",
    name: "",
    description: "",
    users: [] as Array<string>
  }
}
