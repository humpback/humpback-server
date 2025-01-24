interface RuleLengthLimit {
  Min: number
  Max: number
}

interface RuleLength {
  Username: RuleLengthLimit
  TeamName: RuleLengthLimit
  Email: RuleLengthLimit
  Password: RuleLengthLimit
  Phone: RuleLengthLimit
  Description: RuleLengthLimit
  ConfigValue: RuleLengthLimit
}

interface RuleFormat {
  Email: string
  Phone: string
}

export let RuleLength = {} as RuleLength
export let RuleFormat = {} as RuleFormat

export function initRule(lengthLimit: RuleLength, formatRule: RuleFormat) {
  RuleLength = lengthLimit
  RuleFormat = formatRule
}
