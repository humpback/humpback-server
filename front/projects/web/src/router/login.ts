import type { RouteRecordRaw } from "vue-router"
import login from "@/views/login/login.vue"

export default <RouteRecordRaw[]>[
  {
    path: "/login",
    name: "login",
    component: login,
    meta: {}
  }
]
