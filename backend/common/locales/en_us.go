package locales

import (
	"fmt"
)

var baseMsg = map[string]string{
	CodeSucceed:              "Succeed",
	CodeServerErr:            "Server Error, please contact the administrator!",
	CodeRequestParamsInvalid: "Invalid request parameter.",
	CodeLoginExpired:         "Login expired, please log in again.",
	CodeNotLogin:             "Not logged in.",
	CodeNoPermission:         "You don't have permission to operate.",

	CodeUserNameLimitLength: fmt.Sprintf("The user name length limit is %d to %d.", LimitUserName.Min, LimitUserName.Max),
	CodeUserNameIsInvalid:   "The user name is invalid",
	CodeUserNameNotEmpty:    "The user name cannot be empty.",
	CodeUserNotExist:        "The user does not exist.",

	CodeEmailIsInvalid:   "The email address is invalid",
	CodeEmailLimitLength: fmt.Sprintf("The email address length limit is %d.", LimitEmail.Max),

	CodePhoneLimitLength: fmt.Sprintf("The phone number length limit is %d.", LimitPhone.Max),
	CodePhoneIsInvalid:   "The phone number is invalid",

	CodeDescriptionLimitLength: fmt.Sprintf("The description length limit is %d.", LimitDescription.Max),

	CodePasswordLimitLength:    fmt.Sprintf("The password length limit is %d to %d.", LimitPassword.Min, LimitPassword.Max),
	CodePasswordIsWrong:        "Wrong password",
	CodePasswordNotEmpty:       "The password cannot be empty.",
	CodeOldPasswordNotEmpty:    "The old password cannot be empty.",
	CodeOldPasswordLimitLength: fmt.Sprintf("The old password length limit is %d to %d.", LimitPassword.Min, LimitPassword.Max),
	CodeOldPasswordIsWrong:     "The old password is wrong.",
	CodeNewPasswordNotEmpty:    "The new password cannot be empty.",
	CodeNewPasswordLimitLength: fmt.Sprintf("The new password length limit is %d to %d.", LimitPassword.Min, LimitPassword.Max),
}
