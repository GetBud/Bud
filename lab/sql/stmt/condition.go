package stmt

import (
	"github.com/getbud/bud/lab/sql/builder"
	"github.com/getbud/bud/lab/sql/token"
)

// Condition ...
type Condition struct {
	left, right Expression
	operator    token.ComparisonOperator
}

// NewCondition returns a new Condition.
func NewCondition(left, right Expression, operator token.ComparisonOperator) Condition {
	return Condition{
		left:     left,
		right:    right,
		operator: operator,
	}
}

// BuildCondition ...
func (c Condition) BuildCondition(ctx *builder.Context) {
	c.left.BuildExpression(ctx)
	ctx.Write(" ")
	ctx.Write(string(c.operator))
	ctx.Write(" ")
	c.right.BuildExpression(ctx)
}
