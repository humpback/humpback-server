package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/api/handle/models"
	"humpback/api/middleware"
	"humpback/internal/controller"
)

func RouteService(router *gin.RouterGroup) {
	router.POST("/query", serviceQuery)
	router.GET("/total", serviceTotal)
	router.POST("", serviceCreate)
	//router.PUT("", serviceUpdate)
	//router.GET("/info/:serviceId", serviceInfo)
	//router.DELETE("/:serviceId", serviceDelete)
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

func serviceCreate(c *gin.Context) {
	body := new(models.ServiceCreateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	result, err := controller.ServiceCreate(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
