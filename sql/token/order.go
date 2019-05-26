package token

// Possible ordering directions.
const (
	Asc  = Order("ASC")
	Desc = Order("DESC")
)

// Order ...
type Order string

// String ...
func (o Order) String() string {
	return string(o)
}
