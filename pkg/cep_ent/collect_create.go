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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/collect"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// CollectCreate is the builder for creating a Collect entity.
type CollectCreate struct {
	config
	mutation *CollectMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (cc *CollectCreate) SetCreatedBy(i int64) *CollectCreate {
	cc.mutation.SetCreatedBy(i)
	return cc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cc *CollectCreate) SetNillableCreatedBy(i *int64) *CollectCreate {
	if i != nil {
		cc.SetCreatedBy(*i)
	}
	return cc
}

// SetUpdatedBy sets the "updated_by" field.
func (cc *CollectCreate) SetUpdatedBy(i int64) *CollectCreate {
	cc.mutation.SetUpdatedBy(i)
	return cc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cc *CollectCreate) SetNillableUpdatedBy(i *int64) *CollectCreate {
	if i != nil {
		cc.SetUpdatedBy(*i)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CollectCreate) SetCreatedAt(t time.Time) *CollectCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CollectCreate) SetNillableCreatedAt(t *time.Time) *CollectCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CollectCreate) SetUpdatedAt(t time.Time) *CollectCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CollectCreate) SetNillableUpdatedAt(t *time.Time) *CollectCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CollectCreate) SetDeletedAt(t time.Time) *CollectCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CollectCreate) SetNillableDeletedAt(t *time.Time) *CollectCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetURL sets the "url" field.
