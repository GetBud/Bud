package main

import (
	"fmt"

	"github.com/getbud/bud/sql"
	"github.com/getbud/bud/sql/token"
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
		Query()

	fmt.Println(qry)
}
