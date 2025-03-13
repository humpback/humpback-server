package handle

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"humpback/api/handle/models"
	"humpback/api/middleware"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/internal/controller"
)

func RouteGroupContainer(router *gin.RouterGroup) {
	router.PUT("", groupContainerOperate)
	router.POST("/logs", groupContainerQueryLogs)
	router.POST("/performace", groupContainerPerformance)
}

func groupContainerOperate(c *gin.Context) {
	body := new(models.GroupContainerOperateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	groupInfo := middleware.GetGroupInfo(c)
	if slices.Index(groupInfo.Nodes, body.NodeId) == -1 {
		middleware.AbortErr(c, response.NewBadRequestErr(locales.CodeNodesNotExist))
		return
	}
	if err := controller.GroupContainerOperate(body); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}

func groupContainerQueryLogs(c *gin.Context) {
	body := new(models.GroupContainerLogsReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	groupInfo := middleware.GetGroupInfo(c)
	if slices.Index(groupInfo.Nodes, body.NodeId) == -1 {
		middleware.AbortErr(c, response.NewBadRequestErr(locales.CodeNodesNotExist))
		return
	}
	logs, err := controller.GroupContainerQueryLogs(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, logs)
}

func groupContainerPerformance(c *gin.Context) {
	body := make(models.GroupContainerPerformanceReqInfo, 0)
	if !middleware.BindAndCheckBody(c, &body) {
		return
	}
	result, err := controller.GroupContainerPerformances(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
