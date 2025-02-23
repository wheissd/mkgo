package ent

import (
	"reflect"

	"ariga.io/atlas/sql/postgres"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	entfield "entgo.io/ent/schema/field"
	"github.com/samber/lo"
	"github.com/wheissd/mkgo/annotations"
	"github.com/wheissd/mkgo/internal/cases"
	"github.com/wheissd/mkgo/internal/entity"
	"github.com/wheissd/mkgo/lib"
)

func fieldType(e entfield.Type) entity.FieldType {
	switch e {
	case entfield.TypeBool:
		return entity.TypeBool
	case entfield.TypeTime:
		return entity.TypeTime
	case entfield.TypeJSON:
		return entity.TypeJSON
	case entfield.TypeUUID:
		return entity.TypeUUID
	case entfield.TypeBytes:
		return entity.TypeBytes
	case entfield.TypeEnum:
		return entity.TypeEnum
	case entfield.TypeString:
		return entity.TypeString
	case entfield.TypeOther:
		return entity.TypeOther
	case entfield.TypeInt8:
		return entity.TypeInt8
	case entfield.TypeInt16:
		return entity.TypeInt16
	case entfield.TypeInt32:
		return entity.TypeInt32
	case entfield.TypeInt:
		return entity.TypeInt
	case entfield.TypeInt64:
		return entity.TypeInt64
	case entfield.TypeUint8:
		return entity.TypeUint8
	case entfield.TypeUint16:
		return entity.TypeUint16
	case entfield.TypeUint32:
		return entity.TypeUint32
	case entfield.TypeUint:
		return entity.TypeUint
	case entfield.TypeUint64:
		return entity.TypeUint64
	case entfield.TypeFloat32:
		return entity.TypeFloat32
	case entfield.TypeFloat64:
		return entity.TypeFloat64
	}
	return entity.TypeInvalid
}

func rFieldType(t reflect.Kind) entity.FieldType {
	switch t {
	case reflect.Bool:
		return entity.TypeBool
	//case reflect.Time:
	//	return entity.TypeTime
	//case reflect.JSON:
	//	return entity.TypeJSON
	//case reflect.UUID:
	//	return entity.TypeUUID
	//case reflect.Bytes:
	//	return entity.TypeBytes
	//case reflect.Enum:
	//	return entity.TypeEnum
	case reflect.String:
		return entity.TypeString
	//case reflect.Other:
	//	return entity.TypeOther
	case reflect.Int8:
		return entity.TypeInt8
	case reflect.Int16:
		return entity.TypeInt16
	case reflect.Int32:
		return entity.TypeInt32
	case reflect.Int:
		return entity.TypeInt
	case reflect.Int64:
		return entity.TypeInt64
	case reflect.Uint8:
		return entity.TypeUint8
	case reflect.Uint16:
		return entity.TypeUint16
	case reflect.Uint32:
		return entity.TypeUint32
	case reflect.Uint:
		return entity.TypeUint
	case reflect.Uint64:
		return entity.TypeUint64
	case reflect.Float32:
		return entity.TypeFloat32
	case reflect.Float64:
		return entity.TypeFloat64
	}
	return entity.TypeInvalid
}

type parseFieldsResult struct {
	fields       []entity.Field
	protoImports []string
}

func (op *parseOp) parseFields(e lib.PreEntity) parseFieldsResult {
	res := parseFieldsResult{
		protoImports: make([]string, 0),
	}
	rModel := reflect.TypeOf(e.Model)
	schemaFields := lo.SliceToMap(e.Schema.Fields(), func(field ent.Field) (string, entity.Field) {
		f := op.parseField(field)
		if i := op.ParseProtoImport(field); i != nil {
			res.protoImports = append(res.protoImports, *i)
		}
		return field.Descriptor().Name, f
	})
	for _, mix := range e.Schema.Mixin() {
		for _, f := range mix.Fields() {
			fld := op.parseField(f)
			if i := op.ParseProtoImport(f); i != nil {
				res.protoImports = append(res.protoImports, *i)
			}
			schemaFields[f.Descriptor().Name] = fld
		}
	}

	// get phantom fields not from schema, but from templates
	for i := 0; i < rModel.NumField(); i++ {
		rField := rModel.Field(i)
		_, ok := schemaFields[cases.Snake(rField.Name)]
		if rField.IsExported() && !ok &&
			rField.Name != "ID" && rField.Name != "Edges" {
			t := rFieldType(rField.Type.Kind())
			format := openapiParamFormatFromEntityType(t)
			schemaFields[rField.Name] = entity.Field{
				Name:    rField.Name,
				Phantom: true,
				Openapi: entity.Openapi{
					Type:      openapiFieldTypeFromEntityType(t),
					Format:    format,
					HasFormat: format != "",
				},
				RType: rField.Type,
				Type: entity.TypeInfo{
					Type: t,
					Name: t.TemplateName(),
				},
			}
		}
	}

	for _, edge := range e.Schema.Edges() {
		if f, ok := schemaFields[edge.Descriptor().Field]; ok {
			f.Edge = true
			schemaFields[edge.Descriptor().Field] = f
		}
	}
	res.protoImports = lo.Uniq(res.protoImports)
	res.fields = lo.MapToSlice(schemaFields, func(key string, v entity.Field) entity.Field {
		return v
	})
	return res
}

