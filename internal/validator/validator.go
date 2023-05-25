package validator

import (
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrNegativeBalanceNotAllowed = fmt.Errorf("you can't have negative balance")
)

func ValidateGuid(guid string) error {
	_, err := uuid.Parse(guid)
	return err
}

func ValidateBalance(balance float64) error {
	if balance < 0 {
		return ErrNegativeBalanceNotAllowed
	}
	return nil
}
