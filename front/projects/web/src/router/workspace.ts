import type { RouteRecordRaw } from "vue-router"
import { PageLimitRole } from "@/models"

const administrator = <RouteRecordRaw[]>[
  {
    path: "/ws/registries",
    name: "registries",
    component: () => import("@/views/administration/registries/registries.vue"),
    meta: {
      onlyAdmin: true
    }
  },
  {
    path: "/ws/nodes",
    name: "nodes",
    component: () => import("@/views/administration/nodes/nodes.vue"),
    meta: {
      onlyAdmin: true
    }
  },
  {
    path: "/ws/configs",
    name: "configs",
    component: () => import("@/views/administration/configs/configs.vue"),
    meta: {
      onlyAdmin: true
    }
  },
  {
    path: "/ws/user-related/:mode",
    name: "userRelated",
    component: () => import("@/views/administration/user-related/user-related.vue"),
    meta: {
      onlyAdmin: true,
      webTitle: {
        params: "mode"
      }
    }
  }
]

const serviceManagement = <RouteRecordRaw[]>[
  {
    path: "/ws/groups",
    name: "groups",
    component: () => import("@/views/service-management/group/groups.vue"),
    meta: {}
  }
]

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
      ...administrator,
      ...serviceManagement
    ]
  }
]
