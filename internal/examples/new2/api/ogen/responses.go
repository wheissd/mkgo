package ogen

import (
	"github.com/wheissd/mkgo/internal/examples/new2/api/service"
)

func NewHuman(e service.Human) *Human {
	var ret Human
	ret.ID = e.ID
	ret.ID = e.ID
	ret.Name = e.Name

	return &ret
}

func NewHumanList(es []service.Human) []Human {
	if len(es) == 0 {
		return nil
	}
	r := make([]Human, len(es))
	for i, e := range es {
		r[i] = *NewHuman(e)
	}
	return r
}
