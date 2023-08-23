// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/costaccount"
	"cephalon-ent/pkg/cep_ent/costbill"
	"cephalon-ent/pkg/cep_ent/missionconsumeorder"
	"cephalon-ent/pkg/cep_ent/predicate"
	"cephalon-ent/pkg/cep_ent/rechargeorder"
	"cephalon-ent/pkg/cep_ent/user"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CostBillQuery is the builder for querying CostBill entities.
type CostBillQuery struct {
	config
	ctx                     *QueryContext
	order                   []costbill.OrderOption
	inters                  []Interceptor
	predicates              []predicate.CostBill
	withUser                *UserQuery
	withCostAccount         *CostAccountQuery
	withRechargeOrder       *RechargeOrderQuery
	withMissionConsumeOrder *MissionConsumeOrderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CostBillQuery builder.
func (cbq *CostBillQuery) Where(ps ...predicate.CostBill) *CostBillQuery {
	cbq.predicates = append(cbq.predicates, ps...)
	return cbq
}

// Limit the number of records to be returned by this query.
func (cbq *CostBillQuery) Limit(limit int) *CostBillQuery {
	cbq.ctx.Limit = &limit
	return cbq
}

// Offset to start from.
func (cbq *CostBillQuery) Offset(offset int) *CostBillQuery {
	cbq.ctx.Offset = &offset
	return cbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cbq *CostBillQuery) Unique(unique bool) *CostBillQuery {
	cbq.ctx.Unique = &unique
	return cbq
}

// Order specifies how the records should be ordered.
func (cbq *CostBillQuery) Order(o ...costbill.OrderOption) *CostBillQuery {
	cbq.order = append(cbq.order, o...)
	return cbq
}

// QueryUser chains the current query on the "user" edge.
func (cbq *CostBillQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: cbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(costbill.Table, costbill.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, costbill.UserTable, costbill.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(cbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCostAccount chains the current query on the "cost_account" edge.
func (cbq *CostBillQuery) QueryCostAccount() *CostAccountQuery {
	query := (&CostAccountClient{config: cbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(costbill.Table, costbill.FieldID, selector),
			sqlgraph.To(costaccount.Table, costaccount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, costbill.CostAccountTable, costbill.CostAccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(cbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRechargeOrder chains the current query on the "recharge_order" edge.
func (cbq *CostBillQuery) QueryRechargeOrder() *RechargeOrderQuery {
	query := (&RechargeOrderClient{config: cbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(costbill.Table, costbill.FieldID, selector),
			sqlgraph.To(rechargeorder.Table, rechargeorder.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, costbill.RechargeOrderTable, costbill.RechargeOrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(cbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissionConsumeOrder chains the current query on the "mission_consume_order" edge.
func (cbq *CostBillQuery) QueryMissionConsumeOrder() *MissionConsumeOrderQuery {
	query := (&MissionConsumeOrderClient{config: cbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(costbill.Table, costbill.FieldID, selector),
			sqlgraph.To(missionconsumeorder.Table, missionconsumeorder.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, costbill.MissionConsumeOrderTable, costbill.MissionConsumeOrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(cbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CostBill entity from the query.
// Returns a *NotFoundError when no CostBill was found.
func (cbq *CostBillQuery) First(ctx context.Context) (*CostBill, error) {
	nodes, err := cbq.Limit(1).All(setContextOp(ctx, cbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{costbill.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cbq *CostBillQuery) FirstX(ctx context.Context) *CostBill {
	node, err := cbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CostBill ID from the query.
// Returns a *NotFoundError when no CostBill ID was found.
func (cbq *CostBillQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cbq.Limit(1).IDs(setContextOp(ctx, cbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{costbill.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cbq *CostBillQuery) FirstIDX(ctx context.Context) int64 {
	id, err := cbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CostBill entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CostBill entity is found.
// Returns a *NotFoundError when no CostBill entities are found.
func (cbq *CostBillQuery) Only(ctx context.Context) (*CostBill, error) {
	nodes, err := cbq.Limit(2).All(setContextOp(ctx, cbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{costbill.Label}
	default:
		return nil, &NotSingularError{costbill.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cbq *CostBillQuery) OnlyX(ctx context.Context) *CostBill {
	node, err := cbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CostBill ID in the query.
// Returns a *NotSingularError when more than one CostBill ID is found.
// Returns a *NotFoundError when no entities are found.
func (cbq *CostBillQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cbq.Limit(2).IDs(setContextOp(ctx, cbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{costbill.Label}
	default:
		err = &NotSingularError{costbill.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cbq *CostBillQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := cbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CostBills.
func (cbq *CostBillQuery) All(ctx context.Context) ([]*CostBill, error) {
	ctx = setContextOp(ctx, cbq.ctx, "All")
	if err := cbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CostBill, *CostBillQuery]()
	return withInterceptors[[]*CostBill](ctx, cbq, qr, cbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cbq *CostBillQuery) AllX(ctx context.Context) []*CostBill {
	nodes, err := cbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CostBill IDs.
func (cbq *CostBillQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if cbq.ctx.Unique == nil && cbq.path != nil {
		cbq.Unique(true)
	}
	ctx = setContextOp(ctx, cbq.ctx, "IDs")
	if err = cbq.Select(costbill.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cbq *CostBillQuery) IDsX(ctx context.Context) []int64 {
	ids, err := cbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cbq *CostBillQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cbq.ctx, "Count")
	if err := cbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cbq, querierCount[*CostBillQuery](), cbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cbq *CostBillQuery) CountX(ctx context.Context) int {
	count, err := cbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cbq *CostBillQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cbq.ctx, "Exist")
	switch _, err := cbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cbq *CostBillQuery) ExistX(ctx context.Context) bool {
	exist, err := cbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CostBillQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cbq *CostBillQuery) Clone() *CostBillQuery {
	if cbq == nil {
		return nil
	}
	return &CostBillQuery{
		config:                  cbq.config,
		ctx:                     cbq.ctx.Clone(),
		order:                   append([]costbill.OrderOption{}, cbq.order...),
		inters:                  append([]Interceptor{}, cbq.inters...),
		predicates:              append([]predicate.CostBill{}, cbq.predicates...),
		withUser:                cbq.withUser.Clone(),
		withCostAccount:         cbq.withCostAccount.Clone(),
		withRechargeOrder:       cbq.withRechargeOrder.Clone(),
		withMissionConsumeOrder: cbq.withMissionConsumeOrder.Clone(),
		// clone intermediate query.
		sql:  cbq.sql.Clone(),
		path: cbq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (cbq *CostBillQuery) WithUser(opts ...func(*UserQuery)) *CostBillQuery {
	query := (&UserClient{config: cbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cbq.withUser = query
	return cbq
}

// WithCostAccount tells the query-builder to eager-load the nodes that are connected to
// the "cost_account" edge. The optional arguments are used to configure the query builder of the edge.
func (cbq *CostBillQuery) WithCostAccount(opts ...func(*CostAccountQuery)) *CostBillQuery {
	query := (&CostAccountClient{config: cbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cbq.withCostAccount = query
	return cbq
}

// WithRechargeOrder tells the query-builder to eager-load the nodes that are connected to
// the "recharge_order" edge. The optional arguments are used to configure the query builder of the edge.
func (cbq *CostBillQuery) WithRechargeOrder(opts ...func(*RechargeOrderQuery)) *CostBillQuery {
	query := (&RechargeOrderClient{config: cbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cbq.withRechargeOrder = query
	return cbq
}

// WithMissionConsumeOrder tells the query-builder to eager-load the nodes that are connected to
// the "mission_consume_order" edge. The optional arguments are used to configure the query builder of the edge.
func (cbq *CostBillQuery) WithMissionConsumeOrder(opts ...func(*MissionConsumeOrderQuery)) *CostBillQuery {
	query := (&MissionConsumeOrderClient{config: cbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cbq.withMissionConsumeOrder = query
	return cbq
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
//	client.CostBill.Query().
//		GroupBy(costbill.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (cbq *CostBillQuery) GroupBy(field string, fields ...string) *CostBillGroupBy {
	cbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CostBillGroupBy{build: cbq}
	grbuild.flds = &cbq.ctx.Fields
	grbuild.label = costbill.Label
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
//	client.CostBill.Query().
//		Select(costbill.FieldCreatedBy).
//		Scan(ctx, &v)
func (cbq *CostBillQuery) Select(fields ...string) *CostBillSelect {
	cbq.ctx.Fields = append(cbq.ctx.Fields, fields...)
	sbuild := &CostBillSelect{CostBillQuery: cbq}
	sbuild.label = costbill.Label
	sbuild.flds, sbuild.scan = &cbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CostBillSelect configured with the given aggregations.
func (cbq *CostBillQuery) Aggregate(fns ...AggregateFunc) *CostBillSelect {
	return cbq.Select().Aggregate(fns...)
}

func (cbq *CostBillQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cbq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cbq); err != nil {
				return err
			}
		}
	}
	for _, f := range cbq.ctx.Fields {
		if !costbill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if cbq.path != nil {
		prev, err := cbq.path(ctx)
		if err != nil {
			return err
		}
		cbq.sql = prev
	}
	return nil
}

func (cbq *CostBillQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CostBill, error) {
	var (
		nodes       = []*CostBill{}
		_spec       = cbq.querySpec()
		loadedTypes = [4]bool{
			cbq.withUser != nil,
			cbq.withCostAccount != nil,
			cbq.withRechargeOrder != nil,
			cbq.withMissionConsumeOrder != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CostBill).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CostBill{config: cbq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cbq.withUser; query != nil {
		if err := cbq.loadUser(ctx, query, nodes, nil,
			func(n *CostBill, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := cbq.withCostAccount; query != nil {
		if err := cbq.loadCostAccount(ctx, query, nodes, nil,
			func(n *CostBill, e *CostAccount) { n.Edges.CostAccount = e }); err != nil {
			return nil, err
		}
	}
	if query := cbq.withRechargeOrder; query != nil {
		if err := cbq.loadRechargeOrder(ctx, query, nodes, nil,
			func(n *CostBill, e *RechargeOrder) { n.Edges.RechargeOrder = e }); err != nil {
			return nil, err
		}
	}
	if query := cbq.withMissionConsumeOrder; query != nil {
		if err := cbq.loadMissionConsumeOrder(ctx, query, nodes, nil,
			func(n *CostBill, e *MissionConsumeOrder) { n.Edges.MissionConsumeOrder = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cbq *CostBillQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*CostBill, init func(*CostBill), assign func(*CostBill, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*CostBill)
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
func (cbq *CostBillQuery) loadCostAccount(ctx context.Context, query *CostAccountQuery, nodes []*CostBill, init func(*CostBill), assign func(*CostBill, *CostAccount)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*CostBill)
	for i := range nodes {
		fk := nodes[i].CostAccountID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(costaccount.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cost_account_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cbq *CostBillQuery) loadRechargeOrder(ctx context.Context, query *RechargeOrderQuery, nodes []*CostBill, init func(*CostBill), assign func(*CostBill, *RechargeOrder)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*CostBill)
	for i := range nodes {
		fk := nodes[i].ReasonID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(rechargeorder.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "reason_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cbq *CostBillQuery) loadMissionConsumeOrder(ctx context.Context, query *MissionConsumeOrderQuery, nodes []*CostBill, init func(*CostBill), assign func(*CostBill, *MissionConsumeOrder)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*CostBill)
	for i := range nodes {
		fk := nodes[i].ReasonID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(missionconsumeorder.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "reason_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cbq *CostBillQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cbq.querySpec()
	_spec.Node.Columns = cbq.ctx.Fields
	if len(cbq.ctx.Fields) > 0 {
		_spec.Unique = cbq.ctx.Unique != nil && *cbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cbq.driver, _spec)
}

func (cbq *CostBillQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(costbill.Table, costbill.Columns, sqlgraph.NewFieldSpec(costbill.FieldID, field.TypeInt64))
	_spec.From = cbq.sql
	if unique := cbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cbq.path != nil {
		_spec.Unique = true
	}
	if fields := cbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, costbill.FieldID)
		for i := range fields {
			if fields[i] != costbill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cbq.withUser != nil {
			_spec.Node.AddColumnOnce(costbill.FieldUserID)
		}
		if cbq.withCostAccount != nil {
			_spec.Node.AddColumnOnce(costbill.FieldCostAccountID)
		}
		if cbq.withRechargeOrder != nil {
			_spec.Node.AddColumnOnce(costbill.FieldReasonID)
		}
		if cbq.withMissionConsumeOrder != nil {
			_spec.Node.AddColumnOnce(costbill.FieldReasonID)
		}
	}
	if ps := cbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cbq *CostBillQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cbq.driver.Dialect())
	t1 := builder.Table(costbill.Table)
	columns := cbq.ctx.Fields
	if len(columns) == 0 {
		columns = costbill.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cbq.sql != nil {
		selector = cbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cbq.ctx.Unique != nil && *cbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cbq.predicates {
		p(selector)
	}
	for _, p := range cbq.order {
		p(selector)
	}
	if offset := cbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CostBillGroupBy is the group-by builder for CostBill entities.
type CostBillGroupBy struct {
	selector
	build *CostBillQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cbgb *CostBillGroupBy) Aggregate(fns ...AggregateFunc) *CostBillGroupBy {
	cbgb.fns = append(cbgb.fns, fns...)
	return cbgb
}

// Scan applies the selector query and scans the result into the given value.
func (cbgb *CostBillGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cbgb.build.ctx, "GroupBy")
	if err := cbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CostBillQuery, *CostBillGroupBy](ctx, cbgb.build, cbgb, cbgb.build.inters, v)
}

func (cbgb *CostBillGroupBy) sqlScan(ctx context.Context, root *CostBillQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cbgb.fns))
	for _, fn := range cbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cbgb.flds)+len(cbgb.fns))
		for _, f := range *cbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CostBillSelect is the builder for selecting fields of CostBill entities.
type CostBillSelect struct {
	*CostBillQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cbs *CostBillSelect) Aggregate(fns ...AggregateFunc) *CostBillSelect {
	cbs.fns = append(cbs.fns, fns...)
	return cbs
}

// Scan applies the selector query and scans the result into the given value.
func (cbs *CostBillSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cbs.ctx, "Select")
	if err := cbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CostBillQuery, *CostBillSelect](ctx, cbs.CostBillQuery, cbs, cbs.inters, v)
}

func (cbs *CostBillSelect) sqlScan(ctx context.Context, root *CostBillQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cbs.fns))
	for _, fn := range cbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}