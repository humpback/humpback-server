export interface ResponseSuccess {
  statusCode: number
  objectId: string
  msg: string
}

export interface OperateUserBaseInfo {
  name: string
  status: number
  timeAt: number
  userId: string
}

export interface OperateUserInfo {
  createdUserInfo: OperateUserBaseInfo
  updatedUserInfo: OperateUserBaseInfo
  deletedUserInfo?: OperateUserBaseInfo
}

export function NewOperateUserInfo(): OperateUserInfo {
  return {
    createdUserInfo: {
      name: "",
      status: 0,
      timeAt: 0,
      userId: ""
    },
    updatedUserInfo: {
      name: "",
      status: 0,
      timeAt: 0,
      userId: ""
    }
  }
}
