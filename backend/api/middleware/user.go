package middleware

import (
	"github.com/gin-gonic/gin"
	"humpback/types"
)

const (
	UserInfoKey   = "userInfo"
	UserCookieKey = "sessionId"
)

func GetUserInfo(c *gin.Context) *types.UserInfo {
	info, _ := c.Get(UserInfoKey)
	return info.(*types.UserInfo)
}

func SetUserInfo(c *gin.Context, usreInfo *types.UserInfo) {
	c.Set(UserInfoKey, usreInfo)
}

func SetUserSessionId(c *gin.Context, sessionId string) {
	c.Set(UserCookieKey, sessionId)
}

func GetUserSessionId(c *gin.Context) string {
	id, _ := c.Get(UserCookieKey)
	return id.(string)
}

func SetUserCookie(c *gin.Context, sessionId string, maxAge int) {
	if maxAge > 0 {
		c.SetCookie(UserCookieKey, sessionId, maxAge, "/webapi", "", false, true)
		return
	}
	c.SetCookie(UserCookieKey, sessionId, -1, "/webapi", "", false, true)
}

func GetUserCookie(c *gin.Context) (string, error) {
	return c.Cookie(UserCookieKey)
}
