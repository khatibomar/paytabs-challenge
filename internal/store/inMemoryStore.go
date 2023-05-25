package store

import (
	"strconv"
	"sync"

	"github.com/khatibomar/paytabs-challenge/internal/customerrors"
	"github.com/khatibomar/paytabs-challenge/internal/datastructure"
	parser "github.com/khatibomar/paytabs-challenge/internal/parser"
	"github.com/khatibomar/paytabs-challenge/internal/validator"
)

type InMemoryStore struct {
	mu       sync.Mutex
	accounts map[string]*datastructure.Account
}

func New() *InMemoryStore {
	return &InMemoryStore{
		accounts: make(map[string]*datastructure.Account),
	}
}

func (s *InMemoryStore) Add(account *datastructure.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	validator := validator.New()

	err := validator.ValidateGuid(account.Guid)
	if err != nil {
		return err
	}

	err = validator.ValidateBalance(account.Balance)
	if err != nil {
		return err
	}

	_, err = s.Get(account.Guid)
	if err == nil {
		return customerrors.ErrAccountAlreadyExist
	}

	s.accounts[account.Guid] = account
	return nil
}

func (s *InMemoryStore) Get(guid string) (*datastructure.Account, error) {
	account := s.accounts[guid]
	if account == nil {
		return nil, customerrors.ErrAccountDoesNotExist
	}
	return account, nil
}

func (s *InMemoryStore) Count() int {
	return len(s.accounts)
}

func (s *InMemoryStore) All() []*datastructure.Account {
	var accounts []*datastructure.Account
	for key := range s.accounts {
		accounts = append(accounts, s.accounts[key])
	}
	return accounts
}

func (s *InMemoryStore) Seed(path string) error {
	p := parser.New()
	rawAccounts, err := p.ParseJson(path)
	if err != nil {
		return err
	}

	for _, rawAccount := range rawAccounts {
		account := datastructure.Account{}
		account.Guid = rawAccount.ID
		account.Name = rawAccount.Name
		account.Balance, err = strconv.ParseFloat(rawAccount.Balance, 64)
		if err != nil {
			return err
		}
		err = s.Add(&account)
		if err != nil {
			return err
		}
	}

	return nil
}
