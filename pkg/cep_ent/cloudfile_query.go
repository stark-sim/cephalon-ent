// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/cloudfile"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// CloudFileQuery is the builder for querying CloudFile entities.
type CloudFileQuery struct {
	config
	ctx        *QueryContext
	order      []cloudfile.OrderOption
	inters     []Interceptor
	predicates []predicate.CloudFile
	withUser   *UserQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CloudFileQuery builder.
func (cfq *CloudFileQuery) Where(ps ...predicate.CloudFile) *CloudFileQuery {
	cfq.predicates = append(cfq.predicates, ps...)
	return cfq
}

// Limit the number of records to be returned by this query.
func (cfq *CloudFileQuery) Limit(limit int) *CloudFileQuery {
	cfq.ctx.Limit = &limit
	return cfq
}

// Offset to start from.
func (cfq *CloudFileQuery) Offset(offset int) *CloudFileQuery {
	cfq.ctx.Offset = &offset
	return cfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cfq *CloudFileQuery) Unique(unique bool) *CloudFileQuery {
	cfq.ctx.Unique = &unique
	return cfq
}

// Order specifies how the records should be ordered.
func (cfq *CloudFileQuery) Order(o ...cloudfile.OrderOption) *CloudFileQuery {
	cfq.order = append(cfq.order, o...)
	return cfq
}

// QueryUser chains the current query on the "user" edge.
func (cfq *CloudFileQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: cfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cloudfile.Table, cloudfile.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cloudfile.UserTable, cloudfile.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(cfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CloudFile entity from the query.
// Returns a *NotFoundError when no CloudFile was found.
func (cfq *CloudFileQuery) First(ctx context.Context) (*CloudFile, error) {
	nodes, err := cfq.Limit(1).All(setContextOp(ctx, cfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cloudfile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cfq *CloudFileQuery) FirstX(ctx context.Context) *CloudFile {
	node, err := cfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CloudFile ID from the query.
// Returns a *NotFoundError when no CloudFile ID was found.
func (cfq *CloudFileQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cfq.Limit(1).IDs(setContextOp(ctx, cfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cloudfile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cfq *CloudFileQuery) FirstIDX(ctx context.Context) int64 {
	id, err := cfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CloudFile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CloudFile entity is found.
// Returns a *NotFoundError when no CloudFile entities are found.
func (cfq *CloudFileQuery) Only(ctx context.Context) (*CloudFile, error) {
	nodes, err := cfq.Limit(2).All(setContextOp(ctx, cfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cloudfile.Label}
	default:
		return nil, &NotSingularError{cloudfile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cfq *CloudFileQuery) OnlyX(ctx context.Context) *CloudFile {
	node, err := cfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CloudFile ID in the query.
// Returns a *NotSingularError when more than one CloudFile ID is found.
// Returns a *NotFoundError when no entities are found.
func (cfq *CloudFileQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cfq.Limit(2).IDs(setContextOp(ctx, cfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cloudfile.Label}
	default:
		err = &NotSingularError{cloudfile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cfq *CloudFileQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := cfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CloudFiles.
func (cfq *CloudFileQuery) All(ctx context.Context) ([]*CloudFile, error) {
	ctx = setContextOp(ctx, cfq.ctx, "All")
	if err := cfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CloudFile, *CloudFileQuery]()
	return withInterceptors[[]*CloudFile](ctx, cfq, qr, cfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cfq *CloudFileQuery) AllX(ctx context.Context) []*CloudFile {
	nodes, err := cfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CloudFile IDs.
func (cfq *CloudFileQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if cfq.ctx.Unique == nil && cfq.path != nil {
		cfq.Unique(true)
	}
	ctx = setContextOp(ctx, cfq.ctx, "IDs")
	if err = cfq.Select(cloudfile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cfq *CloudFileQuery) IDsX(ctx context.Context) []int64 {
	ids, err := cfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cfq *CloudFileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cfq.ctx, "Count")
	if err := cfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cfq, querierCount[*CloudFileQuery](), cfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cfq *CloudFileQuery) CountX(ctx context.Context) int {
	count, err := cfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cfq *CloudFileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cfq.ctx, "Exist")
	switch _, err := cfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cfq *CloudFileQuery) ExistX(ctx context.Context) bool {
	exist, err := cfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CloudFileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cfq *CloudFileQuery) Clone() *CloudFileQuery {
	if cfq == nil {
		return nil
	}
	return &CloudFileQuery{
		config:     cfq.config,
		ctx:        cfq.ctx.Clone(),
		order:      append([]cloudfile.OrderOption{}, cfq.order...),
		inters:     append([]Interceptor{}, cfq.inters...),
		predicates: append([]predicate.CloudFile{}, cfq.predicates...),
		withUser:   cfq.withUser.Clone(),
		// clone intermediate query.
		sql:  cfq.sql.Clone(),
		path: cfq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (cfq *CloudFileQuery) WithUser(opts ...func(*UserQuery)) *CloudFileQuery {
	query := (&UserClient{config: cfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cfq.withUser = query
	return cfq
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
//	client.CloudFile.Query().
//		GroupBy(cloudfile.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (cfq *CloudFileQuery) GroupBy(field string, fields ...string) *CloudFileGroupBy {
	cfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CloudFileGroupBy{build: cfq}
	grbuild.flds = &cfq.ctx.Fields
	grbuild.label = cloudfile.Label
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
//	client.CloudFile.Query().
//		Select(cloudfile.FieldCreatedBy).
//		Scan(ctx, &v)
func (cfq *CloudFileQuery) Select(fields ...string) *CloudFileSelect {
	cfq.ctx.Fields = append(cfq.ctx.Fields, fields...)
	sbuild := &CloudFileSelect{CloudFileQuery: cfq}
	sbuild.label = cloudfile.Label
	sbuild.flds, sbuild.scan = &cfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CloudFileSelect configured with the given aggregations.
func (cfq *CloudFileQuery) Aggregate(fns ...AggregateFunc) *CloudFileSelect {
	return cfq.Select().Aggregate(fns...)
}

func (cfq *CloudFileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cfq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cfq); err != nil {
				return err
			}
		}
	}
	for _, f := range cfq.ctx.Fields {
		if !cloudfile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if cfq.path != nil {
		prev, err := cfq.path(ctx)
		if err != nil {
			return err
		}
		cfq.sql = prev
	}
	return nil
}

func (cfq *CloudFileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CloudFile, error) {
	var (
		nodes       = []*CloudFile{}
		_spec       = cfq.querySpec()
		loadedTypes = [1]bool{
			cfq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CloudFile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CloudFile{config: cfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cfq.modifiers) > 0 {
		_spec.Modifiers = cfq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cfq.withUser; query != nil {
		if err := cfq.loadUser(ctx, query, nodes, nil,
			func(n *CloudFile, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cfq *CloudFileQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*CloudFile, init func(*CloudFile), assign func(*CloudFile, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*CloudFile)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cfq *CloudFileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cfq.querySpec()
	if len(cfq.modifiers) > 0 {
		_spec.Modifiers = cfq.modifiers
	}
	_spec.Node.Columns = cfq.ctx.Fields
	if len(cfq.ctx.Fields) > 0 {
		_spec.Unique = cfq.ctx.Unique != nil && *cfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cfq.driver, _spec)
}

func (cfq *CloudFileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cloudfile.Table, cloudfile.Columns, sqlgraph.NewFieldSpec(cloudfile.FieldID, field.TypeInt64))
	_spec.From = cfq.sql
	if unique := cfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cfq.path != nil {
		_spec.Unique = true
	}
	if fields := cfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cloudfile.FieldID)
		for i := range fields {
			if fields[i] != cloudfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cfq.withUser != nil {
			_spec.Node.AddColumnOnce(cloudfile.FieldUserID)
		}
	}
	if ps := cfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cfq *CloudFileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cfq.driver.Dialect())
	t1 := builder.Table(cloudfile.Table)
	columns := cfq.ctx.Fields
	if len(columns) == 0 {
		columns = cloudfile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cfq.sql != nil {
		selector = cfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cfq.ctx.Unique != nil && *cfq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cfq.modifiers {
		m(selector)
	}
	for _, p := range cfq.predicates {
		p(selector)
	}
	for _, p := range cfq.order {
		p(selector)
	}
	if offset := cfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cfq *CloudFileQuery) Modify(modifiers ...func(s *sql.Selector)) *CloudFileSelect {
	cfq.modifiers = append(cfq.modifiers, modifiers...)
	return cfq.Select()
}

// CloudFileGroupBy is the group-by builder for CloudFile entities.
type CloudFileGroupBy struct {
	selector
	build *CloudFileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cfgb *CloudFileGroupBy) Aggregate(fns ...AggregateFunc) *CloudFileGroupBy {
	cfgb.fns = append(cfgb.fns, fns...)
	return cfgb
}

// Scan applies the selector query and scans the result into the given value.
func (cfgb *CloudFileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cfgb.build.ctx, "GroupBy")
	if err := cfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CloudFileQuery, *CloudFileGroupBy](ctx, cfgb.build, cfgb, cfgb.build.inters, v)
}

func (cfgb *CloudFileGroupBy) sqlScan(ctx context.Context, root *CloudFileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cfgb.fns))
	for _, fn := range cfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cfgb.flds)+len(cfgb.fns))
		for _, f := range *cfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CloudFileSelect is the builder for selecting fields of CloudFile entities.
type CloudFileSelect struct {
	*CloudFileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cfs *CloudFileSelect) Aggregate(fns ...AggregateFunc) *CloudFileSelect {
	cfs.fns = append(cfs.fns, fns...)
	return cfs
}

// Scan applies the selector query and scans the result into the given value.
func (cfs *CloudFileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cfs.ctx, "Select")
	if err := cfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CloudFileQuery, *CloudFileSelect](ctx, cfs.CloudFileQuery, cfs, cfs.inters, v)
}

func (cfs *CloudFileSelect) sqlScan(ctx context.Context, root *CloudFileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cfs.fns))
	for _, fn := range cfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cfs *CloudFileSelect) Modify(modifiers ...func(s *sql.Selector)) *CloudFileSelect {
	cfs.modifiers = append(cfs.modifiers, modifiers...)
	return cfs
}
