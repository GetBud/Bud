package stmt

// Schema ...
type Schema struct {
	Name string
}

// NewSchema returns a new Schema.
func NewSchema(name string) *Schema {
	return &Schema{
		Name: name,
	}
}

// Table ...
func (s Schema) Table(name string) *Table {
	return NewTable(s, name)
}

// T ...
func (s Schema) T(name string) *Table {
	return s.Table(name)
}

// IsEmpty ...
func (s Schema) IsEmpty() bool {
	return s.Name == ""
}
