package stmt

import (
	"github.com/getbud/bud/lab/sql/builder"
	"github.com/getbud/bud/lab/sql/token"
)

// Select ...
type Select struct {
	distinct          bool
	selectExpressions []SelectExpression
	fromItems         []FromItem
	joins             []*Join
	whereConditions   []Condition
	havingConditions  []Condition
	alias             string
}

// NewSelect ...
func NewSelect(selectExpressions ...SelectExpression) *Select {
	return &Select{
		selectExpressions: selectExpressions,
	}
}

// Distinct ...
func (s *Select) Distinct() *Select {
	s.distinct = true
	return s
}

// Select ...
func (s *Select) Select(expressions ...SelectExpression) *Select {
	s.selectExpressions = append(s.selectExpressions, expressions...)
	return s
}

// From ...
func (s *Select) From(fromItems ...FromItem) *Select {
	s.fromItems = append(s.fromItems, fromItems...)
	return s
}

// Join ...
func (s *Select) Join(join *Join) *Select {
	s.joins = append(s.joins, join)
	return s
}

// InnerJoin ...
func (s *Select) InnerJoin(fromItem FromItem, conditions ...Condition) *Select {
	return s.Join(NewJoin(token.InnerJoin, fromItem).On(conditions...))
}

// Where ...
func (s *Select) Where(conditions ...Condition) *Select {
	s.whereConditions = append(s.whereConditions, conditions...)
	return s
}

// Having ...
func (s *Select) Having(conditions ...Condition) *Select {
	s.havingConditions = append(s.havingConditions, conditions...)
	return s
}

// As ...
func (s *Select) As(alias string) *Select {
	s.alias = alias
	return s
}

// WriteFromItem ...
func (s *Select) WriteFromItem(ctx *builder.Context) {
	ctx.Write("(")
	s.WriteStatement(ctx)
	ctx.Write(") AS ")
	ctx.Write(s.alias)
}

// WriteStatement ...
func (s *Select) WriteStatement(ctx *builder.Context) {
	ctx.Write("SELECT ")

	if s.distinct {
		ctx.Write("DISTINCT ")
	}

	if len(s.selectExpressions) > 0 {
		for i, expr := range s.selectExpressions {
			expr.BuildExpression(ctx)

			if alias := expr.Alias(); alias != "" {
				ctx.Write(" AS ")
				ctx.Write(alias)
			}

			if i < len(s.selectExpressions)-1 {
				ctx.Write(", ")
			}
		}
	} else {
		ctx.Write("*")
	}

	if len(s.fromItems) > 0 {
		ctx.Write(" FROM ")

		for i, fromItem := range s.fromItems {
			fromItem.WriteFromItem(ctx)

			if i < len(s.fromItems)-1 {
				ctx.Write(", ")
			}
		}
	}

	if len(s.joins) > 0 {
		for _, join := range s.joins {
			join.WriteJoin(ctx)
		}
	}

	if len(s.whereConditions) > 0 {
		ctx.Write(" WHERE ")

		for i, wc := range s.whereConditions {
			isList := wc.IsList()

			if isList {
				ctx.Write("(")
			}

			wc.BuildCondition(ctx)

			if isList {
				ctx.Write(")")
			}

			// By default, the relationship is where. If an OR is needed, then wrap using sql.Or.
			if i < len(s.whereConditions)-1 {
				ctx.Write(" AND ")
			}
		}
	}

	if len(s.havingConditions) > 0 {
		ctx.Write(" HAVING ")

		for i, wc := range s.havingConditions {
			isList := wc.IsList()

			if isList {
				ctx.Write("(")
			}

			wc.BuildCondition(ctx)

			if isList {
				ctx.Write(")")
			}

			// By default, the relationship is where. If an OR is needed, then wrap using sql.Or.
			if i < len(s.havingConditions)-1 {
				ctx.Write(" AND ")
			}
		}
	}
}

// Build ...
func (s *Select) Build() (string, []interface{}) {
	ctx := builder.NewContext()

	s.WriteStatement(ctx)

	return ctx.String(), ctx.Args()
}
