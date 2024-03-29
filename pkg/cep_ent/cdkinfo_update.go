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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/cdkinfo"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// CDKInfoUpdate is the builder for updating CDKInfo entities.
type CDKInfoUpdate struct {
	config
	hooks     []Hook
	mutation  *CDKInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CDKInfoUpdate builder.
func (ciu *CDKInfoUpdate) Where(ps ...predicate.CDKInfo) *CDKInfoUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetCreatedBy sets the "created_by" field.
func (ciu *CDKInfoUpdate) SetCreatedBy(i int64) *CDKInfoUpdate {
	ciu.mutation.ResetCreatedBy()
	ciu.mutation.SetCreatedBy(i)
	return ciu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableCreatedBy(i *int64) *CDKInfoUpdate {
	if i != nil {
		ciu.SetCreatedBy(*i)
	}
	return ciu
}

// AddCreatedBy adds i to the "created_by" field.
func (ciu *CDKInfoUpdate) AddCreatedBy(i int64) *CDKInfoUpdate {
	ciu.mutation.AddCreatedBy(i)
	return ciu
}

// SetUpdatedBy sets the "updated_by" field.
func (ciu *CDKInfoUpdate) SetUpdatedBy(i int64) *CDKInfoUpdate {
	ciu.mutation.ResetUpdatedBy()
	ciu.mutation.SetUpdatedBy(i)
	return ciu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableUpdatedBy(i *int64) *CDKInfoUpdate {
	if i != nil {
		ciu.SetUpdatedBy(*i)
	}
	return ciu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ciu *CDKInfoUpdate) AddUpdatedBy(i int64) *CDKInfoUpdate {
	ciu.mutation.AddUpdatedBy(i)
	return ciu
}

// SetUpdatedAt sets the "updated_at" field.
func (ciu *CDKInfoUpdate) SetUpdatedAt(t time.Time) *CDKInfoUpdate {
	ciu.mutation.SetUpdatedAt(t)
	return ciu
}

// SetDeletedAt sets the "deleted_at" field.
func (ciu *CDKInfoUpdate) SetDeletedAt(t time.Time) *CDKInfoUpdate {
	ciu.mutation.SetDeletedAt(t)
	return ciu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableDeletedAt(t *time.Time) *CDKInfoUpdate {
	if t != nil {
		ciu.SetDeletedAt(*t)
	}
	return ciu
}

// SetIssueUserID sets the "issue_user_id" field.
func (ciu *CDKInfoUpdate) SetIssueUserID(i int64) *CDKInfoUpdate {
	ciu.mutation.SetIssueUserID(i)
	return ciu
}

// SetNillableIssueUserID sets the "issue_user_id" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableIssueUserID(i *int64) *CDKInfoUpdate {
	if i != nil {
		ciu.SetIssueUserID(*i)
	}
	return ciu
}

// SetCdkNumber sets the "cdk_number" field.
func (ciu *CDKInfoUpdate) SetCdkNumber(s string) *CDKInfoUpdate {
	ciu.mutation.SetCdkNumber(s)
	return ciu
}

// SetNillableCdkNumber sets the "cdk_number" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableCdkNumber(s *string) *CDKInfoUpdate {
	if s != nil {
		ciu.SetCdkNumber(*s)
	}
	return ciu
}

// SetType sets the "type" field.
func (ciu *CDKInfoUpdate) SetType(et enums.CDKType) *CDKInfoUpdate {
	ciu.mutation.SetType(et)
	return ciu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableType(et *enums.CDKType) *CDKInfoUpdate {
	if et != nil {
		ciu.SetType(*et)
	}
	return ciu
}

// SetGetCep sets the "get_cep" field.
func (ciu *CDKInfoUpdate) SetGetCep(i int64) *CDKInfoUpdate {
	ciu.mutation.ResetGetCep()
	ciu.mutation.SetGetCep(i)
	return ciu
}

// SetNillableGetCep sets the "get_cep" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableGetCep(i *int64) *CDKInfoUpdate {
	if i != nil {
		ciu.SetGetCep(*i)
	}
	return ciu
}

// AddGetCep adds i to the "get_cep" field.
func (ciu *CDKInfoUpdate) AddGetCep(i int64) *CDKInfoUpdate {
	ciu.mutation.AddGetCep(i)
	return ciu
}

// SetGetTime sets the "get_time" field.
func (ciu *CDKInfoUpdate) SetGetTime(i int64) *CDKInfoUpdate {
	ciu.mutation.ResetGetTime()
	ciu.mutation.SetGetTime(i)
	return ciu
}

// SetNillableGetTime sets the "get_time" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableGetTime(i *int64) *CDKInfoUpdate {
	if i != nil {
		ciu.SetGetTime(*i)
	}
	return ciu
}

// AddGetTime adds i to the "get_time" field.
func (ciu *CDKInfoUpdate) AddGetTime(i int64) *CDKInfoUpdate {
	ciu.mutation.AddGetTime(i)
	return ciu
}

// SetBillingType sets the "billing_type" field.
func (ciu *CDKInfoUpdate) SetBillingType(ebt enums.MissionBillingType) *CDKInfoUpdate {
	ciu.mutation.SetBillingType(ebt)
	return ciu
}

// SetNillableBillingType sets the "billing_type" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableBillingType(ebt *enums.MissionBillingType) *CDKInfoUpdate {
	if ebt != nil {
		ciu.SetBillingType(*ebt)
	}
	return ciu
}

