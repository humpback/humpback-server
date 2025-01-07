package app

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"humpback/api"
	"humpback/api/static"
	"humpback/config"
	"humpback/internal/db"
	"humpback/scheduler"
	"humpback/types"
)

type App struct {
	webSite   *api.Router
	scheduler *scheduler.HumpbackScheduler
}

func InitApp() (*App, error) {
	app := &App{
		webSite:   api.InitRouter(),
		scheduler: scheduler.NewHumpbackScheduler(),
	}
	if err := db.InitDB(); err != nil {
		return nil, err
	}
	if err := static.InitStaticsResource(); err != nil {
		return nil, err
	}
	if err := initAccount(); err != nil {
		return nil, err
	}
	return app, nil
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

func initAccount() error {
	adminConfig := config.AdminArgs()
	_, err := db.UserGetById(adminConfig.Id)
	if err != nil {
		if err != db.ErrKeyNotExist {
			return fmt.Errorf("Check admin account failed: %s", err)
		}
		t := time.Now().Unix()
		if err = db.UserUpdate(adminConfig.Id, &types.User{
			UserID:    adminConfig.Id,
			UserName:  adminConfig.Name,
			Email:     "",
			Password:  adminConfig.Password,
			Phone:     "",
			IsAdmin:   true,
			CreatedAt: t,
			UpdatedAt: t,
			Groups:    nil,
		}); err != nil {
			return fmt.Errorf("Create admin account failed: %s", err)
		}
	}
	slog.Info("Admin account check success")
	return nil
}
