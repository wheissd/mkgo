package service

import (
	"context"

	"go.uber.org/fx"

	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/internal/examples/catsadventure/ent/gen"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/internal/examples/catsadventure/ent/gen/breed"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/internal/examples/catsadventure/ent/gen/cat"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/internal/examples/catsadventure/ent/gen/fathercat"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/internal/examples/catsadventure/ent/gen/kitten"
	"go.uber.org/zap"
)

type Service struct {
	client                     *gen.Client
	logger                     *slog.Logger
	BreedListQueryModifier     BreedQueryModifier
	BreedReadQueryModifier     BreedQueryModifier
	CatListQueryModifier       CatQueryModifier
	CatReadQueryModifier       CatQueryModifier
	FatherCatListQueryModifier FatherCatQueryModifier
	FatherCatReadQueryModifier FatherCatQueryModifier
	KittenListQueryModifier    KittenQueryModifier
	KittenReadQueryModifier    KittenQueryModifier
}

func New(client *gen.Client, logger *slog.Logger) *Service {
	return &Service{
		client:                     client,
		logger:                     logger,
		BreedListQueryModifier:     noOpBreedQueryModifier,
		BreedReadQueryModifier:     noOpBreedQueryModifier,
		CatListQueryModifier:       noOpCatQueryModifier,
		CatReadQueryModifier:       noOpCatQueryModifier,
		FatherCatListQueryModifier: noOpFatherCatQueryModifier,
		FatherCatReadQueryModifier: noOpFatherCatQueryModifier,
		KittenListQueryModifier:    noOpKittenQueryModifier,
		KittenReadQueryModifier:    noOpKittenQueryModifier,
	}
}

func (h *Service) CreateBreed(ctx context.Context, req CreateBreed) (Breed, error) {
	b := h.client.Breed.Create()

	// Add all fields.
	b.SetName(req.Name)

	// Set relations

	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return Breed{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Breed.Query().Where(breed.ID(e.ID))

	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return Breed{}, err
	}
	response := NewBreed(e)
	return response, nil
}
func (h *Service) UpdateBreed(ctx context.Context, id uuid.UUID, req UpdateBreed) (Breed, error) {
	b := h.client.Breed.UpdateOneID(id)

	// Add all fields.
	if req.Name.IsSet() {
		b.SetName(req.Name.Get())
	}
	// Set relations.

	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return Breed{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Breed.Query().Where(breed.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return Breed{}, err
	}
	response := NewBreed(e)
	return response, nil
}

// DeleteBreed handles DELETE /breed/{id} requests.
func (h *Service) DeleteBreed(ctx context.Context, id uuid.UUID) error {
	err := h.client.Breed.DeleteOneID(id).Exec(ctx)
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
func (h *Service) ListBreed(ctx context.Context, params BreedListParams) (BreedList, error) {
	q := h.client.Breed.Query()
	if err := h.BreedListQueryModifier(ctx, q); err != nil {
		return BreedList{}, err
	}
	if params.FilterCatsIDs.IsSet() {
		q.Where(breed.HasCatsWith(cat.IDIn(params.FilterCatsIDs.Get()...)))
	}

	if params.With != nil {
		if params.With.Cats != nil {
			q = q.WithCats(applyWithCat(params.With.Cats))
		}
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
			return BreedList{}, err
		case gen.IsNotSingular(err):
			return BreedList{}, err
		default:
			// Let the server handle the error.
			return BreedList{}, err
		}
	}
	r := NewBreedList(es)
	response := BreedList{}
	response.Items = r

	return response, nil
}
func (h *Service) ReadBreed(ctx context.Context, params BreedReadParams) (Breed, error) {

	q := h.client.Breed.Query().Where(breed.IDEQ(params.ID))
	if err := h.BreedReadQueryModifier(ctx, q); err != nil {
		return Breed{}, err
	}

	if params.With != nil {
		if params.With.Cats != nil {
			q = q.WithCats(applyWithCat(params.With.Cats))
		}
	}

	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case gen.IsNotFound(err):
			return Breed{}, err
		case gen.IsNotSingular(err):
			return Breed{}, err
		default:
			// Let the server handle the error.
			return Breed{}, err
		}
	}

	return NewBreed(e), nil
}

