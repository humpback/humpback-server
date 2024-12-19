package app

import (
	"humpback/api"
)

func InitApp() error {
	api.InitRouter()
	return nil
}

func Startup() error {
	return api.Router.Start()
}

func Close() {}
