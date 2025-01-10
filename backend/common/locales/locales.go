package locales

import (
	"regexp"
	"strings"
)

var (
	RegularEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	RegularPhone = regexp.MustCompile(`^\d+$`)
)

type LengthLimit struct {
	Min int
	Max int
}

var (
	LimitUserName    = LengthLimit{Min: 2, Max: 100}
	LimitEmail       = LengthLimit{Min: 0, Max: 200}
	LimitPassword    = LengthLimit{Min: 8, Max: 20}
	LimitPhone       = LengthLimit{Min: 0, Max: 11}
	LimitDescription = LengthLimit{Min: 0, Max: 500}
)

func GetMsg(language string, key string) string {
	msg := baseMsg[key]
	switch strings.ToLower(language) {
	case "zh-cn":
		if v := zhCnMsg[key]; v != "" {
			msg = v
		}
	}
	return msg
}
