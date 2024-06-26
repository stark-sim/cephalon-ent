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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/campaign"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/campaignorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/costbill"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/rechargeorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// CampaignOrderQuery is the builder for querying CampaignOrder entities.
type CampaignOrderQuery struct {
	config
	ctx               *QueryContext
	order             []campaignorder.OrderOption
	inters            []Interceptor
	predicates        []predicate.CampaignOrder
	withUser          *UserQuery
	withCampaign      *CampaignQuery
	withCostBills     *CostBillQuery
	withRechargeOrder *RechargeOrderQuery
	modifiers         []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CampaignOrderQuery builder.
func (coq *CampaignOrderQuery) Where(ps ...predicate.CampaignOrder) *CampaignOrderQuery {
	coq.predicates = append(coq.predicates, ps...)
	return coq
}

// Limit the number of records to be returned by this query.
func (coq *CampaignOrderQuery) Limit(limit int) *CampaignOrderQuery {
	coq.ctx.Limit = &limit
	return coq
}

// Offset to start from.
func (coq *CampaignOrderQuery) Offset(offset int) *CampaignOrderQuery {
	coq.ctx.Offset = &offset
	return coq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (coq *CampaignOrderQuery) Unique(unique bool) *CampaignOrderQuery {
	coq.ctx.Unique = &unique
	return coq
}

// Order specifies how the records should be ordered.
func (coq *CampaignOrderQuery) Order(o ...campaignorder.OrderOption) *CampaignOrderQuery {
	coq.order = append(coq.order, o...)
	return coq
}

// QueryUser chains the current query on the "user" edge.
func (coq *CampaignOrderQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: coq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := coq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := coq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(campaignorder.Table, campaignorder.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, campaignorder.UserTable, campaignorder.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(coq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCampaign chains the current query on the "campaign" edge.
func (coq *CampaignOrderQuery) QueryCampaign() *CampaignQuery {
	query := (&CampaignClient{config: coq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := coq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := coq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(campaignorder.Table, campaignorder.FieldID, selector),
			sqlgraph.To(campaign.Table, campaign.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, campaignorder.CampaignTable, campaignorder.CampaignColumn),
		)
		fromU = sqlgraph.SetNeighbors(coq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCostBills chains the current query on the "cost_bills" edge.
func (coq *CampaignOrderQuery) QueryCostBills() *CostBillQuery {
	query := (&CostBillClient{config: coq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := coq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := coq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(campaignorder.Table, campaignorder.FieldID, selector),
			sqlgraph.To(costbill.Table, costbill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, campaignorder.CostBillsTable, campaignorder.CostBillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(coq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRechargeOrder chains the current query on the "recharge_order" edge.
func (coq *CampaignOrderQuery) QueryRechargeOrder() *RechargeOrderQuery {
	query := (&RechargeOrderClient{config: coq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := coq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := coq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(campaignorder.Table, campaignorder.FieldID, selector),
			sqlgraph.To(rechargeorder.Table, rechargeorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, campaignorder.RechargeOrderTable, campaignorder.RechargeOrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(coq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CampaignOrder entity from the query.
// Returns a *NotFoundError when no CampaignOrder was found.
func (coq *CampaignOrderQuery) First(ctx context.Context) (*CampaignOrder, error) {
	nodes, err := coq.Limit(1).All(setContextOp(ctx, coq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{campaignorder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (coq *CampaignOrderQuery) FirstX(ctx context.Context) *CampaignOrder {
	node, err := coq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CampaignOrder ID from the query.
// Returns a *NotFoundError when no CampaignOrder ID was found.
func (coq *CampaignOrderQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = coq.Limit(1).IDs(setContextOp(ctx, coq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{campaignorder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (coq *CampaignOrderQuery) FirstIDX(ctx context.Context) int64 {
	id, err := coq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CampaignOrder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CampaignOrder entity is found.
// Returns a *NotFoundError when no CampaignOrder entities are found.
func (coq *CampaignOrderQuery) Only(ctx context.Context) (*CampaignOrder, error) {
	nodes, err := coq.Limit(2).All(setContextOp(ctx, coq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{campaignorder.Label}
	default:
		return nil, &NotSingularError{campaignorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (coq *CampaignOrderQuery) OnlyX(ctx context.Context) *CampaignOrder {
	node, err := coq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CampaignOrder ID in the query.
// Returns a *NotSingularError when more than one CampaignOrder ID is found.
// Returns a *NotFoundError when no entities are found.
func (coq *CampaignOrderQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = coq.Limit(2).IDs(setContextOp(ctx, coq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{campaignorder.Label}
	default:
		err = &NotSingularError{campaignorder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (coq *CampaignOrderQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := coq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CampaignOrders.
func (coq *CampaignOrderQuery) All(ctx context.Context) ([]*CampaignOrder, error) {
	ctx = setContextOp(ctx, coq.ctx, "All")
	if err := coq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CampaignOrder, *CampaignOrderQuery]()
	return withInterceptors[[]*CampaignOrder](ctx, coq, qr, coq.inters)
}

// AllX is like All, but panics if an error occurs.
func (coq *CampaignOrderQuery) AllX(ctx context.Context) []*CampaignOrder {
	nodes, err := coq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CampaignOrder IDs.
func (coq *CampaignOrderQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if coq.ctx.Unique == nil && coq.path != nil {
		coq.Unique(true)
	}
	ctx = setContextOp(ctx, coq.ctx, "IDs")
	if err = coq.Select(campaignorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (coq *CampaignOrderQuery) IDsX(ctx context.Context) []int64 {
	ids, err := coq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (coq *CampaignOrderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, coq.ctx, "Count")
	if err := coq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, coq, querierCount[*CampaignOrderQuery](), coq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (coq *CampaignOrderQuery) CountX(ctx context.Context) int {
	count, err := coq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (coq *CampaignOrderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, coq.ctx, "Exist")
	switch _, err := coq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (coq *CampaignOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := coq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CampaignOrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (coq *CampaignOrderQuery) Clone() *CampaignOrderQuery {
	if coq == nil {
		return nil
	}
	return &CampaignOrderQuery{
		config:            coq.config,
		ctx:               coq.ctx.Clone(),
		order:             append([]campaignorder.OrderOption{}, coq.order...),
		inters:            append([]Interceptor{}, coq.inters...),
		predicates:        append([]predicate.CampaignOrder{}, coq.predicates...),
		withUser:          coq.withUser.Clone(),
		withCampaign:      coq.withCampaign.Clone(),
		withCostBills:     coq.withCostBills.Clone(),
		withRechargeOrder: coq.withRechargeOrder.Clone(),
		// clone intermediate query.
		sql:  coq.sql.Clone(),
		path: coq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (coq *CampaignOrderQuery) WithUser(opts ...func(*UserQuery)) *CampaignOrderQuery {
	query := (&UserClient{config: coq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	coq.withUser = query
	return coq
}

// WithCampaign tells the query-builder to eager-load the nodes that are connected to
// the "campaign" edge. The optional arguments are used to configure the query builder of the edge.
func (coq *CampaignOrderQuery) WithCampaign(opts ...func(*CampaignQuery)) *CampaignOrderQuery {
	query := (&CampaignClient{config: coq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	coq.withCampaign = query
	return coq
}

// WithCostBills tells the query-builder to eager-load the nodes that are connected to
// the "cost_bills" edge. The optional arguments are used to configure the query builder of the edge.
func (coq *CampaignOrderQuery) WithCostBills(opts ...func(*CostBillQuery)) *CampaignOrderQuery {
	query := (&CostBillClient{config: coq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	coq.withCostBills = query
	return coq
}

// WithRechargeOrder tells the query-builder to eager-load the nodes that are connected to
// the "recharge_order" edge. The optional arguments are used to configure the query builder of the edge.
func (coq *CampaignOrderQuery) WithRechargeOrder(opts ...func(*RechargeOrderQuery)) *CampaignOrderQuery {
	query := (&RechargeOrderClient{config: coq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	coq.withRechargeOrder = query
	return coq
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
//	client.CampaignOrder.Query().
//		GroupBy(campaignorder.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (coq *CampaignOrderQuery) GroupBy(field string, fields ...string) *CampaignOrderGroupBy {
	coq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CampaignOrderGroupBy{build: coq}
	grbuild.flds = &coq.ctx.Fields
	grbuild.label = campaignorder.Label
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
//	client.CampaignOrder.Query().
//		Select(campaignorder.FieldCreatedBy).
//		Scan(ctx, &v)
func (coq *CampaignOrderQuery) Select(fields ...string) *CampaignOrderSelect {
	coq.ctx.Fields = append(coq.ctx.Fields, fields...)
	sbuild := &CampaignOrderSelect{CampaignOrderQuery: coq}
	sbuild.label = campaignorder.Label
	sbuild.flds, sbuild.scan = &coq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CampaignOrderSelect configured with the given aggregations.
func (coq *CampaignOrderQuery) Aggregate(fns ...AggregateFunc) *CampaignOrderSelect {
	return coq.Select().Aggregate(fns...)
}

func (coq *CampaignOrderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range coq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, coq); err != nil {
				return err
			}
		}
	}
	for _, f := range coq.ctx.Fields {
		if !campaignorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if coq.path != nil {
		prev, err := coq.path(ctx)
		if err != nil {
			return err
		}
		coq.sql = prev
	}
	return nil
}

func (coq *CampaignOrderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CampaignOrder, error) {
	var (
		nodes       = []*CampaignOrder{}
		_spec       = coq.querySpec()
		loadedTypes = [4]bool{
			coq.withUser != nil,
			coq.withCampaign != nil,
			coq.withCostBills != nil,
			coq.withRechargeOrder != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CampaignOrder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CampaignOrder{config: coq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(coq.modifiers) > 0 {
		_spec.Modifiers = coq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, coq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := coq.withUser; query != nil {
		if err := coq.loadUser(ctx, query, nodes, nil,
			func(n *CampaignOrder, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := coq.withCampaign; query != nil {
		if err := coq.loadCampaign(ctx, query, nodes, nil,
			func(n *CampaignOrder, e *Campaign) { n.Edges.Campaign = e }); err != nil {
			return nil, err
		}
	}
	if query := coq.withCostBills; query != nil {
		if err := coq.loadCostBills(ctx, query, nodes,
			func(n *CampaignOrder) { n.Edges.CostBills = []*CostBill{} },
			func(n *CampaignOrder, e *CostBill) { n.Edges.CostBills = append(n.Edges.CostBills, e) }); err != nil {
			return nil, err
		}
	}
	if query := coq.withRechargeOrder; query != nil {
		if err := coq.loadRechargeOrder(ctx, query, nodes, nil,
			func(n *CampaignOrder, e *RechargeOrder) { n.Edges.RechargeOrder = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (coq *CampaignOrderQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*CampaignOrder, init func(*CampaignOrder), assign func(*CampaignOrder, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*CampaignOrder)
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
func (coq *CampaignOrderQuery) loadCampaign(ctx context.Context, query *CampaignQuery, nodes []*CampaignOrder, init func(*CampaignOrder), assign func(*CampaignOrder, *Campaign)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*CampaignOrder)
	for i := range nodes {
		fk := nodes[i].CampaignID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(campaign.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "campaign_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (coq *CampaignOrderQuery) loadCostBills(ctx context.Context, query *CostBillQuery, nodes []*CampaignOrder, init func(*CampaignOrder), assign func(*CampaignOrder, *CostBill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*CampaignOrder)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(costbill.FieldCampaignOrderID)
	}
	query.Where(predicate.CostBill(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(campaignorder.CostBillsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CampaignOrderID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "campaign_order_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (coq *CampaignOrderQuery) loadRechargeOrder(ctx context.Context, query *RechargeOrderQuery, nodes []*CampaignOrder, init func(*CampaignOrder), assign func(*CampaignOrder, *RechargeOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*CampaignOrder)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(rechargeorder.FieldCampaignOrderID)
	}
	query.Where(predicate.RechargeOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(campaignorder.RechargeOrderColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CampaignOrderID
		if fk == nil {
			return fmt.Errorf(`foreign-key "campaign_order_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "campaign_order_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (coq *CampaignOrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := coq.querySpec()
	if len(coq.modifiers) > 0 {
		_spec.Modifiers = coq.modifiers
	}
	_spec.Node.Columns = coq.ctx.Fields
	if len(coq.ctx.Fields) > 0 {
		_spec.Unique = coq.ctx.Unique != nil && *coq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, coq.driver, _spec)
}

func (coq *CampaignOrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(campaignorder.Table, campaignorder.Columns, sqlgraph.NewFieldSpec(campaignorder.FieldID, field.TypeInt64))
	_spec.From = coq.sql
	if unique := coq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if coq.path != nil {
		_spec.Unique = true
	}
	if fields := coq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, campaignorder.FieldID)
		for i := range fields {
			if fields[i] != campaignorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if coq.withUser != nil {
			_spec.Node.AddColumnOnce(campaignorder.FieldUserID)
		}
		if coq.withCampaign != nil {
			_spec.Node.AddColumnOnce(campaignorder.FieldCampaignID)
		}
	}
	if ps := coq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := coq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := coq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := coq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (coq *CampaignOrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(coq.driver.Dialect())
	t1 := builder.Table(campaignorder.Table)
	columns := coq.ctx.Fields
	if len(columns) == 0 {
		columns = campaignorder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if coq.sql != nil {
		selector = coq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if coq.ctx.Unique != nil && *coq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range coq.modifiers {
		m(selector)
	}
	for _, p := range coq.predicates {
		p(selector)
	}
	for _, p := range coq.order {
		p(selector)
	}
	if offset := coq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := coq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (coq *CampaignOrderQuery) Modify(modifiers ...func(s *sql.Selector)) *CampaignOrderSelect {
	coq.modifiers = append(coq.modifiers, modifiers...)
	return coq.Select()
}

// CampaignOrderGroupBy is the group-by builder for CampaignOrder entities.
type CampaignOrderGroupBy struct {
	selector
	build *CampaignOrderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cogb *CampaignOrderGroupBy) Aggregate(fns ...AggregateFunc) *CampaignOrderGroupBy {
	cogb.fns = append(cogb.fns, fns...)
	return cogb
}

// Scan applies the selector query and scans the result into the given value.
func (cogb *CampaignOrderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cogb.build.ctx, "GroupBy")
	if err := cogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CampaignOrderQuery, *CampaignOrderGroupBy](ctx, cogb.build, cogb, cogb.build.inters, v)
}

func (cogb *CampaignOrderGroupBy) sqlScan(ctx context.Context, root *CampaignOrderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cogb.fns))
	for _, fn := range cogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cogb.flds)+len(cogb.fns))
		for _, f := range *cogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CampaignOrderSelect is the builder for selecting fields of CampaignOrder entities.
type CampaignOrderSelect struct {
	*CampaignOrderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cos *CampaignOrderSelect) Aggregate(fns ...AggregateFunc) *CampaignOrderSelect {
	cos.fns = append(cos.fns, fns...)
	return cos
}

// Scan applies the selector query and scans the result into the given value.
func (cos *CampaignOrderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cos.ctx, "Select")
	if err := cos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CampaignOrderQuery, *CampaignOrderSelect](ctx, cos.CampaignOrderQuery, cos, cos.inters, v)
}

func (cos *CampaignOrderSelect) sqlScan(ctx context.Context, root *CampaignOrderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cos.fns))
	for _, fn := range cos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cos *CampaignOrderSelect) Modify(modifiers ...func(s *sql.Selector)) *CampaignOrderSelect {
	cos.modifiers = append(cos.modifiers, modifiers...)
	return cos
}
