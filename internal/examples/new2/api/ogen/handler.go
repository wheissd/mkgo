package ogen

import (
	"context"
	"net/http"

	"github.com/wheissd/mkgo/internal/examples/new2/api/service"
	"github.com/wheissd/mkgo/internal/examples/new2/ent/gen"
	"go.uber.org/zap"
)

type HandlerImpl struct {
	client  *gen.Client
	service *service.Service
	logger  *zap.Logger
}

func NewHandler(
	client *gen.Client,
	srvc *service.Service, logger *zap.Logger,
) *HandlerImpl {
	return &HandlerImpl{
		client:  client,
		service: srvc,
		logger:  logger,
	}
}
func (h *HandlerImpl) NewError(ctx context.Context, err error) *ErrorStatusCode {
	h.logger.Error("Handler error", zap.Error(err))
	return &ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: []ErrorItem{
			{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		},
	}
}

func (h *HandlerImpl) CreateHuman(ctx context.Context, req *CreateHuman) (CreateHumanRes, error) {
	human, err := h.service.CreateHuman(
		ctx,
		service.CreateHuman{
			ID:   req.ID,
			Name: req.Name,
		},
	)
	if err != nil {
		return nil, err
	}
	return NewHuman(human), nil
}

func (h *HandlerImpl) UpdateHuman(ctx context.Context, req *UpdateHuman, params UpdateHumanParams) (UpdateHumanRes, error) {
	update := service.UpdateHuman{}
	if req.Name.IsSet() {
		update.Name.Set(req.Name.Value)
	}
	serviceRes, err := h.service.UpdateHuman(ctx, params.ID, update)
	if err != nil {
		return nil, err
	}
	return NewHuman(serviceRes), nil
}

// DeleteHuman handles DELETE /human/{id} requests.
func (h *HandlerImpl) DeleteHuman(ctx context.Context, params DeleteHumanParams) (DeleteHumanRes, error) {
	err := h.service.DeleteHuman(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return new(DeleteHumanOK), nil
}

func (h *HandlerImpl) ListHuman(ctx context.Context, params ListHumanParams) (ListHumanRes, error) {
	sParams := service.HumanListParams{}
	if params.Page.IsSet() {
		sParams.Page.Set(params.Page.Value)
	}
	if params.ItemsPerPage.IsSet() {
		sParams.ItemsPerPage.Set(params.ItemsPerPage.Value)
	}
	list, err := h.service.ListHuman(ctx, sParams)
	if err != nil {
		return nil, err
	}
	response := HumanListHeaders{
		Response: NewHumanList(list.Items),
	}
	return &response, nil
}

func (h *HandlerImpl) ReadHuman(ctx context.Context, params ReadHumanParams) (ReadHumanRes, error) {
	e, err := h.service.ReadHuman(ctx, service.HumanReadParams{ID: params.ID})
	if err != nil {
		return nil, err
	}
	return NewHuman(e), nil
}
