// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/collect"
	"cephalon-ent/pkg/cep_ent/predicate"
	"cephalon-ent/pkg/cep_ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CollectUpdate is the builder for updating Collect entities.
type CollectUpdate struct {
	config
	hooks    []Hook
	mutation *CollectMutation
}

// Where appends a list predicates to the CollectUpdate builder.
func (cu *CollectUpdate) Where(ps ...predicate.Collect) *CollectUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedBy sets the "created_by" field.
func (cu *CollectUpdate) SetCreatedBy(i int64) *CollectUpdate {
	cu.mutation.ResetCreatedBy()
	cu.mutation.SetCreatedBy(i)
	return cu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cu *CollectUpdate) SetNillableCreatedBy(i *int64) *CollectUpdate {
	if i != nil {
		cu.SetCreatedBy(*i)
	}
	return cu
}

// AddCreatedBy adds i to the "created_by" field.
func (cu *CollectUpdate) AddCreatedBy(i int64) *CollectUpdate {
	cu.mutation.AddCreatedBy(i)
	return cu
}

// SetUpdatedBy sets the "updated_by" field.
func (cu *CollectUpdate) SetUpdatedBy(i int64) *CollectUpdate {
	cu.mutation.ResetUpdatedBy()
	cu.mutation.SetUpdatedBy(i)
	return cu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cu *CollectUpdate) SetNillableUpdatedBy(i *int64) *CollectUpdate {
	if i != nil {
		cu.SetUpdatedBy(*i)
	}
	return cu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (cu *CollectUpdate) AddUpdatedBy(i int64) *CollectUpdate {
	cu.mutation.AddUpdatedBy(i)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CollectUpdate) SetUpdatedAt(t time.Time) *CollectUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CollectUpdate) SetDeletedAt(t time.Time) *CollectUpdate {
	cu.mutation.SetDeletedAt(t)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CollectUpdate) SetNillableDeletedAt(t *time.Time) *CollectUpdate {
	if t != nil {
		cu.SetDeletedAt(*t)
	}
	return cu
}

