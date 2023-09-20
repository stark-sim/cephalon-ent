// Code generated by ent, DO NOT EDIT.

package bill

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldDeletedAt, v))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldOrderID, v))
}

// SymbolID applies equality check predicate on the "symbol_id" field. It's identical to SymbolIDEQ.
func SymbolID(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldSymbolID, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldAmount, v))
}

// TargetUserID applies equality check predicate on the "target_user_id" field. It's identical to TargetUserIDEQ.
func TargetUserID(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldTargetUserID, v))
}

// TargetBeforeAmount applies equality check predicate on the "target_before_amount" field. It's identical to TargetBeforeAmountEQ.
func TargetBeforeAmount(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldTargetBeforeAmount, v))
}

// TargetAfterAmount applies equality check predicate on the "target_after_amount" field. It's identical to TargetAfterAmountEQ.
func TargetAfterAmount(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldTargetAfterAmount, v))
}

// SourceUserID applies equality check predicate on the "source_user_id" field. It's identical to SourceUserIDEQ.
func SourceUserID(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldSourceUserID, v))
}

// SourceBeforeAmount applies equality check predicate on the "source_before_amount" field. It's identical to SourceBeforeAmountEQ.
func SourceBeforeAmount(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldSourceBeforeAmount, v))
}

// SourceAfterAmount applies equality check predicate on the "source_after_amount" field. It's identical to SourceAfterAmountEQ.
func SourceAfterAmount(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldSourceAfterAmount, v))
}

// SerialNumber applies equality check predicate on the "serial_number" field. It's identical to SerialNumberEQ.
func SerialNumber(v string) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldSerialNumber, v))
}

// InviteID applies equality check predicate on the "invite_id" field. It's identical to InviteIDEQ.
func InviteID(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldInviteID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldDeletedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v enums.BillType) predicate.Bill {
	vc := v
	return predicate.Bill(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v enums.BillType) predicate.Bill {
	vc := v
	return predicate.Bill(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...enums.BillType) predicate.Bill {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bill(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...enums.BillType) predicate.Bill {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bill(sql.FieldNotIn(FieldType, v...))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldOrderID, vs...))
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.Bill {
	return predicate.Bill(sql.FieldIsNull(FieldOrderID))
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.Bill {
	return predicate.Bill(sql.FieldNotNull(FieldOrderID))
}

// WayEQ applies the EQ predicate on the "way" field.
func WayEQ(v enums.BillWay) predicate.Bill {
	vc := v
	return predicate.Bill(sql.FieldEQ(FieldWay, vc))
}

// WayNEQ applies the NEQ predicate on the "way" field.
func WayNEQ(v enums.BillWay) predicate.Bill {
	vc := v
	return predicate.Bill(sql.FieldNEQ(FieldWay, vc))
}

// WayIn applies the In predicate on the "way" field.
func WayIn(vs ...enums.BillWay) predicate.Bill {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bill(sql.FieldIn(FieldWay, v...))
}

// WayNotIn applies the NotIn predicate on the "way" field.
func WayNotIn(vs ...enums.BillWay) predicate.Bill {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bill(sql.FieldNotIn(FieldWay, v...))
}

// SymbolIDEQ applies the EQ predicate on the "symbol_id" field.
func SymbolIDEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldSymbolID, v))
}

// SymbolIDNEQ applies the NEQ predicate on the "symbol_id" field.
func SymbolIDNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldSymbolID, v))
}

// SymbolIDIn applies the In predicate on the "symbol_id" field.
func SymbolIDIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldSymbolID, vs...))
}

// SymbolIDNotIn applies the NotIn predicate on the "symbol_id" field.
func SymbolIDNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldSymbolID, vs...))
}

// SymbolIDGT applies the GT predicate on the "symbol_id" field.
func SymbolIDGT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldSymbolID, v))
}

// SymbolIDGTE applies the GTE predicate on the "symbol_id" field.
func SymbolIDGTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldSymbolID, v))
}

