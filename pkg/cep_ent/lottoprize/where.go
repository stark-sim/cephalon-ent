// Code generated by ent, DO NOT EDIT.

package lottoprize

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldDeletedAt, v))
}

// LottoID applies equality check predicate on the "lotto_id" field. It's identical to LottoIDEQ.
func LottoID(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldLottoID, v))
}

// LevelName applies equality check predicate on the "level_name" field. It's identical to LevelNameEQ.
func LevelName(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldLevelName, v))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldWeight, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldName, v))
}

// CepAmount applies equality check predicate on the "cep_amount" field. It's identical to CepAmountEQ.
func CepAmount(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldCepAmount, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLTE(FieldDeletedAt, v))
}

// LottoIDEQ applies the EQ predicate on the "lotto_id" field.
func LottoIDEQ(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldLottoID, v))
}

// LottoIDNEQ applies the NEQ predicate on the "lotto_id" field.
func LottoIDNEQ(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldLottoID, v))
}

// LottoIDIn applies the In predicate on the "lotto_id" field.
func LottoIDIn(vs ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldLottoID, vs...))
}

// LottoIDNotIn applies the NotIn predicate on the "lotto_id" field.
func LottoIDNotIn(vs ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldLottoID, vs...))
}

// LevelNameEQ applies the EQ predicate on the "level_name" field.
func LevelNameEQ(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldLevelName, v))
}

// LevelNameNEQ applies the NEQ predicate on the "level_name" field.
func LevelNameNEQ(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldLevelName, v))
}

// LevelNameIn applies the In predicate on the "level_name" field.
func LevelNameIn(vs ...string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldLevelName, vs...))
}

// LevelNameNotIn applies the NotIn predicate on the "level_name" field.
func LevelNameNotIn(vs ...string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldLevelName, vs...))
}

// LevelNameGT applies the GT predicate on the "level_name" field.
func LevelNameGT(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGT(FieldLevelName, v))
}

// LevelNameGTE applies the GTE predicate on the "level_name" field.
func LevelNameGTE(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGTE(FieldLevelName, v))
}

// LevelNameLT applies the LT predicate on the "level_name" field.
func LevelNameLT(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLT(FieldLevelName, v))
}

// LevelNameLTE applies the LTE predicate on the "level_name" field.
func LevelNameLTE(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLTE(FieldLevelName, v))
}

// LevelNameContains applies the Contains predicate on the "level_name" field.
func LevelNameContains(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldContains(FieldLevelName, v))
}

// LevelNameHasPrefix applies the HasPrefix predicate on the "level_name" field.
func LevelNameHasPrefix(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldHasPrefix(FieldLevelName, v))
}

// LevelNameHasSuffix applies the HasSuffix predicate on the "level_name" field.
func LevelNameHasSuffix(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldHasSuffix(FieldLevelName, v))
}

// LevelNameEqualFold applies the EqualFold predicate on the "level_name" field.
func LevelNameEqualFold(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEqualFold(FieldLevelName, v))
}

// LevelNameContainsFold applies the ContainsFold predicate on the "level_name" field.
func LevelNameContainsFold(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldContainsFold(FieldLevelName, v))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLTE(FieldWeight, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldContainsFold(FieldName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldStatus, vs...))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldType, vs...))
}

// CepAmountEQ applies the EQ predicate on the "cep_amount" field.
func CepAmountEQ(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldEQ(FieldCepAmount, v))
}

// CepAmountNEQ applies the NEQ predicate on the "cep_amount" field.
func CepAmountNEQ(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNEQ(FieldCepAmount, v))
}

// CepAmountIn applies the In predicate on the "cep_amount" field.
func CepAmountIn(vs ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldIn(FieldCepAmount, vs...))
}

// CepAmountNotIn applies the NotIn predicate on the "cep_amount" field.
func CepAmountNotIn(vs ...int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldNotIn(FieldCepAmount, vs...))
}

// CepAmountGT applies the GT predicate on the "cep_amount" field.
func CepAmountGT(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGT(FieldCepAmount, v))
}

// CepAmountGTE applies the GTE predicate on the "cep_amount" field.
func CepAmountGTE(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldGTE(FieldCepAmount, v))
}

// CepAmountLT applies the LT predicate on the "cep_amount" field.
func CepAmountLT(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLT(FieldCepAmount, v))
}

// CepAmountLTE applies the LTE predicate on the "cep_amount" field.
func CepAmountLTE(v int64) predicate.LottoPrize {
	return predicate.LottoPrize(sql.FieldLTE(FieldCepAmount, v))
}

// HasLotto applies the HasEdge predicate on the "lotto" edge.
func HasLotto() predicate.LottoPrize {
	return predicate.LottoPrize(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LottoTable, LottoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLottoWith applies the HasEdge predicate on the "lotto" edge with a given conditions (other predicates).
func HasLottoWith(preds ...predicate.Lotto) predicate.LottoPrize {
	return predicate.LottoPrize(func(s *sql.Selector) {
		step := newLottoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLottoRecords applies the HasEdge predicate on the "lotto_records" edge.
func HasLottoRecords() predicate.LottoPrize {
	return predicate.LottoPrize(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LottoRecordsTable, LottoRecordsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLottoRecordsWith applies the HasEdge predicate on the "lotto_records" edge with a given conditions (other predicates).
func HasLottoRecordsWith(preds ...predicate.LottoRecord) predicate.LottoPrize {
	return predicate.LottoPrize(func(s *sql.Selector) {
		step := newLottoRecordsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LottoPrize) predicate.LottoPrize {
	return predicate.LottoPrize(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LottoPrize) predicate.LottoPrize {
	return predicate.LottoPrize(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LottoPrize) predicate.LottoPrize {
	return predicate.LottoPrize(sql.NotPredicates(p))
}
