// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/apitoken"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/bill"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/invokemodelorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/model"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// InvokeModelOrderUpdate is the builder for updating InvokeModelOrder entities.
type InvokeModelOrderUpdate struct {
	config
	hooks     []Hook
	mutation  *InvokeModelOrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the InvokeModelOrderUpdate builder.
func (imou *InvokeModelOrderUpdate) Where(ps ...predicate.InvokeModelOrder) *InvokeModelOrderUpdate {
	imou.mutation.Where(ps...)
	return imou
}

// SetCreatedBy sets the "created_by" field.
func (imou *InvokeModelOrderUpdate) SetCreatedBy(i int64) *InvokeModelOrderUpdate {
	imou.mutation.ResetCreatedBy()
	imou.mutation.SetCreatedBy(i)
	return imou
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (imou *InvokeModelOrderUpdate) SetNillableCreatedBy(i *int64) *InvokeModelOrderUpdate {
	if i != nil {
		imou.SetCreatedBy(*i)
	}
	return imou
}

// AddCreatedBy adds i to the "created_by" field.
func (imou *InvokeModelOrderUpdate) AddCreatedBy(i int64) *InvokeModelOrderUpdate {
	imou.mutation.AddCreatedBy(i)
	return imou
}

// SetUpdatedBy sets the "updated_by" field.
func (imou *InvokeModelOrderUpdate) SetUpdatedBy(i int64) *InvokeModelOrderUpdate {
	imou.mutation.ResetUpdatedBy()
	imou.mutation.SetUpdatedBy(i)
	return imou
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (imou *InvokeModelOrderUpdate) SetNillableUpdatedBy(i *int64) *InvokeModelOrderUpdate {
	if i != nil {
		imou.SetUpdatedBy(*i)
	}
	return imou
}

// AddUpdatedBy adds i to the "updated_by" field.
func (imou *InvokeModelOrderUpdate) AddUpdatedBy(i int64) *InvokeModelOrderUpdate {
	imou.mutation.AddUpdatedBy(i)
	return imou
}

// SetUpdatedAt sets the "updated_at" field.
func (imou *InvokeModelOrderUpdate) SetUpdatedAt(t time.Time) *InvokeModelOrderUpdate {
	imou.mutation.SetUpdatedAt(t)
	return imou
}

// SetDeletedAt sets the "deleted_at" field.
func (imou *InvokeModelOrderUpdate) SetDeletedAt(t time.Time) *InvokeModelOrderUpdate {
	imou.mutation.SetDeletedAt(t)
	return imou
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (imou *InvokeModelOrderUpdate) SetNillableDeletedAt(t *time.Time) *InvokeModelOrderUpdate {
	if t != nil {
		imou.SetDeletedAt(*t)
	}
	return imou
}

// SetUserID sets the "user_id" field.
func (imou *InvokeModelOrderUpdate) SetUserID(i int64) *InvokeModelOrderUpdate {
	imou.mutation.SetUserID(i)
	return imou
}

// SetModelID sets the "model_id" field.
func (imou *InvokeModelOrderUpdate) SetModelID(i int64) *InvokeModelOrderUpdate {
	imou.mutation.SetModelID(i)
	return imou
}

// SetAPITokenID sets the "api_token_id" field.
func (imou *InvokeModelOrderUpdate) SetAPITokenID(i int64) *InvokeModelOrderUpdate {
	imou.mutation.SetAPITokenID(i)
	return imou
}

// SetInvokeType sets the "invoke_type" field.
func (imou *InvokeModelOrderUpdate) SetInvokeType(et enums.InvokeType) *InvokeModelOrderUpdate {
	imou.mutation.SetInvokeType(et)
	return imou
}

// SetNillableInvokeType sets the "invoke_type" field if the given value is not nil.
func (imou *InvokeModelOrderUpdate) SetNillableInvokeType(et *enums.InvokeType) *InvokeModelOrderUpdate {
	if et != nil {
		imou.SetInvokeType(*et)
	}
	return imou
}

// SetInvokeTimes sets the "invoke_times" field.
func (imou *InvokeModelOrderUpdate) SetInvokeTimes(i int) *InvokeModelOrderUpdate {
	imou.mutation.ResetInvokeTimes()
	imou.mutation.SetInvokeTimes(i)
	return imou
}

// SetNillableInvokeTimes sets the "invoke_times" field if the given value is not nil.
func (imou *InvokeModelOrderUpdate) SetNillableInvokeTimes(i *int) *InvokeModelOrderUpdate {
	if i != nil {
		imou.SetInvokeTimes(*i)
	}
	return imou
}

// AddInvokeTimes adds i to the "invoke_times" field.
func (imou *InvokeModelOrderUpdate) AddInvokeTimes(i int) *InvokeModelOrderUpdate {
	imou.mutation.AddInvokeTimes(i)
	return imou
}

// SetInputTokenCost sets the "input_token_cost" field.
func (imou *InvokeModelOrderUpdate) SetInputTokenCost(i int) *InvokeModelOrderUpdate {
	imou.mutation.ResetInputTokenCost()
	imou.mutation.SetInputTokenCost(i)
	return imou
}

// SetNillableInputTokenCost sets the "input_token_cost" field if the given value is not nil.
func (imou *InvokeModelOrderUpdate) SetNillableInputTokenCost(i *int) *InvokeModelOrderUpdate {
	if i != nil {
		imou.SetInputTokenCost(*i)
	}
	return imou
}

// AddInputTokenCost adds i to the "input_token_cost" field.
func (imou *InvokeModelOrderUpdate) AddInputTokenCost(i int) *InvokeModelOrderUpdate {
	imou.mutation.AddInputTokenCost(i)
	return imou
}

// SetOutputTokenCost sets the "output_token_cost" field.
func (imou *InvokeModelOrderUpdate) SetOutputTokenCost(i int) *InvokeModelOrderUpdate {
	imou.mutation.ResetOutputTokenCost()
	imou.mutation.SetOutputTokenCost(i)
	return imou
}

// SetNillableOutputTokenCost sets the "output_token_cost" field if the given value is not nil.
func (imou *InvokeModelOrderUpdate) SetNillableOutputTokenCost(i *int) *InvokeModelOrderUpdate {
	if i != nil {
		imou.SetOutputTokenCost(*i)
	}
	return imou
}

// AddOutputTokenCost adds i to the "output_token_cost" field.
func (imou *InvokeModelOrderUpdate) AddOutputTokenCost(i int) *InvokeModelOrderUpdate {
	imou.mutation.AddOutputTokenCost(i)
	return imou
}

// SetInputCepCost sets the "input_cep_cost" field.
func (imou *InvokeModelOrderUpdate) SetInputCepCost(i int) *InvokeModelOrderUpdate {
	imou.mutation.ResetInputCepCost()
	imou.mutation.SetInputCepCost(i)
	return imou
}

// SetNillableInputCepCost sets the "input_cep_cost" field if the given value is not nil.
func (imou *InvokeModelOrderUpdate) SetNillableInputCepCost(i *int) *InvokeModelOrderUpdate {
	if i != nil {
		imou.SetInputCepCost(*i)
	}
	return imou
}

// AddInputCepCost adds i to the "input_cep_cost" field.
func (imou *InvokeModelOrderUpdate) AddInputCepCost(i int) *InvokeModelOrderUpdate {
	imou.mutation.AddInputCepCost(i)
	return imou
}

// SetOutputCepCost sets the "output_cep_cost" field.
func (imou *InvokeModelOrderUpdate) SetOutputCepCost(i int) *InvokeModelOrderUpdate {
	imou.mutation.ResetOutputCepCost()
	imou.mutation.SetOutputCepCost(i)
	return imou
}

// SetNillableOutputCepCost sets the "output_cep_cost" field if the given value is not nil.
func (imou *InvokeModelOrderUpdate) SetNillableOutputCepCost(i *int) *InvokeModelOrderUpdate {
	if i != nil {
		imou.SetOutputCepCost(*i)
	}
	return imou
}

// AddOutputCepCost adds i to the "output_cep_cost" field.
func (imou *InvokeModelOrderUpdate) AddOutputCepCost(i int) *InvokeModelOrderUpdate {
	imou.mutation.AddOutputCepCost(i)
	return imou
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (imou *InvokeModelOrderUpdate) AddBillIDs(ids ...int64) *InvokeModelOrderUpdate {
	imou.mutation.AddBillIDs(ids...)
	return imou
}

// AddBills adds the "bills" edges to the Bill entity.
func (imou *InvokeModelOrderUpdate) AddBills(b ...*Bill) *InvokeModelOrderUpdate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return imou.AddBillIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (imou *InvokeModelOrderUpdate) SetUser(u *User) *InvokeModelOrderUpdate {
	return imou.SetUserID(u.ID)
}

// SetModel sets the "model" edge to the Model entity.
func (imou *InvokeModelOrderUpdate) SetModel(m *Model) *InvokeModelOrderUpdate {
	return imou.SetModelID(m.ID)
}

// SetAPIToken sets the "api_token" edge to the ApiToken entity.
func (imou *InvokeModelOrderUpdate) SetAPIToken(a *ApiToken) *InvokeModelOrderUpdate {
	return imou.SetAPITokenID(a.ID)
}

// Mutation returns the InvokeModelOrderMutation object of the builder.
func (imou *InvokeModelOrderUpdate) Mutation() *InvokeModelOrderMutation {
	return imou.mutation
}

// ClearBills clears all "bills" edges to the Bill entity.
func (imou *InvokeModelOrderUpdate) ClearBills() *InvokeModelOrderUpdate {
	imou.mutation.ClearBills()
	return imou
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (imou *InvokeModelOrderUpdate) RemoveBillIDs(ids ...int64) *InvokeModelOrderUpdate {
	imou.mutation.RemoveBillIDs(ids...)
	return imou
}

// RemoveBills removes "bills" edges to Bill entities.
func (imou *InvokeModelOrderUpdate) RemoveBills(b ...*Bill) *InvokeModelOrderUpdate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return imou.RemoveBillIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (imou *InvokeModelOrderUpdate) ClearUser() *InvokeModelOrderUpdate {
	imou.mutation.ClearUser()
	return imou
}

// ClearModel clears the "model" edge to the Model entity.
func (imou *InvokeModelOrderUpdate) ClearModel() *InvokeModelOrderUpdate {
	imou.mutation.ClearModel()
	return imou
}

// ClearAPIToken clears the "api_token" edge to the ApiToken entity.
func (imou *InvokeModelOrderUpdate) ClearAPIToken() *InvokeModelOrderUpdate {
	imou.mutation.ClearAPIToken()
	return imou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (imou *InvokeModelOrderUpdate) Save(ctx context.Context) (int, error) {
	imou.defaults()
	return withHooks(ctx, imou.sqlSave, imou.mutation, imou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (imou *InvokeModelOrderUpdate) SaveX(ctx context.Context) int {
	affected, err := imou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (imou *InvokeModelOrderUpdate) Exec(ctx context.Context) error {
	_, err := imou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (imou *InvokeModelOrderUpdate) ExecX(ctx context.Context) {
	if err := imou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (imou *InvokeModelOrderUpdate) defaults() {
	if _, ok := imou.mutation.UpdatedAt(); !ok {
		v := invokemodelorder.UpdateDefaultUpdatedAt()
		imou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (imou *InvokeModelOrderUpdate) check() error {
	if v, ok := imou.mutation.InvokeType(); ok {
		if err := invokemodelorder.InvokeTypeValidator(v); err != nil {
			return &ValidationError{Name: "invoke_type", err: fmt.Errorf(`cep_ent: validator failed for field "InvokeModelOrder.invoke_type": %w`, err)}
		}
	}
	if _, ok := imou.mutation.UserID(); imou.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "InvokeModelOrder.user"`)
	}
	if _, ok := imou.mutation.ModelID(); imou.mutation.ModelCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "InvokeModelOrder.model"`)
	}
	if _, ok := imou.mutation.APITokenID(); imou.mutation.APITokenCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "InvokeModelOrder.api_token"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (imou *InvokeModelOrderUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InvokeModelOrderUpdate {
	imou.modifiers = append(imou.modifiers, modifiers...)
	return imou
}

func (imou *InvokeModelOrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := imou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(invokemodelorder.Table, invokemodelorder.Columns, sqlgraph.NewFieldSpec(invokemodelorder.FieldID, field.TypeInt64))
	if ps := imou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := imou.mutation.CreatedBy(); ok {
		_spec.SetField(invokemodelorder.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := imou.mutation.AddedCreatedBy(); ok {
		_spec.AddField(invokemodelorder.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := imou.mutation.UpdatedBy(); ok {
		_spec.SetField(invokemodelorder.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := imou.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(invokemodelorder.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := imou.mutation.UpdatedAt(); ok {
		_spec.SetField(invokemodelorder.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := imou.mutation.DeletedAt(); ok {
		_spec.SetField(invokemodelorder.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := imou.mutation.InvokeType(); ok {
		_spec.SetField(invokemodelorder.FieldInvokeType, field.TypeEnum, value)
	}
	if value, ok := imou.mutation.InvokeTimes(); ok {
		_spec.SetField(invokemodelorder.FieldInvokeTimes, field.TypeInt, value)
	}
	if value, ok := imou.mutation.AddedInvokeTimes(); ok {
		_spec.AddField(invokemodelorder.FieldInvokeTimes, field.TypeInt, value)
	}
	if value, ok := imou.mutation.InputTokenCost(); ok {
		_spec.SetField(invokemodelorder.FieldInputTokenCost, field.TypeInt, value)
	}
	if value, ok := imou.mutation.AddedInputTokenCost(); ok {
		_spec.AddField(invokemodelorder.FieldInputTokenCost, field.TypeInt, value)
	}
	if value, ok := imou.mutation.OutputTokenCost(); ok {
		_spec.SetField(invokemodelorder.FieldOutputTokenCost, field.TypeInt, value)
	}
	if value, ok := imou.mutation.AddedOutputTokenCost(); ok {
		_spec.AddField(invokemodelorder.FieldOutputTokenCost, field.TypeInt, value)
	}
	if value, ok := imou.mutation.InputCepCost(); ok {
		_spec.SetField(invokemodelorder.FieldInputCepCost, field.TypeInt, value)
	}
	if value, ok := imou.mutation.AddedInputCepCost(); ok {
		_spec.AddField(invokemodelorder.FieldInputCepCost, field.TypeInt, value)
	}
	if value, ok := imou.mutation.OutputCepCost(); ok {
		_spec.SetField(invokemodelorder.FieldOutputCepCost, field.TypeInt, value)
	}
	if value, ok := imou.mutation.AddedOutputCepCost(); ok {
		_spec.AddField(invokemodelorder.FieldOutputCepCost, field.TypeInt, value)
	}
	if imou.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   invokemodelorder.BillsTable,
			Columns: []string{invokemodelorder.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imou.mutation.RemovedBillsIDs(); len(nodes) > 0 && !imou.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   invokemodelorder.BillsTable,
			Columns: []string{invokemodelorder.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imou.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   invokemodelorder.BillsTable,
			Columns: []string{invokemodelorder.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if imou.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.UserTable,
			Columns: []string{invokemodelorder.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imou.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.UserTable,
			Columns: []string{invokemodelorder.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if imou.mutation.ModelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.ModelTable,
			Columns: []string{invokemodelorder.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imou.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.ModelTable,
			Columns: []string{invokemodelorder.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if imou.mutation.APITokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.APITokenTable,
			Columns: []string{invokemodelorder.APITokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imou.mutation.APITokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.APITokenTable,
			Columns: []string{invokemodelorder.APITokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(imou.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, imou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invokemodelorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	imou.mutation.done = true
	return n, nil
}

// InvokeModelOrderUpdateOne is the builder for updating a single InvokeModelOrder entity.
type InvokeModelOrderUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *InvokeModelOrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (imouo *InvokeModelOrderUpdateOne) SetCreatedBy(i int64) *InvokeModelOrderUpdateOne {
	imouo.mutation.ResetCreatedBy()
	imouo.mutation.SetCreatedBy(i)
	return imouo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (imouo *InvokeModelOrderUpdateOne) SetNillableCreatedBy(i *int64) *InvokeModelOrderUpdateOne {
	if i != nil {
		imouo.SetCreatedBy(*i)
	}
	return imouo
}

// AddCreatedBy adds i to the "created_by" field.
func (imouo *InvokeModelOrderUpdateOne) AddCreatedBy(i int64) *InvokeModelOrderUpdateOne {
	imouo.mutation.AddCreatedBy(i)
	return imouo
}

// SetUpdatedBy sets the "updated_by" field.
func (imouo *InvokeModelOrderUpdateOne) SetUpdatedBy(i int64) *InvokeModelOrderUpdateOne {
	imouo.mutation.ResetUpdatedBy()
	imouo.mutation.SetUpdatedBy(i)
	return imouo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (imouo *InvokeModelOrderUpdateOne) SetNillableUpdatedBy(i *int64) *InvokeModelOrderUpdateOne {
	if i != nil {
		imouo.SetUpdatedBy(*i)
	}
	return imouo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (imouo *InvokeModelOrderUpdateOne) AddUpdatedBy(i int64) *InvokeModelOrderUpdateOne {
	imouo.mutation.AddUpdatedBy(i)
	return imouo
}

// SetUpdatedAt sets the "updated_at" field.
func (imouo *InvokeModelOrderUpdateOne) SetUpdatedAt(t time.Time) *InvokeModelOrderUpdateOne {
	imouo.mutation.SetUpdatedAt(t)
	return imouo
}

// SetDeletedAt sets the "deleted_at" field.
func (imouo *InvokeModelOrderUpdateOne) SetDeletedAt(t time.Time) *InvokeModelOrderUpdateOne {
	imouo.mutation.SetDeletedAt(t)
	return imouo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (imouo *InvokeModelOrderUpdateOne) SetNillableDeletedAt(t *time.Time) *InvokeModelOrderUpdateOne {
	if t != nil {
		imouo.SetDeletedAt(*t)
	}
	return imouo
}

// SetUserID sets the "user_id" field.
func (imouo *InvokeModelOrderUpdateOne) SetUserID(i int64) *InvokeModelOrderUpdateOne {
	imouo.mutation.SetUserID(i)
	return imouo
}

// SetModelID sets the "model_id" field.
func (imouo *InvokeModelOrderUpdateOne) SetModelID(i int64) *InvokeModelOrderUpdateOne {
	imouo.mutation.SetModelID(i)
	return imouo
}

// SetAPITokenID sets the "api_token_id" field.
func (imouo *InvokeModelOrderUpdateOne) SetAPITokenID(i int64) *InvokeModelOrderUpdateOne {
	imouo.mutation.SetAPITokenID(i)
	return imouo
}

// SetInvokeType sets the "invoke_type" field.
func (imouo *InvokeModelOrderUpdateOne) SetInvokeType(et enums.InvokeType) *InvokeModelOrderUpdateOne {
	imouo.mutation.SetInvokeType(et)
	return imouo
}

// SetNillableInvokeType sets the "invoke_type" field if the given value is not nil.
func (imouo *InvokeModelOrderUpdateOne) SetNillableInvokeType(et *enums.InvokeType) *InvokeModelOrderUpdateOne {
	if et != nil {
		imouo.SetInvokeType(*et)
	}
	return imouo
}

// SetInvokeTimes sets the "invoke_times" field.
func (imouo *InvokeModelOrderUpdateOne) SetInvokeTimes(i int) *InvokeModelOrderUpdateOne {
	imouo.mutation.ResetInvokeTimes()
	imouo.mutation.SetInvokeTimes(i)
	return imouo
}

// SetNillableInvokeTimes sets the "invoke_times" field if the given value is not nil.
func (imouo *InvokeModelOrderUpdateOne) SetNillableInvokeTimes(i *int) *InvokeModelOrderUpdateOne {
	if i != nil {
		imouo.SetInvokeTimes(*i)
	}
	return imouo
}

// AddInvokeTimes adds i to the "invoke_times" field.
func (imouo *InvokeModelOrderUpdateOne) AddInvokeTimes(i int) *InvokeModelOrderUpdateOne {
	imouo.mutation.AddInvokeTimes(i)
	return imouo
}

// SetInputTokenCost sets the "input_token_cost" field.
func (imouo *InvokeModelOrderUpdateOne) SetInputTokenCost(i int) *InvokeModelOrderUpdateOne {
	imouo.mutation.ResetInputTokenCost()
	imouo.mutation.SetInputTokenCost(i)
	return imouo
}

// SetNillableInputTokenCost sets the "input_token_cost" field if the given value is not nil.
func (imouo *InvokeModelOrderUpdateOne) SetNillableInputTokenCost(i *int) *InvokeModelOrderUpdateOne {
	if i != nil {
		imouo.SetInputTokenCost(*i)
	}
	return imouo
}

// AddInputTokenCost adds i to the "input_token_cost" field.
func (imouo *InvokeModelOrderUpdateOne) AddInputTokenCost(i int) *InvokeModelOrderUpdateOne {
	imouo.mutation.AddInputTokenCost(i)
	return imouo
}

// SetOutputTokenCost sets the "output_token_cost" field.
func (imouo *InvokeModelOrderUpdateOne) SetOutputTokenCost(i int) *InvokeModelOrderUpdateOne {
	imouo.mutation.ResetOutputTokenCost()
	imouo.mutation.SetOutputTokenCost(i)
	return imouo
}

// SetNillableOutputTokenCost sets the "output_token_cost" field if the given value is not nil.
func (imouo *InvokeModelOrderUpdateOne) SetNillableOutputTokenCost(i *int) *InvokeModelOrderUpdateOne {
	if i != nil {
		imouo.SetOutputTokenCost(*i)
	}
	return imouo
}

// AddOutputTokenCost adds i to the "output_token_cost" field.
func (imouo *InvokeModelOrderUpdateOne) AddOutputTokenCost(i int) *InvokeModelOrderUpdateOne {
	imouo.mutation.AddOutputTokenCost(i)
	return imouo
}

// SetInputCepCost sets the "input_cep_cost" field.
func (imouo *InvokeModelOrderUpdateOne) SetInputCepCost(i int) *InvokeModelOrderUpdateOne {
	imouo.mutation.ResetInputCepCost()
	imouo.mutation.SetInputCepCost(i)
	return imouo
}

// SetNillableInputCepCost sets the "input_cep_cost" field if the given value is not nil.
func (imouo *InvokeModelOrderUpdateOne) SetNillableInputCepCost(i *int) *InvokeModelOrderUpdateOne {
	if i != nil {
		imouo.SetInputCepCost(*i)
	}
	return imouo
}

// AddInputCepCost adds i to the "input_cep_cost" field.
func (imouo *InvokeModelOrderUpdateOne) AddInputCepCost(i int) *InvokeModelOrderUpdateOne {
	imouo.mutation.AddInputCepCost(i)
	return imouo
}

// SetOutputCepCost sets the "output_cep_cost" field.
func (imouo *InvokeModelOrderUpdateOne) SetOutputCepCost(i int) *InvokeModelOrderUpdateOne {
	imouo.mutation.ResetOutputCepCost()
	imouo.mutation.SetOutputCepCost(i)
	return imouo
}

// SetNillableOutputCepCost sets the "output_cep_cost" field if the given value is not nil.
func (imouo *InvokeModelOrderUpdateOne) SetNillableOutputCepCost(i *int) *InvokeModelOrderUpdateOne {
	if i != nil {
		imouo.SetOutputCepCost(*i)
	}
	return imouo
}

// AddOutputCepCost adds i to the "output_cep_cost" field.
func (imouo *InvokeModelOrderUpdateOne) AddOutputCepCost(i int) *InvokeModelOrderUpdateOne {
	imouo.mutation.AddOutputCepCost(i)
	return imouo
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (imouo *InvokeModelOrderUpdateOne) AddBillIDs(ids ...int64) *InvokeModelOrderUpdateOne {
	imouo.mutation.AddBillIDs(ids...)
	return imouo
}

// AddBills adds the "bills" edges to the Bill entity.
func (imouo *InvokeModelOrderUpdateOne) AddBills(b ...*Bill) *InvokeModelOrderUpdateOne {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return imouo.AddBillIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (imouo *InvokeModelOrderUpdateOne) SetUser(u *User) *InvokeModelOrderUpdateOne {
	return imouo.SetUserID(u.ID)
}

// SetModel sets the "model" edge to the Model entity.
func (imouo *InvokeModelOrderUpdateOne) SetModel(m *Model) *InvokeModelOrderUpdateOne {
	return imouo.SetModelID(m.ID)
}

// SetAPIToken sets the "api_token" edge to the ApiToken entity.
func (imouo *InvokeModelOrderUpdateOne) SetAPIToken(a *ApiToken) *InvokeModelOrderUpdateOne {
	return imouo.SetAPITokenID(a.ID)
}

// Mutation returns the InvokeModelOrderMutation object of the builder.
func (imouo *InvokeModelOrderUpdateOne) Mutation() *InvokeModelOrderMutation {
	return imouo.mutation
}

// ClearBills clears all "bills" edges to the Bill entity.
func (imouo *InvokeModelOrderUpdateOne) ClearBills() *InvokeModelOrderUpdateOne {
	imouo.mutation.ClearBills()
	return imouo
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (imouo *InvokeModelOrderUpdateOne) RemoveBillIDs(ids ...int64) *InvokeModelOrderUpdateOne {
	imouo.mutation.RemoveBillIDs(ids...)
	return imouo
}

// RemoveBills removes "bills" edges to Bill entities.
func (imouo *InvokeModelOrderUpdateOne) RemoveBills(b ...*Bill) *InvokeModelOrderUpdateOne {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return imouo.RemoveBillIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (imouo *InvokeModelOrderUpdateOne) ClearUser() *InvokeModelOrderUpdateOne {
	imouo.mutation.ClearUser()
	return imouo
}

// ClearModel clears the "model" edge to the Model entity.
func (imouo *InvokeModelOrderUpdateOne) ClearModel() *InvokeModelOrderUpdateOne {
	imouo.mutation.ClearModel()
	return imouo
}

// ClearAPIToken clears the "api_token" edge to the ApiToken entity.
func (imouo *InvokeModelOrderUpdateOne) ClearAPIToken() *InvokeModelOrderUpdateOne {
	imouo.mutation.ClearAPIToken()
	return imouo
}

// Where appends a list predicates to the InvokeModelOrderUpdate builder.
func (imouo *InvokeModelOrderUpdateOne) Where(ps ...predicate.InvokeModelOrder) *InvokeModelOrderUpdateOne {
	imouo.mutation.Where(ps...)
	return imouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (imouo *InvokeModelOrderUpdateOne) Select(field string, fields ...string) *InvokeModelOrderUpdateOne {
	imouo.fields = append([]string{field}, fields...)
	return imouo
}

// Save executes the query and returns the updated InvokeModelOrder entity.
func (imouo *InvokeModelOrderUpdateOne) Save(ctx context.Context) (*InvokeModelOrder, error) {
	imouo.defaults()
	return withHooks(ctx, imouo.sqlSave, imouo.mutation, imouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (imouo *InvokeModelOrderUpdateOne) SaveX(ctx context.Context) *InvokeModelOrder {
	node, err := imouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (imouo *InvokeModelOrderUpdateOne) Exec(ctx context.Context) error {
	_, err := imouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (imouo *InvokeModelOrderUpdateOne) ExecX(ctx context.Context) {
	if err := imouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (imouo *InvokeModelOrderUpdateOne) defaults() {
	if _, ok := imouo.mutation.UpdatedAt(); !ok {
		v := invokemodelorder.UpdateDefaultUpdatedAt()
		imouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (imouo *InvokeModelOrderUpdateOne) check() error {
	if v, ok := imouo.mutation.InvokeType(); ok {
		if err := invokemodelorder.InvokeTypeValidator(v); err != nil {
			return &ValidationError{Name: "invoke_type", err: fmt.Errorf(`cep_ent: validator failed for field "InvokeModelOrder.invoke_type": %w`, err)}
		}
	}
	if _, ok := imouo.mutation.UserID(); imouo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "InvokeModelOrder.user"`)
	}
	if _, ok := imouo.mutation.ModelID(); imouo.mutation.ModelCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "InvokeModelOrder.model"`)
	}
	if _, ok := imouo.mutation.APITokenID(); imouo.mutation.APITokenCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "InvokeModelOrder.api_token"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (imouo *InvokeModelOrderUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InvokeModelOrderUpdateOne {
	imouo.modifiers = append(imouo.modifiers, modifiers...)
	return imouo
}

func (imouo *InvokeModelOrderUpdateOne) sqlSave(ctx context.Context) (_node *InvokeModelOrder, err error) {
	if err := imouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(invokemodelorder.Table, invokemodelorder.Columns, sqlgraph.NewFieldSpec(invokemodelorder.FieldID, field.TypeInt64))
	id, ok := imouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "InvokeModelOrder.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := imouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invokemodelorder.FieldID)
		for _, f := range fields {
			if !invokemodelorder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != invokemodelorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := imouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := imouo.mutation.CreatedBy(); ok {
		_spec.SetField(invokemodelorder.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := imouo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(invokemodelorder.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := imouo.mutation.UpdatedBy(); ok {
		_spec.SetField(invokemodelorder.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := imouo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(invokemodelorder.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := imouo.mutation.UpdatedAt(); ok {
		_spec.SetField(invokemodelorder.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := imouo.mutation.DeletedAt(); ok {
		_spec.SetField(invokemodelorder.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := imouo.mutation.InvokeType(); ok {
		_spec.SetField(invokemodelorder.FieldInvokeType, field.TypeEnum, value)
	}
	if value, ok := imouo.mutation.InvokeTimes(); ok {
		_spec.SetField(invokemodelorder.FieldInvokeTimes, field.TypeInt, value)
	}
	if value, ok := imouo.mutation.AddedInvokeTimes(); ok {
		_spec.AddField(invokemodelorder.FieldInvokeTimes, field.TypeInt, value)
	}
	if value, ok := imouo.mutation.InputTokenCost(); ok {
		_spec.SetField(invokemodelorder.FieldInputTokenCost, field.TypeInt, value)
	}
	if value, ok := imouo.mutation.AddedInputTokenCost(); ok {
		_spec.AddField(invokemodelorder.FieldInputTokenCost, field.TypeInt, value)
	}
	if value, ok := imouo.mutation.OutputTokenCost(); ok {
		_spec.SetField(invokemodelorder.FieldOutputTokenCost, field.TypeInt, value)
	}
	if value, ok := imouo.mutation.AddedOutputTokenCost(); ok {
		_spec.AddField(invokemodelorder.FieldOutputTokenCost, field.TypeInt, value)
	}
	if value, ok := imouo.mutation.InputCepCost(); ok {
		_spec.SetField(invokemodelorder.FieldInputCepCost, field.TypeInt, value)
	}
	if value, ok := imouo.mutation.AddedInputCepCost(); ok {
		_spec.AddField(invokemodelorder.FieldInputCepCost, field.TypeInt, value)
	}
	if value, ok := imouo.mutation.OutputCepCost(); ok {
		_spec.SetField(invokemodelorder.FieldOutputCepCost, field.TypeInt, value)
	}
	if value, ok := imouo.mutation.AddedOutputCepCost(); ok {
		_spec.AddField(invokemodelorder.FieldOutputCepCost, field.TypeInt, value)
	}
	if imouo.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   invokemodelorder.BillsTable,
			Columns: []string{invokemodelorder.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imouo.mutation.RemovedBillsIDs(); len(nodes) > 0 && !imouo.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   invokemodelorder.BillsTable,
			Columns: []string{invokemodelorder.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imouo.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   invokemodelorder.BillsTable,
			Columns: []string{invokemodelorder.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if imouo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.UserTable,
			Columns: []string{invokemodelorder.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imouo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.UserTable,
			Columns: []string{invokemodelorder.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if imouo.mutation.ModelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.ModelTable,
			Columns: []string{invokemodelorder.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imouo.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.ModelTable,
			Columns: []string{invokemodelorder.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if imouo.mutation.APITokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.APITokenTable,
			Columns: []string{invokemodelorder.APITokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imouo.mutation.APITokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invokemodelorder.APITokenTable,
			Columns: []string{invokemodelorder.APITokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(imouo.modifiers...)
	_node = &InvokeModelOrder{config: imouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, imouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invokemodelorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	imouo.mutation.done = true
	return _node, nil
}
