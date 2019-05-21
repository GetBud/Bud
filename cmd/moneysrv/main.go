package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/getbud/bud/bud"
	"github.com/getbud/bud/recurrence"
)

func main() {
	now := time.Now()

	payRule, err := recurrence.NewRule(recurrence.Monthly,
		recurrence.WithSetPositions(-1),
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

	payPT := bud.PlannedTransaction{
		Description: "Monthly pay",
		Amount:      1000000,
		Recurrence:  &payRule,
	}

	buf := &bytes.Buffer{}

	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(true)

	err = enc.Encode(payPT)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
