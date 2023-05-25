package store

import "github.com/khatibomar/paytabs-challenge/internal/datastructure"

type Store interface {
	Add(*datastructure.Account) error
	Get(string) (*datastructure.Account, error)
	All() []*datastructure.Account
}
