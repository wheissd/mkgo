package grpc

import (
    "context"

    "github.com/google/uuid"
    "github.com/samber/lo"
    "github.com/wheissd/mkgo/internal/examples/catsadventure/api/service"
)


type BreedServiceImpl struct {
    service *service.Service
    UnimplementedBreedServiceServer
}

func NewBreedServiceImpl(srvc *service.Service) *BreedServiceImpl {
    return &BreedServiceImpl{
        service: srvc,
    }
}

func parseWithBreed(e *BreedEdgesRequest) *service.BreedQueryEdges {
	if e == nil {
		return nil
	}
	return &service.BreedQueryEdges{
				Cats: parseWithCat(e.WithCats),
	}
}

    
	func (s *BreedServiceImpl) CreateBreed(ctx context.Context, req *CreateBreedRequest) (*Breed, error) {
		breed, err := s.service.CreateBreed(
			ctx,
			service.CreateBreed {
					Name: req.Name,
			},
			)
			if err != nil {
				return nil, err
			}
			res := NewBreed(breed)
			return res, nil
	}
	
	func (s *BreedServiceImpl) UpdateBreed(ctx context.Context, req *UpdateBreedRequest) (*Breed, error) {
		id, err := uuid.Parse(req.ID)
		if err != nil {
			return nil, err
		}
		update := service.UpdateBreed{}
			if req.Name != nil {
				update.Name.Set(*req.Name)
			}
		serviceRes, err := s.service.UpdateBreed(ctx, id, update)
		if err != nil {
			return nil, err
		}
		res := NewBreed(serviceRes)
		return res, nil
	}
	
	// DeleteBreed handles DELETE /breed/{id} requests.
	func (s *BreedServiceImpl) DeleteBreed(ctx context.Context, params *DeleteBreedRequest) (*DeleteBreedResponse, error) {
		id, err := uuid.Parse(params.ID)
		if err != nil {
			return nil, err
		}
		err = s.service.DeleteBreed(ctx, id)
		if err != nil {
			return nil, err
		}
		return &DeleteBreedResponse{Status: "ok"}, nil
	}
	
	func (s *BreedServiceImpl) ReadManyBreed(ctx context.Context, params *ReadManyBreedRequest) (*ReadManyBreedResponse, error) {
		sParams := service.BreedListParams{}
		if params.Page != nil {
			sParams.Page.Set(int(*params.Page))
		}
		if params.ItemsPerPage != nil {
			sParams.ItemsPerPage.Set(int(*params.ItemsPerPage))
		}
		
		var (
			err error
		)
		sParams.With = parseWithBreed(params.With)
		list, err := s.service.ListBreed(ctx, sParams)
		if err != nil {
			return nil, err
		}
		response := ReadManyBreedResponse{
			Items: NewBreedList(list.Items),
		}
		return &response, nil
	}
	
        func (s *BreedServiceImpl) ReadBreed(ctx context.Context, r *ReadBreedRequest) (*Breed, error) {
            id, err := uuid.Parse(r.ID)
            if err != nil {
                return nil, err
            }
            e, err := s.service.ReadBreed(ctx, service.BreedReadParams{ID: id})
            if err != nil {
                return nil, err
            }
            response := NewBreed(e)
            return response, nil
        }
	

type CatServiceImpl struct {
    service *service.Service
    UnimplementedCatServiceServer
}

func NewCatServiceImpl(srvc *service.Service) *CatServiceImpl {
    return &CatServiceImpl{
        service: srvc,
    }
}

