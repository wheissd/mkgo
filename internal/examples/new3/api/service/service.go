package service

import (
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Service {
	return &Service{
		logger: logger,
	}
}
