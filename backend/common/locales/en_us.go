package locales

import (
	"fmt"

	"humpback/common/enum"
)

var enUSMsg = map[string]string{
	CodeSucceed:              "Succeed",
	CodeServerErr:            "Server Error, please contact the administrator!",
	CodeRequestParamsInvalid: "Invalid request parameter.",
	CodeLoginExpired:         "Login expired, please log in again.",
	CodeNotLogin:             "Not logged in.",
	CodeNoPermission:         "You don't have permission to operate.",

	CodeUserNameLimitLength:  fmt.Sprintf("The user name length limit is %d to %d.", enum.LimitUsername.Min, enum.LimitUsername.Max),
	CodeUserNameIsInvalid:    "The user name is invalid",
	CodeUserNameNotEmpty:     "The user name cannot be empty.",
	CodeUserNotExist:         "The user does not exist.",
	CodeUserIdNotEmpty:       "The user ID cannot be empty.",
	CodeUserAlreadyExist:     "The user already exists.",
	CodeUserNameAlreadyExist: "The user name already exists.",
	CodeUserIsOwner:          "Can't update yourself.",

	CodeUserRoleIsInvalid: "The user role is invalid",

	CodeTeamNameNotEmpty:     "The team name cannot be empty.",
	CodeTeamNameLimitLength:  fmt.Sprintf("The team name length limit is %d to %d.", enum.LimitTeamName.Min, enum.LimitTeamName.Max),
	CodeTeamIdNotEmpty:       "The team ID cannot be empty.",
	CodeTeamNotExist:         "The team does not exist.",
	CodeTeamAlreadyExist:     "The team already exists.",
	CodeTeamNameAlreadyExist: "The team name already exists.",

	CodeEmailIsInvalid:   "The email address is invalid",
	CodeEmailLimitLength: fmt.Sprintf("The email address length limit is %d.", enum.LimitEmail.Max),

	CodePhoneLimitLength: fmt.Sprintf("The phone number length limit is %d.", enum.LimitPhone.Max),
	CodePhoneIsInvalid:   "The phone number is invalid",

	CodeDescriptionLimitLength: fmt.Sprintf("The description length limit is %d.", enum.LimitDescription.Max),

	CodePasswordLimitLength:    fmt.Sprintf("The password length limit is %d to %d.", enum.LimitPassword.Min, enum.LimitPassword.Max),
	CodePasswordIsWrong:        "Wrong password",
	CodePasswordNotEmpty:       "The password cannot be empty.",
	CodeOldPasswordNotEmpty:    "The old password cannot be empty.",
	CodeOldPasswordLimitLength: fmt.Sprintf("The old password length limit is %d to %d.", enum.LimitPassword.Min, enum.LimitPassword.Max),
	CodeOldPasswordIsWrong:     "The old password is wrong.",
	CodeNewPasswordNotEmpty:    "The new password cannot be empty.",
	CodeNewPasswordLimitLength: fmt.Sprintf("The new password length limit is %d to %d.", enum.LimitPassword.Min, enum.LimitPassword.Max),
}
