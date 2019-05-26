package stmt

import "github.com/getbud/bud/sql/rendering"

// Select ...
type Select struct {
	Distinct bool
	Columns  []Column
	Tables   []Table
	OrderBys []OrderBy
	Alias    string
}

// As ...
func (s Select) As(alias string) Select {
	s.Alias = alias
	return s
}

// WriteExpression ...
func (s Select) WriteExpression(w *rendering.Writer) {
	s.WriteStatement(w)

	if s.Alias != "" {
		w.Write(" AS ")
		w.Write(s.Alias)
	}
}

// WriteStatement ...
func (s Select) WriteStatement(w *rendering.Writer) {
	w.Write("SELECT ")

	if len(s.Columns) > 0 {
		for i, col := range s.Columns {
			col.WriteExpression(w)
			if i != len(s.Columns)-1 {
				w.Write(", ")
			}
		}
	} else {
		w.Write("*")
	}

	if len(s.Tables) > 0 {
		w.Write(" FROM ")

		for i, tab := range s.Tables {
			tab.WriteExpression(w)
			if i != len(s.Tables)-1 {
				w.Write(", ")
			}
		}
	}

	if len(s.OrderBys) > 0 {
		w.Write(" ORDER BY ")

		for i, o := range s.OrderBys {
			o.WriteExpression(w)
			if i != len(s.OrderBys)-1 {
				w.Write(", ")
			}
		}
	}
}
