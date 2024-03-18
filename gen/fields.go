package gen

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/samber/lo"
	"github.com/wheissd/mkgo/internal/cases"
	"github.com/wheissd/mkgo/internal/entity"
)

func hasFormat(f entity.Field) bool {
	if f.Openapi.HasFormat != false {
		return true
	}
	hasFormatTypes := []entity.FieldType{
		entity.TypeTime,
		entity.TypeDate,
		entity.TypeDateTime,
		entity.TypeFloat64,
		entity.TypeInt64,
		entity.TypeInt32,
		entity.TypeInt8,
		entity.TypeUUID,
	}
	return lo.Contains(hasFormatTypes, f.Type.Type)
}

func paramFormat(f entity.Field) string {
	if f.Openapi.Format != "" {
		return f.Openapi.Format
	}
	var res string
	switch f.Type.Type {
	case entity.TypeTime:
		res = "time"
	case entity.TypeDate:
		res = "date"
	case entity.TypeDateTime:
		res = "date-time"
	case entity.TypeFloat64:
		res = "double"
	case entity.TypeUUID:
		res = "uuid"
	case entity.TypeInt64:
		res = "int64"
	case entity.TypeInt32:
		res = "int32"
	case entity.TypeInt16:
		res = "int16"
	case entity.TypeInt8:
		res = "int8"
	}
	return res
}

func rFieldType(t reflect.Type) string {
	var res string
	switch t.Kind() {
	case reflect.String:
		res = "string"
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		res = "integer"
	case reflect.Float64, reflect.Float32:
		res = "number"
	}
	return res
}

func setReqFormat(e entity.Entity, f entity.Field) string {
	if f.Type.Type == entity.TypeEnum {
		return fmt.Sprintf("%s.%s(req.%s)", strings.ToLower(e.Name), cases.Pascal(f.Name), cases.Pascal(f.Name))
	}
	return fmt.Sprintf("req.%s", cases.Pascal(f.Name))
}

func updateReqFormat(e entity.Entity, f entity.Field) string {
	if f.Type.Type == entity.TypeEnum {
		return fmt.Sprintf("%s.%s(req.%s.Get())", strings.ToLower(e.Name), cases.Pascal(f.Name), cases.Pascal(f.Name))
	}
	return fmt.Sprintf("req.%s.Get()", cases.Pascal(f.Name))
}

func isEnum(f entity.Field) bool {
	return f.Type.Type == entity.TypeEnum
}

func isFieldPublic(sch *entity.Schema, f entity.Field) bool {
	isPublic := sch.Cfg.FieldsPublicByDefault
	if f.Config != nil && f.Config.GetPublic(sch.Cfg.Mode) != nil {
		isPublic = *(f.Config.GetPublic(sch.Cfg.Mode))
	}
	return isPublic
}

func isIDField(f entity.Field) bool {
	return f.Name == "id"
}

func fieldType(f entity.Field) string {
	res := ""
	//if f.Type.IsPointer {
	//	res += "*"
	//}
	if f.Type.Import != "" {
		res += f.Type.Import + "."
	}
	res += f.Type.Name
	return res
}

func fieldTypeIs(f entity.Field, t string) bool {
	res := ""
	//if f.Type.IsPointer {
	//	res += "*"
	//}
	if f.Type.Import != "" {
		res += f.Type.Import + "."
	}
	res += f.Type.Name
	return res == t
}

func updateFieldType(f entity.Field) string {
	res := ""
	if f.Type.Import != "" {
		res += f.Type.Import + "."
	}
	res += f.Type.Name
	return res
}
