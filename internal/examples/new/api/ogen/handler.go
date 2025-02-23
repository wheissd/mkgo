package ogen

import (
	"context"
	"net/http"

	"github.com/wheissd/mkgo/internal/examples/new/api/service"
	"github.com/wheissd/mkgo/internal/examples/new/ent/gen"
	"go.uber.org/zap"
)

type HandlerImpl struct {
	client  *gen.Client
	service *service.Service
	logger  *slog.Logger
}

func NewHandler(
	client *gen.Client,
	srvc *service.Service, logger *slog.Logger,
) *HandlerImpl {
	return &HandlerImpl{
		client:  client,
		service: srvc,
		logger:  logger,
	}
}
func (h *HandlerImpl) NewError(ctx context.Context, err error) *ErrorStatusCode {
	h.logger.Error("Handler error", slog.Any("error", err))
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

func (h *HandlerImpl) CreateDefaultModel(ctx context.Context, req *CreateDefaultModel) (CreateDefaultModelRes, error) {
	defaultmodel, err := h.service.CreateDefaultModel(
		ctx,
		service.CreateDefaultModel{
			ID:   req.ID,
			Name: req.Name,
		},
	)
	if err != nil {
		return nil, err
	}
	return NewDefaultModel(defaultmodel), nil
}

func (h *HandlerImpl) UpdateDefaultModel(ctx context.Context, req *UpdateDefaultModel, params UpdateDefaultModelParams) (UpdateDefaultModelRes, error) {
	update := service.UpdateDefaultModel{}
	if req.Name.IsSet() {
		update.Name.Set(req.Name.Value)
	}
	serviceRes, err := h.service.UpdateDefaultModel(ctx, params.ID, update)
	if err != nil {
		return nil, err
	}
	return NewDefaultModel(serviceRes), nil
}

// DeleteDefaultModel handles DELETE /defaultmodel/{id} requests.
func (h *HandlerImpl) DeleteDefaultModel(ctx context.Context, params DeleteDefaultModelParams) (DeleteDefaultModelRes, error) {
	err := h.service.DeleteDefaultModel(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return new(DeleteDefaultModelOK), nil
}

func (h *HandlerImpl) ListDefaultModel(ctx context.Context, params ListDefaultModelParams) (ListDefaultModelRes, error) {
	sParams := service.DefaultModelListParams{}
	if params.Page.IsSet() {
		sParams.Page.Set(params.Page.Value)
	}
	if params.ItemsPerPage.IsSet() {
		sParams.ItemsPerPage.Set(params.ItemsPerPage.Value)
	}
	list, err := h.service.ListDefaultModel(ctx, sParams)
	if err != nil {
		return nil, err
	}
	response := DefaultModelListHeaders{
		Response: NewDefaultModelList(list.Items),
	}
	return &response, nil
}

func (h *HandlerImpl) ReadDefaultModel(ctx context.Context, params ReadDefaultModelParams) (ReadDefaultModelRes, error) {
	e, err := h.service.ReadDefaultModel(ctx, service.DefaultModelReadParams{ID: params.ID})
	if err != nil {
		return nil, err
	}
	return NewDefaultModel(e), nil
}
