package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/api/handle/models"
	"humpback/api/middleware"
	"humpback/common/response"
	"humpback/internal/controller"
)

func RouteGroup(router *gin.RouterGroup) {
	router.POST("", middleware.CheckAdminPermissions(), groupCreate)
	router.PUT("", groupUpdate)
	router.GET("/:groupId/info", groupInfo)
	router.POST("/query", groupQuery)
	router.DELETE("/:groupId", middleware.CheckAdminPermissions(), groupDelete)
}

func groupCreate(c *gin.Context) {
	body := new(models.GroupCreateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	id, err := controller.GroupCreate(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func groupUpdate(c *gin.Context) {
	body := new(models.GroupUpdateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	userInfo := middleware.GetUserInfo(c)
	id, err := controller.GroupUpdate(userInfo, body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func groupInfo(c *gin.Context) {
	id := c.Param("groupId")
	userInfo := middleware.GetUserInfo(c)
	info, err := controller.Group(userInfo, id)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, info)
}

func groupQuery(c *gin.Context) {
	queryInfo := new(models.GroupQueryReqInfo)
	if !middleware.BindAndCheckBody(c, queryInfo) {
		return
	}
	queryInfo.UserInfo = middleware.GetUserInfo(c)
	result, err := controller.GroupQuery(queryInfo)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func groupDelete(c *gin.Context) {
	id := c.Param("groupId")
	if err := controller.GroupDelete(id); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}
