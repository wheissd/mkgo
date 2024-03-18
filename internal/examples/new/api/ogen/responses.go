package ogen

import (
	"github.com/wheissd/mkgo/internal/examples/new/api/service"
)

func NewDefaultModel(e service.DefaultModel) *DefaultModel {
	var ret DefaultModel
	ret.ID = e.ID
	ret.ID = e.ID
	ret.Name = e.Name

	return &ret
}

func NewDefaultModelList(es []service.DefaultModel) []DefaultModel {
	if len(es) == 0 {
		return nil
	}
	r := make([]DefaultModel, len(es))
	for i, e := range es {
		r[i] = *NewDefaultModel(e)
	}
	return r
}
