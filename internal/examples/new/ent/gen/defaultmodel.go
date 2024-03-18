// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/new/ent/gen/defaultmodel"
)

// DefaultModel is the model entity for the DefaultModel schema.
type DefaultModel struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DefaultModel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case defaultmodel.FieldName:
			values[i] = new(sql.NullString)
		case defaultmodel.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DefaultModel fields.
func (dm *DefaultModel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case defaultmodel.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				dm.ID = *value
			}
		case defaultmodel.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				dm.Name = value.String
			}
		default:
			dm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DefaultModel.
// This includes values selected through modifiers, order, etc.
func (dm *DefaultModel) Value(name string) (ent.Value, error) {
	return dm.selectValues.Get(name)
}

// Update returns a builder for updating this DefaultModel.
// Note that you need to call DefaultModel.Unwrap() before calling this method if this DefaultModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (dm *DefaultModel) Update() *DefaultModelUpdateOne {
	return NewDefaultModelClient(dm.config).UpdateOne(dm)
}

// Unwrap unwraps the DefaultModel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dm *DefaultModel) Unwrap() *DefaultModel {
	_tx, ok := dm.config.driver.(*txDriver)
	if !ok {
		panic("gen: DefaultModel is not a transactional entity")
	}
	dm.config.driver = _tx.drv
	return dm
}

// String implements the fmt.Stringer.
func (dm *DefaultModel) String() string {
	var builder strings.Builder
	builder.WriteString("DefaultModel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dm.ID))
	builder.WriteString("name=")
	builder.WriteString(dm.Name)
	builder.WriteByte(')')
	return builder.String()
}

// DefaultModels is a parsable slice of DefaultModel.
type DefaultModels []*DefaultModel
