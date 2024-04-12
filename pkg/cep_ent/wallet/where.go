// Code generated by ent, DO NOT EDIT.

package wallet

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUserID, v))
}

// SymbolID applies equality check predicate on the "symbol_id" field. It's identical to SymbolIDEQ.
func SymbolID(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldSymbolID, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldAmount, v))
}

// WithdrawAmount applies equality check predicate on the "withdraw_amount" field. It's identical to WithdrawAmountEQ.
func WithdrawAmount(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldWithdrawAmount, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldDeletedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldUserID, vs...))
}

// SymbolIDEQ applies the EQ predicate on the "symbol_id" field.
func SymbolIDEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldSymbolID, v))
}

// SymbolIDNEQ applies the NEQ predicate on the "symbol_id" field.
func SymbolIDNEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldSymbolID, v))
}

// SymbolIDIn applies the In predicate on the "symbol_id" field.
func SymbolIDIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldSymbolID, vs...))
}

// SymbolIDNotIn applies the NotIn predicate on the "symbol_id" field.
func SymbolIDNotIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldSymbolID, vs...))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldAmount, v))
}

// WithdrawAmountEQ applies the EQ predicate on the "withdraw_amount" field.
func WithdrawAmountEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldWithdrawAmount, v))
}

// WithdrawAmountNEQ applies the NEQ predicate on the "withdraw_amount" field.
func WithdrawAmountNEQ(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldWithdrawAmount, v))
}

// WithdrawAmountIn applies the In predicate on the "withdraw_amount" field.
func WithdrawAmountIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldWithdrawAmount, vs...))
}

// WithdrawAmountNotIn applies the NotIn predicate on the "withdraw_amount" field.
func WithdrawAmountNotIn(vs ...int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldWithdrawAmount, vs...))
}

// WithdrawAmountGT applies the GT predicate on the "withdraw_amount" field.
func WithdrawAmountGT(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldWithdrawAmount, v))
}

// WithdrawAmountGTE applies the GTE predicate on the "withdraw_amount" field.
func WithdrawAmountGTE(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldWithdrawAmount, v))
}

// WithdrawAmountLT applies the LT predicate on the "withdraw_amount" field.
func WithdrawAmountLT(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldWithdrawAmount, v))
}

// WithdrawAmountLTE applies the LTE predicate on the "withdraw_amount" field.
func WithdrawAmountLTE(v int64) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldWithdrawAmount, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSymbol applies the HasEdge predicate on the "symbol" edge.
func HasSymbol() predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SymbolTable, SymbolColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSymbolWith applies the HasEdge predicate on the "symbol" edge with a given conditions (other predicates).
func HasSymbolWith(preds ...predicate.Symbol) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		step := newSymbolStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(sql.NotPredicates(p))
}
