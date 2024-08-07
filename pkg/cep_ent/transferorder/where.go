// Code generated by ent, DO NOT EDIT.

package transferorder

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldDeletedAt, v))
}

// SourceUserID applies equality check predicate on the "source_user_id" field. It's identical to SourceUserIDEQ.
func SourceUserID(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldSourceUserID, v))
}

// TargetUserID applies equality check predicate on the "target_user_id" field. It's identical to TargetUserIDEQ.
func TargetUserID(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldTargetUserID, v))
}

// SymbolID applies equality check predicate on the "symbol_id" field. It's identical to SymbolIDEQ.
func SymbolID(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldSymbolID, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldAmount, v))
}

// SerialNumber applies equality check predicate on the "serial_number" field. It's identical to SerialNumberEQ.
func SerialNumber(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldSerialNumber, v))
}

// SocialID applies equality check predicate on the "social_id" field. It's identical to SocialIDEQ.
func SocialID(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldSocialID, v))
}

// ThirdAPIResp applies equality check predicate on the "third_api_resp" field. It's identical to ThirdAPIRespEQ.
func ThirdAPIResp(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldThirdAPIResp, v))
}

// OutTransactionID applies equality check predicate on the "out_transaction_id" field. It's identical to OutTransactionIDEQ.
func OutTransactionID(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldOutTransactionID, v))
}

// OperateUserID applies equality check predicate on the "operate_user_id" field. It's identical to OperateUserIDEQ.
func OperateUserID(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldOperateUserID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLTE(FieldDeletedAt, v))
}

// SourceUserIDEQ applies the EQ predicate on the "source_user_id" field.
func SourceUserIDEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldSourceUserID, v))
}

// SourceUserIDNEQ applies the NEQ predicate on the "source_user_id" field.
func SourceUserIDNEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldSourceUserID, v))
}

// SourceUserIDIn applies the In predicate on the "source_user_id" field.
func SourceUserIDIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldSourceUserID, vs...))
}

// SourceUserIDNotIn applies the NotIn predicate on the "source_user_id" field.
func SourceUserIDNotIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldSourceUserID, vs...))
}

// TargetUserIDEQ applies the EQ predicate on the "target_user_id" field.
func TargetUserIDEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldTargetUserID, v))
}

// TargetUserIDNEQ applies the NEQ predicate on the "target_user_id" field.
func TargetUserIDNEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldTargetUserID, v))
}

// TargetUserIDIn applies the In predicate on the "target_user_id" field.
func TargetUserIDIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldTargetUserID, vs...))
}

// TargetUserIDNotIn applies the NotIn predicate on the "target_user_id" field.
func TargetUserIDNotIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldTargetUserID, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldStatus, vs...))
}

// SymbolIDEQ applies the EQ predicate on the "symbol_id" field.
func SymbolIDEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldSymbolID, v))
}

// SymbolIDNEQ applies the NEQ predicate on the "symbol_id" field.
func SymbolIDNEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldSymbolID, v))
}

// SymbolIDIn applies the In predicate on the "symbol_id" field.
func SymbolIDIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldSymbolID, vs...))
}

// SymbolIDNotIn applies the NotIn predicate on the "symbol_id" field.
func SymbolIDNotIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldSymbolID, vs...))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLTE(FieldAmount, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v enums.TransferOrderType) predicate.TransferOrder {
	vc := v
	return predicate.TransferOrder(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v enums.TransferOrderType) predicate.TransferOrder {
	vc := v
	return predicate.TransferOrder(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...enums.TransferOrderType) predicate.TransferOrder {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransferOrder(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...enums.TransferOrderType) predicate.TransferOrder {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransferOrder(sql.FieldNotIn(FieldType, v...))
}

// SerialNumberEQ applies the EQ predicate on the "serial_number" field.
func SerialNumberEQ(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldSerialNumber, v))
}

// SerialNumberNEQ applies the NEQ predicate on the "serial_number" field.
func SerialNumberNEQ(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldSerialNumber, v))
}

