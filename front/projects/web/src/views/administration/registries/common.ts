import { NewPageInfo, NewSortInfo, QueryInfo } from "@/types"
import { omitBy } from "lodash-es"

export const sortOptions = ["registryName", "updatedAt", "createdAt"]

export const defaultSort = NewSortInfo("registryName", "asc")
export const defaultPage = NewPageInfo(1, 20)

export class QueryRegistryInfo extends QueryInfo {
  constructor(queryInfo: any) {
    super(queryInfo, ["registryName"], defaultPage, defaultSort, sortOptions, {})
  }

  getQuery() {
    return {
      query: Object.assign({}, this.getBaseQuery())
    }
  }

  getSearch() {
    return omitBy(this, (value, key) => key.startsWith("_"))
  }
}
