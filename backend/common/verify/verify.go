package verify

import (
	"regexp"
	"strings"

	"humpback/common/locales"
	"humpback/common/response"
)

var (
	RegularName     = regexp.MustCompile(`^[\p{Han}a-zA-Z0-9][\p{Han}a-zA-Z0-9\s_@,，-]{0,98}[\p{Han}a-zA-Z0-9]$`)
	RegularEmail    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	RegularPassword = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_\-@#$%+=!]{7,19}$`)
)

type LengthLimit struct {
	Min int
	Max int
}

var (
	LimitUserName = LengthLimit{Min: 2, Max: 100}
	LimitEmail    = LengthLimit{Min: 0, Max: 254}
	LimitPassword = LengthLimit{Min: 8, Max: 20}
)

func isValidName(name string) bool {
	return RegularName.MatchString(name)
}

func IsValidEmail(email string) bool {
	return RegularEmail.MatchString(email)
}

func IsValidPassword(psd string) bool {
	return RegularPassword.MatchString(psd)
}

func CheckIsEmpty(value string, code string) error {
	if strings.TrimSpace(value) == "" {
		return response.NewBadRequestErr(code)
	}
	return nil
}

func CheckUsername(name string) error {
	if !isValidName(name) {
		return response.NewBadRequestErr(locales.CodeUserNameIsInvalid)
	}
	return nil
}

func CheckEmail(email string) error {
	if !IsValidEmail(email) {
		return response.NewBadRequestErr(locales.CodeEmailIsInvalid)
	}
	return nil
}

func CheckPassword(psd string) error {
	if !IsValidPassword(psd) {
		return response.NewBadRequestErr(locales.CodePasswordIsInvalid)
	}
	return nil
}
