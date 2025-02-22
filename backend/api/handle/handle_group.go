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
	router.GET("/info/:id", groupInfo)
	router.POST("/query", groupQuery)
	router.PUT("/nodes", groupUpdateNodes)
	router.POST("/:groupId/query", groupNodeQuery)
	router.DELETE("/:id", middleware.CheckAdminPermissions(), groupDelete)
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

func groupUpdateNodes(c *gin.Context) {
	body := new(models.GroupUpdateNodesReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	userInfo := middleware.GetUserInfo(c)
	id, err := controller.GroupUpdateNodes(userInfo, body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func groupInfo(c *gin.Context) {
	id := c.Param("id")
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
	id := c.Param("id")
	if err := controller.GroupDelete(id); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}

func groupNodeQuery(c *gin.Context) {
	groupId := c.Param("groupId")
	queryInfo := new(models.GroupQueryNodesReqInfo)
	if !middleware.BindAndCheckBody(c, queryInfo) {
		return
	}
	result, err := controller.GroupNodesQuery(groupId, middleware.GetUserInfo(c), queryInfo)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
