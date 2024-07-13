package handlers

import "log/slog"

type HealthCheckHandler struct {
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) Handle() error {

	slog.Info("service is healthy")

	return nil
}
