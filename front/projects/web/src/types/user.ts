import { NewOperateUserInfo, OperateUserInfo } from "#/index.ts"

export interface OrgInfo {
  orgId: string
  address?: {
    street: string
    city: string
    state: string
  }
  companyName: string
  enterpriseCode: string
}

export function NewOrgEmptyInfo(): OrgInfo {
  return {
    orgId: "",
    address: {
      street: "",
      city: "",
      state: ""
    },
    companyName: "",
    enterpriseCode: ""
  }
}

export interface UserInfo extends OperateUserInfo {
  orgId: string
  userId: string
  name: string
  email: string
  phone: string
  role: number
  notes: string
  status: number
  orgInfo: OrgInfo
}

export function NewUserEmptyInfo(): UserInfo {
  return {
    ...NewOperateUserInfo(),
    orgId: "",
    userId: "",
    name: "",
    email: "",
    phone: "",
    role: -1,
    notes: "",
    status: -1,
    orgInfo: NewOrgEmptyInfo()
  }
}

export interface UserBill {
  orgId: string
  userId: string
  orderNO: string
  creditCard: string
  orderAmount: number
  deductionTime: number
}
