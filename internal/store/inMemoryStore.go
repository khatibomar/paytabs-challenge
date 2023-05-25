package store

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/khatibomar/paytabs-challenge/internal/datastructure"
	parser "github.com/khatibomar/paytabs-challenge/internal/parser"
	"github.com/khatibomar/paytabs-challenge/internal/validator"
)

var (
	ErrAccountDoesNotExist = fmt.Errorf("account does not exist")
	ErrAccountAlreadyExist = fmt.Errorf("account already exist")
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
		return ErrAccountAlreadyExist
	}

	s.accounts[account.Guid] = account
	return nil
}

func (s *InMemoryStore) Get(guid string) (*datastructure.Account, error) {
	account := s.accounts[guid]
	if account == nil {
		return nil, ErrAccountDoesNotExist
	}
	return account, nil
}

func (s *InMemoryStore) Count() int {
	return len(s.accounts)
}

func (s *InMemoryStore) All() []*datastructure.Account {
	var accounts []*datastructure.Account
	for _, account := range s.accounts {
		accounts = append(accounts, account)
	}
	return accounts
}

func (s *InMemoryStore) Seed() error {
	p := parser.New()
	path := "../../data/accounts-mock.json"
	rawAccounts, err := p.ParseJson(path)
	if err != nil {
		return err
	}

	account := datastructure.Account{}

	for _, rawAccount := range rawAccounts {
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
