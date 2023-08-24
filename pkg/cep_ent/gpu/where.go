// Code generated by ent, DO NOT EDIT.

package gpu

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldDeletedAt, v))
}

// Power applies equality check predicate on the "power" field. It's identical to PowerEQ.
func Power(v int) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldPower, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.Gpu {
	return predicate.Gpu(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Gpu {
	return predicate.Gpu(sql.FieldLTE(FieldDeletedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v enums.GpuVersion) predicate.Gpu {
	vc := v
	return predicate.Gpu(sql.FieldEQ(FieldVersion, vc))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v enums.GpuVersion) predicate.Gpu {
	vc := v
	return predicate.Gpu(sql.FieldNEQ(FieldVersion, vc))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...enums.GpuVersion) predicate.Gpu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gpu(sql.FieldIn(FieldVersion, v...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...enums.GpuVersion) predicate.Gpu {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gpu(sql.FieldNotIn(FieldVersion, v...))
}

// PowerEQ applies the EQ predicate on the "power" field.
func PowerEQ(v int) predicate.Gpu {
	return predicate.Gpu(sql.FieldEQ(FieldPower, v))
}

// PowerNEQ applies the NEQ predicate on the "power" field.
func PowerNEQ(v int) predicate.Gpu {
	return predicate.Gpu(sql.FieldNEQ(FieldPower, v))
}

// PowerIn applies the In predicate on the "power" field.
func PowerIn(vs ...int) predicate.Gpu {
	return predicate.Gpu(sql.FieldIn(FieldPower, vs...))
}

// PowerNotIn applies the NotIn predicate on the "power" field.
func PowerNotIn(vs ...int) predicate.Gpu {
	return predicate.Gpu(sql.FieldNotIn(FieldPower, vs...))
}

// PowerGT applies the GT predicate on the "power" field.
func PowerGT(v int) predicate.Gpu {
	return predicate.Gpu(sql.FieldGT(FieldPower, v))
}

// PowerGTE applies the GTE predicate on the "power" field.
func PowerGTE(v int) predicate.Gpu {
	return predicate.Gpu(sql.FieldGTE(FieldPower, v))
}

// PowerLT applies the LT predicate on the "power" field.
func PowerLT(v int) predicate.Gpu {
	return predicate.Gpu(sql.FieldLT(FieldPower, v))
}

// PowerLTE applies the LTE predicate on the "power" field.
func PowerLTE(v int) predicate.Gpu {
	return predicate.Gpu(sql.FieldLTE(FieldPower, v))
}

// HasDeviceGpuMissions applies the HasEdge predicate on the "device_gpu_missions" edge.
func HasDeviceGpuMissions() predicate.Gpu {
	return predicate.Gpu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DeviceGpuMissionsTable, DeviceGpuMissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeviceGpuMissionsWith applies the HasEdge predicate on the "device_gpu_missions" edge with a given conditions (other predicates).
func HasDeviceGpuMissionsWith(preds ...predicate.DeviceGpuMission) predicate.Gpu {
	return predicate.Gpu(func(s *sql.Selector) {
		step := newDeviceGpuMissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Gpu) predicate.Gpu {
	return predicate.Gpu(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Gpu) predicate.Gpu {
	return predicate.Gpu(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Gpu) predicate.Gpu {
	return predicate.Gpu(func(s *sql.Selector) {
		p(s.Not())
	})
}
