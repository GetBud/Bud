package stmt

import (
	"github.com/getbud/bud/sql/rendering"
)

// Table ...
type Table struct {
	Schema Schema
	Name   string
	Alias  string
}

// NewTable returns a new Table.
func NewTable(schema Schema, name string) Table {
	return Table{
		Schema: schema,
		Name:   name,
	}
}

// As ...
func (t Table) As(alias string) Table {
	t.Alias = alias
	return t
}

// Column ...
func (t Table) Column(name string) Column {
	return NewColumn(t, name)
}

// IsEmpty ...
func (t Table) IsEmpty() bool {
	return t.Name == ""
}

// WriteExpression ...
func (t Table) WriteExpression(w *rendering.Writer) {
	if t.Schema.Name != "" {
		w.Write(t.Schema.Name)
		w.Write(".")
	}

	w.Write(t.Name)

	if t.Alias != "" {
		w.Write(" AS ")
		w.Write(t.Alias)
	}
}

// WriteStatement ...
func (t Table) WriteReference(w *rendering.Writer) {
	if t.Alias != "" {
		w.Write(t.Alias)
	} else {
		if !t.Schema.IsEmpty() {
			t.Schema.WriteReference(w)
			w.Write(".")
		}

		w.Write(t.Name)
	}
}
