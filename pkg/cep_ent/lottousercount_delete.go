// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lottousercount"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// LottoUserCountDelete is the builder for deleting a LottoUserCount entity.
type LottoUserCountDelete struct {
	config
	hooks    []Hook
	mutation *LottoUserCountMutation
}

// Where appends a list predicates to the LottoUserCountDelete builder.
func (lucd *LottoUserCountDelete) Where(ps ...predicate.LottoUserCount) *LottoUserCountDelete {
	lucd.mutation.Where(ps...)
	return lucd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lucd *LottoUserCountDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lucd.sqlExec, lucd.mutation, lucd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lucd *LottoUserCountDelete) ExecX(ctx context.Context) int {
	n, err := lucd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lucd *LottoUserCountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(lottousercount.Table, sqlgraph.NewFieldSpec(lottousercount.FieldID, field.TypeInt64))
	if ps := lucd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lucd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lucd.mutation.done = true
	return affected, err
}

// LottoUserCountDeleteOne is the builder for deleting a single LottoUserCount entity.
type LottoUserCountDeleteOne struct {
	lucd *LottoUserCountDelete
}

// Where appends a list predicates to the LottoUserCountDelete builder.
func (lucdo *LottoUserCountDeleteOne) Where(ps ...predicate.LottoUserCount) *LottoUserCountDeleteOne {
	lucdo.lucd.mutation.Where(ps...)
	return lucdo
}

// Exec executes the deletion query.
func (lucdo *LottoUserCountDeleteOne) Exec(ctx context.Context) error {
	n, err := lucdo.lucd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{lottousercount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lucdo *LottoUserCountDeleteOne) ExecX(ctx context.Context) {
	if err := lucdo.Exec(ctx); err != nil {
		panic(err)
	}
}
