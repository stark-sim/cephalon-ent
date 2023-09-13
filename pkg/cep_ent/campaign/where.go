// Code generated by ent, DO NOT EDIT.

package campaign

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldName, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldType, v))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldStartedAt, v))
}

// EndedAt applies equality check predicate on the "ended_at" field. It's identical to EndedAtEQ.
func EndedAt(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldEndedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldStatus, v))
}

// InviteID applies equality check predicate on the "invite_id" field. It's identical to InviteIDEQ.
func InviteID(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldInviteID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldDeletedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldContainsFold(FieldType, v))
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldStartedAt, v))
}

// EndedAtEQ applies the EQ predicate on the "ended_at" field.
func EndedAtEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldEndedAt, v))
}

// EndedAtNEQ applies the NEQ predicate on the "ended_at" field.
func EndedAtNEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldEndedAt, v))
}

// EndedAtIn applies the In predicate on the "ended_at" field.
func EndedAtIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldEndedAt, vs...))
}

// EndedAtNotIn applies the NotIn predicate on the "ended_at" field.
func EndedAtNotIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldEndedAt, vs...))
}

// EndedAtGT applies the GT predicate on the "ended_at" field.
func EndedAtGT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldEndedAt, v))
}

// EndedAtGTE applies the GTE predicate on the "ended_at" field.
func EndedAtGTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldEndedAt, v))
}

// EndedAtLT applies the LT predicate on the "ended_at" field.
func EndedAtLT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldEndedAt, v))
}

// EndedAtLTE applies the LTE predicate on the "ended_at" field.
func EndedAtLTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldEndedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldStatus, v))
}

// InviteIDEQ applies the EQ predicate on the "invite_id" field.
func InviteIDEQ(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldInviteID, v))
}

// InviteIDNEQ applies the NEQ predicate on the "invite_id" field.
func InviteIDNEQ(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldInviteID, v))
}

// InviteIDIn applies the In predicate on the "invite_id" field.
func InviteIDIn(vs ...string) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldInviteID, vs...))
}

// InviteIDNotIn applies the NotIn predicate on the "invite_id" field.
func InviteIDNotIn(vs ...string) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldInviteID, vs...))
}

// InviteIDGT applies the GT predicate on the "invite_id" field.
func InviteIDGT(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldInviteID, v))
}

// InviteIDGTE applies the GTE predicate on the "invite_id" field.
func InviteIDGTE(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldInviteID, v))
}

// InviteIDLT applies the LT predicate on the "invite_id" field.
func InviteIDLT(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldInviteID, v))
}

// InviteIDLTE applies the LTE predicate on the "invite_id" field.
func InviteIDLTE(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldInviteID, v))
}

// InviteIDContains applies the Contains predicate on the "invite_id" field.
func InviteIDContains(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldContains(FieldInviteID, v))
}

// InviteIDHasPrefix applies the HasPrefix predicate on the "invite_id" field.
func InviteIDHasPrefix(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldHasPrefix(FieldInviteID, v))
}

// InviteIDHasSuffix applies the HasSuffix predicate on the "invite_id" field.
func InviteIDHasSuffix(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldHasSuffix(FieldInviteID, v))
}

// InviteIDEqualFold applies the EqualFold predicate on the "invite_id" field.
func InviteIDEqualFold(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEqualFold(FieldInviteID, v))
}

// InviteIDContainsFold applies the ContainsFold predicate on the "invite_id" field.
func InviteIDContainsFold(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldContainsFold(FieldInviteID, v))
}

// HasInvites applies the HasEdge predicate on the "invites" edge.
func HasInvites() predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InvitesTable, InvitesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvitesWith applies the HasEdge predicate on the "invites" edge with a given conditions (other predicates).
func HasInvitesWith(preds ...predicate.Invite) predicate.Campaign {
	return predicate.Campaign(func(s *sql.Selector) {
		step := newInvitesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Campaign) predicate.Campaign {
	return predicate.Campaign(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Campaign) predicate.Campaign {
	return predicate.Campaign(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Campaign) predicate.Campaign {
	return predicate.Campaign(sql.NotPredicates(p))
}
