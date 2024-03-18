package main

import (
	"go.uber.org/zap"
)

type cmd struct {
	logger        *zap.Logger
	noTraceLogger *zap.Logger
}
