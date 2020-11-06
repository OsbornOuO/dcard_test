package repository

import (
	"context"
	"dcard/internal/errors"
	"dcard/pkg/model"

	"github.com/go-redis/redis/v8"
)

func (repo *repository) GetRateLimitCount(ctx context.Context, in model.IPRateLimit) (int, error) {
	keys, err := repo.redisClient.Keys(ctx, in.SearchKey()).Result()
	if err != nil {
		return 0, err
	}

	return len(keys), nil
}

func (repo *repository) RateLimitIsAllow(ctx context.Context, in model.IPRateLimit) error {
	_, err := repo.redisClient.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		var (
			err  error
			keys []string
		)
		if keys, err = repo.redisClient.Keys(ctx, in.SearchKey()).Result(); err != nil {
			return err
		}
		if len(keys) >= in.RateCount {
			return errors.ErrIPRateLimiting
		}

		if _, err = pipe.SetEX(ctx, in.GenerateKey(), "", in.RateSec).Result(); err != nil {
			return errors.NewWithMessagef(errors.ErrInternalError, err.Error())
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
