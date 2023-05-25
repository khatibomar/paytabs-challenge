package validator

import (
	"github.com/google/uuid"
	"github.com/khatibomar/paytabs-challenge/internal/customerrors"
)

var ()

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
		return customerrors.ErrNegativeBalanceNotAllowed
	}
	return nil
}

func (v *Validator) ValidateTransactionAmount(amount float64) error {
	if amount < 0 {
		return customerrors.ErrNegativeTransactionAmountNotAllowed
	}
	return nil
}

func (v *Validator) ValidateDeposit(balance, amount float64) error {
	err := v.ValidateTransactionAmount(amount)
	if err != nil {
		return err
	}
	err = v.ValidateBalance(balance + amount)
	if err != nil {
		return err
	}
	return nil
}

func (v *Validator) ValidateWithdrawal(balance, amount float64) error {
	err := v.ValidateTransactionAmount(amount)
	if err != nil {
		return err
	}
	err = v.ValidateBalance(balance - amount)
	if err != nil {
		return err
	}
	return nil
}
