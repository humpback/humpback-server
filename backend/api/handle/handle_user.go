package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/api/handle/models"
	"humpback/api/middleware"
	"humpback/common/response"
	"humpback/config"
	"humpback/internal/controller"
)

func RouteUser(router *gin.RouterGroup) {
	router.POST("/login", login)
	router.POST("/logout", middleware.CheckLogin(), logout)
	router.GET("", middleware.CheckLogin(), user)
	router.PUT("", middleware.CheckLogin(), userUpdate)
	router.PUT("/change-psd", middleware.CheckLogin(), userChangePassword)
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

func logout(c *gin.Context) {
	sessionId := middleware.GetSessionId(c)
	if err := controller.SessionDelete(sessionId); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	middleware.SetCookieSession(c, sessionId, 0)
	c.JSON(http.StatusOK, nil)
}

func user(c *gin.Context) {
	userInfo := middleware.GetUserInfo(c)
	userInfo.Password = ""
	c.JSON(http.StatusOK, userInfo)
}

func userUpdate(c *gin.Context) {
	body := new(models.UserUpdateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	userInfo := middleware.GetUserInfo(c)
	userInfo = body.NewUserInfo(userInfo)
	if err := controller.UserUpdate(userInfo); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}

func userChangePassword(c *gin.Context) {
	body := new(models.UserChangePasswordReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	userInfo := middleware.GetUserInfo(c)
	if err := controller.UserChangePassword(userInfo, body); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	sessionId := middleware.GetSessionId(c)
	middleware.SetCookieSession(c, sessionId, 0)
	c.JSON(http.StatusOK, response.NewRespSucceed())
}
