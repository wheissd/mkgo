package entity

import (
	"github.com/wheissd/mkgo/internal/config"
)

const ImportsEntities = "entities"

type Schema struct {
	Pkg          string
	EntPath      string
	RootPkg      string
	Entities     []*Entity
	Edges        []*Edge
	Imports      map[string][]string
	Cfg          config.GenConfigItem
	WD           string
	ProtoImports []string
}

func (sch *Schema) AddImport(t, imprt string) {
	if sch.Imports == nil {
		sch.Imports = make(map[string][]string)
	}
	if sch.Imports[t] == nil {
		sch.Imports[t] = make([]string, 0)
	}

	sch.Imports[t] = append(sch.Imports[t], imprt)
}
