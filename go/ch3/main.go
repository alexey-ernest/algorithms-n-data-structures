package ch3

type SymbolTable interface {
	Put(key, value string)
	Get(key string) string
	Delete(key string)
	Size() int
	IsEmpty() bool
	Contains(key string) bool
}

type OrderedSymbolTalbe interface {
	SymbolTable
	Min() string
	Max() string
	Floor(key string) string
	Ceiling(key string) string
	Select(rank int) string
	Keys(lo, hi string) []string
}