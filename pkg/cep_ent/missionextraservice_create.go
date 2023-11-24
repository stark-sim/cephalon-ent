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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/extraservice"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionextraservice"
)

// MissionExtraServiceCreate is the builder for creating a MissionExtraService entity.
type MissionExtraServiceCreate struct {
	config
	mutation *MissionExtraServiceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (mesc *MissionExtraServiceCreate) SetCreatedBy(i int64) *MissionExtraServiceCreate {
	mesc.mutation.SetCreatedBy(i)
	return mesc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mesc *MissionExtraServiceCreate) SetNillableCreatedBy(i *int64) *MissionExtraServiceCreate {
	if i != nil {
		mesc.SetCreatedBy(*i)
	}
	return mesc
}

// SetUpdatedBy sets the "updated_by" field.
func (mesc *MissionExtraServiceCreate) SetUpdatedBy(i int64) *MissionExtraServiceCreate {
	mesc.mutation.SetUpdatedBy(i)
	return mesc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mesc *MissionExtraServiceCreate) SetNillableUpdatedBy(i *int64) *MissionExtraServiceCreate {
	if i != nil {
		mesc.SetUpdatedBy(*i)
	}
	return mesc
}

// SetCreatedAt sets the "created_at" field.
func (mesc *MissionExtraServiceCreate) SetCreatedAt(t time.Time) *MissionExtraServiceCreate {
	mesc.mutation.SetCreatedAt(t)
	return mesc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mesc *MissionExtraServiceCreate) SetNillableCreatedAt(t *time.Time) *MissionExtraServiceCreate {
	if t != nil {
		mesc.SetCreatedAt(*t)
	}
	return mesc
}

// SetUpdatedAt sets the "updated_at" field.
func (mesc *MissionExtraServiceCreate) SetUpdatedAt(t time.Time) *MissionExtraServiceCreate {
	mesc.mutation.SetUpdatedAt(t)
	return mesc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mesc *MissionExtraServiceCreate) SetNillableUpdatedAt(t *time.Time) *MissionExtraServiceCreate {
	if t != nil {
		mesc.SetUpdatedAt(*t)
	}
	return mesc
}

// SetDeletedAt sets the "deleted_at" field.
func (mesc *MissionExtraServiceCreate) SetDeletedAt(t time.Time) *MissionExtraServiceCreate {
	mesc.mutation.SetDeletedAt(t)
	return mesc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mesc *MissionExtraServiceCreate) SetNillableDeletedAt(t *time.Time) *MissionExtraServiceCreate {
	if t != nil {
		mesc.SetDeletedAt(*t)
	}
	return mesc
}

// SetMissionID sets the "mission_id" field.
func (mesc *MissionExtraServiceCreate) SetMissionID(i int64) *MissionExtraServiceCreate {
	mesc.mutation.SetMissionID(i)
	return mesc
}

// SetNillableMissionID sets the "mission_id" field if the given value is not nil.
func (mesc *MissionExtraServiceCreate) SetNillableMissionID(i *int64) *MissionExtraServiceCreate {
	if i != nil {
		mesc.SetMissionID(*i)
	}
	return mesc
}

// SetExtraServiceID sets the "extra_service_id" field.
func (mesc *MissionExtraServiceCreate) SetExtraServiceID(i int64) *MissionExtraServiceCreate {
	mesc.mutation.SetExtraServiceID(i)
	return mesc
}

// SetNillableExtraServiceID sets the "extra_service_id" field if the given value is not nil.
func (mesc *MissionExtraServiceCreate) SetNillableExtraServiceID(i *int64) *MissionExtraServiceCreate {
	if i != nil {
		mesc.SetExtraServiceID(*i)
	}
	return mesc
}

// SetID sets the "id" field.
func (mesc *MissionExtraServiceCreate) SetID(i int64) *MissionExtraServiceCreate {
	mesc.mutation.SetID(i)
	return mesc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mesc *MissionExtraServiceCreate) SetNillableID(i *int64) *MissionExtraServiceCreate {
	if i != nil {
		mesc.SetID(*i)
	}
	return mesc
}

// SetMission sets the "mission" edge to the Mission entity.
func (mesc *MissionExtraServiceCreate) SetMission(m *Mission) *MissionExtraServiceCreate {
	return mesc.SetMissionID(m.ID)
}

// SetExtraService sets the "extra_service" edge to the ExtraService entity.
func (mesc *MissionExtraServiceCreate) SetExtraService(e *ExtraService) *MissionExtraServiceCreate {
	return mesc.SetExtraServiceID(e.ID)
}

// Mutation returns the MissionExtraServiceMutation object of the builder.
func (mesc *MissionExtraServiceCreate) Mutation() *MissionExtraServiceMutation {
	return mesc.mutation
}

// Save creates the MissionExtraService in the database.
func (mesc *MissionExtraServiceCreate) Save(ctx context.Context) (*MissionExtraService, error) {
	mesc.defaults()
	return withHooks(ctx, mesc.sqlSave, mesc.mutation, mesc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mesc *MissionExtraServiceCreate) SaveX(ctx context.Context) *MissionExtraService {
	v, err := mesc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mesc *MissionExtraServiceCreate) Exec(ctx context.Context) error {
	_, err := mesc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mesc *MissionExtraServiceCreate) ExecX(ctx context.Context) {
	if err := mesc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mesc *MissionExtraServiceCreate) defaults() {
	if _, ok := mesc.mutation.CreatedBy(); !ok {
		v := missionextraservice.DefaultCreatedBy
		mesc.mutation.SetCreatedBy(v)
	}
	if _, ok := mesc.mutation.UpdatedBy(); !ok {
		v := missionextraservice.DefaultUpdatedBy
		mesc.mutation.SetUpdatedBy(v)
	}
	if _, ok := mesc.mutation.CreatedAt(); !ok {
		v := missionextraservice.DefaultCreatedAt()
		mesc.mutation.SetCreatedAt(v)
	}
	if _, ok := mesc.mutation.UpdatedAt(); !ok {
		v := missionextraservice.DefaultUpdatedAt()
		mesc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mesc.mutation.DeletedAt(); !ok {
		v := missionextraservice.DefaultDeletedAt
		mesc.mutation.SetDeletedAt(v)
	}
	if _, ok := mesc.mutation.MissionID(); !ok {
		v := missionextraservice.DefaultMissionID
		mesc.mutation.SetMissionID(v)
	}
	if _, ok := mesc.mutation.ExtraServiceID(); !ok {
		v := missionextraservice.DefaultExtraServiceID
		mesc.mutation.SetExtraServiceID(v)
	}
	if _, ok := mesc.mutation.ID(); !ok {
		v := missionextraservice.DefaultID()
		mesc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mesc *MissionExtraServiceCreate) check() error {
	if _, ok := mesc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "MissionExtraService.created_by"`)}
	}
	if _, ok := mesc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "MissionExtraService.updated_by"`)}
	}
	if _, ok := mesc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "MissionExtraService.created_at"`)}
	}
	if _, ok := mesc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "MissionExtraService.updated_at"`)}
	}
	if _, ok := mesc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "MissionExtraService.deleted_at"`)}
	}
	if _, ok := mesc.mutation.MissionID(); !ok {
		return &ValidationError{Name: "mission_id", err: errors.New(`cep_ent: missing required field "MissionExtraService.mission_id"`)}
	}
	if _, ok := mesc.mutation.ExtraServiceID(); !ok {
		return &ValidationError{Name: "extra_service_id", err: errors.New(`cep_ent: missing required field "MissionExtraService.extra_service_id"`)}
	}
	if _, ok := mesc.mutation.MissionID(); !ok {
		return &ValidationError{Name: "mission", err: errors.New(`cep_ent: missing required edge "MissionExtraService.mission"`)}
	}
	if _, ok := mesc.mutation.ExtraServiceID(); !ok {
		return &ValidationError{Name: "extra_service", err: errors.New(`cep_ent: missing required edge "MissionExtraService.extra_service"`)}
	}
	return nil
}

func (mesc *MissionExtraServiceCreate) sqlSave(ctx context.Context) (*MissionExtraService, error) {
	if err := mesc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mesc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mesc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mesc.mutation.id = &_node.ID
	mesc.mutation.done = true
	return _node, nil
}

func (mesc *MissionExtraServiceCreate) createSpec() (*MissionExtraService, *sqlgraph.CreateSpec) {
	var (
		_node = &MissionExtraService{config: mesc.config}
		_spec = sqlgraph.NewCreateSpec(missionextraservice.Table, sqlgraph.NewFieldSpec(missionextraservice.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = mesc.conflict
	if id, ok := mesc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mesc.mutation.CreatedBy(); ok {
		_spec.SetField(missionextraservice.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := mesc.mutation.UpdatedBy(); ok {
		_spec.SetField(missionextraservice.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := mesc.mutation.CreatedAt(); ok {
		_spec.SetField(missionextraservice.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mesc.mutation.UpdatedAt(); ok {
		_spec.SetField(missionextraservice.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mesc.mutation.DeletedAt(); ok {
		_spec.SetField(missionextraservice.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := mesc.mutation.MissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionextraservice.MissionTable,
			Columns: []string{missionextraservice.MissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MissionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mesc.mutation.ExtraServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionextraservice.ExtraServiceTable,
			Columns: []string{missionextraservice.ExtraServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(extraservice.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ExtraServiceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MissionExtraService.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MissionExtraServiceUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (mesc *MissionExtraServiceCreate) OnConflict(opts ...sql.ConflictOption) *MissionExtraServiceUpsertOne {
	mesc.conflict = opts
	return &MissionExtraServiceUpsertOne{
		create: mesc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MissionExtraService.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mesc *MissionExtraServiceCreate) OnConflictColumns(columns ...string) *MissionExtraServiceUpsertOne {
	mesc.conflict = append(mesc.conflict, sql.ConflictColumns(columns...))
	return &MissionExtraServiceUpsertOne{
		create: mesc,
	}
}

type (
	// MissionExtraServiceUpsertOne is the builder for "upsert"-ing
	//  one MissionExtraService node.
	MissionExtraServiceUpsertOne struct {
		create *MissionExtraServiceCreate
	}

	// MissionExtraServiceUpsert is the "OnConflict" setter.
	MissionExtraServiceUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *MissionExtraServiceUpsert) SetCreatedBy(v int64) *MissionExtraServiceUpsert {
	u.Set(missionextraservice.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *MissionExtraServiceUpsert) UpdateCreatedBy() *MissionExtraServiceUpsert {
	u.SetExcluded(missionextraservice.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *MissionExtraServiceUpsert) AddCreatedBy(v int64) *MissionExtraServiceUpsert {
	u.Add(missionextraservice.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MissionExtraServiceUpsert) SetUpdatedBy(v int64) *MissionExtraServiceUpsert {
	u.Set(missionextraservice.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MissionExtraServiceUpsert) UpdateUpdatedBy() *MissionExtraServiceUpsert {
	u.SetExcluded(missionextraservice.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MissionExtraServiceUpsert) AddUpdatedBy(v int64) *MissionExtraServiceUpsert {
	u.Add(missionextraservice.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MissionExtraServiceUpsert) SetUpdatedAt(v time.Time) *MissionExtraServiceUpsert {
	u.Set(missionextraservice.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MissionExtraServiceUpsert) UpdateUpdatedAt() *MissionExtraServiceUpsert {
	u.SetExcluded(missionextraservice.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MissionExtraServiceUpsert) SetDeletedAt(v time.Time) *MissionExtraServiceUpsert {
	u.Set(missionextraservice.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MissionExtraServiceUpsert) UpdateDeletedAt() *MissionExtraServiceUpsert {
	u.SetExcluded(missionextraservice.FieldDeletedAt)
	return u
}

// SetMissionID sets the "mission_id" field.
func (u *MissionExtraServiceUpsert) SetMissionID(v int64) *MissionExtraServiceUpsert {
	u.Set(missionextraservice.FieldMissionID, v)
	return u
}

// UpdateMissionID sets the "mission_id" field to the value that was provided on create.
func (u *MissionExtraServiceUpsert) UpdateMissionID() *MissionExtraServiceUpsert {
	u.SetExcluded(missionextraservice.FieldMissionID)
	return u
}

// SetExtraServiceID sets the "extra_service_id" field.
func (u *MissionExtraServiceUpsert) SetExtraServiceID(v int64) *MissionExtraServiceUpsert {
	u.Set(missionextraservice.FieldExtraServiceID, v)
	return u
}

// UpdateExtraServiceID sets the "extra_service_id" field to the value that was provided on create.
func (u *MissionExtraServiceUpsert) UpdateExtraServiceID() *MissionExtraServiceUpsert {
	u.SetExcluded(missionextraservice.FieldExtraServiceID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MissionExtraService.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(missionextraservice.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MissionExtraServiceUpsertOne) UpdateNewValues() *MissionExtraServiceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(missionextraservice.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(missionextraservice.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MissionExtraService.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MissionExtraServiceUpsertOne) Ignore() *MissionExtraServiceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MissionExtraServiceUpsertOne) DoNothing() *MissionExtraServiceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MissionExtraServiceCreate.OnConflict
// documentation for more info.
func (u *MissionExtraServiceUpsertOne) Update(set func(*MissionExtraServiceUpsert)) *MissionExtraServiceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MissionExtraServiceUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *MissionExtraServiceUpsertOne) SetCreatedBy(v int64) *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *MissionExtraServiceUpsertOne) AddCreatedBy(v int64) *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertOne) UpdateCreatedBy() *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MissionExtraServiceUpsertOne) SetUpdatedBy(v int64) *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MissionExtraServiceUpsertOne) AddUpdatedBy(v int64) *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertOne) UpdateUpdatedBy() *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MissionExtraServiceUpsertOne) SetUpdatedAt(v time.Time) *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertOne) UpdateUpdatedAt() *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MissionExtraServiceUpsertOne) SetDeletedAt(v time.Time) *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertOne) UpdateDeletedAt() *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetMissionID sets the "mission_id" field.
func (u *MissionExtraServiceUpsertOne) SetMissionID(v int64) *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetMissionID(v)
	})
}

// UpdateMissionID sets the "mission_id" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertOne) UpdateMissionID() *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateMissionID()
	})
}

// SetExtraServiceID sets the "extra_service_id" field.
func (u *MissionExtraServiceUpsertOne) SetExtraServiceID(v int64) *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetExtraServiceID(v)
	})
}

// UpdateExtraServiceID sets the "extra_service_id" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertOne) UpdateExtraServiceID() *MissionExtraServiceUpsertOne {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateExtraServiceID()
	})
}

// Exec executes the query.
func (u *MissionExtraServiceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for MissionExtraServiceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MissionExtraServiceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MissionExtraServiceUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MissionExtraServiceUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MissionExtraServiceCreateBulk is the builder for creating many MissionExtraService entities in bulk.
type MissionExtraServiceCreateBulk struct {
	config
	err      error
	builders []*MissionExtraServiceCreate
	conflict []sql.ConflictOption
}

// Save creates the MissionExtraService entities in the database.
func (mescb *MissionExtraServiceCreateBulk) Save(ctx context.Context) ([]*MissionExtraService, error) {
	if mescb.err != nil {
		return nil, mescb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mescb.builders))
	nodes := make([]*MissionExtraService, len(mescb.builders))
	mutators := make([]Mutator, len(mescb.builders))
	for i := range mescb.builders {
		func(i int, root context.Context) {
			builder := mescb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MissionExtraServiceMutation)
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
					_, err = mutators[i+1].Mutate(root, mescb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mescb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mescb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mescb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mescb *MissionExtraServiceCreateBulk) SaveX(ctx context.Context) []*MissionExtraService {
	v, err := mescb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mescb *MissionExtraServiceCreateBulk) Exec(ctx context.Context) error {
	_, err := mescb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mescb *MissionExtraServiceCreateBulk) ExecX(ctx context.Context) {
	if err := mescb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MissionExtraService.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MissionExtraServiceUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (mescb *MissionExtraServiceCreateBulk) OnConflict(opts ...sql.ConflictOption) *MissionExtraServiceUpsertBulk {
	mescb.conflict = opts
	return &MissionExtraServiceUpsertBulk{
		create: mescb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MissionExtraService.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mescb *MissionExtraServiceCreateBulk) OnConflictColumns(columns ...string) *MissionExtraServiceUpsertBulk {
	mescb.conflict = append(mescb.conflict, sql.ConflictColumns(columns...))
	return &MissionExtraServiceUpsertBulk{
		create: mescb,
	}
}

// MissionExtraServiceUpsertBulk is the builder for "upsert"-ing
// a bulk of MissionExtraService nodes.
type MissionExtraServiceUpsertBulk struct {
	create *MissionExtraServiceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MissionExtraService.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(missionextraservice.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MissionExtraServiceUpsertBulk) UpdateNewValues() *MissionExtraServiceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(missionextraservice.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(missionextraservice.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MissionExtraService.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MissionExtraServiceUpsertBulk) Ignore() *MissionExtraServiceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MissionExtraServiceUpsertBulk) DoNothing() *MissionExtraServiceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MissionExtraServiceCreateBulk.OnConflict
// documentation for more info.
func (u *MissionExtraServiceUpsertBulk) Update(set func(*MissionExtraServiceUpsert)) *MissionExtraServiceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MissionExtraServiceUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *MissionExtraServiceUpsertBulk) SetCreatedBy(v int64) *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *MissionExtraServiceUpsertBulk) AddCreatedBy(v int64) *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertBulk) UpdateCreatedBy() *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MissionExtraServiceUpsertBulk) SetUpdatedBy(v int64) *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MissionExtraServiceUpsertBulk) AddUpdatedBy(v int64) *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertBulk) UpdateUpdatedBy() *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MissionExtraServiceUpsertBulk) SetUpdatedAt(v time.Time) *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertBulk) UpdateUpdatedAt() *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MissionExtraServiceUpsertBulk) SetDeletedAt(v time.Time) *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertBulk) UpdateDeletedAt() *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetMissionID sets the "mission_id" field.
func (u *MissionExtraServiceUpsertBulk) SetMissionID(v int64) *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetMissionID(v)
	})
}

// UpdateMissionID sets the "mission_id" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertBulk) UpdateMissionID() *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateMissionID()
	})
}

// SetExtraServiceID sets the "extra_service_id" field.
func (u *MissionExtraServiceUpsertBulk) SetExtraServiceID(v int64) *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.SetExtraServiceID(v)
	})
}

// UpdateExtraServiceID sets the "extra_service_id" field to the value that was provided on create.
func (u *MissionExtraServiceUpsertBulk) UpdateExtraServiceID() *MissionExtraServiceUpsertBulk {
	return u.Update(func(s *MissionExtraServiceUpsert) {
		s.UpdateExtraServiceID()
	})
}

// Exec executes the query.
func (u *MissionExtraServiceUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the MissionExtraServiceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for MissionExtraServiceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MissionExtraServiceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
