import type { RouteRecordRaw } from "vue-router"
import { PageGroupDetail, PageLimitRole, PageServiceDetail, PageUserRelated } from "@/models"
import { find } from "lodash-es"

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
    path: "/ws/user-related/:mode",
    name: "userRelated",
    component: () => import("@/views/administration/user-related/user-related.vue"),
    beforeEnter: (to, from, next) => {
      return find([PageUserRelated.Users, PageUserRelated.Teams], x => x === to.params.mode) ? next() : next({ name: "404" })
    },
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
    component: () => import("@/views/service-management/groups.vue"),
    meta: {}
  },
  {
    path: "/ws/group/:groupId/:mode",
    name: "groupDetail",
    component: () => import("@/views/service-management/group-detail.vue"),
    beforeEnter: (to, from, next) => {
      return find([PageGroupDetail.Services, PageGroupDetail.Nodes], x => x === to.params.mode) ? next() : next({ name: "404" })
    },
    meta: {
      currentMenu: "groups",
      webTitle: {
        params: "mode"
      }
    }
  },
  {
    path: "/ws/group/:groupId/service/:serviceId/:mode",
    name: "serviceInfo",
    component: () => import("@/views/service-management/service/detail/service-detail.vue"),
    beforeEnter: (to, from, next) => {
      return find(
        [
          PageServiceDetail.BasicInfo,
          PageServiceDetail.Application,
          PageServiceDetail.Deployment,
          PageServiceDetail.Instances,
          PageServiceDetail.Log,
          PageServiceDetail.Performance
        ],
        x => x === to.params.mode
      )
        ? next()
        : next({ name: "404" })
    },
    meta: {
      currentMenu: "groups"
    }
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
      {
        path: "/ws/configs",
        name: "configs",
        component: () => import("@/views/configs/configs.vue"),
        meta: {}
      },
      ...administrator,
      ...serviceManagement
    ]
  }
]
