package builder

import (
	"github.com/getbud/bud/lab/sql/rendering"
	"github.com/getbud/bud/lab/sql/stmt"
)

// Select ...
type Select struct {
	stmt stmt.Select
}

// NewSelect returns a new Select instance.
func NewSelect() Select {
	return Select{}
}

// Distinct ...
func (s Select) Distinct() Select {
	s.stmt.Distinct = true
	return s
}

// Columns ...
func (s Select) Columns(columns ...stmt.Column) Select {
	s.stmt.Columns = columns
	return s
}

// From ...
func (s Select) From(tables ...stmt.Table) Select {
	s.stmt.Tables = tables
	return s
}

// InnerJoin ...
func (s Select) InnerJoin(table stmt.Table) Select {
	return s
}

// OrderBy ...
func (s Select) OrderBy(orderBys ...stmt.OrderBy) Select {
	s.stmt.OrderBys = orderBys
	return s
}

// Build ...
func (s Select) Build() (string, []interface{}) {
	w := rendering.NewWriter()

	s.stmt.WriteStatement(w)

	return w.String(), nil
}
