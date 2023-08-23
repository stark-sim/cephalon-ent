// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/costaccount"
	"cephalon-ent/pkg/cep_ent/costbill"
	"cephalon-ent/pkg/cep_ent/missionconsumeorder"
	"cephalon-ent/pkg/cep_ent/rechargeorder"
	"cephalon-ent/pkg/cep_ent/user"
	"cephalon-ent/pkg/enums"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CostBillCreate is the builder for creating a CostBill entity.
type CostBillCreate struct {
	config
	mutation *CostBillMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (cbc *CostBillCreate) SetCreatedBy(i int64) *CostBillCreate {
	cbc.mutation.SetCreatedBy(i)
	return cbc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableCreatedBy(i *int64) *CostBillCreate {
	if i != nil {
		cbc.SetCreatedBy(*i)
	}
	return cbc
}

// SetUpdatedBy sets the "updated_by" field.
func (cbc *CostBillCreate) SetUpdatedBy(i int64) *CostBillCreate {
	cbc.mutation.SetUpdatedBy(i)
	return cbc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableUpdatedBy(i *int64) *CostBillCreate {
	if i != nil {
		cbc.SetUpdatedBy(*i)
	}
	return cbc
}

// SetCreatedAt sets the "created_at" field.
func (cbc *CostBillCreate) SetCreatedAt(t time.Time) *CostBillCreate {
	cbc.mutation.SetCreatedAt(t)
	return cbc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableCreatedAt(t *time.Time) *CostBillCreate {
	if t != nil {
		cbc.SetCreatedAt(*t)
	}
	return cbc
}

// SetUpdatedAt sets the "updated_at" field.
func (cbc *CostBillCreate) SetUpdatedAt(t time.Time) *CostBillCreate {
	cbc.mutation.SetUpdatedAt(t)
	return cbc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableUpdatedAt(t *time.Time) *CostBillCreate {
	if t != nil {
		cbc.SetUpdatedAt(*t)
	}
	return cbc
}

// SetDeletedAt sets the "deleted_at" field.
func (cbc *CostBillCreate) SetDeletedAt(t time.Time) *CostBillCreate {
	cbc.mutation.SetDeletedAt(t)
	return cbc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableDeletedAt(t *time.Time) *CostBillCreate {
	if t != nil {
		cbc.SetDeletedAt(*t)
	}
	return cbc
}

// SetType sets the "type" field.
func (cbc *CostBillCreate) SetType(c costbill.Type) *CostBillCreate {
	cbc.mutation.SetType(c)
	return cbc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableType(c *costbill.Type) *CostBillCreate {
	if c != nil {
		cbc.SetType(*c)
	}
	return cbc
}

// SetIsAdd sets the "is_add" field.
func (cbc *CostBillCreate) SetIsAdd(b bool) *CostBillCreate {
	cbc.mutation.SetIsAdd(b)
	return cbc
}

// SetNillableIsAdd sets the "is_add" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableIsAdd(b *bool) *CostBillCreate {
	if b != nil {
		cbc.SetIsAdd(*b)
	}
	return cbc
}

// SetUserID sets the "user_id" field.
func (cbc *CostBillCreate) SetUserID(i int64) *CostBillCreate {
	cbc.mutation.SetUserID(i)
	return cbc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableUserID(i *int64) *CostBillCreate {
	if i != nil {
		cbc.SetUserID(*i)
	}
	return cbc
}

// SetSerialNumber sets the "serial_number" field.
func (cbc *CostBillCreate) SetSerialNumber(s string) *CostBillCreate {
	cbc.mutation.SetSerialNumber(s)
	return cbc
}

// SetNillableSerialNumber sets the "serial_number" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableSerialNumber(s *string) *CostBillCreate {
	if s != nil {
		cbc.SetSerialNumber(*s)
	}
	return cbc
}

// SetCostAccountID sets the "cost_account_id" field.
func (cbc *CostBillCreate) SetCostAccountID(i int64) *CostBillCreate {
	cbc.mutation.SetCostAccountID(i)
	return cbc
}

// SetNillableCostAccountID sets the "cost_account_id" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableCostAccountID(i *int64) *CostBillCreate {
	if i != nil {
		cbc.SetCostAccountID(*i)
	}
	return cbc
}

// SetPureCep sets the "pure_cep" field.
func (cbc *CostBillCreate) SetPureCep(i int64) *CostBillCreate {
	cbc.mutation.SetPureCep(i)
	return cbc
}

// SetNillablePureCep sets the "pure_cep" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillablePureCep(i *int64) *CostBillCreate {
	if i != nil {
		cbc.SetPureCep(*i)
	}
	return cbc
}

// SetGiftCep sets the "gift_cep" field.
func (cbc *CostBillCreate) SetGiftCep(i int64) *CostBillCreate {
	cbc.mutation.SetGiftCep(i)
	return cbc
}

// SetNillableGiftCep sets the "gift_cep" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableGiftCep(i *int64) *CostBillCreate {
	if i != nil {
		cbc.SetGiftCep(*i)
	}
	return cbc
}

// SetReasonID sets the "reason_id" field.
func (cbc *CostBillCreate) SetReasonID(i int64) *CostBillCreate {
	cbc.mutation.SetReasonID(i)
	return cbc
}

// SetNillableReasonID sets the "reason_id" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableReasonID(i *int64) *CostBillCreate {
	if i != nil {
		cbc.SetReasonID(*i)
	}
	return cbc
}

// SetStatus sets the "status" field.
func (cbc *CostBillCreate) SetStatus(es enums.BillStatus) *CostBillCreate {
	cbc.mutation.SetStatus(es)
	return cbc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableStatus(es *enums.BillStatus) *CostBillCreate {
	if es != nil {
		cbc.SetStatus(*es)
	}
	return cbc
}

// SetMarketBillID sets the "market_bill_id" field.
func (cbc *CostBillCreate) SetMarketBillID(i int64) *CostBillCreate {
	cbc.mutation.SetMarketBillID(i)
	return cbc
}

// SetNillableMarketBillID sets the "market_bill_id" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableMarketBillID(i *int64) *CostBillCreate {
	if i != nil {
		cbc.SetMarketBillID(*i)
	}
	return cbc
}

// SetID sets the "id" field.
func (cbc *CostBillCreate) SetID(i int64) *CostBillCreate {
	cbc.mutation.SetID(i)
	return cbc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cbc *CostBillCreate) SetNillableID(i *int64) *CostBillCreate {
	if i != nil {
		cbc.SetID(*i)
	}
	return cbc
}

// SetUser sets the "user" edge to the User entity.
func (cbc *CostBillCreate) SetUser(u *User) *CostBillCreate {
	return cbc.SetUserID(u.ID)
}

// SetCostAccount sets the "cost_account" edge to the CostAccount entity.
func (cbc *CostBillCreate) SetCostAccount(c *CostAccount) *CostBillCreate {
	return cbc.SetCostAccountID(c.ID)
}

// SetRechargeOrderID sets the "recharge_order" edge to the RechargeOrder entity by ID.
func (cbc *CostBillCreate) SetRechargeOrderID(id int64) *CostBillCreate {
	cbc.mutation.SetRechargeOrderID(id)
	return cbc
}

// SetNillableRechargeOrderID sets the "recharge_order" edge to the RechargeOrder entity by ID if the given value is not nil.
func (cbc *CostBillCreate) SetNillableRechargeOrderID(id *int64) *CostBillCreate {
	if id != nil {
		cbc = cbc.SetRechargeOrderID(*id)
	}
	return cbc
}

// SetRechargeOrder sets the "recharge_order" edge to the RechargeOrder entity.
func (cbc *CostBillCreate) SetRechargeOrder(r *RechargeOrder) *CostBillCreate {
	return cbc.SetRechargeOrderID(r.ID)
}

// SetMissionConsumeOrderID sets the "mission_consume_order" edge to the MissionConsumeOrder entity by ID.
func (cbc *CostBillCreate) SetMissionConsumeOrderID(id int64) *CostBillCreate {
	cbc.mutation.SetMissionConsumeOrderID(id)
	return cbc
}

// SetNillableMissionConsumeOrderID sets the "mission_consume_order" edge to the MissionConsumeOrder entity by ID if the given value is not nil.
func (cbc *CostBillCreate) SetNillableMissionConsumeOrderID(id *int64) *CostBillCreate {
	if id != nil {
		cbc = cbc.SetMissionConsumeOrderID(*id)
	}
	return cbc
}

// SetMissionConsumeOrder sets the "mission_consume_order" edge to the MissionConsumeOrder entity.
func (cbc *CostBillCreate) SetMissionConsumeOrder(m *MissionConsumeOrder) *CostBillCreate {
	return cbc.SetMissionConsumeOrderID(m.ID)
}

// Mutation returns the CostBillMutation object of the builder.
func (cbc *CostBillCreate) Mutation() *CostBillMutation {
	return cbc.mutation
}

// Save creates the CostBill in the database.
func (cbc *CostBillCreate) Save(ctx context.Context) (*CostBill, error) {
	cbc.defaults()
	return withHooks(ctx, cbc.sqlSave, cbc.mutation, cbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cbc *CostBillCreate) SaveX(ctx context.Context) *CostBill {
	v, err := cbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cbc *CostBillCreate) Exec(ctx context.Context) error {
	_, err := cbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbc *CostBillCreate) ExecX(ctx context.Context) {
	if err := cbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cbc *CostBillCreate) defaults() {
	if _, ok := cbc.mutation.CreatedBy(); !ok {
		v := costbill.DefaultCreatedBy
		cbc.mutation.SetCreatedBy(v)
	}
	if _, ok := cbc.mutation.UpdatedBy(); !ok {
		v := costbill.DefaultUpdatedBy
		cbc.mutation.SetUpdatedBy(v)
	}
	if _, ok := cbc.mutation.CreatedAt(); !ok {
		v := costbill.DefaultCreatedAt()
		cbc.mutation.SetCreatedAt(v)
	}
	if _, ok := cbc.mutation.UpdatedAt(); !ok {
		v := costbill.DefaultUpdatedAt()
		cbc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cbc.mutation.DeletedAt(); !ok {
		v := costbill.DefaultDeletedAt
		cbc.mutation.SetDeletedAt(v)
	}
	if _, ok := cbc.mutation.GetType(); !ok {
		v := costbill.DefaultType
		cbc.mutation.SetType(v)
	}
	if _, ok := cbc.mutation.IsAdd(); !ok {
		v := costbill.DefaultIsAdd
		cbc.mutation.SetIsAdd(v)
	}
	if _, ok := cbc.mutation.UserID(); !ok {
		v := costbill.DefaultUserID
		cbc.mutation.SetUserID(v)
	}
	if _, ok := cbc.mutation.SerialNumber(); !ok {
		v := costbill.DefaultSerialNumber
		cbc.mutation.SetSerialNumber(v)
	}
	if _, ok := cbc.mutation.CostAccountID(); !ok {
		v := costbill.DefaultCostAccountID
		cbc.mutation.SetCostAccountID(v)
	}
	if _, ok := cbc.mutation.PureCep(); !ok {
		v := costbill.DefaultPureCep
		cbc.mutation.SetPureCep(v)
	}
	if _, ok := cbc.mutation.GiftCep(); !ok {
		v := costbill.DefaultGiftCep
		cbc.mutation.SetGiftCep(v)
	}
	if _, ok := cbc.mutation.ReasonID(); !ok {
		v := costbill.DefaultReasonID
		cbc.mutation.SetReasonID(v)
	}
	if _, ok := cbc.mutation.Status(); !ok {
		v := costbill.DefaultStatus
		cbc.mutation.SetStatus(v)
	}
	if _, ok := cbc.mutation.MarketBillID(); !ok {
		v := costbill.DefaultMarketBillID
		cbc.mutation.SetMarketBillID(v)
	}
	if _, ok := cbc.mutation.ID(); !ok {
		v := costbill.DefaultID()
		cbc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cbc *CostBillCreate) check() error {
	if _, ok := cbc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "CostBill.created_by"`)}
	}
	if _, ok := cbc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "CostBill.updated_by"`)}
	}
	if _, ok := cbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "CostBill.created_at"`)}
	}
	if _, ok := cbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "CostBill.updated_at"`)}
	}
	if _, ok := cbc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "CostBill.deleted_at"`)}
	}
	if _, ok := cbc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`cep_ent: missing required field "CostBill.type"`)}
	}
	if v, ok := cbc.mutation.GetType(); ok {
		if err := costbill.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "CostBill.type": %w`, err)}
		}
	}
	if _, ok := cbc.mutation.IsAdd(); !ok {
		return &ValidationError{Name: "is_add", err: errors.New(`cep_ent: missing required field "CostBill.is_add"`)}
	}
	if _, ok := cbc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`cep_ent: missing required field "CostBill.user_id"`)}
	}
	if _, ok := cbc.mutation.SerialNumber(); !ok {
		return &ValidationError{Name: "serial_number", err: errors.New(`cep_ent: missing required field "CostBill.serial_number"`)}
	}
	if _, ok := cbc.mutation.CostAccountID(); !ok {
		return &ValidationError{Name: "cost_account_id", err: errors.New(`cep_ent: missing required field "CostBill.cost_account_id"`)}
	}
	if _, ok := cbc.mutation.PureCep(); !ok {
		return &ValidationError{Name: "pure_cep", err: errors.New(`cep_ent: missing required field "CostBill.pure_cep"`)}
	}
	if _, ok := cbc.mutation.GiftCep(); !ok {
		return &ValidationError{Name: "gift_cep", err: errors.New(`cep_ent: missing required field "CostBill.gift_cep"`)}
	}
	if _, ok := cbc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`cep_ent: missing required field "CostBill.status"`)}
	}
	if v, ok := cbc.mutation.Status(); ok {
		if err := costbill.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "CostBill.status": %w`, err)}
		}
	}
	if _, ok := cbc.mutation.MarketBillID(); !ok {
		return &ValidationError{Name: "market_bill_id", err: errors.New(`cep_ent: missing required field "CostBill.market_bill_id"`)}
	}
	if _, ok := cbc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`cep_ent: missing required edge "CostBill.user"`)}
	}
	if _, ok := cbc.mutation.CostAccountID(); !ok {
		return &ValidationError{Name: "cost_account", err: errors.New(`cep_ent: missing required edge "CostBill.cost_account"`)}
	}
	return nil
}

func (cbc *CostBillCreate) sqlSave(ctx context.Context) (*CostBill, error) {
	if err := cbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	cbc.mutation.id = &_node.ID
	cbc.mutation.done = true
	return _node, nil
}

func (cbc *CostBillCreate) createSpec() (*CostBill, *sqlgraph.CreateSpec) {
	var (
		_node = &CostBill{config: cbc.config}
		_spec = sqlgraph.NewCreateSpec(costbill.Table, sqlgraph.NewFieldSpec(costbill.FieldID, field.TypeInt64))
	)
	if id, ok := cbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cbc.mutation.CreatedBy(); ok {
		_spec.SetField(costbill.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := cbc.mutation.UpdatedBy(); ok {
		_spec.SetField(costbill.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := cbc.mutation.CreatedAt(); ok {
		_spec.SetField(costbill.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cbc.mutation.UpdatedAt(); ok {
		_spec.SetField(costbill.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cbc.mutation.DeletedAt(); ok {
		_spec.SetField(costbill.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cbc.mutation.GetType(); ok {
		_spec.SetField(costbill.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := cbc.mutation.IsAdd(); ok {
		_spec.SetField(costbill.FieldIsAdd, field.TypeBool, value)
		_node.IsAdd = value
	}
	if value, ok := cbc.mutation.SerialNumber(); ok {
		_spec.SetField(costbill.FieldSerialNumber, field.TypeString, value)
		_node.SerialNumber = value
	}
	if value, ok := cbc.mutation.PureCep(); ok {
		_spec.SetField(costbill.FieldPureCep, field.TypeInt64, value)
		_node.PureCep = value
	}
	if value, ok := cbc.mutation.GiftCep(); ok {
		_spec.SetField(costbill.FieldGiftCep, field.TypeInt64, value)
		_node.GiftCep = value
	}
	if value, ok := cbc.mutation.Status(); ok {
		_spec.SetField(costbill.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := cbc.mutation.MarketBillID(); ok {
		_spec.SetField(costbill.FieldMarketBillID, field.TypeInt64, value)
		_node.MarketBillID = value
	}
	if nodes := cbc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   costbill.UserTable,
			Columns: []string{costbill.UserColumn},
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
	if nodes := cbc.mutation.CostAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   costbill.CostAccountTable,
			Columns: []string{costbill.CostAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(costaccount.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CostAccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cbc.mutation.RechargeOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   costbill.RechargeOrderTable,
			Columns: []string{costbill.RechargeOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rechargeorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ReasonID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cbc.mutation.MissionConsumeOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   costbill.MissionConsumeOrderTable,
			Columns: []string{costbill.MissionConsumeOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionconsumeorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ReasonID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CostBillCreateBulk is the builder for creating many CostBill entities in bulk.
type CostBillCreateBulk struct {
	config
	builders []*CostBillCreate
}

// Save creates the CostBill entities in the database.
func (cbcb *CostBillCreateBulk) Save(ctx context.Context) ([]*CostBill, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cbcb.builders))
	nodes := make([]*CostBill, len(cbcb.builders))
	mutators := make([]Mutator, len(cbcb.builders))
	for i := range cbcb.builders {
		func(i int, root context.Context) {
			builder := cbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CostBillMutation)
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
					_, err = mutators[i+1].Mutate(root, cbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cbcb *CostBillCreateBulk) SaveX(ctx context.Context) []*CostBill {
	v, err := cbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cbcb *CostBillCreateBulk) Exec(ctx context.Context) error {
	_, err := cbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbcb *CostBillCreateBulk) ExecX(ctx context.Context) {
	if err := cbcb.Exec(ctx); err != nil {
		panic(err)
	}
}