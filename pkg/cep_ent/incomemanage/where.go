// Code generated by ent, DO NOT EDIT.

package incomemanage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldUserID, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldPhone, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldAmount, v))
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldReason, v))
}

// CurrentBalance applies equality check predicate on the "current_balance" field. It's identical to CurrentBalanceEQ.
func CurrentBalance(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldCurrentBalance, v))
}

// LastEditedAt applies equality check predicate on the "last_edited_at" field. It's identical to LastEditedAtEQ.
func LastEditedAt(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldLastEditedAt, v))
}

// RejectReason applies equality check predicate on the "reject_reason" field. It's identical to RejectReasonEQ.
func RejectReason(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldRejectReason, v))
}

// ApproveUserID applies equality check predicate on the "approve_user_id" field. It's identical to ApproveUserIDEQ.
func ApproveUserID(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldApproveUserID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldDeletedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldUserID, vs...))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldContainsFold(FieldPhone, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v enums.IncomeManageType) predicate.IncomeManage {
	vc := v
	return predicate.IncomeManage(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v enums.IncomeManageType) predicate.IncomeManage {
	vc := v
	return predicate.IncomeManage(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...enums.IncomeManageType) predicate.IncomeManage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IncomeManage(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...enums.IncomeManageType) predicate.IncomeManage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IncomeManage(sql.FieldNotIn(FieldType, v...))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldAmount, v))
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldReason, v))
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldReason, v))
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldReason, vs...))
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldReason, vs...))
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldReason, v))
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldReason, v))
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldReason, v))
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldReason, v))
}

// ReasonContains applies the Contains predicate on the "reason" field.
func ReasonContains(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldContains(FieldReason, v))
}

// ReasonHasPrefix applies the HasPrefix predicate on the "reason" field.
func ReasonHasPrefix(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldHasPrefix(FieldReason, v))
}

// ReasonHasSuffix applies the HasSuffix predicate on the "reason" field.
func ReasonHasSuffix(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldHasSuffix(FieldReason, v))
}

// ReasonEqualFold applies the EqualFold predicate on the "reason" field.
func ReasonEqualFold(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEqualFold(FieldReason, v))
}

// ReasonContainsFold applies the ContainsFold predicate on the "reason" field.
func ReasonContainsFold(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldContainsFold(FieldReason, v))
}

// CurrentBalanceEQ applies the EQ predicate on the "current_balance" field.
func CurrentBalanceEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldCurrentBalance, v))
}

// CurrentBalanceNEQ applies the NEQ predicate on the "current_balance" field.
func CurrentBalanceNEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldCurrentBalance, v))
}

// CurrentBalanceIn applies the In predicate on the "current_balance" field.
func CurrentBalanceIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldCurrentBalance, vs...))
}

// CurrentBalanceNotIn applies the NotIn predicate on the "current_balance" field.
func CurrentBalanceNotIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldCurrentBalance, vs...))
}

// CurrentBalanceGT applies the GT predicate on the "current_balance" field.
func CurrentBalanceGT(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldCurrentBalance, v))
}

// CurrentBalanceGTE applies the GTE predicate on the "current_balance" field.
func CurrentBalanceGTE(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldCurrentBalance, v))
}

// CurrentBalanceLT applies the LT predicate on the "current_balance" field.
func CurrentBalanceLT(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldCurrentBalance, v))
}

// CurrentBalanceLTE applies the LTE predicate on the "current_balance" field.
func CurrentBalanceLTE(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldCurrentBalance, v))
}

// LastEditedAtEQ applies the EQ predicate on the "last_edited_at" field.
func LastEditedAtEQ(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldLastEditedAt, v))
}

// LastEditedAtNEQ applies the NEQ predicate on the "last_edited_at" field.
func LastEditedAtNEQ(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldLastEditedAt, v))
}

// LastEditedAtIn applies the In predicate on the "last_edited_at" field.
func LastEditedAtIn(vs ...time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldLastEditedAt, vs...))
}

// LastEditedAtNotIn applies the NotIn predicate on the "last_edited_at" field.
func LastEditedAtNotIn(vs ...time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldLastEditedAt, vs...))
}

// LastEditedAtGT applies the GT predicate on the "last_edited_at" field.
func LastEditedAtGT(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldLastEditedAt, v))
}

// LastEditedAtGTE applies the GTE predicate on the "last_edited_at" field.
func LastEditedAtGTE(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldLastEditedAt, v))
}