// SetURL sets the "url" field.
func (cu *CollectUpdate) SetURL(s string) *CollectUpdate {
	cu.mutation.SetURL(s)
	return cu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (cu *CollectUpdate) SetNillableURL(s *string) *CollectUpdate {
	if s != nil {
		cu.SetURL(*s)
	}
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CollectUpdate) SetUserID(i int64) *CollectUpdate {
	cu.mutation.SetUserID(i)
	return cu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cu *CollectUpdate) SetNillableUserID(i *int64) *CollectUpdate {
	if i != nil {
		cu.SetUserID(*i)
	}
	return cu
}

// SetJpgName sets the "jpg_name" field.
func (cu *CollectUpdate) SetJpgName(i int64) *CollectUpdate {
	cu.mutation.ResetJpgName()
	cu.mutation.SetJpgName(i)
	return cu
}

// SetNillableJpgName sets the "jpg_name" field if the given value is not nil.
func (cu *CollectUpdate) SetNillableJpgName(i *int64) *CollectUpdate {
	if i != nil {
		cu.SetJpgName(*i)
	}
	return cu
}

// AddJpgName adds i to the "jpg_name" field.
func (cu *CollectUpdate) AddJpgName(i int64) *CollectUpdate {
	cu.mutation.AddJpgName(i)
	return cu
}

// SetUser sets the "user" edge to the User entity.
func (cu *CollectUpdate) SetUser(u *User) *CollectUpdate {
	return cu.SetUserID(u.ID)
}

// Mutation returns the CollectMutation object of the builder.
func (cu *CollectUpdate) Mutation() *CollectMutation {
	return cu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cu *CollectUpdate) ClearUser() *CollectUpdate {
	cu.mutation.ClearUser()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CollectUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CollectUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CollectUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CollectUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CollectUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := collect.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CollectUpdate) check() error {
	if _, ok := cu.mutation.UserID(); cu.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "Collect.user"`)
	}
	return nil
}

func (cu *CollectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(collect.Table, collect.Columns, sqlgraph.NewFieldSpec(collect.FieldID, field.TypeInt64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedBy(); ok {
		_spec.SetField(collect.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(collect.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.UpdatedBy(); ok {
		_spec.SetField(collect.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(collect.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(collect.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.SetField(collect.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.URL(); ok {
		_spec.SetField(collect.FieldURL, field.TypeString, value)
	}
	if value, ok := cu.mutation.JpgName(); ok {
		_spec.SetField(collect.FieldJpgName, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedJpgName(); ok {
		_spec.AddField(collect.FieldJpgName, field.TypeInt64, value)
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   collect.UserTable,
			Columns: []string{collect.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   collect.UserTable,
			Columns: []string{collect.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collect.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CollectUpdateOne is the builder for updating a single Collect entity.
type CollectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CollectMutation
}

// SetCreatedBy sets the "created_by" field.
func (cuo *CollectUpdateOne) SetCreatedBy(i int64) *CollectUpdateOne {
	cuo.mutation.ResetCreatedBy()
	cuo.mutation.SetCreatedBy(i)
	return cuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cuo *CollectUpdateOne) SetNillableCreatedBy(i *int64) *CollectUpdateOne {
	if i != nil {
		cuo.SetCreatedBy(*i)
	}
	return cuo
}

// AddCreatedBy adds i to the "created_by" field.
func (cuo *CollectUpdateOne) AddCreatedBy(i int64) *CollectUpdateOne {
	cuo.mutation.AddCreatedBy(i)
	return cuo
}

// SetUpdatedBy sets the "updated_by" field.
func (cuo *CollectUpdateOne) SetUpdatedBy(i int64) *CollectUpdateOne {
	cuo.mutation.ResetUpdatedBy()
	cuo.mutation.SetUpdatedBy(i)
	return cuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cuo *CollectUpdateOne) SetNillableUpdatedBy(i *int64) *CollectUpdateOne {
	if i != nil {
		cuo.SetUpdatedBy(*i)
	}
	return cuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (cuo *CollectUpdateOne) AddUpdatedBy(i int64) *CollectUpdateOne {
	cuo.mutation.AddUpdatedBy(i)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CollectUpdateOne) SetUpdatedAt(t time.Time) *CollectUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CollectUpdateOne) SetDeletedAt(t time.Time) *CollectUpdateOne {
	cuo.mutation.SetDeletedAt(t)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CollectUpdateOne) SetNillableDeletedAt(t *time.Time) *CollectUpdateOne {
	if t != nil {
		cuo.SetDeletedAt(*t)
	}
	return cuo
}

// SetURL sets the "url" field.
func (cuo *CollectUpdateOne) SetURL(s string) *CollectUpdateOne {
	cuo.mutation.SetURL(s)
	return cuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (cuo *CollectUpdateOne) SetNillableURL(s *string) *CollectUpdateOne {
	if s != nil {
		cuo.SetURL(*s)
	}
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *CollectUpdateOne) SetUserID(i int64) *CollectUpdateOne {
	cuo.mutation.SetUserID(i)
	return cuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cuo *CollectUpdateOne) SetNillableUserID(i *int64) *CollectUpdateOne {
	if i != nil {
		cuo.SetUserID(*i)
	}
	return cuo
}

// SetJpgName sets the "jpg_name" field.
func (cuo *CollectUpdateOne) SetJpgName(i int64) *CollectUpdateOne {
	cuo.mutation.ResetJpgName()
	cuo.mutation.SetJpgName(i)
	return cuo
}

// SetNillableJpgName sets the "jpg_name" field if the given value is not nil.
func (cuo *CollectUpdateOne) SetNillableJpgName(i *int64) *CollectUpdateOne {
	if i != nil {
		cuo.SetJpgName(*i)
	}
	return cuo
}

// AddJpgName adds i to the "jpg_name" field.
func (cuo *CollectUpdateOne) AddJpgName(i int64) *CollectUpdateOne {
	cuo.mutation.AddJpgName(i)
	return cuo
}

// SetUser sets the "user" edge to the User entity.
func (cuo *CollectUpdateOne) SetUser(u *User) *CollectUpdateOne {
	return cuo.SetUserID(u.ID)
}

// Mutation returns the CollectMutation object of the builder.
func (cuo *CollectUpdateOne) Mutation() *CollectMutation {
	return cuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cuo *CollectUpdateOne) ClearUser() *CollectUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// Where appends a list predicates to the CollectUpdate builder.
func (cuo *CollectUpdateOne) Where(ps ...predicate.Collect) *CollectUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CollectUpdateOne) Select(field string, fields ...string) *CollectUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Collect entity.
func (cuo *CollectUpdateOne) Save(ctx context.Context) (*Collect, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CollectUpdateOne) SaveX(ctx context.Context) *Collect {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CollectUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CollectUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CollectUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := collect.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CollectUpdateOne) check() error {
	if _, ok := cuo.mutation.UserID(); cuo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "Collect.user"`)
	}
	return nil
}

func (cuo *CollectUpdateOne) sqlSave(ctx context.Context) (_node *Collect, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(collect.Table, collect.Columns, sqlgraph.NewFieldSpec(collect.FieldID, field.TypeInt64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "Collect.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, collect.FieldID)
		for _, f := range fields {
			if !collect.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != collect.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedBy(); ok {
		_spec.SetField(collect.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(collect.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.UpdatedBy(); ok {
		_spec.SetField(collect.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(collect.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(collect.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.SetField(collect.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.URL(); ok {
		_spec.SetField(collect.FieldURL, field.TypeString, value)
	}
	if value, ok := cuo.mutation.JpgName(); ok {
		_spec.SetField(collect.FieldJpgName, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedJpgName(); ok {
		_spec.AddField(collect.FieldJpgName, field.TypeInt64, value)
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   collect.UserTable,
			Columns: []string{collect.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   collect.UserTable,
			Columns: []string{collect.UserColumn},
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
	_node = &Collect{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collect.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}