package main

import (
	"fmt"

	"github.com/getbud/bud/lab/sql2"
)

func main() {
	fields := []string{
		"u.id",
		"u.name",
		"u.email",
		"u.mobile",
		"(COUNT(1) > ?) AS count",
	}

	qry, args := sql.NewSelectStatement().
		Select("u.id").
		Select("(COUNT(1) > ?) AS foo", 123).
		Select("(COUNT(1) > ?) AS bar", 456).
		Selects(fields, 789).
		From("users AS u").
		InnerJoin("accounts AS a ON a.userId = a.id AND u.enabled = ?", true).
		GroupBy("u.email, u.mobile").
		OrderBy("u.id ASC").
		Limit(10).Offset(5).
		Build()

	fmt.Println(qry)
	fmt.Println(args)
}
