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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/frpcinfo"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/frpsinfo"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// FrpsInfoUpdate is the builder for updating FrpsInfo entities.
type FrpsInfoUpdate struct {
	config
	hooks     []Hook
	mutation  *FrpsInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FrpsInfoUpdate builder.
func (fiu *FrpsInfoUpdate) Where(ps ...predicate.FrpsInfo) *FrpsInfoUpdate {
	fiu.mutation.Where(ps...)
	return fiu
}

// SetCreatedBy sets the "created_by" field.
func (fiu *FrpsInfoUpdate) SetCreatedBy(i int64) *FrpsInfoUpdate {
	fiu.mutation.ResetCreatedBy()
	fiu.mutation.SetCreatedBy(i)
	return fiu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (fiu *FrpsInfoUpdate) SetNillableCreatedBy(i *int64) *FrpsInfoUpdate {
	if i != nil {
		fiu.SetCreatedBy(*i)
	}
	return fiu
}

// AddCreatedBy adds i to the "created_by" field.
func (fiu *FrpsInfoUpdate) AddCreatedBy(i int64) *FrpsInfoUpdate {
	fiu.mutation.AddCreatedBy(i)
	return fiu
}

// SetUpdatedBy sets the "updated_by" field.
func (fiu *FrpsInfoUpdate) SetUpdatedBy(i int64) *FrpsInfoUpdate {
	fiu.mutation.ResetUpdatedBy()
	fiu.mutation.SetUpdatedBy(i)
	return fiu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fiu *FrpsInfoUpdate) SetNillableUpdatedBy(i *int64) *FrpsInfoUpdate {
	if i != nil {
		fiu.SetUpdatedBy(*i)
	}
	return fiu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (fiu *FrpsInfoUpdate) AddUpdatedBy(i int64) *FrpsInfoUpdate {
	fiu.mutation.AddUpdatedBy(i)
	return fiu
}

// SetUpdatedAt sets the "updated_at" field.
func (fiu *FrpsInfoUpdate) SetUpdatedAt(t time.Time) *FrpsInfoUpdate {
	fiu.mutation.SetUpdatedAt(t)
	return fiu
}

// SetDeletedAt sets the "deleted_at" field.
func (fiu *FrpsInfoUpdate) SetDeletedAt(t time.Time) *FrpsInfoUpdate {
	fiu.mutation.SetDeletedAt(t)
	return fiu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fiu *FrpsInfoUpdate) SetNillableDeletedAt(t *time.Time) *FrpsInfoUpdate {
	if t != nil {
		fiu.SetDeletedAt(*t)
	}
	return fiu
}

// SetTag sets the "tag" field.
func (fiu *FrpsInfoUpdate) SetTag(s string) *FrpsInfoUpdate {
	fiu.mutation.SetTag(s)
	return fiu
}

// SetNillableTag sets the "tag" field if the given value is not nil.
func (fiu *FrpsInfoUpdate) SetNillableTag(s *string) *FrpsInfoUpdate {
	if s != nil {
		fiu.SetTag(*s)
	}
	return fiu
}

// SetServerAddr sets the "server_addr" field.
func (fiu *FrpsInfoUpdate) SetServerAddr(s string) *FrpsInfoUpdate {
	fiu.mutation.SetServerAddr(s)
	return fiu
}

// SetNillableServerAddr sets the "server_addr" field if the given value is not nil.
func (fiu *FrpsInfoUpdate) SetNillableServerAddr(s *string) *FrpsInfoUpdate {
	if s != nil {
		fiu.SetServerAddr(*s)
	}
	return fiu
}

// SetServerPort sets the "server_port" field.
func (fiu *FrpsInfoUpdate) SetServerPort(i int) *FrpsInfoUpdate {
	fiu.mutation.ResetServerPort()
	fiu.mutation.SetServerPort(i)
	return fiu
}

// SetNillableServerPort sets the "server_port" field if the given value is not nil.
func (fiu *FrpsInfoUpdate) SetNillableServerPort(i *int) *FrpsInfoUpdate {
	if i != nil {
		fiu.SetServerPort(*i)
	}
	return fiu
}

// AddServerPort adds i to the "server_port" field.
func (fiu *FrpsInfoUpdate) AddServerPort(i int) *FrpsInfoUpdate {
	fiu.mutation.AddServerPort(i)
	return fiu
}

// SetAuthenticationMethod sets the "authentication_method" field.
func (fiu *FrpsInfoUpdate) SetAuthenticationMethod(s string) *FrpsInfoUpdate {
	fiu.mutation.SetAuthenticationMethod(s)
	return fiu
}

// SetNillableAuthenticationMethod sets the "authentication_method" field if the given value is not nil.
func (fiu *FrpsInfoUpdate) SetNillableAuthenticationMethod(s *string) *FrpsInfoUpdate {
	if s != nil {
		fiu.SetAuthenticationMethod(*s)
	}
	return fiu
}

// SetToken sets the "token" field.
func (fiu *FrpsInfoUpdate) SetToken(s string) *FrpsInfoUpdate {
	fiu.mutation.SetToken(s)
	return fiu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (fiu *FrpsInfoUpdate) SetNillableToken(s *string) *FrpsInfoUpdate {
	if s != nil {
		fiu.SetToken(*s)
	}
	return fiu
}

// SetType sets the "type" field.
func (fiu *FrpsInfoUpdate) SetType(s string) *FrpsInfoUpdate {
	fiu.mutation.SetType(s)
	return fiu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (fiu *FrpsInfoUpdate) SetNillableType(s *string) *FrpsInfoUpdate {
	if s != nil {
		fiu.SetType(*s)
	}
	return fiu
}

// AddFrpcInfoIDs adds the "frpc_infos" edge to the FrpcInfo entity by IDs.
func (fiu *FrpsInfoUpdate) AddFrpcInfoIDs(ids ...int64) *FrpsInfoUpdate {
	fiu.mutation.AddFrpcInfoIDs(ids...)
	return fiu
}

// AddFrpcInfos adds the "frpc_infos" edges to the FrpcInfo entity.
func (fiu *FrpsInfoUpdate) AddFrpcInfos(f ...*FrpcInfo) *FrpsInfoUpdate {
	ids := make([]int64, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fiu.AddFrpcInfoIDs(ids...)
}

// Mutation returns the FrpsInfoMutation object of the builder.
func (fiu *FrpsInfoUpdate) Mutation() *FrpsInfoMutation {
	return fiu.mutation
}

// ClearFrpcInfos clears all "frpc_infos" edges to the FrpcInfo entity.
func (fiu *FrpsInfoUpdate) ClearFrpcInfos() *FrpsInfoUpdate {
	fiu.mutation.ClearFrpcInfos()
	return fiu
}

// RemoveFrpcInfoIDs removes the "frpc_infos" edge to FrpcInfo entities by IDs.
func (fiu *FrpsInfoUpdate) RemoveFrpcInfoIDs(ids ...int64) *FrpsInfoUpdate {
	fiu.mutation.RemoveFrpcInfoIDs(ids...)
	return fiu
}

// RemoveFrpcInfos removes "frpc_infos" edges to FrpcInfo entities.
func (fiu *FrpsInfoUpdate) RemoveFrpcInfos(f ...*FrpcInfo) *FrpsInfoUpdate {
	ids := make([]int64, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fiu.RemoveFrpcInfoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fiu *FrpsInfoUpdate) Save(ctx context.Context) (int, error) {
	fiu.defaults()
	return withHooks(ctx, fiu.sqlSave, fiu.mutation, fiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fiu *FrpsInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := fiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fiu *FrpsInfoUpdate) Exec(ctx context.Context) error {
	_, err := fiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fiu *FrpsInfoUpdate) ExecX(ctx context.Context) {
	if err := fiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fiu *FrpsInfoUpdate) defaults() {
	if _, ok := fiu.mutation.UpdatedAt(); !ok {
		v := frpsinfo.UpdateDefaultUpdatedAt()
		fiu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fiu *FrpsInfoUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FrpsInfoUpdate {
	fiu.modifiers = append(fiu.modifiers, modifiers...)
	return fiu
}

func (fiu *FrpsInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(frpsinfo.Table, frpsinfo.Columns, sqlgraph.NewFieldSpec(frpsinfo.FieldID, field.TypeInt64))
	if ps := fiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fiu.mutation.CreatedBy(); ok {
		_spec.SetField(frpsinfo.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := fiu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(frpsinfo.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := fiu.mutation.UpdatedBy(); ok {
		_spec.SetField(frpsinfo.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := fiu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(frpsinfo.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := fiu.mutation.UpdatedAt(); ok {
		_spec.SetField(frpsinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fiu.mutation.DeletedAt(); ok {
		_spec.SetField(frpsinfo.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := fiu.mutation.Tag(); ok {
		_spec.SetField(frpsinfo.FieldTag, field.TypeString, value)
	}
	if value, ok := fiu.mutation.ServerAddr(); ok {
		_spec.SetField(frpsinfo.FieldServerAddr, field.TypeString, value)
	}
	if value, ok := fiu.mutation.ServerPort(); ok {
		_spec.SetField(frpsinfo.FieldServerPort, field.TypeInt, value)
	}
	if value, ok := fiu.mutation.AddedServerPort(); ok {
		_spec.AddField(frpsinfo.FieldServerPort, field.TypeInt, value)
	}
	if value, ok := fiu.mutation.AuthenticationMethod(); ok {
		_spec.SetField(frpsinfo.FieldAuthenticationMethod, field.TypeString, value)
	}
	if value, ok := fiu.mutation.Token(); ok {
		_spec.SetField(frpsinfo.FieldToken, field.TypeString, value)
	}
	if value, ok := fiu.mutation.GetType(); ok {
		_spec.SetField(frpsinfo.FieldType, field.TypeString, value)
	}
	if fiu.mutation.FrpcInfosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   frpsinfo.FrpcInfosTable,
			Columns: []string{frpsinfo.FrpcInfosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(frpcinfo.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiu.mutation.RemovedFrpcInfosIDs(); len(nodes) > 0 && !fiu.mutation.FrpcInfosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   frpsinfo.FrpcInfosTable,
			Columns: []string{frpsinfo.FrpcInfosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(frpcinfo.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiu.mutation.FrpcInfosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   frpsinfo.FrpcInfosTable,
			Columns: []string{frpsinfo.FrpcInfosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(frpcinfo.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(fiu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, fiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{frpsinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fiu.mutation.done = true
	return n, nil
}

// FrpsInfoUpdateOne is the builder for updating a single FrpsInfo entity.
type FrpsInfoUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FrpsInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (fiuo *FrpsInfoUpdateOne) SetCreatedBy(i int64) *FrpsInfoUpdateOne {
	fiuo.mutation.ResetCreatedBy()
	fiuo.mutation.SetCreatedBy(i)
	return fiuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (fiuo *FrpsInfoUpdateOne) SetNillableCreatedBy(i *int64) *FrpsInfoUpdateOne {
	if i != nil {
		fiuo.SetCreatedBy(*i)
	}
	return fiuo
}

// AddCreatedBy adds i to the "created_by" field.
func (fiuo *FrpsInfoUpdateOne) AddCreatedBy(i int64) *FrpsInfoUpdateOne {
	fiuo.mutation.AddCreatedBy(i)
	return fiuo
}

// SetUpdatedBy sets the "updated_by" field.
func (fiuo *FrpsInfoUpdateOne) SetUpdatedBy(i int64) *FrpsInfoUpdateOne {
	fiuo.mutation.ResetUpdatedBy()
	fiuo.mutation.SetUpdatedBy(i)
	return fiuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fiuo *FrpsInfoUpdateOne) SetNillableUpdatedBy(i *int64) *FrpsInfoUpdateOne {
	if i != nil {
		fiuo.SetUpdatedBy(*i)
	}
	return fiuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (fiuo *FrpsInfoUpdateOne) AddUpdatedBy(i int64) *FrpsInfoUpdateOne {
	fiuo.mutation.AddUpdatedBy(i)
	return fiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fiuo *FrpsInfoUpdateOne) SetUpdatedAt(t time.Time) *FrpsInfoUpdateOne {
	fiuo.mutation.SetUpdatedAt(t)
	return fiuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fiuo *FrpsInfoUpdateOne) SetDeletedAt(t time.Time) *FrpsInfoUpdateOne {
	fiuo.mutation.SetDeletedAt(t)
	return fiuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fiuo *FrpsInfoUpdateOne) SetNillableDeletedAt(t *time.Time) *FrpsInfoUpdateOne {
	if t != nil {
		fiuo.SetDeletedAt(*t)
	}
	return fiuo
}

// SetTag sets the "tag" field.
func (fiuo *FrpsInfoUpdateOne) SetTag(s string) *FrpsInfoUpdateOne {
	fiuo.mutation.SetTag(s)
	return fiuo
}

// SetNillableTag sets the "tag" field if the given value is not nil.
func (fiuo *FrpsInfoUpdateOne) SetNillableTag(s *string) *FrpsInfoUpdateOne {
	if s != nil {
		fiuo.SetTag(*s)
	}
	return fiuo
}

// SetServerAddr sets the "server_addr" field.
func (fiuo *FrpsInfoUpdateOne) SetServerAddr(s string) *FrpsInfoUpdateOne {
	fiuo.mutation.SetServerAddr(s)
	return fiuo
}

// SetNillableServerAddr sets the "server_addr" field if the given value is not nil.
func (fiuo *FrpsInfoUpdateOne) SetNillableServerAddr(s *string) *FrpsInfoUpdateOne {
	if s != nil {
		fiuo.SetServerAddr(*s)
	}
	return fiuo
}

// SetServerPort sets the "server_port" field.
func (fiuo *FrpsInfoUpdateOne) SetServerPort(i int) *FrpsInfoUpdateOne {
	fiuo.mutation.ResetServerPort()
	fiuo.mutation.SetServerPort(i)
	return fiuo
}

// SetNillableServerPort sets the "server_port" field if the given value is not nil.
func (fiuo *FrpsInfoUpdateOne) SetNillableServerPort(i *int) *FrpsInfoUpdateOne {
	if i != nil {
		fiuo.SetServerPort(*i)
	}
	return fiuo
}

// AddServerPort adds i to the "server_port" field.
func (fiuo *FrpsInfoUpdateOne) AddServerPort(i int) *FrpsInfoUpdateOne {
	fiuo.mutation.AddServerPort(i)
	return fiuo
}

// SetAuthenticationMethod sets the "authentication_method" field.
func (fiuo *FrpsInfoUpdateOne) SetAuthenticationMethod(s string) *FrpsInfoUpdateOne {
	fiuo.mutation.SetAuthenticationMethod(s)
	return fiuo
}

// SetNillableAuthenticationMethod sets the "authentication_method" field if the given value is not nil.
func (fiuo *FrpsInfoUpdateOne) SetNillableAuthenticationMethod(s *string) *FrpsInfoUpdateOne {
	if s != nil {
		fiuo.SetAuthenticationMethod(*s)
	}
	return fiuo
}

// SetToken sets the "token" field.
func (fiuo *FrpsInfoUpdateOne) SetToken(s string) *FrpsInfoUpdateOne {
	fiuo.mutation.SetToken(s)
	return fiuo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (fiuo *FrpsInfoUpdateOne) SetNillableToken(s *string) *FrpsInfoUpdateOne {
	if s != nil {
		fiuo.SetToken(*s)
	}
	return fiuo
}

// SetType sets the "type" field.
func (fiuo *FrpsInfoUpdateOne) SetType(s string) *FrpsInfoUpdateOne {
	fiuo.mutation.SetType(s)
	return fiuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (fiuo *FrpsInfoUpdateOne) SetNillableType(s *string) *FrpsInfoUpdateOne {
	if s != nil {
		fiuo.SetType(*s)
	}
	return fiuo
}

// AddFrpcInfoIDs adds the "frpc_infos" edge to the FrpcInfo entity by IDs.
func (fiuo *FrpsInfoUpdateOne) AddFrpcInfoIDs(ids ...int64) *FrpsInfoUpdateOne {
	fiuo.mutation.AddFrpcInfoIDs(ids...)
	return fiuo
}

// AddFrpcInfos adds the "frpc_infos" edges to the FrpcInfo entity.
func (fiuo *FrpsInfoUpdateOne) AddFrpcInfos(f ...*FrpcInfo) *FrpsInfoUpdateOne {
	ids := make([]int64, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fiuo.AddFrpcInfoIDs(ids...)
}

// Mutation returns the FrpsInfoMutation object of the builder.
func (fiuo *FrpsInfoUpdateOne) Mutation() *FrpsInfoMutation {
	return fiuo.mutation
}

// ClearFrpcInfos clears all "frpc_infos" edges to the FrpcInfo entity.
func (fiuo *FrpsInfoUpdateOne) ClearFrpcInfos() *FrpsInfoUpdateOne {
	fiuo.mutation.ClearFrpcInfos()
	return fiuo
}

// RemoveFrpcInfoIDs removes the "frpc_infos" edge to FrpcInfo entities by IDs.
func (fiuo *FrpsInfoUpdateOne) RemoveFrpcInfoIDs(ids ...int64) *FrpsInfoUpdateOne {
	fiuo.mutation.RemoveFrpcInfoIDs(ids...)
	return fiuo
}

// RemoveFrpcInfos removes "frpc_infos" edges to FrpcInfo entities.
func (fiuo *FrpsInfoUpdateOne) RemoveFrpcInfos(f ...*FrpcInfo) *FrpsInfoUpdateOne {
	ids := make([]int64, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fiuo.RemoveFrpcInfoIDs(ids...)
}

// Where appends a list predicates to the FrpsInfoUpdate builder.
func (fiuo *FrpsInfoUpdateOne) Where(ps ...predicate.FrpsInfo) *FrpsInfoUpdateOne {
	fiuo.mutation.Where(ps...)
	return fiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fiuo *FrpsInfoUpdateOne) Select(field string, fields ...string) *FrpsInfoUpdateOne {
	fiuo.fields = append([]string{field}, fields...)
	return fiuo
}

// Save executes the query and returns the updated FrpsInfo entity.
func (fiuo *FrpsInfoUpdateOne) Save(ctx context.Context) (*FrpsInfo, error) {
	fiuo.defaults()
	return withHooks(ctx, fiuo.sqlSave, fiuo.mutation, fiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fiuo *FrpsInfoUpdateOne) SaveX(ctx context.Context) *FrpsInfo {
	node, err := fiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fiuo *FrpsInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := fiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fiuo *FrpsInfoUpdateOne) ExecX(ctx context.Context) {
	if err := fiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fiuo *FrpsInfoUpdateOne) defaults() {
	if _, ok := fiuo.mutation.UpdatedAt(); !ok {
		v := frpsinfo.UpdateDefaultUpdatedAt()
		fiuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fiuo *FrpsInfoUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FrpsInfoUpdateOne {
	fiuo.modifiers = append(fiuo.modifiers, modifiers...)
	return fiuo
}

func (fiuo *FrpsInfoUpdateOne) sqlSave(ctx context.Context) (_node *FrpsInfo, err error) {
	_spec := sqlgraph.NewUpdateSpec(frpsinfo.Table, frpsinfo.Columns, sqlgraph.NewFieldSpec(frpsinfo.FieldID, field.TypeInt64))
	id, ok := fiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "FrpsInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, frpsinfo.FieldID)
		for _, f := range fields {
			if !frpsinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != frpsinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fiuo.mutation.CreatedBy(); ok {
		_spec.SetField(frpsinfo.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := fiuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(frpsinfo.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := fiuo.mutation.UpdatedBy(); ok {
		_spec.SetField(frpsinfo.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := fiuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(frpsinfo.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := fiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(frpsinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fiuo.mutation.DeletedAt(); ok {
		_spec.SetField(frpsinfo.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := fiuo.mutation.Tag(); ok {
		_spec.SetField(frpsinfo.FieldTag, field.TypeString, value)
	}
	if value, ok := fiuo.mutation.ServerAddr(); ok {
		_spec.SetField(frpsinfo.FieldServerAddr, field.TypeString, value)
	}
	if value, ok := fiuo.mutation.ServerPort(); ok {
		_spec.SetField(frpsinfo.FieldServerPort, field.TypeInt, value)
	}
	if value, ok := fiuo.mutation.AddedServerPort(); ok {
		_spec.AddField(frpsinfo.FieldServerPort, field.TypeInt, value)
	}
	if value, ok := fiuo.mutation.AuthenticationMethod(); ok {
		_spec.SetField(frpsinfo.FieldAuthenticationMethod, field.TypeString, value)
	}
	if value, ok := fiuo.mutation.Token(); ok {
		_spec.SetField(frpsinfo.FieldToken, field.TypeString, value)
	}
	if value, ok := fiuo.mutation.GetType(); ok {
		_spec.SetField(frpsinfo.FieldType, field.TypeString, value)
	}
	if fiuo.mutation.FrpcInfosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   frpsinfo.FrpcInfosTable,
			Columns: []string{frpsinfo.FrpcInfosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(frpcinfo.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiuo.mutation.RemovedFrpcInfosIDs(); len(nodes) > 0 && !fiuo.mutation.FrpcInfosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   frpsinfo.FrpcInfosTable,
			Columns: []string{frpsinfo.FrpcInfosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(frpcinfo.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiuo.mutation.FrpcInfosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   frpsinfo.FrpcInfosTable,
			Columns: []string{frpsinfo.FrpcInfosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(frpcinfo.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(fiuo.modifiers...)
	_node = &FrpsInfo{config: fiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{frpsinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fiuo.mutation.done = true
	return _node, nil
}
