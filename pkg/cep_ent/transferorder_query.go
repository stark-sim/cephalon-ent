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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/bill"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/symbol"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/transferorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/vxsocial"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/withdrawrecord"
)

// TransferOrderQuery is the builder for querying TransferOrder entities.
type TransferOrderQuery struct {
	config
	ctx                *QueryContext
	order              []transferorder.OrderOption
	inters             []Interceptor
	predicates         []predicate.TransferOrder
	withSourceUser     *UserQuery
	withTargetUser     *UserQuery
	withBills          *BillQuery
	withVxSocial       *VXSocialQuery
	withSymbol         *SymbolQuery
	withWithdrawRecord *WithdrawRecordQuery
	modifiers          []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TransferOrderQuery builder.
func (toq *TransferOrderQuery) Where(ps ...predicate.TransferOrder) *TransferOrderQuery {
	toq.predicates = append(toq.predicates, ps...)
	return toq
}

// Limit the number of records to be returned by this query.
func (toq *TransferOrderQuery) Limit(limit int) *TransferOrderQuery {
	toq.ctx.Limit = &limit
	return toq
}

// Offset to start from.
func (toq *TransferOrderQuery) Offset(offset int) *TransferOrderQuery {
	toq.ctx.Offset = &offset
	return toq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (toq *TransferOrderQuery) Unique(unique bool) *TransferOrderQuery {
	toq.ctx.Unique = &unique
	return toq
}

// Order specifies how the records should be ordered.
func (toq *TransferOrderQuery) Order(o ...transferorder.OrderOption) *TransferOrderQuery {
	toq.order = append(toq.order, o...)
	return toq
}

// QuerySourceUser chains the current query on the "source_user" edge.
func (toq *TransferOrderQuery) QuerySourceUser() *UserQuery {
	query := (&UserClient{config: toq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := toq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := toq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transferorder.Table, transferorder.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transferorder.SourceUserTable, transferorder.SourceUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(toq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTargetUser chains the current query on the "target_user" edge.
func (toq *TransferOrderQuery) QueryTargetUser() *UserQuery {
	query := (&UserClient{config: toq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := toq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := toq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transferorder.Table, transferorder.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transferorder.TargetUserTable, transferorder.TargetUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(toq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBills chains the current query on the "bills" edge.
func (toq *TransferOrderQuery) QueryBills() *BillQuery {
	query := (&BillClient{config: toq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := toq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := toq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transferorder.Table, transferorder.FieldID, selector),
			sqlgraph.To(bill.Table, bill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, transferorder.BillsTable, transferorder.BillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(toq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVxSocial chains the current query on the "vx_social" edge.
func (toq *TransferOrderQuery) QueryVxSocial() *VXSocialQuery {
	query := (&VXSocialClient{config: toq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := toq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := toq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transferorder.Table, transferorder.FieldID, selector),
			sqlgraph.To(vxsocial.Table, vxsocial.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transferorder.VxSocialTable, transferorder.VxSocialColumn),
		)
		fromU = sqlgraph.SetNeighbors(toq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySymbol chains the current query on the "symbol" edge.
func (toq *TransferOrderQuery) QuerySymbol() *SymbolQuery {
	query := (&SymbolClient{config: toq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := toq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := toq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transferorder.Table, transferorder.FieldID, selector),
			sqlgraph.To(symbol.Table, symbol.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transferorder.SymbolTable, transferorder.SymbolColumn),
		)
		fromU = sqlgraph.SetNeighbors(toq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWithdrawRecord chains the current query on the "withdraw_record" edge.
func (toq *TransferOrderQuery) QueryWithdrawRecord() *WithdrawRecordQuery {
	query := (&WithdrawRecordClient{config: toq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := toq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := toq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transferorder.Table, transferorder.FieldID, selector),
			sqlgraph.To(withdrawrecord.Table, withdrawrecord.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, transferorder.WithdrawRecordTable, transferorder.WithdrawRecordColumn),
		)
		fromU = sqlgraph.SetNeighbors(toq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TransferOrder entity from the query.
// Returns a *NotFoundError when no TransferOrder was found.
func (toq *TransferOrderQuery) First(ctx context.Context) (*TransferOrder, error) {
	nodes, err := toq.Limit(1).All(setContextOp(ctx, toq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{transferorder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (toq *TransferOrderQuery) FirstX(ctx context.Context) *TransferOrder {
	node, err := toq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TransferOrder ID from the query.
// Returns a *NotFoundError when no TransferOrder ID was found.
func (toq *TransferOrderQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = toq.Limit(1).IDs(setContextOp(ctx, toq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{transferorder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (toq *TransferOrderQuery) FirstIDX(ctx context.Context) int64 {
	id, err := toq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TransferOrder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TransferOrder entity is found.
// Returns a *NotFoundError when no TransferOrder entities are found.
func (toq *TransferOrderQuery) Only(ctx context.Context) (*TransferOrder, error) {
	nodes, err := toq.Limit(2).All(setContextOp(ctx, toq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{transferorder.Label}
	default:
		return nil, &NotSingularError{transferorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (toq *TransferOrderQuery) OnlyX(ctx context.Context) *TransferOrder {
	node, err := toq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TransferOrder ID in the query.
// Returns a *NotSingularError when more than one TransferOrder ID is found.
// Returns a *NotFoundError when no entities are found.
func (toq *TransferOrderQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = toq.Limit(2).IDs(setContextOp(ctx, toq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{transferorder.Label}
	default:
		err = &NotSingularError{transferorder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (toq *TransferOrderQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := toq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TransferOrders.
func (toq *TransferOrderQuery) All(ctx context.Context) ([]*TransferOrder, error) {
	ctx = setContextOp(ctx, toq.ctx, "All")
	if err := toq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TransferOrder, *TransferOrderQuery]()
	return withInterceptors[[]*TransferOrder](ctx, toq, qr, toq.inters)
}

// AllX is like All, but panics if an error occurs.
func (toq *TransferOrderQuery) AllX(ctx context.Context) []*TransferOrder {
	nodes, err := toq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TransferOrder IDs.
func (toq *TransferOrderQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if toq.ctx.Unique == nil && toq.path != nil {
		toq.Unique(true)
	}
	ctx = setContextOp(ctx, toq.ctx, "IDs")
	if err = toq.Select(transferorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (toq *TransferOrderQuery) IDsX(ctx context.Context) []int64 {
	ids, err := toq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (toq *TransferOrderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, toq.ctx, "Count")
	if err := toq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, toq, querierCount[*TransferOrderQuery](), toq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (toq *TransferOrderQuery) CountX(ctx context.Context) int {
	count, err := toq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (toq *TransferOrderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, toq.ctx, "Exist")
	switch _, err := toq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (toq *TransferOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := toq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TransferOrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (toq *TransferOrderQuery) Clone() *TransferOrderQuery {
	if toq == nil {
		return nil
	}
	return &TransferOrderQuery{
		config:             toq.config,
		ctx:                toq.ctx.Clone(),
		order:              append([]transferorder.OrderOption{}, toq.order...),
		inters:             append([]Interceptor{}, toq.inters...),
		predicates:         append([]predicate.TransferOrder{}, toq.predicates...),
		withSourceUser:     toq.withSourceUser.Clone(),
		withTargetUser:     toq.withTargetUser.Clone(),
		withBills:          toq.withBills.Clone(),
		withVxSocial:       toq.withVxSocial.Clone(),
		withSymbol:         toq.withSymbol.Clone(),
		withWithdrawRecord: toq.withWithdrawRecord.Clone(),
		// clone intermediate query.
		sql:  toq.sql.Clone(),
		path: toq.path,
	}
}

// WithSourceUser tells the query-builder to eager-load the nodes that are connected to
// the "source_user" edge. The optional arguments are used to configure the query builder of the edge.
func (toq *TransferOrderQuery) WithSourceUser(opts ...func(*UserQuery)) *TransferOrderQuery {
	query := (&UserClient{config: toq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	toq.withSourceUser = query
	return toq
}

// WithTargetUser tells the query-builder to eager-load the nodes that are connected to
// the "target_user" edge. The optional arguments are used to configure the query builder of the edge.
func (toq *TransferOrderQuery) WithTargetUser(opts ...func(*UserQuery)) *TransferOrderQuery {
	query := (&UserClient{config: toq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	toq.withTargetUser = query
	return toq
}

// WithBills tells the query-builder to eager-load the nodes that are connected to
// the "bills" edge. The optional arguments are used to configure the query builder of the edge.
func (toq *TransferOrderQuery) WithBills(opts ...func(*BillQuery)) *TransferOrderQuery {
	query := (&BillClient{config: toq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	toq.withBills = query
	return toq
}

// WithVxSocial tells the query-builder to eager-load the nodes that are connected to
// the "vx_social" edge. The optional arguments are used to configure the query builder of the edge.
func (toq *TransferOrderQuery) WithVxSocial(opts ...func(*VXSocialQuery)) *TransferOrderQuery {
	query := (&VXSocialClient{config: toq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	toq.withVxSocial = query
	return toq
}

// WithSymbol tells the query-builder to eager-load the nodes that are connected to
// the "symbol" edge. The optional arguments are used to configure the query builder of the edge.
func (toq *TransferOrderQuery) WithSymbol(opts ...func(*SymbolQuery)) *TransferOrderQuery {
	query := (&SymbolClient{config: toq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	toq.withSymbol = query
	return toq
}

// WithWithdrawRecord tells the query-builder to eager-load the nodes that are connected to
// the "withdraw_record" edge. The optional arguments are used to configure the query builder of the edge.
func (toq *TransferOrderQuery) WithWithdrawRecord(opts ...func(*WithdrawRecordQuery)) *TransferOrderQuery {
	query := (&WithdrawRecordClient{config: toq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	toq.withWithdrawRecord = query
	return toq
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
//	client.TransferOrder.Query().
//		GroupBy(transferorder.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (toq *TransferOrderQuery) GroupBy(field string, fields ...string) *TransferOrderGroupBy {
	toq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TransferOrderGroupBy{build: toq}
	grbuild.flds = &toq.ctx.Fields
	grbuild.label = transferorder.Label
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
//	client.TransferOrder.Query().
//		Select(transferorder.FieldCreatedBy).
//		Scan(ctx, &v)
func (toq *TransferOrderQuery) Select(fields ...string) *TransferOrderSelect {
	toq.ctx.Fields = append(toq.ctx.Fields, fields...)
	sbuild := &TransferOrderSelect{TransferOrderQuery: toq}
	sbuild.label = transferorder.Label
	sbuild.flds, sbuild.scan = &toq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TransferOrderSelect configured with the given aggregations.
func (toq *TransferOrderQuery) Aggregate(fns ...AggregateFunc) *TransferOrderSelect {
	return toq.Select().Aggregate(fns...)
}

func (toq *TransferOrderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range toq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, toq); err != nil {
				return err
			}
		}
	}
	for _, f := range toq.ctx.Fields {
		if !transferorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if toq.path != nil {
		prev, err := toq.path(ctx)
		if err != nil {
			return err
		}
		toq.sql = prev
	}
	return nil
}

func (toq *TransferOrderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TransferOrder, error) {
	var (
		nodes       = []*TransferOrder{}
		_spec       = toq.querySpec()
		loadedTypes = [6]bool{
			toq.withSourceUser != nil,
			toq.withTargetUser != nil,
			toq.withBills != nil,
			toq.withVxSocial != nil,
			toq.withSymbol != nil,
			toq.withWithdrawRecord != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TransferOrder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TransferOrder{config: toq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(toq.modifiers) > 0 {
		_spec.Modifiers = toq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, toq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := toq.withSourceUser; query != nil {
		if err := toq.loadSourceUser(ctx, query, nodes, nil,
			func(n *TransferOrder, e *User) { n.Edges.SourceUser = e }); err != nil {
			return nil, err
		}
	}
	if query := toq.withTargetUser; query != nil {
		if err := toq.loadTargetUser(ctx, query, nodes, nil,
			func(n *TransferOrder, e *User) { n.Edges.TargetUser = e }); err != nil {
			return nil, err
		}
	}
	if query := toq.withBills; query != nil {
		if err := toq.loadBills(ctx, query, nodes,
			func(n *TransferOrder) { n.Edges.Bills = []*Bill{} },
			func(n *TransferOrder, e *Bill) { n.Edges.Bills = append(n.Edges.Bills, e) }); err != nil {
			return nil, err
		}
	}
	if query := toq.withVxSocial; query != nil {
		if err := toq.loadVxSocial(ctx, query, nodes, nil,
			func(n *TransferOrder, e *VXSocial) { n.Edges.VxSocial = e }); err != nil {
			return nil, err
		}
	}
	if query := toq.withSymbol; query != nil {
		if err := toq.loadSymbol(ctx, query, nodes, nil,
			func(n *TransferOrder, e *Symbol) { n.Edges.Symbol = e }); err != nil {
			return nil, err
		}
	}
	if query := toq.withWithdrawRecord; query != nil {
		if err := toq.loadWithdrawRecord(ctx, query, nodes, nil,
			func(n *TransferOrder, e *WithdrawRecord) { n.Edges.WithdrawRecord = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (toq *TransferOrderQuery) loadSourceUser(ctx context.Context, query *UserQuery, nodes []*TransferOrder, init func(*TransferOrder), assign func(*TransferOrder, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*TransferOrder)
	for i := range nodes {
		fk := nodes[i].SourceUserID
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
			return fmt.Errorf(`unexpected foreign-key "source_user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (toq *TransferOrderQuery) loadTargetUser(ctx context.Context, query *UserQuery, nodes []*TransferOrder, init func(*TransferOrder), assign func(*TransferOrder, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*TransferOrder)
	for i := range nodes {
		fk := nodes[i].TargetUserID
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
			return fmt.Errorf(`unexpected foreign-key "target_user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (toq *TransferOrderQuery) loadBills(ctx context.Context, query *BillQuery, nodes []*TransferOrder, init func(*TransferOrder), assign func(*TransferOrder, *Bill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*TransferOrder)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(bill.FieldOrderID)
	}
	query.Where(predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(transferorder.BillsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrderID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "order_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (toq *TransferOrderQuery) loadVxSocial(ctx context.Context, query *VXSocialQuery, nodes []*TransferOrder, init func(*TransferOrder), assign func(*TransferOrder, *VXSocial)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*TransferOrder)
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
func (toq *TransferOrderQuery) loadSymbol(ctx context.Context, query *SymbolQuery, nodes []*TransferOrder, init func(*TransferOrder), assign func(*TransferOrder, *Symbol)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*TransferOrder)
	for i := range nodes {
		fk := nodes[i].SymbolID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(symbol.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "symbol_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (toq *TransferOrderQuery) loadWithdrawRecord(ctx context.Context, query *WithdrawRecordQuery, nodes []*TransferOrder, init func(*TransferOrder), assign func(*TransferOrder, *WithdrawRecord)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*TransferOrder)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(withdrawrecord.FieldTransferOrderID)
	}
	query.Where(predicate.WithdrawRecord(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(transferorder.WithdrawRecordColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TransferOrderID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "transfer_order_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (toq *TransferOrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := toq.querySpec()
	if len(toq.modifiers) > 0 {
		_spec.Modifiers = toq.modifiers
	}
	_spec.Node.Columns = toq.ctx.Fields
	if len(toq.ctx.Fields) > 0 {
		_spec.Unique = toq.ctx.Unique != nil && *toq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, toq.driver, _spec)
}

func (toq *TransferOrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(transferorder.Table, transferorder.Columns, sqlgraph.NewFieldSpec(transferorder.FieldID, field.TypeInt64))
	_spec.From = toq.sql
	if unique := toq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if toq.path != nil {
		_spec.Unique = true
	}
	if fields := toq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transferorder.FieldID)
		for i := range fields {
			if fields[i] != transferorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if toq.withSourceUser != nil {
			_spec.Node.AddColumnOnce(transferorder.FieldSourceUserID)
		}
		if toq.withTargetUser != nil {
			_spec.Node.AddColumnOnce(transferorder.FieldTargetUserID)
		}
		if toq.withVxSocial != nil {
			_spec.Node.AddColumnOnce(transferorder.FieldSocialID)
		}
		if toq.withSymbol != nil {
			_spec.Node.AddColumnOnce(transferorder.FieldSymbolID)
		}
	}
	if ps := toq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := toq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := toq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := toq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (toq *TransferOrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(toq.driver.Dialect())
	t1 := builder.Table(transferorder.Table)
	columns := toq.ctx.Fields
	if len(columns) == 0 {
		columns = transferorder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if toq.sql != nil {
		selector = toq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if toq.ctx.Unique != nil && *toq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range toq.modifiers {
		m(selector)
	}
	for _, p := range toq.predicates {
		p(selector)
	}
	for _, p := range toq.order {
		p(selector)
	}
	if offset := toq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := toq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (toq *TransferOrderQuery) Modify(modifiers ...func(s *sql.Selector)) *TransferOrderSelect {
	toq.modifiers = append(toq.modifiers, modifiers...)
	return toq.Select()
}

// TransferOrderGroupBy is the group-by builder for TransferOrder entities.
type TransferOrderGroupBy struct {
	selector
	build *TransferOrderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (togb *TransferOrderGroupBy) Aggregate(fns ...AggregateFunc) *TransferOrderGroupBy {
	togb.fns = append(togb.fns, fns...)
	return togb
}

// Scan applies the selector query and scans the result into the given value.
func (togb *TransferOrderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, togb.build.ctx, "GroupBy")
	if err := togb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TransferOrderQuery, *TransferOrderGroupBy](ctx, togb.build, togb, togb.build.inters, v)
}

func (togb *TransferOrderGroupBy) sqlScan(ctx context.Context, root *TransferOrderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(togb.fns))
	for _, fn := range togb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*togb.flds)+len(togb.fns))
		for _, f := range *togb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*togb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := togb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TransferOrderSelect is the builder for selecting fields of TransferOrder entities.
type TransferOrderSelect struct {
	*TransferOrderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tos *TransferOrderSelect) Aggregate(fns ...AggregateFunc) *TransferOrderSelect {
	tos.fns = append(tos.fns, fns...)
	return tos
}

// Scan applies the selector query and scans the result into the given value.
func (tos *TransferOrderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tos.ctx, "Select")
	if err := tos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TransferOrderQuery, *TransferOrderSelect](ctx, tos.TransferOrderQuery, tos, tos.inters, v)
}

func (tos *TransferOrderSelect) sqlScan(ctx context.Context, root *TransferOrderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tos.fns))
	for _, fn := range tos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tos *TransferOrderSelect) Modify(modifiers ...func(s *sql.Selector)) *TransferOrderSelect {
	tos.modifiers = append(tos.modifiers, modifiers...)
	return tos
}
