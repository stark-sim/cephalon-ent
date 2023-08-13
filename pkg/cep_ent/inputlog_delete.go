// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/inputlog"
	"cephalon-ent/pkg/cep_ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InputLogDelete is the builder for deleting a InputLog entity.
type InputLogDelete struct {
	config
	hooks    []Hook
	mutation *InputLogMutation
}

// Where appends a list predicates to the InputLogDelete builder.
func (ild *InputLogDelete) Where(ps ...predicate.InputLog) *InputLogDelete {
	ild.mutation.Where(ps...)
	return ild
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ild *InputLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ild.sqlExec, ild.mutation, ild.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ild *InputLogDelete) ExecX(ctx context.Context) int {
	n, err := ild.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ild *InputLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(inputlog.Table, sqlgraph.NewFieldSpec(inputlog.FieldID, field.TypeInt64))
	if ps := ild.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ild.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ild.mutation.done = true
	return affected, err
}

// InputLogDeleteOne is the builder for deleting a single InputLog entity.
type InputLogDeleteOne struct {
	ild *InputLogDelete
}

// Where appends a list predicates to the InputLogDelete builder.
func (ildo *InputLogDeleteOne) Where(ps ...predicate.InputLog) *InputLogDeleteOne {
	ildo.ild.mutation.Where(ps...)
	return ildo
}

// Exec executes the deletion query.
func (ildo *InputLogDeleteOne) Exec(ctx context.Context) error {
	n, err := ildo.ild.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{inputlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ildo *InputLogDeleteOne) ExecX(ctx context.Context) {
	if err := ildo.Exec(ctx); err != nil {
		panic(err)
	}
}
