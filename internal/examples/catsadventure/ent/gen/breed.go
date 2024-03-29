// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/breed"
)

// Breed is the model entity for the Breed schema.
type Breed struct {
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
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BreedQuery when eager-loading is set.
	Edges        BreedEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BreedEdges holds the relations/edges for other nodes in the graph.
type BreedEdges struct {
	// Cats holds the value of the cats edge.
	Cats []*Cat `json:"cats,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CatsOrErr returns the Cats value or an error if the edge
// was not loaded in eager-loading.
func (e BreedEdges) CatsOrErr() ([]*Cat, error) {
	if e.loadedTypes[0] {
		return e.Cats, nil
	}
	return nil, &NotLoadedError{edge: "cats"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Breed) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case breed.FieldName:
			values[i] = new(sql.NullString)
		case breed.FieldDeletedTime, breed.FieldCreateTime, breed.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case breed.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Breed fields.
func (b *Breed) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case breed.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case breed.FieldDeletedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_time", values[i])
			} else if value.Valid {
				b.DeletedTime = new(time.Time)
				*b.DeletedTime = value.Time
			}
		case breed.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				b.CreateTime = value.Time
			}
		case breed.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				b.UpdateTime = value.Time
			}
		case breed.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Breed.
// This includes values selected through modifiers, order, etc.
func (b *Breed) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryCats queries the "cats" edge of the Breed entity.
func (b *Breed) QueryCats() *CatQuery {
	return NewBreedClient(b.config).QueryCats(b)
}

// Update returns a builder for updating this Breed.
// Note that you need to call Breed.Unwrap() before calling this method if this Breed
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Breed) Update() *BreedUpdateOne {
	return NewBreedClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Breed entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Breed) Unwrap() *Breed {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("gen: Breed is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Breed) String() string {
	var builder strings.Builder
	builder.WriteString("Breed(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	if v := b.DeletedTime; v != nil {
		builder.WriteString("deleted_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(b.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(b.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Breeds is a parsable slice of Breed.
type Breeds []*Breed
