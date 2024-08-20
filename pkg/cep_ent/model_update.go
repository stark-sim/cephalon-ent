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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/invokemodelorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/model"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/modelprice"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/usermodel"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ModelUpdate is the builder for updating Model entities.
type ModelUpdate struct {
	config
	hooks     []Hook
	mutation  *ModelMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ModelUpdate builder.
func (mu *ModelUpdate) Where(ps ...predicate.Model) *ModelUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCreatedBy sets the "created_by" field.
func (mu *ModelUpdate) SetCreatedBy(i int64) *ModelUpdate {
	mu.mutation.ResetCreatedBy()
	mu.mutation.SetCreatedBy(i)
	return mu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableCreatedBy(i *int64) *ModelUpdate {
	if i != nil {
		mu.SetCreatedBy(*i)
	}
	return mu
}

// AddCreatedBy adds i to the "created_by" field.
func (mu *ModelUpdate) AddCreatedBy(i int64) *ModelUpdate {
	mu.mutation.AddCreatedBy(i)
	return mu
}

// SetUpdatedBy sets the "updated_by" field.
func (mu *ModelUpdate) SetUpdatedBy(i int64) *ModelUpdate {
	mu.mutation.ResetUpdatedBy()
	mu.mutation.SetUpdatedBy(i)
	return mu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableUpdatedBy(i *int64) *ModelUpdate {
	if i != nil {
		mu.SetUpdatedBy(*i)
	}
	return mu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (mu *ModelUpdate) AddUpdatedBy(i int64) *ModelUpdate {
	mu.mutation.AddUpdatedBy(i)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *ModelUpdate) SetUpdatedAt(t time.Time) *ModelUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetDeletedAt sets the "deleted_at" field.
func (mu *ModelUpdate) SetDeletedAt(t time.Time) *ModelUpdate {
	mu.mutation.SetDeletedAt(t)
	return mu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableDeletedAt(t *time.Time) *ModelUpdate {
	if t != nil {
		mu.SetDeletedAt(*t)
	}
	return mu
}

// SetName sets the "name" field.
func (mu *ModelUpdate) SetName(s string) *ModelUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableName(s *string) *ModelUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// SetAuthor sets the "author" field.
func (mu *ModelUpdate) SetAuthor(s string) *ModelUpdate {
	mu.mutation.SetAuthor(s)
	return mu
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableAuthor(s *string) *ModelUpdate {
	if s != nil {
		mu.SetAuthor(*s)
	}
	return mu
}

// SetDescription sets the "description" field.
func (mu *ModelUpdate) SetDescription(s string) *ModelUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableDescription(s *string) *ModelUpdate {
	if s != nil {
		mu.SetDescription(*s)
	}
	return mu
}

// SetModelType sets the "model_type" field.
func (mu *ModelUpdate) SetModelType(e enums.Model) *ModelUpdate {
	mu.mutation.SetModelType(e)
	return mu
}

// SetNillableModelType sets the "model_type" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableModelType(e *enums.Model) *ModelUpdate {
	if e != nil {
		mu.SetModelType(*e)
	}
	return mu
}

// SetModelStatus sets the "model_status" field.
func (mu *ModelUpdate) SetModelStatus(es enums.ModelStatus) *ModelUpdate {
	mu.mutation.SetModelStatus(es)
	return mu
}

// SetNillableModelStatus sets the "model_status" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableModelStatus(es *enums.ModelStatus) *ModelUpdate {
	if es != nil {
		mu.SetModelStatus(*es)
	}
	return mu
}

// SetIsOfficial sets the "is_official" field.
func (mu *ModelUpdate) SetIsOfficial(b bool) *ModelUpdate {
	mu.mutation.SetIsOfficial(b)
	return mu
}

// SetNillableIsOfficial sets the "is_official" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableIsOfficial(b *bool) *ModelUpdate {
	if b != nil {
		mu.SetIsOfficial(*b)
	}
	return mu
}

// SetTotalUsage sets the "total_usage" field.
func (mu *ModelUpdate) SetTotalUsage(i int) *ModelUpdate {
	mu.mutation.ResetTotalUsage()
	mu.mutation.SetTotalUsage(i)
	return mu
}

// SetNillableTotalUsage sets the "total_usage" field if the given value is not nil.
func (mu *ModelUpdate) SetNillableTotalUsage(i *int) *ModelUpdate {
	if i != nil {
		mu.SetTotalUsage(*i)
	}
	return mu
}

// AddTotalUsage adds i to the "total_usage" field.
func (mu *ModelUpdate) AddTotalUsage(i int) *ModelUpdate {
	mu.mutation.AddTotalUsage(i)
	return mu
}

// AddModelPriceIDs adds the "model_prices" edge to the ModelPrice entity by IDs.
func (mu *ModelUpdate) AddModelPriceIDs(ids ...int64) *ModelUpdate {
	mu.mutation.AddModelPriceIDs(ids...)
	return mu
}

// AddModelPrices adds the "model_prices" edges to the ModelPrice entity.
func (mu *ModelUpdate) AddModelPrices(m ...*ModelPrice) *ModelUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddModelPriceIDs(ids...)
}

// AddUserModelIDs adds the "user_models" edge to the UserModel entity by IDs.
func (mu *ModelUpdate) AddUserModelIDs(ids ...int64) *ModelUpdate {
	mu.mutation.AddUserModelIDs(ids...)
	return mu
}

// AddUserModels adds the "user_models" edges to the UserModel entity.
func (mu *ModelUpdate) AddUserModels(u ...*UserModel) *ModelUpdate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mu.AddUserModelIDs(ids...)
}

// AddInvokeModelOrderIDs adds the "invoke_model_orders" edge to the InvokeModelOrder entity by IDs.
func (mu *ModelUpdate) AddInvokeModelOrderIDs(ids ...int64) *ModelUpdate {
	mu.mutation.AddInvokeModelOrderIDs(ids...)
	return mu
}

// AddInvokeModelOrders adds the "invoke_model_orders" edges to the InvokeModelOrder entity.
func (mu *ModelUpdate) AddInvokeModelOrders(i ...*InvokeModelOrder) *ModelUpdate {
	ids := make([]int64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return mu.AddInvokeModelOrderIDs(ids...)
}

// Mutation returns the ModelMutation object of the builder.
func (mu *ModelUpdate) Mutation() *ModelMutation {
	return mu.mutation
}

// ClearModelPrices clears all "model_prices" edges to the ModelPrice entity.
func (mu *ModelUpdate) ClearModelPrices() *ModelUpdate {
	mu.mutation.ClearModelPrices()
	return mu
}

// RemoveModelPriceIDs removes the "model_prices" edge to ModelPrice entities by IDs.
func (mu *ModelUpdate) RemoveModelPriceIDs(ids ...int64) *ModelUpdate {
	mu.mutation.RemoveModelPriceIDs(ids...)
	return mu
}

// RemoveModelPrices removes "model_prices" edges to ModelPrice entities.
func (mu *ModelUpdate) RemoveModelPrices(m ...*ModelPrice) *ModelUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveModelPriceIDs(ids...)
}

// ClearUserModels clears all "user_models" edges to the UserModel entity.
func (mu *ModelUpdate) ClearUserModels() *ModelUpdate {
	mu.mutation.ClearUserModels()
	return mu
}

// RemoveUserModelIDs removes the "user_models" edge to UserModel entities by IDs.
func (mu *ModelUpdate) RemoveUserModelIDs(ids ...int64) *ModelUpdate {
	mu.mutation.RemoveUserModelIDs(ids...)
	return mu
}

// RemoveUserModels removes "user_models" edges to UserModel entities.
func (mu *ModelUpdate) RemoveUserModels(u ...*UserModel) *ModelUpdate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mu.RemoveUserModelIDs(ids...)
}

