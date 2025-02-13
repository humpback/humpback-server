interface RuleLengthLimit {
  Min: number
  Max: number
}

interface RuleLength {
  Username: RuleLengthLimit
  TeamName: RuleLengthLimit
  ConfigName: RuleLengthLimit
  Email: RuleLengthLimit
  Password: RuleLengthLimit
  Phone: RuleLengthLimit
  Description: RuleLengthLimit
  ConfigValue: RuleLengthLimit
  RegistryName: RuleLengthLimit
  RegistryUrl: RuleLengthLimit
  RegistryUsername: RuleLengthLimit
  RegistryPassword: RuleLengthLimit
}

interface RuleFormat {
  Email: string
  Phone: string
  IPAddress: string
}

export let RuleLength = {} as RuleLength
export let RuleFormat = {} as RuleFormat

export function initRule(lengthLimit: RuleLength, formatRule: RuleFormat) {
  RuleLength = lengthLimit
  RuleFormat = formatRule
}
