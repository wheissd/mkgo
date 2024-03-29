package ent

import (
	"reflect"
	"sort"
	"strings"

	"github.com/samber/lo"
	"github.com/wheissd/mkgo/annotations"
	"github.com/wheissd/mkgo/internal/entity"
	"github.com/wheissd/mkgo/lib"
)

type parseOp struct {
	sch *entity.Schema
}

func Parse(entities []lib.PreEntity, schema *entity.Schema, mode string) {
	entMap := lo.SliceToMap(entities, func(item lib.PreEntity) (string, lib.PreEntity) {
		rType := reflect.TypeOf(item.Schema)
		return rType.Name(), item
	})
	callbacks := make([]func(), 0)
	op := parseOp{sch: schema}
	res := make(map[string]*entity.Entity, len(entities))
	for _, preEntity := range entities {
		reflectType := reflect.TypeOf(preEntity.Schema)
		name := reflectType.Name()
		var (
			e  *entity.Entity
			ok bool
		)
		if e, ok = res[name]; !ok {
			e = op.parseBase(preEntity)
		}

		for _, edge := range preEntity.Schema.Edges() {
			//for _, antn := range edge.Descriptor().Annotations {
			//if antn.Name() == annotations.EdgeConfigID {
			var (
				edgeEnt  *entity.Entity
				isParsed bool
			)
			edgeType := edge.Descriptor().Type
			if edgeEnt, isParsed = res[edgeType]; !isParsed {
				res[edgeType] = op.parseBase(entMap[edgeType])
				edgeEnt = res[edgeType]
			}

			et := getEdgeType(edge)
			toOne := et == entity.EdgeM2O || et == entity.EdgeO2O
			edg := &entity.Edge{
				Name:       edge.Descriptor().Name,
				FieldName:  edge.Descriptor().Field,
				EntityName: edge.Descriptor().Type,
				Type:       et,
				Fields:     edgeEnt.Fields,
				Inverse:    edge.Descriptor().Inverse,
				ToOne:      toOne,
			}

			var cfgAntn *annotations.EdgeConfig
			for _, antn := range edge.Descriptor().Annotations {
				if antn.Name() == annotations.EdgeConfigID {
					if cfg, ok := antn.(*annotations.EdgeConfig); ok {
						cfgAntn = cfg
					}
				}
			}
			if entity.NeedEdgeRead(schema, cfgAntn) {
				edg.WithRead = true
				e.HasReadEdges = true
			}

			e.Edges = append(e.Edges, edg)
			callbacks = append(callbacks, func() {
				edg.Entity = res[edg.EntityName]
			})
		}
		res[e.Name] = e
	}
	for i := range callbacks {
		callbacks[i]()
	}

	resSlice := lo.MapToSlice(res, func(key string, v *entity.Entity) *entity.Entity {
		return v
	})
	sort.Slice(resSlice, func(i, j int) bool {
		return resSlice[i].Name < resSlice[j].Name
	})

	schema.Entities = resSlice
	return
}

func (op *parseOp) parseBase(preEntity lib.PreEntity) *entity.Entity {
	reflectType := reflect.TypeOf(preEntity.Schema)
	name := reflectType.Name()
	fieldsRes := op.parseFields(preEntity)
	op.sch.ProtoImports = lo.Uniq(append(op.sch.ProtoImports, fieldsRes.protoImports...))
	resEntity := &entity.Entity{
		Model:  preEntity.Model,
		Path:   strings.ToLower(reflectType.Name()),
		Name:   name,
		Edges:  make([]*entity.Edge, 0),
		Fields: fieldsRes.fields,
	}
	for _, antn := range preEntity.Schema.Annotations() {
		if antn.Name() == annotations.EntityConfigID {
			if cfg, ok := antn.(*annotations.EntityConfig); ok {
				resEntity.Config = cfg
			}
		}
	}
	return resEntity
}
