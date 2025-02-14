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
	router.POST("/query", middleware.CheckLogin(), usersQuery)
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
	if err := controller.MeUpdate(userInfo); err != nil {
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
	if err := controller.MeChangePassword(userInfo, body); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	sessionId := middleware.GetSessionId(c)
	middleware.SetCookieSession(c, sessionId, 0)
	c.JSON(http.StatusOK, response.NewRespSucceed())
}

func userCreate(c *gin.Context) {
	body := new(models.UserCreateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	userInfo := middleware.GetUserInfo(c)
	if err := body.CheckCreateRole(userInfo); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	id, err := controller.UserCreate(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func userUpdate(c *gin.Context) {
	body := new(models.UserUpdateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	userInfo := middleware.GetUserInfo(c)
	if err := body.CheckUpdateRole(userInfo); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	id, err := controller.UserUpdate(body, userInfo)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func user(c *gin.Context) {
	id := c.Param("id")
	includePassword := c.Query("p")
	info, err := controller.User(id)
	if includePassword != "true" {
		info.Password = ""
	}
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, info)
}

func usersQuery(c *gin.Context) {
	queryInfo := new(models.UserQueryReqInfo)
	if !middleware.BindAndCheckBody(c, queryInfo) {
		return
	}
	result, err := controller.UsersQuery(queryInfo)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func usersByTeamId(c *gin.Context) {
	teamId := c.Param("teamId")
	users, err := controller.UsersGetByTeamId(teamId)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func userDelete(c *gin.Context) {
	id := c.Param("id")
	userInfo := middleware.GetUserInfo(c)
	if err := controller.UserDelete(id, userInfo); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}
