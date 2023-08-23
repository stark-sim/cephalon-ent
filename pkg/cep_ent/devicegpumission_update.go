// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/device"
	"cephalon-ent/pkg/cep_ent/devicegpumission"
	"cephalon-ent/pkg/cep_ent/gpu"
	"cephalon-ent/pkg/cep_ent/missionkind"
	"cephalon-ent/pkg/cep_ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeviceGpuMissionUpdate is the builder for updating DeviceGpuMission entities.
type DeviceGpuMissionUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceGpuMissionMutation
}

// Where appends a list predicates to the DeviceGpuMissionUpdate builder.
func (dgmu *DeviceGpuMissionUpdate) Where(ps ...predicate.DeviceGpuMission) *DeviceGpuMissionUpdate {
	dgmu.mutation.Where(ps...)
	return dgmu
}

// SetCreatedBy sets the "created_by" field.
func (dgmu *DeviceGpuMissionUpdate) SetCreatedBy(i int64) *DeviceGpuMissionUpdate {
	dgmu.mutation.ResetCreatedBy()
	dgmu.mutation.SetCreatedBy(i)
	return dgmu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dgmu *DeviceGpuMissionUpdate) SetNillableCreatedBy(i *int64) *DeviceGpuMissionUpdate {
	if i != nil {
		dgmu.SetCreatedBy(*i)
	}
	return dgmu
}

// AddCreatedBy adds i to the "created_by" field.
func (dgmu *DeviceGpuMissionUpdate) AddCreatedBy(i int64) *DeviceGpuMissionUpdate {
	dgmu.mutation.AddCreatedBy(i)
	return dgmu
}

