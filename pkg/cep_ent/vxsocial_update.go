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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/rechargeorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/transferorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/vxsocial"
)

// VXSocialUpdate is the builder for updating VXSocial entities.
type VXSocialUpdate struct {
	config
	hooks    []Hook
	mutation *VXSocialMutation
}

// Where appends a list predicates to the VXSocialUpdate builder.
func (vsu *VXSocialUpdate) Where(ps ...predicate.VXSocial) *VXSocialUpdate {
	vsu.mutation.Where(ps...)
	return vsu
}

// SetCreatedBy sets the "created_by" field.
func (vsu *VXSocialUpdate) SetCreatedBy(i int64) *VXSocialUpdate {
	vsu.mutation.ResetCreatedBy()
	vsu.mutation.SetCreatedBy(i)
	return vsu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (vsu *VXSocialUpdate) SetNillableCreatedBy(i *int64) *VXSocialUpdate {
	if i != nil {
		vsu.SetCreatedBy(*i)
	}
	return vsu
}

// AddCreatedBy adds i to the "created_by" field.
func (vsu *VXSocialUpdate) AddCreatedBy(i int64) *VXSocialUpdate {
	vsu.mutation.AddCreatedBy(i)
	return vsu
}

// SetUpdatedBy sets the "updated_by" field.
func (vsu *VXSocialUpdate) SetUpdatedBy(i int64) *VXSocialUpdate {
	vsu.mutation.ResetUpdatedBy()
	vsu.mutation.SetUpdatedBy(i)
	return vsu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (vsu *VXSocialUpdate) SetNillableUpdatedBy(i *int64) *VXSocialUpdate {
	if i != nil {
		vsu.SetUpdatedBy(*i)
	}
	return vsu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (vsu *VXSocialUpdate) AddUpdatedBy(i int64) *VXSocialUpdate {
	vsu.mutation.AddUpdatedBy(i)
	return vsu
}

// SetUpdatedAt sets the "updated_at" field.
func (vsu *VXSocialUpdate) SetUpdatedAt(t time.Time) *VXSocialUpdate {
	vsu.mutation.SetUpdatedAt(t)
	return vsu
}

// SetDeletedAt sets the "deleted_at" field.
func (vsu *VXSocialUpdate) SetDeletedAt(t time.Time) *VXSocialUpdate {
	vsu.mutation.SetDeletedAt(t)
	return vsu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vsu *VXSocialUpdate) SetNillableDeletedAt(t *time.Time) *VXSocialUpdate {
	if t != nil {
		vsu.SetDeletedAt(*t)
	}
	return vsu
}

// SetAppID sets the "app_id" field.
func (vsu *VXSocialUpdate) SetAppID(s string) *VXSocialUpdate {
	vsu.mutation.SetAppID(s)
	return vsu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (vsu *VXSocialUpdate) SetNillableAppID(s *string) *VXSocialUpdate {
	if s != nil {
		vsu.SetAppID(*s)
	}
	return vsu
}

// SetOpenID sets the "open_id" field.
func (vsu *VXSocialUpdate) SetOpenID(s string) *VXSocialUpdate {
	vsu.mutation.SetOpenID(s)
	return vsu
}

// SetNillableOpenID sets the "open_id" field if the given value is not nil.
func (vsu *VXSocialUpdate) SetNillableOpenID(s *string) *VXSocialUpdate {
	if s != nil {
		vsu.SetOpenID(*s)
	}
	return vsu
}

// SetUnionID sets the "union_id" field.
func (vsu *VXSocialUpdate) SetUnionID(s string) *VXSocialUpdate {
	vsu.mutation.SetUnionID(s)
	return vsu
}

// SetNillableUnionID sets the "union_id" field if the given value is not nil.
func (vsu *VXSocialUpdate) SetNillableUnionID(s *string) *VXSocialUpdate {
	if s != nil {
		vsu.SetUnionID(*s)
	}
	return vsu
}

// SetScope sets the "scope" field.
func (vsu *VXSocialUpdate) SetScope(v vxsocial.Scope) *VXSocialUpdate {
	vsu.mutation.SetScope(v)
	return vsu
}

// SetNillableScope sets the "scope" field if the given value is not nil.
func (vsu *VXSocialUpdate) SetNillableScope(v *vxsocial.Scope) *VXSocialUpdate {
	if v != nil {
		vsu.SetScope(*v)
	}
	return vsu
}

// SetSessionKey sets the "session_key" field.
func (vsu *VXSocialUpdate) SetSessionKey(s string) *VXSocialUpdate {
	vsu.mutation.SetSessionKey(s)
	return vsu
}

// SetNillableSessionKey sets the "session_key" field if the given value is not nil.
func (vsu *VXSocialUpdate) SetNillableSessionKey(s *string) *VXSocialUpdate {
	if s != nil {
		vsu.SetSessionKey(*s)
	}
	return vsu
}

// SetAccessToken sets the "access_token" field.
func (vsu *VXSocialUpdate) SetAccessToken(s string) *VXSocialUpdate {
	vsu.mutation.SetAccessToken(s)
	return vsu
}

// SetNillableAccessToken sets the "access_token" field if the given value is not nil.
func (vsu *VXSocialUpdate) SetNillableAccessToken(s *string) *VXSocialUpdate {
	if s != nil {
		vsu.SetAccessToken(*s)
	}
	return vsu
}

// SetRefreshToken sets the "refresh_token" field.
func (vsu *VXSocialUpdate) SetRefreshToken(s string) *VXSocialUpdate {
	vsu.mutation.SetRefreshToken(s)
	return vsu
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (vsu *VXSocialUpdate) SetNillableRefreshToken(s *string) *VXSocialUpdate {
	if s != nil {
		vsu.SetRefreshToken(*s)
	}
	return vsu
}

// SetUserID sets the "user_id" field.
func (vsu *VXSocialUpdate) SetUserID(i int64) *VXSocialUpdate {
	vsu.mutation.SetUserID(i)
	return vsu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (vsu *VXSocialUpdate) SetNillableUserID(i *int64) *VXSocialUpdate {
	if i != nil {
		vsu.SetUserID(*i)
	}
	return vsu
}

// SetUser sets the "user" edge to the User entity.
func (vsu *VXSocialUpdate) SetUser(u *User) *VXSocialUpdate {
	return vsu.SetUserID(u.ID)
}

// AddRechargeOrderIDs adds the "recharge_orders" edge to the RechargeOrder entity by IDs.
func (vsu *VXSocialUpdate) AddRechargeOrderIDs(ids ...int64) *VXSocialUpdate {
	vsu.mutation.AddRechargeOrderIDs(ids...)
	return vsu
}

// AddRechargeOrders adds the "recharge_orders" edges to the RechargeOrder entity.
func (vsu *VXSocialUpdate) AddRechargeOrders(r ...*RechargeOrder) *VXSocialUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return vsu.AddRechargeOrderIDs(ids...)
}

// AddTransferOrderIDs adds the "transfer_orders" edge to the TransferOrder entity by IDs.
func (vsu *VXSocialUpdate) AddTransferOrderIDs(ids ...int64) *VXSocialUpdate {
	vsu.mutation.AddTransferOrderIDs(ids...)
	return vsu
}

// AddTransferOrders adds the "transfer_orders" edges to the TransferOrder entity.
func (vsu *VXSocialUpdate) AddTransferOrders(t ...*TransferOrder) *VXSocialUpdate {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vsu.AddTransferOrderIDs(ids...)
}

// Mutation returns the VXSocialMutation object of the builder.
func (vsu *VXSocialUpdate) Mutation() *VXSocialMutation {
	return vsu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vsu *VXSocialUpdate) ClearUser() *VXSocialUpdate {
	vsu.mutation.ClearUser()
	return vsu
}

// ClearRechargeOrders clears all "recharge_orders" edges to the RechargeOrder entity.
func (vsu *VXSocialUpdate) ClearRechargeOrders() *VXSocialUpdate {
	vsu.mutation.ClearRechargeOrders()
	return vsu
}

// RemoveRechargeOrderIDs removes the "recharge_orders" edge to RechargeOrder entities by IDs.
func (vsu *VXSocialUpdate) RemoveRechargeOrderIDs(ids ...int64) *VXSocialUpdate {
	vsu.mutation.RemoveRechargeOrderIDs(ids...)
	return vsu
}

// RemoveRechargeOrders removes "recharge_orders" edges to RechargeOrder entities.
func (vsu *VXSocialUpdate) RemoveRechargeOrders(r ...*RechargeOrder) *VXSocialUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return vsu.RemoveRechargeOrderIDs(ids...)
}

// ClearTransferOrders clears all "transfer_orders" edges to the TransferOrder entity.
func (vsu *VXSocialUpdate) ClearTransferOrders() *VXSocialUpdate {
	vsu.mutation.ClearTransferOrders()
	return vsu
}

// RemoveTransferOrderIDs removes the "transfer_orders" edge to TransferOrder entities by IDs.
func (vsu *VXSocialUpdate) RemoveTransferOrderIDs(ids ...int64) *VXSocialUpdate {
	vsu.mutation.RemoveTransferOrderIDs(ids...)
	return vsu
}

// RemoveTransferOrders removes "transfer_orders" edges to TransferOrder entities.
func (vsu *VXSocialUpdate) RemoveTransferOrders(t ...*TransferOrder) *VXSocialUpdate {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vsu.RemoveTransferOrderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vsu *VXSocialUpdate) Save(ctx context.Context) (int, error) {
	vsu.defaults()
	return withHooks(ctx, vsu.sqlSave, vsu.mutation, vsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vsu *VXSocialUpdate) SaveX(ctx context.Context) int {
	affected, err := vsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vsu *VXSocialUpdate) Exec(ctx context.Context) error {
	_, err := vsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vsu *VXSocialUpdate) ExecX(ctx context.Context) {
	if err := vsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vsu *VXSocialUpdate) defaults() {
	if _, ok := vsu.mutation.UpdatedAt(); !ok {
		v := vxsocial.UpdateDefaultUpdatedAt()
		vsu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vsu *VXSocialUpdate) check() error {
	if v, ok := vsu.mutation.Scope(); ok {
		if err := vxsocial.ScopeValidator(v); err != nil {
			return &ValidationError{Name: "scope", err: fmt.Errorf(`cep_ent: validator failed for field "VXSocial.scope": %w`, err)}
		}
	}
	if _, ok := vsu.mutation.UserID(); vsu.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "VXSocial.user"`)
	}
	return nil
}

func (vsu *VXSocialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(vxsocial.Table, vxsocial.Columns, sqlgraph.NewFieldSpec(vxsocial.FieldID, field.TypeInt64))
	if ps := vsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vsu.mutation.CreatedBy(); ok {
		_spec.SetField(vxsocial.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := vsu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(vxsocial.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := vsu.mutation.UpdatedBy(); ok {
		_spec.SetField(vxsocial.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := vsu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(vxsocial.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := vsu.mutation.UpdatedAt(); ok {
		_spec.SetField(vxsocial.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vsu.mutation.DeletedAt(); ok {
		_spec.SetField(vxsocial.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := vsu.mutation.AppID(); ok {
		_spec.SetField(vxsocial.FieldAppID, field.TypeString, value)
	}
	if value, ok := vsu.mutation.OpenID(); ok {
		_spec.SetField(vxsocial.FieldOpenID, field.TypeString, value)
	}
	if value, ok := vsu.mutation.UnionID(); ok {
		_spec.SetField(vxsocial.FieldUnionID, field.TypeString, value)
	}
	if value, ok := vsu.mutation.Scope(); ok {
		_spec.SetField(vxsocial.FieldScope, field.TypeEnum, value)
	}
	if value, ok := vsu.mutation.SessionKey(); ok {
		_spec.SetField(vxsocial.FieldSessionKey, field.TypeString, value)
	}
	if value, ok := vsu.mutation.AccessToken(); ok {
		_spec.SetField(vxsocial.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := vsu.mutation.RefreshToken(); ok {
		_spec.SetField(vxsocial.FieldRefreshToken, field.TypeString, value)
	}
	if vsu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vxsocial.UserTable,
			Columns: []string{vxsocial.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vxsocial.UserTable,
			Columns: []string{vxsocial.UserColumn},
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
	if vsu.mutation.RechargeOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.RechargeOrdersTable,
			Columns: []string{vxsocial.RechargeOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rechargeorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsu.mutation.RemovedRechargeOrdersIDs(); len(nodes) > 0 && !vsu.mutation.RechargeOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.RechargeOrdersTable,
			Columns: []string{vxsocial.RechargeOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rechargeorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsu.mutation.RechargeOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.RechargeOrdersTable,
			Columns: []string{vxsocial.RechargeOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rechargeorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vsu.mutation.TransferOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.TransferOrdersTable,
			Columns: []string{vxsocial.TransferOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transferorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsu.mutation.RemovedTransferOrdersIDs(); len(nodes) > 0 && !vsu.mutation.TransferOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.TransferOrdersTable,
			Columns: []string{vxsocial.TransferOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transferorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsu.mutation.TransferOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.TransferOrdersTable,
			Columns: []string{vxsocial.TransferOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transferorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vxsocial.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vsu.mutation.done = true
	return n, nil
}

// VXSocialUpdateOne is the builder for updating a single VXSocial entity.
type VXSocialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VXSocialMutation
}

// SetCreatedBy sets the "created_by" field.
func (vsuo *VXSocialUpdateOne) SetCreatedBy(i int64) *VXSocialUpdateOne {
	vsuo.mutation.ResetCreatedBy()
	vsuo.mutation.SetCreatedBy(i)
	return vsuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (vsuo *VXSocialUpdateOne) SetNillableCreatedBy(i *int64) *VXSocialUpdateOne {
	if i != nil {
		vsuo.SetCreatedBy(*i)
	}
	return vsuo
}

// AddCreatedBy adds i to the "created_by" field.
func (vsuo *VXSocialUpdateOne) AddCreatedBy(i int64) *VXSocialUpdateOne {
	vsuo.mutation.AddCreatedBy(i)
	return vsuo
}

// SetUpdatedBy sets the "updated_by" field.
func (vsuo *VXSocialUpdateOne) SetUpdatedBy(i int64) *VXSocialUpdateOne {
	vsuo.mutation.ResetUpdatedBy()
	vsuo.mutation.SetUpdatedBy(i)
	return vsuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (vsuo *VXSocialUpdateOne) SetNillableUpdatedBy(i *int64) *VXSocialUpdateOne {
	if i != nil {
		vsuo.SetUpdatedBy(*i)
	}
	return vsuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (vsuo *VXSocialUpdateOne) AddUpdatedBy(i int64) *VXSocialUpdateOne {
	vsuo.mutation.AddUpdatedBy(i)
	return vsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (vsuo *VXSocialUpdateOne) SetUpdatedAt(t time.Time) *VXSocialUpdateOne {
	vsuo.mutation.SetUpdatedAt(t)
	return vsuo
}

// SetDeletedAt sets the "deleted_at" field.
func (vsuo *VXSocialUpdateOne) SetDeletedAt(t time.Time) *VXSocialUpdateOne {
	vsuo.mutation.SetDeletedAt(t)
	return vsuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vsuo *VXSocialUpdateOne) SetNillableDeletedAt(t *time.Time) *VXSocialUpdateOne {
	if t != nil {
		vsuo.SetDeletedAt(*t)
	}
	return vsuo
}

// SetAppID sets the "app_id" field.
func (vsuo *VXSocialUpdateOne) SetAppID(s string) *VXSocialUpdateOne {
	vsuo.mutation.SetAppID(s)
	return vsuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (vsuo *VXSocialUpdateOne) SetNillableAppID(s *string) *VXSocialUpdateOne {
	if s != nil {
		vsuo.SetAppID(*s)
	}
	return vsuo
}

// SetOpenID sets the "open_id" field.
func (vsuo *VXSocialUpdateOne) SetOpenID(s string) *VXSocialUpdateOne {
	vsuo.mutation.SetOpenID(s)
	return vsuo
}

// SetNillableOpenID sets the "open_id" field if the given value is not nil.
func (vsuo *VXSocialUpdateOne) SetNillableOpenID(s *string) *VXSocialUpdateOne {
	if s != nil {
		vsuo.SetOpenID(*s)
	}
	return vsuo
}

// SetUnionID sets the "union_id" field.
func (vsuo *VXSocialUpdateOne) SetUnionID(s string) *VXSocialUpdateOne {
	vsuo.mutation.SetUnionID(s)
	return vsuo
}

// SetNillableUnionID sets the "union_id" field if the given value is not nil.
func (vsuo *VXSocialUpdateOne) SetNillableUnionID(s *string) *VXSocialUpdateOne {
	if s != nil {
		vsuo.SetUnionID(*s)
	}
	return vsuo
}

// SetScope sets the "scope" field.
func (vsuo *VXSocialUpdateOne) SetScope(v vxsocial.Scope) *VXSocialUpdateOne {
	vsuo.mutation.SetScope(v)
	return vsuo
}

// SetNillableScope sets the "scope" field if the given value is not nil.
func (vsuo *VXSocialUpdateOne) SetNillableScope(v *vxsocial.Scope) *VXSocialUpdateOne {
	if v != nil {
		vsuo.SetScope(*v)
	}
	return vsuo
}

// SetSessionKey sets the "session_key" field.
func (vsuo *VXSocialUpdateOne) SetSessionKey(s string) *VXSocialUpdateOne {
	vsuo.mutation.SetSessionKey(s)
	return vsuo
}

// SetNillableSessionKey sets the "session_key" field if the given value is not nil.
func (vsuo *VXSocialUpdateOne) SetNillableSessionKey(s *string) *VXSocialUpdateOne {
	if s != nil {
		vsuo.SetSessionKey(*s)
	}
	return vsuo
}

// SetAccessToken sets the "access_token" field.
func (vsuo *VXSocialUpdateOne) SetAccessToken(s string) *VXSocialUpdateOne {
	vsuo.mutation.SetAccessToken(s)
	return vsuo
}

// SetNillableAccessToken sets the "access_token" field if the given value is not nil.
func (vsuo *VXSocialUpdateOne) SetNillableAccessToken(s *string) *VXSocialUpdateOne {
	if s != nil {
		vsuo.SetAccessToken(*s)
	}
	return vsuo
}

// SetRefreshToken sets the "refresh_token" field.
func (vsuo *VXSocialUpdateOne) SetRefreshToken(s string) *VXSocialUpdateOne {
	vsuo.mutation.SetRefreshToken(s)
	return vsuo
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (vsuo *VXSocialUpdateOne) SetNillableRefreshToken(s *string) *VXSocialUpdateOne {
	if s != nil {
		vsuo.SetRefreshToken(*s)
	}
	return vsuo
}

// SetUserID sets the "user_id" field.
func (vsuo *VXSocialUpdateOne) SetUserID(i int64) *VXSocialUpdateOne {
	vsuo.mutation.SetUserID(i)
	return vsuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (vsuo *VXSocialUpdateOne) SetNillableUserID(i *int64) *VXSocialUpdateOne {
	if i != nil {
		vsuo.SetUserID(*i)
	}
	return vsuo
}

// SetUser sets the "user" edge to the User entity.
func (vsuo *VXSocialUpdateOne) SetUser(u *User) *VXSocialUpdateOne {
	return vsuo.SetUserID(u.ID)
}

// AddRechargeOrderIDs adds the "recharge_orders" edge to the RechargeOrder entity by IDs.
func (vsuo *VXSocialUpdateOne) AddRechargeOrderIDs(ids ...int64) *VXSocialUpdateOne {
	vsuo.mutation.AddRechargeOrderIDs(ids...)
	return vsuo
}

// AddRechargeOrders adds the "recharge_orders" edges to the RechargeOrder entity.
func (vsuo *VXSocialUpdateOne) AddRechargeOrders(r ...*RechargeOrder) *VXSocialUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return vsuo.AddRechargeOrderIDs(ids...)
}

// AddTransferOrderIDs adds the "transfer_orders" edge to the TransferOrder entity by IDs.
func (vsuo *VXSocialUpdateOne) AddTransferOrderIDs(ids ...int64) *VXSocialUpdateOne {
	vsuo.mutation.AddTransferOrderIDs(ids...)
	return vsuo
}

// AddTransferOrders adds the "transfer_orders" edges to the TransferOrder entity.
func (vsuo *VXSocialUpdateOne) AddTransferOrders(t ...*TransferOrder) *VXSocialUpdateOne {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vsuo.AddTransferOrderIDs(ids...)
}

// Mutation returns the VXSocialMutation object of the builder.
func (vsuo *VXSocialUpdateOne) Mutation() *VXSocialMutation {
	return vsuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vsuo *VXSocialUpdateOne) ClearUser() *VXSocialUpdateOne {
	vsuo.mutation.ClearUser()
	return vsuo
}

// ClearRechargeOrders clears all "recharge_orders" edges to the RechargeOrder entity.
func (vsuo *VXSocialUpdateOne) ClearRechargeOrders() *VXSocialUpdateOne {
	vsuo.mutation.ClearRechargeOrders()
	return vsuo
}

// RemoveRechargeOrderIDs removes the "recharge_orders" edge to RechargeOrder entities by IDs.
func (vsuo *VXSocialUpdateOne) RemoveRechargeOrderIDs(ids ...int64) *VXSocialUpdateOne {
	vsuo.mutation.RemoveRechargeOrderIDs(ids...)
	return vsuo
}

// RemoveRechargeOrders removes "recharge_orders" edges to RechargeOrder entities.
func (vsuo *VXSocialUpdateOne) RemoveRechargeOrders(r ...*RechargeOrder) *VXSocialUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return vsuo.RemoveRechargeOrderIDs(ids...)
}

// ClearTransferOrders clears all "transfer_orders" edges to the TransferOrder entity.
func (vsuo *VXSocialUpdateOne) ClearTransferOrders() *VXSocialUpdateOne {
	vsuo.mutation.ClearTransferOrders()
	return vsuo
}

// RemoveTransferOrderIDs removes the "transfer_orders" edge to TransferOrder entities by IDs.
func (vsuo *VXSocialUpdateOne) RemoveTransferOrderIDs(ids ...int64) *VXSocialUpdateOne {
	vsuo.mutation.RemoveTransferOrderIDs(ids...)
	return vsuo
}

// RemoveTransferOrders removes "transfer_orders" edges to TransferOrder entities.
func (vsuo *VXSocialUpdateOne) RemoveTransferOrders(t ...*TransferOrder) *VXSocialUpdateOne {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vsuo.RemoveTransferOrderIDs(ids...)
}

// Where appends a list predicates to the VXSocialUpdate builder.
func (vsuo *VXSocialUpdateOne) Where(ps ...predicate.VXSocial) *VXSocialUpdateOne {
	vsuo.mutation.Where(ps...)
	return vsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vsuo *VXSocialUpdateOne) Select(field string, fields ...string) *VXSocialUpdateOne {
	vsuo.fields = append([]string{field}, fields...)
	return vsuo
}

// Save executes the query and returns the updated VXSocial entity.
func (vsuo *VXSocialUpdateOne) Save(ctx context.Context) (*VXSocial, error) {
	vsuo.defaults()
	return withHooks(ctx, vsuo.sqlSave, vsuo.mutation, vsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vsuo *VXSocialUpdateOne) SaveX(ctx context.Context) *VXSocial {
	node, err := vsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vsuo *VXSocialUpdateOne) Exec(ctx context.Context) error {
	_, err := vsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vsuo *VXSocialUpdateOne) ExecX(ctx context.Context) {
	if err := vsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vsuo *VXSocialUpdateOne) defaults() {
	if _, ok := vsuo.mutation.UpdatedAt(); !ok {
		v := vxsocial.UpdateDefaultUpdatedAt()
		vsuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vsuo *VXSocialUpdateOne) check() error {
	if v, ok := vsuo.mutation.Scope(); ok {
		if err := vxsocial.ScopeValidator(v); err != nil {
			return &ValidationError{Name: "scope", err: fmt.Errorf(`cep_ent: validator failed for field "VXSocial.scope": %w`, err)}
		}
	}
	if _, ok := vsuo.mutation.UserID(); vsuo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "VXSocial.user"`)
	}
	return nil
}

func (vsuo *VXSocialUpdateOne) sqlSave(ctx context.Context) (_node *VXSocial, err error) {
	if err := vsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(vxsocial.Table, vxsocial.Columns, sqlgraph.NewFieldSpec(vxsocial.FieldID, field.TypeInt64))
	id, ok := vsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "VXSocial.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vxsocial.FieldID)
		for _, f := range fields {
			if !vxsocial.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != vxsocial.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vsuo.mutation.CreatedBy(); ok {
		_spec.SetField(vxsocial.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := vsuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(vxsocial.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := vsuo.mutation.UpdatedBy(); ok {
		_spec.SetField(vxsocial.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := vsuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(vxsocial.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := vsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(vxsocial.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vsuo.mutation.DeletedAt(); ok {
		_spec.SetField(vxsocial.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := vsuo.mutation.AppID(); ok {
		_spec.SetField(vxsocial.FieldAppID, field.TypeString, value)
	}
	if value, ok := vsuo.mutation.OpenID(); ok {
		_spec.SetField(vxsocial.FieldOpenID, field.TypeString, value)
	}
	if value, ok := vsuo.mutation.UnionID(); ok {
		_spec.SetField(vxsocial.FieldUnionID, field.TypeString, value)
	}
	if value, ok := vsuo.mutation.Scope(); ok {
		_spec.SetField(vxsocial.FieldScope, field.TypeEnum, value)
	}
	if value, ok := vsuo.mutation.SessionKey(); ok {
		_spec.SetField(vxsocial.FieldSessionKey, field.TypeString, value)
	}
	if value, ok := vsuo.mutation.AccessToken(); ok {
		_spec.SetField(vxsocial.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := vsuo.mutation.RefreshToken(); ok {
		_spec.SetField(vxsocial.FieldRefreshToken, field.TypeString, value)
	}
	if vsuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vxsocial.UserTable,
			Columns: []string{vxsocial.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vxsocial.UserTable,
			Columns: []string{vxsocial.UserColumn},
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
	if vsuo.mutation.RechargeOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.RechargeOrdersTable,
			Columns: []string{vxsocial.RechargeOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rechargeorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsuo.mutation.RemovedRechargeOrdersIDs(); len(nodes) > 0 && !vsuo.mutation.RechargeOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.RechargeOrdersTable,
			Columns: []string{vxsocial.RechargeOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rechargeorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsuo.mutation.RechargeOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.RechargeOrdersTable,
			Columns: []string{vxsocial.RechargeOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rechargeorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vsuo.mutation.TransferOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.TransferOrdersTable,
			Columns: []string{vxsocial.TransferOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transferorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsuo.mutation.RemovedTransferOrdersIDs(); len(nodes) > 0 && !vsuo.mutation.TransferOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.TransferOrdersTable,
			Columns: []string{vxsocial.TransferOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transferorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsuo.mutation.TransferOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vxsocial.TransferOrdersTable,
			Columns: []string{vxsocial.TransferOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transferorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &VXSocial{config: vsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vxsocial.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vsuo.mutation.done = true
	return _node, nil
}
