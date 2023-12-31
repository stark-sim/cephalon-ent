// Code generated by ent, DO NOT EDIT.

package rechargeorder

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldUserID, v))
}

// PureCep applies equality check predicate on the "pure_cep" field. It's identical to PureCepEQ.
func PureCep(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldPureCep, v))
}

// GiftCep applies equality check predicate on the "gift_cep" field. It's identical to GiftCepEQ.
func GiftCep(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldGiftCep, v))
}

// SocialID applies equality check predicate on the "social_id" field. It's identical to SocialIDEQ.
func SocialID(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldSocialID, v))
}

// SerialNumber applies equality check predicate on the "serial_number" field. It's identical to SerialNumberEQ.
func SerialNumber(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldSerialNumber, v))
}

// ThirdAPIResp applies equality check predicate on the "third_api_resp" field. It's identical to ThirdAPIRespEQ.
func ThirdAPIResp(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldThirdAPIResp, v))
}

// FromUserID applies equality check predicate on the "from_user_id" field. It's identical to FromUserIDEQ.
func FromUserID(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldFromUserID, v))
}

// OutTransactionID applies equality check predicate on the "out_transaction_id" field. It's identical to OutTransactionIDEQ.
func OutTransactionID(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldOutTransactionID, v))
}

// CampaignOrderID applies equality check predicate on the "campaign_order_id" field. It's identical to CampaignOrderIDEQ.
func CampaignOrderID(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldCampaignOrderID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldDeletedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldUserID, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldStatus, vs...))
}

// PureCepEQ applies the EQ predicate on the "pure_cep" field.
func PureCepEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldPureCep, v))
}

// PureCepNEQ applies the NEQ predicate on the "pure_cep" field.
func PureCepNEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldPureCep, v))
}

// PureCepIn applies the In predicate on the "pure_cep" field.
func PureCepIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldPureCep, vs...))
}

// PureCepNotIn applies the NotIn predicate on the "pure_cep" field.
func PureCepNotIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldPureCep, vs...))
}

// PureCepGT applies the GT predicate on the "pure_cep" field.
func PureCepGT(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldPureCep, v))
}

// PureCepGTE applies the GTE predicate on the "pure_cep" field.
func PureCepGTE(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldPureCep, v))
}

// PureCepLT applies the LT predicate on the "pure_cep" field.
func PureCepLT(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldPureCep, v))
}

// PureCepLTE applies the LTE predicate on the "pure_cep" field.
func PureCepLTE(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldPureCep, v))
}

// GiftCepEQ applies the EQ predicate on the "gift_cep" field.
func GiftCepEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldGiftCep, v))
}

// GiftCepNEQ applies the NEQ predicate on the "gift_cep" field.
func GiftCepNEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldGiftCep, v))
}

// GiftCepIn applies the In predicate on the "gift_cep" field.
func GiftCepIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldGiftCep, vs...))
}

// GiftCepNotIn applies the NotIn predicate on the "gift_cep" field.
func GiftCepNotIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldGiftCep, vs...))
}

// GiftCepGT applies the GT predicate on the "gift_cep" field.
func GiftCepGT(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldGiftCep, v))
}

// GiftCepGTE applies the GTE predicate on the "gift_cep" field.
func GiftCepGTE(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldGiftCep, v))
}

// GiftCepLT applies the LT predicate on the "gift_cep" field.
func GiftCepLT(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldGiftCep, v))
}

// GiftCepLTE applies the LTE predicate on the "gift_cep" field.
func GiftCepLTE(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldGiftCep, v))
}

// SocialIDEQ applies the EQ predicate on the "social_id" field.
func SocialIDEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldSocialID, v))
}

// SocialIDNEQ applies the NEQ predicate on the "social_id" field.
func SocialIDNEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldSocialID, v))
}

// SocialIDIn applies the In predicate on the "social_id" field.
func SocialIDIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldSocialID, vs...))
}

// SocialIDNotIn applies the NotIn predicate on the "social_id" field.
func SocialIDNotIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldSocialID, vs...))
}

// SocialIDIsNil applies the IsNil predicate on the "social_id" field.
func SocialIDIsNil() predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIsNull(FieldSocialID))
}

// SocialIDNotNil applies the NotNil predicate on the "social_id" field.
func SocialIDNotNil() predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotNull(FieldSocialID))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldType, vs...))
}

// SerialNumberEQ applies the EQ predicate on the "serial_number" field.
func SerialNumberEQ(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldSerialNumber, v))
}

// SerialNumberNEQ applies the NEQ predicate on the "serial_number" field.
func SerialNumberNEQ(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldSerialNumber, v))
}

// SerialNumberIn applies the In predicate on the "serial_number" field.
func SerialNumberIn(vs ...string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldSerialNumber, vs...))
}

// SerialNumberNotIn applies the NotIn predicate on the "serial_number" field.
func SerialNumberNotIn(vs ...string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldSerialNumber, vs...))
}

