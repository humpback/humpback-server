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
	router.PUT("/switch", middleware.CheckAdminPermissions(), nodeUpdateSwitch)
	router.GET("/info/:id", node)
	router.POST("/query", nodesQuery)
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

func nodeUpdateSwitch(c *gin.Context) {
	body := new(models.NodeUpdateSwitchReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	id, err := controller.NodeUpdateSwitch(body.NodeId, body.Enable)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func node(c *gin.Context) {
	id := c.Param("id")
	info, err := controller.Node(id)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, info)
}

func nodesQuery(c *gin.Context) {
	queryInfo := new(models.NodeQueryReqInfo)
	if !middleware.BindAndCheckBody(c, queryInfo) {
		return
	}
	result, err := controller.NodesQuery(queryInfo)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func nodeDelete(c *gin.Context) {
	id := c.Param("id")
	if err := controller.NodeDelete(id); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}
