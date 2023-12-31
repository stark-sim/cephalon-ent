// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/enumcondition"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// EnumConditionDelete is the builder for deleting a EnumCondition entity.
type EnumConditionDelete struct {
	config
	hooks    []Hook
	mutation *EnumConditionMutation
}

// Where appends a list predicates to the EnumConditionDelete builder.
func (ecd *EnumConditionDelete) Where(ps ...predicate.EnumCondition) *EnumConditionDelete {
	ecd.mutation.Where(ps...)
	return ecd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ecd *EnumConditionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ecd.sqlExec, ecd.mutation, ecd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ecd *EnumConditionDelete) ExecX(ctx context.Context) int {
	n, err := ecd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ecd *EnumConditionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(enumcondition.Table, sqlgraph.NewFieldSpec(enumcondition.FieldID, field.TypeInt64))
	if ps := ecd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ecd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ecd.mutation.done = true
	return affected, err
}

// EnumConditionDeleteOne is the builder for deleting a single EnumCondition entity.
type EnumConditionDeleteOne struct {
	ecd *EnumConditionDelete
}

// Where appends a list predicates to the EnumConditionDelete builder.
func (ecdo *EnumConditionDeleteOne) Where(ps ...predicate.EnumCondition) *EnumConditionDeleteOne {
	ecdo.ecd.mutation.Where(ps...)
	return ecdo
}

// Exec executes the deletion query.
func (ecdo *EnumConditionDeleteOne) Exec(ctx context.Context) error {
	n, err := ecdo.ecd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{enumcondition.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ecdo *EnumConditionDeleteOne) ExecX(ctx context.Context) {
	if err := ecdo.Exec(ctx); err != nil {
		panic(err)
	}
}
