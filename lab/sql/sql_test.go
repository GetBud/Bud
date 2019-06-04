package sql_test

import (
	"testing"

	"github.com/getbud/bud/lab/sql"

	"gopkg.in/doug-martin/goqu.v5"
	_ "gopkg.in/doug-martin/goqu.v5/adapters/postgres"
)

var (
	// users ...
	users      = sql.Schema("bud").Table("users").As("u")
	usersID    = users.Column("id")
	usersName  = users.Column("name")
	usersEmail = users.Column("email")

	// accounts ...
	accounts   = sql.Schema("bud").Table("accounts").As("a")
	accountsID = accounts.Column("id")
)

func BenchmarkBuildSelect(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		qry, args := sql.
			Select(
				usersID,
				usersName.As("username"),
				usersEmail,
				sql.Function("COUNT", accountsID).As("count"),
			).
			From(users).
			Where(sql.And(
				usersID.Eq(accountsID),
				sql.Function("COUNT", accountsID).Eq(usersID),
				sql.Or(
					usersEmail.Eq(usersName),
					usersEmail.Eq(usersName),
				),
			)).
			Build()

		_, _ = qry, args
	}
}

func BenchmarkGoquBuildSelect(b *testing.B) {
	builder := goqu.New("postgres", nil)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		qry, args, _ := builder.From("bud.users").As("u").
			Select("u.id", "u.name", "u.email").
			Where(goqu.And(
				goqu.I("u.id").Eq("a.id"),
				goqu.COUNT("a.id").Eq("u.id"),
				goqu.Or(
					goqu.I("u.email").Eq("u.name"),
					goqu.I("u.email").Eq("u.id"),
				),
			)).
			Prepared(true).
			ToSql()

		_, _ = qry, args
	}
}
