package locales

// 基础系统code
const (
	CodeSucceed              = "S20000"
	CodeServerErr            = "R50000"
	CodeRequestParamsInvalid = "R40000"
	CodeLoginExpired         = "R40101"
	CodeNotLogin             = "R40102"
	CodeNoPermission         = "RNoPermission-001"

	CodeUserNameNotEmpty    = "R4User-001"
	CodeUserNameIsInvalid   = "R4User-002"
	CodeUserNameLimitLength = "R4User-003"
	CodeUserNotExist        = "R4User-004"
	CodeUserIdNotEmpty      = "R4User-005"
	CodeUserAlreadyExist    = "R4User-006"
	CodeUserIsOwner         = "R4User-007"

	CodeUserRoleIsInvalid = "R4Role-001"

	CodeTeamNameNotEmpty    = "R4Team-001"
	CodeTeamNameLimitLength = "R4Team-002"
	CodeTeamIdNotEmpty      = "R4Team-003"
	CodeTeamNotExist        = "R4Team-004"

	CodeEmailIsInvalid   = "R4Email-001"
	CodeEmailLimitLength = "R4Email-002"

	CodePhoneLimitLength = "R4Phone-001"
	CodePhoneIsInvalid   = "R4Phone-002"

	CodeDescriptionLimitLength = "R4Desc-001"

	CodePasswordLimitLength    = "R4Psd-001"
	CodePasswordIsWrong        = "R4Psd-002"
	CodePasswordNotEmpty       = "R4Psd-003"
	CodeOldPasswordNotEmpty    = "R4Psd-004"
	CodeOldPasswordLimitLength = "R4Psd-005"
	CodeOldPasswordIsWrong     = "R4Psd-006"
	CodeNewPasswordNotEmpty    = "R4Psd-007"
	CodeNewPasswordLimitLength = "R4Psd-008"
)
