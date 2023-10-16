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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/symbol"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/transferorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/wallet"
)

// SymbolQuery is the builder for querying Symbol entities.
type SymbolQuery struct {
	config
	ctx                *QueryContext
	order              []symbol.OrderOption
	inters             []Interceptor
	predicates         []predicate.Symbol
	withWallets        *WalletQuery
	withBills          *BillQuery
	withMissionOrders  *MissionOrderQuery
	withTransferOrders *TransferOrderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SymbolQuery builder.
func (sq *SymbolQuery) Where(ps ...predicate.Symbol) *SymbolQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SymbolQuery) Limit(limit int) *SymbolQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SymbolQuery) Offset(offset int) *SymbolQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SymbolQuery) Unique(unique bool) *SymbolQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SymbolQuery) Order(o ...symbol.OrderOption) *SymbolQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryWallets chains the current query on the "wallets" edge.
func (sq *SymbolQuery) QueryWallets() *WalletQuery {
	query := (&WalletClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(symbol.Table, symbol.FieldID, selector),
			sqlgraph.To(wallet.Table, wallet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, symbol.WalletsTable, symbol.WalletsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBills chains the current query on the "bills" edge.
func (sq *SymbolQuery) QueryBills() *BillQuery {
	query := (&BillClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(symbol.Table, symbol.FieldID, selector),
			sqlgraph.To(bill.Table, bill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, symbol.BillsTable, symbol.BillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissionOrders chains the current query on the "mission_orders" edge.
func (sq *SymbolQuery) QueryMissionOrders() *MissionOrderQuery {
	query := (&MissionOrderClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(symbol.Table, symbol.FieldID, selector),
			sqlgraph.To(missionorder.Table, missionorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, symbol.MissionOrdersTable, symbol.MissionOrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransferOrders chains the current query on the "transfer_orders" edge.
func (sq *SymbolQuery) QueryTransferOrders() *TransferOrderQuery {
	query := (&TransferOrderClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(symbol.Table, symbol.FieldID, selector),
			sqlgraph.To(transferorder.Table, transferorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, symbol.TransferOrdersTable, symbol.TransferOrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Symbol entity from the query.
// Returns a *NotFoundError when no Symbol was found.
func (sq *SymbolQuery) First(ctx context.Context) (*Symbol, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{symbol.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SymbolQuery) FirstX(ctx context.Context) *Symbol {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Symbol ID from the query.
// Returns a *NotFoundError when no Symbol ID was found.
func (sq *SymbolQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{symbol.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SymbolQuery) FirstIDX(ctx context.Context) int64 {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Symbol entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Symbol entity is found.
// Returns a *NotFoundError when no Symbol entities are found.
func (sq *SymbolQuery) Only(ctx context.Context) (*Symbol, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{symbol.Label}
	default:
		return nil, &NotSingularError{symbol.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SymbolQuery) OnlyX(ctx context.Context) *Symbol {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Symbol ID in the query.
// Returns a *NotSingularError when more than one Symbol ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SymbolQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{symbol.Label}
	default:
		err = &NotSingularError{symbol.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SymbolQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Symbols.
func (sq *SymbolQuery) All(ctx context.Context) ([]*Symbol, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Symbol, *SymbolQuery]()
	return withInterceptors[[]*Symbol](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SymbolQuery) AllX(ctx context.Context) []*Symbol {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Symbol IDs.
func (sq *SymbolQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(symbol.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SymbolQuery) IDsX(ctx context.Context) []int64 {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SymbolQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SymbolQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SymbolQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SymbolQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SymbolQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SymbolQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SymbolQuery) Clone() *SymbolQuery {
	if sq == nil {
		return nil
	}
	return &SymbolQuery{
		config:             sq.config,
		ctx:                sq.ctx.Clone(),
		order:              append([]symbol.OrderOption{}, sq.order...),
		inters:             append([]Interceptor{}, sq.inters...),
		predicates:         append([]predicate.Symbol{}, sq.predicates...),
		withWallets:        sq.withWallets.Clone(),
		withBills:          sq.withBills.Clone(),
		withMissionOrders:  sq.withMissionOrders.Clone(),
		withTransferOrders: sq.withTransferOrders.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithWallets tells the query-builder to eager-load the nodes that are connected to
// the "wallets" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SymbolQuery) WithWallets(opts ...func(*WalletQuery)) *SymbolQuery {
	query := (&WalletClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withWallets = query
	return sq
}

// WithBills tells the query-builder to eager-load the nodes that are connected to
// the "bills" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SymbolQuery) WithBills(opts ...func(*BillQuery)) *SymbolQuery {
	query := (&BillClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withBills = query
	return sq
}

// WithMissionOrders tells the query-builder to eager-load the nodes that are connected to
// the "mission_orders" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SymbolQuery) WithMissionOrders(opts ...func(*MissionOrderQuery)) *SymbolQuery {
	query := (&MissionOrderClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withMissionOrders = query
	return sq
}

// WithTransferOrders tells the query-builder to eager-load the nodes that are connected to
// the "transfer_orders" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SymbolQuery) WithTransferOrders(opts ...func(*TransferOrderQuery)) *SymbolQuery {
	query := (&TransferOrderClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withTransferOrders = query
	return sq
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
//	client.Symbol.Query().
//		GroupBy(symbol.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (sq *SymbolQuery) GroupBy(field string, fields ...string) *SymbolGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SymbolGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = symbol.Label
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
//	client.Symbol.Query().
//		Select(symbol.FieldCreatedBy).
//		Scan(ctx, &v)
func (sq *SymbolQuery) Select(fields ...string) *SymbolSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SymbolSelect{SymbolQuery: sq}
	sbuild.label = symbol.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SymbolSelect configured with the given aggregations.
func (sq *SymbolQuery) Aggregate(fns ...AggregateFunc) *SymbolSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SymbolQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !symbol.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SymbolQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Symbol, error) {
	var (
		nodes       = []*Symbol{}
		_spec       = sq.querySpec()
		loadedTypes = [4]bool{
			sq.withWallets != nil,
			sq.withBills != nil,
			sq.withMissionOrders != nil,
			sq.withTransferOrders != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Symbol).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Symbol{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withWallets; query != nil {
		if err := sq.loadWallets(ctx, query, nodes,
			func(n *Symbol) { n.Edges.Wallets = []*Wallet{} },
			func(n *Symbol, e *Wallet) { n.Edges.Wallets = append(n.Edges.Wallets, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withBills; query != nil {
		if err := sq.loadBills(ctx, query, nodes,
			func(n *Symbol) { n.Edges.Bills = []*Bill{} },
			func(n *Symbol, e *Bill) { n.Edges.Bills = append(n.Edges.Bills, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withMissionOrders; query != nil {
		if err := sq.loadMissionOrders(ctx, query, nodes,
			func(n *Symbol) { n.Edges.MissionOrders = []*MissionOrder{} },
			func(n *Symbol, e *MissionOrder) { n.Edges.MissionOrders = append(n.Edges.MissionOrders, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withTransferOrders; query != nil {
		if err := sq.loadTransferOrders(ctx, query, nodes,
			func(n *Symbol) { n.Edges.TransferOrders = []*TransferOrder{} },
			func(n *Symbol, e *TransferOrder) { n.Edges.TransferOrders = append(n.Edges.TransferOrders, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SymbolQuery) loadWallets(ctx context.Context, query *WalletQuery, nodes []*Symbol, init func(*Symbol), assign func(*Symbol, *Wallet)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Symbol)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(wallet.FieldSymbolID)
	}
	query.Where(predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(symbol.WalletsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SymbolID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "symbol_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *SymbolQuery) loadBills(ctx context.Context, query *BillQuery, nodes []*Symbol, init func(*Symbol), assign func(*Symbol, *Bill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Symbol)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(bill.FieldSymbolID)
	}
	query.Where(predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(symbol.BillsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SymbolID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "symbol_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *SymbolQuery) loadMissionOrders(ctx context.Context, query *MissionOrderQuery, nodes []*Symbol, init func(*Symbol), assign func(*Symbol, *MissionOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Symbol)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(missionorder.FieldSymbolID)
	}
	query.Where(predicate.MissionOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(symbol.MissionOrdersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SymbolID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "symbol_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *SymbolQuery) loadTransferOrders(ctx context.Context, query *TransferOrderQuery, nodes []*Symbol, init func(*Symbol), assign func(*Symbol, *TransferOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Symbol)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(transferorder.FieldSymbolID)
	}
	query.Where(predicate.TransferOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(symbol.TransferOrdersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SymbolID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "symbol_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sq *SymbolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SymbolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(symbol.Table, symbol.Columns, sqlgraph.NewFieldSpec(symbol.FieldID, field.TypeInt64))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, symbol.FieldID)
		for i := range fields {
			if fields[i] != symbol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SymbolQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(symbol.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = symbol.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SymbolGroupBy is the group-by builder for Symbol entities.
type SymbolGroupBy struct {
	selector
	build *SymbolQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SymbolGroupBy) Aggregate(fns ...AggregateFunc) *SymbolGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SymbolGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SymbolQuery, *SymbolGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SymbolGroupBy) sqlScan(ctx context.Context, root *SymbolQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SymbolSelect is the builder for selecting fields of Symbol entities.
type SymbolSelect struct {
	*SymbolQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SymbolSelect) Aggregate(fns ...AggregateFunc) *SymbolSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SymbolSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SymbolQuery, *SymbolSelect](ctx, ss.SymbolQuery, ss, ss.inters, v)
}

func (ss *SymbolSelect) sqlScan(ctx context.Context, root *SymbolQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
