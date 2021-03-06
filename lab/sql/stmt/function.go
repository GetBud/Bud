package stmt

import (
	"strings"

	"github.com/getbud/bud/lab/sql/builder"
)

// Function ...
type Function struct {
	comparisonOperations

	name  string
	args  []Expression
	alias string
}

// NewFunction ...
func NewFunction(name string, args ...Expression) *Function {
	fn := &Function{
		name: strings.ToUpper(name),
		args: args,
	}

	fn.comparisonOperations = comparisonOperations{fn}

	return fn
}

// As ...
func (f *Function) As(alias string) *Function {
	f.alias = alias
	return f
}

// Alias ...
func (f *Function) Alias() string {
	return f.alias
}

// BuildExpression ...
func (f *Function) BuildExpression(ctx *builder.Context) {
	ctx.Write(f.name)
	ctx.Write("(")

	for i, arg := range f.args {
		arg.BuildExpression(ctx)

		if i < len(f.args)-1 {
			ctx.Write(", ")
		}
	}

	ctx.Write(")")
}
