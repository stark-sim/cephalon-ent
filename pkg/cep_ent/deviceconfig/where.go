// Code generated by ent, DO NOT EDIT.

package deviceconfig

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldDeletedAt, v))
}

// DeviceID applies equality check predicate on the "device_id" field. It's identical to DeviceIDEQ.
func DeviceID(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldDeviceID, v))
}

// GapBase applies equality check predicate on the "gap_base" field. It's identical to GapBaseEQ.
func GapBase(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldGapBase, v))
}

// GapRandomMax applies equality check predicate on the "gap_random_max" field. It's identical to GapRandomMaxEQ.
func GapRandomMax(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldGapRandomMax, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLTE(FieldDeletedAt, v))
}

// DeviceIDEQ applies the EQ predicate on the "device_id" field.
func DeviceIDEQ(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldDeviceID, v))
}

// DeviceIDNEQ applies the NEQ predicate on the "device_id" field.
func DeviceIDNEQ(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNEQ(FieldDeviceID, v))
}

// DeviceIDIn applies the In predicate on the "device_id" field.
func DeviceIDIn(vs ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldIn(FieldDeviceID, vs...))
}

// DeviceIDNotIn applies the NotIn predicate on the "device_id" field.
func DeviceIDNotIn(vs ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNotIn(FieldDeviceID, vs...))
}

// GapBaseEQ applies the EQ predicate on the "gap_base" field.
func GapBaseEQ(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldGapBase, v))
}

// GapBaseNEQ applies the NEQ predicate on the "gap_base" field.
func GapBaseNEQ(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNEQ(FieldGapBase, v))
}

// GapBaseIn applies the In predicate on the "gap_base" field.
func GapBaseIn(vs ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldIn(FieldGapBase, vs...))
}

// GapBaseNotIn applies the NotIn predicate on the "gap_base" field.
func GapBaseNotIn(vs ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNotIn(FieldGapBase, vs...))
}

// GapBaseGT applies the GT predicate on the "gap_base" field.
func GapBaseGT(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGT(FieldGapBase, v))
}

// GapBaseGTE applies the GTE predicate on the "gap_base" field.
func GapBaseGTE(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGTE(FieldGapBase, v))
}

// GapBaseLT applies the LT predicate on the "gap_base" field.
func GapBaseLT(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLT(FieldGapBase, v))
}

// GapBaseLTE applies the LTE predicate on the "gap_base" field.
func GapBaseLTE(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLTE(FieldGapBase, v))
}

// GapRandomMaxEQ applies the EQ predicate on the "gap_random_max" field.
func GapRandomMaxEQ(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldEQ(FieldGapRandomMax, v))
}

// GapRandomMaxNEQ applies the NEQ predicate on the "gap_random_max" field.
func GapRandomMaxNEQ(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNEQ(FieldGapRandomMax, v))
}

// GapRandomMaxIn applies the In predicate on the "gap_random_max" field.
func GapRandomMaxIn(vs ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldIn(FieldGapRandomMax, vs...))
}

// GapRandomMaxNotIn applies the NotIn predicate on the "gap_random_max" field.
func GapRandomMaxNotIn(vs ...int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldNotIn(FieldGapRandomMax, vs...))
}

// GapRandomMaxGT applies the GT predicate on the "gap_random_max" field.
func GapRandomMaxGT(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGT(FieldGapRandomMax, v))
}

// GapRandomMaxGTE applies the GTE predicate on the "gap_random_max" field.
func GapRandomMaxGTE(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldGTE(FieldGapRandomMax, v))
}

// GapRandomMaxLT applies the LT predicate on the "gap_random_max" field.
func GapRandomMaxLT(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLT(FieldGapRandomMax, v))
}

// GapRandomMaxLTE applies the LTE predicate on the "gap_random_max" field.
func GapRandomMaxLTE(v int64) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.FieldLTE(FieldGapRandomMax, v))
}

// HasDevice applies the HasEdge predicate on the "device" edge.
func HasDevice() predicate.DeviceConfig {
	return predicate.DeviceConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, DeviceTable, DeviceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeviceWith applies the HasEdge predicate on the "device" edge with a given conditions (other predicates).
func HasDeviceWith(preds ...predicate.Device) predicate.DeviceConfig {
	return predicate.DeviceConfig(func(s *sql.Selector) {
		step := newDeviceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeviceConfig) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeviceConfig) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DeviceConfig) predicate.DeviceConfig {
	return predicate.DeviceConfig(sql.NotPredicates(p))
}
