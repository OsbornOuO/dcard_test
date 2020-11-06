package cmd

import (
	"context"
	"dcard/configuration"
	"dcard/internal/http"
	"dcard/internal/redis"
	pkgHTTP "dcard/pkg/delivery/http"
	repository "dcard/pkg/repository"
	service "dcard/pkg/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/rs/zerolog/log"
	cobra "github.com/spf13/cobra"
	"gitlab.com/howmay/gopher/zlog"
	fx "go.uber.org/fx"
)

// ServerCmd 是此程式的Service入口點
var ServerCmd = &cobra.Command{
	Run: run,
	Use: "server",
}

func run(cmd *cobra.Command, args []string) {
	defer cmdRecover()

	config, err := configuration.New()
	if err != nil {
		os.Exit(0)
		return
	}

	zlog.InitV2(config.Log)

	app := fx.New(
		fx.Supply(*config),
		fx.Provide(
			redis.New,
			repository.NewRepository,
			service.NewService,
			pkgHTTP.NewHandler,
			http.StartEcho,
		),
		fx.Invoke(pkgHTTP.SetRoutes),
	)

	exitCode := 0
	if err := app.Start(context.Background()); err != nil {
		log.Error().Msg(err.Error())
		os.Exit(exitCode)
		return
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM)
	<-stopChan
	log.Info().Msg("main: shutting down server...")

	stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Error().Msg(err.Error())
	}

	os.Exit(exitCode)
}
