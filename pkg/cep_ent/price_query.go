// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/price"
)

// PriceQuery is the builder for querying Price entities.
type PriceQuery struct {
	config
	ctx        *QueryContext
	order      []price.OrderOption
	inters     []Interceptor
	predicates []predicate.Price
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PriceQuery builder.
func (pq *PriceQuery) Where(ps ...predicate.Price) *PriceQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PriceQuery) Limit(limit int) *PriceQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PriceQuery) Offset(offset int) *PriceQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PriceQuery) Unique(unique bool) *PriceQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PriceQuery) Order(o ...price.OrderOption) *PriceQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// First returns the first Price entity from the query.
// Returns a *NotFoundError when no Price was found.
func (pq *PriceQuery) First(ctx context.Context) (*Price, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{price.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PriceQuery) FirstX(ctx context.Context) *Price {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Price ID from the query.
// Returns a *NotFoundError when no Price ID was found.
func (pq *PriceQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{price.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PriceQuery) FirstIDX(ctx context.Context) int64 {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Price entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Price entity is found.
// Returns a *NotFoundError when no Price entities are found.
func (pq *PriceQuery) Only(ctx context.Context) (*Price, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{price.Label}
	default:
		return nil, &NotSingularError{price.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PriceQuery) OnlyX(ctx context.Context) *Price {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Price ID in the query.
// Returns a *NotSingularError when more than one Price ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PriceQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{price.Label}
	default:
		err = &NotSingularError{price.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PriceQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Prices.
func (pq *PriceQuery) All(ctx context.Context) ([]*Price, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Price, *PriceQuery]()
	return withInterceptors[[]*Price](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PriceQuery) AllX(ctx context.Context) []*Price {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Price IDs.
func (pq *PriceQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(price.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PriceQuery) IDsX(ctx context.Context) []int64 {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PriceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PriceQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PriceQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PriceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PriceQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PriceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PriceQuery) Clone() *PriceQuery {
	if pq == nil {
		return nil
	}
	return &PriceQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]price.OrderOption{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Price{}, pq.predicates...),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
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
//	client.Price.Query().
//		GroupBy(price.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (pq *PriceQuery) GroupBy(field string, fields ...string) *PriceGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PriceGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = price.Label
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
//	client.Price.Query().
//		Select(price.FieldCreatedBy).
//		Scan(ctx, &v)
func (pq *PriceQuery) Select(fields ...string) *PriceSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PriceSelect{PriceQuery: pq}
	sbuild.label = price.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PriceSelect configured with the given aggregations.
func (pq *PriceQuery) Aggregate(fns ...AggregateFunc) *PriceSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PriceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !price.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PriceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Price, error) {
	var (
		nodes = []*Price{}
		_spec = pq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Price).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Price{config: pq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pq *PriceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PriceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(price.Table, price.Columns, sqlgraph.NewFieldSpec(price.FieldID, field.TypeInt64))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, price.FieldID)
		for i := range fields {
			if fields[i] != price.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PriceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(price.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = price.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PriceGroupBy is the group-by builder for Price entities.
type PriceGroupBy struct {
	selector
	build *PriceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PriceGroupBy) Aggregate(fns ...AggregateFunc) *PriceGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PriceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PriceQuery, *PriceGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PriceGroupBy) sqlScan(ctx context.Context, root *PriceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PriceSelect is the builder for selecting fields of Price entities.
type PriceSelect struct {
	*PriceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PriceSelect) Aggregate(fns ...AggregateFunc) *PriceSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PriceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PriceQuery, *PriceSelect](ctx, ps.PriceQuery, ps, ps.inters, v)
}

func (ps *PriceSelect) sqlScan(ctx context.Context, root *PriceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
