export interface UserInfo {
  userId: string
  userName: string
  email: string
  password: string
  phone: string
  isAdmin: boolean
  createdAt: number
  updatedAt: number
  groups: string[]
}

export function NewUserEmptyInfo(): UserInfo {
  return {
    userId: "",
    userName: "",
    email: "",
    password: "",
    phone: "",
    isAdmin: false,
    createdAt: 0,
    updatedAt: 0,
    groups: []
  }
}