// SerialNumberGT applies the GT predicate on the "serial_number" field.
func SerialNumberGT(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldSerialNumber, v))
}

// SerialNumberGTE applies the GTE predicate on the "serial_number" field.
func SerialNumberGTE(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldSerialNumber, v))
}

// SerialNumberLT applies the LT predicate on the "serial_number" field.
func SerialNumberLT(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldSerialNumber, v))
}

// SerialNumberLTE applies the LTE predicate on the "serial_number" field.
func SerialNumberLTE(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldSerialNumber, v))
}

// SerialNumberContains applies the Contains predicate on the "serial_number" field.
func SerialNumberContains(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldContains(FieldSerialNumber, v))
}

// SerialNumberHasPrefix applies the HasPrefix predicate on the "serial_number" field.
func SerialNumberHasPrefix(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldHasPrefix(FieldSerialNumber, v))
}

// SerialNumberHasSuffix applies the HasSuffix predicate on the "serial_number" field.
func SerialNumberHasSuffix(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldHasSuffix(FieldSerialNumber, v))
}

// SerialNumberEqualFold applies the EqualFold predicate on the "serial_number" field.
func SerialNumberEqualFold(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEqualFold(FieldSerialNumber, v))
}

// SerialNumberContainsFold applies the ContainsFold predicate on the "serial_number" field.
func SerialNumberContainsFold(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldContainsFold(FieldSerialNumber, v))
}

// ThirdAPIRespEQ applies the EQ predicate on the "third_api_resp" field.
func ThirdAPIRespEQ(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldThirdAPIResp, v))
}

// ThirdAPIRespNEQ applies the NEQ predicate on the "third_api_resp" field.
func ThirdAPIRespNEQ(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldThirdAPIResp, v))
}

// ThirdAPIRespIn applies the In predicate on the "third_api_resp" field.
func ThirdAPIRespIn(vs ...string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldThirdAPIResp, vs...))
}

// ThirdAPIRespNotIn applies the NotIn predicate on the "third_api_resp" field.
func ThirdAPIRespNotIn(vs ...string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldThirdAPIResp, vs...))
}

// ThirdAPIRespGT applies the GT predicate on the "third_api_resp" field.
func ThirdAPIRespGT(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldThirdAPIResp, v))
}

// ThirdAPIRespGTE applies the GTE predicate on the "third_api_resp" field.
func ThirdAPIRespGTE(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldThirdAPIResp, v))
}

// ThirdAPIRespLT applies the LT predicate on the "third_api_resp" field.
func ThirdAPIRespLT(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldThirdAPIResp, v))
}

// ThirdAPIRespLTE applies the LTE predicate on the "third_api_resp" field.
func ThirdAPIRespLTE(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldThirdAPIResp, v))
}

// ThirdAPIRespContains applies the Contains predicate on the "third_api_resp" field.
func ThirdAPIRespContains(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldContains(FieldThirdAPIResp, v))
}

// ThirdAPIRespHasPrefix applies the HasPrefix predicate on the "third_api_resp" field.
func ThirdAPIRespHasPrefix(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldHasPrefix(FieldThirdAPIResp, v))
}

// ThirdAPIRespHasSuffix applies the HasSuffix predicate on the "third_api_resp" field.
func ThirdAPIRespHasSuffix(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldHasSuffix(FieldThirdAPIResp, v))
}

// ThirdAPIRespEqualFold applies the EqualFold predicate on the "third_api_resp" field.
func ThirdAPIRespEqualFold(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEqualFold(FieldThirdAPIResp, v))
}

// ThirdAPIRespContainsFold applies the ContainsFold predicate on the "third_api_resp" field.
func ThirdAPIRespContainsFold(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldContainsFold(FieldThirdAPIResp, v))
}

// FromUserIDEQ applies the EQ predicate on the "from_user_id" field.
func FromUserIDEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldFromUserID, v))
}

// FromUserIDNEQ applies the NEQ predicate on the "from_user_id" field.
func FromUserIDNEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldFromUserID, v))
}

// FromUserIDIn applies the In predicate on the "from_user_id" field.
func FromUserIDIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldFromUserID, vs...))
}

// FromUserIDNotIn applies the NotIn predicate on the "from_user_id" field.
func FromUserIDNotIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldFromUserID, vs...))
}

// FromUserIDGT applies the GT predicate on the "from_user_id" field.
func FromUserIDGT(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldFromUserID, v))
}

// FromUserIDGTE applies the GTE predicate on the "from_user_id" field.
func FromUserIDGTE(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldFromUserID, v))
}

// FromUserIDLT applies the LT predicate on the "from_user_id" field.
func FromUserIDLT(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldFromUserID, v))
}

// FromUserIDLTE applies the LTE predicate on the "from_user_id" field.
func FromUserIDLTE(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldFromUserID, v))
}

// OutTransactionIDEQ applies the EQ predicate on the "out_transaction_id" field.
func OutTransactionIDEQ(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldOutTransactionID, v))
}

