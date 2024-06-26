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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/devicereboottime"
)

// DeviceRebootTimeCreate is the builder for creating a DeviceRebootTime entity.
type DeviceRebootTimeCreate struct {
	config
	mutation *DeviceRebootTimeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (drtc *DeviceRebootTimeCreate) SetCreatedBy(i int64) *DeviceRebootTimeCreate {
	drtc.mutation.SetCreatedBy(i)
	return drtc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableCreatedBy(i *int64) *DeviceRebootTimeCreate {
	if i != nil {
		drtc.SetCreatedBy(*i)
	}
	return drtc
}

// SetUpdatedBy sets the "updated_by" field.
func (drtc *DeviceRebootTimeCreate) SetUpdatedBy(i int64) *DeviceRebootTimeCreate {
	drtc.mutation.SetUpdatedBy(i)
	return drtc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableUpdatedBy(i *int64) *DeviceRebootTimeCreate {
	if i != nil {
		drtc.SetUpdatedBy(*i)
	}
	return drtc
}

// SetCreatedAt sets the "created_at" field.
func (drtc *DeviceRebootTimeCreate) SetCreatedAt(t time.Time) *DeviceRebootTimeCreate {
	drtc.mutation.SetCreatedAt(t)
	return drtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableCreatedAt(t *time.Time) *DeviceRebootTimeCreate {
	if t != nil {
		drtc.SetCreatedAt(*t)
	}
	return drtc
}

// SetUpdatedAt sets the "updated_at" field.
func (drtc *DeviceRebootTimeCreate) SetUpdatedAt(t time.Time) *DeviceRebootTimeCreate {
	drtc.mutation.SetUpdatedAt(t)
	return drtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableUpdatedAt(t *time.Time) *DeviceRebootTimeCreate {
	if t != nil {
		drtc.SetUpdatedAt(*t)
	}
	return drtc
}

// SetDeletedAt sets the "deleted_at" field.
func (drtc *DeviceRebootTimeCreate) SetDeletedAt(t time.Time) *DeviceRebootTimeCreate {
	drtc.mutation.SetDeletedAt(t)
	return drtc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableDeletedAt(t *time.Time) *DeviceRebootTimeCreate {
	if t != nil {
		drtc.SetDeletedAt(*t)
	}
	return drtc
}

// SetDeviceID sets the "device_id" field.
func (drtc *DeviceRebootTimeCreate) SetDeviceID(i int64) *DeviceRebootTimeCreate {
	drtc.mutation.SetDeviceID(i)
	return drtc
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableDeviceID(i *int64) *DeviceRebootTimeCreate {
	if i != nil {
		drtc.SetDeviceID(*i)
	}
	return drtc
}

// SetStartTime sets the "start_time" field.
func (drtc *DeviceRebootTimeCreate) SetStartTime(t time.Time) *DeviceRebootTimeCreate {
	drtc.mutation.SetStartTime(t)
	return drtc
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableStartTime(t *time.Time) *DeviceRebootTimeCreate {
	if t != nil {
		drtc.SetStartTime(*t)
	}
	return drtc
}

// SetEndTime sets the "end_time" field.
func (drtc *DeviceRebootTimeCreate) SetEndTime(t time.Time) *DeviceRebootTimeCreate {
	drtc.mutation.SetEndTime(t)
	return drtc
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableEndTime(t *time.Time) *DeviceRebootTimeCreate {
	if t != nil {
		drtc.SetEndTime(*t)
	}
	return drtc
}

// SetNowTime sets the "now_time" field.
func (drtc *DeviceRebootTimeCreate) SetNowTime(t time.Time) *DeviceRebootTimeCreate {
	drtc.mutation.SetNowTime(t)
	return drtc
}

// SetNillableNowTime sets the "now_time" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableNowTime(t *time.Time) *DeviceRebootTimeCreate {
	if t != nil {
		drtc.SetNowTime(*t)
	}
	return drtc
}

// SetOnlineTime sets the "online_time" field.
func (drtc *DeviceRebootTimeCreate) SetOnlineTime(s string) *DeviceRebootTimeCreate {
	drtc.mutation.SetOnlineTime(s)
	return drtc
}

// SetNillableOnlineTime sets the "online_time" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableOnlineTime(s *string) *DeviceRebootTimeCreate {
	if s != nil {
		drtc.SetOnlineTime(*s)
	}
	return drtc
}

// SetOfflineTime sets the "offline_time" field.
func (drtc *DeviceRebootTimeCreate) SetOfflineTime(s string) *DeviceRebootTimeCreate {
	drtc.mutation.SetOfflineTime(s)
	return drtc
}

// SetNillableOfflineTime sets the "offline_time" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableOfflineTime(s *string) *DeviceRebootTimeCreate {
	if s != nil {
		drtc.SetOfflineTime(*s)
	}
	return drtc
}

// SetID sets the "id" field.
func (drtc *DeviceRebootTimeCreate) SetID(i int64) *DeviceRebootTimeCreate {
	drtc.mutation.SetID(i)
	return drtc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (drtc *DeviceRebootTimeCreate) SetNillableID(i *int64) *DeviceRebootTimeCreate {
	if i != nil {
		drtc.SetID(*i)
	}
	return drtc
}

// SetDevice sets the "device" edge to the Device entity.
func (drtc *DeviceRebootTimeCreate) SetDevice(d *Device) *DeviceRebootTimeCreate {
	return drtc.SetDeviceID(d.ID)
}

// Mutation returns the DeviceRebootTimeMutation object of the builder.
func (drtc *DeviceRebootTimeCreate) Mutation() *DeviceRebootTimeMutation {
	return drtc.mutation
}

// Save creates the DeviceRebootTime in the database.
func (drtc *DeviceRebootTimeCreate) Save(ctx context.Context) (*DeviceRebootTime, error) {
	drtc.defaults()
	return withHooks(ctx, drtc.sqlSave, drtc.mutation, drtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (drtc *DeviceRebootTimeCreate) SaveX(ctx context.Context) *DeviceRebootTime {
	v, err := drtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (drtc *DeviceRebootTimeCreate) Exec(ctx context.Context) error {
	_, err := drtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drtc *DeviceRebootTimeCreate) ExecX(ctx context.Context) {
	if err := drtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (drtc *DeviceRebootTimeCreate) defaults() {
	if _, ok := drtc.mutation.CreatedBy(); !ok {
		v := devicereboottime.DefaultCreatedBy
		drtc.mutation.SetCreatedBy(v)
	}
	if _, ok := drtc.mutation.UpdatedBy(); !ok {
		v := devicereboottime.DefaultUpdatedBy
		drtc.mutation.SetUpdatedBy(v)
	}
	if _, ok := drtc.mutation.CreatedAt(); !ok {
		v := devicereboottime.DefaultCreatedAt()
		drtc.mutation.SetCreatedAt(v)
	}
	if _, ok := drtc.mutation.UpdatedAt(); !ok {
		v := devicereboottime.DefaultUpdatedAt()
		drtc.mutation.SetUpdatedAt(v)
	}
	if _, ok := drtc.mutation.DeletedAt(); !ok {
		v := devicereboottime.DefaultDeletedAt
		drtc.mutation.SetDeletedAt(v)
	}
	if _, ok := drtc.mutation.DeviceID(); !ok {
		v := devicereboottime.DefaultDeviceID
		drtc.mutation.SetDeviceID(v)
	}
	if _, ok := drtc.mutation.StartTime(); !ok {
		v := devicereboottime.DefaultStartTime
		drtc.mutation.SetStartTime(v)
	}
	if _, ok := drtc.mutation.EndTime(); !ok {
		v := devicereboottime.DefaultEndTime
		drtc.mutation.SetEndTime(v)
	}
	if _, ok := drtc.mutation.NowTime(); !ok {
		v := devicereboottime.DefaultNowTime
		drtc.mutation.SetNowTime(v)
	}
	if _, ok := drtc.mutation.OnlineTime(); !ok {
		v := devicereboottime.DefaultOnlineTime
		drtc.mutation.SetOnlineTime(v)
	}
	if _, ok := drtc.mutation.OfflineTime(); !ok {
		v := devicereboottime.DefaultOfflineTime
		drtc.mutation.SetOfflineTime(v)
	}
	if _, ok := drtc.mutation.ID(); !ok {
		v := devicereboottime.DefaultID()
		drtc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (drtc *DeviceRebootTimeCreate) check() error {
	if _, ok := drtc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "DeviceRebootTime.created_by"`)}
	}
	if _, ok := drtc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "DeviceRebootTime.updated_by"`)}
	}
	if _, ok := drtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "DeviceRebootTime.created_at"`)}
	}
	if _, ok := drtc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "DeviceRebootTime.updated_at"`)}
	}
	if _, ok := drtc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "DeviceRebootTime.deleted_at"`)}
	}
	if _, ok := drtc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device_id", err: errors.New(`cep_ent: missing required field "DeviceRebootTime.device_id"`)}
	}
	if _, ok := drtc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`cep_ent: missing required field "DeviceRebootTime.start_time"`)}
	}
	if _, ok := drtc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`cep_ent: missing required field "DeviceRebootTime.end_time"`)}
	}
	if _, ok := drtc.mutation.NowTime(); !ok {
		return &ValidationError{Name: "now_time", err: errors.New(`cep_ent: missing required field "DeviceRebootTime.now_time"`)}
	}
	if _, ok := drtc.mutation.OnlineTime(); !ok {
		return &ValidationError{Name: "online_time", err: errors.New(`cep_ent: missing required field "DeviceRebootTime.online_time"`)}
	}
	if _, ok := drtc.mutation.OfflineTime(); !ok {
		return &ValidationError{Name: "offline_time", err: errors.New(`cep_ent: missing required field "DeviceRebootTime.offline_time"`)}
	}
	if _, ok := drtc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device", err: errors.New(`cep_ent: missing required edge "DeviceRebootTime.device"`)}
	}
	return nil
}

func (drtc *DeviceRebootTimeCreate) sqlSave(ctx context.Context) (*DeviceRebootTime, error) {
	if err := drtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := drtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, drtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	drtc.mutation.id = &_node.ID
	drtc.mutation.done = true
	return _node, nil
}

func (drtc *DeviceRebootTimeCreate) createSpec() (*DeviceRebootTime, *sqlgraph.CreateSpec) {
	var (
		_node = &DeviceRebootTime{config: drtc.config}
		_spec = sqlgraph.NewCreateSpec(devicereboottime.Table, sqlgraph.NewFieldSpec(devicereboottime.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = drtc.conflict
	if id, ok := drtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := drtc.mutation.CreatedBy(); ok {
		_spec.SetField(devicereboottime.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := drtc.mutation.UpdatedBy(); ok {
		_spec.SetField(devicereboottime.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := drtc.mutation.CreatedAt(); ok {
		_spec.SetField(devicereboottime.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := drtc.mutation.UpdatedAt(); ok {
		_spec.SetField(devicereboottime.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := drtc.mutation.DeletedAt(); ok {
		_spec.SetField(devicereboottime.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := drtc.mutation.StartTime(); ok {
		_spec.SetField(devicereboottime.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := drtc.mutation.EndTime(); ok {
		_spec.SetField(devicereboottime.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := drtc.mutation.NowTime(); ok {
		_spec.SetField(devicereboottime.FieldNowTime, field.TypeTime, value)
		_node.NowTime = value
	}
	if value, ok := drtc.mutation.OnlineTime(); ok {
		_spec.SetField(devicereboottime.FieldOnlineTime, field.TypeString, value)
		_node.OnlineTime = value
	}
	if value, ok := drtc.mutation.OfflineTime(); ok {
		_spec.SetField(devicereboottime.FieldOfflineTime, field.TypeString, value)
		_node.OfflineTime = value
	}
	if nodes := drtc.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicereboottime.DeviceTable,
			Columns: []string{devicereboottime.DeviceColumn},
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
//	client.DeviceRebootTime.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceRebootTimeUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (drtc *DeviceRebootTimeCreate) OnConflict(opts ...sql.ConflictOption) *DeviceRebootTimeUpsertOne {
	drtc.conflict = opts
	return &DeviceRebootTimeUpsertOne{
		create: drtc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceRebootTime.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (drtc *DeviceRebootTimeCreate) OnConflictColumns(columns ...string) *DeviceRebootTimeUpsertOne {
	drtc.conflict = append(drtc.conflict, sql.ConflictColumns(columns...))
	return &DeviceRebootTimeUpsertOne{
		create: drtc,
	}
}

type (
	// DeviceRebootTimeUpsertOne is the builder for "upsert"-ing
	//  one DeviceRebootTime node.
	DeviceRebootTimeUpsertOne struct {
		create *DeviceRebootTimeCreate
	}

	// DeviceRebootTimeUpsert is the "OnConflict" setter.
	DeviceRebootTimeUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *DeviceRebootTimeUpsert) SetCreatedBy(v int64) *DeviceRebootTimeUpsert {
	u.Set(devicereboottime.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsert) UpdateCreatedBy() *DeviceRebootTimeUpsert {
	u.SetExcluded(devicereboottime.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceRebootTimeUpsert) AddCreatedBy(v int64) *DeviceRebootTimeUpsert {
	u.Add(devicereboottime.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceRebootTimeUpsert) SetUpdatedBy(v int64) *DeviceRebootTimeUpsert {
	u.Set(devicereboottime.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsert) UpdateUpdatedBy() *DeviceRebootTimeUpsert {
	u.SetExcluded(devicereboottime.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceRebootTimeUpsert) AddUpdatedBy(v int64) *DeviceRebootTimeUpsert {
	u.Add(devicereboottime.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceRebootTimeUpsert) SetUpdatedAt(v time.Time) *DeviceRebootTimeUpsert {
	u.Set(devicereboottime.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsert) UpdateUpdatedAt() *DeviceRebootTimeUpsert {
	u.SetExcluded(devicereboottime.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceRebootTimeUpsert) SetDeletedAt(v time.Time) *DeviceRebootTimeUpsert {
	u.Set(devicereboottime.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsert) UpdateDeletedAt() *DeviceRebootTimeUpsert {
	u.SetExcluded(devicereboottime.FieldDeletedAt)
	return u
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceRebootTimeUpsert) SetDeviceID(v int64) *DeviceRebootTimeUpsert {
	u.Set(devicereboottime.FieldDeviceID, v)
	return u
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsert) UpdateDeviceID() *DeviceRebootTimeUpsert {
	u.SetExcluded(devicereboottime.FieldDeviceID)
	return u
}

// SetStartTime sets the "start_time" field.
func (u *DeviceRebootTimeUpsert) SetStartTime(v time.Time) *DeviceRebootTimeUpsert {
	u.Set(devicereboottime.FieldStartTime, v)
	return u
}

// UpdateStartTime sets the "start_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsert) UpdateStartTime() *DeviceRebootTimeUpsert {
	u.SetExcluded(devicereboottime.FieldStartTime)
	return u
}

// SetEndTime sets the "end_time" field.
func (u *DeviceRebootTimeUpsert) SetEndTime(v time.Time) *DeviceRebootTimeUpsert {
	u.Set(devicereboottime.FieldEndTime, v)
	return u
}

// UpdateEndTime sets the "end_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsert) UpdateEndTime() *DeviceRebootTimeUpsert {
	u.SetExcluded(devicereboottime.FieldEndTime)
	return u
}

// SetNowTime sets the "now_time" field.
func (u *DeviceRebootTimeUpsert) SetNowTime(v time.Time) *DeviceRebootTimeUpsert {
	u.Set(devicereboottime.FieldNowTime, v)
	return u
}

// UpdateNowTime sets the "now_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsert) UpdateNowTime() *DeviceRebootTimeUpsert {
	u.SetExcluded(devicereboottime.FieldNowTime)
	return u
}

// SetOnlineTime sets the "online_time" field.
func (u *DeviceRebootTimeUpsert) SetOnlineTime(v string) *DeviceRebootTimeUpsert {
	u.Set(devicereboottime.FieldOnlineTime, v)
	return u
}

// UpdateOnlineTime sets the "online_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsert) UpdateOnlineTime() *DeviceRebootTimeUpsert {
	u.SetExcluded(devicereboottime.FieldOnlineTime)
	return u
}

// SetOfflineTime sets the "offline_time" field.
func (u *DeviceRebootTimeUpsert) SetOfflineTime(v string) *DeviceRebootTimeUpsert {
	u.Set(devicereboottime.FieldOfflineTime, v)
	return u
}

// UpdateOfflineTime sets the "offline_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsert) UpdateOfflineTime() *DeviceRebootTimeUpsert {
	u.SetExcluded(devicereboottime.FieldOfflineTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DeviceRebootTime.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(devicereboottime.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceRebootTimeUpsertOne) UpdateNewValues() *DeviceRebootTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(devicereboottime.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(devicereboottime.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceRebootTime.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeviceRebootTimeUpsertOne) Ignore() *DeviceRebootTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceRebootTimeUpsertOne) DoNothing() *DeviceRebootTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceRebootTimeCreate.OnConflict
// documentation for more info.
func (u *DeviceRebootTimeUpsertOne) Update(set func(*DeviceRebootTimeUpsert)) *DeviceRebootTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceRebootTimeUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *DeviceRebootTimeUpsertOne) SetCreatedBy(v int64) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceRebootTimeUpsertOne) AddCreatedBy(v int64) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertOne) UpdateCreatedBy() *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceRebootTimeUpsertOne) SetUpdatedBy(v int64) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceRebootTimeUpsertOne) AddUpdatedBy(v int64) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertOne) UpdateUpdatedBy() *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceRebootTimeUpsertOne) SetUpdatedAt(v time.Time) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertOne) UpdateUpdatedAt() *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceRebootTimeUpsertOne) SetDeletedAt(v time.Time) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertOne) UpdateDeletedAt() *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceRebootTimeUpsertOne) SetDeviceID(v int64) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertOne) UpdateDeviceID() *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateDeviceID()
	})
}

// SetStartTime sets the "start_time" field.
func (u *DeviceRebootTimeUpsertOne) SetStartTime(v time.Time) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetStartTime(v)
	})
}

// UpdateStartTime sets the "start_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertOne) UpdateStartTime() *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateStartTime()
	})
}

// SetEndTime sets the "end_time" field.
func (u *DeviceRebootTimeUpsertOne) SetEndTime(v time.Time) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetEndTime(v)
	})
}

// UpdateEndTime sets the "end_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertOne) UpdateEndTime() *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateEndTime()
	})
}

// SetNowTime sets the "now_time" field.
func (u *DeviceRebootTimeUpsertOne) SetNowTime(v time.Time) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetNowTime(v)
	})
}

// UpdateNowTime sets the "now_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertOne) UpdateNowTime() *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateNowTime()
	})
}

// SetOnlineTime sets the "online_time" field.
func (u *DeviceRebootTimeUpsertOne) SetOnlineTime(v string) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetOnlineTime(v)
	})
}

// UpdateOnlineTime sets the "online_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertOne) UpdateOnlineTime() *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateOnlineTime()
	})
}

// SetOfflineTime sets the "offline_time" field.
func (u *DeviceRebootTimeUpsertOne) SetOfflineTime(v string) *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetOfflineTime(v)
	})
}

// UpdateOfflineTime sets the "offline_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertOne) UpdateOfflineTime() *DeviceRebootTimeUpsertOne {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateOfflineTime()
	})
}

// Exec executes the query.
func (u *DeviceRebootTimeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for DeviceRebootTimeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceRebootTimeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeviceRebootTimeUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeviceRebootTimeUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeviceRebootTimeCreateBulk is the builder for creating many DeviceRebootTime entities in bulk.
type DeviceRebootTimeCreateBulk struct {
	config
	err      error
	builders []*DeviceRebootTimeCreate
	conflict []sql.ConflictOption
}

// Save creates the DeviceRebootTime entities in the database.
func (drtcb *DeviceRebootTimeCreateBulk) Save(ctx context.Context) ([]*DeviceRebootTime, error) {
	if drtcb.err != nil {
		return nil, drtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(drtcb.builders))
	nodes := make([]*DeviceRebootTime, len(drtcb.builders))
	mutators := make([]Mutator, len(drtcb.builders))
	for i := range drtcb.builders {
		func(i int, root context.Context) {
			builder := drtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceRebootTimeMutation)
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
					_, err = mutators[i+1].Mutate(root, drtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = drtcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, drtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, drtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (drtcb *DeviceRebootTimeCreateBulk) SaveX(ctx context.Context) []*DeviceRebootTime {
	v, err := drtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (drtcb *DeviceRebootTimeCreateBulk) Exec(ctx context.Context) error {
	_, err := drtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drtcb *DeviceRebootTimeCreateBulk) ExecX(ctx context.Context) {
	if err := drtcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceRebootTime.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceRebootTimeUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (drtcb *DeviceRebootTimeCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeviceRebootTimeUpsertBulk {
	drtcb.conflict = opts
	return &DeviceRebootTimeUpsertBulk{
		create: drtcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceRebootTime.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (drtcb *DeviceRebootTimeCreateBulk) OnConflictColumns(columns ...string) *DeviceRebootTimeUpsertBulk {
	drtcb.conflict = append(drtcb.conflict, sql.ConflictColumns(columns...))
	return &DeviceRebootTimeUpsertBulk{
		create: drtcb,
	}
}

// DeviceRebootTimeUpsertBulk is the builder for "upsert"-ing
// a bulk of DeviceRebootTime nodes.
type DeviceRebootTimeUpsertBulk struct {
	create *DeviceRebootTimeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeviceRebootTime.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(devicereboottime.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceRebootTimeUpsertBulk) UpdateNewValues() *DeviceRebootTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(devicereboottime.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(devicereboottime.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceRebootTime.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeviceRebootTimeUpsertBulk) Ignore() *DeviceRebootTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceRebootTimeUpsertBulk) DoNothing() *DeviceRebootTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceRebootTimeCreateBulk.OnConflict
// documentation for more info.
func (u *DeviceRebootTimeUpsertBulk) Update(set func(*DeviceRebootTimeUpsert)) *DeviceRebootTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceRebootTimeUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *DeviceRebootTimeUpsertBulk) SetCreatedBy(v int64) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceRebootTimeUpsertBulk) AddCreatedBy(v int64) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertBulk) UpdateCreatedBy() *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceRebootTimeUpsertBulk) SetUpdatedBy(v int64) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceRebootTimeUpsertBulk) AddUpdatedBy(v int64) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertBulk) UpdateUpdatedBy() *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceRebootTimeUpsertBulk) SetUpdatedAt(v time.Time) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertBulk) UpdateUpdatedAt() *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceRebootTimeUpsertBulk) SetDeletedAt(v time.Time) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertBulk) UpdateDeletedAt() *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceRebootTimeUpsertBulk) SetDeviceID(v int64) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertBulk) UpdateDeviceID() *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateDeviceID()
	})
}

// SetStartTime sets the "start_time" field.
func (u *DeviceRebootTimeUpsertBulk) SetStartTime(v time.Time) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetStartTime(v)
	})
}

// UpdateStartTime sets the "start_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertBulk) UpdateStartTime() *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateStartTime()
	})
}

// SetEndTime sets the "end_time" field.
func (u *DeviceRebootTimeUpsertBulk) SetEndTime(v time.Time) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetEndTime(v)
	})
}

// UpdateEndTime sets the "end_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertBulk) UpdateEndTime() *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateEndTime()
	})
}

// SetNowTime sets the "now_time" field.
func (u *DeviceRebootTimeUpsertBulk) SetNowTime(v time.Time) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetNowTime(v)
	})
}

// UpdateNowTime sets the "now_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertBulk) UpdateNowTime() *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateNowTime()
	})
}

// SetOnlineTime sets the "online_time" field.
func (u *DeviceRebootTimeUpsertBulk) SetOnlineTime(v string) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetOnlineTime(v)
	})
}

// UpdateOnlineTime sets the "online_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertBulk) UpdateOnlineTime() *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateOnlineTime()
	})
}

// SetOfflineTime sets the "offline_time" field.
func (u *DeviceRebootTimeUpsertBulk) SetOfflineTime(v string) *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.SetOfflineTime(v)
	})
}

// UpdateOfflineTime sets the "offline_time" field to the value that was provided on create.
func (u *DeviceRebootTimeUpsertBulk) UpdateOfflineTime() *DeviceRebootTimeUpsertBulk {
	return u.Update(func(s *DeviceRebootTimeUpsert) {
		s.UpdateOfflineTime()
	})
}

// Exec executes the query.
func (u *DeviceRebootTimeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the DeviceRebootTimeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for DeviceRebootTimeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceRebootTimeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
