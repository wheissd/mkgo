package entity

import (
	"reflect"

	"github.com/wheissd/mkgo/annotations"
)

type Field struct {
	Name       string
	Primary    bool
	Required   bool
	Optional   bool
	Immutable  bool
	AutoUpdate bool
	Edge       bool
	Openapi    Openapi
	Phantom    bool
	RType      reflect.Type
	Type       TypeInfo
	Enum       []any
	Config     *annotations.FieldConfig
	Order      int
}

type TypeInfo struct {
	Type FieldType
	Name string
	// import name for composite e.g. uuid.UUID
	Import       string
	ProtoImport  *string
	ProtoPointer bool
	IsPointer    bool
}

func (f FieldType) OApiTypeName() string {
	var res string
	switch f {
	case TypeBool:
		res = "boolean"
	case TypeString, TypeUUID, TypeEnum:
		res = "string"
	case TypeInt, TypeInt64, TypeInt32, TypeInt16, TypeInt8:
		res = "integer"
	case TypeTime, TypeDate, TypeDateTime:
		res = "string"
	case TypeFloat64, TypeFloat32:
		res = "number"
	}
	return res
}

func (f FieldType) ProtoTypeName() string {
	var res string
	switch f {
	case TypeBool:
		res = "bool"
	case TypeString, TypeUUID, TypeEnum:
		res = "string"
	case TypeInt:
		res = "int64"
	case TypeInt64:
		res = "int64"
	case TypeInt32:
		res = "int32"
	case TypeInt16:
		res = "int16"
	case TypeInt8:
		res = "int8"
	case TypeTime, TypeDate, TypeDateTime:
		res = "google.protobuf.Timestamp"
	case TypeFloat64:
		res = "float64"
	case TypeFloat32:
		res = "float32"
	}
	return res
}

type FieldType uint8

const (
	TypeInvalid FieldType = iota
	TypeBool
	TypeTime
	TypeDate
	TypeDateTime
	TypeJSON
	TypeUUID
	TypeBytes
	TypeEnum
	TypeString
	TypeOther
	TypeInt8
	TypeInt16
	TypeInt32
	TypeInt
	TypeInt64
	TypeUint8
	TypeUint16
	TypeUint32
	TypeUint
	TypeUint64
	TypeFloat32
	TypeFloat64
)

func (t FieldType) TemplateName() string {
	switch t {
	case TypeInvalid:
		return "invalid"
	case TypeBool:
		return "bool"
	case TypeTime:
		return "Time"
	case TypeDate:
		return "date"
	case TypeDateTime:
		return "datetime"
	case TypeJSON:
		return "json"
	case TypeUUID:
		return "uuid"
	case TypeBytes:
		return "bytes"
	case TypeEnum:
		return "string"
	case TypeString:
		return "string"
	case TypeOther:
		return "other"
	case TypeInt8:
		return "int8"
	case TypeInt16:
		return "int16"
	case TypeInt32:
		return "int32"
	case TypeInt:
		return "int"
	case TypeInt64:
		return "int64"
	case TypeUint8:
		return "uint8"
	case TypeUint16:
		return "uint16"
	case TypeUint32:
		return "uint32"
	case TypeUint:
		return "uint"
	case TypeUint64:
		return "uint64"
	case TypeFloat32:
		return "float32"
	case TypeFloat64:
		return "float64"
	}
	return ""
}

func (t FieldType) Name() string {
	switch t {
	case TypeInvalid:
		return "invalid"
	case TypeBool:
		return "bool"
	case TypeTime:
		return "Time"
	case TypeDate:
		return "date"
	case TypeDateTime:
		return "datetime"
	case TypeJSON:
		return "json"
	case TypeUUID:
		return "uuid"
	case TypeBytes:
		return "bytes"
	case TypeEnum:
		return "enum"
	case TypeString:
		return "string"
	case TypeOther:
		return "other"
	case TypeInt8:
		return "int8"
	case TypeInt16:
		return "int16"
	case TypeInt32:
		return "int32"
	case TypeInt:
		return "int"
	case TypeInt64:
		return "int64"
	case TypeUint8:
		return "uint8"
	case TypeUint16:
		return "uint16"
	case TypeUint32:
		return "uint32"
	case TypeUint:
		return "uint"
	case TypeUint64:
		return "uint64"
	case TypeFloat32:
		return "float32"
	case TypeFloat64:
		return "float64"
	}
	return ""
}

func (t FieldType) CanRange() bool {
	switch t {
	case TypeInt, TypeInt8, TypeInt16, TypeInt32, TypeInt64, TypeFloat32,
		TypeFloat64, TypeDate, TypeDateTime, TypeTime:
		return true
	default:
		return false
	}
}