// SetUpdatedBy sets the "updated_by" field.
func (dgmu *DeviceGpuMissionUpdate) SetUpdatedBy(i int64) *DeviceGpuMissionUpdate {
	dgmu.mutation.ResetUpdatedBy()
	dgmu.mutation.SetUpdatedBy(i)
	return dgmu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dgmu *DeviceGpuMissionUpdate) SetNillableUpdatedBy(i *int64) *DeviceGpuMissionUpdate {
	if i != nil {
		dgmu.SetUpdatedBy(*i)
	}
	return dgmu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (dgmu *DeviceGpuMissionUpdate) AddUpdatedBy(i int64) *DeviceGpuMissionUpdate {
	dgmu.mutation.AddUpdatedBy(i)
	return dgmu
}

// SetUpdatedAt sets the "updated_at" field.
func (dgmu *DeviceGpuMissionUpdate) SetUpdatedAt(t time.Time) *DeviceGpuMissionUpdate {
	dgmu.mutation.SetUpdatedAt(t)
	return dgmu
}

// SetDeletedAt sets the "deleted_at" field.
func (dgmu *DeviceGpuMissionUpdate) SetDeletedAt(t time.Time) *DeviceGpuMissionUpdate {
	dgmu.mutation.SetDeletedAt(t)
	return dgmu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dgmu *DeviceGpuMissionUpdate) SetNillableDeletedAt(t *time.Time) *DeviceGpuMissionUpdate {
	if t != nil {
		dgmu.SetDeletedAt(*t)
	}
	return dgmu
}

// SetDeviceID sets the "device_id" field.
func (dgmu *DeviceGpuMissionUpdate) SetDeviceID(i int64) *DeviceGpuMissionUpdate {
	dgmu.mutation.SetDeviceID(i)
	return dgmu
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (dgmu *DeviceGpuMissionUpdate) SetNillableDeviceID(i *int64) *DeviceGpuMissionUpdate {
	if i != nil {
		dgmu.SetDeviceID(*i)
	}
	return dgmu
}

// SetGpuID sets the "gpu_id" field.
func (dgmu *DeviceGpuMissionUpdate) SetGpuID(i int64) *DeviceGpuMissionUpdate {
	dgmu.mutation.SetGpuID(i)
	return dgmu
}

// SetNillableGpuID sets the "gpu_id" field if the given value is not nil.
func (dgmu *DeviceGpuMissionUpdate) SetNillableGpuID(i *int64) *DeviceGpuMissionUpdate {
	if i != nil {
		dgmu.SetGpuID(*i)
	}
	return dgmu
}

// SetMissionKindID sets the "mission_kind_id" field.
func (dgmu *DeviceGpuMissionUpdate) SetMissionKindID(i int64) *DeviceGpuMissionUpdate {
	dgmu.mutation.SetMissionKindID(i)
	return dgmu
}

// SetNillableMissionKindID sets the "mission_kind_id" field if the given value is not nil.
func (dgmu *DeviceGpuMissionUpdate) SetNillableMissionKindID(i *int64) *DeviceGpuMissionUpdate {
	if i != nil {
		dgmu.SetMissionKindID(*i)
	}
	return dgmu
}

// SetDevice sets the "device" edge to the Device entity.
func (dgmu *DeviceGpuMissionUpdate) SetDevice(d *Device) *DeviceGpuMissionUpdate {
	return dgmu.SetDeviceID(d.ID)
}

// SetMissionKind sets the "mission_kind" edge to the MissionKind entity.
func (dgmu *DeviceGpuMissionUpdate) SetMissionKind(m *MissionKind) *DeviceGpuMissionUpdate {
	return dgmu.SetMissionKindID(m.ID)
}

// SetGpu sets the "gpu" edge to the Gpu entity.
func (dgmu *DeviceGpuMissionUpdate) SetGpu(g *Gpu) *DeviceGpuMissionUpdate {
	return dgmu.SetGpuID(g.ID)
}

// Mutation returns the DeviceGpuMissionMutation object of the builder.
func (dgmu *DeviceGpuMissionUpdate) Mutation() *DeviceGpuMissionMutation {
	return dgmu.mutation
}

// ClearDevice clears the "device" edge to the Device entity.
func (dgmu *DeviceGpuMissionUpdate) ClearDevice() *DeviceGpuMissionUpdate {
	dgmu.mutation.ClearDevice()
	return dgmu
}

// ClearMissionKind clears the "mission_kind" edge to the MissionKind entity.
func (dgmu *DeviceGpuMissionUpdate) ClearMissionKind() *DeviceGpuMissionUpdate {
	dgmu.mutation.ClearMissionKind()
	return dgmu
}

// ClearGpu clears the "gpu" edge to the Gpu entity.
func (dgmu *DeviceGpuMissionUpdate) ClearGpu() *DeviceGpuMissionUpdate {
	dgmu.mutation.ClearGpu()
	return dgmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dgmu *DeviceGpuMissionUpdate) Save(ctx context.Context) (int, error) {
	dgmu.defaults()
	return withHooks(ctx, dgmu.sqlSave, dgmu.mutation, dgmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dgmu *DeviceGpuMissionUpdate) SaveX(ctx context.Context) int {
	affected, err := dgmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dgmu *DeviceGpuMissionUpdate) Exec(ctx context.Context) error {
	_, err := dgmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dgmu *DeviceGpuMissionUpdate) ExecX(ctx context.Context) {
	if err := dgmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dgmu *DeviceGpuMissionUpdate) defaults() {
	if _, ok := dgmu.mutation.UpdatedAt(); !ok {
		v := devicegpumission.UpdateDefaultUpdatedAt()
		dgmu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dgmu *DeviceGpuMissionUpdate) check() error {
	if _, ok := dgmu.mutation.DeviceID(); dgmu.mutation.DeviceCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "DeviceGpuMission.device"`)
	}
	if _, ok := dgmu.mutation.MissionKindID(); dgmu.mutation.MissionKindCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "DeviceGpuMission.mission_kind"`)
	}
	if _, ok := dgmu.mutation.GpuID(); dgmu.mutation.GpuCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "DeviceGpuMission.gpu"`)
	}
	return nil
}

func (dgmu *DeviceGpuMissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dgmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(devicegpumission.Table, devicegpumission.Columns, sqlgraph.NewFieldSpec(devicegpumission.FieldID, field.TypeInt64))
	if ps := dgmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dgmu.mutation.CreatedBy(); ok {
		_spec.SetField(devicegpumission.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := dgmu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(devicegpumission.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := dgmu.mutation.UpdatedBy(); ok {
		_spec.SetField(devicegpumission.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := dgmu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(devicegpumission.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := dgmu.mutation.UpdatedAt(); ok {
		_spec.SetField(devicegpumission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := dgmu.mutation.DeletedAt(); ok {
		_spec.SetField(devicegpumission.FieldDeletedAt, field.TypeTime, value)
	}
	if dgmu.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.DeviceTable,
			Columns: []string{devicegpumission.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgmu.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.DeviceTable,
			Columns: []string{devicegpumission.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dgmu.mutation.MissionKindCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.MissionKindTable,
			Columns: []string{devicegpumission.MissionKindColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionkind.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgmu.mutation.MissionKindIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.MissionKindTable,
			Columns: []string{devicegpumission.MissionKindColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionkind.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dgmu.mutation.GpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.GpuTable,
			Columns: []string{devicegpumission.GpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gpu.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgmu.mutation.GpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.GpuTable,
			Columns: []string{devicegpumission.GpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gpu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dgmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{devicegpumission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dgmu.mutation.done = true
	return n, nil
}

// DeviceGpuMissionUpdateOne is the builder for updating a single DeviceGpuMission entity.
type DeviceGpuMissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceGpuMissionMutation
}

// SetCreatedBy sets the "created_by" field.
func (dgmuo *DeviceGpuMissionUpdateOne) SetCreatedBy(i int64) *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.ResetCreatedBy()
	dgmuo.mutation.SetCreatedBy(i)
	return dgmuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dgmuo *DeviceGpuMissionUpdateOne) SetNillableCreatedBy(i *int64) *DeviceGpuMissionUpdateOne {
	if i != nil {
		dgmuo.SetCreatedBy(*i)
	}
	return dgmuo
}

// AddCreatedBy adds i to the "created_by" field.
func (dgmuo *DeviceGpuMissionUpdateOne) AddCreatedBy(i int64) *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.AddCreatedBy(i)
	return dgmuo
}

// SetUpdatedBy sets the "updated_by" field.
func (dgmuo *DeviceGpuMissionUpdateOne) SetUpdatedBy(i int64) *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.ResetUpdatedBy()
	dgmuo.mutation.SetUpdatedBy(i)
	return dgmuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dgmuo *DeviceGpuMissionUpdateOne) SetNillableUpdatedBy(i *int64) *DeviceGpuMissionUpdateOne {
	if i != nil {
		dgmuo.SetUpdatedBy(*i)
	}
	return dgmuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (dgmuo *DeviceGpuMissionUpdateOne) AddUpdatedBy(i int64) *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.AddUpdatedBy(i)
	return dgmuo
}

// SetUpdatedAt sets the "updated_at" field.
func (dgmuo *DeviceGpuMissionUpdateOne) SetUpdatedAt(t time.Time) *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.SetUpdatedAt(t)
	return dgmuo
}

// SetDeletedAt sets the "deleted_at" field.
func (dgmuo *DeviceGpuMissionUpdateOne) SetDeletedAt(t time.Time) *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.SetDeletedAt(t)
	return dgmuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dgmuo *DeviceGpuMissionUpdateOne) SetNillableDeletedAt(t *time.Time) *DeviceGpuMissionUpdateOne {
	if t != nil {
		dgmuo.SetDeletedAt(*t)
	}
	return dgmuo
}

// SetDeviceID sets the "device_id" field.
func (dgmuo *DeviceGpuMissionUpdateOne) SetDeviceID(i int64) *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.SetDeviceID(i)
	return dgmuo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (dgmuo *DeviceGpuMissionUpdateOne) SetNillableDeviceID(i *int64) *DeviceGpuMissionUpdateOne {
	if i != nil {
		dgmuo.SetDeviceID(*i)
	}
	return dgmuo
}

// SetGpuID sets the "gpu_id" field.
func (dgmuo *DeviceGpuMissionUpdateOne) SetGpuID(i int64) *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.SetGpuID(i)
	return dgmuo
}

// SetNillableGpuID sets the "gpu_id" field if the given value is not nil.
func (dgmuo *DeviceGpuMissionUpdateOne) SetNillableGpuID(i *int64) *DeviceGpuMissionUpdateOne {
	if i != nil {
		dgmuo.SetGpuID(*i)
	}
	return dgmuo
}

// SetMissionKindID sets the "mission_kind_id" field.
func (dgmuo *DeviceGpuMissionUpdateOne) SetMissionKindID(i int64) *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.SetMissionKindID(i)
	return dgmuo
}

// SetNillableMissionKindID sets the "mission_kind_id" field if the given value is not nil.
func (dgmuo *DeviceGpuMissionUpdateOne) SetNillableMissionKindID(i *int64) *DeviceGpuMissionUpdateOne {
	if i != nil {
		dgmuo.SetMissionKindID(*i)
	}
	return dgmuo
}

// SetDevice sets the "device" edge to the Device entity.
func (dgmuo *DeviceGpuMissionUpdateOne) SetDevice(d *Device) *DeviceGpuMissionUpdateOne {
	return dgmuo.SetDeviceID(d.ID)
}

// SetMissionKind sets the "mission_kind" edge to the MissionKind entity.
func (dgmuo *DeviceGpuMissionUpdateOne) SetMissionKind(m *MissionKind) *DeviceGpuMissionUpdateOne {
	return dgmuo.SetMissionKindID(m.ID)
}

// SetGpu sets the "gpu" edge to the Gpu entity.
func (dgmuo *DeviceGpuMissionUpdateOne) SetGpu(g *Gpu) *DeviceGpuMissionUpdateOne {
	return dgmuo.SetGpuID(g.ID)
}

// Mutation returns the DeviceGpuMissionMutation object of the builder.
func (dgmuo *DeviceGpuMissionUpdateOne) Mutation() *DeviceGpuMissionMutation {
	return dgmuo.mutation
}

// ClearDevice clears the "device" edge to the Device entity.
func (dgmuo *DeviceGpuMissionUpdateOne) ClearDevice() *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.ClearDevice()
	return dgmuo
}

// ClearMissionKind clears the "mission_kind" edge to the MissionKind entity.
func (dgmuo *DeviceGpuMissionUpdateOne) ClearMissionKind() *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.ClearMissionKind()
	return dgmuo
}

// ClearGpu clears the "gpu" edge to the Gpu entity.
func (dgmuo *DeviceGpuMissionUpdateOne) ClearGpu() *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.ClearGpu()
	return dgmuo
}

// Where appends a list predicates to the DeviceGpuMissionUpdate builder.
func (dgmuo *DeviceGpuMissionUpdateOne) Where(ps ...predicate.DeviceGpuMission) *DeviceGpuMissionUpdateOne {
	dgmuo.mutation.Where(ps...)
	return dgmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dgmuo *DeviceGpuMissionUpdateOne) Select(field string, fields ...string) *DeviceGpuMissionUpdateOne {
	dgmuo.fields = append([]string{field}, fields...)
	return dgmuo
}

// Save executes the query and returns the updated DeviceGpuMission entity.
func (dgmuo *DeviceGpuMissionUpdateOne) Save(ctx context.Context) (*DeviceGpuMission, error) {
	dgmuo.defaults()
	return withHooks(ctx, dgmuo.sqlSave, dgmuo.mutation, dgmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dgmuo *DeviceGpuMissionUpdateOne) SaveX(ctx context.Context) *DeviceGpuMission {
	node, err := dgmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dgmuo *DeviceGpuMissionUpdateOne) Exec(ctx context.Context) error {
	_, err := dgmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dgmuo *DeviceGpuMissionUpdateOne) ExecX(ctx context.Context) {
	if err := dgmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dgmuo *DeviceGpuMissionUpdateOne) defaults() {
	if _, ok := dgmuo.mutation.UpdatedAt(); !ok {
		v := devicegpumission.UpdateDefaultUpdatedAt()
		dgmuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dgmuo *DeviceGpuMissionUpdateOne) check() error {
	if _, ok := dgmuo.mutation.DeviceID(); dgmuo.mutation.DeviceCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "DeviceGpuMission.device"`)
	}
	if _, ok := dgmuo.mutation.MissionKindID(); dgmuo.mutation.MissionKindCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "DeviceGpuMission.mission_kind"`)
	}
	if _, ok := dgmuo.mutation.GpuID(); dgmuo.mutation.GpuCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "DeviceGpuMission.gpu"`)
	}
	return nil
}

