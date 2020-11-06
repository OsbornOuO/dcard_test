package http

import (
	"dcard/pkg/model"
	"time"

	"github.com/labstack/echo/v4"
)

// IPRateLimitMiddleware ...
func (h *Handler) IPRateLimitMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			ctx := c.Request().Context()

			if err := h.service.RateLimitIsAllow(ctx, model.IPRateLimit{
				IP:        c.RealIP(),
				RateCount: h.cfg.IPRatelimitingCount,
				RateSec:   time.Duration(h.cfg.IPRatelimitingSec) * time.Second,
			}); err != nil {
				return err
			}

			return next(c)
		}
	}
}
