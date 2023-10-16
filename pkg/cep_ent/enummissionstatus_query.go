// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/enummissionstatus"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// EnumMissionStatusQuery is the builder for querying EnumMissionStatus entities.
type EnumMissionStatusQuery struct {
	config
	ctx        *QueryContext
	order      []enummissionstatus.OrderOption
	inters     []Interceptor
	predicates []predicate.EnumMissionStatus
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EnumMissionStatusQuery builder.
func (emsq *EnumMissionStatusQuery) Where(ps ...predicate.EnumMissionStatus) *EnumMissionStatusQuery {
	emsq.predicates = append(emsq.predicates, ps...)
	return emsq
}

// Limit the number of records to be returned by this query.
func (emsq *EnumMissionStatusQuery) Limit(limit int) *EnumMissionStatusQuery {
	emsq.ctx.Limit = &limit
	return emsq
}

// Offset to start from.
func (emsq *EnumMissionStatusQuery) Offset(offset int) *EnumMissionStatusQuery {
	emsq.ctx.Offset = &offset
	return emsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (emsq *EnumMissionStatusQuery) Unique(unique bool) *EnumMissionStatusQuery {
	emsq.ctx.Unique = &unique
	return emsq
}

// Order specifies how the records should be ordered.
func (emsq *EnumMissionStatusQuery) Order(o ...enummissionstatus.OrderOption) *EnumMissionStatusQuery {
	emsq.order = append(emsq.order, o...)
	return emsq
}

// First returns the first EnumMissionStatus entity from the query.
// Returns a *NotFoundError when no EnumMissionStatus was found.
func (emsq *EnumMissionStatusQuery) First(ctx context.Context) (*EnumMissionStatus, error) {
	nodes, err := emsq.Limit(1).All(setContextOp(ctx, emsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{enummissionstatus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (emsq *EnumMissionStatusQuery) FirstX(ctx context.Context) *EnumMissionStatus {
	node, err := emsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EnumMissionStatus ID from the query.
// Returns a *NotFoundError when no EnumMissionStatus ID was found.
func (emsq *EnumMissionStatusQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = emsq.Limit(1).IDs(setContextOp(ctx, emsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{enummissionstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (emsq *EnumMissionStatusQuery) FirstIDX(ctx context.Context) int64 {
	id, err := emsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EnumMissionStatus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EnumMissionStatus entity is found.
// Returns a *NotFoundError when no EnumMissionStatus entities are found.
func (emsq *EnumMissionStatusQuery) Only(ctx context.Context) (*EnumMissionStatus, error) {
	nodes, err := emsq.Limit(2).All(setContextOp(ctx, emsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{enummissionstatus.Label}
	default:
		return nil, &NotSingularError{enummissionstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (emsq *EnumMissionStatusQuery) OnlyX(ctx context.Context) *EnumMissionStatus {
	node, err := emsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EnumMissionStatus ID in the query.
// Returns a *NotSingularError when more than one EnumMissionStatus ID is found.
// Returns a *NotFoundError when no entities are found.
func (emsq *EnumMissionStatusQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = emsq.Limit(2).IDs(setContextOp(ctx, emsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{enummissionstatus.Label}
	default:
		err = &NotSingularError{enummissionstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (emsq *EnumMissionStatusQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := emsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EnumMissionStatusSlice.
func (emsq *EnumMissionStatusQuery) All(ctx context.Context) ([]*EnumMissionStatus, error) {
	ctx = setContextOp(ctx, emsq.ctx, "All")
	if err := emsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EnumMissionStatus, *EnumMissionStatusQuery]()
	return withInterceptors[[]*EnumMissionStatus](ctx, emsq, qr, emsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (emsq *EnumMissionStatusQuery) AllX(ctx context.Context) []*EnumMissionStatus {
	nodes, err := emsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EnumMissionStatus IDs.
func (emsq *EnumMissionStatusQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if emsq.ctx.Unique == nil && emsq.path != nil {
		emsq.Unique(true)
	}
	ctx = setContextOp(ctx, emsq.ctx, "IDs")
	if err = emsq.Select(enummissionstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (emsq *EnumMissionStatusQuery) IDsX(ctx context.Context) []int64 {
	ids, err := emsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (emsq *EnumMissionStatusQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, emsq.ctx, "Count")
	if err := emsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, emsq, querierCount[*EnumMissionStatusQuery](), emsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (emsq *EnumMissionStatusQuery) CountX(ctx context.Context) int {
	count, err := emsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (emsq *EnumMissionStatusQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, emsq.ctx, "Exist")
	switch _, err := emsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (emsq *EnumMissionStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := emsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EnumMissionStatusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (emsq *EnumMissionStatusQuery) Clone() *EnumMissionStatusQuery {
	if emsq == nil {
		return nil
	}
	return &EnumMissionStatusQuery{
		config:     emsq.config,
		ctx:        emsq.ctx.Clone(),
		order:      append([]enummissionstatus.OrderOption{}, emsq.order...),
		inters:     append([]Interceptor{}, emsq.inters...),
		predicates: append([]predicate.EnumMissionStatus{}, emsq.predicates...),
		// clone intermediate query.
		sql:  emsq.sql.Clone(),
		path: emsq.path,
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
//	client.EnumMissionStatus.Query().
//		GroupBy(enummissionstatus.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (emsq *EnumMissionStatusQuery) GroupBy(field string, fields ...string) *EnumMissionStatusGroupBy {
	emsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EnumMissionStatusGroupBy{build: emsq}
	grbuild.flds = &emsq.ctx.Fields
	grbuild.label = enummissionstatus.Label
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
//	client.EnumMissionStatus.Query().
//		Select(enummissionstatus.FieldCreatedBy).
//		Scan(ctx, &v)
func (emsq *EnumMissionStatusQuery) Select(fields ...string) *EnumMissionStatusSelect {
	emsq.ctx.Fields = append(emsq.ctx.Fields, fields...)
	sbuild := &EnumMissionStatusSelect{EnumMissionStatusQuery: emsq}
	sbuild.label = enummissionstatus.Label
	sbuild.flds, sbuild.scan = &emsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EnumMissionStatusSelect configured with the given aggregations.
func (emsq *EnumMissionStatusQuery) Aggregate(fns ...AggregateFunc) *EnumMissionStatusSelect {
	return emsq.Select().Aggregate(fns...)
}

func (emsq *EnumMissionStatusQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range emsq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, emsq); err != nil {
				return err
			}
		}
	}
	for _, f := range emsq.ctx.Fields {
		if !enummissionstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if emsq.path != nil {
		prev, err := emsq.path(ctx)
		if err != nil {
			return err
		}
		emsq.sql = prev
	}
	return nil
}

func (emsq *EnumMissionStatusQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EnumMissionStatus, error) {
	var (
		nodes = []*EnumMissionStatus{}
		_spec = emsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EnumMissionStatus).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EnumMissionStatus{config: emsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, emsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (emsq *EnumMissionStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := emsq.querySpec()
	_spec.Node.Columns = emsq.ctx.Fields
	if len(emsq.ctx.Fields) > 0 {
		_spec.Unique = emsq.ctx.Unique != nil && *emsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, emsq.driver, _spec)
}

func (emsq *EnumMissionStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(enummissionstatus.Table, enummissionstatus.Columns, sqlgraph.NewFieldSpec(enummissionstatus.FieldID, field.TypeInt64))
	_spec.From = emsq.sql
	if unique := emsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if emsq.path != nil {
		_spec.Unique = true
	}
	if fields := emsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, enummissionstatus.FieldID)
		for i := range fields {
			if fields[i] != enummissionstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := emsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := emsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := emsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := emsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (emsq *EnumMissionStatusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(emsq.driver.Dialect())
	t1 := builder.Table(enummissionstatus.Table)
	columns := emsq.ctx.Fields
	if len(columns) == 0 {
		columns = enummissionstatus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if emsq.sql != nil {
		selector = emsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if emsq.ctx.Unique != nil && *emsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range emsq.predicates {
		p(selector)
	}
	for _, p := range emsq.order {
		p(selector)
	}
	if offset := emsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := emsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EnumMissionStatusGroupBy is the group-by builder for EnumMissionStatus entities.
type EnumMissionStatusGroupBy struct {
	selector
	build *EnumMissionStatusQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (emsgb *EnumMissionStatusGroupBy) Aggregate(fns ...AggregateFunc) *EnumMissionStatusGroupBy {
	emsgb.fns = append(emsgb.fns, fns...)
	return emsgb
}

// Scan applies the selector query and scans the result into the given value.
func (emsgb *EnumMissionStatusGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, emsgb.build.ctx, "GroupBy")
	if err := emsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnumMissionStatusQuery, *EnumMissionStatusGroupBy](ctx, emsgb.build, emsgb, emsgb.build.inters, v)
}

func (emsgb *EnumMissionStatusGroupBy) sqlScan(ctx context.Context, root *EnumMissionStatusQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(emsgb.fns))
	for _, fn := range emsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*emsgb.flds)+len(emsgb.fns))
		for _, f := range *emsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*emsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := emsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EnumMissionStatusSelect is the builder for selecting fields of EnumMissionStatus entities.
type EnumMissionStatusSelect struct {
	*EnumMissionStatusQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (emss *EnumMissionStatusSelect) Aggregate(fns ...AggregateFunc) *EnumMissionStatusSelect {
	emss.fns = append(emss.fns, fns...)
	return emss
}

// Scan applies the selector query and scans the result into the given value.
func (emss *EnumMissionStatusSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, emss.ctx, "Select")
	if err := emss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnumMissionStatusQuery, *EnumMissionStatusSelect](ctx, emss.EnumMissionStatusQuery, emss, emss.inters, v)
}

func (emss *EnumMissionStatusSelect) sqlScan(ctx context.Context, root *EnumMissionStatusQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(emss.fns))
	for _, fn := range emss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*emss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := emss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