func (dgmuo *DeviceGpuMissionUpdateOne) sqlSave(ctx context.Context) (_node *DeviceGpuMission, err error) {
	if err := dgmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(devicegpumission.Table, devicegpumission.Columns, sqlgraph.NewFieldSpec(devicegpumission.FieldID, field.TypeInt64))
	id, ok := dgmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "DeviceGpuMission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dgmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, devicegpumission.FieldID)
		for _, f := range fields {
			if !devicegpumission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != devicegpumission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dgmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dgmuo.mutation.CreatedBy(); ok {
		_spec.SetField(devicegpumission.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := dgmuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(devicegpumission.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := dgmuo.mutation.UpdatedBy(); ok {
		_spec.SetField(devicegpumission.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := dgmuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(devicegpumission.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := dgmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(devicegpumission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := dgmuo.mutation.DeletedAt(); ok {
		_spec.SetField(devicegpumission.FieldDeletedAt, field.TypeTime, value)
	}
	if dgmuo.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.DeviceTable,
			Columns: []string{devicegpumission.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgmuo.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.DeviceTable,
			Columns: []string{devicegpumission.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dgmuo.mutation.MissionKindCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.MissionKindTable,
			Columns: []string{devicegpumission.MissionKindColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionkind.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgmuo.mutation.MissionKindIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.MissionKindTable,
			Columns: []string{devicegpumission.MissionKindColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionkind.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dgmuo.mutation.GpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.GpuTable,
			Columns: []string{devicegpumission.GpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gpu.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgmuo.mutation.GpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.GpuTable,
			Columns: []string{devicegpumission.GpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gpu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DeviceGpuMission{config: dgmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dgmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{devicegpumission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dgmuo.mutation.done = true
	return _node, nil
}