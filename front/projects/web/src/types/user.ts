export interface UserInfo {
  userId: string
  username: string
  email: string
  password: string
  description: string
  phone: string
  role: number
  createdAt: number
  updatedAt: number
  teams: string[]
}

export function NewUserEmptyInfo(): UserInfo {
  return {
    userId: "",
    username: "",
    email: "",
    password: "",
    description: "",
    phone: "",
    role: UserRole.User,
    createdAt: 0,
    updatedAt: 0,
    teams: []
  }
}

export interface TeamInfo {
  teamId: string
  name: string
  description: string
  createdAt: number
  updatedAt: number
  users: string[]
}

export function NewTeamEmptyInfo(): TeamInfo {
  return {
    teamId: "",
    name: "",
    description: "",
    createdAt: 0,
    updatedAt: 0,
    users: []
  }
}
