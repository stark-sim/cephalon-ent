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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/artwork"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/artworklike"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// ArtworkQuery is the builder for querying Artwork entities.
type ArtworkQuery struct {
	config
	ctx              *QueryContext
	order            []artwork.OrderOption
	inters           []Interceptor
	predicates       []predicate.Artwork
	withAuthor       *UserQuery
	withArtworkLikes *ArtworkLikeQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArtworkQuery builder.
func (aq *ArtworkQuery) Where(ps ...predicate.Artwork) *ArtworkQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *ArtworkQuery) Limit(limit int) *ArtworkQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *ArtworkQuery) Offset(offset int) *ArtworkQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *ArtworkQuery) Unique(unique bool) *ArtworkQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *ArtworkQuery) Order(o ...artwork.OrderOption) *ArtworkQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryAuthor chains the current query on the "author" edge.
func (aq *ArtworkQuery) QueryAuthor() *UserQuery {
	query := (&UserClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(artwork.Table, artwork.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, artwork.AuthorTable, artwork.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryArtworkLikes chains the current query on the "artwork_likes" edge.
func (aq *ArtworkQuery) QueryArtworkLikes() *ArtworkLikeQuery {
	query := (&ArtworkLikeClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(artwork.Table, artwork.FieldID, selector),
			sqlgraph.To(artworklike.Table, artworklike.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, artwork.ArtworkLikesTable, artwork.ArtworkLikesColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Artwork entity from the query.
// Returns a *NotFoundError when no Artwork was found.
func (aq *ArtworkQuery) First(ctx context.Context) (*Artwork, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{artwork.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *ArtworkQuery) FirstX(ctx context.Context) *Artwork {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Artwork ID from the query.
// Returns a *NotFoundError when no Artwork ID was found.
func (aq *ArtworkQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{artwork.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *ArtworkQuery) FirstIDX(ctx context.Context) int64 {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Artwork entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Artwork entity is found.
// Returns a *NotFoundError when no Artwork entities are found.
func (aq *ArtworkQuery) Only(ctx context.Context) (*Artwork, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{artwork.Label}
	default:
		return nil, &NotSingularError{artwork.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *ArtworkQuery) OnlyX(ctx context.Context) *Artwork {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Artwork ID in the query.
// Returns a *NotSingularError when more than one Artwork ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *ArtworkQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{artwork.Label}
	default:
		err = &NotSingularError{artwork.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *ArtworkQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Artworks.
func (aq *ArtworkQuery) All(ctx context.Context) ([]*Artwork, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Artwork, *ArtworkQuery]()
	return withInterceptors[[]*Artwork](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *ArtworkQuery) AllX(ctx context.Context) []*Artwork {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Artwork IDs.
func (aq *ArtworkQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(artwork.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *ArtworkQuery) IDsX(ctx context.Context) []int64 {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *ArtworkQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*ArtworkQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *ArtworkQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *ArtworkQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *ArtworkQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArtworkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *ArtworkQuery) Clone() *ArtworkQuery {
	if aq == nil {
		return nil
	}
	return &ArtworkQuery{
		config:           aq.config,
		ctx:              aq.ctx.Clone(),
		order:            append([]artwork.OrderOption{}, aq.order...),
		inters:           append([]Interceptor{}, aq.inters...),
		predicates:       append([]predicate.Artwork{}, aq.predicates...),
		withAuthor:       aq.withAuthor.Clone(),
		withArtworkLikes: aq.withArtworkLikes.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ArtworkQuery) WithAuthor(opts ...func(*UserQuery)) *ArtworkQuery {
	query := (&UserClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withAuthor = query
	return aq
}

// WithArtworkLikes tells the query-builder to eager-load the nodes that are connected to
// the "artwork_likes" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ArtworkQuery) WithArtworkLikes(opts ...func(*ArtworkLikeQuery)) *ArtworkQuery {
	query := (&ArtworkLikeClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withArtworkLikes = query
	return aq
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
//	client.Artwork.Query().
//		GroupBy(artwork.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (aq *ArtworkQuery) GroupBy(field string, fields ...string) *ArtworkGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ArtworkGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = artwork.Label
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
//	client.Artwork.Query().
//		Select(artwork.FieldCreatedBy).
//		Scan(ctx, &v)
func (aq *ArtworkQuery) Select(fields ...string) *ArtworkSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &ArtworkSelect{ArtworkQuery: aq}
	sbuild.label = artwork.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ArtworkSelect configured with the given aggregations.
func (aq *ArtworkQuery) Aggregate(fns ...AggregateFunc) *ArtworkSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *ArtworkQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !artwork.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *ArtworkQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Artwork, error) {
	var (
		nodes       = []*Artwork{}
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withAuthor != nil,
			aq.withArtworkLikes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Artwork).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Artwork{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withAuthor; query != nil {
		if err := aq.loadAuthor(ctx, query, nodes, nil,
			func(n *Artwork, e *User) { n.Edges.Author = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withArtworkLikes; query != nil {
		if err := aq.loadArtworkLikes(ctx, query, nodes,
			func(n *Artwork) { n.Edges.ArtworkLikes = []*ArtworkLike{} },
			func(n *Artwork, e *ArtworkLike) { n.Edges.ArtworkLikes = append(n.Edges.ArtworkLikes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *ArtworkQuery) loadAuthor(ctx context.Context, query *UserQuery, nodes []*Artwork, init func(*Artwork), assign func(*Artwork, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Artwork)
	for i := range nodes {
		fk := nodes[i].AuthorID
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
			return fmt.Errorf(`unexpected foreign-key "author_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *ArtworkQuery) loadArtworkLikes(ctx context.Context, query *ArtworkLikeQuery, nodes []*Artwork, init func(*Artwork), assign func(*Artwork, *ArtworkLike)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Artwork)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(artworklike.FieldArtworkID)
	}
	query.Where(predicate.ArtworkLike(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(artwork.ArtworkLikesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ArtworkID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "artwork_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *ArtworkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *ArtworkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(artwork.Table, artwork.Columns, sqlgraph.NewFieldSpec(artwork.FieldID, field.TypeInt64))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artwork.FieldID)
		for i := range fields {
			if fields[i] != artwork.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if aq.withAuthor != nil {
			_spec.Node.AddColumnOnce(artwork.FieldAuthorID)
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *ArtworkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(artwork.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = artwork.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range aq.modifiers {
		m(selector)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aq *ArtworkQuery) Modify(modifiers ...func(s *sql.Selector)) *ArtworkSelect {
	aq.modifiers = append(aq.modifiers, modifiers...)
	return aq.Select()
}

// ArtworkGroupBy is the group-by builder for Artwork entities.
type ArtworkGroupBy struct {
	selector
	build *ArtworkQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *ArtworkGroupBy) Aggregate(fns ...AggregateFunc) *ArtworkGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *ArtworkGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArtworkQuery, *ArtworkGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *ArtworkGroupBy) sqlScan(ctx context.Context, root *ArtworkQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ArtworkSelect is the builder for selecting fields of Artwork entities.
type ArtworkSelect struct {
	*ArtworkQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *ArtworkSelect) Aggregate(fns ...AggregateFunc) *ArtworkSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *ArtworkSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArtworkQuery, *ArtworkSelect](ctx, as.ArtworkQuery, as, as.inters, v)
}

func (as *ArtworkSelect) sqlScan(ctx context.Context, root *ArtworkQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (as *ArtworkSelect) Modify(modifiers ...func(s *sql.Selector)) *ArtworkSelect {
	as.modifiers = append(as.modifiers, modifiers...)
	return as
}