// SymbolIDLT applies the LT predicate on the "symbol_id" field.
func SymbolIDLT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldSymbolID, v))
}

// SymbolIDLTE applies the LTE predicate on the "symbol_id" field.
func SymbolIDLTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldSymbolID, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldAmount, v))
}

// TargetUserIDEQ applies the EQ predicate on the "target_user_id" field.
func TargetUserIDEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldTargetUserID, v))
}

// TargetUserIDNEQ applies the NEQ predicate on the "target_user_id" field.
func TargetUserIDNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldTargetUserID, v))
}

// TargetUserIDIn applies the In predicate on the "target_user_id" field.
func TargetUserIDIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldTargetUserID, vs...))
}

// TargetUserIDNotIn applies the NotIn predicate on the "target_user_id" field.
func TargetUserIDNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldTargetUserID, vs...))
}

// TargetBeforeAmountEQ applies the EQ predicate on the "target_before_amount" field.
func TargetBeforeAmountEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldTargetBeforeAmount, v))
}

// TargetBeforeAmountNEQ applies the NEQ predicate on the "target_before_amount" field.
func TargetBeforeAmountNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldTargetBeforeAmount, v))
}

// TargetBeforeAmountIn applies the In predicate on the "target_before_amount" field.
func TargetBeforeAmountIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldTargetBeforeAmount, vs...))
}

// TargetBeforeAmountNotIn applies the NotIn predicate on the "target_before_amount" field.
func TargetBeforeAmountNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldTargetBeforeAmount, vs...))
}

// TargetBeforeAmountGT applies the GT predicate on the "target_before_amount" field.
func TargetBeforeAmountGT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldTargetBeforeAmount, v))
}

// TargetBeforeAmountGTE applies the GTE predicate on the "target_before_amount" field.
func TargetBeforeAmountGTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldTargetBeforeAmount, v))
}

// TargetBeforeAmountLT applies the LT predicate on the "target_before_amount" field.
func TargetBeforeAmountLT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldTargetBeforeAmount, v))
}

// TargetBeforeAmountLTE applies the LTE predicate on the "target_before_amount" field.
func TargetBeforeAmountLTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldTargetBeforeAmount, v))
}

// TargetAfterAmountEQ applies the EQ predicate on the "target_after_amount" field.
func TargetAfterAmountEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldTargetAfterAmount, v))
}

// TargetAfterAmountNEQ applies the NEQ predicate on the "target_after_amount" field.
func TargetAfterAmountNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldTargetAfterAmount, v))
}

// TargetAfterAmountIn applies the In predicate on the "target_after_amount" field.
func TargetAfterAmountIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldTargetAfterAmount, vs...))
}

// TargetAfterAmountNotIn applies the NotIn predicate on the "target_after_amount" field.
func TargetAfterAmountNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldTargetAfterAmount, vs...))
}

// TargetAfterAmountGT applies the GT predicate on the "target_after_amount" field.
func TargetAfterAmountGT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldTargetAfterAmount, v))
}

// TargetAfterAmountGTE applies the GTE predicate on the "target_after_amount" field.
func TargetAfterAmountGTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldTargetAfterAmount, v))
}

// TargetAfterAmountLT applies the LT predicate on the "target_after_amount" field.
func TargetAfterAmountLT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldTargetAfterAmount, v))
}

// TargetAfterAmountLTE applies the LTE predicate on the "target_after_amount" field.
func TargetAfterAmountLTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldTargetAfterAmount, v))
}

// SourceUserIDEQ applies the EQ predicate on the "source_user_id" field.
func SourceUserIDEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldSourceUserID, v))
}

// SourceUserIDNEQ applies the NEQ predicate on the "source_user_id" field.
func SourceUserIDNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldSourceUserID, v))
}

// SourceUserIDIn applies the In predicate on the "source_user_id" field.
func SourceUserIDIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldSourceUserID, vs...))
}

