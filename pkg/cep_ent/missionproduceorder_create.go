// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/bill"
	"cephalon-ent/pkg/cep_ent/device"
	"cephalon-ent/pkg/cep_ent/mission"
	"cephalon-ent/pkg/cep_ent/missionbatch"
	"cephalon-ent/pkg/cep_ent/missionconsumeorder"
	"cephalon-ent/pkg/cep_ent/missionproduceorder"
	"cephalon-ent/pkg/cep_ent/missionproduction"
	"cephalon-ent/pkg/cep_ent/user"
	"cephalon-ent/pkg/enums"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MissionProduceOrderCreate is the builder for creating a MissionProduceOrder entity.
type MissionProduceOrderCreate struct {
	config
	mutation *MissionProduceOrderMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (mpoc *MissionProduceOrderCreate) SetCreatedBy(i int64) *MissionProduceOrderCreate {
	mpoc.mutation.SetCreatedBy(i)
	return mpoc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableCreatedBy(i *int64) *MissionProduceOrderCreate {
	if i != nil {
		mpoc.SetCreatedBy(*i)
	}
	return mpoc
}

// SetUpdatedBy sets the "updated_by" field.
func (mpoc *MissionProduceOrderCreate) SetUpdatedBy(i int64) *MissionProduceOrderCreate {
	mpoc.mutation.SetUpdatedBy(i)
	return mpoc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableUpdatedBy(i *int64) *MissionProduceOrderCreate {
	if i != nil {
		mpoc.SetUpdatedBy(*i)
	}
	return mpoc
}

// SetCreatedAt sets the "created_at" field.
func (mpoc *MissionProduceOrderCreate) SetCreatedAt(t time.Time) *MissionProduceOrderCreate {
	mpoc.mutation.SetCreatedAt(t)
	return mpoc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableCreatedAt(t *time.Time) *MissionProduceOrderCreate {
	if t != nil {
		mpoc.SetCreatedAt(*t)
	}
	return mpoc
}

// SetUpdatedAt sets the "updated_at" field.
func (mpoc *MissionProduceOrderCreate) SetUpdatedAt(t time.Time) *MissionProduceOrderCreate {
	mpoc.mutation.SetUpdatedAt(t)
	return mpoc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableUpdatedAt(t *time.Time) *MissionProduceOrderCreate {
	if t != nil {
		mpoc.SetUpdatedAt(*t)
	}
	return mpoc
}

// SetDeletedAt sets the "deleted_at" field.
func (mpoc *MissionProduceOrderCreate) SetDeletedAt(t time.Time) *MissionProduceOrderCreate {
	mpoc.mutation.SetDeletedAt(t)
	return mpoc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableDeletedAt(t *time.Time) *MissionProduceOrderCreate {
	if t != nil {
		mpoc.SetDeletedAt(*t)
	}
	return mpoc
}

// SetUserID sets the "user_id" field.
func (mpoc *MissionProduceOrderCreate) SetUserID(i int64) *MissionProduceOrderCreate {
	mpoc.mutation.SetUserID(i)
	return mpoc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableUserID(i *int64) *MissionProduceOrderCreate {
	if i != nil {
		mpoc.SetUserID(*i)
	}
	return mpoc
}

// SetMissionID sets the "mission_id" field.
func (mpoc *MissionProduceOrderCreate) SetMissionID(i int64) *MissionProduceOrderCreate {
	mpoc.mutation.SetMissionID(i)
	return mpoc
}

// SetNillableMissionID sets the "mission_id" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableMissionID(i *int64) *MissionProduceOrderCreate {
	if i != nil {
		mpoc.SetMissionID(*i)
	}
	return mpoc
}

// SetMissionProductionID sets the "mission_production_id" field.
func (mpoc *MissionProduceOrderCreate) SetMissionProductionID(i int64) *MissionProduceOrderCreate {
	mpoc.mutation.SetMissionProductionID(i)
	return mpoc
}

// SetNillableMissionProductionID sets the "mission_production_id" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableMissionProductionID(i *int64) *MissionProduceOrderCreate {
	if i != nil {
		mpoc.SetMissionProductionID(*i)
	}
	return mpoc
}

// SetStatus sets the "status" field.
func (mpoc *MissionProduceOrderCreate) SetStatus(m missionproduceorder.Status) *MissionProduceOrderCreate {
	mpoc.mutation.SetStatus(m)
	return mpoc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableStatus(m *missionproduceorder.Status) *MissionProduceOrderCreate {
	if m != nil {
		mpoc.SetStatus(*m)
	}
	return mpoc
}

// SetCep sets the "cep" field.
func (mpoc *MissionProduceOrderCreate) SetCep(i int64) *MissionProduceOrderCreate {
	mpoc.mutation.SetCep(i)
	return mpoc
}

// SetNillableCep sets the "cep" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableCep(i *int64) *MissionProduceOrderCreate {
	if i != nil {
		mpoc.SetCep(*i)
	}
	return mpoc
}

// SetType sets the "type" field.
func (mpoc *MissionProduceOrderCreate) SetType(et enums.MissionType) *MissionProduceOrderCreate {
	mpoc.mutation.SetType(et)
	return mpoc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableType(et *enums.MissionType) *MissionProduceOrderCreate {
	if et != nil {
		mpoc.SetType(*et)
	}
	return mpoc
}

// SetIsTime sets the "is_time" field.
func (mpoc *MissionProduceOrderCreate) SetIsTime(b bool) *MissionProduceOrderCreate {
	mpoc.mutation.SetIsTime(b)
	return mpoc
}

// SetNillableIsTime sets the "is_time" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableIsTime(b *bool) *MissionProduceOrderCreate {
	if b != nil {
		mpoc.SetIsTime(*b)
	}
	return mpoc
}

// SetDeviceID sets the "device_id" field.
func (mpoc *MissionProduceOrderCreate) SetDeviceID(i int64) *MissionProduceOrderCreate {
	mpoc.mutation.SetDeviceID(i)
	return mpoc
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableDeviceID(i *int64) *MissionProduceOrderCreate {
	if i != nil {
		mpoc.SetDeviceID(*i)
	}
	return mpoc
}

// SetSerialNumber sets the "serial_number" field.
func (mpoc *MissionProduceOrderCreate) SetSerialNumber(s string) *MissionProduceOrderCreate {
	mpoc.mutation.SetSerialNumber(s)
	return mpoc
}

// SetNillableSerialNumber sets the "serial_number" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableSerialNumber(s *string) *MissionProduceOrderCreate {
	if s != nil {
		mpoc.SetSerialNumber(*s)
	}
	return mpoc
}

// SetMissionConsumeOrderID sets the "mission_consume_order_id" field.
func (mpoc *MissionProduceOrderCreate) SetMissionConsumeOrderID(i int64) *MissionProduceOrderCreate {
	mpoc.mutation.SetMissionConsumeOrderID(i)
	return mpoc
}

// SetNillableMissionConsumeOrderID sets the "mission_consume_order_id" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableMissionConsumeOrderID(i *int64) *MissionProduceOrderCreate {
	if i != nil {
		mpoc.SetMissionConsumeOrderID(*i)
	}
	return mpoc
}

// SetMissionBatchID sets the "mission_batch_id" field.
func (mpoc *MissionProduceOrderCreate) SetMissionBatchID(i int64) *MissionProduceOrderCreate {
	mpoc.mutation.SetMissionBatchID(i)
	return mpoc
}

// SetNillableMissionBatchID sets the "mission_batch_id" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableMissionBatchID(i *int64) *MissionProduceOrderCreate {
	if i != nil {
		mpoc.SetMissionBatchID(*i)
	}
	return mpoc
}

// SetID sets the "id" field.
func (mpoc *MissionProduceOrderCreate) SetID(i int64) *MissionProduceOrderCreate {
	mpoc.mutation.SetID(i)
	return mpoc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mpoc *MissionProduceOrderCreate) SetNillableID(i *int64) *MissionProduceOrderCreate {
	if i != nil {
		mpoc.SetID(*i)
	}
	return mpoc
}

// SetUser sets the "user" edge to the User entity.
func (mpoc *MissionProduceOrderCreate) SetUser(u *User) *MissionProduceOrderCreate {
	return mpoc.SetUserID(u.ID)
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (mpoc *MissionProduceOrderCreate) AddBillIDs(ids ...int64) *MissionProduceOrderCreate {
	mpoc.mutation.AddBillIDs(ids...)
	return mpoc
}

// AddBills adds the "bills" edges to the Bill entity.
func (mpoc *MissionProduceOrderCreate) AddBills(b ...*Bill) *MissionProduceOrderCreate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return mpoc.AddBillIDs(ids...)
}

// SetDevice sets the "device" edge to the Device entity.
func (mpoc *MissionProduceOrderCreate) SetDevice(d *Device) *MissionProduceOrderCreate {
	return mpoc.SetDeviceID(d.ID)
}

// SetMissionConsumeOrder sets the "mission_consume_order" edge to the MissionConsumeOrder entity.
func (mpoc *MissionProduceOrderCreate) SetMissionConsumeOrder(m *MissionConsumeOrder) *MissionProduceOrderCreate {
	return mpoc.SetMissionConsumeOrderID(m.ID)
}

// SetMission sets the "mission" edge to the Mission entity.
func (mpoc *MissionProduceOrderCreate) SetMission(m *Mission) *MissionProduceOrderCreate {
	return mpoc.SetMissionID(m.ID)
}

// SetMissionProduction sets the "mission_production" edge to the MissionProduction entity.
func (mpoc *MissionProduceOrderCreate) SetMissionProduction(m *MissionProduction) *MissionProduceOrderCreate {
	return mpoc.SetMissionProductionID(m.ID)
}

// SetMissionBatch sets the "mission_batch" edge to the MissionBatch entity.
func (mpoc *MissionProduceOrderCreate) SetMissionBatch(m *MissionBatch) *MissionProduceOrderCreate {
	return mpoc.SetMissionBatchID(m.ID)
}

// Mutation returns the MissionProduceOrderMutation object of the builder.
func (mpoc *MissionProduceOrderCreate) Mutation() *MissionProduceOrderMutation {
	return mpoc.mutation
}

// Save creates the MissionProduceOrder in the database.
func (mpoc *MissionProduceOrderCreate) Save(ctx context.Context) (*MissionProduceOrder, error) {
	mpoc.defaults()
	return withHooks(ctx, mpoc.sqlSave, mpoc.mutation, mpoc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mpoc *MissionProduceOrderCreate) SaveX(ctx context.Context) *MissionProduceOrder {
	v, err := mpoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mpoc *MissionProduceOrderCreate) Exec(ctx context.Context) error {
	_, err := mpoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpoc *MissionProduceOrderCreate) ExecX(ctx context.Context) {
	if err := mpoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpoc *MissionProduceOrderCreate) defaults() {
	if _, ok := mpoc.mutation.CreatedBy(); !ok {
		v := missionproduceorder.DefaultCreatedBy
		mpoc.mutation.SetCreatedBy(v)
	}
	if _, ok := mpoc.mutation.UpdatedBy(); !ok {
		v := missionproduceorder.DefaultUpdatedBy
		mpoc.mutation.SetUpdatedBy(v)
	}
	if _, ok := mpoc.mutation.CreatedAt(); !ok {
		v := missionproduceorder.DefaultCreatedAt()
		mpoc.mutation.SetCreatedAt(v)
	}
	if _, ok := mpoc.mutation.UpdatedAt(); !ok {
		v := missionproduceorder.DefaultUpdatedAt()
		mpoc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mpoc.mutation.DeletedAt(); !ok {
		v := missionproduceorder.DefaultDeletedAt
		mpoc.mutation.SetDeletedAt(v)
	}
	if _, ok := mpoc.mutation.UserID(); !ok {
		v := missionproduceorder.DefaultUserID
		mpoc.mutation.SetUserID(v)
	}
	if _, ok := mpoc.mutation.MissionID(); !ok {
		v := missionproduceorder.DefaultMissionID
		mpoc.mutation.SetMissionID(v)
	}
	if _, ok := mpoc.mutation.MissionProductionID(); !ok {
		v := missionproduceorder.DefaultMissionProductionID
		mpoc.mutation.SetMissionProductionID(v)
	}
	if _, ok := mpoc.mutation.Status(); !ok {
		v := missionproduceorder.DefaultStatus
		mpoc.mutation.SetStatus(v)
	}
	if _, ok := mpoc.mutation.Cep(); !ok {
		v := missionproduceorder.DefaultCep
		mpoc.mutation.SetCep(v)
	}
	if _, ok := mpoc.mutation.GetType(); !ok {
		v := missionproduceorder.DefaultType
		mpoc.mutation.SetType(v)
	}
	if _, ok := mpoc.mutation.IsTime(); !ok {
		v := missionproduceorder.DefaultIsTime
		mpoc.mutation.SetIsTime(v)
	}
	if _, ok := mpoc.mutation.DeviceID(); !ok {
		v := missionproduceorder.DefaultDeviceID
		mpoc.mutation.SetDeviceID(v)
	}
	if _, ok := mpoc.mutation.SerialNumber(); !ok {
		v := missionproduceorder.DefaultSerialNumber
		mpoc.mutation.SetSerialNumber(v)
	}
	if _, ok := mpoc.mutation.MissionConsumeOrderID(); !ok {
		v := missionproduceorder.DefaultMissionConsumeOrderID
		mpoc.mutation.SetMissionConsumeOrderID(v)
	}
	if _, ok := mpoc.mutation.MissionBatchID(); !ok {
		v := missionproduceorder.DefaultMissionBatchID
		mpoc.mutation.SetMissionBatchID(v)
	}
	if _, ok := mpoc.mutation.ID(); !ok {
		v := missionproduceorder.DefaultID()
		mpoc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mpoc *MissionProduceOrderCreate) check() error {
	if _, ok := mpoc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.created_by"`)}
	}
	if _, ok := mpoc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.updated_by"`)}
	}
	if _, ok := mpoc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.created_at"`)}
	}
	if _, ok := mpoc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.updated_at"`)}
	}
	if _, ok := mpoc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.deleted_at"`)}
	}
	if _, ok := mpoc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.user_id"`)}
	}
	if _, ok := mpoc.mutation.MissionID(); !ok {
		return &ValidationError{Name: "mission_id", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.mission_id"`)}
	}
	if _, ok := mpoc.mutation.MissionProductionID(); !ok {
		return &ValidationError{Name: "mission_production_id", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.mission_production_id"`)}
	}
	if _, ok := mpoc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.status"`)}
	}
	if v, ok := mpoc.mutation.Status(); ok {
		if err := missionproduceorder.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "MissionProduceOrder.status": %w`, err)}
		}
	}
	if _, ok := mpoc.mutation.Cep(); !ok {
		return &ValidationError{Name: "cep", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.cep"`)}
	}
	if _, ok := mpoc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.type"`)}
	}
	if v, ok := mpoc.mutation.GetType(); ok {
		if err := missionproduceorder.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "MissionProduceOrder.type": %w`, err)}
		}
	}
	if _, ok := mpoc.mutation.IsTime(); !ok {
		return &ValidationError{Name: "is_time", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.is_time"`)}
	}
	if _, ok := mpoc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device_id", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.device_id"`)}
	}
	if _, ok := mpoc.mutation.SerialNumber(); !ok {
		return &ValidationError{Name: "serial_number", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.serial_number"`)}
	}
	if _, ok := mpoc.mutation.MissionConsumeOrderID(); !ok {
		return &ValidationError{Name: "mission_consume_order_id", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.mission_consume_order_id"`)}
	}
	if _, ok := mpoc.mutation.MissionBatchID(); !ok {
		return &ValidationError{Name: "mission_batch_id", err: errors.New(`cep_ent: missing required field "MissionProduceOrder.mission_batch_id"`)}
	}
	if _, ok := mpoc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`cep_ent: missing required edge "MissionProduceOrder.user"`)}
	}
	if _, ok := mpoc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device", err: errors.New(`cep_ent: missing required edge "MissionProduceOrder.device"`)}
	}
	if _, ok := mpoc.mutation.MissionConsumeOrderID(); !ok {
		return &ValidationError{Name: "mission_consume_order", err: errors.New(`cep_ent: missing required edge "MissionProduceOrder.mission_consume_order"`)}
	}
	if _, ok := mpoc.mutation.MissionID(); !ok {
		return &ValidationError{Name: "mission", err: errors.New(`cep_ent: missing required edge "MissionProduceOrder.mission"`)}
	}
	if _, ok := mpoc.mutation.MissionProductionID(); !ok {
		return &ValidationError{Name: "mission_production", err: errors.New(`cep_ent: missing required edge "MissionProduceOrder.mission_production"`)}
	}
	if _, ok := mpoc.mutation.MissionBatchID(); !ok {
		return &ValidationError{Name: "mission_batch", err: errors.New(`cep_ent: missing required edge "MissionProduceOrder.mission_batch"`)}
	}
	return nil
}

