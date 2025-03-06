package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/api/handle/models"
	"humpback/api/middleware"
	"humpback/common/response"
	"humpback/internal/controller"
)

func RouteService(router *gin.RouterGroup) {
	router.POST("/query", serviceQuery)
	router.GET("/total", serviceTotal)
	router.POST("", serviceCreate)
	router.PUT("", serviceUpdate)
	router.GET("/:serviceId/info", serviceInfo)
	router.PUT("/operate", serviceOperate)
	router.DELETE("/:serviceId", serviceDelete)
}

func serviceQuery(c *gin.Context) {
	queryInfo := new(models.ServiceQueryReqInfo)
	if !middleware.BindAndCheckBody(c, queryInfo) {
		return
	}
	groupId := c.Param("groupId")
	queryInfo.UserInfo = middleware.GetUserInfo(c)
	result, err := controller.ServiceQuery(groupId, queryInfo)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func serviceTotal(c *gin.Context) {
	groupId := c.Param("groupId")
	result, err := controller.ServiceTotal(groupId)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func serviceInfo(c *gin.Context) {
	groupId := c.Param("groupId")
	serviceId := c.Param("serviceId")
	result, err := controller.Service(groupId, serviceId)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func serviceCreate(c *gin.Context) {
	body := new(models.ServiceCreateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	body.GroupId = c.Param("groupId")
	result, err := controller.ServiceCreate(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func serviceUpdate(c *gin.Context) {
	body := new(models.ServiceUpdateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	body.GroupId = c.Param("groupId")
	result, err := controller.ServiceUpdate(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func serviceOperate(c *gin.Context) {
	body := new(models.ServiceOperateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	body.GroupId = c.Param("groupId")
	result, err := controller.ServiceOperate(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func serviceDelete(c *gin.Context) {
	serviceId := c.Param("serviceId")
	groupId := c.Param("groupId")
	if err := controller.ServiceDelete(groupId, serviceId); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}
