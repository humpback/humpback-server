import { includes, toLower } from "lodash-es"
import { UserRole } from "@/models"

export function TableHeight(invalidHeight: number, minHeight = 500) {
  const pageStore = usePageStore()
  const height = pageStore.screenHeight - invalidHeight
  if (height < minHeight) {
    return "auto"
  }
  return height
}

export function IncludesIgnoreCase(str?: string, subStr?: string): boolean {
  if (!str) {
    return false
  }
  return includes(toLower(str), toLower(subStr || ""))
}

export function IsSupperAdmin(role: number) {
  return role === UserRole.SupperAdmin
}

export function IsAdmin(role: number) {
  return role === UserRole.Admin
}

export function IsUser(role: number) {
  return role === UserRole.User
}

export function GetUserRole(role: number) {
  switch (role) {
    case UserRole.SupperAdmin:
      return UserRole.SupperAdmin
    case UserRole.Admin:
      return UserRole.Admin
    default:
      return UserRole.User
  }
}
