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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/incomemanage"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// IncomeManageUpdate is the builder for updating IncomeManage entities.
type IncomeManageUpdate struct {
	config
	hooks     []Hook
	mutation  *IncomeManageMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the IncomeManageUpdate builder.
func (imu *IncomeManageUpdate) Where(ps ...predicate.IncomeManage) *IncomeManageUpdate {
	imu.mutation.Where(ps...)
	return imu
}

// SetCreatedBy sets the "created_by" field.
func (imu *IncomeManageUpdate) SetCreatedBy(i int64) *IncomeManageUpdate {
	imu.mutation.ResetCreatedBy()
	imu.mutation.SetCreatedBy(i)
	return imu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableCreatedBy(i *int64) *IncomeManageUpdate {
	if i != nil {
		imu.SetCreatedBy(*i)
	}
	return imu
}

// AddCreatedBy adds i to the "created_by" field.
func (imu *IncomeManageUpdate) AddCreatedBy(i int64) *IncomeManageUpdate {
	imu.mutation.AddCreatedBy(i)
	return imu
}

// SetUpdatedBy sets the "updated_by" field.
func (imu *IncomeManageUpdate) SetUpdatedBy(i int64) *IncomeManageUpdate {
	imu.mutation.ResetUpdatedBy()
	imu.mutation.SetUpdatedBy(i)
	return imu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableUpdatedBy(i *int64) *IncomeManageUpdate {
	if i != nil {
		imu.SetUpdatedBy(*i)
	}
	return imu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (imu *IncomeManageUpdate) AddUpdatedBy(i int64) *IncomeManageUpdate {
	imu.mutation.AddUpdatedBy(i)
	return imu
}

// SetUpdatedAt sets the "updated_at" field.
func (imu *IncomeManageUpdate) SetUpdatedAt(t time.Time) *IncomeManageUpdate {
	imu.mutation.SetUpdatedAt(t)
	return imu
}

// SetDeletedAt sets the "deleted_at" field.
func (imu *IncomeManageUpdate) SetDeletedAt(t time.Time) *IncomeManageUpdate {
	imu.mutation.SetDeletedAt(t)
	return imu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableDeletedAt(t *time.Time) *IncomeManageUpdate {
	if t != nil {
		imu.SetDeletedAt(*t)
	}
	return imu
}

// SetUserID sets the "user_id" field.
func (imu *IncomeManageUpdate) SetUserID(i int64) *IncomeManageUpdate {
	imu.mutation.SetUserID(i)
	return imu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableUserID(i *int64) *IncomeManageUpdate {
	if i != nil {
		imu.SetUserID(*i)
	}
	return imu
}

// SetPhone sets the "phone" field.
func (imu *IncomeManageUpdate) SetPhone(s string) *IncomeManageUpdate {
	imu.mutation.SetPhone(s)
	return imu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillablePhone(s *string) *IncomeManageUpdate {
	if s != nil {
		imu.SetPhone(*s)
	}
	return imu
}

// SetType sets the "type" field.
func (imu *IncomeManageUpdate) SetType(emt enums.IncomeManageType) *IncomeManageUpdate {
	imu.mutation.SetType(emt)
	return imu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableType(emt *enums.IncomeManageType) *IncomeManageUpdate {
	if emt != nil {
		imu.SetType(*emt)
	}
	return imu
}

// SetAmount sets the "amount" field.
func (imu *IncomeManageUpdate) SetAmount(i int64) *IncomeManageUpdate {
	imu.mutation.ResetAmount()
	imu.mutation.SetAmount(i)
	return imu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableAmount(i *int64) *IncomeManageUpdate {
	if i != nil {
		imu.SetAmount(*i)
	}
	return imu
}

// AddAmount adds i to the "amount" field.
func (imu *IncomeManageUpdate) AddAmount(i int64) *IncomeManageUpdate {
	imu.mutation.AddAmount(i)
	return imu
}

// SetReason sets the "reason" field.
func (imu *IncomeManageUpdate) SetReason(s string) *IncomeManageUpdate {
	imu.mutation.SetReason(s)
	return imu
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableReason(s *string) *IncomeManageUpdate {
	if s != nil {
		imu.SetReason(*s)
	}
	return imu
}

// SetCurrentBalance sets the "current_balance" field.
func (imu *IncomeManageUpdate) SetCurrentBalance(i int64) *IncomeManageUpdate {
	imu.mutation.ResetCurrentBalance()
	imu.mutation.SetCurrentBalance(i)
	return imu
}

// SetNillableCurrentBalance sets the "current_balance" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableCurrentBalance(i *int64) *IncomeManageUpdate {
	if i != nil {
		imu.SetCurrentBalance(*i)
	}
	return imu
}

// AddCurrentBalance adds i to the "current_balance" field.
func (imu *IncomeManageUpdate) AddCurrentBalance(i int64) *IncomeManageUpdate {
	imu.mutation.AddCurrentBalance(i)
	return imu
}

// SetLastEditedAt sets the "last_edited_at" field.
func (imu *IncomeManageUpdate) SetLastEditedAt(t time.Time) *IncomeManageUpdate {
	imu.mutation.SetLastEditedAt(t)
	return imu
}

// SetNillableLastEditedAt sets the "last_edited_at" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableLastEditedAt(t *time.Time) *IncomeManageUpdate {
	if t != nil {
		imu.SetLastEditedAt(*t)
	}
	return imu
}

// SetRejectReason sets the "reject_reason" field.
func (imu *IncomeManageUpdate) SetRejectReason(s string) *IncomeManageUpdate {
	imu.mutation.SetRejectReason(s)
	return imu
}

// SetNillableRejectReason sets the "reject_reason" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableRejectReason(s *string) *IncomeManageUpdate {
	if s != nil {
		imu.SetRejectReason(*s)
	}
	return imu
}

// SetStatus sets the "status" field.
func (imu *IncomeManageUpdate) SetStatus(ems enums.IncomeManageStatus) *IncomeManageUpdate {
	imu.mutation.SetStatus(ems)
	return imu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableStatus(ems *enums.IncomeManageStatus) *IncomeManageUpdate {
	if ems != nil {
		imu.SetStatus(*ems)
	}
	return imu
}

// SetApproveUserID sets the "approve_user_id" field.
func (imu *IncomeManageUpdate) SetApproveUserID(i int64) *IncomeManageUpdate {
	imu.mutation.SetApproveUserID(i)
	return imu
}

// SetNillableApproveUserID sets the "approve_user_id" field if the given value is not nil.
func (imu *IncomeManageUpdate) SetNillableApproveUserID(i *int64) *IncomeManageUpdate {
	if i != nil {
		imu.SetApproveUserID(*i)
	}
	return imu
}

// SetUser sets the "user" edge to the User entity.
func (imu *IncomeManageUpdate) SetUser(u *User) *IncomeManageUpdate {
	return imu.SetUserID(u.ID)
}

// SetApproveUser sets the "approve_user" edge to the User entity.
func (imu *IncomeManageUpdate) SetApproveUser(u *User) *IncomeManageUpdate {
	return imu.SetApproveUserID(u.ID)
}

// Mutation returns the IncomeManageMutation object of the builder.
func (imu *IncomeManageUpdate) Mutation() *IncomeManageMutation {
	return imu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (imu *IncomeManageUpdate) ClearUser() *IncomeManageUpdate {
	imu.mutation.ClearUser()
	return imu
}

// ClearApproveUser clears the "approve_user" edge to the User entity.
func (imu *IncomeManageUpdate) ClearApproveUser() *IncomeManageUpdate {
	imu.mutation.ClearApproveUser()
	return imu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (imu *IncomeManageUpdate) Save(ctx context.Context) (int, error) {
	imu.defaults()
	return withHooks(ctx, imu.sqlSave, imu.mutation, imu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (imu *IncomeManageUpdate) SaveX(ctx context.Context) int {
	affected, err := imu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (imu *IncomeManageUpdate) Exec(ctx context.Context) error {
	_, err := imu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (imu *IncomeManageUpdate) ExecX(ctx context.Context) {
	if err := imu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (imu *IncomeManageUpdate) defaults() {
	if _, ok := imu.mutation.UpdatedAt(); !ok {
		v := incomemanage.UpdateDefaultUpdatedAt()
		imu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (imu *IncomeManageUpdate) check() error {
	if v, ok := imu.mutation.GetType(); ok {
		if err := incomemanage.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "IncomeManage.type": %w`, err)}
		}
	}
	if v, ok := imu.mutation.Status(); ok {
		if err := incomemanage.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "IncomeManage.status": %w`, err)}
		}
	}
	if _, ok := imu.mutation.UserID(); imu.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "IncomeManage.user"`)
	}
	if _, ok := imu.mutation.ApproveUserID(); imu.mutation.ApproveUserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "IncomeManage.approve_user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (imu *IncomeManageUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *IncomeManageUpdate {
	imu.modifiers = append(imu.modifiers, modifiers...)
	return imu
}

func (imu *IncomeManageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := imu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(incomemanage.Table, incomemanage.Columns, sqlgraph.NewFieldSpec(incomemanage.FieldID, field.TypeInt64))
	if ps := imu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := imu.mutation.CreatedBy(); ok {
		_spec.SetField(incomemanage.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := imu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(incomemanage.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := imu.mutation.UpdatedBy(); ok {
		_spec.SetField(incomemanage.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := imu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(incomemanage.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := imu.mutation.UpdatedAt(); ok {
		_spec.SetField(incomemanage.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := imu.mutation.DeletedAt(); ok {
		_spec.SetField(incomemanage.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := imu.mutation.Phone(); ok {
		_spec.SetField(incomemanage.FieldPhone, field.TypeString, value)
	}
	if value, ok := imu.mutation.GetType(); ok {
		_spec.SetField(incomemanage.FieldType, field.TypeEnum, value)
	}
	if value, ok := imu.mutation.Amount(); ok {
		_spec.SetField(incomemanage.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := imu.mutation.AddedAmount(); ok {
		_spec.AddField(incomemanage.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := imu.mutation.Reason(); ok {
		_spec.SetField(incomemanage.FieldReason, field.TypeString, value)
	}
	if value, ok := imu.mutation.CurrentBalance(); ok {
		_spec.SetField(incomemanage.FieldCurrentBalance, field.TypeInt64, value)
	}
	if value, ok := imu.mutation.AddedCurrentBalance(); ok {
		_spec.AddField(incomemanage.FieldCurrentBalance, field.TypeInt64, value)
	}
	if value, ok := imu.mutation.LastEditedAt(); ok {
		_spec.SetField(incomemanage.FieldLastEditedAt, field.TypeTime, value)
	}
	if value, ok := imu.mutation.RejectReason(); ok {
		_spec.SetField(incomemanage.FieldRejectReason, field.TypeString, value)
	}
	if value, ok := imu.mutation.Status(); ok {
		_spec.SetField(incomemanage.FieldStatus, field.TypeEnum, value)
	}
	if imu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incomemanage.UserTable,
			Columns: []string{incomemanage.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incomemanage.UserTable,
			Columns: []string{incomemanage.UserColumn},
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
	if imu.mutation.ApproveUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incomemanage.ApproveUserTable,
			Columns: []string{incomemanage.ApproveUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imu.mutation.ApproveUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incomemanage.ApproveUserTable,
			Columns: []string{incomemanage.ApproveUserColumn},
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
	_spec.AddModifiers(imu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, imu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{incomemanage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	imu.mutation.done = true
	return n, nil
}

// IncomeManageUpdateOne is the builder for updating a single IncomeManage entity.
type IncomeManageUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *IncomeManageMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (imuo *IncomeManageUpdateOne) SetCreatedBy(i int64) *IncomeManageUpdateOne {
	imuo.mutation.ResetCreatedBy()
	imuo.mutation.SetCreatedBy(i)
	return imuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableCreatedBy(i *int64) *IncomeManageUpdateOne {
	if i != nil {
		imuo.SetCreatedBy(*i)
	}
	return imuo
}

// AddCreatedBy adds i to the "created_by" field.
func (imuo *IncomeManageUpdateOne) AddCreatedBy(i int64) *IncomeManageUpdateOne {
	imuo.mutation.AddCreatedBy(i)
	return imuo
}

// SetUpdatedBy sets the "updated_by" field.
func (imuo *IncomeManageUpdateOne) SetUpdatedBy(i int64) *IncomeManageUpdateOne {
	imuo.mutation.ResetUpdatedBy()
	imuo.mutation.SetUpdatedBy(i)
	return imuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableUpdatedBy(i *int64) *IncomeManageUpdateOne {
	if i != nil {
		imuo.SetUpdatedBy(*i)
	}
	return imuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (imuo *IncomeManageUpdateOne) AddUpdatedBy(i int64) *IncomeManageUpdateOne {
	imuo.mutation.AddUpdatedBy(i)
	return imuo
}

// SetUpdatedAt sets the "updated_at" field.
func (imuo *IncomeManageUpdateOne) SetUpdatedAt(t time.Time) *IncomeManageUpdateOne {
	imuo.mutation.SetUpdatedAt(t)
	return imuo
}

// SetDeletedAt sets the "deleted_at" field.
func (imuo *IncomeManageUpdateOne) SetDeletedAt(t time.Time) *IncomeManageUpdateOne {
	imuo.mutation.SetDeletedAt(t)
	return imuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableDeletedAt(t *time.Time) *IncomeManageUpdateOne {
	if t != nil {
		imuo.SetDeletedAt(*t)
	}
	return imuo
}

// SetUserID sets the "user_id" field.
func (imuo *IncomeManageUpdateOne) SetUserID(i int64) *IncomeManageUpdateOne {
	imuo.mutation.SetUserID(i)
	return imuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableUserID(i *int64) *IncomeManageUpdateOne {
	if i != nil {
		imuo.SetUserID(*i)
	}
	return imuo
}

// SetPhone sets the "phone" field.
func (imuo *IncomeManageUpdateOne) SetPhone(s string) *IncomeManageUpdateOne {
	imuo.mutation.SetPhone(s)
	return imuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillablePhone(s *string) *IncomeManageUpdateOne {
	if s != nil {
		imuo.SetPhone(*s)
	}
	return imuo
}

// SetType sets the "type" field.
func (imuo *IncomeManageUpdateOne) SetType(emt enums.IncomeManageType) *IncomeManageUpdateOne {
	imuo.mutation.SetType(emt)
	return imuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableType(emt *enums.IncomeManageType) *IncomeManageUpdateOne {
	if emt != nil {
		imuo.SetType(*emt)
	}
	return imuo
}

// SetAmount sets the "amount" field.
func (imuo *IncomeManageUpdateOne) SetAmount(i int64) *IncomeManageUpdateOne {
	imuo.mutation.ResetAmount()
	imuo.mutation.SetAmount(i)
	return imuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableAmount(i *int64) *IncomeManageUpdateOne {
	if i != nil {
		imuo.SetAmount(*i)
	}
	return imuo
}

// AddAmount adds i to the "amount" field.
func (imuo *IncomeManageUpdateOne) AddAmount(i int64) *IncomeManageUpdateOne {
	imuo.mutation.AddAmount(i)
	return imuo
}

// SetReason sets the "reason" field.
func (imuo *IncomeManageUpdateOne) SetReason(s string) *IncomeManageUpdateOne {
	imuo.mutation.SetReason(s)
	return imuo
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableReason(s *string) *IncomeManageUpdateOne {
	if s != nil {
		imuo.SetReason(*s)
	}
	return imuo
}

// SetCurrentBalance sets the "current_balance" field.
func (imuo *IncomeManageUpdateOne) SetCurrentBalance(i int64) *IncomeManageUpdateOne {
	imuo.mutation.ResetCurrentBalance()
	imuo.mutation.SetCurrentBalance(i)
	return imuo
}

// SetNillableCurrentBalance sets the "current_balance" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableCurrentBalance(i *int64) *IncomeManageUpdateOne {
	if i != nil {
		imuo.SetCurrentBalance(*i)
	}
	return imuo
}

// AddCurrentBalance adds i to the "current_balance" field.
func (imuo *IncomeManageUpdateOne) AddCurrentBalance(i int64) *IncomeManageUpdateOne {
	imuo.mutation.AddCurrentBalance(i)
	return imuo
}

// SetLastEditedAt sets the "last_edited_at" field.
func (imuo *IncomeManageUpdateOne) SetLastEditedAt(t time.Time) *IncomeManageUpdateOne {
	imuo.mutation.SetLastEditedAt(t)
	return imuo
}

// SetNillableLastEditedAt sets the "last_edited_at" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableLastEditedAt(t *time.Time) *IncomeManageUpdateOne {
	if t != nil {
		imuo.SetLastEditedAt(*t)
	}
	return imuo
}

// SetRejectReason sets the "reject_reason" field.
func (imuo *IncomeManageUpdateOne) SetRejectReason(s string) *IncomeManageUpdateOne {
	imuo.mutation.SetRejectReason(s)
	return imuo
}

// SetNillableRejectReason sets the "reject_reason" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableRejectReason(s *string) *IncomeManageUpdateOne {
	if s != nil {
		imuo.SetRejectReason(*s)
	}
	return imuo
}

// SetStatus sets the "status" field.
func (imuo *IncomeManageUpdateOne) SetStatus(ems enums.IncomeManageStatus) *IncomeManageUpdateOne {
	imuo.mutation.SetStatus(ems)
	return imuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableStatus(ems *enums.IncomeManageStatus) *IncomeManageUpdateOne {
	if ems != nil {
		imuo.SetStatus(*ems)
	}
	return imuo
}

// SetApproveUserID sets the "approve_user_id" field.
func (imuo *IncomeManageUpdateOne) SetApproveUserID(i int64) *IncomeManageUpdateOne {
	imuo.mutation.SetApproveUserID(i)
	return imuo
}

// SetNillableApproveUserID sets the "approve_user_id" field if the given value is not nil.
func (imuo *IncomeManageUpdateOne) SetNillableApproveUserID(i *int64) *IncomeManageUpdateOne {
	if i != nil {
		imuo.SetApproveUserID(*i)
	}
	return imuo
}

// SetUser sets the "user" edge to the User entity.
func (imuo *IncomeManageUpdateOne) SetUser(u *User) *IncomeManageUpdateOne {
	return imuo.SetUserID(u.ID)
}

// SetApproveUser sets the "approve_user" edge to the User entity.
func (imuo *IncomeManageUpdateOne) SetApproveUser(u *User) *IncomeManageUpdateOne {
	return imuo.SetApproveUserID(u.ID)
}

// Mutation returns the IncomeManageMutation object of the builder.
func (imuo *IncomeManageUpdateOne) Mutation() *IncomeManageMutation {
	return imuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (imuo *IncomeManageUpdateOne) ClearUser() *IncomeManageUpdateOne {
	imuo.mutation.ClearUser()
	return imuo
}

// ClearApproveUser clears the "approve_user" edge to the User entity.
func (imuo *IncomeManageUpdateOne) ClearApproveUser() *IncomeManageUpdateOne {
	imuo.mutation.ClearApproveUser()
	return imuo
}

// Where appends a list predicates to the IncomeManageUpdate builder.
func (imuo *IncomeManageUpdateOne) Where(ps ...predicate.IncomeManage) *IncomeManageUpdateOne {
	imuo.mutation.Where(ps...)
	return imuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (imuo *IncomeManageUpdateOne) Select(field string, fields ...string) *IncomeManageUpdateOne {
	imuo.fields = append([]string{field}, fields...)
	return imuo
}

// Save executes the query and returns the updated IncomeManage entity.
func (imuo *IncomeManageUpdateOne) Save(ctx context.Context) (*IncomeManage, error) {
	imuo.defaults()
	return withHooks(ctx, imuo.sqlSave, imuo.mutation, imuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (imuo *IncomeManageUpdateOne) SaveX(ctx context.Context) *IncomeManage {
	node, err := imuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (imuo *IncomeManageUpdateOne) Exec(ctx context.Context) error {
	_, err := imuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (imuo *IncomeManageUpdateOne) ExecX(ctx context.Context) {
	if err := imuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (imuo *IncomeManageUpdateOne) defaults() {
	if _, ok := imuo.mutation.UpdatedAt(); !ok {
		v := incomemanage.UpdateDefaultUpdatedAt()
		imuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (imuo *IncomeManageUpdateOne) check() error {
	if v, ok := imuo.mutation.GetType(); ok {
		if err := incomemanage.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "IncomeManage.type": %w`, err)}
		}
	}
	if v, ok := imuo.mutation.Status(); ok {
		if err := incomemanage.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "IncomeManage.status": %w`, err)}
		}
	}
	if _, ok := imuo.mutation.UserID(); imuo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "IncomeManage.user"`)
	}
	if _, ok := imuo.mutation.ApproveUserID(); imuo.mutation.ApproveUserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "IncomeManage.approve_user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (imuo *IncomeManageUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *IncomeManageUpdateOne {
	imuo.modifiers = append(imuo.modifiers, modifiers...)
	return imuo
}

func (imuo *IncomeManageUpdateOne) sqlSave(ctx context.Context) (_node *IncomeManage, err error) {
	if err := imuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(incomemanage.Table, incomemanage.Columns, sqlgraph.NewFieldSpec(incomemanage.FieldID, field.TypeInt64))
	id, ok := imuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "IncomeManage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := imuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, incomemanage.FieldID)
		for _, f := range fields {
			if !incomemanage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != incomemanage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := imuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := imuo.mutation.CreatedBy(); ok {
		_spec.SetField(incomemanage.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := imuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(incomemanage.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := imuo.mutation.UpdatedBy(); ok {
		_spec.SetField(incomemanage.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := imuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(incomemanage.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := imuo.mutation.UpdatedAt(); ok {
		_spec.SetField(incomemanage.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := imuo.mutation.DeletedAt(); ok {
		_spec.SetField(incomemanage.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := imuo.mutation.Phone(); ok {
		_spec.SetField(incomemanage.FieldPhone, field.TypeString, value)
	}
	if value, ok := imuo.mutation.GetType(); ok {
		_spec.SetField(incomemanage.FieldType, field.TypeEnum, value)
	}
	if value, ok := imuo.mutation.Amount(); ok {
		_spec.SetField(incomemanage.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := imuo.mutation.AddedAmount(); ok {
		_spec.AddField(incomemanage.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := imuo.mutation.Reason(); ok {
		_spec.SetField(incomemanage.FieldReason, field.TypeString, value)
	}
	if value, ok := imuo.mutation.CurrentBalance(); ok {
		_spec.SetField(incomemanage.FieldCurrentBalance, field.TypeInt64, value)
	}
	if value, ok := imuo.mutation.AddedCurrentBalance(); ok {
		_spec.AddField(incomemanage.FieldCurrentBalance, field.TypeInt64, value)
	}
	if value, ok := imuo.mutation.LastEditedAt(); ok {
		_spec.SetField(incomemanage.FieldLastEditedAt, field.TypeTime, value)
	}
	if value, ok := imuo.mutation.RejectReason(); ok {
		_spec.SetField(incomemanage.FieldRejectReason, field.TypeString, value)
	}
	if value, ok := imuo.mutation.Status(); ok {
		_spec.SetField(incomemanage.FieldStatus, field.TypeEnum, value)
	}
	if imuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incomemanage.UserTable,
			Columns: []string{incomemanage.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incomemanage.UserTable,
			Columns: []string{incomemanage.UserColumn},
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
	if imuo.mutation.ApproveUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incomemanage.ApproveUserTable,
			Columns: []string{incomemanage.ApproveUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := imuo.mutation.ApproveUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incomemanage.ApproveUserTable,
			Columns: []string{incomemanage.ApproveUserColumn},
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
	_spec.AddModifiers(imuo.modifiers...)
	_node = &IncomeManage{config: imuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, imuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{incomemanage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	imuo.mutation.done = true
	return _node, nil
}