package handler

import (
	"github.com/wheissd/mkgo/internal/examples/catsadventure/api/ogen"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/api/service"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/config"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen"
	"go.uber.org/zap"
)

var _ ogen.Handler = &Handler{}

type Handler struct {
	*ogen.HandlerImpl
	client  *gen.Client
	config  config.Config
	service *service.Service
}

func New(
	client *gen.Client,
	logger *slog.Logger,
	config config.Config,
	svc *service.Service,
) *Handler {
	hi := ogen.NewHandler(client, svc, logger)
	return &Handler{
		HandlerImpl: hi,
		client:      client,
		config:      config,
		service:     svc,
	}
}
