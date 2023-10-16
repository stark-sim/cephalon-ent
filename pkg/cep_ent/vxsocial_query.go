// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/rechargeorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/transferorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/vxsocial"
)

// VXSocialQuery is the builder for querying VXSocial entities.
type VXSocialQuery struct {
	config
	ctx                *QueryContext
	order              []vxsocial.OrderOption
	inters             []Interceptor
	predicates         []predicate.VXSocial
	withUser           *UserQuery
	withRechargeOrders *RechargeOrderQuery
	withTransferOrders *TransferOrderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VXSocialQuery builder.
func (vsq *VXSocialQuery) Where(ps ...predicate.VXSocial) *VXSocialQuery {
	vsq.predicates = append(vsq.predicates, ps...)
	return vsq
}

// Limit the number of records to be returned by this query.
func (vsq *VXSocialQuery) Limit(limit int) *VXSocialQuery {
	vsq.ctx.Limit = &limit
	return vsq
}

// Offset to start from.
func (vsq *VXSocialQuery) Offset(offset int) *VXSocialQuery {
	vsq.ctx.Offset = &offset
	return vsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vsq *VXSocialQuery) Unique(unique bool) *VXSocialQuery {
	vsq.ctx.Unique = &unique
	return vsq
}

// Order specifies how the records should be ordered.
func (vsq *VXSocialQuery) Order(o ...vxsocial.OrderOption) *VXSocialQuery {
	vsq.order = append(vsq.order, o...)
	return vsq
}

// QueryUser chains the current query on the "user" edge.
func (vsq *VXSocialQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: vsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vxsocial.Table, vxsocial.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, vxsocial.UserTable, vxsocial.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(vsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRechargeOrders chains the current query on the "recharge_orders" edge.
func (vsq *VXSocialQuery) QueryRechargeOrders() *RechargeOrderQuery {
	query := (&RechargeOrderClient{config: vsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vxsocial.Table, vxsocial.FieldID, selector),
			sqlgraph.To(rechargeorder.Table, rechargeorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, vxsocial.RechargeOrdersTable, vxsocial.RechargeOrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(vsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransferOrders chains the current query on the "transfer_orders" edge.
func (vsq *VXSocialQuery) QueryTransferOrders() *TransferOrderQuery {
	query := (&TransferOrderClient{config: vsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vxsocial.Table, vxsocial.FieldID, selector),
			sqlgraph.To(transferorder.Table, transferorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, vxsocial.TransferOrdersTable, vxsocial.TransferOrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(vsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first VXSocial entity from the query.
// Returns a *NotFoundError when no VXSocial was found.
func (vsq *VXSocialQuery) First(ctx context.Context) (*VXSocial, error) {
	nodes, err := vsq.Limit(1).All(setContextOp(ctx, vsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vxsocial.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vsq *VXSocialQuery) FirstX(ctx context.Context) *VXSocial {
	node, err := vsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VXSocial ID from the query.
// Returns a *NotFoundError when no VXSocial ID was found.
func (vsq *VXSocialQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vsq.Limit(1).IDs(setContextOp(ctx, vsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vxsocial.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vsq *VXSocialQuery) FirstIDX(ctx context.Context) int64 {
	id, err := vsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VXSocial entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VXSocial entity is found.
// Returns a *NotFoundError when no VXSocial entities are found.
func (vsq *VXSocialQuery) Only(ctx context.Context) (*VXSocial, error) {
	nodes, err := vsq.Limit(2).All(setContextOp(ctx, vsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vxsocial.Label}
	default:
		return nil, &NotSingularError{vxsocial.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vsq *VXSocialQuery) OnlyX(ctx context.Context) *VXSocial {
	node, err := vsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VXSocial ID in the query.
// Returns a *NotSingularError when more than one VXSocial ID is found.
// Returns a *NotFoundError when no entities are found.
func (vsq *VXSocialQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vsq.Limit(2).IDs(setContextOp(ctx, vsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vxsocial.Label}
	default:
		err = &NotSingularError{vxsocial.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vsq *VXSocialQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := vsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VXSocials.
func (vsq *VXSocialQuery) All(ctx context.Context) ([]*VXSocial, error) {
	ctx = setContextOp(ctx, vsq.ctx, "All")
	if err := vsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VXSocial, *VXSocialQuery]()
	return withInterceptors[[]*VXSocial](ctx, vsq, qr, vsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vsq *VXSocialQuery) AllX(ctx context.Context) []*VXSocial {
	nodes, err := vsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VXSocial IDs.
func (vsq *VXSocialQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if vsq.ctx.Unique == nil && vsq.path != nil {
		vsq.Unique(true)
	}
	ctx = setContextOp(ctx, vsq.ctx, "IDs")
	if err = vsq.Select(vxsocial.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vsq *VXSocialQuery) IDsX(ctx context.Context) []int64 {
	ids, err := vsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vsq *VXSocialQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vsq.ctx, "Count")
	if err := vsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vsq, querierCount[*VXSocialQuery](), vsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vsq *VXSocialQuery) CountX(ctx context.Context) int {
	count, err := vsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vsq *VXSocialQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vsq.ctx, "Exist")
	switch _, err := vsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vsq *VXSocialQuery) ExistX(ctx context.Context) bool {
	exist, err := vsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VXSocialQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vsq *VXSocialQuery) Clone() *VXSocialQuery {
	if vsq == nil {
		return nil
	}
	return &VXSocialQuery{
		config:             vsq.config,
		ctx:                vsq.ctx.Clone(),
		order:              append([]vxsocial.OrderOption{}, vsq.order...),
		inters:             append([]Interceptor{}, vsq.inters...),
		predicates:         append([]predicate.VXSocial{}, vsq.predicates...),
		withUser:           vsq.withUser.Clone(),
		withRechargeOrders: vsq.withRechargeOrders.Clone(),
		withTransferOrders: vsq.withTransferOrders.Clone(),
		// clone intermediate query.
		sql:  vsq.sql.Clone(),
		path: vsq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (vsq *VXSocialQuery) WithUser(opts ...func(*UserQuery)) *VXSocialQuery {
	query := (&UserClient{config: vsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vsq.withUser = query
	return vsq
}

// WithRechargeOrders tells the query-builder to eager-load the nodes that are connected to
// the "recharge_orders" edge. The optional arguments are used to configure the query builder of the edge.
func (vsq *VXSocialQuery) WithRechargeOrders(opts ...func(*RechargeOrderQuery)) *VXSocialQuery {
	query := (&RechargeOrderClient{config: vsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vsq.withRechargeOrders = query
	return vsq
}

// WithTransferOrders tells the query-builder to eager-load the nodes that are connected to
// the "transfer_orders" edge. The optional arguments are used to configure the query builder of the edge.
func (vsq *VXSocialQuery) WithTransferOrders(opts ...func(*TransferOrderQuery)) *VXSocialQuery {
	query := (&TransferOrderClient{config: vsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vsq.withTransferOrders = query
	return vsq
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
//	client.VXSocial.Query().
//		GroupBy(vxsocial.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (vsq *VXSocialQuery) GroupBy(field string, fields ...string) *VXSocialGroupBy {
	vsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VXSocialGroupBy{build: vsq}
	grbuild.flds = &vsq.ctx.Fields
	grbuild.label = vxsocial.Label
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
//	client.VXSocial.Query().
//		Select(vxsocial.FieldCreatedBy).
//		Scan(ctx, &v)
func (vsq *VXSocialQuery) Select(fields ...string) *VXSocialSelect {
	vsq.ctx.Fields = append(vsq.ctx.Fields, fields...)
	sbuild := &VXSocialSelect{VXSocialQuery: vsq}
	sbuild.label = vxsocial.Label
	sbuild.flds, sbuild.scan = &vsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VXSocialSelect configured with the given aggregations.
func (vsq *VXSocialQuery) Aggregate(fns ...AggregateFunc) *VXSocialSelect {
	return vsq.Select().Aggregate(fns...)
}

func (vsq *VXSocialQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vsq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vsq); err != nil {
				return err
			}
		}
	}
	for _, f := range vsq.ctx.Fields {
		if !vxsocial.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if vsq.path != nil {
		prev, err := vsq.path(ctx)
		if err != nil {
			return err
		}
		vsq.sql = prev
	}
	return nil
}

func (vsq *VXSocialQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VXSocial, error) {
	var (
		nodes       = []*VXSocial{}
		_spec       = vsq.querySpec()
		loadedTypes = [3]bool{
			vsq.withUser != nil,
			vsq.withRechargeOrders != nil,
			vsq.withTransferOrders != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VXSocial).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VXSocial{config: vsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vsq.withUser; query != nil {
		if err := vsq.loadUser(ctx, query, nodes, nil,
			func(n *VXSocial, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := vsq.withRechargeOrders; query != nil {
		if err := vsq.loadRechargeOrders(ctx, query, nodes,
			func(n *VXSocial) { n.Edges.RechargeOrders = []*RechargeOrder{} },
			func(n *VXSocial, e *RechargeOrder) { n.Edges.RechargeOrders = append(n.Edges.RechargeOrders, e) }); err != nil {
			return nil, err
		}
	}
	if query := vsq.withTransferOrders; query != nil {
		if err := vsq.loadTransferOrders(ctx, query, nodes,
			func(n *VXSocial) { n.Edges.TransferOrders = []*TransferOrder{} },
			func(n *VXSocial, e *TransferOrder) { n.Edges.TransferOrders = append(n.Edges.TransferOrders, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vsq *VXSocialQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*VXSocial, init func(*VXSocial), assign func(*VXSocial, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*VXSocial)
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
func (vsq *VXSocialQuery) loadRechargeOrders(ctx context.Context, query *RechargeOrderQuery, nodes []*VXSocial, init func(*VXSocial), assign func(*VXSocial, *RechargeOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*VXSocial)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(rechargeorder.FieldSocialID)
	}
	query.Where(predicate.RechargeOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(vxsocial.RechargeOrdersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SocialID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "social_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vsq *VXSocialQuery) loadTransferOrders(ctx context.Context, query *TransferOrderQuery, nodes []*VXSocial, init func(*VXSocial), assign func(*VXSocial, *TransferOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*VXSocial)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(transferorder.FieldSocialID)
	}
	query.Where(predicate.TransferOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(vxsocial.TransferOrdersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SocialID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "social_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (vsq *VXSocialQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vsq.querySpec()
	_spec.Node.Columns = vsq.ctx.Fields
	if len(vsq.ctx.Fields) > 0 {
		_spec.Unique = vsq.ctx.Unique != nil && *vsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vsq.driver, _spec)
}

func (vsq *VXSocialQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(vxsocial.Table, vxsocial.Columns, sqlgraph.NewFieldSpec(vxsocial.FieldID, field.TypeInt64))
	_spec.From = vsq.sql
	if unique := vsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vsq.path != nil {
		_spec.Unique = true
	}
	if fields := vsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vxsocial.FieldID)
		for i := range fields {
			if fields[i] != vxsocial.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if vsq.withUser != nil {
			_spec.Node.AddColumnOnce(vxsocial.FieldUserID)
		}
	}
	if ps := vsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vsq *VXSocialQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vsq.driver.Dialect())
	t1 := builder.Table(vxsocial.Table)
	columns := vsq.ctx.Fields
	if len(columns) == 0 {
		columns = vxsocial.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vsq.sql != nil {
		selector = vsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vsq.ctx.Unique != nil && *vsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vsq.predicates {
		p(selector)
	}
	for _, p := range vsq.order {
		p(selector)
	}
	if offset := vsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VXSocialGroupBy is the group-by builder for VXSocial entities.
type VXSocialGroupBy struct {
	selector
	build *VXSocialQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vsgb *VXSocialGroupBy) Aggregate(fns ...AggregateFunc) *VXSocialGroupBy {
	vsgb.fns = append(vsgb.fns, fns...)
	return vsgb
}

// Scan applies the selector query and scans the result into the given value.
func (vsgb *VXSocialGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vsgb.build.ctx, "GroupBy")
	if err := vsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VXSocialQuery, *VXSocialGroupBy](ctx, vsgb.build, vsgb, vsgb.build.inters, v)
}

func (vsgb *VXSocialGroupBy) sqlScan(ctx context.Context, root *VXSocialQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vsgb.fns))
	for _, fn := range vsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vsgb.flds)+len(vsgb.fns))
		for _, f := range *vsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VXSocialSelect is the builder for selecting fields of VXSocial entities.
type VXSocialSelect struct {
	*VXSocialQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vss *VXSocialSelect) Aggregate(fns ...AggregateFunc) *VXSocialSelect {
	vss.fns = append(vss.fns, fns...)
	return vss
}

// Scan applies the selector query and scans the result into the given value.
func (vss *VXSocialSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vss.ctx, "Select")
	if err := vss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VXSocialQuery, *VXSocialSelect](ctx, vss.VXSocialQuery, vss, vss.inters, v)
}

func (vss *VXSocialSelect) sqlScan(ctx context.Context, root *VXSocialQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vss.fns))
	for _, fn := range vss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
