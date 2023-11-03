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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/campaignorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/costbill"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/rechargeorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/vxsocial"
)

// RechargeOrderQuery is the builder for querying RechargeOrder entities.
type RechargeOrderQuery struct {
	config
	ctx               *QueryContext
	order             []rechargeorder.OrderOption
	inters            []Interceptor
	predicates        []predicate.RechargeOrder
	withUser          *UserQuery
	withCostBills     *CostBillQuery
	withVxSocial      *VXSocialQuery
	withCampaignOrder *CampaignOrderQuery
	modifiers         []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RechargeOrderQuery builder.
func (roq *RechargeOrderQuery) Where(ps ...predicate.RechargeOrder) *RechargeOrderQuery {
	roq.predicates = append(roq.predicates, ps...)
	return roq
}

// Limit the number of records to be returned by this query.
func (roq *RechargeOrderQuery) Limit(limit int) *RechargeOrderQuery {
	roq.ctx.Limit = &limit
	return roq
}

// Offset to start from.
func (roq *RechargeOrderQuery) Offset(offset int) *RechargeOrderQuery {
	roq.ctx.Offset = &offset
	return roq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (roq *RechargeOrderQuery) Unique(unique bool) *RechargeOrderQuery {
	roq.ctx.Unique = &unique
	return roq
}

// Order specifies how the records should be ordered.
func (roq *RechargeOrderQuery) Order(o ...rechargeorder.OrderOption) *RechargeOrderQuery {
	roq.order = append(roq.order, o...)
	return roq
}

// QueryUser chains the current query on the "user" edge.
func (roq *RechargeOrderQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: roq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := roq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := roq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rechargeorder.Table, rechargeorder.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rechargeorder.UserTable, rechargeorder.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(roq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCostBills chains the current query on the "cost_bills" edge.
func (roq *RechargeOrderQuery) QueryCostBills() *CostBillQuery {
	query := (&CostBillClient{config: roq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := roq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := roq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rechargeorder.Table, rechargeorder.FieldID, selector),
			sqlgraph.To(costbill.Table, costbill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, rechargeorder.CostBillsTable, rechargeorder.CostBillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(roq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVxSocial chains the current query on the "vx_social" edge.
func (roq *RechargeOrderQuery) QueryVxSocial() *VXSocialQuery {
	query := (&VXSocialClient{config: roq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := roq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := roq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rechargeorder.Table, rechargeorder.FieldID, selector),
			sqlgraph.To(vxsocial.Table, vxsocial.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rechargeorder.VxSocialTable, rechargeorder.VxSocialColumn),
		)
		fromU = sqlgraph.SetNeighbors(roq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCampaignOrder chains the current query on the "campaign_order" edge.
func (roq *RechargeOrderQuery) QueryCampaignOrder() *CampaignOrderQuery {
	query := (&CampaignOrderClient{config: roq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := roq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := roq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rechargeorder.Table, rechargeorder.FieldID, selector),
			sqlgraph.To(campaignorder.Table, campaignorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, rechargeorder.CampaignOrderTable, rechargeorder.CampaignOrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(roq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RechargeOrder entity from the query.
// Returns a *NotFoundError when no RechargeOrder was found.
func (roq *RechargeOrderQuery) First(ctx context.Context) (*RechargeOrder, error) {
	nodes, err := roq.Limit(1).All(setContextOp(ctx, roq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rechargeorder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (roq *RechargeOrderQuery) FirstX(ctx context.Context) *RechargeOrder {
	node, err := roq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RechargeOrder ID from the query.
// Returns a *NotFoundError when no RechargeOrder ID was found.
func (roq *RechargeOrderQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = roq.Limit(1).IDs(setContextOp(ctx, roq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rechargeorder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (roq *RechargeOrderQuery) FirstIDX(ctx context.Context) int64 {
	id, err := roq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RechargeOrder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RechargeOrder entity is found.
// Returns a *NotFoundError when no RechargeOrder entities are found.
func (roq *RechargeOrderQuery) Only(ctx context.Context) (*RechargeOrder, error) {
	nodes, err := roq.Limit(2).All(setContextOp(ctx, roq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rechargeorder.Label}
	default:
		return nil, &NotSingularError{rechargeorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (roq *RechargeOrderQuery) OnlyX(ctx context.Context) *RechargeOrder {
	node, err := roq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RechargeOrder ID in the query.
// Returns a *NotSingularError when more than one RechargeOrder ID is found.
// Returns a *NotFoundError when no entities are found.
func (roq *RechargeOrderQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = roq.Limit(2).IDs(setContextOp(ctx, roq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rechargeorder.Label}
	default:
		err = &NotSingularError{rechargeorder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (roq *RechargeOrderQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := roq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RechargeOrders.
func (roq *RechargeOrderQuery) All(ctx context.Context) ([]*RechargeOrder, error) {
	ctx = setContextOp(ctx, roq.ctx, "All")
	if err := roq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RechargeOrder, *RechargeOrderQuery]()
	return withInterceptors[[]*RechargeOrder](ctx, roq, qr, roq.inters)
}

// AllX is like All, but panics if an error occurs.
func (roq *RechargeOrderQuery) AllX(ctx context.Context) []*RechargeOrder {
	nodes, err := roq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RechargeOrder IDs.
func (roq *RechargeOrderQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if roq.ctx.Unique == nil && roq.path != nil {
		roq.Unique(true)
	}
	ctx = setContextOp(ctx, roq.ctx, "IDs")
	if err = roq.Select(rechargeorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (roq *RechargeOrderQuery) IDsX(ctx context.Context) []int64 {
	ids, err := roq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (roq *RechargeOrderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, roq.ctx, "Count")
	if err := roq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, roq, querierCount[*RechargeOrderQuery](), roq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (roq *RechargeOrderQuery) CountX(ctx context.Context) int {
	count, err := roq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (roq *RechargeOrderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, roq.ctx, "Exist")
	switch _, err := roq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (roq *RechargeOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := roq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RechargeOrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (roq *RechargeOrderQuery) Clone() *RechargeOrderQuery {
	if roq == nil {
		return nil
	}
	return &RechargeOrderQuery{
		config:            roq.config,
		ctx:               roq.ctx.Clone(),
		order:             append([]rechargeorder.OrderOption{}, roq.order...),
		inters:            append([]Interceptor{}, roq.inters...),
		predicates:        append([]predicate.RechargeOrder{}, roq.predicates...),
		withUser:          roq.withUser.Clone(),
		withCostBills:     roq.withCostBills.Clone(),
		withVxSocial:      roq.withVxSocial.Clone(),
		withCampaignOrder: roq.withCampaignOrder.Clone(),
		// clone intermediate query.
		sql:  roq.sql.Clone(),
		path: roq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (roq *RechargeOrderQuery) WithUser(opts ...func(*UserQuery)) *RechargeOrderQuery {
	query := (&UserClient{config: roq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	roq.withUser = query
	return roq
}

// WithCostBills tells the query-builder to eager-load the nodes that are connected to
// the "cost_bills" edge. The optional arguments are used to configure the query builder of the edge.
func (roq *RechargeOrderQuery) WithCostBills(opts ...func(*CostBillQuery)) *RechargeOrderQuery {
	query := (&CostBillClient{config: roq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	roq.withCostBills = query
	return roq
}

// WithVxSocial tells the query-builder to eager-load the nodes that are connected to
// the "vx_social" edge. The optional arguments are used to configure the query builder of the edge.
func (roq *RechargeOrderQuery) WithVxSocial(opts ...func(*VXSocialQuery)) *RechargeOrderQuery {
	query := (&VXSocialClient{config: roq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	roq.withVxSocial = query
	return roq
}

// WithCampaignOrder tells the query-builder to eager-load the nodes that are connected to
// the "campaign_order" edge. The optional arguments are used to configure the query builder of the edge.
func (roq *RechargeOrderQuery) WithCampaignOrder(opts ...func(*CampaignOrderQuery)) *RechargeOrderQuery {
	query := (&CampaignOrderClient{config: roq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	roq.withCampaignOrder = query
	return roq
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
//	client.RechargeOrder.Query().
//		GroupBy(rechargeorder.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (roq *RechargeOrderQuery) GroupBy(field string, fields ...string) *RechargeOrderGroupBy {
	roq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RechargeOrderGroupBy{build: roq}
	grbuild.flds = &roq.ctx.Fields
	grbuild.label = rechargeorder.Label
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
//	client.RechargeOrder.Query().
//		Select(rechargeorder.FieldCreatedBy).
//		Scan(ctx, &v)
func (roq *RechargeOrderQuery) Select(fields ...string) *RechargeOrderSelect {
	roq.ctx.Fields = append(roq.ctx.Fields, fields...)
	sbuild := &RechargeOrderSelect{RechargeOrderQuery: roq}
	sbuild.label = rechargeorder.Label
	sbuild.flds, sbuild.scan = &roq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RechargeOrderSelect configured with the given aggregations.
func (roq *RechargeOrderQuery) Aggregate(fns ...AggregateFunc) *RechargeOrderSelect {
	return roq.Select().Aggregate(fns...)
}

func (roq *RechargeOrderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range roq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, roq); err != nil {
				return err
			}
		}
	}
	for _, f := range roq.ctx.Fields {
		if !rechargeorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if roq.path != nil {
		prev, err := roq.path(ctx)
		if err != nil {
			return err
		}
		roq.sql = prev
	}
	return nil
}

func (roq *RechargeOrderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RechargeOrder, error) {
	var (
		nodes       = []*RechargeOrder{}
		_spec       = roq.querySpec()
		loadedTypes = [4]bool{
			roq.withUser != nil,
			roq.withCostBills != nil,
			roq.withVxSocial != nil,
			roq.withCampaignOrder != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RechargeOrder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RechargeOrder{config: roq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(roq.modifiers) > 0 {
		_spec.Modifiers = roq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, roq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := roq.withUser; query != nil {
		if err := roq.loadUser(ctx, query, nodes, nil,
			func(n *RechargeOrder, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := roq.withCostBills; query != nil {
		if err := roq.loadCostBills(ctx, query, nodes,
			func(n *RechargeOrder) { n.Edges.CostBills = []*CostBill{} },
			func(n *RechargeOrder, e *CostBill) { n.Edges.CostBills = append(n.Edges.CostBills, e) }); err != nil {
			return nil, err
		}
	}
	if query := roq.withVxSocial; query != nil {
		if err := roq.loadVxSocial(ctx, query, nodes, nil,
			func(n *RechargeOrder, e *VXSocial) { n.Edges.VxSocial = e }); err != nil {
			return nil, err
		}
	}
	if query := roq.withCampaignOrder; query != nil {
		if err := roq.loadCampaignOrder(ctx, query, nodes, nil,
			func(n *RechargeOrder, e *CampaignOrder) { n.Edges.CampaignOrder = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (roq *RechargeOrderQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*RechargeOrder, init func(*RechargeOrder), assign func(*RechargeOrder, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*RechargeOrder)
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
func (roq *RechargeOrderQuery) loadCostBills(ctx context.Context, query *CostBillQuery, nodes []*RechargeOrder, init func(*RechargeOrder), assign func(*RechargeOrder, *CostBill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*RechargeOrder)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(costbill.FieldReasonID)
	}
	query.Where(predicate.CostBill(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(rechargeorder.CostBillsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ReasonID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "reason_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (roq *RechargeOrderQuery) loadVxSocial(ctx context.Context, query *VXSocialQuery, nodes []*RechargeOrder, init func(*RechargeOrder), assign func(*RechargeOrder, *VXSocial)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*RechargeOrder)
	for i := range nodes {
		fk := nodes[i].SocialID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(vxsocial.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "social_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (roq *RechargeOrderQuery) loadCampaignOrder(ctx context.Context, query *CampaignOrderQuery, nodes []*RechargeOrder, init func(*RechargeOrder), assign func(*RechargeOrder, *CampaignOrder)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*RechargeOrder)
	for i := range nodes {
		if nodes[i].CampaignOrderID == nil {
			continue
		}
		fk := *nodes[i].CampaignOrderID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(campaignorder.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "campaign_order_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (roq *RechargeOrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := roq.querySpec()
	if len(roq.modifiers) > 0 {
		_spec.Modifiers = roq.modifiers
	}
	_spec.Node.Columns = roq.ctx.Fields
	if len(roq.ctx.Fields) > 0 {
		_spec.Unique = roq.ctx.Unique != nil && *roq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, roq.driver, _spec)
}

func (roq *RechargeOrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(rechargeorder.Table, rechargeorder.Columns, sqlgraph.NewFieldSpec(rechargeorder.FieldID, field.TypeInt64))
	_spec.From = roq.sql
	if unique := roq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if roq.path != nil {
		_spec.Unique = true
	}
	if fields := roq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rechargeorder.FieldID)
		for i := range fields {
			if fields[i] != rechargeorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if roq.withUser != nil {
			_spec.Node.AddColumnOnce(rechargeorder.FieldUserID)
		}
		if roq.withVxSocial != nil {
			_spec.Node.AddColumnOnce(rechargeorder.FieldSocialID)
		}
		if roq.withCampaignOrder != nil {
			_spec.Node.AddColumnOnce(rechargeorder.FieldCampaignOrderID)
		}
	}
	if ps := roq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := roq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := roq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := roq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (roq *RechargeOrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(roq.driver.Dialect())
	t1 := builder.Table(rechargeorder.Table)
	columns := roq.ctx.Fields
	if len(columns) == 0 {
		columns = rechargeorder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if roq.sql != nil {
		selector = roq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if roq.ctx.Unique != nil && *roq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range roq.modifiers {
		m(selector)
	}
	for _, p := range roq.predicates {
		p(selector)
	}
	for _, p := range roq.order {
		p(selector)
	}
	if offset := roq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := roq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (roq *RechargeOrderQuery) Modify(modifiers ...func(s *sql.Selector)) *RechargeOrderSelect {
	roq.modifiers = append(roq.modifiers, modifiers...)
	return roq.Select()
}

// RechargeOrderGroupBy is the group-by builder for RechargeOrder entities.
type RechargeOrderGroupBy struct {
	selector
	build *RechargeOrderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rogb *RechargeOrderGroupBy) Aggregate(fns ...AggregateFunc) *RechargeOrderGroupBy {
	rogb.fns = append(rogb.fns, fns...)
	return rogb
}

// Scan applies the selector query and scans the result into the given value.
func (rogb *RechargeOrderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rogb.build.ctx, "GroupBy")
	if err := rogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RechargeOrderQuery, *RechargeOrderGroupBy](ctx, rogb.build, rogb, rogb.build.inters, v)
}

func (rogb *RechargeOrderGroupBy) sqlScan(ctx context.Context, root *RechargeOrderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rogb.fns))
	for _, fn := range rogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rogb.flds)+len(rogb.fns))
		for _, f := range *rogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RechargeOrderSelect is the builder for selecting fields of RechargeOrder entities.
type RechargeOrderSelect struct {
	*RechargeOrderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ros *RechargeOrderSelect) Aggregate(fns ...AggregateFunc) *RechargeOrderSelect {
	ros.fns = append(ros.fns, fns...)
	return ros
}

// Scan applies the selector query and scans the result into the given value.
func (ros *RechargeOrderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ros.ctx, "Select")
	if err := ros.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RechargeOrderQuery, *RechargeOrderSelect](ctx, ros.RechargeOrderQuery, ros, ros.inters, v)
}

func (ros *RechargeOrderSelect) sqlScan(ctx context.Context, root *RechargeOrderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ros.fns))
	for _, fn := range ros.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ros.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ros.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ros *RechargeOrderSelect) Modify(modifiers ...func(s *sql.Selector)) *RechargeOrderSelect {
	ros.modifiers = append(ros.modifiers, modifiers...)
	return ros
}
