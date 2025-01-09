import type { RouteRecordRaw } from "vue-router"
import { PageLimitRole } from "@/models"

export default <RouteRecordRaw[]>[
  {
    path: "",
    name: "workspace",
    redirect: "/ws/dashboard",
    component: () => import("@/views/layout/layout.vue"),
    children: [
      {
        path: "/ws/user-profile",
        name: "userProfile",
        component: () => import("@/views/user-profile/user-profile.vue"),
        meta: {}
      },
      {
        path: "/ws/dashboard",
        name: "dashboard",
        component: () => import("@/views/dashboard/dashboard.vue"),
        meta: {
          loginLimit: PageLimitRole.Login
        }
      }
    ]
  }
]
