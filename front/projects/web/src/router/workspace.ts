import type { RouteRecordRaw } from "vue-router"

export default <RouteRecordRaw[]>[
  {
    path: "",
    name: "workspace",
    redirect: "/ws/dashboard",
    component: () => import("@/views/layout/layout.vue"),
    children: [
      {
        path: "/ws/user-profile/:mode",
        name: "userProfile",
        component: () => import("@/views/user-profile/user-profile.vue"),
        meta: {}
      },
      {
        path: "/ws/dashboard",
        name: "dashboard",
        component: () => import("@/views/dashboard/dashboard.vue"),
        meta: {}
      }
    ]
  }
]
