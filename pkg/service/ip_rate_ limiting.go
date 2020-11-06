package service

import (
	"context"
	"dcard/pkg/model"
)

func (s *service) GetRateLimitCount(ctx context.Context, in model.IPRateLimit) (int, error) {
	return s.repo.GetRateLimitCount(ctx, in)
}

func (s *service) RateLimitIsAllow(ctx context.Context, in model.IPRateLimit) error {
	return s.repo.RateLimitIsAllow(ctx, in)
}
