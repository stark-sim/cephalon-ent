// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/rechargeorder"
)

// RechargeOrderDelete is the builder for deleting a RechargeOrder entity.
type RechargeOrderDelete struct {
	config
	hooks    []Hook
	mutation *RechargeOrderMutation
}

// Where appends a list predicates to the RechargeOrderDelete builder.
func (rod *RechargeOrderDelete) Where(ps ...predicate.RechargeOrder) *RechargeOrderDelete {
	rod.mutation.Where(ps...)
	return rod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rod *RechargeOrderDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rod.sqlExec, rod.mutation, rod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rod *RechargeOrderDelete) ExecX(ctx context.Context) int {
	n, err := rod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rod *RechargeOrderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rechargeorder.Table, sqlgraph.NewFieldSpec(rechargeorder.FieldID, field.TypeInt64))
	if ps := rod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rod.mutation.done = true
	return affected, err
}

// RechargeOrderDeleteOne is the builder for deleting a single RechargeOrder entity.
type RechargeOrderDeleteOne struct {
	rod *RechargeOrderDelete
}

// Where appends a list predicates to the RechargeOrderDelete builder.
func (rodo *RechargeOrderDeleteOne) Where(ps ...predicate.RechargeOrder) *RechargeOrderDeleteOne {
	rodo.rod.mutation.Where(ps...)
	return rodo
}

// Exec executes the deletion query.
func (rodo *RechargeOrderDeleteOne) Exec(ctx context.Context) error {
	n, err := rodo.rod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rechargeorder.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rodo *RechargeOrderDeleteOne) ExecX(ctx context.Context) {
	if err := rodo.Exec(ctx); err != nil {
		panic(err)
	}
}
