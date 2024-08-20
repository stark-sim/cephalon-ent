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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/usermodel"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// UserModelUpdate is the builder for updating UserModel entities.
type UserModelUpdate struct {
	config
	hooks     []Hook
	mutation  *UserModelMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserModelUpdate builder.
func (umu *UserModelUpdate) Where(ps ...predicate.UserModel) *UserModelUpdate {
	umu.mutation.Where(ps...)
	return umu
}

// SetCreatedBy sets the "created_by" field.
func (umu *UserModelUpdate) SetCreatedBy(i int64) *UserModelUpdate {
	umu.mutation.ResetCreatedBy()
	umu.mutation.SetCreatedBy(i)
	return umu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (umu *UserModelUpdate) SetNillableCreatedBy(i *int64) *UserModelUpdate {
	if i != nil {
		umu.SetCreatedBy(*i)
	}
	return umu
}

// AddCreatedBy adds i to the "created_by" field.
func (umu *UserModelUpdate) AddCreatedBy(i int64) *UserModelUpdate {
	umu.mutation.AddCreatedBy(i)
	return umu
}

// SetUpdatedBy sets the "updated_by" field.
func (umu *UserModelUpdate) SetUpdatedBy(i int64) *UserModelUpdate {
	umu.mutation.ResetUpdatedBy()
	umu.mutation.SetUpdatedBy(i)
	return umu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (umu *UserModelUpdate) SetNillableUpdatedBy(i *int64) *UserModelUpdate {
	if i != nil {
		umu.SetUpdatedBy(*i)
	}
	return umu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (umu *UserModelUpdate) AddUpdatedBy(i int64) *UserModelUpdate {
	umu.mutation.AddUpdatedBy(i)
	return umu
}

// SetUpdatedAt sets the "updated_at" field.
func (umu *UserModelUpdate) SetUpdatedAt(t time.Time) *UserModelUpdate {
	umu.mutation.SetUpdatedAt(t)
	return umu
}

// SetDeletedAt sets the "deleted_at" field.
func (umu *UserModelUpdate) SetDeletedAt(t time.Time) *UserModelUpdate {
	umu.mutation.SetDeletedAt(t)
	return umu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (umu *UserModelUpdate) SetNillableDeletedAt(t *time.Time) *UserModelUpdate {
	if t != nil {
		umu.SetDeletedAt(*t)
	}
	return umu
}

// SetUserID sets the "user_id" field.
func (umu *UserModelUpdate) SetUserID(i int64) *UserModelUpdate {
	umu.mutation.SetUserID(i)
	return umu
}

// SetModelID sets the "model_id" field.
func (umu *UserModelUpdate) SetModelID(i int64) *UserModelUpdate {
	umu.mutation.SetModelID(i)
	return umu
}

// SetRelation sets the "relation" field.
func (umu *UserModelUpdate) SetRelation(emr enums.UserModelRelation) *UserModelUpdate {
	umu.mutation.SetRelation(emr)
	return umu
}

// SetNillableRelation sets the "relation" field if the given value is not nil.
func (umu *UserModelUpdate) SetNillableRelation(emr *enums.UserModelRelation) *UserModelUpdate {
	if emr != nil {
		umu.SetRelation(*emr)
	}
	return umu
}

// SetStatus sets the "status" field.
func (umu *UserModelUpdate) SetStatus(emrs enums.UserModelRelationStatus) *UserModelUpdate {
	umu.mutation.SetStatus(emrs)
	return umu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (umu *UserModelUpdate) SetNillableStatus(emrs *enums.UserModelRelationStatus) *UserModelUpdate {
	if emrs != nil {
		umu.SetStatus(*emrs)
	}
	return umu
}

// SetUser sets the "user" edge to the User entity.
func (umu *UserModelUpdate) SetUser(u *User) *UserModelUpdate {
	return umu.SetUserID(u.ID)
}

// SetModel sets the "model" edge to the Model entity.
func (umu *UserModelUpdate) SetModel(m *Model) *UserModelUpdate {
	return umu.SetModelID(m.ID)
}

// Mutation returns the UserModelMutation object of the builder.
func (umu *UserModelUpdate) Mutation() *UserModelMutation {
	return umu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (umu *UserModelUpdate) ClearUser() *UserModelUpdate {
	umu.mutation.ClearUser()
	return umu
}

// ClearModel clears the "model" edge to the Model entity.
func (umu *UserModelUpdate) ClearModel() *UserModelUpdate {
	umu.mutation.ClearModel()
	return umu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (umu *UserModelUpdate) Save(ctx context.Context) (int, error) {
	umu.defaults()
	return withHooks(ctx, umu.sqlSave, umu.mutation, umu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (umu *UserModelUpdate) SaveX(ctx context.Context) int {
	affected, err := umu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (umu *UserModelUpdate) Exec(ctx context.Context) error {
	_, err := umu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umu *UserModelUpdate) ExecX(ctx context.Context) {
	if err := umu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (umu *UserModelUpdate) defaults() {
	if _, ok := umu.mutation.UpdatedAt(); !ok {
		v := usermodel.UpdateDefaultUpdatedAt()
		umu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (umu *UserModelUpdate) check() error {
	if v, ok := umu.mutation.Relation(); ok {
		if err := usermodel.RelationValidator(v); err != nil {
			return &ValidationError{Name: "relation", err: fmt.Errorf(`cep_ent: validator failed for field "UserModel.relation": %w`, err)}
		}
	}
	if v, ok := umu.mutation.Status(); ok {
		if err := usermodel.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "UserModel.status": %w`, err)}
		}
	}
	if _, ok := umu.mutation.UserID(); umu.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "UserModel.user"`)
	}
	if _, ok := umu.mutation.ModelID(); umu.mutation.ModelCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "UserModel.model"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (umu *UserModelUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserModelUpdate {
	umu.modifiers = append(umu.modifiers, modifiers...)
	return umu
}

func (umu *UserModelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := umu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usermodel.Table, usermodel.Columns, sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt64))
	if ps := umu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := umu.mutation.CreatedBy(); ok {
		_spec.SetField(usermodel.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := umu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(usermodel.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := umu.mutation.UpdatedBy(); ok {
		_spec.SetField(usermodel.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := umu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(usermodel.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := umu.mutation.UpdatedAt(); ok {
		_spec.SetField(usermodel.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := umu.mutation.DeletedAt(); ok {
		_spec.SetField(usermodel.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := umu.mutation.Relation(); ok {
		_spec.SetField(usermodel.FieldRelation, field.TypeEnum, value)
	}
	if value, ok := umu.mutation.Status(); ok {
		_spec.SetField(usermodel.FieldStatus, field.TypeEnum, value)
	}
	if umu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermodel.UserTable,
			Columns: []string{usermodel.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := umu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermodel.UserTable,
			Columns: []string{usermodel.UserColumn},
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
	if umu.mutation.ModelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermodel.ModelTable,
			Columns: []string{usermodel.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := umu.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermodel.ModelTable,
			Columns: []string{usermodel.ModelColumn},
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
	_spec.AddModifiers(umu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, umu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usermodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	umu.mutation.done = true
	return n, nil
}

// UserModelUpdateOne is the builder for updating a single UserModel entity.
type UserModelUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserModelMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (umuo *UserModelUpdateOne) SetCreatedBy(i int64) *UserModelUpdateOne {
	umuo.mutation.ResetCreatedBy()
	umuo.mutation.SetCreatedBy(i)
	return umuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (umuo *UserModelUpdateOne) SetNillableCreatedBy(i *int64) *UserModelUpdateOne {
	if i != nil {
		umuo.SetCreatedBy(*i)
	}
	return umuo
}

// AddCreatedBy adds i to the "created_by" field.
func (umuo *UserModelUpdateOne) AddCreatedBy(i int64) *UserModelUpdateOne {
	umuo.mutation.AddCreatedBy(i)
	return umuo
}

// SetUpdatedBy sets the "updated_by" field.
func (umuo *UserModelUpdateOne) SetUpdatedBy(i int64) *UserModelUpdateOne {
	umuo.mutation.ResetUpdatedBy()
	umuo.mutation.SetUpdatedBy(i)
	return umuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (umuo *UserModelUpdateOne) SetNillableUpdatedBy(i *int64) *UserModelUpdateOne {
	if i != nil {
		umuo.SetUpdatedBy(*i)
	}
	return umuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (umuo *UserModelUpdateOne) AddUpdatedBy(i int64) *UserModelUpdateOne {
	umuo.mutation.AddUpdatedBy(i)
	return umuo
}

// SetUpdatedAt sets the "updated_at" field.
func (umuo *UserModelUpdateOne) SetUpdatedAt(t time.Time) *UserModelUpdateOne {
	umuo.mutation.SetUpdatedAt(t)
	return umuo
}

// SetDeletedAt sets the "deleted_at" field.
func (umuo *UserModelUpdateOne) SetDeletedAt(t time.Time) *UserModelUpdateOne {
	umuo.mutation.SetDeletedAt(t)
	return umuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (umuo *UserModelUpdateOne) SetNillableDeletedAt(t *time.Time) *UserModelUpdateOne {
	if t != nil {
		umuo.SetDeletedAt(*t)
	}
	return umuo
}

// SetUserID sets the "user_id" field.
func (umuo *UserModelUpdateOne) SetUserID(i int64) *UserModelUpdateOne {
	umuo.mutation.SetUserID(i)
	return umuo
}

// SetModelID sets the "model_id" field.
func (umuo *UserModelUpdateOne) SetModelID(i int64) *UserModelUpdateOne {
	umuo.mutation.SetModelID(i)
	return umuo
}

// SetRelation sets the "relation" field.
func (umuo *UserModelUpdateOne) SetRelation(emr enums.UserModelRelation) *UserModelUpdateOne {
	umuo.mutation.SetRelation(emr)
	return umuo
}

// SetNillableRelation sets the "relation" field if the given value is not nil.
func (umuo *UserModelUpdateOne) SetNillableRelation(emr *enums.UserModelRelation) *UserModelUpdateOne {
	if emr != nil {
		umuo.SetRelation(*emr)
	}
	return umuo
}

// SetStatus sets the "status" field.
func (umuo *UserModelUpdateOne) SetStatus(emrs enums.UserModelRelationStatus) *UserModelUpdateOne {
	umuo.mutation.SetStatus(emrs)
	return umuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (umuo *UserModelUpdateOne) SetNillableStatus(emrs *enums.UserModelRelationStatus) *UserModelUpdateOne {
	if emrs != nil {
		umuo.SetStatus(*emrs)
	}
	return umuo
}

// SetUser sets the "user" edge to the User entity.
func (umuo *UserModelUpdateOne) SetUser(u *User) *UserModelUpdateOne {
	return umuo.SetUserID(u.ID)
}

// SetModel sets the "model" edge to the Model entity.
func (umuo *UserModelUpdateOne) SetModel(m *Model) *UserModelUpdateOne {
	return umuo.SetModelID(m.ID)
}

// Mutation returns the UserModelMutation object of the builder.
func (umuo *UserModelUpdateOne) Mutation() *UserModelMutation {
	return umuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (umuo *UserModelUpdateOne) ClearUser() *UserModelUpdateOne {
	umuo.mutation.ClearUser()
	return umuo
}

// ClearModel clears the "model" edge to the Model entity.
func (umuo *UserModelUpdateOne) ClearModel() *UserModelUpdateOne {
	umuo.mutation.ClearModel()
	return umuo
}

// Where appends a list predicates to the UserModelUpdate builder.
func (umuo *UserModelUpdateOne) Where(ps ...predicate.UserModel) *UserModelUpdateOne {
	umuo.mutation.Where(ps...)
	return umuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (umuo *UserModelUpdateOne) Select(field string, fields ...string) *UserModelUpdateOne {
	umuo.fields = append([]string{field}, fields...)
	return umuo
}

// Save executes the query and returns the updated UserModel entity.
func (umuo *UserModelUpdateOne) Save(ctx context.Context) (*UserModel, error) {
	umuo.defaults()
	return withHooks(ctx, umuo.sqlSave, umuo.mutation, umuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (umuo *UserModelUpdateOne) SaveX(ctx context.Context) *UserModel {
	node, err := umuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (umuo *UserModelUpdateOne) Exec(ctx context.Context) error {
	_, err := umuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umuo *UserModelUpdateOne) ExecX(ctx context.Context) {
	if err := umuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (umuo *UserModelUpdateOne) defaults() {
	if _, ok := umuo.mutation.UpdatedAt(); !ok {
		v := usermodel.UpdateDefaultUpdatedAt()
		umuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (umuo *UserModelUpdateOne) check() error {
	if v, ok := umuo.mutation.Relation(); ok {
		if err := usermodel.RelationValidator(v); err != nil {
			return &ValidationError{Name: "relation", err: fmt.Errorf(`cep_ent: validator failed for field "UserModel.relation": %w`, err)}
		}
	}
	if v, ok := umuo.mutation.Status(); ok {
		if err := usermodel.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "UserModel.status": %w`, err)}
		}
	}
	if _, ok := umuo.mutation.UserID(); umuo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "UserModel.user"`)
	}
	if _, ok := umuo.mutation.ModelID(); umuo.mutation.ModelCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "UserModel.model"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (umuo *UserModelUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserModelUpdateOne {
	umuo.modifiers = append(umuo.modifiers, modifiers...)
	return umuo
}

func (umuo *UserModelUpdateOne) sqlSave(ctx context.Context) (_node *UserModel, err error) {
	if err := umuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usermodel.Table, usermodel.Columns, sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt64))
	id, ok := umuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "UserModel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := umuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usermodel.FieldID)
		for _, f := range fields {
			if !usermodel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != usermodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := umuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := umuo.mutation.CreatedBy(); ok {
		_spec.SetField(usermodel.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := umuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(usermodel.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := umuo.mutation.UpdatedBy(); ok {
		_spec.SetField(usermodel.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := umuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(usermodel.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := umuo.mutation.UpdatedAt(); ok {
		_spec.SetField(usermodel.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := umuo.mutation.DeletedAt(); ok {
		_spec.SetField(usermodel.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := umuo.mutation.Relation(); ok {
		_spec.SetField(usermodel.FieldRelation, field.TypeEnum, value)
	}
	if value, ok := umuo.mutation.Status(); ok {
		_spec.SetField(usermodel.FieldStatus, field.TypeEnum, value)
	}
	if umuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermodel.UserTable,
			Columns: []string{usermodel.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := umuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermodel.UserTable,
			Columns: []string{usermodel.UserColumn},
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
	if umuo.mutation.ModelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermodel.ModelTable,
			Columns: []string{usermodel.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := umuo.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermodel.ModelTable,
			Columns: []string{usermodel.ModelColumn},
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
	_spec.AddModifiers(umuo.modifiers...)
	_node = &UserModel{config: umuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, umuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usermodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	umuo.mutation.done = true
	return _node, nil
}