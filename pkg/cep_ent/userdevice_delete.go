// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/userdevice"
)

// UserDeviceDelete is the builder for deleting a UserDevice entity.
type UserDeviceDelete struct {
	config
	hooks    []Hook
	mutation *UserDeviceMutation
}

// Where appends a list predicates to the UserDeviceDelete builder.
func (udd *UserDeviceDelete) Where(ps ...predicate.UserDevice) *UserDeviceDelete {
	udd.mutation.Where(ps...)
	return udd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (udd *UserDeviceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, udd.sqlExec, udd.mutation, udd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (udd *UserDeviceDelete) ExecX(ctx context.Context) int {
	n, err := udd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (udd *UserDeviceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userdevice.Table, sqlgraph.NewFieldSpec(userdevice.FieldID, field.TypeInt64))
	if ps := udd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, udd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	udd.mutation.done = true
	return affected, err
}

// UserDeviceDeleteOne is the builder for deleting a single UserDevice entity.
type UserDeviceDeleteOne struct {
	udd *UserDeviceDelete
}

// Where appends a list predicates to the UserDeviceDelete builder.
func (uddo *UserDeviceDeleteOne) Where(ps ...predicate.UserDevice) *UserDeviceDeleteOne {
	uddo.udd.mutation.Where(ps...)
	return uddo
}

// Exec executes the deletion query.
func (uddo *UserDeviceDeleteOne) Exec(ctx context.Context) error {
	n, err := uddo.udd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userdevice.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uddo *UserDeviceDeleteOne) ExecX(ctx context.Context) {
	if err := uddo.Exec(ctx); err != nil {
		panic(err)
	}
}
