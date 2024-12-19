package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/api/middleware"
	"humpback/api/static"
	"humpback/config"
)

var Router RouterInterface

type RouterInterface interface {
	Start() error
}

type router struct {
	engine *gin.Engine
}

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	r := &router{
		engine: gin.New(),
	}
	r.setMiddleware()
	r.setRoute()
	Router = r
}

func (api *router) Start() error {
	slog.Info("[Api] init front static resource to cache start...")
	//if err := static.InitStaticsResource(); err != nil {
	//	return fmt.Errorf("init front static resource to cache failed: %s", err)
	//}
	slog.Info("[Api] init front static resource to cache complted.")
	listeningAddress := fmt.Sprintf("%s:%s", config.NodeArgs().HostIp, config.NodeArgs().Port)
	slog.Info("[Api] listening...", "Address", listeningAddress)
	if err := api.engine.Run(listeningAddress); err != nil {
		return fmt.Errorf("listening %s failed: %s", listeningAddress, err)
	}
	return nil
}

func (api *router) setMiddleware() {
	api.engine.Use(gin.Recovery(), middleware.Log(), middleware.CorsCheck(), middleware.HandleError())
}

func (api *router) setRoute() {
	var routes = map[string]map[string][]any{
		"/webapi": {
			"/config": {func(c *gin.Context) {
				c.JSON(http.StatusOK, config.Config())
			}},
		},
	}

	for group, list := range routes {
		for path, fList := range list {
			routerGroup := api.engine.Group(fmt.Sprintf("%s%s", group, path), parseSliceAnyToSliceFunc(fList[:len(fList)-1])...)
			groupFunc := fList[len(fList)-1].(func(*gin.RouterGroup))
			groupFunc(routerGroup)
		}
	}
	api.engine.NoRoute(static.Web)
}

func parseSliceAnyToSliceFunc(functions []any) []gin.HandlerFunc {
	result := make([]gin.HandlerFunc, 0)
	for _, f := range functions {
		if fun, ok := f.(gin.HandlerFunc); ok {
			result = append(result, fun)
		}
	}
	return result
}
