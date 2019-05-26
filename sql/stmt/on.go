package stmt

import "github.com/getbud/bud/sql/rendering"

// On ...
type On struct {
	Left, Right Column
}

// NewOn returns a new On.
func NewOn(left, right Column) On {
	return On{
		Left:  left,
		Right: right,
	}
}

// Write ...
func (o On) Write(w *rendering.Writer) {

}
