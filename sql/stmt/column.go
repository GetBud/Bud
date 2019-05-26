package stmt

import (
	"github.com/getbud/bud/sql/rendering"
)

// Column ...
type Column struct {
	Table Table
	Name  string
	Alias string
}

// NewColumn returns a new Column.
func NewColumn(table Table, name string) Column {
	return Column{
		Table: table,
		Name:  name,
	}
}

// As ...
func (c Column) As(alias string) Column {
	c.Alias = alias
	return c
}

// WriteExpression ...
func (c Column) WriteExpression(w *rendering.Writer) {
	if !c.Table.IsEmpty() {
		c.Table.WriteReference(w)
		w.Write(".")
	}

	w.Write(c.Name)

	if c.Alias != "" {
		w.Write(" AS ")
		w.Write(c.Alias)
	}
}

// WriteOnExpression ...
func (c Column) WriteOnExpression(w *rendering.Writer) {
	c.WriteExpression(w)
}

// WriteStatement ...
func (c Column) WriteReference(w *rendering.Writer) {
	if c.Alias != "" {
		w.Write(c.Alias)
	} else {
		c.Table.WriteReference(w)
		w.Write(".")
		w.Write(c.Name)
	}
}
