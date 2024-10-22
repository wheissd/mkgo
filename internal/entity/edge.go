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
	RefName    string
	FieldName  string
	EntityName string
	Order      int
	Entity     *Entity
	Type       EdgeType
	Fields     []Field
	Inverse    bool
	ToOne      bool
	WithRead   bool
	WithCreate bool
	WithUpdate bool
	WithDelete bool
}

func NeedEdgeRead(sch *Schema, cfg *annotations.EdgeConfig) bool {
	enabled := sch.Cfg.EnableEdgeReadByDefault
	if cfg != nil && cfg.GetReadEnabled(sch.Cfg.Mode) != nil {
		enabled = *(cfg.GetReadEnabled(sch.Cfg.Mode))
	}
	return enabled
}

func NeedEdgeCreate(sch *Schema, cfg *annotations.EdgeConfig) bool {
	enabled := sch.Cfg.EnableEdgeCreateByDefault
	if cfg != nil && cfg.GetCreateEnabled(sch.Cfg.Mode) != nil {
		enabled = *(cfg.GetCreateEnabled(sch.Cfg.Mode))
	}
	return enabled
}

func NeedEdgeUpdate(sch *Schema, cfg *annotations.EdgeConfig) bool {
	enabled := sch.Cfg.EnableEdgeUpdateByDefault
	if cfg != nil && cfg.GetUpdateEnabled(sch.Cfg.Mode) != nil {
		enabled = *(cfg.GetUpdateEnabled(sch.Cfg.Mode))
	}
	return enabled
}

func NeedEdgeDelete(sch *Schema, cfg *annotations.EdgeConfig) bool {
	enabled := sch.Cfg.EnableEdgeDeleteByDefault
	if cfg != nil && cfg.GetDeleteEnabled(sch.Cfg.Mode) != nil {
		enabled = *(cfg.GetDeleteEnabled(sch.Cfg.Mode))
	}
	return enabled
}

func (e *Edge) ToMany() bool {
	if e.Type == EdgeO2M || e.Type == EdgeM2M {
		return true
	}
	return false
}
