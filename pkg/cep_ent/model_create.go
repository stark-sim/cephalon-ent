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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/modelprice"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/modlestar"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ModelCreate is the builder for creating a Model entity.
type ModelCreate struct {
	config
	mutation *ModelMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (mc *ModelCreate) SetCreatedBy(i int64) *ModelCreate {
	mc.mutation.SetCreatedBy(i)
	return mc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mc *ModelCreate) SetNillableCreatedBy(i *int64) *ModelCreate {
	if i != nil {
		mc.SetCreatedBy(*i)
	}
	return mc
}

// SetUpdatedBy sets the "updated_by" field.
func (mc *ModelCreate) SetUpdatedBy(i int64) *ModelCreate {
	mc.mutation.SetUpdatedBy(i)
	return mc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mc *ModelCreate) SetNillableUpdatedBy(i *int64) *ModelCreate {
	if i != nil {
		mc.SetUpdatedBy(*i)
	}
	return mc
}

// SetCreatedAt sets the "created_at" field.
func (mc *ModelCreate) SetCreatedAt(t time.Time) *ModelCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *ModelCreate) SetNillableCreatedAt(t *time.Time) *ModelCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *ModelCreate) SetUpdatedAt(t time.Time) *ModelCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *ModelCreate) SetNillableUpdatedAt(t *time.Time) *ModelCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetDeletedAt sets the "deleted_at" field.
func (mc *ModelCreate) SetDeletedAt(t time.Time) *ModelCreate {
	mc.mutation.SetDeletedAt(t)
	return mc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mc *ModelCreate) SetNillableDeletedAt(t *time.Time) *ModelCreate {
	if t != nil {
		mc.SetDeletedAt(*t)
	}
	return mc
}

// SetName sets the "name" field.
func (mc *ModelCreate) SetName(s string) *ModelCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mc *ModelCreate) SetNillableName(s *string) *ModelCreate {
	if s != nil {
		mc.SetName(*s)
	}
	return mc
}

// SetAuthor sets the "author" field.
func (mc *ModelCreate) SetAuthor(s string) *ModelCreate {
	mc.mutation.SetAuthor(s)
	return mc
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (mc *ModelCreate) SetNillableAuthor(s *string) *ModelCreate {
	if s != nil {
		mc.SetAuthor(*s)
	}
	return mc
}

// SetModelType sets the "model_type" field.
func (mc *ModelCreate) SetModelType(e enums.Model) *ModelCreate {
	mc.mutation.SetModelType(e)
	return mc
}

// SetNillableModelType sets the "model_type" field if the given value is not nil.
func (mc *ModelCreate) SetNillableModelType(e *enums.Model) *ModelCreate {
	if e != nil {
		mc.SetModelType(*e)
	}
	return mc
}

// SetModelStatus sets the "model_status" field.
func (mc *ModelCreate) SetModelStatus(es enums.ModelStatus) *ModelCreate {
	mc.mutation.SetModelStatus(es)
	return mc
}

// SetNillableModelStatus sets the "model_status" field if the given value is not nil.
func (mc *ModelCreate) SetNillableModelStatus(es *enums.ModelStatus) *ModelCreate {
	if es != nil {
		mc.SetModelStatus(*es)
	}
	return mc
}

// SetIsOfficial sets the "is_official" field.
func (mc *ModelCreate) SetIsOfficial(b bool) *ModelCreate {
	mc.mutation.SetIsOfficial(b)
	return mc
}

// SetNillableIsOfficial sets the "is_official" field if the given value is not nil.
func (mc *ModelCreate) SetNillableIsOfficial(b *bool) *ModelCreate {
	if b != nil {
		mc.SetIsOfficial(*b)
	}
	return mc
}

// SetTotalUsage sets the "total_usage" field.
func (mc *ModelCreate) SetTotalUsage(i int) *ModelCreate {
	mc.mutation.SetTotalUsage(i)
	return mc
}

// SetNillableTotalUsage sets the "total_usage" field if the given value is not nil.
func (mc *ModelCreate) SetNillableTotalUsage(i *int) *ModelCreate {
	if i != nil {
		mc.SetTotalUsage(*i)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *ModelCreate) SetID(i int64) *ModelCreate {
	mc.mutation.SetID(i)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *ModelCreate) SetNillableID(i *int64) *ModelCreate {
	if i != nil {
		mc.SetID(*i)
	}
	return mc
}

// AddModelPriceIDs adds the "model_prices" edge to the ModelPrice entity by IDs.
func (mc *ModelCreate) AddModelPriceIDs(ids ...int64) *ModelCreate {
	mc.mutation.AddModelPriceIDs(ids...)
	return mc
}

// AddModelPrices adds the "model_prices" edges to the ModelPrice entity.
func (mc *ModelCreate) AddModelPrices(m ...*ModelPrice) *ModelCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddModelPriceIDs(ids...)
}

// AddStarUserIDs adds the "star_user" edge to the User entity by IDs.
func (mc *ModelCreate) AddStarUserIDs(ids ...int64) *ModelCreate {
	mc.mutation.AddStarUserIDs(ids...)
	return mc
}

// AddStarUser adds the "star_user" edges to the User entity.
func (mc *ModelCreate) AddStarUser(u ...*User) *ModelCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mc.AddStarUserIDs(ids...)
}

