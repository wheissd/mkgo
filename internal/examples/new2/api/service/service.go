package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/wheissd/mkgo/internal/examples/new2/ent/gen"
	"github.com/wheissd/mkgo/internal/examples/new2/ent/gen/human"
	"go.uber.org/zap"
)

type Service struct {
	client                 *gen.Client
	logger                 *slog.Logger
	HumanListQueryModifier HumanQueryModifier
	HumanReadQueryModifier HumanQueryModifier
}

func New(client *gen.Client, logger *slog.Logger) *Service {
	return &Service{
		client:                 client,
		logger:                 logger,
		HumanListQueryModifier: noOpHumanQueryModifier,
		HumanReadQueryModifier: noOpHumanQueryModifier,
	}
}

func (h *Service) CreateHuman(ctx context.Context, req CreateHuman) (Human, error) {
	b := h.client.Human.Create()

	// Add all fields.
	b.SetID(req.ID)
	b.SetName(req.Name)

	// Set relations
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return Human{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Human.Query().Where(human.ID(e.ID))

	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return Human{}, err
	}
	response := NewHuman(e)
	return response, nil
}

func (h *Service) UpdateHuman(ctx context.Context, id uuid.UUID, req UpdateHuman) (Human, error) {
	b := h.client.Human.UpdateOneID(id)

	// Add all fields.
	if req.Name.IsSet() {
		b.SetName(req.Name.Get())
	}
	// Set relations.
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return Human{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Human.Query().Where(human.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return Human{}, err
	}
	response := NewHuman(e)
	return response, nil
}

// DeleteHuman handles DELETE /human/{id} requests.
func (h *Service) DeleteHuman(ctx context.Context, id uuid.UUID) error {
	err := h.client.Human.DeleteOneID(id).Exec(ctx)
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

func (h *Service) ListHuman(ctx context.Context, params HumanListParams) (HumanList, error) {
	q := h.client.Human.Query()
	if err := h.HumanListQueryModifier(ctx, q); err != nil {
		return HumanList{}, err
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
			return HumanList{}, err
		case gen.IsNotSingular(err):
			return HumanList{}, err
		default:
			// Let the server handle the error.
			return HumanList{}, err
		}
	}
	r := NewHumanList(es)
	response := HumanList{}
	response.Items = r

	return response, nil
}

func (h *Service) ReadHuman(ctx context.Context, params HumanReadParams) (Human, error) {
	q := h.client.Human.Query().Where(human.IDEQ(params.ID))
	if err := h.HumanReadQueryModifier(ctx, q); err != nil {
		return Human{}, err
	}
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case gen.IsNotFound(err):
			return Human{}, err
		case gen.IsNotSingular(err):
			return Human{}, err
		default:
			// Let the server handle the error.
			return Human{}, err
		}
	}

	return NewHuman(e), nil
}
