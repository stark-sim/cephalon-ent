// Code generated by ent, DO NOT EDIT.

package modelprice

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldDeletedAt, v))
}

// ModelID applies equality check predicate on the "model_id" field. It's identical to ModelIDEQ.
func ModelID(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldModelID, v))
}

// InputGpuPrice applies equality check predicate on the "input_gpu_price" field. It's identical to InputGpuPriceEQ.
func InputGpuPrice(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldInputGpuPrice, v))
}

// OutputGpuPrice applies equality check predicate on the "output_gpu_price" field. It's identical to OutputGpuPriceEQ.
func OutputGpuPrice(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldOutputGpuPrice, v))
}

// InputModelPrice applies equality check predicate on the "input_model_price" field. It's identical to InputModelPriceEQ.
func InputModelPrice(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldInputModelPrice, v))
}

// OutputModelPrice applies equality check predicate on the "output_model_price" field. It's identical to OutputModelPriceEQ.
func OutputModelPrice(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldOutputModelPrice, v))
}

// TokenPerCep applies equality check predicate on the "token_per_cep" field. It's identical to TokenPerCepEQ.
func TokenPerCep(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldTokenPerCep, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLTE(FieldDeletedAt, v))
}

// ModelIDEQ applies the EQ predicate on the "model_id" field.
func ModelIDEQ(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldModelID, v))
}

// ModelIDNEQ applies the NEQ predicate on the "model_id" field.
func ModelIDNEQ(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldModelID, v))
}

// ModelIDIn applies the In predicate on the "model_id" field.
func ModelIDIn(vs ...int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldModelID, vs...))
}

// ModelIDNotIn applies the NotIn predicate on the "model_id" field.
func ModelIDNotIn(vs ...int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldModelID, vs...))
}

// InvokeTypeEQ applies the EQ predicate on the "invoke_type" field.
func InvokeTypeEQ(v enums.InvokeType) predicate.ModelPrice {
	vc := v
	return predicate.ModelPrice(sql.FieldEQ(FieldInvokeType, vc))
}

// InvokeTypeNEQ applies the NEQ predicate on the "invoke_type" field.
func InvokeTypeNEQ(v enums.InvokeType) predicate.ModelPrice {
	vc := v
	return predicate.ModelPrice(sql.FieldNEQ(FieldInvokeType, vc))
}

// InvokeTypeIn applies the In predicate on the "invoke_type" field.
func InvokeTypeIn(vs ...enums.InvokeType) predicate.ModelPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ModelPrice(sql.FieldIn(FieldInvokeType, v...))
}

// InvokeTypeNotIn applies the NotIn predicate on the "invoke_type" field.
func InvokeTypeNotIn(vs ...enums.InvokeType) predicate.ModelPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ModelPrice(sql.FieldNotIn(FieldInvokeType, v...))
}

// GpuVersionEQ applies the EQ predicate on the "gpu_version" field.
func GpuVersionEQ(v enums.GpuVersion) predicate.ModelPrice {
	vc := v
	return predicate.ModelPrice(sql.FieldEQ(FieldGpuVersion, vc))
}

// GpuVersionNEQ applies the NEQ predicate on the "gpu_version" field.
func GpuVersionNEQ(v enums.GpuVersion) predicate.ModelPrice {
	vc := v
	return predicate.ModelPrice(sql.FieldNEQ(FieldGpuVersion, vc))
}

// GpuVersionIn applies the In predicate on the "gpu_version" field.
func GpuVersionIn(vs ...enums.GpuVersion) predicate.ModelPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ModelPrice(sql.FieldIn(FieldGpuVersion, v...))
}

// GpuVersionNotIn applies the NotIn predicate on the "gpu_version" field.
func GpuVersionNotIn(vs ...enums.GpuVersion) predicate.ModelPrice {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ModelPrice(sql.FieldNotIn(FieldGpuVersion, v...))
}

// InputGpuPriceEQ applies the EQ predicate on the "input_gpu_price" field.
func InputGpuPriceEQ(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldInputGpuPrice, v))
}

// InputGpuPriceNEQ applies the NEQ predicate on the "input_gpu_price" field.
func InputGpuPriceNEQ(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldInputGpuPrice, v))
}

// InputGpuPriceIn applies the In predicate on the "input_gpu_price" field.
func InputGpuPriceIn(vs ...int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldInputGpuPrice, vs...))
}

// InputGpuPriceNotIn applies the NotIn predicate on the "input_gpu_price" field.
func InputGpuPriceNotIn(vs ...int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldInputGpuPrice, vs...))
}

// InputGpuPriceGT applies the GT predicate on the "input_gpu_price" field.
func InputGpuPriceGT(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGT(FieldInputGpuPrice, v))
}

// InputGpuPriceGTE applies the GTE predicate on the "input_gpu_price" field.
func InputGpuPriceGTE(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGTE(FieldInputGpuPrice, v))
}

// InputGpuPriceLT applies the LT predicate on the "input_gpu_price" field.
func InputGpuPriceLT(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLT(FieldInputGpuPrice, v))
}

// InputGpuPriceLTE applies the LTE predicate on the "input_gpu_price" field.
func InputGpuPriceLTE(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLTE(FieldInputGpuPrice, v))
}

// OutputGpuPriceEQ applies the EQ predicate on the "output_gpu_price" field.
func OutputGpuPriceEQ(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldOutputGpuPrice, v))
}

// OutputGpuPriceNEQ applies the NEQ predicate on the "output_gpu_price" field.
func OutputGpuPriceNEQ(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldOutputGpuPrice, v))
}

// OutputGpuPriceIn applies the In predicate on the "output_gpu_price" field.
func OutputGpuPriceIn(vs ...int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldOutputGpuPrice, vs...))
}

