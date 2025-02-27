import { BaseInfo, NewBaseEmptyInfo } from "#/base.ts"
import { ServiceVolumeType } from "@/models"

export interface ServiceInfo extends BaseInfo {
  serviceId: string
  groupId: string
  serviceName: string
  description: string
  version: string
  action: string
  isEnabled: boolean
  isDelete: boolean
  status: string
  meta: ServiceMetaDockerInfo
  deployment: ServiceDeploymentInfo
  containers: ServiceContainerStatusInfo[]
}

export interface ServiceMetaDockerInfo {
  image: string
  alwaysPull: boolean
  command: string
  env: string[]
  labels: { [key: string]: string }
  privileged: boolean
  capabilities: string[]
  volumes: ServiceVolumeInfo[]
  network: ServiceNetworkInfo
  restartPolicy: ServiceRestartPolicyInfo
}

export interface ServiceVolumeInfo {
  type: ServiceVolumeType.VolumeTypeBind | ServiceVolumeType.VolumeTypeVolume
  target: string
  source: string
  "readonly": boolean
}

export interface ServiceNetworkInfo {
  mode: string
  hostname: string
  networkName: string
  useMachineHostname: boolean
  ports: ServicePortInfo[]
}

export interface ServicePortInfo {
  hostPort: number
  containerPort: number
  protocol: string
}

export interface ServiceRestartPolicyInfo {
  mode: string
  maxRetryCount: number
}

export interface ServiceDeploymentInfo {
  type: string
  mode: string
  replicas: number
  placements: ServicePlacementInfo[]
  schedule: ServiceScheduleInfo
}

export interface ServicePlacementInfo {
  mode: string
  key: string
  value: string
  isEqual: string
}

export interface ServiceScheduleInfo {
  timeout: string
  rules: string[]
}

export interface ServiceContainerStatusInfo {
  containerId: string
  containerName: string
  nodeId: string
  status: string
  statusInfo: string
  errorMsg: string
  image: string
  command: string
  network: string
  created: number
  started: number
  lastHeartbeat: number
}

export function NewServiceEmptyInfo(): ServiceInfo {
  return {
    ...NewBaseEmptyInfo(),
    serviceId: "",
    groupId: "",
    serviceName: "",
    description: "",
    version: "",
    action: "",
    isEnabled: false,
    isDelete: false,
    status: "",
    meta: NewServiceMetaDockerEmptyInfo(),
    deployment: NewServiceDeploymentInfo(),
    containers: []
  }
}

export function NewServiceMetaDockerEmptyInfo(): ServiceMetaDockerInfo {
  return {
    image: "",
    alwaysPull: false,
    command: "",
    env: [],
    labels: {},
    volumes: [],
    privileged: false,
    capabilities: [],
    network: {
      mode: ServiceNetworkMode.NetworkModeHost,
      hostname: "",
      networkName: "",
      useMachineHostname: false,
      ports: []
    },
    restartPolicy: {
      mode: ServiceRestartPolicyMode.RestartPolicyModeNo,
      maxRetryCount: 0
    }
  }
}

export function NewServiceDeploymentInfo(): ServiceDeploymentInfo {
  return {
    type: "",
    mode: "",
    replicas: 0,
    placements: [],
    schedule: {
      timeout: "",
      rules: []
    }
  }
}

//
// type DeployMode string
// type DeployType string
//
// var (
// 	DeployModeGlobal    DeployMode = "global"
// 	DeployModeReplicate DeployMode = "replicate"
// )
//
// var (
// 	DeployTypeSchedule   DeployType = "schedule"
// 	DeployTypeBackground DeployType = "background"
// )
//
// type PlacementMode string
//
// var (
// 	PlacementModeLabel PlacementMode = "label"
// 	PlacementModeIP    PlacementMode = "ip"
// )

// type NetworkMode string
//
// var (
// 	NetworkModeHost   NetworkMode = "host"
// 	NetworkModeBridge NetworkMode = "bridge"
// 	NetworkModeCustom NetworkMode = "custom"
// )

//
//   type Service struct {
// 	Containers  []*ContainerStatus `json:"containers"`
// }
