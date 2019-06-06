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
	subquery := sql.Select(usersIDCol).From(usersTable)

	qry, _ := sql.
		Select(
			usersIDCol,
			usersNameCol.As("username"),
			usersEmailCol,
			sql.Function("COUNT", accountsIDCol).As("count"),
			sql.Function("COUNT", sql.Int(1)).As("count2"),
			sql.Count(sql.Int(1)).As("count3"),
			sql.String("sup").As("greeting"),
			sql.Int(42).As("meaning_of_life"),
		).
		From(usersTable).
		From(subquery.As("hello")).
		From(subquery.As("hello2")).
		Where(sql.And(
			usersIDCol.Eq(accountsIDCol),
			sql.Function("COUNT", accountsIDCol).Eq(sql.Int(123)),
		)).
		Where(sql.Or(
			usersEmailCol.Eq(usersNameCol),
			usersEmailCol.Eq(usersIDCol),
		)).
		Where(usersNameCol.Eq(sql.String("seeruk"))).
		Where(usersNameCol.NotBetween(sql.String("hmmmm"), sql.String("siodfiusdf"))).
		Where(usersNameCol.Is(sql.Null())).
		Having(usersNameCol.Eq(sql.String("seeruk"))).
		Build()

	fmt.Println(qry)
}
