package stmt

import (
	"github.com/getbud/bud/lab/sql/token"
)

// comparisonOperations ...
type comparisonOperations struct {
	expr Expression
}

// Eq ...
func (c *comparisonOperations) Eq(expr Expression) Condition {
	return NewComparisonCondition(token.Equal, c.expr, expr)
}

// Ne ...
func (c *comparisonOperations) Ne(expr Expression) Condition {
	return NewComparisonCondition(token.NotEqual, c.expr, expr)
}

// Is ...
func (c *comparisonOperations) Is(expr Expression) Condition {
	return NewComparisonCondition(token.Is, c.expr, expr)
}

// IsNot ...
func (c *comparisonOperations) IsNot(expr Expression) Condition {
	return NewComparisonCondition(token.IsNot, c.expr, expr)
}

// Gt ...
func (c *comparisonOperations) Gt(expr Expression) Condition {
	return NewComparisonCondition(token.GreaterThan, c.expr, expr)
}

// Gte ...
func (c *comparisonOperations) Gte(expr Expression) Condition {
	return NewComparisonCondition(token.GreaterThanOrEqual, c.expr, expr)
}

// Lt ...
func (c *comparisonOperations) Lt(expr Expression) Condition {
	return NewComparisonCondition(token.LessThan, c.expr, expr)
}

// Lte ...
func (c *comparisonOperations) Lte(expr Expression) Condition {
	return NewComparisonCondition(token.LessThanOrEqual, c.expr, expr)
}

// In ...
func (c *comparisonOperations) In(expr Expression) Condition {
	return NewComparisonCondition(token.In, c.expr, expr)
}

// NotIn ...
func (c *comparisonOperations) NotIn(expr Expression) Condition {
	return NewComparisonCondition(token.NotIn, c.expr, expr)
}

// Like ...
func (c *comparisonOperations) Like(expr Expression) Condition {
	return NewComparisonCondition(token.Like, c.expr, expr)
}

// NotLike ...
func (c *comparisonOperations) NotLike(expr Expression) Condition {
	return NewComparisonCondition(token.NotLike, c.expr, expr)
}

// ILike ...
func (c *comparisonOperations) ILike(expr Expression) Condition {
	return NewComparisonCondition(token.ILike, c.expr, expr)
}

// NotILike ...
func (c *comparisonOperations) NotILike(expr Expression) Condition {
	return NewComparisonCondition(token.NotILike, c.expr, expr)
}

// Between ...
func (c *comparisonOperations) Between(x, y Expression) Condition {
	return NewBetweenCondition(c.expr, x, y)
}

// NotBetween ...
func (c *comparisonOperations) NotBetween(x, y Expression) Condition {
	return NewNotBetweenCondition(c.expr, x, y)
}
