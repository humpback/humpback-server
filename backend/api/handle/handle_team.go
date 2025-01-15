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
	router.POST("", teamCreate)
	router.PUT("", teamUpdate)
	router.GET("/info/:id", team)
	router.POST("/query", teamsQuery)
	router.GET("/query-by-user/:userId", teamsByUserId)
	router.DELETE("/:id", teamDelete)
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
	teams, err := controller.TeamsByUserId(userId)
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
