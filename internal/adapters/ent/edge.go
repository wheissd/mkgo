package ent

import (
	"entgo.io/ent"
	"github.com/wheissd/mkgo/internal/entity"
)

func getEdgeType(edge ent.Edge) entity.EdgeType {
	d := edge.Descriptor()
	if d.Inverse {
		if d.Unique {
			return entity.EdgeM2O
		} else {
			return entity.EdgeO2M
		}
	}
	if d.Unique {
		return entity.EdgeO2O
	}
	return entity.EdgeO2M
}
