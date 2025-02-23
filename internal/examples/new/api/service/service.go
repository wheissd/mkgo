package service

import (
	"context"

	"go.uber.org/fx"

	"github.com/wheissd/mkgo/internal/examples/new/ent/gen"
	"github.com/wheissd/mkgo/internal/examples/new/ent/gen/defaultmodel"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

type Service struct {
	client                        *gen.Client
	logger                        *slog.Logger
	DefaultModelListQueryModifier DefaultModelQueryModifier
	DefaultModelReadQueryModifier DefaultModelQueryModifier
}

func New(client *gen.Client, logger *slog.Logger) *Service {
	return &Service{
		client:                        client,
		logger:                        logger,
		DefaultModelListQueryModifier: noOpDefaultModelQueryModifier,
		DefaultModelReadQueryModifier: noOpDefaultModelQueryModifier,
	}
}

func (h *Service) CreateDefaultModel(ctx context.Context, req CreateDefaultModel) (DefaultModel, error) {
	b := h.client.DefaultModel.Create()

	// Add all fields.
	b.SetID(req.ID)
	b.SetName(req.Name)

	// Set relations
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return DefaultModel{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.DefaultModel.Query().Where(defaultmodel.ID(e.ID))

	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return DefaultModel{}, err
	}
	response := NewDefaultModel(e)
	return response, nil
}

func (h *Service) UpdateDefaultModel(ctx context.Context, id uuid.UUID, req UpdateDefaultModel) (DefaultModel, error) {
	b := h.client.DefaultModel.UpdateOneID(id)

	// Add all fields.
	if req.Name.IsSet() {
		b.SetName(req.Name.Get())
	}
	// Set relations.
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return DefaultModel{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.DefaultModel.Query().Where(defaultmodel.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return DefaultModel{}, err
	}
	response := NewDefaultModel(e)
	return response, nil
}

// DeleteDefaultModel handles DELETE /defaultmodel/{id} requests.
func (h *Service) DeleteDefaultModel(ctx context.Context, id uuid.UUID) error {
	err := h.client.DefaultModel.DeleteOneID(id).Exec(ctx)
	if err != nil {
		switch {
		case gen.IsNotFound(err):
			return err
		case gen.IsConstraintError(err):
			return err
		default:
			// Let the server handle the error.
			return err
		}
	}
	return nil
}

func (h *Service) ListDefaultModel(ctx context.Context, params DefaultModelListParams) (DefaultModelList, error) {
	q := h.client.DefaultModel.Query()
	if err := h.DefaultModelListQueryModifier(ctx, q); err != nil {
		return DefaultModelList{}, err
	}
	page := 1
	if params.Page.IsSet() {
		page = params.Page.Get()
	}
	itemsPerPage := 30
	if params.ItemsPerPage.IsSet() {
		itemsPerPage = params.ItemsPerPage.Get()
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case gen.IsNotFound(err):
			return DefaultModelList{}, err
		case gen.IsNotSingular(err):
			return DefaultModelList{}, err
		default:
			// Let the server handle the error.
			return DefaultModelList{}, err
		}
	}
	r := NewDefaultModelList(es)
	response := DefaultModelList{}
	response.Items = r

	return response, nil
}

func (h *Service) ReadDefaultModel(ctx context.Context, params DefaultModelReadParams) (DefaultModel, error) {
	q := h.client.DefaultModel.Query().Where(defaultmodel.IDEQ(params.ID))
	if err := h.DefaultModelReadQueryModifier(ctx, q); err != nil {
		return DefaultModel{}, err
	}
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case gen.IsNotFound(err):
			return DefaultModel{}, err
		case gen.IsNotSingular(err):
			return DefaultModel{}, err
		default:
			// Let the server handle the error.
			return DefaultModel{}, err
		}
	}

	return NewDefaultModel(e), nil
}

var Module = fx.Module(
	"service",
	fx.Provide(
		New,
	),
)
