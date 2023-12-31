// Code generated by ent, DO NOT EDIT.

package profitaccount

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldUserID, v))
}

// SumCep applies equality check predicate on the "sum_cep" field. It's identical to SumCepEQ.
func SumCep(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldSumCep, v))
}

// RemainCep applies equality check predicate on the "remain_cep" field. It's identical to RemainCepEQ.
func RemainCep(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldRemainCep, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLTE(FieldDeletedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNotIn(FieldUserID, vs...))
}

// SumCepEQ applies the EQ predicate on the "sum_cep" field.
func SumCepEQ(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldSumCep, v))
}

// SumCepNEQ applies the NEQ predicate on the "sum_cep" field.
func SumCepNEQ(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNEQ(FieldSumCep, v))
}

// SumCepIn applies the In predicate on the "sum_cep" field.
func SumCepIn(vs ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldIn(FieldSumCep, vs...))
}

// SumCepNotIn applies the NotIn predicate on the "sum_cep" field.
func SumCepNotIn(vs ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNotIn(FieldSumCep, vs...))
}

// SumCepGT applies the GT predicate on the "sum_cep" field.
func SumCepGT(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGT(FieldSumCep, v))
}

// SumCepGTE applies the GTE predicate on the "sum_cep" field.
func SumCepGTE(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGTE(FieldSumCep, v))
}

// SumCepLT applies the LT predicate on the "sum_cep" field.
func SumCepLT(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLT(FieldSumCep, v))
}

// SumCepLTE applies the LTE predicate on the "sum_cep" field.
func SumCepLTE(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLTE(FieldSumCep, v))
}

// RemainCepEQ applies the EQ predicate on the "remain_cep" field.
func RemainCepEQ(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldEQ(FieldRemainCep, v))
}

// RemainCepNEQ applies the NEQ predicate on the "remain_cep" field.
func RemainCepNEQ(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNEQ(FieldRemainCep, v))
}

// RemainCepIn applies the In predicate on the "remain_cep" field.
func RemainCepIn(vs ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldIn(FieldRemainCep, vs...))
}

// RemainCepNotIn applies the NotIn predicate on the "remain_cep" field.
func RemainCepNotIn(vs ...int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldNotIn(FieldRemainCep, vs...))
}

// RemainCepGT applies the GT predicate on the "remain_cep" field.
func RemainCepGT(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGT(FieldRemainCep, v))
}

// RemainCepGTE applies the GTE predicate on the "remain_cep" field.
func RemainCepGTE(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldGTE(FieldRemainCep, v))
}

// RemainCepLT applies the LT predicate on the "remain_cep" field.
func RemainCepLT(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLT(FieldRemainCep, v))
}

// RemainCepLTE applies the LTE predicate on the "remain_cep" field.
func RemainCepLTE(v int64) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.FieldLTE(FieldRemainCep, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ProfitAccount {
	return predicate.ProfitAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ProfitAccount {
	return predicate.ProfitAccount(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEarnBills applies the HasEdge predicate on the "earn_bills" edge.
func HasEarnBills() predicate.ProfitAccount {
	return predicate.ProfitAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EarnBillsTable, EarnBillsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEarnBillsWith applies the HasEdge predicate on the "earn_bills" edge with a given conditions (other predicates).
func HasEarnBillsWith(preds ...predicate.EarnBill) predicate.ProfitAccount {
	return predicate.ProfitAccount(func(s *sql.Selector) {
		step := newEarnBillsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProfitAccount) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProfitAccount) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProfitAccount) predicate.ProfitAccount {
	return predicate.ProfitAccount(sql.NotPredicates(p))
}
