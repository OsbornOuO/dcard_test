package pkg

import (
	"context"
	"dcard/pkg/model"
)

// IPRateLimitingService 抽象，請在這定義要實作的方法
type IPRateLimitingService interface {
	GetRateLimitCount(ctx context.Context, in model.IPRateLimit) (int, error)
	RateLimitIsAllow(ctx context.Context, in model.IPRateLimit) error
}

// IPRateLimitingRepository 抽象，請在這定義要實作的方法
type IPRateLimitingRepository interface {
	GetRateLimitCount(ctx context.Context, in model.IPRateLimit) (int, error)
	RateLimitIsAllow(ctx context.Context, in model.IPRateLimit) error
}
