package stmt

import "github.com/getbud/bud/lab/sql/rendering"

// OnExpression ...
type OnExpression interface {
	WriteOnExpression(w *rendering.Writer)
}

// On ...
type On struct {
	Left, Right OnExpression
}

// NewOn returns a new On.
func NewOn(left, right OnExpression) On {
	return On{
		Left:  left,
		Right: right,
	}
}

// Write ...
func (o On) Write(w *rendering.Writer) {

}