// SourceUserIDNotIn applies the NotIn predicate on the "source_user_id" field.
func SourceUserIDNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldSourceUserID, vs...))
}

// SourceBeforeAmountEQ applies the EQ predicate on the "source_before_amount" field.
func SourceBeforeAmountEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldSourceBeforeAmount, v))
}

// SourceBeforeAmountNEQ applies the NEQ predicate on the "source_before_amount" field.
func SourceBeforeAmountNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldSourceBeforeAmount, v))
}

// SourceBeforeAmountIn applies the In predicate on the "source_before_amount" field.
func SourceBeforeAmountIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldSourceBeforeAmount, vs...))
}

// SourceBeforeAmountNotIn applies the NotIn predicate on the "source_before_amount" field.
func SourceBeforeAmountNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldSourceBeforeAmount, vs...))
}

// SourceBeforeAmountGT applies the GT predicate on the "source_before_amount" field.
func SourceBeforeAmountGT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldSourceBeforeAmount, v))
}

// SourceBeforeAmountGTE applies the GTE predicate on the "source_before_amount" field.
func SourceBeforeAmountGTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldSourceBeforeAmount, v))
}

// SourceBeforeAmountLT applies the LT predicate on the "source_before_amount" field.
func SourceBeforeAmountLT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldSourceBeforeAmount, v))
}

// SourceBeforeAmountLTE applies the LTE predicate on the "source_before_amount" field.
func SourceBeforeAmountLTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldSourceBeforeAmount, v))
}

// SourceAfterAmountEQ applies the EQ predicate on the "source_after_amount" field.
func SourceAfterAmountEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldSourceAfterAmount, v))
}

// SourceAfterAmountNEQ applies the NEQ predicate on the "source_after_amount" field.
func SourceAfterAmountNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldSourceAfterAmount, v))
}

// SourceAfterAmountIn applies the In predicate on the "source_after_amount" field.
func SourceAfterAmountIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldSourceAfterAmount, vs...))
}

// SourceAfterAmountNotIn applies the NotIn predicate on the "source_after_amount" field.
func SourceAfterAmountNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldSourceAfterAmount, vs...))
}

// SourceAfterAmountGT applies the GT predicate on the "source_after_amount" field.
func SourceAfterAmountGT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldSourceAfterAmount, v))
}

// SourceAfterAmountGTE applies the GTE predicate on the "source_after_amount" field.
func SourceAfterAmountGTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldSourceAfterAmount, v))
}

// SourceAfterAmountLT applies the LT predicate on the "source_after_amount" field.
func SourceAfterAmountLT(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldSourceAfterAmount, v))
}

// SourceAfterAmountLTE applies the LTE predicate on the "source_after_amount" field.
func SourceAfterAmountLTE(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldSourceAfterAmount, v))
}

// SerialNumberEQ applies the EQ predicate on the "serial_number" field.
func SerialNumberEQ(v string) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldSerialNumber, v))
}

// SerialNumberNEQ applies the NEQ predicate on the "serial_number" field.
func SerialNumberNEQ(v string) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldSerialNumber, v))
}

// SerialNumberIn applies the In predicate on the "serial_number" field.
func SerialNumberIn(vs ...string) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldSerialNumber, vs...))
}

// SerialNumberNotIn applies the NotIn predicate on the "serial_number" field.
func SerialNumberNotIn(vs ...string) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldSerialNumber, vs...))
}

// SerialNumberGT applies the GT predicate on the "serial_number" field.
func SerialNumberGT(v string) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldSerialNumber, v))
}

// SerialNumberGTE applies the GTE predicate on the "serial_number" field.
func SerialNumberGTE(v string) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldSerialNumber, v))
}

// SerialNumberLT applies the LT predicate on the "serial_number" field.
func SerialNumberLT(v string) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldSerialNumber, v))
}

// SerialNumberLTE applies the LTE predicate on the "serial_number" field.
func SerialNumberLTE(v string) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldSerialNumber, v))
}

