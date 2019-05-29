package stmt

import (
	"github.com/getbud/bud/lab/sql/builder"
	"github.com/getbud/bud/lab/sql/token"
)

// Column ...
type Column struct {
	table Table
	name  string
	alias string
}

// NewColumn returns a new Column.
func NewColumn(table Table, name string) Column {
	return Column{
		table: table,
		name:  name,
	}
}

// Alias ...
func (c Column) Alias() string {
	return c.alias
}

func (c Column) As(alias string) Column {
	c.alias = alias
	return c
}

// BuildExpression ...
func (c Column) BuildExpression(ctx *builder.Context) {
	if !c.table.IsEmpty() {
		if c.table.alias != "" {
			ctx.Write(c.table.alias)
		} else {
			if !c.table.schema.IsEmpty() {
				ctx.Write(c.table.schema.Name)
				ctx.Write(".")
			}

			ctx.Write(c.table.name)
		}

		ctx.Write(".")
	}

	ctx.Write(c.name)
}

// Eq ...
func (c Column) Eq(expr Expression) Condition {
	return NewComparisonCondition(token.Equal, c, expr)
}
