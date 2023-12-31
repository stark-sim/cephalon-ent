// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/extraserviceprice"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ExtraServicePriceDelete is the builder for deleting a ExtraServicePrice entity.
type ExtraServicePriceDelete struct {
	config
	hooks    []Hook
	mutation *ExtraServicePriceMutation
}

// Where appends a list predicates to the ExtraServicePriceDelete builder.
func (espd *ExtraServicePriceDelete) Where(ps ...predicate.ExtraServicePrice) *ExtraServicePriceDelete {
	espd.mutation.Where(ps...)
	return espd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (espd *ExtraServicePriceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, espd.sqlExec, espd.mutation, espd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (espd *ExtraServicePriceDelete) ExecX(ctx context.Context) int {
	n, err := espd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (espd *ExtraServicePriceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(extraserviceprice.Table, sqlgraph.NewFieldSpec(extraserviceprice.FieldID, field.TypeInt64))
	if ps := espd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, espd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	espd.mutation.done = true
	return affected, err
}

// ExtraServicePriceDeleteOne is the builder for deleting a single ExtraServicePrice entity.
type ExtraServicePriceDeleteOne struct {
	espd *ExtraServicePriceDelete
}

// Where appends a list predicates to the ExtraServicePriceDelete builder.
func (espdo *ExtraServicePriceDeleteOne) Where(ps ...predicate.ExtraServicePrice) *ExtraServicePriceDeleteOne {
	espdo.espd.mutation.Where(ps...)
	return espdo
}

// Exec executes the deletion query.
func (espdo *ExtraServicePriceDeleteOne) Exec(ctx context.Context) error {
	n, err := espdo.espd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{extraserviceprice.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (espdo *ExtraServicePriceDeleteOne) ExecX(ctx context.Context) {
	if err := espdo.Exec(ctx); err != nil {
		panic(err)
	}
}
