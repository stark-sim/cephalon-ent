// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/outputlog"
	"cephalon-ent/pkg/cep_ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OutputLogUpdate is the builder for updating OutputLog entities.
type OutputLogUpdate struct {
	config
	hooks    []Hook
	mutation *OutputLogMutation
}

// Where appends a list predicates to the OutputLogUpdate builder.
func (olu *OutputLogUpdate) Where(ps ...predicate.OutputLog) *OutputLogUpdate {
	olu.mutation.Where(ps...)
	return olu
}

// SetCreatedBy sets the "created_by" field.
func (olu *OutputLogUpdate) SetCreatedBy(i int64) *OutputLogUpdate {
	olu.mutation.ResetCreatedBy()
	olu.mutation.SetCreatedBy(i)
	return olu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (olu *OutputLogUpdate) SetNillableCreatedBy(i *int64) *OutputLogUpdate {
	if i != nil {
		olu.SetCreatedBy(*i)
	}
	return olu
}

// AddCreatedBy adds i to the "created_by" field.
func (olu *OutputLogUpdate) AddCreatedBy(i int64) *OutputLogUpdate {
	olu.mutation.AddCreatedBy(i)
	return olu
}

// SetUpdatedBy sets the "updated_by" field.
func (olu *OutputLogUpdate) SetUpdatedBy(i int64) *OutputLogUpdate {
	olu.mutation.ResetUpdatedBy()
	olu.mutation.SetUpdatedBy(i)
	return olu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (olu *OutputLogUpdate) SetNillableUpdatedBy(i *int64) *OutputLogUpdate {
	if i != nil {
		olu.SetUpdatedBy(*i)
	}
	return olu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (olu *OutputLogUpdate) AddUpdatedBy(i int64) *OutputLogUpdate {
	olu.mutation.AddUpdatedBy(i)
	return olu
}

// SetUpdatedAt sets the "updated_at" field.
func (olu *OutputLogUpdate) SetUpdatedAt(t time.Time) *OutputLogUpdate {
	olu.mutation.SetUpdatedAt(t)
	return olu
}

// SetDeletedAt sets the "deleted_at" field.
func (olu *OutputLogUpdate) SetDeletedAt(t time.Time) *OutputLogUpdate {
	olu.mutation.SetDeletedAt(t)
	return olu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (olu *OutputLogUpdate) SetNillableDeletedAt(t *time.Time) *OutputLogUpdate {
	if t != nil {
		olu.SetDeletedAt(*t)
	}
	return olu
}

// SetHeaders sets the "headers" field.
func (olu *OutputLogUpdate) SetHeaders(s string) *OutputLogUpdate {
	olu.mutation.SetHeaders(s)
	return olu
}

// SetBody sets the "body" field.
func (olu *OutputLogUpdate) SetBody(s string) *OutputLogUpdate {
	olu.mutation.SetBody(s)
	return olu
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (olu *OutputLogUpdate) SetNillableBody(s *string) *OutputLogUpdate {
	if s != nil {
		olu.SetBody(*s)
	}
	return olu
}

// ClearBody clears the value of the "body" field.
func (olu *OutputLogUpdate) ClearBody() *OutputLogUpdate {
	olu.mutation.ClearBody()
	return olu
}

// SetURL sets the "url" field.
func (olu *OutputLogUpdate) SetURL(s string) *OutputLogUpdate {
	olu.mutation.SetURL(s)
	return olu
}

// SetIP sets the "ip" field.
func (olu *OutputLogUpdate) SetIP(s string) *OutputLogUpdate {
	olu.mutation.SetIP(s)
	return olu
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (olu *OutputLogUpdate) SetNillableIP(s *string) *OutputLogUpdate {
	if s != nil {
		olu.SetIP(*s)
	}
	return olu
}

// ClearIP clears the value of the "ip" field.
func (olu *OutputLogUpdate) ClearIP() *OutputLogUpdate {
	olu.mutation.ClearIP()
	return olu
}

// SetCaller sets the "caller" field.
func (olu *OutputLogUpdate) SetCaller(s string) *OutputLogUpdate {
	olu.mutation.SetCaller(s)
	return olu
}

// SetNillableCaller sets the "caller" field if the given value is not nil.
func (olu *OutputLogUpdate) SetNillableCaller(s *string) *OutputLogUpdate {
	if s != nil {
		olu.SetCaller(*s)
	}
	return olu
}

// SetStatus sets the "status" field.
func (olu *OutputLogUpdate) SetStatus(i int16) *OutputLogUpdate {
	olu.mutation.ResetStatus()
	olu.mutation.SetStatus(i)
	return olu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (olu *OutputLogUpdate) SetNillableStatus(i *int16) *OutputLogUpdate {
	if i != nil {
		olu.SetStatus(*i)
	}
	return olu
}

// AddStatus adds i to the "status" field.
func (olu *OutputLogUpdate) AddStatus(i int16) *OutputLogUpdate {
	olu.mutation.AddStatus(i)
	return olu
}

// ClearStatus clears the value of the "status" field.
func (olu *OutputLogUpdate) ClearStatus() *OutputLogUpdate {
	olu.mutation.ClearStatus()
	return olu
}

// SetHmacKey sets the "hmac_key" field.
func (olu *OutputLogUpdate) SetHmacKey(s string) *OutputLogUpdate {
	olu.mutation.SetHmacKey(s)
	return olu
}

// SetNillableHmacKey sets the "hmac_key" field if the given value is not nil.
func (olu *OutputLogUpdate) SetNillableHmacKey(s *string) *OutputLogUpdate {
	if s != nil {
		olu.SetHmacKey(*s)
	}
	return olu
}

// Mutation returns the OutputLogMutation object of the builder.
func (olu *OutputLogUpdate) Mutation() *OutputLogMutation {
	return olu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (olu *OutputLogUpdate) Save(ctx context.Context) (int, error) {
	olu.defaults()
	return withHooks(ctx, olu.sqlSave, olu.mutation, olu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (olu *OutputLogUpdate) SaveX(ctx context.Context) int {
	affected, err := olu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (olu *OutputLogUpdate) Exec(ctx context.Context) error {
	_, err := olu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olu *OutputLogUpdate) ExecX(ctx context.Context) {
	if err := olu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (olu *OutputLogUpdate) defaults() {
	if _, ok := olu.mutation.UpdatedAt(); !ok {
		v := outputlog.UpdateDefaultUpdatedAt()
		olu.mutation.SetUpdatedAt(v)
	}
}

func (olu *OutputLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(outputlog.Table, outputlog.Columns, sqlgraph.NewFieldSpec(outputlog.FieldID, field.TypeInt64))
	if ps := olu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := olu.mutation.CreatedBy(); ok {
		_spec.SetField(outputlog.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := olu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(outputlog.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := olu.mutation.UpdatedBy(); ok {
		_spec.SetField(outputlog.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := olu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(outputlog.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := olu.mutation.UpdatedAt(); ok {
		_spec.SetField(outputlog.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := olu.mutation.DeletedAt(); ok {
		_spec.SetField(outputlog.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := olu.mutation.Headers(); ok {
		_spec.SetField(outputlog.FieldHeaders, field.TypeString, value)
	}
	if value, ok := olu.mutation.Body(); ok {
		_spec.SetField(outputlog.FieldBody, field.TypeString, value)
	}
	if olu.mutation.BodyCleared() {
		_spec.ClearField(outputlog.FieldBody, field.TypeString)
	}
	if value, ok := olu.mutation.URL(); ok {
		_spec.SetField(outputlog.FieldURL, field.TypeString, value)
	}
	if value, ok := olu.mutation.IP(); ok {
		_spec.SetField(outputlog.FieldIP, field.TypeString, value)
	}
	if olu.mutation.IPCleared() {
		_spec.ClearField(outputlog.FieldIP, field.TypeString)
	}
	if value, ok := olu.mutation.Caller(); ok {
		_spec.SetField(outputlog.FieldCaller, field.TypeString, value)
	}
	if value, ok := olu.mutation.Status(); ok {
		_spec.SetField(outputlog.FieldStatus, field.TypeInt16, value)
	}
	if value, ok := olu.mutation.AddedStatus(); ok {
		_spec.AddField(outputlog.FieldStatus, field.TypeInt16, value)
	}
	if olu.mutation.StatusCleared() {
		_spec.ClearField(outputlog.FieldStatus, field.TypeInt16)
	}
	if value, ok := olu.mutation.HmacKey(); ok {
		_spec.SetField(outputlog.FieldHmacKey, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, olu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outputlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	olu.mutation.done = true
	return n, nil
}

// OutputLogUpdateOne is the builder for updating a single OutputLog entity.
type OutputLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutputLogMutation
}

// SetCreatedBy sets the "created_by" field.
func (oluo *OutputLogUpdateOne) SetCreatedBy(i int64) *OutputLogUpdateOne {
	oluo.mutation.ResetCreatedBy()
	oluo.mutation.SetCreatedBy(i)
	return oluo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (oluo *OutputLogUpdateOne) SetNillableCreatedBy(i *int64) *OutputLogUpdateOne {
	if i != nil {
		oluo.SetCreatedBy(*i)
	}
	return oluo
}

// AddCreatedBy adds i to the "created_by" field.
func (oluo *OutputLogUpdateOne) AddCreatedBy(i int64) *OutputLogUpdateOne {
	oluo.mutation.AddCreatedBy(i)
	return oluo
}

// SetUpdatedBy sets the "updated_by" field.
func (oluo *OutputLogUpdateOne) SetUpdatedBy(i int64) *OutputLogUpdateOne {
	oluo.mutation.ResetUpdatedBy()
	oluo.mutation.SetUpdatedBy(i)
	return oluo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oluo *OutputLogUpdateOne) SetNillableUpdatedBy(i *int64) *OutputLogUpdateOne {
	if i != nil {
		oluo.SetUpdatedBy(*i)
	}
	return oluo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (oluo *OutputLogUpdateOne) AddUpdatedBy(i int64) *OutputLogUpdateOne {
	oluo.mutation.AddUpdatedBy(i)
	return oluo
}

// SetUpdatedAt sets the "updated_at" field.
func (oluo *OutputLogUpdateOne) SetUpdatedAt(t time.Time) *OutputLogUpdateOne {
	oluo.mutation.SetUpdatedAt(t)
	return oluo
}

// SetDeletedAt sets the "deleted_at" field.
func (oluo *OutputLogUpdateOne) SetDeletedAt(t time.Time) *OutputLogUpdateOne {
	oluo.mutation.SetDeletedAt(t)
	return oluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oluo *OutputLogUpdateOne) SetNillableDeletedAt(t *time.Time) *OutputLogUpdateOne {
	if t != nil {
		oluo.SetDeletedAt(*t)
	}
	return oluo
}

// SetHeaders sets the "headers" field.
func (oluo *OutputLogUpdateOne) SetHeaders(s string) *OutputLogUpdateOne {
	oluo.mutation.SetHeaders(s)
	return oluo
}

// SetBody sets the "body" field.
func (oluo *OutputLogUpdateOne) SetBody(s string) *OutputLogUpdateOne {
	oluo.mutation.SetBody(s)
	return oluo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (oluo *OutputLogUpdateOne) SetNillableBody(s *string) *OutputLogUpdateOne {
	if s != nil {
		oluo.SetBody(*s)
	}
	return oluo
}

// ClearBody clears the value of the "body" field.
func (oluo *OutputLogUpdateOne) ClearBody() *OutputLogUpdateOne {
	oluo.mutation.ClearBody()
	return oluo
}

// SetURL sets the "url" field.
func (oluo *OutputLogUpdateOne) SetURL(s string) *OutputLogUpdateOne {
	oluo.mutation.SetURL(s)
	return oluo
}

// SetIP sets the "ip" field.
func (oluo *OutputLogUpdateOne) SetIP(s string) *OutputLogUpdateOne {
	oluo.mutation.SetIP(s)
	return oluo
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (oluo *OutputLogUpdateOne) SetNillableIP(s *string) *OutputLogUpdateOne {
	if s != nil {
		oluo.SetIP(*s)
	}
	return oluo
}

// ClearIP clears the value of the "ip" field.
func (oluo *OutputLogUpdateOne) ClearIP() *OutputLogUpdateOne {
	oluo.mutation.ClearIP()
	return oluo
}

// SetCaller sets the "caller" field.
func (oluo *OutputLogUpdateOne) SetCaller(s string) *OutputLogUpdateOne {
	oluo.mutation.SetCaller(s)
	return oluo
}

// SetNillableCaller sets the "caller" field if the given value is not nil.
func (oluo *OutputLogUpdateOne) SetNillableCaller(s *string) *OutputLogUpdateOne {
	if s != nil {
		oluo.SetCaller(*s)
	}
	return oluo
}

// SetStatus sets the "status" field.
func (oluo *OutputLogUpdateOne) SetStatus(i int16) *OutputLogUpdateOne {
	oluo.mutation.ResetStatus()
	oluo.mutation.SetStatus(i)
	return oluo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oluo *OutputLogUpdateOne) SetNillableStatus(i *int16) *OutputLogUpdateOne {
	if i != nil {
		oluo.SetStatus(*i)
	}
	return oluo
}

// AddStatus adds i to the "status" field.
func (oluo *OutputLogUpdateOne) AddStatus(i int16) *OutputLogUpdateOne {
	oluo.mutation.AddStatus(i)
	return oluo
}

// ClearStatus clears the value of the "status" field.
func (oluo *OutputLogUpdateOne) ClearStatus() *OutputLogUpdateOne {
	oluo.mutation.ClearStatus()
	return oluo
}

// SetHmacKey sets the "hmac_key" field.
func (oluo *OutputLogUpdateOne) SetHmacKey(s string) *OutputLogUpdateOne {
	oluo.mutation.SetHmacKey(s)
	return oluo
}

// SetNillableHmacKey sets the "hmac_key" field if the given value is not nil.
func (oluo *OutputLogUpdateOne) SetNillableHmacKey(s *string) *OutputLogUpdateOne {
	if s != nil {
		oluo.SetHmacKey(*s)
	}
	return oluo
}

// Mutation returns the OutputLogMutation object of the builder.
func (oluo *OutputLogUpdateOne) Mutation() *OutputLogMutation {
	return oluo.mutation
}

// Where appends a list predicates to the OutputLogUpdate builder.
func (oluo *OutputLogUpdateOne) Where(ps ...predicate.OutputLog) *OutputLogUpdateOne {
	oluo.mutation.Where(ps...)
	return oluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oluo *OutputLogUpdateOne) Select(field string, fields ...string) *OutputLogUpdateOne {
	oluo.fields = append([]string{field}, fields...)
	return oluo
}

// Save executes the query and returns the updated OutputLog entity.
func (oluo *OutputLogUpdateOne) Save(ctx context.Context) (*OutputLog, error) {
	oluo.defaults()
	return withHooks(ctx, oluo.sqlSave, oluo.mutation, oluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oluo *OutputLogUpdateOne) SaveX(ctx context.Context) *OutputLog {
	node, err := oluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oluo *OutputLogUpdateOne) Exec(ctx context.Context) error {
	_, err := oluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oluo *OutputLogUpdateOne) ExecX(ctx context.Context) {
	if err := oluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oluo *OutputLogUpdateOne) defaults() {
	if _, ok := oluo.mutation.UpdatedAt(); !ok {
		v := outputlog.UpdateDefaultUpdatedAt()
		oluo.mutation.SetUpdatedAt(v)
	}
}

func (oluo *OutputLogUpdateOne) sqlSave(ctx context.Context) (_node *OutputLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(outputlog.Table, outputlog.Columns, sqlgraph.NewFieldSpec(outputlog.FieldID, field.TypeInt64))
	id, ok := oluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "OutputLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outputlog.FieldID)
		for _, f := range fields {
			if !outputlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != outputlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oluo.mutation.CreatedBy(); ok {
		_spec.SetField(outputlog.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := oluo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(outputlog.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := oluo.mutation.UpdatedBy(); ok {
		_spec.SetField(outputlog.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := oluo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(outputlog.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := oluo.mutation.UpdatedAt(); ok {
		_spec.SetField(outputlog.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := oluo.mutation.DeletedAt(); ok {
		_spec.SetField(outputlog.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := oluo.mutation.Headers(); ok {
		_spec.SetField(outputlog.FieldHeaders, field.TypeString, value)
	}
	if value, ok := oluo.mutation.Body(); ok {
		_spec.SetField(outputlog.FieldBody, field.TypeString, value)
	}
	if oluo.mutation.BodyCleared() {
		_spec.ClearField(outputlog.FieldBody, field.TypeString)
	}
	if value, ok := oluo.mutation.URL(); ok {
		_spec.SetField(outputlog.FieldURL, field.TypeString, value)
	}
	if value, ok := oluo.mutation.IP(); ok {
		_spec.SetField(outputlog.FieldIP, field.TypeString, value)
	}
	if oluo.mutation.IPCleared() {
		_spec.ClearField(outputlog.FieldIP, field.TypeString)
	}
	if value, ok := oluo.mutation.Caller(); ok {
		_spec.SetField(outputlog.FieldCaller, field.TypeString, value)
	}
	if value, ok := oluo.mutation.Status(); ok {
		_spec.SetField(outputlog.FieldStatus, field.TypeInt16, value)
	}
	if value, ok := oluo.mutation.AddedStatus(); ok {
		_spec.AddField(outputlog.FieldStatus, field.TypeInt16, value)
	}
	if oluo.mutation.StatusCleared() {
		_spec.ClearField(outputlog.FieldStatus, field.TypeInt16)
	}
	if value, ok := oluo.mutation.HmacKey(); ok {
		_spec.SetField(outputlog.FieldHmacKey, field.TypeString, value)
	}
	_node = &OutputLog{config: oluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outputlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oluo.mutation.done = true
	return _node, nil
}