package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"humpback/app"
	"humpback/config"
	"humpback/pkg/glog"
	"humpback/pkg/utils"
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
	application := app.InitApp()

	slog.Info("[APP] new completed.")

	// Start server
	application.Startup()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := application.Close(ctx); err != nil {
		slog.Error(err.Error())
	}

	slog.Info("[App] quit...")

}