// LastEditedAtLT applies the LT predicate on the "last_edited_at" field.
func LastEditedAtLT(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldLastEditedAt, v))
}

// LastEditedAtLTE applies the LTE predicate on the "last_edited_at" field.
func LastEditedAtLTE(v time.Time) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldLastEditedAt, v))
}

// RejectReasonEQ applies the EQ predicate on the "reject_reason" field.
func RejectReasonEQ(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldRejectReason, v))
}

// RejectReasonNEQ applies the NEQ predicate on the "reject_reason" field.
func RejectReasonNEQ(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldRejectReason, v))
}

// RejectReasonIn applies the In predicate on the "reject_reason" field.
func RejectReasonIn(vs ...string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldRejectReason, vs...))
}

// RejectReasonNotIn applies the NotIn predicate on the "reject_reason" field.
func RejectReasonNotIn(vs ...string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldRejectReason, vs...))
}

// RejectReasonGT applies the GT predicate on the "reject_reason" field.
func RejectReasonGT(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGT(FieldRejectReason, v))
}

// RejectReasonGTE applies the GTE predicate on the "reject_reason" field.
func RejectReasonGTE(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldGTE(FieldRejectReason, v))
}

// RejectReasonLT applies the LT predicate on the "reject_reason" field.
func RejectReasonLT(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLT(FieldRejectReason, v))
}

// RejectReasonLTE applies the LTE predicate on the "reject_reason" field.
func RejectReasonLTE(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldLTE(FieldRejectReason, v))
}

// RejectReasonContains applies the Contains predicate on the "reject_reason" field.
func RejectReasonContains(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldContains(FieldRejectReason, v))
}

// RejectReasonHasPrefix applies the HasPrefix predicate on the "reject_reason" field.
func RejectReasonHasPrefix(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldHasPrefix(FieldRejectReason, v))
}

// RejectReasonHasSuffix applies the HasSuffix predicate on the "reject_reason" field.
func RejectReasonHasSuffix(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldHasSuffix(FieldRejectReason, v))
}

// RejectReasonEqualFold applies the EqualFold predicate on the "reject_reason" field.
func RejectReasonEqualFold(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEqualFold(FieldRejectReason, v))
}

// RejectReasonContainsFold applies the ContainsFold predicate on the "reject_reason" field.
func RejectReasonContainsFold(v string) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldContainsFold(FieldRejectReason, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v enums.IncomeManageStatus) predicate.IncomeManage {
	vc := v
	return predicate.IncomeManage(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v enums.IncomeManageStatus) predicate.IncomeManage {
	vc := v
	return predicate.IncomeManage(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...enums.IncomeManageStatus) predicate.IncomeManage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IncomeManage(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...enums.IncomeManageStatus) predicate.IncomeManage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IncomeManage(sql.FieldNotIn(FieldStatus, v...))
}

// ApproveUserIDEQ applies the EQ predicate on the "approve_user_id" field.
func ApproveUserIDEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldEQ(FieldApproveUserID, v))
}

// ApproveUserIDNEQ applies the NEQ predicate on the "approve_user_id" field.
func ApproveUserIDNEQ(v int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNEQ(FieldApproveUserID, v))
}

// ApproveUserIDIn applies the In predicate on the "approve_user_id" field.
func ApproveUserIDIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldIn(FieldApproveUserID, vs...))
}

// ApproveUserIDNotIn applies the NotIn predicate on the "approve_user_id" field.
func ApproveUserIDNotIn(vs ...int64) predicate.IncomeManage {
	return predicate.IncomeManage(sql.FieldNotIn(FieldApproveUserID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.IncomeManage {
	return predicate.IncomeManage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.IncomeManage {
	return predicate.IncomeManage(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasApproveUser applies the HasEdge predicate on the "approve_user" edge.
func HasApproveUser() predicate.IncomeManage {
	return predicate.IncomeManage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApproveUserTable, ApproveUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApproveUserWith applies the HasEdge predicate on the "approve_user" edge with a given conditions (other predicates).
func HasApproveUserWith(preds ...predicate.User) predicate.IncomeManage {
	return predicate.IncomeManage(func(s *sql.Selector) {
		step := newApproveUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.IncomeManage) predicate.IncomeManage {
	return predicate.IncomeManage(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.IncomeManage) predicate.IncomeManage {
	return predicate.IncomeManage(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.IncomeManage) predicate.IncomeManage {
	return predicate.IncomeManage(sql.NotPredicates(p))
}
