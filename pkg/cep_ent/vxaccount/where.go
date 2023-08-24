// Code generated by ent, DO NOT EDIT.

package vxaccount

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldDeletedAt, v))
}

// OpenID applies equality check predicate on the "open_id" field. It's identical to OpenIDEQ.
func OpenID(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldOpenID, v))
}

// UnionID applies equality check predicate on the "union_id" field. It's identical to UnionIDEQ.
func UnionID(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldUnionID, v))
}

// Scope applies equality check predicate on the "scope" field. It's identical to ScopeEQ.
func Scope(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldScope, v))
}

// SessionKey applies equality check predicate on the "session_key" field. It's identical to SessionKeyEQ.
func SessionKey(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldSessionKey, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldUserID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLTE(FieldDeletedAt, v))
}

// OpenIDEQ applies the EQ predicate on the "open_id" field.
func OpenIDEQ(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldOpenID, v))
}

// OpenIDNEQ applies the NEQ predicate on the "open_id" field.
func OpenIDNEQ(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNEQ(FieldOpenID, v))
}

// OpenIDIn applies the In predicate on the "open_id" field.
func OpenIDIn(vs ...string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldIn(FieldOpenID, vs...))
}

// OpenIDNotIn applies the NotIn predicate on the "open_id" field.
func OpenIDNotIn(vs ...string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNotIn(FieldOpenID, vs...))
}

// OpenIDGT applies the GT predicate on the "open_id" field.
func OpenIDGT(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGT(FieldOpenID, v))
}

// OpenIDGTE applies the GTE predicate on the "open_id" field.
func OpenIDGTE(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGTE(FieldOpenID, v))
}

// OpenIDLT applies the LT predicate on the "open_id" field.
func OpenIDLT(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLT(FieldOpenID, v))
}

// OpenIDLTE applies the LTE predicate on the "open_id" field.
func OpenIDLTE(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLTE(FieldOpenID, v))
}

// OpenIDContains applies the Contains predicate on the "open_id" field.
func OpenIDContains(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldContains(FieldOpenID, v))
}

// OpenIDHasPrefix applies the HasPrefix predicate on the "open_id" field.
func OpenIDHasPrefix(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldHasPrefix(FieldOpenID, v))
}

// OpenIDHasSuffix applies the HasSuffix predicate on the "open_id" field.
func OpenIDHasSuffix(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldHasSuffix(FieldOpenID, v))
}

// OpenIDEqualFold applies the EqualFold predicate on the "open_id" field.
func OpenIDEqualFold(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEqualFold(FieldOpenID, v))
}

// OpenIDContainsFold applies the ContainsFold predicate on the "open_id" field.
func OpenIDContainsFold(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldContainsFold(FieldOpenID, v))
}

// UnionIDEQ applies the EQ predicate on the "union_id" field.
func UnionIDEQ(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldUnionID, v))
}

// UnionIDNEQ applies the NEQ predicate on the "union_id" field.
func UnionIDNEQ(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNEQ(FieldUnionID, v))
}

// UnionIDIn applies the In predicate on the "union_id" field.
func UnionIDIn(vs ...string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldIn(FieldUnionID, vs...))
}

// UnionIDNotIn applies the NotIn predicate on the "union_id" field.
func UnionIDNotIn(vs ...string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNotIn(FieldUnionID, vs...))
}

// UnionIDGT applies the GT predicate on the "union_id" field.
func UnionIDGT(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGT(FieldUnionID, v))
}

// UnionIDGTE applies the GTE predicate on the "union_id" field.
func UnionIDGTE(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGTE(FieldUnionID, v))
}

// UnionIDLT applies the LT predicate on the "union_id" field.
func UnionIDLT(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLT(FieldUnionID, v))
}

// UnionIDLTE applies the LTE predicate on the "union_id" field.
func UnionIDLTE(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLTE(FieldUnionID, v))
}

// UnionIDContains applies the Contains predicate on the "union_id" field.
func UnionIDContains(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldContains(FieldUnionID, v))
}

// UnionIDHasPrefix applies the HasPrefix predicate on the "union_id" field.
func UnionIDHasPrefix(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldHasPrefix(FieldUnionID, v))
}

// UnionIDHasSuffix applies the HasSuffix predicate on the "union_id" field.
func UnionIDHasSuffix(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldHasSuffix(FieldUnionID, v))
}

// UnionIDEqualFold applies the EqualFold predicate on the "union_id" field.
func UnionIDEqualFold(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEqualFold(FieldUnionID, v))
}

