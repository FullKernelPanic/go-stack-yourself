package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"go-stack-yourself/src/common/infrastructure"
	webCommon "go-stack-yourself/src/common/ports/web"
	"go-stack-yourself/src/common/services"
	webRoll "go-stack-yourself/src/roll/ports/web"
)

var (
	config *services.Config
)

func init() {
	services.SetConfigPath(os.Getenv("APP_CONFIG_PATH"))
	config = services.GetConfig()
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	otelShutdown, err := infrastructure.SetupOTelSDK(ctx, config.Infrastructure.Monitoring.OtelCollectorHost)

	if err != nil {
		return
	}

	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	srv := infrastructure.NewHTTPServer(ctx, config.App.Web.Port, newHTTPHandler())

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	select {
	case err = <-srvErr:
		return
	case <-ctx.Done():
		stop()
	}

	err = srv.Shutdown(context.Background())
	return
}

func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", webRoll.RollHomeHandler)
	mux.Handle("/rolldice/", http.StripPrefix("/rolldice", webRoll.NewRouting()))

	return http.StripPrefix(config.App.Web.BasePath, webCommon.DecorateWithMiddlewares(mux))
}
