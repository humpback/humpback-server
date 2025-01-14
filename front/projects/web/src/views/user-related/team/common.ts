import { NewPageInfo, NewSortInfo, QueryInfo } from "@/types"

export const sortOptions = ["name", "updatedAt", "createdAt"]

export const defaultSort = NewSortInfo("name", "asc")
export const defaultPage = NewPageInfo(1, 20)

export class QueryTeamInfo extends QueryInfo {
  constructor(queryInfo: any) {
    super(queryInfo, ["name"], defaultPage, defaultSort, sortOptions, {})
  }

  getQuery() {
    return {
      query: Object.assign({}, this.getBaseQuery())
    }
  }
}
