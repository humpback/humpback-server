import { NewPageInfo, NewSortInfo, QueryInfo } from "@/types"
import { find, map, omitBy } from "lodash-es"
import { NodeStatus, NodeSwitch } from "@/models"

export const sortOptions = ["ip", "hostname", "updatedAt", "createdAt"]

export const defaultSort = NewSortInfo("ip", "asc")
export const defaultPage = NewPageInfo(1, 20)
export const defaultFilter = { status: "" }

export const statusOptions = [
  { label: "label.all", value: "" },
  { label: "label.enabled", value: NodeSwitch.Enabled },
  { label: "label.disabled", value: NodeSwitch.Disabled },
  { label: "label.healthy", value: NodeStatus.Online },
  { label: "label.deadly", value: NodeStatus.Offline }
]

export const modeOptions = [
  { label: "label.keywords", value: "keywords" },
  { label: "label.label", value: "label" }
]

export class QueryNodesInfo extends QueryInfo {
  constructor(queryInfo: any, groupOptions: any[]) {
    super(
      queryInfo,
      map(modeOptions, x => x.value),
      defaultPage,
      defaultSort,
      sortOptions,
      defaultFilter
    )
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
  return `docker run -d --name=humpback-agent
--net=host
--restart=always
-e ip=${ip}
-v /etc/localtime:/etc/localtime
humpback:latest
`
}
