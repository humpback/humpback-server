import { NewPageInfo, NewSortInfo, QueryInfo } from "@/types"
import { cloneDeep, find, omitBy } from "lodash-es"
import { NodeStatus, NodeSwitch } from "@/models"

export const sortOptions = ["ipAddress", "name", "updatedAt", "createdAt"]

export const defaultSort = NewSortInfo("ipAddress", "asc")
export const defaultPage = NewPageInfo(1, 20)
export const defaultFilter = { status: "" }

export const statusOptions: Array<{ type?: "primary" | "success" | "warning" | "danger" | "info"; label: string; value: string }> = [
  { label: "label.all", value: "" },
  { type: "success", label: "label.enabled", value: NodeSwitch.Enabled },
  { type: "info", label: "label.disabled", value: NodeSwitch.Disabled },
  { type: "success", label: "label.healthy", value: NodeStatus.Online },
  { type: "danger", label: "label.deadly", value: NodeStatus.Offline }
]

export class QueryNodesInfo extends QueryInfo {
  constructor(queryInfo: any, groupOptions: any[]) {
    super(queryInfo, ["keywords"], defaultPage, defaultSort, sortOptions, cloneDeep(defaultFilter))
    const statusInfo = find(statusOptions, x => x.value === (queryInfo["status"] as string))
    this.filter.status = statusInfo?.value || ""
  }

  urlQuery() {
    return {
      query: Object.assign(
        {},
        {
          group: this.filter.group || undefined,
          status: this.filter.status || undefined
        },
        this.getBaseQuery()
      )
    }
  }

  searchParams() {
    return omitBy(this, (value, key) => key.startsWith("_"))
  }
}

export function NewCommand(ip: string, isUninstall?: boolean) {
  if (isUninstall) {
    return `docker rm -f humpback-agent`
  }
  return `docker run -d --name=humpback-agent \\
--net=host \\
--restart=always \\
-e ip=${ip} \\
-v /etc/localtime:/etc/localtime \\
humpback:latest`
}
