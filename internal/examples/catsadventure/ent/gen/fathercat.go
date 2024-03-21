// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/fathercat"
)

// FatherCat is the model entity for the FatherCat schema.
type FatherCat struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// DeletedTime holds the value of the "deleted_time" field.
	DeletedTime *time.Time `json:"deleted_time,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FatherCat) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case fathercat.FieldName:
			values[i] = new(sql.NullString)
		case fathercat.FieldDeletedTime, fathercat.FieldCreateTime, fathercat.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case fathercat.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FatherCat fields.
func (fc *FatherCat) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fathercat.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fc.ID = *value
			}
		case fathercat.FieldDeletedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_time", values[i])
			} else if value.Valid {
				fc.DeletedTime = new(time.Time)
				*fc.DeletedTime = value.Time
			}
		case fathercat.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				fc.CreateTime = value.Time
			}
		case fathercat.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				fc.UpdateTime = value.Time
			}
		case fathercat.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				fc.Name = value.String
			}
		default:
			fc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FatherCat.
// This includes values selected through modifiers, order, etc.
func (fc *FatherCat) Value(name string) (ent.Value, error) {
	return fc.selectValues.Get(name)
}

// Update returns a builder for updating this FatherCat.
// Note that you need to call FatherCat.Unwrap() before calling this method if this FatherCat
// was returned from a transaction, and the transaction was committed or rolled back.
func (fc *FatherCat) Update() *FatherCatUpdateOne {
	return NewFatherCatClient(fc.config).UpdateOne(fc)
}

// Unwrap unwraps the FatherCat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fc *FatherCat) Unwrap() *FatherCat {
	_tx, ok := fc.config.driver.(*txDriver)
	if !ok {
		panic("gen: FatherCat is not a transactional entity")
	}
	fc.config.driver = _tx.drv
	return fc
}

// String implements the fmt.Stringer.
func (fc *FatherCat) String() string {
	var builder strings.Builder
	builder.WriteString("FatherCat(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fc.ID))
	if v := fc.DeletedTime; v != nil {
		builder.WriteString("deleted_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(fc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(fc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(fc.Name)
	builder.WriteByte(')')
	return builder.String()
}

// FatherCats is a parsable slice of FatherCat.
type FatherCats []*FatherCat