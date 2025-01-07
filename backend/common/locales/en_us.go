package locales

var baseMsg = map[string]string{
	CodeSucceed:              "Succeed",
	CodeServerErr:            "Server Error, please contact the administrator!",
	CodeRequestParamsInvalid: "Invalid request parameter.",
	CodeLoginExpired:         "Login expired, please log in again.",
	CodeNotLogin:             "Not logged in.",
	CodeNoPermission:         "You don't have permission to operate.",

	CodeUserNameIsInvalid: "The user name is invalid",
	CodeUserNameNotEmpty:  "The user name cannot be empty.",
	CodeUserNotExist:      "The user does not exist.",
	CodeUserIsInvalid:     "The user is invalid.",

	CodeEmailIsInvalid:    "The email address is invalid",
	CodePasswordIsInvalid: "The password is invalid",
	CodePasswordIsWrong:   "Wrong password",
	CodePasswordNotEmpty:  "The password cannot be empty.",
}