// SerialNumberContains applies the Contains predicate on the "serial_number" field.
func SerialNumberContains(v string) predicate.Bill {
	return predicate.Bill(sql.FieldContains(FieldSerialNumber, v))
}

// SerialNumberHasPrefix applies the HasPrefix predicate on the "serial_number" field.
func SerialNumberHasPrefix(v string) predicate.Bill {
	return predicate.Bill(sql.FieldHasPrefix(FieldSerialNumber, v))
}

// SerialNumberHasSuffix applies the HasSuffix predicate on the "serial_number" field.
func SerialNumberHasSuffix(v string) predicate.Bill {
	return predicate.Bill(sql.FieldHasSuffix(FieldSerialNumber, v))
}

// SerialNumberEqualFold applies the EqualFold predicate on the "serial_number" field.
func SerialNumberEqualFold(v string) predicate.Bill {
	return predicate.Bill(sql.FieldEqualFold(FieldSerialNumber, v))
}

// SerialNumberContainsFold applies the ContainsFold predicate on the "serial_number" field.
func SerialNumberContainsFold(v string) predicate.Bill {
	return predicate.Bill(sql.FieldContainsFold(FieldSerialNumber, v))
}

// InviteIDEQ applies the EQ predicate on the "invite_id" field.
func InviteIDEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldInviteID, v))
}

// InviteIDNEQ applies the NEQ predicate on the "invite_id" field.
func InviteIDNEQ(v int64) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldInviteID, v))
}

// InviteIDIn applies the In predicate on the "invite_id" field.
func InviteIDIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldInviteID, vs...))
}

// InviteIDNotIn applies the NotIn predicate on the "invite_id" field.
func InviteIDNotIn(vs ...int64) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldInviteID, vs...))
}

// HasSourceUser applies the HasEdge predicate on the "source_user" edge.
func HasSourceUser() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SourceUserTable, SourceUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSourceUserWith applies the HasEdge predicate on the "source_user" edge with a given conditions (other predicates).
func HasSourceUserWith(preds ...predicate.User) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := newSourceUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTargetUser applies the HasEdge predicate on the "target_user" edge.
func HasTargetUser() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TargetUserTable, TargetUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTargetUserWith applies the HasEdge predicate on the "target_user" edge with a given conditions (other predicates).
func HasTargetUserWith(preds ...predicate.User) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := newTargetUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTransferOrder applies the HasEdge predicate on the "transfer_order" edge.
func HasTransferOrder() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TransferOrderTable, TransferOrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransferOrderWith applies the HasEdge predicate on the "transfer_order" edge with a given conditions (other predicates).
func HasTransferOrderWith(preds ...predicate.TransferOrder) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := newTransferOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMissionConsumeOrder applies the HasEdge predicate on the "mission_consume_order" edge.
func HasMissionConsumeOrder() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MissionConsumeOrderTable, MissionConsumeOrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionConsumeOrderWith applies the HasEdge predicate on the "mission_consume_order" edge with a given conditions (other predicates).
func HasMissionConsumeOrderWith(preds ...predicate.MissionConsumeOrder) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := newMissionConsumeOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMissionProduceOrder applies the HasEdge predicate on the "mission_produce_order" edge.
func HasMissionProduceOrder() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MissionProduceOrderTable, MissionProduceOrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionProduceOrderWith applies the HasEdge predicate on the "mission_produce_order" edge with a given conditions (other predicates).
func HasMissionProduceOrderWith(preds ...predicate.MissionProduceOrder) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := newMissionProduceOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInvite applies the HasEdge predicate on the "invite" edge.
func HasInvite() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InviteTable, InviteColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInviteWith applies the HasEdge predicate on the "invite" edge with a given conditions (other predicates).
func HasInviteWith(preds ...predicate.Invite) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := newInviteStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Bill) predicate.Bill {
	return predicate.Bill(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Bill) predicate.Bill {
	return predicate.Bill(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Bill) predicate.Bill {
	return predicate.Bill(sql.NotPredicates(p))
}
