package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/getbud/bud/recurrence"
)

func main() {
	now := time.Now()

	payRule, err := recurrence.NewRule(recurrence.Monthly,
		recurrence.WithOccurrences(-1),
		recurrence.WithStartDate(now.Year(), now.Month(), now.Day()),
		recurrence.WithWeekdays(
			time.Monday,
			time.Tuesday,
			time.Wednesday,
			time.Thursday,
			time.Friday,
		),
	)

	if err != nil {
		panic(err)
	}

	occurrences := payRule.Between(time.Now(), time.Now().AddDate(0, 1, 0))

	spew.Dump(occurrences)

	fmt.Println("payRule", payRule)

	rentRule, err := recurrence.NewRule(recurrence.Monthly,
		recurrence.WithMonthdays(10),
		recurrence.WithStartDate(2019, time.June, 10),
		recurrence.WithEndDate(2019, time.October, 10),
	)

	if err != nil {
		panic(err)
	}

	occurrences = rentRule.Between(time.Now(), time.Now().AddDate(1, 0, 0))

	spew.Dump(occurrences)

	fmt.Println("rentRule", rentRule)
}
