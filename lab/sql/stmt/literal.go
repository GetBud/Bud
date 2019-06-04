package stmt

import (
	"strconv"

	"github.com/getbud/bud/lab/sql/builder"
)

// String ...
type String struct {
	val   string
	alias string
}

// NewString ...
func NewString(val string) String {
	return String{
		val: val,
	}
}

// Alias ...
func (s String) Alias() string {
	return s.alias
}

// As ...
func (s String) As(alias string) String {
	s.alias = alias
	return s
}

// BuildExpression ...
func (s String) BuildExpression(ctx *builder.Context) {
	ctx.Write("'")
	ctx.Write(s.val)
	ctx.Write("'")
}

// Int ...
type Int struct {
	val   int
	alias string
}

// NewInt ...
func NewInt(val int) Int {
	return Int{
		val: val,
	}
}

// Alias ...
func (i Int) Alias() string {
	return i.alias
}

// As ...
func (i Int) As(alias string) Int {
	i.alias = alias
	return i
}

// BuildExpression ...
func (i Int) BuildExpression(ctx *builder.Context) {
	ctx.Write(strconv.Itoa(i.val))
}
