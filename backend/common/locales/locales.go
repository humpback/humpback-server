package locales

import (
	"strings"
)

func GetMsg(language string, key string) string {
	msg := enUSMsg[key]
	switch strings.ToLower(language) {
	case "zh-cn":
		if v := zhCnMsg[key]; v != "" {
			msg = v
		}
	}
	return msg
}

const (
	CodeSucceed              = "S20000"
	CodeServerErr            = "R50000"
	CodeRequestParamsInvalid = "R40000"
	CodeLoginExpired         = "R40101"
	CodeNotLogin             = "R40102"
	CodeNoPermission         = "RNoPermission-001"

	CodeUserNameNotEmpty     = "R4User-001"
	CodeUserNameIsInvalid    = "R4User-002"
	CodeUserNameLimitLength  = "R4User-003"
	CodeUserNotExist         = "R4User-004"
	CodeUserIdNotEmpty       = "R4User-005"
	CodeUserAlreadyExist     = "R4User-006"
	CodeUserIsOwner          = "R4User-007"
	CodeUserNameAlreadyExist = "R4User-008"

	CodeUserRoleIsInvalid = "R4Role-001"

	CodeTeamNameNotEmpty     = "R4Team-001"
	CodeTeamNameLimitLength  = "R4Team-002"
	CodeTeamIdNotEmpty       = "R4Team-003"
	CodeTeamNotExist         = "R4Team-004"
	CodeTeamAlreadyExist     = "R4Team-005"
	CodeTeamNameAlreadyExist = "R4Team-006"

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

	CodeConfigNameNotEmpty           = "R4Cfg-001"
	CodeConfigNameLimitLength        = "R4Cfg-002"
	CodeConfigNameAlreadyExist       = "R4Cfg-003"
	CodeConfigNotExist               = "R4Cfg-004"
	CodeConfigValueNotEmpty          = "R4Cfg-005"
	CodeConfigStaticValueLimitLength = "R4Cfg-006"
	CodeConfigVolumeValueLimitLength = "R4Cfg-007"
	CodeConfigTypeIsInvlaid          = "R4Cfg-008"
	CodeConfigIdNotEmpty             = "R4Cfg-009"

	CodeRegistryNameNotEmpty        = "R4Registry-001"
	CodeRegistryNameLimitLength     = "R4Registry-002"
	CodeRegistryNameAlreadyExist    = "R4Registry-003"
	CodeRegistryNotExist            = "R4Registry-004"
	CodeRegistryIdNotEmpty          = "R4Registry-005"
	CodeRegistryUrlNotEmpty         = "R4Registry-006"
	CodeRegistryUrlLimitLength      = "R4Registry-007"
	CodeRegistryUrlAlreadyExist     = "R4Registry-008"
	CodeRegistryUsernameLimitLength = "R4Registry-009"
	CodeRegistryPasswordLimitLength = "R4Registry-010"
)
