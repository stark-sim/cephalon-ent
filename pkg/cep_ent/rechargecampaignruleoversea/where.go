// Code generated by ent, DO NOT EDIT.

package rechargecampaignruleoversea

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldDeletedAt, v))
}

// DollarPrice applies equality check predicate on the "dollar_price" field. It's identical to DollarPriceEQ.
func DollarPrice(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldDollarPrice, v))
}

// RmbPrice applies equality check predicate on the "rmb_price" field. It's identical to RmbPriceEQ.
func RmbPrice(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldRmbPrice, v))
}

// OriginalRmbPrice applies equality check predicate on the "original_rmb_price" field. It's identical to OriginalRmbPriceEQ.
func OriginalRmbPrice(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldOriginalRmbPrice, v))
}

// TotalCep applies equality check predicate on the "total_cep" field. It's identical to TotalCepEQ.
func TotalCep(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldTotalCep, v))
}

// BeforeDiscountCep applies equality check predicate on the "before_discount_cep" field. It's identical to BeforeDiscountCepEQ.
func BeforeDiscountCep(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldBeforeDiscountCep, v))
}

// DiscountRatio applies equality check predicate on the "discount_ratio" field. It's identical to DiscountRatioEQ.
func DiscountRatio(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldDiscountRatio, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldDeletedAt, v))
}

// DollarPriceEQ applies the EQ predicate on the "dollar_price" field.
func DollarPriceEQ(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldDollarPrice, v))
}

// DollarPriceNEQ applies the NEQ predicate on the "dollar_price" field.
func DollarPriceNEQ(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldDollarPrice, v))
}

// DollarPriceIn applies the In predicate on the "dollar_price" field.
func DollarPriceIn(vs ...string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldDollarPrice, vs...))
}

// DollarPriceNotIn applies the NotIn predicate on the "dollar_price" field.
func DollarPriceNotIn(vs ...string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldDollarPrice, vs...))
}

// DollarPriceGT applies the GT predicate on the "dollar_price" field.
func DollarPriceGT(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldDollarPrice, v))
}

// DollarPriceGTE applies the GTE predicate on the "dollar_price" field.
func DollarPriceGTE(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldDollarPrice, v))
}

// DollarPriceLT applies the LT predicate on the "dollar_price" field.
func DollarPriceLT(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldDollarPrice, v))
}

// DollarPriceLTE applies the LTE predicate on the "dollar_price" field.
func DollarPriceLTE(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldDollarPrice, v))
}

// DollarPriceContains applies the Contains predicate on the "dollar_price" field.
func DollarPriceContains(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldContains(FieldDollarPrice, v))
}

// DollarPriceHasPrefix applies the HasPrefix predicate on the "dollar_price" field.
func DollarPriceHasPrefix(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldHasPrefix(FieldDollarPrice, v))
}

// DollarPriceHasSuffix applies the HasSuffix predicate on the "dollar_price" field.
func DollarPriceHasSuffix(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldHasSuffix(FieldDollarPrice, v))
}

// DollarPriceEqualFold applies the EqualFold predicate on the "dollar_price" field.
func DollarPriceEqualFold(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEqualFold(FieldDollarPrice, v))
}

// DollarPriceContainsFold applies the ContainsFold predicate on the "dollar_price" field.
func DollarPriceContainsFold(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldContainsFold(FieldDollarPrice, v))
}

// RmbPriceEQ applies the EQ predicate on the "rmb_price" field.
func RmbPriceEQ(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldRmbPrice, v))
}

// RmbPriceNEQ applies the NEQ predicate on the "rmb_price" field.
func RmbPriceNEQ(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldRmbPrice, v))
}

// RmbPriceIn applies the In predicate on the "rmb_price" field.
func RmbPriceIn(vs ...string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldRmbPrice, vs...))
}

// RmbPriceNotIn applies the NotIn predicate on the "rmb_price" field.
func RmbPriceNotIn(vs ...string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldRmbPrice, vs...))
}