func (cc *CollectCreate) SetURL(s string) *CollectCreate {
	cc.mutation.SetURL(s)
	return cc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (cc *CollectCreate) SetNillableURL(s *string) *CollectCreate {
	if s != nil {
		cc.SetURL(*s)
	}
	return cc
}

// SetUserID sets the "user_id" field.
func (cc *CollectCreate) SetUserID(i int64) *CollectCreate {
	cc.mutation.SetUserID(i)
	return cc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cc *CollectCreate) SetNillableUserID(i *int64) *CollectCreate {
	if i != nil {
		cc.SetUserID(*i)
	}
	return cc
}

// SetJpgName sets the "jpg_name" field.
func (cc *CollectCreate) SetJpgName(i int64) *CollectCreate {
	cc.mutation.SetJpgName(i)
	return cc
}

// SetNillableJpgName sets the "jpg_name" field if the given value is not nil.
func (cc *CollectCreate) SetNillableJpgName(i *int64) *CollectCreate {
	if i != nil {
		cc.SetJpgName(*i)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CollectCreate) SetID(i int64) *CollectCreate {
	cc.mutation.SetID(i)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CollectCreate) SetNillableID(i *int64) *CollectCreate {
	if i != nil {
		cc.SetID(*i)
	}
	return cc
}

// SetUser sets the "user" edge to the User entity.
func (cc *CollectCreate) SetUser(u *User) *CollectCreate {
	return cc.SetUserID(u.ID)
}

// Mutation returns the CollectMutation object of the builder.
func (cc *CollectCreate) Mutation() *CollectMutation {
	return cc.mutation
}

// Save creates the Collect in the database.
func (cc *CollectCreate) Save(ctx context.Context) (*Collect, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CollectCreate) SaveX(ctx context.Context) *Collect {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CollectCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CollectCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CollectCreate) defaults() {
	if _, ok := cc.mutation.CreatedBy(); !ok {
		v := collect.DefaultCreatedBy
		cc.mutation.SetCreatedBy(v)
	}
	if _, ok := cc.mutation.UpdatedBy(); !ok {
		v := collect.DefaultUpdatedBy
		cc.mutation.SetUpdatedBy(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := collect.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := collect.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		v := collect.DefaultDeletedAt
		cc.mutation.SetDeletedAt(v)
	}
	if _, ok := cc.mutation.URL(); !ok {
		v := collect.DefaultURL
		cc.mutation.SetURL(v)
	}
	if _, ok := cc.mutation.UserID(); !ok {
		v := collect.DefaultUserID
		cc.mutation.SetUserID(v)
	}
	if _, ok := cc.mutation.JpgName(); !ok {
		v := collect.DefaultJpgName
		cc.mutation.SetJpgName(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := collect.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CollectCreate) check() error {
	if _, ok := cc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "Collect.created_by"`)}
	}
	if _, ok := cc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "Collect.updated_by"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "Collect.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "Collect.updated_at"`)}
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "Collect.deleted_at"`)}
	}
	if _, ok := cc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`cep_ent: missing required field "Collect.url"`)}
	}
	if _, ok := cc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`cep_ent: missing required field "Collect.user_id"`)}
	}
	if _, ok := cc.mutation.JpgName(); !ok {
		return &ValidationError{Name: "jpg_name", err: errors.New(`cep_ent: missing required field "Collect.jpg_name"`)}
	}
	if _, ok := cc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`cep_ent: missing required edge "Collect.user"`)}
	}
	return nil
}

func (cc *CollectCreate) sqlSave(ctx context.Context) (*Collect, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CollectCreate) createSpec() (*Collect, *sqlgraph.CreateSpec) {
	var (
		_node = &Collect{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(collect.Table, sqlgraph.NewFieldSpec(collect.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedBy(); ok {
		_spec.SetField(collect.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := cc.mutation.UpdatedBy(); ok {
		_spec.SetField(collect.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(collect.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(collect.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(collect.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.URL(); ok {
		_spec.SetField(collect.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := cc.mutation.JpgName(); ok {
		_spec.SetField(collect.FieldJpgName, field.TypeInt64, value)
		_node.JpgName = value
	}
	if nodes := cc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Collect.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CollectUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (cc *CollectCreate) OnConflict(opts ...sql.ConflictOption) *CollectUpsertOne {
	cc.conflict = opts
	return &CollectUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Collect.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CollectCreate) OnConflictColumns(columns ...string) *CollectUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CollectUpsertOne{
		create: cc,
	}
}

type (
	// CollectUpsertOne is the builder for "upsert"-ing
	//  one Collect node.
	CollectUpsertOne struct {
		create *CollectCreate
	}

	// CollectUpsert is the "OnConflict" setter.
	CollectUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *CollectUpsert) SetCreatedBy(v int64) *CollectUpsert {
	u.Set(collect.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *CollectUpsert) UpdateCreatedBy() *CollectUpsert {
	u.SetExcluded(collect.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *CollectUpsert) AddCreatedBy(v int64) *CollectUpsert {
	u.Add(collect.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *CollectUpsert) SetUpdatedBy(v int64) *CollectUpsert {
	u.Set(collect.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *CollectUpsert) UpdateUpdatedBy() *CollectUpsert {
	u.SetExcluded(collect.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *CollectUpsert) AddUpdatedBy(v int64) *CollectUpsert {
	u.Add(collect.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CollectUpsert) SetUpdatedAt(v time.Time) *CollectUpsert {
	u.Set(collect.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CollectUpsert) UpdateUpdatedAt() *CollectUpsert {
	u.SetExcluded(collect.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CollectUpsert) SetDeletedAt(v time.Time) *CollectUpsert {
	u.Set(collect.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CollectUpsert) UpdateDeletedAt() *CollectUpsert {
	u.SetExcluded(collect.FieldDeletedAt)
	return u
}

// SetURL sets the "url" field.
func (u *CollectUpsert) SetURL(v string) *CollectUpsert {
	u.Set(collect.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *CollectUpsert) UpdateURL() *CollectUpsert {
	u.SetExcluded(collect.FieldURL)
	return u
}

// SetUserID sets the "user_id" field.
func (u *CollectUpsert) SetUserID(v int64) *CollectUpsert {
	u.Set(collect.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CollectUpsert) UpdateUserID() *CollectUpsert {
	u.SetExcluded(collect.FieldUserID)
	return u
}

// SetJpgName sets the "jpg_name" field.
func (u *CollectUpsert) SetJpgName(v int64) *CollectUpsert {
	u.Set(collect.FieldJpgName, v)
	return u
}

// UpdateJpgName sets the "jpg_name" field to the value that was provided on create.
func (u *CollectUpsert) UpdateJpgName() *CollectUpsert {
	u.SetExcluded(collect.FieldJpgName)
	return u
}

// AddJpgName adds v to the "jpg_name" field.
func (u *CollectUpsert) AddJpgName(v int64) *CollectUpsert {
	u.Add(collect.FieldJpgName, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Collect.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(collect.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CollectUpsertOne) UpdateNewValues() *CollectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(collect.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(collect.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Collect.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CollectUpsertOne) Ignore() *CollectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CollectUpsertOne) DoNothing() *CollectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CollectCreate.OnConflict
// documentation for more info.
func (u *CollectUpsertOne) Update(set func(*CollectUpsert)) *CollectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CollectUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *CollectUpsertOne) SetCreatedBy(v int64) *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *CollectUpsertOne) AddCreatedBy(v int64) *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *CollectUpsertOne) UpdateCreatedBy() *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *CollectUpsertOne) SetUpdatedBy(v int64) *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *CollectUpsertOne) AddUpdatedBy(v int64) *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *CollectUpsertOne) UpdateUpdatedBy() *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CollectUpsertOne) SetUpdatedAt(v time.Time) *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CollectUpsertOne) UpdateUpdatedAt() *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CollectUpsertOne) SetDeletedAt(v time.Time) *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CollectUpsertOne) UpdateDeletedAt() *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetURL sets the "url" field.
func (u *CollectUpsertOne) SetURL(v string) *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *CollectUpsertOne) UpdateURL() *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateURL()
	})
}

// SetUserID sets the "user_id" field.
func (u *CollectUpsertOne) SetUserID(v int64) *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CollectUpsertOne) UpdateUserID() *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateUserID()
	})
}

// SetJpgName sets the "jpg_name" field.
func (u *CollectUpsertOne) SetJpgName(v int64) *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.SetJpgName(v)
	})
}

// AddJpgName adds v to the "jpg_name" field.
func (u *CollectUpsertOne) AddJpgName(v int64) *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.AddJpgName(v)
	})
}

// UpdateJpgName sets the "jpg_name" field to the value that was provided on create.
func (u *CollectUpsertOne) UpdateJpgName() *CollectUpsertOne {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateJpgName()
	})
}

// Exec executes the query.
func (u *CollectUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for CollectCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CollectUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CollectUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CollectUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CollectCreateBulk is the builder for creating many Collect entities in bulk.
type CollectCreateBulk struct {
	config
	err      error
	builders []*CollectCreate
	conflict []sql.ConflictOption
}

// Save creates the Collect entities in the database.
func (ccb *CollectCreateBulk) Save(ctx context.Context) ([]*Collect, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Collect, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CollectMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CollectCreateBulk) SaveX(ctx context.Context) []*Collect {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CollectCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CollectCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Collect.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CollectUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (ccb *CollectCreateBulk) OnConflict(opts ...sql.ConflictOption) *CollectUpsertBulk {
	ccb.conflict = opts
	return &CollectUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Collect.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CollectCreateBulk) OnConflictColumns(columns ...string) *CollectUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CollectUpsertBulk{
		create: ccb,
	}
}

// CollectUpsertBulk is the builder for "upsert"-ing
// a bulk of Collect nodes.
type CollectUpsertBulk struct {
	create *CollectCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Collect.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(collect.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CollectUpsertBulk) UpdateNewValues() *CollectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(collect.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(collect.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Collect.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CollectUpsertBulk) Ignore() *CollectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CollectUpsertBulk) DoNothing() *CollectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CollectCreateBulk.OnConflict
// documentation for more info.
func (u *CollectUpsertBulk) Update(set func(*CollectUpsert)) *CollectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CollectUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *CollectUpsertBulk) SetCreatedBy(v int64) *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *CollectUpsertBulk) AddCreatedBy(v int64) *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *CollectUpsertBulk) UpdateCreatedBy() *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *CollectUpsertBulk) SetUpdatedBy(v int64) *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *CollectUpsertBulk) AddUpdatedBy(v int64) *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *CollectUpsertBulk) UpdateUpdatedBy() *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CollectUpsertBulk) SetUpdatedAt(v time.Time) *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CollectUpsertBulk) UpdateUpdatedAt() *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CollectUpsertBulk) SetDeletedAt(v time.Time) *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CollectUpsertBulk) UpdateDeletedAt() *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetURL sets the "url" field.
func (u *CollectUpsertBulk) SetURL(v string) *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *CollectUpsertBulk) UpdateURL() *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateURL()
	})
}

// SetUserID sets the "user_id" field.
func (u *CollectUpsertBulk) SetUserID(v int64) *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CollectUpsertBulk) UpdateUserID() *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateUserID()
	})
}

// SetJpgName sets the "jpg_name" field.
func (u *CollectUpsertBulk) SetJpgName(v int64) *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.SetJpgName(v)
	})
}

// AddJpgName adds v to the "jpg_name" field.
func (u *CollectUpsertBulk) AddJpgName(v int64) *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.AddJpgName(v)
	})
}

// UpdateJpgName sets the "jpg_name" field to the value that was provided on create.
func (u *CollectUpsertBulk) UpdateJpgName() *CollectUpsertBulk {
	return u.Update(func(s *CollectUpsert) {
		s.UpdateJpgName()
	})
}

// Exec executes the query.
func (u *CollectUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the CollectCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for CollectCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CollectUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