// SerialNumberIn applies the In predicate on the "serial_number" field.
func SerialNumberIn(vs ...string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldSerialNumber, vs...))
}

// SerialNumberNotIn applies the NotIn predicate on the "serial_number" field.
func SerialNumberNotIn(vs ...string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldSerialNumber, vs...))
}

// SerialNumberGT applies the GT predicate on the "serial_number" field.
func SerialNumberGT(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGT(FieldSerialNumber, v))
}

// SerialNumberGTE applies the GTE predicate on the "serial_number" field.
func SerialNumberGTE(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGTE(FieldSerialNumber, v))
}

// SerialNumberLT applies the LT predicate on the "serial_number" field.
func SerialNumberLT(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLT(FieldSerialNumber, v))
}

// SerialNumberLTE applies the LTE predicate on the "serial_number" field.
func SerialNumberLTE(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLTE(FieldSerialNumber, v))
}

// SerialNumberContains applies the Contains predicate on the "serial_number" field.
func SerialNumberContains(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldContains(FieldSerialNumber, v))
}

// SerialNumberHasPrefix applies the HasPrefix predicate on the "serial_number" field.
func SerialNumberHasPrefix(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldHasPrefix(FieldSerialNumber, v))
}

// SerialNumberHasSuffix applies the HasSuffix predicate on the "serial_number" field.
func SerialNumberHasSuffix(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldHasSuffix(FieldSerialNumber, v))
}

// SerialNumberEqualFold applies the EqualFold predicate on the "serial_number" field.
func SerialNumberEqualFold(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEqualFold(FieldSerialNumber, v))
}

// SerialNumberContainsFold applies the ContainsFold predicate on the "serial_number" field.
func SerialNumberContainsFold(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldContainsFold(FieldSerialNumber, v))
}

// SocialIDEQ applies the EQ predicate on the "social_id" field.
func SocialIDEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldSocialID, v))
}

// SocialIDNEQ applies the NEQ predicate on the "social_id" field.
func SocialIDNEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldSocialID, v))
}

// SocialIDIn applies the In predicate on the "social_id" field.
func SocialIDIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldSocialID, vs...))
}

// SocialIDNotIn applies the NotIn predicate on the "social_id" field.
func SocialIDNotIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldSocialID, vs...))
}

// SocialIDIsNil applies the IsNil predicate on the "social_id" field.
func SocialIDIsNil() predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIsNull(FieldSocialID))
}

// SocialIDNotNil applies the NotNil predicate on the "social_id" field.
func SocialIDNotNil() predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotNull(FieldSocialID))
}

// ThirdAPIRespEQ applies the EQ predicate on the "third_api_resp" field.
func ThirdAPIRespEQ(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldThirdAPIResp, v))
}

// ThirdAPIRespNEQ applies the NEQ predicate on the "third_api_resp" field.
func ThirdAPIRespNEQ(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldThirdAPIResp, v))
}

// ThirdAPIRespIn applies the In predicate on the "third_api_resp" field.
func ThirdAPIRespIn(vs ...string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldThirdAPIResp, vs...))
}

// ThirdAPIRespNotIn applies the NotIn predicate on the "third_api_resp" field.
func ThirdAPIRespNotIn(vs ...string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldThirdAPIResp, vs...))
}

// ThirdAPIRespGT applies the GT predicate on the "third_api_resp" field.
func ThirdAPIRespGT(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGT(FieldThirdAPIResp, v))
}

// ThirdAPIRespGTE applies the GTE predicate on the "third_api_resp" field.
func ThirdAPIRespGTE(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGTE(FieldThirdAPIResp, v))
}

// ThirdAPIRespLT applies the LT predicate on the "third_api_resp" field.
func ThirdAPIRespLT(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLT(FieldThirdAPIResp, v))
}

// ThirdAPIRespLTE applies the LTE predicate on the "third_api_resp" field.
func ThirdAPIRespLTE(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLTE(FieldThirdAPIResp, v))
}

