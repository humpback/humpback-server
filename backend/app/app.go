package app

import (
	"context"

	"humpback/api"
	"humpback/internal/db"
	"humpback/scheduler"
)

type App struct {
	webSite   *api.Router
	scheduler *scheduler.HumpbackScheduler
}

func InitApp() *App {
	app := &App{
		webSite:   api.InitRouter(),
		scheduler: scheduler.NewHumpbackScheduler(),
	}

	db.InitDB()
	return app
}

func (app *App) Startup() {
	app.scheduler.Start()
	app.webSite.Start()
}

func (app *App) Close(c context.Context) error {
	if err := app.webSite.Close(c); err != nil {
		return err
	}
	if err := app.scheduler.Close(c); err != nil {
		return err
	}
	return nil
}
