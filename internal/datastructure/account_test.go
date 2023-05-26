package datastructure_test

import (
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/khatibomar/paytabs-challenge/internal/customerrors"
	"github.com/khatibomar/paytabs-challenge/internal/datastructure"
	"github.com/stretchr/testify/require"
)

func TestDepositAccount(t *testing.T) {
	a := &datastructure.Account{
		Guid:    uuid.New().String(),
		Balance: 0,
		Name:    faker.Name(),
	}
	err := a.Deposit(-1000)
	require.ErrorIs(t, err, customerrors.ErrNegativeTransactionAmountNotAllowed)
	err = a.Withdraw(-100)
	require.ErrorIs(t, err, customerrors.ErrNegativeTransactionAmountNotAllowed)
	err = a.Withdraw(1)
	require.ErrorIs(t, err, customerrors.ErrNegativeBalanceNotAllowed)
	err = a.Deposit(100)
	require.NoError(t, err)
	require.Equal(t, a.Balance, float64(100))
	err = a.Withdraw(50)
	require.NoError(t, err)
	require.Equal(t, a.Balance, float64(50))
	err = a.Withdraw(50)
	require.NoError(t, err)
	require.Equal(t, a.Balance, float64(0))
	err = a.Withdraw(1)
	require.ErrorIs(t, err, customerrors.ErrNegativeBalanceNotAllowed)
}

func TestTransferBetweenTwoAccounts(t *testing.T) {
	fromAccount := &datastructure.Account{
		Guid:    uuid.New().String(),
		Balance: 0,
		Name:    faker.Name(),
	}

	toAccount := &datastructure.Account{
		Guid:    uuid.New().String(),
		Balance: 0,
		Name:    faker.Name(),
	}

	err := fromAccount.Transfer(toAccount, 100)
	require.ErrorIs(t, err, customerrors.ErrNegativeBalanceNotAllowed)
	err = fromAccount.Deposit(100)
	require.NoError(t, err)
	err = fromAccount.Transfer(toAccount, 50)
	require.NoError(t, err)
	require.Equal(t, fromAccount.Balance, float64(50))
	require.Equal(t, toAccount.Balance, float64(50))
}
