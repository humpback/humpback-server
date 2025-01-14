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
        path: "/ws/my-account",
        name: "myAccount",
        component: () => import("@/views/my-account/my-account.vue"),
        meta: {}
      },
      {
        path: "/ws/dashboard",
        name: "dashboard",
        component: () => import("@/views/dashboard/dashboard.vue"),
        meta: {
          loginLimit: PageLimitRole.Login
        }
      },
      {
        path: "/ws/user-related/:mode",
        name: "userRelated",
        component: () => import("@/views/user-related/user-related.vue"),
        meta: {
          onlyAdmin: true,
          webTitle: {
            params: "mode"
          }
        }
      }
    ]
  }
]
