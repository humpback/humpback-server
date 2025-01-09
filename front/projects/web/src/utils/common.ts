import { includes, toLower } from "lodash"
import { UserRole } from "@/models"

export function TableHeight(invalidHeight: number, minHeight = 500): number {
  const pageStore = usePageStore()
  return Math.max(pageStore.screenHeight - invalidHeight, minHeight)
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

export function IsNormal(role: number) {
  return role === UserRole.Normal
}

export function GetUserRole(role: number) {
  switch (role) {
    case UserRole.SupperAdmin:
      return UserRole.SupperAdmin
    case UserRole.Admin:
      return UserRole.Admin
    default:
      return UserRole.Normal
  }
}
