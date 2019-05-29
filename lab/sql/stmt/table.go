package stmt

import "github.com/getbud/bud/lab/sql/builder"

// Table ...
type Table struct {
	schema Schema
	name   string
	alias  string
}

// NewTable returns a new Table.
func NewTable(schema Schema, name string) Table {
	return Table{
		schema: schema,
		name:   name,
	}
}

// As ...
func (t Table) As(alias string) Table {
	t.alias = alias
	return t
}

// Column ...
func (t Table) Column(name string) Column {
	return NewColumn(t, name)
}

// IsEmpty ...
func (t Table) IsEmpty() bool {
	return t.name == ""
}

// WriteFromItem ...
func (t Table) WriteFromItem(ctx *builder.Context) {
	ctx.Write(`"`)
	ctx.Write(t.name)
	ctx.Write(`"`)

	if t.alias != "" {
		ctx.Write(" AS ")
		ctx.Write(t.alias)
	}
}
