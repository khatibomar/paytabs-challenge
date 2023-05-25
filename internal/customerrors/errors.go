package customerrors

import "fmt"

var (
	ErrNegativeBalanceNotAllowed           = fmt.Errorf("you can't have negative balance")
	ErrNegativeTransactionAmountNotAllowed = fmt.Errorf("transactions amount should always be positive")
	ErrAccountDoesNotExist                 = fmt.Errorf("account does not exist")
	ErrAccountAlreadyExist                 = fmt.Errorf("account already exist")
)
