package datastructure_test

import (
	"testing"

	"github.com/khatibomar/paytabs-challenge/internal/customerrors"
	"github.com/khatibomar/paytabs-challenge/internal/datastructure"
	"github.com/stretchr/testify/require"
)

func TestDepositAccount(t *testing.T) {
	a := &datastructure.Account{}
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