func parseWithCat(e *CatEdgesRequest) *service.CatQueryEdges {
	if e == nil {
		return nil
	}
	return &service.CatQueryEdges{
				Kittens: parseWithKitten(e.WithKittens),
				Breed: parseWithBreed(e.WithBreed),
	}
}

    
	func (s *CatServiceImpl) CreateCat(ctx context.Context, req *CreateCatRequest) (*Cat, error) {
				breedID, err := uuid.Parse(req.BreedID)
				if err != nil {
					return nil, err
				}
		cat, err := s.service.CreateCat(
			ctx,
			service.CreateCat {
					Name: req.Name,
					BreedID: breedID,
					Speed: req.Speed,
					DateFrom: req.DateFrom.AsTime(),
					OtherType: catOtherTypeProtoToServiceConverterreq.OtherType),
					Type: catTypeProtoToServiceConverterreq.Type),
			},
			)
			if err != nil {
				return nil, err
			}
			res := NewCat(cat)
			return res, nil
	}
	
	func (s *CatServiceImpl) UpdateCat(ctx context.Context, req *UpdateCatRequest) (*Cat, error) {
		id, err := uuid.Parse(req.ID)
		if err != nil {
			return nil, err
		}
		update := service.UpdateCat{}
			if req.Name != nil {
				update.Name.Set(*req.Name)
			}
			if req.BreedID != nil {
					breedID, err := uuid.Parse(*req.BreedID)
					if err != nil {
						return nil, err
					}
				update.BreedID.Set(breedID)
			}
			if req.Speed != nil {
				update.Speed.Set(*req.Speed)
			}
			if req.DateFrom != nil {
				update.DateFrom.Set(req.DateFrom.AsTime())
			}
			if req.OtherType != nil {
				update.OtherType.Set(*req.OtherType)
			}
			if req.Type != nil {
				update.Type.Set(*req.Type)
			}
		serviceRes, err := s.service.UpdateCat(ctx, id, update)
		if err != nil {
			return nil, err
		}
		res := NewCat(serviceRes)
		return res, nil
	}
	
	// DeleteCat handles DELETE /cat/{id} requests.
	func (s *CatServiceImpl) DeleteCat(ctx context.Context, params *DeleteCatRequest) (*DeleteCatResponse, error) {
		id, err := uuid.Parse(params.ID)
		if err != nil {
			return nil, err
		}
		err = s.service.DeleteCat(ctx, id)
		if err != nil {
			return nil, err
		}
		return &DeleteCatResponse{Status: "ok"}, nil
	}
	
	func (s *CatServiceImpl) ReadManyCat(ctx context.Context, params *ReadManyCatRequest) (*ReadManyCatResponse, error) {
		sParams := service.CatListParams{}
		if params.Page != nil {
			sParams.Page.Set(int(*params.Page))
		}
		if params.ItemsPerPage != nil {
			sParams.ItemsPerPage.Set(int(*params.ItemsPerPage))
		}
		
		var (
			err error
		)
		sParams.With = parseWithCat(params.With)
		list, err := s.service.ListCat(ctx, sParams)
		if err != nil {
			return nil, err
		}
		response := ReadManyCatResponse{
			Items: NewCatList(list.Items),
		}
		return &response, nil
	}
	
        func (s *CatServiceImpl) ReadCat(ctx context.Context, r *ReadCatRequest) (*Cat, error) {
            id, err := uuid.Parse(r.ID)
            if err != nil {
                return nil, err
            }
            e, err := s.service.ReadCat(ctx, service.CatReadParams{ID: id})
            if err != nil {
                return nil, err
            }
            response := NewCat(e)
            return response, nil
        }
	

type FatherCatServiceImpl struct {
    service *service.Service
    UnimplementedFatherCatServiceServer
}

func NewFatherCatServiceImpl(srvc *service.Service) *FatherCatServiceImpl {
    return &FatherCatServiceImpl{
        service: srvc,
    }
}

func parseWithFatherCat(e *FatherCatEdgesRequest) *service.FatherCatQueryEdges {
	if e == nil {
		return nil
	}
	return &service.FatherCatQueryEdges{
	}
}

    
	func (s *FatherCatServiceImpl) CreateFatherCat(ctx context.Context, req *CreateFatherCatRequest) (*FatherCat, error) {
		fathercat, err := s.service.CreateFatherCat(
			ctx,
			service.CreateFatherCat {
					Name: req.Name,
			},
			)
			if err != nil {
				return nil, err
			}
			res := NewFatherCat(fathercat)
			return res, nil
	}
	
	func (s *FatherCatServiceImpl) UpdateFatherCat(ctx context.Context, req *UpdateFatherCatRequest) (*FatherCat, error) {
		id, err := uuid.Parse(req.ID)
		if err != nil {
			return nil, err
		}
		update := service.UpdateFatherCat{}
			if req.Name != nil {
				update.Name.Set(*req.Name)
			}
		serviceRes, err := s.service.UpdateFatherCat(ctx, id, update)
		if err != nil {
			return nil, err
		}
		res := NewFatherCat(serviceRes)
		return res, nil
	}
	
	// DeleteFatherCat handles DELETE /fathercat/{id} requests.
	func (s *FatherCatServiceImpl) DeleteFatherCat(ctx context.Context, params *DeleteFatherCatRequest) (*DeleteFatherCatResponse, error) {
		id, err := uuid.Parse(params.ID)
		if err != nil {
			return nil, err
		}
		err = s.service.DeleteFatherCat(ctx, id)
		if err != nil {
			return nil, err
		}
		return &DeleteFatherCatResponse{Status: "ok"}, nil
	}
	
	func (s *FatherCatServiceImpl) ReadManyFatherCat(ctx context.Context, params *ReadManyFatherCatRequest) (*ReadManyFatherCatResponse, error) {
		sParams := service.FatherCatListParams{}
		if params.Page != nil {
			sParams.Page.Set(int(*params.Page))
		}
		if params.ItemsPerPage != nil {
			sParams.ItemsPerPage.Set(int(*params.ItemsPerPage))
		}
		
		list, err := s.service.ListFatherCat(ctx, sParams)
		if err != nil {
			return nil, err
		}
		response := ReadManyFatherCatResponse{
			Items: NewFatherCatList(list.Items),
		}
		return &response, nil
	}
	
        func (s *FatherCatServiceImpl) ReadFatherCat(ctx context.Context, r *ReadFatherCatRequest) (*FatherCat, error) {
            id, err := uuid.Parse(r.ID)
            if err != nil {
                return nil, err
            }
            e, err := s.service.ReadFatherCat(ctx, service.FatherCatReadParams{ID: id})
            if err != nil {
                return nil, err
            }
            response := NewFatherCat(e)
            return response, nil
        }
	

