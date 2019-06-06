package stmt

import (
	"strconv"

	"github.com/getbud/bud/lab/sql/builder"
)

// String ...
type String struct {
	comparisonOperations
	val   string
	alias string
}

// NewString ...
func NewString(val string) *String {
	s := &String{val: val}
	s.comparisonOperations = comparisonOperations{s}

	return s
}

// Alias ...
func (s *String) Alias() string {
	return s.alias
}

// As ...
func (s *String) As(alias string) *String {
	s.alias = alias
	return s
}

// BuildExpression ...
func (s *String) BuildExpression(ctx *builder.Context) {
	ctx.Write("'")
	ctx.Write(s.val)
	ctx.Write("'")
}

// Int ...
type Int struct {
	comparisonOperations
	val   int
	alias string
}

// NewInt ...
func NewInt(val int) *Int {
	i := &Int{val: val}
	i.comparisonOperations = comparisonOperations{i}

	return i
}

// Alias ...
func (i *Int) Alias() string {
	return i.alias
}

// As ...
func (i *Int) As(alias string) *Int {
	i.alias = alias
	return i
}

// BuildExpression ...
func (i *Int) BuildExpression(ctx *builder.Context) {
	ctx.Write(strconv.Itoa(i.val))
}

type Null struct {
	comparisonOperations
	alias string
}

// NewNull ...
func NewNull() *Null {
	n := &Null{}
	n.comparisonOperations = comparisonOperations{n}

	return n
}

// Alias ...
func (n *Null) Alias() string {
	return n.alias
}

// As ...
func (n *Null) As(alias string) *Null {
	n.alias = alias
	return n
}

// BuildExpression ...
func (n *Null) BuildExpression(ctx *builder.Context) {
	ctx.Write("NULL")
}
