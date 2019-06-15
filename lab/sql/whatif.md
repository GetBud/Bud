# What if..?

This SQL query building malarkey is pretty complex. It's very easy to mess up, as in, not include
the means to accomplish some kind of query building goal. So what if I just ignored all of that and
turned everything into strings. You could pass in args as variadiac args at the end of strings to
fill in placeholders even to still collect them all up for easy quering. Basically, this would just
be something that let you build the parts of a query in any order you want... and that's it. With
this approach, we'd have main builder methods for SQL keywords like `FROM` and `INNER JOIN`, etc.
but then let the user just fill in the rest as they see fit.

For example:

```go
// in app:

func buildQuery() {
    qry, args := sql.NewSelect().
        // Functions would just append to slices.
        Select("u.id").
        Select("COUNT(1) > ?", sql.Int(53)).
        From("users AS u").
        InnerJoin("accounts AS a ON ")
}

// in stmt:

func (s *Select) Select(expr string, args ...interface{}) *Select {
    s.selections = append(s.selections, stmt.NewExpression(expr, args...))
    return s
}
```
