// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionconsumeorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// MissionConsumeOrderDelete is the builder for deleting a MissionConsumeOrder entity.
type MissionConsumeOrderDelete struct {
	config
	hooks    []Hook
	mutation *MissionConsumeOrderMutation
}

// Where appends a list predicates to the MissionConsumeOrderDelete builder.
func (mcod *MissionConsumeOrderDelete) Where(ps ...predicate.MissionConsumeOrder) *MissionConsumeOrderDelete {
	mcod.mutation.Where(ps...)
	return mcod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mcod *MissionConsumeOrderDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mcod.sqlExec, mcod.mutation, mcod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mcod *MissionConsumeOrderDelete) ExecX(ctx context.Context) int {
	n, err := mcod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mcod *MissionConsumeOrderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(missionconsumeorder.Table, sqlgraph.NewFieldSpec(missionconsumeorder.FieldID, field.TypeInt64))
	if ps := mcod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mcod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mcod.mutation.done = true
	return affected, err
}

// MissionConsumeOrderDeleteOne is the builder for deleting a single MissionConsumeOrder entity.
type MissionConsumeOrderDeleteOne struct {
	mcod *MissionConsumeOrderDelete
}

// Where appends a list predicates to the MissionConsumeOrderDelete builder.
func (mcodo *MissionConsumeOrderDeleteOne) Where(ps ...predicate.MissionConsumeOrder) *MissionConsumeOrderDeleteOne {
	mcodo.mcod.mutation.Where(ps...)
	return mcodo
}

// Exec executes the deletion query.
func (mcodo *MissionConsumeOrderDeleteOne) Exec(ctx context.Context) error {
	n, err := mcodo.mcod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{missionconsumeorder.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mcodo *MissionConsumeOrderDeleteOne) ExecX(ctx context.Context) {
	if err := mcodo.Exec(ctx); err != nil {
		panic(err)
	}
}
