package stmt

import "github.com/getbud/bud/lab/sql/builder"

// Condition ...
type Condition interface {
	WriteCondition(ctx *builder.Context)
}
