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
	router.GET("/me", middleware.CheckLogin(), me)
	router.PUT("/me", middleware.CheckLogin(), meUpdate)
	router.PUT("/me/change-psd", middleware.CheckLogin(), meChangePassword)

	router.POST("", middleware.CheckLogin(), middleware.CheckAdminPermissions(), userCreate)
	router.PUT("", middleware.CheckLogin(), middleware.CheckAdminPermissions(), userUpdate)
	router.GET("/info/:id", middleware.CheckLogin(), middleware.CheckAdminPermissions(), user)
	router.POST("/query", middleware.CheckLogin(), middleware.CheckAdminPermissions(), users)
	router.GET("/query-by-team/:teamId", middleware.CheckLogin(), middleware.CheckAdminPermissions(), usersByTeamId)
	router.DELETE("/:id", middleware.CheckLogin(), middleware.CheckAdminPermissions(), userDelete)
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

func me(c *gin.Context) {
	userInfo := middleware.GetUserInfo(c)
	userInfo.Password = ""
	c.JSON(http.StatusOK, userInfo)
}

func meUpdate(c *gin.Context) {
	body := new(models.MeUpdateReqInfo)
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

func meChangePassword(c *gin.Context) {
	body := new(models.MeChangePasswordReqInfo)
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

func userCreate(c *gin.Context) {

}

func userUpdate(c *gin.Context) {

}

func user(c *gin.Context) {

}

func users(c *gin.Context) {

}

func usersByTeamId(c *gin.Context) {

}

func userDelete(c *gin.Context) {

}
