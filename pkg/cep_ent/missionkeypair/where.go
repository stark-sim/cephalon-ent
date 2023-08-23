// Code generated by ent, DO NOT EDIT.

package missionkeypair

import (
	"cephalon-ent/pkg/cep_ent/predicate"
	"cephalon-ent/pkg/enums"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldDeletedAt, v))
}

// MissionID applies equality check predicate on the "mission_id" field. It's identical to MissionIDEQ.
func MissionID(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldMissionID, v))
}

// KeyPairID applies equality check predicate on the "key_pair_id" field. It's identical to KeyPairIDEQ.
func KeyPairID(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldKeyPairID, v))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldStartedAt, v))
}

// FinishedAt applies equality check predicate on the "finished_at" field. It's identical to FinishedAtEQ.
func FinishedAt(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldFinishedAt, v))
}

// DeviceID applies equality check predicate on the "device_id" field. It's identical to DeviceIDEQ.
func DeviceID(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldDeviceID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLTE(FieldDeletedAt, v))
}

// MissionIDEQ applies the EQ predicate on the "mission_id" field.
func MissionIDEQ(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldMissionID, v))
}

// MissionIDNEQ applies the NEQ predicate on the "mission_id" field.
func MissionIDNEQ(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldMissionID, v))
}

// MissionIDIn applies the In predicate on the "mission_id" field.
func MissionIDIn(vs ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIn(FieldMissionID, vs...))
}

// MissionIDNotIn applies the NotIn predicate on the "mission_id" field.
func MissionIDNotIn(vs ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldMissionID, vs...))
}

// KeyPairIDEQ applies the EQ predicate on the "key_pair_id" field.
func KeyPairIDEQ(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldKeyPairID, v))
}

// KeyPairIDNEQ applies the NEQ predicate on the "key_pair_id" field.
func KeyPairIDNEQ(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldKeyPairID, v))
}

// KeyPairIDIn applies the In predicate on the "key_pair_id" field.
func KeyPairIDIn(vs ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIn(FieldKeyPairID, vs...))
}

// KeyPairIDNotIn applies the NotIn predicate on the "key_pair_id" field.
func KeyPairIDNotIn(vs ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldKeyPairID, vs...))
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLTE(FieldStartedAt, v))
}

// FinishedAtEQ applies the EQ predicate on the "finished_at" field.
func FinishedAtEQ(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldFinishedAt, v))
}

// FinishedAtNEQ applies the NEQ predicate on the "finished_at" field.
func FinishedAtNEQ(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldFinishedAt, v))
}

// FinishedAtIn applies the In predicate on the "finished_at" field.
func FinishedAtIn(vs ...time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIn(FieldFinishedAt, vs...))
}

// FinishedAtNotIn applies the NotIn predicate on the "finished_at" field.
func FinishedAtNotIn(vs ...time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldFinishedAt, vs...))
}

// FinishedAtGT applies the GT predicate on the "finished_at" field.
func FinishedAtGT(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGT(FieldFinishedAt, v))
}

// FinishedAtGTE applies the GTE predicate on the "finished_at" field.
func FinishedAtGTE(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGTE(FieldFinishedAt, v))
}

// FinishedAtLT applies the LT predicate on the "finished_at" field.
func FinishedAtLT(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLT(FieldFinishedAt, v))
}

// FinishedAtLTE applies the LTE predicate on the "finished_at" field.
func FinishedAtLTE(v time.Time) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLTE(FieldFinishedAt, v))
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v enums.MissionResult) predicate.MissionKeyPair {
	vc := v
	return predicate.MissionKeyPair(sql.FieldEQ(FieldResult, vc))
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v enums.MissionResult) predicate.MissionKeyPair {
	vc := v
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldResult, vc))
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...enums.MissionResult) predicate.MissionKeyPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionKeyPair(sql.FieldIn(FieldResult, v...))
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...enums.MissionResult) predicate.MissionKeyPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldResult, v...))
}

// DeviceIDEQ applies the EQ predicate on the "device_id" field.
func DeviceIDEQ(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldEQ(FieldDeviceID, v))
}

// DeviceIDNEQ applies the NEQ predicate on the "device_id" field.
func DeviceIDNEQ(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNEQ(FieldDeviceID, v))
}

// DeviceIDIn applies the In predicate on the "device_id" field.
func DeviceIDIn(vs ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIn(FieldDeviceID, vs...))
}

// DeviceIDNotIn applies the NotIn predicate on the "device_id" field.
func DeviceIDNotIn(vs ...int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotIn(FieldDeviceID, vs...))
}

// DeviceIDGT applies the GT predicate on the "device_id" field.
func DeviceIDGT(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGT(FieldDeviceID, v))
}

// DeviceIDGTE applies the GTE predicate on the "device_id" field.
func DeviceIDGTE(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldGTE(FieldDeviceID, v))
}

// DeviceIDLT applies the LT predicate on the "device_id" field.
func DeviceIDLT(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLT(FieldDeviceID, v))
}

// DeviceIDLTE applies the LTE predicate on the "device_id" field.
func DeviceIDLTE(v int64) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldLTE(FieldDeviceID, v))
}

// ResultUrlsIsNil applies the IsNil predicate on the "result_urls" field.
func ResultUrlsIsNil() predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldIsNull(FieldResultUrls))
}

// ResultUrlsNotNil applies the NotNil predicate on the "result_urls" field.
func ResultUrlsNotNil() predicate.MissionKeyPair {
	return predicate.MissionKeyPair(sql.FieldNotNull(FieldResultUrls))
}

// HasMission applies the HasEdge predicate on the "mission" edge.
func HasMission() predicate.MissionKeyPair {
	return predicate.MissionKeyPair(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MissionTable, MissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionWith applies the HasEdge predicate on the "mission" edge with a given conditions (other predicates).
func HasMissionWith(preds ...predicate.Mission) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(func(s *sql.Selector) {
		step := newMissionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasKeyPair applies the HasEdge predicate on the "key_pair" edge.
func HasKeyPair() predicate.MissionKeyPair {
	return predicate.MissionKeyPair(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, KeyPairTable, KeyPairColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasKeyPairWith applies the HasEdge predicate on the "key_pair" edge with a given conditions (other predicates).
func HasKeyPairWith(preds ...predicate.HmacKeyPair) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(func(s *sql.Selector) {
		step := newKeyPairStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MissionKeyPair) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MissionKeyPair) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(func(s *sql.Selector) {
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
func Not(p predicate.MissionKeyPair) predicate.MissionKeyPair {
	return predicate.MissionKeyPair(func(s *sql.Selector) {
		p(s.Not())
	})
}