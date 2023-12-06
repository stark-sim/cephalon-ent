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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/extraservice"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/extraserviceprice"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ExtraServicePriceUpdate is the builder for updating ExtraServicePrice entities.
type ExtraServicePriceUpdate struct {
	config
	hooks     []Hook
	mutation  *ExtraServicePriceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ExtraServicePriceUpdate builder.
func (espu *ExtraServicePriceUpdate) Where(ps ...predicate.ExtraServicePrice) *ExtraServicePriceUpdate {
	espu.mutation.Where(ps...)
	return espu
}

// SetCreatedBy sets the "created_by" field.
func (espu *ExtraServicePriceUpdate) SetCreatedBy(i int64) *ExtraServicePriceUpdate {
	espu.mutation.ResetCreatedBy()
	espu.mutation.SetCreatedBy(i)
	return espu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (espu *ExtraServicePriceUpdate) SetNillableCreatedBy(i *int64) *ExtraServicePriceUpdate {
	if i != nil {
		espu.SetCreatedBy(*i)
	}
	return espu
}

// AddCreatedBy adds i to the "created_by" field.
func (espu *ExtraServicePriceUpdate) AddCreatedBy(i int64) *ExtraServicePriceUpdate {
	espu.mutation.AddCreatedBy(i)
	return espu
}

// SetUpdatedBy sets the "updated_by" field.
func (espu *ExtraServicePriceUpdate) SetUpdatedBy(i int64) *ExtraServicePriceUpdate {
	espu.mutation.ResetUpdatedBy()
	espu.mutation.SetUpdatedBy(i)
	return espu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (espu *ExtraServicePriceUpdate) SetNillableUpdatedBy(i *int64) *ExtraServicePriceUpdate {
	if i != nil {
		espu.SetUpdatedBy(*i)
	}
	return espu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (espu *ExtraServicePriceUpdate) AddUpdatedBy(i int64) *ExtraServicePriceUpdate {
	espu.mutation.AddUpdatedBy(i)
	return espu
}

// SetUpdatedAt sets the "updated_at" field.
func (espu *ExtraServicePriceUpdate) SetUpdatedAt(t time.Time) *ExtraServicePriceUpdate {
	espu.mutation.SetUpdatedAt(t)
	return espu
}

// SetDeletedAt sets the "deleted_at" field.
func (espu *ExtraServicePriceUpdate) SetDeletedAt(t time.Time) *ExtraServicePriceUpdate {
	espu.mutation.SetDeletedAt(t)
	return espu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (espu *ExtraServicePriceUpdate) SetNillableDeletedAt(t *time.Time) *ExtraServicePriceUpdate {
	if t != nil {
		espu.SetDeletedAt(*t)
	}
	return espu
}

// SetExtraServiceType sets the "extra_service_type" field.
func (espu *ExtraServicePriceUpdate) SetExtraServiceType(est enums.ExtraServiceType) *ExtraServicePriceUpdate {
	espu.mutation.SetExtraServiceType(est)
	return espu
}

// SetNillableExtraServiceType sets the "extra_service_type" field if the given value is not nil.
func (espu *ExtraServicePriceUpdate) SetNillableExtraServiceType(est *enums.ExtraServiceType) *ExtraServicePriceUpdate {
	if est != nil {
		espu.SetExtraServiceType(*est)
	}
	return espu
}

// SetExtraServiceBillingType sets the "extra_service_billing_type" field.
func (espu *ExtraServicePriceUpdate) SetExtraServiceBillingType(esbt enums.ExtraServiceBillingType) *ExtraServicePriceUpdate {
	espu.mutation.SetExtraServiceBillingType(esbt)
	return espu
}

// SetNillableExtraServiceBillingType sets the "extra_service_billing_type" field if the given value is not nil.
func (espu *ExtraServicePriceUpdate) SetNillableExtraServiceBillingType(esbt *enums.ExtraServiceBillingType) *ExtraServicePriceUpdate {
	if esbt != nil {
		espu.SetExtraServiceBillingType(*esbt)
	}
	return espu
}

// SetExtraServiceID sets the "extra_service_id" field.
func (espu *ExtraServicePriceUpdate) SetExtraServiceID(i int64) *ExtraServicePriceUpdate {
	espu.mutation.SetExtraServiceID(i)
	return espu
}

// SetNillableExtraServiceID sets the "extra_service_id" field if the given value is not nil.
func (espu *ExtraServicePriceUpdate) SetNillableExtraServiceID(i *int64) *ExtraServicePriceUpdate {
	if i != nil {
		espu.SetExtraServiceID(*i)
	}
	return espu
}

// SetCep sets the "cep" field.
func (espu *ExtraServicePriceUpdate) SetCep(i int64) *ExtraServicePriceUpdate {
	espu.mutation.ResetCep()
	espu.mutation.SetCep(i)
	return espu
}

// SetNillableCep sets the "cep" field if the given value is not nil.
func (espu *ExtraServicePriceUpdate) SetNillableCep(i *int64) *ExtraServicePriceUpdate {
	if i != nil {
		espu.SetCep(*i)
	}
	return espu
}

// AddCep adds i to the "cep" field.
func (espu *ExtraServicePriceUpdate) AddCep(i int64) *ExtraServicePriceUpdate {
	espu.mutation.AddCep(i)
	return espu
}

// SetStartedAt sets the "started_at" field.
func (espu *ExtraServicePriceUpdate) SetStartedAt(t time.Time) *ExtraServicePriceUpdate {
	espu.mutation.SetStartedAt(t)
	return espu
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (espu *ExtraServicePriceUpdate) SetNillableStartedAt(t *time.Time) *ExtraServicePriceUpdate {
	if t != nil {
		espu.SetStartedAt(*t)
	}
	return espu
}

// ClearStartedAt clears the value of the "started_at" field.
func (espu *ExtraServicePriceUpdate) ClearStartedAt() *ExtraServicePriceUpdate {
	espu.mutation.ClearStartedAt()
	return espu
}

// SetFinishedAt sets the "finished_at" field.
func (espu *ExtraServicePriceUpdate) SetFinishedAt(t time.Time) *ExtraServicePriceUpdate {
	espu.mutation.SetFinishedAt(t)
	return espu
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (espu *ExtraServicePriceUpdate) SetNillableFinishedAt(t *time.Time) *ExtraServicePriceUpdate {
	if t != nil {
		espu.SetFinishedAt(*t)
	}
	return espu
}

// ClearFinishedAt clears the value of the "finished_at" field.
func (espu *ExtraServicePriceUpdate) ClearFinishedAt() *ExtraServicePriceUpdate {
	espu.mutation.ClearFinishedAt()
	return espu
}

// SetIsDeprecated sets the "is_deprecated" field.
func (espu *ExtraServicePriceUpdate) SetIsDeprecated(b bool) *ExtraServicePriceUpdate {
	espu.mutation.SetIsDeprecated(b)
	return espu
}

// SetNillableIsDeprecated sets the "is_deprecated" field if the given value is not nil.
func (espu *ExtraServicePriceUpdate) SetNillableIsDeprecated(b *bool) *ExtraServicePriceUpdate {
	if b != nil {
		espu.SetIsDeprecated(*b)
	}
	return espu
}

// SetIsSensitive sets the "is_sensitive" field.
func (espu *ExtraServicePriceUpdate) SetIsSensitive(b bool) *ExtraServicePriceUpdate {
	espu.mutation.SetIsSensitive(b)
	return espu
}

// SetNillableIsSensitive sets the "is_sensitive" field if the given value is not nil.
func (espu *ExtraServicePriceUpdate) SetNillableIsSensitive(b *bool) *ExtraServicePriceUpdate {
	if b != nil {
		espu.SetIsSensitive(*b)
	}
	return espu
}

// SetExtraService sets the "extra_service" edge to the ExtraService entity.
func (espu *ExtraServicePriceUpdate) SetExtraService(e *ExtraService) *ExtraServicePriceUpdate {
	return espu.SetExtraServiceID(e.ID)
}

// Mutation returns the ExtraServicePriceMutation object of the builder.
func (espu *ExtraServicePriceUpdate) Mutation() *ExtraServicePriceMutation {
	return espu.mutation
}

// ClearExtraService clears the "extra_service" edge to the ExtraService entity.
func (espu *ExtraServicePriceUpdate) ClearExtraService() *ExtraServicePriceUpdate {
	espu.mutation.ClearExtraService()
	return espu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (espu *ExtraServicePriceUpdate) Save(ctx context.Context) (int, error) {
	espu.defaults()
	return withHooks(ctx, espu.sqlSave, espu.mutation, espu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (espu *ExtraServicePriceUpdate) SaveX(ctx context.Context) int {
	affected, err := espu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (espu *ExtraServicePriceUpdate) Exec(ctx context.Context) error {
	_, err := espu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (espu *ExtraServicePriceUpdate) ExecX(ctx context.Context) {
	if err := espu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (espu *ExtraServicePriceUpdate) defaults() {
	if _, ok := espu.mutation.UpdatedAt(); !ok {
		v := extraserviceprice.UpdateDefaultUpdatedAt()
		espu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (espu *ExtraServicePriceUpdate) check() error {
	if v, ok := espu.mutation.ExtraServiceType(); ok {
		if err := extraserviceprice.ExtraServiceTypeValidator(v); err != nil {
			return &ValidationError{Name: "extra_service_type", err: fmt.Errorf(`cep_ent: validator failed for field "ExtraServicePrice.extra_service_type": %w`, err)}
		}
	}
	if v, ok := espu.mutation.ExtraServiceBillingType(); ok {
		if err := extraserviceprice.ExtraServiceBillingTypeValidator(v); err != nil {
			return &ValidationError{Name: "extra_service_billing_type", err: fmt.Errorf(`cep_ent: validator failed for field "ExtraServicePrice.extra_service_billing_type": %w`, err)}
		}
	}
	if _, ok := espu.mutation.ExtraServiceID(); espu.mutation.ExtraServiceCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ExtraServicePrice.extra_service"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (espu *ExtraServicePriceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ExtraServicePriceUpdate {
	espu.modifiers = append(espu.modifiers, modifiers...)
	return espu
}

func (espu *ExtraServicePriceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := espu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(extraserviceprice.Table, extraserviceprice.Columns, sqlgraph.NewFieldSpec(extraserviceprice.FieldID, field.TypeInt64))
	if ps := espu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := espu.mutation.CreatedBy(); ok {
		_spec.SetField(extraserviceprice.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := espu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(extraserviceprice.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := espu.mutation.UpdatedBy(); ok {
		_spec.SetField(extraserviceprice.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := espu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(extraserviceprice.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := espu.mutation.UpdatedAt(); ok {
		_spec.SetField(extraserviceprice.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := espu.mutation.DeletedAt(); ok {
		_spec.SetField(extraserviceprice.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := espu.mutation.ExtraServiceType(); ok {
		_spec.SetField(extraserviceprice.FieldExtraServiceType, field.TypeEnum, value)
	}
	if value, ok := espu.mutation.ExtraServiceBillingType(); ok {
		_spec.SetField(extraserviceprice.FieldExtraServiceBillingType, field.TypeEnum, value)
	}
	if value, ok := espu.mutation.Cep(); ok {
		_spec.SetField(extraserviceprice.FieldCep, field.TypeInt64, value)
	}
	if value, ok := espu.mutation.AddedCep(); ok {
		_spec.AddField(extraserviceprice.FieldCep, field.TypeInt64, value)
	}
	if value, ok := espu.mutation.StartedAt(); ok {
		_spec.SetField(extraserviceprice.FieldStartedAt, field.TypeTime, value)
	}
	if espu.mutation.StartedAtCleared() {
		_spec.ClearField(extraserviceprice.FieldStartedAt, field.TypeTime)
	}
	if value, ok := espu.mutation.FinishedAt(); ok {
		_spec.SetField(extraserviceprice.FieldFinishedAt, field.TypeTime, value)
	}
	if espu.mutation.FinishedAtCleared() {
		_spec.ClearField(extraserviceprice.FieldFinishedAt, field.TypeTime)
	}
	if value, ok := espu.mutation.IsDeprecated(); ok {
		_spec.SetField(extraserviceprice.FieldIsDeprecated, field.TypeBool, value)
	}
	if value, ok := espu.mutation.IsSensitive(); ok {
		_spec.SetField(extraserviceprice.FieldIsSensitive, field.TypeBool, value)
	}
	if espu.mutation.ExtraServiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceprice.ExtraServiceTable,
			Columns: []string{extraserviceprice.ExtraServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(extraservice.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := espu.mutation.ExtraServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceprice.ExtraServiceTable,
			Columns: []string{extraserviceprice.ExtraServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(extraservice.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(espu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, espu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{extraserviceprice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	espu.mutation.done = true
	return n, nil
}

// ExtraServicePriceUpdateOne is the builder for updating a single ExtraServicePrice entity.
type ExtraServicePriceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ExtraServicePriceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (espuo *ExtraServicePriceUpdateOne) SetCreatedBy(i int64) *ExtraServicePriceUpdateOne {
	espuo.mutation.ResetCreatedBy()
	espuo.mutation.SetCreatedBy(i)
	return espuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (espuo *ExtraServicePriceUpdateOne) SetNillableCreatedBy(i *int64) *ExtraServicePriceUpdateOne {
	if i != nil {
		espuo.SetCreatedBy(*i)
	}
	return espuo
}

// AddCreatedBy adds i to the "created_by" field.
func (espuo *ExtraServicePriceUpdateOne) AddCreatedBy(i int64) *ExtraServicePriceUpdateOne {
	espuo.mutation.AddCreatedBy(i)
	return espuo
}

// SetUpdatedBy sets the "updated_by" field.
func (espuo *ExtraServicePriceUpdateOne) SetUpdatedBy(i int64) *ExtraServicePriceUpdateOne {
	espuo.mutation.ResetUpdatedBy()
	espuo.mutation.SetUpdatedBy(i)
	return espuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (espuo *ExtraServicePriceUpdateOne) SetNillableUpdatedBy(i *int64) *ExtraServicePriceUpdateOne {
	if i != nil {
		espuo.SetUpdatedBy(*i)
	}
	return espuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (espuo *ExtraServicePriceUpdateOne) AddUpdatedBy(i int64) *ExtraServicePriceUpdateOne {
	espuo.mutation.AddUpdatedBy(i)
	return espuo
}

// SetUpdatedAt sets the "updated_at" field.
func (espuo *ExtraServicePriceUpdateOne) SetUpdatedAt(t time.Time) *ExtraServicePriceUpdateOne {
	espuo.mutation.SetUpdatedAt(t)
	return espuo
}

// SetDeletedAt sets the "deleted_at" field.
func (espuo *ExtraServicePriceUpdateOne) SetDeletedAt(t time.Time) *ExtraServicePriceUpdateOne {
	espuo.mutation.SetDeletedAt(t)
	return espuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (espuo *ExtraServicePriceUpdateOne) SetNillableDeletedAt(t *time.Time) *ExtraServicePriceUpdateOne {
	if t != nil {
		espuo.SetDeletedAt(*t)
	}
	return espuo
}

// SetExtraServiceType sets the "extra_service_type" field.
func (espuo *ExtraServicePriceUpdateOne) SetExtraServiceType(est enums.ExtraServiceType) *ExtraServicePriceUpdateOne {
	espuo.mutation.SetExtraServiceType(est)
	return espuo
}

// SetNillableExtraServiceType sets the "extra_service_type" field if the given value is not nil.
func (espuo *ExtraServicePriceUpdateOne) SetNillableExtraServiceType(est *enums.ExtraServiceType) *ExtraServicePriceUpdateOne {
	if est != nil {
		espuo.SetExtraServiceType(*est)
	}
	return espuo
}

// SetExtraServiceBillingType sets the "extra_service_billing_type" field.
func (espuo *ExtraServicePriceUpdateOne) SetExtraServiceBillingType(esbt enums.ExtraServiceBillingType) *ExtraServicePriceUpdateOne {
	espuo.mutation.SetExtraServiceBillingType(esbt)
	return espuo
}

// SetNillableExtraServiceBillingType sets the "extra_service_billing_type" field if the given value is not nil.
func (espuo *ExtraServicePriceUpdateOne) SetNillableExtraServiceBillingType(esbt *enums.ExtraServiceBillingType) *ExtraServicePriceUpdateOne {
	if esbt != nil {
		espuo.SetExtraServiceBillingType(*esbt)
	}
	return espuo
}

// SetExtraServiceID sets the "extra_service_id" field.
func (espuo *ExtraServicePriceUpdateOne) SetExtraServiceID(i int64) *ExtraServicePriceUpdateOne {
	espuo.mutation.SetExtraServiceID(i)
	return espuo
}

// SetNillableExtraServiceID sets the "extra_service_id" field if the given value is not nil.
func (espuo *ExtraServicePriceUpdateOne) SetNillableExtraServiceID(i *int64) *ExtraServicePriceUpdateOne {
	if i != nil {
		espuo.SetExtraServiceID(*i)
	}
	return espuo
}

// SetCep sets the "cep" field.
func (espuo *ExtraServicePriceUpdateOne) SetCep(i int64) *ExtraServicePriceUpdateOne {
	espuo.mutation.ResetCep()
	espuo.mutation.SetCep(i)
	return espuo
}

// SetNillableCep sets the "cep" field if the given value is not nil.
func (espuo *ExtraServicePriceUpdateOne) SetNillableCep(i *int64) *ExtraServicePriceUpdateOne {
	if i != nil {
		espuo.SetCep(*i)
	}
	return espuo
}

// AddCep adds i to the "cep" field.
func (espuo *ExtraServicePriceUpdateOne) AddCep(i int64) *ExtraServicePriceUpdateOne {
	espuo.mutation.AddCep(i)
	return espuo
}

// SetStartedAt sets the "started_at" field.
func (espuo *ExtraServicePriceUpdateOne) SetStartedAt(t time.Time) *ExtraServicePriceUpdateOne {
	espuo.mutation.SetStartedAt(t)
	return espuo
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (espuo *ExtraServicePriceUpdateOne) SetNillableStartedAt(t *time.Time) *ExtraServicePriceUpdateOne {
	if t != nil {
		espuo.SetStartedAt(*t)
	}
	return espuo
}

// ClearStartedAt clears the value of the "started_at" field.
func (espuo *ExtraServicePriceUpdateOne) ClearStartedAt() *ExtraServicePriceUpdateOne {
	espuo.mutation.ClearStartedAt()
	return espuo
}

// SetFinishedAt sets the "finished_at" field.
func (espuo *ExtraServicePriceUpdateOne) SetFinishedAt(t time.Time) *ExtraServicePriceUpdateOne {
	espuo.mutation.SetFinishedAt(t)
	return espuo
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (espuo *ExtraServicePriceUpdateOne) SetNillableFinishedAt(t *time.Time) *ExtraServicePriceUpdateOne {
	if t != nil {
		espuo.SetFinishedAt(*t)
	}
	return espuo
}

// ClearFinishedAt clears the value of the "finished_at" field.
func (espuo *ExtraServicePriceUpdateOne) ClearFinishedAt() *ExtraServicePriceUpdateOne {
	espuo.mutation.ClearFinishedAt()
	return espuo
}

// SetIsDeprecated sets the "is_deprecated" field.
func (espuo *ExtraServicePriceUpdateOne) SetIsDeprecated(b bool) *ExtraServicePriceUpdateOne {
	espuo.mutation.SetIsDeprecated(b)
	return espuo
}

// SetNillableIsDeprecated sets the "is_deprecated" field if the given value is not nil.
func (espuo *ExtraServicePriceUpdateOne) SetNillableIsDeprecated(b *bool) *ExtraServicePriceUpdateOne {
	if b != nil {
		espuo.SetIsDeprecated(*b)
	}
	return espuo
}

// SetIsSensitive sets the "is_sensitive" field.
func (espuo *ExtraServicePriceUpdateOne) SetIsSensitive(b bool) *ExtraServicePriceUpdateOne {
	espuo.mutation.SetIsSensitive(b)
	return espuo
}

// SetNillableIsSensitive sets the "is_sensitive" field if the given value is not nil.
func (espuo *ExtraServicePriceUpdateOne) SetNillableIsSensitive(b *bool) *ExtraServicePriceUpdateOne {
	if b != nil {
		espuo.SetIsSensitive(*b)
	}
	return espuo
}

// SetExtraService sets the "extra_service" edge to the ExtraService entity.
func (espuo *ExtraServicePriceUpdateOne) SetExtraService(e *ExtraService) *ExtraServicePriceUpdateOne {
	return espuo.SetExtraServiceID(e.ID)
}

// Mutation returns the ExtraServicePriceMutation object of the builder.
func (espuo *ExtraServicePriceUpdateOne) Mutation() *ExtraServicePriceMutation {
	return espuo.mutation
}

// ClearExtraService clears the "extra_service" edge to the ExtraService entity.
func (espuo *ExtraServicePriceUpdateOne) ClearExtraService() *ExtraServicePriceUpdateOne {
	espuo.mutation.ClearExtraService()
	return espuo
}

// Where appends a list predicates to the ExtraServicePriceUpdate builder.
func (espuo *ExtraServicePriceUpdateOne) Where(ps ...predicate.ExtraServicePrice) *ExtraServicePriceUpdateOne {
	espuo.mutation.Where(ps...)
	return espuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (espuo *ExtraServicePriceUpdateOne) Select(field string, fields ...string) *ExtraServicePriceUpdateOne {
	espuo.fields = append([]string{field}, fields...)
	return espuo
}

// Save executes the query and returns the updated ExtraServicePrice entity.
func (espuo *ExtraServicePriceUpdateOne) Save(ctx context.Context) (*ExtraServicePrice, error) {
	espuo.defaults()
	return withHooks(ctx, espuo.sqlSave, espuo.mutation, espuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (espuo *ExtraServicePriceUpdateOne) SaveX(ctx context.Context) *ExtraServicePrice {
	node, err := espuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (espuo *ExtraServicePriceUpdateOne) Exec(ctx context.Context) error {
	_, err := espuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (espuo *ExtraServicePriceUpdateOne) ExecX(ctx context.Context) {
	if err := espuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (espuo *ExtraServicePriceUpdateOne) defaults() {
	if _, ok := espuo.mutation.UpdatedAt(); !ok {
		v := extraserviceprice.UpdateDefaultUpdatedAt()
		espuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (espuo *ExtraServicePriceUpdateOne) check() error {
	if v, ok := espuo.mutation.ExtraServiceType(); ok {
		if err := extraserviceprice.ExtraServiceTypeValidator(v); err != nil {
			return &ValidationError{Name: "extra_service_type", err: fmt.Errorf(`cep_ent: validator failed for field "ExtraServicePrice.extra_service_type": %w`, err)}
		}
	}
	if v, ok := espuo.mutation.ExtraServiceBillingType(); ok {
		if err := extraserviceprice.ExtraServiceBillingTypeValidator(v); err != nil {
			return &ValidationError{Name: "extra_service_billing_type", err: fmt.Errorf(`cep_ent: validator failed for field "ExtraServicePrice.extra_service_billing_type": %w`, err)}
		}
	}
	if _, ok := espuo.mutation.ExtraServiceID(); espuo.mutation.ExtraServiceCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ExtraServicePrice.extra_service"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (espuo *ExtraServicePriceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ExtraServicePriceUpdateOne {
	espuo.modifiers = append(espuo.modifiers, modifiers...)
	return espuo
}

func (espuo *ExtraServicePriceUpdateOne) sqlSave(ctx context.Context) (_node *ExtraServicePrice, err error) {
	if err := espuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(extraserviceprice.Table, extraserviceprice.Columns, sqlgraph.NewFieldSpec(extraserviceprice.FieldID, field.TypeInt64))
	id, ok := espuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "ExtraServicePrice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := espuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, extraserviceprice.FieldID)
		for _, f := range fields {
			if !extraserviceprice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != extraserviceprice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := espuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := espuo.mutation.CreatedBy(); ok {
		_spec.SetField(extraserviceprice.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := espuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(extraserviceprice.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := espuo.mutation.UpdatedBy(); ok {
		_spec.SetField(extraserviceprice.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := espuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(extraserviceprice.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := espuo.mutation.UpdatedAt(); ok {
		_spec.SetField(extraserviceprice.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := espuo.mutation.DeletedAt(); ok {
		_spec.SetField(extraserviceprice.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := espuo.mutation.ExtraServiceType(); ok {
		_spec.SetField(extraserviceprice.FieldExtraServiceType, field.TypeEnum, value)
	}
	if value, ok := espuo.mutation.ExtraServiceBillingType(); ok {
		_spec.SetField(extraserviceprice.FieldExtraServiceBillingType, field.TypeEnum, value)
	}
	if value, ok := espuo.mutation.Cep(); ok {
		_spec.SetField(extraserviceprice.FieldCep, field.TypeInt64, value)
	}
	if value, ok := espuo.mutation.AddedCep(); ok {
		_spec.AddField(extraserviceprice.FieldCep, field.TypeInt64, value)
	}
	if value, ok := espuo.mutation.StartedAt(); ok {
		_spec.SetField(extraserviceprice.FieldStartedAt, field.TypeTime, value)
	}
	if espuo.mutation.StartedAtCleared() {
		_spec.ClearField(extraserviceprice.FieldStartedAt, field.TypeTime)
	}
	if value, ok := espuo.mutation.FinishedAt(); ok {
		_spec.SetField(extraserviceprice.FieldFinishedAt, field.TypeTime, value)
	}
	if espuo.mutation.FinishedAtCleared() {
		_spec.ClearField(extraserviceprice.FieldFinishedAt, field.TypeTime)
	}
	if value, ok := espuo.mutation.IsDeprecated(); ok {
		_spec.SetField(extraserviceprice.FieldIsDeprecated, field.TypeBool, value)
	}
	if value, ok := espuo.mutation.IsSensitive(); ok {
		_spec.SetField(extraserviceprice.FieldIsSensitive, field.TypeBool, value)
	}
	if espuo.mutation.ExtraServiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceprice.ExtraServiceTable,
			Columns: []string{extraserviceprice.ExtraServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(extraservice.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := espuo.mutation.ExtraServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceprice.ExtraServiceTable,
			Columns: []string{extraserviceprice.ExtraServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(extraservice.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(espuo.modifiers...)
	_node = &ExtraServicePrice{config: espuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, espuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{extraserviceprice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	espuo.mutation.done = true
	return _node, nil
}
