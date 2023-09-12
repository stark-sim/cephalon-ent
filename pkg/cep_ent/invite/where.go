// Code generated by ent, DO NOT EDIT.

package invite

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldDeletedAt, v))
}

// InviteCode applies equality check predicate on the "invite_code" field. It's identical to InviteCodeEQ.
func InviteCode(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldInviteCode, v))
}

// ShareCep applies equality check predicate on the "share_cep" field. It's identical to ShareCepEQ.
func ShareCep(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldShareCep, v))
}

// RegCep applies equality check predicate on the "reg_cep" field. It's identical to RegCepEQ.
func RegCep(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldRegCep, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldType, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldUserID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldDeletedAt, v))
}

// InviteCodeEQ applies the EQ predicate on the "invite_code" field.
func InviteCodeEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldInviteCode, v))
}

// InviteCodeNEQ applies the NEQ predicate on the "invite_code" field.
func InviteCodeNEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldInviteCode, v))
}

// InviteCodeIn applies the In predicate on the "invite_code" field.
func InviteCodeIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldInviteCode, vs...))
}

// InviteCodeNotIn applies the NotIn predicate on the "invite_code" field.
func InviteCodeNotIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldInviteCode, vs...))
}

// InviteCodeGT applies the GT predicate on the "invite_code" field.
func InviteCodeGT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldInviteCode, v))
}

// InviteCodeGTE applies the GTE predicate on the "invite_code" field.
func InviteCodeGTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldInviteCode, v))
}

// InviteCodeLT applies the LT predicate on the "invite_code" field.
func InviteCodeLT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldInviteCode, v))
}

// InviteCodeLTE applies the LTE predicate on the "invite_code" field.
func InviteCodeLTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldInviteCode, v))
}

// InviteCodeContains applies the Contains predicate on the "invite_code" field.
func InviteCodeContains(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContains(FieldInviteCode, v))
}

// InviteCodeHasPrefix applies the HasPrefix predicate on the "invite_code" field.
func InviteCodeHasPrefix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasPrefix(FieldInviteCode, v))
}

// InviteCodeHasSuffix applies the HasSuffix predicate on the "invite_code" field.
func InviteCodeHasSuffix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasSuffix(FieldInviteCode, v))
}

// InviteCodeEqualFold applies the EqualFold predicate on the "invite_code" field.
func InviteCodeEqualFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEqualFold(FieldInviteCode, v))
}

// InviteCodeContainsFold applies the ContainsFold predicate on the "invite_code" field.
func InviteCodeContainsFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContainsFold(FieldInviteCode, v))
}

// ShareCepEQ applies the EQ predicate on the "share_cep" field.
func ShareCepEQ(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldShareCep, v))
}

// ShareCepNEQ applies the NEQ predicate on the "share_cep" field.
func ShareCepNEQ(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldShareCep, v))
}

// ShareCepIn applies the In predicate on the "share_cep" field.
func ShareCepIn(vs ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldShareCep, vs...))
}

// ShareCepNotIn applies the NotIn predicate on the "share_cep" field.
func ShareCepNotIn(vs ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldShareCep, vs...))
}

// ShareCepGT applies the GT predicate on the "share_cep" field.
func ShareCepGT(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldShareCep, v))
}

// ShareCepGTE applies the GTE predicate on the "share_cep" field.
func ShareCepGTE(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldShareCep, v))
}

// ShareCepLT applies the LT predicate on the "share_cep" field.
func ShareCepLT(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldShareCep, v))
}

// ShareCepLTE applies the LTE predicate on the "share_cep" field.
func ShareCepLTE(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldShareCep, v))
}

// RegCepEQ applies the EQ predicate on the "reg_cep" field.
func RegCepEQ(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldRegCep, v))
}

// RegCepNEQ applies the NEQ predicate on the "reg_cep" field.
func RegCepNEQ(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldRegCep, v))
}

// RegCepIn applies the In predicate on the "reg_cep" field.
func RegCepIn(vs ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldRegCep, vs...))
}

// RegCepNotIn applies the NotIn predicate on the "reg_cep" field.
func RegCepNotIn(vs ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldRegCep, vs...))
}

// RegCepGT applies the GT predicate on the "reg_cep" field.
func RegCepGT(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldRegCep, v))
}

// RegCepGTE applies the GTE predicate on the "reg_cep" field.
func RegCepGTE(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldRegCep, v))
}

// RegCepLT applies the LT predicate on the "reg_cep" field.
func RegCepLT(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldRegCep, v))
}

// RegCepLTE applies the LTE predicate on the "reg_cep" field.
func RegCepLTE(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldRegCep, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContainsFold(FieldType, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldUserID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Invite {
	return predicate.Invite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Invite {
	return predicate.Invite(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Invite) predicate.Invite {
	return predicate.Invite(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Invite) predicate.Invite {
	return predicate.Invite(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Invite) predicate.Invite {
	return predicate.Invite(sql.NotPredicates(p))
}
