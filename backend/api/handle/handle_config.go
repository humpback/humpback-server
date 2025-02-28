package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/api/handle/models"
	"humpback/api/middleware"
	"humpback/common/response"
	"humpback/internal/controller"
)

func RouteConfig(router *gin.RouterGroup) {
	router.POST("", configCreate)
	router.PUT("", configUpdate)
	router.GET("/info/:id", configInfo)
	router.POST("/query", configQuery)
	router.DELETE("/:id", configDelete)
}

func configCreate(c *gin.Context) {
	body := new(models.ConfigCreateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	id, err := controller.ConfigCreate(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func configUpdate(c *gin.Context) {
	body := new(models.ConfigUpdateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	id, err := controller.ConfigUpdate(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func configInfo(c *gin.Context) {
	id := c.Param("id")
	info, err := controller.Config(id)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, info)
}

func configQuery(c *gin.Context) {
	queryInfo := new(models.ConfigQueryReqInfo)
	if !middleware.BindAndCheckBody(c, queryInfo) {
		return
	}
	result, err := controller.ConfigQuery(queryInfo)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func configDelete(c *gin.Context) {
	id := c.Param("id")
	if err := controller.ConfigDelete(id); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}