type KittenServiceImpl struct {
    service *service.Service
    UnimplementedKittenServiceServer
}

func NewKittenServiceImpl(srvc *service.Service) *KittenServiceImpl {
    return &KittenServiceImpl{
        service: srvc,
    }
}

func parseWithKitten(e *KittenEdgesRequest) *service.KittenQueryEdges {
	if e == nil {
		return nil
	}
	return &service.KittenQueryEdges{
				Mother: parseWithCat(e.WithMother),
	}
}

    
	func (s *KittenServiceImpl) CreateKitten(ctx context.Context, req *CreateKittenRequest) (*Kitten, error) {
				motherID, err := uuid.Parse(req.MotherID)
				if err != nil {
					return nil, err
				}
		kitten, err := s.service.CreateKitten(
			ctx,
			service.CreateKitten {
					Name: req.Name,
					MotherID: motherID,
			},
			)
			if err != nil {
				return nil, err
			}
			res := NewKitten(kitten)
			return res, nil
	}
	
	func (s *KittenServiceImpl) UpdateKitten(ctx context.Context, req *UpdateKittenRequest) (*Kitten, error) {
		id, err := uuid.Parse(req.ID)
		if err != nil {
			return nil, err
		}
		update := service.UpdateKitten{}
			if req.Name != nil {
				update.Name.Set(*req.Name)
			}
			if req.MotherID != nil {
					motherID, err := uuid.Parse(*req.MotherID)
					if err != nil {
						return nil, err
					}
				update.MotherID.Set(motherID)
			}
		serviceRes, err := s.service.UpdateKitten(ctx, id, update)
		if err != nil {
			return nil, err
		}
		res := NewKitten(serviceRes)
		return res, nil
	}
	
	// DeleteKitten handles DELETE /kitten/{id} requests.
	func (s *KittenServiceImpl) DeleteKitten(ctx context.Context, params *DeleteKittenRequest) (*DeleteKittenResponse, error) {
		id, err := uuid.Parse(params.ID)
		if err != nil {
			return nil, err
		}
		err = s.service.DeleteKitten(ctx, id)
		if err != nil {
			return nil, err
		}
		return &DeleteKittenResponse{Status: "ok"}, nil
	}
	
	func (s *KittenServiceImpl) ReadManyKitten(ctx context.Context, params *ReadManyKittenRequest) (*ReadManyKittenResponse, error) {
		sParams := service.KittenListParams{}
		if params.Page != nil {
			sParams.Page.Set(int(*params.Page))
		}
		if params.ItemsPerPage != nil {
			sParams.ItemsPerPage.Set(int(*params.ItemsPerPage))
		}
		
		var (
			err error
		)
		sParams.With = parseWithKitten(params.With)
		list, err := s.service.ListKitten(ctx, sParams)
		if err != nil {
			return nil, err
		}
		response := ReadManyKittenResponse{
			Items: NewKittenList(list.Items),
		}
		return &response, nil
	}
	
        func (s *KittenServiceImpl) ReadKitten(ctx context.Context, r *ReadKittenRequest) (*Kitten, error) {
            id, err := uuid.Parse(r.ID)
            if err != nil {
                return nil, err
            }
            e, err := s.service.ReadKitten(ctx, service.KittenReadParams{ID: id})
            if err != nil {
                return nil, err
            }
            response := NewKitten(e)
            return response, nil
        }
	

