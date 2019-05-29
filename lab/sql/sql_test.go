package sql_test

import (
	"testing"

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

func BenchmarkBuildSelect(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
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
				sql.Function("COUNT", accountsIDCol).Eq(usersIDCol),
				sql.Or(
					usersEmailCol.Eq(usersNameCol),
					usersEmailCol.Eq(usersIDCol),
				),
			)).
			Build()

		_, _ = qry, args
	}
}
