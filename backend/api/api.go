package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"humpback/api/handle"
	"humpback/api/middleware"
	"humpback/api/static"
	"humpback/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine  *gin.Engine
	httpSrv *http.Server
}

func InitRouter() *Router {
	gin.SetMode(gin.ReleaseMode)
	gin.Default()
	r := &Router{engine: gin.New()}
	r.setMiddleware()
	r.setRoute()
	return r
}

func (api *Router) Start() {
	go func() {
		listeningAddress := fmt.Sprintf("%s:%s", config.NodeArgs().HostIp, config.NodeArgs().SitePort)
		slog.Info("[Api] Listening...", "Address", listeningAddress)
		api.httpSrv = &http.Server{
			Addr:    listeningAddress,
			Handler: api.engine,
		}
		if err := api.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error(fmt.Sprintf("[Api] Listening %s failed: %s", listeningAddress, err))
		}
	}()
}

func (api *Router) Close(c context.Context) error {
	return api.httpSrv.Shutdown(c)
}

func (api *Router) setMiddleware() {
	api.engine.Use(middleware.Log(), middleware.CorsCheck(), middleware.HandleError())
}

func (api *Router) setRoute() {
	var routes = map[string]map[string][]any{
		"/webapi": {
			"/user": {handle.RouteUser},
			"/team": {middleware.CheckLogin(), handle.RouteTeam},
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
