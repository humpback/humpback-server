package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/api/handle/models"
	"humpback/api/middleware"
	"humpback/common/response"
	"humpback/internal/controller"
)

func RouteNodes(router *gin.RouterGroup) {
	router.POST("", middleware.CheckAdminPermissions(), nodesCreate)
	router.PUT("/labels", middleware.CheckAdminPermissions(), nodeUpdateLabels)
	router.GET("/info/:id", registryGet)
	router.POST("/query", registryQuery)
	router.DELETE("/:id", middleware.CheckAdminPermissions(), nodeDelete)
}

func nodesCreate(c *gin.Context) {
	nodes := make(models.NodesCreateReqInfo, 0)
	if !middleware.BindAndCheckBody(c, &nodes) {
		return
	}
	if err := controller.NodeCreate(nodes); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}

func nodeUpdateLabels(c *gin.Context) {
	body := new(models.NodeUpdateLabelReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	id, err := controller.NodeUpdateLabel(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func nodeDelete(c *gin.Context) {
	id := c.Param("id")
	if err := controller.NodeDelete(id); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}
