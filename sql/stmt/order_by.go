package stmt

import (
	"github.com/getbud/bud/sql/rendering"
	"github.com/getbud/bud/sql/token"
)

// OrderBy ...
type OrderBy struct {
	Column    Column
	Direction token.Order
}

// NewOrderBy ...
func NewOrderBy(column Column, direction token.Order) OrderBy {
	return OrderBy{
		Column:    column,
		Direction: direction,
	}
}

// WriteExpression ...
func (o OrderBy) WriteExpression(w *rendering.Writer) {
	o.Column.WriteReference(w)
	w.Write(" ")
	w.Write(o.Direction.String())
}
