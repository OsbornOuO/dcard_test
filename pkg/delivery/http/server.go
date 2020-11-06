package http

import (
	"dcard/configuration"
	server "dcard/pkg"
)

// Handler http handler ...
type Handler struct {
	cfg     *configuration.App
	service server.IService
}

// NewHandler http handler 依賴注入
func NewHandler(service server.IService, cfg *configuration.App) *Handler {
	server := Handler{
		service: service,
		cfg:     cfg,
	}
	return &server
}
