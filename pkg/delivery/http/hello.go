package http

import (
	"dcard/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello ...
func (h *Handler) Hello(c echo.Context) error {
	ctx := c.Request().Context()
	count, err := h.service.GetRateLimitCount(ctx, model.IPRateLimit{
		IP: c.RealIP(),
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": echo.Map{
			"count": count,
		},
	})
}
