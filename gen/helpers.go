package gen

import (
	"errors"
	"strings"

	"github.com/wheissd/mkgo/internal/cases"
	"github.com/wheissd/mkgo/internal/entity"
)

func needReadOneOp(sch *entity.Schema, e *entity.Entity) bool {
	enabled := sch.Cfg.EnableDefaultReadOne
	if e.Config != nil && e.Config.GetReadOneOpEnabled(sch.Cfg.Mode) != nil {
		enabled = *(e.Config.GetReadOneOpEnabled(sch.Cfg.Mode))
	}
	return enabled
}

func needEntity(sch *entity.Schema, e *entity.Entity) bool {
	enabled := needReadOneOp(sch, e) || needReadManyOp(sch, e) || needCreateOp(sch, e) || needUpdateOp(sch, e)
	if enabled {
		return true
	}
	for _, e := range e.Edges {
		if e.Inverse {
			for _, revEdge := range e.Entity.Edges {
				if revEdge.Name == e.RefName && revEdge.WithRead {
					return true
				}
			}
		}
	}
	return false
}

func needReadManyOp(sch *entity.Schema, e *entity.Entity) bool {
	enabled := sch.Cfg.EnableDefaultReadMany
	if e.Config != nil && e.Config.GetReadManyOpEnabled(sch.Cfg.Mode) != nil {
		enabled = *(e.Config.GetReadManyOpEnabled(sch.Cfg.Mode))
	}
	return enabled
}

func needCreateOp(sch *entity.Schema, e *entity.Entity) bool {
	enabled := sch.Cfg.EnableDefaultCreate
	if e.Config != nil && e.Config.GetCreateOpEnabled(sch.Cfg.Mode) != nil {
		enabled = *(e.Config.GetCreateOpEnabled(sch.Cfg.Mode))
	}
	return enabled
}

func needCreateEntity(sch *entity.Schema, e *entity.Entity) bool {
	enabled := sch.Cfg.EnableDefaultCreate
	if e.Config != nil && e.Config.GetCreateOpEnabled(sch.Cfg.Mode) != nil {
		enabled = *(e.Config.GetCreateOpEnabled(sch.Cfg.Mode))
	}
	for _, e := range e.Edges {
		if e.Inverse {
			for _, revEdge := range e.Entity.Edges {
				if revEdge.Name == e.RefName && revEdge.WithCreate {
					return true
				}
			}
		}
	}
	return enabled
}

func needUpdateOp(sch *entity.Schema, e *entity.Entity) bool {
	enabled := sch.Cfg.EnableDefaultUpdate
	if e.Config != nil && e.Config.GetUpdateOpEnabled(sch.Cfg.Mode) != nil {
		enabled = *(e.Config.GetUpdateOpEnabled(sch.Cfg.Mode))
	}
	return enabled
}

func needUpdateEntity(sch *entity.Schema, e *entity.Entity) bool {
	enabled := sch.Cfg.EnableDefaultUpdate
	if e.Config != nil && e.Config.GetUpdateOpEnabled(sch.Cfg.Mode) != nil {
		enabled = *(e.Config.GetUpdateOpEnabled(sch.Cfg.Mode))
	}
	for _, e := range e.Edges {
		if e.Inverse {
			for _, revEdge := range e.Entity.Edges {
				if revEdge.Name == e.RefName && revEdge.WithUpdate {
					return true
				}
			}
		}
	}
	return enabled
}

func needDeleteOp(sch *entity.Schema, e *entity.Entity) bool {
	enabled := sch.Cfg.EnableDefaultDelete
	if e.Config != nil && e.Config.GetDeleteOpEnabled(sch.Cfg.Mode) != nil {
		enabled = *(e.Config.GetDeleteOpEnabled(sch.Cfg.Mode))
	}
	return enabled
}

var hasOp bool

func hasOps(sch *entity.Schema) bool {
	if hasOp {
		return hasOp
	}
	for _, e := range sch.Entities {
		if needCreateOp(sch, e) {
			hasOp = true
			return hasOp
		}
		if needUpdateOp(sch, e) {
			hasOp = true
			return hasOp
		}
		if needDeleteOp(sch, e) {
			hasOp = true
			return hasOp
		}
		if needReadManyOp(sch, e) {
			hasOp = true
			return hasOp
		}
		if needReadOneOp(sch, e) {
			hasOp = true
			return hasOp
		}
	}
	return hasOp
}

func needFilter(sch *entity.Schema, f *entity.Field) bool {
	enabled := sch.Cfg.EnableDefaultReadMany
	if f.Config != nil && f.Config.GetEnableFilter(sch.Cfg.Mode) != nil {
		enabled = *(f.Config.GetEnableFilter(sch.Cfg.Mode))
	}
	return enabled
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func concat(strs ...string) string {
	b := strings.Builder{}
	for i := range strs {
		b.WriteString(strs[i])
	}
	return b.String()
}

func techField(f *entity.Field) bool {
	switch f.Name {
	case "create_time", "update_time", "deleted_time":
		return true
	}
	return false
}

func protoToServiceField(f *entity.Field) string {
	if f.Type.Type == entity.TypeTime {
		return "req." + cases.Pascal(f.Name) + ".AsTime()"
	} else {
		return "req." + cases.Pascal(f.Name)
	}
}

func protoToServiceFieldUpdate(f *entity.Field) string {
	if f.Type.Type == entity.TypeTime {
		return "req." + cases.Pascal(f.Name) + ".AsTime()"
	} else {
		return "*req." + cases.Pascal(f.Name)
	}
}

func protoToServiceFieldFilter(f *entity.Field) string {
	if f.Type.Type == entity.TypeTime {
		return "params.Filter" + cases.Pascal(f.Name) + ".AsTime()"
	} else {
		return "*params.Filter" + cases.Pascal(f.Name)
	}
}

func serviceToProtoField(f entity.Field) string {
	get := ""
	if !f.Required {
		get = ".Get()"
	}
	switch f.Type.Type {
	case entity.TypeTime:
		return "timestamppb.New(e." + cases.Pascal(f.Name) + get + ")"
	case entity.TypeUUID:
		return "e." + cases.Pascal(f.Name) + get + ".String()"
	default:
		return "e." + cases.Pascal(f.Name) + get
	}
}
