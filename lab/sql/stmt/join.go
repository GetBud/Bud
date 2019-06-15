package stmt

import (
	"github.com/getbud/bud/lab/sql/builder"
	"github.com/getbud/bud/lab/sql/token"
)

// Join ...
type Join struct {
	isNatural  bool
	joinType   token.JoinType
	fromItem   FromItem
	conditions []Condition
}

// NewJoin returns a new Join instance.
func NewJoin(joinType token.JoinType, fromItem FromItem) *Join {
	return &Join{
		joinType: joinType,
		fromItem: fromItem,
	}
}

// Natural ...
func (j *Join) Natural() *Join {
	j.isNatural = true
	return j
}

// On ...
func (j *Join) On(conditions ...Condition) *Join {
	j.conditions = conditions
	return j
}

// WriteJoin ...
func (j *Join) WriteJoin(ctx *builder.Context) {
	if j.isNatural {
		ctx.Write("NATURAL ")
	}

	ctx.Write(j.joinType.String())
	ctx.Write(" ")

	j.fromItem.WriteFromItem(ctx)
}
