export enum PageLimitRole {
  Ingore = 0,
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
  Delete = "delete",
  View = "view"
}

export enum NodeStatus {
  Online = "Online",
  Offline = "Offline"
}
