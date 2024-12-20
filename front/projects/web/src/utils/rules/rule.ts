import { GetI18nMessage } from "@/locales"

export const RegularName = /^[\u4e00-\u9fa5a-zA-Z0-9][\u4e00-\u9fa5a-zA-Z0-9\s_@,，-]{0,198}[\u4e00-\u9fa5a-zA-Z0-9]$/
export const RegularPassword = /^[a-zA-Z0-9][a-zA-Z0-9_\-@#$%+=!]{7,15}$/
export const RegularEnterpriseCode = /^[a-zA-Z]{4}$/

export function IsEmpty(value: any): boolean {
  return !value || !(value as string).trim()
}

export function IsValidName(name: string): boolean {
  return RegularName.test(name)
}

export function IsValidPassword(psd: string): boolean {
  return RegularPassword.test(psd)
}

// --------------------通用规则定义-----------------------

export function RuleCannotBeEmpty(rule: any, value: any, callback: any) {
  return IsEmpty(value) ? callback(new Error(GetI18nMessage("rules.cannotBeEmpty"))) : callback()
}

export function RuleIsRequired(fieldI18nName: string) {
  return (rule: any, value: any, callback: any) => {
    return IsEmpty(value) ? callback(new Error(`${GetI18nMessage(fieldI18nName)} ${GetI18nMessage("rules.isRequired")}`)) : callback()
  }
}

export function RuleLimitMax(max: number) {
  return (rule: any, value: any, callback: any) => {
    return (value as string).length > max ? callback(new Error(GetI18nMessage("rules.limitLengthMax", { max: max }))) : callback()
  }
}

export function RuleLimitRange(min: number, max: number) {
  return (rule: any, value: any, callback: any) => {
    const len = (value as string).length
    if (len < min || len > max) {
      return callback(new Error(GetI18nMessage("rules.limitLengthRange", { min: min, max: max })))
    }
    callback()
  }
}