// OutTransactionIDNEQ applies the NEQ predicate on the "out_transaction_id" field.
func OutTransactionIDNEQ(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldOutTransactionID, v))
}

// OutTransactionIDIn applies the In predicate on the "out_transaction_id" field.
func OutTransactionIDIn(vs ...string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldOutTransactionID, vs...))
}

// OutTransactionIDNotIn applies the NotIn predicate on the "out_transaction_id" field.
func OutTransactionIDNotIn(vs ...string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldOutTransactionID, vs...))
}

// OutTransactionIDGT applies the GT predicate on the "out_transaction_id" field.
func OutTransactionIDGT(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGT(FieldOutTransactionID, v))
}

// OutTransactionIDGTE applies the GTE predicate on the "out_transaction_id" field.
func OutTransactionIDGTE(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldGTE(FieldOutTransactionID, v))
}

// OutTransactionIDLT applies the LT predicate on the "out_transaction_id" field.
func OutTransactionIDLT(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLT(FieldOutTransactionID, v))
}

// OutTransactionIDLTE applies the LTE predicate on the "out_transaction_id" field.
func OutTransactionIDLTE(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldLTE(FieldOutTransactionID, v))
}

// OutTransactionIDContains applies the Contains predicate on the "out_transaction_id" field.
func OutTransactionIDContains(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldContains(FieldOutTransactionID, v))
}

// OutTransactionIDHasPrefix applies the HasPrefix predicate on the "out_transaction_id" field.
func OutTransactionIDHasPrefix(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldHasPrefix(FieldOutTransactionID, v))
}

// OutTransactionIDHasSuffix applies the HasSuffix predicate on the "out_transaction_id" field.
func OutTransactionIDHasSuffix(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldHasSuffix(FieldOutTransactionID, v))
}

// OutTransactionIDEqualFold applies the EqualFold predicate on the "out_transaction_id" field.
func OutTransactionIDEqualFold(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEqualFold(FieldOutTransactionID, v))
}

// OutTransactionIDContainsFold applies the ContainsFold predicate on the "out_transaction_id" field.
func OutTransactionIDContainsFold(v string) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldContainsFold(FieldOutTransactionID, v))
}

// CampaignOrderIDEQ applies the EQ predicate on the "campaign_order_id" field.
func CampaignOrderIDEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldEQ(FieldCampaignOrderID, v))
}

// CampaignOrderIDNEQ applies the NEQ predicate on the "campaign_order_id" field.
func CampaignOrderIDNEQ(v int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNEQ(FieldCampaignOrderID, v))
}

// CampaignOrderIDIn applies the In predicate on the "campaign_order_id" field.
func CampaignOrderIDIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIn(FieldCampaignOrderID, vs...))
}

// CampaignOrderIDNotIn applies the NotIn predicate on the "campaign_order_id" field.
func CampaignOrderIDNotIn(vs ...int64) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotIn(FieldCampaignOrderID, vs...))
}

// CampaignOrderIDIsNil applies the IsNil predicate on the "campaign_order_id" field.
func CampaignOrderIDIsNil() predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldIsNull(FieldCampaignOrderID))
}

// CampaignOrderIDNotNil applies the NotNil predicate on the "campaign_order_id" field.
func CampaignOrderIDNotNil() predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.FieldNotNull(FieldCampaignOrderID))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.RechargeOrder {
	return predicate.RechargeOrder(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.RechargeOrder {
	return predicate.RechargeOrder(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCostBills applies the HasEdge predicate on the "cost_bills" edge.
func HasCostBills() predicate.RechargeOrder {
	return predicate.RechargeOrder(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CostBillsTable, CostBillsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCostBillsWith applies the HasEdge predicate on the "cost_bills" edge with a given conditions (other predicates).
func HasCostBillsWith(preds ...predicate.CostBill) predicate.RechargeOrder {
	return predicate.RechargeOrder(func(s *sql.Selector) {
		step := newCostBillsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVxSocial applies the HasEdge predicate on the "vx_social" edge.
func HasVxSocial() predicate.RechargeOrder {
	return predicate.RechargeOrder(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VxSocialTable, VxSocialColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVxSocialWith applies the HasEdge predicate on the "vx_social" edge with a given conditions (other predicates).
func HasVxSocialWith(preds ...predicate.VXSocial) predicate.RechargeOrder {
	return predicate.RechargeOrder(func(s *sql.Selector) {
		step := newVxSocialStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCampaignOrder applies the HasEdge predicate on the "campaign_order" edge.
func HasCampaignOrder() predicate.RechargeOrder {
	return predicate.RechargeOrder(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CampaignOrderTable, CampaignOrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCampaignOrderWith applies the HasEdge predicate on the "campaign_order" edge with a given conditions (other predicates).
func HasCampaignOrderWith(preds ...predicate.CampaignOrder) predicate.RechargeOrder {
	return predicate.RechargeOrder(func(s *sql.Selector) {
		step := newCampaignOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RechargeOrder) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RechargeOrder) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RechargeOrder) predicate.RechargeOrder {
	return predicate.RechargeOrder(sql.NotPredicates(p))
}
