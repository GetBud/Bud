package main

import (
	"fmt"

	"github.com/getbud/bud/lab/sql"
	"github.com/getbud/bud/lab/sql/token"
)

var (
	// usersTable ...
	usersTable = sql.Schema("bud").Table("users").As("u")

	usersIDCol    = usersTable.Column("id")
	usersNameCol  = usersTable.Column("name")
	usersEmailCol = usersTable.Column("email")

	accountsTable = sql.Schema("bud").Table("accounts").As("a")
)

func main() {
	qry, _ := sql.Select().
		Columns(usersIDCol, usersNameCol.As("username"), usersEmailCol).
		From(usersTable).
		InnerJoin(accountsTable).
		OrderBy(
			sql.OrderBy(usersNameCol, token.Asc),
			sql.OrderBy(usersEmailCol, token.Desc),
		).
		Build()

	fmt.Println(qry)
}
