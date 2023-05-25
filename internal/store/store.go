package store

type Store interface {
	Add(string, string, float64) error
	Get(string) error
}
