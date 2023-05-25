package datastructure

import (
	"fmt"
	"sync"

	"github.com/khatibomar/paytabs-challenge/internal/validator"
)

var (
	ErrAccountDoesNotExist = fmt.Errorf("account does not exist")
	ErrAccountAlreadyExist = fmt.Errorf("account already exist")
)

type Store struct {
	mu       sync.Mutex
	Accounts map[string]*Account
}

func New() *Store {
	return &Store{}
}

func (s *Store) Add(guid, name string, balance float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

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

	account := &Account{
		Guid:    guid,
		Name:    name,
		Balance: balance,
	}
	s.Accounts[guid] = account
	return nil
}

func (s *Store) Get(guid string) (*Account, error) {
	account := s.Accounts[guid]
	if account == nil {
		return nil, ErrAccountDoesNotExist
	}
	return account, nil
}

func (s *Store) Count() int {
	return len(s.Accounts)
}