func (op *parseOp) parseField(field ent.Field) entity.Field {
	required := !field.Descriptor().Optional && field.Descriptor().Default == nil
	f := entity.Field{
		Name:     field.Descriptor().Name,
		Primary:  field.Descriptor().Unique,
		Required: required,
		Optional: field.Descriptor().Optional,
		Openapi: entity.Openapi{
			Type:      openapiFieldType(field.Descriptor()),
			HasFormat: openapiHasFormat(field.Descriptor()),
			Format:    openapiParamFormat(field.Descriptor()),
		},
		Immutable: field.Descriptor().Immutable,
		Type: entity.TypeInfo{
			Type:      fieldType(field.Descriptor().Info.Type),
			Name:      parseName(field),
			Import:    op.parseImport(field),
			IsPointer: field.Descriptor().Nillable,
		},
		Enum: lo.Map(field.Descriptor().Enums, func(item struct{ N, V string }, _ int) any {
			return item.V
		}),
		AutoUpdate: field.Descriptor().UpdateDefault != nil,
	}
	for _, antn := range field.Descriptor().Annotations {
		if antn.Name() == annotations.FieldConfigID {
			if cfg, ok := antn.(*annotations.FieldConfig); ok {
				f.Config = cfg
			}
		}
	}
	return f
}

func parseName(f ent.Field) string {
	name := f.Descriptor().Info.Ident
	// primitive type
	if name == "" {
		name = fieldType(f.Descriptor().Info.Type).TemplateName()
	}
	if f.Descriptor().Info.RType != nil {
		name = f.Descriptor().Info.RType.Name
	}
	return name
}

func (op *parseOp) parseImport(f ent.Field) string {
	imprt := f.Descriptor().Info.PkgName
	if imprt == "" && f.Descriptor().Info.PkgPath != "" {
		imprt = f.Descriptor().Info.PkgPath
		op.sch.AddImport(entity.ImportsEntities, f.Descriptor().Info.PkgPath)
	}
	return imprt
}

func (op *parseOp) ParseProtoImport(f ent.Field) *string {
	if f.Descriptor().Info.Type == entfield.TypeTime {
		return lo.ToPtr("google/protobuf/timestamp.proto")
	}
	return nil
}

func openapiHasFormat(d *entfield.Descriptor) bool {
	hasFormatTypes := []entfield.Type{
		entfield.TypeTime,
		entfield.TypeFloat64,
		entfield.TypeInt64,
		entfield.TypeInt32,
		entfield.TypeInt8,
		entfield.TypeUUID,
	}
	return lo.Contains(hasFormatTypes, d.Info.Type)
}

func openapiParamFormat(d *entfield.Descriptor) string {
	var res string
	switch d.Info.Type {
	case entfield.TypeTime:
		if d.SchemaType[dialect.Postgres] == postgres.TypeDate {
			res = "date"
			break
		}
		res = "date-time"
	case entfield.TypeFloat64:
		res = "double"
	case entfield.TypeUUID:
		res = "uuid"
	case entfield.TypeInt64:
		res = "int64"
	case entfield.TypeInt32:
		res = "int32"
	case entfield.TypeInt16:
		res = "int16"
	case entfield.TypeInt8:
		res = "int8"
	}
	return res
}

func openapiFieldType(d *entfield.Descriptor) string {
	var res string
	switch d.Info.Type {
	case entfield.TypeBool:
		res = "boolean"
	case entfield.TypeString, entfield.TypeUUID, entfield.TypeEnum:
		res = "string"
	case entfield.TypeInt, entfield.TypeInt64, entfield.TypeInt32, entfield.TypeInt16, entfield.TypeInt8:
		res = "integer"
	case entfield.TypeTime:
		res = "string"
	case entfield.TypeFloat64, entfield.TypeFloat32:
		res = "number"
	}
	return res
}

func openapiFieldTypeFromEntityType(t entity.FieldType) string {
	var res string
	switch t {
	case entity.TypeBool:
		res = "boolean"
	case entity.TypeString, entity.TypeUUID, entity.TypeEnum:
		res = "string"
	case entity.TypeInt, entity.TypeInt64, entity.TypeInt32, entity.TypeInt16, entity.TypeInt8:
		res = "integer"
	case entity.TypeTime:
		res = "string"
	case entity.TypeFloat64, entity.TypeFloat32:
		res = "number"
	}
	return res
}

func openapiParamFormatFromEntityType(t entity.FieldType) string {
	var res string
	switch t {
	case entity.TypeTime:
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
