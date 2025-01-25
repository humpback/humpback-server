import { ConfigType } from "@/models"
import { BaseInfo, NewBaseEmptyInfo } from "#/base.ts"

export interface ConfigInfo extends BaseInfo {
  configId: string
  configName: string
  description: string
  configValue: string
  configType: number
}

export function NewConfigEmptyInfo(): ConfigInfo {
  return {
    ...NewBaseEmptyInfo(),
    configId: "",
    configName: "",
    description: "",
    configValue: "",
    configType: ConfigType.Static
  }
}

export interface RegistryInfo extends BaseInfo {
  registryId: string
  registryName: string
  url: string
  isDefault: boolean
  username: string
  password: string
  hasAuth?: boolean
}

export function NewRegistryEmptyInfo(): RegistryInfo {
  return {
    ...NewBaseEmptyInfo(),
    registryId: "",
    registryName: "",
    url: "",
    isDefault: false,
    username: "",
    password: ""
  }
}

export interface NodeInfo extends BaseInfo {
  nodeId: string
  name: string
  ipAddress: string
  port: number
  status: string
  isEnable: boolean
  cpuUsage: number
  memoryUsage: number
  labels: { [key: string]: string }
}

export function NewNodeEmptyInfo(): NodeInfo {
  return {
    ...NewBaseEmptyInfo(),
    nodeId: "",
    name: "",
    ipAddress: "",
    port: 0,
    status: "",
    isEnable: false,
    cpuUsage: 0,
    memoryUsage: 0,
    labels: {}
  }
}
