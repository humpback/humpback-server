const ChannelKey = "HUMPBACK_USER_CHANGE"
const bc = new BroadcastChannel(ChannelKey)
const code = GenerateUUID()

export enum ChangeEventType {
  Login = "login",
  Logout = "logout"
}

export function GetChannelMessage(handleFunc: (data: any) => void) {
  bc.onmessage = function(e) {
    if (e.data?.code === code) {
      return
    }
    handleFunc(e.data)
  }
}

export function SendChannelMessage(type: ChangeEventType, value?: any) {
  bc.postMessage({ code: code, type: type, value: value })
}

export function CloseChannelMessage() {
  bc.close()
}

export function GenerateUUID(): string {
  let uuid = ""
  const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

  for (let i = 0; i < 32; i++) {
    const index = Math.floor(Math.random() * chars.length)
    uuid += chars[index]
    if (i === 7 || i === 11 || i === 15 || i === 19) {
      uuid += "-"
    }
  }
  return uuid
}