// OutputGpuPriceNotIn applies the NotIn predicate on the "output_gpu_price" field.
func OutputGpuPriceNotIn(vs ...int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldOutputGpuPrice, vs...))
}

// OutputGpuPriceGT applies the GT predicate on the "output_gpu_price" field.
func OutputGpuPriceGT(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGT(FieldOutputGpuPrice, v))
}

// OutputGpuPriceGTE applies the GTE predicate on the "output_gpu_price" field.
func OutputGpuPriceGTE(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGTE(FieldOutputGpuPrice, v))
}

// OutputGpuPriceLT applies the LT predicate on the "output_gpu_price" field.
func OutputGpuPriceLT(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLT(FieldOutputGpuPrice, v))
}

// OutputGpuPriceLTE applies the LTE predicate on the "output_gpu_price" field.
func OutputGpuPriceLTE(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLTE(FieldOutputGpuPrice, v))
}

// InputModelPriceEQ applies the EQ predicate on the "input_model_price" field.
func InputModelPriceEQ(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldInputModelPrice, v))
}

// InputModelPriceNEQ applies the NEQ predicate on the "input_model_price" field.
func InputModelPriceNEQ(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldInputModelPrice, v))
}

// InputModelPriceIn applies the In predicate on the "input_model_price" field.
func InputModelPriceIn(vs ...int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldInputModelPrice, vs...))
}

// InputModelPriceNotIn applies the NotIn predicate on the "input_model_price" field.
func InputModelPriceNotIn(vs ...int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldInputModelPrice, vs...))
}

// InputModelPriceGT applies the GT predicate on the "input_model_price" field.
func InputModelPriceGT(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGT(FieldInputModelPrice, v))
}

// InputModelPriceGTE applies the GTE predicate on the "input_model_price" field.
func InputModelPriceGTE(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGTE(FieldInputModelPrice, v))
}

// InputModelPriceLT applies the LT predicate on the "input_model_price" field.
func InputModelPriceLT(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLT(FieldInputModelPrice, v))
}

// InputModelPriceLTE applies the LTE predicate on the "input_model_price" field.
func InputModelPriceLTE(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLTE(FieldInputModelPrice, v))
}

// OutputModelPriceEQ applies the EQ predicate on the "output_model_price" field.
func OutputModelPriceEQ(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldOutputModelPrice, v))
}

// OutputModelPriceNEQ applies the NEQ predicate on the "output_model_price" field.
func OutputModelPriceNEQ(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldOutputModelPrice, v))
}

// OutputModelPriceIn applies the In predicate on the "output_model_price" field.
func OutputModelPriceIn(vs ...int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldOutputModelPrice, vs...))
}

// OutputModelPriceNotIn applies the NotIn predicate on the "output_model_price" field.
func OutputModelPriceNotIn(vs ...int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldOutputModelPrice, vs...))
}

// OutputModelPriceGT applies the GT predicate on the "output_model_price" field.
func OutputModelPriceGT(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGT(FieldOutputModelPrice, v))
}

// OutputModelPriceGTE applies the GTE predicate on the "output_model_price" field.
func OutputModelPriceGTE(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGTE(FieldOutputModelPrice, v))
}

// OutputModelPriceLT applies the LT predicate on the "output_model_price" field.
func OutputModelPriceLT(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLT(FieldOutputModelPrice, v))
}

// OutputModelPriceLTE applies the LTE predicate on the "output_model_price" field.
func OutputModelPriceLTE(v int) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLTE(FieldOutputModelPrice, v))
}

// TokenPerCepEQ applies the EQ predicate on the "token_per_cep" field.
func TokenPerCepEQ(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldEQ(FieldTokenPerCep, v))
}

// TokenPerCepNEQ applies the NEQ predicate on the "token_per_cep" field.
func TokenPerCepNEQ(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNEQ(FieldTokenPerCep, v))
}

// TokenPerCepIn applies the In predicate on the "token_per_cep" field.
func TokenPerCepIn(vs ...int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldIn(FieldTokenPerCep, vs...))
}

// TokenPerCepNotIn applies the NotIn predicate on the "token_per_cep" field.
func TokenPerCepNotIn(vs ...int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldNotIn(FieldTokenPerCep, vs...))
}

// TokenPerCepGT applies the GT predicate on the "token_per_cep" field.
func TokenPerCepGT(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGT(FieldTokenPerCep, v))
}

// TokenPerCepGTE applies the GTE predicate on the "token_per_cep" field.
func TokenPerCepGTE(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldGTE(FieldTokenPerCep, v))
}

// TokenPerCepLT applies the LT predicate on the "token_per_cep" field.
func TokenPerCepLT(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLT(FieldTokenPerCep, v))
}

// TokenPerCepLTE applies the LTE predicate on the "token_per_cep" field.
func TokenPerCepLTE(v int64) predicate.ModelPrice {
	return predicate.ModelPrice(sql.FieldLTE(FieldTokenPerCep, v))
}

// HasModel applies the HasEdge predicate on the "model" edge.
func HasModel() predicate.ModelPrice {
	return predicate.ModelPrice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ModelTable, ModelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModelWith applies the HasEdge predicate on the "model" edge with a given conditions (other predicates).
func HasModelWith(preds ...predicate.Model) predicate.ModelPrice {
	return predicate.ModelPrice(func(s *sql.Selector) {
		step := newModelStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ModelPrice) predicate.ModelPrice {
	return predicate.ModelPrice(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ModelPrice) predicate.ModelPrice {
	return predicate.ModelPrice(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ModelPrice) predicate.ModelPrice {
	return predicate.ModelPrice(sql.NotPredicates(p))
}