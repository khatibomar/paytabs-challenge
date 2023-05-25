package model

import "sync"

type Account struct {
	mu      sync.Mutex
	Guid    string
	name    string
	balance float64
}
