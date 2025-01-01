/// <reference types="vite/client" />
/// <reference types="element-plus" />

declare const __APP_VERSION_: string

declare module "*.vue" {
  import Vue from "vue"
  export default Vue
}

declare module "v3-easyui"

import "vue-router"
//用于定义router中的meta结构
declare module "vue-router" {
  interface RouteMeta {
    currentMenu?: string
    isNew?: boolean
    rolesLimit?: number[] // undefined：所有角色都可以进入；有值：限定用户才可以进入。 枚举： admin(99),companyAdmin(1),dispatcher(2)
    loginLimit?: number // 0|undefined： 必须登录； -1: 忽略是否登录；1: 必须登出
    webTitleParam?: string
    webTitle?: string
  }
}
