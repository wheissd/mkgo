package main

import (
	"net/http"

	"go.uber.org/zap"
)

type cmd struct {
	logger         *zap.Logger
	noTraceLogger  *zap.Logger
	depCheckClient *http.Client
}