// RmbPriceGT applies the GT predicate on the "rmb_price" field.
func RmbPriceGT(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldRmbPrice, v))
}

// RmbPriceGTE applies the GTE predicate on the "rmb_price" field.
func RmbPriceGTE(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldRmbPrice, v))
}

// RmbPriceLT applies the LT predicate on the "rmb_price" field.
func RmbPriceLT(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldRmbPrice, v))
}

// RmbPriceLTE applies the LTE predicate on the "rmb_price" field.
func RmbPriceLTE(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldRmbPrice, v))
}

// RmbPriceContains applies the Contains predicate on the "rmb_price" field.
func RmbPriceContains(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldContains(FieldRmbPrice, v))
}

// RmbPriceHasPrefix applies the HasPrefix predicate on the "rmb_price" field.
func RmbPriceHasPrefix(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldHasPrefix(FieldRmbPrice, v))
}

// RmbPriceHasSuffix applies the HasSuffix predicate on the "rmb_price" field.
func RmbPriceHasSuffix(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldHasSuffix(FieldRmbPrice, v))
}

// RmbPriceEqualFold applies the EqualFold predicate on the "rmb_price" field.
func RmbPriceEqualFold(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEqualFold(FieldRmbPrice, v))
}

// RmbPriceContainsFold applies the ContainsFold predicate on the "rmb_price" field.
func RmbPriceContainsFold(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldContainsFold(FieldRmbPrice, v))
}

// OriginalRmbPriceEQ applies the EQ predicate on the "original_rmb_price" field.
func OriginalRmbPriceEQ(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldOriginalRmbPrice, v))
}

// OriginalRmbPriceNEQ applies the NEQ predicate on the "original_rmb_price" field.
func OriginalRmbPriceNEQ(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldOriginalRmbPrice, v))
}

// OriginalRmbPriceIn applies the In predicate on the "original_rmb_price" field.
func OriginalRmbPriceIn(vs ...string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldOriginalRmbPrice, vs...))
}

// OriginalRmbPriceNotIn applies the NotIn predicate on the "original_rmb_price" field.
func OriginalRmbPriceNotIn(vs ...string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldOriginalRmbPrice, vs...))
}

// OriginalRmbPriceGT applies the GT predicate on the "original_rmb_price" field.
func OriginalRmbPriceGT(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldOriginalRmbPrice, v))
}

// OriginalRmbPriceGTE applies the GTE predicate on the "original_rmb_price" field.
func OriginalRmbPriceGTE(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldOriginalRmbPrice, v))
}

// OriginalRmbPriceLT applies the LT predicate on the "original_rmb_price" field.
func OriginalRmbPriceLT(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldOriginalRmbPrice, v))
}

// OriginalRmbPriceLTE applies the LTE predicate on the "original_rmb_price" field.
func OriginalRmbPriceLTE(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldOriginalRmbPrice, v))
}

// OriginalRmbPriceContains applies the Contains predicate on the "original_rmb_price" field.
func OriginalRmbPriceContains(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldContains(FieldOriginalRmbPrice, v))
}

// OriginalRmbPriceHasPrefix applies the HasPrefix predicate on the "original_rmb_price" field.
func OriginalRmbPriceHasPrefix(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldHasPrefix(FieldOriginalRmbPrice, v))
}

// OriginalRmbPriceHasSuffix applies the HasSuffix predicate on the "original_rmb_price" field.
func OriginalRmbPriceHasSuffix(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldHasSuffix(FieldOriginalRmbPrice, v))
}

// OriginalRmbPriceEqualFold applies the EqualFold predicate on the "original_rmb_price" field.
func OriginalRmbPriceEqualFold(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEqualFold(FieldOriginalRmbPrice, v))
}

// OriginalRmbPriceContainsFold applies the ContainsFold predicate on the "original_rmb_price" field.
func OriginalRmbPriceContainsFold(v string) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldContainsFold(FieldOriginalRmbPrice, v))
}

// TotalCepEQ applies the EQ predicate on the "total_cep" field.
func TotalCepEQ(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldTotalCep, v))
}

