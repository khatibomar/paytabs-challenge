package store_test

import (
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/khatibomar/paytabs-challenge/internal/datastructure"
	"github.com/khatibomar/paytabs-challenge/internal/store"
	"github.com/stretchr/testify/require"
)

func TestInMemoryStoreValid(t *testing.T) {
	s := store.New()

	var validTests = []*datastructure.Account{
		{Guid: uuid.New().String(), Name: faker.Name(), Balance: faker.Float64Range(0, 1000)},
		{Guid: uuid.New().String(), Name: faker.Name(), Balance: faker.Float64Range(0, 1000)},
		{Guid: uuid.New().String(), Name: faker.Name(), Balance: faker.Float64Range(0, 1000)},
	}

	for _, test := range validTests {
		err := s.Add(test.Guid, test.Name, test.Balance)
		require.NoError(t, err)
	}
}
