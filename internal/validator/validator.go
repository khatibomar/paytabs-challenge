package validator

import (
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrNegativeBalanceNotAllowed           = fmt.Errorf("you can't have negative balance")
	ErrNegativeTransactionAmountNotAllowed = fmt.Errorf("transactions amount should always be positive")
)

type Validator struct {
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateGuid(guid string) error {
	_, err := uuid.Parse(guid)
	return err
}

func (v *Validator) ValidateBalance(balance float64) error {
	if balance < 0 {
		return ErrNegativeBalanceNotAllowed
	}
	return nil
}

func (v *Validator) ValidateTransactionAmount(amount float64) error {
	if amount < 0 {
		return ErrNegativeTransactionAmountNotAllowed
	}
	return nil
}
