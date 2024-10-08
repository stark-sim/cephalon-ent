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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/giftmissionconfig"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// GiftMissionConfigUpdate is the builder for updating GiftMissionConfig entities.
type GiftMissionConfigUpdate struct {
	config
	hooks     []Hook
	mutation  *GiftMissionConfigMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GiftMissionConfigUpdate builder.
func (gmcu *GiftMissionConfigUpdate) Where(ps ...predicate.GiftMissionConfig) *GiftMissionConfigUpdate {
	gmcu.mutation.Where(ps...)
	return gmcu
}

// SetCreatedBy sets the "created_by" field.
func (gmcu *GiftMissionConfigUpdate) SetCreatedBy(i int64) *GiftMissionConfigUpdate {
	gmcu.mutation.ResetCreatedBy()
	gmcu.mutation.SetCreatedBy(i)
	return gmcu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gmcu *GiftMissionConfigUpdate) SetNillableCreatedBy(i *int64) *GiftMissionConfigUpdate {
	if i != nil {
		gmcu.SetCreatedBy(*i)
	}
	return gmcu
}

// AddCreatedBy adds i to the "created_by" field.
func (gmcu *GiftMissionConfigUpdate) AddCreatedBy(i int64) *GiftMissionConfigUpdate {
	gmcu.mutation.AddCreatedBy(i)
	return gmcu
}

// SetUpdatedBy sets the "updated_by" field.
func (gmcu *GiftMissionConfigUpdate) SetUpdatedBy(i int64) *GiftMissionConfigUpdate {
	gmcu.mutation.ResetUpdatedBy()
	gmcu.mutation.SetUpdatedBy(i)
	return gmcu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gmcu *GiftMissionConfigUpdate) SetNillableUpdatedBy(i *int64) *GiftMissionConfigUpdate {
	if i != nil {
		gmcu.SetUpdatedBy(*i)
	}
	return gmcu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (gmcu *GiftMissionConfigUpdate) AddUpdatedBy(i int64) *GiftMissionConfigUpdate {
	gmcu.mutation.AddUpdatedBy(i)
	return gmcu
}

// SetUpdatedAt sets the "updated_at" field.
func (gmcu *GiftMissionConfigUpdate) SetUpdatedAt(t time.Time) *GiftMissionConfigUpdate {
	gmcu.mutation.SetUpdatedAt(t)
	return gmcu
}

// SetDeletedAt sets the "deleted_at" field.
func (gmcu *GiftMissionConfigUpdate) SetDeletedAt(t time.Time) *GiftMissionConfigUpdate {
	gmcu.mutation.SetDeletedAt(t)
	return gmcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gmcu *GiftMissionConfigUpdate) SetNillableDeletedAt(t *time.Time) *GiftMissionConfigUpdate {
	if t != nil {
		gmcu.SetDeletedAt(*t)
	}
	return gmcu
}

// SetStabilityLevel sets the "stability_level" field.
func (gmcu *GiftMissionConfigUpdate) SetStabilityLevel(est enums.DeviceStabilityType) *GiftMissionConfigUpdate {
	gmcu.mutation.SetStabilityLevel(est)
	return gmcu
}

// SetNillableStabilityLevel sets the "stability_level" field if the given value is not nil.
func (gmcu *GiftMissionConfigUpdate) SetNillableStabilityLevel(est *enums.DeviceStabilityType) *GiftMissionConfigUpdate {
	if est != nil {
		gmcu.SetStabilityLevel(*est)
	}
	return gmcu
}

// SetGpuVersion sets the "gpu_version" field.
func (gmcu *GiftMissionConfigUpdate) SetGpuVersion(ev enums.GpuVersion) *GiftMissionConfigUpdate {
	gmcu.mutation.SetGpuVersion(ev)
	return gmcu
}

// SetNillableGpuVersion sets the "gpu_version" field if the given value is not nil.
func (gmcu *GiftMissionConfigUpdate) SetNillableGpuVersion(ev *enums.GpuVersion) *GiftMissionConfigUpdate {
	if ev != nil {
		gmcu.SetGpuVersion(*ev)
	}
	return gmcu
}

// SetGapBase sets the "gap_base" field.
func (gmcu *GiftMissionConfigUpdate) SetGapBase(i int64) *GiftMissionConfigUpdate {
	gmcu.mutation.ResetGapBase()
	gmcu.mutation.SetGapBase(i)
	return gmcu
}

// SetNillableGapBase sets the "gap_base" field if the given value is not nil.
func (gmcu *GiftMissionConfigUpdate) SetNillableGapBase(i *int64) *GiftMissionConfigUpdate {
	if i != nil {
		gmcu.SetGapBase(*i)
	}
	return gmcu
}

// AddGapBase adds i to the "gap_base" field.
func (gmcu *GiftMissionConfigUpdate) AddGapBase(i int64) *GiftMissionConfigUpdate {
	gmcu.mutation.AddGapBase(i)
	return gmcu
}

// SetGapRandomMax sets the "gap_random_max" field.
func (gmcu *GiftMissionConfigUpdate) SetGapRandomMax(i int64) *GiftMissionConfigUpdate {
	gmcu.mutation.ResetGapRandomMax()
	gmcu.mutation.SetGapRandomMax(i)
	return gmcu
}

// SetNillableGapRandomMax sets the "gap_random_max" field if the given value is not nil.
func (gmcu *GiftMissionConfigUpdate) SetNillableGapRandomMax(i *int64) *GiftMissionConfigUpdate {
	if i != nil {
		gmcu.SetGapRandomMax(*i)
	}
	return gmcu
}

// AddGapRandomMax adds i to the "gap_random_max" field.
func (gmcu *GiftMissionConfigUpdate) AddGapRandomMax(i int64) *GiftMissionConfigUpdate {
	gmcu.mutation.AddGapRandomMax(i)
	return gmcu
}

// SetGapRandomMin sets the "gap_random_min" field.
func (gmcu *GiftMissionConfigUpdate) SetGapRandomMin(i int64) *GiftMissionConfigUpdate {
	gmcu.mutation.ResetGapRandomMin()
	gmcu.mutation.SetGapRandomMin(i)
	return gmcu
}

// SetNillableGapRandomMin sets the "gap_random_min" field if the given value is not nil.
func (gmcu *GiftMissionConfigUpdate) SetNillableGapRandomMin(i *int64) *GiftMissionConfigUpdate {
	if i != nil {
		gmcu.SetGapRandomMin(*i)
	}
	return gmcu
}

// AddGapRandomMin adds i to the "gap_random_min" field.
func (gmcu *GiftMissionConfigUpdate) AddGapRandomMin(i int64) *GiftMissionConfigUpdate {
	gmcu.mutation.AddGapRandomMin(i)
	return gmcu
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (gmcu *GiftMissionConfigUpdate) AddDeviceIDs(ids ...int64) *GiftMissionConfigUpdate {
	gmcu.mutation.AddDeviceIDs(ids...)
	return gmcu
}

// AddDevices adds the "devices" edges to the Device entity.
func (gmcu *GiftMissionConfigUpdate) AddDevices(d ...*Device) *GiftMissionConfigUpdate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gmcu.AddDeviceIDs(ids...)
}

// Mutation returns the GiftMissionConfigMutation object of the builder.
func (gmcu *GiftMissionConfigUpdate) Mutation() *GiftMissionConfigMutation {
	return gmcu.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (gmcu *GiftMissionConfigUpdate) ClearDevices() *GiftMissionConfigUpdate {
	gmcu.mutation.ClearDevices()
	return gmcu
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (gmcu *GiftMissionConfigUpdate) RemoveDeviceIDs(ids ...int64) *GiftMissionConfigUpdate {
	gmcu.mutation.RemoveDeviceIDs(ids...)
	return gmcu
}

// RemoveDevices removes "devices" edges to Device entities.
func (gmcu *GiftMissionConfigUpdate) RemoveDevices(d ...*Device) *GiftMissionConfigUpdate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gmcu.RemoveDeviceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gmcu *GiftMissionConfigUpdate) Save(ctx context.Context) (int, error) {
	gmcu.defaults()
	return withHooks(ctx, gmcu.sqlSave, gmcu.mutation, gmcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmcu *GiftMissionConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := gmcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gmcu *GiftMissionConfigUpdate) Exec(ctx context.Context) error {
	_, err := gmcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmcu *GiftMissionConfigUpdate) ExecX(ctx context.Context) {
	if err := gmcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmcu *GiftMissionConfigUpdate) defaults() {
	if _, ok := gmcu.mutation.UpdatedAt(); !ok {
		v := giftmissionconfig.UpdateDefaultUpdatedAt()
		gmcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmcu *GiftMissionConfigUpdate) check() error {
	if v, ok := gmcu.mutation.StabilityLevel(); ok {
		if err := giftmissionconfig.StabilityLevelValidator(v); err != nil {
			return &ValidationError{Name: "stability_level", err: fmt.Errorf(`cep_ent: validator failed for field "GiftMissionConfig.stability_level": %w`, err)}
		}
	}
	if v, ok := gmcu.mutation.GpuVersion(); ok {
		if err := giftmissionconfig.GpuVersionValidator(v); err != nil {
			return &ValidationError{Name: "gpu_version", err: fmt.Errorf(`cep_ent: validator failed for field "GiftMissionConfig.gpu_version": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gmcu *GiftMissionConfigUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GiftMissionConfigUpdate {
	gmcu.modifiers = append(gmcu.modifiers, modifiers...)
	return gmcu
}

func (gmcu *GiftMissionConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gmcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(giftmissionconfig.Table, giftmissionconfig.Columns, sqlgraph.NewFieldSpec(giftmissionconfig.FieldID, field.TypeInt64))
	if ps := gmcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmcu.mutation.CreatedBy(); ok {
		_spec.SetField(giftmissionconfig.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := gmcu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(giftmissionconfig.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := gmcu.mutation.UpdatedBy(); ok {
		_spec.SetField(giftmissionconfig.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := gmcu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(giftmissionconfig.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := gmcu.mutation.UpdatedAt(); ok {
		_spec.SetField(giftmissionconfig.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gmcu.mutation.DeletedAt(); ok {
		_spec.SetField(giftmissionconfig.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := gmcu.mutation.StabilityLevel(); ok {
		_spec.SetField(giftmissionconfig.FieldStabilityLevel, field.TypeEnum, value)
	}
	if value, ok := gmcu.mutation.GpuVersion(); ok {
		_spec.SetField(giftmissionconfig.FieldGpuVersion, field.TypeEnum, value)
	}
	if value, ok := gmcu.mutation.GapBase(); ok {
		_spec.SetField(giftmissionconfig.FieldGapBase, field.TypeInt64, value)
	}
	if value, ok := gmcu.mutation.AddedGapBase(); ok {
		_spec.AddField(giftmissionconfig.FieldGapBase, field.TypeInt64, value)
	}
	if value, ok := gmcu.mutation.GapRandomMax(); ok {
		_spec.SetField(giftmissionconfig.FieldGapRandomMax, field.TypeInt64, value)
	}
	if value, ok := gmcu.mutation.AddedGapRandomMax(); ok {
		_spec.AddField(giftmissionconfig.FieldGapRandomMax, field.TypeInt64, value)
	}
	if value, ok := gmcu.mutation.GapRandomMin(); ok {
		_spec.SetField(giftmissionconfig.FieldGapRandomMin, field.TypeInt64, value)
	}
	if value, ok := gmcu.mutation.AddedGapRandomMin(); ok {
		_spec.AddField(giftmissionconfig.FieldGapRandomMin, field.TypeInt64, value)
	}
	if gmcu.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   giftmissionconfig.DevicesTable,
			Columns: []string{giftmissionconfig.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmcu.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !gmcu.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   giftmissionconfig.DevicesTable,
			Columns: []string{giftmissionconfig.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmcu.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   giftmissionconfig.DevicesTable,
			Columns: []string{giftmissionconfig.DevicesColumn},
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
	_spec.AddModifiers(gmcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, gmcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{giftmissionconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gmcu.mutation.done = true
	return n, nil
}

// GiftMissionConfigUpdateOne is the builder for updating a single GiftMissionConfig entity.
type GiftMissionConfigUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GiftMissionConfigMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (gmcuo *GiftMissionConfigUpdateOne) SetCreatedBy(i int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.ResetCreatedBy()
	gmcuo.mutation.SetCreatedBy(i)
	return gmcuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gmcuo *GiftMissionConfigUpdateOne) SetNillableCreatedBy(i *int64) *GiftMissionConfigUpdateOne {
	if i != nil {
		gmcuo.SetCreatedBy(*i)
	}
	return gmcuo
}

// AddCreatedBy adds i to the "created_by" field.
func (gmcuo *GiftMissionConfigUpdateOne) AddCreatedBy(i int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.AddCreatedBy(i)
	return gmcuo
}

// SetUpdatedBy sets the "updated_by" field.
func (gmcuo *GiftMissionConfigUpdateOne) SetUpdatedBy(i int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.ResetUpdatedBy()
	gmcuo.mutation.SetUpdatedBy(i)
	return gmcuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gmcuo *GiftMissionConfigUpdateOne) SetNillableUpdatedBy(i *int64) *GiftMissionConfigUpdateOne {
	if i != nil {
		gmcuo.SetUpdatedBy(*i)
	}
	return gmcuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (gmcuo *GiftMissionConfigUpdateOne) AddUpdatedBy(i int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.AddUpdatedBy(i)
	return gmcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (gmcuo *GiftMissionConfigUpdateOne) SetUpdatedAt(t time.Time) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.SetUpdatedAt(t)
	return gmcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (gmcuo *GiftMissionConfigUpdateOne) SetDeletedAt(t time.Time) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.SetDeletedAt(t)
	return gmcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gmcuo *GiftMissionConfigUpdateOne) SetNillableDeletedAt(t *time.Time) *GiftMissionConfigUpdateOne {
	if t != nil {
		gmcuo.SetDeletedAt(*t)
	}
	return gmcuo
}

// SetStabilityLevel sets the "stability_level" field.
func (gmcuo *GiftMissionConfigUpdateOne) SetStabilityLevel(est enums.DeviceStabilityType) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.SetStabilityLevel(est)
	return gmcuo
}

// SetNillableStabilityLevel sets the "stability_level" field if the given value is not nil.
func (gmcuo *GiftMissionConfigUpdateOne) SetNillableStabilityLevel(est *enums.DeviceStabilityType) *GiftMissionConfigUpdateOne {
	if est != nil {
		gmcuo.SetStabilityLevel(*est)
	}
	return gmcuo
}

// SetGpuVersion sets the "gpu_version" field.
func (gmcuo *GiftMissionConfigUpdateOne) SetGpuVersion(ev enums.GpuVersion) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.SetGpuVersion(ev)
	return gmcuo
}

// SetNillableGpuVersion sets the "gpu_version" field if the given value is not nil.
func (gmcuo *GiftMissionConfigUpdateOne) SetNillableGpuVersion(ev *enums.GpuVersion) *GiftMissionConfigUpdateOne {
	if ev != nil {
		gmcuo.SetGpuVersion(*ev)
	}
	return gmcuo
}

// SetGapBase sets the "gap_base" field.
func (gmcuo *GiftMissionConfigUpdateOne) SetGapBase(i int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.ResetGapBase()
	gmcuo.mutation.SetGapBase(i)
	return gmcuo
}

// SetNillableGapBase sets the "gap_base" field if the given value is not nil.
func (gmcuo *GiftMissionConfigUpdateOne) SetNillableGapBase(i *int64) *GiftMissionConfigUpdateOne {
	if i != nil {
		gmcuo.SetGapBase(*i)
	}
	return gmcuo
}

// AddGapBase adds i to the "gap_base" field.
func (gmcuo *GiftMissionConfigUpdateOne) AddGapBase(i int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.AddGapBase(i)
	return gmcuo
}

// SetGapRandomMax sets the "gap_random_max" field.
func (gmcuo *GiftMissionConfigUpdateOne) SetGapRandomMax(i int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.ResetGapRandomMax()
	gmcuo.mutation.SetGapRandomMax(i)
	return gmcuo
}

// SetNillableGapRandomMax sets the "gap_random_max" field if the given value is not nil.
func (gmcuo *GiftMissionConfigUpdateOne) SetNillableGapRandomMax(i *int64) *GiftMissionConfigUpdateOne {
	if i != nil {
		gmcuo.SetGapRandomMax(*i)
	}
	return gmcuo
}

// AddGapRandomMax adds i to the "gap_random_max" field.
func (gmcuo *GiftMissionConfigUpdateOne) AddGapRandomMax(i int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.AddGapRandomMax(i)
	return gmcuo
}

// SetGapRandomMin sets the "gap_random_min" field.
func (gmcuo *GiftMissionConfigUpdateOne) SetGapRandomMin(i int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.ResetGapRandomMin()
	gmcuo.mutation.SetGapRandomMin(i)
	return gmcuo
}

// SetNillableGapRandomMin sets the "gap_random_min" field if the given value is not nil.
func (gmcuo *GiftMissionConfigUpdateOne) SetNillableGapRandomMin(i *int64) *GiftMissionConfigUpdateOne {
	if i != nil {
		gmcuo.SetGapRandomMin(*i)
	}
	return gmcuo
}

// AddGapRandomMin adds i to the "gap_random_min" field.
func (gmcuo *GiftMissionConfigUpdateOne) AddGapRandomMin(i int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.AddGapRandomMin(i)
	return gmcuo
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (gmcuo *GiftMissionConfigUpdateOne) AddDeviceIDs(ids ...int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.AddDeviceIDs(ids...)
	return gmcuo
}

// AddDevices adds the "devices" edges to the Device entity.
func (gmcuo *GiftMissionConfigUpdateOne) AddDevices(d ...*Device) *GiftMissionConfigUpdateOne {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gmcuo.AddDeviceIDs(ids...)
}

// Mutation returns the GiftMissionConfigMutation object of the builder.
func (gmcuo *GiftMissionConfigUpdateOne) Mutation() *GiftMissionConfigMutation {
	return gmcuo.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (gmcuo *GiftMissionConfigUpdateOne) ClearDevices() *GiftMissionConfigUpdateOne {
	gmcuo.mutation.ClearDevices()
	return gmcuo
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (gmcuo *GiftMissionConfigUpdateOne) RemoveDeviceIDs(ids ...int64) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.RemoveDeviceIDs(ids...)
	return gmcuo
}

// RemoveDevices removes "devices" edges to Device entities.
func (gmcuo *GiftMissionConfigUpdateOne) RemoveDevices(d ...*Device) *GiftMissionConfigUpdateOne {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gmcuo.RemoveDeviceIDs(ids...)
}

// Where appends a list predicates to the GiftMissionConfigUpdate builder.
func (gmcuo *GiftMissionConfigUpdateOne) Where(ps ...predicate.GiftMissionConfig) *GiftMissionConfigUpdateOne {
	gmcuo.mutation.Where(ps...)
	return gmcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gmcuo *GiftMissionConfigUpdateOne) Select(field string, fields ...string) *GiftMissionConfigUpdateOne {
	gmcuo.fields = append([]string{field}, fields...)
	return gmcuo
}

// Save executes the query and returns the updated GiftMissionConfig entity.
func (gmcuo *GiftMissionConfigUpdateOne) Save(ctx context.Context) (*GiftMissionConfig, error) {
	gmcuo.defaults()
	return withHooks(ctx, gmcuo.sqlSave, gmcuo.mutation, gmcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmcuo *GiftMissionConfigUpdateOne) SaveX(ctx context.Context) *GiftMissionConfig {
	node, err := gmcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gmcuo *GiftMissionConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := gmcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmcuo *GiftMissionConfigUpdateOne) ExecX(ctx context.Context) {
	if err := gmcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmcuo *GiftMissionConfigUpdateOne) defaults() {
	if _, ok := gmcuo.mutation.UpdatedAt(); !ok {
		v := giftmissionconfig.UpdateDefaultUpdatedAt()
		gmcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmcuo *GiftMissionConfigUpdateOne) check() error {
	if v, ok := gmcuo.mutation.StabilityLevel(); ok {
		if err := giftmissionconfig.StabilityLevelValidator(v); err != nil {
			return &ValidationError{Name: "stability_level", err: fmt.Errorf(`cep_ent: validator failed for field "GiftMissionConfig.stability_level": %w`, err)}
		}
	}
	if v, ok := gmcuo.mutation.GpuVersion(); ok {
		if err := giftmissionconfig.GpuVersionValidator(v); err != nil {
			return &ValidationError{Name: "gpu_version", err: fmt.Errorf(`cep_ent: validator failed for field "GiftMissionConfig.gpu_version": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gmcuo *GiftMissionConfigUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GiftMissionConfigUpdateOne {
	gmcuo.modifiers = append(gmcuo.modifiers, modifiers...)
	return gmcuo
}

func (gmcuo *GiftMissionConfigUpdateOne) sqlSave(ctx context.Context) (_node *GiftMissionConfig, err error) {
	if err := gmcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(giftmissionconfig.Table, giftmissionconfig.Columns, sqlgraph.NewFieldSpec(giftmissionconfig.FieldID, field.TypeInt64))
	id, ok := gmcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "GiftMissionConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gmcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, giftmissionconfig.FieldID)
		for _, f := range fields {
			if !giftmissionconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != giftmissionconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gmcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmcuo.mutation.CreatedBy(); ok {
		_spec.SetField(giftmissionconfig.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := gmcuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(giftmissionconfig.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := gmcuo.mutation.UpdatedBy(); ok {
		_spec.SetField(giftmissionconfig.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := gmcuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(giftmissionconfig.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := gmcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(giftmissionconfig.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gmcuo.mutation.DeletedAt(); ok {
		_spec.SetField(giftmissionconfig.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := gmcuo.mutation.StabilityLevel(); ok {
		_spec.SetField(giftmissionconfig.FieldStabilityLevel, field.TypeEnum, value)
	}
	if value, ok := gmcuo.mutation.GpuVersion(); ok {
		_spec.SetField(giftmissionconfig.FieldGpuVersion, field.TypeEnum, value)
	}
	if value, ok := gmcuo.mutation.GapBase(); ok {
		_spec.SetField(giftmissionconfig.FieldGapBase, field.TypeInt64, value)
	}
	if value, ok := gmcuo.mutation.AddedGapBase(); ok {
		_spec.AddField(giftmissionconfig.FieldGapBase, field.TypeInt64, value)
	}
	if value, ok := gmcuo.mutation.GapRandomMax(); ok {
		_spec.SetField(giftmissionconfig.FieldGapRandomMax, field.TypeInt64, value)
	}
	if value, ok := gmcuo.mutation.AddedGapRandomMax(); ok {
		_spec.AddField(giftmissionconfig.FieldGapRandomMax, field.TypeInt64, value)
	}
	if value, ok := gmcuo.mutation.GapRandomMin(); ok {
		_spec.SetField(giftmissionconfig.FieldGapRandomMin, field.TypeInt64, value)
	}
	if value, ok := gmcuo.mutation.AddedGapRandomMin(); ok {
		_spec.AddField(giftmissionconfig.FieldGapRandomMin, field.TypeInt64, value)
	}
	if gmcuo.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   giftmissionconfig.DevicesTable,
			Columns: []string{giftmissionconfig.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmcuo.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !gmcuo.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   giftmissionconfig.DevicesTable,
			Columns: []string{giftmissionconfig.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmcuo.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   giftmissionconfig.DevicesTable,
			Columns: []string{giftmissionconfig.DevicesColumn},
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
	_spec.AddModifiers(gmcuo.modifiers...)
	_node = &GiftMissionConfig{config: gmcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gmcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{giftmissionconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gmcuo.mutation.done = true
	return _node, nil
}
