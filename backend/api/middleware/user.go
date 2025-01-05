package middleware

import (
	"humpback/types"

	"github.com/gin-gonic/gin"
)

const (
	UserInfoKey   = "userInfo"
	UserCookieKey = "sessionId"
)

func GetUserInfo(c *gin.Context) *types.User {
	info, _ := c.Get(UserInfoKey)
	return info.(*types.User)
}

func SetUserInfo(c *gin.Context, usreInfo *types.User) {
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
