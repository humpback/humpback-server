import { shallowRef } from "vue"
import { NewPageInfo, NewSortInfo, QueryInfo } from "@/types"
import { cloneDeep, find, omitBy } from "lodash-es"
import { NodeSwitch, ServiceStatus } from "@/models"

export const sortOptions = ["serviceName", "updatedAt", "createdAt"]

export const defaultSort = NewSortInfo("serviceName", "asc")
export const defaultPage = NewPageInfo(1, 20)
export const defaultFilter = { status: "", schedule: "" }

export class QueryServicesInfo extends QueryInfo {
  constructor(queryInfo: any) {
    super(queryInfo, ["keywords"], defaultPage, defaultSort, sortOptions, cloneDeep(defaultFilter))
    const status = find(
      [NodeSwitch.Enabled, NodeSwitch.Disabled, ServiceStatus.ServiceStatusNotReady, ServiceStatus.ServiceStatusRunning, ServiceStatus.ServiceStatusFailed],
      x => x === (queryInfo["status"] as string)
    )
    this.filter.status = status || ""
    const schedule = find(["Yes", "No"], x => x === (queryInfo["schedule"] as string))
    this.filter.schedule = schedule || ""
  }

  urlQuery() {
    return {
      query: Object.assign(
        {},
        {
          status: this.filter.status !== defaultFilter.status ? this.filter.status : undefined,
          schedule: this.filter.schedule !== defaultFilter.schedule ? this.filter.schedule : undefined
        },
        this.getBaseQuery()
      )
    }
  }

  searchParams() {
    return omitBy(this, (value, key) => key.startsWith("_"))
  }
}

export const ActionOptions: Array<{
  action: string
  type: "default" | "info" | "success" | "primary" | "text" | "warning" | "danger"
  i18nLabel: string
  icon: any
}> = [
  { action: "Enable", type: "success", i18nLabel: "btn.enable", icon: shallowRef(IconMdiPlay) },
  { action: "Disable", type: "info", i18nLabel: "btn.disable", icon: shallowRef(IconMdiSquare) },
  { action: "Start", type: "success", i18nLabel: "btn.start", icon: shallowRef(IconMdiPlay) },
  { action: "Restart", type: "success", i18nLabel: "btn.restart", icon: shallowRef(IconMdiRestart) },
  { action: "Stop", type: "primary", i18nLabel: "btn.stop", icon: shallowRef(IconMdiSquare) }
]
