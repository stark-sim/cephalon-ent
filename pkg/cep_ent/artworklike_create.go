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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/artwork"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/artworklike"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// ArtworkLikeCreate is the builder for creating a ArtworkLike entity.
type ArtworkLikeCreate struct {
	config
	mutation *ArtworkLikeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (alc *ArtworkLikeCreate) SetCreatedBy(i int64) *ArtworkLikeCreate {
	alc.mutation.SetCreatedBy(i)
	return alc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (alc *ArtworkLikeCreate) SetNillableCreatedBy(i *int64) *ArtworkLikeCreate {
	if i != nil {
		alc.SetCreatedBy(*i)
	}
	return alc
}

// SetUpdatedBy sets the "updated_by" field.
func (alc *ArtworkLikeCreate) SetUpdatedBy(i int64) *ArtworkLikeCreate {
	alc.mutation.SetUpdatedBy(i)
	return alc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (alc *ArtworkLikeCreate) SetNillableUpdatedBy(i *int64) *ArtworkLikeCreate {
	if i != nil {
		alc.SetUpdatedBy(*i)
	}
	return alc
}

// SetCreatedAt sets the "created_at" field.
func (alc *ArtworkLikeCreate) SetCreatedAt(t time.Time) *ArtworkLikeCreate {
	alc.mutation.SetCreatedAt(t)
	return alc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (alc *ArtworkLikeCreate) SetNillableCreatedAt(t *time.Time) *ArtworkLikeCreate {
	if t != nil {
		alc.SetCreatedAt(*t)
	}
	return alc
}

// SetUpdatedAt sets the "updated_at" field.
func (alc *ArtworkLikeCreate) SetUpdatedAt(t time.Time) *ArtworkLikeCreate {
	alc.mutation.SetUpdatedAt(t)
	return alc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (alc *ArtworkLikeCreate) SetNillableUpdatedAt(t *time.Time) *ArtworkLikeCreate {
	if t != nil {
		alc.SetUpdatedAt(*t)
	}
	return alc
}

// SetDeletedAt sets the "deleted_at" field.
func (alc *ArtworkLikeCreate) SetDeletedAt(t time.Time) *ArtworkLikeCreate {
	alc.mutation.SetDeletedAt(t)
	return alc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (alc *ArtworkLikeCreate) SetNillableDeletedAt(t *time.Time) *ArtworkLikeCreate {
	if t != nil {
		alc.SetDeletedAt(*t)
	}
	return alc
}

// SetUserID sets the "user_id" field.
func (alc *ArtworkLikeCreate) SetUserID(i int64) *ArtworkLikeCreate {
	alc.mutation.SetUserID(i)
	return alc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (alc *ArtworkLikeCreate) SetNillableUserID(i *int64) *ArtworkLikeCreate {
	if i != nil {
		alc.SetUserID(*i)
	}
	return alc
}

// SetArtworkID sets the "artwork_id" field.
func (alc *ArtworkLikeCreate) SetArtworkID(i int64) *ArtworkLikeCreate {
	alc.mutation.SetArtworkID(i)
	return alc
}

// SetNillableArtworkID sets the "artwork_id" field if the given value is not nil.
func (alc *ArtworkLikeCreate) SetNillableArtworkID(i *int64) *ArtworkLikeCreate {
	if i != nil {
		alc.SetArtworkID(*i)
	}
	return alc
}

// SetDate sets the "date" field.
func (alc *ArtworkLikeCreate) SetDate(i int32) *ArtworkLikeCreate {
	alc.mutation.SetDate(i)
	return alc
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (alc *ArtworkLikeCreate) SetNillableDate(i *int32) *ArtworkLikeCreate {
	if i != nil {
		alc.SetDate(*i)
	}
	return alc
}

// SetID sets the "id" field.
func (alc *ArtworkLikeCreate) SetID(i int64) *ArtworkLikeCreate {
	alc.mutation.SetID(i)
	return alc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (alc *ArtworkLikeCreate) SetNillableID(i *int64) *ArtworkLikeCreate {
	if i != nil {
		alc.SetID(*i)
	}
	return alc
}

// SetUser sets the "user" edge to the User entity.
func (alc *ArtworkLikeCreate) SetUser(u *User) *ArtworkLikeCreate {
	return alc.SetUserID(u.ID)
}

// SetArtwork sets the "artwork" edge to the Artwork entity.
func (alc *ArtworkLikeCreate) SetArtwork(a *Artwork) *ArtworkLikeCreate {
	return alc.SetArtworkID(a.ID)
}

// Mutation returns the ArtworkLikeMutation object of the builder.
func (alc *ArtworkLikeCreate) Mutation() *ArtworkLikeMutation {
	return alc.mutation
}

// Save creates the ArtworkLike in the database.
func (alc *ArtworkLikeCreate) Save(ctx context.Context) (*ArtworkLike, error) {
	alc.defaults()
	return withHooks(ctx, alc.sqlSave, alc.mutation, alc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (alc *ArtworkLikeCreate) SaveX(ctx context.Context) *ArtworkLike {
	v, err := alc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alc *ArtworkLikeCreate) Exec(ctx context.Context) error {
	_, err := alc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alc *ArtworkLikeCreate) ExecX(ctx context.Context) {
	if err := alc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (alc *ArtworkLikeCreate) defaults() {
	if _, ok := alc.mutation.CreatedBy(); !ok {
		v := artworklike.DefaultCreatedBy
		alc.mutation.SetCreatedBy(v)
	}
	if _, ok := alc.mutation.UpdatedBy(); !ok {
		v := artworklike.DefaultUpdatedBy
		alc.mutation.SetUpdatedBy(v)
	}
	if _, ok := alc.mutation.CreatedAt(); !ok {
		v := artworklike.DefaultCreatedAt()
		alc.mutation.SetCreatedAt(v)
	}
	if _, ok := alc.mutation.UpdatedAt(); !ok {
		v := artworklike.DefaultUpdatedAt()
		alc.mutation.SetUpdatedAt(v)
	}
	if _, ok := alc.mutation.DeletedAt(); !ok {
		v := artworklike.DefaultDeletedAt
		alc.mutation.SetDeletedAt(v)
	}
	if _, ok := alc.mutation.UserID(); !ok {
		v := artworklike.DefaultUserID
		alc.mutation.SetUserID(v)
	}
	if _, ok := alc.mutation.ArtworkID(); !ok {
		v := artworklike.DefaultArtworkID
		alc.mutation.SetArtworkID(v)
	}
	if _, ok := alc.mutation.Date(); !ok {
		v := artworklike.DefaultDate
		alc.mutation.SetDate(v)
	}
	if _, ok := alc.mutation.ID(); !ok {
		v := artworklike.DefaultID()
		alc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (alc *ArtworkLikeCreate) check() error {
	if _, ok := alc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "ArtworkLike.created_by"`)}
	}
	if _, ok := alc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "ArtworkLike.updated_by"`)}
	}
	if _, ok := alc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "ArtworkLike.created_at"`)}
	}
	if _, ok := alc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "ArtworkLike.updated_at"`)}
	}
	if _, ok := alc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "ArtworkLike.deleted_at"`)}
	}
	if _, ok := alc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`cep_ent: missing required field "ArtworkLike.user_id"`)}
	}
	if _, ok := alc.mutation.ArtworkID(); !ok {
		return &ValidationError{Name: "artwork_id", err: errors.New(`cep_ent: missing required field "ArtworkLike.artwork_id"`)}
	}
	if _, ok := alc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`cep_ent: missing required field "ArtworkLike.date"`)}
	}
	if _, ok := alc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`cep_ent: missing required edge "ArtworkLike.user"`)}
	}
	if _, ok := alc.mutation.ArtworkID(); !ok {
		return &ValidationError{Name: "artwork", err: errors.New(`cep_ent: missing required edge "ArtworkLike.artwork"`)}
	}
	return nil
}

func (alc *ArtworkLikeCreate) sqlSave(ctx context.Context) (*ArtworkLike, error) {
	if err := alc.check(); err != nil {
		return nil, err
	}
	_node, _spec := alc.createSpec()
	if err := sqlgraph.CreateNode(ctx, alc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	alc.mutation.id = &_node.ID
	alc.mutation.done = true
	return _node, nil
}

func (alc *ArtworkLikeCreate) createSpec() (*ArtworkLike, *sqlgraph.CreateSpec) {
	var (
		_node = &ArtworkLike{config: alc.config}
		_spec = sqlgraph.NewCreateSpec(artworklike.Table, sqlgraph.NewFieldSpec(artworklike.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = alc.conflict
	if id, ok := alc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := alc.mutation.CreatedBy(); ok {
		_spec.SetField(artworklike.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := alc.mutation.UpdatedBy(); ok {
		_spec.SetField(artworklike.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := alc.mutation.CreatedAt(); ok {
		_spec.SetField(artworklike.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := alc.mutation.UpdatedAt(); ok {
		_spec.SetField(artworklike.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := alc.mutation.DeletedAt(); ok {
		_spec.SetField(artworklike.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := alc.mutation.Date(); ok {
		_spec.SetField(artworklike.FieldDate, field.TypeInt32, value)
		_node.Date = value
	}
	if nodes := alc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artworklike.UserTable,
			Columns: []string{artworklike.UserColumn},
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
	if nodes := alc.mutation.ArtworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artworklike.ArtworkTable,
			Columns: []string{artworklike.ArtworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artwork.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ArtworkID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ArtworkLike.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArtworkLikeUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (alc *ArtworkLikeCreate) OnConflict(opts ...sql.ConflictOption) *ArtworkLikeUpsertOne {
	alc.conflict = opts
	return &ArtworkLikeUpsertOne{
		create: alc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ArtworkLike.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (alc *ArtworkLikeCreate) OnConflictColumns(columns ...string) *ArtworkLikeUpsertOne {
	alc.conflict = append(alc.conflict, sql.ConflictColumns(columns...))
	return &ArtworkLikeUpsertOne{
		create: alc,
	}
}

type (
	// ArtworkLikeUpsertOne is the builder for "upsert"-ing
	//  one ArtworkLike node.
	ArtworkLikeUpsertOne struct {
		create *ArtworkLikeCreate
	}

	// ArtworkLikeUpsert is the "OnConflict" setter.
	ArtworkLikeUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *ArtworkLikeUpsert) SetCreatedBy(v int64) *ArtworkLikeUpsert {
	u.Set(artworklike.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *ArtworkLikeUpsert) UpdateCreatedBy() *ArtworkLikeUpsert {
	u.SetExcluded(artworklike.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *ArtworkLikeUpsert) AddCreatedBy(v int64) *ArtworkLikeUpsert {
	u.Add(artworklike.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *ArtworkLikeUpsert) SetUpdatedBy(v int64) *ArtworkLikeUpsert {
	u.Set(artworklike.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *ArtworkLikeUpsert) UpdateUpdatedBy() *ArtworkLikeUpsert {
	u.SetExcluded(artworklike.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *ArtworkLikeUpsert) AddUpdatedBy(v int64) *ArtworkLikeUpsert {
	u.Add(artworklike.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArtworkLikeUpsert) SetUpdatedAt(v time.Time) *ArtworkLikeUpsert {
	u.Set(artworklike.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArtworkLikeUpsert) UpdateUpdatedAt() *ArtworkLikeUpsert {
	u.SetExcluded(artworklike.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ArtworkLikeUpsert) SetDeletedAt(v time.Time) *ArtworkLikeUpsert {
	u.Set(artworklike.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ArtworkLikeUpsert) UpdateDeletedAt() *ArtworkLikeUpsert {
	u.SetExcluded(artworklike.FieldDeletedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *ArtworkLikeUpsert) SetUserID(v int64) *ArtworkLikeUpsert {
	u.Set(artworklike.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ArtworkLikeUpsert) UpdateUserID() *ArtworkLikeUpsert {
	u.SetExcluded(artworklike.FieldUserID)
	return u
}

// SetArtworkID sets the "artwork_id" field.
func (u *ArtworkLikeUpsert) SetArtworkID(v int64) *ArtworkLikeUpsert {
	u.Set(artworklike.FieldArtworkID, v)
	return u
}

// UpdateArtworkID sets the "artwork_id" field to the value that was provided on create.
func (u *ArtworkLikeUpsert) UpdateArtworkID() *ArtworkLikeUpsert {
	u.SetExcluded(artworklike.FieldArtworkID)
	return u
}

// SetDate sets the "date" field.
func (u *ArtworkLikeUpsert) SetDate(v int32) *ArtworkLikeUpsert {
	u.Set(artworklike.FieldDate, v)
	return u
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *ArtworkLikeUpsert) UpdateDate() *ArtworkLikeUpsert {
	u.SetExcluded(artworklike.FieldDate)
	return u
}

// AddDate adds v to the "date" field.
func (u *ArtworkLikeUpsert) AddDate(v int32) *ArtworkLikeUpsert {
	u.Add(artworklike.FieldDate, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ArtworkLike.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(artworklike.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ArtworkLikeUpsertOne) UpdateNewValues() *ArtworkLikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(artworklike.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(artworklike.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ArtworkLike.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ArtworkLikeUpsertOne) Ignore() *ArtworkLikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArtworkLikeUpsertOne) DoNothing() *ArtworkLikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArtworkLikeCreate.OnConflict
// documentation for more info.
func (u *ArtworkLikeUpsertOne) Update(set func(*ArtworkLikeUpsert)) *ArtworkLikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArtworkLikeUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *ArtworkLikeUpsertOne) SetCreatedBy(v int64) *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *ArtworkLikeUpsertOne) AddCreatedBy(v int64) *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *ArtworkLikeUpsertOne) UpdateCreatedBy() *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *ArtworkLikeUpsertOne) SetUpdatedBy(v int64) *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *ArtworkLikeUpsertOne) AddUpdatedBy(v int64) *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *ArtworkLikeUpsertOne) UpdateUpdatedBy() *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArtworkLikeUpsertOne) SetUpdatedAt(v time.Time) *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArtworkLikeUpsertOne) UpdateUpdatedAt() *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ArtworkLikeUpsertOne) SetDeletedAt(v time.Time) *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ArtworkLikeUpsertOne) UpdateDeletedAt() *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *ArtworkLikeUpsertOne) SetUserID(v int64) *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ArtworkLikeUpsertOne) UpdateUserID() *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateUserID()
	})
}

// SetArtworkID sets the "artwork_id" field.
func (u *ArtworkLikeUpsertOne) SetArtworkID(v int64) *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetArtworkID(v)
	})
}

// UpdateArtworkID sets the "artwork_id" field to the value that was provided on create.
func (u *ArtworkLikeUpsertOne) UpdateArtworkID() *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateArtworkID()
	})
}

// SetDate sets the "date" field.
func (u *ArtworkLikeUpsertOne) SetDate(v int32) *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetDate(v)
	})
}

// AddDate adds v to the "date" field.
func (u *ArtworkLikeUpsertOne) AddDate(v int32) *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.AddDate(v)
	})
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *ArtworkLikeUpsertOne) UpdateDate() *ArtworkLikeUpsertOne {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateDate()
	})
}

// Exec executes the query.
func (u *ArtworkLikeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for ArtworkLikeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArtworkLikeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ArtworkLikeUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ArtworkLikeUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ArtworkLikeCreateBulk is the builder for creating many ArtworkLike entities in bulk.
type ArtworkLikeCreateBulk struct {
	config
	err      error
	builders []*ArtworkLikeCreate
	conflict []sql.ConflictOption
}

// Save creates the ArtworkLike entities in the database.
func (alcb *ArtworkLikeCreateBulk) Save(ctx context.Context) ([]*ArtworkLike, error) {
	if alcb.err != nil {
		return nil, alcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(alcb.builders))
	nodes := make([]*ArtworkLike, len(alcb.builders))
	mutators := make([]Mutator, len(alcb.builders))
	for i := range alcb.builders {
		func(i int, root context.Context) {
			builder := alcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArtworkLikeMutation)
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
					_, err = mutators[i+1].Mutate(root, alcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = alcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, alcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, alcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (alcb *ArtworkLikeCreateBulk) SaveX(ctx context.Context) []*ArtworkLike {
	v, err := alcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alcb *ArtworkLikeCreateBulk) Exec(ctx context.Context) error {
	_, err := alcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alcb *ArtworkLikeCreateBulk) ExecX(ctx context.Context) {
	if err := alcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ArtworkLike.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArtworkLikeUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (alcb *ArtworkLikeCreateBulk) OnConflict(opts ...sql.ConflictOption) *ArtworkLikeUpsertBulk {
	alcb.conflict = opts
	return &ArtworkLikeUpsertBulk{
		create: alcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ArtworkLike.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (alcb *ArtworkLikeCreateBulk) OnConflictColumns(columns ...string) *ArtworkLikeUpsertBulk {
	alcb.conflict = append(alcb.conflict, sql.ConflictColumns(columns...))
	return &ArtworkLikeUpsertBulk{
		create: alcb,
	}
}

// ArtworkLikeUpsertBulk is the builder for "upsert"-ing
// a bulk of ArtworkLike nodes.
type ArtworkLikeUpsertBulk struct {
	create *ArtworkLikeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ArtworkLike.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(artworklike.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ArtworkLikeUpsertBulk) UpdateNewValues() *ArtworkLikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(artworklike.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(artworklike.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ArtworkLike.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ArtworkLikeUpsertBulk) Ignore() *ArtworkLikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArtworkLikeUpsertBulk) DoNothing() *ArtworkLikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArtworkLikeCreateBulk.OnConflict
// documentation for more info.
func (u *ArtworkLikeUpsertBulk) Update(set func(*ArtworkLikeUpsert)) *ArtworkLikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArtworkLikeUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *ArtworkLikeUpsertBulk) SetCreatedBy(v int64) *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *ArtworkLikeUpsertBulk) AddCreatedBy(v int64) *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *ArtworkLikeUpsertBulk) UpdateCreatedBy() *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *ArtworkLikeUpsertBulk) SetUpdatedBy(v int64) *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *ArtworkLikeUpsertBulk) AddUpdatedBy(v int64) *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *ArtworkLikeUpsertBulk) UpdateUpdatedBy() *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArtworkLikeUpsertBulk) SetUpdatedAt(v time.Time) *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArtworkLikeUpsertBulk) UpdateUpdatedAt() *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ArtworkLikeUpsertBulk) SetDeletedAt(v time.Time) *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ArtworkLikeUpsertBulk) UpdateDeletedAt() *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *ArtworkLikeUpsertBulk) SetUserID(v int64) *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ArtworkLikeUpsertBulk) UpdateUserID() *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateUserID()
	})
}

// SetArtworkID sets the "artwork_id" field.
func (u *ArtworkLikeUpsertBulk) SetArtworkID(v int64) *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetArtworkID(v)
	})
}

// UpdateArtworkID sets the "artwork_id" field to the value that was provided on create.
func (u *ArtworkLikeUpsertBulk) UpdateArtworkID() *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateArtworkID()
	})
}

// SetDate sets the "date" field.
func (u *ArtworkLikeUpsertBulk) SetDate(v int32) *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.SetDate(v)
	})
}

// AddDate adds v to the "date" field.
func (u *ArtworkLikeUpsertBulk) AddDate(v int32) *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.AddDate(v)
	})
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *ArtworkLikeUpsertBulk) UpdateDate() *ArtworkLikeUpsertBulk {
	return u.Update(func(s *ArtworkLikeUpsert) {
		s.UpdateDate()
	})
}

// Exec executes the query.
func (u *ArtworkLikeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the ArtworkLikeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for ArtworkLikeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArtworkLikeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
