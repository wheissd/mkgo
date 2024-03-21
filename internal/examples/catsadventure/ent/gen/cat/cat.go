// Code generated by ent, DO NOT EDIT.

package cat

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the cat type in the database.
	Label = "cat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeletedTime holds the string denoting the deleted_time field in the database.
	FieldDeletedTime = "deleted_time"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBreedID holds the string denoting the breed_id field in the database.
	FieldBreedID = "breed_id"
	// FieldSpeed holds the string denoting the speed field in the database.
	FieldSpeed = "speed"
	// FieldDateFrom holds the string denoting the date_from field in the database.
	FieldDateFrom = "date_from"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldOtherType holds the string denoting the other_type field in the database.
	FieldOtherType = "other_type"
	// EdgeKittens holds the string denoting the kittens edge name in mutations.
	EdgeKittens = "kittens"
	// EdgeBreed holds the string denoting the breed edge name in mutations.
	EdgeBreed = "breed"
	// Table holds the table name of the cat in the database.
	Table = "cats"
	// KittensTable is the table that holds the kittens relation/edge.
	KittensTable = "kittens"
	// KittensInverseTable is the table name for the Kitten entity.
	// It exists in this package in order to avoid circular dependency with the "kitten" package.
	KittensInverseTable = "kittens"
	// KittensColumn is the table column denoting the kittens relation/edge.
	KittensColumn = "mother_id"
	// BreedTable is the table that holds the breed relation/edge.
	BreedTable = "cats"
	// BreedInverseTable is the table name for the Breed entity.
	// It exists in this package in order to avoid circular dependency with the "breed" package.
	BreedInverseTable = "breeds"
	// BreedColumn is the table column denoting the breed relation/edge.
	BreedColumn = "breed_id"
)

// Columns holds all SQL columns for cat fields.
var Columns = []string{
	FieldID,
	FieldDeletedTime,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldBreedID,
	FieldSpeed,
	FieldDateFrom,
	FieldType,
	FieldOtherType,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeMerch      Type = "merch"
	TypeHotel      Type = "hotel"
	TypeTournament Type = "tournament"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeMerch, TypeHotel, TypeTournament:
		return nil
	default:
		return fmt.Errorf("cat: invalid enum value for type field: %q", _type)
	}
}

// OtherType defines the type for the "other_type" enum field.
type OtherType string

// OtherType values.
const (
	OtherTypeMerch      OtherType = "merch"
	OtherTypeHotel      OtherType = "hotel"
	OtherTypeTournament OtherType = "tournament"
)

func (ot OtherType) String() string {
	return string(ot)
}

// OtherTypeValidator is a validator for the "other_type" field enum values. It is called by the builders before save.
func OtherTypeValidator(ot OtherType) error {
	switch ot {
	case OtherTypeMerch, OtherTypeHotel, OtherTypeTournament:
		return nil
	default:
		return fmt.Errorf("cat: invalid enum value for other_type field: %q", ot)
	}
}

// OrderOption defines the ordering options for the Cat queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDeletedTime orders the results by the deleted_time field.
func ByDeletedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedTime, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByBreedID orders the results by the breed_id field.
func ByBreedID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBreedID, opts...).ToFunc()
}

// BySpeed orders the results by the speed field.
func BySpeed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSpeed, opts...).ToFunc()
}

// ByDateFrom orders the results by the date_from field.
func ByDateFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateFrom, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByOtherType orders the results by the other_type field.
func ByOtherType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOtherType, opts...).ToFunc()
}

// ByKittensCount orders the results by kittens count.
func ByKittensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newKittensStep(), opts...)
	}
}

// ByKittens orders the results by kittens terms.
func ByKittens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newKittensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBreedField orders the results by breed field.
func ByBreedField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBreedStep(), sql.OrderByField(field, opts...))
	}
}
func newKittensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(KittensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, KittensTable, KittensColumn),
	)
}
func newBreedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BreedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BreedTable, BreedColumn),
	)
}