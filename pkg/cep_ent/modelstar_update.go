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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/model"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/modelstar"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ModelStarUpdate is the builder for updating ModelStar entities.
type ModelStarUpdate struct {
	config
	hooks     []Hook
	mutation  *ModelStarMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ModelStarUpdate builder.
func (msu *ModelStarUpdate) Where(ps ...predicate.ModelStar) *ModelStarUpdate {
	msu.mutation.Where(ps...)
	return msu
}

// SetCreatedBy sets the "created_by" field.
func (msu *ModelStarUpdate) SetCreatedBy(i int64) *ModelStarUpdate {
	msu.mutation.ResetCreatedBy()
	msu.mutation.SetCreatedBy(i)
	return msu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (msu *ModelStarUpdate) SetNillableCreatedBy(i *int64) *ModelStarUpdate {
	if i != nil {
		msu.SetCreatedBy(*i)
	}
	return msu
}

// AddCreatedBy adds i to the "created_by" field.
func (msu *ModelStarUpdate) AddCreatedBy(i int64) *ModelStarUpdate {
	msu.mutation.AddCreatedBy(i)
	return msu
}

// SetUpdatedBy sets the "updated_by" field.
func (msu *ModelStarUpdate) SetUpdatedBy(i int64) *ModelStarUpdate {
	msu.mutation.ResetUpdatedBy()
	msu.mutation.SetUpdatedBy(i)
	return msu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (msu *ModelStarUpdate) SetNillableUpdatedBy(i *int64) *ModelStarUpdate {
	if i != nil {
		msu.SetUpdatedBy(*i)
	}
	return msu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (msu *ModelStarUpdate) AddUpdatedBy(i int64) *ModelStarUpdate {
	msu.mutation.AddUpdatedBy(i)
	return msu
}

// SetUpdatedAt sets the "updated_at" field.
func (msu *ModelStarUpdate) SetUpdatedAt(t time.Time) *ModelStarUpdate {
	msu.mutation.SetUpdatedAt(t)
	return msu
}

// SetDeletedAt sets the "deleted_at" field.
func (msu *ModelStarUpdate) SetDeletedAt(t time.Time) *ModelStarUpdate {
	msu.mutation.SetDeletedAt(t)
	return msu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (msu *ModelStarUpdate) SetNillableDeletedAt(t *time.Time) *ModelStarUpdate {
	if t != nil {
		msu.SetDeletedAt(*t)
	}
	return msu
}

// SetUserID sets the "user_id" field.
func (msu *ModelStarUpdate) SetUserID(i int64) *ModelStarUpdate {
	msu.mutation.SetUserID(i)
	return msu
}

// SetModelID sets the "model_id" field.
func (msu *ModelStarUpdate) SetModelID(i int64) *ModelStarUpdate {
	msu.mutation.SetModelID(i)
	return msu
}

// SetStatus sets the "status" field.
func (msu *ModelStarUpdate) SetStatus(es enums.StarStatus) *ModelStarUpdate {
	msu.mutation.SetStatus(es)
	return msu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (msu *ModelStarUpdate) SetNillableStatus(es *enums.StarStatus) *ModelStarUpdate {
	if es != nil {
		msu.SetStatus(*es)
	}
	return msu
}

// SetUser sets the "user" edge to the User entity.
func (msu *ModelStarUpdate) SetUser(u *User) *ModelStarUpdate {
	return msu.SetUserID(u.ID)
}

// SetModel sets the "model" edge to the Model entity.
func (msu *ModelStarUpdate) SetModel(m *Model) *ModelStarUpdate {
	return msu.SetModelID(m.ID)
}

// Mutation returns the ModelStarMutation object of the builder.
func (msu *ModelStarUpdate) Mutation() *ModelStarMutation {
	return msu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (msu *ModelStarUpdate) ClearUser() *ModelStarUpdate {
	msu.mutation.ClearUser()
	return msu
}

// ClearModel clears the "model" edge to the Model entity.
func (msu *ModelStarUpdate) ClearModel() *ModelStarUpdate {
	msu.mutation.ClearModel()
	return msu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (msu *ModelStarUpdate) Save(ctx context.Context) (int, error) {
	msu.defaults()
	return withHooks(ctx, msu.sqlSave, msu.mutation, msu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (msu *ModelStarUpdate) SaveX(ctx context.Context) int {
	affected, err := msu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (msu *ModelStarUpdate) Exec(ctx context.Context) error {
	_, err := msu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msu *ModelStarUpdate) ExecX(ctx context.Context) {
	if err := msu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (msu *ModelStarUpdate) defaults() {
	if _, ok := msu.mutation.UpdatedAt(); !ok {
		v := modelstar.UpdateDefaultUpdatedAt()
		msu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msu *ModelStarUpdate) check() error {
	if v, ok := msu.mutation.Status(); ok {
		if err := modelstar.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "ModelStar.status": %w`, err)}
		}
	}
	if _, ok := msu.mutation.UserID(); msu.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ModelStar.user"`)
	}
	if _, ok := msu.mutation.ModelID(); msu.mutation.ModelCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ModelStar.model"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (msu *ModelStarUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ModelStarUpdate {
	msu.modifiers = append(msu.modifiers, modifiers...)
	return msu
}

func (msu *ModelStarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := msu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(modelstar.Table, modelstar.Columns, sqlgraph.NewFieldSpec(modelstar.FieldID, field.TypeInt64))
	if ps := msu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := msu.mutation.CreatedBy(); ok {
		_spec.SetField(modelstar.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := msu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(modelstar.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := msu.mutation.UpdatedBy(); ok {
		_spec.SetField(modelstar.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := msu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(modelstar.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := msu.mutation.UpdatedAt(); ok {
		_spec.SetField(modelstar.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := msu.mutation.DeletedAt(); ok {
		_spec.SetField(modelstar.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := msu.mutation.Status(); ok {
		_spec.SetField(modelstar.FieldStatus, field.TypeEnum, value)
	}
	if msu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   modelstar.UserTable,
			Columns: []string{modelstar.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   modelstar.UserTable,
			Columns: []string{modelstar.UserColumn},
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
	if msu.mutation.ModelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   modelstar.ModelTable,
			Columns: []string{modelstar.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msu.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   modelstar.ModelTable,
			Columns: []string{modelstar.ModelColumn},
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
	_spec.AddModifiers(msu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, msu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{modelstar.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	msu.mutation.done = true
	return n, nil
}

// ModelStarUpdateOne is the builder for updating a single ModelStar entity.
type ModelStarUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ModelStarMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (msuo *ModelStarUpdateOne) SetCreatedBy(i int64) *ModelStarUpdateOne {
	msuo.mutation.ResetCreatedBy()
	msuo.mutation.SetCreatedBy(i)
	return msuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (msuo *ModelStarUpdateOne) SetNillableCreatedBy(i *int64) *ModelStarUpdateOne {
	if i != nil {
		msuo.SetCreatedBy(*i)
	}
	return msuo
}

// AddCreatedBy adds i to the "created_by" field.
func (msuo *ModelStarUpdateOne) AddCreatedBy(i int64) *ModelStarUpdateOne {
	msuo.mutation.AddCreatedBy(i)
	return msuo
}

// SetUpdatedBy sets the "updated_by" field.
func (msuo *ModelStarUpdateOne) SetUpdatedBy(i int64) *ModelStarUpdateOne {
	msuo.mutation.ResetUpdatedBy()
	msuo.mutation.SetUpdatedBy(i)
	return msuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (msuo *ModelStarUpdateOne) SetNillableUpdatedBy(i *int64) *ModelStarUpdateOne {
	if i != nil {
		msuo.SetUpdatedBy(*i)
	}
	return msuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (msuo *ModelStarUpdateOne) AddUpdatedBy(i int64) *ModelStarUpdateOne {
	msuo.mutation.AddUpdatedBy(i)
	return msuo
}

// SetUpdatedAt sets the "updated_at" field.
func (msuo *ModelStarUpdateOne) SetUpdatedAt(t time.Time) *ModelStarUpdateOne {
	msuo.mutation.SetUpdatedAt(t)
	return msuo
}

// SetDeletedAt sets the "deleted_at" field.
func (msuo *ModelStarUpdateOne) SetDeletedAt(t time.Time) *ModelStarUpdateOne {
	msuo.mutation.SetDeletedAt(t)
	return msuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (msuo *ModelStarUpdateOne) SetNillableDeletedAt(t *time.Time) *ModelStarUpdateOne {
	if t != nil {
		msuo.SetDeletedAt(*t)
	}
	return msuo
}

// SetUserID sets the "user_id" field.
func (msuo *ModelStarUpdateOne) SetUserID(i int64) *ModelStarUpdateOne {
	msuo.mutation.SetUserID(i)
	return msuo
}

// SetModelID sets the "model_id" field.
func (msuo *ModelStarUpdateOne) SetModelID(i int64) *ModelStarUpdateOne {
	msuo.mutation.SetModelID(i)
	return msuo
}

// SetStatus sets the "status" field.
func (msuo *ModelStarUpdateOne) SetStatus(es enums.StarStatus) *ModelStarUpdateOne {
	msuo.mutation.SetStatus(es)
	return msuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (msuo *ModelStarUpdateOne) SetNillableStatus(es *enums.StarStatus) *ModelStarUpdateOne {
	if es != nil {
		msuo.SetStatus(*es)
	}
	return msuo
}

// SetUser sets the "user" edge to the User entity.
func (msuo *ModelStarUpdateOne) SetUser(u *User) *ModelStarUpdateOne {
	return msuo.SetUserID(u.ID)
}

// SetModel sets the "model" edge to the Model entity.
func (msuo *ModelStarUpdateOne) SetModel(m *Model) *ModelStarUpdateOne {
	return msuo.SetModelID(m.ID)
}

// Mutation returns the ModelStarMutation object of the builder.
func (msuo *ModelStarUpdateOne) Mutation() *ModelStarMutation {
	return msuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (msuo *ModelStarUpdateOne) ClearUser() *ModelStarUpdateOne {
	msuo.mutation.ClearUser()
	return msuo
}

// ClearModel clears the "model" edge to the Model entity.
func (msuo *ModelStarUpdateOne) ClearModel() *ModelStarUpdateOne {
	msuo.mutation.ClearModel()
	return msuo
}

// Where appends a list predicates to the ModelStarUpdate builder.
func (msuo *ModelStarUpdateOne) Where(ps ...predicate.ModelStar) *ModelStarUpdateOne {
	msuo.mutation.Where(ps...)
	return msuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (msuo *ModelStarUpdateOne) Select(field string, fields ...string) *ModelStarUpdateOne {
	msuo.fields = append([]string{field}, fields...)
	return msuo
}

// Save executes the query and returns the updated ModelStar entity.
func (msuo *ModelStarUpdateOne) Save(ctx context.Context) (*ModelStar, error) {
	msuo.defaults()
	return withHooks(ctx, msuo.sqlSave, msuo.mutation, msuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (msuo *ModelStarUpdateOne) SaveX(ctx context.Context) *ModelStar {
	node, err := msuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (msuo *ModelStarUpdateOne) Exec(ctx context.Context) error {
	_, err := msuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msuo *ModelStarUpdateOne) ExecX(ctx context.Context) {
	if err := msuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (msuo *ModelStarUpdateOne) defaults() {
	if _, ok := msuo.mutation.UpdatedAt(); !ok {
		v := modelstar.UpdateDefaultUpdatedAt()
		msuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msuo *ModelStarUpdateOne) check() error {
	if v, ok := msuo.mutation.Status(); ok {
		if err := modelstar.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "ModelStar.status": %w`, err)}
		}
	}
	if _, ok := msuo.mutation.UserID(); msuo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ModelStar.user"`)
	}
	if _, ok := msuo.mutation.ModelID(); msuo.mutation.ModelCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ModelStar.model"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (msuo *ModelStarUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ModelStarUpdateOne {
	msuo.modifiers = append(msuo.modifiers, modifiers...)
	return msuo
}

func (msuo *ModelStarUpdateOne) sqlSave(ctx context.Context) (_node *ModelStar, err error) {
	if err := msuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(modelstar.Table, modelstar.Columns, sqlgraph.NewFieldSpec(modelstar.FieldID, field.TypeInt64))
	id, ok := msuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "ModelStar.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := msuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, modelstar.FieldID)
		for _, f := range fields {
			if !modelstar.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != modelstar.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := msuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := msuo.mutation.CreatedBy(); ok {
		_spec.SetField(modelstar.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := msuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(modelstar.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := msuo.mutation.UpdatedBy(); ok {
		_spec.SetField(modelstar.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := msuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(modelstar.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := msuo.mutation.UpdatedAt(); ok {
		_spec.SetField(modelstar.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := msuo.mutation.DeletedAt(); ok {
		_spec.SetField(modelstar.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := msuo.mutation.Status(); ok {
		_spec.SetField(modelstar.FieldStatus, field.TypeEnum, value)
	}
	if msuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   modelstar.UserTable,
			Columns: []string{modelstar.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   modelstar.UserTable,
			Columns: []string{modelstar.UserColumn},
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
	if msuo.mutation.ModelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   modelstar.ModelTable,
			Columns: []string{modelstar.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msuo.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   modelstar.ModelTable,
			Columns: []string{modelstar.ModelColumn},
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
	_spec.AddModifiers(msuo.modifiers...)
	_node = &ModelStar{config: msuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, msuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{modelstar.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	msuo.mutation.done = true
	return _node, nil
}
