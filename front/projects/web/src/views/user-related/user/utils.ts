import { NewPageInfo, NewSortInfo } from "@/types"

export const sortOptions = ["username", "updatedAt", "createdAt"]

export const defaultSort = NewSortInfo("username", "asc")

export const defaultPage = NewPageInfo(1, 20)