func (h *Service) CreateCat(ctx context.Context, req CreateCat) (Cat, error) {
	b := h.client.Cat.Create()

	// Add all fields.
	b.SetName(req.Name)
	b.SetSpeed(req.Speed)
	b.SetType(cat.Type(req.Type))
	b.SetBreedID(req.BreedID)
	b.SetDateFrom(req.DateFrom)
	b.SetOtherType(cat.OtherType(req.OtherType))

	// Set relations

	b.SetBreedID(req.BreedID)

	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return Cat{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Cat.Query().Where(cat.ID(e.ID))

	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return Cat{}, err
	}
	response := NewCat(e)
	return response, nil
}
func (h *Service) UpdateCat(ctx context.Context, id uuid.UUID, req UpdateCat) (Cat, error) {
	b := h.client.Cat.UpdateOneID(id)

	// Add all fields.
	if req.Name.IsSet() {
		b.SetName(req.Name.Get())
	}
	if req.Speed.IsSet() {
		b.SetSpeed(req.Speed.Get())
	}
	if req.Type.IsSet() {
		b.SetType(cat.Type(req.Type.Get()))
	}
	if req.BreedID.IsSet() {
		b.SetBreedID(req.BreedID.Get())
	}
	if req.DateFrom.IsSet() {
		b.SetDateFrom(req.DateFrom.Get())
	}
	if req.OtherType.IsSet() {
		b.SetOtherType(cat.OtherType(req.OtherType.Get()))
	}
	// Set relations.

	if req.BreedID.IsSet() {
		b.SetBreedID(req.BreedID.Get())
	}

	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return Cat{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Cat.Query().Where(cat.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return Cat{}, err
	}
	response := NewCat(e)
	return response, nil
}

// DeleteCat handles DELETE /cat/{id} requests.
func (h *Service) DeleteCat(ctx context.Context, id uuid.UUID) error {
	err := h.client.Cat.DeleteOneID(id).Exec(ctx)
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
func (h *Service) ListCat(ctx context.Context, params CatListParams) (CatList, error) {
	q := h.client.Cat.Query()
	if err := h.CatListQueryModifier(ctx, q); err != nil {
		return CatList{}, err
	}
	if params.FilterKittensIDs.IsSet() {
		q.Where(cat.HasKittensWith(kitten.IDIn(params.FilterKittensIDs.Get()...)))
	}
	if params.FilterBreedID.IsSet() {
		q.Where(cat.BreedIDEQ(params.FilterBreedID.Get()))
	}

	if params.With != nil {
		if params.With.Kittens != nil {
			q = q.WithKittens(applyWithKitten(params.With.Kittens))
		}
		if params.With.Breed != nil {
			q = q.WithBreed(applyWithBreed(params.With.Breed))
		}
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
			return CatList{}, err
		case gen.IsNotSingular(err):
			return CatList{}, err
		default:
			// Let the server handle the error.
			return CatList{}, err
		}
	}
	r := NewCatList(es)
	response := CatList{}
	response.Items = r

	return response, nil
}
func (h *Service) ReadCat(ctx context.Context, params CatReadParams) (Cat, error) {

	q := h.client.Cat.Query().Where(cat.IDEQ(params.ID))
	if err := h.CatReadQueryModifier(ctx, q); err != nil {
		return Cat{}, err
	}

	if params.With != nil {
		if params.With.Kittens != nil {
			q = q.WithKittens(applyWithKitten(params.With.Kittens))
		}
		if params.With.Breed != nil {
			q = q.WithBreed(applyWithBreed(params.With.Breed))
		}
	}

	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case gen.IsNotFound(err):
			return Cat{}, err
		case gen.IsNotSingular(err):
			return Cat{}, err
		default:
			// Let the server handle the error.
			return Cat{}, err
		}
	}

	return NewCat(e), nil
}

