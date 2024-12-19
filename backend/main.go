package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"humpback/app"
	"humpback/config"
	"humpback/pkg/glog"
	"humpback/pkg/utils"
)

type (
	StartupFunc    func() error
	SignalQuitFunc func()
)

func init() {
	glog.Open(glog.WithOutputSource(glog.OutputTypeStd), glog.WithOutputFormat(glog.OutputFormatDefault))

	if err := config.InitConfig(); err != nil {
		panic(err)
	}
	slog.Info("[Config] init completed.")
	utils.PrintJson(config.Config())
}

func main() {
	if err := app.InitApp(); err != nil {
		panic(err)
	}
	slog.Info("[APP] new completed.")
	InitSignal(app.Startup, app.Close)
}

func InitSignal(startupF StartupFunc, closeF SignalQuitFunc) {
	defer glog.Close()
	defer slog.Info("[Exit] Process Exit... Over")
	if startupF == nil {
		slog.Error("[Startup] StartupFunc Is Nil.")
		return
	}
	ch := make(chan os.Signal, 1)
	defer close(ch)
	go func() {
		if err := startupF(); err != nil {
			slog.Error("[Startup] " + err.Error())
			ch <- syscall.SIGQUIT
		}
	}()
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	if closeF != nil {
		closeF()
	}
}
