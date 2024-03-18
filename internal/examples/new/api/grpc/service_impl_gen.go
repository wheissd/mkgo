package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/new/api/service"
)

type DefaultModelServiceImpl struct {
	service *service.Service
	UnimplementedDefaultModelServiceServer
}

func New(srvc *service.Service) *DefaultModelServiceImpl {
	return &DefaultModelServiceImpl{
		service: srvc,
	}
}

func NewDefaultModel(v service.DefaultModel) DefaultModel {
	return DefaultModel{
		ID:   v.ID.String(),
		Name: v.Name,
	}
}

func (s *DefaultModelServiceImpl) ReadDefaultModel(ctx context.Context, r *ReadDefaultModelRequest) (*DefaultModel, error) {
	id, err := uuid.Parse(r.ID)
	if err != nil {
		return nil, err
	}
	e, err := s.service.ReadDefaultModel(ctx, service.DefaultModelReadParams{ID: id})
	if err != nil {
		return nil, err
	}
	response := NewDefaultModel(e)
	return &response, nil
}
