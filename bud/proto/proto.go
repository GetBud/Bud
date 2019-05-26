package proto

import (
	"regexp"

	"github.com/getbud/bud/bud"
	"github.com/getbud/bud/recurrence"
	"github.com/golang/protobuf/ptypes"
	"github.com/icelolly/go-errors"
)

//go:generate protoc bud.proto -I $GOPATH/src -I . --gofast_out=.

// BudAccount converts a proto Account to bud.Account.
func BudAccount(account *Account) bud.Account {
	return bud.Account{
		UUID:    account.UUID,
		Name:    account.Name,
		Balance: int(account.Balance),
	}
}

// ProtoAccount converts a bud.Account to User.
func ProtoAccount(account bud.Account) *Account {
	return &Account{
		UUID:    account.UUID,
		Name:    account.Name,
		Balance: int64(account.Balance),
	}
}

// BudCategory converts a proto Category to bud.Category.
func BudCategory(category *Category) bud.Category {
	return bud.Category{
		UUID: category.UUID,
		Name: category.Name,
	}
}

// ProtoCategory converts a bud.Category to Category.
func ProtoCategory(category bud.Category) *Category {
	return &Category{
		UUID: category.UUID,
		Name: category.Name,
	}
}

// BudPlannedTransaction converts a proto PlannedTransaction to bud.PlannedTransaction.
func BudPlannedTransaction(plannedTransaction *PlannedTransaction) bud.PlannedTransaction {
	recurrenceRule, err := recurrence.NewRuleFromString(plannedTransaction.Recurrence)
	if err != nil {
		// TODO: Perhaps we can do something more appropriate here. Though if things are working
		// properly we should never see this happen...
		errors.Fatal(err)
	}

	return bud.PlannedTransaction{
		UUID:         plannedTransaction.UUID,
		AccountUUID:  plannedTransaction.AccountUUID,
		CategoryUUID: plannedTransaction.CategoryUUID,
		Description:  plannedTransaction.Description,
		Amount:       int(plannedTransaction.Amount),
		Recurrence:   recurrenceRule,
	}
}

// ProtoPlannedTransaction converts a bud.PlannedTransaction to PlannedTransaction.
func ProtoPlannedTransaction(plannedTransaction bud.PlannedTransaction) *PlannedTransaction {
	return &PlannedTransaction{
		UUID:         plannedTransaction.UUID,
		AccountUUID:  plannedTransaction.AccountUUID,
		CategoryUUID: plannedTransaction.CategoryUUID,
		Description:  plannedTransaction.Description,
		Amount:       int64(plannedTransaction.Amount),
		Recurrence:   plannedTransaction.Recurrence.String(),
	}
}

// BudTransaction converts a proto Transaction to bud.Transaction.
func BudTransaction(transaction *Transaction) bud.Transaction {
	transactedAt, err := ptypes.Timestamp(transaction.TransactedAt)
	if err != nil {
		// TODO: Perhaps we can do something more appropriate here. Though if things are working
		// properly we should never see this happen...
		errors.Fatal(err)
	}

	return bud.Transaction{
		UUID:                   transaction.UUID,
		AccountUUID:            transaction.AccountUUID,
		CategoryUUID:           transaction.CategoryUUID,
		PlannedTransactionUUID: transaction.PlannedTransactionUUID,
		TransformationUUID:     transaction.TransformationUUID,
		Description:            transaction.Description,
		Amount:                 int(transaction.Amount),
		TransactedAt:           transactedAt,
	}
}

// ProtoTransaction converts a bud.Transaction to Transaction.
func ProtoTransaction(transaction bud.Transaction) *Transaction {
	transactedAt, err := ptypes.TimestampProto(transaction.TransactedAt)
	if err != nil {
		// TODO: Perhaps we can do something more appropriate here. Though if things are working
		// properly we should never see this happen...
		errors.Fatal(err)
	}

	return &Transaction{
		UUID:                   transaction.UUID,
		AccountUUID:            transaction.AccountUUID,
		CategoryUUID:           transaction.CategoryUUID,
		PlannedTransactionUUID: transaction.PlannedTransactionUUID,
		TransformationUUID:     transaction.TransformationUUID,
		Description:            transaction.Description,
		Amount:                 int64(transaction.Amount),
		TransactedAt:           transactedAt,
	}

}

// BudTransformation converts a proto Transformation to bud.Transformation.
func BudTransformation(transformation *Transformation) bud.Transformation {
	pattern, err := regexp.Compile(transformation.Pattern)
	if err != nil {
		// TODO: Perhaps we can do something more appropriate here. Though if things are working
		// properly we should never see this happen...
		errors.Fatal(err)
	}

	return bud.Transformation{
		UUID:                   transformation.UUID,
		Description:            transformation.Description,
		Pattern:                pattern,
		CategoryUUID:           transformation.CategoryUUID,
		PlannedTransactionUUID: transformation.PlannedTransactionUUID,
		DescriptionFormat:      transformation.DescriptionFormat,
	}
}

// ProtoTransformation converts a bud.Transformation to Transformation.
func ProtoTransformation(transformation bud.Transformation) *Transformation {
	return &Transformation{
		UUID:                   transformation.UUID,
		Description:            transformation.Description,
		Pattern:                transformation.Pattern.String(),
		CategoryUUID:           transformation.CategoryUUID,
		PlannedTransactionUUID: transformation.PlannedTransactionUUID,
		DescriptionFormat:      transformation.DescriptionFormat,
	}
}
