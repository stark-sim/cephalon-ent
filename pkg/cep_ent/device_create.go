// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/device"
	"cephalon-ent/pkg/cep_ent/devicegpumission"
	"cephalon-ent/pkg/cep_ent/missionproduceorder"
	"cephalon-ent/pkg/cep_ent/user"
	"cephalon-ent/pkg/cep_ent/userdevice"
	"cephalon-ent/pkg/enums"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeviceCreate is the builder for creating a Device entity.
type DeviceCreate struct {
	config
	mutation *DeviceMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (dc *DeviceCreate) SetCreatedBy(i int64) *DeviceCreate {
	dc.mutation.SetCreatedBy(i)
	return dc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableCreatedBy(i *int64) *DeviceCreate {
	if i != nil {
		dc.SetCreatedBy(*i)
	}
	return dc
}

// SetUpdatedBy sets the "updated_by" field.
func (dc *DeviceCreate) SetUpdatedBy(i int64) *DeviceCreate {
	dc.mutation.SetUpdatedBy(i)
	return dc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableUpdatedBy(i *int64) *DeviceCreate {
	if i != nil {
		dc.SetUpdatedBy(*i)
	}
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DeviceCreate) SetCreatedAt(t time.Time) *DeviceCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableCreatedAt(t *time.Time) *DeviceCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DeviceCreate) SetUpdatedAt(t time.Time) *DeviceCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableUpdatedAt(t *time.Time) *DeviceCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetDeletedAt sets the "deleted_at" field.
func (dc *DeviceCreate) SetDeletedAt(t time.Time) *DeviceCreate {
	dc.mutation.SetDeletedAt(t)
	return dc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableDeletedAt(t *time.Time) *DeviceCreate {
	if t != nil {
		dc.SetDeletedAt(*t)
	}
	return dc
}

// SetUserID sets the "user_id" field.
func (dc *DeviceCreate) SetUserID(i int64) *DeviceCreate {
	dc.mutation.SetUserID(i)
	return dc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableUserID(i *int64) *DeviceCreate {
	if i != nil {
		dc.SetUserID(*i)
	}
	return dc
}

// SetSerialNumber sets the "serial_number" field.
func (dc *DeviceCreate) SetSerialNumber(s string) *DeviceCreate {
	dc.mutation.SetSerialNumber(s)
	return dc
}

// SetNillableSerialNumber sets the "serial_number" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableSerialNumber(s *string) *DeviceCreate {
	if s != nil {
		dc.SetSerialNumber(*s)
	}
	return dc
}

// SetState sets the "state" field.
func (dc *DeviceCreate) SetState(d device.State) *DeviceCreate {
	dc.mutation.SetState(d)
	return dc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableState(d *device.State) *DeviceCreate {
	if d != nil {
		dc.SetState(*d)
	}
	return dc
}

// SetSumCep sets the "sum_cep" field.
func (dc *DeviceCreate) SetSumCep(i int64) *DeviceCreate {
	dc.mutation.SetSumCep(i)
	return dc
}

// SetNillableSumCep sets the "sum_cep" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableSumCep(i *int64) *DeviceCreate {
	if i != nil {
		dc.SetSumCep(*i)
	}
	return dc
}

// SetLinking sets the "linking" field.
func (dc *DeviceCreate) SetLinking(b bool) *DeviceCreate {
	dc.mutation.SetLinking(b)
	return dc
}

// SetNillableLinking sets the "linking" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableLinking(b *bool) *DeviceCreate {
	if b != nil {
		dc.SetLinking(*b)
	}
	return dc
}

// SetBindingStatus sets the "binding_status" field.
func (dc *DeviceCreate) SetBindingStatus(ebs enums.DeviceBindingStatus) *DeviceCreate {
	dc.mutation.SetBindingStatus(ebs)
	return dc
}

// SetNillableBindingStatus sets the "binding_status" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableBindingStatus(ebs *enums.DeviceBindingStatus) *DeviceCreate {
	if ebs != nil {
		dc.SetBindingStatus(*ebs)
	}
	return dc
}

// SetStatus sets the "status" field.
func (dc *DeviceCreate) SetStatus(d device.Status) *DeviceCreate {
	dc.mutation.SetStatus(d)
	return dc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableStatus(d *device.Status) *DeviceCreate {
	if d != nil {
		dc.SetStatus(*d)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DeviceCreate) SetID(i int64) *DeviceCreate {
	dc.mutation.SetID(i)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableID(i *int64) *DeviceCreate {
	if i != nil {
		dc.SetID(*i)
	}
	return dc
}

// SetUser sets the "user" edge to the User entity.
func (dc *DeviceCreate) SetUser(u *User) *DeviceCreate {
	return dc.SetUserID(u.ID)
}

// AddMissionProduceOrderIDs adds the "mission_produce_orders" edge to the MissionProduceOrder entity by IDs.
func (dc *DeviceCreate) AddMissionProduceOrderIDs(ids ...int64) *DeviceCreate {
	dc.mutation.AddMissionProduceOrderIDs(ids...)
	return dc
}

// AddMissionProduceOrders adds the "mission_produce_orders" edges to the MissionProduceOrder entity.
func (dc *DeviceCreate) AddMissionProduceOrders(m ...*MissionProduceOrder) *DeviceCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return dc.AddMissionProduceOrderIDs(ids...)
}

// AddUserDeviceIDs adds the "user_devices" edge to the UserDevice entity by IDs.
func (dc *DeviceCreate) AddUserDeviceIDs(ids ...int64) *DeviceCreate {
	dc.mutation.AddUserDeviceIDs(ids...)
	return dc
}

// AddUserDevices adds the "user_devices" edges to the UserDevice entity.
func (dc *DeviceCreate) AddUserDevices(u ...*UserDevice) *DeviceCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserDeviceIDs(ids...)
}

// AddDeviceGpuMissionIDs adds the "device_gpu_missions" edge to the DeviceGpuMission entity by IDs.
func (dc *DeviceCreate) AddDeviceGpuMissionIDs(ids ...int64) *DeviceCreate {
	dc.mutation.AddDeviceGpuMissionIDs(ids...)
	return dc
}

// AddDeviceGpuMissions adds the "device_gpu_missions" edges to the DeviceGpuMission entity.
func (dc *DeviceCreate) AddDeviceGpuMissions(d ...*DeviceGpuMission) *DeviceCreate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDeviceGpuMissionIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (dc *DeviceCreate) Mutation() *DeviceMutation {
	return dc.mutation
}

// Save creates the Device in the database.
func (dc *DeviceCreate) Save(ctx context.Context) (*Device, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeviceCreate) SaveX(ctx context.Context) *Device {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeviceCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeviceCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeviceCreate) defaults() {
	if _, ok := dc.mutation.CreatedBy(); !ok {
		v := device.DefaultCreatedBy
		dc.mutation.SetCreatedBy(v)
	}
	if _, ok := dc.mutation.UpdatedBy(); !ok {
		v := device.DefaultUpdatedBy
		dc.mutation.SetUpdatedBy(v)
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := device.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := device.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.DeletedAt(); !ok {
		v := device.DefaultDeletedAt
		dc.mutation.SetDeletedAt(v)
	}
	if _, ok := dc.mutation.UserID(); !ok {
		v := device.DefaultUserID
		dc.mutation.SetUserID(v)
	}
	if _, ok := dc.mutation.SerialNumber(); !ok {
		v := device.DefaultSerialNumber
		dc.mutation.SetSerialNumber(v)
	}
	if _, ok := dc.mutation.State(); !ok {
		v := device.DefaultState
		dc.mutation.SetState(v)
	}
	if _, ok := dc.mutation.SumCep(); !ok {
		v := device.DefaultSumCep
		dc.mutation.SetSumCep(v)
	}
	if _, ok := dc.mutation.Linking(); !ok {
		v := device.DefaultLinking
		dc.mutation.SetLinking(v)
	}
	if _, ok := dc.mutation.BindingStatus(); !ok {
		v := device.DefaultBindingStatus
		dc.mutation.SetBindingStatus(v)
	}
	if _, ok := dc.mutation.Status(); !ok {
		v := device.DefaultStatus
		dc.mutation.SetStatus(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := device.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeviceCreate) check() error {
	if _, ok := dc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "Device.created_by"`)}
	}
	if _, ok := dc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "Device.updated_by"`)}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "Device.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "Device.updated_at"`)}
	}
	if _, ok := dc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "Device.deleted_at"`)}
	}
	if _, ok := dc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`cep_ent: missing required field "Device.user_id"`)}
	}
	if _, ok := dc.mutation.SerialNumber(); !ok {
		return &ValidationError{Name: "serial_number", err: errors.New(`cep_ent: missing required field "Device.serial_number"`)}
	}
	if _, ok := dc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`cep_ent: missing required field "Device.state"`)}
	}
	if v, ok := dc.mutation.State(); ok {
		if err := device.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`cep_ent: validator failed for field "Device.state": %w`, err)}
		}
	}
	if _, ok := dc.mutation.SumCep(); !ok {
		return &ValidationError{Name: "sum_cep", err: errors.New(`cep_ent: missing required field "Device.sum_cep"`)}
	}
	if _, ok := dc.mutation.Linking(); !ok {
		return &ValidationError{Name: "linking", err: errors.New(`cep_ent: missing required field "Device.linking"`)}
	}
	if _, ok := dc.mutation.BindingStatus(); !ok {
		return &ValidationError{Name: "binding_status", err: errors.New(`cep_ent: missing required field "Device.binding_status"`)}
	}
	if v, ok := dc.mutation.BindingStatus(); ok {
		if err := device.BindingStatusValidator(v); err != nil {
			return &ValidationError{Name: "binding_status", err: fmt.Errorf(`cep_ent: validator failed for field "Device.binding_status": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`cep_ent: missing required field "Device.status"`)}
	}
	if v, ok := dc.mutation.Status(); ok {
		if err := device.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "Device.status": %w`, err)}
		}
	}
	if _, ok := dc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`cep_ent: missing required edge "Device.user"`)}
	}
	return nil
}

func (dc *DeviceCreate) sqlSave(ctx context.Context) (*Device, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeviceCreate) createSpec() (*Device, *sqlgraph.CreateSpec) {
	var (
		_node = &Device{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(device.Table, sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedBy(); ok {
		_spec.SetField(device.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := dc.mutation.UpdatedBy(); ok {
		_spec.SetField(device.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(device.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(device.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.DeletedAt(); ok {
		_spec.SetField(device.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := dc.mutation.SerialNumber(); ok {
		_spec.SetField(device.FieldSerialNumber, field.TypeString, value)
		_node.SerialNumber = value
	}
	if value, ok := dc.mutation.State(); ok {
		_spec.SetField(device.FieldState, field.TypeEnum, value)
		_node.State = value
	}
	if value, ok := dc.mutation.SumCep(); ok {
		_spec.SetField(device.FieldSumCep, field.TypeInt64, value)
		_node.SumCep = value
	}
	if value, ok := dc.mutation.Linking(); ok {
		_spec.SetField(device.FieldLinking, field.TypeBool, value)
		_node.Linking = value
	}
	if value, ok := dc.mutation.BindingStatus(); ok {
		_spec.SetField(device.FieldBindingStatus, field.TypeEnum, value)
		_node.BindingStatus = value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.SetField(device.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := dc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.UserTable,
			Columns: []string{device.UserColumn},
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
	if nodes := dc.mutation.MissionProduceOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.MissionProduceOrdersTable,
			Columns: []string{device.MissionProduceOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionproduceorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.UserDevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.UserDevicesTable,
			Columns: []string{device.UserDevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userdevice.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.DeviceGpuMissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.DeviceGpuMissionsTable,
			Columns: []string{device.DeviceGpuMissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(devicegpumission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeviceCreateBulk is the builder for creating many Device entities in bulk.
type DeviceCreateBulk struct {
	config
	builders []*DeviceCreate
}

// Save creates the Device entities in the database.
func (dcb *DeviceCreateBulk) Save(ctx context.Context) ([]*Device, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Device, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeviceCreateBulk) SaveX(ctx context.Context) []*Device {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeviceCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}