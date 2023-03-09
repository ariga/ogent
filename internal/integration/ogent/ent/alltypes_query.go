// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"ariga.io/ogent/internal/integration/ogent/ent/alltypes"
	"ariga.io/ogent/internal/integration/ogent/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AllTypesQuery is the builder for querying AllTypes entities.
type AllTypesQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.AllTypes
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AllTypesQuery builder.
func (atq *AllTypesQuery) Where(ps ...predicate.AllTypes) *AllTypesQuery {
	atq.predicates = append(atq.predicates, ps...)
	return atq
}

// Limit the number of records to be returned by this query.
func (atq *AllTypesQuery) Limit(limit int) *AllTypesQuery {
	atq.ctx.Limit = &limit
	return atq
}

// Offset to start from.
func (atq *AllTypesQuery) Offset(offset int) *AllTypesQuery {
	atq.ctx.Offset = &offset
	return atq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (atq *AllTypesQuery) Unique(unique bool) *AllTypesQuery {
	atq.ctx.Unique = &unique
	return atq
}

// Order specifies how the records should be ordered.
func (atq *AllTypesQuery) Order(o ...OrderFunc) *AllTypesQuery {
	atq.order = append(atq.order, o...)
	return atq
}

// First returns the first AllTypes entity from the query.
// Returns a *NotFoundError when no AllTypes was found.
func (atq *AllTypesQuery) First(ctx context.Context) (*AllTypes, error) {
	nodes, err := atq.Limit(1).All(setContextOp(ctx, atq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{alltypes.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (atq *AllTypesQuery) FirstX(ctx context.Context) *AllTypes {
	node, err := atq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AllTypes ID from the query.
// Returns a *NotFoundError when no AllTypes ID was found.
func (atq *AllTypesQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = atq.Limit(1).IDs(setContextOp(ctx, atq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{alltypes.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (atq *AllTypesQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := atq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AllTypes entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AllTypes entity is found.
// Returns a *NotFoundError when no AllTypes entities are found.
func (atq *AllTypesQuery) Only(ctx context.Context) (*AllTypes, error) {
	nodes, err := atq.Limit(2).All(setContextOp(ctx, atq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{alltypes.Label}
	default:
		return nil, &NotSingularError{alltypes.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (atq *AllTypesQuery) OnlyX(ctx context.Context) *AllTypes {
	node, err := atq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AllTypes ID in the query.
// Returns a *NotSingularError when more than one AllTypes ID is found.
// Returns a *NotFoundError when no entities are found.
func (atq *AllTypesQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = atq.Limit(2).IDs(setContextOp(ctx, atq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{alltypes.Label}
	default:
		err = &NotSingularError{alltypes.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (atq *AllTypesQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := atq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AllTypesSlice.
func (atq *AllTypesQuery) All(ctx context.Context) ([]*AllTypes, error) {
	ctx = setContextOp(ctx, atq.ctx, "All")
	if err := atq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AllTypes, *AllTypesQuery]()
	return withInterceptors[[]*AllTypes](ctx, atq, qr, atq.inters)
}

// AllX is like All, but panics if an error occurs.
func (atq *AllTypesQuery) AllX(ctx context.Context) []*AllTypes {
	nodes, err := atq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AllTypes IDs.
func (atq *AllTypesQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if atq.ctx.Unique == nil && atq.path != nil {
		atq.Unique(true)
	}
	ctx = setContextOp(ctx, atq.ctx, "IDs")
	if err = atq.Select(alltypes.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (atq *AllTypesQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := atq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (atq *AllTypesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, atq.ctx, "Count")
	if err := atq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, atq, querierCount[*AllTypesQuery](), atq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (atq *AllTypesQuery) CountX(ctx context.Context) int {
	count, err := atq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (atq *AllTypesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, atq.ctx, "Exist")
	switch _, err := atq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (atq *AllTypesQuery) ExistX(ctx context.Context) bool {
	exist, err := atq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AllTypesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (atq *AllTypesQuery) Clone() *AllTypesQuery {
	if atq == nil {
		return nil
	}
	return &AllTypesQuery{
		config:     atq.config,
		ctx:        atq.ctx.Clone(),
		order:      append([]OrderFunc{}, atq.order...),
		inters:     append([]Interceptor{}, atq.inters...),
		predicates: append([]predicate.AllTypes{}, atq.predicates...),
		// clone intermediate query.
		sql:  atq.sql.Clone(),
		path: atq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AllTypes.Query().
//		GroupBy(alltypes.FieldInt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (atq *AllTypesQuery) GroupBy(field string, fields ...string) *AllTypesGroupBy {
	atq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AllTypesGroupBy{build: atq}
	grbuild.flds = &atq.ctx.Fields
	grbuild.label = alltypes.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//	}
//
//	client.AllTypes.Query().
//		Select(alltypes.FieldInt).
//		Scan(ctx, &v)
func (atq *AllTypesQuery) Select(fields ...string) *AllTypesSelect {
	atq.ctx.Fields = append(atq.ctx.Fields, fields...)
	sbuild := &AllTypesSelect{AllTypesQuery: atq}
	sbuild.label = alltypes.Label
	sbuild.flds, sbuild.scan = &atq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AllTypesSelect configured with the given aggregations.
func (atq *AllTypesQuery) Aggregate(fns ...AggregateFunc) *AllTypesSelect {
	return atq.Select().Aggregate(fns...)
}

func (atq *AllTypesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range atq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, atq); err != nil {
				return err
			}
		}
	}
	for _, f := range atq.ctx.Fields {
		if !alltypes.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if atq.path != nil {
		prev, err := atq.path(ctx)
		if err != nil {
			return err
		}
		atq.sql = prev
	}
	return nil
}

func (atq *AllTypesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AllTypes, error) {
	var (
		nodes = []*AllTypes{}
		_spec = atq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AllTypes).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AllTypes{config: atq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, atq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (atq *AllTypesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := atq.querySpec()
	_spec.Node.Columns = atq.ctx.Fields
	if len(atq.ctx.Fields) > 0 {
		_spec.Unique = atq.ctx.Unique != nil && *atq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, atq.driver, _spec)
}

func (atq *AllTypesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(alltypes.Table, alltypes.Columns, sqlgraph.NewFieldSpec(alltypes.FieldID, field.TypeUint32))
	_spec.From = atq.sql
	if unique := atq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if atq.path != nil {
		_spec.Unique = true
	}
	if fields := atq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, alltypes.FieldID)
		for i := range fields {
			if fields[i] != alltypes.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := atq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := atq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := atq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := atq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (atq *AllTypesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(atq.driver.Dialect())
	t1 := builder.Table(alltypes.Table)
	columns := atq.ctx.Fields
	if len(columns) == 0 {
		columns = alltypes.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if atq.sql != nil {
		selector = atq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if atq.ctx.Unique != nil && *atq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range atq.predicates {
		p(selector)
	}
	for _, p := range atq.order {
		p(selector)
	}
	if offset := atq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := atq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AllTypesGroupBy is the group-by builder for AllTypes entities.
type AllTypesGroupBy struct {
	selector
	build *AllTypesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atgb *AllTypesGroupBy) Aggregate(fns ...AggregateFunc) *AllTypesGroupBy {
	atgb.fns = append(atgb.fns, fns...)
	return atgb
}

// Scan applies the selector query and scans the result into the given value.
func (atgb *AllTypesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, atgb.build.ctx, "GroupBy")
	if err := atgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AllTypesQuery, *AllTypesGroupBy](ctx, atgb.build, atgb, atgb.build.inters, v)
}

func (atgb *AllTypesGroupBy) sqlScan(ctx context.Context, root *AllTypesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(atgb.fns))
	for _, fn := range atgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*atgb.flds)+len(atgb.fns))
		for _, f := range *atgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*atgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AllTypesSelect is the builder for selecting fields of AllTypes entities.
type AllTypesSelect struct {
	*AllTypesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ats *AllTypesSelect) Aggregate(fns ...AggregateFunc) *AllTypesSelect {
	ats.fns = append(ats.fns, fns...)
	return ats
}

// Scan applies the selector query and scans the result into the given value.
func (ats *AllTypesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ats.ctx, "Select")
	if err := ats.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AllTypesQuery, *AllTypesSelect](ctx, ats.AllTypesQuery, ats, ats.inters, v)
}

func (ats *AllTypesSelect) sqlScan(ctx context.Context, root *AllTypesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ats.fns))
	for _, fn := range ats.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ats.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
