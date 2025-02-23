package main

import (
	"log/slog"
	"net/http"
)

type cmd struct {
	logger         *slog.Logger
	depCheckClient *http.Client
}
