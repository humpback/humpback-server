export enum PageLimitRole {
  Ignore = 0,
  Login = 1,
  Logout = -1
}

export enum UserRole {
  SupperAdmin = 1,
  Admin = 2,
  User = 3
}

export enum ConfigType {
  Static = 1,
  Volume = 2
}

export enum SortType {
  Asc = "asc",
  Desc = "desc"
}

export const PageSizeOptions = [10, 20, 30, 50, 100]

export enum Action {
  Add = "add",
  Edit = "edit",
  EditLabel = "editLabel",
  Enable = "enable",
  Disable = "Disable",
  Delete = "delete",
  View = "view",
  Start = "Start",
  ReStart = "ReStart",
  Stop = "Stop"
}

export enum NodeStatus {
  Online = "Online",
  Offline = "Offline"
}

export enum NodeSwitch {
  Enabled = "Enabled",
  Disabled = "Disabled"
}

export enum ServiceStatus {
  ServiceStatusNotReady = "NotReady",
  ServiceStatusRunning = "Running",
  ServiceStatusFailed = "Failed"
}

export enum ServiceDeployMode {
  DeployModeGlobal = "global",
  DeployModeReplicate = "replicate"
}

export enum ServiceDeployType {
  DeployTypeSchedule = "schedule",
  DeployTypeBackground = "background"
}

export enum ServicePlacementMode {
  PlacementModeLabel = "label",
  PlacementModeIP = "ip"
}

export enum ServiceNetworkMode {
  NetworkModeHost = "host",
  NetworkModeBridge = "bridge",
  NetworkModeCustom = "custom"
}

export enum ServiceNetworkProtocol {
  NetworkProtocolTCP = "TCP",
  NetworkProtocolUDP = "UDP"
}

export enum ServiceRestartPolicyMode {
  RestartPolicyModeNo = "no",
  RestartPolicyModeAlways = "always",
  RestartPolicyModeOnFail = "on-failure",
  RestartPolicyModeUnlessStopped = "unless-stopped"
}
