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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/device"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/deviceconfig"
)

// DeviceConfigCreate is the builder for creating a DeviceConfig entity.
type DeviceConfigCreate struct {
	config
	mutation *DeviceConfigMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (dcc *DeviceConfigCreate) SetCreatedBy(i int64) *DeviceConfigCreate {
	dcc.mutation.SetCreatedBy(i)
	return dcc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dcc *DeviceConfigCreate) SetNillableCreatedBy(i *int64) *DeviceConfigCreate {
	if i != nil {
		dcc.SetCreatedBy(*i)
	}
	return dcc
}

// SetUpdatedBy sets the "updated_by" field.
func (dcc *DeviceConfigCreate) SetUpdatedBy(i int64) *DeviceConfigCreate {
	dcc.mutation.SetUpdatedBy(i)
	return dcc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dcc *DeviceConfigCreate) SetNillableUpdatedBy(i *int64) *DeviceConfigCreate {
	if i != nil {
		dcc.SetUpdatedBy(*i)
	}
	return dcc
}

// SetCreatedAt sets the "created_at" field.
func (dcc *DeviceConfigCreate) SetCreatedAt(t time.Time) *DeviceConfigCreate {
	dcc.mutation.SetCreatedAt(t)
	return dcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dcc *DeviceConfigCreate) SetNillableCreatedAt(t *time.Time) *DeviceConfigCreate {
	if t != nil {
		dcc.SetCreatedAt(*t)
	}
	return dcc
}

// SetUpdatedAt sets the "updated_at" field.
func (dcc *DeviceConfigCreate) SetUpdatedAt(t time.Time) *DeviceConfigCreate {
	dcc.mutation.SetUpdatedAt(t)
	return dcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dcc *DeviceConfigCreate) SetNillableUpdatedAt(t *time.Time) *DeviceConfigCreate {
	if t != nil {
		dcc.SetUpdatedAt(*t)
	}
	return dcc
}

// SetDeletedAt sets the "deleted_at" field.
func (dcc *DeviceConfigCreate) SetDeletedAt(t time.Time) *DeviceConfigCreate {
	dcc.mutation.SetDeletedAt(t)
	return dcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dcc *DeviceConfigCreate) SetNillableDeletedAt(t *time.Time) *DeviceConfigCreate {
	if t != nil {
		dcc.SetDeletedAt(*t)
	}
	return dcc
}

// SetDeviceID sets the "device_id" field.
func (dcc *DeviceConfigCreate) SetDeviceID(i int64) *DeviceConfigCreate {
	dcc.mutation.SetDeviceID(i)
	return dcc
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (dcc *DeviceConfigCreate) SetNillableDeviceID(i *int64) *DeviceConfigCreate {
	if i != nil {
		dcc.SetDeviceID(*i)
	}
	return dcc
}

// SetGapBase sets the "gap_base" field.
func (dcc *DeviceConfigCreate) SetGapBase(i int64) *DeviceConfigCreate {
	dcc.mutation.SetGapBase(i)
	return dcc
}

// SetNillableGapBase sets the "gap_base" field if the given value is not nil.
func (dcc *DeviceConfigCreate) SetNillableGapBase(i *int64) *DeviceConfigCreate {
	if i != nil {
		dcc.SetGapBase(*i)
	}
	return dcc
}

// SetGapRandomMax sets the "gap_random_max" field.
func (dcc *DeviceConfigCreate) SetGapRandomMax(i int64) *DeviceConfigCreate {
	dcc.mutation.SetGapRandomMax(i)
	return dcc
}

// SetNillableGapRandomMax sets the "gap_random_max" field if the given value is not nil.
func (dcc *DeviceConfigCreate) SetNillableGapRandomMax(i *int64) *DeviceConfigCreate {
	if i != nil {
		dcc.SetGapRandomMax(*i)
	}
	return dcc
}

// SetGapRandomMin sets the "gap_random_min" field.
func (dcc *DeviceConfigCreate) SetGapRandomMin(i int64) *DeviceConfigCreate {
	dcc.mutation.SetGapRandomMin(i)
	return dcc
}

// SetNillableGapRandomMin sets the "gap_random_min" field if the given value is not nil.
func (dcc *DeviceConfigCreate) SetNillableGapRandomMin(i *int64) *DeviceConfigCreate {
	if i != nil {
		dcc.SetGapRandomMin(*i)
	}
	return dcc
}

// SetID sets the "id" field.
func (dcc *DeviceConfigCreate) SetID(i int64) *DeviceConfigCreate {
	dcc.mutation.SetID(i)
	return dcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dcc *DeviceConfigCreate) SetNillableID(i *int64) *DeviceConfigCreate {
	if i != nil {
		dcc.SetID(*i)
	}
	return dcc
}

// SetDevice sets the "device" edge to the Device entity.
func (dcc *DeviceConfigCreate) SetDevice(d *Device) *DeviceConfigCreate {
	return dcc.SetDeviceID(d.ID)
}

// Mutation returns the DeviceConfigMutation object of the builder.
func (dcc *DeviceConfigCreate) Mutation() *DeviceConfigMutation {
	return dcc.mutation
}

// Save creates the DeviceConfig in the database.
func (dcc *DeviceConfigCreate) Save(ctx context.Context) (*DeviceConfig, error) {
	dcc.defaults()
	return withHooks(ctx, dcc.sqlSave, dcc.mutation, dcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dcc *DeviceConfigCreate) SaveX(ctx context.Context) *DeviceConfig {
	v, err := dcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcc *DeviceConfigCreate) Exec(ctx context.Context) error {
	_, err := dcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcc *DeviceConfigCreate) ExecX(ctx context.Context) {
	if err := dcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dcc *DeviceConfigCreate) defaults() {
	if _, ok := dcc.mutation.CreatedBy(); !ok {
		v := deviceconfig.DefaultCreatedBy
		dcc.mutation.SetCreatedBy(v)
	}
	if _, ok := dcc.mutation.UpdatedBy(); !ok {
		v := deviceconfig.DefaultUpdatedBy
		dcc.mutation.SetUpdatedBy(v)
	}
	if _, ok := dcc.mutation.CreatedAt(); !ok {
		v := deviceconfig.DefaultCreatedAt()
		dcc.mutation.SetCreatedAt(v)
	}
	if _, ok := dcc.mutation.UpdatedAt(); !ok {
		v := deviceconfig.DefaultUpdatedAt()
		dcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dcc.mutation.DeletedAt(); !ok {
		v := deviceconfig.DefaultDeletedAt
		dcc.mutation.SetDeletedAt(v)
	}
	if _, ok := dcc.mutation.DeviceID(); !ok {
		v := deviceconfig.DefaultDeviceID
		dcc.mutation.SetDeviceID(v)
	}
	if _, ok := dcc.mutation.GapBase(); !ok {
		v := deviceconfig.DefaultGapBase
		dcc.mutation.SetGapBase(v)
	}
	if _, ok := dcc.mutation.GapRandomMax(); !ok {
		v := deviceconfig.DefaultGapRandomMax
		dcc.mutation.SetGapRandomMax(v)
	}
	if _, ok := dcc.mutation.GapRandomMin(); !ok {
		v := deviceconfig.DefaultGapRandomMin
		dcc.mutation.SetGapRandomMin(v)
	}
	if _, ok := dcc.mutation.ID(); !ok {
		v := deviceconfig.DefaultID()
		dcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dcc *DeviceConfigCreate) check() error {
	if _, ok := dcc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "DeviceConfig.created_by"`)}
	}
	if _, ok := dcc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "DeviceConfig.updated_by"`)}
	}
	if _, ok := dcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "DeviceConfig.created_at"`)}
	}
	if _, ok := dcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "DeviceConfig.updated_at"`)}
	}
	if _, ok := dcc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "DeviceConfig.deleted_at"`)}
	}
	if _, ok := dcc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device_id", err: errors.New(`cep_ent: missing required field "DeviceConfig.device_id"`)}
	}
	if _, ok := dcc.mutation.GapBase(); !ok {
		return &ValidationError{Name: "gap_base", err: errors.New(`cep_ent: missing required field "DeviceConfig.gap_base"`)}
	}
	if _, ok := dcc.mutation.GapRandomMax(); !ok {
		return &ValidationError{Name: "gap_random_max", err: errors.New(`cep_ent: missing required field "DeviceConfig.gap_random_max"`)}
	}
	if _, ok := dcc.mutation.GapRandomMin(); !ok {
		return &ValidationError{Name: "gap_random_min", err: errors.New(`cep_ent: missing required field "DeviceConfig.gap_random_min"`)}
	}
	if _, ok := dcc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device", err: errors.New(`cep_ent: missing required edge "DeviceConfig.device"`)}
	}
	return nil
}

func (dcc *DeviceConfigCreate) sqlSave(ctx context.Context) (*DeviceConfig, error) {
	if err := dcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	dcc.mutation.id = &_node.ID
	dcc.mutation.done = true
	return _node, nil
}

func (dcc *DeviceConfigCreate) createSpec() (*DeviceConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &DeviceConfig{config: dcc.config}
		_spec = sqlgraph.NewCreateSpec(deviceconfig.Table, sqlgraph.NewFieldSpec(deviceconfig.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = dcc.conflict
	if id, ok := dcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dcc.mutation.CreatedBy(); ok {
		_spec.SetField(deviceconfig.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := dcc.mutation.UpdatedBy(); ok {
		_spec.SetField(deviceconfig.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := dcc.mutation.CreatedAt(); ok {
		_spec.SetField(deviceconfig.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dcc.mutation.UpdatedAt(); ok {
		_spec.SetField(deviceconfig.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dcc.mutation.DeletedAt(); ok {
		_spec.SetField(deviceconfig.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := dcc.mutation.GapBase(); ok {
		_spec.SetField(deviceconfig.FieldGapBase, field.TypeInt64, value)
		_node.GapBase = value
	}
	if value, ok := dcc.mutation.GapRandomMax(); ok {
		_spec.SetField(deviceconfig.FieldGapRandomMax, field.TypeInt64, value)
		_node.GapRandomMax = value
	}
	if value, ok := dcc.mutation.GapRandomMin(); ok {
		_spec.SetField(deviceconfig.FieldGapRandomMin, field.TypeInt64, value)
		_node.GapRandomMin = value
	}
	if nodes := dcc.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   deviceconfig.DeviceTable,
			Columns: []string{deviceconfig.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DeviceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceConfig.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceConfigUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (dcc *DeviceConfigCreate) OnConflict(opts ...sql.ConflictOption) *DeviceConfigUpsertOne {
	dcc.conflict = opts
	return &DeviceConfigUpsertOne{
		create: dcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dcc *DeviceConfigCreate) OnConflictColumns(columns ...string) *DeviceConfigUpsertOne {
	dcc.conflict = append(dcc.conflict, sql.ConflictColumns(columns...))
	return &DeviceConfigUpsertOne{
		create: dcc,
	}
}

type (
	// DeviceConfigUpsertOne is the builder for "upsert"-ing
	//  one DeviceConfig node.
	DeviceConfigUpsertOne struct {
		create *DeviceConfigCreate
	}

	// DeviceConfigUpsert is the "OnConflict" setter.
	DeviceConfigUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *DeviceConfigUpsert) SetCreatedBy(v int64) *DeviceConfigUpsert {
	u.Set(deviceconfig.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceConfigUpsert) UpdateCreatedBy() *DeviceConfigUpsert {
	u.SetExcluded(deviceconfig.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceConfigUpsert) AddCreatedBy(v int64) *DeviceConfigUpsert {
	u.Add(deviceconfig.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceConfigUpsert) SetUpdatedBy(v int64) *DeviceConfigUpsert {
	u.Set(deviceconfig.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceConfigUpsert) UpdateUpdatedBy() *DeviceConfigUpsert {
	u.SetExcluded(deviceconfig.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceConfigUpsert) AddUpdatedBy(v int64) *DeviceConfigUpsert {
	u.Add(deviceconfig.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceConfigUpsert) SetUpdatedAt(v time.Time) *DeviceConfigUpsert {
	u.Set(deviceconfig.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceConfigUpsert) UpdateUpdatedAt() *DeviceConfigUpsert {
	u.SetExcluded(deviceconfig.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceConfigUpsert) SetDeletedAt(v time.Time) *DeviceConfigUpsert {
	u.Set(deviceconfig.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceConfigUpsert) UpdateDeletedAt() *DeviceConfigUpsert {
	u.SetExcluded(deviceconfig.FieldDeletedAt)
	return u
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceConfigUpsert) SetDeviceID(v int64) *DeviceConfigUpsert {
	u.Set(deviceconfig.FieldDeviceID, v)
	return u
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceConfigUpsert) UpdateDeviceID() *DeviceConfigUpsert {
	u.SetExcluded(deviceconfig.FieldDeviceID)
	return u
}

// SetGapBase sets the "gap_base" field.
func (u *DeviceConfigUpsert) SetGapBase(v int64) *DeviceConfigUpsert {
	u.Set(deviceconfig.FieldGapBase, v)
	return u
}

// UpdateGapBase sets the "gap_base" field to the value that was provided on create.
func (u *DeviceConfigUpsert) UpdateGapBase() *DeviceConfigUpsert {
	u.SetExcluded(deviceconfig.FieldGapBase)
	return u
}

// AddGapBase adds v to the "gap_base" field.
func (u *DeviceConfigUpsert) AddGapBase(v int64) *DeviceConfigUpsert {
	u.Add(deviceconfig.FieldGapBase, v)
	return u
}

// SetGapRandomMax sets the "gap_random_max" field.
func (u *DeviceConfigUpsert) SetGapRandomMax(v int64) *DeviceConfigUpsert {
	u.Set(deviceconfig.FieldGapRandomMax, v)
	return u
}

// UpdateGapRandomMax sets the "gap_random_max" field to the value that was provided on create.
func (u *DeviceConfigUpsert) UpdateGapRandomMax() *DeviceConfigUpsert {
	u.SetExcluded(deviceconfig.FieldGapRandomMax)
	return u
}

// AddGapRandomMax adds v to the "gap_random_max" field.
func (u *DeviceConfigUpsert) AddGapRandomMax(v int64) *DeviceConfigUpsert {
	u.Add(deviceconfig.FieldGapRandomMax, v)
	return u
}

// SetGapRandomMin sets the "gap_random_min" field.
func (u *DeviceConfigUpsert) SetGapRandomMin(v int64) *DeviceConfigUpsert {
	u.Set(deviceconfig.FieldGapRandomMin, v)
	return u
}

// UpdateGapRandomMin sets the "gap_random_min" field to the value that was provided on create.
func (u *DeviceConfigUpsert) UpdateGapRandomMin() *DeviceConfigUpsert {
	u.SetExcluded(deviceconfig.FieldGapRandomMin)
	return u
}

// AddGapRandomMin adds v to the "gap_random_min" field.
func (u *DeviceConfigUpsert) AddGapRandomMin(v int64) *DeviceConfigUpsert {
	u.Add(deviceconfig.FieldGapRandomMin, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DeviceConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deviceconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceConfigUpsertOne) UpdateNewValues() *DeviceConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(deviceconfig.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(deviceconfig.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceConfig.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeviceConfigUpsertOne) Ignore() *DeviceConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceConfigUpsertOne) DoNothing() *DeviceConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceConfigCreate.OnConflict
// documentation for more info.
func (u *DeviceConfigUpsertOne) Update(set func(*DeviceConfigUpsert)) *DeviceConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *DeviceConfigUpsertOne) SetCreatedBy(v int64) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceConfigUpsertOne) AddCreatedBy(v int64) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceConfigUpsertOne) UpdateCreatedBy() *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceConfigUpsertOne) SetUpdatedBy(v int64) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceConfigUpsertOne) AddUpdatedBy(v int64) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceConfigUpsertOne) UpdateUpdatedBy() *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceConfigUpsertOne) SetUpdatedAt(v time.Time) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceConfigUpsertOne) UpdateUpdatedAt() *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceConfigUpsertOne) SetDeletedAt(v time.Time) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceConfigUpsertOne) UpdateDeletedAt() *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceConfigUpsertOne) SetDeviceID(v int64) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceConfigUpsertOne) UpdateDeviceID() *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateDeviceID()
	})
}

// SetGapBase sets the "gap_base" field.
func (u *DeviceConfigUpsertOne) SetGapBase(v int64) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetGapBase(v)
	})
}

// AddGapBase adds v to the "gap_base" field.
func (u *DeviceConfigUpsertOne) AddGapBase(v int64) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.AddGapBase(v)
	})
}

// UpdateGapBase sets the "gap_base" field to the value that was provided on create.
func (u *DeviceConfigUpsertOne) UpdateGapBase() *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateGapBase()
	})
}

// SetGapRandomMax sets the "gap_random_max" field.
func (u *DeviceConfigUpsertOne) SetGapRandomMax(v int64) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetGapRandomMax(v)
	})
}

// AddGapRandomMax adds v to the "gap_random_max" field.
func (u *DeviceConfigUpsertOne) AddGapRandomMax(v int64) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.AddGapRandomMax(v)
	})
}

// UpdateGapRandomMax sets the "gap_random_max" field to the value that was provided on create.
func (u *DeviceConfigUpsertOne) UpdateGapRandomMax() *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateGapRandomMax()
	})
}

// SetGapRandomMin sets the "gap_random_min" field.
func (u *DeviceConfigUpsertOne) SetGapRandomMin(v int64) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetGapRandomMin(v)
	})
}

// AddGapRandomMin adds v to the "gap_random_min" field.
func (u *DeviceConfigUpsertOne) AddGapRandomMin(v int64) *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.AddGapRandomMin(v)
	})
}

// UpdateGapRandomMin sets the "gap_random_min" field to the value that was provided on create.
func (u *DeviceConfigUpsertOne) UpdateGapRandomMin() *DeviceConfigUpsertOne {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateGapRandomMin()
	})
}

// Exec executes the query.
func (u *DeviceConfigUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for DeviceConfigCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceConfigUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeviceConfigUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeviceConfigUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeviceConfigCreateBulk is the builder for creating many DeviceConfig entities in bulk.
type DeviceConfigCreateBulk struct {
	config
	err      error
	builders []*DeviceConfigCreate
	conflict []sql.ConflictOption
}

// Save creates the DeviceConfig entities in the database.
func (dccb *DeviceConfigCreateBulk) Save(ctx context.Context) ([]*DeviceConfig, error) {
	if dccb.err != nil {
		return nil, dccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dccb.builders))
	nodes := make([]*DeviceConfig, len(dccb.builders))
	mutators := make([]Mutator, len(dccb.builders))
	for i := range dccb.builders {
		func(i int, root context.Context) {
			builder := dccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceConfigMutation)
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
					_, err = mutators[i+1].Mutate(root, dccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dccb *DeviceConfigCreateBulk) SaveX(ctx context.Context) []*DeviceConfig {
	v, err := dccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dccb *DeviceConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := dccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dccb *DeviceConfigCreateBulk) ExecX(ctx context.Context) {
	if err := dccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceConfig.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceConfigUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (dccb *DeviceConfigCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeviceConfigUpsertBulk {
	dccb.conflict = opts
	return &DeviceConfigUpsertBulk{
		create: dccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dccb *DeviceConfigCreateBulk) OnConflictColumns(columns ...string) *DeviceConfigUpsertBulk {
	dccb.conflict = append(dccb.conflict, sql.ConflictColumns(columns...))
	return &DeviceConfigUpsertBulk{
		create: dccb,
	}
}

// DeviceConfigUpsertBulk is the builder for "upsert"-ing
// a bulk of DeviceConfig nodes.
type DeviceConfigUpsertBulk struct {
	create *DeviceConfigCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeviceConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deviceconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceConfigUpsertBulk) UpdateNewValues() *DeviceConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(deviceconfig.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(deviceconfig.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceConfig.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeviceConfigUpsertBulk) Ignore() *DeviceConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceConfigUpsertBulk) DoNothing() *DeviceConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceConfigCreateBulk.OnConflict
// documentation for more info.
func (u *DeviceConfigUpsertBulk) Update(set func(*DeviceConfigUpsert)) *DeviceConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *DeviceConfigUpsertBulk) SetCreatedBy(v int64) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceConfigUpsertBulk) AddCreatedBy(v int64) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceConfigUpsertBulk) UpdateCreatedBy() *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceConfigUpsertBulk) SetUpdatedBy(v int64) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceConfigUpsertBulk) AddUpdatedBy(v int64) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceConfigUpsertBulk) UpdateUpdatedBy() *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceConfigUpsertBulk) SetUpdatedAt(v time.Time) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceConfigUpsertBulk) UpdateUpdatedAt() *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceConfigUpsertBulk) SetDeletedAt(v time.Time) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceConfigUpsertBulk) UpdateDeletedAt() *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceConfigUpsertBulk) SetDeviceID(v int64) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceConfigUpsertBulk) UpdateDeviceID() *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateDeviceID()
	})
}

// SetGapBase sets the "gap_base" field.
func (u *DeviceConfigUpsertBulk) SetGapBase(v int64) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetGapBase(v)
	})
}

// AddGapBase adds v to the "gap_base" field.
func (u *DeviceConfigUpsertBulk) AddGapBase(v int64) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.AddGapBase(v)
	})
}

// UpdateGapBase sets the "gap_base" field to the value that was provided on create.
func (u *DeviceConfigUpsertBulk) UpdateGapBase() *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateGapBase()
	})
}

// SetGapRandomMax sets the "gap_random_max" field.
func (u *DeviceConfigUpsertBulk) SetGapRandomMax(v int64) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetGapRandomMax(v)
	})
}

// AddGapRandomMax adds v to the "gap_random_max" field.
func (u *DeviceConfigUpsertBulk) AddGapRandomMax(v int64) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.AddGapRandomMax(v)
	})
}

// UpdateGapRandomMax sets the "gap_random_max" field to the value that was provided on create.
func (u *DeviceConfigUpsertBulk) UpdateGapRandomMax() *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateGapRandomMax()
	})
}

// SetGapRandomMin sets the "gap_random_min" field.
func (u *DeviceConfigUpsertBulk) SetGapRandomMin(v int64) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.SetGapRandomMin(v)
	})
}

// AddGapRandomMin adds v to the "gap_random_min" field.
func (u *DeviceConfigUpsertBulk) AddGapRandomMin(v int64) *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.AddGapRandomMin(v)
	})
}

// UpdateGapRandomMin sets the "gap_random_min" field to the value that was provided on create.
func (u *DeviceConfigUpsertBulk) UpdateGapRandomMin() *DeviceConfigUpsertBulk {
	return u.Update(func(s *DeviceConfigUpsert) {
		s.UpdateGapRandomMin()
	})
}

// Exec executes the query.
func (u *DeviceConfigUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the DeviceConfigCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for DeviceConfigCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceConfigUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}