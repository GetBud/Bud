package stmt

import (
	"github.com/getbud/bud/lab/sql/builder"
	"github.com/getbud/bud/lab/sql/token"
)

// Condition ...
type Condition interface {
	IsList() bool
	BuildCondition(ctx *builder.Context)
}

// ConditionList ...
type ConditionList struct {
	operator   token.LogicalOperator
	conditions []Condition
}

func NewConditionList(operator token.LogicalOperator, conditions ...Condition) *ConditionList {
	return &ConditionList{
		operator:   operator,
		conditions: conditions,
	}
}

// BuildCondition ...
func (c *ConditionList) BuildCondition(ctx *builder.Context) {
	for i, condition := range c.conditions {
		isList := condition.IsList()

		if isList {
			ctx.Write("(")
		}

		condition.BuildCondition(ctx)

		if isList {
			ctx.Write(")")
		}

		if i < len(c.conditions)-1 {
			ctx.Write(" ")
			ctx.Write(string(c.operator))
			ctx.Write(" ")
		}
	}
}

// IsList ...
func (c *ConditionList) IsList() bool {
	return true
}

// ComparisonCondition ...
type ComparisonCondition struct {
	operator    token.ComparisonOperator
	left, right Expression
}

// NewComparisonCondition returns a new ComparisonCondition.
func NewComparisonCondition(operator token.ComparisonOperator, left, right Expression) *ComparisonCondition {
	return &ComparisonCondition{
		operator: operator,
		left:     left,
		right:    right,
	}
}

// BuildCondition ...
func (c *ComparisonCondition) BuildCondition(ctx *builder.Context) {
	c.left.BuildExpression(ctx)
	ctx.Write(" ")
	ctx.Write(string(c.operator))
	ctx.Write(" ")
	c.right.BuildExpression(ctx)
}

// IsList ...
func (c *ComparisonCondition) IsList() bool {
	return false
}

// BetweenCondition ...
type BetweenCondition struct {
	o       token.ComparisonOperator
	a, x, y Expression
}

// NewBetweenCondition returns a new BetweenCondition.
func NewBetweenCondition(a, x, y Expression) *BetweenCondition {
	return &BetweenCondition{
		o: token.Between,
		a: a,
		x: x,
		y: y,
	}
}

// NewNotBetweenCondition returns a new BetweenCondition.
func NewNotBetweenCondition(a, x, y Expression) *BetweenCondition {
	return &BetweenCondition{
		o: token.NotBetween,
		a: a,
		x: x,
		y: y,
	}
}

// BuildCondition ...
func (c *BetweenCondition) BuildCondition(ctx *builder.Context) {
	c.a.BuildExpression(ctx)
	ctx.Write(" ")
	ctx.Write(string(c.o))
	ctx.Write(" ")
	c.x.BuildExpression(ctx)
	ctx.Write(" ")
	ctx.Write(string(token.And))
	ctx.Write(" ")
	c.y.BuildExpression(ctx)
}

// IsList ...
func (c *BetweenCondition) IsList() bool {
	return false
}
