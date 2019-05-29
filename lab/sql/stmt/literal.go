package stmt

import (
	"strconv"

	"github.com/getbud/bud/lab/sql/builder"
)

// String ...
type String struct {
	val string
}

// NewString ...
func NewString(val string) String {
	return String{
		val: val,
	}
}

// BuildExpression ...
func (s String) BuildExpression(ctx *builder.Context) {
	ctx.Write("'")
	ctx.Write(s.val)
	ctx.Write("'")
}

// Int ...
type Int struct {
	val int
}

// NewInt ...
func NewInt(val int) Int {
	return Int{
		val: val,
	}
}

// BuildExpression ...
func (i Int) BuildExpression(ctx *builder.Context) {
	ctx.Write(strconv.Itoa(i.val))
}
