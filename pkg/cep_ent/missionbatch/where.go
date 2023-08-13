// Code generated by ent, DO NOT EDIT.

package missionbatch

import (
	"cephalon-ent/pkg/cep_ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldDeletedAt, v))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldNumber, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldUserID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLTE(FieldDeletedAt, v))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldLTE(FieldNumber, v))
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldContains(FieldNumber, v))
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldHasPrefix(FieldNumber, v))
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldHasSuffix(FieldNumber, v))
}

// NumberEqualFold applies the EqualFold predicate on the "number" field.
func NumberEqualFold(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEqualFold(FieldNumber, v))
}

// NumberContainsFold applies the ContainsFold predicate on the "number" field.
func NumberContainsFold(v string) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldContainsFold(FieldNumber, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.MissionBatch {
	return predicate.MissionBatch(sql.FieldNotIn(FieldUserID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.MissionBatch {
	return predicate.MissionBatch(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.MissionBatch {
	return predicate.MissionBatch(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMissions applies the HasEdge predicate on the "missions" edge.
func HasMissions() predicate.MissionBatch {
	return predicate.MissionBatch(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MissionsTable, MissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionsWith applies the HasEdge predicate on the "missions" edge with a given conditions (other predicates).
func HasMissionsWith(preds ...predicate.Mission) predicate.MissionBatch {
	return predicate.MissionBatch(func(s *sql.Selector) {
		step := newMissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMissionConsumeOrders applies the HasEdge predicate on the "mission_consume_orders" edge.
func HasMissionConsumeOrders() predicate.MissionBatch {
	return predicate.MissionBatch(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MissionConsumeOrdersTable, MissionConsumeOrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionConsumeOrdersWith applies the HasEdge predicate on the "mission_consume_orders" edge with a given conditions (other predicates).
func HasMissionConsumeOrdersWith(preds ...predicate.MissionConsumeOrder) predicate.MissionBatch {
	return predicate.MissionBatch(func(s *sql.Selector) {
		step := newMissionConsumeOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMissionProduceOrders applies the HasEdge predicate on the "mission_produce_orders" edge.
func HasMissionProduceOrders() predicate.MissionBatch {
	return predicate.MissionBatch(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MissionProduceOrdersTable, MissionProduceOrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionProduceOrdersWith applies the HasEdge predicate on the "mission_produce_orders" edge with a given conditions (other predicates).
func HasMissionProduceOrdersWith(preds ...predicate.MissionProduceOrder) predicate.MissionBatch {
	return predicate.MissionBatch(func(s *sql.Selector) {
		step := newMissionProduceOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MissionBatch) predicate.MissionBatch {
	return predicate.MissionBatch(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MissionBatch) predicate.MissionBatch {
	return predicate.MissionBatch(func(s *sql.Selector) {
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
func Not(p predicate.MissionBatch) predicate.MissionBatch {
	return predicate.MissionBatch(func(s *sql.Selector) {
		p(s.Not())
	})
}
