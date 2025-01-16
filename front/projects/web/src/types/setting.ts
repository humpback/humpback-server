import { ConfigType } from "@/models"

export interface ConfigInfo {
  configId: string
  configName: string
  description: string
  configValue: any
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
