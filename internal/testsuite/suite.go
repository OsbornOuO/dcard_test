package testsuite

import (
	"context"
	"dcard/configuration"
	"dcard/internal/redis"
	"os"
	"strconv"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/rs/zerolog/log"
	"gitlab.com/howmay/gopher/zlog"
	"go.uber.org/fx"
)

// Suite ...
type Suite struct {
	app            *fx.App
	redisContainer *dockertest.Resource
	pool           *dockertest.Pool
	t              testing.T
}

var suite Suite

// Initialize 初始化 suite
func Initialize(fxOption ...fx.Option) error {
	if os.Getenv("CONFIG_NAME") == "" {
		_ = os.Setenv("CONFIG_NAME", "app-test")
	}
	configuration, err := configuration.New()
	if err != nil {
		return err
	}

	suite.pool, err = dockertest.NewPool("")
	if err != nil {
		log.Error().Msgf("Could not connect to docker: %s", err)
		return err
	}
	suite.redisContainer, err = suite.pool.Run("redis", "latest", []string{})
	if err != nil {
		log.Error().Msgf("Could not start resource: %s", err)
		return err
	}
	redisPort, _ := strconv.Atoi(suite.redisContainer.GetPort("6379/tcp"))

	configuration.Redis = &redis.Config{
		Address: "localhost:" + strconv.Itoa(redisPort),
	}

	base := []fx.Option{
		fx.Supply(*configuration),
		fx.Provide(
			redis.New,
		),
		fx.Invoke(zlog.InitV2),
	}

	base = append(base, fxOption...)

	app := fx.New(
		base...,
	)

	suite.app = app
	return app.Start(context.Background())
}

// Close 停止 container
func Close() {
	log.Info().Msg("close app")
	if err := suite.pool.Purge(suite.redisContainer); err != nil {
		log.Error().Msgf("Could not purge resource: %s", err)
	}
}