func (h *Service) CreateFatherCat(ctx context.Context, req CreateFatherCat) (FatherCat, error) {
	b := h.client.FatherCat.Create()

	// Add all fields.
	b.SetName(req.Name)

	// Set relations
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return FatherCat{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.FatherCat.Query().Where(fathercat.ID(e.ID))

	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return FatherCat{}, err
	}
	response := NewFatherCat(e)
	return response, nil
}
func (h *Service) UpdateFatherCat(ctx context.Context, id uuid.UUID, req UpdateFatherCat) (FatherCat, error) {
	b := h.client.FatherCat.UpdateOneID(id)

	// Add all fields.
	if req.Name.IsSet() {
		b.SetName(req.Name.Get())
	}
	// Set relations.
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return FatherCat{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.FatherCat.Query().Where(fathercat.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return FatherCat{}, err
	}
	response := NewFatherCat(e)
	return response, nil
}

// DeleteFatherCat handles DELETE /fathercat/{id} requests.
func (h *Service) DeleteFatherCat(ctx context.Context, id uuid.UUID) error {
	err := h.client.FatherCat.DeleteOneID(id).Exec(ctx)
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
func (h *Service) ListFatherCat(ctx context.Context, params FatherCatListParams) (FatherCatList, error) {
	q := h.client.FatherCat.Query()
	if err := h.FatherCatListQueryModifier(ctx, q); err != nil {
		return FatherCatList{}, err
	}

	if params.With != nil {
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
			return FatherCatList{}, err
		case gen.IsNotSingular(err):
			return FatherCatList{}, err
		default:
			// Let the server handle the error.
			return FatherCatList{}, err
		}
	}
	r := NewFatherCatList(es)
	response := FatherCatList{}
	response.Items = r

	return response, nil
}
func (h *Service) ReadFatherCat(ctx context.Context, params FatherCatReadParams) (FatherCat, error) {

	q := h.client.FatherCat.Query().Where(fathercat.IDEQ(params.ID))
	if err := h.FatherCatReadQueryModifier(ctx, q); err != nil {
		return FatherCat{}, err
	}

	if params.With != nil {
	}

	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case gen.IsNotFound(err):
			return FatherCat{}, err
		case gen.IsNotSingular(err):
			return FatherCat{}, err
		default:
			// Let the server handle the error.
			return FatherCat{}, err
		}
	}

	return NewFatherCat(e), nil
}

func (h *Service) CreateKitten(ctx context.Context, req CreateKitten) (Kitten, error) {
	b := h.client.Kitten.Create()

	// Add all fields.
	b.SetName(req.Name)
	b.SetMotherID(req.MotherID)

	// Set relations

	b.SetMotherID(req.MotherID)

	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return Kitten{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Kitten.Query().Where(kitten.ID(e.ID))

	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return Kitten{}, err
	}
	response := NewKitten(e)
	return response, nil
}
func (h *Service) UpdateKitten(ctx context.Context, id uuid.UUID, req UpdateKitten) (Kitten, error) {
	b := h.client.Kitten.UpdateOneID(id)

	// Add all fields.
	if req.Name.IsSet() {
		b.SetName(req.Name.Get())
	}
	if req.MotherID.IsSet() {
		b.SetMotherID(req.MotherID.Get())
	}
	// Set relations.

	if req.MotherID.IsSet() {
		b.SetMotherID(req.MotherID.Get())
	}

	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		return Kitten{}, err
	}

	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Kitten.Query().Where(kitten.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return Kitten{}, err
	}
	response := NewKitten(e)
	return response, nil
}

// DeleteKitten handles DELETE /kitten/{id} requests.
func (h *Service) DeleteKitten(ctx context.Context, id uuid.UUID) error {
	err := h.client.Kitten.DeleteOneID(id).Exec(ctx)
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
func (h *Service) ListKitten(ctx context.Context, params KittenListParams) (KittenList, error) {
	q := h.client.Kitten.Query()
	if err := h.KittenListQueryModifier(ctx, q); err != nil {
		return KittenList{}, err
	}
	if params.FilterMotherID.IsSet() {
		q.Where(kitten.MotherIDEQ(params.FilterMotherID.Get()))
	}

	if params.With != nil {
		if params.With.Mother != nil {
			q = q.WithMother(applyWithCat(params.With.Mother))
		}
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
			return KittenList{}, err
		case gen.IsNotSingular(err):
			return KittenList{}, err
		default:
			// Let the server handle the error.
			return KittenList{}, err
		}
	}
	r := NewKittenList(es)
	response := KittenList{}
	response.Items = r

	return response, nil
}
func (h *Service) ReadKitten(ctx context.Context, params KittenReadParams) (Kitten, error) {

	q := h.client.Kitten.Query().Where(kitten.IDEQ(params.ID))
	if err := h.KittenReadQueryModifier(ctx, q); err != nil {
		return Kitten{}, err
	}

	if params.With != nil {
		if params.With.Mother != nil {
			q = q.WithMother(applyWithCat(params.With.Mother))
		}
	}

	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case gen.IsNotFound(err):
			return Kitten{}, err
		case gen.IsNotSingular(err):
			return Kitten{}, err
		default:
			// Let the server handle the error.
			return Kitten{}, err
		}
	}

	return NewKitten(e), nil
}

var Module = fx.Module(
	"service",
	fx.Provide(
		New,
	),
)
