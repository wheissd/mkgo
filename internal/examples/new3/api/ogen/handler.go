package ogen

import (
	"go.uber.org/zap"
)

type HandlerImpl struct {
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger,
) *HandlerImpl {
	return &HandlerImpl{
		logger: logger,
	}
}
