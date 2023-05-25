package store

import (
	"fmt"
	"sync"

	"github.com/khatibomar/paytabs-challenge/internal/datastructure"
	"github.com/khatibomar/paytabs-challenge/internal/validator"
)

var (
	ErrAccountDoesNotExist = fmt.Errorf("account does not exist")
	ErrAccountAlreadyExist = fmt.Errorf("account already exist")
)

type InMemoryStore struct {
	mu       sync.Mutex
	Accounts map[string]*datastructure.Account
}

func New() *InMemoryStore {
	return &InMemoryStore{}
}

func (s *InMemoryStore) Add(guid, name string, balance float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	validator := validator.New()

	err := validator.ValidateGuid(guid)
	if err != nil {
		return err
	}

	err = validator.ValidateBalance(balance)
	if err != nil {
		return err
	}

	_, err = s.Get(guid)
	if err == nil {
		return ErrAccountAlreadyExist
	}

	account := &datastructure.Account{
		Guid:    guid,
		Name:    name,
		Balance: balance,
	}
	s.Accounts[guid] = account
	return nil
}

func (s *InMemoryStore) Get(guid string) (*datastructure.Account, error) {
	account := s.Accounts[guid]
	if account == nil {
		return nil, ErrAccountDoesNotExist
	}
	return account, nil
}

func (s *InMemoryStore) Count() int {
	return len(s.Accounts)
}