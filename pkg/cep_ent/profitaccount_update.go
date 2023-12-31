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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/earnbill"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/profitaccount"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// ProfitAccountUpdate is the builder for updating ProfitAccount entities.
type ProfitAccountUpdate struct {
	config
	hooks     []Hook
	mutation  *ProfitAccountMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ProfitAccountUpdate builder.
func (pau *ProfitAccountUpdate) Where(ps ...predicate.ProfitAccount) *ProfitAccountUpdate {
	pau.mutation.Where(ps...)
	return pau
}

// SetCreatedBy sets the "created_by" field.
func (pau *ProfitAccountUpdate) SetCreatedBy(i int64) *ProfitAccountUpdate {
	pau.mutation.ResetCreatedBy()
	pau.mutation.SetCreatedBy(i)
	return pau
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pau *ProfitAccountUpdate) SetNillableCreatedBy(i *int64) *ProfitAccountUpdate {
	if i != nil {
		pau.SetCreatedBy(*i)
	}
	return pau
}

// AddCreatedBy adds i to the "created_by" field.
func (pau *ProfitAccountUpdate) AddCreatedBy(i int64) *ProfitAccountUpdate {
	pau.mutation.AddCreatedBy(i)
	return pau
}

// SetUpdatedBy sets the "updated_by" field.
func (pau *ProfitAccountUpdate) SetUpdatedBy(i int64) *ProfitAccountUpdate {
	pau.mutation.ResetUpdatedBy()
	pau.mutation.SetUpdatedBy(i)
	return pau
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pau *ProfitAccountUpdate) SetNillableUpdatedBy(i *int64) *ProfitAccountUpdate {
	if i != nil {
		pau.SetUpdatedBy(*i)
	}
	return pau
}

// AddUpdatedBy adds i to the "updated_by" field.
func (pau *ProfitAccountUpdate) AddUpdatedBy(i int64) *ProfitAccountUpdate {
	pau.mutation.AddUpdatedBy(i)
	return pau
}

// SetUpdatedAt sets the "updated_at" field.
func (pau *ProfitAccountUpdate) SetUpdatedAt(t time.Time) *ProfitAccountUpdate {
	pau.mutation.SetUpdatedAt(t)
	return pau
}

// SetDeletedAt sets the "deleted_at" field.
func (pau *ProfitAccountUpdate) SetDeletedAt(t time.Time) *ProfitAccountUpdate {
	pau.mutation.SetDeletedAt(t)
	return pau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pau *ProfitAccountUpdate) SetNillableDeletedAt(t *time.Time) *ProfitAccountUpdate {
	if t != nil {
		pau.SetDeletedAt(*t)
	}
	return pau
}

// SetUserID sets the "user_id" field.
func (pau *ProfitAccountUpdate) SetUserID(i int64) *ProfitAccountUpdate {
	pau.mutation.SetUserID(i)
	return pau
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pau *ProfitAccountUpdate) SetNillableUserID(i *int64) *ProfitAccountUpdate {
	if i != nil {
		pau.SetUserID(*i)
	}
	return pau
}

// SetSumCep sets the "sum_cep" field.
func (pau *ProfitAccountUpdate) SetSumCep(i int64) *ProfitAccountUpdate {
	pau.mutation.ResetSumCep()
	pau.mutation.SetSumCep(i)
	return pau
}

// SetNillableSumCep sets the "sum_cep" field if the given value is not nil.
func (pau *ProfitAccountUpdate) SetNillableSumCep(i *int64) *ProfitAccountUpdate {
	if i != nil {
		pau.SetSumCep(*i)
	}
	return pau
}

// AddSumCep adds i to the "sum_cep" field.
func (pau *ProfitAccountUpdate) AddSumCep(i int64) *ProfitAccountUpdate {
	pau.mutation.AddSumCep(i)
	return pau
}

// SetRemainCep sets the "remain_cep" field.
func (pau *ProfitAccountUpdate) SetRemainCep(i int64) *ProfitAccountUpdate {
	pau.mutation.ResetRemainCep()
	pau.mutation.SetRemainCep(i)
	return pau
}

// SetNillableRemainCep sets the "remain_cep" field if the given value is not nil.
func (pau *ProfitAccountUpdate) SetNillableRemainCep(i *int64) *ProfitAccountUpdate {
	if i != nil {
		pau.SetRemainCep(*i)
	}
	return pau
}

// AddRemainCep adds i to the "remain_cep" field.
func (pau *ProfitAccountUpdate) AddRemainCep(i int64) *ProfitAccountUpdate {
	pau.mutation.AddRemainCep(i)
	return pau
}

// SetUser sets the "user" edge to the User entity.
func (pau *ProfitAccountUpdate) SetUser(u *User) *ProfitAccountUpdate {
	return pau.SetUserID(u.ID)
}

// AddEarnBillIDs adds the "earn_bills" edge to the EarnBill entity by IDs.
func (pau *ProfitAccountUpdate) AddEarnBillIDs(ids ...int64) *ProfitAccountUpdate {
	pau.mutation.AddEarnBillIDs(ids...)
	return pau
}

// AddEarnBills adds the "earn_bills" edges to the EarnBill entity.
func (pau *ProfitAccountUpdate) AddEarnBills(e ...*EarnBill) *ProfitAccountUpdate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pau.AddEarnBillIDs(ids...)
}

