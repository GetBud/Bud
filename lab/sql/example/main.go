package main

import (
	"fmt"

	"github.com/getbud/bud/lab/sql"
)

var (
	// usersTable ...
	usersTable    = sql.Schema("bud").Table("users")
	usersIDCol    = usersTable.Column("id")
	usersNameCol  = usersTable.Column("name")
	usersEmailCol = usersTable.Column("email")

	// accountsTable ...
	accountsTable = sql.Schema("bud").Table("accounts")
	accountsIDCol = accountsTable.Column("id")
)

func main() {
	subquery := sql.Select(usersIDCol).From(usersTable)

	qry, _ := sql.
		Select(
			usersIDCol,
			usersNameCol.As("username"),
			usersEmailCol,
			sql.Count(sql.Int(1)).As("count3"),
			sql.String("sup").As("greeting"),
			sql.Column("sq.id"),
		).
		From(usersTable.As("u")).
		From(subquery.As("sq")).
		InnerJoin(usersTable).
		Where(sql.And(
			usersIDCol.Eq(accountsIDCol),
			sql.Count(accountsIDCol).Eq(sql.Int(123)),
		)).
		Where(sql.Or(
			usersEmailCol.Eq(usersNameCol),
			usersEmailCol.Eq(usersIDCol),
		)).
		Where(usersNameCol.Eq(sql.String("seeruk"))).
		Where(usersNameCol.NotBetween(sql.String("a"), sql.String("z"))).
		Where(usersNameCol.Is(sql.Null())).
		Having(usersNameCol.Eq(sql.String("seeruk"))).
		Build()

	fmt.Println(qry)
}
