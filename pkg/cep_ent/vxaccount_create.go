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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/vxaccount"
)

// VXAccountCreate is the builder for creating a VXAccount entity.
type VXAccountCreate struct {
	config
	mutation *VXAccountMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (vac *VXAccountCreate) SetCreatedBy(i int64) *VXAccountCreate {
	vac.mutation.SetCreatedBy(i)
	return vac
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (vac *VXAccountCreate) SetNillableCreatedBy(i *int64) *VXAccountCreate {
	if i != nil {
		vac.SetCreatedBy(*i)
	}
	return vac
}

// SetUpdatedBy sets the "updated_by" field.
func (vac *VXAccountCreate) SetUpdatedBy(i int64) *VXAccountCreate {
	vac.mutation.SetUpdatedBy(i)
	return vac
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (vac *VXAccountCreate) SetNillableUpdatedBy(i *int64) *VXAccountCreate {
	if i != nil {
		vac.SetUpdatedBy(*i)
	}
	return vac
}

// SetCreatedAt sets the "created_at" field.
func (vac *VXAccountCreate) SetCreatedAt(t time.Time) *VXAccountCreate {
	vac.mutation.SetCreatedAt(t)
	return vac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vac *VXAccountCreate) SetNillableCreatedAt(t *time.Time) *VXAccountCreate {
	if t != nil {
		vac.SetCreatedAt(*t)
	}
	return vac
}

// SetUpdatedAt sets the "updated_at" field.
func (vac *VXAccountCreate) SetUpdatedAt(t time.Time) *VXAccountCreate {
	vac.mutation.SetUpdatedAt(t)
	return vac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vac *VXAccountCreate) SetNillableUpdatedAt(t *time.Time) *VXAccountCreate {
	if t != nil {
		vac.SetUpdatedAt(*t)
	}
	return vac
}

// SetDeletedAt sets the "deleted_at" field.
func (vac *VXAccountCreate) SetDeletedAt(t time.Time) *VXAccountCreate {
	vac.mutation.SetDeletedAt(t)
	return vac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vac *VXAccountCreate) SetNillableDeletedAt(t *time.Time) *VXAccountCreate {
	if t != nil {
		vac.SetDeletedAt(*t)
	}
	return vac
}

// SetOpenID sets the "open_id" field.
func (vac *VXAccountCreate) SetOpenID(s string) *VXAccountCreate {
	vac.mutation.SetOpenID(s)
	return vac
}

// SetNillableOpenID sets the "open_id" field if the given value is not nil.
func (vac *VXAccountCreate) SetNillableOpenID(s *string) *VXAccountCreate {
	if s != nil {
		vac.SetOpenID(*s)
	}
	return vac
}

// SetUnionID sets the "union_id" field.
func (vac *VXAccountCreate) SetUnionID(s string) *VXAccountCreate {
	vac.mutation.SetUnionID(s)
	return vac
}

// SetNillableUnionID sets the "union_id" field if the given value is not nil.
func (vac *VXAccountCreate) SetNillableUnionID(s *string) *VXAccountCreate {
	if s != nil {
		vac.SetUnionID(*s)
	}
	return vac
}

// SetScope sets the "scope" field.
func (vac *VXAccountCreate) SetScope(s string) *VXAccountCreate {
	vac.mutation.SetScope(s)
	return vac
}

// SetNillableScope sets the "scope" field if the given value is not nil.
func (vac *VXAccountCreate) SetNillableScope(s *string) *VXAccountCreate {
	if s != nil {
		vac.SetScope(*s)
	}
	return vac
}

// SetSessionKey sets the "session_key" field.
func (vac *VXAccountCreate) SetSessionKey(s string) *VXAccountCreate {
	vac.mutation.SetSessionKey(s)
	return vac
}

// SetNillableSessionKey sets the "session_key" field if the given value is not nil.
func (vac *VXAccountCreate) SetNillableSessionKey(s *string) *VXAccountCreate {
	if s != nil {
		vac.SetSessionKey(*s)
	}
	return vac
}

// SetUserID sets the "user_id" field.
func (vac *VXAccountCreate) SetUserID(i int64) *VXAccountCreate {
	vac.mutation.SetUserID(i)
	return vac
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (vac *VXAccountCreate) SetNillableUserID(i *int64) *VXAccountCreate {
	if i != nil {
		vac.SetUserID(*i)
	}
	return vac
}

// SetID sets the "id" field.
func (vac *VXAccountCreate) SetID(i int64) *VXAccountCreate {
	vac.mutation.SetID(i)
	return vac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vac *VXAccountCreate) SetNillableID(i *int64) *VXAccountCreate {
	if i != nil {
		vac.SetID(*i)
	}
	return vac
}

// SetUser sets the "user" edge to the User entity.
func (vac *VXAccountCreate) SetUser(u *User) *VXAccountCreate {
	return vac.SetUserID(u.ID)
}

// Mutation returns the VXAccountMutation object of the builder.
func (vac *VXAccountCreate) Mutation() *VXAccountMutation {
	return vac.mutation
}

// Save creates the VXAccount in the database.
func (vac *VXAccountCreate) Save(ctx context.Context) (*VXAccount, error) {
	vac.defaults()
	return withHooks(ctx, vac.sqlSave, vac.mutation, vac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vac *VXAccountCreate) SaveX(ctx context.Context) *VXAccount {
	v, err := vac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vac *VXAccountCreate) Exec(ctx context.Context) error {
	_, err := vac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vac *VXAccountCreate) ExecX(ctx context.Context) {
	if err := vac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vac *VXAccountCreate) defaults() {
	if _, ok := vac.mutation.CreatedBy(); !ok {
		v := vxaccount.DefaultCreatedBy
		vac.mutation.SetCreatedBy(v)
	}
	if _, ok := vac.mutation.UpdatedBy(); !ok {
		v := vxaccount.DefaultUpdatedBy
		vac.mutation.SetUpdatedBy(v)
	}
	if _, ok := vac.mutation.CreatedAt(); !ok {
		v := vxaccount.DefaultCreatedAt()
		vac.mutation.SetCreatedAt(v)
	}
	if _, ok := vac.mutation.UpdatedAt(); !ok {
		v := vxaccount.DefaultUpdatedAt()
		vac.mutation.SetUpdatedAt(v)
	}
	if _, ok := vac.mutation.DeletedAt(); !ok {
		v := vxaccount.DefaultDeletedAt
		vac.mutation.SetDeletedAt(v)
	}
	if _, ok := vac.mutation.OpenID(); !ok {
		v := vxaccount.DefaultOpenID
		vac.mutation.SetOpenID(v)
	}
	if _, ok := vac.mutation.UnionID(); !ok {
		v := vxaccount.DefaultUnionID
		vac.mutation.SetUnionID(v)
	}
	if _, ok := vac.mutation.Scope(); !ok {
		v := vxaccount.DefaultScope
		vac.mutation.SetScope(v)
	}
	if _, ok := vac.mutation.SessionKey(); !ok {
		v := vxaccount.DefaultSessionKey
		vac.mutation.SetSessionKey(v)
	}
	if _, ok := vac.mutation.UserID(); !ok {
		v := vxaccount.DefaultUserID
		vac.mutation.SetUserID(v)
	}
	if _, ok := vac.mutation.ID(); !ok {
		v := vxaccount.DefaultID()
		vac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vac *VXAccountCreate) check() error {
	if _, ok := vac.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "VXAccount.created_by"`)}
	}
	if _, ok := vac.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "VXAccount.updated_by"`)}
	}
	if _, ok := vac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "VXAccount.created_at"`)}
	}
	if _, ok := vac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "VXAccount.updated_at"`)}
	}
	if _, ok := vac.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "VXAccount.deleted_at"`)}
	}
	if _, ok := vac.mutation.OpenID(); !ok {
		return &ValidationError{Name: "open_id", err: errors.New(`cep_ent: missing required field "VXAccount.open_id"`)}
	}
	if _, ok := vac.mutation.UnionID(); !ok {
		return &ValidationError{Name: "union_id", err: errors.New(`cep_ent: missing required field "VXAccount.union_id"`)}
	}
	if _, ok := vac.mutation.Scope(); !ok {
		return &ValidationError{Name: "scope", err: errors.New(`cep_ent: missing required field "VXAccount.scope"`)}
	}
	if _, ok := vac.mutation.SessionKey(); !ok {
		return &ValidationError{Name: "session_key", err: errors.New(`cep_ent: missing required field "VXAccount.session_key"`)}
	}
	if _, ok := vac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`cep_ent: missing required field "VXAccount.user_id"`)}
	}
	if _, ok := vac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`cep_ent: missing required edge "VXAccount.user"`)}
	}
	return nil
}

func (vac *VXAccountCreate) sqlSave(ctx context.Context) (*VXAccount, error) {
	if err := vac.check(); err != nil {
		return nil, err
	}
	_node, _spec := vac.createSpec()
	if err := sqlgraph.CreateNode(ctx, vac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	vac.mutation.id = &_node.ID
	vac.mutation.done = true
	return _node, nil
}

func (vac *VXAccountCreate) createSpec() (*VXAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &VXAccount{config: vac.config}
		_spec = sqlgraph.NewCreateSpec(vxaccount.Table, sqlgraph.NewFieldSpec(vxaccount.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = vac.conflict
	if id, ok := vac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vac.mutation.CreatedBy(); ok {
		_spec.SetField(vxaccount.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := vac.mutation.UpdatedBy(); ok {
		_spec.SetField(vxaccount.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := vac.mutation.CreatedAt(); ok {
		_spec.SetField(vxaccount.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vac.mutation.UpdatedAt(); ok {
		_spec.SetField(vxaccount.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := vac.mutation.DeletedAt(); ok {
		_spec.SetField(vxaccount.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := vac.mutation.OpenID(); ok {
		_spec.SetField(vxaccount.FieldOpenID, field.TypeString, value)
		_node.OpenID = value
	}
	if value, ok := vac.mutation.UnionID(); ok {
		_spec.SetField(vxaccount.FieldUnionID, field.TypeString, value)
		_node.UnionID = value
	}
	if value, ok := vac.mutation.Scope(); ok {
		_spec.SetField(vxaccount.FieldScope, field.TypeString, value)
		_node.Scope = value
	}
	if value, ok := vac.mutation.SessionKey(); ok {
		_spec.SetField(vxaccount.FieldSessionKey, field.TypeString, value)
		_node.SessionKey = value
	}
	if nodes := vac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vxaccount.UserTable,
			Columns: []string{vxaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.VXAccount.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VXAccountUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (vac *VXAccountCreate) OnConflict(opts ...sql.ConflictOption) *VXAccountUpsertOne {
	vac.conflict = opts
	return &VXAccountUpsertOne{
		create: vac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.VXAccount.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vac *VXAccountCreate) OnConflictColumns(columns ...string) *VXAccountUpsertOne {
	vac.conflict = append(vac.conflict, sql.ConflictColumns(columns...))
	return &VXAccountUpsertOne{
		create: vac,
	}
}

type (
	// VXAccountUpsertOne is the builder for "upsert"-ing
	//  one VXAccount node.
	VXAccountUpsertOne struct {
		create *VXAccountCreate
	}

	// VXAccountUpsert is the "OnConflict" setter.
	VXAccountUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *VXAccountUpsert) SetCreatedBy(v int64) *VXAccountUpsert {
	u.Set(vxaccount.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *VXAccountUpsert) UpdateCreatedBy() *VXAccountUpsert {
	u.SetExcluded(vxaccount.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *VXAccountUpsert) AddCreatedBy(v int64) *VXAccountUpsert {
	u.Add(vxaccount.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *VXAccountUpsert) SetUpdatedBy(v int64) *VXAccountUpsert {
	u.Set(vxaccount.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *VXAccountUpsert) UpdateUpdatedBy() *VXAccountUpsert {
	u.SetExcluded(vxaccount.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *VXAccountUpsert) AddUpdatedBy(v int64) *VXAccountUpsert {
	u.Add(vxaccount.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *VXAccountUpsert) SetUpdatedAt(v time.Time) *VXAccountUpsert {
	u.Set(vxaccount.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *VXAccountUpsert) UpdateUpdatedAt() *VXAccountUpsert {
	u.SetExcluded(vxaccount.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *VXAccountUpsert) SetDeletedAt(v time.Time) *VXAccountUpsert {
	u.Set(vxaccount.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *VXAccountUpsert) UpdateDeletedAt() *VXAccountUpsert {
	u.SetExcluded(vxaccount.FieldDeletedAt)
	return u
}

// SetOpenID sets the "open_id" field.
func (u *VXAccountUpsert) SetOpenID(v string) *VXAccountUpsert {
	u.Set(vxaccount.FieldOpenID, v)
	return u
}

// UpdateOpenID sets the "open_id" field to the value that was provided on create.
func (u *VXAccountUpsert) UpdateOpenID() *VXAccountUpsert {
	u.SetExcluded(vxaccount.FieldOpenID)
	return u
}

// SetUnionID sets the "union_id" field.
func (u *VXAccountUpsert) SetUnionID(v string) *VXAccountUpsert {
	u.Set(vxaccount.FieldUnionID, v)
	return u
}

// UpdateUnionID sets the "union_id" field to the value that was provided on create.
func (u *VXAccountUpsert) UpdateUnionID() *VXAccountUpsert {
	u.SetExcluded(vxaccount.FieldUnionID)
	return u
}

// SetScope sets the "scope" field.
func (u *VXAccountUpsert) SetScope(v string) *VXAccountUpsert {
	u.Set(vxaccount.FieldScope, v)
	return u
}

// UpdateScope sets the "scope" field to the value that was provided on create.
func (u *VXAccountUpsert) UpdateScope() *VXAccountUpsert {
	u.SetExcluded(vxaccount.FieldScope)
	return u
}

// SetSessionKey sets the "session_key" field.
func (u *VXAccountUpsert) SetSessionKey(v string) *VXAccountUpsert {
	u.Set(vxaccount.FieldSessionKey, v)
	return u
}

// UpdateSessionKey sets the "session_key" field to the value that was provided on create.
func (u *VXAccountUpsert) UpdateSessionKey() *VXAccountUpsert {
	u.SetExcluded(vxaccount.FieldSessionKey)
	return u
}

// SetUserID sets the "user_id" field.
func (u *VXAccountUpsert) SetUserID(v int64) *VXAccountUpsert {
	u.Set(vxaccount.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *VXAccountUpsert) UpdateUserID() *VXAccountUpsert {
	u.SetExcluded(vxaccount.FieldUserID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.VXAccount.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(vxaccount.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *VXAccountUpsertOne) UpdateNewValues() *VXAccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(vxaccount.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(vxaccount.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.VXAccount.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *VXAccountUpsertOne) Ignore() *VXAccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VXAccountUpsertOne) DoNothing() *VXAccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VXAccountCreate.OnConflict
// documentation for more info.
func (u *VXAccountUpsertOne) Update(set func(*VXAccountUpsert)) *VXAccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VXAccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *VXAccountUpsertOne) SetCreatedBy(v int64) *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *VXAccountUpsertOne) AddCreatedBy(v int64) *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *VXAccountUpsertOne) UpdateCreatedBy() *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *VXAccountUpsertOne) SetUpdatedBy(v int64) *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *VXAccountUpsertOne) AddUpdatedBy(v int64) *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *VXAccountUpsertOne) UpdateUpdatedBy() *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *VXAccountUpsertOne) SetUpdatedAt(v time.Time) *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *VXAccountUpsertOne) UpdateUpdatedAt() *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *VXAccountUpsertOne) SetDeletedAt(v time.Time) *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *VXAccountUpsertOne) UpdateDeletedAt() *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetOpenID sets the "open_id" field.
func (u *VXAccountUpsertOne) SetOpenID(v string) *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetOpenID(v)
	})
}

// UpdateOpenID sets the "open_id" field to the value that was provided on create.
func (u *VXAccountUpsertOne) UpdateOpenID() *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateOpenID()
	})
}

// SetUnionID sets the "union_id" field.
func (u *VXAccountUpsertOne) SetUnionID(v string) *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetUnionID(v)
	})
}

// UpdateUnionID sets the "union_id" field to the value that was provided on create.
func (u *VXAccountUpsertOne) UpdateUnionID() *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateUnionID()
	})
}

// SetScope sets the "scope" field.
func (u *VXAccountUpsertOne) SetScope(v string) *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetScope(v)
	})
}

// UpdateScope sets the "scope" field to the value that was provided on create.
func (u *VXAccountUpsertOne) UpdateScope() *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateScope()
	})
}

// SetSessionKey sets the "session_key" field.
func (u *VXAccountUpsertOne) SetSessionKey(v string) *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetSessionKey(v)
	})
}

// UpdateSessionKey sets the "session_key" field to the value that was provided on create.
func (u *VXAccountUpsertOne) UpdateSessionKey() *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateSessionKey()
	})
}

// SetUserID sets the "user_id" field.
func (u *VXAccountUpsertOne) SetUserID(v int64) *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *VXAccountUpsertOne) UpdateUserID() *VXAccountUpsertOne {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *VXAccountUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for VXAccountCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VXAccountUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *VXAccountUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *VXAccountUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// VXAccountCreateBulk is the builder for creating many VXAccount entities in bulk.
type VXAccountCreateBulk struct {
	config
	err      error
	builders []*VXAccountCreate
	conflict []sql.ConflictOption
}

// Save creates the VXAccount entities in the database.
func (vacb *VXAccountCreateBulk) Save(ctx context.Context) ([]*VXAccount, error) {
	if vacb.err != nil {
		return nil, vacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vacb.builders))
	nodes := make([]*VXAccount, len(vacb.builders))
	mutators := make([]Mutator, len(vacb.builders))
	for i := range vacb.builders {
		func(i int, root context.Context) {
			builder := vacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VXAccountMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = vacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vacb *VXAccountCreateBulk) SaveX(ctx context.Context) []*VXAccount {
	v, err := vacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vacb *VXAccountCreateBulk) Exec(ctx context.Context) error {
	_, err := vacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vacb *VXAccountCreateBulk) ExecX(ctx context.Context) {
	if err := vacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.VXAccount.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VXAccountUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (vacb *VXAccountCreateBulk) OnConflict(opts ...sql.ConflictOption) *VXAccountUpsertBulk {
	vacb.conflict = opts
	return &VXAccountUpsertBulk{
		create: vacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.VXAccount.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vacb *VXAccountCreateBulk) OnConflictColumns(columns ...string) *VXAccountUpsertBulk {
	vacb.conflict = append(vacb.conflict, sql.ConflictColumns(columns...))
	return &VXAccountUpsertBulk{
		create: vacb,
	}
}

// VXAccountUpsertBulk is the builder for "upsert"-ing
// a bulk of VXAccount nodes.
type VXAccountUpsertBulk struct {
	create *VXAccountCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.VXAccount.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(vxaccount.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *VXAccountUpsertBulk) UpdateNewValues() *VXAccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(vxaccount.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(vxaccount.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.VXAccount.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *VXAccountUpsertBulk) Ignore() *VXAccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VXAccountUpsertBulk) DoNothing() *VXAccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VXAccountCreateBulk.OnConflict
// documentation for more info.
func (u *VXAccountUpsertBulk) Update(set func(*VXAccountUpsert)) *VXAccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VXAccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *VXAccountUpsertBulk) SetCreatedBy(v int64) *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *VXAccountUpsertBulk) AddCreatedBy(v int64) *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *VXAccountUpsertBulk) UpdateCreatedBy() *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *VXAccountUpsertBulk) SetUpdatedBy(v int64) *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *VXAccountUpsertBulk) AddUpdatedBy(v int64) *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *VXAccountUpsertBulk) UpdateUpdatedBy() *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *VXAccountUpsertBulk) SetUpdatedAt(v time.Time) *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *VXAccountUpsertBulk) UpdateUpdatedAt() *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *VXAccountUpsertBulk) SetDeletedAt(v time.Time) *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *VXAccountUpsertBulk) UpdateDeletedAt() *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetOpenID sets the "open_id" field.
func (u *VXAccountUpsertBulk) SetOpenID(v string) *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetOpenID(v)
	})
}

// UpdateOpenID sets the "open_id" field to the value that was provided on create.
func (u *VXAccountUpsertBulk) UpdateOpenID() *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateOpenID()
	})
}

// SetUnionID sets the "union_id" field.
func (u *VXAccountUpsertBulk) SetUnionID(v string) *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetUnionID(v)
	})
}

// UpdateUnionID sets the "union_id" field to the value that was provided on create.
func (u *VXAccountUpsertBulk) UpdateUnionID() *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateUnionID()
	})
}

// SetScope sets the "scope" field.
func (u *VXAccountUpsertBulk) SetScope(v string) *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetScope(v)
	})
}

// UpdateScope sets the "scope" field to the value that was provided on create.
func (u *VXAccountUpsertBulk) UpdateScope() *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateScope()
	})
}

// SetSessionKey sets the "session_key" field.
func (u *VXAccountUpsertBulk) SetSessionKey(v string) *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetSessionKey(v)
	})
}

// UpdateSessionKey sets the "session_key" field to the value that was provided on create.
func (u *VXAccountUpsertBulk) UpdateSessionKey() *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateSessionKey()
	})
}

// SetUserID sets the "user_id" field.
func (u *VXAccountUpsertBulk) SetUserID(v int64) *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *VXAccountUpsertBulk) UpdateUserID() *VXAccountUpsertBulk {
	return u.Update(func(s *VXAccountUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *VXAccountUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the VXAccountCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for VXAccountCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VXAccountUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
