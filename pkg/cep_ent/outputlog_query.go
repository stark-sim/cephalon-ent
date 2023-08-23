// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/outputlog"
	"cephalon-ent/pkg/cep_ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OutputLogQuery is the builder for querying OutputLog entities.
type OutputLogQuery struct {
	config
	ctx        *QueryContext
	order      []outputlog.OrderOption
	inters     []Interceptor
	predicates []predicate.OutputLog
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutputLogQuery builder.
func (olq *OutputLogQuery) Where(ps ...predicate.OutputLog) *OutputLogQuery {
	olq.predicates = append(olq.predicates, ps...)
	return olq
}

// Limit the number of records to be returned by this query.
func (olq *OutputLogQuery) Limit(limit int) *OutputLogQuery {
	olq.ctx.Limit = &limit
	return olq
}

// Offset to start from.
func (olq *OutputLogQuery) Offset(offset int) *OutputLogQuery {
	olq.ctx.Offset = &offset
	return olq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (olq *OutputLogQuery) Unique(unique bool) *OutputLogQuery {
	olq.ctx.Unique = &unique
	return olq
}

// Order specifies how the records should be ordered.
func (olq *OutputLogQuery) Order(o ...outputlog.OrderOption) *OutputLogQuery {
	olq.order = append(olq.order, o...)
	return olq
}

// First returns the first OutputLog entity from the query.
// Returns a *NotFoundError when no OutputLog was found.
func (olq *OutputLogQuery) First(ctx context.Context) (*OutputLog, error) {
	nodes, err := olq.Limit(1).All(setContextOp(ctx, olq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outputlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (olq *OutputLogQuery) FirstX(ctx context.Context) *OutputLog {
	node, err := olq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutputLog ID from the query.
// Returns a *NotFoundError when no OutputLog ID was found.
func (olq *OutputLogQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = olq.Limit(1).IDs(setContextOp(ctx, olq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outputlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (olq *OutputLogQuery) FirstIDX(ctx context.Context) int64 {
	id, err := olq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutputLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutputLog entity is found.
// Returns a *NotFoundError when no OutputLog entities are found.
func (olq *OutputLogQuery) Only(ctx context.Context) (*OutputLog, error) {
	nodes, err := olq.Limit(2).All(setContextOp(ctx, olq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outputlog.Label}
	default:
		return nil, &NotSingularError{outputlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (olq *OutputLogQuery) OnlyX(ctx context.Context) *OutputLog {
	node, err := olq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutputLog ID in the query.
// Returns a *NotSingularError when more than one OutputLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (olq *OutputLogQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = olq.Limit(2).IDs(setContextOp(ctx, olq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outputlog.Label}
	default:
		err = &NotSingularError{outputlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (olq *OutputLogQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := olq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutputLogs.
func (olq *OutputLogQuery) All(ctx context.Context) ([]*OutputLog, error) {
	ctx = setContextOp(ctx, olq.ctx, "All")
	if err := olq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OutputLog, *OutputLogQuery]()
	return withInterceptors[[]*OutputLog](ctx, olq, qr, olq.inters)
}

// AllX is like All, but panics if an error occurs.
func (olq *OutputLogQuery) AllX(ctx context.Context) []*OutputLog {
	nodes, err := olq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutputLog IDs.
func (olq *OutputLogQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if olq.ctx.Unique == nil && olq.path != nil {
		olq.Unique(true)
	}
	ctx = setContextOp(ctx, olq.ctx, "IDs")
	if err = olq.Select(outputlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (olq *OutputLogQuery) IDsX(ctx context.Context) []int64 {
	ids, err := olq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (olq *OutputLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, olq.ctx, "Count")
	if err := olq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, olq, querierCount[*OutputLogQuery](), olq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (olq *OutputLogQuery) CountX(ctx context.Context) int {
	count, err := olq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (olq *OutputLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, olq.ctx, "Exist")
	switch _, err := olq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (olq *OutputLogQuery) ExistX(ctx context.Context) bool {
	exist, err := olq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutputLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (olq *OutputLogQuery) Clone() *OutputLogQuery {
	if olq == nil {
		return nil
	}
	return &OutputLogQuery{
		config:     olq.config,
		ctx:        olq.ctx.Clone(),
		order:      append([]outputlog.OrderOption{}, olq.order...),
		inters:     append([]Interceptor{}, olq.inters...),
		predicates: append([]predicate.OutputLog{}, olq.predicates...),
		// clone intermediate query.
		sql:  olq.sql.Clone(),
		path: olq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutputLog.Query().
//		GroupBy(outputlog.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (olq *OutputLogQuery) GroupBy(field string, fields ...string) *OutputLogGroupBy {
	olq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OutputLogGroupBy{build: olq}
	grbuild.flds = &olq.ctx.Fields
	grbuild.label = outputlog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//	}
//
//	client.OutputLog.Query().
//		Select(outputlog.FieldCreatedBy).
//		Scan(ctx, &v)
func (olq *OutputLogQuery) Select(fields ...string) *OutputLogSelect {
	olq.ctx.Fields = append(olq.ctx.Fields, fields...)
	sbuild := &OutputLogSelect{OutputLogQuery: olq}
	sbuild.label = outputlog.Label
	sbuild.flds, sbuild.scan = &olq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OutputLogSelect configured with the given aggregations.
func (olq *OutputLogQuery) Aggregate(fns ...AggregateFunc) *OutputLogSelect {
	return olq.Select().Aggregate(fns...)
}

func (olq *OutputLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range olq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, olq); err != nil {
				return err
			}
		}
	}
	for _, f := range olq.ctx.Fields {
		if !outputlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if olq.path != nil {
		prev, err := olq.path(ctx)
		if err != nil {
			return err
		}
		olq.sql = prev
	}
	return nil
}

func (olq *OutputLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutputLog, error) {
	var (
		nodes = []*OutputLog{}
		_spec = olq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OutputLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OutputLog{config: olq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, olq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (olq *OutputLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := olq.querySpec()
	_spec.Node.Columns = olq.ctx.Fields
	if len(olq.ctx.Fields) > 0 {
		_spec.Unique = olq.ctx.Unique != nil && *olq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, olq.driver, _spec)
}

func (olq *OutputLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(outputlog.Table, outputlog.Columns, sqlgraph.NewFieldSpec(outputlog.FieldID, field.TypeInt64))
	_spec.From = olq.sql
	if unique := olq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if olq.path != nil {
		_spec.Unique = true
	}
	if fields := olq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outputlog.FieldID)
		for i := range fields {
			if fields[i] != outputlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := olq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := olq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := olq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := olq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (olq *OutputLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(olq.driver.Dialect())
	t1 := builder.Table(outputlog.Table)
	columns := olq.ctx.Fields
	if len(columns) == 0 {
		columns = outputlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if olq.sql != nil {
		selector = olq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if olq.ctx.Unique != nil && *olq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range olq.predicates {
		p(selector)
	}
	for _, p := range olq.order {
		p(selector)
	}
	if offset := olq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := olq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutputLogGroupBy is the group-by builder for OutputLog entities.
type OutputLogGroupBy struct {
	selector
	build *OutputLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (olgb *OutputLogGroupBy) Aggregate(fns ...AggregateFunc) *OutputLogGroupBy {
	olgb.fns = append(olgb.fns, fns...)
	return olgb
}

// Scan applies the selector query and scans the result into the given value.
func (olgb *OutputLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, olgb.build.ctx, "GroupBy")
	if err := olgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OutputLogQuery, *OutputLogGroupBy](ctx, olgb.build, olgb, olgb.build.inters, v)
}

func (olgb *OutputLogGroupBy) sqlScan(ctx context.Context, root *OutputLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(olgb.fns))
	for _, fn := range olgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*olgb.flds)+len(olgb.fns))
		for _, f := range *olgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*olgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := olgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OutputLogSelect is the builder for selecting fields of OutputLog entities.
type OutputLogSelect struct {
	*OutputLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ols *OutputLogSelect) Aggregate(fns ...AggregateFunc) *OutputLogSelect {
	ols.fns = append(ols.fns, fns...)
	return ols
}

// Scan applies the selector query and scans the result into the given value.
func (ols *OutputLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ols.ctx, "Select")
	if err := ols.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OutputLogQuery, *OutputLogSelect](ctx, ols.OutputLogQuery, ols, ols.inters, v)
}

func (ols *OutputLogSelect) sqlScan(ctx context.Context, root *OutputLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ols.fns))
	for _, fn := range ols.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ols.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ols.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}