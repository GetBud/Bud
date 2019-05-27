package stmt

import "github.com/getbud/bud/lab/sql/rendering"

// Schema ...
type Schema struct {
	Name string
}

// NewSchema returns a new Schema.
func NewSchema(name string) Schema {
	return Schema{
		Name: name,
	}
}

// Table ...
func (s Schema) Table(name string) Table {
	return NewTable(s, name)
}

// IsEmpty ...
func (s Schema) IsEmpty() bool {
	return s.Name == ""
}

// WriteExpression ...
func (s Schema) WriteExpression(w *rendering.Writer) {
	w.Write(s.Name)
}

// WriteStatement ...
func (s Schema) WriteReference(w *rendering.Writer) {
	w.Write(s.Name)
}
