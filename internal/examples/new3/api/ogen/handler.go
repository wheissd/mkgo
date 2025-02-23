package ogen

import (
	"go.uber.org/zap"
)

type HandlerImpl struct {
	logger *slog.Logger
}

func NewHandler(logger *slog.Logger,
) *HandlerImpl {
	return &HandlerImpl{
		logger: logger,
	}
}
