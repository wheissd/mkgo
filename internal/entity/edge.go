package entity

import (
	"github.com/wheissd/mkgo/annotations"
)

type EdgeType int

const (
	EdgeO2O = iota
	EdgeO2M
	EdgeM2O
	EdgeM2M
)

type Edge struct {
	Name       string
	EntityName string
	Order      int
	Entity     *Entity
	Type       EdgeType
	Fields     []Field
	Inverse    bool
	ToOne      bool
	WithRead   bool
}

func NeedEdgeRead(sch *Schema, cfg *annotations.EdgeConfig) bool {
	enabled := sch.Cfg.EnableEdgeReadByDefault
	if cfg != nil && cfg.GetReadEnabled(sch.Cfg.Mode) != nil {
		enabled = *(cfg.GetReadEnabled(sch.Cfg.Mode))
	}
	return enabled
}
