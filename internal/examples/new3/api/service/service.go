package service

import (
	"go.uber.org/zap"
)

type Service struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Service {
	return &Service{
		logger: logger,
	}
}
