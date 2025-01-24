import { ConfigType } from "@/models"

export interface ConfigInfo {
  configId: string
  configName: string
  description: string
  configValue: string
  configType: number
  createdAt: number
  updatedAt: number
}

export function NewConfigEmptyInfo(): ConfigInfo {
  return {
    configId: "",
    configName: "",
    description: "",
    configValue: "",
    configType: ConfigType.Static,
    createdAt: 0,
    updatedAt: 0
  }
}

export interface RegistryInfo {
  registryId: string
  registryName: string
  url: string
  isDefault: boolean
  username: string
  password: string
  createdAt: number
  updatedAt: number
  hasAuth?: boolean
}

export function NewRegistryEmptyInfo(): RegistryInfo {
  return {
    registryId: "",
    registryName: "",
    url: "",
    isDefault: false,
    username: "",
    password: "",
    createdAt: 0,
    updatedAt: 0
  }
}
