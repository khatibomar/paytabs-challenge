package datastructure

import (
	"sync"
)

type Account struct {
	mu      sync.Mutex
	Guid    string
	Name    string
	Balance float64
}
