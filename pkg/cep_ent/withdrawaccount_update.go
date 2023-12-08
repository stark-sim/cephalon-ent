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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/withdrawaccount"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// WithdrawAccountUpdate is the builder for updating WithdrawAccount entities.
type WithdrawAccountUpdate struct {
	config
	hooks     []Hook
	mutation  *WithdrawAccountMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the WithdrawAccountUpdate builder.
func (wau *WithdrawAccountUpdate) Where(ps ...predicate.WithdrawAccount) *WithdrawAccountUpdate {
	wau.mutation.Where(ps...)
	return wau
}

// SetCreatedBy sets the "created_by" field.
func (wau *WithdrawAccountUpdate) SetCreatedBy(i int64) *WithdrawAccountUpdate {
	wau.mutation.ResetCreatedBy()
	wau.mutation.SetCreatedBy(i)
	return wau
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (wau *WithdrawAccountUpdate) SetNillableCreatedBy(i *int64) *WithdrawAccountUpdate {
	if i != nil {
		wau.SetCreatedBy(*i)
	}
	return wau
}

// AddCreatedBy adds i to the "created_by" field.
func (wau *WithdrawAccountUpdate) AddCreatedBy(i int64) *WithdrawAccountUpdate {
	wau.mutation.AddCreatedBy(i)
	return wau
}

// SetUpdatedBy sets the "updated_by" field.
func (wau *WithdrawAccountUpdate) SetUpdatedBy(i int64) *WithdrawAccountUpdate {
	wau.mutation.ResetUpdatedBy()
	wau.mutation.SetUpdatedBy(i)
	return wau
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (wau *WithdrawAccountUpdate) SetNillableUpdatedBy(i *int64) *WithdrawAccountUpdate {
	if i != nil {
		wau.SetUpdatedBy(*i)
	}
	return wau
}

// AddUpdatedBy adds i to the "updated_by" field.
func (wau *WithdrawAccountUpdate) AddUpdatedBy(i int64) *WithdrawAccountUpdate {
	wau.mutation.AddUpdatedBy(i)
	return wau
}

// SetUpdatedAt sets the "updated_at" field.
func (wau *WithdrawAccountUpdate) SetUpdatedAt(t time.Time) *WithdrawAccountUpdate {
	wau.mutation.SetUpdatedAt(t)
	return wau
}

// SetDeletedAt sets the "deleted_at" field.
func (wau *WithdrawAccountUpdate) SetDeletedAt(t time.Time) *WithdrawAccountUpdate {
	wau.mutation.SetDeletedAt(t)
	return wau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wau *WithdrawAccountUpdate) SetNillableDeletedAt(t *time.Time) *WithdrawAccountUpdate {
	if t != nil {
		wau.SetDeletedAt(*t)
	}
	return wau
}

// SetUserID sets the "user_id" field.
func (wau *WithdrawAccountUpdate) SetUserID(i int64) *WithdrawAccountUpdate {
	wau.mutation.SetUserID(i)
	return wau
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wau *WithdrawAccountUpdate) SetNillableUserID(i *int64) *WithdrawAccountUpdate {
	if i != nil {
		wau.SetUserID(*i)
	}
	return wau
}

// SetBusinessName sets the "business_name" field.
func (wau *WithdrawAccountUpdate) SetBusinessName(s string) *WithdrawAccountUpdate {
	wau.mutation.SetBusinessName(s)
	return wau
}

// SetNillableBusinessName sets the "business_name" field if the given value is not nil.
func (wau *WithdrawAccountUpdate) SetNillableBusinessName(s *string) *WithdrawAccountUpdate {
	if s != nil {
		wau.SetBusinessName(*s)
	}
	return wau
}

// SetBusinessType sets the "business_type" field.
func (wau *WithdrawAccountUpdate) SetBusinessType(et enums.BusinessType) *WithdrawAccountUpdate {
	wau.mutation.SetBusinessType(et)
	return wau
}

// SetNillableBusinessType sets the "business_type" field if the given value is not nil.
func (wau *WithdrawAccountUpdate) SetNillableBusinessType(et *enums.BusinessType) *WithdrawAccountUpdate {
	if et != nil {
		wau.SetBusinessType(*et)
	}
	return wau
}

// SetBusinessID sets the "business_id" field.
func (wau *WithdrawAccountUpdate) SetBusinessID(i int64) *WithdrawAccountUpdate {
	wau.mutation.ResetBusinessID()
	wau.mutation.SetBusinessID(i)
	return wau
}

// SetNillableBusinessID sets the "business_id" field if the given value is not nil.
func (wau *WithdrawAccountUpdate) SetNillableBusinessID(i *int64) *WithdrawAccountUpdate {
	if i != nil {
		wau.SetBusinessID(*i)
	}
	return wau
}

// AddBusinessID adds i to the "business_id" field.
func (wau *WithdrawAccountUpdate) AddBusinessID(i int64) *WithdrawAccountUpdate {
	wau.mutation.AddBusinessID(i)
	return wau
}

// SetUser sets the "user" edge to the User entity.
func (wau *WithdrawAccountUpdate) SetUser(u *User) *WithdrawAccountUpdate {
	return wau.SetUserID(u.ID)
}

// Mutation returns the WithdrawAccountMutation object of the builder.
func (wau *WithdrawAccountUpdate) Mutation() *WithdrawAccountMutation {
	return wau.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wau *WithdrawAccountUpdate) ClearUser() *WithdrawAccountUpdate {
	wau.mutation.ClearUser()
	return wau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wau *WithdrawAccountUpdate) Save(ctx context.Context) (int, error) {
	wau.defaults()
	return withHooks(ctx, wau.sqlSave, wau.mutation, wau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wau *WithdrawAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := wau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wau *WithdrawAccountUpdate) Exec(ctx context.Context) error {
	_, err := wau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wau *WithdrawAccountUpdate) ExecX(ctx context.Context) {
	if err := wau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wau *WithdrawAccountUpdate) defaults() {
	if _, ok := wau.mutation.UpdatedAt(); !ok {
		v := withdrawaccount.UpdateDefaultUpdatedAt()
		wau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wau *WithdrawAccountUpdate) check() error {
	if v, ok := wau.mutation.BusinessType(); ok {
		if err := withdrawaccount.BusinessTypeValidator(v); err != nil {
			return &ValidationError{Name: "business_type", err: fmt.Errorf(`cep_ent: validator failed for field "WithdrawAccount.business_type": %w`, err)}
		}
	}
	if _, ok := wau.mutation.UserID(); wau.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "WithdrawAccount.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wau *WithdrawAccountUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WithdrawAccountUpdate {
	wau.modifiers = append(wau.modifiers, modifiers...)
	return wau
}

func (wau *WithdrawAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(withdrawaccount.Table, withdrawaccount.Columns, sqlgraph.NewFieldSpec(withdrawaccount.FieldID, field.TypeInt64))
	if ps := wau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wau.mutation.CreatedBy(); ok {
		_spec.SetField(withdrawaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := wau.mutation.AddedCreatedBy(); ok {
		_spec.AddField(withdrawaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := wau.mutation.UpdatedBy(); ok {
		_spec.SetField(withdrawaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := wau.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(withdrawaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := wau.mutation.UpdatedAt(); ok {
		_spec.SetField(withdrawaccount.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wau.mutation.DeletedAt(); ok {
		_spec.SetField(withdrawaccount.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := wau.mutation.BusinessName(); ok {
		_spec.SetField(withdrawaccount.FieldBusinessName, field.TypeString, value)
	}
	if value, ok := wau.mutation.BusinessType(); ok {
		_spec.SetField(withdrawaccount.FieldBusinessType, field.TypeEnum, value)
	}
	if value, ok := wau.mutation.BusinessID(); ok {
		_spec.SetField(withdrawaccount.FieldBusinessID, field.TypeInt64, value)
	}
	if value, ok := wau.mutation.AddedBusinessID(); ok {
		_spec.AddField(withdrawaccount.FieldBusinessID, field.TypeInt64, value)
	}
	if wau.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   withdrawaccount.UserTable,
			Columns: []string{withdrawaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wau.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   withdrawaccount.UserTable,
			Columns: []string{withdrawaccount.UserColumn},
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
	_spec.AddModifiers(wau.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, wau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withdrawaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wau.mutation.done = true
	return n, nil
}

// WithdrawAccountUpdateOne is the builder for updating a single WithdrawAccount entity.
type WithdrawAccountUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *WithdrawAccountMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (wauo *WithdrawAccountUpdateOne) SetCreatedBy(i int64) *WithdrawAccountUpdateOne {
	wauo.mutation.ResetCreatedBy()
	wauo.mutation.SetCreatedBy(i)
	return wauo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (wauo *WithdrawAccountUpdateOne) SetNillableCreatedBy(i *int64) *WithdrawAccountUpdateOne {
	if i != nil {
		wauo.SetCreatedBy(*i)
	}
	return wauo
}

// AddCreatedBy adds i to the "created_by" field.
func (wauo *WithdrawAccountUpdateOne) AddCreatedBy(i int64) *WithdrawAccountUpdateOne {
	wauo.mutation.AddCreatedBy(i)
	return wauo
}

// SetUpdatedBy sets the "updated_by" field.
func (wauo *WithdrawAccountUpdateOne) SetUpdatedBy(i int64) *WithdrawAccountUpdateOne {
	wauo.mutation.ResetUpdatedBy()
	wauo.mutation.SetUpdatedBy(i)
	return wauo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (wauo *WithdrawAccountUpdateOne) SetNillableUpdatedBy(i *int64) *WithdrawAccountUpdateOne {
	if i != nil {
		wauo.SetUpdatedBy(*i)
	}
	return wauo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (wauo *WithdrawAccountUpdateOne) AddUpdatedBy(i int64) *WithdrawAccountUpdateOne {
	wauo.mutation.AddUpdatedBy(i)
	return wauo
}

// SetUpdatedAt sets the "updated_at" field.
func (wauo *WithdrawAccountUpdateOne) SetUpdatedAt(t time.Time) *WithdrawAccountUpdateOne {
	wauo.mutation.SetUpdatedAt(t)
	return wauo
}

// SetDeletedAt sets the "deleted_at" field.
func (wauo *WithdrawAccountUpdateOne) SetDeletedAt(t time.Time) *WithdrawAccountUpdateOne {
	wauo.mutation.SetDeletedAt(t)
	return wauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wauo *WithdrawAccountUpdateOne) SetNillableDeletedAt(t *time.Time) *WithdrawAccountUpdateOne {
	if t != nil {
		wauo.SetDeletedAt(*t)
	}
	return wauo
}

// SetUserID sets the "user_id" field.
func (wauo *WithdrawAccountUpdateOne) SetUserID(i int64) *WithdrawAccountUpdateOne {
	wauo.mutation.SetUserID(i)
	return wauo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wauo *WithdrawAccountUpdateOne) SetNillableUserID(i *int64) *WithdrawAccountUpdateOne {
	if i != nil {
		wauo.SetUserID(*i)
	}
	return wauo
}

// SetBusinessName sets the "business_name" field.
func (wauo *WithdrawAccountUpdateOne) SetBusinessName(s string) *WithdrawAccountUpdateOne {
	wauo.mutation.SetBusinessName(s)
	return wauo
}

// SetNillableBusinessName sets the "business_name" field if the given value is not nil.
func (wauo *WithdrawAccountUpdateOne) SetNillableBusinessName(s *string) *WithdrawAccountUpdateOne {
	if s != nil {
		wauo.SetBusinessName(*s)
	}
	return wauo
}

// SetBusinessType sets the "business_type" field.
func (wauo *WithdrawAccountUpdateOne) SetBusinessType(et enums.BusinessType) *WithdrawAccountUpdateOne {
	wauo.mutation.SetBusinessType(et)
	return wauo
}

// SetNillableBusinessType sets the "business_type" field if the given value is not nil.
func (wauo *WithdrawAccountUpdateOne) SetNillableBusinessType(et *enums.BusinessType) *WithdrawAccountUpdateOne {
	if et != nil {
		wauo.SetBusinessType(*et)
	}
	return wauo
}

// SetBusinessID sets the "business_id" field.
func (wauo *WithdrawAccountUpdateOne) SetBusinessID(i int64) *WithdrawAccountUpdateOne {
	wauo.mutation.ResetBusinessID()
	wauo.mutation.SetBusinessID(i)
	return wauo
}

// SetNillableBusinessID sets the "business_id" field if the given value is not nil.
func (wauo *WithdrawAccountUpdateOne) SetNillableBusinessID(i *int64) *WithdrawAccountUpdateOne {
	if i != nil {
		wauo.SetBusinessID(*i)
	}
	return wauo
}

// AddBusinessID adds i to the "business_id" field.
func (wauo *WithdrawAccountUpdateOne) AddBusinessID(i int64) *WithdrawAccountUpdateOne {
	wauo.mutation.AddBusinessID(i)
	return wauo
}

// SetUser sets the "user" edge to the User entity.
func (wauo *WithdrawAccountUpdateOne) SetUser(u *User) *WithdrawAccountUpdateOne {
	return wauo.SetUserID(u.ID)
}

// Mutation returns the WithdrawAccountMutation object of the builder.
func (wauo *WithdrawAccountUpdateOne) Mutation() *WithdrawAccountMutation {
	return wauo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wauo *WithdrawAccountUpdateOne) ClearUser() *WithdrawAccountUpdateOne {
	wauo.mutation.ClearUser()
	return wauo
}

// Where appends a list predicates to the WithdrawAccountUpdate builder.
func (wauo *WithdrawAccountUpdateOne) Where(ps ...predicate.WithdrawAccount) *WithdrawAccountUpdateOne {
	wauo.mutation.Where(ps...)
	return wauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wauo *WithdrawAccountUpdateOne) Select(field string, fields ...string) *WithdrawAccountUpdateOne {
	wauo.fields = append([]string{field}, fields...)
	return wauo
}

// Save executes the query and returns the updated WithdrawAccount entity.
func (wauo *WithdrawAccountUpdateOne) Save(ctx context.Context) (*WithdrawAccount, error) {
	wauo.defaults()
	return withHooks(ctx, wauo.sqlSave, wauo.mutation, wauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wauo *WithdrawAccountUpdateOne) SaveX(ctx context.Context) *WithdrawAccount {
	node, err := wauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wauo *WithdrawAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := wauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wauo *WithdrawAccountUpdateOne) ExecX(ctx context.Context) {
	if err := wauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wauo *WithdrawAccountUpdateOne) defaults() {
	if _, ok := wauo.mutation.UpdatedAt(); !ok {
		v := withdrawaccount.UpdateDefaultUpdatedAt()
		wauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wauo *WithdrawAccountUpdateOne) check() error {
	if v, ok := wauo.mutation.BusinessType(); ok {
		if err := withdrawaccount.BusinessTypeValidator(v); err != nil {
			return &ValidationError{Name: "business_type", err: fmt.Errorf(`cep_ent: validator failed for field "WithdrawAccount.business_type": %w`, err)}
		}
	}
	if _, ok := wauo.mutation.UserID(); wauo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "WithdrawAccount.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wauo *WithdrawAccountUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WithdrawAccountUpdateOne {
	wauo.modifiers = append(wauo.modifiers, modifiers...)
	return wauo
}

func (wauo *WithdrawAccountUpdateOne) sqlSave(ctx context.Context) (_node *WithdrawAccount, err error) {
	if err := wauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(withdrawaccount.Table, withdrawaccount.Columns, sqlgraph.NewFieldSpec(withdrawaccount.FieldID, field.TypeInt64))
	id, ok := wauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "WithdrawAccount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withdrawaccount.FieldID)
		for _, f := range fields {
			if !withdrawaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != withdrawaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wauo.mutation.CreatedBy(); ok {
		_spec.SetField(withdrawaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := wauo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(withdrawaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := wauo.mutation.UpdatedBy(); ok {
		_spec.SetField(withdrawaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := wauo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(withdrawaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := wauo.mutation.UpdatedAt(); ok {
		_spec.SetField(withdrawaccount.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wauo.mutation.DeletedAt(); ok {
		_spec.SetField(withdrawaccount.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := wauo.mutation.BusinessName(); ok {
		_spec.SetField(withdrawaccount.FieldBusinessName, field.TypeString, value)
	}
	if value, ok := wauo.mutation.BusinessType(); ok {
		_spec.SetField(withdrawaccount.FieldBusinessType, field.TypeEnum, value)
	}
	if value, ok := wauo.mutation.BusinessID(); ok {
		_spec.SetField(withdrawaccount.FieldBusinessID, field.TypeInt64, value)
	}
	if value, ok := wauo.mutation.AddedBusinessID(); ok {
		_spec.AddField(withdrawaccount.FieldBusinessID, field.TypeInt64, value)
	}
	if wauo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   withdrawaccount.UserTable,
			Columns: []string{withdrawaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wauo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   withdrawaccount.UserTable,
			Columns: []string{withdrawaccount.UserColumn},
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
	_spec.AddModifiers(wauo.modifiers...)
	_node = &WithdrawAccount{config: wauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withdrawaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wauo.mutation.done = true
	return _node, nil
}
