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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionproduceorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionproduction"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// MissionProductionUpdate is the builder for updating MissionProduction entities.
type MissionProductionUpdate struct {
	config
	hooks     []Hook
	mutation  *MissionProductionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MissionProductionUpdate builder.
func (mpu *MissionProductionUpdate) Where(ps ...predicate.MissionProduction) *MissionProductionUpdate {
	mpu.mutation.Where(ps...)
	return mpu
}

// SetCreatedBy sets the "created_by" field.
func (mpu *MissionProductionUpdate) SetCreatedBy(i int64) *MissionProductionUpdate {
	mpu.mutation.ResetCreatedBy()
	mpu.mutation.SetCreatedBy(i)
	return mpu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableCreatedBy(i *int64) *MissionProductionUpdate {
	if i != nil {
		mpu.SetCreatedBy(*i)
	}
	return mpu
}

// AddCreatedBy adds i to the "created_by" field.
func (mpu *MissionProductionUpdate) AddCreatedBy(i int64) *MissionProductionUpdate {
	mpu.mutation.AddCreatedBy(i)
	return mpu
}

// SetUpdatedBy sets the "updated_by" field.
func (mpu *MissionProductionUpdate) SetUpdatedBy(i int64) *MissionProductionUpdate {
	mpu.mutation.ResetUpdatedBy()
	mpu.mutation.SetUpdatedBy(i)
	return mpu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableUpdatedBy(i *int64) *MissionProductionUpdate {
	if i != nil {
		mpu.SetUpdatedBy(*i)
	}
	return mpu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (mpu *MissionProductionUpdate) AddUpdatedBy(i int64) *MissionProductionUpdate {
	mpu.mutation.AddUpdatedBy(i)
	return mpu
}

// SetUpdatedAt sets the "updated_at" field.
func (mpu *MissionProductionUpdate) SetUpdatedAt(t time.Time) *MissionProductionUpdate {
	mpu.mutation.SetUpdatedAt(t)
	return mpu
}

// SetDeletedAt sets the "deleted_at" field.
func (mpu *MissionProductionUpdate) SetDeletedAt(t time.Time) *MissionProductionUpdate {
	mpu.mutation.SetDeletedAt(t)
	return mpu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableDeletedAt(t *time.Time) *MissionProductionUpdate {
	if t != nil {
		mpu.SetDeletedAt(*t)
	}
	return mpu
}

// SetMissionID sets the "mission_id" field.
func (mpu *MissionProductionUpdate) SetMissionID(i int64) *MissionProductionUpdate {
	mpu.mutation.SetMissionID(i)
	return mpu
}

// SetUserID sets the "user_id" field.
func (mpu *MissionProductionUpdate) SetUserID(i int64) *MissionProductionUpdate {
	mpu.mutation.SetUserID(i)
	return mpu
}

// SetStartedAt sets the "started_at" field.
func (mpu *MissionProductionUpdate) SetStartedAt(t time.Time) *MissionProductionUpdate {
	mpu.mutation.SetStartedAt(t)
	return mpu
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableStartedAt(t *time.Time) *MissionProductionUpdate {
	if t != nil {
		mpu.SetStartedAt(*t)
	}
	return mpu
}

// SetFinishedAt sets the "finished_at" field.
func (mpu *MissionProductionUpdate) SetFinishedAt(t time.Time) *MissionProductionUpdate {
	mpu.mutation.SetFinishedAt(t)
	return mpu
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableFinishedAt(t *time.Time) *MissionProductionUpdate {
	if t != nil {
		mpu.SetFinishedAt(*t)
	}
	return mpu
}

// SetState sets the "state" field.
func (mpu *MissionProductionUpdate) SetState(es enums.MissionState) *MissionProductionUpdate {
	mpu.mutation.SetState(es)
	return mpu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableState(es *enums.MissionState) *MissionProductionUpdate {
	if es != nil {
		mpu.SetState(*es)
	}
	return mpu
}

// SetDeviceID sets the "device_id" field.
func (mpu *MissionProductionUpdate) SetDeviceID(i int64) *MissionProductionUpdate {
	mpu.mutation.ResetDeviceID()
	mpu.mutation.SetDeviceID(i)
	return mpu
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableDeviceID(i *int64) *MissionProductionUpdate {
	if i != nil {
		mpu.SetDeviceID(*i)
	}
	return mpu
}

// AddDeviceID adds i to the "device_id" field.
func (mpu *MissionProductionUpdate) AddDeviceID(i int64) *MissionProductionUpdate {
	mpu.mutation.AddDeviceID(i)
	return mpu
}

// SetGpuVersion sets the "gpu_version" field.
func (mpu *MissionProductionUpdate) SetGpuVersion(ev enums.GpuVersion) *MissionProductionUpdate {
	mpu.mutation.SetGpuVersion(ev)
	return mpu
}

// SetNillableGpuVersion sets the "gpu_version" field if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableGpuVersion(ev *enums.GpuVersion) *MissionProductionUpdate {
	if ev != nil {
		mpu.SetGpuVersion(*ev)
	}
	return mpu
}

// SetUrls sets the "urls" field.
func (mpu *MissionProductionUpdate) SetUrls(s string) *MissionProductionUpdate {
	mpu.mutation.SetUrls(s)
	return mpu
}

// SetNillableUrls sets the "urls" field if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableUrls(s *string) *MissionProductionUpdate {
	if s != nil {
		mpu.SetUrls(*s)
	}
	return mpu
}

// SetRespStatusCode sets the "resp_status_code" field.
func (mpu *MissionProductionUpdate) SetRespStatusCode(i int32) *MissionProductionUpdate {
	mpu.mutation.ResetRespStatusCode()
	mpu.mutation.SetRespStatusCode(i)
	return mpu
}

// SetNillableRespStatusCode sets the "resp_status_code" field if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableRespStatusCode(i *int32) *MissionProductionUpdate {
	if i != nil {
		mpu.SetRespStatusCode(*i)
	}
	return mpu
}

// AddRespStatusCode adds i to the "resp_status_code" field.
func (mpu *MissionProductionUpdate) AddRespStatusCode(i int32) *MissionProductionUpdate {
	mpu.mutation.AddRespStatusCode(i)
	return mpu
}

// SetRespBody sets the "resp_body" field.
func (mpu *MissionProductionUpdate) SetRespBody(s string) *MissionProductionUpdate {
	mpu.mutation.SetRespBody(s)
	return mpu
}

// SetNillableRespBody sets the "resp_body" field if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableRespBody(s *string) *MissionProductionUpdate {
	if s != nil {
		mpu.SetRespBody(*s)
	}
	return mpu
}

// SetMission sets the "mission" edge to the Mission entity.
func (mpu *MissionProductionUpdate) SetMission(m *Mission) *MissionProductionUpdate {
	return mpu.SetMissionID(m.ID)
}

// SetUser sets the "user" edge to the User entity.
func (mpu *MissionProductionUpdate) SetUser(u *User) *MissionProductionUpdate {
	return mpu.SetUserID(u.ID)
}

// SetMissionProduceOrderID sets the "mission_produce_order" edge to the MissionProduceOrder entity by ID.
func (mpu *MissionProductionUpdate) SetMissionProduceOrderID(id int64) *MissionProductionUpdate {
	mpu.mutation.SetMissionProduceOrderID(id)
	return mpu
}

// SetNillableMissionProduceOrderID sets the "mission_produce_order" edge to the MissionProduceOrder entity by ID if the given value is not nil.
func (mpu *MissionProductionUpdate) SetNillableMissionProduceOrderID(id *int64) *MissionProductionUpdate {
	if id != nil {
		mpu = mpu.SetMissionProduceOrderID(*id)
	}
	return mpu
}

// SetMissionProduceOrder sets the "mission_produce_order" edge to the MissionProduceOrder entity.
func (mpu *MissionProductionUpdate) SetMissionProduceOrder(m *MissionProduceOrder) *MissionProductionUpdate {
	return mpu.SetMissionProduceOrderID(m.ID)
}

// Mutation returns the MissionProductionMutation object of the builder.
func (mpu *MissionProductionUpdate) Mutation() *MissionProductionMutation {
	return mpu.mutation
}

// ClearMission clears the "mission" edge to the Mission entity.
func (mpu *MissionProductionUpdate) ClearMission() *MissionProductionUpdate {
	mpu.mutation.ClearMission()
	return mpu
}

// ClearUser clears the "user" edge to the User entity.
func (mpu *MissionProductionUpdate) ClearUser() *MissionProductionUpdate {
	mpu.mutation.ClearUser()
	return mpu
}

// ClearMissionProduceOrder clears the "mission_produce_order" edge to the MissionProduceOrder entity.
func (mpu *MissionProductionUpdate) ClearMissionProduceOrder() *MissionProductionUpdate {
	mpu.mutation.ClearMissionProduceOrder()
	return mpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mpu *MissionProductionUpdate) Save(ctx context.Context) (int, error) {
	mpu.defaults()
	return withHooks(ctx, mpu.sqlSave, mpu.mutation, mpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpu *MissionProductionUpdate) SaveX(ctx context.Context) int {
	affected, err := mpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mpu *MissionProductionUpdate) Exec(ctx context.Context) error {
	_, err := mpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpu *MissionProductionUpdate) ExecX(ctx context.Context) {
	if err := mpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpu *MissionProductionUpdate) defaults() {
	if _, ok := mpu.mutation.UpdatedAt(); !ok {
		v := missionproduction.UpdateDefaultUpdatedAt()
		mpu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mpu *MissionProductionUpdate) check() error {
	if v, ok := mpu.mutation.State(); ok {
		if err := missionproduction.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`cep_ent: validator failed for field "MissionProduction.state": %w`, err)}
		}
	}
	if v, ok := mpu.mutation.GpuVersion(); ok {
		if err := missionproduction.GpuVersionValidator(v); err != nil {
			return &ValidationError{Name: "gpu_version", err: fmt.Errorf(`cep_ent: validator failed for field "MissionProduction.gpu_version": %w`, err)}
		}
	}
	if _, ok := mpu.mutation.MissionID(); mpu.mutation.MissionCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "MissionProduction.mission"`)
	}
	if _, ok := mpu.mutation.UserID(); mpu.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "MissionProduction.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mpu *MissionProductionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MissionProductionUpdate {
	mpu.modifiers = append(mpu.modifiers, modifiers...)
	return mpu
}

func (mpu *MissionProductionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(missionproduction.Table, missionproduction.Columns, sqlgraph.NewFieldSpec(missionproduction.FieldID, field.TypeInt64))
	if ps := mpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mpu.mutation.CreatedBy(); ok {
		_spec.SetField(missionproduction.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mpu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(missionproduction.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mpu.mutation.UpdatedBy(); ok {
		_spec.SetField(missionproduction.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mpu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(missionproduction.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mpu.mutation.UpdatedAt(); ok {
		_spec.SetField(missionproduction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mpu.mutation.DeletedAt(); ok {
		_spec.SetField(missionproduction.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := mpu.mutation.StartedAt(); ok {
		_spec.SetField(missionproduction.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := mpu.mutation.FinishedAt(); ok {
		_spec.SetField(missionproduction.FieldFinishedAt, field.TypeTime, value)
	}
	if value, ok := mpu.mutation.State(); ok {
		_spec.SetField(missionproduction.FieldState, field.TypeEnum, value)
	}
	if value, ok := mpu.mutation.DeviceID(); ok {
		_spec.SetField(missionproduction.FieldDeviceID, field.TypeInt64, value)
	}
	if value, ok := mpu.mutation.AddedDeviceID(); ok {
		_spec.AddField(missionproduction.FieldDeviceID, field.TypeInt64, value)
	}
	if value, ok := mpu.mutation.GpuVersion(); ok {
		_spec.SetField(missionproduction.FieldGpuVersion, field.TypeEnum, value)
	}
	if value, ok := mpu.mutation.Urls(); ok {
		_spec.SetField(missionproduction.FieldUrls, field.TypeString, value)
	}
	if value, ok := mpu.mutation.RespStatusCode(); ok {
		_spec.SetField(missionproduction.FieldRespStatusCode, field.TypeInt32, value)
	}
	if value, ok := mpu.mutation.AddedRespStatusCode(); ok {
		_spec.AddField(missionproduction.FieldRespStatusCode, field.TypeInt32, value)
	}
	if value, ok := mpu.mutation.RespBody(); ok {
		_spec.SetField(missionproduction.FieldRespBody, field.TypeString, value)
	}
	if mpu.mutation.MissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduction.MissionTable,
			Columns: []string{missionproduction.MissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpu.mutation.MissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduction.MissionTable,
			Columns: []string{missionproduction.MissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mpu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduction.UserTable,
			Columns: []string{missionproduction.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduction.UserTable,
			Columns: []string{missionproduction.UserColumn},
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
	if mpu.mutation.MissionProduceOrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   missionproduction.MissionProduceOrderTable,
			Columns: []string{missionproduction.MissionProduceOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionproduceorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpu.mutation.MissionProduceOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   missionproduction.MissionProduceOrderTable,
			Columns: []string{missionproduction.MissionProduceOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionproduceorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mpu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{missionproduction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mpu.mutation.done = true
	return n, nil
}

// MissionProductionUpdateOne is the builder for updating a single MissionProduction entity.
type MissionProductionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MissionProductionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (mpuo *MissionProductionUpdateOne) SetCreatedBy(i int64) *MissionProductionUpdateOne {
	mpuo.mutation.ResetCreatedBy()
	mpuo.mutation.SetCreatedBy(i)
	return mpuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableCreatedBy(i *int64) *MissionProductionUpdateOne {
	if i != nil {
		mpuo.SetCreatedBy(*i)
	}
	return mpuo
}

// AddCreatedBy adds i to the "created_by" field.
func (mpuo *MissionProductionUpdateOne) AddCreatedBy(i int64) *MissionProductionUpdateOne {
	mpuo.mutation.AddCreatedBy(i)
	return mpuo
}

// SetUpdatedBy sets the "updated_by" field.
func (mpuo *MissionProductionUpdateOne) SetUpdatedBy(i int64) *MissionProductionUpdateOne {
	mpuo.mutation.ResetUpdatedBy()
	mpuo.mutation.SetUpdatedBy(i)
	return mpuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableUpdatedBy(i *int64) *MissionProductionUpdateOne {
	if i != nil {
		mpuo.SetUpdatedBy(*i)
	}
	return mpuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (mpuo *MissionProductionUpdateOne) AddUpdatedBy(i int64) *MissionProductionUpdateOne {
	mpuo.mutation.AddUpdatedBy(i)
	return mpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (mpuo *MissionProductionUpdateOne) SetUpdatedAt(t time.Time) *MissionProductionUpdateOne {
	mpuo.mutation.SetUpdatedAt(t)
	return mpuo
}

// SetDeletedAt sets the "deleted_at" field.
func (mpuo *MissionProductionUpdateOne) SetDeletedAt(t time.Time) *MissionProductionUpdateOne {
	mpuo.mutation.SetDeletedAt(t)
	return mpuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableDeletedAt(t *time.Time) *MissionProductionUpdateOne {
	if t != nil {
		mpuo.SetDeletedAt(*t)
	}
	return mpuo
}

// SetMissionID sets the "mission_id" field.
func (mpuo *MissionProductionUpdateOne) SetMissionID(i int64) *MissionProductionUpdateOne {
	mpuo.mutation.SetMissionID(i)
	return mpuo
}

// SetUserID sets the "user_id" field.
func (mpuo *MissionProductionUpdateOne) SetUserID(i int64) *MissionProductionUpdateOne {
	mpuo.mutation.SetUserID(i)
	return mpuo
}

// SetStartedAt sets the "started_at" field.
func (mpuo *MissionProductionUpdateOne) SetStartedAt(t time.Time) *MissionProductionUpdateOne {
	mpuo.mutation.SetStartedAt(t)
	return mpuo
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableStartedAt(t *time.Time) *MissionProductionUpdateOne {
	if t != nil {
		mpuo.SetStartedAt(*t)
	}
	return mpuo
}

// SetFinishedAt sets the "finished_at" field.
func (mpuo *MissionProductionUpdateOne) SetFinishedAt(t time.Time) *MissionProductionUpdateOne {
	mpuo.mutation.SetFinishedAt(t)
	return mpuo
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableFinishedAt(t *time.Time) *MissionProductionUpdateOne {
	if t != nil {
		mpuo.SetFinishedAt(*t)
	}
	return mpuo
}

// SetState sets the "state" field.
func (mpuo *MissionProductionUpdateOne) SetState(es enums.MissionState) *MissionProductionUpdateOne {
	mpuo.mutation.SetState(es)
	return mpuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableState(es *enums.MissionState) *MissionProductionUpdateOne {
	if es != nil {
		mpuo.SetState(*es)
	}
	return mpuo
}

// SetDeviceID sets the "device_id" field.
func (mpuo *MissionProductionUpdateOne) SetDeviceID(i int64) *MissionProductionUpdateOne {
	mpuo.mutation.ResetDeviceID()
	mpuo.mutation.SetDeviceID(i)
	return mpuo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableDeviceID(i *int64) *MissionProductionUpdateOne {
	if i != nil {
		mpuo.SetDeviceID(*i)
	}
	return mpuo
}

// AddDeviceID adds i to the "device_id" field.
func (mpuo *MissionProductionUpdateOne) AddDeviceID(i int64) *MissionProductionUpdateOne {
	mpuo.mutation.AddDeviceID(i)
	return mpuo
}

// SetGpuVersion sets the "gpu_version" field.
func (mpuo *MissionProductionUpdateOne) SetGpuVersion(ev enums.GpuVersion) *MissionProductionUpdateOne {
	mpuo.mutation.SetGpuVersion(ev)
	return mpuo
}

// SetNillableGpuVersion sets the "gpu_version" field if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableGpuVersion(ev *enums.GpuVersion) *MissionProductionUpdateOne {
	if ev != nil {
		mpuo.SetGpuVersion(*ev)
	}
	return mpuo
}

// SetUrls sets the "urls" field.
func (mpuo *MissionProductionUpdateOne) SetUrls(s string) *MissionProductionUpdateOne {
	mpuo.mutation.SetUrls(s)
	return mpuo
}

// SetNillableUrls sets the "urls" field if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableUrls(s *string) *MissionProductionUpdateOne {
	if s != nil {
		mpuo.SetUrls(*s)
	}
	return mpuo
}

// SetRespStatusCode sets the "resp_status_code" field.
func (mpuo *MissionProductionUpdateOne) SetRespStatusCode(i int32) *MissionProductionUpdateOne {
	mpuo.mutation.ResetRespStatusCode()
	mpuo.mutation.SetRespStatusCode(i)
	return mpuo
}

// SetNillableRespStatusCode sets the "resp_status_code" field if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableRespStatusCode(i *int32) *MissionProductionUpdateOne {
	if i != nil {
		mpuo.SetRespStatusCode(*i)
	}
	return mpuo
}

// AddRespStatusCode adds i to the "resp_status_code" field.
func (mpuo *MissionProductionUpdateOne) AddRespStatusCode(i int32) *MissionProductionUpdateOne {
	mpuo.mutation.AddRespStatusCode(i)
	return mpuo
}

// SetRespBody sets the "resp_body" field.
func (mpuo *MissionProductionUpdateOne) SetRespBody(s string) *MissionProductionUpdateOne {
	mpuo.mutation.SetRespBody(s)
	return mpuo
}

// SetNillableRespBody sets the "resp_body" field if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableRespBody(s *string) *MissionProductionUpdateOne {
	if s != nil {
		mpuo.SetRespBody(*s)
	}
	return mpuo
}

// SetMission sets the "mission" edge to the Mission entity.
func (mpuo *MissionProductionUpdateOne) SetMission(m *Mission) *MissionProductionUpdateOne {
	return mpuo.SetMissionID(m.ID)
}

// SetUser sets the "user" edge to the User entity.
func (mpuo *MissionProductionUpdateOne) SetUser(u *User) *MissionProductionUpdateOne {
	return mpuo.SetUserID(u.ID)
}

// SetMissionProduceOrderID sets the "mission_produce_order" edge to the MissionProduceOrder entity by ID.
func (mpuo *MissionProductionUpdateOne) SetMissionProduceOrderID(id int64) *MissionProductionUpdateOne {
	mpuo.mutation.SetMissionProduceOrderID(id)
	return mpuo
}

// SetNillableMissionProduceOrderID sets the "mission_produce_order" edge to the MissionProduceOrder entity by ID if the given value is not nil.
func (mpuo *MissionProductionUpdateOne) SetNillableMissionProduceOrderID(id *int64) *MissionProductionUpdateOne {
	if id != nil {
		mpuo = mpuo.SetMissionProduceOrderID(*id)
	}
	return mpuo
}

// SetMissionProduceOrder sets the "mission_produce_order" edge to the MissionProduceOrder entity.
func (mpuo *MissionProductionUpdateOne) SetMissionProduceOrder(m *MissionProduceOrder) *MissionProductionUpdateOne {
	return mpuo.SetMissionProduceOrderID(m.ID)
}

// Mutation returns the MissionProductionMutation object of the builder.
func (mpuo *MissionProductionUpdateOne) Mutation() *MissionProductionMutation {
	return mpuo.mutation
}

// ClearMission clears the "mission" edge to the Mission entity.
func (mpuo *MissionProductionUpdateOne) ClearMission() *MissionProductionUpdateOne {
	mpuo.mutation.ClearMission()
	return mpuo
}

// ClearUser clears the "user" edge to the User entity.
func (mpuo *MissionProductionUpdateOne) ClearUser() *MissionProductionUpdateOne {
	mpuo.mutation.ClearUser()
	return mpuo
}

// ClearMissionProduceOrder clears the "mission_produce_order" edge to the MissionProduceOrder entity.
func (mpuo *MissionProductionUpdateOne) ClearMissionProduceOrder() *MissionProductionUpdateOne {
	mpuo.mutation.ClearMissionProduceOrder()
	return mpuo
}

// Where appends a list predicates to the MissionProductionUpdate builder.
func (mpuo *MissionProductionUpdateOne) Where(ps ...predicate.MissionProduction) *MissionProductionUpdateOne {
	mpuo.mutation.Where(ps...)
	return mpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mpuo *MissionProductionUpdateOne) Select(field string, fields ...string) *MissionProductionUpdateOne {
	mpuo.fields = append([]string{field}, fields...)
	return mpuo
}

// Save executes the query and returns the updated MissionProduction entity.
func (mpuo *MissionProductionUpdateOne) Save(ctx context.Context) (*MissionProduction, error) {
	mpuo.defaults()
	return withHooks(ctx, mpuo.sqlSave, mpuo.mutation, mpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpuo *MissionProductionUpdateOne) SaveX(ctx context.Context) *MissionProduction {
	node, err := mpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mpuo *MissionProductionUpdateOne) Exec(ctx context.Context) error {
	_, err := mpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpuo *MissionProductionUpdateOne) ExecX(ctx context.Context) {
	if err := mpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpuo *MissionProductionUpdateOne) defaults() {
	if _, ok := mpuo.mutation.UpdatedAt(); !ok {
		v := missionproduction.UpdateDefaultUpdatedAt()
		mpuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mpuo *MissionProductionUpdateOne) check() error {
	if v, ok := mpuo.mutation.State(); ok {
		if err := missionproduction.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`cep_ent: validator failed for field "MissionProduction.state": %w`, err)}
		}
	}
	if v, ok := mpuo.mutation.GpuVersion(); ok {
		if err := missionproduction.GpuVersionValidator(v); err != nil {
			return &ValidationError{Name: "gpu_version", err: fmt.Errorf(`cep_ent: validator failed for field "MissionProduction.gpu_version": %w`, err)}
		}
	}
	if _, ok := mpuo.mutation.MissionID(); mpuo.mutation.MissionCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "MissionProduction.mission"`)
	}
	if _, ok := mpuo.mutation.UserID(); mpuo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "MissionProduction.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mpuo *MissionProductionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MissionProductionUpdateOne {
	mpuo.modifiers = append(mpuo.modifiers, modifiers...)
	return mpuo
}

func (mpuo *MissionProductionUpdateOne) sqlSave(ctx context.Context) (_node *MissionProduction, err error) {
	if err := mpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(missionproduction.Table, missionproduction.Columns, sqlgraph.NewFieldSpec(missionproduction.FieldID, field.TypeInt64))
	id, ok := mpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "MissionProduction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, missionproduction.FieldID)
		for _, f := range fields {
			if !missionproduction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != missionproduction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mpuo.mutation.CreatedBy(); ok {
		_spec.SetField(missionproduction.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mpuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(missionproduction.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mpuo.mutation.UpdatedBy(); ok {
		_spec.SetField(missionproduction.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mpuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(missionproduction.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(missionproduction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mpuo.mutation.DeletedAt(); ok {
		_spec.SetField(missionproduction.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := mpuo.mutation.StartedAt(); ok {
		_spec.SetField(missionproduction.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := mpuo.mutation.FinishedAt(); ok {
		_spec.SetField(missionproduction.FieldFinishedAt, field.TypeTime, value)
	}
	if value, ok := mpuo.mutation.State(); ok {
		_spec.SetField(missionproduction.FieldState, field.TypeEnum, value)
	}
	if value, ok := mpuo.mutation.DeviceID(); ok {
		_spec.SetField(missionproduction.FieldDeviceID, field.TypeInt64, value)
	}
	if value, ok := mpuo.mutation.AddedDeviceID(); ok {
		_spec.AddField(missionproduction.FieldDeviceID, field.TypeInt64, value)
	}
	if value, ok := mpuo.mutation.GpuVersion(); ok {
		_spec.SetField(missionproduction.FieldGpuVersion, field.TypeEnum, value)
	}
	if value, ok := mpuo.mutation.Urls(); ok {
		_spec.SetField(missionproduction.FieldUrls, field.TypeString, value)
	}
	if value, ok := mpuo.mutation.RespStatusCode(); ok {
		_spec.SetField(missionproduction.FieldRespStatusCode, field.TypeInt32, value)
	}
	if value, ok := mpuo.mutation.AddedRespStatusCode(); ok {
		_spec.AddField(missionproduction.FieldRespStatusCode, field.TypeInt32, value)
	}
	if value, ok := mpuo.mutation.RespBody(); ok {
		_spec.SetField(missionproduction.FieldRespBody, field.TypeString, value)
	}
	if mpuo.mutation.MissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduction.MissionTable,
			Columns: []string{missionproduction.MissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpuo.mutation.MissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduction.MissionTable,
			Columns: []string{missionproduction.MissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mpuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduction.UserTable,
			Columns: []string{missionproduction.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduction.UserTable,
			Columns: []string{missionproduction.UserColumn},
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
	if mpuo.mutation.MissionProduceOrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   missionproduction.MissionProduceOrderTable,
			Columns: []string{missionproduction.MissionProduceOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionproduceorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpuo.mutation.MissionProduceOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   missionproduction.MissionProduceOrderTable,
			Columns: []string{missionproduction.MissionProduceOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionproduceorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mpuo.modifiers...)
	_node = &MissionProduction{config: mpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{missionproduction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mpuo.mutation.done = true
	return _node, nil
}
