import { GetI18nMessage } from "@/locales"
import { ElNotification } from "element-plus"

const defaultDuration = 3000

// function showMessage(type: "info" | "success" | "warning" | "error", message: string, showClose?: boolean) {
//   ElMessage({
//     showClose: showClose,
//     message: message,
//     duration: defaultDuration,
//     type: type,
//     customClass: "messageBox"
//   })
// }

function notifyMsg(type: "info" | "success" | "warning" | "error", message: string, title?: string, showClose?: boolean) {
  ElNotification({
    type: type,
    title: title,
    message: message,
    showClose: showClose,
    duration: defaultDuration,
    customClass: "messageBox"
  })
}

export function ShowInfoMsg(message: string, title: string = GetI18nMessage("message.info")) {
  // showMessage("info", message, true)
  notifyMsg("info", message, title, true)
}

export function ShowSuccessMsg(message: string, title: string = GetI18nMessage("message.succeed")) {
  // showMessage("success", message, false)
  notifyMsg("success", message, title, false)
}

export function ShowWarningMsg(message: string, title: string = GetI18nMessage("message.warning")) {
  // showMessage("warning", message, true)
  notifyMsg("warning", message, title, true)
}

export function ShowErrMsg(message: string, title: string = GetI18nMessage("message.error")) {
  // showMessage("error", message, true)
  notifyMsg("error", message, title, true)
}

export function ShowSystemErrMsg() {
  // showMessage("error", GetI18nMessage("err.systemError"), true)
  notifyMsg("error", GetI18nMessage("err.systemError"), GetI18nMessage("message.error"), true)
}
