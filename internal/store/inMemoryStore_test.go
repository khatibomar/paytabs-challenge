package store_test

import (
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/khatibomar/paytabs-challenge/internal/datastructure"
	"github.com/khatibomar/paytabs-challenge/internal/store"
	"github.com/khatibomar/paytabs-challenge/internal/validator"
	"github.com/stretchr/testify/require"
)

func TestInMemoryStoreValid(t *testing.T) {
	s := store.New()

	var validAccounts = []*datastructure.Account{
		{Guid: uuid.New().String(), Name: faker.Name(), Balance: faker.Float64Range(0, 1000)},
		{Guid: uuid.New().String(), Name: faker.Name(), Balance: faker.Float64Range(0, 1000)},
		{Guid: uuid.New().String(), Name: faker.Name(), Balance: faker.Float64Range(0, 1000)},
	}

	for _, account := range validAccounts {
		err := s.Add(account)
		require.NoError(t, err)
	}
}

func TestInMemoryStoreWithDuplicateGuid(t *testing.T) {
	s := store.New()
	guid := uuid.New().String()
	account := &datastructure.Account{
		Guid:    guid,
		Name:    faker.Name(),
		Balance: faker.Float64Range(0, 1000),
	}
	err := s.Add(account)
	require.NoError(t, err)
	err = s.Add(account)
	require.ErrorIs(t, err, store.ErrAccountAlreadyExist)
}

func TestInMemoryStoreWithNegativeBalance(t *testing.T) {
	s := store.New()

	var validTests = []*datastructure.Account{
		{Guid: uuid.New().String(), Name: faker.Name(), Balance: faker.Float64Range(-2000, -1000)},
		{Guid: uuid.New().String(), Name: faker.Name(), Balance: faker.Float64Range(-2000, -1000)},
		{Guid: uuid.New().String(), Name: faker.Name(), Balance: faker.Float64Range(-2000, -1000)},
	}

	for _, account := range validTests {
		err := s.Add(account)
		require.ErrorIs(t, err, validator.ErrNegativeBalanceNotAllowed)
	}
}

func TestInMemoryStoreGetExistingAccount(t *testing.T) {
	s := store.New()

	account := &datastructure.Account{
		Guid:    uuid.New().String(),
		Name:    faker.Name(),
		Balance: faker.Float64Range(1000, 2000),
	}

	err := s.Add(account)
	require.NoError(t, err)

	acc, err := s.Get(account.Guid)
	require.NoError(t, err)
	if acc != account {
		t.Fatal("Should return same account, this should not happen")
	}
}

func TestInMemoryStoreGetNotExistingAccount(t *testing.T) {
	s := store.New()

	account := &datastructure.Account{
		Guid:    uuid.New().String(),
		Name:    faker.Name(),
		Balance: faker.Float64Range(1000, 2000),
	}

	err := s.Add(account)
	require.NoError(t, err)

	_, err = s.Get(uuid.New().String())
	require.ErrorIs(t, err, store.ErrAccountDoesNotExist)
}

func TestStoreSeed(t *testing.T) {
	s := store.New()
	err := s.Seed()
	require.NoError(t, err)
	require.Equal(t, s.Count(), 500)
}