// ClearInvokeModelOrders clears all "invoke_model_orders" edges to the InvokeModelOrder entity.
func (mu *ModelUpdate) ClearInvokeModelOrders() *ModelUpdate {
	mu.mutation.ClearInvokeModelOrders()
	return mu
}

// RemoveInvokeModelOrderIDs removes the "invoke_model_orders" edge to InvokeModelOrder entities by IDs.
func (mu *ModelUpdate) RemoveInvokeModelOrderIDs(ids ...int64) *ModelUpdate {
	mu.mutation.RemoveInvokeModelOrderIDs(ids...)
	return mu
}

// RemoveInvokeModelOrders removes "invoke_model_orders" edges to InvokeModelOrder entities.
func (mu *ModelUpdate) RemoveInvokeModelOrders(i ...*InvokeModelOrder) *ModelUpdate {
	ids := make([]int64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return mu.RemoveInvokeModelOrderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *ModelUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *ModelUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *ModelUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *ModelUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *ModelUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := model.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *ModelUpdate) check() error {
	if v, ok := mu.mutation.ModelType(); ok {
		if err := model.ModelTypeValidator(v); err != nil {
			return &ValidationError{Name: "model_type", err: fmt.Errorf(`cep_ent: validator failed for field "Model.model_type": %w`, err)}
		}
	}
	if v, ok := mu.mutation.ModelStatus(); ok {
		if err := model.ModelStatusValidator(v); err != nil {
			return &ValidationError{Name: "model_status", err: fmt.Errorf(`cep_ent: validator failed for field "Model.model_status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *ModelUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ModelUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *ModelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(model.Table, model.Columns, sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.CreatedBy(); ok {
		_spec.SetField(model.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(model.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.UpdatedBy(); ok {
		_spec.SetField(model.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(model.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(model.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.DeletedAt(); ok {
		_spec.SetField(model.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(model.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Author(); ok {
		_spec.SetField(model.FieldAuthor, field.TypeString, value)
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.SetField(model.FieldDescription, field.TypeString, value)
	}
	if value, ok := mu.mutation.ModelType(); ok {
		_spec.SetField(model.FieldModelType, field.TypeEnum, value)
	}
	if value, ok := mu.mutation.ModelStatus(); ok {
		_spec.SetField(model.FieldModelStatus, field.TypeEnum, value)
	}
	if value, ok := mu.mutation.IsOfficial(); ok {
		_spec.SetField(model.FieldIsOfficial, field.TypeBool, value)
	}
	if value, ok := mu.mutation.TotalUsage(); ok {
		_spec.SetField(model.FieldTotalUsage, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedTotalUsage(); ok {
		_spec.AddField(model.FieldTotalUsage, field.TypeInt, value)
	}
	if mu.mutation.ModelPricesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ModelPricesTable,
			Columns: []string{model.ModelPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelprice.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedModelPricesIDs(); len(nodes) > 0 && !mu.mutation.ModelPricesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ModelPricesTable,
			Columns: []string{model.ModelPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelprice.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ModelPricesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ModelPricesTable,
			Columns: []string{model.ModelPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelprice.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.UserModelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.UserModelsTable,
			Columns: []string{model.UserModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedUserModelsIDs(); len(nodes) > 0 && !mu.mutation.UserModelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.UserModelsTable,
			Columns: []string{model.UserModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.UserModelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.UserModelsTable,
			Columns: []string{model.UserModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.InvokeModelOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.InvokeModelOrdersTable,
			Columns: []string{model.InvokeModelOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invokemodelorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedInvokeModelOrdersIDs(); len(nodes) > 0 && !mu.mutation.InvokeModelOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.InvokeModelOrdersTable,
			Columns: []string{model.InvokeModelOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invokemodelorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.InvokeModelOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.InvokeModelOrdersTable,
			Columns: []string{model.InvokeModelOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invokemodelorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{model.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// ModelUpdateOne is the builder for updating a single Model entity.
type ModelUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ModelMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (muo *ModelUpdateOne) SetCreatedBy(i int64) *ModelUpdateOne {
	muo.mutation.ResetCreatedBy()
	muo.mutation.SetCreatedBy(i)
	return muo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableCreatedBy(i *int64) *ModelUpdateOne {
	if i != nil {
		muo.SetCreatedBy(*i)
	}
	return muo
}

// AddCreatedBy adds i to the "created_by" field.
func (muo *ModelUpdateOne) AddCreatedBy(i int64) *ModelUpdateOne {
	muo.mutation.AddCreatedBy(i)
	return muo
}

// SetUpdatedBy sets the "updated_by" field.
func (muo *ModelUpdateOne) SetUpdatedBy(i int64) *ModelUpdateOne {
	muo.mutation.ResetUpdatedBy()
	muo.mutation.SetUpdatedBy(i)
	return muo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableUpdatedBy(i *int64) *ModelUpdateOne {
	if i != nil {
		muo.SetUpdatedBy(*i)
	}
	return muo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (muo *ModelUpdateOne) AddUpdatedBy(i int64) *ModelUpdateOne {
	muo.mutation.AddUpdatedBy(i)
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *ModelUpdateOne) SetUpdatedAt(t time.Time) *ModelUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetDeletedAt sets the "deleted_at" field.
func (muo *ModelUpdateOne) SetDeletedAt(t time.Time) *ModelUpdateOne {
	muo.mutation.SetDeletedAt(t)
	return muo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableDeletedAt(t *time.Time) *ModelUpdateOne {
	if t != nil {
		muo.SetDeletedAt(*t)
	}
	return muo
}

// SetName sets the "name" field.
func (muo *ModelUpdateOne) SetName(s string) *ModelUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableName(s *string) *ModelUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// SetAuthor sets the "author" field.
func (muo *ModelUpdateOne) SetAuthor(s string) *ModelUpdateOne {
	muo.mutation.SetAuthor(s)
	return muo
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableAuthor(s *string) *ModelUpdateOne {
	if s != nil {
		muo.SetAuthor(*s)
	}
	return muo
}

// SetDescription sets the "description" field.
func (muo *ModelUpdateOne) SetDescription(s string) *ModelUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableDescription(s *string) *ModelUpdateOne {
	if s != nil {
		muo.SetDescription(*s)
	}
	return muo
}

// SetModelType sets the "model_type" field.
func (muo *ModelUpdateOne) SetModelType(e enums.Model) *ModelUpdateOne {
	muo.mutation.SetModelType(e)
	return muo
}

// SetNillableModelType sets the "model_type" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableModelType(e *enums.Model) *ModelUpdateOne {
	if e != nil {
		muo.SetModelType(*e)
	}
	return muo
}

// SetModelStatus sets the "model_status" field.
func (muo *ModelUpdateOne) SetModelStatus(es enums.ModelStatus) *ModelUpdateOne {
	muo.mutation.SetModelStatus(es)
	return muo
}

// SetNillableModelStatus sets the "model_status" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableModelStatus(es *enums.ModelStatus) *ModelUpdateOne {
	if es != nil {
		muo.SetModelStatus(*es)
	}
	return muo
}

// SetIsOfficial sets the "is_official" field.
func (muo *ModelUpdateOne) SetIsOfficial(b bool) *ModelUpdateOne {
	muo.mutation.SetIsOfficial(b)
	return muo
}

// SetNillableIsOfficial sets the "is_official" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableIsOfficial(b *bool) *ModelUpdateOne {
	if b != nil {
		muo.SetIsOfficial(*b)
	}
	return muo
}

// SetTotalUsage sets the "total_usage" field.
func (muo *ModelUpdateOne) SetTotalUsage(i int) *ModelUpdateOne {
	muo.mutation.ResetTotalUsage()
	muo.mutation.SetTotalUsage(i)
	return muo
}

// SetNillableTotalUsage sets the "total_usage" field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableTotalUsage(i *int) *ModelUpdateOne {
	if i != nil {
		muo.SetTotalUsage(*i)
	}
	return muo
}

// AddTotalUsage adds i to the "total_usage" field.
func (muo *ModelUpdateOne) AddTotalUsage(i int) *ModelUpdateOne {
	muo.mutation.AddTotalUsage(i)
	return muo
}

// AddModelPriceIDs adds the "model_prices" edge to the ModelPrice entity by IDs.
func (muo *ModelUpdateOne) AddModelPriceIDs(ids ...int64) *ModelUpdateOne {
	muo.mutation.AddModelPriceIDs(ids...)
	return muo
}

// AddModelPrices adds the "model_prices" edges to the ModelPrice entity.
func (muo *ModelUpdateOne) AddModelPrices(m ...*ModelPrice) *ModelUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddModelPriceIDs(ids...)
}

// AddUserModelIDs adds the "user_models" edge to the UserModel entity by IDs.
func (muo *ModelUpdateOne) AddUserModelIDs(ids ...int64) *ModelUpdateOne {
	muo.mutation.AddUserModelIDs(ids...)
	return muo
}

// AddUserModels adds the "user_models" edges to the UserModel entity.
func (muo *ModelUpdateOne) AddUserModels(u ...*UserModel) *ModelUpdateOne {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return muo.AddUserModelIDs(ids...)
}

// AddInvokeModelOrderIDs adds the "invoke_model_orders" edge to the InvokeModelOrder entity by IDs.
func (muo *ModelUpdateOne) AddInvokeModelOrderIDs(ids ...int64) *ModelUpdateOne {
	muo.mutation.AddInvokeModelOrderIDs(ids...)
	return muo
}

// AddInvokeModelOrders adds the "invoke_model_orders" edges to the InvokeModelOrder entity.
func (muo *ModelUpdateOne) AddInvokeModelOrders(i ...*InvokeModelOrder) *ModelUpdateOne {
	ids := make([]int64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return muo.AddInvokeModelOrderIDs(ids...)
}

// Mutation returns the ModelMutation object of the builder.
func (muo *ModelUpdateOne) Mutation() *ModelMutation {
	return muo.mutation
}

// ClearModelPrices clears all "model_prices" edges to the ModelPrice entity.
func (muo *ModelUpdateOne) ClearModelPrices() *ModelUpdateOne {
	muo.mutation.ClearModelPrices()
	return muo
}

// RemoveModelPriceIDs removes the "model_prices" edge to ModelPrice entities by IDs.
func (muo *ModelUpdateOne) RemoveModelPriceIDs(ids ...int64) *ModelUpdateOne {
	muo.mutation.RemoveModelPriceIDs(ids...)
	return muo
}

// RemoveModelPrices removes "model_prices" edges to ModelPrice entities.
func (muo *ModelUpdateOne) RemoveModelPrices(m ...*ModelPrice) *ModelUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveModelPriceIDs(ids...)
}

// ClearUserModels clears all "user_models" edges to the UserModel entity.
func (muo *ModelUpdateOne) ClearUserModels() *ModelUpdateOne {
	muo.mutation.ClearUserModels()
	return muo
}

// RemoveUserModelIDs removes the "user_models" edge to UserModel entities by IDs.
func (muo *ModelUpdateOne) RemoveUserModelIDs(ids ...int64) *ModelUpdateOne {
	muo.mutation.RemoveUserModelIDs(ids...)
	return muo
}

// RemoveUserModels removes "user_models" edges to UserModel entities.
func (muo *ModelUpdateOne) RemoveUserModels(u ...*UserModel) *ModelUpdateOne {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return muo.RemoveUserModelIDs(ids...)
}

// ClearInvokeModelOrders clears all "invoke_model_orders" edges to the InvokeModelOrder entity.
func (muo *ModelUpdateOne) ClearInvokeModelOrders() *ModelUpdateOne {
	muo.mutation.ClearInvokeModelOrders()
	return muo
}

// RemoveInvokeModelOrderIDs removes the "invoke_model_orders" edge to InvokeModelOrder entities by IDs.
func (muo *ModelUpdateOne) RemoveInvokeModelOrderIDs(ids ...int64) *ModelUpdateOne {
	muo.mutation.RemoveInvokeModelOrderIDs(ids...)
	return muo
}

// RemoveInvokeModelOrders removes "invoke_model_orders" edges to InvokeModelOrder entities.
func (muo *ModelUpdateOne) RemoveInvokeModelOrders(i ...*InvokeModelOrder) *ModelUpdateOne {
	ids := make([]int64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return muo.RemoveInvokeModelOrderIDs(ids...)
}

// Where appends a list predicates to the ModelUpdate builder.
func (muo *ModelUpdateOne) Where(ps ...predicate.Model) *ModelUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *ModelUpdateOne) Select(field string, fields ...string) *ModelUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Model entity.
func (muo *ModelUpdateOne) Save(ctx context.Context) (*Model, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *ModelUpdateOne) SaveX(ctx context.Context) *Model {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *ModelUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *ModelUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *ModelUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := model.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *ModelUpdateOne) check() error {
	if v, ok := muo.mutation.ModelType(); ok {
		if err := model.ModelTypeValidator(v); err != nil {
			return &ValidationError{Name: "model_type", err: fmt.Errorf(`cep_ent: validator failed for field "Model.model_type": %w`, err)}
		}
	}
	if v, ok := muo.mutation.ModelStatus(); ok {
		if err := model.ModelStatusValidator(v); err != nil {
			return &ValidationError{Name: "model_status", err: fmt.Errorf(`cep_ent: validator failed for field "Model.model_status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *ModelUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ModelUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *ModelUpdateOne) sqlSave(ctx context.Context) (_node *Model, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(model.Table, model.Columns, sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "Model.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, model.FieldID)
		for _, f := range fields {
			if !model.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != model.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.CreatedBy(); ok {
		_spec.SetField(model.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(model.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.UpdatedBy(); ok {
		_spec.SetField(model.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(model.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(model.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.DeletedAt(); ok {
		_spec.SetField(model.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(model.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Author(); ok {
		_spec.SetField(model.FieldAuthor, field.TypeString, value)
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.SetField(model.FieldDescription, field.TypeString, value)
	}
	if value, ok := muo.mutation.ModelType(); ok {
		_spec.SetField(model.FieldModelType, field.TypeEnum, value)
	}
	if value, ok := muo.mutation.ModelStatus(); ok {
		_spec.SetField(model.FieldModelStatus, field.TypeEnum, value)
	}
	if value, ok := muo.mutation.IsOfficial(); ok {
		_spec.SetField(model.FieldIsOfficial, field.TypeBool, value)
	}
	if value, ok := muo.mutation.TotalUsage(); ok {
		_spec.SetField(model.FieldTotalUsage, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedTotalUsage(); ok {
		_spec.AddField(model.FieldTotalUsage, field.TypeInt, value)
	}
	if muo.mutation.ModelPricesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ModelPricesTable,
			Columns: []string{model.ModelPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelprice.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedModelPricesIDs(); len(nodes) > 0 && !muo.mutation.ModelPricesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ModelPricesTable,
			Columns: []string{model.ModelPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelprice.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ModelPricesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.ModelPricesTable,
			Columns: []string{model.ModelPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(modelprice.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.UserModelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.UserModelsTable,
			Columns: []string{model.UserModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedUserModelsIDs(); len(nodes) > 0 && !muo.mutation.UserModelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.UserModelsTable,
			Columns: []string{model.UserModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.UserModelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.UserModelsTable,
			Columns: []string{model.UserModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.InvokeModelOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.InvokeModelOrdersTable,
			Columns: []string{model.InvokeModelOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invokemodelorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedInvokeModelOrdersIDs(); len(nodes) > 0 && !muo.mutation.InvokeModelOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.InvokeModelOrdersTable,
			Columns: []string{model.InvokeModelOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invokemodelorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.InvokeModelOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.InvokeModelOrdersTable,
			Columns: []string{model.InvokeModelOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invokemodelorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(muo.modifiers...)
	_node = &Model{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{model.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}