// TotalCepNEQ applies the NEQ predicate on the "total_cep" field.
func TotalCepNEQ(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldTotalCep, v))
}

// TotalCepIn applies the In predicate on the "total_cep" field.
func TotalCepIn(vs ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldTotalCep, vs...))
}

// TotalCepNotIn applies the NotIn predicate on the "total_cep" field.
func TotalCepNotIn(vs ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldTotalCep, vs...))
}

// TotalCepGT applies the GT predicate on the "total_cep" field.
func TotalCepGT(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldTotalCep, v))
}

// TotalCepGTE applies the GTE predicate on the "total_cep" field.
func TotalCepGTE(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldTotalCep, v))
}

// TotalCepLT applies the LT predicate on the "total_cep" field.
func TotalCepLT(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldTotalCep, v))
}

// TotalCepLTE applies the LTE predicate on the "total_cep" field.
func TotalCepLTE(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldTotalCep, v))
}

// BeforeDiscountCepEQ applies the EQ predicate on the "before_discount_cep" field.
func BeforeDiscountCepEQ(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldBeforeDiscountCep, v))
}

// BeforeDiscountCepNEQ applies the NEQ predicate on the "before_discount_cep" field.
func BeforeDiscountCepNEQ(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldBeforeDiscountCep, v))
}

// BeforeDiscountCepIn applies the In predicate on the "before_discount_cep" field.
func BeforeDiscountCepIn(vs ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldBeforeDiscountCep, vs...))
}

// BeforeDiscountCepNotIn applies the NotIn predicate on the "before_discount_cep" field.
func BeforeDiscountCepNotIn(vs ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldBeforeDiscountCep, vs...))
}

// BeforeDiscountCepGT applies the GT predicate on the "before_discount_cep" field.
func BeforeDiscountCepGT(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldBeforeDiscountCep, v))
}

// BeforeDiscountCepGTE applies the GTE predicate on the "before_discount_cep" field.
func BeforeDiscountCepGTE(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldBeforeDiscountCep, v))
}

// BeforeDiscountCepLT applies the LT predicate on the "before_discount_cep" field.
func BeforeDiscountCepLT(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldBeforeDiscountCep, v))
}

// BeforeDiscountCepLTE applies the LTE predicate on the "before_discount_cep" field.
func BeforeDiscountCepLTE(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldBeforeDiscountCep, v))
}

// DiscountRatioEQ applies the EQ predicate on the "discount_ratio" field.
func DiscountRatioEQ(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldEQ(FieldDiscountRatio, v))
}

// DiscountRatioNEQ applies the NEQ predicate on the "discount_ratio" field.
func DiscountRatioNEQ(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNEQ(FieldDiscountRatio, v))
}

// DiscountRatioIn applies the In predicate on the "discount_ratio" field.
func DiscountRatioIn(vs ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldIn(FieldDiscountRatio, vs...))
}

// DiscountRatioNotIn applies the NotIn predicate on the "discount_ratio" field.
func DiscountRatioNotIn(vs ...int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldNotIn(FieldDiscountRatio, vs...))
}

// DiscountRatioGT applies the GT predicate on the "discount_ratio" field.
func DiscountRatioGT(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGT(FieldDiscountRatio, v))
}

// DiscountRatioGTE applies the GTE predicate on the "discount_ratio" field.
func DiscountRatioGTE(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldGTE(FieldDiscountRatio, v))
}

// DiscountRatioLT applies the LT predicate on the "discount_ratio" field.
func DiscountRatioLT(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLT(FieldDiscountRatio, v))
}

// DiscountRatioLTE applies the LTE predicate on the "discount_ratio" field.
func DiscountRatioLTE(v int64) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.FieldLTE(FieldDiscountRatio, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RechargeCampaignRuleOversea) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RechargeCampaignRuleOversea) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RechargeCampaignRuleOversea) predicate.RechargeCampaignRuleOversea {
	return predicate.RechargeCampaignRuleOversea(sql.NotPredicates(p))
}
