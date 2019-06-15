package sql

import (
	"strconv"
	"strings"
)

// expression is used to join together a string, and any arguments that may be used in that string
// for placeholders, this allows parts of a query to be cleared without affecting other parts of the
// query (e.g. clearing selections without affecting the args passed to where conditions).
type expression struct {
	expr string
	args []interface{}
}

// SelectStatement is the builder type for SQL SELECT queries.
type SelectStatement struct {
	selections []expression
	fromItems  []expression
	joins      []expression
	groupBys   []expression
	orderBys   []expression
	limit      int
	offset     int
}

// NewSelectStatement returns a new SelectStatement.
func NewSelectStatement() *SelectStatement {
	return &SelectStatement{}
}

// Select ...
func (s *SelectStatement) Select(expr string, args ...interface{}) *SelectStatement {
	s.selections = append(s.selections, expression{
		expr: strings.Join(strings.Fields(expr), " "),
		args: args,
	})

	return s
}

// From ...
func (s *SelectStatement) From(fromItem string, args ...interface{}) *SelectStatement {
	s.fromItems = append(s.fromItems, expression{
		expr: strings.Join(strings.Fields(fromItem), " "),
		args: args,
	})

	return s
}

// InnerJoin ...
func (s *SelectStatement) InnerJoin(join string, args ...interface{}) *SelectStatement {
	s.joins = append(s.joins, expression{
		expr: "INNER JOIN " + strings.Join(strings.Fields(join), " "),
		args: args,
	})

	return s
}

// GroupBy ...
func (s *SelectStatement) GroupBy(expr string, args ...interface{}) *SelectStatement {
	s.groupBys = append(s.groupBys, expression{
		expr: strings.Join(strings.Fields(expr), " "),
		args: args,
	})

	return s
}

// OrderBy ...
func (s *SelectStatement) OrderBy(expr string, args ...interface{}) *SelectStatement {
	s.orderBys = append(s.orderBys, expression{
		expr: strings.Join(strings.Fields(expr), " "),
		args: args,
	})

	return s
}

// Limit ...
func (s *SelectStatement) Limit(limit int) *SelectStatement {
	s.limit = limit
	return s
}

// Offset ...
func (s *SelectStatement) Offset(offset int) *SelectStatement {
	s.offset = offset
	return s
}

// Build returns a SQL string, and any arguments collected for prepared statements.
func (s *SelectStatement) Build() (string, []interface{}) {
	ctx := NewContext()

	ctx.Write("SELECT ")

	if len(s.selections) > 0 {
		for i, selection := range s.selections {
			ctx.Write(selection.expr)
			ctx.AddArgs(selection.args...)

			if i < len(s.selections)-1 {
				ctx.Write(", ")
			}
		}
	} else {
		ctx.Write("*")
	}

	if len(s.fromItems) > 0 {
		ctx.Write(" FROM ")
		for i, fromItem := range s.fromItems {
			ctx.Write(fromItem.expr)
			ctx.AddArgs(fromItem.args...)

			if i < len(s.fromItems)-1 {
				ctx.Write(", ")
			}
		}
	}

	if len(s.joins) > 0 {
		for _, join := range s.joins {
			ctx.Write(" ")
			ctx.Write(join.expr)
			ctx.AddArgs(join.args...)
		}
	}

	if len(s.groupBys) > 0 {
		ctx.Write(" GROUP BY ")
		for i, groupBy := range s.groupBys {
			ctx.Write(groupBy.expr)
			ctx.AddArgs(groupBy.args...)

			if i < len(s.groupBys)-1 {
				ctx.Write(", ")
			}
		}
	}

	if len(s.orderBys) > 0 {
		ctx.Write(" ORDER BY ")
		for i, orderBy := range s.orderBys {
			ctx.Write(orderBy.expr)
			ctx.AddArgs(orderBy.args...)

			if i < len(s.orderBys)-1 {
				ctx.Write(", ")
			}
		}
	}

	if s.limit > 0 {
		ctx.Write(" LIMIT ")
		ctx.Write(strconv.Itoa(s.limit))

		if s.offset > 0 {
			ctx.Write(" OFFSET ")
			ctx.Write(strconv.Itoa(s.offset))
		}
	}

	return ctx.String(), ctx.Args()
}