// UnionIDContainsFold applies the ContainsFold predicate on the "union_id" field.
func UnionIDContainsFold(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldContainsFold(FieldUnionID, v))
}

// ScopeEQ applies the EQ predicate on the "scope" field.
func ScopeEQ(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldScope, v))
}

// ScopeNEQ applies the NEQ predicate on the "scope" field.
func ScopeNEQ(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNEQ(FieldScope, v))
}

// ScopeIn applies the In predicate on the "scope" field.
func ScopeIn(vs ...string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldIn(FieldScope, vs...))
}

// ScopeNotIn applies the NotIn predicate on the "scope" field.
func ScopeNotIn(vs ...string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNotIn(FieldScope, vs...))
}

// ScopeGT applies the GT predicate on the "scope" field.
func ScopeGT(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGT(FieldScope, v))
}

// ScopeGTE applies the GTE predicate on the "scope" field.
func ScopeGTE(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGTE(FieldScope, v))
}

// ScopeLT applies the LT predicate on the "scope" field.
func ScopeLT(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLT(FieldScope, v))
}

// ScopeLTE applies the LTE predicate on the "scope" field.
func ScopeLTE(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLTE(FieldScope, v))
}

// ScopeContains applies the Contains predicate on the "scope" field.
func ScopeContains(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldContains(FieldScope, v))
}

// ScopeHasPrefix applies the HasPrefix predicate on the "scope" field.
func ScopeHasPrefix(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldHasPrefix(FieldScope, v))
}

// ScopeHasSuffix applies the HasSuffix predicate on the "scope" field.
func ScopeHasSuffix(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldHasSuffix(FieldScope, v))
}

// ScopeEqualFold applies the EqualFold predicate on the "scope" field.
func ScopeEqualFold(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEqualFold(FieldScope, v))
}

// ScopeContainsFold applies the ContainsFold predicate on the "scope" field.
func ScopeContainsFold(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldContainsFold(FieldScope, v))
}

// SessionKeyEQ applies the EQ predicate on the "session_key" field.
func SessionKeyEQ(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldSessionKey, v))
}

// SessionKeyNEQ applies the NEQ predicate on the "session_key" field.
func SessionKeyNEQ(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNEQ(FieldSessionKey, v))
}

// SessionKeyIn applies the In predicate on the "session_key" field.
func SessionKeyIn(vs ...string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldIn(FieldSessionKey, vs...))
}

// SessionKeyNotIn applies the NotIn predicate on the "session_key" field.
func SessionKeyNotIn(vs ...string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNotIn(FieldSessionKey, vs...))
}

// SessionKeyGT applies the GT predicate on the "session_key" field.
func SessionKeyGT(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGT(FieldSessionKey, v))
}

// SessionKeyGTE applies the GTE predicate on the "session_key" field.
func SessionKeyGTE(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldGTE(FieldSessionKey, v))
}

// SessionKeyLT applies the LT predicate on the "session_key" field.
func SessionKeyLT(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLT(FieldSessionKey, v))
}

// SessionKeyLTE applies the LTE predicate on the "session_key" field.
func SessionKeyLTE(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldLTE(FieldSessionKey, v))
}

// SessionKeyContains applies the Contains predicate on the "session_key" field.
func SessionKeyContains(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldContains(FieldSessionKey, v))
}

// SessionKeyHasPrefix applies the HasPrefix predicate on the "session_key" field.
func SessionKeyHasPrefix(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldHasPrefix(FieldSessionKey, v))
}

// SessionKeyHasSuffix applies the HasSuffix predicate on the "session_key" field.
func SessionKeyHasSuffix(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldHasSuffix(FieldSessionKey, v))
}

// SessionKeyEqualFold applies the EqualFold predicate on the "session_key" field.
func SessionKeyEqualFold(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEqualFold(FieldSessionKey, v))
}

// SessionKeyContainsFold applies the ContainsFold predicate on the "session_key" field.
func SessionKeyContainsFold(v string) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldContainsFold(FieldSessionKey, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.VXAccount {
	return predicate.VXAccount(sql.FieldNotIn(FieldUserID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.VXAccount {
	return predicate.VXAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.VXAccount {
	return predicate.VXAccount(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VXAccount) predicate.VXAccount {
	return predicate.VXAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VXAccount) predicate.VXAccount {
	return predicate.VXAccount(func(s *sql.Selector) {
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
func Not(p predicate.VXAccount) predicate.VXAccount {
	return predicate.VXAccount(func(s *sql.Selector) {
		p(s.Not())
	})
}
