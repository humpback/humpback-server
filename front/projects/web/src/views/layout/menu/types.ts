import { shallowRef } from "vue"

export interface MenuInfo {
  icon: any
  name: string
  params?: any
  query?: any
  path?: string
  children?: MenuInfo[]
  rolesLimit?: number[]
}

export const menuI18nPrefix = "menu"

export const menuList: MenuInfo[] = [
  {
    icon: shallowRef(IconMdiViewDashboard),
    name: "dashboard"
  },
  {
    icon: shallowRef(IconMdiCompany),
    name: "serviceManagement"
  },
  {
    icon: shallowRef(IconMdiCogOutline),
    name: "administration",
    rolesLimit: [UserRole.SupperAdmin, UserRole.Admin],
    children: [
      {
        icon: shallowRef(IconMdiAlphaCBoxOutline),
        name: "registries"
      },
      {
        icon: shallowRef(IconMdiTextBoxOutline),
        name: "nodes"
      },
      {
        icon: shallowRef(IconMdiTextBoxOutline),
        name: "configs"
      },
      {
        icon: shallowRef(IconMdiAccount),
        name: "userRelated"
      }
    ]
  }
]
