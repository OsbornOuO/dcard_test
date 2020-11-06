package repository

import (
	pkg "dcard/pkg"

	"github.com/go-redis/redis/v8"
)

type repository struct {
	redisClient *redis.Client
}

// NewRepository 依賴注入
func NewRepository(redisClient *redis.Client) pkg.IRepository {
	return &repository{
		redisClient: redisClient,
	}
}