// ThirdAPIRespContains applies the Contains predicate on the "third_api_resp" field.
func ThirdAPIRespContains(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldContains(FieldThirdAPIResp, v))
}

// ThirdAPIRespHasPrefix applies the HasPrefix predicate on the "third_api_resp" field.
func ThirdAPIRespHasPrefix(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldHasPrefix(FieldThirdAPIResp, v))
}

// ThirdAPIRespHasSuffix applies the HasSuffix predicate on the "third_api_resp" field.
func ThirdAPIRespHasSuffix(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldHasSuffix(FieldThirdAPIResp, v))
}

// ThirdAPIRespEqualFold applies the EqualFold predicate on the "third_api_resp" field.
func ThirdAPIRespEqualFold(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEqualFold(FieldThirdAPIResp, v))
}

// ThirdAPIRespContainsFold applies the ContainsFold predicate on the "third_api_resp" field.
func ThirdAPIRespContainsFold(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldContainsFold(FieldThirdAPIResp, v))
}

// OutTransactionIDEQ applies the EQ predicate on the "out_transaction_id" field.
func OutTransactionIDEQ(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldOutTransactionID, v))
}

// OutTransactionIDNEQ applies the NEQ predicate on the "out_transaction_id" field.
func OutTransactionIDNEQ(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldOutTransactionID, v))
}

// OutTransactionIDIn applies the In predicate on the "out_transaction_id" field.
func OutTransactionIDIn(vs ...string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldOutTransactionID, vs...))
}

// OutTransactionIDNotIn applies the NotIn predicate on the "out_transaction_id" field.
func OutTransactionIDNotIn(vs ...string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldOutTransactionID, vs...))
}

// OutTransactionIDGT applies the GT predicate on the "out_transaction_id" field.
func OutTransactionIDGT(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGT(FieldOutTransactionID, v))
}

// OutTransactionIDGTE applies the GTE predicate on the "out_transaction_id" field.
func OutTransactionIDGTE(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGTE(FieldOutTransactionID, v))
}

// OutTransactionIDLT applies the LT predicate on the "out_transaction_id" field.
func OutTransactionIDLT(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLT(FieldOutTransactionID, v))
}

// OutTransactionIDLTE applies the LTE predicate on the "out_transaction_id" field.
func OutTransactionIDLTE(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLTE(FieldOutTransactionID, v))
}

// OutTransactionIDContains applies the Contains predicate on the "out_transaction_id" field.
func OutTransactionIDContains(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldContains(FieldOutTransactionID, v))
}

// OutTransactionIDHasPrefix applies the HasPrefix predicate on the "out_transaction_id" field.
func OutTransactionIDHasPrefix(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldHasPrefix(FieldOutTransactionID, v))
}

// OutTransactionIDHasSuffix applies the HasSuffix predicate on the "out_transaction_id" field.
func OutTransactionIDHasSuffix(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldHasSuffix(FieldOutTransactionID, v))
}

// OutTransactionIDEqualFold applies the EqualFold predicate on the "out_transaction_id" field.
func OutTransactionIDEqualFold(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEqualFold(FieldOutTransactionID, v))
}

// OutTransactionIDContainsFold applies the ContainsFold predicate on the "out_transaction_id" field.
func OutTransactionIDContainsFold(v string) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldContainsFold(FieldOutTransactionID, v))
}

// OperateUserIDEQ applies the EQ predicate on the "operate_user_id" field.
func OperateUserIDEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldEQ(FieldOperateUserID, v))
}

// OperateUserIDNEQ applies the NEQ predicate on the "operate_user_id" field.
func OperateUserIDNEQ(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNEQ(FieldOperateUserID, v))
}

// OperateUserIDIn applies the In predicate on the "operate_user_id" field.
func OperateUserIDIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldIn(FieldOperateUserID, vs...))
}

// OperateUserIDNotIn applies the NotIn predicate on the "operate_user_id" field.
func OperateUserIDNotIn(vs ...int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldNotIn(FieldOperateUserID, vs...))
}

// OperateUserIDGT applies the GT predicate on the "operate_user_id" field.
func OperateUserIDGT(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGT(FieldOperateUserID, v))
}

