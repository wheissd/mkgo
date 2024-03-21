// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateBreed implements createBreed operation.
//
// Create a new Breed  and persists it to storage.
//
// POST /breed
func (UnimplementedHandler) CreateBreed(ctx context.Context, req *CreateBreed) (r CreateBreedRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateCat implements createCat operation.
//
// Create a new Cat  and persists it to storage.
//
// POST /cat
func (UnimplementedHandler) CreateCat(ctx context.Context, req *CreateCat) (r CreateCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateFatherCat implements createFatherCat operation.
//
// Create a new FatherCat  and persists it to storage.
//
// POST /fathercat
func (UnimplementedHandler) CreateFatherCat(ctx context.Context, req *CreateFatherCat) (r CreateFatherCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateKitten implements createKitten operation.
//
// Create a new Kitten  and persists it to storage.
//
// POST /kitten
func (UnimplementedHandler) CreateKitten(ctx context.Context, req *CreateKitten) (r CreateKittenRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteBreed implements deleteBreed operation.
//
// Delete Breed.
//
// DELETE /breed/{id}
func (UnimplementedHandler) DeleteBreed(ctx context.Context, params DeleteBreedParams) (r DeleteBreedRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteCat implements deleteCat operation.
//
// Delete Cat.
//
// DELETE /cat/{id}
func (UnimplementedHandler) DeleteCat(ctx context.Context, params DeleteCatParams) (r DeleteCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteFatherCat implements deleteFatherCat operation.
//
// Delete FatherCat.
//
// DELETE /fathercat/{id}
func (UnimplementedHandler) DeleteFatherCat(ctx context.Context, params DeleteFatherCatParams) (r DeleteFatherCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteKitten implements deleteKitten operation.
//
// Delete Kitten.
//
// DELETE /kitten/{id}
func (UnimplementedHandler) DeleteKitten(ctx context.Context, params DeleteKittenParams) (r DeleteKittenRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListBreed implements listBreed operation.
//
// List for Breeds.
//
// GET /breed
func (UnimplementedHandler) ListBreed(ctx context.Context, params ListBreedParams) (r ListBreedRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListCat implements listCat operation.
//
// List for Cats.
//
// GET /cat
func (UnimplementedHandler) ListCat(ctx context.Context, params ListCatParams) (r ListCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListFatherCat implements listFatherCat operation.
//
// List for FatherCats.
//
// GET /fathercat
func (UnimplementedHandler) ListFatherCat(ctx context.Context, params ListFatherCatParams) (r ListFatherCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListKitten implements listKitten operation.
//
// List for Kittens.
//
// GET /kitten
func (UnimplementedHandler) ListKitten(ctx context.Context, params ListKittenParams) (r ListKittenRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadBreed implements readBreed operation.
//
// Finds the Breed with the requested ID and returns it.
//
// GET /breed/{id}
func (UnimplementedHandler) ReadBreed(ctx context.Context, params ReadBreedParams) (r ReadBreedRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadCat implements readCat operation.
//
// Finds the Cat with the requested ID and returns it.
//
// GET /cat/{id}
func (UnimplementedHandler) ReadCat(ctx context.Context, params ReadCatParams) (r ReadCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadFatherCat implements readFatherCat operation.
//
// Finds the FatherCat with the requested ID and returns it.
//
// GET /fathercat/{id}
func (UnimplementedHandler) ReadFatherCat(ctx context.Context, params ReadFatherCatParams) (r ReadFatherCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadKitten implements readKitten operation.
//
// Finds the Kitten with the requested ID and returns it.
//
// GET /kitten/{id}
func (UnimplementedHandler) ReadKitten(ctx context.Context, params ReadKittenParams) (r ReadKittenRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateBreed implements updateBreed operation.
//
// Update Breed  and persists it to storage.
//
// PUT /breed/{id}
func (UnimplementedHandler) UpdateBreed(ctx context.Context, req *UpdateBreed, params UpdateBreedParams) (r UpdateBreedRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateCat implements updateCat operation.
//
// Update Cat  and persists it to storage.
//
// PUT /cat/{id}
func (UnimplementedHandler) UpdateCat(ctx context.Context, req *UpdateCat, params UpdateCatParams) (r UpdateCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateFatherCat implements updateFatherCat operation.
//
// Update FatherCat  and persists it to storage.
//
// PUT /fathercat/{id}
func (UnimplementedHandler) UpdateFatherCat(ctx context.Context, req *UpdateFatherCat, params UpdateFatherCatParams) (r UpdateFatherCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateKitten implements updateKitten operation.
//
// Update Kitten  and persists it to storage.
//
// PUT /kitten/{id}
func (UnimplementedHandler) UpdateKitten(ctx context.Context, req *UpdateKitten, params UpdateKittenParams) (r UpdateKittenRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}