package stmt

import (
	"github.com/getbud/bud/lab/sql/builder"
)

// FromItem ...
type FromItem interface {
	WriteFromItem(ctx *builder.Context)
}
