// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/campaign"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/invite"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// InviteQuery is the builder for querying Invite entities.
type InviteQuery struct {
	config
	ctx          *QueryContext
	order        []invite.OrderOption
	inters       []Interceptor
	predicates   []predicate.Invite
	withUser     *UserQuery
	withCampaign *CampaignQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InviteQuery builder.
func (iq *InviteQuery) Where(ps ...predicate.Invite) *InviteQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *InviteQuery) Limit(limit int) *InviteQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *InviteQuery) Offset(offset int) *InviteQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InviteQuery) Unique(unique bool) *InviteQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *InviteQuery) Order(o ...invite.OrderOption) *InviteQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryUser chains the current query on the "user" edge.
func (iq *InviteQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invite.Table, invite.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, invite.UserTable, invite.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCampaign chains the current query on the "campaign" edge.
func (iq *InviteQuery) QueryCampaign() *CampaignQuery {
	query := (&CampaignClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invite.Table, invite.FieldID, selector),
			sqlgraph.To(campaign.Table, campaign.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, invite.CampaignTable, invite.CampaignColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Invite entity from the query.
// Returns a *NotFoundError when no Invite was found.
func (iq *InviteQuery) First(ctx context.Context) (*Invite, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{invite.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InviteQuery) FirstX(ctx context.Context) *Invite {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Invite ID from the query.
// Returns a *NotFoundError when no Invite ID was found.
func (iq *InviteQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{invite.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InviteQuery) FirstIDX(ctx context.Context) int64 {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Invite entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Invite entity is found.
// Returns a *NotFoundError when no Invite entities are found.
func (iq *InviteQuery) Only(ctx context.Context) (*Invite, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{invite.Label}
	default:
		return nil, &NotSingularError{invite.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InviteQuery) OnlyX(ctx context.Context) *Invite {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Invite ID in the query.
// Returns a *NotSingularError when more than one Invite ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InviteQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{invite.Label}
	default:
		err = &NotSingularError{invite.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InviteQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Invites.
func (iq *InviteQuery) All(ctx context.Context) ([]*Invite, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Invite, *InviteQuery]()
	return withInterceptors[[]*Invite](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *InviteQuery) AllX(ctx context.Context) []*Invite {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Invite IDs.
func (iq *InviteQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(invite.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InviteQuery) IDsX(ctx context.Context) []int64 {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InviteQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*InviteQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InviteQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InviteQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InviteQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InviteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InviteQuery) Clone() *InviteQuery {
	if iq == nil {
		return nil
	}
	return &InviteQuery{
		config:       iq.config,
		ctx:          iq.ctx.Clone(),
		order:        append([]invite.OrderOption{}, iq.order...),
		inters:       append([]Interceptor{}, iq.inters...),
		predicates:   append([]predicate.Invite{}, iq.predicates...),
		withUser:     iq.withUser.Clone(),
		withCampaign: iq.withCampaign.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InviteQuery) WithUser(opts ...func(*UserQuery)) *InviteQuery {
	query := (&UserClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withUser = query
	return iq
}

// WithCampaign tells the query-builder to eager-load the nodes that are connected to
// the "campaign" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InviteQuery) WithCampaign(opts ...func(*CampaignQuery)) *InviteQuery {
	query := (&CampaignClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withCampaign = query
	return iq
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
//	client.Invite.Query().
//		GroupBy(invite.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (iq *InviteQuery) GroupBy(field string, fields ...string) *InviteGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InviteGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = invite.Label
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
//	client.Invite.Query().
//		Select(invite.FieldCreatedBy).
//		Scan(ctx, &v)
func (iq *InviteQuery) Select(fields ...string) *InviteSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &InviteSelect{InviteQuery: iq}
	sbuild.label = invite.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InviteSelect configured with the given aggregations.
func (iq *InviteQuery) Aggregate(fns ...AggregateFunc) *InviteSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *InviteQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !invite.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InviteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Invite, error) {
	var (
		nodes       = []*Invite{}
		_spec       = iq.querySpec()
		loadedTypes = [2]bool{
			iq.withUser != nil,
			iq.withCampaign != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Invite).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Invite{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withUser; query != nil {
		if err := iq.loadUser(ctx, query, nodes, nil,
			func(n *Invite, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withCampaign; query != nil {
		if err := iq.loadCampaign(ctx, query, nodes, nil,
			func(n *Invite, e *Campaign) { n.Edges.Campaign = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *InviteQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Invite, init func(*Invite), assign func(*Invite, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Invite)
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
func (iq *InviteQuery) loadCampaign(ctx context.Context, query *CampaignQuery, nodes []*Invite, init func(*Invite), assign func(*Invite, *Campaign)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Invite)
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

func (iq *InviteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InviteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(invite.Table, invite.Columns, sqlgraph.NewFieldSpec(invite.FieldID, field.TypeInt64))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invite.FieldID)
		for i := range fields {
			if fields[i] != invite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if iq.withUser != nil {
			_spec.Node.AddColumnOnce(invite.FieldUserID)
		}
		if iq.withCampaign != nil {
			_spec.Node.AddColumnOnce(invite.FieldCampaignID)
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InviteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(invite.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = invite.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InviteGroupBy is the group-by builder for Invite entities.
type InviteGroupBy struct {
	selector
	build *InviteQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InviteGroupBy) Aggregate(fns ...AggregateFunc) *InviteGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *InviteGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InviteQuery, *InviteGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *InviteGroupBy) sqlScan(ctx context.Context, root *InviteQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InviteSelect is the builder for selecting fields of Invite entities.
type InviteSelect struct {
	*InviteQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *InviteSelect) Aggregate(fns ...AggregateFunc) *InviteSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *InviteSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InviteQuery, *InviteSelect](ctx, is.InviteQuery, is, is.inters, v)
}

func (is *InviteSelect) sqlScan(ctx context.Context, root *InviteQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}