// SetExpiredAt sets the "expired_at" field.
func (ciu *CDKInfoUpdate) SetExpiredAt(t time.Time) *CDKInfoUpdate {
	ciu.mutation.SetExpiredAt(t)
	return ciu
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableExpiredAt(t *time.Time) *CDKInfoUpdate {
	if t != nil {
		ciu.SetExpiredAt(*t)
	}
	return ciu
}

// ClearExpiredAt clears the value of the "expired_at" field.
func (ciu *CDKInfoUpdate) ClearExpiredAt() *CDKInfoUpdate {
	ciu.mutation.ClearExpiredAt()
	return ciu
}

// SetUseTimes sets the "use_times" field.
func (ciu *CDKInfoUpdate) SetUseTimes(i int64) *CDKInfoUpdate {
	ciu.mutation.ResetUseTimes()
	ciu.mutation.SetUseTimes(i)
	return ciu
}

// SetNillableUseTimes sets the "use_times" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableUseTimes(i *int64) *CDKInfoUpdate {
	if i != nil {
		ciu.SetUseTimes(*i)
	}
	return ciu
}

// AddUseTimes adds i to the "use_times" field.
func (ciu *CDKInfoUpdate) AddUseTimes(i int64) *CDKInfoUpdate {
	ciu.mutation.AddUseTimes(i)
	return ciu
}

// SetStatus sets the "status" field.
func (ciu *CDKInfoUpdate) SetStatus(es enums.CDKStatus) *CDKInfoUpdate {
	ciu.mutation.SetStatus(es)
	return ciu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableStatus(es *enums.CDKStatus) *CDKInfoUpdate {
	if es != nil {
		ciu.SetStatus(*es)
	}
	return ciu
}

// SetUseUserID sets the "use_user_id" field.
func (ciu *CDKInfoUpdate) SetUseUserID(i int64) *CDKInfoUpdate {
	ciu.mutation.SetUseUserID(i)
	return ciu
}

// SetNillableUseUserID sets the "use_user_id" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableUseUserID(i *int64) *CDKInfoUpdate {
	if i != nil {
		ciu.SetUseUserID(*i)
	}
	return ciu
}

// SetUsedAt sets the "used_at" field.
func (ciu *CDKInfoUpdate) SetUsedAt(t time.Time) *CDKInfoUpdate {
	ciu.mutation.SetUsedAt(t)
	return ciu
}

// SetNillableUsedAt sets the "used_at" field if the given value is not nil.
func (ciu *CDKInfoUpdate) SetNillableUsedAt(t *time.Time) *CDKInfoUpdate {
	if t != nil {
		ciu.SetUsedAt(*t)
	}
	return ciu
}

