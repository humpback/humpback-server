package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/api/handle/models"
	"humpback/api/middleware"
	"humpback/common/response"
	"humpback/internal/controller"
)

func RouteTeam(router *gin.RouterGroup) {
	router.POST("", middleware.CheckAdminPermissions(), teamCreate)
	router.PUT("", middleware.CheckAdminPermissions(), teamUpdate)
	router.GET("/info/:id", middleware.CheckAdminPermissions(), team)
	router.POST("/query", teamsQuery)
	router.GET("/query-by-user/:userId", middleware.CheckAdminPermissions(), teamsByUserId)
	router.DELETE("/:id", middleware.CheckAdminPermissions(), teamDelete)
}

func teamCreate(c *gin.Context) {
	body := new(models.TeamCreateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	id, err := controller.TeamCreate(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func teamUpdate(c *gin.Context) {
	body := new(models.TeamUpdateReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	id, err := controller.TeamUpdate(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func team(c *gin.Context) {
	id := c.Param("id")
	info, err := controller.Team(id)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, info)
}

func teamsQuery(c *gin.Context) {
	queryInfo := new(models.TeamQueryReqInfo)
	if !middleware.BindAndCheckBody(c, queryInfo) {
		return
	}
	result, err := controller.TeamQuery(queryInfo)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func teamsByUserId(c *gin.Context) {
	userId := c.Param("userId")
	teams, err := controller.TeamsGetByUserId(userId)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, teams)
}

func teamDelete(c *gin.Context) {
	id := c.Param("id")
	if err := controller.TeamDelete(id); err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}
