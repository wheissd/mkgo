package ogen

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/wheissd/mkgo/internal/examples/catsadventure/api/service"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/internal/examples/catsadventure/ent/gen"
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
	var errStatusCode *ErrorStatusCode
	if errors.As(err, &errStatusCode) {
		return errStatusCode
	}
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

func (h *HandlerImpl) CreateBreed(ctx context.Context, req *CreateBreed) (CreateBreedRes, error) {
	breed, err := h.service.CreateBreed(
		ctx,
		service.CreateBreed{
			Name: req.Name,
		},
	)
	if err != nil {
		return nil, err
	}
	res := NewBreed(breed)
	return &res, nil
}

func (h *HandlerImpl) UpdateBreed(ctx context.Context, req *UpdateBreed, params UpdateBreedParams) (UpdateBreedRes, error) {
	update := service.UpdateBreed{}
	if req.Name.IsSet() {
		update.Name.Set(req.Name.Value)
	}
	serviceRes, err := h.service.UpdateBreed(ctx, params.ID, update)
	if err != nil {
		return nil, err
	}
	res := NewBreed(serviceRes)
	return &res, nil
}

// DeleteBreed handles DELETE /breed/{id} requests.
func (h *HandlerImpl) DeleteBreed(ctx context.Context, params DeleteBreedParams) (DeleteBreedRes, error) {
	err := h.service.DeleteBreed(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return new(DeleteBreedOK), nil
}

func (h *HandlerImpl) ListBreed(ctx context.Context, params ListBreedParams) (ListBreedRes, error) {
	sParams := service.BreedListParams{}
	if params.Page.IsSet() {
		sParams.Page.Set(params.Page.Value)
	}
	if params.ItemsPerPage.IsSet() {
		sParams.ItemsPerPage.Set(params.ItemsPerPage.Value)
	}

	var (
		err error
	)
	if params.With.IsSet() {
		topLevelEdges := strings.Split(params.With.Value, ",")
		for i := range topLevelEdges {
			switch topLevelEdges[i] {
			case "cats":
				sParams.With.Cats, err = parseWithCat(0, topLevelEdges[i], nil)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	list, err := h.service.ListBreed(ctx, sParams)
	if err != nil {
		return nil, err
	}
	response := BreedListHeaders{
		Response: NewBreedList(list.Items),
	}
	return &response, nil
}

func (h *HandlerImpl) ReadBreed(ctx context.Context, params ReadBreedParams) (ReadBreedRes, error) {
	e, err := h.service.ReadBreed(ctx, service.BreedReadParams{ID: params.ID})
	if err != nil {
		return nil, err
	}
	res := NewBreed(e)
	return &res, nil
}

func (h *HandlerImpl) CreateCat(ctx context.Context, req *CreateCat) (CreateCatRes, error) {
	cat, err := h.service.CreateCat(
		ctx,
		service.CreateCat{
			Name:      req.Name,
			Speed:     req.Speed,
			Type:      req.Type,
			BreedID:   req.BreedID,
			DateFrom:  req.DateFrom,
			OtherType: req.OtherType,
		},
	)
	if err != nil {
		return nil, err
	}
	res := NewCat(cat)
	return &res, nil
}

func (h *HandlerImpl) UpdateCat(ctx context.Context, req *UpdateCat, params UpdateCatParams) (UpdateCatRes, error) {
	update := service.UpdateCat{}
	if req.Name.IsSet() {
		update.Name.Set(req.Name.Value)
	}
	if req.Speed.IsSet() {
		update.Speed.Set(req.Speed.Value)
	}
	if req.Type.IsSet() {
		update.Type.Set(req.Type.Value)
	}
	if req.BreedID.IsSet() {
		update.BreedID.Set(req.BreedID.Value)
	}
	if req.DateFrom.IsSet() {
		update.DateFrom.Set(req.DateFrom.Value)
	}
	if req.OtherType.IsSet() {
		update.OtherType.Set(req.OtherType.Value)
	}
	serviceRes, err := h.service.UpdateCat(ctx, params.ID, update)
	if err != nil {
		return nil, err
	}
	res := NewCat(serviceRes)
	return &res, nil
}

// DeleteCat handles DELETE /cat/{id} requests.
func (h *HandlerImpl) DeleteCat(ctx context.Context, params DeleteCatParams) (DeleteCatRes, error) {
	err := h.service.DeleteCat(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return new(DeleteCatOK), nil
}

func (h *HandlerImpl) ListCat(ctx context.Context, params ListCatParams) (ListCatRes, error) {
	sParams := service.CatListParams{}
	if params.Page.IsSet() {
		sParams.Page.Set(params.Page.Value)
	}
	if params.ItemsPerPage.IsSet() {
		sParams.ItemsPerPage.Set(params.ItemsPerPage.Value)
	}

	var (
		err error
	)
	if params.With.IsSet() {
		topLevelEdges := strings.Split(params.With.Value, ",")
		for i := range topLevelEdges {
			switch topLevelEdges[i] {
			case "kittens":
				sParams.With.Kittens, err = parseWithKitten(0, topLevelEdges[i], nil)
				if err != nil {
					return nil, err
				}
			case "breed":
				sParams.With.Breed, err = parseWithBreed(0, topLevelEdges[i], nil)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	list, err := h.service.ListCat(ctx, sParams)
	if err != nil {
		return nil, err
	}
	response := CatListHeaders{
		Response: NewCatList(list.Items),
	}
	return &response, nil
}

func (h *HandlerImpl) ReadCat(ctx context.Context, params ReadCatParams) (ReadCatRes, error) {
	e, err := h.service.ReadCat(ctx, service.CatReadParams{ID: params.ID})
	if err != nil {
		return nil, err
	}
	res := NewCat(e)
	return &res, nil
}

func (h *HandlerImpl) CreateFatherCat(ctx context.Context, req *CreateFatherCat) (CreateFatherCatRes, error) {
	fathercat, err := h.service.CreateFatherCat(
		ctx,
		service.CreateFatherCat{
			Name: req.Name,
		},
	)
	if err != nil {
		return nil, err
	}
	res := NewFatherCat(fathercat)
	return &res, nil
}

func (h *HandlerImpl) UpdateFatherCat(ctx context.Context, req *UpdateFatherCat, params UpdateFatherCatParams) (UpdateFatherCatRes, error) {
	update := service.UpdateFatherCat{}
	if req.Name.IsSet() {
		update.Name.Set(req.Name.Value)
	}
	serviceRes, err := h.service.UpdateFatherCat(ctx, params.ID, update)
	if err != nil {
		return nil, err
	}
	res := NewFatherCat(serviceRes)
	return &res, nil
}

// DeleteFatherCat handles DELETE /fathercat/{id} requests.
func (h *HandlerImpl) DeleteFatherCat(ctx context.Context, params DeleteFatherCatParams) (DeleteFatherCatRes, error) {
	err := h.service.DeleteFatherCat(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return new(DeleteFatherCatOK), nil
}

func (h *HandlerImpl) ListFatherCat(ctx context.Context, params ListFatherCatParams) (ListFatherCatRes, error) {
	sParams := service.FatherCatListParams{}
	if params.Page.IsSet() {
		sParams.Page.Set(params.Page.Value)
	}
	if params.ItemsPerPage.IsSet() {
		sParams.ItemsPerPage.Set(params.ItemsPerPage.Value)
	}

	list, err := h.service.ListFatherCat(ctx, sParams)
	if err != nil {
		return nil, err
	}
	response := FatherCatListHeaders{
		Response: NewFatherCatList(list.Items),
	}
	return &response, nil
}

func (h *HandlerImpl) ReadFatherCat(ctx context.Context, params ReadFatherCatParams) (ReadFatherCatRes, error) {
	e, err := h.service.ReadFatherCat(ctx, service.FatherCatReadParams{ID: params.ID})
	if err != nil {
		return nil, err
	}
	res := NewFatherCat(e)
	return &res, nil
}

func (h *HandlerImpl) CreateKitten(ctx context.Context, req *CreateKitten) (CreateKittenRes, error) {
	kitten, err := h.service.CreateKitten(
		ctx,
		service.CreateKitten{
			Name:     req.Name,
			MotherID: req.MotherID,
		},
	)
	if err != nil {
		return nil, err
	}
	res := NewKitten(kitten)
	return &res, nil
}

func (h *HandlerImpl) UpdateKitten(ctx context.Context, req *UpdateKitten, params UpdateKittenParams) (UpdateKittenRes, error) {
	update := service.UpdateKitten{}
	if req.Name.IsSet() {
		update.Name.Set(req.Name.Value)
	}
	if req.MotherID.IsSet() {
		update.MotherID.Set(req.MotherID.Value)
	}
	serviceRes, err := h.service.UpdateKitten(ctx, params.ID, update)
	if err != nil {
		return nil, err
	}
	res := NewKitten(serviceRes)
	return &res, nil
}

// DeleteKitten handles DELETE /kitten/{id} requests.
func (h *HandlerImpl) DeleteKitten(ctx context.Context, params DeleteKittenParams) (DeleteKittenRes, error) {
	err := h.service.DeleteKitten(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return new(DeleteKittenOK), nil
}

func (h *HandlerImpl) ListKitten(ctx context.Context, params ListKittenParams) (ListKittenRes, error) {
	sParams := service.KittenListParams{}
	if params.Page.IsSet() {
		sParams.Page.Set(params.Page.Value)
	}
	if params.ItemsPerPage.IsSet() {
		sParams.ItemsPerPage.Set(params.ItemsPerPage.Value)
	}

	var (
		err error
	)
	if params.With.IsSet() {
		topLevelEdges := strings.Split(params.With.Value, ",")
		for i := range topLevelEdges {
			switch topLevelEdges[i] {
			case "mother":
				sParams.With.Mother, err = parseWithCat(0, topLevelEdges[i], nil)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	list, err := h.service.ListKitten(ctx, sParams)
	if err != nil {
		return nil, err
	}
	response := KittenListHeaders{
		Response: NewKittenList(list.Items),
	}
	return &response, nil
}

func (h *HandlerImpl) ReadKitten(ctx context.Context, params ReadKittenParams) (ReadKittenRes, error) {
	e, err := h.service.ReadKitten(ctx, service.KittenReadParams{ID: params.ID})
	if err != nil {
		return nil, err
	}
	res := NewKitten(e)
	return &res, nil
}
