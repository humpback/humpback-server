export const LimitUserName = { Min: 2, Max: 100 }
export const LimitTeamName = { Min: 2, Max: 100 }
export const LimitEmail = { Min: 0, Max: 200 }
export const LimitPhone = { Min: 0, Max: 11 }
export const LimitPassword = { Min: 8, Max: 20 }
export const LimitDescription = { Min: 0, Max: 500 }
export const LimitConfigValue = { Min: 1, Max: 10000 }

// interface RuleLengthLimit {
//   Min: number
//   Max: number
// }

// interface RuleLength {
//   Username: RuleLengthLimit
//   TeamName: RuleLengthLimit
//   Email: RuleLengthLimit
//   Password: RuleLengthLimit
//   Phone: RuleLengthLimit
//   Description: RuleLengthLimit
// }
//
// interface RuleFormat {
//   Email: string
//   Phone: string
// }
//
// function initRule(info: any) {}
