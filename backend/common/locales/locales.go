package locales

import "strings"

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
