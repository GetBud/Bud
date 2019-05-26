package token

// Possible logical operators.
const (
	And = LogicalOperator("AND")
	Or  = LogicalOperator("OR")
	Not = LogicalOperator("NOT")
)

// LogicalOperator ...
type LogicalOperator string

// String ...
func (o LogicalOperator) String() string {
	return string(o)
}

// Possible comparison operators.
const (
	Equal              = ComparisonOperator("=")
	NotEqual           = ComparisonOperator("!=")
	Is                 = ComparisonOperator("IS")
	IsNot              = ComparisonOperator("IS NOT")
	GreaterThan        = ComparisonOperator(">")
	GreaterThanOrEqual = ComparisonOperator(">=")
	LessThan           = ComparisonOperator("<")
	LessThanOrEqual    = ComparisonOperator("<=")
	In                 = ComparisonOperator("IN")
	NotIn              = ComparisonOperator("NOT IN")
	Like               = ComparisonOperator("LIKE")
	NotLike            = ComparisonOperator("NOT LIKE")
	ILike              = ComparisonOperator("ILIKE")
	NotILike           = ComparisonOperator("NOT ILIKE")
	Between            = ComparisonOperator("BETWEEN")
	NotBetween         = ComparisonOperator("NOT BETWEEN")
)

// ComparisonOperator ...
type ComparisonOperator string

// String ...
func (o ComparisonOperator) String() string {
	return string(o)
}
