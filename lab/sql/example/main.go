package main

import (
	"fmt"

	"github.com/getbud/bud/lab/sql"
)

var (
	// usersTable ...
	usersTable = sql.Schema("bud").Table("users").As("u")

	usersIDCol    = usersTable.Column("id")
	usersNameCol  = usersTable.Column("name")
	usersEmailCol = usersTable.Column("email")

	accountsTable = sql.Schema("bud").Table("accounts").As("a")

	accountsIDCol = accountsTable.Column("id")
)

func main() {
	qry, args := sql.
		Select(
			usersIDCol,
			usersNameCol.As("username"),
			usersEmailCol,
			sql.Function("COUNT", accountsIDCol).As("count"),
		).
		From(usersTable).
		Build()

	fmt.Println(qry)
	fmt.Println(args)
}