// ClearUsedAt clears the value of the "used_at" field.
func (ciu *CDKInfoUpdate) ClearUsedAt() *CDKInfoUpdate {
	ciu.mutation.ClearUsedAt()
	return ciu
}

// SetIssueUser sets the "issue_user" edge to the User entity.
func (ciu *CDKInfoUpdate) SetIssueUser(u *User) *CDKInfoUpdate {
	return ciu.SetIssueUserID(u.ID)
}

// SetUseUser sets the "use_user" edge to the User entity.
func (ciu *CDKInfoUpdate) SetUseUser(u *User) *CDKInfoUpdate {
	return ciu.SetUseUserID(u.ID)
}

// Mutation returns the CDKInfoMutation object of the builder.
func (ciu *CDKInfoUpdate) Mutation() *CDKInfoMutation {
	return ciu.mutation
}

// ClearIssueUser clears the "issue_user" edge to the User entity.
func (ciu *CDKInfoUpdate) ClearIssueUser() *CDKInfoUpdate {
	ciu.mutation.ClearIssueUser()
	return ciu
}

// ClearUseUser clears the "use_user" edge to the User entity.
func (ciu *CDKInfoUpdate) ClearUseUser() *CDKInfoUpdate {
	ciu.mutation.ClearUseUser()
	return ciu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *CDKInfoUpdate) Save(ctx context.Context) (int, error) {
	ciu.defaults()
	return withHooks(ctx, ciu.sqlSave, ciu.mutation, ciu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *CDKInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *CDKInfoUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *CDKInfoUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ciu *CDKInfoUpdate) defaults() {
	if _, ok := ciu.mutation.UpdatedAt(); !ok {
		v := cdkinfo.UpdateDefaultUpdatedAt()
		ciu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciu *CDKInfoUpdate) check() error {
	if v, ok := ciu.mutation.GetType(); ok {
		if err := cdkinfo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "CDKInfo.type": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.BillingType(); ok {
		if err := cdkinfo.BillingTypeValidator(v); err != nil {
			return &ValidationError{Name: "billing_type", err: fmt.Errorf(`cep_ent: validator failed for field "CDKInfo.billing_type": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.Status(); ok {
		if err := cdkinfo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "CDKInfo.status": %w`, err)}
		}
	}
	if _, ok := ciu.mutation.IssueUserID(); ciu.mutation.IssueUserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "CDKInfo.issue_user"`)
	}
	if _, ok := ciu.mutation.UseUserID(); ciu.mutation.UseUserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "CDKInfo.use_user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ciu *CDKInfoUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CDKInfoUpdate {
	ciu.modifiers = append(ciu.modifiers, modifiers...)
	return ciu
}

func (ciu *CDKInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ciu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(cdkinfo.Table, cdkinfo.Columns, sqlgraph.NewFieldSpec(cdkinfo.FieldID, field.TypeInt64))
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.CreatedBy(); ok {
		_spec.SetField(cdkinfo.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ciu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(cdkinfo.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ciu.mutation.UpdatedBy(); ok {
		_spec.SetField(cdkinfo.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ciu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(cdkinfo.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ciu.mutation.UpdatedAt(); ok {
		_spec.SetField(cdkinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ciu.mutation.DeletedAt(); ok {
		_spec.SetField(cdkinfo.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := ciu.mutation.CdkNumber(); ok {
		_spec.SetField(cdkinfo.FieldCdkNumber, field.TypeString, value)
	}
	if value, ok := ciu.mutation.GetType(); ok {
		_spec.SetField(cdkinfo.FieldType, field.TypeEnum, value)
	}
	if value, ok := ciu.mutation.GetCep(); ok {
		_spec.SetField(cdkinfo.FieldGetCep, field.TypeInt64, value)
	}
	if value, ok := ciu.mutation.AddedGetCep(); ok {
		_spec.AddField(cdkinfo.FieldGetCep, field.TypeInt64, value)
	}
	if value, ok := ciu.mutation.GetTime(); ok {
		_spec.SetField(cdkinfo.FieldGetTime, field.TypeInt64, value)
	}
	if value, ok := ciu.mutation.AddedGetTime(); ok {
		_spec.AddField(cdkinfo.FieldGetTime, field.TypeInt64, value)
	}
	if value, ok := ciu.mutation.BillingType(); ok {
		_spec.SetField(cdkinfo.FieldBillingType, field.TypeEnum, value)
	}
	if value, ok := ciu.mutation.ExpiredAt(); ok {
		_spec.SetField(cdkinfo.FieldExpiredAt, field.TypeTime, value)
	}
	if ciu.mutation.ExpiredAtCleared() {
		_spec.ClearField(cdkinfo.FieldExpiredAt, field.TypeTime)
	}
	if value, ok := ciu.mutation.UseTimes(); ok {
		_spec.SetField(cdkinfo.FieldUseTimes, field.TypeInt64, value)
	}
	if value, ok := ciu.mutation.AddedUseTimes(); ok {
		_spec.AddField(cdkinfo.FieldUseTimes, field.TypeInt64, value)
	}
	if value, ok := ciu.mutation.Status(); ok {
		_spec.SetField(cdkinfo.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ciu.mutation.UsedAt(); ok {
		_spec.SetField(cdkinfo.FieldUsedAt, field.TypeTime, value)
	}
	if ciu.mutation.UsedAtCleared() {
		_spec.ClearField(cdkinfo.FieldUsedAt, field.TypeTime)
	}
	if ciu.mutation.IssueUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cdkinfo.IssueUserTable,
			Columns: []string{cdkinfo.IssueUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.IssueUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cdkinfo.IssueUserTable,
			Columns: []string{cdkinfo.IssueUserColumn},
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
	if ciu.mutation.UseUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cdkinfo.UseUserTable,
			Columns: []string{cdkinfo.UseUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.UseUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cdkinfo.UseUserTable,
			Columns: []string{cdkinfo.UseUserColumn},
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
	_spec.AddModifiers(ciu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cdkinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ciu.mutation.done = true
	return n, nil
}

// CDKInfoUpdateOne is the builder for updating a single CDKInfo entity.
type CDKInfoUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CDKInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (ciuo *CDKInfoUpdateOne) SetCreatedBy(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.ResetCreatedBy()
	ciuo.mutation.SetCreatedBy(i)
	return ciuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableCreatedBy(i *int64) *CDKInfoUpdateOne {
	if i != nil {
		ciuo.SetCreatedBy(*i)
	}
	return ciuo
}

// AddCreatedBy adds i to the "created_by" field.
func (ciuo *CDKInfoUpdateOne) AddCreatedBy(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.AddCreatedBy(i)
	return ciuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ciuo *CDKInfoUpdateOne) SetUpdatedBy(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.ResetUpdatedBy()
	ciuo.mutation.SetUpdatedBy(i)
	return ciuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableUpdatedBy(i *int64) *CDKInfoUpdateOne {
	if i != nil {
		ciuo.SetUpdatedBy(*i)
	}
	return ciuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ciuo *CDKInfoUpdateOne) AddUpdatedBy(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.AddUpdatedBy(i)
	return ciuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ciuo *CDKInfoUpdateOne) SetUpdatedAt(t time.Time) *CDKInfoUpdateOne {
	ciuo.mutation.SetUpdatedAt(t)
	return ciuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ciuo *CDKInfoUpdateOne) SetDeletedAt(t time.Time) *CDKInfoUpdateOne {
	ciuo.mutation.SetDeletedAt(t)
	return ciuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableDeletedAt(t *time.Time) *CDKInfoUpdateOne {
	if t != nil {
		ciuo.SetDeletedAt(*t)
	}
	return ciuo
}

// SetIssueUserID sets the "issue_user_id" field.
func (ciuo *CDKInfoUpdateOne) SetIssueUserID(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.SetIssueUserID(i)
	return ciuo
}

// SetNillableIssueUserID sets the "issue_user_id" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableIssueUserID(i *int64) *CDKInfoUpdateOne {
	if i != nil {
		ciuo.SetIssueUserID(*i)
	}
	return ciuo
}

// SetCdkNumber sets the "cdk_number" field.
func (ciuo *CDKInfoUpdateOne) SetCdkNumber(s string) *CDKInfoUpdateOne {
	ciuo.mutation.SetCdkNumber(s)
	return ciuo
}

// SetNillableCdkNumber sets the "cdk_number" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableCdkNumber(s *string) *CDKInfoUpdateOne {
	if s != nil {
		ciuo.SetCdkNumber(*s)
	}
	return ciuo
}

// SetType sets the "type" field.
func (ciuo *CDKInfoUpdateOne) SetType(et enums.CDKType) *CDKInfoUpdateOne {
	ciuo.mutation.SetType(et)
	return ciuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableType(et *enums.CDKType) *CDKInfoUpdateOne {
	if et != nil {
		ciuo.SetType(*et)
	}
	return ciuo
}

// SetGetCep sets the "get_cep" field.
func (ciuo *CDKInfoUpdateOne) SetGetCep(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.ResetGetCep()
	ciuo.mutation.SetGetCep(i)
	return ciuo
}

// SetNillableGetCep sets the "get_cep" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableGetCep(i *int64) *CDKInfoUpdateOne {
	if i != nil {
		ciuo.SetGetCep(*i)
	}
	return ciuo
}

// AddGetCep adds i to the "get_cep" field.
func (ciuo *CDKInfoUpdateOne) AddGetCep(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.AddGetCep(i)
	return ciuo
}

// SetGetTime sets the "get_time" field.
func (ciuo *CDKInfoUpdateOne) SetGetTime(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.ResetGetTime()
	ciuo.mutation.SetGetTime(i)
	return ciuo
}

// SetNillableGetTime sets the "get_time" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableGetTime(i *int64) *CDKInfoUpdateOne {
	if i != nil {
		ciuo.SetGetTime(*i)
	}
	return ciuo
}

// AddGetTime adds i to the "get_time" field.
func (ciuo *CDKInfoUpdateOne) AddGetTime(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.AddGetTime(i)
	return ciuo
}

// SetBillingType sets the "billing_type" field.
func (ciuo *CDKInfoUpdateOne) SetBillingType(ebt enums.MissionBillingType) *CDKInfoUpdateOne {
	ciuo.mutation.SetBillingType(ebt)
	return ciuo
}

// SetNillableBillingType sets the "billing_type" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableBillingType(ebt *enums.MissionBillingType) *CDKInfoUpdateOne {
	if ebt != nil {
		ciuo.SetBillingType(*ebt)
	}
	return ciuo
}

// SetExpiredAt sets the "expired_at" field.
func (ciuo *CDKInfoUpdateOne) SetExpiredAt(t time.Time) *CDKInfoUpdateOne {
	ciuo.mutation.SetExpiredAt(t)
	return ciuo
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableExpiredAt(t *time.Time) *CDKInfoUpdateOne {
	if t != nil {
		ciuo.SetExpiredAt(*t)
	}
	return ciuo
}

// ClearExpiredAt clears the value of the "expired_at" field.
func (ciuo *CDKInfoUpdateOne) ClearExpiredAt() *CDKInfoUpdateOne {
	ciuo.mutation.ClearExpiredAt()
	return ciuo
}

// SetUseTimes sets the "use_times" field.
func (ciuo *CDKInfoUpdateOne) SetUseTimes(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.ResetUseTimes()
	ciuo.mutation.SetUseTimes(i)
	return ciuo
}

// SetNillableUseTimes sets the "use_times" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableUseTimes(i *int64) *CDKInfoUpdateOne {
	if i != nil {
		ciuo.SetUseTimes(*i)
	}
	return ciuo
}

// AddUseTimes adds i to the "use_times" field.
func (ciuo *CDKInfoUpdateOne) AddUseTimes(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.AddUseTimes(i)
	return ciuo
}

// SetStatus sets the "status" field.
func (ciuo *CDKInfoUpdateOne) SetStatus(es enums.CDKStatus) *CDKInfoUpdateOne {
	ciuo.mutation.SetStatus(es)
	return ciuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableStatus(es *enums.CDKStatus) *CDKInfoUpdateOne {
	if es != nil {
		ciuo.SetStatus(*es)
	}
	return ciuo
}

// SetUseUserID sets the "use_user_id" field.
func (ciuo *CDKInfoUpdateOne) SetUseUserID(i int64) *CDKInfoUpdateOne {
	ciuo.mutation.SetUseUserID(i)
	return ciuo
}

// SetNillableUseUserID sets the "use_user_id" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableUseUserID(i *int64) *CDKInfoUpdateOne {
	if i != nil {
		ciuo.SetUseUserID(*i)
	}
	return ciuo
}

// SetUsedAt sets the "used_at" field.
func (ciuo *CDKInfoUpdateOne) SetUsedAt(t time.Time) *CDKInfoUpdateOne {
	ciuo.mutation.SetUsedAt(t)
	return ciuo
}

// SetNillableUsedAt sets the "used_at" field if the given value is not nil.
func (ciuo *CDKInfoUpdateOne) SetNillableUsedAt(t *time.Time) *CDKInfoUpdateOne {
	if t != nil {
		ciuo.SetUsedAt(*t)
	}
	return ciuo
}

// ClearUsedAt clears the value of the "used_at" field.
func (ciuo *CDKInfoUpdateOne) ClearUsedAt() *CDKInfoUpdateOne {
	ciuo.mutation.ClearUsedAt()
	return ciuo
}

// SetIssueUser sets the "issue_user" edge to the User entity.
func (ciuo *CDKInfoUpdateOne) SetIssueUser(u *User) *CDKInfoUpdateOne {
	return ciuo.SetIssueUserID(u.ID)
}

// SetUseUser sets the "use_user" edge to the User entity.
func (ciuo *CDKInfoUpdateOne) SetUseUser(u *User) *CDKInfoUpdateOne {
	return ciuo.SetUseUserID(u.ID)
}

// Mutation returns the CDKInfoMutation object of the builder.
func (ciuo *CDKInfoUpdateOne) Mutation() *CDKInfoMutation {
	return ciuo.mutation
}

// ClearIssueUser clears the "issue_user" edge to the User entity.
func (ciuo *CDKInfoUpdateOne) ClearIssueUser() *CDKInfoUpdateOne {
	ciuo.mutation.ClearIssueUser()
	return ciuo
}

// ClearUseUser clears the "use_user" edge to the User entity.
func (ciuo *CDKInfoUpdateOne) ClearUseUser() *CDKInfoUpdateOne {
	ciuo.mutation.ClearUseUser()
	return ciuo
}

// Where appends a list predicates to the CDKInfoUpdate builder.
func (ciuo *CDKInfoUpdateOne) Where(ps ...predicate.CDKInfo) *CDKInfoUpdateOne {
	ciuo.mutation.Where(ps...)
	return ciuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *CDKInfoUpdateOne) Select(field string, fields ...string) *CDKInfoUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated CDKInfo entity.
func (ciuo *CDKInfoUpdateOne) Save(ctx context.Context) (*CDKInfo, error) {
	ciuo.defaults()
	return withHooks(ctx, ciuo.sqlSave, ciuo.mutation, ciuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *CDKInfoUpdateOne) SaveX(ctx context.Context) *CDKInfo {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *CDKInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *CDKInfoUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ciuo *CDKInfoUpdateOne) defaults() {
	if _, ok := ciuo.mutation.UpdatedAt(); !ok {
		v := cdkinfo.UpdateDefaultUpdatedAt()
		ciuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciuo *CDKInfoUpdateOne) check() error {
	if v, ok := ciuo.mutation.GetType(); ok {
		if err := cdkinfo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "CDKInfo.type": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.BillingType(); ok {
		if err := cdkinfo.BillingTypeValidator(v); err != nil {
			return &ValidationError{Name: "billing_type", err: fmt.Errorf(`cep_ent: validator failed for field "CDKInfo.billing_type": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.Status(); ok {
		if err := cdkinfo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "CDKInfo.status": %w`, err)}
		}
	}
	if _, ok := ciuo.mutation.IssueUserID(); ciuo.mutation.IssueUserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "CDKInfo.issue_user"`)
	}
	if _, ok := ciuo.mutation.UseUserID(); ciuo.mutation.UseUserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "CDKInfo.use_user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ciuo *CDKInfoUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CDKInfoUpdateOne {
	ciuo.modifiers = append(ciuo.modifiers, modifiers...)
	return ciuo
}

func (ciuo *CDKInfoUpdateOne) sqlSave(ctx context.Context) (_node *CDKInfo, err error) {
	if err := ciuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(cdkinfo.Table, cdkinfo.Columns, sqlgraph.NewFieldSpec(cdkinfo.FieldID, field.TypeInt64))
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "CDKInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cdkinfo.FieldID)
		for _, f := range fields {
			if !cdkinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != cdkinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.CreatedBy(); ok {
		_spec.SetField(cdkinfo.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ciuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(cdkinfo.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ciuo.mutation.UpdatedBy(); ok {
		_spec.SetField(cdkinfo.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ciuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(cdkinfo.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ciuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cdkinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ciuo.mutation.DeletedAt(); ok {
		_spec.SetField(cdkinfo.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := ciuo.mutation.CdkNumber(); ok {
		_spec.SetField(cdkinfo.FieldCdkNumber, field.TypeString, value)
	}
	if value, ok := ciuo.mutation.GetType(); ok {
		_spec.SetField(cdkinfo.FieldType, field.TypeEnum, value)
	}
	if value, ok := ciuo.mutation.GetCep(); ok {
		_spec.SetField(cdkinfo.FieldGetCep, field.TypeInt64, value)
	}
	if value, ok := ciuo.mutation.AddedGetCep(); ok {
		_spec.AddField(cdkinfo.FieldGetCep, field.TypeInt64, value)
	}
	if value, ok := ciuo.mutation.GetTime(); ok {
		_spec.SetField(cdkinfo.FieldGetTime, field.TypeInt64, value)
	}
	if value, ok := ciuo.mutation.AddedGetTime(); ok {
		_spec.AddField(cdkinfo.FieldGetTime, field.TypeInt64, value)
	}
	if value, ok := ciuo.mutation.BillingType(); ok {
		_spec.SetField(cdkinfo.FieldBillingType, field.TypeEnum, value)
	}
	if value, ok := ciuo.mutation.ExpiredAt(); ok {
		_spec.SetField(cdkinfo.FieldExpiredAt, field.TypeTime, value)
	}
	if ciuo.mutation.ExpiredAtCleared() {
		_spec.ClearField(cdkinfo.FieldExpiredAt, field.TypeTime)
	}
	if value, ok := ciuo.mutation.UseTimes(); ok {
		_spec.SetField(cdkinfo.FieldUseTimes, field.TypeInt64, value)
	}
	if value, ok := ciuo.mutation.AddedUseTimes(); ok {
		_spec.AddField(cdkinfo.FieldUseTimes, field.TypeInt64, value)
	}
	if value, ok := ciuo.mutation.Status(); ok {
		_spec.SetField(cdkinfo.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ciuo.mutation.UsedAt(); ok {
		_spec.SetField(cdkinfo.FieldUsedAt, field.TypeTime, value)
	}
	if ciuo.mutation.UsedAtCleared() {
		_spec.ClearField(cdkinfo.FieldUsedAt, field.TypeTime)
	}
	if ciuo.mutation.IssueUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cdkinfo.IssueUserTable,
			Columns: []string{cdkinfo.IssueUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.IssueUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cdkinfo.IssueUserTable,
			Columns: []string{cdkinfo.IssueUserColumn},
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
	if ciuo.mutation.UseUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cdkinfo.UseUserTable,
			Columns: []string{cdkinfo.UseUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.UseUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cdkinfo.UseUserTable,
			Columns: []string{cdkinfo.UseUserColumn},
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
	_spec.AddModifiers(ciuo.modifiers...)
	_node = &CDKInfo{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cdkinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ciuo.mutation.done = true
	return _node, nil
}
