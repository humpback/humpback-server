package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/api/handle/models"
	"humpback/api/middleware"
	"humpback/config"
	"humpback/internal/controller"
)

func RouteUser(router *gin.RouterGroup) {
	router.POST("/login", login)
	router.GET("", middleware.CheckLogin(), user)
	router.POST("/logout", middleware.CheckLogin(), logout)
}

func login(c *gin.Context) {
	body := new(models.UserLoginReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	userInfo, sessionId, err := controller.UserLogin(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	userInfo.Password = ""
	middleware.SetCookieSession(c, sessionId, int(config.DBArgs().SessionTimeout.Seconds()))
	c.JSON(http.StatusOK, userInfo)
}

func user(c *gin.Context) {
	c.JSON(http.StatusOK, middleware.GetUserInfo(c))
}

func logout(c *gin.Context) {
	sessionId := middleware.GetSessionId(c)
	if err := controller.SessionDelete(sessionId); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	middleware.SetCookieSession(c, sessionId, 0)
	c.JSON(http.StatusOK, nil)
}