func (mpoc *MissionProduceOrderCreate) sqlSave(ctx context.Context) (*MissionProduceOrder, error) {
	if err := mpoc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mpoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mpoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mpoc.mutation.id = &_node.ID
	mpoc.mutation.done = true
	return _node, nil
}

func (mpoc *MissionProduceOrderCreate) createSpec() (*MissionProduceOrder, *sqlgraph.CreateSpec) {
	var (
		_node = &MissionProduceOrder{config: mpoc.config}
		_spec = sqlgraph.NewCreateSpec(missionproduceorder.Table, sqlgraph.NewFieldSpec(missionproduceorder.FieldID, field.TypeInt64))
	)
	if id, ok := mpoc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mpoc.mutation.CreatedBy(); ok {
		_spec.SetField(missionproduceorder.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := mpoc.mutation.UpdatedBy(); ok {
		_spec.SetField(missionproduceorder.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := mpoc.mutation.CreatedAt(); ok {
		_spec.SetField(missionproduceorder.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mpoc.mutation.UpdatedAt(); ok {
		_spec.SetField(missionproduceorder.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mpoc.mutation.DeletedAt(); ok {
		_spec.SetField(missionproduceorder.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := mpoc.mutation.Status(); ok {
		_spec.SetField(missionproduceorder.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := mpoc.mutation.Cep(); ok {
		_spec.SetField(missionproduceorder.FieldCep, field.TypeInt64, value)
		_node.Cep = value
	}
	if value, ok := mpoc.mutation.GetType(); ok {
		_spec.SetField(missionproduceorder.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := mpoc.mutation.IsTime(); ok {
		_spec.SetField(missionproduceorder.FieldIsTime, field.TypeBool, value)
		_node.IsTime = value
	}
	if value, ok := mpoc.mutation.SerialNumber(); ok {
		_spec.SetField(missionproduceorder.FieldSerialNumber, field.TypeString, value)
		_node.SerialNumber = value
	}
	if nodes := mpoc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduceorder.UserTable,
			Columns: []string{missionproduceorder.UserColumn},
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
	if nodes := mpoc.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   missionproduceorder.BillsTable,
			Columns: []string{missionproduceorder.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mpoc.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduceorder.DeviceTable,
			Columns: []string{missionproduceorder.DeviceColumn},
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
	if nodes := mpoc.mutation.MissionConsumeOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduceorder.MissionConsumeOrderTable,
			Columns: []string{missionproduceorder.MissionConsumeOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionconsumeorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MissionConsumeOrderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mpoc.mutation.MissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduceorder.MissionTable,
			Columns: []string{missionproduceorder.MissionColumn},
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
	if nodes := mpoc.mutation.MissionProductionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   missionproduceorder.MissionProductionTable,
			Columns: []string{missionproduceorder.MissionProductionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionproduction.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MissionProductionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mpoc.mutation.MissionBatchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   missionproduceorder.MissionBatchTable,
			Columns: []string{missionproduceorder.MissionBatchColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionbatch.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MissionBatchID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MissionProduceOrderCreateBulk is the builder for creating many MissionProduceOrder entities in bulk.
type MissionProduceOrderCreateBulk struct {
	config
	builders []*MissionProduceOrderCreate
}

// Save creates the MissionProduceOrder entities in the database.
func (mpocb *MissionProduceOrderCreateBulk) Save(ctx context.Context) ([]*MissionProduceOrder, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mpocb.builders))
	nodes := make([]*MissionProduceOrder, len(mpocb.builders))
	mutators := make([]Mutator, len(mpocb.builders))
	for i := range mpocb.builders {
		func(i int, root context.Context) {
			builder := mpocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MissionProduceOrderMutation)
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
					_, err = mutators[i+1].Mutate(root, mpocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mpocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mpocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mpocb *MissionProduceOrderCreateBulk) SaveX(ctx context.Context) []*MissionProduceOrder {
	v, err := mpocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mpocb *MissionProduceOrderCreateBulk) Exec(ctx context.Context) error {
	_, err := mpocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpocb *MissionProduceOrderCreateBulk) ExecX(ctx context.Context) {
	if err := mpocb.Exec(ctx); err != nil {
		panic(err)
	}
}
