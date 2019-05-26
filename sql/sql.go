package sql

import (
	"github.com/getbud/bud/sql/builder"
	"github.com/getbud/bud/sql/stmt"
	"github.com/getbud/bud/sql/token"
)

// Select ...
func Select() builder.Select {
	return builder.NewSelect()
}

// Schema ...
func Schema(name string) stmt.Schema {
	return stmt.NewSchema(name)
}

// Table ...
func Table(name string) stmt.Table {
	return stmt.NewTable(stmt.Schema{}, name)
}

// Column ...
func Column(name string) stmt.Column {
	return stmt.NewColumn(stmt.Table{}, name)
}

func OrderBy(column stmt.Column, direction token.Order) stmt.OrderBy {
	return stmt.NewOrderBy(column, direction)
}