// OperateUserIDGTE applies the GTE predicate on the "operate_user_id" field.
func OperateUserIDGTE(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldGTE(FieldOperateUserID, v))
}

// OperateUserIDLT applies the LT predicate on the "operate_user_id" field.
func OperateUserIDLT(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLT(FieldOperateUserID, v))
}

// OperateUserIDLTE applies the LTE predicate on the "operate_user_id" field.
func OperateUserIDLTE(v int64) predicate.TransferOrder {
	return predicate.TransferOrder(sql.FieldLTE(FieldOperateUserID, v))
}

// GiftTypeEQ applies the EQ predicate on the "gift_type" field.
func GiftTypeEQ(v enums.TransferOrderGiftType) predicate.TransferOrder {
	vc := v
	return predicate.TransferOrder(sql.FieldEQ(FieldGiftType, vc))
}

// GiftTypeNEQ applies the NEQ predicate on the "gift_type" field.
func GiftTypeNEQ(v enums.TransferOrderGiftType) predicate.TransferOrder {
	vc := v
	return predicate.TransferOrder(sql.FieldNEQ(FieldGiftType, vc))
}

// GiftTypeIn applies the In predicate on the "gift_type" field.
func GiftTypeIn(vs ...enums.TransferOrderGiftType) predicate.TransferOrder {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransferOrder(sql.FieldIn(FieldGiftType, v...))
}

// GiftTypeNotIn applies the NotIn predicate on the "gift_type" field.
func GiftTypeNotIn(vs ...enums.TransferOrderGiftType) predicate.TransferOrder {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransferOrder(sql.FieldNotIn(FieldGiftType, v...))
}

// HasSourceUser applies the HasEdge predicate on the "source_user" edge.
func HasSourceUser() predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SourceUserTable, SourceUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSourceUserWith applies the HasEdge predicate on the "source_user" edge with a given conditions (other predicates).
func HasSourceUserWith(preds ...predicate.User) predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := newSourceUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTargetUser applies the HasEdge predicate on the "target_user" edge.
func HasTargetUser() predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TargetUserTable, TargetUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTargetUserWith applies the HasEdge predicate on the "target_user" edge with a given conditions (other predicates).
func HasTargetUserWith(preds ...predicate.User) predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := newTargetUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBills applies the HasEdge predicate on the "bills" edge.
func HasBills() predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BillsTable, BillsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillsWith applies the HasEdge predicate on the "bills" edge with a given conditions (other predicates).
func HasBillsWith(preds ...predicate.Bill) predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := newBillsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVxSocial applies the HasEdge predicate on the "vx_social" edge.
func HasVxSocial() predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VxSocialTable, VxSocialColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVxSocialWith applies the HasEdge predicate on the "vx_social" edge with a given conditions (other predicates).
func HasVxSocialWith(preds ...predicate.VXSocial) predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := newVxSocialStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSymbol applies the HasEdge predicate on the "symbol" edge.
func HasSymbol() predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SymbolTable, SymbolColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSymbolWith applies the HasEdge predicate on the "symbol" edge with a given conditions (other predicates).
func HasSymbolWith(preds ...predicate.Symbol) predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := newSymbolStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWithdrawRecord applies the HasEdge predicate on the "withdraw_record" edge.
func HasWithdrawRecord() predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, WithdrawRecordTable, WithdrawRecordColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWithdrawRecordWith applies the HasEdge predicate on the "withdraw_record" edge with a given conditions (other predicates).
func HasWithdrawRecordWith(preds ...predicate.WithdrawRecord) predicate.TransferOrder {
	return predicate.TransferOrder(func(s *sql.Selector) {
		step := newWithdrawRecordStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TransferOrder) predicate.TransferOrder {
	return predicate.TransferOrder(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TransferOrder) predicate.TransferOrder {
	return predicate.TransferOrder(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TransferOrder) predicate.TransferOrder {
	return predicate.TransferOrder(sql.NotPredicates(p))
}
