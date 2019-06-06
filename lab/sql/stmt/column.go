package stmt

import (
	"github.com/getbud/bud/lab/sql/builder"
)

// Column ...
type Column struct {
	comparisonOperations

	table Table
	name  string
	alias string
}

// NewColumn returns a new Column.
func NewColumn(table Table, name string) *Column {
	col := &Column{
		table: table,
		name:  name,
	}

	col.comparisonOperations = comparisonOperations{col}

	return col
}

// Alias ...
func (c *Column) Alias() string {
	return c.alias
}

func (c *Column) As(alias string) *Column {
	c.alias = alias
	return c
}

// BuildExpression ...
func (c *Column) BuildExpression(ctx *builder.Context) {
	ctx.Write(`"`)

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
	ctx.Write(`"`)
}
