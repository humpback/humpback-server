package app

import (
	"context"

	"humpback/api"
	"humpback/api/static"
	"humpback/internal/controller"
	"humpback/internal/db"
	"humpback/scheduler"
)

type App struct {
	webSite   *api.Router
	scheduler *scheduler.HumpbackScheduler
	stopCh    chan struct{}
}

func InitApp() (*App, error) {
	scheduler := scheduler.NewHumpbackScheduler()
	app := &App{
		webSite:   api.InitRouter(scheduler.NodeHeartbeatChan, scheduler.ServiceChangeChan),
		scheduler: scheduler,
		stopCh:    make(chan struct{}),
	}
	if err := db.InitDB(); err != nil {
		return nil, err
	}
	if err := static.InitStaticsResource(); err != nil {
		return nil, err
	}
	if err := controller.InitData(); err != nil {
		return nil, err
	}
	return app, nil
}

func (app *App) Startup() {
	app.scheduler.Start()
	app.webSite.Start()
	controller.Start(app.stopCh)
}

func (app *App) Close(c context.Context) error {
	close(app.stopCh)
	if err := app.webSite.Close(c); err != nil {
		return err
	}
	if err := app.scheduler.Close(c); err != nil {
		return err
	}
	return nil
}
