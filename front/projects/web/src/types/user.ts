export interface UserInfo {
  userId: string
  username: string
  email: string
  password?: string
  description: string
  phone: string
  role: number
  createdAt: number
  updatedAt: number
  groups: string[]
}

export function NewUserEmptyInfo(): UserInfo {
  return {
    userId: "",
    username: "",
    email: "",
    password: "",
    description: "",
    phone: "",
    role: 0,
    createdAt: 0,
    updatedAt: 0,
    groups: []
  }
}
