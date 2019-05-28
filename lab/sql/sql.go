package sql

import (
	"github.com/getbud/bud/lab/sql/stmt"
)

// SelectBuilder ...
func Select(selectExpressions ...stmt.SelectExpression) stmt.Select {
	return stmt.NewSelect(selectExpressions...)
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

// Function ...
// TODO: We can shortcut a ton of common functions, and make it easier to not mess them up too by
// making functions for many of them. May be best to be in another package though?
func Function(name string, args ...stmt.Expression) stmt.Function {
	return stmt.NewFunction(name, args...)
}
