// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/profitaccount"
)

// ProfitAccountDelete is the builder for deleting a ProfitAccount entity.
type ProfitAccountDelete struct {
	config
	hooks    []Hook
	mutation *ProfitAccountMutation
}

// Where appends a list predicates to the ProfitAccountDelete builder.
func (pad *ProfitAccountDelete) Where(ps ...predicate.ProfitAccount) *ProfitAccountDelete {
	pad.mutation.Where(ps...)
	return pad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pad *ProfitAccountDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pad.sqlExec, pad.mutation, pad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pad *ProfitAccountDelete) ExecX(ctx context.Context) int {
	n, err := pad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pad *ProfitAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(profitaccount.Table, sqlgraph.NewFieldSpec(profitaccount.FieldID, field.TypeInt64))
	if ps := pad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pad.mutation.done = true
	return affected, err
}

// ProfitAccountDeleteOne is the builder for deleting a single ProfitAccount entity.
type ProfitAccountDeleteOne struct {
	pad *ProfitAccountDelete
}

// Where appends a list predicates to the ProfitAccountDelete builder.
func (pado *ProfitAccountDeleteOne) Where(ps ...predicate.ProfitAccount) *ProfitAccountDeleteOne {
	pado.pad.mutation.Where(ps...)
	return pado
}

// Exec executes the deletion query.
func (pado *ProfitAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := pado.pad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{profitaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pado *ProfitAccountDeleteOne) ExecX(ctx context.Context) {
	if err := pado.Exec(ctx); err != nil {
		panic(err)
	}
}
