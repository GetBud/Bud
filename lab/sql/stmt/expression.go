package stmt

import "github.com/getbud/bud/lab/sql/builder"

// Expression ...
type Expression interface {
	WriteExpression(ctx *builder.Context)
}

// SelectExpression ...
type SelectExpression interface {
	Expression
	Alias() string
}
