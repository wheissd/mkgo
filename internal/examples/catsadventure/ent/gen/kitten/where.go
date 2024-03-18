// Code generated by ent, DO NOT EDIT.

package kitten

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldLTE(FieldID, id))
}

// DeletedTime applies equality check predicate on the "deleted_time" field. It's identical to DeletedTimeEQ.
func DeletedTime(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldDeletedTime, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldName, v))
}

// MotherID applies equality check predicate on the "mother_id" field. It's identical to MotherIDEQ.
func MotherID(v uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldMotherID, v))
}

// DeletedTimeEQ applies the EQ predicate on the "deleted_time" field.
func DeletedTimeEQ(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldDeletedTime, v))
}

// DeletedTimeNEQ applies the NEQ predicate on the "deleted_time" field.
func DeletedTimeNEQ(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldNEQ(FieldDeletedTime, v))
}

// DeletedTimeIn applies the In predicate on the "deleted_time" field.
func DeletedTimeIn(vs ...time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldIn(FieldDeletedTime, vs...))
}

// DeletedTimeNotIn applies the NotIn predicate on the "deleted_time" field.
func DeletedTimeNotIn(vs ...time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldNotIn(FieldDeletedTime, vs...))
}

// DeletedTimeGT applies the GT predicate on the "deleted_time" field.
func DeletedTimeGT(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldGT(FieldDeletedTime, v))
}

// DeletedTimeGTE applies the GTE predicate on the "deleted_time" field.
func DeletedTimeGTE(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldGTE(FieldDeletedTime, v))
}

// DeletedTimeLT applies the LT predicate on the "deleted_time" field.
func DeletedTimeLT(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldLT(FieldDeletedTime, v))
}

// DeletedTimeLTE applies the LTE predicate on the "deleted_time" field.
func DeletedTimeLTE(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldLTE(FieldDeletedTime, v))
}

// DeletedTimeIsNil applies the IsNil predicate on the "deleted_time" field.
func DeletedTimeIsNil() predicate.Kitten {
	return predicate.Kitten(sql.FieldIsNull(FieldDeletedTime))
}

// DeletedTimeNotNil applies the NotNil predicate on the "deleted_time" field.
func DeletedTimeNotNil() predicate.Kitten {
	return predicate.Kitten(sql.FieldNotNull(FieldDeletedTime))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Kitten {
	return predicate.Kitten(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Kitten {
	return predicate.Kitten(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Kitten {
	return predicate.Kitten(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Kitten {
	return predicate.Kitten(sql.FieldContainsFold(FieldName, v))
}

// MotherIDEQ applies the EQ predicate on the "mother_id" field.
func MotherIDEQ(v uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldEQ(FieldMotherID, v))
}

// MotherIDNEQ applies the NEQ predicate on the "mother_id" field.
func MotherIDNEQ(v uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldNEQ(FieldMotherID, v))
}

// MotherIDIn applies the In predicate on the "mother_id" field.
func MotherIDIn(vs ...uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldIn(FieldMotherID, vs...))
}

// MotherIDNotIn applies the NotIn predicate on the "mother_id" field.
func MotherIDNotIn(vs ...uuid.UUID) predicate.Kitten {
	return predicate.Kitten(sql.FieldNotIn(FieldMotherID, vs...))
}

// HasMother applies the HasEdge predicate on the "mother" edge.
func HasMother() predicate.Kitten {
	return predicate.Kitten(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MotherTable, MotherColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMotherWith applies the HasEdge predicate on the "mother" edge with a given conditions (other predicates).
func HasMotherWith(preds ...predicate.Cat) predicate.Kitten {
	return predicate.Kitten(func(s *sql.Selector) {
		step := newMotherStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Kitten) predicate.Kitten {
	return predicate.Kitten(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Kitten) predicate.Kitten {
	return predicate.Kitten(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Kitten) predicate.Kitten {
	return predicate.Kitten(sql.NotPredicates(p))
}
