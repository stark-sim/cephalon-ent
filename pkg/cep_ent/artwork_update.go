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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// ArtworkUpdate is the builder for updating Artwork entities.
type ArtworkUpdate struct {
	config
	hooks     []Hook
	mutation  *ArtworkMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ArtworkUpdate builder.
func (au *ArtworkUpdate) Where(ps ...predicate.Artwork) *ArtworkUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedBy sets the "created_by" field.
func (au *ArtworkUpdate) SetCreatedBy(i int64) *ArtworkUpdate {
	au.mutation.ResetCreatedBy()
	au.mutation.SetCreatedBy(i)
	return au
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (au *ArtworkUpdate) SetNillableCreatedBy(i *int64) *ArtworkUpdate {
	if i != nil {
		au.SetCreatedBy(*i)
	}
	return au
}

// AddCreatedBy adds i to the "created_by" field.
func (au *ArtworkUpdate) AddCreatedBy(i int64) *ArtworkUpdate {
	au.mutation.AddCreatedBy(i)
	return au
}

// SetUpdatedBy sets the "updated_by" field.
func (au *ArtworkUpdate) SetUpdatedBy(i int64) *ArtworkUpdate {
	au.mutation.ResetUpdatedBy()
	au.mutation.SetUpdatedBy(i)
	return au
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (au *ArtworkUpdate) SetNillableUpdatedBy(i *int64) *ArtworkUpdate {
	if i != nil {
		au.SetUpdatedBy(*i)
	}
	return au
}

// AddUpdatedBy adds i to the "updated_by" field.
func (au *ArtworkUpdate) AddUpdatedBy(i int64) *ArtworkUpdate {
	au.mutation.AddUpdatedBy(i)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ArtworkUpdate) SetUpdatedAt(t time.Time) *ArtworkUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *ArtworkUpdate) SetDeletedAt(t time.Time) *ArtworkUpdate {
	au.mutation.SetDeletedAt(t)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *ArtworkUpdate) SetNillableDeletedAt(t *time.Time) *ArtworkUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// SetName sets the "name" field.
func (au *ArtworkUpdate) SetName(s string) *ArtworkUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *ArtworkUpdate) SetNillableName(s *string) *ArtworkUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetURL sets the "url" field.
func (au *ArtworkUpdate) SetURL(s string) *ArtworkUpdate {
	au.mutation.SetURL(s)
	return au
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (au *ArtworkUpdate) SetNillableURL(s *string) *ArtworkUpdate {
	if s != nil {
		au.SetURL(*s)
	}
	return au
}

// SetAuthorID sets the "author_id" field.
func (au *ArtworkUpdate) SetAuthorID(i int64) *ArtworkUpdate {
	au.mutation.SetAuthorID(i)
	return au
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (au *ArtworkUpdate) SetNillableAuthorID(i *int64) *ArtworkUpdate {
	if i != nil {
		au.SetAuthorID(*i)
	}
	return au
}

// SetAuthor sets the "author" edge to the User entity.
func (au *ArtworkUpdate) SetAuthor(u *User) *ArtworkUpdate {
	return au.SetAuthorID(u.ID)
}

// AddArtworkLikeIDs adds the "artwork_likes" edge to the ArtworkLike entity by IDs.
func (au *ArtworkUpdate) AddArtworkLikeIDs(ids ...int64) *ArtworkUpdate {
	au.mutation.AddArtworkLikeIDs(ids...)
	return au
}

// AddArtworkLikes adds the "artwork_likes" edges to the ArtworkLike entity.
func (au *ArtworkUpdate) AddArtworkLikes(a ...*ArtworkLike) *ArtworkUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddArtworkLikeIDs(ids...)
}

// Mutation returns the ArtworkMutation object of the builder.
func (au *ArtworkUpdate) Mutation() *ArtworkMutation {
	return au.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (au *ArtworkUpdate) ClearAuthor() *ArtworkUpdate {
	au.mutation.ClearAuthor()
	return au
}

// ClearArtworkLikes clears all "artwork_likes" edges to the ArtworkLike entity.
func (au *ArtworkUpdate) ClearArtworkLikes() *ArtworkUpdate {
	au.mutation.ClearArtworkLikes()
	return au
}

// RemoveArtworkLikeIDs removes the "artwork_likes" edge to ArtworkLike entities by IDs.
func (au *ArtworkUpdate) RemoveArtworkLikeIDs(ids ...int64) *ArtworkUpdate {
	au.mutation.RemoveArtworkLikeIDs(ids...)
	return au
}

// RemoveArtworkLikes removes "artwork_likes" edges to ArtworkLike entities.
func (au *ArtworkUpdate) RemoveArtworkLikes(a ...*ArtworkLike) *ArtworkUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveArtworkLikeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArtworkUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArtworkUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArtworkUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArtworkUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *ArtworkUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := artwork.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ArtworkUpdate) check() error {
	if _, ok := au.mutation.AuthorID(); au.mutation.AuthorCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "Artwork.author"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *ArtworkUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ArtworkUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *ArtworkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(artwork.Table, artwork.Columns, sqlgraph.NewFieldSpec(artwork.FieldID, field.TypeInt64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedBy(); ok {
		_spec.SetField(artwork.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedCreatedBy(); ok {
		_spec.AddField(artwork.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := au.mutation.UpdatedBy(); ok {
		_spec.SetField(artwork.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(artwork.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(artwork.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(artwork.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(artwork.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.URL(); ok {
		_spec.SetField(artwork.FieldURL, field.TypeString, value)
	}
	if au.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artwork.AuthorTable,
			Columns: []string{artwork.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artwork.AuthorTable,
			Columns: []string{artwork.AuthorColumn},
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
	if au.mutation.ArtworkLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artwork.ArtworkLikesTable,
			Columns: []string{artwork.ArtworkLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artworklike.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedArtworkLikesIDs(); len(nodes) > 0 && !au.mutation.ArtworkLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artwork.ArtworkLikesTable,
			Columns: []string{artwork.ArtworkLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artworklike.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ArtworkLikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artwork.ArtworkLikesTable,
			Columns: []string{artwork.ArtworkLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artworklike.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artwork.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArtworkUpdateOne is the builder for updating a single Artwork entity.
type ArtworkUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ArtworkMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (auo *ArtworkUpdateOne) SetCreatedBy(i int64) *ArtworkUpdateOne {
	auo.mutation.ResetCreatedBy()
	auo.mutation.SetCreatedBy(i)
	return auo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (auo *ArtworkUpdateOne) SetNillableCreatedBy(i *int64) *ArtworkUpdateOne {
	if i != nil {
		auo.SetCreatedBy(*i)
	}
	return auo
}

// AddCreatedBy adds i to the "created_by" field.
func (auo *ArtworkUpdateOne) AddCreatedBy(i int64) *ArtworkUpdateOne {
	auo.mutation.AddCreatedBy(i)
	return auo
}

// SetUpdatedBy sets the "updated_by" field.
func (auo *ArtworkUpdateOne) SetUpdatedBy(i int64) *ArtworkUpdateOne {
	auo.mutation.ResetUpdatedBy()
	auo.mutation.SetUpdatedBy(i)
	return auo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (auo *ArtworkUpdateOne) SetNillableUpdatedBy(i *int64) *ArtworkUpdateOne {
	if i != nil {
		auo.SetUpdatedBy(*i)
	}
	return auo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (auo *ArtworkUpdateOne) AddUpdatedBy(i int64) *ArtworkUpdateOne {
	auo.mutation.AddUpdatedBy(i)
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ArtworkUpdateOne) SetUpdatedAt(t time.Time) *ArtworkUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *ArtworkUpdateOne) SetDeletedAt(t time.Time) *ArtworkUpdateOne {
	auo.mutation.SetDeletedAt(t)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *ArtworkUpdateOne) SetNillableDeletedAt(t *time.Time) *ArtworkUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// SetName sets the "name" field.
func (auo *ArtworkUpdateOne) SetName(s string) *ArtworkUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *ArtworkUpdateOne) SetNillableName(s *string) *ArtworkUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetURL sets the "url" field.
func (auo *ArtworkUpdateOne) SetURL(s string) *ArtworkUpdateOne {
	auo.mutation.SetURL(s)
	return auo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (auo *ArtworkUpdateOne) SetNillableURL(s *string) *ArtworkUpdateOne {
	if s != nil {
		auo.SetURL(*s)
	}
	return auo
}

// SetAuthorID sets the "author_id" field.
func (auo *ArtworkUpdateOne) SetAuthorID(i int64) *ArtworkUpdateOne {
	auo.mutation.SetAuthorID(i)
	return auo
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (auo *ArtworkUpdateOne) SetNillableAuthorID(i *int64) *ArtworkUpdateOne {
	if i != nil {
		auo.SetAuthorID(*i)
	}
	return auo
}

// SetAuthor sets the "author" edge to the User entity.
func (auo *ArtworkUpdateOne) SetAuthor(u *User) *ArtworkUpdateOne {
	return auo.SetAuthorID(u.ID)
}

// AddArtworkLikeIDs adds the "artwork_likes" edge to the ArtworkLike entity by IDs.
func (auo *ArtworkUpdateOne) AddArtworkLikeIDs(ids ...int64) *ArtworkUpdateOne {
	auo.mutation.AddArtworkLikeIDs(ids...)
	return auo
}

// AddArtworkLikes adds the "artwork_likes" edges to the ArtworkLike entity.
func (auo *ArtworkUpdateOne) AddArtworkLikes(a ...*ArtworkLike) *ArtworkUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddArtworkLikeIDs(ids...)
}

// Mutation returns the ArtworkMutation object of the builder.
func (auo *ArtworkUpdateOne) Mutation() *ArtworkMutation {
	return auo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (auo *ArtworkUpdateOne) ClearAuthor() *ArtworkUpdateOne {
	auo.mutation.ClearAuthor()
	return auo
}

// ClearArtworkLikes clears all "artwork_likes" edges to the ArtworkLike entity.
func (auo *ArtworkUpdateOne) ClearArtworkLikes() *ArtworkUpdateOne {
	auo.mutation.ClearArtworkLikes()
	return auo
}

// RemoveArtworkLikeIDs removes the "artwork_likes" edge to ArtworkLike entities by IDs.
func (auo *ArtworkUpdateOne) RemoveArtworkLikeIDs(ids ...int64) *ArtworkUpdateOne {
	auo.mutation.RemoveArtworkLikeIDs(ids...)
	return auo
}

// RemoveArtworkLikes removes "artwork_likes" edges to ArtworkLike entities.
func (auo *ArtworkUpdateOne) RemoveArtworkLikes(a ...*ArtworkLike) *ArtworkUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveArtworkLikeIDs(ids...)
}

// Where appends a list predicates to the ArtworkUpdate builder.
func (auo *ArtworkUpdateOne) Where(ps ...predicate.Artwork) *ArtworkUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArtworkUpdateOne) Select(field string, fields ...string) *ArtworkUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Artwork entity.
func (auo *ArtworkUpdateOne) Save(ctx context.Context) (*Artwork, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArtworkUpdateOne) SaveX(ctx context.Context) *Artwork {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArtworkUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArtworkUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *ArtworkUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := artwork.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ArtworkUpdateOne) check() error {
	if _, ok := auo.mutation.AuthorID(); auo.mutation.AuthorCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "Artwork.author"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *ArtworkUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ArtworkUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *ArtworkUpdateOne) sqlSave(ctx context.Context) (_node *Artwork, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(artwork.Table, artwork.Columns, sqlgraph.NewFieldSpec(artwork.FieldID, field.TypeInt64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "Artwork.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artwork.FieldID)
		for _, f := range fields {
			if !artwork.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != artwork.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedBy(); ok {
		_spec.SetField(artwork.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(artwork.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.UpdatedBy(); ok {
		_spec.SetField(artwork.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(artwork.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(artwork.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(artwork.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(artwork.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.URL(); ok {
		_spec.SetField(artwork.FieldURL, field.TypeString, value)
	}
	if auo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artwork.AuthorTable,
			Columns: []string{artwork.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artwork.AuthorTable,
			Columns: []string{artwork.AuthorColumn},
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
	if auo.mutation.ArtworkLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artwork.ArtworkLikesTable,
			Columns: []string{artwork.ArtworkLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artworklike.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedArtworkLikesIDs(); len(nodes) > 0 && !auo.mutation.ArtworkLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artwork.ArtworkLikesTable,
			Columns: []string{artwork.ArtworkLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artworklike.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ArtworkLikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artwork.ArtworkLikesTable,
			Columns: []string{artwork.ArtworkLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artworklike.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(auo.modifiers...)
	_node = &Artwork{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artwork.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