// Mutation returns the ProfitAccountMutation object of the builder.
func (pau *ProfitAccountUpdate) Mutation() *ProfitAccountMutation {
	return pau.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pau *ProfitAccountUpdate) ClearUser() *ProfitAccountUpdate {
	pau.mutation.ClearUser()
	return pau
}

// ClearEarnBills clears all "earn_bills" edges to the EarnBill entity.
func (pau *ProfitAccountUpdate) ClearEarnBills() *ProfitAccountUpdate {
	pau.mutation.ClearEarnBills()
	return pau
}

// RemoveEarnBillIDs removes the "earn_bills" edge to EarnBill entities by IDs.
func (pau *ProfitAccountUpdate) RemoveEarnBillIDs(ids ...int64) *ProfitAccountUpdate {
	pau.mutation.RemoveEarnBillIDs(ids...)
	return pau
}

// RemoveEarnBills removes "earn_bills" edges to EarnBill entities.
func (pau *ProfitAccountUpdate) RemoveEarnBills(e ...*EarnBill) *ProfitAccountUpdate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pau.RemoveEarnBillIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pau *ProfitAccountUpdate) Save(ctx context.Context) (int, error) {
	pau.defaults()
	return withHooks(ctx, pau.sqlSave, pau.mutation, pau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pau *ProfitAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := pau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pau *ProfitAccountUpdate) Exec(ctx context.Context) error {
	_, err := pau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pau *ProfitAccountUpdate) ExecX(ctx context.Context) {
	if err := pau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pau *ProfitAccountUpdate) defaults() {
	if _, ok := pau.mutation.UpdatedAt(); !ok {
		v := profitaccount.UpdateDefaultUpdatedAt()
		pau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pau *ProfitAccountUpdate) check() error {
	if _, ok := pau.mutation.UserID(); pau.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ProfitAccount.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pau *ProfitAccountUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProfitAccountUpdate {
	pau.modifiers = append(pau.modifiers, modifiers...)
	return pau
}

func (pau *ProfitAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(profitaccount.Table, profitaccount.Columns, sqlgraph.NewFieldSpec(profitaccount.FieldID, field.TypeInt64))
	if ps := pau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pau.mutation.CreatedBy(); ok {
		_spec.SetField(profitaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := pau.mutation.AddedCreatedBy(); ok {
		_spec.AddField(profitaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := pau.mutation.UpdatedBy(); ok {
		_spec.SetField(profitaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := pau.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(profitaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := pau.mutation.UpdatedAt(); ok {
		_spec.SetField(profitaccount.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pau.mutation.DeletedAt(); ok {
		_spec.SetField(profitaccount.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := pau.mutation.SumCep(); ok {
		_spec.SetField(profitaccount.FieldSumCep, field.TypeInt64, value)
	}
	if value, ok := pau.mutation.AddedSumCep(); ok {
		_spec.AddField(profitaccount.FieldSumCep, field.TypeInt64, value)
	}
	if value, ok := pau.mutation.RemainCep(); ok {
		_spec.SetField(profitaccount.FieldRemainCep, field.TypeInt64, value)
	}
	if value, ok := pau.mutation.AddedRemainCep(); ok {
		_spec.AddField(profitaccount.FieldRemainCep, field.TypeInt64, value)
	}
	if pau.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profitaccount.UserTable,
			Columns: []string{profitaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profitaccount.UserTable,
			Columns: []string{profitaccount.UserColumn},
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
	if pau.mutation.EarnBillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profitaccount.EarnBillsTable,
			Columns: []string{profitaccount.EarnBillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earnbill.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.RemovedEarnBillsIDs(); len(nodes) > 0 && !pau.mutation.EarnBillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profitaccount.EarnBillsTable,
			Columns: []string{profitaccount.EarnBillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earnbill.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.EarnBillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profitaccount.EarnBillsTable,
			Columns: []string{profitaccount.EarnBillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earnbill.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pau.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profitaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pau.mutation.done = true
	return n, nil
}

// ProfitAccountUpdateOne is the builder for updating a single ProfitAccount entity.
type ProfitAccountUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ProfitAccountMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (pauo *ProfitAccountUpdateOne) SetCreatedBy(i int64) *ProfitAccountUpdateOne {
	pauo.mutation.ResetCreatedBy()
	pauo.mutation.SetCreatedBy(i)
	return pauo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pauo *ProfitAccountUpdateOne) SetNillableCreatedBy(i *int64) *ProfitAccountUpdateOne {
	if i != nil {
		pauo.SetCreatedBy(*i)
	}
	return pauo
}

// AddCreatedBy adds i to the "created_by" field.
func (pauo *ProfitAccountUpdateOne) AddCreatedBy(i int64) *ProfitAccountUpdateOne {
	pauo.mutation.AddCreatedBy(i)
	return pauo
}

// SetUpdatedBy sets the "updated_by" field.
func (pauo *ProfitAccountUpdateOne) SetUpdatedBy(i int64) *ProfitAccountUpdateOne {
	pauo.mutation.ResetUpdatedBy()
	pauo.mutation.SetUpdatedBy(i)
	return pauo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pauo *ProfitAccountUpdateOne) SetNillableUpdatedBy(i *int64) *ProfitAccountUpdateOne {
	if i != nil {
		pauo.SetUpdatedBy(*i)
	}
	return pauo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (pauo *ProfitAccountUpdateOne) AddUpdatedBy(i int64) *ProfitAccountUpdateOne {
	pauo.mutation.AddUpdatedBy(i)
	return pauo
}

// SetUpdatedAt sets the "updated_at" field.
func (pauo *ProfitAccountUpdateOne) SetUpdatedAt(t time.Time) *ProfitAccountUpdateOne {
	pauo.mutation.SetUpdatedAt(t)
	return pauo
}

// SetDeletedAt sets the "deleted_at" field.
func (pauo *ProfitAccountUpdateOne) SetDeletedAt(t time.Time) *ProfitAccountUpdateOne {
	pauo.mutation.SetDeletedAt(t)
	return pauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pauo *ProfitAccountUpdateOne) SetNillableDeletedAt(t *time.Time) *ProfitAccountUpdateOne {
	if t != nil {
		pauo.SetDeletedAt(*t)
	}
	return pauo
}

// SetUserID sets the "user_id" field.
func (pauo *ProfitAccountUpdateOne) SetUserID(i int64) *ProfitAccountUpdateOne {
	pauo.mutation.SetUserID(i)
	return pauo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pauo *ProfitAccountUpdateOne) SetNillableUserID(i *int64) *ProfitAccountUpdateOne {
	if i != nil {
		pauo.SetUserID(*i)
	}
	return pauo
}

// SetSumCep sets the "sum_cep" field.
func (pauo *ProfitAccountUpdateOne) SetSumCep(i int64) *ProfitAccountUpdateOne {
	pauo.mutation.ResetSumCep()
	pauo.mutation.SetSumCep(i)
	return pauo
}

// SetNillableSumCep sets the "sum_cep" field if the given value is not nil.
func (pauo *ProfitAccountUpdateOne) SetNillableSumCep(i *int64) *ProfitAccountUpdateOne {
	if i != nil {
		pauo.SetSumCep(*i)
	}
	return pauo
}

// AddSumCep adds i to the "sum_cep" field.
func (pauo *ProfitAccountUpdateOne) AddSumCep(i int64) *ProfitAccountUpdateOne {
	pauo.mutation.AddSumCep(i)
	return pauo
}

// SetRemainCep sets the "remain_cep" field.
func (pauo *ProfitAccountUpdateOne) SetRemainCep(i int64) *ProfitAccountUpdateOne {
	pauo.mutation.ResetRemainCep()
	pauo.mutation.SetRemainCep(i)
	return pauo
}

// SetNillableRemainCep sets the "remain_cep" field if the given value is not nil.
func (pauo *ProfitAccountUpdateOne) SetNillableRemainCep(i *int64) *ProfitAccountUpdateOne {
	if i != nil {
		pauo.SetRemainCep(*i)
	}
	return pauo
}

// AddRemainCep adds i to the "remain_cep" field.
func (pauo *ProfitAccountUpdateOne) AddRemainCep(i int64) *ProfitAccountUpdateOne {
	pauo.mutation.AddRemainCep(i)
	return pauo
}

// SetUser sets the "user" edge to the User entity.
func (pauo *ProfitAccountUpdateOne) SetUser(u *User) *ProfitAccountUpdateOne {
	return pauo.SetUserID(u.ID)
}

// AddEarnBillIDs adds the "earn_bills" edge to the EarnBill entity by IDs.
func (pauo *ProfitAccountUpdateOne) AddEarnBillIDs(ids ...int64) *ProfitAccountUpdateOne {
	pauo.mutation.AddEarnBillIDs(ids...)
	return pauo
}

// AddEarnBills adds the "earn_bills" edges to the EarnBill entity.
func (pauo *ProfitAccountUpdateOne) AddEarnBills(e ...*EarnBill) *ProfitAccountUpdateOne {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pauo.AddEarnBillIDs(ids...)
}

// Mutation returns the ProfitAccountMutation object of the builder.
func (pauo *ProfitAccountUpdateOne) Mutation() *ProfitAccountMutation {
	return pauo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pauo *ProfitAccountUpdateOne) ClearUser() *ProfitAccountUpdateOne {
	pauo.mutation.ClearUser()
	return pauo
}

// ClearEarnBills clears all "earn_bills" edges to the EarnBill entity.
func (pauo *ProfitAccountUpdateOne) ClearEarnBills() *ProfitAccountUpdateOne {
	pauo.mutation.ClearEarnBills()
	return pauo
}

// RemoveEarnBillIDs removes the "earn_bills" edge to EarnBill entities by IDs.
func (pauo *ProfitAccountUpdateOne) RemoveEarnBillIDs(ids ...int64) *ProfitAccountUpdateOne {
	pauo.mutation.RemoveEarnBillIDs(ids...)
	return pauo
}

// RemoveEarnBills removes "earn_bills" edges to EarnBill entities.
func (pauo *ProfitAccountUpdateOne) RemoveEarnBills(e ...*EarnBill) *ProfitAccountUpdateOne {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pauo.RemoveEarnBillIDs(ids...)
}

// Where appends a list predicates to the ProfitAccountUpdate builder.
func (pauo *ProfitAccountUpdateOne) Where(ps ...predicate.ProfitAccount) *ProfitAccountUpdateOne {
	pauo.mutation.Where(ps...)
	return pauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pauo *ProfitAccountUpdateOne) Select(field string, fields ...string) *ProfitAccountUpdateOne {
	pauo.fields = append([]string{field}, fields...)
	return pauo
}

// Save executes the query and returns the updated ProfitAccount entity.
func (pauo *ProfitAccountUpdateOne) Save(ctx context.Context) (*ProfitAccount, error) {
	pauo.defaults()
	return withHooks(ctx, pauo.sqlSave, pauo.mutation, pauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pauo *ProfitAccountUpdateOne) SaveX(ctx context.Context) *ProfitAccount {
	node, err := pauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pauo *ProfitAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := pauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pauo *ProfitAccountUpdateOne) ExecX(ctx context.Context) {
	if err := pauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pauo *ProfitAccountUpdateOne) defaults() {
	if _, ok := pauo.mutation.UpdatedAt(); !ok {
		v := profitaccount.UpdateDefaultUpdatedAt()
		pauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pauo *ProfitAccountUpdateOne) check() error {
	if _, ok := pauo.mutation.UserID(); pauo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ProfitAccount.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pauo *ProfitAccountUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProfitAccountUpdateOne {
	pauo.modifiers = append(pauo.modifiers, modifiers...)
	return pauo
}

func (pauo *ProfitAccountUpdateOne) sqlSave(ctx context.Context) (_node *ProfitAccount, err error) {
	if err := pauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(profitaccount.Table, profitaccount.Columns, sqlgraph.NewFieldSpec(profitaccount.FieldID, field.TypeInt64))
	id, ok := pauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "ProfitAccount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profitaccount.FieldID)
		for _, f := range fields {
			if !profitaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != profitaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pauo.mutation.CreatedBy(); ok {
		_spec.SetField(profitaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := pauo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(profitaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := pauo.mutation.UpdatedBy(); ok {
		_spec.SetField(profitaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := pauo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(profitaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := pauo.mutation.UpdatedAt(); ok {
		_spec.SetField(profitaccount.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pauo.mutation.DeletedAt(); ok {
		_spec.SetField(profitaccount.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := pauo.mutation.SumCep(); ok {
		_spec.SetField(profitaccount.FieldSumCep, field.TypeInt64, value)
	}
	if value, ok := pauo.mutation.AddedSumCep(); ok {
		_spec.AddField(profitaccount.FieldSumCep, field.TypeInt64, value)
	}
	if value, ok := pauo.mutation.RemainCep(); ok {
		_spec.SetField(profitaccount.FieldRemainCep, field.TypeInt64, value)
	}
	if value, ok := pauo.mutation.AddedRemainCep(); ok {
		_spec.AddField(profitaccount.FieldRemainCep, field.TypeInt64, value)
	}
	if pauo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profitaccount.UserTable,
			Columns: []string{profitaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profitaccount.UserTable,
			Columns: []string{profitaccount.UserColumn},
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
	if pauo.mutation.EarnBillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profitaccount.EarnBillsTable,
			Columns: []string{profitaccount.EarnBillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earnbill.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.RemovedEarnBillsIDs(); len(nodes) > 0 && !pauo.mutation.EarnBillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profitaccount.EarnBillsTable,
			Columns: []string{profitaccount.EarnBillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earnbill.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.EarnBillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profitaccount.EarnBillsTable,
			Columns: []string{profitaccount.EarnBillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earnbill.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pauo.modifiers...)
	_node = &ProfitAccount{config: pauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profitaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pauo.mutation.done = true
	return _node, nil
}
