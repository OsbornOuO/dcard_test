package redis

import (
	"context"
	"time"

	"github.com/cenk/backoff"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

// Config ...
type Config struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// New ...
func New(cfg *Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Username: cfg.Username,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	err := backoff.Retry(func() error {
		status := redisClient.Ping(context.Background())
		if status.Err() != nil {
			log.Error().Msgf("redis: %+v", status.Err())
			return status.Err()
		}

		return nil
	}, bo)
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}
