package bud

import (
	"regexp"
	"time"

	"github.com/getbud/bud/recurrence"
)

// Account represents a bank account.
type Account struct {
	// UUID is a unique identifier for this Account.
	UUID string `json:"uuid"`
	// Name is the name of this Account.
	Name string `json:"name"`
	// Balance is the current balance of this Account.
	Balance int `json:"balance"`
}

// Category represents a grouping of Transactions of a certain kind. A Transaction may belong to
// only one Category, otherwise it would make visualisations far less useful.
type Category struct {
	// UUID is a unique identifier for this Category.
	UUID string `json:"uuid"`
	// Name is the name of this Category.
	Name string `json:"name"`
}

// PlannedTransaction represents an upcoming Transaction that has not yet occurred. It may be a
// one-off payment, or something that is recurring.
type PlannedTransaction struct {
	// UUID is a unique identifier for this PlannedTransaction.
	UUID string `json:"uuid"`
	// AccountUUID denotes the Account that this PlannedTransaction belongs to.
	AccountUUID string `json:"account_uuid"`
	// CategoryUUID denotes the Category that this PlannedTransaction belongs to.
	CategoryUUID string `json:"category_uuid"`
	// Description is a description of this PlannedTransaction. The format should be similar to the
	// Description found on a regular Transaction.
	Description string `json:"description"`
	// Amount is the expected amount that will be applied to the Balance of the linked Account for
	// this PlannedTransaction. It may be either positive or negative.
	Amount int `json:"amount"`
	// Recurrence defines when this PlannedTransaction is expected to be applied. Even if the
	// Transaction is only planned to happen once, this will still be set to define the expected
	// date of the Transaction.
	Recurrence recurrence.Rule `json:"recurrence"`
}

// Transaction represents a Transaction that has happened, i.e. it has been reflected in a
// statement that has been imported.
type Transaction struct {
	// UUID is a unique identifier for this Transaction.
	UUID string `json:"uuid"`
	// AccountUUID denotes the Account that this Transaction belongs to.
	AccountUUID string `json:"account_uuid"`
	// CategoryUUID denotes the Category that this Transaction belongs to.
	CategoryUUID string `json:"category_uuid"`
	// PlannedTransactionUUID is optionally set if this Transaction matched a Transformation that is
	// linked to a PlannedTransaction. This allows Bud to identify transactions that have or have
	// not happened.
	PlannedTransactionUUID string `json:"planned_transaction_uuid,omitempty"`
	// TransformationUUID is optionally set if this Transaction matched a Transformation.
	TransformationUUID string `json:"transformation_uuid,omitempty"`
	// Description is a description of this Transaction. It may have been changed since it was
	// originally imported.
	Description string `json:"description"`
	// Amount specifies the amount to apply to the balance of the Account linked to this
	// Transaction. It may be either positive or negative.
	Amount int `json:"amount"`
	// TransactedAt specifies the date that this Transaction happened on.
	TransactedAt time.Time `json:"transacted_at"`
}

// Transformation represents a description of how a Transaction may be transformed, given a certain
// input, e.g. a Transaction with a matching description may be changed to have a new description.
type Transformation struct {
	// UUID is a unique identifier for this Transformation.
	UUID string `json:"uuid"`
	// Description provides a description for what Transactions this Transformation applies to.
	Description string `json:"description"`
	// Pattern is a regular expression pattern that is applied to the description of a Transaction.
	// If the regular expression matches on a Transaction's description then the Transformation will
	// be applied to it.
	Pattern *regexp.Regexp `json:"pattern"`

	// CategoryUUID denotes which Category will be applied to a matching Transaction.
	CategoryUUID string `json:"category_uuid"`
	// PlannedTransactionUUID is an optional PlannedTransaction to link this Transformation to. If
	// a Transaction matches this Transformation then the PlannedTransactionUUID will be given to
	// that matching Transaction, linking the two.
	PlannedTransactionUUID string `json:"planned_transaction_uuid,omitempty"`
	// DescriptionFormat presents a format that may be used for creating a new name for a matching
	// Transaction. The format may include numbered 'variables' captured by this Transformation's
	// pattern via Regular Expression capture groups. For example; given the following regular
	// expression: '^SUZUKI FINANCE (\d+)$', the description format could be 'Suzuki Finance: $1'.
	DescriptionFormat string
}
