package bud

import "github.com/getbud/bud/recurrence"

// PlannedTransaction represents an upcoming Transaction that has not yet occurred. It may be a
// one-off payment, or something that is recurring.
type PlannedTransaction struct {
	// TODO: Account to apply transaction to.
	Amount     int              `json:"amount"`
	Recurrence *recurrence.Rule `json:"recurrence,omitempty"`
}

// Transaction represents a transaction that has happened, i.e. it has been reflected in a
// statement that has been imported.
type Transaction struct {
	// TODO: ...
	Amount int `json:"amount"`
}
