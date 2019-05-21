package recurrence

import (
	"time"

	"github.com/icelolly/go-errors"
	"github.com/teambition/rrule-go"
)

// The various Frequency values.
const (
	Yearly Frequency = iota
	Monthly
	Weekly
	Daily
)

// Frequency denotes the period upon which a recurrence rule is evaluated. It is a required
// parameter to a recurrence rule.
type Frequency int

// frequencyMap is a map used to return an rrule frequency from our custom frequency type.
var frequencyMap = map[Frequency]rrule.Frequency{
	Yearly:  rrule.YEARLY,
	Monthly: rrule.MONTHLY,
	Weekly:  rrule.WEEKLY,
	Daily:   rrule.DAILY,
}

// weekdayMap is a map used to return an rrule weekday from a stdlib time weekday.
var weekdayMap = map[time.Weekday]rrule.Weekday{
	time.Monday:    rrule.MO,
	time.Tuesday:   rrule.TU,
	time.Wednesday: rrule.WE,
	time.Thursday:  rrule.TH,
	time.Friday:    rrule.FR,
	time.Saturday:  rrule.SA,
	time.Sunday:    rrule.SU,
}

// Rule ...
type Rule struct {
	rrule *rrule.RRule
}

// NewRule returns a new recurrence rule using the given values to describe when a recurrence should
// be. This uses a subset of all possible recurrence rule values specified in RFC 2445.
func NewRule(frequency Frequency, options ...Option) (Rule, error) {
	rule := Rule{}

	freq, ok := frequencyMap[frequency]
	if !ok {
		return rule, errors.New("recurrence: invalid frequency value")
	}

	ropt, err := buildRRuleROption(freq, options...)
	if err != nil {
		return rule, errors.Wrap(err, "recurrence: failed to build options")
	}

	rr, err := rrule.NewRRule(ropt)
	if err != nil {
		return rule, errors.Wrap(err, "recurrence: failed to build rule")
	}

	rule.rrule = rr

	return rule, nil
}

// All returns all occurrences that apply to this recurrence rule.
func (r *Rule) All() []time.Time {
	return r.rrule.All()
}

// Between returns all occurrences that apply to this recurrence rule between the given dates.
func (r *Rule) Between(after, before time.Time) []time.Time {
	return r.rrule.Between(after, before, true)
}

// Next returns the next occurrence from this rule, if there is one. If there isn't one, the bool
// return value will be false.
func (r *Rule) Next() (time.Time, bool) {
	next := r.rrule.Iterator()
	return next()
}

// MarshalJSON allows a Rule to be marshaled as JSON.
func (r Rule) MarshalJSON() ([]byte, error) {
	return []byte(`"` + r.String() + `"`), nil
}

// String returns this recurrence rule as a string.
func (r Rule) String() string {
	return r.rrule.String()
}

// UnmarshalJSON allows a Rule to be un-marshaled from JSON.
func (r *Rule) UnmarshalJSON(bs []byte) error {
	if len(bs) <= 2 {
		return errors.New("recurrence: expected at least 2 bytes")
	}

	str := string(bs[1 : len(bs)-1])

	rr, err := rrule.StrToRRule(str)
	if err != nil {
		return errors.Wrap(err, "recurrence: failed to parse rule string")
	}

	*r = Rule{
		rrule: rr,
	}

	return nil
}

// buildRRuleROption ...
func buildRRuleROption(freq rrule.Frequency, options ...Option) (rrule.ROption, error) {
	ropt := rrule.ROption{
		Freq: freq,
	}

	for _, option := range options {
		switch option.kind {
		case optionKindStartTime:
			ropt.Dtstart = option.starting
		case optionKindEndTime:
			ropt.Until = option.ending
		case optionKindMonthdays:
			ropt.Bymonthday = option.monthdays
		case optionKindWeekdays:
			wds := make([]rrule.Weekday, 0, len(option.weekdays))

			for _, weekday := range option.weekdays {
				wd, ok := weekdayMap[weekday]
				if !ok {
					return ropt, errors.New("recurrence: invalid weekday value")
				}

				wds = append(wds, wd)
			}

			ropt.Byweekday = wds
		case optionKindSetPositions:
			ropt.Bysetpos = option.setPositions
		case optionKindInterval:
			ropt.Interval = option.interval
		}
	}

	return ropt, nil
}

// The various optionKind values.
const (
	optionKindStartTime optionKind = iota
	optionKindEndTime
	optionKindMonthdays
	optionKindWeekdays
	optionKindSetPositions
	optionKindInterval
)

// optionKind helps us branch off to apply the various options we allow.
type optionKind int

// Option ...
type Option struct {
	kind         optionKind
	starting     time.Time
	ending       time.Time
	monthdays    []int
	weekdays     []time.Weekday
	setPositions []int
	interval     int
}

// WithStartDate ...
func WithStartDate(year int, month time.Month, day int) Option {
	starting := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	return Option{
		kind:     optionKindStartTime,
		starting: starting,
	}
}

// WithEndDate ...
func WithEndDate(year int, month time.Month, day int) Option {
	ending := time.Date(year, month, day, 23, 59, 59, 999999999, time.UTC)

	return Option{
		kind:   optionKindEndTime,
		ending: ending,
	}
}

// WithMonthdays ...
func WithMonthdays(monthdays ...int) Option {
	return Option{
		kind:      optionKindMonthdays,
		monthdays: monthdays,
	}
}

// WithWeekdays ...
func WithWeekdays(weekdays ...time.Weekday) Option {
	return Option{
		kind:     optionKindWeekdays,
		weekdays: weekdays,
	}
}

// WithSetPositions ...
func WithSetPositions(setPositions ...int) Option {
	return Option{
		kind:         optionKindSetPositions,
		setPositions: setPositions,
	}
}

// WithInterval ...
func WithInterval(interval int) Option {
	return Option{
		kind:     optionKindInterval,
		interval: interval,
	}
}
