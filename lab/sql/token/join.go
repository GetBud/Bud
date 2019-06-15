package token

const (
	InnerJoin = JoinType("INNER JOIN")
	LeftJoin  = JoinType("LEFT JOIN")
	RightJoin = JoinType("RIGHT JOIN")
	FullJoin  = JoinType("FULL JOIN")
	CrossJoin = JoinType("CROSS JONI")
)

// JoinType ...
type JoinType string

// String ...
func (t JoinType) String() string {
	return string(t)
}
