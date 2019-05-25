package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/getbud/bud/bud"
	"github.com/getbud/bud/recurrence"
)

func main() {
	r := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	matches := r.FindStringSubmatch(`2015-05-27`)

	args := make([]interface{}, 0, len(matches))
	for _, match := range matches {
		args = append(args, match)
	}

	fmt.Printf("%s %s %s %s\n", args...)

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
		Recurrence:  payRule,
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
