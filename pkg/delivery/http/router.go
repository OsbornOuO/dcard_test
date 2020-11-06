package http

import (
	"github.com/labstack/echo/v4"
)

// SetRoutes ...
func SetRoutes(e *echo.Echo, h *Handler) {
	rootV1 := e.Group("/v1", h.IPRateLimitMiddleware())

	rootV1.GET("/hello", h.Hello)
}
