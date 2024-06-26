// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/enumcondition"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// EnumConditionQuery is the builder for querying EnumCondition entities.
type EnumConditionQuery struct {
	config
	ctx        *QueryContext
	order      []enumcondition.OrderOption
	inters     []Interceptor
	predicates []predicate.EnumCondition
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EnumConditionQuery builder.
func (ecq *EnumConditionQuery) Where(ps ...predicate.EnumCondition) *EnumConditionQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit the number of records to be returned by this query.
func (ecq *EnumConditionQuery) Limit(limit int) *EnumConditionQuery {
	ecq.ctx.Limit = &limit
	return ecq
}

// Offset to start from.
func (ecq *EnumConditionQuery) Offset(offset int) *EnumConditionQuery {
	ecq.ctx.Offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *EnumConditionQuery) Unique(unique bool) *EnumConditionQuery {
	ecq.ctx.Unique = &unique
	return ecq
}

// Order specifies how the records should be ordered.
func (ecq *EnumConditionQuery) Order(o ...enumcondition.OrderOption) *EnumConditionQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// First returns the first EnumCondition entity from the query.
// Returns a *NotFoundError when no EnumCondition was found.
func (ecq *EnumConditionQuery) First(ctx context.Context) (*EnumCondition, error) {
	nodes, err := ecq.Limit(1).All(setContextOp(ctx, ecq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{enumcondition.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *EnumConditionQuery) FirstX(ctx context.Context) *EnumCondition {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EnumCondition ID from the query.
// Returns a *NotFoundError when no EnumCondition ID was found.
func (ecq *EnumConditionQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ecq.Limit(1).IDs(setContextOp(ctx, ecq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{enumcondition.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *EnumConditionQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EnumCondition entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EnumCondition entity is found.
// Returns a *NotFoundError when no EnumCondition entities are found.
func (ecq *EnumConditionQuery) Only(ctx context.Context) (*EnumCondition, error) {
	nodes, err := ecq.Limit(2).All(setContextOp(ctx, ecq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{enumcondition.Label}
	default:
		return nil, &NotSingularError{enumcondition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *EnumConditionQuery) OnlyX(ctx context.Context) *EnumCondition {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EnumCondition ID in the query.
// Returns a *NotSingularError when more than one EnumCondition ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *EnumConditionQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ecq.Limit(2).IDs(setContextOp(ctx, ecq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{enumcondition.Label}
	default:
		err = &NotSingularError{enumcondition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *EnumConditionQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EnumConditions.
func (ecq *EnumConditionQuery) All(ctx context.Context) ([]*EnumCondition, error) {
	ctx = setContextOp(ctx, ecq.ctx, "All")
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EnumCondition, *EnumConditionQuery]()
	return withInterceptors[[]*EnumCondition](ctx, ecq, qr, ecq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ecq *EnumConditionQuery) AllX(ctx context.Context) []*EnumCondition {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EnumCondition IDs.
func (ecq *EnumConditionQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if ecq.ctx.Unique == nil && ecq.path != nil {
		ecq.Unique(true)
	}
	ctx = setContextOp(ctx, ecq.ctx, "IDs")
	if err = ecq.Select(enumcondition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *EnumConditionQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *EnumConditionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Count")
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ecq, querierCount[*EnumConditionQuery](), ecq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *EnumConditionQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *EnumConditionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Exist")
	switch _, err := ecq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ecq *EnumConditionQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EnumConditionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *EnumConditionQuery) Clone() *EnumConditionQuery {
	if ecq == nil {
		return nil
	}
	return &EnumConditionQuery{
		config:     ecq.config,
		ctx:        ecq.ctx.Clone(),
		order:      append([]enumcondition.OrderOption{}, ecq.order...),
		inters:     append([]Interceptor{}, ecq.inters...),
		predicates: append([]predicate.EnumCondition{}, ecq.predicates...),
		// clone intermediate query.
		sql:  ecq.sql.Clone(),
		path: ecq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by,string"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EnumCondition.Query().
//		GroupBy(enumcondition.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (ecq *EnumConditionQuery) GroupBy(field string, fields ...string) *EnumConditionGroupBy {
	ecq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EnumConditionGroupBy{build: ecq}
	grbuild.flds = &ecq.ctx.Fields
	grbuild.label = enumcondition.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by,string"`
//	}
//
//	client.EnumCondition.Query().
//		Select(enumcondition.FieldCreatedBy).
//		Scan(ctx, &v)
func (ecq *EnumConditionQuery) Select(fields ...string) *EnumConditionSelect {
	ecq.ctx.Fields = append(ecq.ctx.Fields, fields...)
	sbuild := &EnumConditionSelect{EnumConditionQuery: ecq}
	sbuild.label = enumcondition.Label
	sbuild.flds, sbuild.scan = &ecq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EnumConditionSelect configured with the given aggregations.
func (ecq *EnumConditionQuery) Aggregate(fns ...AggregateFunc) *EnumConditionSelect {
	return ecq.Select().Aggregate(fns...)
}

func (ecq *EnumConditionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ecq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ecq); err != nil {
				return err
			}
		}
	}
	for _, f := range ecq.ctx.Fields {
		if !enumcondition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if ecq.path != nil {
		prev, err := ecq.path(ctx)
		if err != nil {
			return err
		}
		ecq.sql = prev
	}
	return nil
}

func (ecq *EnumConditionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EnumCondition, error) {
	var (
		nodes = []*EnumCondition{}
		_spec = ecq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EnumCondition).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EnumCondition{config: ecq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(ecq.modifiers) > 0 {
		_spec.Modifiers = ecq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ecq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ecq *EnumConditionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	if len(ecq.modifiers) > 0 {
		_spec.Modifiers = ecq.modifiers
	}
	_spec.Node.Columns = ecq.ctx.Fields
	if len(ecq.ctx.Fields) > 0 {
		_spec.Unique = ecq.ctx.Unique != nil && *ecq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *EnumConditionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(enumcondition.Table, enumcondition.Columns, sqlgraph.NewFieldSpec(enumcondition.FieldID, field.TypeInt64))
	_spec.From = ecq.sql
	if unique := ecq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ecq.path != nil {
		_spec.Unique = true
	}
	if fields := ecq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, enumcondition.FieldID)
		for i := range fields {
			if fields[i] != enumcondition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ecq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ecq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ecq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ecq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ecq *EnumConditionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(enumcondition.Table)
	columns := ecq.ctx.Fields
	if len(columns) == 0 {
		columns = enumcondition.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ecq.sql != nil {
		selector = ecq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ecq.ctx.Unique != nil && *ecq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ecq.modifiers {
		m(selector)
	}
	for _, p := range ecq.predicates {
		p(selector)
	}
	for _, p := range ecq.order {
		p(selector)
	}
	if offset := ecq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ecq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ecq *EnumConditionQuery) Modify(modifiers ...func(s *sql.Selector)) *EnumConditionSelect {
	ecq.modifiers = append(ecq.modifiers, modifiers...)
	return ecq.Select()
}

// EnumConditionGroupBy is the group-by builder for EnumCondition entities.
type EnumConditionGroupBy struct {
	selector
	build *EnumConditionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *EnumConditionGroupBy) Aggregate(fns ...AggregateFunc) *EnumConditionGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the selector query and scans the result into the given value.
func (ecgb *EnumConditionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecgb.build.ctx, "GroupBy")
	if err := ecgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnumConditionQuery, *EnumConditionGroupBy](ctx, ecgb.build, ecgb, ecgb.build.inters, v)
}

func (ecgb *EnumConditionGroupBy) sqlScan(ctx context.Context, root *EnumConditionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ecgb.fns))
	for _, fn := range ecgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ecgb.flds)+len(ecgb.fns))
		for _, f := range *ecgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ecgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EnumConditionSelect is the builder for selecting fields of EnumCondition entities.
type EnumConditionSelect struct {
	*EnumConditionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ecs *EnumConditionSelect) Aggregate(fns ...AggregateFunc) *EnumConditionSelect {
	ecs.fns = append(ecs.fns, fns...)
	return ecs
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *EnumConditionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecs.ctx, "Select")
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnumConditionQuery, *EnumConditionSelect](ctx, ecs.EnumConditionQuery, ecs, ecs.inters, v)
}

func (ecs *EnumConditionSelect) sqlScan(ctx context.Context, root *EnumConditionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ecs.fns))
	for _, fn := range ecs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ecs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ecs *EnumConditionSelect) Modify(modifiers ...func(s *sql.Selector)) *EnumConditionSelect {
	ecs.modifiers = append(ecs.modifiers, modifiers...)
	return ecs
}
