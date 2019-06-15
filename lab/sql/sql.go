package sql

import (
	"github.com/getbud/bud/lab/sql/stmt"
	"github.com/getbud/bud/lab/sql/token"
)

// SelectBuilder ...
func Select(selectExpressions ...stmt.SelectExpression) *stmt.Select {
	return stmt.NewSelect(selectExpressions...)
}

// Schema ...
func Schema(name string) *stmt.Schema {
	return stmt.NewSchema(name)
}

// Table ...
func Table(name string) *stmt.Table {
	return stmt.NewTable(stmt.Schema{}, name)
}

// Column ...
func Column(name string) *stmt.Column {
	return stmt.NewColumn(stmt.Table{}, name)
}

// Join ...
func Join(joinType token.JoinType, fromItem stmt.FromItem) *stmt.Join {
	return stmt.NewJoin(joinType, fromItem)
}

// Function ...
// TODO: We can shortcut a ton of common functions, and make it easier to not mess them up too by
// making functions for many of them. May be best to be in another package though?
func Function(name string, args ...stmt.Expression) *stmt.Function {
	return stmt.NewFunction(name, args...)
}

// And ...
func And(conditions ...stmt.Condition) *stmt.ConditionList {
	return stmt.NewConditionList(token.And, conditions...)
}

// Or ...
func Or(conditions ...stmt.Condition) *stmt.ConditionList {
	return stmt.NewConditionList(token.Or, conditions...)
}

// Literals ...

// Int ...
func Int(val int) *stmt.Int {
	return stmt.NewInt(val)
}

// Null ...
func Null() *stmt.Null {
	return stmt.NewNull()
}

// String ...
func String(val string) *stmt.String {
	return stmt.NewString(val)
}

// Functions ...

// Count ...
func Count(expr stmt.Expression) *stmt.Function {
	return stmt.NewFunction("COUNT", expr)
}
