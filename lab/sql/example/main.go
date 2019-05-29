package main

import (
	"fmt"

	"github.com/getbud/bud/lab/sql"
)

var (
	// usersTable ...
	usersTable    = sql.Schema("bud").Table("users").As("u")
	usersIDCol    = usersTable.Column("id")
	usersNameCol  = usersTable.Column("name")
	usersEmailCol = usersTable.Column("email")

	// accountsTable ...
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
		Where(sql.And(
			usersIDCol.Eq(accountsIDCol),
			sql.Function("COUNT", accountsIDCol).Eq(sql.Int(123)),
		)).
		Where(sql.Or(
			usersEmailCol.Eq(usersNameCol),
			usersEmailCol.Eq(usersIDCol),
		)).
		Where(usersNameCol.Eq(sql.String("seeruk"))).
		Build()

	fmt.Println(qry)
	fmt.Println(args)
}