// AddStarModelIDs adds the "star_model" edge to the ModleStar entity by IDs.
func (mc *ModelCreate) AddStarModelIDs(ids ...int64) *ModelCreate {
	mc.mutation.AddStarModelIDs(ids...)
	return mc
}

// AddStarModel adds the "star_model" edges to the ModleStar entity.
func (mc *ModelCreate) AddStarModel(m ...*ModleStar) *ModelCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddStarModelIDs(ids...)
}

// Mutation returns the ModelMutation object of the builder.
func (mc *ModelCreate) Mutation() *ModelMutation {
	return mc.mutation
}

// Save creates the Model in the database.
func (mc *ModelCreate) Save(ctx context.Context) (*Model, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *ModelCreate) SaveX(ctx context.Context) *Model {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *ModelCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *ModelCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *ModelCreate) defaults() {
	if _, ok := mc.mutation.CreatedBy(); !ok {
		v := model.DefaultCreatedBy
		mc.mutation.SetCreatedBy(v)
	}
	if _, ok := mc.mutation.UpdatedBy(); !ok {
		v := model.DefaultUpdatedBy
		mc.mutation.SetUpdatedBy(v)
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := model.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := model.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.DeletedAt(); !ok {
		v := model.DefaultDeletedAt
		mc.mutation.SetDeletedAt(v)
	}
	if _, ok := mc.mutation.Name(); !ok {
		v := model.DefaultName
		mc.mutation.SetName(v)
	}
	if _, ok := mc.mutation.Author(); !ok {
		v := model.DefaultAuthor
		mc.mutation.SetAuthor(v)
	}
	if _, ok := mc.mutation.ModelType(); !ok {
		v := model.DefaultModelType
		mc.mutation.SetModelType(v)
	}
	if _, ok := mc.mutation.ModelStatus(); !ok {
		v := model.DefaultModelStatus
		mc.mutation.SetModelStatus(v)
	}
	if _, ok := mc.mutation.IsOfficial(); !ok {
		v := model.DefaultIsOfficial
		mc.mutation.SetIsOfficial(v)
	}
	if _, ok := mc.mutation.TotalUsage(); !ok {
		v := model.DefaultTotalUsage
		mc.mutation.SetTotalUsage(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := model.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *ModelCreate) check() error {
	if _, ok := mc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "Model.created_by"`)}
	}
	if _, ok := mc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "Model.updated_by"`)}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "Model.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "Model.updated_at"`)}
	}
	if _, ok := mc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "Model.deleted_at"`)}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`cep_ent: missing required field "Model.name"`)}
	}
	if _, ok := mc.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`cep_ent: missing required field "Model.author"`)}
	}
	if _, ok := mc.mutation.ModelType(); !ok {
		return &ValidationError{Name: "model_type", err: errors.New(`cep_ent: missing required field "Model.model_type"`)}
	}
	if v, ok := mc.mutation.ModelType(); ok {
		if err := model.ModelTypeValidator(v); err != nil {
			return &ValidationError{Name: "model_type", err: fmt.Errorf(`cep_ent: validator failed for field "Model.model_type": %w`, err)}
		}
	}
	if _, ok := mc.mutation.ModelStatus(); !ok {
		return &ValidationError{Name: "model_status", err: errors.New(`cep_ent: missing required field "Model.model_status"`)}
	}
	if v, ok := mc.mutation.ModelStatus(); ok {
		if err := model.ModelStatusValidator(v); err != nil {
			return &ValidationError{Name: "model_status", err: fmt.Errorf(`cep_ent: validator failed for field "Model.model_status": %w`, err)}
		}
	}
	if _, ok := mc.mutation.IsOfficial(); !ok {
		return &ValidationError{Name: "is_official", err: errors.New(`cep_ent: missing required field "Model.is_official"`)}
	}
	if _, ok := mc.mutation.TotalUsage(); !ok {
		return &ValidationError{Name: "total_usage", err: errors.New(`cep_ent: missing required field "Model.total_usage"`)}
	}
	return nil
}

func (mc *ModelCreate) sqlSave(ctx context.Context) (*Model, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *ModelCreate) createSpec() (*Model, *sqlgraph.CreateSpec) {
	var (
		_node = &Model{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(model.Table, sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedBy(); ok {
		_spec.SetField(model.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := mc.mutation.UpdatedBy(); ok {
		_spec.SetField(model.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(model.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(model.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.DeletedAt(); ok {
		_spec.SetField(model.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(model.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.Author(); ok {
		_spec.SetField(model.FieldAuthor, field.TypeString, value)
		_node.Author = value
	}
	if value, ok := mc.mutation.ModelType(); ok {
		_spec.SetField(model.FieldModelType, field.TypeEnum, value)
		_node.ModelType = value
	}
	if value, ok := mc.mutation.ModelStatus(); ok {
		_spec.SetField(model.FieldModelStatus, field.TypeEnum, value)
		_node.ModelStatus = value
	}
	if value, ok := mc.mutation.IsOfficial(); ok {
		_spec.SetField(model.FieldIsOfficial, field.TypeBool, value)
		_node.IsOfficial = value
	}
	if value, ok := mc.mutation.TotalUsage(); ok {
		_spec.SetField(model.FieldTotalUsage, field.TypeInt, value)
		_node.TotalUsage = value
	}
	if nodes := mc.mutation.ModelPricesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ModelPricesTable,
			Columns: []string{model.ModelPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelprice.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.StarUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   model.StarUserTable,
			Columns: model.StarUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ModleStarCreate{config: mc.config, mutation: newModleStarMutation(mc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.StarModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   model.StarModelTable,
			Columns: []string{model.StarModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modlestar.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Model.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ModelUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (mc *ModelCreate) OnConflict(opts ...sql.ConflictOption) *ModelUpsertOne {
	mc.conflict = opts
	return &ModelUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Model.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *ModelCreate) OnConflictColumns(columns ...string) *ModelUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &ModelUpsertOne{
		create: mc,
	}
}

type (
	// ModelUpsertOne is the builder for "upsert"-ing
	//  one Model node.
	ModelUpsertOne struct {
		create *ModelCreate
	}

	// ModelUpsert is the "OnConflict" setter.
	ModelUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *ModelUpsert) SetCreatedBy(v int64) *ModelUpsert {
	u.Set(model.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *ModelUpsert) UpdateCreatedBy() *ModelUpsert {
	u.SetExcluded(model.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *ModelUpsert) AddCreatedBy(v int64) *ModelUpsert {
	u.Add(model.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *ModelUpsert) SetUpdatedBy(v int64) *ModelUpsert {
	u.Set(model.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *ModelUpsert) UpdateUpdatedBy() *ModelUpsert {
	u.SetExcluded(model.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *ModelUpsert) AddUpdatedBy(v int64) *ModelUpsert {
	u.Add(model.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ModelUpsert) SetUpdatedAt(v time.Time) *ModelUpsert {
	u.Set(model.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ModelUpsert) UpdateUpdatedAt() *ModelUpsert {
	u.SetExcluded(model.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ModelUpsert) SetDeletedAt(v time.Time) *ModelUpsert {
	u.Set(model.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ModelUpsert) UpdateDeletedAt() *ModelUpsert {
	u.SetExcluded(model.FieldDeletedAt)
	return u
}

// SetName sets the "name" field.
func (u *ModelUpsert) SetName(v string) *ModelUpsert {
	u.Set(model.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ModelUpsert) UpdateName() *ModelUpsert {
	u.SetExcluded(model.FieldName)
	return u
}

// SetAuthor sets the "author" field.
func (u *ModelUpsert) SetAuthor(v string) *ModelUpsert {
	u.Set(model.FieldAuthor, v)
	return u
}

// UpdateAuthor sets the "author" field to the value that was provided on create.
func (u *ModelUpsert) UpdateAuthor() *ModelUpsert {
	u.SetExcluded(model.FieldAuthor)
	return u
}

// SetModelType sets the "model_type" field.
func (u *ModelUpsert) SetModelType(v enums.Model) *ModelUpsert {
	u.Set(model.FieldModelType, v)
	return u
}

// UpdateModelType sets the "model_type" field to the value that was provided on create.
func (u *ModelUpsert) UpdateModelType() *ModelUpsert {
	u.SetExcluded(model.FieldModelType)
	return u
}

// SetModelStatus sets the "model_status" field.
func (u *ModelUpsert) SetModelStatus(v enums.ModelStatus) *ModelUpsert {
	u.Set(model.FieldModelStatus, v)
	return u
}

// UpdateModelStatus sets the "model_status" field to the value that was provided on create.
func (u *ModelUpsert) UpdateModelStatus() *ModelUpsert {
	u.SetExcluded(model.FieldModelStatus)
	return u
}

// SetIsOfficial sets the "is_official" field.
func (u *ModelUpsert) SetIsOfficial(v bool) *ModelUpsert {
	u.Set(model.FieldIsOfficial, v)
	return u
}

// UpdateIsOfficial sets the "is_official" field to the value that was provided on create.
func (u *ModelUpsert) UpdateIsOfficial() *ModelUpsert {
	u.SetExcluded(model.FieldIsOfficial)
	return u
}

// SetTotalUsage sets the "total_usage" field.
func (u *ModelUpsert) SetTotalUsage(v int) *ModelUpsert {
	u.Set(model.FieldTotalUsage, v)
	return u
}

// UpdateTotalUsage sets the "total_usage" field to the value that was provided on create.
func (u *ModelUpsert) UpdateTotalUsage() *ModelUpsert {
	u.SetExcluded(model.FieldTotalUsage)
	return u
}

// AddTotalUsage adds v to the "total_usage" field.
func (u *ModelUpsert) AddTotalUsage(v int) *ModelUpsert {
	u.Add(model.FieldTotalUsage, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Model.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(model.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ModelUpsertOne) UpdateNewValues() *ModelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(model.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(model.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Model.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ModelUpsertOne) Ignore() *ModelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ModelUpsertOne) DoNothing() *ModelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ModelCreate.OnConflict
// documentation for more info.
func (u *ModelUpsertOne) Update(set func(*ModelUpsert)) *ModelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ModelUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *ModelUpsertOne) SetCreatedBy(v int64) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *ModelUpsertOne) AddCreatedBy(v int64) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *ModelUpsertOne) UpdateCreatedBy() *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *ModelUpsertOne) SetUpdatedBy(v int64) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *ModelUpsertOne) AddUpdatedBy(v int64) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *ModelUpsertOne) UpdateUpdatedBy() *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ModelUpsertOne) SetUpdatedAt(v time.Time) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ModelUpsertOne) UpdateUpdatedAt() *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ModelUpsertOne) SetDeletedAt(v time.Time) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ModelUpsertOne) UpdateDeletedAt() *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *ModelUpsertOne) SetName(v string) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ModelUpsertOne) UpdateName() *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateName()
	})
}

// SetAuthor sets the "author" field.
func (u *ModelUpsertOne) SetAuthor(v string) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.SetAuthor(v)
	})
}

// UpdateAuthor sets the "author" field to the value that was provided on create.
func (u *ModelUpsertOne) UpdateAuthor() *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateAuthor()
	})
}

// SetModelType sets the "model_type" field.
func (u *ModelUpsertOne) SetModelType(v enums.Model) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.SetModelType(v)
	})
}

// UpdateModelType sets the "model_type" field to the value that was provided on create.
func (u *ModelUpsertOne) UpdateModelType() *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateModelType()
	})
}

// SetModelStatus sets the "model_status" field.
func (u *ModelUpsertOne) SetModelStatus(v enums.ModelStatus) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.SetModelStatus(v)
	})
}

// UpdateModelStatus sets the "model_status" field to the value that was provided on create.
func (u *ModelUpsertOne) UpdateModelStatus() *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateModelStatus()
	})
}

// SetIsOfficial sets the "is_official" field.
func (u *ModelUpsertOne) SetIsOfficial(v bool) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.SetIsOfficial(v)
	})
}

// UpdateIsOfficial sets the "is_official" field to the value that was provided on create.
func (u *ModelUpsertOne) UpdateIsOfficial() *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateIsOfficial()
	})
}

// SetTotalUsage sets the "total_usage" field.
func (u *ModelUpsertOne) SetTotalUsage(v int) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.SetTotalUsage(v)
	})
}

// AddTotalUsage adds v to the "total_usage" field.
func (u *ModelUpsertOne) AddTotalUsage(v int) *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.AddTotalUsage(v)
	})
}

// UpdateTotalUsage sets the "total_usage" field to the value that was provided on create.
func (u *ModelUpsertOne) UpdateTotalUsage() *ModelUpsertOne {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateTotalUsage()
	})
}

// Exec executes the query.
func (u *ModelUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for ModelCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ModelUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ModelUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ModelUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ModelCreateBulk is the builder for creating many Model entities in bulk.
type ModelCreateBulk struct {
	config
	err      error
	builders []*ModelCreate
	conflict []sql.ConflictOption
}

// Save creates the Model entities in the database.
func (mcb *ModelCreateBulk) Save(ctx context.Context) ([]*Model, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Model, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ModelMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *ModelCreateBulk) SaveX(ctx context.Context) []*Model {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *ModelCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *ModelCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Model.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ModelUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (mcb *ModelCreateBulk) OnConflict(opts ...sql.ConflictOption) *ModelUpsertBulk {
	mcb.conflict = opts
	return &ModelUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Model.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *ModelCreateBulk) OnConflictColumns(columns ...string) *ModelUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &ModelUpsertBulk{
		create: mcb,
	}
}

// ModelUpsertBulk is the builder for "upsert"-ing
// a bulk of Model nodes.
type ModelUpsertBulk struct {
	create *ModelCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Model.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(model.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ModelUpsertBulk) UpdateNewValues() *ModelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(model.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(model.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Model.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ModelUpsertBulk) Ignore() *ModelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ModelUpsertBulk) DoNothing() *ModelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ModelCreateBulk.OnConflict
// documentation for more info.
func (u *ModelUpsertBulk) Update(set func(*ModelUpsert)) *ModelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ModelUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *ModelUpsertBulk) SetCreatedBy(v int64) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *ModelUpsertBulk) AddCreatedBy(v int64) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *ModelUpsertBulk) UpdateCreatedBy() *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *ModelUpsertBulk) SetUpdatedBy(v int64) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *ModelUpsertBulk) AddUpdatedBy(v int64) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *ModelUpsertBulk) UpdateUpdatedBy() *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ModelUpsertBulk) SetUpdatedAt(v time.Time) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ModelUpsertBulk) UpdateUpdatedAt() *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ModelUpsertBulk) SetDeletedAt(v time.Time) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ModelUpsertBulk) UpdateDeletedAt() *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *ModelUpsertBulk) SetName(v string) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ModelUpsertBulk) UpdateName() *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateName()
	})
}

// SetAuthor sets the "author" field.
func (u *ModelUpsertBulk) SetAuthor(v string) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.SetAuthor(v)
	})
}

// UpdateAuthor sets the "author" field to the value that was provided on create.
func (u *ModelUpsertBulk) UpdateAuthor() *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateAuthor()
	})
}

// SetModelType sets the "model_type" field.
func (u *ModelUpsertBulk) SetModelType(v enums.Model) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.SetModelType(v)
	})
}

// UpdateModelType sets the "model_type" field to the value that was provided on create.
func (u *ModelUpsertBulk) UpdateModelType() *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateModelType()
	})
}

// SetModelStatus sets the "model_status" field.
func (u *ModelUpsertBulk) SetModelStatus(v enums.ModelStatus) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.SetModelStatus(v)
	})
}

// UpdateModelStatus sets the "model_status" field to the value that was provided on create.
func (u *ModelUpsertBulk) UpdateModelStatus() *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateModelStatus()
	})
}

// SetIsOfficial sets the "is_official" field.
func (u *ModelUpsertBulk) SetIsOfficial(v bool) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.SetIsOfficial(v)
	})
}

// UpdateIsOfficial sets the "is_official" field to the value that was provided on create.
func (u *ModelUpsertBulk) UpdateIsOfficial() *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateIsOfficial()
	})
}

// SetTotalUsage sets the "total_usage" field.
func (u *ModelUpsertBulk) SetTotalUsage(v int) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.SetTotalUsage(v)
	})
}

// AddTotalUsage adds v to the "total_usage" field.
func (u *ModelUpsertBulk) AddTotalUsage(v int) *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.AddTotalUsage(v)
	})
}

// UpdateTotalUsage sets the "total_usage" field to the value that was provided on create.
func (u *ModelUpsertBulk) UpdateTotalUsage() *ModelUpsertBulk {
	return u.Update(func(s *ModelUpsert) {
		s.UpdateTotalUsage()
	})
}

// Exec executes the query.
func (u *ModelUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the ModelCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for ModelCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ModelUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
