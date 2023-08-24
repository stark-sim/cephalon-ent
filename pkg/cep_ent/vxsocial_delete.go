// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/vxsocial"
)

// VXSocialDelete is the builder for deleting a VXSocial entity.
type VXSocialDelete struct {
	config
	hooks    []Hook
	mutation *VXSocialMutation
}

// Where appends a list predicates to the VXSocialDelete builder.
func (vsd *VXSocialDelete) Where(ps ...predicate.VXSocial) *VXSocialDelete {
	vsd.mutation.Where(ps...)
	return vsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vsd *VXSocialDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vsd.sqlExec, vsd.mutation, vsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vsd *VXSocialDelete) ExecX(ctx context.Context) int {
	n, err := vsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vsd *VXSocialDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(vxsocial.Table, sqlgraph.NewFieldSpec(vxsocial.FieldID, field.TypeInt64))
	if ps := vsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vsd.mutation.done = true
	return affected, err
}

// VXSocialDeleteOne is the builder for deleting a single VXSocial entity.
type VXSocialDeleteOne struct {
	vsd *VXSocialDelete
}

// Where appends a list predicates to the VXSocialDelete builder.
func (vsdo *VXSocialDeleteOne) Where(ps ...predicate.VXSocial) *VXSocialDeleteOne {
	vsdo.vsd.mutation.Where(ps...)
	return vsdo
}

// Exec executes the deletion query.
func (vsdo *VXSocialDeleteOne) Exec(ctx context.Context) error {
	n, err := vsdo.vsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vxsocial.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vsdo *VXSocialDeleteOne) ExecX(ctx context.Context) {
	if err := vsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
